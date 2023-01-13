package onedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/connector/uploadsession"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
	msdrives "github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
)

const (
	// downloadUrlKey is used to find the download URL in a
	// DriveItem response
	downloadURLKey = "@microsoft.graph.downloadUrl"
)

// sharePointItemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
func sharePointItemReader(
	ctx context.Context,
	item models.DriveItemable,
) (details.ItemInfo, io.ReadCloser, error) {
	url, ok := item.GetAdditionalData()[downloadURLKey].(*string)
	if !ok {
		return details.ItemInfo{}, nil, fmt.Errorf("failed to get url for %s", *item.GetName())
	}

	rc, err := driveItemReader(ctx, *url)
	if err != nil {
		return details.ItemInfo{}, nil, err
	}

	dii := details.ItemInfo{
		SharePoint: sharePointItemInfo(item, *item.GetSize()),
	}

	return dii, rc, nil
}

// oneDriveItemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
func oneDriveItemReader(
	ctx context.Context,
	item models.DriveItemable,
) (details.ItemInfo, io.ReadCloser, error) {
	url, ok := item.GetAdditionalData()[downloadURLKey].(*string)
	if !ok {
		return details.ItemInfo{}, nil, fmt.Errorf("failed to get url for %s", *item.GetName())
	}

	rc, err := driveItemReader(ctx, *url)
	if err != nil {
		return details.ItemInfo{}, nil, err
	}

	dii := details.ItemInfo{
		OneDrive: oneDriveItemInfo(item, *item.GetSize()),
	}

	return dii, rc, nil
}

// driveItemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
func driveItemReader(
	ctx context.Context,
	url string,
) (io.ReadCloser, error) {
	httpClient := graph.CreateHTTPClient()
	httpClient.Timeout = 0 // infinite timeout for pulling large files

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to download file from %s", url)
	}

	return resp.Body, nil
}

// oneDriveItemInfo will populate a details.OneDriveInfo struct
// with properties from the drive item.  ItemSize is specified
// separately for restore processes because the local itemable
// doesn't have its size value updated as a side effect of creation,
// and kiota drops any SetSize update.
func oneDriveItemInfo(di models.DriveItemable, itemSize int64) *details.OneDriveInfo {
	var email, parent string

	if di.GetCreatedBy().GetUser() != nil {
		// User is sometimes not available when created via some
		// external applications (like backup/restore solutions)
		ed, ok := di.GetCreatedBy().GetUser().GetAdditionalData()["email"]
		if ok {
			email = *ed.(*string)
		}
	}

	if di.GetParentReference() != nil {
		if di.GetParentReference().GetName() != nil {
			// EndPoint is not always populated from external apps
			parent = *di.GetParentReference().GetName()
		}
	}

	return &details.OneDriveInfo{
		ItemType:  details.OneDriveItem,
		ItemName:  *di.GetName(),
		Created:   *di.GetCreatedDateTime(),
		Modified:  *di.GetLastModifiedDateTime(),
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
) (io.ReadCloser, int64, error) {
	itemID := di.GetId()

	perm, err := service.Client().DrivesById(driveID).ItemsById(*itemID).Permissions().Get(ctx, nil)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "failed to get item permissions %s", *itemID)
	}

	up := []UserPermission{}

	for _, p := range perm.GetValue() {
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
			ID:         *p.GetId(),
			Roles:      roles,
			Email:      *p.GetGrantedToV2().GetUser().GetAdditionalData()["email"].(*string),
			Expiration: p.GetExpirationDateTime(),
		})
	}

	c, err := json.Marshal(Metadata{Permissions: up})
	if err != nil {
		return nil, 0, err
	}

	return io.NopCloser(bytes.NewReader(c)), int64(len(c)), nil
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
		ItemName:  *di.GetName(),
		Created:   *di.GetCreatedDateTime(),
		Modified:  *di.GetLastModifiedDateTime(),
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

	r, err := service.Client().DrivesById(driveID).ItemsById(itemID).CreateUploadSession().Post(ctx, session, nil)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"failed to create upload session for item %s. details: %s",
			itemID,
			support.ConnectorStackErrorTrace(err),
		)
	}

	url := *r.GetUploadUrl()

	logger.Ctx(ctx).Debugf("Created an upload session for item %s. URL: %s", itemID, url)

	return uploadsession.NewWriter(itemID, url, itemSize), nil
}
