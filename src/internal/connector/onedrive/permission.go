package onedrive

import (
	"context"

	"github.com/alcionai/clues"
	msdrive "github.com/microsoftgraph/msgraph-sdk-go/drive"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/path"
)

func getParentMetadata(
	parentPath path.Path,
	metas map[string]Metadata,
) (Metadata, error) {
	parentMeta, ok := metas[parentPath.String()]
	if !ok {
		onedrivePath, err := path.ToOneDrivePath(parentPath)
		if err != nil {
			return Metadata{}, errors.Wrap(err, "invalid restore path")
		}

		if len(onedrivePath.Folders) != 0 {
			return Metadata{}, errors.Wrap(err, "computing item permissions")
		}

		parentMeta = Metadata{}
	}

	return parentMeta, nil
}

func getCollectionMetadata(
	ctx context.Context,
	drivePath *path.DrivePath,
	dc data.RestoreCollection,
	metas map[string]Metadata,
	backupVersion int,
	restorePerms bool,
) (Metadata, error) {
	if !restorePerms || backupVersion < version.OneDrive1DataAndMetaFiles {
		return Metadata{}, nil
	}

	var (
		err            error
		collectionPath = dc.FullPath()
	)

	if len(drivePath.Folders) == 0 {
		// No permissions for root folder
		return Metadata{}, nil
	}

	if backupVersion < version.OneDrive4DirIncludesPermissions {
		colMeta, err := getParentMetadata(collectionPath, metas)
		if err != nil {
			return Metadata{}, clues.Wrap(err, "collection metadata")
		}

		return colMeta, nil
	}

	// Root folder doesn't have a metadata file associated with it.
	folders := collectionPath.Folders()
	metaName := folders[len(folders)-1] + DirMetaFileSuffix

	if backupVersion >= version.OneDrive5DirMetaNoName {
		metaName = DirMetaFileSuffix
	}

	meta, err := fetchAndReadMetadata(ctx, dc, metaName)
	if err != nil {
		return Metadata{}, clues.Wrap(err, "collection metadata")
	}

	return meta, nil
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
	folderMetadata Metadata,
	permissionIDMappings map[string]string,
) (string, error) {
	id, err := CreateRestoreFolders(ctx, service, driveID, restoreFolders)
	if err != nil {
		return "", err
	}

	err = restorePermissions(
		ctx,
		service,
		driveID,
		id,
		folderMetadata,
		permissionIDMappings)

	return id, err
}

// isSame checks equality of two string slices
func isSame(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}

	for _, f := range first {
		found := false

		for _, s := range second {
			if f == s {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func diffPermissions(
	before, after []UserPermission,
	permissionIDMappings map[string]string,
) ([]UserPermission, []UserPermission) {
	var (
		added   = []UserPermission{}
		removed = []UserPermission{}
	)

	for _, cp := range after {
		found := false

		for _, pp := range before {
			if isSame(cp.Roles, pp.Roles) &&
				cp.EntityID == pp.EntityID {
				found = true
				break
			}
		}

		if !found {
			added = append(added, cp)
		}
	}

	for _, pp := range before {
		found := false

		for _, cp := range after {
			if isSame(cp.Roles, pp.Roles) &&
				cp.EntityID == pp.EntityID {
				found = true
				break
			}
		}

		if !found {
			removed = append(removed, pp)
		}
	}

	return added, removed
}

// restorePermissions takes in the permissions that were added and the
// removed(ones present in parent but not in child) and adds/removes
// the necessary permissions on onedrive objects.
func restorePermissions(
	ctx context.Context,
	service graph.Servicer,
	driveID string,
	itemID string,
	meta Metadata,
	permissionIDMappings map[string]string,
) error {
	if meta.SharingMode == SharingModeInherited {
		return nil
	}

	ctx = clues.Add(ctx, "permission_item_id", itemID)

	// TODO(meain): Compute this from the data that we have instead of fetching from graph
	currentPermissions, err := oneDriveItemPermissionInfo(ctx, service, driveID, itemID)
	if err != nil {
		return graph.Wrap(ctx, err, "fetching current permissions")
	}

	permAdded, permRemoved := diffPermissions(currentPermissions, meta.Permissions, permissionIDMappings)

	for _, p := range permRemoved {
		if len(p.Roles) == 1 && p.Roles[0] == "owner" {
			// Since we can't restore owner permissions, this will be
			// original owner which we don't/can't delete.
			continue
		}

		err := service.Client().
			DrivesById(driveID).
			ItemsById(itemID).
			PermissionsById(p.ID).
			Delete(ctx, nil)
		if err != nil {
			return graph.Wrap(ctx, err, "removing permissions")
		}
	}

	for _, p := range permAdded {
		// We are not able to restore permissions when there are no
		// roles or for owner, this seems to be restriction in graph
		roles := p.Roles
		for _, r := range roles {
			if r == "owner" {
				continue
			}

			roles = append(roles, r)
		}

		if len(roles) == 0 {
			continue
		}

		pbody := msdrive.NewItemsItemInvitePostRequestBody()
		pbody.SetRoles(roles)

		if p.Expiration != nil {
			expiry := p.Expiration.String()
			pbody.SetExpirationDateTime(&expiry)
		}

		si := false
		pbody.SetSendInvitation(&si)

		rs := true
		pbody.SetRequireSignIn(&rs)

		rec := models.NewDriveRecipient()
		if p.EntityID != "" {
			rec.SetObjectId(&p.EntityID)
		} else {
			// Previous versions used to only store email for a
			// permissions. Use that if id is not found.
			rec.SetEmail(&p.Email)
		}

		pbody.SetRecipients([]models.DriveRecipientable{rec})

		np, err := service.Client().DrivesById(driveID).ItemsById(itemID).Invite().Post(ctx, pbody, nil)
		if err != nil {
			return graph.Wrap(ctx, err, "setting permissions")
		}

		permissionIDMappings[p.ID] = ptr.Val(np.GetValue()[0].GetId())
	}

	return nil
}
