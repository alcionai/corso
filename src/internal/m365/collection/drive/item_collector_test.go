package drive

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/mock"
)

type ItemCollectorUnitSuite struct {
	tester.Suite
}

func TestOneDriveUnitSuite(t *testing.T) {
	suite.Run(t, &ItemCollectorUnitSuite{Suite: tester.NewUnitSuite(t)})
}

const (
	userMysiteURLNotFound = "BadRequest Unable to retrieve user's mysite URL"
	userMysiteNotFound    = "ResourceNotFound User's mysite not found"
)

func odErr(code string) *odataerrors.ODataError {
	odErr := odataerrors.NewODataError()
	merr := odataerrors.NewMainError()
	merr.SetCode(&code)
	odErr.SetErrorEscaped(merr)

	return odErr
}

func (suite *ItemCollectorUnitSuite) TestDrives() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	numDriveResults := 4
	emptyLink := ""
	link := "foo"

	// These errors won't be the "correct" format when compared to what graph
	// returns, but they're close enough to have the same info when the inner
	// details are extracted via support package.
	mySiteURLNotFound := odErr(userMysiteURLNotFound)
	mySiteNotFound := odErr(userMysiteNotFound)

	resultDrives := make([]models.Driveable, 0, numDriveResults)

	for i := 0; i < numDriveResults; i++ {
		d := models.NewDrive()
		id := uuid.NewString()
		d.SetId(&id)

		resultDrives = append(resultDrives, d)
	}

	tooManyRetries := make([]mock.PagerResult[models.Driveable], 0, maxDrivesRetries+1)

	for i := 0; i < maxDrivesRetries+1; i++ {
		tooManyRetries = append(tooManyRetries, mock.PagerResult[models.Driveable]{
			Err: context.DeadlineExceeded,
		})
	}

	table := []struct {
		name            string
		pagerResults    []mock.PagerResult[models.Driveable]
		retry           bool
		expectedErr     assert.ErrorAssertionFunc
		expectedResults []models.Driveable
	}{
		{
			name: "AllOneResultNilNextLink",
			pagerResults: []mock.PagerResult[models.Driveable]{
				{
					Values:   resultDrives,
					NextLink: nil,
					Err:      nil,
				},
			},
			retry:           false,
			expectedErr:     assert.NoError,
			expectedResults: resultDrives,
		},
		{
			name: "AllOneResultEmptyNextLink",
			pagerResults: []mock.PagerResult[models.Driveable]{
				{
					Values:   resultDrives,
					NextLink: &emptyLink,
					Err:      nil,
				},
			},
			retry:           false,
			expectedErr:     assert.NoError,
			expectedResults: resultDrives,
		},
		{
			name: "SplitResultsNilNextLink",
			pagerResults: []mock.PagerResult[models.Driveable]{
				{
					Values:   resultDrives[:numDriveResults/2],
					NextLink: &link,
					Err:      nil,
				},
				{
					Values:   resultDrives[numDriveResults/2:],
					NextLink: nil,
					Err:      nil,
				},
			},
			retry:           false,
			expectedErr:     assert.NoError,
			expectedResults: resultDrives,
		},
		{
			name: "SplitResultsEmptyNextLink",
			pagerResults: []mock.PagerResult[models.Driveable]{
				{
					Values:   resultDrives[:numDriveResults/2],
					NextLink: &link,
					Err:      nil,
				},
				{
					Values:   resultDrives[numDriveResults/2:],
					NextLink: &emptyLink,
					Err:      nil,
				},
			},
			retry:           false,
			expectedErr:     assert.NoError,
			expectedResults: resultDrives,
		},
		{
			name: "NonRetryableError",
			pagerResults: []mock.PagerResult[models.Driveable]{
				{
					Values:   resultDrives,
					NextLink: &link,
					Err:      nil,
				},
				{
					Values:   nil,
					NextLink: nil,
					Err:      assert.AnError,
				},
			},
			retry:           true,
			expectedErr:     assert.Error,
			expectedResults: nil,
		},
		{
			name: "MySiteURLNotFound",
			pagerResults: []mock.PagerResult[models.Driveable]{
				{
					Values:   nil,
					NextLink: nil,
					Err:      graph.Stack(ctx, mySiteURLNotFound),
				},
			},
			retry:           true,
			expectedErr:     assert.NoError,
			expectedResults: nil,
		},
		{
			name: "MySiteNotFound",
			pagerResults: []mock.PagerResult[models.Driveable]{
				{
					Values:   nil,
					NextLink: nil,
					Err:      graph.Stack(ctx, mySiteNotFound),
				},
			},
			retry:           true,
			expectedErr:     assert.NoError,
			expectedResults: nil,
		},
		{
			name: "SplitResultsContextTimeoutWithRetries",
			pagerResults: []mock.PagerResult[models.Driveable]{
				{
					Values:   resultDrives[:numDriveResults/2],
					NextLink: &link,
					Err:      nil,
				},
				{
					Values:   nil,
					NextLink: nil,
					Err:      context.DeadlineExceeded,
				},
				{
					Values:   resultDrives[numDriveResults/2:],
					NextLink: &emptyLink,
					Err:      nil,
				},
			},
			retry:           true,
			expectedErr:     assert.NoError,
			expectedResults: resultDrives,
		},
		{
			name: "SplitResultsContextTimeoutNoRetries",
			pagerResults: []mock.PagerResult[models.Driveable]{
				{
					Values:   resultDrives[:numDriveResults/2],
					NextLink: &link,
					Err:      nil,
				},
				{
					Values:   nil,
					NextLink: nil,
					Err:      context.DeadlineExceeded,
				},
				{
					Values:   resultDrives[numDriveResults/2:],
					NextLink: &emptyLink,
					Err:      nil,
				},
			},
			retry:           false,
			expectedErr:     assert.Error,
			expectedResults: nil,
		},
		{
			name: "TooManyRetries",
			pagerResults: append(
				[]mock.PagerResult[models.Driveable]{
					{
						Values:   resultDrives[:numDriveResults/2],
						NextLink: &link,
						Err:      nil,
					},
				},
				tooManyRetries...,
			),
			retry:           true,
			expectedErr:     assert.Error,
			expectedResults: nil,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			pager := &mock.Pager[models.Driveable]{
				ToReturn: test.pagerResults,
			}

			drives, err := api.GetAllDrives(ctx, pager, test.retry, maxDrivesRetries)
			test.expectedErr(t, err, clues.ToCore(err))

			assert.ElementsMatch(t, test.expectedResults, drives)
		})
	}
}

// Integration tests

type OneDriveIntgSuite struct {
	tester.Suite
	userID string
	creds  account.M365Config
	ac     api.Client
}

func TestOneDriveSuite(t *testing.T) {
	suite.Run(t, &OneDriveIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *OneDriveIntgSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	suite.userID = tconfig.SecondaryM365UserID(t)

	acct := tconfig.NewM365Account(t)
	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.creds = creds

	suite.ac, err = api.NewClient(creds, control.DefaultOptions())
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *OneDriveIntgSuite) TestOneDriveNewCollections() {
	creds, err := tconfig.NewM365Account(suite.T()).M365Config()
	require.NoError(suite.T(), err, clues.ToCore(err))

	tests := []struct {
		name, user string
	}{
		{
			name: "Test User w/ Drive",
			user: suite.userID,
		},
		{
			name: "Test User w/out Drive",
			user: "testevents@10rqc2.onmicrosoft.com",
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				service = loadTestService(t)
				scope   = selectors.
					NewOneDriveBackup([]string{test.user}).
					AllData()[0]
			)

			colls := NewCollections(
				&itemBackupHandler{suite.ac.Drives(), test.user, scope},
				creds.AzureTenantID,
				test.user,
				service.updateStatus,
				control.Options{
					ToggleFeatures: control.Toggles{},
				})

			ssmb := prefixmatcher.NewStringSetBuilder()

			odcs, _, err := colls.Get(ctx, nil, ssmb, fault.New(true))
			assert.NoError(t, err, clues.ToCore(err))
			// Don't expect excludes as this isn't an incremental backup.
			assert.True(t, ssmb.Empty())

			for _, entry := range odcs {
				assert.NotEmpty(t, entry.FullPath())
			}
		})
	}
}
