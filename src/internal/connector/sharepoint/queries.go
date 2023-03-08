package sharepoint

import (
	"context"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	mssite "github.com/microsoftgraph/msgraph-sdk-go/sites"

	"github.com/alcionai/corso/src/internal/connector/graph"
)

// GetAllSitesForTenant makes a GraphQuery request retrieving all sites in the tenant.
// Due to restrictions in filter capabilities for site queries, the returned iterable
// will contain all personal sites for all users in the org.
func GetAllSitesForTenant(ctx context.Context, gs graph.Servicer) (absser.Parsable, error) {
	options := &mssite.SitesRequestBuilderGetRequestConfiguration{
		QueryParameters: &mssite.SitesRequestBuilderGetQueryParameters{
			Select: []string{"id", "name", "weburl"},
		},
	}

	sites, err := gs.Client().Sites().Get(ctx, options)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "getting sites")
	}

	return sites, nil
}
