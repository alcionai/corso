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
	"golang.org/x/exp/maps"

	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/syncd"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	deeTD "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func runBasicDriveishBackupTests(
	suite tester.Suite,
	service path.ServiceType,
	opts control.Options,
	sel selectors.Selector,
) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		tenID   = tconfig.M365TenantID(t)
		mb      = evmock.NewBus()
		counter = count.New()
		ws      = deeTD.DriveIDFromRepoRef
	)

	bo, bod := prepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup, counter)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)

	bID := bo.Results.BackupID

	_, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bID,
		tenID,
		bod.sel.ID(),
		service,
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

func runIncrementalDriveishBackupTest(
	suite tester.Suite,
	opts control.Options,
	owner, permissionsUser string,
	service path.ServiceType,
	category path.CategoryType,
	includeContainers func([]string) selectors.Selector,
	getTestDriveID func(*testing.T, context.Context) string,
	getTestSiteID func(*testing.T, context.Context) string,
	getRestoreHandler func(api.Client) drive.RestoreHandler,
	skipPermissionsTests bool,
) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acct    = tconfig.NewM365Account(t)
		mb      = evmock.NewBus()
		counter = count.New()
		ws      = deeTD.DriveIDFromRepoRef

		// `now` has to be formatted with SimpleDateTimeTesting as
		// some drives cannot have `:` in file/folder names
		now = dttm.FormatNow(dttm.SafeForTesting)

		categories      = map[path.CategoryType][][]string{}
		container1      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 1, now)
		container2      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 2, now)
		container3      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 3, now)
		containerRename = "renamed_folder"

		genDests = []string{container1, container2}

		// container3 does not exist yet. It will get created later on
		// during the tests.
		containers = []string{container1, container2, container3}
	)

	if service == path.GroupsService && category == path.LibrariesCategory {
		categories[category] = [][]string{{odConsts.SitesPathDir, bupMD.PreviousPathFileName}}
	} else {
		categories[category] = [][]string{{bupMD.DeltaURLsFileName}, {bupMD.PreviousPathFileName}}
	}

	sel := includeContainers(containers)

	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	ctrl, sel := ControllerWithSelector(t, ctx, acct, sel, nil, nil, counter)
	ac := ctrl.AC.Drives()
	rh := getRestoreHandler(ctrl.AC)

	roidn := inMock.NewProvider(sel.ID(), sel.Name())

	var (
		atid    = creds.AzureTenantID
		driveID = getTestDriveID(t, ctx)
		siteID  = ""
		fileDBF = func(id, timeStamp, subject, body string) []byte {
			return []byte(id + subject)
		}
		makeLocRef = func(flds ...string) *path.Builder {
			elems := append([]string{"root:"}, flds...)
			return path.Builder{}.Append(elems...)
		}
	)

	// Will only be available for groups
	if getTestSiteID != nil {
		siteID = getTestSiteID(t, ctx)
	}

	rrPfx, err := path.BuildPrefix(atid, roidn.ID(), service, category)
	require.NoError(t, err, clues.ToCore(err))

	// strip the category from the prefix; we primarily want the tenant and resource owner.
	expectDeets := deeTD.NewInDeets(rrPfx.ToBuilder().Dir().String())

	type containerInfo struct {
		id     string
		locRef *path.Builder
	}

	containerInfos := map[string]containerInfo{}

	mustGetExpectedContainerItems := func(
		t *testing.T,
		driveID, destName string,
		locRef *path.Builder,
	) {
		// Use path-based indexing to get the folder's ID.
		itemURL := fmt.Sprintf(
			"https://graph.microsoft.com/v1.0/drives/%s/root:/%s",
			driveID,
			locRef.String())
		resp, err := drives.
			NewItemItemsDriveItemItemRequestBuilder(itemURL, ctrl.AC.Stable.Adapter()).
			Get(ctx, nil)
		require.NoError(
			t,
			err,
			"getting drive folder ID for %s: %v",
			locRef.String(),
			clues.ToCore(err))

		containerInfos[destName] = containerInfo{
			id:     ptr.Val(resp.GetId()),
			locRef: makeLocRef(locRef.Elements()...),
		}
		dest := containerInfos[destName]

		items, err := ac.GetItemsInContainerByCollisionKey(
			ctx,
			driveID,
			ptr.Val(resp.GetId()))
		require.NoError(
			t,
			err,
			"getting container %s items: %v",
			locRef.String(),
			clues.ToCore(err))

		// Add the directory and all its ancestors to the cache so we can compare
		// folders.
		for pb := dest.locRef; len(pb.Elements()) > 0; pb = pb.Dir() {
			expectDeets.AddLocation(driveID, pb.String())
		}

		for _, item := range items {
			if item.IsFolder {
				continue
			}

			expectDeets.AddItem(
				driveID,
				dest.locRef.String(),
				item.ItemID)
		}
	}

	// Populate initial test data.
	// Generate 2 new folders with two items each. Only the first two
	// folders will be part of the initial backup and
	// incrementals. The third folder will be introduced partway
	// through the changes. This should be enough to cover most delta
	// actions.
	for _, destName := range genDests {
		generateContainerOfItems(
			t,
			ctx,
			ctrl,
			service,
			category,
			sel,
			atid, roidn.ID(), siteID, driveID, destName,
			2,
			// Use an old backup version so we don't need metadata files.
			0,
			fileDBF)

		// The way we generate containers causes it to duplicate the destName.
		locRef := path.Builder{}.Append(destName, destName)
		mustGetExpectedContainerItems(
			t,
			driveID,
			destName,
			locRef)
	}

	bo, bod := prepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup, counter)
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

		permissionIDMappings = syncd.NewMapTo[string]()
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
				targetContainer := containerInfos[container1]
				driveItem := models.NewDriveItem()
				driveItem.SetName(&newFileName)
				driveItem.SetFile(models.NewFile())
				newFile, err = ac.PostItemInContainer(
					ctx,
					driveID,
					targetContainer.id,
					driveItem,
					control.Copy)
				require.NoErrorf(t, err, "creating new file %v", clues.ToCore(err))

				newFileID = ptr.Val(newFile.GetId())

				expectDeets.AddItem(driveID, targetContainer.locRef.String(), newFileID)
			},
			itemsRead:           1, // .data file for newitem
			itemsWritten:        4, // .data and .meta for newitem, .dirmeta for parent and ancestor
			nonMetaItemsWritten: 1, // .data file for newitem
		},
		{
			name: "add permission to new file",
			updateFiles: func(t *testing.T, ctx context.Context) {
				err = drive.UpdatePermissions(
					ctx,
					rh,
					driveID,
					ptr.Val(newFile.GetId()),
					[]metadata.Permission{writePerm},
					[]metadata.Permission{},
					permissionIDMappings,
					fault.New(true))
				require.NoErrorf(t, err, "adding permission to file %v", clues.ToCore(err))
				// no expectedDeets: metadata isn't tracked
			},
			itemsRead:           1, // .data file for newitem
			itemsWritten:        3, // .meta for newitem, .dirmeta for parent (.data is not written as it is not updated)
			nonMetaItemsWritten: 0, // none because the file is considered cached instead of written.
		},
		{
			name: "remove permission from new file",
			updateFiles: func(t *testing.T, ctx context.Context) {
				err = drive.UpdatePermissions(
					ctx,
					rh,
					driveID,
					*newFile.GetId(),
					[]metadata.Permission{},
					[]metadata.Permission{writePerm},
					permissionIDMappings,
					fault.New(true))
				require.NoErrorf(t, err, "removing permission from file %v", clues.ToCore(err))
				// no expectedDeets: metadata isn't tracked
			},
			itemsRead:           1, // .data file for newitem
			itemsWritten:        3, // .meta for newitem, .dirmeta for parent (.data is not written as it is not updated)
			nonMetaItemsWritten: 0, // none because the file is considered cached instead of written.
		},
		{
			name: "add permission to container",
			updateFiles: func(t *testing.T, ctx context.Context) {
				targetContainer := containerInfos[container1].id
				err = drive.UpdatePermissions(
					ctx,
					rh,
					driveID,
					targetContainer,
					[]metadata.Permission{writePerm},
					[]metadata.Permission{},
					permissionIDMappings,
					fault.New(true))
				require.NoErrorf(t, err, "adding permission to container %v", clues.ToCore(err))
				// no expectedDeets: metadata isn't tracked
			},
			itemsRead:           0,
			itemsWritten:        2, // .dirmeta for collection
			nonMetaItemsWritten: 0, // no files updated as update on container
		},
		{
			name: "remove permission from container",
			updateFiles: func(t *testing.T, ctx context.Context) {
				targetContainer := containerInfos[container1].id
				err = drive.UpdatePermissions(
					ctx,
					rh,
					driveID,
					targetContainer,
					[]metadata.Permission{},
					[]metadata.Permission{writePerm},
					permissionIDMappings,
					fault.New(true))
				require.NoErrorf(t, err, "removing permission from container %v", clues.ToCore(err))
				// no expectedDeets: metadata isn't tracked
			},
			itemsRead:           0,
			itemsWritten:        2, // .dirmeta for collection
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
			itemsWritten:        4, // .data and .meta for newitem, .dirmeta for parent
			nonMetaItemsWritten: 1, // .data  file for newitem
		},
		{
			name: "rename a file",
			updateFiles: func(t *testing.T, ctx context.Context) {
				driveItem := models.NewDriveItem()
				name := "renamed_new_file.txt"
				driveItem.SetName(&name)

				err := ac.PatchItem(
					ctx,
					driveID,
					ptr.Val(newFile.GetId()),
					driveItem)
				require.NoError(t, err, "renaming file %v", clues.ToCore(err))
			},
			itemsRead:           1, // .data file for newitem
			itemsWritten:        4, // .data and .meta for newitem, .dirmeta for parent
			nonMetaItemsWritten: 1, // .data file for newitem
			// no expectedDeets: neither file id nor location changed
		},
		{
			name: "move a file between folders",
			updateFiles: func(t *testing.T, ctx context.Context) {
				dest := containerInfos[container2]

				driveItem := models.NewDriveItem()
				parentRef := models.NewItemReference()
				parentRef.SetId(&dest.id)
				driveItem.SetParentReference(parentRef)

				err := ac.PatchItem(
					ctx,
					driveID,
					ptr.Val(newFile.GetId()),
					driveItem)
				require.NoErrorf(t, err, "moving file between folders %v", clues.ToCore(err))

				expectDeets.MoveItem(
					driveID,
					containerInfos[container1].locRef.String(),
					dest.locRef.String(),
					ptr.Val(newFile.GetId()))
			},
			itemsRead:           1, // .data file for newitem
			itemsWritten:        4, // .data and .meta for newitem, .dirmeta for parent
			nonMetaItemsWritten: 1, // .data file for moved item
		},
		{
			name: "boomerang a file",
			updateFiles: func(t *testing.T, ctx context.Context) {
				dest := containerInfos[container2]
				temp := containerInfos[container1]

				driveItem := models.NewDriveItem()
				parentRef := models.NewItemReference()
				parentRef.SetId(&temp.id)
				driveItem.SetParentReference(parentRef)

				err := ac.PatchItem(
					ctx,
					driveID,
					ptr.Val(newFile.GetId()),
					driveItem)
				require.NoErrorf(t, err, "moving file to temporary folder %v", clues.ToCore(err))

				parentRef.SetId(&dest.id)
				driveItem.SetParentReference(parentRef)

				err = ac.PatchItem(
					ctx,
					driveID,
					ptr.Val(newFile.GetId()),
					driveItem)
				require.NoErrorf(t, err, "moving file back to folder %v", clues.ToCore(err))
			},
			itemsRead:           1, // .data file for newitem
			itemsWritten:        3, // .data and .meta for newitem, .dirmeta for parent
			nonMetaItemsWritten: 0, // non because the file is considered cached instead of written.
		},
		{
			name: "delete file",
			updateFiles: func(t *testing.T, ctx context.Context) {
				err := ac.DeleteItem(
					ctx,
					driveID,
					ptr.Val(newFile.GetId()))
				require.NoErrorf(t, err, "deleting file %v", clues.ToCore(err))

				expectDeets.RemoveItem(
					driveID,
					containerInfos[container2].locRef.String(),
					ptr.Val(newFile.GetId()))
			},
			itemsRead:           0,
			itemsWritten:        0,
			nonMetaItemsWritten: 0,
		},
		{
			name: "move a folder to a subfolder",
			updateFiles: func(t *testing.T, ctx context.Context) {
				parent := containerInfos[container1]
				child := containerInfos[container2]

				driveItem := models.NewDriveItem()
				parentRef := models.NewItemReference()
				parentRef.SetId(&parent.id)
				driveItem.SetParentReference(parentRef)

				err := ac.PatchItem(
					ctx,
					driveID,
					child.id,
					driveItem)
				require.NoError(t, err, "moving folder", clues.ToCore(err))

				expectDeets.MoveLocation(
					driveID,
					child.locRef.String(),
					parent.locRef.String())

				// Remove parent of moved folder since it's now empty.
				expectDeets.RemoveLocation(driveID, child.locRef.Dir().String())

				// Update in-memory cache with new location.
				child.locRef = path.Builder{}.Append(append(
					parent.locRef.Elements(),
					child.locRef.LastElem())...)
				containerInfos[container2] = child
			},
			itemsRead:           0,
			itemsWritten:        7, // 2*2(data and meta of 2 files) + 3 (dirmeta of two moved folders and target)
			nonMetaItemsWritten: 0,
		},
		{
			name: "rename a folder",
			updateFiles: func(t *testing.T, ctx context.Context) {
				child := containerInfos[container2]

				driveItem := models.NewDriveItem()
				driveItem.SetName(&containerRename)

				err := ac.PatchItem(
					ctx,
					driveID,
					child.id,
					driveItem)
				require.NoError(t, err, "renaming folder", clues.ToCore(err))

				containerInfos[containerRename] = containerInfo{
					id:     child.id,
					locRef: child.locRef.Dir().Append(containerRename),
				}

				expectDeets.RenameLocation(
					driveID,
					child.locRef.String(),
					containerInfos[containerRename].locRef.String())

				delete(containerInfos, container2)
			},
			itemsRead:           0,
			itemsWritten:        7, // 2*2(data and meta of 2 files) + 3 (dirmeta of two moved folders and target)
			nonMetaItemsWritten: 0,
		},
		{
			name: "delete a folder",
			updateFiles: func(t *testing.T, ctx context.Context) {
				container := containerInfos[containerRename]
				err := ac.DeleteItem(
					ctx,
					driveID,
					container.id)
				require.NoError(t, err, "deleting folder", clues.ToCore(err))

				expectDeets.RemoveLocation(driveID, container.locRef.String())

				delete(containerInfos, containerRename)
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
					atid, roidn.ID(), siteID, driveID, container3,
					2,
					0,
					fileDBF)

				locRef := path.Builder{}.Append(container3, container3)
				mustGetExpectedContainerItems(
					t,
					driveID,
					container3,
					locRef)
			},
			itemsRead:           2, // 2 .data for 2 files
			itemsWritten:        6, // read items + 2 directory meta
			nonMetaItemsWritten: 2, // 2 .data for 2 files
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			cleanCtrl, err := m365.NewController(
				ctx,
				acct,
				sel.PathService(),
				control.DefaultOptions(),
				count.New())
			require.NoError(t, err, clues.ToCore(err))

			bod.ctrl = cleanCtrl

			var (
				t       = suite.T()
				incMB   = evmock.NewBus()
				counter = count.New()
				incBO   = newTestBackupOp(
					t,
					ctx,
					bod,
					incMB,
					opts,
					counter)
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
			var (
				expectWrites        = test.itemsWritten
				expectNonMetaWrites = test.nonMetaItemsWritten
				expectReads         = test.itemsRead
				assertReadWrite     = assert.Equal
			)

			if service == path.GroupsService && category == path.LibrariesCategory {
				// Groups SharePoint have an extra metadata file at
				// /libraries/sites/previouspath
				expectWrites++
				expectReads++

				// +2 on read/writes to account for metadata: 1 delta and 1 path (for each site)
				sites, err := ac.Groups().GetAllSites(ctx, owner, fault.New(true))
				require.NoError(t, err, clues.ToCore(err))

				expectWrites += len(sites) * 2
				expectReads += len(sites) * 2
			} else {
				// +2 on read/writes to account for metadata: 1 delta and 1 path.
				expectWrites += 2
				expectReads += 2
			}

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
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
		})
	}
}

func runDriveishBackupWithExtensionsTests(
	suite tester.Suite,
	service path.ServiceType,
	opts control.Options,
	sel selectors.Selector,
) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		tenID   = tconfig.M365TenantID(t)
		mb      = evmock.NewBus()
		counter = count.New()
		ws      = deeTD.DriveIDFromRepoRef
	)

	opts.ItemExtensionFactory = getTestExtensionFactories()

	bo, bod := prepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup, counter)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)

	bID := bo.Results.BackupID

	deets, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bID,
		tenID,
		bod.sel.ID(),
		service,
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
			verifyExtensionData(t, ent.ItemInfo, service)
		}
	}
}
