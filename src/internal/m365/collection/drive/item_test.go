package drive

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"syscall"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type ItemIntegrationSuite struct {
	tester.Suite
	user        string
	userDriveID string
	service     *oneDriveService
}

func TestItemIntegrationSuite(t *testing.T) {
	suite.Run(t, &ItemIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *ItemIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.service = loadTestService(t)
	suite.user = tconfig.SecondaryM365UserID(t)

	pager := suite.service.ac.Drives().NewUserDrivePager(suite.user, nil)

	odDrives, err := api.GetAllDrives(ctx, pager, true, maxDrivesRetries)
	require.NoError(t, err, clues.ToCore(err))
	// Test Requirement 1: Need a drive
	require.Greaterf(t, len(odDrives), 0, "user %s does not have a drive", suite.user)

	// Pick the first drive
	suite.userDriveID = ptr.Val(odDrives[0].GetId())
}

// TestItemReader is an integration test that makes a few assumptions
// about the test environment
// 1) It assumes the test user has a drive
// 2) It assumes the drive has a file it can use to test `driveItemReader`
// The test checks these in below
func (suite *ItemIntegrationSuite) TestItemReader_oneDrive() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var driveItem models.DriveItemable
	// This item collector tries to find "a" drive item that is a non-empty
	// file to test the reader function
	itemCollector := func(
		_ context.Context,
		_, _ string,
		items []models.DriveItemable,
		_ map[string]string,
		_ map[string]string,
		_ map[string]struct{},
		_ map[string]map[string]string,
		_ bool,
		_ *fault.Bus,
	) error {
		if driveItem != nil {
			return nil
		}

		for _, item := range items {
			if item.GetFile() != nil && ptr.Val(item.GetSize()) > 0 {
				driveItem = item
				break
			}
		}

		return nil
	}

	ip := suite.service.ac.
		Drives().
		NewDriveItemDeltaPager(suite.userDriveID, "", api.DriveItemSelectDefault())

	_, _, _, err := collectItems(
		ctx,
		ip,
		suite.userDriveID,
		"General",
		itemCollector,
		map[string]string{},
		"",
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	// Test Requirement 2: Need a file
	require.NotEmpty(
		t,
		driveItem,
		"no file item found for user %s drive %s",
		suite.user,
		suite.userDriveID)

	bh := itemBackupHandler{
		suite.service.ac.Drives(),
		suite.user,
		(&selectors.OneDriveBackup{}).Folders(selectors.Any())[0],
	}

	// Read data for the file
	itemData, err := downloadItem(ctx, bh, driveItem)
	require.NoError(t, err, clues.ToCore(err))

	size, err := io.Copy(io.Discard, itemData)
	require.NoError(t, err, clues.ToCore(err))
	require.NotZero(t, size)
}

// TestItemWriter is an integration test for uploading data to OneDrive
// It creates a new folder with a new item and writes data to it
func (suite *ItemIntegrationSuite) TestItemWriter() {
	table := []struct {
		name    string
		driveID string
	}{
		{
			name:    "",
			driveID: suite.userDriveID,
		},
		// {
		// 	name:   "sharePoint",
		// 	driveID: suite.siteDriveID,
		// },
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			rh := NewRestoreHandler(suite.service.ac)

			ctx, flush := tester.NewContext(t)
			defer flush()

			root, err := suite.service.ac.Drives().GetRootFolder(ctx, test.driveID)
			require.NoError(t, err, clues.ToCore(err))

			newFolderName := testdata.DefaultRestoreConfig("folder").Location
			t.Logf("creating folder %s", newFolderName)

			newFolder, err := rh.PostItemInContainer(
				ctx,
				test.driveID,
				ptr.Val(root.GetId()),
				newItem(newFolderName, true),
				control.Copy)
			require.NoError(t, err, clues.ToCore(err))
			require.NotNil(t, newFolder.GetId())

			newItemName := "testItem_" + dttm.FormatNow(dttm.SafeForTesting)
			t.Logf("creating item %s", newItemName)

			newItem, err := rh.PostItemInContainer(
				ctx,
				test.driveID,
				ptr.Val(newFolder.GetId()),
				newItem(newItemName, false),
				control.Copy)
			require.NoError(t, err, clues.ToCore(err))
			require.NotNil(t, newItem.GetId())

			// HACK: Leveraging this to test getFolder behavior for a file. `getFolder()` on the
			// newly created item should fail because it's a file not a folder
			_, err = suite.service.ac.Drives().GetFolderByName(
				ctx,
				test.driveID,
				ptr.Val(newFolder.GetId()),
				newItemName)
			require.ErrorIs(t, err, api.ErrFolderNotFound, clues.ToCore(err))

			// Initialize a 100KB mockDataProvider
			td, writeSize := mockDataReader(int64(100 * 1024))

			w, _, err := driveItemWriter(
				ctx,
				rh,
				test.driveID,
				ptr.Val(newItem.GetId()),
				writeSize)
			require.NoError(t, err, clues.ToCore(err))

			// Using a 32 KB buffer for the copy allows us to validate the
			// multi-part upload. `io.CopyBuffer` will only write 32 KB at
			// a time
			copyBuffer := make([]byte, 32*1024)

			size, err := io.CopyBuffer(w, td, copyBuffer)
			require.NoError(t, err, clues.ToCore(err))

			require.Equal(t, writeSize, size)
		})
	}
}

func mockDataReader(size int64) (io.Reader, int64) {
	data := bytes.Repeat([]byte("D"), int(size))
	return bytes.NewReader(data), size
}

func (suite *ItemIntegrationSuite) TestDriveGetFolder() {
	table := []struct {
		name    string
		driveID string
	}{
		{
			name:    "oneDrive",
			driveID: suite.userDriveID,
		},
		// {
		// 	name:   "sharePoint",
		// 	driveID: suite.siteDriveID,
		// },
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			root, err := suite.service.ac.Drives().GetRootFolder(ctx, test.driveID)
			require.NoError(t, err, clues.ToCore(err))

			// Lookup a folder that doesn't exist
			_, err = suite.service.ac.Drives().GetFolderByName(
				ctx,
				test.driveID,
				ptr.Val(root.GetId()),
				"FolderDoesNotExist")
			require.ErrorIs(t, err, api.ErrFolderNotFound, clues.ToCore(err))

			// Lookup a folder that does exist
			_, err = suite.service.ac.Drives().GetFolderByName(
				ctx,
				test.driveID,
				ptr.Val(root.GetId()),
				"")
			require.NoError(t, err, clues.ToCore(err))
		})
	}
}

// Unit tests

type mockGetter struct {
	GetFunc func(ctx context.Context, url string) (*http.Response, error)
}

func (m mockGetter) Get(
	ctx context.Context,
	url string,
	headers map[string]string,
) (*http.Response, error) {
	return m.GetFunc(ctx, url)
}

type ItemUnitTestSuite struct {
	tester.Suite
}

func TestItemUnitTestSuite(t *testing.T) {
	suite.Run(t, &ItemUnitTestSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ItemUnitTestSuite) TestDownloadItem() {
	testRc := io.NopCloser(bytes.NewReader([]byte("test")))
	url := "https://example.com"

	table := []struct {
		name          string
		itemFunc      func() models.DriveItemable
		GetFunc       func(ctx context.Context, url string) (*http.Response, error)
		errorExpected require.ErrorAssertionFunc
		rcExpected    require.ValueAssertionFunc
		label         string
	}{
		{
			name: "nil item",
			itemFunc: func() models.DriveItemable {
				return nil
			},
			GetFunc: func(ctx context.Context, url string) (*http.Response, error) {
				return nil, nil
			},
			errorExpected: require.Error,
			rcExpected:    require.Nil,
		},
		{
			name: "success",
			itemFunc: func() models.DriveItemable {
				di := newItem("test", false)
				di.SetAdditionalData(map[string]any{
					"@microsoft.graph.downloadUrl": url,
				})

				return di
			},
			GetFunc: func(ctx context.Context, url string) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       testRc,
				}, nil
			},
			errorExpected: require.NoError,
			rcExpected:    require.NotNil,
		},
		{
			name: "success, content url set instead of download url",
			itemFunc: func() models.DriveItemable {
				di := newItem("test", false)
				di.SetAdditionalData(map[string]any{
					"@content.downloadUrl": url,
				})

				return di
			},
			GetFunc: func(ctx context.Context, url string) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       testRc,
				}, nil
			},
			errorExpected: require.NoError,
			rcExpected:    require.NotNil,
		},
		{
			name: "api getter returns error",
			itemFunc: func() models.DriveItemable {
				di := newItem("test", false)
				di.SetAdditionalData(map[string]any{
					"@microsoft.graph.downloadUrl": url,
				})

				return di
			},
			GetFunc: func(ctx context.Context, url string) (*http.Response, error) {
				return nil, clues.New("test error")
			},
			errorExpected: require.Error,
			rcExpected:    require.Nil,
		},
		{
			name: "download url is empty",
			itemFunc: func() models.DriveItemable {
				di := newItem("test", false)
				return di
			},
			GetFunc: func(ctx context.Context, url string) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       testRc,
				}, nil
			},
			errorExpected: require.Error,
			rcExpected:    require.Nil,
		},
		{
			name: "malware",
			itemFunc: func() models.DriveItemable {
				di := newItem("test", false)
				di.SetAdditionalData(map[string]any{
					"@microsoft.graph.downloadUrl": url,
				})

				return di
			},
			GetFunc: func(ctx context.Context, url string) (*http.Response, error) {
				return &http.Response{
					Header: http.Header{
						"X-Virus-Infected": []string{"true"},
					},
					StatusCode: http.StatusOK,
					Body:       testRc,
				}, nil
			},
			errorExpected: require.Error,
			rcExpected:    require.Nil,
		},
		{
			name: "non-2xx http response",
			itemFunc: func() models.DriveItemable {
				di := newItem("test", false)
				di.SetAdditionalData(map[string]any{
					"@microsoft.graph.downloadUrl": url,
				})

				return di
			},
			GetFunc: func(ctx context.Context, url string) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusNotFound,
					Body:       nil,
				}, nil
			},
			errorExpected: require.Error,
			rcExpected:    require.Nil,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			ctx, flush := tester.NewContext(t)
			defer flush()

			mg := mockGetter{
				GetFunc: test.GetFunc,
			}
			rc, err := downloadItem(ctx, mg, test.itemFunc())
			test.errorExpected(t, err, clues.ToCore(err))
			test.rcExpected(t, rc)
		})
	}
}

type errReader struct{}

func (r errReader) Read(p []byte) (int, error) {
	return 0, syscall.ECONNRESET
}

func (suite *ItemUnitTestSuite) TestDownloadItem_ConnectionResetErrorOnFirstRead() {
	var (
		callCount int

		testData = []byte("test")
		testRc   = io.NopCloser(bytes.NewReader(testData))
		url      = "https://example.com"

		itemFunc = func() models.DriveItemable {
			di := newItem("test", false)
			di.SetAdditionalData(map[string]any{
				"@microsoft.graph.downloadUrl": url,
			})

			return di
		}

		GetFunc = func(ctx context.Context, url string) (*http.Response, error) {
			defer func() {
				callCount++
			}()

			if callCount == 0 {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(errReader{}),
				}, nil
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       testRc,
			}, nil
		}
		errorExpected = require.NoError
		rcExpected    = require.NotNil
	)

	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	mg := mockGetter{
		GetFunc: GetFunc,
	}
	rc, err := downloadItem(ctx, mg, itemFunc())
	errorExpected(t, err, clues.ToCore(err))
	rcExpected(t, rc)

	data, err := io.ReadAll(rc)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, testData, data)
}
