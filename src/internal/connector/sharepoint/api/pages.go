package api

import (
	"context"

	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk/sites"
	"github.com/alcionai/corso/src/internal/connector/support"
)

// GetSitePages retrieves a collection of Pages related to the give Site.
// Returns error if error experienced during the call
func GetSitePage(
	ctx context.Context,
	serv *api.BetaService,
	siteID string,
	pages []string,
) ([]models.SitePageable, error) {
	col := make([]models.SitePageable, 0)
	opts := retrieveSitePageOptions()

	for _, entry := range pages {
		page, err := serv.Client().SitesById(siteID).PagesById(entry).Get(ctx, opts)
		if err != nil {
			return nil, support.ConnectorStackErrorTraceWrap(err, "fetching page: "+entry)
		}

		col = append(col, page)
	}

	return col, nil
}

// fetchPages utility function to return the tuple of item
func FetchPages(ctx context.Context, bs *api.BetaService, siteID string) ([]Tuple, error) {
	var (
		builder    = bs.Client().SitesById(siteID).Pages()
		opts       = fetchPageOptions()
		pageTuples = make([]Tuple, 0)
	)

	for {
		resp, err := builder.Get(ctx, opts)
		if err != nil {
			return nil, support.ConnectorStackErrorTraceWrap(err, "failed fetching site page")
		}

		for _, entry := range resp.GetValue() {
			pid := *entry.GetId()
			temp := Tuple{pid, pid}

			if entry.GetName() != nil {
				temp.Name = *entry.GetName()
			}

			pageTuples = append(pageTuples, temp)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = sites.NewItemPagesRequestBuilder(*resp.GetOdataNextLink(), bs.Client().Adapter())
	}

	return pageTuples, nil
}

// fetchPageOptions is used to return minimal information reltating to Site Pages
// Pages API: https://learn.microsoft.com/en-us/graph/api/resources/sitepage?view=graph-rest-beta
func fetchPageOptions() *sites.ItemPagesRequestBuilderGetRequestConfiguration {
	fields := []string{"id", "name"}
	options := &sites.ItemPagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemPagesRequestBuilderGetQueryParameters{
			Select: fields,
		},
	}

	return options
}

// DeleteSitePage removes the selected page from the SharePoint Site
// https://learn.microsoft.com/en-us/graph/api/sitepage-delete?view=graph-rest-beta
func DeleteSitePage(
	ctx context.Context,
	serv *betasdk.Service,
	siteID, pageID string,
) error {
	err := serv.Client().SitesById(siteID).PagesById(pageID).Delete(ctx, nil)
	if err != nil {
		return support.ConnectorStackErrorTraceWrap(err, "deleting page: "+pageID)
	}

	return nil
}

// retrievePageOptions returns options to expand
func retrieveSitePageOptions() *sites.ItemPagesSitePageItemRequestBuilderGetRequestConfiguration {
	fields := []string{"canvasLayout"}
	options := &sites.ItemPagesSitePageItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.ItemPagesSitePageItemRequestBuilderGetQueryParameters{
			Expand: fields,
		},
	}

	return options
}
