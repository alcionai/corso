package sharepoint

import (
	"context"

	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
)

type sharePointService struct {
	client      msgraphsdk.GraphServiceClient
	adapter     msgraphsdk.GraphRequestAdapter
	failFast    bool // if true service will exit sequence upon encountering an error
	credentials account.M365Config
}

///------------------------------------------------------------
// Functions to comply with graph.Service Interface
//-------------------------------------------------------

func (es *sharePointService) Client() *msgraphsdk.GraphServiceClient {
	return &es.client
}

func (es *sharePointService) Adapter() *msgraphsdk.GraphRequestAdapter {
	return &es.adapter
}

func (es *sharePointService) ErrPolicy() bool {
	return es.failFast
}

// createService internal constructor for sharePointService struct returns an error
// iff the params for the entry are incorrect (e.g. len(TenantID) == 0, etc.)
// NOTE: Incorrect account information will result in errors on subsequent queries.
func createService(credentials account.M365Config, shouldFailFast bool) (*sharePointService, error) {
	adapter, err := graph.CreateAdapter(
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret,
	)
	if err != nil {
		return nil, errors.Wrap(err, "creating microsoft graph service")
	}

	service := sharePointService{
		adapter:     *adapter,
		client:      *msgraphsdk.NewGraphServiceClient(adapter),
		failFast:    shouldFailFast,
		credentials: credentials,
	}

	return &service, nil
}

// PopulateContainerResolver gets a container resolver if one is available for
// this category of data. If one is not available, returns nil so that other
// logic in the caller can complete as long as they check if the resolver is not
// nil. If an error occurs populating the resolver, returns an error.
func PopulateContainerResolver(
	ctx context.Context,
	qp graph.QueryParams,
) (graph.ContainerResolver, error) {
	return nil, nil
	// var (
	// 	c            graph.ContainerPopulater
	// 	service, err = createService(qp.Credentials, qp.FailFast)
	// 	cacheRoot    string
	// )

	// if err != nil {
	// 	return nil, err
	// }

	// switch qp.Category {
	// case path.FilesCategory:
	// 	c = &driveCache{
	// 		siteID: qp.ResourceOwner,
	// 		gs:     service,
	// 	}
	// 	cacheRoot = "root"

	// default:
	// 	return nil, fmt.Errorf("ContainerResolver not present for %s type", qp.Category)
	// }

	// if err := c.Populate(ctx, cacheRoot); err != nil {
	// 	return nil, errors.Wrap(err, "populating container resolver")
	// }

	// return c, nil
}
