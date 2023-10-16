package driveish

import (
	"context"

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
) *common.Sanitree[models.DriveItemable, models.DriveItemable] {
	common.Infof(ctx, "building sanitree for drive: %s", driveID)

	root, err := ac.Drives().GetRootFolder(ctx, driveID)
	if err != nil {
		common.Fatal(ctx, "getting drive root folder", err)
	}

	stree := &common.Sanitree[models.DriveItemable, models.DriveItemable]{
		Self:     root,
		ID:       ptr.Val(root.GetId()),
		Name:     ptr.Val(root.GetName()),
		Leaves:   map[string]*common.Sanileaf[models.DriveItemable, models.DriveItemable]{},
		Children: map[string]*common.Sanitree[models.DriveItemable, models.DriveItemable]{},
	}

	recursivelyBuildTree(
		ctx,
		ac,
		driveID,
		stree.Name+"/",
		stree)

	return stree
}

func recursivelyBuildTree(
	ctx context.Context,
	ac api.Client,
	driveID, location string,
	stree *common.Sanitree[models.DriveItemable, models.DriveItemable],
) {
	common.Debugf(ctx, "adding: %s", location)

	children, err := ac.Drives().GetFolderChildren(ctx, driveID, stree.ID)
	if err != nil {
		common.Fatal(ctx, "getting drive children by id", err)
	}

	for _, driveItem := range children {
		var (
			itemID   = ptr.Val(driveItem.GetId())
			itemName = ptr.Val(driveItem.GetName())
		)

		if driveItem.GetFolder() != nil {
			// currently we don't restore blank folders.
			// skip permission check for empty folders
			if ptr.Val(driveItem.GetFolder().GetChildCount()) == 0 {
				common.Infof(ctx, "skipped empty folder: %s/%s", location, itemName)
				continue
			}

			branch := &common.Sanitree[models.DriveItemable, models.DriveItemable]{
				Parent: stree,
				Self:   driveItem,
				ID:     itemID,
				Name:   itemName,
				Expand: map[string]any{
					expandPermissions: permissionIn(ctx, ac, driveID, itemID),
				},
				Leaves:   map[string]*common.Sanileaf[models.DriveItemable, models.DriveItemable]{},
				Children: map[string]*common.Sanitree[models.DriveItemable, models.DriveItemable]{},
			}

			stree.Children[itemName] = branch

			recursivelyBuildTree(
				ctx,
				ac,
				driveID,
				location+branch.Name+"/",
				branch)
		}

		if driveItem.GetFile() != nil {
			stree.Leaves[itemName] = &common.Sanileaf[models.DriveItemable, models.DriveItemable]{
				Parent: stree,
				Self:   driveItem,
				ID:     itemID,
				Name:   itemName,
				Size:   ptr.Val(driveItem.GetSize()),
			}
		}
	}
}
