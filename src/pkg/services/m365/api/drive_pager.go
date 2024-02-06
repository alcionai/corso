package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/sites"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	onedrive "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

type DriveItemIDType struct {
	ItemID   string
	IsFolder bool
}

// ---------------------------------------------------------------------------
// non-delta item pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.DriveItemable] = &driveItemPageCtrl{}

type driveItemPageCtrl struct {
	gs      graph.Servicer
	builder *drives.ItemItemsItemChildrenRequestBuilder
	options *drives.ItemItemsItemChildrenRequestBuilderGetRequestConfiguration
}

func (c Drives) NewDriveItemPager(
	driveID, containerID string,
	selectProps ...string,
) pagers.NonDeltaHandler[models.DriveItemable] {
	options := &drives.ItemItemsItemChildrenRequestBuilderGetRequestConfiguration{
		QueryParameters: &drives.ItemItemsItemChildrenRequestBuilderGetQueryParameters{
			Top: ptr.To(maxNonDeltaPageSize),
		},
	}

	if len(selectProps) > 0 {
		options.QueryParameters.Select = selectProps
	}

	builder := c.Stable.
		Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(containerID).
		Children()

	return &driveItemPageCtrl{c.Stable, builder, options}
}

func (p *driveItemPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.DriveItemable], error) {
	page, err := p.builder.Get(ctx, p.options)
	return page, clues.Stack(err).OrNil()
}

func (p *driveItemPageCtrl) SetNextLink(nextLink string) {
	p.builder = drives.NewItemItemsItemChildrenRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *driveItemPageCtrl) ValidModTimes() bool {
	return true
}

func (c Drives) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	driveID, containerID string,
) (map[string]DriveItemIDType, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewDriveItemPager(driveID, containerID, idAnd("name")...)

	items, err := pagers.BatchEnumerateItems(ctx, pager)
	if err != nil {
		return nil, clues.Wrap(err, "enumerating drive items")
	}

	m := map[string]DriveItemIDType{}

	for _, item := range items {
		m[DriveItemCollisionKey(item)] = DriveItemIDType{
			ItemID:   ptr.Val(item.GetId()),
			IsFolder: item.GetFolder() != nil,
		}
	}

	return m, nil
}

func (c Drives) GetItemIDsInContainer(
	ctx context.Context,
	driveID, containerID string,
) (map[string]DriveItemIDType, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewDriveItemPager(driveID, containerID, idAnd("file", "folder")...)

	items, err := pagers.BatchEnumerateItems(ctx, pager)
	if err != nil {
		return nil, clues.Wrap(err, "enumerating contacts")
	}

	m := map[string]DriveItemIDType{}

	for _, item := range items {
		m[ptr.Val(item.GetId())] = DriveItemIDType{
			ItemID:   ptr.Val(item.GetId()),
			IsFolder: item.GetFolder() != nil,
		}
	}

	return m, nil
}

// ---------------------------------------------------------------------------
// delta item pager
// ---------------------------------------------------------------------------

var _ pagers.DeltaHandler[models.DriveItemable] = &DriveItemDeltaPageCtrl{}

type DriveItemDeltaPageCtrl struct {
	gs      graph.Servicer
	driveID string
	builder *drives.ItemItemsItemDeltaRequestBuilder
	options *drives.ItemItemsItemDeltaRequestBuilderGetRequestConfiguration
}

func (c Drives) newDriveItemDeltaPager(
	driveID, prevDeltaLink string,
	cc CallConfig,
) *DriveItemDeltaPageCtrl {
	preferHeaderItems := []string{
		"deltashowremovedasdeleted",
		"deltatraversepermissiongaps",
		"deltashowsharingchanges",
		"hierarchicalsharing",
	}

	options := &drives.ItemItemsItemDeltaRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(preferHeaderItems...),
		QueryParameters: &drives.ItemItemsItemDeltaRequestBuilderGetQueryParameters{
			Top: ptr.To(maxDeltaPageSize),
		},
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	builder := c.Stable.
		Client().
		Drives().
		ByDriveId(driveID).
		Items().
		ByDriveItemId(onedrive.RootID).
		Delta()

	if len(prevDeltaLink) > 0 {
		builder = drives.NewItemItemsItemDeltaRequestBuilder(prevDeltaLink, c.Stable.Adapter())
	}

	res := &DriveItemDeltaPageCtrl{
		gs:      c.Stable,
		driveID: driveID,
		options: options,
		builder: builder,
	}

	return res
}

func (p *DriveItemDeltaPageCtrl) GetPage(
	ctx context.Context,
) (pagers.DeltaLinkValuer[models.DriveItemable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, clues.Stack(err).OrNil()
}

func (p *DriveItemDeltaPageCtrl) SetNextLink(link string) {
	p.builder = drives.NewItemItemsItemDeltaRequestBuilder(link, p.gs.Adapter())
}

func (p *DriveItemDeltaPageCtrl) Reset(context.Context) {
	p.builder = p.gs.Client().
		Drives().
		ByDriveId(p.driveID).
		Items().
		ByDriveItemId(onedrive.RootID).
		Delta()
}

func (p *DriveItemDeltaPageCtrl) ValidModTimes() bool {
	return true
}

// EnumerateDriveItems will enumerate all items in the specified drive and stream them page
// by page, along with the delta update and any errors, to the provided channel.
func (c Drives) EnumerateDriveItemsDelta(
	ctx context.Context,
	driveID string,
	prevDeltaLink string,
	cc CallConfig,
) pagers.NextPageResulter[models.DriveItemable] {
	deltaPager := c.newDriveItemDeltaPager(
		driveID,
		prevDeltaLink,
		cc)

	npr := pagers.NewNextPageResults[models.DriveItemable]()

	// asynchronously enumerate pages on the caller's behalf.
	// they only need to consume the pager and call Results at
	// the end.
	go pagers.DeltaEnumerateItems[models.DriveItemable](
		ctx,
		deltaPager,
		npr,
		prevDeltaLink)

	return npr
}

// ---------------------------------------------------------------------------
// user's drives pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.Driveable] = &userDrivePager{}

type userDrivePager struct {
	userID  string
	gs      graph.Servicer
	builder *users.ItemDrivesRequestBuilder
	options *users.ItemDrivesRequestBuilderGetRequestConfiguration
}

func (c Drives) NewUserDrivePager(
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
		gs:      c.Stable,
		options: requestConfig,
		builder: c.Stable.
			Client().
			Users().
			ByUserId(userID).
			Drives(),
	}

	return res
}

type nopUserDrivePage struct {
	drive models.Driveable
}

func (nl nopUserDrivePage) GetValue() []models.Driveable {
	return []models.Driveable{nl.drive}
}

func (nl nopUserDrivePage) GetOdataNextLink() *string {
	return nil
}

func (p *userDrivePager) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.Driveable], error) {
	// we only ever want to return the user's default drive.
	d, err := p.gs.
		Client().
		Users().
		ByUserId(p.userID).
		Drive().
		Get(ctx, nil)

	return &nopUserDrivePage{drive: d}, clues.Stack(err).OrNil()
}

func (p *userDrivePager) SetNextLink(link string) {
	p.builder = users.NewItemDrivesRequestBuilder(link, p.gs.Adapter())
}

func (p *userDrivePager) ValidModTimes() bool {
	return true
}

// ---------------------------------------------------------------------------
// site's libraries pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.Driveable] = &siteDrivePager{}

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
func (c Drives) NewSiteDrivePager(
	siteID string,
	fields []string,
) *siteDrivePager {
	requestConfig := &sites.ItemDrivesRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemDrivesRequestBuilderGetQueryParameters{
			Select: fields,
		},
	}

	res := &siteDrivePager{
		gs:      c.Stable,
		options: requestConfig,
		builder: c.Stable.
			Client().
			Sites().
			BySiteId(siteID).
			Drives(),
	}

	return res
}

func (p *siteDrivePager) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.Driveable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, clues.Stack(err).OrNil()
}

func (p *siteDrivePager) SetNextLink(link string) {
	p.builder = sites.NewItemDrivesRequestBuilder(link, p.gs.Adapter())
}

func (p *siteDrivePager) ValidModTimes() bool {
	return true
}

// ---------------------------------------------------------------------------
// drive pager
// ---------------------------------------------------------------------------

// GetAllDrives fetches all drives for the given pager
func GetAllDrives(
	ctx context.Context,
	pager pagers.NonDeltaHandler[models.Driveable],
) ([]models.Driveable, error) {
	ds, err := pagers.BatchEnumerateItems(ctx, pager)

	// no license or drives available.
	// return a non-error and let the caller assume an empty result set.
	// TODO: is this the best way to handle this?
	// what about returning a ResourceNotFound error as is standard elsewhere?
	if err != nil &&
		(clues.HasLabel(err, graph.LabelsMysiteNotFound) || clues.HasLabel(err, graph.LabelsNoSharePointLicense)) {
		logger.CtxErr(ctx, err).Infof("resource owner does not have a drive")

		return make([]models.Driveable, 0), nil
	}

	return ds, clues.Stack(err).OrNil()
}
