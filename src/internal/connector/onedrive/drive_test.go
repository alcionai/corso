package onedrive

import (
	"context"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type mockPageLinker struct {
	link *string
}

func (pl *mockPageLinker) GetOdataNextLink() *string {
	return pl.link
}

type pagerResult struct {
	drives   []models.Driveable
	nextLink *string
	err      error
}

type mockDrivePager struct {
	toReturn []pagerResult
	getIdx   int
}

func (p *mockDrivePager) GetPage(context.Context) (api.PageLinker, error) {
	if len(p.toReturn) <= p.getIdx {
		return nil, assert.AnError
	}

	idx := p.getIdx
	p.getIdx++

	return &mockPageLinker{p.toReturn[idx].nextLink}, p.toReturn[idx].err
}

func (p *mockDrivePager) SetNext(string) {}

func (p *mockDrivePager) ValuesIn(api.PageLinker) ([]models.Driveable, error) {
	idx := p.getIdx
	if idx > 0 {
		// Return values lag by one since we increment in GetPage().
		idx--
	}

	if len(p.toReturn) <= idx {
		return nil, assert.AnError
	}

	return p.toReturn[idx].drives, nil
}

// Unit tests
type OneDriveUnitSuite struct {
	tester.Suite
}

func TestOneDriveUnitSuite(t *testing.T) {
	suite.Run(t, &OneDriveUnitSuite{Suite: tester.NewUnitSuite(t)})
}

const (
	userMysiteURLNotFound = "BadRequest Unable to retrieve user's mysite URL"
	userMysiteNotFound    = "ResourceNotFound User's mysite not found"
)

func odErr(code string) *odataerrors.ODataError {
	odErr := &odataerrors.ODataError{}
	merr := odataerrors.MainError{}
	merr.SetCode(&code)
	odErr.SetError(&merr)

	return odErr
}

func (suite *OneDriveUnitSuite) TestDrives() {
	// ctx, flush := tester.NewContext()
	// defer flush()

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

	tooManyRetries := make([]pagerResult, 0, getDrivesRetries+1)

	for i := 0; i < getDrivesRetries+1; i++ {
		tooManyRetries = append(tooManyRetries, pagerResult{
			err: context.DeadlineExceeded,
		})
	}

	table := []struct {
		name            string
		pagerResults    []pagerResult
		retry           bool
		expectedErr     assert.ErrorAssertionFunc
		expectedResults []models.Driveable
	}{
		{
			name: "AllOneResultNilNextLink",
			pagerResults: []pagerResult{
				{
					drives:   resultDrives,
					nextLink: nil,
					err:      nil,
				},
			},
			retry:           false,
			expectedErr:     assert.NoError,
			expectedResults: resultDrives,
		},
		{
			name: "AllOneResultEmptyNextLink",
			pagerResults: []pagerResult{
				{
					drives:   resultDrives,
					nextLink: &emptyLink,
					err:      nil,
				},
			},
			retry:           false,
			expectedErr:     assert.NoError,
			expectedResults: resultDrives,
		},
		{
			name: "SplitResultsNilNextLink",
			pagerResults: []pagerResult{
				{
					drives:   resultDrives[:numDriveResults/2],
					nextLink: &link,
					err:      nil,
				},
				{
					drives:   resultDrives[numDriveResults/2:],
					nextLink: nil,
					err:      nil,
				},
			},
			retry:           false,
			expectedErr:     assert.NoError,
			expectedResults: resultDrives,
		},
		{
			name: "SplitResultsEmptyNextLink",
			pagerResults: []pagerResult{
				{
					drives:   resultDrives[:numDriveResults/2],
					nextLink: &link,
					err:      nil,
				},
				{
					drives:   resultDrives[numDriveResults/2:],
					nextLink: &emptyLink,
					err:      nil,
				},
			},
			retry:           false,
			expectedErr:     assert.NoError,
			expectedResults: resultDrives,
		},
		{
			name: "NonRetryableError",
			pagerResults: []pagerResult{
				{
					drives:   resultDrives,
					nextLink: &link,
					err:      nil,
				},
				{
					drives:   nil,
					nextLink: nil,
					err:      assert.AnError,
				},
			},
			retry:           true,
			expectedErr:     assert.Error,
			expectedResults: nil,
		},
		{
			name: "SiteURLNotFound",
			pagerResults: []pagerResult{
				{
					drives:   nil,
					nextLink: nil,
					err:      graph.Stack(ctx, mySiteURLNotFound),
				},
			},
			retry:           true,
			expectedErr:     assert.NoError,
			expectedResults: nil,
		},
		{
			name: "SiteNotFound",
			pagerResults: []pagerResult{
				{
					drives:   nil,
					nextLink: nil,
					err:      graph.Stack(ctx, mySiteNotFound),
				},
			},
			retry:           true,
			expectedErr:     assert.NoError,
			expectedResults: nil,
		},
		{
			name: "SplitResultsContextTimeoutWithRetries",
			pagerResults: []pagerResult{
				{
					drives:   resultDrives[:numDriveResults/2],
					nextLink: &link,
					err:      nil,
				},
				{
					drives:   nil,
					nextLink: nil,
					err:      context.DeadlineExceeded,
				},
				{
					drives:   resultDrives[numDriveResults/2:],
					nextLink: &emptyLink,
					err:      nil,
				},
			},
			retry:           true,
			expectedErr:     assert.NoError,
			expectedResults: resultDrives,
		},
		{
			name: "SplitResultsContextTimeoutNoRetries",
			pagerResults: []pagerResult{
				{
					drives:   resultDrives[:numDriveResults/2],
					nextLink: &link,
					err:      nil,
				},
				{
					drives:   nil,
					nextLink: nil,
					err:      deadlineExceeded,
				},
				{
					drives:   resultDrives[numDriveResults/2:],
					nextLink: &emptyLink,
					err:      nil,
				},
			},
			retry:           false,
			expectedErr:     assert.Error,
			expectedResults: nil,
		},
		{
			name: "TooManyRetries",
			pagerResults: append(
				[]pagerResult{
					{
						drives:   resultDrives[:numDriveResults/2],
						nextLink: &link,
						err:      nil,
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

			ctx, flush := tester.NewContext()
			defer flush()

			pager := &mockDrivePager{
				toReturn: test.pagerResults,
			}

			drives, err := drives(ctx, pager, test.retry)
			test.expectedErr(t, err)

			assert.ElementsMatch(t, test.expectedResults, drives)
		})
	}
}

// Integration tests

type OneDriveSuite struct {
	tester.Suite
	userID string
}

func TestOneDriveDriveSuite(t *testing.T) {
	suite.Run(t, &OneDriveSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
			tester.CorsoGraphConnectorTests,
			tester.CorsoGraphConnectorOneDriveTests),
	})
}

func (suite *OneDriveSuite) SetupSuite() {
	suite.userID = tester.SecondaryM365UserID(suite.T())
}

func (suite *OneDriveSuite) TestCreateGetDeleteFolder() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	folderIDs := []string{}
	folderName1 := "Corso_Folder_Test_" + common.FormatNow(common.SimpleTimeTesting)
	folderElements := []string{folderName1}
	gs := loadTestService(t)

	pager, err := PagerForSource(OneDriveSource, gs, suite.userID, nil)
	require.NoError(t, err)

	drives, err := drives(ctx, pager, true)
	require.NoError(t, err)
	require.NotEmpty(t, drives)

	// TODO: Verify the intended drive
	driveID := *drives[0].GetId()

	defer func() {
		for _, id := range folderIDs {
			err := DeleteItem(ctx, gs, driveID, id)
			if err != nil {
				logger.Ctx(ctx).Warnw("deleting folder", "id", id, "error", err)
			}
		}
	}()

	folderID, err := CreateRestoreFolders(ctx, gs, driveID, folderElements)
	require.NoError(t, err)

	folderIDs = append(folderIDs, folderID)

	folderName2 := "Corso_Folder_Test_" + common.FormatNow(common.SimpleTimeTesting)
	folderElements = append(folderElements, folderName2)

	folderID, err = CreateRestoreFolders(ctx, gs, driveID, folderElements)
	require.NoError(t, err)

	folderIDs = append(folderIDs, folderID)

	table := []struct {
		name   string
		prefix string
	}{
		{
			name:   "NoPrefix",
			prefix: "",
		},
		{
			name:   "Prefix",
			prefix: "Corso_Folder_Test",
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			pager, err := PagerForSource(OneDriveSource, gs, suite.userID, nil)
			require.NoError(t, err)

			allFolders, err := GetAllFolders(ctx, gs, pager, test.prefix, fault.New(true))
			require.NoError(t, err)

			foundFolderIDs := []string{}

			for _, f := range allFolders {

				if *f.GetName() == folderName1 || *f.GetName() == folderName2 {
					foundFolderIDs = append(foundFolderIDs, *f.GetId())
				}

				assert.True(t, strings.HasPrefix(*f.GetName(), test.prefix), "folder prefix")
			}

			assert.ElementsMatch(t, folderIDs, foundFolderIDs)
		})
	}
}

type testFolderMatcher struct {
	scope selectors.OneDriveScope
}

func (fm testFolderMatcher) IsAny() bool {
	return fm.scope.IsAny(selectors.OneDriveFolder)
}

func (fm testFolderMatcher) Matches(path string) bool {
	return fm.scope.Matches(selectors.OneDriveFolder, path)
}

func (suite *OneDriveSuite) TestOneDriveNewCollections() {
	ctx, flush := tester.NewContext()
	defer flush()

	creds, err := tester.NewM365Account(suite.T()).M365Config()
	require.NoError(suite.T(), err)

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

			service := loadTestService(t)
			scope := selectors.
				NewOneDriveBackup([]string{test.user}).
				AllData()[0]
			odcs, excludes, err := NewCollections(
				graph.HTTPClient(graph.NoTimeout()),
				creds.AzureTenantID,
				test.user,
				OneDriveSource,
				testFolderMatcher{scope},
				service,
				service.updateStatus,
				control.Options{ToggleFeatures: control.Toggles{EnablePermissionsBackup: true}},
			).Get(ctx, nil, fault.New(true))
			assert.NoError(t, err)
			// Don't expect excludes as this isn't an incremental backup.
			assert.Empty(t, excludes)

			for _, entry := range odcs {
				assert.NotEmpty(t, entry.FullPath())
			}
		})
	}
}
