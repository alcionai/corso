package onedrive

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/delta"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/connector/graph"
	"github.com/alcionai/corso/internal/connector/support"
	"github.com/alcionai/corso/pkg/logger"
)

const (
	// nextLinkKey is used to find the next link in a paged
	// graph response
	nextLinkKey = "@odata.nextLink"
)

// Enumerates the drives for the specified user
func drives(ctx context.Context, service graph.Service, user string) ([]models.Driveable, error) {
	r, err := service.Client().UsersById(user).Drives().Get()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve user drives. user: %s, details: %s",
			user, support.ConnectorStackErrorTrace(err))
	}
	logger.Ctx(ctx).Debugf("Found %d drives for user %s", len(r.GetValue()), user)

	return r.GetValue(), nil
}

// itemCollector functions collect the items found in a drive
type itemCollector func(ctx context.Context, driveID string, items []models.DriveItemable) error

// collectItems will enumerate all items in the specified drive and hand them to the
// provided `collector` method
func collectItems(
	ctx context.Context,
	service graph.Service,
	driveID string,
	collector itemCollector,
) error {
	// TODO: Specify a timestamp in the delta query
	// https://docs.microsoft.com/en-us/graph/api/driveitem-delta?
	// view=graph-rest-1.0&tabs=http#example-4-retrieving-delta-results-using-a-timestamp
	builder := service.Client().DrivesById(driveID).Root().Delta()
	for {
		r, err := builder.Get()
		if err != nil {
			return errors.Wrapf(
				err,
				"failed to query drive items. details: %s",
				support.ConnectorStackErrorTrace(err),
			)
		}

		err = collector(ctx, driveID, r.GetValue())
		if err != nil {
			return err
		}

		// Check if there are more items
		if _, found := r.GetAdditionalData()[nextLinkKey]; !found {
			break
		}
		nextLink := r.GetAdditionalData()[nextLinkKey].(*string)
		logger.Ctx(ctx).Debugf("Found %s nextLink", *nextLink)
		builder = delta.NewDeltaRequestBuilder(*nextLink, service.Adapter())
	}
	return nil
}
