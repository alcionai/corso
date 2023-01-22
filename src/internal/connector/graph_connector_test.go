package connector

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// Unit tests
// ---------------------------------------------------------------------------

type GraphConnectorUnitSuite struct {
	suite.Suite
}

func TestGraphConnectorUnitSuite(t *testing.T) {
	suite.Run(t, new(GraphConnectorUnitSuite))
}

func (suite *GraphConnectorUnitSuite) TestUnionSiteIDsAndWebURLs() {
	const (
		url1  = "www.foo.com/bar"
		url2  = "www.fnords.com/smarf"
		path1 = "bar"
		path2 = "/smarf"
		id1   = "site-id-1"
		id2   = "site-id-2"
	)

	gc := &GraphConnector{
		// must be populated, else the func will try to make a graph call
		// to retrieve site data.
		Sites: map[string]string{
			url1: id1,
			url2: id2,
		},
	}

	table := []struct {
		name   string
		ids    []string
		urls   []string
		expect []string
	}{
		{
			name: "nil",
		},
		{
			name:   "empty",
			ids:    []string{},
			urls:   []string{},
			expect: []string{},
		},
		{
			name:   "ids only",
			ids:    []string{id1, id2},
			urls:   []string{},
			expect: []string{id1, id2},
		},
		{
			name:   "urls only",
			ids:    []string{},
			urls:   []string{url1, url2},
			expect: []string{id1, id2},
		},
		{
			name:   "url suffix only",
			ids:    []string{},
			urls:   []string{path1, path2},
			expect: []string{id1, id2},
		},
		{
			name:   "url and suffix overlap",
			ids:    []string{},
			urls:   []string{url1, url2, path1, path2},
			expect: []string{id1, id2},
		},
		{
			name:   "ids and urls, no overlap",
			ids:    []string{id1},
			urls:   []string{url2},
			expect: []string{id1, id2},
		},
		{
			name:   "ids and urls, overlap",
			ids:    []string{id1, id2},
			urls:   []string{url1, url2},
			expect: []string{id1, id2},
		},
		{
			name:   "partial non-match on path",
			ids:    []string{},
			urls:   []string{path1[2:], path2[2:]},
			expect: []string{},
		},
		{
			name:   "partial non-match on url",
			ids:    []string{},
			urls:   []string{url1[5:], url2[5:]},
			expect: []string{},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			//nolint
			result, err := gc.UnionSiteIDsAndWebURLs(context.Background(), test.ids, test.urls)
			assert.NoError(t, err)
			assert.ElementsMatch(t, test.expect, result)
		})
	}
}

// ---------------------------------------------------------------------------
// Integration tests
// ---------------------------------------------------------------------------

type GraphConnectorIntegrationSuite struct {
	suite.Suite
	connector     *GraphConnector
	user          string
	secondaryUser string
	acct          account.Account
}

func TestGraphConnectorIntegrationSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests)

	suite.Run(t, new(GraphConnectorIntegrationSuite))
}

func (suite *GraphConnectorIntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext()
	defer flush()

	tester.MustGetEnvSets(suite.T(), tester.M365AcctCredEnvs)

	suite.connector = loadConnector(ctx, suite.T(), Users)
	suite.user = tester.M365UserID(suite.T())
	suite.secondaryUser = tester.SecondaryM365UserID(suite.T())
	suite.acct = tester.NewM365Account(suite.T())

	tester.LogTimeOfTest(suite.T())
}

// TestSetTenantUsers verifies GraphConnector's ability to query
// the users associated with the credentials
func (suite *GraphConnectorIntegrationSuite) TestSetTenantUsers() {
	newConnector := GraphConnector{
		tenant:      "test_tenant",
		Users:       make(map[string]string, 0),
		credentials: suite.connector.credentials,
	}

	ctx, flush := tester.NewContext()
	defer flush()

	service, err := newConnector.createService()
	require.NoError(suite.T(), err)

	newConnector.Service = service

	suite.Empty(len(newConnector.Users))
	err = newConnector.setTenantUsers(ctx)
	suite.NoError(err)
	suite.Less(0, len(newConnector.Users))
}

// TestSetTenantUsers verifies GraphConnector's ability to query
// the sites associated with the credentials
func (suite *GraphConnectorIntegrationSuite) TestSetTenantSites() {
	newConnector := GraphConnector{
		tenant:      "test_tenant",
		Sites:       make(map[string]string, 0),
		credentials: suite.connector.credentials,
	}

	ctx, flush := tester.NewContext()
	defer flush()

	service, err := newConnector.createService()
	require.NoError(suite.T(), err)

	newConnector.Service = service

	suite.Equal(0, len(newConnector.Sites))
	err = newConnector.setTenantSites(ctx)
	suite.NoError(err)
	suite.Less(0, len(newConnector.Sites))

	for _, site := range newConnector.Sites {
		suite.NotContains("sharepoint.com/personal/", site)
	}
}

func (suite *GraphConnectorIntegrationSuite) TestRestoreFailsBadService() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t    = suite.T()
		acct = tester.NewM365Account(t)
		dest = tester.DefaultTestRestoreDestination()
		sel  = selectors.Selector{
			Service: selectors.ServiceUnknown,
		}
	)

	deets, err := suite.connector.RestoreDataCollections(ctx, acct, sel, dest, nil)
	assert.Error(t, err)
	assert.NotNil(t, deets)

	status := suite.connector.AwaitStatus()
	assert.Equal(t, 0, status.ObjectCount)
	assert.Equal(t, 0, status.FolderCount)
	assert.Equal(t, 0, status.Successful)
}

func (suite *GraphConnectorIntegrationSuite) TestEmptyCollections() {
	dest := tester.DefaultTestRestoreDestination()
	table := []struct {
		name string
		col  []data.Collection
		sel  selectors.Selector
	}{
		{
			name: "ExchangeNil",
			col:  nil,
			sel: selectors.Selector{
				Service: selectors.ServiceExchange,
			},
		},
		{
			name: "ExchangeEmpty",
			col:  []data.Collection{},
			sel: selectors.Selector{
				Service: selectors.ServiceExchange,
			},
		},
		{
			name: "OneDriveNil",
			col:  nil,
			sel: selectors.Selector{
				Service: selectors.ServiceOneDrive,
			},
		},
		{
			name: "OneDriveEmpty",
			col:  []data.Collection{},
			sel: selectors.Selector{
				Service: selectors.ServiceOneDrive,
			},
		},
		{
			name: "SharePointNil",
			col:  nil,
			sel: selectors.Selector{
				Service: selectors.ServiceSharePoint,
			},
		},
		{
			name: "SharePointEmpty",
			col:  []data.Collection{},
			sel: selectors.Selector{
				Service: selectors.ServiceSharePoint,
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			deets, err := suite.connector.RestoreDataCollections(
				ctx,
				suite.acct,
				test.sel,
				dest,
				test.col)
			require.NoError(t, err)
			assert.NotNil(t, deets)

			stats := suite.connector.AwaitStatus()
			assert.Zero(t, stats.ObjectCount)
			assert.Zero(t, stats.FolderCount)
			assert.Zero(t, stats.Successful)
		})
	}
}

//-------------------------------------------------------------
// Exchange Functions
//-------------------------------------------------------------

//revive:disable:context-as-argument
func mustGetDefaultDriveID(
	t *testing.T,
	ctx context.Context,
	service graph.Servicer,
	userID string,
) string {
	//revive:enable:context-as-argument
	d, err := service.Client().UsersById(userID).Drive().Get(ctx, nil)
	if err != nil {
		err = errors.Wrapf(
			err,
			"failed to retrieve default user drive. user: %s, details: %s",
			userID,
			support.ConnectorStackErrorTrace(err),
		)
	}

	require.NoError(t, err)
	require.NotNil(t, d.GetId())
	require.NotEmpty(t, *d.GetId())

	return *d.GetId()
}

func runRestoreBackupTest(
	t *testing.T,
	acct account.Account,
	test restoreBackupInfo,
	tenant string,
	resourceOwners []string,
) {
	var (
		collections      []data.Collection
		expectedData     = map[string]map[string][]byte{}
		totalItems       = 0
		toatalKopiaItems = 0
		// Get a dest per test so they're independent.
		dest = tester.DefaultTestRestoreDestination()
	)

	ctx, flush := tester.NewContext()
	defer flush()

	for _, owner := range resourceOwners {
		numItems, kopiaItems, ownerCollections, userExpectedData := collectionsForInfo(
			t,
			test.service,
			tenant,
			owner,
			dest,
			test.collections,
		)

		collections = append(collections, ownerCollections...)
		totalItems += numItems
		toatalKopiaItems += kopiaItems

		maps.Copy(expectedData, userExpectedData)
	}

	t.Logf(
		"Restoring collections to %s for resourceOwners(s) %v\n",
		dest.ContainerName,
		resourceOwners,
	)

	start := time.Now()

	restoreGC := loadConnector(ctx, t, test.resource)
	restoreSel := getSelectorWith(t, test.service, resourceOwners, true)
	deets, err := restoreGC.RestoreDataCollections(
		ctx,
		acct,
		restoreSel,
		dest,
		collections)
	require.NoError(t, err)
	assert.NotNil(t, deets)

	status := restoreGC.AwaitStatus()
	runTime := time.Since(start)

	assert.Equal(t, totalItems, status.ObjectCount, "status.ObjectCount")
	assert.Equal(t, totalItems, status.Successful, "status.Successful")
	assert.Len(
		t,
		deets.Entries,
		totalItems,
		"details entries contains same item count as total successful items restored")

	t.Logf("Restore complete in %v\n", runTime)

	// Run a backup and compare its output with what we put in.
	cats := make(map[path.CategoryType]struct{}, len(test.collections))
	for _, c := range test.collections {
		cats[c.category] = struct{}{}
	}

	expectedDests := make([]destAndCats, 0, len(resourceOwners))
	for _, ro := range resourceOwners {
		expectedDests = append(expectedDests, destAndCats{
			resourceOwner: ro,
			dest:          dest.ContainerName,
			cats:          cats,
		})
	}

	backupGC := loadConnector(ctx, t, test.resource)
	backupSel := backupSelectorForExpected(t, test.service, expectedDests)
	t.Logf("Selective backup of %s\n", backupSel)

	start = time.Now()
	dcs, err := backupGC.DataCollections(ctx, backupSel, nil, control.Options{})
	require.NoError(t, err)

	t.Logf("Backup enumeration complete in %v\n", time.Since(start))

	// Pull the data prior to waiting for the status as otherwise it will
	// deadlock.
	skipped := checkCollections(t, toatalKopiaItems, expectedData, dcs)

	status = backupGC.AwaitStatus()
	assert.Equal(t, totalItems+skipped, status.ObjectCount, "status.ObjectCount")
	assert.Equal(t, totalItems+skipped, status.Successful, "status.Successful")
}

func getTestMetaJson(t *testing.T, id, user string, roles []string) []byte {
	testMeta := onedrive.Metadata{Permissions: []onedrive.UserPermission{
		{ID: id, Roles: roles, Email: user},
	}}
	testMetaJson, err := json.Marshal(testMeta)
	if err != nil {
		t.Fatal("unable to marshall test permissions", err)
	}
	return testMetaJson
}

func (suite *GraphConnectorIntegrationSuite) TestRestoreAndBackup() {
	bodyText := "This email has some text. However, all the text is on the same line."
	subjectText := "Test message for restore"

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
			name:     "EmailsWithAttachments",
			service:  path.ExchangeService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{"Inbox"},
					category:     path.EmailCategory,
					items: []itemInfo{
						{
							name: "someencodeditemID",
							data: mockconnector.GetMockMessageWithDirectAttachment(
								subjectText + "-1",
							),
							lookupKey: subjectText + "-1",
						},
						{
							name: "someencodeditemID2",
							data: mockconnector.GetMockMessageWithTwoAttachments(
								subjectText + "-2",
							),
							lookupKey: subjectText + "-2",
						},
					},
				},
			},
		},
		{
			name:     "MultipleEmailsMultipleFolders",
			service:  path.ExchangeService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{"Inbox"},
					category:     path.EmailCategory,
					items: []itemInfo{
						{
							name: "someencodeditemID",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-1",
								bodyText+" 1.",
								bodyText+" 1.",
							),
							lookupKey: subjectText + "-1",
						},
					},
				},
				{
					pathElements: []string{"Work"},
					category:     path.EmailCategory,
					items: []itemInfo{
						{
							name: "someencodeditemID2",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-2",
								bodyText+" 2.",
								bodyText+" 2.",
							),
							lookupKey: subjectText + "-2",
						},
						{
							name: "someencodeditemID3",
							data: mockconnector.GetMockMessageWithBodyBytes(
								subjectText+"-3",
								bodyText+" 3.",
								bodyText+" 3.",
							),
							lookupKey: subjectText + "-3",
						},
					},
				},
			},
		},
		{
			name:     "MultipleContactsSingleFolder",
			service:  path.ExchangeService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{"Contacts"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockContactBytes("Irgot"),
							lookupKey: "Irgot",
						},
						{
							name:      "someencodeditemID3",
							data:      mockconnector.GetMockContactBytes("Jannes"),
							lookupKey: "Jannes",
						},
					},
				},
			},
		},
		{
			name:     "MultipleContactsMultipleFolders",
			service:  path.ExchangeService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{"Work"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockContactBytes("Irgot"),
							lookupKey: "Irgot",
						},
						{
							name:      "someencodeditemID3",
							data:      mockconnector.GetMockContactBytes("Jannes"),
							lookupKey: "Jannes",
						},
					},
				},
				{
					pathElements: []string{"Personal"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID4",
							data:      mockconnector.GetMockContactBytes("Argon"),
							lookupKey: "Argon",
						},
						{
							name:      "someencodeditemID5",
							data:      mockconnector.GetMockContactBytes("Bernard"),
							lookupKey: "Bernard",
						},
					},
				},
			},
		},
		// {
		// 	name:    "MultipleEventsSingleCalendar",
		// 	service: path.ExchangeService,
		// 	collections: []colInfo{
		// 		{
		// 			pathElements: []string{"Work"},
		// 			category:     path.EventsCategory,
		// 			items: []itemInfo{
		// 				{
		// 					name:      "someencodeditemID",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Ghimley"),
		// 					lookupKey: "Ghimley",
		// 				},
		// 				{
		// 					name:      "someencodeditemID2",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Irgot"),
		// 					lookupKey: "Irgot",
		// 				},
		// 				{
		// 					name:      "someencodeditemID3",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Jannes"),
		// 					lookupKey: "Jannes",
		// 				},
		// 			},
		// 		},
		// 	},
		// },
		// {
		// 	name:    "MultipleEventsMultipleCalendars",
		// 	service: path.ExchangeService,
		// 	collections: []colInfo{
		// 		{
		// 			pathElements: []string{"Work"},
		// 			category:     path.EventsCategory,
		// 			items: []itemInfo{
		// 				{
		// 					name:      "someencodeditemID",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Ghimley"),
		// 					lookupKey: "Ghimley",
		// 				},
		// 				{
		// 					name:      "someencodeditemID2",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Irgot"),
		// 					lookupKey: "Irgot",
		// 				},
		// 				{
		// 					name:      "someencodeditemID3",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Jannes"),
		// 					lookupKey: "Jannes",
		// 				},
		// 			},
		// 		},
		// 		{
		// 			pathElements: []string{"Personal"},
		// 			category:     path.EventsCategory,
		// 			items: []itemInfo{
		// 				{
		// 					name:      "someencodeditemID4",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Argon"),
		// 					lookupKey: "Argon",
		// 				},
		// 				{
		// 					name:      "someencodeditemID5",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Bernard"),
		// 					lookupKey: "Bernard",
		// 				},
		// 			},
		// 		},
		// 	},
		// },
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
							data:      getTestMetaJson(suite.T(), "testMeta2ID", suite.secondaryUser, []string{"write"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
						{
							name:      "b" + onedrive.DirMetaFileSuffix,
							data:      getTestMetaJson(suite.T(), "testMetaID", suite.secondaryUser, []string{"read"}),
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
							data:      getTestMetaJson(suite.T(), "testMetaID", suite.secondaryUser, []string{"read"}),
							lookupKey: "test-file.txt" + onedrive.MetaFileSuffix,
						},
					},
				},
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			runRestoreBackupTest(t, suite.acct, test, suite.connector.tenant, []string{suite.user})
		})
	}
}

func (suite *GraphConnectorIntegrationSuite) TestMultiFolderBackupDifferentNames() {
	table := []restoreBackupInfo{
		{
			name:     "Contacts",
			service:  path.ExchangeService,
			resource: Users,
			collections: []colInfo{
				{
					pathElements: []string{"Work"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID",
							data:      mockconnector.GetMockContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
					},
				},
				{
					pathElements: []string{"Personal"},
					category:     path.ContactsCategory,
					items: []itemInfo{
						{
							name:      "someencodeditemID2",
							data:      mockconnector.GetMockContactBytes("Irgot"),
							lookupKey: "Irgot",
						},
					},
				},
			},
		},
		// {
		// 	name:    "Events",
		// 	service: path.ExchangeService,
		// 	collections: []colInfo{
		// 		{
		// 			pathElements: []string{"Work"},
		// 			category:     path.EventsCategory,
		// 			items: []itemInfo{
		// 				{
		// 					name:      "someencodeditemID",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Ghimley"),
		// 					lookupKey: "Ghimley",
		// 				},
		// 			},
		// 		},
		// 		{
		// 			pathElements: []string{"Personal"},
		// 			category:     path.EventsCategory,
		// 			items: []itemInfo{
		// 				{
		// 					name:      "someencodeditemID2",
		// 					data:      mockconnector.GetMockEventWithSubjectBytes("Irgot"),
		// 					lookupKey: "Irgot",
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			restoreSel := getSelectorWith(t, test.service, []string{suite.user}, true)
			expectedDests := make([]destAndCats, 0, len(test.collections))
			allItems := 0
			allExpectedData := map[string]map[string][]byte{}

			for i, collection := range test.collections {
				// Get a dest per collection so they're independent.
				dest := tester.DefaultTestRestoreDestination()
				expectedDests = append(expectedDests, destAndCats{
					resourceOwner: suite.user,
					dest:          dest.ContainerName,
					cats: map[path.CategoryType]struct{}{
						collection.category: {},
					},
				})

				totalItems, _, collections, expectedData := collectionsForInfo(
					t,
					test.service,
					suite.connector.tenant,
					suite.user,
					dest,
					[]colInfo{collection},
				)
				allItems += totalItems

				for k, v := range expectedData {
					allExpectedData[k] = v
				}

				t.Logf(
					"Restoring %v/%v collections to %s\n",
					i+1,
					len(test.collections),
					dest.ContainerName,
				)

				restoreGC := loadConnector(ctx, t, test.resource)
				deets, err := restoreGC.RestoreDataCollections(ctx, suite.acct, restoreSel, dest, collections)
				require.NoError(t, err)
				require.NotNil(t, deets)

				status := restoreGC.AwaitStatus()
				// Always just 1 because it's just 1 collection.
				assert.Equal(t, totalItems, status.ObjectCount, "status.ObjectCount")
				assert.Equal(t, totalItems, status.Successful, "status.Successful")
				assert.Equal(
					t, totalItems, len(deets.Entries),
					"details entries contains same item count as total successful items restored")

				t.Log("Restore complete")
			}

			// Run a backup and compare its output with what we put in.

			backupGC := loadConnector(ctx, t, test.resource)
			backupSel := backupSelectorForExpected(t, test.service, expectedDests)
			t.Log("Selective backup of", backupSel)

			dcs, err := backupGC.DataCollections(ctx, backupSel, nil, control.Options{})
			require.NoError(t, err)

			t.Log("Backup enumeration complete")

			// Pull the data prior to waiting for the status as otherwise it will
			// deadlock.
			skipped := checkCollections(t, allItems, allExpectedData, dcs)

			status := backupGC.AwaitStatus()
			assert.Equal(t, allItems+skipped, status.ObjectCount, "status.ObjectCount")
			assert.Equal(t, allItems+skipped, status.Successful, "status.Successful")
		})
	}
}
