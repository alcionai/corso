package api

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	mssites "github.com/microsoftgraph/msgraph-sdk-go/sites"
	msusers "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/api"
)

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
	return p.builder.Get(ctx, p.options)
}

func (p *userDrivePager) SetNext(link string) {
	p.builder = msusers.NewItemDrivesRequestBuilder(link, p.gs.Adapter())
}

func (p *userDrivePager) ValuesIn(l api.PageLinker) ([]models.Driveable, error) {
	page, ok := l.(interface{ GetValue() []models.Driveable })
	if !ok {
		return nil, errors.Errorf(
			"response of type [%T] does not comply with GetValue() interface",
			l,
		)
	}

	return page.GetValue(), nil
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
	return p.builder.Get(ctx, p.options)
}

func (p *siteDrivePager) SetNext(link string) {
	p.builder = mssites.NewItemDrivesRequestBuilder(link, p.gs.Adapter())
}

func (p *siteDrivePager) ValuesIn(l api.PageLinker) ([]models.Driveable, error) {
	page, ok := l.(interface{ GetValue() []models.Driveable })
	if !ok {
		return nil, errors.Errorf(
			"response of type [%T] does not comply with GetValue() interface",
			l,
		)
	}

	return page.GetValue(), nil
}
