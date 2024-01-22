package driveish

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/cmd/sanity_test/common"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const (
	expandPermissions = "expand_permissions"
	owner             = "owner"
)

// sanitree population will grab a superset of data in the drive.
// this increases the chance that we'll run into a race collision with
// the cleanup script.  Sometimes that's okay (deleting old data that
// isn't scrutinized in the test), other times it's not.  We mark whether
// that's okay to do or not by specifying the folder that's being
// scrutinized for the test.  Any errors within that folder should cause
// a fatal exit.  Errors outside of that folder get ignored.
//
// since we're using folder names, requireNoErrorsWithinFolderName will
// work best (ie: have the fewest collisions/side-effects) if the folder
// name is very specific.  Standard sanity tests should include timestamps,
// which should help ensure that.  Be warned if you try to use it with
// a more generic name: unintended effects could occur.
func populateSanitree(
	ctx context.Context,
	ac api.Client,
	driveID, requireNoErrorsWithinFolderName string,
) *common.Sanitree[models.DriveItemable, models.DriveItemable] {
	common.Infof(ctx, "building sanitree for drive: %s", driveID)

	root, err := ac.Drives().GetRootFolder(ctx, driveID)
	if err != nil {
		common.Fatal(ctx, "getting drive root folder", err)
	}

	rootName := ptr.Val(root.GetName())

	stree := &common.Sanitree[models.DriveItemable, models.DriveItemable]{
		Self:     root,
		ID:       ptr.Val(root.GetId()),
		Name:     rootName,
		Leaves:   map[string]*common.Sanileaf[models.DriveItemable, models.DriveItemable]{},
		Children: map[string]*common.Sanitree[models.DriveItemable, models.DriveItemable]{},
	}

	recursivelyBuildTree(
		ctx,
		ac,
		driveID,
		stree.Name+"/",
		requireNoErrorsWithinFolderName,
		rootName == requireNoErrorsWithinFolderName,
		stree)

	return stree
}

func recursivelyBuildTree(
	ctx context.Context,
	ac api.Client,
	driveID, location, requireNoErrorsWithinFolderName string,
	isChildOfFolderRequiringNoErrors bool,
	stree *common.Sanitree[models.DriveItemable, models.DriveItemable],
) {
	common.Debugf(ctx, "adding: %s", location)

	children, err := ac.Drives().GetFolderChildren(ctx, driveID, stree.ID)
	if err != nil {
		if isChildOfFolderRequiringNoErrors {
			common.Fatal(ctx, "getting drive children by id", err)
		}

		common.Infof(
			ctx,
			"ignoring error getting children in directory %q because it is not within directory %q\nerror: %s\n%+v",
			location,
			requireNoErrorsWithinFolderName,
			err.Error(),
			clues.ToCore(err))

		return
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

			cannotAllowErrors := isChildOfFolderRequiringNoErrors || itemName == requireNoErrorsWithinFolderName

			branch := &common.Sanitree[models.DriveItemable, models.DriveItemable]{
				Parent: stree,
				Self:   driveItem,
				ID:     itemID,
				Name:   itemName,
				Expand: map[string]any{
					expandPermissions: permissionIn(ctx, ac, driveID, itemID, cannotAllowErrors),
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
				requireNoErrorsWithinFolderName,
				cannotAllowErrors,
				branch)
		}

		if driveItem.GetFile() != nil {
			stree.CountLeaves++
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
