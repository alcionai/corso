package onedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// downloadUrlKeys is used to find the download URL in a DriveItem response.
var downloadURLKeys = []string{
	"@microsoft.graph.downloadUrl",
	"@content.downloadUrl",
}

type augmenterAndGetter interface {
	ItemInfoAugmenter
	api.Getter
}

func downloadItem(
	ctx context.Context,
	aag augmenterAndGetter,
	item models.DriveItemable,
) (details.ItemInfo, io.ReadCloser, error) {
	var (
		rc     io.ReadCloser
		isFile = item.GetFile() != nil
	)

	if isFile {
		var url string

		for _, key := range downloadURLKeys {
			tmp, ok := item.GetAdditionalData()[key].(*string)
			if ok {
				url = ptr.Val(tmp)
				break
			}
		}

		if len(url) == 0 {
			return details.ItemInfo{}, nil, clues.New("extracting file url")
		}

		resp, err := aag.Get(ctx, url, nil)
		if err != nil {
			return details.ItemInfo{}, nil, clues.Wrap(err, "getting item")
		}

		if graph.IsMalwareResp(ctx, resp) {
			return details.ItemInfo{}, nil, clues.New("malware detected").Label(graph.LabelsMalware)
		}

		if (resp.StatusCode / 100) != 2 {
			// upstream error checks can compare the status with
			// clues.HasLabel(err, graph.LabelStatus(http.KnownStatusCode))
			return details.ItemInfo{}, nil, clues.
				Wrap(clues.New(resp.Status), "non-2xx http response").
				Label(graph.LabelStatus(resp.StatusCode))
		}

		rc = resp.Body
	}

	dii := aag.AugmentItemInfo(
		details.ItemInfo{},
		item,
		ptr.Val(item.GetSize()),
		nil)

	return dii, rc, nil
}

func downloadItemMeta(
	ctx context.Context,
	gip GetItemPermissioner,
	driveID string,
	item models.DriveItemable,
) (io.ReadCloser, int, error) {
	meta := metadata.Metadata{FileName: ptr.Val(item.GetName())}

	if item.GetShared() == nil {
		meta.SharingMode = metadata.SharingModeInherited
	} else {
		meta.SharingMode = metadata.SharingModeCustom
	}

	if meta.SharingMode == metadata.SharingModeCustom {
		perm, err := gip.GetItemPermission(ctx, driveID, ptr.Val(item.GetId()))
		if err != nil {
			return nil, 0, err
		}

		meta.Permissions = filterUserPermissions(ctx, perm.GetValue())
	}

	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return nil, 0, clues.Wrap(err, "serializing item metadata").WithClues(ctx)
	}

	return io.NopCloser(bytes.NewReader(metaJSON)), len(metaJSON), nil
}

func filterUserPermissions(ctx context.Context, perms []models.Permissionable) []metadata.Permission {
	up := []metadata.Permission{}

	for _, p := range perms {
		if p.GetGrantedToV2() == nil {
			// For link shares, we get permissions without a user
			// specified
			continue
		}

		var (
			// Below are the mapping from roles to "Advanced" permissions
			// screen entries:
			//
			// owner - Full Control
			// write - Design | Edit | Contribute (no difference in /permissions api)
			// read  - Read
			// empty - Restricted View
			//
			// helpful docs:
			// https://devblogs.microsoft.com/microsoft365dev/controlling-app-access-on-specific-sharepoint-site-collections/
			roles    = p.GetRoles()
			gv2      = p.GetGrantedToV2()
			entityID string
			gv2t     metadata.GV2Type
		)

		switch true {
		case gv2.GetUser() != nil:
			gv2t = metadata.GV2User
			entityID = ptr.Val(gv2.GetUser().GetId())
		case gv2.GetSiteUser() != nil:
			gv2t = metadata.GV2SiteUser
			entityID = ptr.Val(gv2.GetSiteUser().GetId())
		case gv2.GetGroup() != nil:
			gv2t = metadata.GV2Group
			entityID = ptr.Val(gv2.GetGroup().GetId())
		case gv2.GetSiteGroup() != nil:
			gv2t = metadata.GV2SiteGroup
			entityID = ptr.Val(gv2.GetSiteGroup().GetId())
		case gv2.GetApplication() != nil:
			gv2t = metadata.GV2App
			entityID = ptr.Val(gv2.GetApplication().GetId())
		case gv2.GetDevice() != nil:
			gv2t = metadata.GV2Device
			entityID = ptr.Val(gv2.GetDevice().GetId())
		default:
			logger.Ctx(ctx).Info("untracked permission")
		}

		// Technically GrantedToV2 can also contain devices, but the
		// documentation does not mention about devices in permissions
		if entityID == "" {
			// This should ideally not be hit
			continue
		}

		up = append(up, metadata.Permission{
			ID:         ptr.Val(p.GetId()),
			Roles:      roles,
			EntityID:   entityID,
			EntityType: gv2t,
			Expiration: p.GetExpirationDateTime(),
		})
	}

	return up
}

// driveItemWriter is used to initialize and return an io.Writer to upload data for the specified item
// It does so by creating an upload session and using that URL to initialize an `itemWriter`
// TODO: @vkamra verify if var session is the desired input
func driveItemWriter(
	ctx context.Context,
	rh RestoreHandler,
	driveID, itemID string,
	itemSize int64,
) (io.Writer, string, error) {
	ctx = clues.Add(ctx, "upload_item_id", itemID)

	r, err := rh.ItemPoster().PostItem(ctx, driveID, itemID)
	if err != nil {
		return nil, "", clues.Stack(err)
	}

	iw := graph.NewLargeItemWriter(itemID, ptr.Val(r.GetUploadUrl()), itemSize)

	return iw, ptr.Val(r.GetUploadUrl()), nil
}

func setName(orig models.ItemReferenceable, driveName string) models.ItemReferenceable {
	if orig == nil {
		return nil
	}

	orig.SetName(&driveName)

	return orig
}
