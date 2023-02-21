package onedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/alcionai/clues"
	msdrives "github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/uploadsession"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	// downloadUrlKey is used to find the download URL in a
	// DriveItem response
	downloadURLKey = "@microsoft.graph.downloadUrl"
)

// generic drive item getter
func getDriveItem(
	ctx context.Context,
	srv graph.Servicer,
	driveID, itemID string,
) (models.DriveItemable, error) {
	di, err := srv.Client().DrivesById(driveID).ItemsById(itemID).Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting item")
	}

	return di, nil
}

// sharePointItemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
// TODO: Add metadata fetching to SharePoint
func sharePointItemReader(
	ctx context.Context,
	hc *http.Client,
	item models.DriveItemable,
) (details.ItemInfo, io.ReadCloser, error) {
	resp, err := downloadItem(ctx, hc, item)
	if err != nil {
		return details.ItemInfo{}, nil, errors.Wrap(err, "downloading item")
	}

	dii := details.ItemInfo{
		SharePoint: sharePointItemInfo(item, *item.GetSize()),
	}

	return dii, resp.Body, nil
}

func oneDriveItemMetaReader(
	ctx context.Context,
	service graph.Servicer,
	driveID string,
	item models.DriveItemable,
) (io.ReadCloser, int, error) {
	meta, err := oneDriveItemMetaInfo(ctx, service, driveID, item)
	if err != nil {
		return nil, 0, err
	}

	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return nil, 0, clues.Wrap(err, "marshalling json").WithClues(ctx)
	}

	return io.NopCloser(bytes.NewReader(metaJSON)), len(metaJSON), nil
}

// oneDriveItemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
func oneDriveItemReader(
	ctx context.Context,
	hc *http.Client,
	item models.DriveItemable,
) (details.ItemInfo, io.ReadCloser, error) {
	var (
		rc     io.ReadCloser
		isFile = item.GetFile() != nil
	)

	if isFile {
		resp, err := downloadItem(ctx, hc, item)
		if err != nil {
			return details.ItemInfo{}, nil, errors.Wrap(err, "downloading item")
		}

		rc = resp.Body
	}

	dii := details.ItemInfo{
		OneDrive: oneDriveItemInfo(item, *item.GetSize()),
	}

	return dii, rc, nil
}

func downloadItem(ctx context.Context, hc *http.Client, item models.DriveItemable) (*http.Response, error) {
	url, ok := item.GetAdditionalData()[downloadURLKey].(*string)
	if !ok {
		return nil, clues.New("extracting file url").With("item_id", ptr.Val(item.GetId()))
	}

	req, err := http.NewRequest(http.MethodGet, *url, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "new request")
	}

	//nolint:lll
	// Decorate the traffic
	// See https://learn.microsoft.com/en-us/sharepoint/dev/general-development/how-to-avoid-getting-throttled-or-blocked-in-sharepoint-online#how-to-decorate-your-http-traffic
	req.Header.Set("User-Agent", "ISV|Alcion|Corso/"+version.Version)

	resp, err := hc.Do(req)
	if err != nil {
		return nil, err
	}

	if (resp.StatusCode / 100) == 2 {
		return resp, nil
	}

	// upstream error checks can compare the status with
	// clues.HasLabel(err, graph.LabelStatus(http.KnownStatusCode))
	cerr := clues.Wrap(clues.New(resp.Status), "non-2xx http response").
		Label(graph.LabelStatus(resp.StatusCode))

	return resp, cerr
}

// oneDriveItemInfo will populate a details.OneDriveInfo struct
// with properties from the drive item.  ItemSize is specified
// separately for restore processes because the local itemable
// doesn't have its size value updated as a side effect of creation,
// and kiota drops any SetSize update.
func oneDriveItemInfo(di models.DriveItemable, itemSize int64) *details.OneDriveInfo {
	var email, parent string

	if di.GetCreatedBy() != nil && di.GetCreatedBy().GetUser() != nil {
		// User is sometimes not available when created via some
		// external applications (like backup/restore solutions)
		ed, ok := di.GetCreatedBy().GetUser().GetAdditionalData()["email"]
		if ok {
			email = *ed.(*string)
		}
	}

	if di.GetParentReference() != nil && di.GetParentReference().GetName() != nil {
		// EndPoint is not always populated from external apps
		parent = *di.GetParentReference().GetName()
	}

	return &details.OneDriveInfo{
		ItemType:  details.OneDriveItem,
		ItemName:  ptr.Val(di.GetName()),
		Created:   ptr.Val(di.GetCreatedDateTime()),
		Modified:  ptr.Val(di.GetLastModifiedDateTime()),
		DriveName: parent,
		Size:      itemSize,
		Owner:     email,
	}
}

// oneDriveItemMetaInfo will fetch the meta information for a drive
// item. As of now, it only adds the permissions applicable for a
// onedrive item.
func oneDriveItemMetaInfo(
	ctx context.Context, service graph.Servicer,
	driveID string, di models.DriveItemable,
) (Metadata, error) {
	perm, err := service.Client().
		DrivesById(driveID).
		ItemsById(ptr.Val(di.GetId())).
		Permissions().
		Get(ctx, nil)
	if err != nil {
		return Metadata{}, graph.Wrap(ctx, err, "getting item metadata")
	}

	uperms := filterUserPermissions(perm.GetValue())

	return Metadata{Permissions: uperms}, nil
}

func filterUserPermissions(perms []models.Permissionable) []UserPermission {
	up := []UserPermission{}

	for _, p := range perms {
		if p.GetGrantedToV2() == nil {
			// For link shares, we get permissions without a user
			// specified
			continue
		}

		roles := []string{}

		for _, r := range p.GetRoles() {
			// Skip if the only role available in owner
			if r != "owner" {
				roles = append(roles, r)
			}
		}

		if len(roles) == 0 {
			continue
		}

		up = append(up, UserPermission{
			ID:         ptr.Val(p.GetId()),
			Roles:      roles,
			Email:      *p.GetGrantedToV2().GetUser().GetAdditionalData()["email"].(*string),
			Expiration: p.GetExpirationDateTime(),
		})
	}

	return up
}

// sharePointItemInfo will populate a details.SharePointInfo struct
// with properties from the drive item.  ItemSize is specified
// separately for restore processes because the local itemable
// doesn't have its size value updated as a side effect of creation,
// and kiota drops any SetSize update.
// TODO: Update drive name during Issue #2071
func sharePointItemInfo(di models.DriveItemable, itemSize int64) *details.SharePointInfo {
	var (
		id, parent, url string
		reference       = di.GetParentReference()
	)

	// TODO: we rely on this info for details/restore lookups,
	// so if it's nil we have an issue, and will need an alternative
	// way to source the data.
	gsi := di.GetSharepointIds()
	if gsi != nil {
		if gsi.GetSiteId() != nil {
			id = *gsi.GetSiteId()
		}

		if gsi.GetSiteUrl() != nil {
			url = *gsi.GetSiteUrl()
		}
	}

	if reference != nil {
		parent = *reference.GetDriveId()

		if reference.GetName() != nil {
			// EndPoint is not always populated from external apps
			temp := *reference.GetName()
			temp = strings.TrimSpace(temp)

			if temp != "" {
				parent = temp
			}
		}
	}

	return &details.SharePointInfo{
		ItemType:  details.OneDriveItem,
		ItemName:  ptr.Val(di.GetName()),
		Created:   ptr.Val(di.GetCreatedDateTime()),
		Modified:  ptr.Val(di.GetLastModifiedDateTime()),
		DriveName: parent,
		Size:      itemSize,
		Owner:     id,
		WebURL:    url,
	}
}

// driveItemWriter is used to initialize and return an io.Writer to upload data for the specified item
// It does so by creating an upload session and using that URL to initialize an `itemWriter`
// TODO: @vkamra verify if var session is the desired input
func driveItemWriter(
	ctx context.Context,
	service graph.Servicer,
	driveID, itemID string,
	itemSize int64,
) (io.Writer, error) {
	session := msdrives.NewItemItemsItemCreateUploadSessionPostRequestBody()
	ctx = clues.Add(ctx, "upload_item_id", itemID)

	r, err := service.Client().DrivesById(driveID).ItemsById(itemID).CreateUploadSession().Post(ctx, session, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating item upload session")
	}

	logger.Ctx(ctx).Debug("created an upload session")

	url := ptr.Val(r.GetUploadUrl())

	return uploadsession.NewWriter(itemID, url, itemSize), nil
}
