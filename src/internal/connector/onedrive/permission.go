package onedrive

import (
	"context"

	"github.com/alcionai/clues"
	msdrive "github.com/microsoftgraph/msgraph-sdk-go/drive"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/path"
)

func getParentMetadata(
	parentPath path.Path,
	metas map[string]metadata.Metadata,
) (metadata.Metadata, error) {
	parentMeta, ok := metas[parentPath.String()]
	if !ok {
		drivePath, err := path.ToDrivePath(parentPath)
		if err != nil {
			return metadata.Metadata{}, clues.Wrap(err, "invalid restore path")
		}

		if len(drivePath.Folders) != 0 {
			return metadata.Metadata{}, clues.Wrap(err, "computing item permissions")
		}

		parentMeta = metadata.Metadata{}
	}

	return parentMeta, nil
}

func getCollectionMetadata(
	ctx context.Context,
	drivePath *path.DrivePath,
	dc data.RestoreCollection,
	metas map[string]metadata.Metadata,
	backupVersion int,
	restorePerms bool,
) (metadata.Metadata, error) {
	if !restorePerms || backupVersion < version.OneDrive1DataAndMetaFiles {
		return metadata.Metadata{}, nil
	}

	var (
		err            error
		collectionPath = dc.FullPath()
	)

	if len(drivePath.Folders) == 0 {
		// No permissions for root folder
		return metadata.Metadata{}, nil
	}

	if backupVersion < version.OneDrive4DirIncludesPermissions {
		colMeta, err := getParentMetadata(collectionPath, metas)
		if err != nil {
			return metadata.Metadata{}, clues.Wrap(err, "collection metadata")
		}

		return colMeta, nil
	}

	// Root folder doesn't have a metadata file associated with it.
	folders := collectionPath.Folders()
	metaName := folders[len(folders)-1] + metadata.DirMetaFileSuffix

	if backupVersion >= version.OneDrive5DirMetaNoName {
		metaName = metadata.DirMetaFileSuffix
	}

	meta, err := fetchAndReadMetadata(ctx, dc, metaName)
	if err != nil {
		return metadata.Metadata{}, clues.Wrap(err, "collection metadata")
	}

	return meta, nil
}

// computeParentPermissions computes the parent permissions by
// traversing folderMetas and finding the first item with custom
// permissions. folderMetas is expected to have all the parent
// directory metas for this to work.
func computeParentPermissions(
	itemPath path.Path,
	folderMetas map[string]metadata.Metadata,
) (metadata.Metadata, error) {
	var (
		parent path.Path
		meta   metadata.Metadata

		err error
		ok  bool
	)

	parent = itemPath

	for {
		parent, err = parent.Dir()
		if err != nil {
			return metadata.Metadata{}, clues.New("getting parent")
		}

		drivePath, err := path.ToDrivePath(parent)
		if err != nil {
			return metadata.Metadata{}, clues.New("get parent path")
		}

		if len(drivePath.Folders) == 0 {
			return metadata.Metadata{}, nil
		}

		meta, ok = folderMetas[parent.String()]
		if !ok {
			return metadata.Metadata{}, clues.New("no parent meta")
		}

		if meta.SharingMode == metadata.SharingModeCustom {
			return meta, nil
		}
	}
}

// UpdatePermissions takes in the set of permission to be added and
// removed from an item to bring it to the desired state.
func UpdatePermissions(
	ctx context.Context,
	creds account.M365Config,
	service graph.Servicer,
	driveID string,
	itemID string,
	permAdded, permRemoved []metadata.Permission,
	permissionIDMappings map[string]string,
) error {
	// The ordering of the operations is important here. We first
	// remove all the removed permissions and then add the added ones.
	for _, p := range permRemoved {
		// deletes require unique http clients
		// https://github.com/alcionai/corso/issues/2707
		// this is bad citizenship, and could end up consuming a lot of
		// system resources if servicers leak client connections (sockets, etc).
		a, err := graph.CreateAdapter(creds.AzureTenantID, creds.AzureClientID, creds.AzureClientSecret)
		if err != nil {
			return graph.Wrap(ctx, err, "creating delete client")
		}

		pid, ok := permissionIDMappings[p.ID]
		if !ok {
			return clues.New("no new permission id").WithClues(ctx)
		}

		err = graph.NewService(a).
			Client().
			DrivesById(driveID).
			ItemsById(itemID).
			PermissionsById(pid).
			Delete(ctx, nil)
		if err != nil {
			return graph.Wrap(ctx, err, "removing permissions")
		}
	}

	for _, p := range permAdded {
		// We are not able to restore permissions when there are no
		// roles or for owner, this seems to be restriction in graph
		roles := []string{}

		for _, r := range p.Roles {
			if r != "owner" {
				roles = append(roles, r)
			}
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

// RestorePermissions takes in the permissions of an item, computes
// what permissions need to added and removed based on the parent
// folder metas and uses that to add/remove the necessary permissions
// on onedrive items.
func RestorePermissions(
	ctx context.Context,
	creds account.M365Config,
	service graph.Servicer,
	driveID string,
	itemID string,
	itemPath path.Path,
	meta metadata.Metadata,
	folderMetas map[string]metadata.Metadata,
	permissionIDMappings map[string]string,
) error {
	if meta.SharingMode == metadata.SharingModeInherited {
		return nil
	}

	ctx = clues.Add(ctx, "permission_item_id", itemID)

	parentPermissions, err := computeParentPermissions(itemPath, folderMetas)
	if err != nil {
		return clues.Wrap(err, "parent permissions").WithClues(ctx)
	}

	permAdded, permRemoved := metadata.DiffPermissions(parentPermissions.Permissions, meta.Permissions)

	return UpdatePermissions(ctx, creds, service, driveID, itemID, permAdded, permRemoved, permissionIDMappings)
}
