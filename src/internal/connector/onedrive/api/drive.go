package api

import (
	"context"

	msdrives "github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	mssites "github.com/microsoftgraph/msgraph-sdk-go/sites"
	msusers "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/api"
)

func getValues[T any](l api.PageLinker) ([]T, error) {
	page, ok := l.(interface{ GetValue() []T })
	if !ok {
		return nil, errors.Errorf(
			"response of type [%T] does not comply with GetValue() interface",
			l,
		)
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
	requestConfig := &msdrives.ItemRootDeltaRequestBuilderGetRequestConfiguration{
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

	err = graph.RunWithRetry(func() error {
		resp, err = p.builder.Get(ctx, p.options)
		return err
	})

	return resp, err
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

	err = graph.RunWithRetry(func() error {
		resp, err = p.builder.Get(ctx, p.options)
		return err
	})

	return resp, err
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

	err = graph.RunWithRetry(func() error {
		resp, err = p.builder.Get(ctx, p.options)
		return err
	})

	return resp, err
}

func (p *siteDrivePager) SetNext(link string) {
	p.builder = mssites.NewItemDrivesRequestBuilder(link, p.gs.Adapter())
}

func (p *siteDrivePager) ValuesIn(l api.PageLinker) ([]models.Driveable, error) {
	return getValues[models.Driveable](l)
}
