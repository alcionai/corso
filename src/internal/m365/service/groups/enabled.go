package groups

import (
	"context"

	"github.com/alcionai/clues"
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
	_, err := gbi.GetByID(ctx, resource)
	if err != nil {
		// TODO(meain): check for error message in case groups are
		// not enabled at all similar to sharepoint
		return false, clues.Stack(err)
	}

	return true, nil
}
