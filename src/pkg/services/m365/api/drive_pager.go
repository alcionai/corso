package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/sites"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	onedrive "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// non-delta item pager
// ---------------------------------------------------------------------------

var _ Pager[models.DriveItemable] = &driveItemPageCtrl{}

type driveItemPageCtrl struct {
	gs      graph.Servicer
	builder *drives.ItemItemsItemChildrenRequestBuilder
	options *drives.ItemItemsItemChildrenRequestBuilderGetRequestConfiguration
}

func (c Drives) NewDriveItemPager(
	driveID, containerID string,
	selectProps ...string,
) Pager[models.DriveItemable] {
	options := &drives.ItemItemsItemChildrenRequestBuilderGetRequestConfiguration{
		QueryParameters: &drives.ItemItemsItemChildrenRequestBuilderGetQueryParameters{},
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
) (NextLinkValuer[models.DriveItemable], error) {
	page, err := p.builder.Get(ctx, p.options)
	return page, graph.Stack(ctx, err).OrNil()
}

func (p *driveItemPageCtrl) SetNextLink(nextLink string) {
	p.builder = drives.NewItemItemsItemChildrenRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *driveItemPageCtrl) ValidModTimes() bool {
	return true
}

type DriveItemIDType struct {
	ItemID   string
	IsFolder bool
}

func (c Drives) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	driveID, containerID string,
) (map[string]DriveItemIDType, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewDriveItemPager(driveID, containerID, idAnd("name")...)

	items, err := enumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating drive items")
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

	items, err := enumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating contacts")
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

var _ DeltaPager[models.DriveItemable] = &DriveItemDeltaPageCtrl{}

type DriveItemDeltaPageCtrl struct {
	gs      graph.Servicer
	driveID string
	builder *drives.ItemItemsItemDeltaRequestBuilder
	options *drives.ItemItemsItemDeltaRequestBuilderGetRequestConfiguration
}

func (c Drives) NewDriveItemDeltaPager(
	driveID, link string,
	selectFields []string,
) *DriveItemDeltaPageCtrl {
	preferHeaderItems := []string{
		"deltashowremovedasdeleted",
		"deltatraversepermissiongaps",
		"deltashowsharingchanges",
		"hierarchicalsharing",
	}

	requestConfig := &drives.ItemItemsItemDeltaRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(preferHeaderItems...),
		QueryParameters: &drives.ItemItemsItemDeltaRequestBuilderGetQueryParameters{
			Select: selectFields,
		},
	}

	res := &DriveItemDeltaPageCtrl{
		gs:      c.Stable,
		driveID: driveID,
		options: requestConfig,
		builder: c.Stable.
			Client().
			Drives().
			ByDriveId(driveID).
			Items().
			ByDriveItemId(onedrive.RootID).
			Delta(),
	}

	if len(link) > 0 {
		res.builder = drives.NewItemItemsItemDeltaRequestBuilder(link, c.Stable.Adapter())
	}

	return res
}

func (p *DriveItemDeltaPageCtrl) GetPage(
	ctx context.Context,
) (DeltaLinkValuer[models.DriveItemable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
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

// ---------------------------------------------------------------------------
// user's drives pager
// ---------------------------------------------------------------------------

var _ Pager[models.Driveable] = &userDrivePager{}

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
) (NextLinkValuer[models.Driveable], error) {
	// we only ever want to return the user's default drive.
	d, err := p.gs.
		Client().
		Users().
		ByUserId(p.userID).
		Drive().
		Get(ctx, nil)

	return &nopUserDrivePage{drive: d}, graph.Stack(ctx, err).OrNil()
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

var _ Pager[models.Driveable] = &siteDrivePager{}

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
) (NextLinkValuer[models.Driveable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
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
	pager Pager[models.Driveable],
) ([]models.Driveable, error) {
	ds, err := enumerateItems(ctx, pager)
	if err != nil && (clues.HasLabel(err, graph.LabelsMysiteNotFound) ||
		clues.HasLabel(err, graph.LabelsNoSharePointLicense)) {
		logger.CtxErr(ctx, err).Infof("resource owner does not have a drive")
		return make([]models.Driveable, 0), nil // no license or drives.
	}

	return ds, graph.Stack(ctx, err).OrNil()
}
