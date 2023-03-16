package sharepoint

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/sites"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
)

// GetAllSitesForTenant makes a GraphQuery request retrieving all sites in the tenant.
// Due to restrictions in filter capabilities for site queries, the returned iterable
// will contain all personal sites for all users in the org.
func GetAllSitesForTenant(ctx context.Context, gs graph.Servicer) (serialization.Parsable, error) {
	// url := "https://graph.microsoft.com/beta/sites" +
	// 	"?$top=1000" +
	// 	"&$filter=displayname ne null AND NOT(contains(weburl, 'sharepoint.com/personal/'))"
	// s, err := sites.NewItemSitesRequestBuilder(url, gs.Adapter()).Get(ctx, nil)
	// if err != nil {
	// 	return nil, graph.Wrap(ctx, err, "getting sites")
	// }

	options := &sites.SitesRequestBuilderGetRequestConfiguration{
		QueryParameters: &sites.SitesRequestBuilderGetQueryParameters{
			// Select: []string{"id", "name", "weburl"},
			Filter: ptr.To("NOT(contains(webUrl,'sharepoint.com/personal/'))"),
		},
	}

	s, err := gs.Client().Sites().Get(ctx, options)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting sites")
	}

	return s, nil
}
