package onedrive

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type getDefaultDriver interface {
	GetDefaultDrive(ctx context.Context, userID string) (models.Driveable, error)
}

func IsServiceEnabled(
	ctx context.Context,
	gdd getDefaultDriver,
	resource string,
) (bool, error) {
	_, err := gdd.GetDefaultDrive(ctx, resource)
	if err != nil {
		// we consider this a non-error case, since it
		// answers the question the caller is asking.
		if clues.HasLabel(err, graph.LabelsMysiteNotFound) || clues.HasLabel(err, graph.LabelsNoSharePointLicense) {
			return false, nil
		}

		if graph.IsErrUserNotFound(err) {
			return false, clues.Stack(core.ErrResourceOwnerNotFound, err)
		}

		if graph.IsErrResourceLocked(err) {
			return false, clues.Stack(graph.ErrResourceLocked, err)
		}

		return false, clues.Stack(err)
	}

	return true, nil
}
