package discovery

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

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
// api
// ---------------------------------------------------------------------------

// Users fetches all users in the tenant.
func Users(
	ctx context.Context,
	acct account.Account,
	errs *fault.Bus,
) ([]models.Userable, error) {
	m365, err := acct.M365Config()
	if err != nil {
		return nil, clues.Wrap(err, "retrieving m365 account configuration").WithClues(ctx)
	}

	client, err := api.NewClient(m365)
	if err != nil {
		return nil, clues.Wrap(err, "creating api client").WithClues(ctx)
	}

	return client.Users().GetAll(ctx, errs)
}

func User(ctx context.Context, gwi getWithInfoer, userID string) (models.Userable, *api.UserInfo, error) {
	u, err := gwi.GetByID(ctx, userID)
	if err != nil {
		if graph.IsErrUserNotFound(err) {
			return nil, nil, fmt.Errorf("resource owner [%s] not found within tenant", userID)
		}

		return nil, nil, errors.Wrap(err, "getting user")
	}

	ui, err := gwi.GetInfo(ctx, userID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting user info")
	}

	return u, ui, nil
}
