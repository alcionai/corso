package exchange

import (
	"context"
	"fmt"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// carries details about delta retrieval in aggregators
type deltaUpdate struct {
	// the deltaLink itself
	url string
	// true if the old delta was marked as invalid
	reset bool
}

// filterContainersAndFillCollections is a utility function
// that places the M365 object ids belonging to specific directories
// into a Collection. Messages outside of those directories are omitted.
// @param collection is filled with during this function.
// Supports all exchange applications: Contacts, Events, and Mail
func filterContainersAndFillCollections(
	ctx context.Context,
	qp graph.QueryParams,
	collections map[string]data.Collection,
	statusUpdater support.StatusUpdater,
	resolver graph.ContainerResolver,
	scope selectors.ExchangeScope,
	dps DeltaPaths,
	ctrlOpts control.Options,
) error {
	var (
		errs error
		oi   = CategoryToOptionIdentifier(qp.Category)
		// folder ID -> delta url or folder path lookups
		deltaURLs = map[string]string{}
		currPaths = map[string]string{}
		// copy of previousPaths.  any folder found in the resolver get
		// deleted from this map, leaving only the deleted folders behind
		tombstones = makeTombstones(dps)
	)

	getJobs, err := getFetchIDFunc(qp.Category)
	if err != nil {
		return support.WrapAndAppend(qp.ResourceOwner, err, errs)
	}

	for _, c := range resolver.Items() {
		if ctrlOpts.FailFast && errs != nil {
			return errs
		}

		// cannot be moved out of the loop,
		// else we run into state issues.
		service, err := createService(qp.Credentials)
		if err != nil {
			errs = support.WrapAndAppend(qp.ResourceOwner, err, errs)
			continue
		}

		cID := *c.GetId()

		currPath, ok := includeContainer(qp, c, scope)
		if currPath != nil {
			// this path exists, do not delete it, even if it isn't
			// included in this backup.
			delete(tombstones, currPath.String())
		}
		// Only create a collection if the path matches the scope.
		if !ok {
			continue
		}

		var (
			dp          = dps[cID]
			prevDelta   = dp.delta
			prevPathStr = dp.path
			prevPath    path.Path
		)

		if len(prevPathStr) > 0 {
			if prevPath, err = pathFromPrevString(prevPathStr); err != nil {
				logger.Ctx(ctx).Error(err)
			}
		}

		jobs, newDelta, err := getJobs(ctx, service, qp.ResourceOwner, cID, prevDelta)
		if err != nil {
			// race conditions happen, containers might get deleted while
			// this process is in flight.  If it was deleted, we remake the
			// tombstone, just to be sure it gets deleted from storage.
			if graph.IsErrDeletedInFlight(err) == nil {
				errs = support.WrapAndAppend(qp.ResourceOwner, err, errs)
			} else {
				tombstones[currPath.String()] = struct{}{}
			}

			continue
		}

		if len(newDelta.url) > 0 {
			deltaURLs[cID] = newDelta.url
		}

		if newDelta.reset {
			// the previous delta was invalid or otherwise unusable.
			// We need to mark the collection as New, and tombstone
			// the path to clear out any data for a clean refresh.
			prevPath = nil
			tombstones[currPath.String()] = struct{}{}
		}

		edc := NewCollection(
			qp.ResourceOwner,
			currPath,
			prevPath,
			oi,
			service,
			statusUpdater,
			ctrlOpts,
		)

		collections[cID] = &edc
		edc.jobs = append(edc.jobs, jobs...)

		// add the current path for the container ID to be used in the next backup
		// as the "previous path", for reference in case of a rename or relocation.
		currPaths[cID] = currPath.String()
	}

	// A tombstone is a collection path that needs to be marked for deletion.
	// Tombstones can occur for a number of reasons: the delta token expired,
	// the container was deleted in flight, or the user deleted the container
	// between backup runs.  If events combine to both delete and write content
	// to the same container (ex: container_1 gets deleted, then container_2
	// gets created with the same name), it is assumed that the backup consumer
	// processes deletions before creations, making the combined operation safe.
	for p := range tombstones {
		service, err := createService(qp.Credentials)
		if err != nil {
			errs = support.WrapAndAppend(qp.ResourceOwner, err, errs)
			continue
		}

		prevPath, err := pathFromPrevString(p)
		if err != nil {
			logger.Ctx(ctx).Error(err)
			continue
		}

		edc := NewCollection(
			qp.ResourceOwner,
			nil, // marks the collection as deleted
			prevPath,
			oi,
			service,
			statusUpdater,
			ctrlOpts,
		)
		collections[p] = &edc
	}

	entries := []graph.MetadataCollectionEntry{
		graph.NewMetadataEntry(graph.PreviousPathFileName, currPaths),
	}

	if len(deltaURLs) > 0 {
		entries = append(entries, graph.NewMetadataEntry(graph.DeltaURLsFileName, deltaURLs))
	}

	if col, err := graph.MakeMetadataCollection(
		qp.Credentials.AzureTenantID,
		qp.ResourceOwner,
		path.ExchangeService,
		qp.Category,
		entries,
		statusUpdater,
	); err != nil {
		errs = support.WrapAndAppend("making metadata collection", err, errs)
	} else if col != nil {
		collections["metadata"] = col
	}

	return errs
}

// produces a set keyed by path strings from the deltapaths map.
// Each entry in the set will, if not removed, produce a collection
// that will delete the tombstone by path.
func makeTombstones(dps DeltaPaths) map[string]struct{} {
	r := make(map[string]struct{}, len(dps))

	for _, v := range dps {
		r[v.path] = struct{}{}
	}

	return r
}

func pathFromPrevString(ps string) (path.Path, error) {
	p, err := path.FromDataLayerPath(ps, false)
	if err != nil {
		return nil, errors.Wrap(err, "parsing previous path string")
	}

	return p, nil
}

func IterativeCollectContactContainers(
	containers map[string]graph.Container,
	nameContains string,
	errUpdater func(string, error),
) func(any) bool {
	return func(entry any) bool {
		folder, ok := entry.(models.ContactFolderable)
		if !ok {
			errUpdater("iterateCollectContactContainers",
				errors.New("casting item to models.ContactFolderable"))
			return false
		}

		include := len(nameContains) == 0 ||
			strings.Contains(*folder.GetDisplayName(), nameContains)

		if include {
			containers[*folder.GetDisplayName()] = folder
		}

		return true
	}
}

func IterativeCollectCalendarContainers(
	containers map[string]graph.Container,
	nameContains string,
	errUpdater func(string, error),
) func(any) bool {
	return func(entry any) bool {
		cal, ok := entry.(models.Calendarable)
		if !ok {
			errUpdater("iterativeCollectCalendarContainers",
				errors.New("casting item to models.Calendarable"))
			return false
		}

		include := len(nameContains) == 0 ||
			strings.Contains(*cal.GetName(), nameContains)
		if include {
			temp := CreateCalendarDisplayable(cal)
			containers[*temp.GetDisplayName()] = temp
		}

		return true
	}
}

// FetchIDFunc collection of helper functions which return a list of all item
// IDs in the given container and a delta token for future requests if the
// container supports fetching delta records.
type FetchIDFunc func(
	ctx context.Context,
	gs graph.Servicer,
	user, containerID, oldDeltaToken string,
) ([]string, deltaUpdate, error)

func getFetchIDFunc(category path.CategoryType) (FetchIDFunc, error) {
	switch category {
	case path.EmailCategory:
		return FetchMessageIDsFromDirectory, nil
	case path.EventsCategory:
		return FetchEventIDsFromCalendar, nil
	case path.ContactsCategory:
		return FetchContactIDsFromDirectory, nil
	default:
		return nil, fmt.Errorf("category %s not supported by getFetchIDFunc", category)
	}
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// FetchEventIDsFromCalendar returns a list of all M365IDs of events of the targeted Calendar.
func FetchEventIDsFromCalendar(
	ctx context.Context,
	gs graph.Servicer,
	user, calendarID, oldDelta string,
) ([]string, deltaUpdate, error) {
	var (
		errs *multierror.Error
		ids  []string
	)

	options, err := optionsForEventsByCalendar([]string{"id"})
	if err != nil {
		return nil, deltaUpdate{}, err
	}

	builder := gs.Client().
		UsersById(user).
		CalendarsById(calendarID).
		Events()

	for {
		resp, err := builder.Get(ctx, options)
		if err != nil {
			if err := graph.IsErrDeletedInFlight(err); err != nil {
				return nil, deltaUpdate{}, err
			}

			return nil, deltaUpdate{}, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, item := range resp.GetValue() {
			if item.GetId() == nil {
				errs = multierror.Append(
					errs,
					errors.Errorf("event with nil ID in calendar %s", calendarID),
				)

				// TODO(ashmrtn): Handle fail-fast.
				continue
			}

			ids = append(ids, *item.GetId())
		}

		nextLink := resp.GetOdataNextLink()
		if nextLink == nil || len(*nextLink) == 0 {
			break
		}

		builder = msuser.NewItemCalendarsItemEventsRequestBuilder(*nextLink, gs.Adapter())
	}

	// Events don't have a delta endpoint so just return an empty string.
	return ids, deltaUpdate{}, errs.ErrorOrNil()
}

// ---------------------------------------------------------------------------
// contacts
// ---------------------------------------------------------------------------

// FetchContactIDsFromDirectory function that returns a list of  all the m365IDs of the contacts
// of the targeted directory
func FetchContactIDsFromDirectory(
	ctx context.Context,
	gs graph.Servicer,
	user, directoryID, oldDelta string,
) ([]string, deltaUpdate, error) {
	var (
		errs       *multierror.Error
		ids        []string
		deltaURL   string
		resetDelta bool
	)

	options, err := optionsForContactFoldersItemDelta([]string{"parentFolderId"})
	if err != nil {
		return nil, deltaUpdate{}, errors.Wrap(err, "getting query options")
	}

	getIDs := func(builder *msuser.ItemContactFoldersItemContactsDeltaRequestBuilder) error {
		for {
			resp, err := builder.Get(ctx, options)
			if err != nil {
				if err := graph.IsErrDeletedInFlight(err); err != nil {
					return err
				}

				if err := graph.IsErrInvalidDelta(err); err != nil {
					return err
				}

				return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
			}

			for _, item := range resp.GetValue() {
				if item.GetId() == nil {
					errs = multierror.Append(
						errs,
						errors.Errorf("item with nil ID in folder %s", directoryID),
					)

					// TODO(ashmrtn): Handle fail-fast.
					continue
				}

				ids = append(ids, *item.GetId())
			}

			delta := resp.GetOdataDeltaLink()
			if delta != nil && len(*delta) > 0 {
				deltaURL = *delta
			}

			nextLink := resp.GetOdataNextLink()
			if nextLink == nil || len(*nextLink) == 0 {
				break
			}

			builder = msuser.NewItemContactFoldersItemContactsDeltaRequestBuilder(*nextLink, gs.Adapter())
		}

		return nil
	}

	if len(oldDelta) > 0 {
		err := getIDs(msuser.NewItemContactFoldersItemContactsDeltaRequestBuilder(oldDelta, gs.Adapter()))
		// happy path
		if err == nil {
			return ids, deltaUpdate{deltaURL, false}, errs.ErrorOrNil()
		}
		// only return on error if it is NOT a delta issue.
		// otherwise we'll retry the call with the regular builder
		if graph.IsErrInvalidDelta(err) == nil {
			return nil, deltaUpdate{}, err
		}

		resetDelta = true
		errs = nil
	}

	builder := gs.Client().
		UsersById(user).
		ContactFoldersById(directoryID).
		Contacts().
		Delta()

	if err := getIDs(builder); err != nil {
		return nil, deltaUpdate{}, err
	}

	return ids, deltaUpdate{deltaURL, resetDelta}, errs.ErrorOrNil()
}

// ---------------------------------------------------------------------------
// messages
// ---------------------------------------------------------------------------

// FetchMessageIDsFromDirectory function that returns a list of  all the m365IDs of the exchange.Mail
// of the targeted directory
func FetchMessageIDsFromDirectory(
	ctx context.Context,
	gs graph.Servicer,
	user, directoryID, oldDelta string,
) ([]string, deltaUpdate, error) {
	var (
		errs       *multierror.Error
		ids        []string
		deltaURL   string
		resetDelta bool
	)

	options, err := optionsForFolderMessagesDelta([]string{"isRead"})
	if err != nil {
		return nil, deltaUpdate{}, errors.Wrap(err, "getting query options")
	}

	getIDs := func(builder *msuser.ItemMailFoldersItemMessagesDeltaRequestBuilder) error {
		for {
			resp, err := builder.Get(ctx, options)
			if err != nil {
				if err := graph.IsErrDeletedInFlight(err); err != nil {
					return err
				}

				if err := graph.IsErrInvalidDelta(err); err != nil {
					return err
				}

				return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
			}

			for _, item := range resp.GetValue() {
				if item.GetId() == nil {
					errs = multierror.Append(
						errs,
						errors.Errorf("item with nil ID in folder %s", directoryID),
					)

					// TODO(ashmrtn): Handle fail-fast.
					continue
				}

				ids = append(ids, *item.GetId())
			}

			delta := resp.GetOdataDeltaLink()
			if delta != nil && len(*delta) > 0 {
				deltaURL = *delta
			}

			nextLink := resp.GetOdataNextLink()
			if nextLink == nil || len(*nextLink) == 0 {
				break
			}

			builder = msuser.NewItemMailFoldersItemMessagesDeltaRequestBuilder(*nextLink, gs.Adapter())
		}

		return nil
	}

	if len(oldDelta) > 0 {
		err := getIDs(msuser.NewItemMailFoldersItemMessagesDeltaRequestBuilder(oldDelta, gs.Adapter()))
		// happy path
		if err == nil {
			return ids, deltaUpdate{deltaURL, false}, errs.ErrorOrNil()
		}
		// only return on error if it is NOT a delta issue.
		// otherwise we'll retry the call with the regular builder
		if graph.IsErrInvalidDelta(err) == nil {
			return nil, deltaUpdate{}, err
		}

		resetDelta = true
		errs = nil
	}

	builder := gs.Client().
		UsersById(user).
		MailFoldersById(directoryID).
		Messages().
		Delta()

	if err := getIDs(builder); err != nil {
		return nil, deltaUpdate{}, err
	}

	return ids, deltaUpdate{deltaURL, resetDelta}, errs.ErrorOrNil()
}
