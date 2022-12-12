package sharepoint

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

// RestoreCollections will restore the specified data collections into OneDrive
func RestoreCollections(
	ctx context.Context,
	service graph.Servicer,
	dest control.RestoreDestination,
	dcs []data.Collection,
	deets *details.Details,
) (*support.ConnectorOperationStatus, error) {
	var (
		restoreMetrics support.CollectionMetrics
		restoreErrors  error
	)

	errUpdater := func(id string, err error) {
		restoreErrors = support.WrapAndAppend(id, err, restoreErrors)
	}

	// Iterate through the data collections and restore the contents of each
	for _, dc := range dcs {
		var (
			metrics  support.CollectionMetrics
			canceled bool
		)

		switch dc.FullPath().Category() {
		case path.LibrariesCategory:
			metrics, canceled = onedrive.RestoreCollection(
				ctx,
				service,
				dc,
				onedrive.OneDriveSource,
				dest.ContainerName,
				deets,
				errUpdater)
		default:
			return nil, errors.Errorf("category %s not supported", dc.FullPath().Category())
		}

		restoreMetrics.Combine(metrics)

		if canceled {
			break
		}
	}

	return support.CreateStatus(
			ctx,
			support.Restore,
			len(dcs),
			restoreMetrics,
			restoreErrors,
			dest.ContainerName),
		nil
}
