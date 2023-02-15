package onedrive

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
	msdrive "github.com/microsoftgraph/msgraph-sdk-go/drive"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
)

func getParentPermissions(
	parentPath path.Path,
	parentPermissions map[string][]UserPermission,
) ([]UserPermission, error) {
	parentPerms, ok := parentPermissions[parentPath.String()]
	if !ok {
		onedrivePath, err := path.ToOneDrivePath(parentPath)
		if err != nil {
			return nil, errors.Wrap(err, "invalid restore path")
		}

		if len(onedrivePath.Folders) != 0 {
			return nil, errors.Wrap(err, "unable to compute item permissions")
		}

		parentPerms = []UserPermission{}
	}

	return parentPerms, nil
}

func getParentAndCollectionPermissions(
	drivePath *path.DrivePath,
	collectionPath path.Path,
	permissions map[string][]UserPermission,
	restorePerms bool,
) ([]UserPermission, []UserPermission, error) {
	if !restorePerms {
		return nil, nil, nil
	}

	var (
		err         error
		parentPerms []UserPermission
		colPerms    []UserPermission
	)

	// Only get parent permissions if we're not restoring the root.
	if len(drivePath.Folders) > 0 {
		parentPath, err := collectionPath.Dir()
		if err != nil {
			return nil, nil, clues.Wrap(err, "getting parent path")
		}

		parentPerms, err = getParentPermissions(parentPath, permissions)
		if err != nil {
			return nil, nil, clues.Wrap(err, "getting parent permissions")
		}
	}

	// TODO(ashmrtn): For versions after this pull the permissions from the
	// current collection with Fetch().
	colPerms, err = getParentPermissions(collectionPath, permissions)
	if err != nil {
		return nil, nil, clues.Wrap(err, "getting collection permissions")
	}

	return parentPerms, colPerms, nil
}

// createRestoreFoldersWithPermissions creates the restore folder hierarchy in
// the specified drive and returns the folder ID of the last folder entry in the
// hierarchy. Permissions are only applied to the last folder in the hierarchy.
// Passing nil for the permissions results in just creating the folder(s).
func createRestoreFoldersWithPermissions(
	ctx context.Context,
	service graph.Servicer,
	driveID string,
	restoreFolders []string,
	parentPermissions []UserPermission,
	folderPermissions []UserPermission,
	permissionIDMappings map[string]string,
) (string, map[string]string, error) {
	id, err := CreateRestoreFolders(ctx, service, driveID, restoreFolders)
	if err != nil {
		return "", permissionIDMappings, err
	}

	permissionIDMappings, err = restorePermissions(
		ctx,
		service,
		driveID,
		id,
		parentPermissions,
		folderPermissions,
		permissionIDMappings)

	return id, permissionIDMappings, err
}

// getChildPermissions is to filter out permissions present in the
// parent from the ones that are available for child. This is
// necessary as we store the nested permissions in the child. We
// cannot avoid storing the nested permissions as it is possible that
// a file in a folder can remove the nested permission that is present
// on itself.
func getChildPermissions(childPermissions, parentPermissions []UserPermission) ([]UserPermission, []UserPermission) {
	addedPermissions := []UserPermission{}
	removedPermissions := []UserPermission{}

	for _, cp := range childPermissions {
		found := false

		for _, pp := range parentPermissions {
			if cp.ID == pp.ID {
				found = true
				break
			}
		}

		if !found {
			addedPermissions = append(addedPermissions, cp)
		}
	}

	for _, pp := range parentPermissions {
		found := false

		for _, cp := range childPermissions {
			if pp.ID == cp.ID {
				found = true
				break
			}
		}

		if !found {
			removedPermissions = append(removedPermissions, pp)
		}
	}

	return addedPermissions, removedPermissions
}

// restorePermissions takes in the permissions that were added and the
// removed(ones present in parent but not in child) and adds/removes
// the necessary permissions on onedrive objects.
func restorePermissions(
	ctx context.Context,
	service graph.Servicer,
	driveID string,
	itemID string,
	parentPerms []UserPermission,
	childPerms []UserPermission,
	permissionIDMappings map[string]string,
) (map[string]string, error) {
	permAdded, permRemoved := getChildPermissions(childPerms, parentPerms)

	for _, p := range permRemoved {
		err := service.Client().DrivesById(driveID).ItemsById(itemID).
			PermissionsById(permissionIDMappings[p.ID]).Delete(ctx, nil)
		if err != nil {
			return permissionIDMappings, errors.Wrapf(
				err,
				"failed to remove permission for item %s. details: %s",
				itemID,
				support.ConnectorStackErrorTrace(err),
			)
		}
	}

	for _, p := range permAdded {
		pbody := msdrive.NewItemsItemInvitePostRequestBody()
		pbody.SetRoles(p.Roles)

		if p.Expiration != nil {
			expiry := p.Expiration.String()
			pbody.SetExpirationDateTime(&expiry)
		}

		si := false
		pbody.SetSendInvitation(&si)

		rs := true
		pbody.SetRequireSignIn(&rs)

		rec := models.NewDriveRecipient()
		rec.SetEmail(&p.Email)
		pbody.SetRecipients([]models.DriveRecipientable{rec})

		np, err := service.Client().DrivesById(driveID).ItemsById(itemID).Invite().Post(ctx, pbody, nil)
		if err != nil {
			return permissionIDMappings, errors.Wrapf(
				err,
				"failed to set permission for item %s. details: %s",
				itemID,
				support.ConnectorStackErrorTrace(err),
			)
		}

		permissionIDMappings[p.ID] = *np.GetValue()[0].GetId()
	}

	return permissionIDMappings, nil
}
