package sharepoint

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type getSiteRooter interface {
	GetRoot(ctx context.Context, cc api.CallConfig) (models.Siteable, error)
}

func IsServiceEnabled(
	ctx context.Context,
	gsr getSiteRooter,
	resource string,
) (bool, error) {
	_, err := gsr.GetRoot(ctx, api.CallConfig{})
	if err != nil {
		if clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
			return false, nil
		}

		return false, clues.Stack(err)
	}

	return true, nil
}
