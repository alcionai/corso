package groups

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

type getByIDer interface {
	GetByID(ctx context.Context, identifier string) (models.Groupable, error)
}

func IsServiceEnabled(
	ctx context.Context,
	gbi getByIDer,
	resource string,
) (bool, error) {
	// TODO(meain): check for error message in case groups are
	// not enabled at all similar to sharepoint
	return true, nil
}
