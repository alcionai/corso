package driveish

import (
	"context"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const (
	expandPermissions = "expand_permissions"
	owner             = "owner"
)

func populateSanitree(
	ctx context.Context,
	ac api.Client,
	driveID string,
	startTime time.Time,
) *common.Sanitree[models.DriveItemable] {
	root, err := ac.Drives().GetRootFolder(ctx, driveID)
	if err != nil {
		common.Fatal(ctx, "getting drive root folder", err)
	}

	stree := &common.Sanitree[models.DriveItemable]{
		Self:     root,
		ID:       ptr.Val(root.GetId()),
		Name:     ptr.Val(root.GetName()),
		Leaves:   map[string]*common.Sanileaf[models.DriveItemable]{},
		Children: map[string]*common.Sanitree[models.DriveItemable]{},
	}

	recursivelyBuildTree(
		ctx,
		ac,
		driveID,
		stree,
		startTime)

	return stree
}

func recursivelyBuildTree(
	ctx context.Context,
	ac api.Client,
	driveID string,
	stree *common.Sanitree[models.DriveItemable],
	startTime time.Time,
) {
	children, err := ac.Drives().GetFolderChildren(ctx, driveID, "root")
	if err != nil {
		common.Fatal(ctx, "getting drive children by id", err)
	}

	for _, driveItem := range children {
		var (
			itemID   = ptr.Val(driveItem.GetId())
			itemName = ptr.Val(driveItem.GetName())
		)

		if driveItem.GetFolder() != nil {
			if driveItem.GetPackageEscaped() == nil {
				common.LogAndPrint(ctx, "skipped unescaped package: %s", itemName)
				continue
			}

			// currently we don't restore blank folders.
			// skip permission check for empty folders
			if ptr.Val(driveItem.GetFolder().GetChildCount()) == 0 {
				common.LogAndPrint(ctx, "skipped empty folder: %s", itemName)
				continue
			}

			stree.Children[itemName] = &common.Sanitree[models.DriveItemable]{
				Parent: stree,
				Self:   driveItem,
				ID:     itemID,
				Name:   itemName,
				Expand: map[string]any{
					expandPermissions: permissionIn(ctx, ac, driveID, itemID),
				},
			}

			recursivelyBuildTree(
				ctx,
				ac,
				driveID,
				stree,
				startTime)
		}

		if driveItem.GetFile() != nil {
			stree.Leaves[itemName] = &common.Sanileaf[models.DriveItemable]{
				Self: driveItem,
				ID:   itemID,
				Name: itemName,
				Size: ptr.Val(driveItem.GetSize()),
			}
		}
	}
}
