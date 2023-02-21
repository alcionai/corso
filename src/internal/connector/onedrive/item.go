package onedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/alcionai/clues"
	msdrives "github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
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

// generic drive item getter
func getDriveItem(
	ctx context.Context,
	srv graph.Servicer,
	driveID, itemID string,
) (models.DriveItemable, error) {
	return srv.Client().DrivesById(driveID).ItemsById(itemID).Get(ctx, nil)
}

// sharePointItemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
// TODO: Add metadata fetching to SharePoint
func sharePointItemReader(
	hc *http.Client,
	item models.DriveItemable,
) (details.ItemInfo, io.ReadCloser, error) {
	resp, err := downloadItem(hc, item)
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
	fetchPermissions bool,
) (io.ReadCloser, int, error) {
	meta := Metadata{
		FileName: *item.GetName(),
	}

	perms, err := oneDriveItemPermissionInfo(ctx, service, driveID, item, fetchPermissions)
	if err != nil {
		// Keep this in an if-block because if it's not then we have a weird issue
		// of having no value in error but golang thinking it's non nil because of
		// the way interfaces work.
		err = clues.Wrap(err, "fetching item permissions")
	} else {
		meta.Permissions = perms
	}

	metaJSON, serializeErr := json.Marshal(meta)
	if serializeErr != nil {
		serializeErr = clues.Wrap(serializeErr, "serializing item metadata")

		// Need to check if err was already non-nil since it doesn't filter nil
		// values out in calls to Stack().
		if err != nil {
			err = clues.Stack(err, serializeErr)
		} else {
			err = serializeErr
		}

		return nil, 0, err
	}

	r := io.NopCloser(bytes.NewReader(metaJSON))

	return r, len(metaJSON), err
}

// oneDriveItemReader will return a io.ReadCloser for the specified item
// It crafts this by querying M365 for a download URL for the item
// and using a http client to initialize a reader
func oneDriveItemReader(
	hc *http.Client,
	item models.DriveItemable,
) (details.ItemInfo, io.ReadCloser, error) {
	var (
		rc     io.ReadCloser
		isFile = item.GetFile() != nil
	)

	if isFile {
		resp, err := downloadItem(hc, item)
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

func downloadItem(hc *http.Client, item models.DriveItemable) (*http.Response, error) {
	url, ok := item.GetAdditionalData()[downloadURLKey].(*string)
	if !ok {
		return nil, fmt.Errorf("extracting file url: file %s", *item.GetId())
	}

	req, err := http.NewRequest(http.MethodGet, *url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new request")
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

	if resp.StatusCode == http.StatusTooManyRequests {
		return resp, graph.Err429TooManyRequests
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return resp, graph.Err401Unauthorized
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return resp, graph.Err500InternalServerError
	}

	if resp.StatusCode == http.StatusServiceUnavailable {
		return resp, graph.Err503ServiceUnavailable
	}

	return resp, errors.New("non-2xx http response: " + resp.Status)
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

// oneDriveItemPermissionInfo will fetch the permission information for a drive
// item.
func oneDriveItemPermissionInfo(
	ctx context.Context,
	service graph.Servicer,
	driveID string,
	di models.DriveItemable,
	fetchPermissions bool,
) ([]UserPermission, error) {
	if !fetchPermissions {
		return nil, nil
	}

	perm, err := service.
		Client().
		DrivesById(driveID).
		ItemsById(*di.GetId()).
		Permissions().
		Get(ctx, nil)
	if err != nil {
		return nil, err
	}

	uperms := filterUserPermissions(perm.GetValue())

	return uperms, nil
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
			ID:         *p.GetId(),
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
		id = ptr.Val(gsi.GetSiteId())
		url = ptr.Val(gsi.GetSiteUrl())

		if len(url) == 0 {
			url = constructWebURL(di.GetAdditionalData())
		}
	}

	if reference != nil {
		parent = ptr.Val(reference.GetDriveId())
		temp := ptr.Val(reference.GetName())
		temp = strings.TrimSpace(temp)

		if temp != "" {
			parent = temp
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

// constructWebURL helper function for recreating the webURL
// for the originating SharePoint site. Uses additional data map
// from a models.DriveItemable that possesses a downloadURL within the map.
// Returns "" if map nil or key is not present.
func constructWebURL(adtl map[string]any) string {
	var (
		desiredKey = "@microsoft.graph.downloadUrl"
		del        = `/_layouts`
		url        string
	)

	if adtl == nil {
		return url
	}

	r := adtl[desiredKey]
	point, ok := r.(*string)

	if !ok {
		return url
	}

	value := ptr.Val(point)
	if len(value) == 0 {
		return url
	}

	temp := strings.Split(value, del)
	url = temp[0]

	return url
}
