package api

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alcionai/clues"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	"github.com/microsoftgraph/msgraph-sdk-go/drive"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/sites"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/api"
	"github.com/alcionai/corso/src/pkg/logger"
)

func getValues[T any](l api.PageLinker) ([]T, error) {
	page, ok := l.(interface{ GetValue() []T })
	if !ok {
		return nil, clues.New("page does not comply with GetValue() interface").With("page_item_type", fmt.Sprintf("%T", l))
	}

	return page.GetValue(), nil
}

// max we can do is 999
const pageSize = int32(999)

type driveItemPager struct {
	gs      graph.Servicer
	driveID string
	builder *drives.ItemRootDeltaRequestBuilder
	options *drives.ItemRootDeltaRequestBuilderGetRequestConfiguration
}

func NewItemPager(
	gs graph.Servicer,
	driveID, link string,
	fields []string,
) *driveItemPager {
	pageCount := pageSize

	headers := abstractions.NewRequestHeaders()
	preferHeaderItems := []string{
		"deltashowremovedasdeleted",
		"deltatraversepermissiongaps",
		"deltashowsharingchanges",
		"hierarchicalsharing",
	}
	headers.Add("Prefer", strings.Join(preferHeaderItems, ","))

	requestConfig := &drives.ItemRootDeltaRequestBuilderGetRequestConfiguration{
		Headers: headers,
		QueryParameters: &drives.ItemRootDeltaRequestBuilderGetQueryParameters{
			Top:    &pageCount,
			Select: fields,
		},
	}

	res := &driveItemPager{
		gs:      gs,
		driveID: driveID,
		options: requestConfig,
		builder: gs.Client().DrivesById(driveID).Root().Delta(),
	}

	if len(link) > 0 {
		res.builder = drives.NewItemRootDeltaRequestBuilder(link, gs.Adapter())
	}

	return res
}

func (p *driveItemPager) GetPage(ctx context.Context) (api.DeltaPageLinker, error) {
	var (
		resp api.DeltaPageLinker
		err  error
	)

	resp, err = p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

func (p *driveItemPager) SetNext(link string) {
	p.builder = drives.NewItemRootDeltaRequestBuilder(link, p.gs.Adapter())
}

func (p *driveItemPager) Reset() {
	p.builder = p.gs.Client().DrivesById(p.driveID).Root().Delta()
}

func (p *driveItemPager) ValuesIn(l api.DeltaPageLinker) ([]models.DriveItemable, error) {
	return getValues[models.DriveItemable](l)
}

type userDrivePager struct {
	userID  string
	gs      graph.Servicer
	builder *users.ItemDrivesRequestBuilder
	options *users.ItemDrivesRequestBuilderGetRequestConfiguration
}

func NewUserDrivePager(
	gs graph.Servicer,
	userID string,
	fields []string,
) *userDrivePager {
	requestConfig := &users.ItemDrivesRequestBuilderGetRequestConfiguration{
		QueryParameters: &users.ItemDrivesRequestBuilderGetQueryParameters{
			Select: fields,
		},
	}

	res := &userDrivePager{
		userID:  userID,
		gs:      gs,
		options: requestConfig,
		builder: gs.Client().UsersById(userID).Drives(),
	}

	return res
}

type nopUserDrivePageLinker struct {
	drive models.Driveable
}

func (nl nopUserDrivePageLinker) GetOdataNextLink() *string { return nil }

func (p *userDrivePager) GetPage(ctx context.Context) (api.PageLinker, error) {
	var (
		resp api.PageLinker
		err  error
	)

	d, err := p.gs.Client().UsersById(p.userID).Drive().Get(ctx, nil)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	resp = &nopUserDrivePageLinker{drive: d}

	// TODO(keepers): turn back on when we can separate drive enumeration
	// from default drive lookup.

	// resp, err = p.builder.Get(ctx, p.options)
	// if err != nil {
	// 	return nil, graph.Stack(ctx, err)
	// }

	return resp, nil
}

func (p *userDrivePager) SetNext(link string) {
	p.builder = users.NewItemDrivesRequestBuilder(link, p.gs.Adapter())
}

func (p *userDrivePager) ValuesIn(l api.PageLinker) ([]models.Driveable, error) {
	nl, ok := l.(*nopUserDrivePageLinker)
	if !ok || nl == nil {
		return nil, clues.New(fmt.Sprintf("improper page linker struct for user drives: %T", l))
	}

	// TODO(keepers): turn back on when we can separate drive enumeration
	// from default drive lookup.

	// return getValues[models.Driveable](l)

	return []models.Driveable{nl.drive}, nil
}

type siteDrivePager struct {
	gs      graph.Servicer
	builder *sites.ItemDrivesRequestBuilder
	options *sites.ItemDrivesRequestBuilderGetRequestConfiguration
}

// NewSiteDrivePager is a constructor for creating a siteDrivePager
// fields are the associated site drive fields that are desired to be returned
// in a query.  NOTE: Fields are case-sensitive. Incorrect field settings will
// cause errors during later paging.
// Available fields: https://learn.microsoft.com/en-us/graph/api/resources/drive?view=graph-rest-1.0
func NewSiteDrivePager(
	gs graph.Servicer,
	siteID string,
	fields []string,
) *siteDrivePager {
	requestConfig := &sites.ItemDrivesRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemDrivesRequestBuilderGetQueryParameters{
			Select: fields,
		},
	}

	res := &siteDrivePager{
		gs:      gs,
		options: requestConfig,
		builder: gs.Client().SitesById(siteID).Drives(),
	}

	return res
}

func (p *siteDrivePager) GetPage(ctx context.Context) (api.PageLinker, error) {
	var (
		resp api.PageLinker
		err  error
	)

	resp, err = p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

func (p *siteDrivePager) SetNext(link string) {
	p.builder = sites.NewItemDrivesRequestBuilder(link, p.gs.Adapter())
}

func (p *siteDrivePager) ValuesIn(l api.PageLinker) ([]models.Driveable, error) {
	return getValues[models.Driveable](l)
}

// DrivePager pages through different types of drive owners
type DrivePager interface {
	GetPage(context.Context) (api.PageLinker, error)
	SetNext(nextLink string)
	ValuesIn(api.PageLinker) ([]models.Driveable, error)
}

// GetAllDrives fetches all drives for the given pager
func GetAllDrives(
	ctx context.Context,
	pager DrivePager,
	retry bool,
	maxRetryCount int,
) ([]models.Driveable, error) {
	ds := []models.Driveable{}

	if !retry {
		maxRetryCount = 0
	}

	// Loop through all pages returned by Graph API.
	for {
		var (
			err  error
			page api.PageLinker
		)

		// Retry Loop for Drive retrieval. Request can timeout
		for i := 0; i <= maxRetryCount; i++ {
			page, err = pager.GetPage(ctx)
			if err != nil {
				if clues.HasLabel(err, graph.LabelsMysiteNotFound) {
					logger.Ctx(ctx).Infof("resource owner does not have a drive")
					return make([]models.Driveable, 0), nil // no license or drives.
				}

				if graph.IsErrTimeout(err) && i < maxRetryCount {
					time.Sleep(time.Duration(3*(i+1)) * time.Second)
					continue
				}

				return nil, graph.Wrap(ctx, err, "retrieving drives")
			}

			// No error encountered, break the retry loop so we can extract results
			// and see if there's another page to fetch.
			break
		}

		tmp, err := pager.ValuesIn(page)
		if err != nil {
			return nil, graph.Wrap(ctx, err, "extracting drives from response")
		}

		ds = append(ds, tmp...)

		nextLink := ptr.Val(page.GetOdataNextLink())
		if len(nextLink) == 0 {
			break
		}

		pager.SetNext(nextLink)
	}

	logger.Ctx(ctx).Debugf("retrieved %d valid drives", len(ds))

	return ds, nil
}

// generic drive item getter
func GetDriveItem(
	ctx context.Context,
	srv graph.Servicer,
	driveID, itemID string,
) (models.DriveItemable, error) {
	di, err := srv.Client().
		DrivesById(driveID).
		ItemsById(itemID).
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting item")
	}

	return di, nil
}

func GetItemPermission(
	ctx context.Context,
	service graph.Servicer,
	driveID, itemID string,
) (models.PermissionCollectionResponseable, error) {
	perm, err := service.
		Client().
		DrivesById(driveID).
		ItemsById(itemID).
		Permissions().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting item metadata").With("item_id", itemID)
	}

	return perm, nil
}

func GetUsersDrive(
	ctx context.Context,
	srv graph.Servicer,
	user string,
) (models.Driveable, error) {
	d, err := srv.Client().
		UsersById(user).
		Drive().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting user's drive")
	}

	return d, nil
}

func GetSitesDefaultDrive(
	ctx context.Context,
	srv graph.Servicer,
	site string,
) (models.Driveable, error) {
	d, err := srv.Client().
		SitesById(site).
		Drive().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting site's drive")
	}

	return d, nil
}

func GetDriveRoot(
	ctx context.Context,
	srv graph.Servicer,
	driveID string,
) (models.DriveItemable, error) {
	root, err := srv.Client().
		DrivesById(driveID).
		Root().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting drive root")
	}

	return root, nil
}

const itemByPathRawURLFmt = "https://graph.microsoft.com/v1.0/drives/%s/items/%s:/%s"

var ErrFolderNotFound = clues.New("folder not found")

// GetFolderByName will lookup the specified folder by name within the parentFolderID folder.
func GetFolderByName(
	ctx context.Context,
	service graph.Servicer,
	driveID, parentFolderID, folder string,
) (models.DriveItemable, error) {
	// The `Children().Get()` API doesn't yet support $filter, so using that to find a folder
	// will be sub-optimal.
	// Instead, we leverage OneDrive path-based addressing -
	// https://learn.microsoft.com/en-us/graph/onedrive-addressing-driveitems#path-based-addressing
	// - which allows us to lookup an item by its path relative to the parent ID
	rawURL := fmt.Sprintf(itemByPathRawURLFmt, driveID, parentFolderID, folder)
	builder := drive.NewItemsDriveItemItemRequestBuilder(rawURL, service.Adapter())

	foundItem, err := builder.Get(ctx, nil)
	if err != nil {
		if graph.IsErrDeletedInFlight(err) {
			return nil, graph.Stack(ctx, clues.Stack(ErrFolderNotFound, err))
		}

		return nil, graph.Wrap(ctx, err, "getting folder")
	}

	// Check if the item found is a folder, fail the call if not
	if foundItem.GetFolder() == nil {
		return nil, graph.Wrap(ctx, ErrFolderNotFound, "item is not a folder")
	}

	return foundItem, nil
}

func PostItemPermissionUpdate(
	ctx context.Context,
	service graph.Servicer,
	driveID, itemID string,
	body *drive.ItemsItemInvitePostRequestBody,
) (drives.ItemItemsItemInviteResponseable, error) {
	ctx = graph.ConsumeNTokens(ctx, graph.PermissionsLC)

	itm, err := service.Client().
		DrivesById(driveID).
		ItemsById(itemID).
		Invite().
		Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "posting permissions")
	}

	return itm, nil
}
