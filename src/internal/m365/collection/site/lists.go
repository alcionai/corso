package site

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/sites"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

// ListToSPInfo translates models.Listable metadata into searchable content
// List Details: https://learn.microsoft.com/en-us/graph/api/resources/list?view=graph-rest-1.0
func ListToSPInfo(lst models.Listable) *details.SharePointInfo {
	var (
		name     = ptr.Val(lst.GetDisplayName())
		webURL   = ptr.Val(lst.GetWebUrl())
		created  = ptr.Val(lst.GetCreatedDateTime())
		modified = ptr.Val(lst.GetLastModifiedDateTime())
		count    = len(lst.GetItems())
	)

	return &details.SharePointInfo{
		ItemType:  details.SharePointList,
		ItemName:  name,
		ItemCount: int64(count),
		Created:   created,
		Modified:  modified,
		WebURL:    webURL,
	}
}

type ListTuple struct {
	ID   string
	Name string
}

func preFetchListOptions() *sites.ItemListsRequestBuilderGetRequestConfiguration {
	selecting := []string{"id", "displayName"}
	queryOptions := sites.ItemListsRequestBuilderGetQueryParameters{
		Select: selecting,
	}
	options := &sites.ItemListsRequestBuilderGetRequestConfiguration{
		QueryParameters: &queryOptions,
	}

	return options
}

func PreFetchLists(
	ctx context.Context,
	gs graph.Servicer,
	siteID string,
) ([]ListTuple, error) {
	var (
		builder    = gs.Client().Sites().BySiteId(siteID).Lists()
		options    = preFetchListOptions()
		listTuples = make([]ListTuple, 0)
	)

	for {
		resp, err := builder.Get(ctx, options)
		if err != nil {
			return nil, graph.Wrap(ctx, err, "getting lists")
		}

		for _, entry := range resp.GetValue() {
			var (
				id   = ptr.Val(entry.GetId())
				name = ptr.Val(entry.GetDisplayName())
				temp = ListTuple{ID: id, Name: name}
			)

			if len(name) == 0 {
				temp.Name = id
			}

			listTuples = append(listTuples, temp)
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = sites.NewItemListsRequestBuilder(link, gs.Adapter())
	}

	return listTuples, nil
}

// DeleteList removes a list object from a site.
// deletes require unique http clients
// https://github.com/alcionai/corso/issues/2707
func DeleteList(
	ctx context.Context,
	gs graph.Servicer,
	siteID, listID string,
) error {
	err := gs.Client().Sites().BySiteId(siteID).Lists().ByListId(listID).Delete(ctx, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting list")
	}

	return nil
}
