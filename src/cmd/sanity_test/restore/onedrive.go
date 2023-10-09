package restore

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alcionai/clues"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const (
	owner = "owner"
)

func CheckOneDriveRestoration(
	ctx context.Context,
	ac api.Client,
	envs common.Envs,
) {
	drive, err := ac.Users().GetDefaultDrive(ctx, envs.UserID)
	if err != nil {
		common.Fatal(ctx, "getting the drive:", err)
	}

	checkDriveRestoration(
		ctx,
		ac,
		path.OneDriveService,
		envs.FolderName,
		ptr.Val(drive.GetId()),
		ptr.Val(drive.GetName()),
		envs.DataFolder,
		envs.StartTime,
		false)
}

func checkDriveRestoration(
	ctx context.Context,
	ac api.Client,
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
		folderPermissions         = make(map[string][]common.PermissionInfo)
		restoreFile               = make(map[string]int64)
		restoredFolderPermissions = make(map[string][]common.PermissionInfo)
	)

	ctx = clues.Add(ctx, "drive_id", driveID, "drive_name", driveName)

	restoreFolderID := PopulateDriveDetails(
		ctx,
		ac,
		driveID,
		folderName,
		dataFolder,
		fileSizes,
		folderPermissions,
		startTime)

	getRestoredDrive(
		ctx,
		ac,
		driveID,
		restoreFolderID,
		restoreFile,
		restoredFolderPermissions,
		startTime)

	checkRestoredDriveItemPermissions(
		ctx,
		service,
		skipPermissionTest,
		folderPermissions,
		restoredFolderPermissions)

	for fileName, expected := range fileSizes {
		common.LogAndPrint(ctx, "checking for file: %s", fileName)

		got := restoreFile[fileName]

		common.Assert(
			ctx,
			func() bool { return expected == got },
			fmt.Sprintf("different file size: %s", fileName),
			expected,
			got)
	}

	fmt.Println("Success")
}

func PopulateDriveDetails(
	ctx context.Context,
	ac api.Client,
	driveID, folderName, dataFolder string,
	fileSizes map[string]int64,
	folderPermissions map[string][]common.PermissionInfo,
	startTime time.Time,
) string {
	var restoreFolderID string

	children, err := ac.Drives().GetFolderChildren(ctx, driveID, "root")
	if err != nil {
		common.Fatal(ctx, "getting drive by id", err)
	}

	for _, driveItem := range children {
		var (
			itemID   = ptr.Val(driveItem.GetId())
			itemName = ptr.Val(driveItem.GetName())
		)

		if itemName == folderName {
			restoreFolderID = itemID
			continue
		}

		if itemName != dataFolder {
			common.LogAndPrint(ctx, "test data for folder: %s", dataFolder)
			continue
		}

		// if it's a file check the size
		if driveItem.GetFile() != nil {
			fileSizes[itemName] = ptr.Val(driveItem.GetSize())
		}

		if driveItem.GetFolder() == nil && driveItem.GetPackageEscaped() == nil {
			continue
		}

		// currently we don't restore blank folders.
		// skip permission check for empty folders
		if ptr.Val(driveItem.GetFolder().GetChildCount()) == 0 {
			common.LogAndPrint(ctx, "skipped empty folder: %s", itemName)
			continue
		}

		folderPermissions[itemName] = permissionIn(ctx, ac, driveID, itemID)

		getOneDriveChildFolder(
			ctx,
			ac,
			driveID,
			itemID,
			itemName,
			fileSizes,
			folderPermissions,
			startTime)
	}

	return restoreFolderID
}

func checkRestoredDriveItemPermissions(
	ctx context.Context,
	service path.ServiceType,
	skip bool,
	folderPermissions map[string][]common.PermissionInfo,
	restoredFolderPermissions map[string][]common.PermissionInfo,
) {
	if skip {
		return
	}

	/**
		TODO: replace this check with testElementsMatch
		from internal/connecter/graph_connector_helper_test.go
	**/

	for folderName, permissions := range folderPermissions {
		common.LogAndPrint(ctx, "checking for folder: %s", folderName)

		restoreFolderPerm := restoredFolderPermissions[folderName]

		if len(permissions) < 1 {
			common.LogAndPrint(ctx, "no permissions found in: %s", folderName)
			continue
		}

		permCheck := func() bool { return len(permissions) == len(restoreFolderPerm) }

		if service == path.SharePointService {
			permCheck = func() bool { return len(permissions) <= len(restoreFolderPerm) }
		}

		common.Assert(
			ctx,
			permCheck,
			fmt.Sprintf("wrong number of restored permissions: %s", folderName),
			permissions,
			restoreFolderPerm)

		for _, perm := range permissions {
			eqID := func(pi common.PermissionInfo) bool { return strings.EqualFold(pi.EntityID, perm.EntityID) }
			i := slices.IndexFunc(restoreFolderPerm, eqID)

			common.Assert(
				ctx,
				func() bool { return i >= 0 },
				fmt.Sprintf("permission was restored in: %s", folderName),
				perm.EntityID,
				restoreFolderPerm)

			// permissions should be sorted, so a by-index comparison works
			restored := restoreFolderPerm[i]

			common.Assert(
				ctx,
				func() bool { return slices.Equal(perm.Roles, restored.Roles) },
				fmt.Sprintf("different roles restored: %s", folderName),
				perm.Roles,
				restored.Roles)
		}
	}
}

func getOneDriveChildFolder(
	ctx context.Context,
	ac api.Client,
	driveID, itemID, parentName string,
	fileSizes map[string]int64,
	folderPermission map[string][]common.PermissionInfo,
	startTime time.Time,
) {
	children, err := ac.Drives().GetFolderChildren(ctx, driveID, itemID)
	if err != nil {
		common.Fatal(ctx, "getting child folder", err)
	}

	for _, driveItem := range children {
		var (
			itemID   = ptr.Val(driveItem.GetId())
			itemName = ptr.Val(driveItem.GetName())
			fullName = parentName + "/" + itemName
		)

		folderTime, hasTime := common.MustGetTimeFromName(ctx, itemName)
		if !common.IsWithinTimeBound(ctx, startTime, folderTime, hasTime) {
			continue
		}

		// if it's a file check the size
		if driveItem.GetFile() != nil {
			fileSizes[fullName] = ptr.Val(driveItem.GetSize())
		}

		if driveItem.GetFolder() == nil && driveItem.GetPackageEscaped() == nil {
			continue
		}

		// currently we don't restore blank folders.
		// skip permission check for empty folders
		if ptr.Val(driveItem.GetFolder().GetChildCount()) == 0 {
			common.LogAndPrint(ctx, "skipped empty folder: %s", fullName)

			continue
		}

		folderPermission[fullName] = permissionIn(ctx, ac, driveID, itemID)
		getOneDriveChildFolder(
			ctx,
			ac,
			driveID,
			itemID,
			fullName,
			fileSizes,
			folderPermission,
			startTime)
	}
}

func getRestoredDrive(
	ctx context.Context,
	ac api.Client,
	driveID, restoreFolderID string,
	restoreFile map[string]int64,
	restoreFolder map[string][]common.PermissionInfo,
	startTime time.Time,
) {
	children, err := ac.Drives().GetFolderChildren(ctx, driveID, restoreFolderID)
	if err != nil {
		common.Fatal(ctx, "getting child folder", err)
	}

	for _, item := range children {
		var (
			itemID   = ptr.Val(item.GetId())
			itemName = ptr.Val(item.GetName())
			itemSize = ptr.Val(item.GetSize())
		)

		if item.GetFile() != nil {
			restoreFile[itemName] = itemSize
			continue
		}

		if item.GetFolder() == nil && item.GetPackageEscaped() == nil {
			continue
		}

		restoreFolder[itemName] = permissionIn(ctx, ac, driveID, itemID)
		getOneDriveChildFolder(
			ctx,
			ac,
			driveID,
			itemID,
			itemName,
			restoreFile,
			restoreFolder,
			startTime)
	}
}

// ---------------------------------------------------------------------------
// permission helpers
// ---------------------------------------------------------------------------

func permissionIn(
	ctx context.Context,
	ac api.Client,
	driveID, itemID string,
) []common.PermissionInfo {
	pi := []common.PermissionInfo{}

	pcr, err := ac.Drives().GetItemPermission(ctx, driveID, itemID)
	if err != nil {
		common.Fatal(ctx, "getting permission", err)
	}

	for _, perm := range pcr.GetValue() {
		if perm.GetGrantedToV2() == nil {
			continue
		}

		var (
			gv2      = perm.GetGrantedToV2()
			permInfo = common.PermissionInfo{}
			entityID string
		)

		// TODO: replace with filterUserPermissions in onedrive item.go
		if gv2.GetUser() != nil {
			entityID = ptr.Val(gv2.GetUser().GetId())
		} else if gv2.GetGroup() != nil {
			entityID = ptr.Val(gv2.GetGroup().GetId())
		}

		roles := common.FilterSlice(perm.GetRoles(), owner)
		for _, role := range roles {
			permInfo.EntityID = entityID
			permInfo.Roles = append(permInfo.Roles, role)
		}

		if len(roles) > 0 {
			slices.Sort(permInfo.Roles)
			pi = append(pi, permInfo)
		}
	}

	return pi
}
