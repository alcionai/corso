package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	msdrives "github.com/microsoftgraph/msgraph-sdk-go/drives"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	mssites "github.com/microsoftgraph/msgraph-sdk-go/sites"
	msusers "github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/graph/api"
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
	headers.Add("Prefer", "deltashowremovedasdeleted,deltatraversepermissiongaps,deltashowsharingchanges,hierarchicalsharing")

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

// GetDriveIDByName is a helper function to retrieve the M365ID of a site drive.
// Returns "" if the folder is not within the drive.
// Dependency: Requires "name" and "id" to be part of the given options
func (p *siteDrivePager) GetDriveIDByName(ctx context.Context, driveName string) (string, error) {
	var empty string

	for {
		resp, err := p.builder.Get(ctx, p.options)
		if err != nil {
			return empty, graph.Stack(ctx, err)
		}

		for _, entry := range resp.GetValue() {
			if ptr.Val(entry.GetName()) == driveName {
				return ptr.Val(entry.GetId()), nil
			}
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		p.builder = mssites.NewItemDrivesRequestBuilder(link, p.gs.Adapter())
	}

	return empty, nil
}

// GetFolderIDByName is a helper function to retrieve the M365ID of a folder within a site document library.
// Returns "" if the folder is not within the drive
func (p *siteDrivePager) GetFolderIDByName(ctx context.Context, driveID, folderName string) (string, error) {
	var empty string

	// *msdrives.ItemRootChildrenRequestBuilder
	builder := p.gs.Client().DrivesById(driveID).Root().Children()
	option := &msdrives.ItemRootChildrenRequestBuilderGetRequestConfiguration{
		QueryParameters: &msdrives.ItemRootChildrenRequestBuilderGetQueryParameters{
			Select: []string{"id", "name", "folder"},
		},
	}

	for {
		resp, err := builder.Get(ctx, option)
		if err != nil {
			return empty, graph.Stack(ctx, err)
		}

		for _, entry := range resp.GetValue() {
			if entry.GetFolder() == nil {
				continue
			}

			if ptr.Val(entry.GetName()) == folderName {
				return ptr.Val(entry.GetId()), nil
			}
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = msdrives.NewItemRootChildrenRequestBuilder(link, p.gs.Adapter())
	}

	return empty, nil
}
