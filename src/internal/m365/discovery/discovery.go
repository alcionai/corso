package discovery

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

type getter interface {
	GetByID(context.Context, string) (models.Userable, error)
}

type GetInfoer interface {
	GetInfo(context.Context, string) (*api.UserInfo, error)
}

type getWithInfoer interface {
	getter
	GetInfoer
}

type GetDefaultDriver interface {
	GetDefaultDrive(ctx context.Context, userID string) (models.Driveable, error)
}

type getAller interface {
	GetAll(ctx context.Context, errs *fault.Bus) ([]models.Userable, error)
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
// users
// ---------------------------------------------------------------------------

// Users fetches all users in the tenant.
func Users(
	ctx context.Context,
	ga getAller,
	errs *fault.Bus,
) ([]models.Userable, error) {
	users, err := ga.GetAll(ctx, errs)
	if err != nil {
		return nil, clues.Wrap(err, "getting all users")
	}

	return users, nil
}

// UserDetails fetches detailed info like - userPurpose for all users in the tenant.
func GetUserInfo(
	ctx context.Context,
	acct account.Account,
	userID string,
	errs *fault.Bus,
) (*api.UserInfo, error) {
	client, err := apiClient(ctx, acct)
	if err != nil {
		return nil, err
	}

	return client.Users().GetInfo(ctx, userID)
}

// User fetches a single user's data.
func User(
	ctx context.Context,
	gwi getWithInfoer,
	userID string,
) (models.Userable, *api.UserInfo, error) {
	u, err := gwi.GetByID(ctx, userID)
	if err != nil {
		if graph.IsErrUserNotFound(err) {
			return nil, nil, clues.Stack(graph.ErrResourceOwnerNotFound, err).With("user_id", userID)
		}

		return nil, nil, clues.Wrap(err, "getting user")
	}

	ui, err := gwi.GetInfo(ctx, userID)
	if err != nil {
		return nil, nil, clues.Wrap(err, "getting user info")
	}

	return u, ui, nil
}

// UserInfo produces extensible user info: metadata that is relevant
// or identified in Corso, but not in m365.
func UserInfo(
	ctx context.Context,
	gi GetInfoer,
	userID string,
) (*api.UserInfo, error) {
	ui, err := gi.GetInfo(ctx, userID)
	if err != nil {
		return nil, clues.Wrap(err, "getting user info")
	}

	return ui, nil
}

// ---------------------------------------------------------------------------
// sites
// ---------------------------------------------------------------------------

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
