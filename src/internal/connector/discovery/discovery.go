package discovery

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/discovery/api"
	"github.com/alcionai/corso/src/pkg/fault"
)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

type getAller interface {
	GetAll(context.Context, *fault.Errors) ([]models.Userable, error)
}

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
func Users(ctx context.Context, ga getAller, errs *fault.Errors) ([]models.Userable, error) {
	return ga.GetAll(ctx, errs)
}

func User(ctx context.Context, gwi getWithInfoer, userID string) (models.Userable, *api.UserInfo, error) {
	u, err := gwi.GetByID(ctx, userID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting user")
	}

	ui, err := gwi.GetInfo(ctx, userID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting user info")
	}

	return u, ui, nil
}
