package test_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/dttm"
	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/onedrive"
	"github.com/alcionai/corso/src/internal/m365/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	deeTD "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/control"
	ctrlTD "github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type OneDriveBackupIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestOneDriveBackupIntgSuite(t *testing.T) {
	suite.Run(t, &OneDriveBackupIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *OneDriveBackupIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *OneDriveBackupIntgSuite) TestBackup_Run_oneDrive() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		tenID  = tconfig.M365TenantID(t)
		mb     = evmock.NewBus()
		userID = tconfig.SecondaryM365UserID(t)
		osel   = selectors.NewOneDriveBackup([]string{userID})
		ws     = deeTD.DriveIDFromRepoRef
		svc    = path.OneDriveService
		opts   = control.Defaults()
	)

	osel.Include(selTD.OneDriveBackupFolderScope(osel))

	bo, bod := prepNewTestBackupOp(t, ctx, mb, osel.Selector, opts, version.Backup)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)

	bID := bo.Results.BackupID

	_, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bID,
		tenID,
		bod.sel.ID(),
		svc,
		ws,
		bod.kms,
		bod.sss)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bID,
		ws,
		bod.kms,
		bod.sss,
		expectDeets,
		false)
}

func (suite *OneDriveBackupIntgSuite) TestBackup_Run_incrementalOneDrive() {
	sel := selectors.NewOneDriveRestore([]string{suite.its.userID})

	ic := func(cs []string) selectors.Selector {
		sel.Include(sel.Folders(cs, selectors.PrefixMatch()))
		return sel.Selector
	}

	gtdi := func(
		t *testing.T,
		ctx context.Context,
	) string {
		d, err := suite.its.ac.Users().GetDefaultDrive(ctx, suite.its.userID)
		if err != nil {
			err = graph.Wrap(ctx, err, "retrieving default user drive").
				With("user", suite.its.userID)
		}

		require.NoError(t, err, clues.ToCore(err))

		id := ptr.Val(d.GetId())
		require.NotEmpty(t, id, "drive ID")

		return id
	}

	grh := func(ac api.Client) onedrive.RestoreHandler {
		return onedrive.NewRestoreHandler(ac)
	}

	runDriveIncrementalTest(
		suite,
		suite.its.userID,
		suite.its.userID,
		resource.Users,
		path.OneDriveService,
		path.FilesCategory,
		ic,
		gtdi,
		grh,
		false)
}

func runDriveIncrementalTest(
	suite tester.Suite,
	owner, permissionsUser string,
	rc resource.Category,
	service path.ServiceType,
	category path.CategoryType,
	includeContainers func([]string) selectors.Selector,
	getTestDriveID func(*testing.T, context.Context) string,
	getRestoreHandler func(api.Client) onedrive.RestoreHandler,
	skipPermissionsTests bool,
) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acct = tconfig.NewM365Account(t)
		opts = control.Defaults()
		mb   = evmock.NewBus()
		ws   = deeTD.DriveIDFromRepoRef

		// `now` has to be formatted with SimpleDateTimeTesting as
		// some drives cannot have `:` in file/folder names
		now = dttm.FormatNow(dttm.SafeForTesting)

		categories = map[path.CategoryType][]string{
			category: {graph.DeltaURLsFileName, graph.PreviousPathFileName},
		}
		container1      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 1, now)
		container2      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 2, now)
		container3      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 3, now)
		containerRename = "renamed_folder"

		genDests = []string{container1, container2}

		// container3 does not exist yet. It will get created later on
		// during the tests.
		containers = []string{container1, container2, container3}
	)

	sel := includeContainers(containers)

	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	ctrl, sel := ControllerWithSelector(t, ctx, acct, rc, sel, nil, nil)
	ac := ctrl.AC.Drives()
	rh := getRestoreHandler(ctrl.AC)

	roidn := inMock.NewProvider(sel.ID(), sel.Name())

	var (
		atid    = creds.AzureTenantID
		driveID = getTestDriveID(t, ctx)
		fileDBF = func(id, timeStamp, subject, body string) []byte {
			return []byte(id + subject)
		}
		makeLocRef = func(flds ...string) string {
			elems := append([]string{driveID, "root:"}, flds...)
			return path.Builder{}.Append(elems...).String()
		}
	)

	rrPfx, err := path.ServicePrefix(atid, roidn.ID(), service, category)
	require.NoError(t, err, clues.ToCore(err))

	// strip the category from the prefix; we primarily want the tenant and resource owner.
	expectDeets := deeTD.NewInDeets(rrPfx.ToBuilder().Dir().String())

	// Populate initial test data.
	// Generate 2 new folders with two items each. Only the first two
	// folders will be part of the initial backup and
	// incrementals. The third folder will be introduced partway
	// through the changes. This should be enough to cover most delta
	// actions.
	for _, destName := range genDests {
		deets := generateContainerOfItems(
			t,
			ctx,
			ctrl,
			service,
			category,
			sel,
			atid, roidn.ID(), driveID, destName,
			2,
			// Use an old backup version so we don't need metadata files.
			0,
			fileDBF)

		for _, ent := range deets.Entries {
			if ent.Folder != nil {
				continue
			}

			expectDeets.AddItem(driveID, makeLocRef(destName), ent.ItemRef)
		}
	}

	containerIDs := map[string]string{}

	// verify test data was populated, and track it for comparisons
	for _, destName := range genDests {
		// Use path-based indexing to get the folder's ID. This is sourced from the
		// onedrive package `getFolder` function.
		itemURL := fmt.Sprintf("https://graph.microsoft.com/v1.0/drives/%s/root:/%s", driveID, destName)
		resp, err := drives.
			NewItemItemsDriveItemItemRequestBuilder(itemURL, ctrl.AC.Stable.Adapter()).
			Get(ctx, nil)
		require.NoError(t, err, "getting drive folder ID", "folder name", destName, clues.ToCore(err))

		containerIDs[destName] = ptr.Val(resp.GetId())
	}

	bo, bod := prepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup)
	defer bod.close(t, ctx)

	sel = bod.sel

	// run the initial backup
	runAndCheckBackup(t, ctx, &bo, mb, false)

	// precheck to ensure the expectedDeets are correct.
	// if we fail here, the expectedDeets were populated incorrectly.
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bo.Results.BackupID,
		ws,
		bod.kms,
		bod.sss,
		expectDeets,
		true)

	var (
		newFile     models.DriveItemable
		newFileName = "new_file.txt"
		newFileID   string

		permissionIDMappings = map[string]string{}
		writePerm            = metadata.Permission{
			ID:       "perm-id",
			Roles:    []string{"write"},
			EntityID: permissionsUser,
		}
	)

	// Although established as a table, these tests are not isolated from each other.
	// Assume that every test's side effects cascade to all following test cases.
	// The changes are split across the table so that we can monitor the deltas
	// in isolation, rather than debugging one change from the rest of a series.
	table := []struct {
		name string
		// performs the incremental update required for the test.
		//revive:disable-next-line:context-as-argument
		updateFiles         func(t *testing.T, ctx context.Context)
		itemsRead           int
		itemsWritten        int
		nonMetaItemsWritten int
	}{
		{
			name:         "clean incremental, no changes",
			updateFiles:  func(t *testing.T, ctx context.Context) {},
			itemsRead:    0,
			itemsWritten: 0,
		},
		{
			name: "create a new file",
			updateFiles: func(t *testing.T, ctx context.Context) {
				targetContainer := containerIDs[container1]
				driveItem := models.NewDriveItem()
				driveItem.SetName(&newFileName)
				driveItem.SetFile(models.NewFile())
				newFile, err = ac.PostItemInContainer(
					ctx,
					driveID,
					targetContainer,
					driveItem,
					control.Copy)
				require.NoErrorf(t, err, "creating new file %v", clues.ToCore(err))

				newFileID = ptr.Val(newFile.GetId())

				expectDeets.AddItem(driveID, makeLocRef(container1), newFileID)
			},
			itemsRead:           1, // .data file for newitem
			itemsWritten:        3, // .data and .meta for newitem, .dirmeta for parent
			nonMetaItemsWritten: 1, // .data file for newitem
		},
		{
			name: "add permission to new file",
			updateFiles: func(t *testing.T, ctx context.Context) {
				err = onedrive.UpdatePermissions(
					ctx,
					rh,
					driveID,
					ptr.Val(newFile.GetId()),
					[]metadata.Permission{writePerm},
					[]metadata.Permission{},
					permissionIDMappings)
				require.NoErrorf(t, err, "adding permission to file %v", clues.ToCore(err))
				// no expectedDeets: metadata isn't tracked
			},
			itemsRead:           1, // .data file for newitem
			itemsWritten:        2, // .meta for newitem, .dirmeta for parent (.data is not written as it is not updated)
			nonMetaItemsWritten: 1, // the file for which permission was updated
		},
		{
			name: "remove permission from new file",
			updateFiles: func(t *testing.T, ctx context.Context) {
				err = onedrive.UpdatePermissions(
					ctx,
					rh,
					driveID,
					*newFile.GetId(),
					[]metadata.Permission{},
					[]metadata.Permission{writePerm},
					permissionIDMappings)
				require.NoErrorf(t, err, "removing permission from file %v", clues.ToCore(err))
				// no expectedDeets: metadata isn't tracked
			},
			itemsRead:           1, // .data file for newitem
			itemsWritten:        2, // .meta for newitem, .dirmeta for parent (.data is not written as it is not updated)
			nonMetaItemsWritten: 1, //.data file for newitem
		},
		{
			name: "add permission to container",
			updateFiles: func(t *testing.T, ctx context.Context) {
				targetContainer := containerIDs[container1]
				err = onedrive.UpdatePermissions(
					ctx,
					rh,
					driveID,
					targetContainer,
					[]metadata.Permission{writePerm},
					[]metadata.Permission{},
					permissionIDMappings)
				require.NoErrorf(t, err, "adding permission to container %v", clues.ToCore(err))
				// no expectedDeets: metadata isn't tracked
			},
			itemsRead:           0,
			itemsWritten:        1, // .dirmeta for collection
			nonMetaItemsWritten: 0, // no files updated as update on container
		},
		{
			name: "remove permission from container",
			updateFiles: func(t *testing.T, ctx context.Context) {
				targetContainer := containerIDs[container1]
				err = onedrive.UpdatePermissions(
					ctx,
					rh,
					driveID,
					targetContainer,
					[]metadata.Permission{},
					[]metadata.Permission{writePerm},
					permissionIDMappings)
				require.NoErrorf(t, err, "removing permission from container %v", clues.ToCore(err))
				// no expectedDeets: metadata isn't tracked
			},
			itemsRead:           0,
			itemsWritten:        1, // .dirmeta for collection
			nonMetaItemsWritten: 0, // no files updated
		},
		{
			name: "update contents of a file",
			updateFiles: func(t *testing.T, ctx context.Context) {
				err := ac.PutItemContent(
					ctx,
					driveID,
					ptr.Val(newFile.GetId()),
					[]byte("new content"))
				require.NoErrorf(t, err, "updating file contents: %v", clues.ToCore(err))
				// no expectedDeets: neither file id nor location changed
			},
			itemsRead:           1, // .data file for newitem
			itemsWritten:        3, // .data and .meta for newitem, .dirmeta for parent
			nonMetaItemsWritten: 1, // .data  file for newitem
		},
		{
			name: "rename a file",
			updateFiles: func(t *testing.T, ctx context.Context) {
				container := containerIDs[container1]

				driveItem := models.NewDriveItem()
				name := "renamed_new_file.txt"
				driveItem.SetName(&name)
				parentRef := models.NewItemReference()
				parentRef.SetId(&container)
				driveItem.SetParentReference(parentRef)

				err := ac.PatchItem(
					ctx,
					driveID,
					ptr.Val(newFile.GetId()),
					driveItem)
				require.NoError(t, err, "renaming file %v", clues.ToCore(err))
			},
			itemsRead:           1, // .data file for newitem
			itemsWritten:        3, // .data and .meta for newitem, .dirmeta for parent
			nonMetaItemsWritten: 1, // .data file for newitem
			// no expectedDeets: neither file id nor location changed
		},
		{
			name: "move a file between folders",
			updateFiles: func(t *testing.T, ctx context.Context) {
				dest := containerIDs[container2]

				driveItem := models.NewDriveItem()
				driveItem.SetName(&newFileName)
				parentRef := models.NewItemReference()
				parentRef.SetId(&dest)
				driveItem.SetParentReference(parentRef)

				err := ac.PatchItem(
					ctx,
					driveID,
					ptr.Val(newFile.GetId()),
					driveItem)
				require.NoErrorf(t, err, "moving file between folders %v", clues.ToCore(err))

				expectDeets.MoveItem(
					driveID,
					makeLocRef(container1),
					makeLocRef(container2),
					ptr.Val(newFile.GetId()))
			},
			itemsRead:           1, // .data file for newitem
			itemsWritten:        3, // .data and .meta for newitem, .dirmeta for parent
			nonMetaItemsWritten: 1, // .data file for new item
		},
		{
			name: "delete file",
			updateFiles: func(t *testing.T, ctx context.Context) {
				err := ac.DeleteItem(
					ctx,
					driveID,
					ptr.Val(newFile.GetId()))
				require.NoErrorf(t, err, "deleting file %v", clues.ToCore(err))

				expectDeets.RemoveItem(driveID, makeLocRef(container2), ptr.Val(newFile.GetId()))
			},
			itemsRead:           0,
			itemsWritten:        0,
			nonMetaItemsWritten: 0,
		},
		{
			name: "move a folder to a subfolder",
			updateFiles: func(t *testing.T, ctx context.Context) {
				parent := containerIDs[container1]
				child := containerIDs[container2]

				driveItem := models.NewDriveItem()
				driveItem.SetName(&container2)
				parentRef := models.NewItemReference()
				parentRef.SetId(&parent)
				driveItem.SetParentReference(parentRef)

				err := ac.PatchItem(
					ctx,
					driveID,
					child,
					driveItem)
				require.NoError(t, err, "moving folder", clues.ToCore(err))

				expectDeets.MoveLocation(
					driveID,
					makeLocRef(container2),
					makeLocRef(container1))
			},
			itemsRead:           0,
			itemsWritten:        7, // 2*2(data and meta of 2 files) + 3 (dirmeta of two moved folders and target)
			nonMetaItemsWritten: 0,
		},
		{
			name: "rename a folder",
			updateFiles: func(t *testing.T, ctx context.Context) {
				parent := containerIDs[container1]
				child := containerIDs[container2]

				driveItem := models.NewDriveItem()
				driveItem.SetName(&containerRename)
				parentRef := models.NewItemReference()
				parentRef.SetId(&parent)
				driveItem.SetParentReference(parentRef)

				err := ac.PatchItem(
					ctx,
					driveID,
					child,
					driveItem)
				require.NoError(t, err, "renaming folder", clues.ToCore(err))

				containerIDs[containerRename] = containerIDs[container2]

				expectDeets.RenameLocation(
					driveID,
					makeLocRef(container1, container2),
					makeLocRef(container1, containerRename))
			},
			itemsRead:           0,
			itemsWritten:        7, // 2*2(data and meta of 2 files) + 3 (dirmeta of two moved folders and target)
			nonMetaItemsWritten: 0,
		},
		{
			name: "delete a folder",
			updateFiles: func(t *testing.T, ctx context.Context) {
				container := containerIDs[containerRename]
				err := ac.DeleteItem(
					ctx,
					driveID,
					container)
				require.NoError(t, err, "deleting folder", clues.ToCore(err))

				expectDeets.RemoveLocation(driveID, makeLocRef(container1, containerRename))
			},
			itemsRead:           0,
			itemsWritten:        0,
			nonMetaItemsWritten: 0,
		},
		{
			name: "add a new folder",
			updateFiles: func(t *testing.T, ctx context.Context) {
				generateContainerOfItems(
					t,
					ctx,
					ctrl,
					service,
					category,
					sel,
					atid, roidn.ID(), driveID, container3,
					2,
					0,
					fileDBF)

				// Validate creation
				itemURL := fmt.Sprintf(
					"https://graph.microsoft.com/v1.0/drives/%s/root:/%s",
					driveID,
					container3)
				resp, err := drives.NewItemItemsDriveItemItemRequestBuilder(itemURL, ctrl.AC.Stable.Adapter()).
					Get(ctx, nil)
				require.NoError(t, err, "getting drive folder ID", "folder name", container3, clues.ToCore(err))

				containerIDs[container3] = ptr.Val(resp.GetId())

				expectDeets.AddLocation(driveID, container3)
			},
			itemsRead:           2, // 2 .data for 2 files
			itemsWritten:        6, // read items + 2 directory meta
			nonMetaItemsWritten: 2, // 2 .data for 2 files
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			cleanCtrl, err := m365.NewController(ctx, acct, rc, sel.PathService(), control.Defaults())
			require.NoError(t, err, clues.ToCore(err))

			bod.ctrl = cleanCtrl

			var (
				t     = suite.T()
				incMB = evmock.NewBus()
				incBO = newTestBackupOp(
					t,
					ctx,
					bod,
					incMB,
					opts)
			)

			ctx, flush := tester.WithContext(t, ctx)
			defer flush()

			suite.Run("PreTestSetup", func() {
				t := suite.T()

				ctx, flush := tester.WithContext(t, ctx)
				defer flush()

				test.updateFiles(t, ctx)
			})

			err = incBO.Run(ctx)
			require.NoError(t, err, clues.ToCore(err))

			bupID := incBO.Results.BackupID

			checkBackupIsInManifests(
				t,
				ctx,
				bod.kw,
				bod.sw,
				&incBO,
				sel,
				roidn.ID(),
				maps.Keys(categories)...)
			checkMetadataFilesExist(
				t,
				ctx,
				bupID,
				bod.kw,
				bod.kms,
				atid,
				roidn.ID(),
				service,
				categories)
			deeTD.CheckBackupDetails(
				t,
				ctx,
				bupID,
				ws,
				bod.kms,
				bod.sss,
				expectDeets,
				true)

			// do some additional checks to ensure the incremental dealt with fewer items.
			// +2 on read/writes to account for metadata: 1 delta and 1 path.
			var (
				expectWrites        = test.itemsWritten + 2
				expectNonMetaWrites = test.nonMetaItemsWritten
				expectReads         = test.itemsRead + 2
				assertReadWrite     = assert.Equal
			)

			// Sharepoint can produce a superset of permissions by nature of
			// its drive type.  Since this counter comparison is a bit hacky
			// to begin with, it's easiest to assert a <= comparison instead
			// of fine tuning each test case.
			if service == path.SharePointService {
				assertReadWrite = assert.LessOrEqual
			}

			assertReadWrite(t, expectWrites, incBO.Results.ItemsWritten, "incremental items written")
			assertReadWrite(t, expectNonMetaWrites, incBO.Results.NonMetaItemsWritten, "incremental non-meta items written")
			assertReadWrite(t, expectReads, incBO.Results.ItemsRead, "incremental items read")

			assert.NoError(t, incBO.Errors.Failure(), "incremental non-recoverable error", clues.ToCore(incBO.Errors.Failure()))
			assert.Empty(t, incBO.Errors.Recovered(), "incremental recoverable/iteration errors")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "incremental backup-start events")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
			assert.Equal(t,
				incMB.CalledWith[events.BackupStart][0][events.BackupID],
				bupID, "incremental backupID pre-declaration")
		})
	}
}

func (suite *OneDriveBackupIntgSuite) TestBackup_Run_oneDriveOwnerMigration() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acct = tconfig.NewM365Account(t)
		opts = control.Defaults()
		mb   = evmock.NewBus()

		categories = map[path.CategoryType][]string{
			path.FilesCategory: {graph.DeltaURLsFileName, graph.PreviousPathFileName},
		}
	)

	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	ctrl, err := m365.NewController(
		ctx,
		acct,
		resource.Users,
		path.OneDriveService,
		control.Defaults())
	require.NoError(t, err, clues.ToCore(err))

	userable, err := ctrl.AC.Users().GetByID(ctx, suite.its.userID)
	require.NoError(t, err, clues.ToCore(err))

	uid := ptr.Val(userable.GetId())
	uname := ptr.Val(userable.GetUserPrincipalName())

	oldsel := selectors.NewOneDriveBackup([]string{uname})
	oldsel.Include(selTD.OneDriveBackupFolderScope(oldsel))

	bo, bod := prepNewTestBackupOp(t, ctx, mb, oldsel.Selector, opts, 0)
	defer bod.close(t, ctx)

	sel := bod.sel

	// ensure the initial owner uses name in both cases
	bo.ResourceOwner = sel.SetDiscreteOwnerIDName(uname, uname)
	// required, otherwise we don't run the migration
	bo.BackupVersion = version.All8MigrateUserPNToID - 1

	require.Equalf(
		t,
		bo.ResourceOwner.Name(),
		bo.ResourceOwner.ID(),
		"historical representation of user id [%s] should match pn [%s]",
		bo.ResourceOwner.ID(),
		bo.ResourceOwner.Name())

	// run the initial backup
	runAndCheckBackup(t, ctx, &bo, mb, false)

	newsel := selectors.NewOneDriveBackup([]string{uid})
	newsel.Include(selTD.OneDriveBackupFolderScope(newsel))
	sel = newsel.SetDiscreteOwnerIDName(uid, uname)

	var (
		incMB = evmock.NewBus()
		// the incremental backup op should have a proper user ID for the id.
		incBO = newTestBackupOp(t, ctx, bod, incMB, opts)
	)

	require.NotEqualf(
		t,
		incBO.ResourceOwner.Name(),
		incBO.ResourceOwner.ID(),
		"current representation of user: id [%s] should differ from PN [%s]",
		incBO.ResourceOwner.ID(),
		incBO.ResourceOwner.Name())

	err = incBO.Run(ctx)
	require.NoError(t, err, clues.ToCore(err))
	checkBackupIsInManifests(
		t,
		ctx,
		bod.kw,
		bod.sw,
		&incBO,
		sel,
		uid,
		maps.Keys(categories)...)
	checkMetadataFilesExist(
		t,
		ctx,
		incBO.Results.BackupID,
		bod.kw,
		bod.kms,
		creds.AzureTenantID,
		uid,
		path.OneDriveService,
		categories)

	// 2 on read/writes to account for metadata: 1 delta and 1 path.
	assert.LessOrEqual(t, 2, incBO.Results.ItemsWritten, "items written")
	assert.LessOrEqual(t, 1, incBO.Results.NonMetaItemsWritten, "non meta items written")
	assert.LessOrEqual(t, 2, incBO.Results.ItemsRead, "items read")
	assert.NoError(t, incBO.Errors.Failure(), "non-recoverable error", clues.ToCore(incBO.Errors.Failure()))
	assert.Empty(t, incBO.Errors.Recovered(), "recoverable/iteration errors")
	assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "backup-start events")
	assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "backup-end events")
	assert.Equal(t,
		incMB.CalledWith[events.BackupStart][0][events.BackupID],
		incBO.Results.BackupID, "backupID pre-declaration")

	bid := incBO.Results.BackupID
	bup := &backup.Backup{}

	err = bod.kms.Get(ctx, model.BackupSchema, bid, bup)
	require.NoError(t, err, clues.ToCore(err))

	var (
		ssid  = bup.StreamStoreID
		deets details.Details
		ss    = streamstore.NewStreamer(bod.kw, creds.AzureTenantID, path.OneDriveService)
	)

	err = ss.Read(ctx, ssid, streamstore.DetailsReader(details.UnmarshalTo(&deets)), fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	for _, ent := range deets.Entries {
		// 46 is the tenant uuid + "onedrive" + two slashes
		if len(ent.RepoRef) > 46 {
			assert.Contains(t, ent.RepoRef, uid)
		}
	}
}

func (suite *OneDriveBackupIntgSuite) TestBackup_Run_oneDriveExtensions() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		tenID  = tconfig.M365TenantID(t)
		mb     = evmock.NewBus()
		userID = tconfig.SecondaryM365UserID(t)
		osel   = selectors.NewOneDriveBackup([]string{userID})
		ws     = deeTD.DriveIDFromRepoRef
		svc    = path.OneDriveService
		opts   = control.Defaults()
	)

	opts.ItemExtensionFactory = getTestExtensionFactories()

	osel.Include(selTD.OneDriveBackupFolderScope(osel))

	bo, bod := prepNewTestBackupOp(t, ctx, mb, osel.Selector, opts, version.Backup)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)

	bID := bo.Results.BackupID

	deets, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bID,
		tenID,
		bod.sel.ID(),
		svc,
		ws,
		bod.kms,
		bod.sss)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bID,
		ws,
		bod.kms,
		bod.sss,
		expectDeets,
		false)

	// Check that the extensions are in the backup
	for _, ent := range deets.Entries {
		if ent.Folder == nil {
			verifyExtensionData(t, ent.ItemInfo, path.OneDriveService)
		}
	}
}

type OneDriveRestoreIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestOneDriveRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &OneDriveRestoreIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *OneDriveRestoreIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *OneDriveRestoreIntgSuite) TestRestore_Run_onedriveWithAdvancedOptions() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// a backup is required to run restores

	baseSel := selectors.NewOneDriveBackup([]string{suite.its.userID})
	baseSel.Include(selTD.OneDriveBackupFolderScope(baseSel))

	baseSel.DiscreteOwner = suite.its.userID

	var (
		mb   = evmock.NewBus()
		opts = control.Defaults()
	)

	bo, bod := prepNewTestBackupOp(t, ctx, mb, baseSel.Selector, opts, version.Backup)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)

	rsel, err := baseSel.ToOneDriveRestore()
	require.NoError(t, err, clues.ToCore(err))

	var (
		restoreCfg  = ctrlTD.DefaultRestoreConfig("onedrive_adv_restore")
		sel         = rsel.Selector
		userDriveID = suite.its.userDriveID
		containerID string
		collKeys    = map[string]api.DriveCollisionItem{}
		acd         = suite.its.ac.Drives()
	)

	// initial restore

	suite.Run("baseline", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()

		restoreCfg.OnCollision = control.Copy

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			count.New(),
			sel,
			opts,
			restoreCfg)

		runAndCheckRestore(t, ctx, &ro, mb, -1, false)

		// get all files in folder, use these as the base
		// set of files to compare against.
		contGC, err := acd.GetFolderByName(ctx, userDriveID, suite.its.userDriveRootFolderID, restoreCfg.Location)
		require.NoError(t, err, clues.ToCore(err))

		contGC, err = acd.GetFolderByName(ctx, userDriveID, ptr.Val(contGC.GetId()), selTD.TestFolderName)
		require.NoError(t, err, clues.ToCore(err))

		containerID = ptr.Val(contGC.GetId())

		collKeys, err = acd.GetItemsInContainerByCollisionKey(
			ctx,
			userDriveID,
			containerID)
		require.NoError(t, err, clues.ToCore(err))
	})

	// skip restore

	suite.Run("skip collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Skip

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		deets := runAndCheckRestore(t, ctx, &ro, mb, 0, false)

		assert.Equal(
			t,
			int64(len(collKeys)),
			ctr.Total(count.CollisionSkip),
			"all attempted item restores should have been skipped")
		assert.Zero(
			t,
			len(deets.Entries),
			"no items should have been restored")

		// get all files in folder, use these as the base
		// set of files to compare against.

		result := filterCollisionKeyResults(
			t,
			ctx,
			userDriveID,
			containerID,
			giicbcker[api.DriveCollisionItem](acd),
			collKeys)

		assert.Len(t, result, 0, "no new items should get added")
	})

	// replace restore

	suite.Run("replace collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Replace

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		deets := runAndCheckRestore(t, ctx, &ro, mb, len(collKeys), false)
		filtEnts := []details.Entry{}

		for _, e := range deets.Entries {
			if e.Folder == nil {
				filtEnts = append(filtEnts, e)
			}
		}

		assert.Zero(
			t,
			ctr.Total(count.CollisionSkip),
			"no attempted item restores should have been skipped")
		assert.Equal(
			t,
			len(filtEnts),
			len(collKeys),
			"every item should have been replaced: %+v", filtEnts)

		result := filterCollisionKeyResults(
			t,
			ctx,
			userDriveID,
			containerID,
			giicbcker[api.DriveCollisionItem](acd),
			collKeys)

		assert.Len(t, result, 0, "all items should have been replaced")

		for k, v := range result {
			assert.NotEqual(t, v, collKeys[k], "replaced items should have new IDs")
		}
	})

	// copy restore

	suite.Run("copy collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Copy

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		deets := runAndCheckRestore(t, ctx, &ro, mb, len(collKeys), false)
		filtEnts := []details.Entry{}

		for _, e := range deets.Entries {
			if e.Folder == nil {
				filtEnts = append(filtEnts, e)
			}
		}

		assert.Zero(
			t,
			ctr.Total(count.CollisionSkip),
			"no attempted item restores should have been skipped")
		assert.Equal(
			t,
			len(filtEnts),
			len(collKeys),
			"every item should have been copied: %+v", filtEnts)

		result := filterCollisionKeyResults(
			t,
			ctx,
			userDriveID,
			containerID,
			giicbcker[api.DriveCollisionItem](acd),
			collKeys)

		assert.Len(t, result, len(collKeys), "all items should have been added as copies")
	})
}
