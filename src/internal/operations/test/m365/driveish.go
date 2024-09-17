package m365

import (
	"context"
	"fmt"
	"io"
	"sync/atomic"
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
	"github.com/alcionai/corso/src/internal/tester/its"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	deeTD "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	ctrlTD "github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/extensions"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func RunBasicDriveishBackupTests(
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

	bo, bod := PrepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup, counter)
	defer bod.Close(t, ctx)

	RunAndCheckBackup(t, ctx, &bo, mb, false)

	bID := bo.Results.BackupID

	_, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bID,
		tenID,
		bod.Sel.ID(),
		service,
		ws,
		bod.KMS,
		bod.SSS)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bID,
		ws,
		bod.KMS,
		bod.SSS,
		expectDeets,
		false)
}

func RunIncrementalDriveishBackupTest(
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
		container1      = fmt.Sprintf("%s%d_%s", IncrementalsDestContainerPrefix, 1, now)
		container2      = fmt.Sprintf("%s%d_%s", IncrementalsDestContainerPrefix, 2, now)
		container3      = fmt.Sprintf("%s%d_%s", IncrementalsDestContainerPrefix, 3, now)
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
		GenerateContainerOfItems(
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

	bo, bod := PrepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup, counter)
	defer bod.Close(t, ctx)

	sel = bod.Sel

	// run the initial backup
	RunAndCheckBackup(t, ctx, &bo, mb, false)

	// precheck to ensure the expectedDeets are correct.
	// if we fail here, the expectedDeets were populated incorrectly.
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bo.Results.BackupID,
		ws,
		bod.KMS,
		bod.SSS,
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

		// TODO: Temporary mechanism to skip permissions
		// related tests. Remove once we figure out the issue.
		skipChecks bool
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
			skipChecks:          true,
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
			skipChecks:          true,
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
			skipChecks:          true,
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
			skipChecks:          true,
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
				GenerateContainerOfItems(
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

			bod.Ctrl = cleanCtrl

			var (
				t       = suite.T()
				incMB   = evmock.NewBus()
				counter = count.New()
				incBO   = NewTestBackupOp(
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

			CheckBackupIsInManifests(
				t,
				ctx,
				bod.KW,
				bod.SW,
				&incBO,
				sel,
				roidn.ID(),
				maps.Keys(categories)...)
			CheckMetadataFilesExist(
				t,
				ctx,
				bupID,
				bod.KW,
				bod.KMS,
				atid,
				roidn.ID(),
				service,
				categories)
			deeTD.CheckBackupDetails(
				t,
				ctx,
				bupID,
				ws,
				bod.KMS,
				bod.SSS,
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

			if !test.skipChecks {
				assertReadWrite(t, expectWrites, incBO.Results.ItemsWritten, "incremental items written")
				assertReadWrite(t, expectNonMetaWrites, incBO.Results.NonMetaItemsWritten, "incremental non-meta items written")
				assertReadWrite(t, expectReads, incBO.Results.ItemsRead, "incremental items read")
			}

			assert.NoError(t, incBO.Errors.Failure(), "incremental non-recoverable error", clues.ToCore(incBO.Errors.Failure()))
			assert.Empty(t, incBO.Errors.Recovered(), "incremental recoverable/iteration errors")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
		})
	}
}

func RunDriveishBackupWithExtensionsTests(
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

	opts.ItemExtensionFactory = GetTestExtensionFactories()

	bo, bod := PrepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup, counter)
	defer bod.Close(t, ctx)

	RunAndCheckBackup(t, ctx, &bo, mb, false)

	bID := bo.Results.BackupID

	deets, expectDeets := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bID,
		tenID,
		bod.Sel.ID(),
		service,
		ws,
		bod.KMS,
		bod.SSS)
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bID,
		ws,
		bod.KMS,
		bod.SSS,
		expectDeets,
		false)

	// Check that the extensions are in the backup
	for _, ent := range deets.Entries {
		if ent.Folder == nil {
			VerifyExtensionData(t, ent.ItemInfo, service)
		}
	}
}

var (
	_ io.ReadCloser                    = &FailFirstRead{}
	_ extensions.CreateItemExtensioner = &CreateFailFirstRead{}
)

// FailFirstRead fails the first read on a file being uploaded during a
// snapshot. Only one file is failed during the snapshot even if it the snapshot
// contains multiple files.
type FailFirstRead struct {
	firstFile *atomic.Bool
	io.ReadCloser
}

func (e *FailFirstRead) Read(p []byte) (int, error) {
	if e.firstFile.CompareAndSwap(true, false) {
		// This is the first file being read, return an error for it.
		return 0, clues.New("injected error for testing")
	}

	return e.ReadCloser.Read(p)
}

type CreateFailFirstRead struct {
	firstItem *atomic.Bool
}

func (ce *CreateFailFirstRead) CreateItemExtension(
	_ context.Context,
	r io.ReadCloser,
	_ details.ItemInfo,
	_ *details.ExtensionData,
) (io.ReadCloser, error) {
	return &FailFirstRead{
		firstFile:  ce.firstItem,
		ReadCloser: r,
	}, nil
}

func NewCreateSingleFileFailExtension() *CreateFailFirstRead {
	firstItem := &atomic.Bool{}
	firstItem.Store(true)

	return &CreateFailFirstRead{
		firstItem: firstItem,
	}
}

func RunDriveAssistBaseGroupsUpdate(
	suite tester.Suite,
	sel selectors.Selector,
	expectCached bool,
) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		whatSet = deeTD.CategoryFromRepoRef
		mb      = evmock.NewBus()
		counter = count.New()
		opts    = control.DefaultOptions()
	)

	opts.ItemExtensionFactory = []extensions.CreateItemExtensioner{
		NewCreateSingleFileFailExtension(),
	}

	// Creating out here so bod lasts for full test and isn't closed until the
	// test is compltely done.
	bo, bod := PrepNewTestBackupOp(
		t,
		ctx,
		mb,
		sel,
		opts,
		version.All8MigrateUserPNToID,
		counter)
	defer bod.Close(t, ctx)

	suite.Run("makeAssistBackup", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		// Need to run manually cause runAndCheckBackup assumes success for the most
		// part.
		err := bo.Run(ctx)
		assert.Error(t, err, clues.ToCore(err))
		assert.NotEmpty(t, bo.Results, "backup had non-zero results")
		assert.NotEmpty(t, bo.Results.BackupID, "backup generated an ID")
		assert.NotZero(t, bo.Results.ItemsWritten)

		// TODO(ashmrtn): Check that the base is marked as an assist base.
		t.Logf("base error: %v\n", err)
	})

	// Don't run the below if we've already failed since it won't make sense
	// anymore.
	if suite.T().Failed() {
		return
	}

	suite.Run("makeIncrementalBackup", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		var (
			mb      = evmock.NewBus()
			counter = count.New()
			opts    = control.DefaultOptions()
		)

		forcedFull := NewTestBackupOp(
			t,
			ctx,
			bod,
			mb,
			opts,
			counter)
		forcedFull.BackupVersion = version.Groups9Update

		RunAndCheckBackup(t, ctx, &forcedFull, mb, false)

		reasons, err := bod.Sel.Reasons(bod.Acct.ID(), false)
		require.NoError(t, err, clues.ToCore(err))

		for _, reason := range reasons {
			CheckBackupIsInManifests(
				t,
				ctx,
				bod.KW,
				bod.SW,
				&forcedFull,
				bod.Sel,
				bod.Sel.ID(),
				reason.Category())
		}

		_, expectDeets := deeTD.GetDeetsInBackup(
			t,
			ctx,
			forcedFull.Results.BackupID,
			bod.Acct.ID(),
			bod.Sel.ID(),
			bod.Sel.PathService(),
			whatSet,
			bod.KMS,
			bod.SSS)
		deeTD.CheckBackupDetails(
			t,
			ctx,
			forcedFull.Results.BackupID,
			whatSet,
			bod.KMS,
			bod.SSS,
			expectDeets,
			false)

		// For groups the forced full backup shouldn't have any cached items. For
		// OneDrive and SharePoint it should since they shouldn't be forcing full
		// backups.
		cachedCheck := assert.NotZero
		if !expectCached {
			cachedCheck = assert.Zero
		}

		cachedCheck(
			t,
			forcedFull.Results.Counts[string(count.PersistedCachedFiles)],
			"kopia cached items")
	})
}

func RunDriveRestoreToAlternateProtectedResource(
	t *testing.T,
	suite tester.Suite,
	ac api.Client,
	sel selectors.Selector, // owner should match 'from', both Restore and Backup types work.
	driveFrom, driveTo its.IDs,
	toResource string,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	// a backup is required to run restores

	var (
		mb      = evmock.NewBus()
		counter = count.New()
		opts    = control.DefaultOptions()
	)

	bo, bod := PrepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup, counter)
	defer bod.Close(t, ctx)

	RunAndCheckBackup(t, ctx, &bo, mb, false)

	var (
		restoreCfg        = ctrlTD.DefaultRestoreConfig("drive_restore_to_resource")
		fromCollisionKeys map[string]api.DriveItemIDType
		fromItemIDs       map[string]api.DriveItemIDType
		acd               = ac.Drives()
	)

	// first restore to the 'from' resource

	suite.Run("restore original resource", func() {
		mb = evmock.NewBus()
		fromCtr := count.New()
		driveID := driveFrom.DriveID
		rootFolderID := driveFrom.DriveRootFolderID
		restoreCfg.OnCollision = control.Copy

		ro, _ := PrepNewTestRestoreOp(
			t,
			ctx,
			bod.St,
			bo.Results.BackupID,
			mb,
			fromCtr,
			sel,
			opts,
			restoreCfg)

		RunAndCheckRestore(t, ctx, &ro, mb, false)

		// get all files in folder, use these as the base
		// set of files to compare against.
		fromItemIDs, fromCollisionKeys = GetDriveCollKeysAndItemIDs(
			t,
			ctx,
			acd,
			driveID,
			rootFolderID,
			restoreCfg.Location,
			selTD.TestFolderName)
	})

	// then restore to the 'to' resource
	var (
		toCollisionKeys map[string]api.DriveItemIDType
		toItemIDs       map[string]api.DriveItemIDType
	)

	suite.Run("restore to alternate resource", func() {
		mb = evmock.NewBus()
		toCtr := count.New()
		driveID := driveTo.DriveID
		rootFolderID := driveTo.DriveRootFolderID
		restoreCfg.ProtectedResource = toResource

		ro, _ := PrepNewTestRestoreOp(
			t,
			ctx,
			bod.St,
			bo.Results.BackupID,
			mb,
			toCtr,
			sel,
			opts,
			restoreCfg)

		RunAndCheckRestore(t, ctx, &ro, mb, false)

		// get all files in folder, use these as the base
		// set of files to compare against.
		toItemIDs, toCollisionKeys = GetDriveCollKeysAndItemIDs(
			t,
			ctx,
			acd,
			driveID,
			rootFolderID,
			restoreCfg.Location,
			selTD.TestFolderName)
	})

	// compare restore results
	assert.Equal(t, len(fromItemIDs), len(toItemIDs))
	assert.ElementsMatch(t, maps.Keys(fromCollisionKeys), maps.Keys(toCollisionKeys))
}

type GetItemsKeysAndFolderByNameer interface {
	GetItemIDsInContainer(
		ctx context.Context,
		driveID, containerID string,
	) (map[string]api.DriveItemIDType, error)
	GetFolderByName(
		ctx context.Context,
		driveID, parentFolderID, folderName string,
	) (models.DriveItemable, error)
	GetItemsInContainerByCollisionKey(
		ctx context.Context,
		driveID, containerID string,
	) (map[string]api.DriveItemIDType, error)
}

func GetDriveCollKeysAndItemIDs(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	gikafbn GetItemsKeysAndFolderByNameer,
	driveID, parentContainerID string,
	containerNames ...string,
) (map[string]api.DriveItemIDType, map[string]api.DriveItemIDType) {
	var (
		c   models.DriveItemable
		err error
		cID string
	)

	for _, cn := range containerNames {
		pcid := parentContainerID

		if len(cID) != 0 {
			pcid = cID
		}

		c, err = gikafbn.GetFolderByName(ctx, driveID, pcid, cn)
		require.NoError(t, err, clues.ToCore(err))

		cID = ptr.Val(c.GetId())
	}

	itemIDs, err := gikafbn.GetItemIDsInContainer(ctx, driveID, cID)
	require.NoError(t, err, clues.ToCore(err))

	collisionKeys, err := gikafbn.GetItemsInContainerByCollisionKey(ctx, driveID, cID)
	require.NoError(t, err, clues.ToCore(err))

	return itemIDs, collisionKeys
}

func RunDriveRestoreWithAdvancedOptions(
	t *testing.T,
	suite tester.Suite,
	ac api.Client,
	sel selectors.Selector, // both Restore and Backup types work.
	driveID, rootFolderID string,
) {
	ctx, flush := tester.NewContext(t)
	defer flush()

	// a backup is required to run restores

	var (
		mb      = evmock.NewBus()
		counter = count.New()
		opts    = control.DefaultOptions()
	)

	bo, bod := PrepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup, counter)
	defer bod.Close(t, ctx)

	RunAndCheckBackup(t, ctx, &bo, mb, false)

	var (
		restoreCfg          = ctrlTD.DefaultRestoreConfig("drive_adv_restore")
		containerID         string
		countItemsInRestore int
		collKeys            = map[string]api.DriveItemIDType{}
		fileIDs             map[string]api.DriveItemIDType
		acd                 = ac.Drives()
	)

	// initial restore

	suite.Run("baseline", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Copy

		ro, _ := PrepNewTestRestoreOp(
			t,
			ctx,
			bod.St,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		RunAndCheckRestore(t, ctx, &ro, mb, false)

		// get all files in folder, use these as the base
		// set of files to compare against.
		contGC, err := acd.GetFolderByName(ctx, driveID, rootFolderID, restoreCfg.Location)
		require.NoError(t, err, clues.ToCore(err))

		// the folder containing the files is a child of the folder created by the restore.
		contGC, err = acd.GetFolderByName(ctx, driveID, ptr.Val(contGC.GetId()), selTD.TestFolderName)
		require.NoError(t, err, clues.ToCore(err))

		containerID = ptr.Val(contGC.GetId())

		collKeys, err = acd.GetItemsInContainerByCollisionKey(
			ctx,
			driveID,
			containerID)
		require.NoError(t, err, clues.ToCore(err))

		countItemsInRestore = len(collKeys)

		CheckRestoreCounts(t, ctr, 0, 0, countItemsInRestore)

		fileIDs, err = acd.GetItemIDsInContainer(ctx, driveID, containerID)
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

		ro, _ := PrepNewTestRestoreOp(
			t,
			ctx,
			bod.St,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		deets := RunAndCheckRestore(t, ctx, &ro, mb, false)

		CheckRestoreCounts(t, ctr, countItemsInRestore, 0, 0)
		assert.Zero(
			t,
			len(deets.Entries),
			"no items should have been restored")

		// get all files in folder, use these as the base
		// set of files to compare against.

		result := FilterCollisionKeyResults(
			t,
			ctx,
			driveID,
			containerID,
			GetItemsInContainerByCollisionKeyer[api.DriveItemIDType](acd),
			collKeys)

		assert.Len(t, result, 0, "no new items should get added")

		currentFileIDs, err := acd.GetItemIDsInContainer(ctx, driveID, containerID)
		require.NoError(t, err, clues.ToCore(err))

		assert.Equal(t, fileIDs, currentFileIDs, "ids are equal")
	})

	// replace restore

	suite.Run("replace collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Replace

		ro, _ := PrepNewTestRestoreOp(
			t,
			ctx,
			bod.St,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		deets := RunAndCheckRestore(t, ctx, &ro, mb, false)
		filtEnts := []details.Entry{}

		for _, e := range deets.Entries {
			if e.Folder == nil {
				filtEnts = append(filtEnts, e)
			}
		}

		CheckRestoreCounts(t, ctr, 0, countItemsInRestore, 0)
		assert.Len(
			t,
			filtEnts,
			countItemsInRestore,
			"every item should have been replaced")

		result := FilterCollisionKeyResults(
			t,
			ctx,
			driveID,
			containerID,
			GetItemsInContainerByCollisionKeyer[api.DriveItemIDType](acd),
			collKeys)

		assert.Len(t, result, 0, "all items should have been replaced")

		for k, v := range result {
			assert.NotEqual(t, v, collKeys[k], "replaced items should have new IDs")
		}

		currentFileIDs, err := acd.GetItemIDsInContainer(ctx, driveID, containerID)
		require.NoError(t, err, clues.ToCore(err))

		assert.Equal(t, len(fileIDs), len(currentFileIDs), "count of ids ids are equal")
		for orig := range fileIDs {
			assert.NotContains(t, currentFileIDs, orig, "original item should not exist after replacement")
		}

		fileIDs = currentFileIDs
	})

	// copy restore

	suite.Run("copy collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Copy

		ro, _ := PrepNewTestRestoreOp(
			t,
			ctx,
			bod.St,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		deets := RunAndCheckRestore(t, ctx, &ro, mb, false)
		filtEnts := []details.Entry{}

		for _, e := range deets.Entries {
			if e.Folder == nil {
				filtEnts = append(filtEnts, e)
			}
		}

		CheckRestoreCounts(t, ctr, 0, 0, countItemsInRestore)
		assert.Len(
			t,
			filtEnts,
			countItemsInRestore,
			"every item should have been copied")

		result := FilterCollisionKeyResults(
			t,
			ctx,
			driveID,
			containerID,
			GetItemsInContainerByCollisionKeyer[api.DriveItemIDType](acd),
			collKeys)

		assert.Len(t, result, len(collKeys), "all items should have been added as copies")

		currentFileIDs, err := acd.GetItemIDsInContainer(ctx, driveID, containerID)
		require.NoError(t, err, clues.ToCore(err))

		assert.Equal(t, 2*len(fileIDs), len(currentFileIDs), "count of ids should be double from before")
		assert.Subset(t, maps.Keys(currentFileIDs), maps.Keys(fileIDs), "original item should exist after copy")
	})
}
