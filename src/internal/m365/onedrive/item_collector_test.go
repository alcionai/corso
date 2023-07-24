package onedrive

import (
	"context"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
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
	odErr.SetError(merr)

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

	tooManyRetries := make([]mock.PagerResult, 0, maxDrivesRetries+1)

	for i := 0; i < maxDrivesRetries+1; i++ {
		tooManyRetries = append(tooManyRetries, mock.PagerResult{
			Err: context.DeadlineExceeded,
		})
	}

	table := []struct {
		name            string
		pagerResults    []mock.PagerResult
		retry           bool
		expectedErr     assert.ErrorAssertionFunc
		expectedResults []models.Driveable
	}{
		{
			name: "AllOneResultNilNextLink",
			pagerResults: []mock.PagerResult{
				{
					Drives:   resultDrives,
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
			pagerResults: []mock.PagerResult{
				{
					Drives:   resultDrives,
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
			pagerResults: []mock.PagerResult{
				{
					Drives:   resultDrives[:numDriveResults/2],
					NextLink: &link,
					Err:      nil,
				},
				{
					Drives:   resultDrives[numDriveResults/2:],
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
			pagerResults: []mock.PagerResult{
				{
					Drives:   resultDrives[:numDriveResults/2],
					NextLink: &link,
					Err:      nil,
				},
				{
					Drives:   resultDrives[numDriveResults/2:],
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
			pagerResults: []mock.PagerResult{
				{
					Drives:   resultDrives,
					NextLink: &link,
					Err:      nil,
				},
				{
					Drives:   nil,
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
			pagerResults: []mock.PagerResult{
				{
					Drives:   nil,
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
			pagerResults: []mock.PagerResult{
				{
					Drives:   nil,
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
			pagerResults: []mock.PagerResult{
				{
					Drives:   resultDrives[:numDriveResults/2],
					NextLink: &link,
					Err:      nil,
				},
				{
					Drives:   nil,
					NextLink: nil,
					Err:      context.DeadlineExceeded,
				},
				{
					Drives:   resultDrives[numDriveResults/2:],
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
			pagerResults: []mock.PagerResult{
				{
					Drives:   resultDrives[:numDriveResults/2],
					NextLink: &link,
					Err:      nil,
				},
				{
					Drives:   nil,
					NextLink: nil,
					Err:      context.DeadlineExceeded,
				},
				{
					Drives:   resultDrives[numDriveResults/2:],
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
				[]mock.PagerResult{
					{
						Drives:   resultDrives[:numDriveResults/2],
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

			pager := &mock.DrivePager{
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

	suite.ac, err = api.NewClient(creds, control.Defaults())
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *OneDriveIntgSuite) TestCreateGetDeleteFolder() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		folderIDs      = []string{}
		folderName1    = "Corso_Folder_Test_" + dttm.FormatNow(dttm.SafeForTesting)
		folderElements = []string{folderName1}
	)

	pager := suite.ac.Drives().NewUserDrivePager(suite.userID, nil)

	drives, err := api.GetAllDrives(ctx, pager, true, maxDrivesRetries)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, drives)

	// TODO: Verify the intended drive
	driveID := ptr.Val(drives[0].GetId())

	defer func() {
		for _, id := range folderIDs {
			ictx := clues.Add(ctx, "folder_id", id)

			// deletes require unique http clients
			// https://github.com/alcionai/corso/issues/2707
			err := suite.ac.Drives().DeleteItem(ictx, driveID, id)
			if err != nil {
				logger.CtxErr(ictx, err).Errorw("deleting folder")
			}
		}
	}()

	rootFolder, err := suite.ac.Drives().GetRootFolder(ctx, driveID)
	require.NoError(t, err, clues.ToCore(err))

	restoreDir := path.Builder{}.Append(folderElements...)
	drivePath := path.DrivePath{
		DriveID: driveID,
		Root:    "root:",
		Folders: folderElements,
	}

	caches := NewRestoreCaches(nil)
	caches.DriveIDToDriveInfo[driveID] = driveInfo{rootFolderID: ptr.Val(rootFolder.GetId())}

	rh := NewRestoreHandler(suite.ac)

	folderID, err := createRestoreFolders(ctx, rh, &drivePath, restoreDir, caches)
	require.NoError(t, err, clues.ToCore(err))

	folderIDs = append(folderIDs, folderID)

	folderName2 := "Corso_Folder_Test_" + dttm.FormatNow(dttm.SafeForTesting)
	restoreDir = restoreDir.Append(folderName2)

	folderID, err = createRestoreFolders(ctx, rh, &drivePath, restoreDir, caches)
	require.NoError(t, err, clues.ToCore(err))

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
			bh := itemBackupHandler{
				suite.ac.Drives(),
				(&selectors.OneDriveBackup{}).Folders(selectors.Any())[0],
			}
			pager := suite.ac.Drives().NewUserDrivePager(suite.userID, nil)

			ctx, flush := tester.NewContext(t)
			defer flush()

			allFolders, err := GetAllFolders(ctx, bh, pager, test.prefix, fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			foundFolderIDs := []string{}

			for _, f := range allFolders {

				if ptr.Val(f.GetName()) == folderName1 || ptr.Val(f.GetName()) == folderName2 {
					foundFolderIDs = append(foundFolderIDs, ptr.Val(f.GetId()))
				}

				assert.True(t, strings.HasPrefix(ptr.Val(f.GetName()), test.prefix), "folder prefix")
			}

			assert.ElementsMatch(t, folderIDs, foundFolderIDs)
		})
	}
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
				&itemBackupHandler{suite.ac.Drives(), scope},
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
