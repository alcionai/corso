package discovery

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

type getter interface {
	GetByID(context.Context, string) (models.Userable, error)
}

type getInfoer interface {
	GetInfo(context.Context, string) (*api.UserInfo, error)
}

type getWithInfoer interface {
	getter
	getInfoer
}

// ---------------------------------------------------------------------------
// helpers
// ---------------------------------------------------------------------------

func apiClient(ctx context.Context, acct account.Account) (api.Client, error) {
	m365, err := acct.M365Config()
	if err != nil {
		return api.Client{}, clues.Wrap(err, "retrieving m365 account configuration").WithClues(ctx)
	}

	client, err := api.NewClient(m365)
	if err != nil {
		return api.Client{}, clues.Wrap(err, "creating api client").WithClues(ctx)
	}

	return client, nil
}

// ---------------------------------------------------------------------------
// api
// ---------------------------------------------------------------------------

// Users fetches all users in the tenant.
func Users(
	ctx context.Context,
	acct account.Account,
	errs *fault.Bus,
) ([]models.Userable, error) {
	client, err := apiClient(ctx, acct)
	if err != nil {
		return nil, err
	}

	return client.Users().GetAll(ctx, errs)
}

// UsersDetails fetches detailed info like - userPurpose for all users in the tenant.
func UsersDetails(
	ctx context.Context,
	acct account.Account,
	userID string,
	errs *fault.Bus,
) (string, error) {
	client, err := apiClient(ctx, acct)
	if err != nil {
		return "", err
	}

	return client.Users().GetUserPurpose(ctx, userID)
}

// User fetches a single user's data.
func User(ctx context.Context, gwi getWithInfoer, userID string) (models.Userable, *api.UserInfo, error) {
	u, err := gwi.GetByID(ctx, userID)
	if err != nil {
		if graph.IsErrUserNotFound(err) {
			return nil, nil, clues.New("resource owner not found within tenant").With("user_id", userID)
		}

		return nil, nil, clues.Wrap(err, "getting user")
	}

	ui, err := gwi.GetInfo(ctx, userID)
	if err != nil {
		return nil, nil, clues.Wrap(err, "getting user info")
	}

	return u, ui, nil
}

// Sites fetches all sharepoint sites in the tenant
func Sites(
	ctx context.Context,
	acct account.Account,
	errs *fault.Bus,
) ([]models.Siteable, error) {
	client, err := apiClient(ctx, acct)
	if err != nil {
		return nil, err
	}

	return client.Sites().GetAll(ctx, errs)
}
