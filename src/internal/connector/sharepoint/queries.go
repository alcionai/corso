package sharepoint

import (
	"context"

	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	mssite "github.com/microsoftgraph/msgraph-sdk-go/sites"

	"github.com/alcionai/corso/src/internal/connector/graph"
)

// GetAllSitesForTenant makes a GraphQuery request retrieving all sites in the tenant.
func GetAllSitesForTenant(ctx context.Context, gs graph.Service) (absser.Parsable, error) {
	options := &mssite.SitesRequestBuilderGetRequestConfiguration{
		QueryParameters: &mssite.SitesRequestBuilderGetQueryParameters{
			Select: []string{"id", "name", "weburl"},
		},
	}

	return gs.Client().Sites().Get(ctx, options)
}
