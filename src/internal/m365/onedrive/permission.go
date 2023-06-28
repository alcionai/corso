package onedrive

import (
	"context"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

func getParentMetadata(
	parentPath path.Path,
	parentDirToMeta map[string]metadata.Metadata,
) (metadata.Metadata, error) {
	parentMeta, ok := parentDirToMeta[parentPath.String()]
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
	caches *restoreCaches,
	backupVersion int,
	restorePerms bool,
) (metadata.Metadata, error) {
	if !restorePerms || backupVersion < version.OneDrive1DataAndMetaFiles {
		return metadata.Metadata{}, nil
	}

	var (
		err      error
		fullPath = dc.FullPath()
	)

	if len(drivePath.Folders) == 0 {
		// No permissions for root folder
		return metadata.Metadata{}, nil
	}

	if backupVersion < version.OneDrive4DirIncludesPermissions {
		colMeta, err := getParentMetadata(fullPath, caches.ParentDirToMeta)
		if err != nil {
			return metadata.Metadata{}, clues.Wrap(err, "collection metadata")
		}

		return colMeta, nil
	}

	folders := fullPath.Folders()
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

// computePreviousMetadata computes the parent permissions by
// traversing parentMetas and finding the first item with custom
// permissions. parentMetas is expected to have all the parent
// directory metas for this to work.
func computePreviousMetadata(
	ctx context.Context,
	originDir path.Path,
	// map parent dir -> parent's metadata
	parentMetas map[string]metadata.Metadata,
) (metadata.Metadata, error) {
	var (
		parent path.Path
		meta   metadata.Metadata

		err error
		ok  bool
	)

	parent = originDir

	for {
		parent, err = parent.Dir()
		if err != nil {
			return metadata.Metadata{}, clues.New("getting parent").WithClues(ctx)
		}

		ictx := clues.Add(ctx, "parent_dir", parent)

		drivePath, err := path.ToDrivePath(parent)
		if err != nil {
			return metadata.Metadata{}, clues.New("transforming dir to drivePath").WithClues(ictx)
		}

		if len(drivePath.Folders) == 0 {
			return metadata.Metadata{}, nil
		}

		meta, ok = parentMetas[parent.String()]
		if !ok {
			return metadata.Metadata{}, clues.New("no metadata found for parent folder: " + parent.String()).WithClues(ictx)
		}

		if meta.SharingMode == metadata.SharingModeCustom {
			return meta, nil
		}
	}
}

type updateDeleteItemPermissioner interface {
	DeleteItemPermissioner
	UpdateItemPermissioner
}

// UpdatePermissions takes in the set of permission to be added and
// removed from an item to bring it to the desired state.
func UpdatePermissions(
	ctx context.Context,
	udip updateDeleteItemPermissioner,
	driveID string,
	itemID string,
	permAdded, permRemoved []metadata.Permission,
	oldPermIDToNewID map[string]string,
) error {
	// The ordering of the operations is important here. We first
	// remove all the removed permissions and then add the added ones.
	for _, p := range permRemoved {
		ictx := clues.Add(
			ctx,
			"permission_entity_type", p.EntityType,
			"permission_entity_id", clues.Hide(p.EntityID))

		// deletes require unique http clients
		// https://github.com/alcionai/corso/issues/2707
		// this is bad citizenship, and could end up consuming a lot of
		// system resources if servicers leak client connections (sockets, etc).

		pid, ok := oldPermIDToNewID[p.ID]
		if !ok {
			return clues.New("no new permission id").WithClues(ctx)
		}

		err := udip.DeleteItemPermission(
			ictx,
			driveID,
			itemID,
			pid)
		if err != nil {
			return clues.Stack(err)
		}
	}

	for _, p := range permAdded {
		ictx := clues.Add(
			ctx,
			"permission_entity_type", p.EntityType,
			"permission_entity_id", clues.Hide(p.EntityID))

		// We are not able to restore permissions when there are no
		// roles or for owner, this seems to be restriction in graph
		roles := []string{}

		for _, r := range p.Roles {
			if r != "owner" {
				roles = append(roles, r)
			}
		}

		// TODO: sitegroup support.  Currently errors with "One or more users could not be resolved",
		// likely due to the site group entityID consisting of a single integer (ex: 4)
		if len(roles) == 0 || p.EntityType == metadata.GV2SiteGroup {
			continue
		}

		pbody := drives.NewItemItemsItemInvitePostRequestBody()
		pbody.SetRoles(roles)

		if p.Expiration != nil {
			expiry := p.Expiration.String()
			pbody.SetExpirationDateTime(&expiry)
		}

		pbody.SetSendInvitation(ptr.To(false))
		pbody.SetRequireSignIn(ptr.To(true))

		rec := models.NewDriveRecipient()
		if len(p.EntityID) > 0 {
			rec.SetObjectId(&p.EntityID)
		} else {
			// Previous versions used to only store email for a
			// permissions. Use that if id is not found.
			rec.SetEmail(&p.Email)
		}

		pbody.SetRecipients([]models.DriveRecipientable{rec})

		newPerm, err := udip.PostItemPermissionUpdate(ictx, driveID, itemID, pbody)
		if err != nil {
			return clues.Stack(err)
		}

		oldPermIDToNewID[p.ID] = ptr.Val(newPerm.GetValue()[0].GetId())
	}

	return nil
}

type updateDeleteItemLinkSharer interface {
	DeleteItemPermissioner // Deletion logic is same as permissions
	UpdateItemLinkSharer
}

func UpdateLinkShares(
	ctx context.Context,
	upils updateDeleteItemLinkSharer,
	driveID string,
	itemID string,
	lsAdded, lsRemoved []metadata.LinkShare,
	oldLinkShareIDToNewID map[string]string,
) (bool, error) {
	// You can only delete inherited sharing links the didReset time you
	// create a sharing link which is done using
	// `retainInheritedPermissions`. We get a 204 as a response if we
	// try to delete an inherited link share, but it does not get
	// deleted.
	didReset := false

	for _, ls := range lsAdded {
		ictx := clues.Add(ctx, "link_share_id", ls.ID)

		// Links with password are not shared with a specific user
		// even when we select a particular user, plus we are not
		// able to get the password or retain the original link and
		// so restoring them makes no sense.
		if ls.HasPassword {
			continue
		}

		idens := []map[string]string{}
		entities := []string{}

		for _, iden := range ls.Entities {
			// TODO: sitegroup support.  Currently errors with "One or more users could not be resolved",
			// likely due to the site group entityID consisting of a single integer (ex: 4)
			if iden.EntityType == metadata.GV2SiteGroup {
				continue
			}

			// Using DriveRecipient seems to error out on Graph end
			idens = append(idens, map[string]string{"objectId": iden.ID})
			entities = append(entities, iden.ID)
		}

		ictx = clues.Add(ictx, "link_share_entity_ids", strings.Join(entities, ","))

		// https://learn.microsoft.com/en-us/graph/api/driveitem-createlink?view=graph-rest-beta&tabs=http
		// v1.0 version of the graph API does not support creating a
		// link without sending a notification to the user and so we
		// use the beta API. Since we use the v1.0 API, we have to
		// stuff some of the data into the AdditionalData fields as
		// the actual fields don't exist in the stable sdk.
		// Here is the data that we have to send:
		// {
		//   "type": "view",
		//   "scope": "anonymous",
		//   "password": "String",
		//   "expirationDateTime": "...",
		//   "recipients": [{"@odata.type": "microsoft.graph.driveRecipient"}],
		//   "sendNotification": true,
		//   "retainInheritedPermissions": false
		// }
		lsbody := drives.NewItemItemsItemCreateLinkPostRequestBody()
		lsbody.SetType(ptr.To(ls.Link.Type))
		lsbody.SetScope(ptr.To(ls.Link.Scope))
		lsbody.SetExpirationDateTime(ls.Expiration)

		ad := map[string]any{
			"sendNotification": false,
			"recipients":       idens,
		}
		lsbody.SetAdditionalData(ad)

		if !didReset {
			// The only way to delete any is to use this and so if
			// we have any deleted items, we can be sure that all the
			// inherited permissions would have been removed.
			lsbody.SetRetainInheritedPermissions(ptr.To(len(lsRemoved) == 0))

			// This value only effective on the first call, but lets
			// make sure to not send it on followups.
			didReset = true
		}

		newLS, err := upils.PostItemLinkShareUpdate(ictx, driveID, itemID, lsbody)
		if err != nil {
			return didReset, clues.Stack(err)
		}

		oldLinkShareIDToNewID[ls.ID] = ptr.Val(newLS.GetId())
	}

	// It is possible to have empty link shares even though we should
	// have inherited one if the user creates a link using
	// `retainInheritedPermissions` as false, but then deleted it.  We
	// can recreate this by creating an anonymous one and deleting it.
	if len(lsRemoved) > 0 && len(lsAdded) == 0 {
		lsbody := drives.NewItemItemsItemCreateLinkPostRequestBody()
		lsbody.SetType(ptr.To("view"))
		lsbody.SetScope(ptr.To("anonymous"))
		lsbody.SetRetainInheritedPermissions(ptr.To(false))

		newLS, err := upils.PostItemLinkShareUpdate(ctx, driveID, itemID, lsbody)
		if err != nil {
			return didReset, clues.Stack(err)
		}

		didReset = true

		err = upils.DeleteItemPermission(ctx, driveID, itemID, ptr.Val(newLS.GetId()))
		if err != nil {
			return didReset, clues.Stack(err)
		}
	}

	return didReset, nil
}

// RestorePermissions takes in the permissions of an item, computes
// what permissions need to added and removed based on the parent
// folder metas and uses that to add/remove the necessary permissions
// on onedrive items.
func RestorePermissions(
	ctx context.Context,
	rh RestoreHandler,
	driveID string,
	itemID string,
	itemPath path.Path,
	current metadata.Metadata,
	caches *restoreCaches,
) error {
	if current.SharingMode == metadata.SharingModeInherited {
		return nil
	}

	ctx = clues.Add(ctx, "permission_item_id", itemID)

	previous, err := computePreviousMetadata(ctx, itemPath, caches.ParentDirToMeta)
	if err != nil {
		return clues.Wrap(err, "previous metadata")
	}

	lsAdded, lsRemoved := metadata.DiffLinkShares(previous.LinkShares, current.LinkShares)

	didReset, err := UpdateLinkShares(
		ctx,
		rh,
		driveID,
		itemID,
		lsAdded,
		lsRemoved,
		caches.OldLinkShareIDToNewID)
	if err != nil {
		return clues.Wrap(err, "updating link shares")
	}

	permRemoved := []metadata.Permission{}
	permAdded := current.Permissions
	if !didReset {
		// In case we did a reset of permissions when restoring link
		// shares, we have to make sure to restore all the permissions
		// that an item has as they too will be removed.
		logger.Ctx(ctx).Debug("link share creation reset all inheritance")
		permAdded, permRemoved = metadata.DiffPermissions(previous.Permissions, current.Permissions)
	}

	err = UpdatePermissions(
		ctx,
		rh,
		driveID,
		itemID,
		permAdded,
		permRemoved,
		caches.OldPermIDToNewID)
	if err != nil {
		return clues.Wrap(err, "updating permissions")
	}

	return nil
}
