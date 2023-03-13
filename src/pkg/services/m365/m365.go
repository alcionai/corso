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
	"github.com/alcionai/corso/src/pkg/fault"
)

type User struct {
	PrincipalName string
	ID            string
	Name          string
}

// UsersCompat returns a list of users in the specified M365 tenant.
// TODO(ashmrtn): Remove when upstream consumers of the SDK support the fault
// package.
func UsersCompat(ctx context.Context, acct account.Account) ([]*User, error) {
	errs := fault.New(true)

	users, err := Users(ctx, acct, errs)
	if err != nil {
		return nil, err
	}

	return users, errs.Failure()
}

// Users returns a list of users in the specified M365 tenant
// TODO: Implement paging support
func Users(ctx context.Context, acct account.Account, errs *fault.Bus) ([]*User, error) {
	users, err := discovery.Users(ctx, acct, errs)
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

func UserIDs(ctx context.Context, acct account.Account, errs *fault.Bus) ([]string, error) {
	users, err := Users(ctx, acct, errs)
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
func UserPNs(ctx context.Context, acct account.Account, errs *fault.Bus) ([]string, error) {
	users, err := Users(ctx, acct, errs)
	if err != nil {
		return nil, err
	}

	ret := make([]string, 0, len(users))
	for _, u := range users {
		ret = append(ret, u.PrincipalName)
	}

	return ret, nil
}

type Site struct {
	// WebURL that displays the item in the browser
	WebURL string

	// ID is of the format: <site collection hostname>.<site collection unique id>.<site unique id>
	// for example: contoso.sharepoint.com,abcdeab3-0ccc-4ce1-80ae-b32912c9468d,xyzud296-9f7c-44e1-af81-3c06d0d43007
	ID string
}

// Sites returns a list of Sites in a specified M365 tenant
func Sites(ctx context.Context, acct account.Account, errs *fault.Bus) ([]*Site, error) {
	gc, err := connector.NewGraphConnector(ctx, graph.HTTPClient(graph.NoTimeout()), acct, connector.Sites, errs)
	if err != nil {
		return nil, errors.Wrap(err, "initializing M365 graph connection")
	}

	// gc.Sites is a map with keys: SiteURL, values: ID
	ret := make([]*Site, 0, len(gc.Sites))
	for k, v := range gc.Sites {
		ret = append(ret, &Site{
			WebURL: k,
			ID:     v,
		})
	}

	return ret, nil
}

// SiteURLs returns a list of SharePoint site WebURLs in the specified M365 tenant
func SiteURLs(ctx context.Context, acct account.Account, errs *fault.Bus) ([]string, error) {
	gc, err := connector.NewGraphConnector(ctx, graph.HTTPClient(graph.NoTimeout()), acct, connector.Sites, errs)
	if err != nil {
		return nil, errors.Wrap(err, "initializing M365 graph connection")
	}

	return gc.GetSiteWebURLs(), nil
}

// SiteIDs returns a list of SharePoint sites IDs in the specified M365 tenant
func SiteIDs(ctx context.Context, acct account.Account, errs *fault.Bus) ([]string, error) {
	gc, err := connector.NewGraphConnector(ctx, graph.HTTPClient(graph.NoTimeout()), acct, connector.Sites, errs)
	if err != nil {
		return nil, errors.Wrap(err, "initializing graph connection")
	}

	return gc.GetSiteIDs(), nil
}

// parseUser extracts information from `models.Userable` we care about
func parseUser(item models.Userable) (*User, error) {
	if item.GetUserPrincipalName() == nil {
		return nil, clues.New("user missing principal name").
			With("user_id", *item.GetId()) // TODO: pii
	}

	u := &User{PrincipalName: *item.GetUserPrincipalName(), ID: *item.GetId()}

	if item.GetDisplayName() != nil {
		u.Name = *item.GetDisplayName()
	}

	return u, nil
}
