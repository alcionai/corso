package onedrive

import (
	"context"
	"strings"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

const (
	maxDrivesRetries = 3

	// nextLinkKey is used to find the next link in a paged
	// graph response
	nextLinkKey           = "@odata.nextLink"
	itemNotFoundErrorCode = "itemNotFound"
)

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
	pager api.DriveItemEnumerator,
	driveID, driveName string,
	collector itemCollector,
	oldPaths map[string]string,
	prevDelta string,
	errs *fault.Bus,
) (DeltaUpdate, map[string]string, map[string]struct{}, error) {
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

		vals, err := pager.ValuesIn(page)
		if err != nil {
			return DeltaUpdate{}, nil, nil, graph.Wrap(ctx, err, "extracting items from response")
		}

		err = collector(
			ctx,
			driveID,
			driveName,
			vals,
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
func newItem(name string, folder bool) models.DriveItemable {
	itemToCreate := models.NewDriveItem()
	itemToCreate.SetName(&name)

	if folder {
		itemToCreate.SetFolder(models.NewFolder())
	} else {
		itemToCreate.SetFile(models.NewFile())
	}

	return itemToCreate
}

type Displayable struct {
	models.DriveItemable
}

func (op *Displayable) GetDisplayName() *string {
	return op.GetName()
}

// GetAllFolders returns all folders in all drives for the given user. If a
// prefix is given, returns all folders with that prefix, regardless of if they
// are a subfolder or top-level folder in the hierarchy.
func GetAllFolders(
	ctx context.Context,
	bh BackupHandler,
	pager api.DrivePager,
	prefix string,
	errs *fault.Bus,
) ([]*Displayable, error) {
	ds, err := api.GetAllDrives(ctx, pager, true, maxDrivesRetries)
	if err != nil {
		return nil, clues.Wrap(err, "getting OneDrive folders")
	}

	var (
		folders = map[string]*Displayable{}
		el      = errs.Local()
	)

	for _, drive := range ds {
		if el.Failure() != nil {
			break
		}

		var (
			id   = ptr.Val(drive.GetId())
			name = ptr.Val(drive.GetName())
		)

		ictx := clues.Add(ctx, "drive_id", id, "drive_name", clues.Hide(name))
		collector := func(
			_ context.Context,
			_, _ string,
			items []models.DriveItemable,
			_ map[string]string,
			_ map[string]string,
			_ map[string]struct{},
			_ map[string]map[string]string,
			_ bool,
			_ *fault.Bus,
		) error {
			for _, item := range items {
				// Skip the root item.
				if item.GetRoot() != nil {
					continue
				}

				// Only selecting folders right now, not packages.
				if item.GetFolder() == nil {
					continue
				}

				itemID := ptr.Val(item.GetId())
				if len(itemID) == 0 {
					logger.Ctx(ctx).Info("folder missing ID")
					continue
				}

				if !strings.HasPrefix(ptr.Val(item.GetName()), prefix) {
					continue
				}

				// Add the item instead of the folder because the item has more
				// functionality.
				folders[itemID] = &Displayable{item}
			}

			return nil
		}

		_, _, _, err = collectItems(
			ictx,
			bh.ItemPager(id, "", nil),
			id,
			name,
			collector,
			map[string]string{},
			"",
			errs)
		if err != nil {
			el.AddRecoverable(clues.Wrap(err, "enumerating items in drive"))
		}
	}

	res := make([]*Displayable, 0, len(folders))

	for _, f := range folders {
		res = append(res, f)
	}

	return res, el.Failure()
}
