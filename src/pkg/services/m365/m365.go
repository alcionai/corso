package m365

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/pkg/account"
)

// Users returns a list of users in the specified M365 tenant
// TODO: Implement paging support
func Users(ctx context.Context, m365Account account.Account) ([]string, error) {
	gc, err := connector.NewGraphConnector(ctx, m365Account, connector.Users)
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize M365 graph connection")
	}

	return gc.GetUsers(), nil
}

// UserIDs returns a list of user IDs for the specified M365 tenant
// TODO: Implement paging support
func UserIDs(ctx context.Context, m365Account account.Account) ([]string, error) {
	gc, err := connector.NewGraphConnector(ctx, m365Account, connector.Users)
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize M365 graph connection")
	}

	return gc.GetUsersIds(), nil
}

func GetEmailAndUserID(ctx context.Context, m365Account account.Account) (map[string]string, error) {
	gc, err := connector.NewGraphConnector(ctx, m365Account, connector.Users)
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize M365 graph connection")
	}

	return gc.Users, nil
}

// SiteURLs returns a list of SharePoint site WebURLs in the specified M365 tenant
func SiteURLs(ctx context.Context, m365Account account.Account) ([]string, error) {
	gc, err := connector.NewGraphConnector(ctx, m365Account, connector.Sites)
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize M365 graph connection")
	}

	return gc.GetSiteWebURLs(), nil
}

// SiteURLs returns a list of SharePoint sites IDs in the specified M365 tenant
func SiteIDs(ctx context.Context, m365Account account.Account) ([]string, error) {
	gc, err := connector.NewGraphConnector(ctx, m365Account, connector.Sites)
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize M365 graph connection")
	}

	return gc.GetSiteIDs(), nil
}
