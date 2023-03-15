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
	"github.com/alcionai/corso/src/pkg/fault"
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

func getDriveItemContent(
	ctx context.Context,
	srv graph.Servicer,
	driveID string,
	item models.DriveItemable,
	itemReader itemReaderFunc,
	itemClient *http.Client,
	errs *fault.Bus,
) (io.ReadCloser, error) {
	var (
		itemID   = ptr.Val(item.GetId())
		itemName = ptr.Val(item.GetName())
		el       = errs.Local()

		itemData io.ReadCloser
		err      error
	)

	// Initial try with url from delta + 2 retries
	for i := 1; i <= maxRetries; i++ {
		_, itemData, err = itemReader(ctx, itemClient, item)

		if err == nil || !graph.IsErrUnauthorized(err) {
			break
		}

		// Assume unauthorized requests are a sign of an expired jwt
		// token, and that we've overrun the available window to
		// download the actual file.  Re-downloading the item will
		// refresh that download url.
		di, diErr := getDriveItem(ctx, srv, driveID, itemID)
		if diErr != nil {
			return nil, errors.Wrap(diErr, "retrieving expired item")
		}
		item = di
	}

	// check for errors following retries
	if err != nil {
		if item.GetMalware() != nil || clues.HasLabel(err, graph.LabelsMalware) {
			logger.Ctx(ctx).With("error", err.Error(), "malware", true).Error("downloading item")
			el.AddSkip(fault.FileSkip(fault.SkipMalware, itemID, itemName, graph.MalwareInfo(item)))
		} else {
			logger.Ctx(ctx).With("error", err.Error()).Error("downloading item")
			el.AddRecoverable(clues.Stack(err).WithClues(ctx).Label(fault.LabelForceNoBackupCreation))
		}

		// return err, not el.Err(), because the lazy reader needs to communicate to
		// the data consumer that this item is unreadable, regardless of the fault state.
		return nil, err
	}

	return itemData, nil
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
		SharePoint: sharePointItemInfo(item, ptr.Val(item.GetSize())),
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
		FileName: ptr.Val(item.GetName()),
	}

	if item.GetShared() == nil {
		meta.SharingMode = SharingModeInherited
	} else {
		meta.SharingMode = SharingModeCustom
	}

	var (
		perms []UserPermission
		err   error
	)

	if meta.SharingMode == SharingModeCustom && fetchPermissions {
		perms, err = oneDriveItemPermissionInfo(ctx, service, driveID, ptr.Val(item.GetId()))
		if err != nil {
			// Keep this in an if-block because if it's not then we have a weird issue
			// of having no value in error but golang thinking it's non nil because of
			// the way interfaces work.
			err = clues.Wrap(err, "fetching item permissions")
		} else {
			meta.Permissions = perms
		}
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
		OneDrive: oneDriveItemInfo(item, ptr.Val(item.GetSize())),
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
		return nil, graph.Wrap(ctx, err, "new item download request")
	}

	//nolint:lll
	// Decorate the traffic
	// See https://learn.microsoft.com/en-us/sharepoint/dev/general-development/how-to-avoid-getting-throttled-or-blocked-in-sharepoint-online#how-to-decorate-your-http-traffic
	req.Header.Set("User-Agent", "ISV|Alcion|Corso/"+version.Version)

	resp, err := hc.Do(req)
	if err != nil {
		cerr := graph.Wrap(ctx, err, "downloading item")

		if graph.IsMalware(err) {
			cerr = cerr.Label(graph.LabelsMalware)
		}

		return nil, cerr
	}

	if (resp.StatusCode / 100) == 2 {
		return resp, nil
	}

	if graph.IsMalwareResp(context.Background(), resp) {
		return nil, clues.New("malware detected").Label(graph.LabelsMalware)
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
	var email, driveName, driveID string

	if di.GetCreatedBy() != nil && di.GetCreatedBy().GetUser() != nil {
		// User is sometimes not available when created via some
		// external applications (like backup/restore solutions)
		ed, ok := di.GetCreatedBy().GetUser().GetAdditionalData()["email"]
		if ok {
			email = *ed.(*string)
		}
	}

	if di.GetParentReference() != nil {
		driveID = ptr.Val(di.GetParentReference().GetDriveId())
		driveName = strings.TrimSpace(ptr.Val(di.GetParentReference().GetName()))
	}

	return &details.OneDriveInfo{
		Created:   ptr.Val(di.GetCreatedDateTime()),
		DriveID:   driveID,
		DriveName: driveName,
		ItemName:  ptr.Val(di.GetName()),
		ItemType:  details.OneDriveItem,
		Modified:  ptr.Val(di.GetLastModifiedDateTime()),
		Owner:     email,
		Size:      itemSize,
	}
}

// OneDriveItemPermissionInfo will fetch the permission information
// for a drive item given a drive and item id.
func oneDriveItemPermissionInfo(
	ctx context.Context,
	service graph.Servicer,
	driveID string,
	itemID string,
) ([]UserPermission, error) {
	perm, err := service.
		Client().
		DrivesById(driveID).
		ItemsById(itemID).
		Permissions().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting item metadata").With("item_id", itemID)
	}

	uperms := filterUserPermissions(ctx, perm.GetValue())

	return uperms, nil
}

func filterUserPermissions(ctx context.Context, perms []models.Permissionable) []UserPermission {
	up := []UserPermission{}

	for _, p := range perms {
		if p.GetGrantedToV2() == nil {
			// For link shares, we get permissions without a user
			// specified
			continue
		}

		gv2 := p.GetGrantedToV2()
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

		entityID := ""
		if gv2.GetUser() != nil {
			entityID = ptr.Val(gv2.GetUser().GetId())
		} else if gv2.GetGroup() != nil {
			entityID = ptr.Val(gv2.GetGroup().GetId())
		} else {
			// TODO Add appliction permissions when adding permissions for SharePoint
			// https://devblogs.microsoft.com/microsoft365dev/controlling-app-access-on-specific-sharepoint-site-collections/
			logm := logger.Ctx(ctx)
			if gv2.GetApplication() != nil {
				logm.With("application_id", ptr.Val(gv2.GetApplication().GetId()))
			}
			if gv2.GetDevice() != nil {
				logm.With("application_id", ptr.Val(gv2.GetDevice().GetId()))
			}
			logm.Warn("untracked permission")
		}

		// Technically GrantedToV2 can also contain devices, but the
		// documentation does not mention about devices in permissions
		if entityID == "" {
			// This should ideally not be hit
			continue
		}

		up = append(up, UserPermission{
			ID:         ptr.Val(p.GetId()),
			Roles:      roles,
			EntityID:   entityID,
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
	var id, driveName, driveID, weburl string

	// TODO: we rely on this info for details/restore lookups,
	// so if it's nil we have an issue, and will need an alternative
	// way to source the data.

	gsi := di.GetSharepointIds()
	if gsi != nil {
		id = ptr.Val(gsi.GetSiteId())
		weburl = ptr.Val(gsi.GetSiteUrl())

		if len(weburl) == 0 {
			weburl = constructWebURL(di.GetAdditionalData())
		}
	}

	if di.GetParentReference() != nil {
		driveID = ptr.Val(di.GetParentReference().GetDriveId())
		driveName = strings.TrimSpace(ptr.Val(di.GetParentReference().GetName()))
	}

	return &details.SharePointInfo{
		ItemType:  details.OneDriveItem,
		ItemName:  ptr.Val(di.GetName()),
		Created:   ptr.Val(di.GetCreatedDateTime()),
		Modified:  ptr.Val(di.GetLastModifiedDateTime()),
		DriveID:   driveID,
		DriveName: driveName,
		Size:      itemSize,
		Owner:     id,
		WebURL:    weburl,
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

// constructWebURL helper function for recreating the webURL
// for the originating SharePoint site. Uses additional data map
// from a models.DriveItemable that possesses a downloadURL within the map.
// Returns "" if map nil or key is not present.
func constructWebURL(adtl map[string]any) string {
	var (
		desiredKey = "@microsoft.graph.downloadUrl"
		sep        = `/_layouts`
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

	temp := strings.Split(value, sep)
	url = temp[0]

	return url
}

func setName(orig models.ItemReferenceable, driveName string) models.ItemReferenceable {
	if orig == nil {
		return nil
	}

	orig.SetName(&driveName)

	return orig
}
