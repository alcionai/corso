package sharepoint

import (
	"context"

	bmodel "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/microsoftgraph/msgraph-beta-sdk-go/sites"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/account"
)

// GetSitePages retrieves a collection of Pages related to the give Site.
// Returns error if error experienced during the call
func GetSitePage(
	ctx context.Context,
	creds account.M365Config,
	siteID string,
	pages []string,
) ([]bmodel.SitePageable, error) {
	adpt, err := graph.CreateBetaAdapter(creds.AzureTenantID, creds.AzureClientID, creds.AzureClientSecret)
	if err != nil {
		return nil, support.ConnectorStackErrorTraceWrap(err, "fetching beta adapter")
	}

	service := graph.NewBetaService(adpt)
	col := make([]bmodel.SitePageable, 0)

	for _, entry := range pages {
		page, err := service.Client().SitesById(siteID).PagesById(entry).Get(ctx, nil)
		if err != nil {
			return nil, support.ConnectorStackErrorTraceWrap(err, "fetching page: "+entry)
		}

		col = append(col, page)
	}

	return col, nil
}

// fetchPages utility function to return the tuple of item
func fetchPages(ctx context.Context, bs graph.BetaService, siteID string) ([]listTuple, error) {
	var (
		builder    = bs.Client().SitesById(siteID).Pages()
		opts       = fetchPageOptions()
		pageTuples = make([]listTuple, 0)
	)

	for {
		resp, err := builder.Get(ctx, opts)
		if err != nil {
			return nil, support.ConnectorStackErrorTraceWrap(err, "failed fetching site page")
		}

		for _, entry := range resp.GetValue() {
			pid := *entry.GetId()
			temp := listTuple{id: pid, name: pid}

			if entry.GetName() != nil {
				temp.name = *entry.GetName()
			}

			pageTuples = append(pageTuples, temp)
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = sites.NewItemPagesRequestBuilder(*resp.GetOdataNextLink(), bs.Adapter())
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
