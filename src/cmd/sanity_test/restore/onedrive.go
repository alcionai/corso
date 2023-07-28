package restore

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alcionai/clues"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/cmd/sanity_test/utils"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	owner = "owner"
)

type permissionInfo struct {
	entityID string
	roles    []string
}

func CheckOneDriveRestoration(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	userID, folderName, dataFolder string,
	startTime time.Time,
) {
	drive, err := client.
		Users().
		ByUserId(userID).
		Drive().
		Get(ctx, nil)
	if err != nil {
		utils.Fatal(ctx, "getting the drive:", err)
	}

	checkDriveRestoration(
		ctx,
		client,
		path.OneDriveService,
		folderName,
		ptr.Val(drive.GetId()),
		ptr.Val(drive.GetName()),
		dataFolder,
		startTime,
		false)
}

func checkDriveRestoration(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	service path.ServiceType,
	folderName,
	driveID,
	driveName,
	dataFolder string,
	startTime time.Time,
	skipPermissionTest bool,
) {
	var (
		// map itemID -> item size
		fileSizes = make(map[string]int64)
		// map itemID -> permission id -> []permission roles
		folderPermissions         = make(map[string][]permissionInfo)
		restoreFile               = make(map[string]int64)
		restoredFolderPermissions = make(map[string][]permissionInfo)
	)

	var restoreFolderID string

	ctx = clues.Add(ctx, "drive_id", driveID, "drive_name", driveName)

	response, err := client.
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId("root").
		Children().
		Get(ctx, nil)
	if err != nil {
		utils.Fatal(ctx, "getting drive by id", err)
	}

	for _, driveItem := range response.GetValue() {
		var (
			itemID   = ptr.Val(driveItem.GetId())
			itemName = ptr.Val(driveItem.GetName())
		)

		if itemName == folderName {
			restoreFolderID = itemID
			continue
		}

		if itemName != dataFolder {
			utils.LogAndPrint(ctx, "test data for folder: %s", dataFolder)
			continue
		}

		// if it's a file check the size
		if driveItem.GetFile() != nil {
			fileSizes[itemName] = ptr.Val(driveItem.GetSize())
		}

		if driveItem.GetFolder() == nil && driveItem.GetPackage() == nil {
			continue
		}

		// currently we don't restore blank folders.
		// skip permission check for empty folders
		if ptr.Val(driveItem.GetFolder().GetChildCount()) == 0 {
			utils.LogAndPrint(ctx, "skipped empty folder: %s", itemName)
			continue
		}

		folderPermissions[itemName] = permissionIn(ctx, client, driveID, itemID)
		getOneDriveChildFolder(ctx, client, driveID, itemID, itemName, fileSizes, folderPermissions, startTime)
	}

	getRestoredDrive(ctx, client, driveID, restoreFolderID, restoreFile, restoredFolderPermissions, startTime)

	checkRestoredDriveItemPermissions(
		ctx,
		service,
		skipPermissionTest,
		folderPermissions,
		restoredFolderPermissions)

	for fileName, expected := range fileSizes {
		utils.LogAndPrint(ctx, "checking for file: %s", fileName)

		got := restoreFile[fileName]

		utils.Assert(
			ctx,
			func() bool { return expected == got },
			fmt.Sprintf("different file size: %s", fileName),
			expected,
			got)
	}

	fmt.Println("Success")
}

func checkRestoredDriveItemPermissions(
	ctx context.Context,
	service path.ServiceType,
	skip bool,
	folderPermissions map[string][]permissionInfo,
	restoredFolderPermissions map[string][]permissionInfo,
) {
	if skip {
		return
	}

	/**
		TODO: replace this check with testElementsMatch
		from internal/connecter/graph_connector_helper_test.go
	**/

	for folderName, permissions := range folderPermissions {
		utils.LogAndPrint(ctx, "checking for folder: %s", folderName)

		restoreFolderPerm := restoredFolderPermissions[folderName]

		if len(permissions) < 1 {
			utils.LogAndPrint(ctx, "no permissions found in: %s", folderName)
			continue
		}

		permCheck := func() bool { return len(permissions) == len(restoreFolderPerm) }

		if service == path.SharePointService {
			permCheck = func() bool { return len(permissions) <= len(restoreFolderPerm) }
		}

		utils.Assert(
			ctx,
			permCheck,
			fmt.Sprintf("wrong number of restored permissions: %s", folderName),
			permissions,
			restoreFolderPerm)

		for _, perm := range permissions {
			eqID := func(pi permissionInfo) bool { return strings.EqualFold(pi.entityID, perm.entityID) }
			i := slices.IndexFunc(restoreFolderPerm, eqID)

			utils.Assert(
				ctx,
				func() bool { return i >= 0 },
				fmt.Sprintf("permission was restored in: %s", folderName),
				perm.entityID,
				restoreFolderPerm)

			// permissions should be sorted, so a by-index comparison works
			restored := restoreFolderPerm[i]

			utils.Assert(
				ctx,
				func() bool { return slices.Equal(perm.roles, restored.roles) },
				fmt.Sprintf("different roles restored: %s", folderName),
				perm.roles,
				restored.roles)
		}
	}
}

func getOneDriveChildFolder(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	driveID, itemID, parentName string,
	fileSizes map[string]int64,
	folderPermission map[string][]permissionInfo,
	startTime time.Time,
) {
	response, err := client.Drives().ByDriveId(driveID).Items().ByDriveItemId(itemID).Children().Get(ctx, nil)
	if err != nil {
		utils.Fatal(ctx, "getting child folder", err)
	}

	for _, driveItem := range response.GetValue() {
		var (
			itemID   = ptr.Val(driveItem.GetId())
			itemName = ptr.Val(driveItem.GetName())
			fullName = parentName + "/" + itemName
		)

		folderTime, hasTime := utils.MustGetTimeFromName(ctx, itemName)
		if !utils.IsWithinTimeBound(ctx, startTime, folderTime, hasTime) {
			continue
		}

		// if it's a file check the size
		if driveItem.GetFile() != nil {
			fileSizes[fullName] = ptr.Val(driveItem.GetSize())
		}

		if driveItem.GetFolder() == nil && driveItem.GetPackage() == nil {
			continue
		}

		// currently we don't restore blank folders.
		// skip permission check for empty folders
		if ptr.Val(driveItem.GetFolder().GetChildCount()) == 0 {
			utils.LogAndPrint(ctx, "skipped empty folder: %s", fullName)

			continue
		}

		folderPermission[fullName] = permissionIn(ctx, client, driveID, itemID)
		getOneDriveChildFolder(ctx, client, driveID, itemID, fullName, fileSizes, folderPermission, startTime)
	}
}

func getRestoredDrive(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	driveID, restoreFolderID string,
	restoreFile map[string]int64,
	restoreFolder map[string][]permissionInfo,
	startTime time.Time,
) {
	restored, err := client.
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(restoreFolderID).
		Children().
		Get(ctx, nil)
	if err != nil {
		utils.Fatal(ctx, "getting child folder", err)
	}

	for _, item := range restored.GetValue() {
		var (
			itemID   = ptr.Val(item.GetId())
			itemName = ptr.Val(item.GetName())
			itemSize = ptr.Val(item.GetSize())
		)

		if item.GetFile() != nil {
			restoreFile[itemName] = itemSize
			continue
		}

		if item.GetFolder() == nil && item.GetPackage() == nil {
			continue
		}

		restoreFolder[itemName] = permissionIn(ctx, client, driveID, itemID)
		getOneDriveChildFolder(ctx, client, driveID, itemID, itemName, restoreFile, restoreFolder, startTime)
	}
}

// ---------------------------------------------------------------------------
// permission helpers
// ---------------------------------------------------------------------------

func permissionIn(
	ctx context.Context,
	client *msgraphsdk.GraphServiceClient,
	driveID, itemID string,
) []permissionInfo {
	pi := []permissionInfo{}

	pcr, err := client.
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(itemID).
		Permissions().
		Get(ctx, nil)
	if err != nil {
		utils.Fatal(ctx, "getting permission", err)
	}

	for _, perm := range pcr.GetValue() {
		if perm.GetGrantedToV2() == nil {
			continue
		}

		var (
			gv2      = perm.GetGrantedToV2()
			permInfo = permissionInfo{}
			entityID string
		)

		// TODO: replace with filterUserPermissions in onedrive item.go
		if gv2.GetUser() != nil {
			entityID = ptr.Val(gv2.GetUser().GetId())
		} else if gv2.GetGroup() != nil {
			entityID = ptr.Val(gv2.GetGroup().GetId())
		}

		roles := utils.FilterSlice(perm.GetRoles(), owner)
		for _, role := range roles {
			permInfo.entityID = entityID
			permInfo.roles = append(permInfo.roles, role)
		}

		if len(roles) > 0 {
			slices.Sort(permInfo.roles)
			pi = append(pi, permInfo)
		}
	}

	return pi
}
