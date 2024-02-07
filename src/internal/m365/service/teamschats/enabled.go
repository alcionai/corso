package teamschats

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/canario/src/pkg/services/m365/api"
)

func IsServiceEnabled(
	ctx context.Context,
	gbi api.GetByIDer[models.Userable],
	resource string,
) (bool, error) {
	// TODO(rkeepers): investgate service enablement checks
	return true, nil
}
