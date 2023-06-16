package api

import (
	"context"
	"fmt"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/sites"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	onedrive "github.com/alcionai/corso/src/internal/m365/onedrive/consts"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// non-delta item pager
// ---------------------------------------------------------------------------

var _ itemPager[models.DriveItemable] = &driveItemPageCtrl{}

type driveItemPageCtrl struct {
	gs      graph.Servicer
	builder *drives.ItemItemsItemChildrenRequestBuilder
	options *drives.ItemItemsItemChildrenRequestBuilderGetRequestConfiguration
}

func (c Drives) NewDriveItemPager(
	driveID, containerID string,
	selectProps ...string,
) itemPager[models.DriveItemable] {
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

func (p *driveItemPageCtrl) getPage(ctx context.Context) (PageLinkValuer[models.DriveItemable], error) {
	page, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return EmptyDeltaLinker[models.DriveItemable]{PageLinkValuer: page}, nil
}

func (p *driveItemPageCtrl) setNext(nextLink string) {
	p.builder = drives.NewItemItemsItemChildrenRequestBuilder(nextLink, p.gs.Adapter())
}

func (c Drives) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	driveID, containerID string,
) (map[string]string, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewDriveItemPager(driveID, containerID, idAnd("name", "file", "folder")...)

	items, err := enumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating drive items")
	}

	m := map[string]string{}

	for _, item := range items {
		// folders are returned in this query, we only want the files
		if item.GetFile() == nil {
			continue
		}

		m[DriveItemCollisionKey(item)] = ptr.Val(item.GetId())
	}

	return m, nil
}

// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// delta item pager
// ---------------------------------------------------------------------------

type DriveItemDeltaEnumerator interface {
	GetPage(context.Context) (DeltaPageLinker, error)
	SetNext(nextLink string)
	Reset()
	ValuesIn(DeltaPageLinker) ([]models.DriveItemable, error)
}

var _ DriveItemDeltaEnumerator = &DriveItemDeltaPageCtrl{}

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
			Top:    ptr.To(maxDeltaPageSize),
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

func (p *DriveItemDeltaPageCtrl) GetPage(ctx context.Context) (DeltaPageLinker, error) {
	var (
		resp DeltaPageLinker
		err  error
	)

	resp, err = p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

func (p *DriveItemDeltaPageCtrl) SetNext(link string) {
	p.builder = drives.NewItemItemsItemDeltaRequestBuilder(link, p.gs.Adapter())
}

func (p *DriveItemDeltaPageCtrl) Reset() {
	p.builder = p.gs.Client().
		Drives().
		ByDriveId(p.driveID).
		Items().
		ByDriveItemId(onedrive.RootID).
		Delta()
}

func (p *DriveItemDeltaPageCtrl) ValuesIn(l DeltaPageLinker) ([]models.DriveItemable, error) {
	return getValues[models.DriveItemable](l)
}

// ---------------------------------------------------------------------------
// user's drives pager
// ---------------------------------------------------------------------------

var _ DrivePager = &userDrivePager{}

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

type nopUserDrivePageLinker struct {
	drive models.Driveable
}

func (nl nopUserDrivePageLinker) GetOdataNextLink() *string { return nil }

func (p *userDrivePager) GetPage(ctx context.Context) (PageLinker, error) {
	var (
		resp PageLinker
		err  error
	)

	d, err := p.gs.
		Client().
		Users().
		ByUserId(p.userID).
		Drive().
		Get(ctx, nil)
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

func (p *userDrivePager) ValuesIn(l PageLinker) ([]models.Driveable, error) {
	nl, ok := l.(*nopUserDrivePageLinker)
	if !ok || nl == nil {
		return nil, clues.New(fmt.Sprintf("improper page linker struct for user drives: %T", l))
	}

	// TODO(keepers): turn back on when we can separate drive enumeration
	// from default drive lookup.

	// return getValues[models.Driveable](l)

	return []models.Driveable{nl.drive}, nil
}

// ---------------------------------------------------------------------------
// site's libraries pager
// ---------------------------------------------------------------------------

var _ DrivePager = &siteDrivePager{}

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

func (p *siteDrivePager) GetPage(ctx context.Context) (PageLinker, error) {
	var (
		resp PageLinker
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

func (p *siteDrivePager) ValuesIn(l PageLinker) ([]models.Driveable, error) {
	return getValues[models.Driveable](l)
}

// ---------------------------------------------------------------------------
// drive pager
// ---------------------------------------------------------------------------

// DrivePager pages through different types of drive owners
type DrivePager interface {
	GetPage(context.Context) (PageLinker, error)
	SetNext(nextLink string)
	ValuesIn(PageLinker) ([]models.Driveable, error)
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
			page PageLinker
		)

		// Retry Loop for Drive retrieval. Request can timeout
		for i := 0; i <= maxRetryCount; i++ {
			page, err = pager.GetPage(ctx)
			if err != nil {
				if clues.HasLabel(err, graph.LabelsMysiteNotFound) || clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
					logger.CtxErr(ctx, err).Infof("resource owner does not have a drive")
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

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func getValues[T any](l PageLinker) ([]T, error) {
	page, ok := l.(interface{ GetValue() []T })
	if !ok {
		return nil, clues.New("page does not comply with GetValue() interface").
			With("page_item_type", fmt.Sprintf("%T", l))
	}

	return page.GetValue(), nil
}
