package onedrive

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	msdrives "github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
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

// sharePointItemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
func sharePointItemReader(
	hc *http.Client,
	item models.DriveItemable,
) (details.ItemInfo, io.ReadCloser, error) {
	url, ok := item.GetAdditionalData()[downloadURLKey].(*string)
	if !ok {
		return details.ItemInfo{}, nil, fmt.Errorf("failed to get url for %s", *item.GetName())
	}

	resp, err := hc.Get(*url)
	if err != nil {
		return details.ItemInfo{}, nil, err
	}

	dii := details.ItemInfo{
		SharePoint: sharePointItemInfo(item, *item.GetSize()),
	}

	return dii, resp.Body, nil
}

// oneDriveItemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
func oneDriveItemReader(
	hc *http.Client,
	item models.DriveItemable,
) (details.ItemInfo, io.ReadCloser, error) {
	url, ok := item.GetAdditionalData()[downloadURLKey].(*string)
	if !ok {
		return details.ItemInfo{}, nil, fmt.Errorf("failed to get url for %s", *item.GetName())
	}

	req, err := http.NewRequest(http.MethodGet, *url, nil)
	if err != nil {
		return details.ItemInfo{}, nil, err
	}

	// Decorate the traffic
	// See https://learn.microsoft.com/en-us/sharepoint/dev/general-development/how-to-avoid-getting-throttled-or-blocked-in-sharepoint-online#how-to-decorate-your-http-traffic
	req.Header.Set("User-Agent", "ISV|Alcion|Corso/"+version.Version)

	resp, err := hc.Do(req)
	if err != nil {
		return details.ItemInfo{}, nil, err
	}

	dii := details.ItemInfo{
		OneDrive: oneDriveItemInfo(item, *item.GetSize()),
	}

	return dii, resp.Body, nil
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
		ItemName:  *di.GetName(),
		Created:   *di.GetCreatedDateTime(),
		Modified:  *di.GetLastModifiedDateTime(),
		DriveName: parent,
		Size:      itemSize,
		Owner:     email,
	}
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
