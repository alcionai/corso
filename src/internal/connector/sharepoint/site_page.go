package sharepoint

import (
	"context"

	bmodel "github.com/microsoftgraph/msgraph-beta-sdk-go/models"

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
