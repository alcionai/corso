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

	"github.com/alcionai/corso/src/internal/common/idname"
	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/mock"
	"github.com/alcionai/corso/src/internal/m365/resource"
	exchMock "github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/corso/src/internal/m365/stub"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
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
		lookup = &resourceGetter{
			enum:   resource.Users,
			getter: &mock.IDNameGetter{ID: id, Name: name},
		}
		noLookup = &resourceGetter{enum: resource.Users, getter: &mock.IDNameGetter{}}
	)

	table := []struct {
		name              string
		protectedResource string
		ins               inMock.Cache
		rc                *resourceGetter
		expectID          string
		expectName        string
		expectErr         require.ErrorAssertionFunc
		expectNil         require.ValueAssertionFunc
	}{
		{
			name:              "nil ins",
			protectedResource: id,
			rc:                lookup,
			expectID:          id,
			expectName:        name,
			expectErr:         require.NoError,
			expectNil:         require.NotNil,
		},
		{
			name:              "nil ins no lookup",
			protectedResource: id,
			rc:                noLookup,
			expectID:          "",
			expectName:        "",
			expectErr:         require.Error,
			expectNil:         require.Nil,
		},
		{
			name:              "only id map with owner id",
			protectedResource: id,
			ins:               inMock.NewCache(itn, nil),
			rc:                noLookup,
			expectID:          id,
			expectName:        name,
			expectErr:         require.NoError,
			expectNil:         require.NotNil,
		},
		{
			name:              "only name map with owner id",
			protectedResource: id,
			ins:               inMock.NewCache(nil, nti),
			rc:                noLookup,
			expectID:          "",
			expectName:        "",
			expectErr:         require.Error,
			expectNil:         require.Nil,
		},
		{
			name:              "only name map with owner id and lookup",
			protectedResource: id,
			ins:               inMock.NewCache(nil, nti),
			rc:                lookup,
			expectID:          id,
			expectName:        name,
			expectErr:         require.NoError,
			expectNil:         require.NotNil,
		},
		{
			name:              "only id map with owner name",
			protectedResource: name,
			ins:               inMock.NewCache(itn, nil),
			rc:                lookup,
			expectID:          id,
			expectName:        name,
			expectErr:         require.NoError,
			expectNil:         require.NotNil,
		},
		{
			name:              "only name map with owner name",
			protectedResource: name,
			ins:               inMock.NewCache(nil, nti),
			rc:                noLookup,
			expectID:          id,
			expectName:        name,
			expectErr:         require.NoError,
			expectNil:         require.NotNil,
		},
		{
			name:              "only id map with owner name",
			protectedResource: name,
			ins:               inMock.NewCache(itn, nil),
			rc:                noLookup,
			expectID:          "",
			expectName:        "",
			expectErr:         require.Error,
			expectNil:         require.Nil,
		},
		{
			name:              "only id map with owner name and lookup",
			protectedResource: name,
			ins:               inMock.NewCache(itn, nil),
			rc:                lookup,
			expectID:          id,
			expectName:        name,
			expectErr:         require.NoError,
			expectNil:         require.NotNil,
		},
		{
			name:              "both maps with owner id",
			protectedResource: id,
			ins:               inMock.NewCache(itn, nti),
			rc:                noLookup,
			expectID:          id,
			expectName:        name,
			expectErr:         require.NoError,
			expectNil:         require.NotNil,
		},
		{
			name:              "both maps with owner name",
			protectedResource: name,
			ins:               inMock.NewCache(itn, nti),
			rc:                noLookup,
			expectID:          id,
			expectName:        name,
			expectErr:         require.NoError,
			expectNil:         require.NotNil,
		},
		{
			name:              "non-matching maps with owner id",
			protectedResource: id,
			ins: inMock.NewCache(
				map[string]string{"foo": "bar"},
				map[string]string{"fnords": "smarf"}),
			rc:         noLookup,
			expectID:   "",
			expectName: "",
			expectErr:  require.Error,
			expectNil:  require.Nil,
		},
		{
			name:              "non-matching with owner name",
			protectedResource: name,
			ins: inMock.NewCache(
				map[string]string{"foo": "bar"},
				map[string]string{"fnords": "smarf"}),
			rc:         noLookup,
			expectID:   "",
			expectName: "",
			expectErr:  require.Error,
			expectNil:  require.Nil,
		},
		{
			name:              "non-matching maps with owner id and lookup",
			protectedResource: id,
			ins: inMock.NewCache(
				map[string]string{"foo": "bar"},
				map[string]string{"fnords": "smarf"}),
			rc:         lookup,
			expectID:   id,
			expectName: name,
			expectErr:  require.NoError,
			expectNil:  require.NotNil,
		},
		{
			name:              "non-matching with owner name and lookup",
			protectedResource: name,
			ins: inMock.NewCache(
				map[string]string{"foo": "bar"},
				map[string]string{"fnords": "smarf"}),
			rc:         lookup,
			expectID:   id,
			expectName: name,
			expectErr:  require.NoError,
			expectNil:  require.NotNil,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrl := &Controller{resourceHandler: test.rc}

			resource, err := ctrl.PopulateProtectedResourceIDAndName(ctx, test.protectedResource, test.ins)
			test.expectErr(t, err, clues.ToCore(err))
			test.expectNil(t, resource)

			if err != nil {
				return
			}

			assert.Equal(t, test.expectID, resource.ID(), "id")
			assert.Equal(t, test.expectName, resource.Name(), "name")
		})
	}
}

func (suite *ControllerUnitSuite) TestPopulateOwnerIDAndNamesFrom_nilCheck() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	ctrl := &Controller{resourceHandler: nil}

	_, err := ctrl.PopulateProtectedResourceIDAndName(ctx, "", nil)
	require.ErrorIs(t, err, ErrNoResourceLookup, clues.ToCore(err))
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

func (suite *ControllerUnitSuite) TestController_CacheItemInfo() {
	var (
		odid   = "od-id"
		odname = "od-name"
		spid   = "sp-id"
		spname = "sp-name"
		spsid  = "sp-sid"
		spurl  = "sp-url"
		gpid   = "gp-id"
		gpname = "gp-name"
		// intentionally declared outside the test loop
		ctrl = &Controller{
			wg:                 &sync.WaitGroup{},
			region:             &trace.Region{},
			backupDriveIDNames: idname.NewCache(nil),
			backupSiteIDWebURL: idname.NewCache(nil),
		}
	)

	table := []struct {
		name             string
		service          path.ServiceType
		cat              path.CategoryType
		dii              details.ItemInfo
		expectDriveID    string
		expectDriveName  string
		expectSiteID     string
		expectSiteWebURL string
	}{
		{
			name: "exchange",
			dii: details.ItemInfo{
				Exchange: &details.ExchangeInfo{},
			},
			expectDriveID:   "",
			expectDriveName: "",
		},
		{
			name: "folder",
			dii: details.ItemInfo{
				Folder: &details.FolderInfo{},
			},
			expectDriveID:   "",
			expectDriveName: "",
		},
		{
			name: "onedrive",
			dii: details.ItemInfo{
				OneDrive: &details.OneDriveInfo{
					DriveID:   odid,
					DriveName: odname,
				},
			},
			expectDriveID:   odid,
			expectDriveName: odname,
		},
		{
			name: "sharepoint",
			dii: details.ItemInfo{
				SharePoint: &details.SharePointInfo{
					DriveID:   spid,
					DriveName: spname,
					SiteID:    spsid,
					WebURL:    spurl,
				},
			},
			expectDriveID:    spid,
			expectDriveName:  spname,
			expectSiteID:     spsid,
			expectSiteWebURL: spurl,
		},
		{
			name: "groups",
			dii: details.ItemInfo{
				Groups: &details.GroupsInfo{
					DriveID:   gpid,
					DriveName: gpname,
					SiteID:    spsid,
					WebURL:    spurl,
				},
			},
			expectDriveID:    gpid,
			expectDriveName:  gpname,
			expectSiteID:     spsid,
			expectSiteWebURL: spurl,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctrl.CacheItemInfo(test.dii)

			name, _ := ctrl.backupDriveIDNames.NameOf(test.expectDriveID)
			assert.Equal(t, test.expectDriveName, name)

			id, _ := ctrl.backupDriveIDNames.IDOf(test.expectDriveName)
			assert.Equal(t, test.expectDriveID, id)

			url, _ := ctrl.backupSiteIDWebURL.NameOf(test.expectSiteID)
			assert.Equal(t, test.expectSiteWebURL, url)

			sid, _ := ctrl.backupSiteIDWebURL.IDOf(test.expectSiteWebURL)
			assert.Equal(t, test.expectSiteID, sid)
		})
	}
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
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *ControllerIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.ctrl = newController(ctx, t, path.ExchangeService)
	suite.user = tconfig.M365UserID(t)
	suite.secondaryUser = tconfig.SecondaryM365UserID(t)

	tester.LogTimeOfTest(t)
}

func (suite *ControllerIntegrationSuite) TestEmptyCollections() {
	restoreCfg := testdata.DefaultRestoreConfig("")
	restoreCfg.IncludePermissions = true

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
		{
			name: "GroupsNil",
			col:  nil,
			sel: selectors.Selector{
				Service: selectors.ServiceGroups,
			},
		},
		{
			name: "GroupsEmpty",
			col:  []data.RestoreCollection{},
			sel: selectors.Selector{
				Service: selectors.ServiceGroups,
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			rcc := inject.RestoreConsumerConfig{
				BackupVersion:     version.Backup,
				Options:           control.DefaultOptions(),
				ProtectedResource: test.sel,
				RestoreConfig:     restoreCfg,
				Selector:          test.sel,
			}

			handler, err := suite.ctrl.NewServiceHandler(test.sel.PathService())
			require.NoError(t, err, clues.ToCore(err))

			deets, _, err := handler.ConsumeRestoreCollections(
				ctx,
				rcc,
				test.col,
				fault.New(true),
				count.New())
			require.Error(t, err, clues.ToCore(err))
			assert.Nil(t, deets)
		})
	}
}

//-------------------------------------------------------------
// Exchange Functions
//-------------------------------------------------------------

func runRestore(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	sci stub.ConfigInfo,
	backupVersion int,
	collections []data.RestoreCollection,
	numRestoreItems int,
) {
	t.Logf(
		"Restoring collections to %s for resourceOwners(s) %v\n",
		sci.RestoreCfg.Location,
		sci.ResourceOwners)

	start := time.Now()

	restoreCtrl := newController(ctx, t, path.ExchangeService)
	restoreSel := getSelectorWith(t, sci.Service, sci.ResourceOwners, true)

	rcc := inject.RestoreConsumerConfig{
		BackupVersion:     backupVersion,
		Options:           control.DefaultOptions(),
		ProtectedResource: restoreSel,
		RestoreConfig:     sci.RestoreCfg,
		Selector:          restoreSel,
	}

	handler, err := restoreCtrl.NewServiceHandler(sci.Service)
	require.NoError(t, err, clues.ToCore(err))

	deets, status, err := handler.ConsumeRestoreCollections(
		ctx,
		rcc,
		collections,
		fault.New(true),
		count.New())
	require.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, deets)

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
	sci stub.ConfigInfo,
	expectedData map[string]map[string][]byte,
	totalItems int,
	totalKopiaItems int,
	inputCollections []stub.ColInfo,
) {
	t.Helper()

	// Run a backup and compare its output with what we put in.
	cats := make(map[path.CategoryType]struct{}, len(inputCollections))
	for _, c := range inputCollections {
		cats[c.Category] = struct{}{}
	}

	var (
		expectedDests = make([]destAndCats, 0, len(sci.ResourceOwners))
		idToName      = map[string]string{}
		nameToID      = map[string]string{}
	)

	for _, ro := range sci.ResourceOwners {
		expectedDests = append(expectedDests, destAndCats{
			resourceOwner: ro,
			dest:          sci.RestoreCfg.Location,
			cats:          cats,
		})

		idToName[ro] = ro
		nameToID[ro] = ro
	}

	backupCtrl := newController(ctx, t, path.ExchangeService)
	backupCtrl.IDNameLookup = inMock.NewCache(idToName, nameToID)

	backupSel := backupSelectorForExpected(t, sci.Service, expectedDests)
	t.Logf("Selective backup of %s\n", backupSel)

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           sci.Opts,
		ProtectedResource: backupSel,
		Selector:          backupSel,
	}

	start := time.Now()
	dcs, excludes, canUsePreviousBackup, err := backupCtrl.ProduceBackupCollections(
		ctx,
		bpc,
		count.New(),
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
		sci)

	status := backupCtrl.Wait()

	assert.Equalf(t, totalItems+skipped, status.Objects,
		"backup status.Objects; wanted %d items + %d skipped", totalItems, skipped)
	assert.Equalf(t, totalItems+skipped, status.Successes,
		"backup status.Successes; wanted %d items + %d skipped", totalItems, skipped)
}

func runRestoreBackupTest(
	t *testing.T,
	test restoreBackupInfo,
	cfg stub.ConfigInfo,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	totalItems, totalKopiaItems, collections, expectedData, err := stub.GetCollectionsAndExpected(
		cfg,
		test.collections,
		version.Backup)

	require.NoError(t, err)

	runRestore(
		t,
		ctx,
		cfg,
		version.Backup,
		collections,
		totalItems)

	runBackupAndCompare(
		t,
		ctx,
		cfg,
		expectedData,
		totalItems,
		totalKopiaItems,
		test.collections)
}

// runRestoreTest restores with data using the test's backup version
func runRestoreTestWithVersion(
	t *testing.T,
	test restoreBackupInfoMultiVersion,
	cfg stub.ConfigInfo,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	totalItems, _, collections, _, err := stub.GetCollectionsAndExpected(
		cfg,
		test.collectionsPrevious,
		test.backupVersion)
	require.NoError(t, err)

	runRestore(
		t,
		ctx,
		cfg,
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
	cfg stub.ConfigInfo,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	totalItems, _, collections, _, err := stub.GetCollectionsAndExpected(
		cfg,
		test.collectionsPrevious,
		test.backupVersion)
	require.NoError(t, err)

	runRestore(
		t,
		ctx,
		cfg,
		test.backupVersion,
		collections,
		totalItems)

	// Get expected output for new version.
	totalItems, totalKopiaItems, _, expectedData, err := stub.GetCollectionsAndExpected(
		cfg,
		test.collectionsLatest,
		version.Backup)
	require.NoError(t, err)

	runBackupAndCompare(
		t,
		ctx,
		cfg,
		expectedData,
		totalItems,
		totalKopiaItems,
		test.collectionsLatest)
}

func (suite *ControllerIntegrationSuite) TestRestoreAndBackup_core() {
	bodyText := "This email has some text. However, all the text is on the same line."
	subjectText := "Test message for restore"

	table := []restoreBackupInfo{
		{
			name:    "EmailsWithAttachments",
			service: path.ExchangeService,
			collections: []stub.ColInfo{
				{
					PathElements: []string{api.MailInbox},
					Category:     path.EmailCategory,
					Items: []stub.ItemInfo{
						{
							Name: "someencodeditemID",
							Data: exchMock.MessageWithDirectAttachment(
								subjectText + "-1"),
							LookupKey: subjectText + "-1",
						},
						{
							Name: "someencodeditemID2",
							Data: exchMock.MessageWithTwoAttachments(
								subjectText + "-2"),
							LookupKey: subjectText + "-2",
						},
					},
				},
			},
		},
		{
			name:    "MultipleEmailsMultipleFolders",
			service: path.ExchangeService,
			collections: []stub.ColInfo{
				{
					PathElements: []string{api.MailInbox},
					Category:     path.EmailCategory,
					Items: []stub.ItemInfo{
						{
							Name: "someencodeditemID",
							Data: exchMock.MessageWithBodyBytes(
								subjectText+"-1",
								bodyText+" 1.",
								bodyText+" 1."),
							LookupKey: subjectText + "-1",
						},
					},
				},
				{
					PathElements: []string{"Work"},
					Category:     path.EmailCategory,
					Items: []stub.ItemInfo{
						{
							Name: "someencodeditemID2",
							Data: exchMock.MessageWithBodyBytes(
								subjectText+"-2",
								bodyText+" 2.",
								bodyText+" 2."),
							LookupKey: subjectText + "-2",
						},
						{
							Name: "someencodeditemID3",
							Data: exchMock.MessageWithBodyBytes(
								subjectText+"-3",
								bodyText+" 3.",
								bodyText+" 3."),
							LookupKey: subjectText + "-3",
						},
					},
				},
				{
					PathElements: []string{"Work", api.MailInbox},
					Category:     path.EmailCategory,
					Items: []stub.ItemInfo{
						{
							Name: "someencodeditemID4",
							Data: exchMock.MessageWithBodyBytes(
								subjectText+"-4",
								bodyText+" 4.",
								bodyText+" 4."),
							LookupKey: subjectText + "-4",
						},
					},
				},
				{
					PathElements: []string{"Work", api.MailInbox, "Work"},
					Category:     path.EmailCategory,
					Items: []stub.ItemInfo{
						{
							Name: "someencodeditemID5",
							Data: exchMock.MessageWithBodyBytes(
								subjectText+"-5",
								bodyText+" 5.",
								bodyText+" 5."),
							LookupKey: subjectText + "-5",
						},
					},
				},
			},
		},
		{
			name:    "MultipleContactsInRestoreFolder",
			service: path.ExchangeService,
			collections: []stub.ColInfo{
				{
					PathElements: []string{"Contacts"},
					Category:     path.ContactsCategory,
					Items: []stub.ItemInfo{
						{
							Name:      "someencodeditemID",
							Data:      exchMock.ContactBytes("Ghimley"),
							LookupKey: "Ghimley",
						},
						{
							Name:      "someencodeditemID2",
							Data:      exchMock.ContactBytes("Irgot"),
							LookupKey: "Irgot",
						},
						{
							Name:      "someencodeditemID3",
							Data:      exchMock.ContactBytes("Jannes"),
							LookupKey: "Jannes",
						},
					},
				},
			},
		},
		// TODO(ashmrtn): Re-enable when we can restore contacts to nested folders.
		//{
		//	name:    "MultipleContactsSingleFolder",
		//	service: path.ExchangeService,
		//	collections: []stub.ColInfo{
		//		{
		//			PathElements: []string{"Contacts"},
		//			Category:     path.ContactsCategory,
		//			Items: []stub.ItemInfo{
		//				{
		//					Name:      "someencodeditemID",
		//					Data:      exchMock.ContactBytes("Ghimley"),
		//					LookupKey: "Ghimley",
		//				},
		//				{
		//					Name:      "someencodeditemID2",
		//					Data:      exchMock.ContactBytes("Irgot"),
		//					LookupKey: "Irgot",
		//				},
		//				{
		//					Name:      "someencodeditemID3",
		//					Data:      exchMock.ContactBytes("Jannes"),
		//					LookupKey: "Jannes",
		//				},
		//			},
		//		},
		//	},
		//},
		//{
		//	name:    "MultipleContactsMultipleFolders",
		//	service: path.ExchangeService,
		//	collections: []stub.ColInfo{
		//		{
		//			PathElements: []string{"Work"},
		//			Category:     path.ContactsCategory,
		//			Items: []stub.ItemInfo{
		//				{
		//					Name:      "someencodeditemID",
		//					Data:      exchMock.ContactBytes("Ghimley"),
		//					LookupKey: "Ghimley",
		//				},
		//				{
		//					Name:      "someencodeditemID2",
		//					Data:      exchMock.ContactBytes("Irgot"),
		//					LookupKey: "Irgot",
		//				},
		//				{
		//					Name:      "someencodeditemID3",
		//					Data:      exchMock.ContactBytes("Jannes"),
		//					LookupKey: "Jannes",
		//				},
		//			},
		//		},
		//		{
		//			PathElements: []string{"Personal"},
		//			Category:     path.ContactsCategory,
		//			Items: []stub.ItemInfo{
		//				{
		//					Name:      "someencodeditemID4",
		//					Data:      exchMock.ContactBytes("Argon"),
		//					LookupKey: "Argon",
		//				},
		//				{
		//					Name:      "someencodeditemID5",
		//					Data:      exchMock.ContactBytes("Bernard"),
		//					LookupKey: "Bernard",
		//				},
		//			},
		//		},
		//	},
		//},
		// {
		// 	name:    "MultipleEventsSingleCalendar",
		// 	service: path.ExchangeService,
		// 	collections: []stub.ColInfo{
		// 		{
		// 			PathElements: []string{"Work"},
		// 			Category:     path.EventsCategory,
		// 			Items: []stub.ItemInfo{
		// 				{
		// 					Name:      "someencodeditemID",
		// 					Data:      exchMock.EventWithSubjectBytes("Ghimley"),
		// 					LookupKey: "Ghimley",
		// 				},
		// 				{
		// 					Name:      "someencodeditemID2",
		// 					Data:      exchMock.EventWithSubjectBytes("Irgot"),
		// 					LookupKey: "Irgot",
		// 				},
		// 				{
		// 					Name:      "someencodeditemID3",
		// 					Data:      exchMock.EventWithSubjectBytes("Jannes"),
		// 					LookupKey: "Jannes",
		// 				},
		// 			},
		// 		},
		// 	},
		// },
		// {
		// 	name:    "MultipleEventsMultipleCalendars",
		// 	service: path.ExchangeService,
		// 	collections: []stub.ColInfo{
		// 		{
		// 			PathElements: []string{"Work"},
		// 			Category:     path.EventsCategory,
		// 			Items: []stub.ItemInfo{
		// 				{
		// 					Name:      "someencodeditemID",
		// 					Data:      exchMock.EventWithSubjectBytes("Ghimley"),
		// 					LookupKey: "Ghimley",
		// 				},
		// 				{
		// 					Name:      "someencodeditemID2",
		// 					Data:      exchMock.EventWithSubjectBytes("Irgot"),
		// 					LookupKey: "Irgot",
		// 				},
		// 				{
		// 					Name:      "someencodeditemID3",
		// 					Data:      exchMock.EventWithSubjectBytes("Jannes"),
		// 					LookupKey: "Jannes",
		// 				},
		// 			},
		// 		},
		// 		{
		// 			PathElements: []string{"Personal"},
		// 			Category:     path.EventsCategory,
		// 			Items: []stub.ItemInfo{
		// 				{
		// 					Name:      "someencodeditemID4",
		// 					Data:      exchMock.EventWithSubjectBytes("Argon"),
		// 					LookupKey: "Argon",
		// 				},
		// 				{
		// 					Name:      "someencodeditemID5",
		// 					Data:      exchMock.EventWithSubjectBytes("Bernard"),
		// 					LookupKey: "Bernard",
		// 				},
		// 			},
		// 		},
		// 	},
		// },
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			cfg := stub.ConfigInfo{
				Tenant:         suite.ctrl.tenant,
				ResourceOwners: []string{suite.user},
				Service:        test.service,
				Opts:           control.DefaultOptions(),
				RestoreCfg:     control.DefaultRestoreConfig(dttm.SafeForTesting),
			}

			runRestoreBackupTest(suite.T(), test, cfg)
		})
	}
}

func (suite *ControllerIntegrationSuite) TestMultiFolderBackupDifferentNames() {
	table := []restoreBackupInfo{
		// TODO(ashmrtn): Re-enable when we can restore contacts to nested folders.
		//{
		//	name:    "Contacts",
		//	service: path.ExchangeService,
		//	collections: []stub.ColInfo{
		//		{
		//			PathElements: []string{"Work"},
		//			Category:     path.ContactsCategory,
		//			Items: []stub.ItemInfo{
		//				{
		//					Name:      "someencodeditemID",
		//					Data:      exchMock.ContactBytes("Ghimley"),
		//					LookupKey: "Ghimley",
		//				},
		//			},
		//		},
		//		{
		//			PathElements: []string{"Personal"},
		//			Category:     path.ContactsCategory,
		//			Items: []stub.ItemInfo{
		//				{
		//					Name:      "someencodeditemID2",
		//					Data:      exchMock.ContactBytes("Irgot"),
		//					LookupKey: "Irgot",
		//				},
		//			},
		//		},
		//	},
		//},
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
				restoreCfg.IncludePermissions = true

				expectedDests = append(expectedDests, destAndCats{
					resourceOwner: suite.user,
					dest:          restoreCfg.Location,
					cats: map[path.CategoryType]struct{}{
						collection.Category: {},
					},
				})

				totalItems, _, collections, expectedData, err := stub.CollectionsForInfo(
					test.service,
					suite.ctrl.tenant,
					suite.user,
					restoreCfg,
					[]stub.ColInfo{collection},
					version.Backup)
				require.NoError(t, err)

				allItems += totalItems

				for k, v := range expectedData {
					allExpectedData[k] = v
				}

				t.Logf(
					"Restoring %v/%v collections to %s\n",
					i+1,
					len(test.collections),
					restoreCfg.Location)

				restoreCtrl := newController(ctx, t, path.ExchangeService)

				rcc := inject.RestoreConsumerConfig{
					BackupVersion:     version.Backup,
					Options:           control.DefaultOptions(),
					ProtectedResource: restoreSel,
					RestoreConfig:     restoreCfg,
					Selector:          restoreSel,
				}

				handler, err := restoreCtrl.NewServiceHandler(test.service)
				require.NoError(t, err, clues.ToCore(err))

				deets, status, err := handler.ConsumeRestoreCollections(
					ctx,
					rcc,
					collections,
					fault.New(true),
					count.New())
				require.NoError(t, err, clues.ToCore(err))
				require.NotNil(t, deets)

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

			backupCtrl := newController(ctx, t, path.ExchangeService)
			backupSel := backupSelectorForExpected(t, test.service, expectedDests)
			t.Log("Selective backup of", backupSel)

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: version.NoBackup,
				Options:           control.DefaultOptions(),
				ProtectedResource: backupSel,
				Selector:          backupSel,
			}

			dcs, excludes, canUsePreviousBackup, err := backupCtrl.ProduceBackupCollections(
				ctx,
				bpc,
				count.New(),
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			assert.True(t, canUsePreviousBackup, "can use previous backup")
			// No excludes yet because this isn't an incremental backup.
			assert.True(t, excludes.Empty())

			t.Log("Backup enumeration complete")

			restoreCfg := control.DefaultRestoreConfig(dttm.SafeForTesting)
			restoreCfg.IncludePermissions = true

			ci := stub.ConfigInfo{
				Opts: control.DefaultOptions(),
				// Alright to be empty, needed for OneDrive.
				RestoreCfg: restoreCfg,
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
		name:    "EmailsWithLargeAttachments",
		service: path.ExchangeService,
		collections: []stub.ColInfo{
			{
				PathElements: []string{api.MailInbox},
				Category:     path.EmailCategory,
				Items: []stub.ItemInfo{
					{
						Name:      "35mbAttachment",
						Data:      exchMock.MessageWithSizedAttachment(subjectText, 35),
						LookupKey: subjectText,
					},
				},
			},
		},
	}

	restoreCfg := control.DefaultRestoreConfig(dttm.SafeForTesting)
	restoreCfg.IncludePermissions = true

	cfg := stub.ConfigInfo{
		Tenant:         suite.ctrl.tenant,
		ResourceOwners: []string{suite.user},
		Service:        test.service,
		Opts:           control.DefaultOptions(),
		RestoreCfg:     restoreCfg,
	}

	runRestoreBackupTest(suite.T(), test, cfg)
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
					sel.MailFolders([]string{selectors.NoneTgt}))

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
				sel := selectors.NewSharePointBackup([]string{tconfig.M365SiteID(t)})
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
		{
			name:        "Groups",
			resourceCat: resource.Sites,
			selectorFunc: func(t *testing.T) selectors.Selector {
				sel := selectors.NewGroupsBackup([]string{tconfig.M365TeamID(t)})
				sel.Include(
					sel.LibraryFolders([]string{selectors.NoneTgt}),
					// not yet in use
					//  sel.Pages([]string{selectors.NoneTgt}),
					//  sel.Lists([]string{selectors.NoneTgt}),
				)

				return sel.Selector
			},
			service: path.GroupsService,
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
				backupCtrl = newController(ctx, t, test.service)
				backupSel  = test.selectorFunc(t)
				errs       = fault.New(true)
				start      = time.Now()
			)

			resource, err := backupCtrl.PopulateProtectedResourceIDAndName(ctx, backupSel.DiscreteOwner, nil)
			require.NoError(t, err, clues.ToCore(err))

			backupSel.SetDiscreteOwnerIDName(resource.ID(), resource.Name())

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: version.NoBackup,
				Options:           control.DefaultOptions(),
				ProtectedResource: resource,
				Selector:          backupSel,
			}

			dcs, excludes, canUsePreviousBackup, err := backupCtrl.ProduceBackupCollections(
				ctx,
				bpc,
				count.New(),
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
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

type DisconnectedUnitSuite struct {
	tester.Suite
}

func TestDisconnectedUnitSuite(t *testing.T) {
	s := &DisconnectedUnitSuite{
		Suite: tester.NewUnitSuite(t),
	}

	suite.Run(t, s)
}

func statusTestTask(
	t *testing.T,
	ctrl *Controller,
	objects, success, folder int,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	status := support.CreateStatus(
		ctx,
		support.Restore, folder,
		support.CollectionMetrics{
			Objects:   objects,
			Successes: success,
			Bytes:     0,
		},
		"statusTestTask")
	ctrl.UpdateStatus(status)
}

func (suite *DisconnectedUnitSuite) TestController_Status() {
	t := suite.T()
	ctrl := Controller{wg: &sync.WaitGroup{}}

	// Two tasks
	ctrl.incrementAwaitingMessages()
	ctrl.incrementAwaitingMessages()

	// Each helper task processes 4 objects, 1 success, 3 errors, 1 folders
	go statusTestTask(t, &ctrl, 4, 1, 1)
	go statusTestTask(t, &ctrl, 4, 1, 1)

	stats := ctrl.Wait()

	assert.NotEmpty(t, ctrl.PrintableStatus())
	// Expect 8 objects
	assert.Equal(t, 8, stats.Objects)
	// Expect 2 success
	assert.Equal(t, 2, stats.Successes)
	// Expect 2 folders
	assert.Equal(t, 2, stats.Folders)
}

func (suite *DisconnectedUnitSuite) TestVerifyBackupInputs_allServices() {
	sites := []string{"abc.site.foo", "bar.site.baz"}
	groups := []string{"123", "456"}

	tests := []struct {
		name       string
		excludes   func(t *testing.T) selectors.Selector
		filters    func(t *testing.T) selectors.Selector
		includes   func(t *testing.T) selectors.Selector
		cachedIDs  []string
		checkError assert.ErrorAssertionFunc
	}{
		{
			name:       "Valid User",
			checkError: assert.NoError,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"elliotReid@someHospital.org", "foo@SomeCompany.org"})
				sel.Exclude(selTD.OneDriveBackupFolderScope(sel))
				sel.DiscreteOwner = "elliotReid@someHospital.org"
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"elliotReid@someHospital.org", "foo@SomeCompany.org"})
				sel.Filter(selTD.OneDriveBackupFolderScope(sel))
				sel.DiscreteOwner = "elliotReid@someHospital.org"
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"elliotReid@someHospital.org", "foo@SomeCompany.org"})
				sel.Include(selTD.OneDriveBackupFolderScope(sel))
				sel.DiscreteOwner = "elliotReid@someHospital.org"
				return sel.Selector
			},
		},
		{
			name:       "Invalid User",
			checkError: assert.NoError,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"foo@SomeCompany.org"})
				sel.Exclude(selTD.OneDriveBackupFolderScope(sel))
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"foo@SomeCompany.org"})
				sel.Filter(selTD.OneDriveBackupFolderScope(sel))
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup([]string{"foo@SomeCompany.org"})
				sel.Include(selTD.OneDriveBackupFolderScope(sel))
				return sel.Selector
			},
		},
		{
			name:       "valid sites",
			checkError: assert.NoError,
			cachedIDs:  sites,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"abc.site.foo", "bar.site.baz"})
				sel.DiscreteOwner = "abc.site.foo"
				sel.Exclude(sel.AllData())
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"abc.site.foo", "bar.site.baz"})
				sel.DiscreteOwner = "abc.site.foo"
				sel.Filter(sel.AllData())
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"abc.site.foo", "bar.site.baz"})
				sel.DiscreteOwner = "abc.site.foo"
				sel.Include(sel.AllData())
				return sel.Selector
			},
		},
		{
			name:       "invalid sites",
			checkError: assert.Error,
			cachedIDs:  sites,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"fnords.smarfs.brawnhilda"})
				sel.Exclude(sel.AllData())
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"fnords.smarfs.brawnhilda"})
				sel.Filter(sel.AllData())
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup([]string{"fnords.smarfs.brawnhilda"})
				sel.Include(sel.AllData())
				return sel.Selector
			},
		},

		{
			name:       "valid groups",
			checkError: assert.NoError,
			cachedIDs:  groups,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewGroupsBackup([]string{"123", "456"})
				sel.DiscreteOwner = "123"
				sel.Exclude(sel.AllData())
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewGroupsBackup([]string{"123", "456"})
				sel.DiscreteOwner = "123"
				sel.Filter(sel.AllData())
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewGroupsBackup([]string{"123", "456"})
				sel.DiscreteOwner = "123"
				sel.Include(sel.AllData())
				return sel.Selector
			},
		},
		{
			name:       "invalid groups",
			checkError: assert.Error,
			cachedIDs:  groups,
			excludes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewGroupsBackup([]string{"789"})
				sel.Exclude(sel.AllData())
				return sel.Selector
			},
			filters: func(t *testing.T) selectors.Selector {
				sel := selectors.NewGroupsBackup([]string{"789"})
				sel.Filter(sel.AllData())
				return sel.Selector
			},
			includes: func(t *testing.T) selectors.Selector {
				sel := selectors.NewGroupsBackup([]string{"789"})
				sel.Include(sel.AllData())
				return sel.Selector
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			err := verifyBackupInputs(test.excludes(t), test.cachedIDs)
			test.checkError(t, err, clues.ToCore(err))
			err = verifyBackupInputs(test.filters(t), test.cachedIDs)
			test.checkError(t, err, clues.ToCore(err))
			err = verifyBackupInputs(test.includes(t), test.cachedIDs)
			test.checkError(t, err, clues.ToCore(err))
		})
	}
}
