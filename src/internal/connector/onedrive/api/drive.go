package api

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/alcionai/clues"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	msdrives "github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	mssites "github.com/microsoftgraph/msgraph-sdk-go/sites"
	msusers "github.com/microsoftgraph/msgraph-sdk-go/users"

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
	builder *msdrives.ItemRootDeltaRequestBuilder
	options *msdrives.ItemRootDeltaRequestBuilderGetRequestConfiguration
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

	requestConfig := &msdrives.ItemRootDeltaRequestBuilderGetRequestConfiguration{
		Headers: headers,
		QueryParameters: &msdrives.ItemRootDeltaRequestBuilderGetQueryParameters{
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
		res.builder = msdrives.NewItemRootDeltaRequestBuilder(link, gs.Adapter())
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
	p.builder = msdrives.NewItemRootDeltaRequestBuilder(link, p.gs.Adapter())
}

func (p *driveItemPager) Reset() {
	p.builder = p.gs.Client().DrivesById(p.driveID).Root().Delta()
}

func (p *driveItemPager) ValuesIn(l api.DeltaPageLinker) ([]models.DriveItemable, error) {
	return getValues[models.DriveItemable](l)
}

type userDrivePager struct {
	gs      graph.Servicer
	builder *msusers.ItemDrivesRequestBuilder
	options *msusers.ItemDrivesRequestBuilderGetRequestConfiguration
}

func NewUserDrivePager(
	gs graph.Servicer,
	userID string,
	fields []string,
) *userDrivePager {
	requestConfig := &msusers.ItemDrivesRequestBuilderGetRequestConfiguration{
		QueryParameters: &msusers.ItemDrivesRequestBuilderGetQueryParameters{
			Select: fields,
		},
	}

	res := &userDrivePager{
		gs:      gs,
		options: requestConfig,
		builder: gs.Client().UsersById(userID).Drives(),
	}

	return res
}

func (p *userDrivePager) GetPage(ctx context.Context) (api.PageLinker, error) {
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

func (p *userDrivePager) SetNext(link string) {
	p.builder = msusers.NewItemDrivesRequestBuilder(link, p.gs.Adapter())
}

func (p *userDrivePager) ValuesIn(l api.PageLinker) ([]models.Driveable, error) {
	return getValues[models.Driveable](l)
}

type siteDrivePager struct {
	gs      graph.Servicer
	builder *mssites.ItemDrivesRequestBuilder
	options *mssites.ItemDrivesRequestBuilderGetRequestConfiguration
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
	requestConfig := &mssites.ItemDrivesRequestBuilderGetRequestConfiguration{
		QueryParameters: &mssites.ItemDrivesRequestBuilderGetQueryParameters{
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
	p.builder = mssites.NewItemDrivesRequestBuilder(link, p.gs.Adapter())
}

func (p *siteDrivePager) ValuesIn(l api.PageLinker) ([]models.Driveable, error) {
	return getValues[models.Driveable](l)
}

// ---------------------------------------------------------------------------
// Drive Paging
// ---------------------------------------------------------------------------

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
	drives := []models.Driveable{}

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

		drives = append(drives, tmp...)

		nextLink := ptr.Val(page.GetOdataNextLink())
		if len(nextLink) == 0 {
			break
		}

		pager.SetNext(nextLink)
	}

	logger.Ctx(ctx).Debugf("retrieved %d valid drives", len(drives))

	return drives, nil
}
