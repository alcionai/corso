package sharepoint

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/m365/graph"
)

type getSiteRooter interface {
	GetRoot(ctx context.Context) (models.Siteable, error)
}

func IsServiceEnabled(
	ctx context.Context,
	gsr getSiteRooter,
	resource string,
) (bool, error) {
	_, err := gsr.GetRoot(ctx)
	if err != nil {
		if clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
			return false, nil
		}

		return false, clues.Stack(err)
	}

	return true, nil
}
