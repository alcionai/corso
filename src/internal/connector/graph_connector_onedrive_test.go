package connector

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

func getTestMetaJSON(t *testing.T, user string, roles []string) []byte {
	id := base64.StdEncoding.EncodeToString([]byte(user + strings.Join(roles, "+")))
	testMeta := onedrive.Metadata{Permissions: []onedrive.UserPermission{
		{ID: id, Roles: roles, Email: user},
	}}

	testMetaJSON, err := json.Marshal(testMeta)
	if err != nil {
		t.Fatal("unable to marshall test permissions", err)
	}

	return testMetaJSON
}

type GraphConnectorOneDriveIntegrationSuite struct {
	suite.Suite
	connector     *GraphConnector
	user          string
	secondaryUser string
	acct          account.Account
}

func TestGraphConnectorOneDriveIntegrationSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests)

	suite.Run(t, new(GraphConnectorOneDriveIntegrationSuite))
}

func (suite *GraphConnectorOneDriveIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	tester.MustGetEnvSets(suite.T(), tester.M365AcctCredEnvs)

	suite.connector = loadConnector(ctx, suite.T(), graph.HTTPClient(graph.NoTimeout()), Users)
	suite.user = tester.M365UserID(suite.T())
	suite.secondaryUser = tester.SecondaryM365UserID(suite.T())
	suite.acct = tester.NewM365Account(suite.T())

	tester.LogTimeOfTest(suite.T())
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestRestoreAndBackup() {
	ctx, flush := tester.NewContext()
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.connector.Service,
		suite.user,
	)

	table := []restoreBackupInfo{
		{
			name:     "OneDriveMultipleFoldersAndFiles",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("a", 33)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
						{
							name:      "folder-a" + onedrive.DirMetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "folder-a" + onedrive.DirMetaFileSuffix,
						},
						{
							name:      "b" + onedrive.DirMetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "b" + onedrive.DirMetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("b", 65)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
						{
							name:      "b" + onedrive.DirMetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "b" + onedrive.DirMetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("c", 129)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
						{
							name:      "folder-a" + onedrive.DirMetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "folder-a" + onedrive.DirMetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
						"b",
						"folder-a",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("d", 257)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("e", 257)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
			},
		},
		{
			name:     "OneDriveFoldersAndFilesWithMetadata",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("a", 33)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
						{
							name:      "b" + onedrive.DirMetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
							lookupKey: "b" + onedrive.DirMetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("e", 66)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			runRestoreBackupTest(
				t,
				suite.acct,
				test,
				suite.connector.tenant,
				[]string{suite.user},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestRestoreAndBackupVersion0() {
	ctx, flush := tester.NewContext()
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.connector.Service,
		suite.user,
	)

	table := []restoreBackupInfoMultiVersion{
		{
			name:     "OneDriveMultipleFoldersAndFiles",
			service:  path.OneDriveService,
			resource: Users,

			collectionsPrevious: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt",
							data:      []byte(strings.Repeat("a", 33)),
							lookupKey: "test-file.txt",
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt",
							data:      []byte(strings.Repeat("b", 65)),
							lookupKey: "test-file.txt",
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt",
							data:      []byte(strings.Repeat("c", 129)),
							lookupKey: "test-file.txt",
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
						"b",
						"folder-a",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt",
							data:      []byte(strings.Repeat("d", 257)),
							lookupKey: "test-file.txt",
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt",
							data:      []byte(strings.Repeat("e", 257)),
							lookupKey: "test-file.txt",
						},
					},
				},
			},

			collectionsLatest: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("a", 33)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
						{
							name:      "folder-a" + onedrive.DirMetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "folder-a" + onedrive.DirMetaFileSuffix,
						},
						{
							name:      "b" + onedrive.DirMetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "b" + onedrive.DirMetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("b", 65)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
						{
							name:      "b" + onedrive.DirMetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "b" + onedrive.DirMetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("c", 129)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
						{
							name:      "folder-a" + onedrive.DirMetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "folder-a" + onedrive.DirMetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"folder-a",
						"b",
						"folder-a",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("d", 257)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("e", 257)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			runRestoreBackupTestVersions(
				t,
				suite.acct,
				test,
				suite.connector.tenant,
				[]string{suite.user},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestPermissionsRestoreAndBackup() {
	ctx, flush := tester.NewContext()
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.connector.Service,
		suite.user,
	)

	table := []restoreBackupInfo{
		{
			name:     "FilePermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("a", 33)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
			},
		},

		{
			name:     "FileInsideFolderPermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("a", 33)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
						{
							name:      "b" + onedrive.DirMetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "b" + onedrive.DirMetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("e", 66)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
			},
		},

		{
			name:     "FileAndFolderPermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("a", 33)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
						{
							name:      "b" + onedrive.DirMetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
							lookupKey: "b" + onedrive.DirMetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("e", 66)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
			},
		},

		{
			name:     "FileAndFolderSeparatePermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "b" + onedrive.DirMetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
							lookupKey: "b" + onedrive.DirMetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("e", 66)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
			},
		},

		{
			name:     "FolderAndNoChildPermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "b" + onedrive.DirMetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"read"}),
							lookupKey: "b" + onedrive.DirMetaFileSuffix,
						},
					},
				},
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
						"b",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("e", 66)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      []byte("{}"),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			runRestoreBackupTest(t,
				suite.acct,
				test,
				suite.connector.tenant,
				[]string{suite.user},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}

func (suite *GraphConnectorOneDriveIntegrationSuite) TestPermissionsBackupAndNoRestore() {
	ctx, flush := tester.NewContext()
	defer flush()

	// Get the default drive ID for the test user.
	driveID := mustGetDefaultDriveID(
		suite.T(),
		ctx,
		suite.connector.Service,
		suite.user,
	)

	table := []restoreBackupInfo{
		{
			name:     "FilePermissionsRestore",
			service:  path.OneDriveService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{
						"drives",
						driveID,
						"root:",
					},
					category: path.FilesCategory,
					items: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.DataFileSuffix,
							data:      []byte(strings.Repeat("a", 33)),
							lookupKey: "test-file.txt" + onedrive.DataFileSuffix,
						},
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
					auxItems: []itemInfo{
						{
							name:      "test-file.txt" + onedrive.MetaFileSuffix,
							data:      getTestMetaJSON(suite.T(), suite.secondaryUser, []string{"write"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			runRestoreBackupTest(
				t,
				suite.acct,
				test,
				suite.connector.tenant,
				[]string{suite.user},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
			)
		})
	}
}
