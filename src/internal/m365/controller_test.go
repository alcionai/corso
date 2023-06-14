package m365

import (
	"context"
	"runtime/trace"
	"sync"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/data"
	exchMock "github.com/alcionai/corso/src/internal/m365/exchange/mock"
	"github.com/alcionai/corso/src/internal/m365/mock"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ---------------------------------------------------------------------------
// Unit tests
// ---------------------------------------------------------------------------

type ControllerUnitSuite struct {
	tester.Suite
}

func TestControllerUnitSuite(t *testing.T) {
	suite.Run(t, &ControllerUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ControllerUnitSuite) TestPopulateOwnerIDAndNamesFrom() {
	const (
		id   = "owner-id"
		name = "owner-name"
	)

	var (
		itn    = map[string]string{id: name}
		nti    = map[string]string{name: id}
		lookup = &resourceClient{
			enum:   resource.Users,
			getter: &mock.IDNameGetter{ID: id, Name: name},
		}
		noLookup = &resourceClient{enum: resource.Users, getter: &mock.IDNameGetter{}}
	)

	table := []struct {
		name       string
		owner      string
		ins        inMock.Cache
		rc         *resourceClient
		expectID   string
		expectName string
		expectErr  require.ErrorAssertionFunc
	}{
		{
			name:       "nil ins",
			owner:      id,
			rc:         lookup,
			expectID:   id,
			expectName: name,
			expectErr:  require.NoError,
		},
		{
			name:       "nil ins no lookup",
			owner:      id,
			rc:         noLookup,
			expectID:   "",
			expectName: "",
			expectErr:  require.Error,
		},
		{
			name:       "only id map with owner id",
			owner:      id,
			ins:        inMock.NewCache(itn, nil),
			rc:         noLookup,
			expectID:   id,
			expectName: name,
			expectErr:  require.NoError,
		},
		{
			name:       "only name map with owner id",
			owner:      id,
			ins:        inMock.NewCache(nil, nti),
			rc:         noLookup,
			expectID:   "",
			expectName: "",
			expectErr:  require.Error,
		},
		{
			name:       "only name map with owner id and lookup",
			owner:      id,
			ins:        inMock.NewCache(nil, nti),
			rc:         lookup,
			expectID:   id,
			expectName: name,
			expectErr:  require.NoError,
		},
		{
			name:       "only id map with owner name",
			owner:      name,
			ins:        inMock.NewCache(itn, nil),
			rc:         lookup,
			expectID:   id,
			expectName: name,
			expectErr:  require.NoError,
		},
		{
			name:       "only name map with owner name",
			owner:      name,
			ins:        inMock.NewCache(nil, nti),
			rc:         noLookup,
			expectID:   id,
			expectName: name,
			expectErr:  require.NoError,
		},
		{
			name:       "only id map with owner name",
			owner:      name,
			ins:        inMock.NewCache(itn, nil),
			rc:         noLookup,
			expectID:   "",
			expectName: "",
			expectErr:  require.Error,
		},
		{
			name:       "only id map with owner name and lookup",
			owner:      name,
			ins:        inMock.NewCache(itn, nil),
			rc:         lookup,
			expectID:   id,
			expectName: name,
			expectErr:  require.NoError,
		},
		{
			name:       "both maps with owner id",
			owner:      id,
			ins:        inMock.NewCache(itn, nti),
			rc:         noLookup,
			expectID:   id,
			expectName: name,
			expectErr:  require.NoError,
		},
		{
			name:       "both maps with owner name",
			owner:      name,
			ins:        inMock.NewCache(itn, nti),
			rc:         noLookup,
			expectID:   id,
			expectName: name,
			expectErr:  require.NoError,
		},
		{
			name:  "non-matching maps with owner id",
			owner: id,
			ins: inMock.NewCache(
				map[string]string{"foo": "bar"},
				map[string]string{"fnords": "smarf"}),
			rc:         noLookup,
			expectID:   "",
			expectName: "",
			expectErr:  require.Error,
		},
		{
			name:  "non-matching with owner name",
			owner: name,
			ins: inMock.NewCache(
				map[string]string{"foo": "bar"},
				map[string]string{"fnords": "smarf"}),
			rc:         noLookup,
			expectID:   "",
			expectName: "",
			expectErr:  require.Error,
		},
		{
			name:  "non-matching maps with owner id and lookup",
			owner: id,
			ins: inMock.NewCache(
				map[string]string{"foo": "bar"},
				map[string]string{"fnords": "smarf"}),
			rc:         lookup,
			expectID:   id,
			expectName: name,
			expectErr:  require.NoError,
		},
		{
			name:  "non-matching with owner name and lookup",
			owner: name,
			ins: inMock.NewCache(
				map[string]string{"foo": "bar"},
				map[string]string{"fnords": "smarf"}),
			rc:         lookup,
			expectID:   id,
			expectName: name,
			expectErr:  require.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrl := &Controller{ownerLookup: test.rc}

			rID, rName, err := ctrl.PopulateOwnerIDAndNamesFrom(ctx, test.owner, test.ins)
			test.expectErr(t, err, clues.ToCore(err))
			assert.Equal(t, test.expectID, rID, "id")
			assert.Equal(t, test.expectName, rName, "name")
		})
	}
}

func (suite *ControllerUnitSuite) TestController_Wait() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		ctrl = &Controller{
			wg:     &sync.WaitGroup{},
			region: &trace.Region{},
		}
		metrics = support.CollectionMetrics{
			Objects:   2,
			Successes: 3,
			Bytes:     4,
		}
		status = support.CreateStatus(ctx, support.Backup, 1, metrics, "details")
	)

	ctrl.wg.Add(1)
	ctrl.UpdateStatus(status)

	result := ctrl.Wait()
	require.NotNil(t, result)
	assert.Nil(t, ctrl.region, "region")
	assert.Empty(t, ctrl.status, "status")
	assert.Equal(t, 1, result.Folders)
	assert.Equal(t, 2, result.Objects)
	assert.Equal(t, 3, result.Successes)
	assert.Equal(t, int64(4), result.Bytes)
}

// ---------------------------------------------------------------------------
// Integration tests
// ---------------------------------------------------------------------------

type ControllerIntegrationSuite struct {
	tester.Suite
	ctrl          *Controller
	user          string
	secondaryUser string
}

func TestControllerIntegrationSuite(t *testing.T) {
	suite.Run(t, &ControllerIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *ControllerIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.ctrl = loadController(ctx, t, resource.Users)
	suite.user = tester.M365UserID(t)
	suite.secondaryUser = tester.SecondaryM365UserID(t)

	tester.LogTimeOfTest(t)
}

func (suite *ControllerIntegrationSuite) TestRestoreFailsBadService() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		restoreCfg = testdata.DefaultRestoreConfig("")
		sel        = selectors.Selector{
			Service: selectors.ServiceUnknown,
		}
	)

	deets, err := suite.ctrl.ConsumeRestoreCollections(
		ctx,
		version.Backup,
		sel,
		restoreCfg,
		control.Options{
			RestorePermissions: true,
			ToggleFeatures:     control.Toggles{},
		},
		nil,
		fault.New(true))
	assert.Error(t, err, clues.ToCore(err))
	assert.NotNil(t, deets)

	status := suite.ctrl.Wait()
	assert.Equal(t, 0, status.Objects)
	assert.Equal(t, 0, status.Folders)
	assert.Equal(t, 0, status.Successes)
}

func (suite *ControllerIntegrationSuite) TestEmptyCollections() {
	restoreCfg := testdata.DefaultRestoreConfig("")
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
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			deets, err := suite.ctrl.ConsumeRestoreCollections(
				ctx,
				version.Backup,
				test.sel,
				restoreCfg,
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{},
				},
				test.col,
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			assert.NotNil(t, deets)

			stats := suite.ctrl.Wait()
			assert.Zero(t, stats.Objects)
			assert.Zero(t, stats.Folders)
			assert.Zero(t, stats.Successes)
		})
	}
}

//-------------------------------------------------------------
// Exchange Functions
//-------------------------------------------------------------

func runRestore(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	config ConfigInfo,
	backupVersion int,
	collections []data.RestoreCollection,
	numRestoreItems int,
) {
	t.Logf(
		"Restoring collections to %s for resourceOwners(s) %v\n",
		config.RestoreCfg.Location,
		config.ResourceOwners)

	start := time.Now()

	restoreCtrl := loadController(ctx, t, config.Resource)
	restoreSel := getSelectorWith(t, config.Service, config.ResourceOwners, true)
	deets, err := restoreCtrl.ConsumeRestoreCollections(
		ctx,
		backupVersion,
		restoreSel,
		config.RestoreCfg,
		config.Opts,
		collections,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, deets)

	status := restoreCtrl.Wait()
	runTime := time.Since(start)

	assert.Equal(t, numRestoreItems, status.Objects, "restored status.Objects")
	assert.Equal(t, numRestoreItems, status.Successes, "restored status.Successes")
	assert.Len(
		t,
		// Don't check folders as those are now added to details.
		deets.Items(),
		numRestoreItems,
		"details entries contains same item count as total successful items restored")

	t.Logf("Restore complete in %v\n", runTime)
}

func runBackupAndCompare(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	config ConfigInfo,
	expectedData map[string]map[string][]byte,
	totalItems int,
	totalKopiaItems int,
	inputCollections []ColInfo,
) {
	t.Helper()

	// Run a backup and compare its output with what we put in.
	cats := make(map[path.CategoryType]struct{}, len(inputCollections))
	for _, c := range inputCollections {
		cats[c.Category] = struct{}{}
	}

	var (
		expectedDests = make([]destAndCats, 0, len(config.ResourceOwners))
		idToName      = map[string]string{}
		nameToID      = map[string]string{}
	)

	for _, ro := range config.ResourceOwners {
		expectedDests = append(expectedDests, destAndCats{
			resourceOwner: ro,
			dest:          config.RestoreCfg.Location,
			cats:          cats,
		})

		idToName[ro] = ro
		nameToID[ro] = ro
	}

	backupCtrl := loadController(ctx, t, config.Resource)
	backupCtrl.IDNameLookup = inMock.NewCache(idToName, nameToID)

	backupSel := backupSelectorForExpected(t, config.Service, expectedDests)
	t.Logf("Selective backup of %s\n", backupSel)

	start := time.Now()
	dcs, excludes, canUsePreviousBackup, err := backupCtrl.ProduceBackupCollections(
		ctx,
		backupSel,
		backupSel,
		nil,
		version.NoBackup,
		config.Opts,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, canUsePreviousBackup, "can use previous backup")
	// No excludes yet because this isn't an incremental backup.
	assert.True(t, excludes.Empty())

	t.Logf("Backup enumeration complete in %v\n", time.Since(start))

	// Pull the data prior to waiting for the status as otherwise it will
	// deadlock.
	skipped := checkCollections(
		t,
		ctx,
		totalKopiaItems,
		expectedData,
		dcs,
		config)

	status := backupCtrl.Wait()

	assert.Equalf(t, totalItems+skipped, status.Objects,
		"backup status.Objects; wanted %d items + %d skipped", totalItems, skipped)
	assert.Equalf(t, totalItems+skipped, status.Successes,
		"backup status.Successes; wanted %d items + %d skipped", totalItems, skipped)
}

func runRestoreBackupTest(
	t *testing.T,
	test restoreBackupInfo,
	tenant string,
	resourceOwners []string,
	opts control.Options,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	config := ConfigInfo{
		Opts:           opts,
		Resource:       test.resourceCat,
		Service:        test.service,
		Tenant:         tenant,
		ResourceOwners: resourceOwners,
		RestoreCfg:     testdata.DefaultRestoreConfig(""),
	}

	totalItems, totalKopiaItems, collections, expectedData, err := GetCollectionsAndExpected(
		config,
		test.collections,
		version.Backup)

	require.NoError(t, err)

	runRestore(
		t,
		ctx,
		config,
		version.Backup,
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

// runRestoreTest restores with data using the test's backup version
func runRestoreTestWithVersion(
	t *testing.T,
	test restoreBackupInfoMultiVersion,
	tenant string,
	resourceOwners []string,
	opts control.Options,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	config := ConfigInfo{
		Opts:           opts,
		Resource:       test.resource,
		Service:        test.service,
		Tenant:         tenant,
		ResourceOwners: resourceOwners,
		RestoreCfg:     testdata.DefaultRestoreConfig(""),
	}

	totalItems, _, collections, _, err := GetCollectionsAndExpected(
		config,
		test.collectionsPrevious,
		test.backupVersion)
	require.NoError(t, err)

	runRestore(
		t,
		ctx,
		config,
		test.backupVersion,
		collections,
		totalItems)
}

// runRestoreBackupTestVersions restores with data from an older
// version of the backup and check the restored data against the
// something that would be in the form of a newer backup.
func runRestoreBackupTestVersions(
	t *testing.T,
	test restoreBackupInfoMultiVersion,
	tenant string,
	resourceOwners []string,
	opts control.Options,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	config := ConfigInfo{
		Opts:           opts,
		Resource:       test.resource,
		Service:        test.service,
		Tenant:         tenant,
		ResourceOwners: resourceOwners,
		RestoreCfg:     testdata.DefaultRestoreConfig(""),
	}

	totalItems, _, collections, _, err := GetCollectionsAndExpected(
		config,
		test.collectionsPrevious,
		test.backupVersion)
	require.NoError(t, err)

	runRestore(
		t,
		ctx,
		config,
		test.backupVersion,
		collections,
		totalItems)

	// Get expected output for new version.
	totalItems, totalKopiaItems, _, expectedData, err := GetCollectionsAndExpected(
		config,
		test.collectionsLatest,
		version.Backup)
	require.NoError(t, err)

	runBackupAndCompare(
		t,
		ctx,
		config,
		expectedData,
		totalItems,
		totalKopiaItems,
		test.collectionsLatest)
}

func (suite *ControllerIntegrationSuite) TestRestoreAndBackup() {
	bodyText := "This email has some text. However, all the text is on the same line."
	subjectText := "Test message for restore"

	table := []restoreBackupInfo{
		{
			name:        "EmailsWithAttachments",
			service:     path.ExchangeService,
			resourceCat: resource.Users,
			collections: []ColInfo{
				{
					PathElements: []string{"Inbox"},
					Category:     path.EmailCategory,
					Items: []ItemInfo{
						{
							name: "someencodeditemID",
							data: exchMock.MessageWithDirectAttachment(
								subjectText + "-1",
							),
							lookupKey: subjectText + "-1",
						},
						{
							name: "someencodeditemID2",
							data: exchMock.MessageWithTwoAttachments(
								subjectText + "-2",
							),
							lookupKey: subjectText + "-2",
						},
					},
				},
			},
		},
		{
			name:        "MultipleEmailsMultipleFolders",
			service:     path.ExchangeService,
			resourceCat: resource.Users,
			collections: []ColInfo{
				{
					PathElements: []string{"Inbox"},
					Category:     path.EmailCategory,
					Items: []ItemInfo{
						{
							name: "someencodeditemID",
							data: exchMock.MessageWithBodyBytes(
								subjectText+"-1",
								bodyText+" 1.",
								bodyText+" 1.",
							),
							lookupKey: subjectText + "-1",
						},
					},
				},
				{
					PathElements: []string{"Work"},
					Category:     path.EmailCategory,
					Items: []ItemInfo{
						{
							name: "someencodeditemID2",
							data: exchMock.MessageWithBodyBytes(
								subjectText+"-2",
								bodyText+" 2.",
								bodyText+" 2.",
							),
							lookupKey: subjectText + "-2",
						},
						{
							name: "someencodeditemID3",
							data: exchMock.MessageWithBodyBytes(
								subjectText+"-3",
								bodyText+" 3.",
								bodyText+" 3.",
							),
							lookupKey: subjectText + "-3",
						},
					},
				},
				{
					PathElements: []string{"Work", "Inbox"},
					Category:     path.EmailCategory,
					Items: []ItemInfo{
						{
							name: "someencodeditemID4",
							data: exchMock.MessageWithBodyBytes(
								subjectText+"-4",
								bodyText+" 4.",
								bodyText+" 4.",
							),
							lookupKey: subjectText + "-4",
						},
					},
				},
				{
					PathElements: []string{"Work", "Inbox", "Work"},
					Category:     path.EmailCategory,
					Items: []ItemInfo{
						{
							name: "someencodeditemID5",
							data: exchMock.MessageWithBodyBytes(
								subjectText+"-5",
								bodyText+" 5.",
								bodyText+" 5.",
							),
							lookupKey: subjectText + "-5",
						},
					},
				},
			},
		},
		{
			name:        "MultipleContactsSingleFolder",
			service:     path.ExchangeService,
			resourceCat: resource.Users,
			collections: []ColInfo{
				{
					PathElements: []string{"Contacts"},
					Category:     path.ContactsCategory,
					Items: []ItemInfo{
						{
							name:      "someencodeditemID",
							data:      exchMock.ContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
						{
							name:      "someencodeditemID2",
							data:      exchMock.ContactBytes("Irgot"),
							lookupKey: "Irgot",
						},
						{
							name:      "someencodeditemID3",
							data:      exchMock.ContactBytes("Jannes"),
							lookupKey: "Jannes",
						},
					},
				},
			},
		},
		{
			name:        "MultipleContactsMultipleFolders",
			service:     path.ExchangeService,
			resourceCat: resource.Users,
			collections: []ColInfo{
				{
					PathElements: []string{"Work"},
					Category:     path.ContactsCategory,
					Items: []ItemInfo{
						{
							name:      "someencodeditemID",
							data:      exchMock.ContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
						{
							name:      "someencodeditemID2",
							data:      exchMock.ContactBytes("Irgot"),
							lookupKey: "Irgot",
						},
						{
							name:      "someencodeditemID3",
							data:      exchMock.ContactBytes("Jannes"),
							lookupKey: "Jannes",
						},
					},
				},
				{
					PathElements: []string{"Personal"},
					Category:     path.ContactsCategory,
					Items: []ItemInfo{
						{
							name:      "someencodeditemID4",
							data:      exchMock.ContactBytes("Argon"),
							lookupKey: "Argon",
						},
						{
							name:      "someencodeditemID5",
							data:      exchMock.ContactBytes("Bernard"),
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
		// 					data:      exchMock.EventWithSubjectBytes("Ghimley"),
		// 					lookupKey: "Ghimley",
		// 				},
		// 				{
		// 					name:      "someencodeditemID2",
		// 					data:      exchMock.EventWithSubjectBytes("Irgot"),
		// 					lookupKey: "Irgot",
		// 				},
		// 				{
		// 					name:      "someencodeditemID3",
		// 					data:      exchMock.EventWithSubjectBytes("Jannes"),
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
		// 					data:      exchMock.EventWithSubjectBytes("Ghimley"),
		// 					lookupKey: "Ghimley",
		// 				},
		// 				{
		// 					name:      "someencodeditemID2",
		// 					data:      exchMock.EventWithSubjectBytes("Irgot"),
		// 					lookupKey: "Irgot",
		// 				},
		// 				{
		// 					name:      "someencodeditemID3",
		// 					data:      exchMock.EventWithSubjectBytes("Jannes"),
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
		// 					data:      exchMock.EventWithSubjectBytes("Argon"),
		// 					lookupKey: "Argon",
		// 				},
		// 				{
		// 					name:      "someencodeditemID5",
		// 					data:      exchMock.EventWithSubjectBytes("Bernard"),
		// 					lookupKey: "Bernard",
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			runRestoreBackupTest(
				suite.T(),
				test,
				suite.ctrl.tenant,
				[]string{suite.user},
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{},
				})
		})
	}
}

func (suite *ControllerIntegrationSuite) TestMultiFolderBackupDifferentNames() {
	table := []restoreBackupInfo{
		{
			name:        "Contacts",
			service:     path.ExchangeService,
			resourceCat: resource.Users,
			collections: []ColInfo{
				{
					PathElements: []string{"Work"},
					Category:     path.ContactsCategory,
					Items: []ItemInfo{
						{
							name:      "someencodeditemID",
							data:      exchMock.ContactBytes("Ghimley"),
							lookupKey: "Ghimley",
						},
					},
				},
				{
					PathElements: []string{"Personal"},
					Category:     path.ContactsCategory,
					Items: []ItemInfo{
						{
							name:      "someencodeditemID2",
							data:      exchMock.ContactBytes("Irgot"),
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
		// 					data:      exchMock.EventWithSubjectBytes("Ghimley"),
		// 					lookupKey: "Ghimley",
		// 				},
		// 			},
		// 		},
		// 		{
		// 			PathElements: []string{"Personal"},
		// 			Category:     path.EventsCategory,
		// 			Items: []ItemInfo{
		// 				{
		// 					name:      "someencodeditemID2",
		// 					data:      exchMock.EventWithSubjectBytes("Irgot"),
		// 					lookupKey: "Irgot",
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			restoreSel := getSelectorWith(t, test.service, []string{suite.user}, true)
			expectedDests := make([]destAndCats, 0, len(test.collections))
			allItems := 0
			allExpectedData := map[string]map[string][]byte{}

			for i, collection := range test.collections {
				// Get a restoreCfg per collection so they're independent.
				restoreCfg := testdata.DefaultRestoreConfig("")
				expectedDests = append(expectedDests, destAndCats{
					resourceOwner: suite.user,
					dest:          restoreCfg.Location,
					cats: map[path.CategoryType]struct{}{
						collection.Category: {},
					},
				})

				totalItems, _, collections, expectedData, err := collectionsForInfo(
					test.service,
					suite.ctrl.tenant,
					suite.user,
					restoreCfg,
					[]ColInfo{collection},
					version.Backup,
				)
				require.NoError(t, err)

				allItems += totalItems

				for k, v := range expectedData {
					allExpectedData[k] = v
				}

				t.Logf(
					"Restoring %v/%v collections to %s\n",
					i+1,
					len(test.collections),
					restoreCfg.Location,
				)

				restoreCtrl := loadController(ctx, t, test.resourceCat)
				deets, err := restoreCtrl.ConsumeRestoreCollections(
					ctx,
					version.Backup,
					restoreSel,
					restoreCfg,
					control.Options{
						RestorePermissions: true,
						ToggleFeatures:     control.Toggles{},
					},
					collections,
					fault.New(true))
				require.NoError(t, err, clues.ToCore(err))
				require.NotNil(t, deets)

				status := restoreCtrl.Wait()
				// Always just 1 because it's just 1 collection.
				assert.Equal(t, totalItems, status.Objects, "status.Objects")
				assert.Equal(t, totalItems, status.Successes, "status.Successes")
				assert.Len(
					t,
					deets.Items(),
					totalItems,
					"details entries contains same item count as total successful items restored")

				t.Log("Restore complete")
			}

			// Run a backup and compare its output with what we put in.

			backupCtrl := loadController(ctx, t, test.resourceCat)
			backupSel := backupSelectorForExpected(t, test.service, expectedDests)
			t.Log("Selective backup of", backupSel)

			dcs, excludes, canUsePreviousBackup, err := backupCtrl.ProduceBackupCollections(
				ctx,
				backupSel,
				backupSel,
				nil,
				version.NoBackup,
				control.Options{
					RestorePermissions: true,
					ToggleFeatures:     control.Toggles{},
				},
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			assert.True(t, canUsePreviousBackup, "can use previous backup")
			// No excludes yet because this isn't an incremental backup.
			assert.True(t, excludes.Empty())

			t.Log("Backup enumeration complete")

			ci := ConfigInfo{
				Opts: control.Options{RestorePermissions: true},
				// Alright to be empty, needed for OneDrive.
				RestoreCfg: control.RestoreConfig{},
			}

			// Pull the data prior to waiting for the status as otherwise it will
			// deadlock.
			skipped := checkCollections(t, ctx, allItems, allExpectedData, dcs, ci)

			status := backupCtrl.Wait()
			assert.Equal(t, allItems+skipped, status.Objects, "status.Objects")
			assert.Equal(t, allItems+skipped, status.Successes, "status.Successes")
		})
	}
}

// TODO: this should only be run during smoke tests, not part of the standard CI.
// That's why it's set aside instead of being included in the other test set.
func (suite *ControllerIntegrationSuite) TestRestoreAndBackup_largeMailAttachment() {
	subjectText := "Test message for restore with large attachment"

	test := restoreBackupInfo{
		name:        "EmailsWithLargeAttachments",
		service:     path.ExchangeService,
		resourceCat: resource.Users,
		collections: []ColInfo{
			{
				PathElements: []string{"Inbox"},
				Category:     path.EmailCategory,
				Items: []ItemInfo{
					{
						name:      "35mbAttachment",
						data:      exchMock.MessageWithSizedAttachment(subjectText, 35),
						lookupKey: subjectText,
					},
				},
			},
		},
	}

	runRestoreBackupTest(
		suite.T(),
		test,
		suite.ctrl.tenant,
		[]string{suite.user},
		control.Options{
			RestorePermissions: true,
			ToggleFeatures:     control.Toggles{},
		},
	)
}

func (suite *ControllerIntegrationSuite) TestBackup_CreatesPrefixCollections() {
	table := []struct {
		name         string
		resourceCat  resource.Category
		selectorFunc func(t *testing.T) selectors.Selector
		service      path.ServiceType
		categories   []string
	}{
		{
			name:        "Exchange",
			resourceCat: resource.Users,
			selectorFunc: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup([]string{suite.user})
				sel.Include(
					sel.ContactFolders([]string{selectors.NoneTgt}),
					sel.EventCalendars([]string{selectors.NoneTgt}),
					sel.MailFolders([]string{selectors.NoneTgt}),
				)

				return sel.Selector
			},
			service: path.ExchangeService,
			categories: []string{
				path.EmailCategory.String(),
				path.ContactsCategory.String(),
				path.EventsCategory.String(),
			},
		},
		{
			name:        "OneDrive",
			resourceCat: resource.Users,
			selectorFunc: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{suite.user})
				sel.Include(sel.Folders([]string{selectors.NoneTgt}))

				return sel.Selector
			},
			service: path.OneDriveService,
			categories: []string{
				path.FilesCategory.String(),
			},
		},
		{
			name:        "SharePoint",
			resourceCat: resource.Sites,
			selectorFunc: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{tester.M365SiteID(t)})
				sel.Include(
					sel.LibraryFolders([]string{selectors.NoneTgt}),
					// not yet in use
					//  sel.Pages([]string{selectors.NoneTgt}),
					//  sel.Lists([]string{selectors.NoneTgt}),
				)

				return sel.Selector
			},
			service: path.SharePointService,
			categories: []string{
				path.LibrariesCategory.String(),
				// not yet in use
				// path.PagesCategory.String(),
				// path.ListsCategory.String(),
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				backupCtrl = loadController(ctx, t, test.resourceCat)
				backupSel  = test.selectorFunc(t)
				errs       = fault.New(true)
				start      = time.Now()
			)

			id, name, err := backupCtrl.PopulateOwnerIDAndNamesFrom(ctx, backupSel.DiscreteOwner, nil)
			require.NoError(t, err, clues.ToCore(err))

			backupSel.SetDiscreteOwnerIDName(id, name)

			dcs, excludes, canUsePreviousBackup, err := backupCtrl.ProduceBackupCollections(
				ctx,
				inMock.NewProvider(id, name),
				backupSel,
				nil,
				version.NoBackup,
				control.Options{
					RestorePermissions: false,
					ToggleFeatures:     control.Toggles{},
				},
				fault.New(true))
			require.NoError(t, err)
			assert.True(t, canUsePreviousBackup, "can use previous backup")
			// No excludes yet because this isn't an incremental backup.
			assert.True(t, excludes.Empty())

			t.Logf("Backup enumeration complete in %v\n", time.Since(start))

			// Use a map to find duplicates.
			foundCategories := []string{}
			for _, col := range dcs {
				// TODO(ashmrtn): We should be able to remove the below if we change how
				// status updates are done. Ideally we shouldn't have to fetch items in
				// these collections to avoid deadlocking.
				var found int

				// Need to iterate through this before the continue below else we'll
				// hang checking the status.
				for range col.Items(ctx, errs) {
					found++
				}

				// Ignore metadata collections.
				fullPath := col.FullPath()
				if fullPath.Service() != test.service {
					continue
				}

				assert.Empty(t, fullPath.Folders(), "non-prefix collection")
				assert.NotEqual(t, col.State(), data.NewState, "prefix collection marked as new")
				foundCategories = append(foundCategories, fullPath.Category().String())

				assert.Zero(t, found, "non-empty collection")
			}

			assert.ElementsMatch(t, test.categories, foundCategories)

			backupCtrl.Wait()

			assert.NoError(t, errs.Failure())
		})
	}
}
