package pathtransformer_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/operations/pathtransformer"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type RestorePathTransformerUnitSuite struct {
	tester.Suite
}

func TestRestorePathTransformerUnitSuite(t *testing.T) {
	suite.Run(t, &RestorePathTransformerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RestorePathTransformerUnitSuite) TestGetPaths() {
	type expectPaths struct {
		storage         string
		restore         string
		isRestorePrefix bool
	}

	toRestore := func(
		repoRef path.Path,
		unescapedFolders ...string,
	) string {
		pfx, _ := repoRef.Halves()
		return pfx.
			Append(unescapedFolders...).
			String()
	}

	var (
		driveID                = "some-drive-id"
		extraItemName          = "some-item"
		SharePointRootItemPath = testdata.SharePointRootPath.MustAppend(extraItemName, true)
	)

	table := []struct {
		name          string
		backupVersion int
		input         []*details.Entry
		expectErr     assert.ErrorAssertionFunc
		expected      []expectPaths
	}{
		{
			name: "SharePoint List Errors",
			// No version bump for the change so we always have to check for this.
			backupVersion: version.All8MigrateUserPNToID,
			input: []*details.Entry{
				{
					RepoRef:     SharePointRootItemPath.RR.String(),
					LocationRef: SharePointRootItemPath.Loc.String(),
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType: details.SharePointList,
						},
					},
				},
			},
			expectErr: assert.Error,
		},
		{
			name: "SharePoint Page Errors",
			// No version bump for the change so we always have to check for this.
			backupVersion: version.All8MigrateUserPNToID,
			input: []*details.Entry{
				{
					RepoRef:     SharePointRootItemPath.RR.String(),
					LocationRef: SharePointRootItemPath.Loc.String(),
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType: details.SharePointPage,
						},
					},
				},
			},
			expectErr: assert.Error,
		},
		{
			name: "SharePoint old format, item in root",
			// No version bump for the change so we always have to check for this.
			backupVersion: version.All8MigrateUserPNToID,
			input: []*details.Entry{
				{
					RepoRef:     SharePointRootItemPath.RR.String(),
					LocationRef: SharePointRootItemPath.Loc.String(),
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType: details.OneDriveItem,
							DriveID:  driveID,
						},
					},
				},
			},
			expectErr: assert.NoError,
			expected: []expectPaths{
				{
					storage: SharePointRootItemPath.RR.String(),
					restore: toRestore(
						SharePointRootItemPath.RR,
						append(
							[]string{"drives", driveID},
							SharePointRootItemPath.Loc.Elements()...)...),
				},
			},
		},
		{
			name:          "SharePoint, no LocationRef, no DriveID, item in root",
			backupVersion: version.OneDrive6NameInMeta,
			input: []*details.Entry{
				{
					RepoRef: SharePointRootItemPath.RR.String(),
					ItemInfo: details.ItemInfo{
						SharePoint: &details.SharePointInfo{
							ItemType: details.SharePointLibrary,
						},
					},
				},
			},
			expectErr: assert.NoError,
			expected: []expectPaths{
				{
					storage: SharePointRootItemPath.RR.String(),
					restore: toRestore(
						SharePointRootItemPath.RR,
						append(
							[]string{"drives"},
							// testdata path has '.d' on the drives folder we need to remove.
							SharePointRootItemPath.RR.Folders()[1:]...)...),
				},
			},
		},
		{
			name:          "OneDrive, nested item",
			backupVersion: version.All8MigrateUserPNToID,
			input: []*details.Entry{
				{
					RepoRef:     testdata.OneDriveItemPath2.RR.String(),
					LocationRef: testdata.OneDriveItemPath2.Loc.String(),
					ItemInfo: details.ItemInfo{
						OneDrive: &details.OneDriveInfo{
							ItemType: details.OneDriveItem,
							DriveID:  driveID,
						},
					},
				},
			},
			expectErr: assert.NoError,
			expected: []expectPaths{
				{
					storage: testdata.OneDriveItemPath2.RR.String(),
					restore: toRestore(
						testdata.OneDriveItemPath2.RR,
						append(
							[]string{"drives", driveID},
							testdata.OneDriveItemPath2.Loc.Elements()...)...),
				},
			},
		},
		{
			name:          "Exchange Email, extra / in path",
			backupVersion: version.All8MigrateUserPNToID,
			input: []*details.Entry{
				{
					RepoRef:     testdata.ExchangeEmailItemPath3.RR.String(),
					LocationRef: testdata.ExchangeEmailItemPath3.Loc.String(),
					ItemInfo: details.ItemInfo{
						Exchange: &details.ExchangeInfo{
							ItemType: details.ExchangeMail,
						},
					},
				},
			},
			expectErr: assert.NoError,
			expected: []expectPaths{
				{
					storage: testdata.ExchangeEmailItemPath3.RR.String(),
					restore: toRestore(
						testdata.ExchangeEmailItemPath3.RR,
						testdata.ExchangeEmailItemPath3.Loc.Elements()...),
				},
			},
		},
		{
			name:          "Exchange Email, no LocationRef, extra / in path",
			backupVersion: version.OneDrive7LocationRef,
			input: []*details.Entry{
				{
					RepoRef:     testdata.ExchangeEmailItemPath3.RR.String(),
					LocationRef: testdata.ExchangeEmailItemPath3.Loc.String(),
					ItemInfo: details.ItemInfo{
						Exchange: &details.ExchangeInfo{
							ItemType: details.ExchangeMail,
						},
					},
				},
			},
			expectErr: assert.NoError,
			expected: []expectPaths{
				{
					storage: testdata.ExchangeEmailItemPath3.RR.String(),
					restore: toRestore(
						testdata.ExchangeEmailItemPath3.RR,
						testdata.ExchangeEmailItemPath3.Loc.Elements()...),
				},
			},
		},
		{
			name:          "Exchange Contact",
			backupVersion: version.All8MigrateUserPNToID,
			input: []*details.Entry{
				{
					RepoRef:     testdata.ExchangeContactsItemPath1.RR.String(),
					LocationRef: testdata.ExchangeContactsItemPath1.Loc.String(),
					ItemInfo: details.ItemInfo{
						Exchange: &details.ExchangeInfo{
							ItemType: details.ExchangeContact,
						},
					},
				},
			},
			expectErr: assert.NoError,
			expected: []expectPaths{
				{
					storage: testdata.ExchangeContactsItemPath1.RR.String(),
					restore: toRestore(
						testdata.ExchangeContactsItemPath1.RR,
						testdata.ExchangeContactsItemPath1.Loc.Elements()...),
				},
			},
		},
		{
			name:          "Exchange Contact, root dir",
			backupVersion: version.All8MigrateUserPNToID,
			input: []*details.Entry{
				{
					RepoRef: testdata.ExchangeContactsItemPath1.RR.String(),
					ItemInfo: details.ItemInfo{
						Exchange: &details.ExchangeInfo{
							ItemType: details.ExchangeContact,
						},
					},
				},
			},
			expectErr: assert.NoError,
			expected: []expectPaths{
				{
					storage:         testdata.ExchangeContactsItemPath1.RR.String(),
					restore:         toRestore(testdata.ExchangeContactsItemPath1.RR, "tmp"),
					isRestorePrefix: true,
				},
			},
		},
		{
			name:          "Exchange Event",
			backupVersion: version.All8MigrateUserPNToID,
			input: []*details.Entry{
				{
					RepoRef:     testdata.ExchangeEmailItemPath3.RR.String(),
					LocationRef: testdata.ExchangeEmailItemPath3.Loc.String(),
					ItemInfo: details.ItemInfo{
						Exchange: &details.ExchangeInfo{
							ItemType: details.ExchangeMail,
						},
					},
				},
			},
			expectErr: assert.NoError,
			expected: []expectPaths{
				{
					storage: testdata.ExchangeEmailItemPath3.RR.String(),
					restore: toRestore(
						testdata.ExchangeEmailItemPath3.RR,
						testdata.ExchangeEmailItemPath3.Loc.Elements()...),
				},
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			paths, err := pathtransformer.GetPaths(
				ctx,
				test.backupVersion,
				test.input,
				fault.New(true))
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			expected := make([]path.RestorePaths, 0, len(test.expected))

			for _, e := range test.expected {
				tmp := path.RestorePaths{}
				p, err := path.FromDataLayerPath(e.storage, true)
				require.NoError(t, err, "parsing expected storage path", clues.ToCore(err))

				tmp.StoragePath = p

				p, err = path.FromDataLayerPath(e.restore, false)
				require.NoError(t, err, "parsing expected restore path", clues.ToCore(err))

				if e.isRestorePrefix {
					p, err = p.Dir()
					require.NoError(t, err, "getting service prefix", clues.ToCore(err))
				}

				tmp.RestorePath = p

				expected = append(expected, tmp)
			}

			assert.ElementsMatch(t, expected, paths)
		})
	}
}
