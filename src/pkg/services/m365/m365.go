package m365

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/internal/connector/discovery"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
)

type User struct {
	PrincipalName string
	ID            string
	Name          string
}

// Users returns a list of users in the specified M365 tenant
// TODO: Implement paging support
func Users(ctx context.Context, acct account.Account) ([]*User, error) {
	gc, err := connector.NewGraphConnector(ctx, graph.HTTPClient(graph.NoTimeout()), acct, connector.Users)
	if err != nil {
		return nil, errors.Wrap(err, "initializing M365 graph connection")
	}

	users, err := discovery.Users(ctx, gc.Owners.Users())
	if err != nil {
		return nil, err
	}

	ret := make([]*User, 0, len(users))

	for _, u := range users {
		pu, err := parseUser(u)
		if err != nil {
			return nil, errors.Wrap(err, "parsing userable")
		}

		ret = append(ret, pu)
	}

	return ret, nil
}

func UserIDs(ctx context.Context, acct account.Account) ([]string, error) {
	users, err := Users(ctx, acct)
	if err != nil {
		return nil, err
	}

	ret := make([]string, 0, len(users))
	for _, u := range users {
		ret = append(ret, u.ID)
	}

	return ret, nil
}

// UserPNs retrieves all user principleNames in the tenant.  Principle Names
// can be used analogous userIDs in graph API queries.
func UserPNs(ctx context.Context, acct account.Account) ([]string, error) {
	users, err := Users(ctx, acct)
	if err != nil {
		return nil, err
	}

	ret := make([]string, 0, len(users))
	for _, u := range users {
		ret = append(ret, u.PrincipalName)
	}

	return ret, nil
}

// SiteURLs returns a list of SharePoint site WebURLs in the specified M365 tenant
func SiteURLs(ctx context.Context, acct account.Account) ([]string, error) {
	gc, err := connector.NewGraphConnector(ctx, graph.HTTPClient(graph.NoTimeout()), acct, connector.Sites)
	if err != nil {
		return nil, errors.Wrap(err, "initializing M365 graph connection")
	}

	return gc.GetSiteWebURLs(), nil
}

// SiteURLs returns a list of SharePoint sites IDs in the specified M365 tenant
func SiteIDs(ctx context.Context, acct account.Account) ([]string, error) {
	gc, err := connector.NewGraphConnector(ctx, graph.HTTPClient(graph.NoTimeout()), acct, connector.Sites)
	if err != nil {
		return nil, errors.Wrap(err, "initializing graph connection")
	}

	return gc.GetSiteIDs(), nil
}

// parseUser extracts information from `models.Userable` we care about
func parseUser(item models.Userable) (*User, error) {
	if item.GetUserPrincipalName() == nil {
		return nil, clues.New("userable missing principal name").
			With("user_id", *item.GetId()) // TODO: pii
	}

	u := &User{PrincipalName: *item.GetUserPrincipalName(), ID: *item.GetId()}

	if item.GetDisplayName() != nil {
		u.Name = *item.GetDisplayName()
	}

	return u, nil
}
