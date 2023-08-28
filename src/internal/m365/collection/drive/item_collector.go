package drive

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const maxDrivesRetries = 3

// DeltaUpdate holds the results of a current delta token.  It normally
// gets produced when aggregating the addition and removal of items in
// a delta-queryable folder.
// FIXME: This is same as exchange.api.DeltaUpdate
type DeltaUpdate struct {
	// the deltaLink itself
	URL string
	// true if the old delta was marked as invalid
	Reset bool
}

// itemCollector functions collect the items found in a drive
type itemCollector func(
	ctx context.Context,
	driveID, driveName string,
	driveItems []models.DriveItemable,
	oldPaths map[string]string,
	newPaths map[string]string,
	excluded map[string]struct{},
	itemCollections map[string]map[string]string,
	validPrevDelta bool,
	errs *fault.Bus,
) error

// collectItems will enumerate all items in the specified drive and hand them to the
// provided `collector` method
func collectItems(
	ctx context.Context,
	pager api.DeltaPager[models.DriveItemable],
	driveID, driveName string,
	collector itemCollector,
	oldPaths map[string]string,
	prevDelta string,
	errs *fault.Bus,
) (
	DeltaUpdate,
	map[string]string, // newPaths
	map[string]struct{}, // excluded
	error,
) {
	var (
		newDeltaURL      = ""
		newPaths         = map[string]string{}
		excluded         = map[string]struct{}{}
		invalidPrevDelta = len(prevDelta) == 0

		// itemCollection is used to identify which collection a
		// file belongs to. This is useful to delete a file from the
		// collection it was previously in, in case it was moved to a
		// different collection within the same delta query
		// drive ID -> item ID -> item ID
		itemCollection = map[string]map[string]string{
			driveID: {},
		}
	)

	if !invalidPrevDelta {
		maps.Copy(newPaths, oldPaths)
		pager.SetNext(prevDelta)
	}

	for {
		// assume delta urls here, which allows single-token consumption
		page, err := pager.GetPage(graph.ConsumeNTokens(ctx, graph.SingleGetOrDeltaLC))

		if graph.IsErrInvalidDelta(err) {
			logger.Ctx(ctx).Infow("Invalid previous delta link", "link", prevDelta)

			invalidPrevDelta = true
			newPaths = map[string]string{}

			pager.Reset()

			continue
		}

		if err != nil {
			return DeltaUpdate{}, nil, nil, graph.Wrap(ctx, err, "getting page")
		}

		err = collector(
			ctx,
			driveID,
			driveName,
			page.GetValue(),
			oldPaths,
			newPaths,
			excluded,
			itemCollection,
			invalidPrevDelta,
			errs)
		if err != nil {
			return DeltaUpdate{}, nil, nil, err
		}

		nextLink, deltaLink := api.NextAndDeltaLink(page)

		if len(deltaLink) > 0 {
			newDeltaURL = deltaLink
		}

		// Check if there are more items
		if len(nextLink) == 0 {
			break
		}

		logger.Ctx(ctx).Debugw("Found nextLink", "link", nextLink)
		pager.SetNext(nextLink)
	}

	return DeltaUpdate{URL: newDeltaURL, Reset: invalidPrevDelta}, newPaths, excluded, nil
}

// newItem initializes a `models.DriveItemable` that can be used as input to `createItem`
func newItem(name string, folder bool) *models.DriveItem {
	itemToCreate := models.NewDriveItem()
	itemToCreate.SetName(&name)

	if folder {
		itemToCreate.SetFolder(models.NewFolder())
	} else {
		itemToCreate.SetFile(models.NewFile())
	}

	return itemToCreate
}
