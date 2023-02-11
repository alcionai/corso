package connector

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
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
			ctx, flush := tester.NewContext()
			defer flush()

			result, err := gc.UnionSiteIDsAndWebURLs(ctx, test.ids, test.urls, fault.New(true))
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

	suite.connector = loadConnector(ctx, suite.T(), graph.HTTPClient(graph.NoTimeout()), Users)
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

	owners, err := api.NewClient(suite.connector.credentials)
	require.NoError(suite.T(), err)

	newConnector.Owners = owners
	suite.Empty(len(newConnector.Users))

	errs := fault.New(true)

	err = newConnector.setTenantUsers(ctx, errs)
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

	t := suite.T()

	service, err := newConnector.createService()
	require.NoError(t, err)

	newConnector.Service = service
	assert.Equal(t, 0, len(newConnector.Sites))

	err = newConnector.setTenantSites(ctx, fault.New(true))
	assert.NoError(t, err)
	assert.Less(t, 0, len(newConnector.Sites))

	for _, site := range newConnector.Sites {
		assert.NotContains(t, "sharepoint.com/personal/", site)
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

	deets, err := suite.connector.RestoreDataCollections(
		ctx,
		backup.Version,
		acct,
		sel,
		dest,
		control.Options{
			RestorePermissions: true,
			ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
		},
		nil,
		fault.New(true))
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
		col  []data.RestoreCollection
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
			col:  []data.RestoreCollection{},
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
			col:  []data.RestoreCollection{},
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
			col:  []data.RestoreCollection{},
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
				backup.Version,
				suite.acct,
				test.sel,
				dest,
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
				test.col,
				fault.New(true))
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

func getCollectionsAndExpected(
	t *testing.T,
	config configInfo,
	testCollections []colInfo,
	countMeta bool,
) (int, int, []data.RestoreCollection, map[string]map[string][]byte) {
	t.Helper()

	var (
		collections     []data.RestoreCollection
		expectedData    = map[string]map[string][]byte{}
		totalItems      = 0
		totalKopiaItems = 0
	)

	for _, owner := range config.resourceOwners {
		numItems, kopiaItems, ownerCollections, userExpectedData := collectionsForInfo(
			t,
			config.service,
			config.tenant,
			owner,
			config.dest,
			testCollections,
			countMeta,
		)

		collections = append(collections, ownerCollections...)
		totalItems += numItems
		totalKopiaItems += kopiaItems

		maps.Copy(expectedData, userExpectedData)
	}

	return totalItems, totalKopiaItems, collections, expectedData
}

func runRestore(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	config configInfo,
	backupVersion int,
	collections []data.RestoreCollection,
	numRestoreItems int,
) {
	t.Logf(
		"Restoring collections to %s for resourceOwners(s) %v\n",
		config.dest.ContainerName,
		config.resourceOwners)

	start := time.Now()

	restoreGC := loadConnector(ctx, t, graph.HTTPClient(graph.NoTimeout()), config.resource)
	restoreSel := getSelectorWith(t, config.service, config.resourceOwners, true)
	deets, err := restoreGC.RestoreDataCollections(
		ctx,
		backupVersion,
		config.acct,
		restoreSel,
		config.dest,
		config.opts,
		collections,
		fault.New(true))
	require.NoError(t, err)
	assert.NotNil(t, deets)

	status := restoreGC.AwaitStatus()
	runTime := time.Since(start)

	assert.NoError(t, status.Err, "restored status.Err")
	assert.Zero(t, status.ErrorCount, "restored status.ErrorCount")
	assert.Equal(t, numRestoreItems, status.ObjectCount, "restored status.ObjectCount")
	assert.Equal(t, numRestoreItems, status.Successful, "restored status.Successful")
	assert.Len(
		t,
		deets.Entries,
		numRestoreItems,
		"details entries contains same item count as total successful items restored")

	t.Logf("Restore complete in %v\n", runTime)
}

func runBackupAndCompare(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	config configInfo,
	expectedData map[string]map[string][]byte,
	totalItems int,
	totalKopiaItems int,
	inputCollections []colInfo,
) {
	t.Helper()

	// Run a backup and compare its output with what we put in.
	cats := make(map[path.CategoryType]struct{}, len(inputCollections))
	for _, c := range inputCollections {
		cats[c.category] = struct{}{}
	}

	expectedDests := make([]destAndCats, 0, len(config.resourceOwners))
	for _, ro := range config.resourceOwners {
		expectedDests = append(expectedDests, destAndCats{
			resourceOwner: ro,
			dest:          config.dest.ContainerName,
			cats:          cats,
		})
	}

	backupGC := loadConnector(ctx, t, graph.HTTPClient(graph.NoTimeout()), config.resource)
	backupSel := backupSelectorForExpected(t, config.service, expectedDests)
	t.Logf("Selective backup of %s\n", backupSel)

	start := time.Now()
	dcs, excludes, err := backupGC.DataCollections(
		ctx,
		backupSel,
		nil,
		control.Options{
			RestorePermissions: true,
			ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
		},
		fault.New(true))
	require.NoError(t, err)
	// No excludes yet because this isn't an incremental backup.
	assert.Empty(t, excludes)

	t.Logf("Backup enumeration complete in %v\n", time.Since(start))

	// Pull the data prior to waiting for the status as otherwise it will
	// deadlock.
	skipped := checkCollections(t, totalKopiaItems, expectedData, dcs, config.opts.RestorePermissions)

	status := backupGC.AwaitStatus()

	assert.NoError(t, status.Err, "backup status.Err")
	assert.Zero(t, status.ErrorCount, "backup status.ErrorCount")
	assert.Equalf(t, totalItems+skipped, status.ObjectCount,
		"backup status.ObjectCount; wanted %d items + %d skipped", totalItems, skipped)
	assert.Equalf(t, totalItems+skipped, status.Successful,
		"backup status.Successful; wanted %d items + %d skipped", totalItems, skipped)
}

func runRestoreBackupTest(
	t *testing.T,
	acct account.Account,
	test restoreBackupInfo,
	tenant string,
	resourceOwners []string,
	opts control.Options,
) {
	ctx, flush := tester.NewContext()
	defer flush()

	config := configInfo{
		acct:           acct,
		opts:           opts,
		resource:       test.resource,
		service:        test.service,
		tenant:         tenant,
		resourceOwners: resourceOwners,
		dest:           tester.DefaultTestRestoreDestination(),
	}

	totalItems, totalKopiaItems, collections, expectedData := getCollectionsAndExpected(
		t,
		config,
		test.collections,
		false)

	runRestore(
		t,
		ctx,
		config,
		backup.Version,
		collections,
		totalItems)

	runBackupAndCompare(
		t,
		ctx,
		config,
		expectedData,
		totalItems,
		totalKopiaItems,
		test.collections)
}

// runRestoreBackupTestVersions restores with data from an older
// version of the backup and check the restored data against the
// something that would be in the form of a newer backup.
func runRestoreBackupTestVersion0(
	t *testing.T,
	acct account.Account,
	test restoreBackupInfoMultiVersion,
	tenant string,
	resourceOwners []string,
	opts control.Options,
) {
	ctx, flush := tester.NewContext()
	defer flush()

	config := configInfo{
		acct:           acct,
		opts:           opts,
		resource:       test.resource,
		service:        test.service,
		tenant:         tenant,
		resourceOwners: resourceOwners,
		dest:           tester.DefaultTestRestoreDestination(),
	}

	totalItems, _, collections, _ := getCollectionsAndExpected(
		t,
		config,
		test.collectionsPrevious,
		true)

	runRestore(
		t,
		ctx,
		config,
		0, // The OG version ;)
		collections,
		totalItems)

	// Get expected output for new version.
	totalItems, totalKopiaItems, _, expectedData := getCollectionsAndExpected(
		t,
		config,
		test.collectionsLatest,
		false)

	runBackupAndCompare(
		t,
		ctx,
		config,
		expectedData,
		totalItems,
		totalKopiaItems,
		test.collectionsLatest)
}

func (suite *GraphConnectorIntegrationSuite) TestRestoreAndBackup() {
	bodyText := "This email has some text. However, all the text is on the same line."
	subjectText := "Test message for restore"

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
					false,
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

				restoreGC := loadConnector(ctx, t, graph.HTTPClient(graph.NoTimeout()), test.resource)
				deets, err := restoreGC.RestoreDataCollections(
					ctx,
					backup.Version,
					suite.acct,
					restoreSel,
					dest,
					control.Options{
						RestorePermissions: true,
						ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
					},
					collections,
					fault.New(true))
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

			backupGC := loadConnector(ctx, t, graph.HTTPClient(graph.NoTimeout()), test.resource)
			backupSel := backupSelectorForExpected(t, test.service, expectedDests)
			t.Log("Selective backup of", backupSel)

			dcs, excludes, err := backupGC.DataCollections(
				ctx,
				backupSel,
				nil,
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
				},
				fault.New(true))
			require.NoError(t, err)
			// No excludes yet because this isn't an incremental backup.
			assert.Empty(t, excludes)

			t.Log("Backup enumeration complete")

			// Pull the data prior to waiting for the status as otherwise it will
			// deadlock.
			skipped := checkCollections(t, allItems, allExpectedData, dcs, true)

			status := backupGC.AwaitStatus()
			assert.Equal(t, allItems+skipped, status.ObjectCount, "status.ObjectCount")
			assert.Equal(t, allItems+skipped, status.Successful, "status.Successful")
		})
	}
}

// TODO: this should only be run during smoke tests, not part of the standard CI.
// That's why it's set aside instead of being included in the other test set.
func (suite *GraphConnectorIntegrationSuite) TestRestoreAndBackup_largeMailAttachment() {
	subjectText := "Test message for restore with large attachment"

	test := restoreBackupInfo{
		name:     "EmailsWithLargeAttachments",
		service:  path.ExchangeService,
		resource: Users,
		collections: []colInfo{
			{
				pathElements: []string{"Inbox"},
				category:     path.EmailCategory,
				items: []itemInfo{
					{
						name:      "35mbAttachment",
						data:      mockconnector.GetMockMessageWithSizedAttachment(subjectText, 35),
						lookupKey: subjectText,
					},
				},
			},
		},
	}

	runRestoreBackupTest(
		suite.T(),
		suite.acct,
		test,
		suite.connector.tenant,
		[]string{suite.user},
		control.Options{
			RestorePermissions: true,
			ToggleFeatures:     control.Toggles{EnablePermissionsBackup: true},
		},
	)
}
