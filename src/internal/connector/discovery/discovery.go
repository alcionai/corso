package discovery

import (
	"context"

	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	userSelectID            = "id"
	userSelectPrincipalName = "userPrincipalName"
	userSelectDisplayName   = "displayName"
)

func Users(ctx context.Context, gs graph.Servicer, tenantID string) ([]models.Userable, error) {
	users := make([]models.Userable, 0)

	options := &msuser.UsersRequestBuilderGetRequestConfiguration{
		QueryParameters: &msuser.UsersRequestBuilderGetQueryParameters{
			Select: []string{userSelectID, userSelectPrincipalName, userSelectDisplayName},
		},
	}

	response, err := gs.Client().Users().Get(ctx, options)
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"retrieving resources for tenant %s: %s",
			tenantID,
			support.ConnectorStackErrorTrace(err),
		)
	}

	iter, err := msgraphgocore.NewPageIterator(response, gs.Adapter(),
		models.CreateUserCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	var iterErrs error

	callbackFunc := func(item interface{}) bool {
		u, err := parseUser(item)
		if err != nil {
			iterErrs = support.WrapAndAppend("discovering users: ", err, iterErrs)
			return true
		}

		users = append(users, u)

		return true
	}

	if err := iter.Iterate(ctx, callbackFunc); err != nil {
		return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	return users, iterErrs
}

type UserInfo struct {
	DiscoveredServices map[path.ServiceType]struct{}
}

func User(ctx context.Context, gs graph.Servicer, userID string) (models.Userable, *UserInfo, error) {
	user, err := gs.Client().UsersById(userID).Get(ctx, nil)
	if err != nil {
		return nil, nil, errors.Wrapf(
			err,
			"retrieving resource for tenant: %s",
			support.ConnectorStackErrorTrace(err),
		)
	}

	// Assume all services are enabled
	userInfo := &UserInfo{
		DiscoveredServices: map[path.ServiceType]struct{}{
			path.ExchangeService: {},
			path.OneDriveService: {},
		},
	}

	// Discover which services the user has enabled

	// Exchange: Query `MailFolders`
	_, err = gs.Client().UsersById(userID).MailFolders().Get(ctx, nil)
	if err != nil {
		if !graph.IsErrExchangeMailFolderNotFound(err) {
			return nil, nil, errors.Wrapf(
				err,
				"retrieving mail folders for tenant: %s",
				support.ConnectorStackErrorTrace(err),
			)
		}

		delete(userInfo.DiscoveredServices, path.ExchangeService)
	}

	// TODO: OneDrive

	return user, userInfo, nil
}

// parseUser extracts information from `models.Userable` we care about
func parseUser(item interface{}) (models.Userable, error) {
	m, ok := item.(models.Userable)
	if !ok {
		return nil, errors.New("iteration retrieved non-User item")
	}

	if m.GetId() == nil {
		return nil, errors.Errorf("no ID for User")
	}

	if m.GetUserPrincipalName() == nil {
		return nil, errors.Errorf("no principal name for User: %s", *m.GetId())
	}

	return m, nil
}
