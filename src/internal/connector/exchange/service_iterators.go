package exchange

import (
	"context"
	"fmt"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

const (
	errEmailFolderNotFound = "ErrorSyncFolderNotFound"
	errItemNotFound        = "ErrorItemNotFound"
)

var errContainerDeleted = errors.New("container deleted")

func hasErrorCode(err error, code string) bool {
	var oDataError *odataerrors.ODataError
	if !errors.As(err, &oDataError) {
		return false
	}

	return oDataError.GetError() != nil &&
		oDataError.GetError().GetCode() != nil &&
		*oDataError.GetError().GetCode() == code
}

// FilterContainersAndFillCollections is a utility function
// that places the M365 object ids belonging to specific directories
// into a Collection. Messages outside of those directories are omitted.
// @param collection is filled with during this function.
// Supports all exchange applications: Contacts, Events, and Mail
func FilterContainersAndFillCollections(
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
		// deleted from this map, leaving only the deleted maps behind
		deletedPaths = common.CopyMap(dps.paths)
	)

	for _, c := range resolver.Items() {
		if ctrlOpts.FailFast && errs != nil {
			return errs
		}

		cID := *c.GetId()

		delete(deletedPaths, cID)

		// Only create a collection if the path matches the scope.
		currPath, ok := pathAndMatch(qp, c, scope)
		if !ok {
			continue
		}

		var prevPath path.Path

		if p, ok := dps.paths[cID]; ok {
			var err error
			if prevPath, err = pathFromPrevString(p); err != nil {
				logger.Ctx(ctx).Error(err)
			}
		}

		service, err := createService(qp.Credentials)
		if err != nil {
			errs = support.WrapAndAppend(qp.ResourceOwner, err, errs)
			continue
		}

		fetchFunc, err := getFetchIDFunc(qp.Category)
		if err != nil {
			errs = support.WrapAndAppend(qp.ResourceOwner, err, errs)
			continue
		}

		var deletedInFlight bool

		jobs, delta, err := fetchFunc(ctx, service, qp.ResourceOwner, cID, dps.deltas[cID])
		if err != nil && !errors.Is(err, errContainerDeleted) {
			deletedInFlight = true
			errs = support.WrapAndAppend(qp.ResourceOwner, err, errs)
		}

		if len(delta) > 0 {
			deltaURLs[cID] = delta
		}

		// Delay creating the new container so we can handle setting the current
		// path correctly if the folder was deleted.
		if deletedInFlight {
			currPath = nil
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

		if deletedInFlight {
			continue
		}

		edc.jobs = append(edc.jobs, jobs...)

		// add the current path for the container ID to be used in the next backup
		// as the "previous path", for reference in case of a rename or relocation.
		currPaths[cID] = currPath.String()
	}

	// any path that wasn't present in the resolver was deleted by the user.
	// relocations and renames will have removed the dir by id earlier.  What's
	// left in deletedPaths are only the previous paths that did not appear as
	// children of the root.
	for fID, ps := range deletedPaths {
		service, err := createService(qp.Credentials)
		if err != nil {
			errs = support.WrapAndAppend(qp.ResourceOwner, err, errs)
			continue
		}

		prevPath, err := pathFromPrevString(ps)
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
		collections[fID] = &edc
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
) ([]string, string, error)

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

// FetchEventIDsFromCalendar returns a list of all M365IDs of events of the targeted Calendar.
func FetchEventIDsFromCalendar(
	ctx context.Context,
	gs graph.Servicer,
	user, calendarID, oldDelta string,
) ([]string, string, error) {
	var (
		errs *multierror.Error
		ids  []string
	)

	options, err := optionsForEventsByCalendar([]string{"id"})
	if err != nil {
		return nil, "", err
	}

	builder := gs.Client().
		UsersById(user).
		CalendarsById(calendarID).
		Events()

	for {
		resp, err := builder.Get(ctx, options)
		if err != nil {
			if hasErrorCode(err, errItemNotFound) {
				// The folder was deleted between the time we populated the container
				// cache and when we tried to fetch data for it. All we can do is
				// return no jobs because we've only pulled basic info about each
				// item.
				return nil, "", errors.WithStack(errContainerDeleted)
			}

			return nil, "", errors.Wrap(err, support.ConnectorStackErrorTrace(err))
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
	return ids, "", errs.ErrorOrNil()
}

// FetchContactIDsFromDirectory function that returns a list of  all the m365IDs of the contacts
// of the targeted directory
func FetchContactIDsFromDirectory(
	ctx context.Context,
	gs graph.Servicer,
	user, directoryID, oldDelta string,
) ([]string, string, error) {
	var (
		errs     *multierror.Error
		ids      []string
		deltaURL string
	)

	options, err := optionsForContactFoldersItemDelta([]string{"parentFolderId"})
	if err != nil {
		return nil, deltaURL, errors.Wrap(err, "getting query options")
	}

	builder := gs.Client().
		UsersById(user).
		ContactFoldersById(directoryID).
		Contacts().
		Delta()

	// TODO(rkeepers): Awaiting full integration of incremental support, else this
	// will cause unexpected behavior/errors.
	// if len(oldDelta) > 0 {
	// 	builder = msuser.NewUsersItemContactFoldersItemContactsDeltaRequestBuilder(oldDelta, gs.Adapter())
	// }

	for {
		resp, err := builder.Get(ctx, options)
		if err != nil {
			if hasErrorCode(err, errItemNotFound) {
				// The folder was deleted between the time we populated the container
				// cache and when we tried to fetch data for it. All we can do is
				// return no jobs because we've only pulled basic info about each
				// item.
				return nil, "", errors.WithStack(errContainerDeleted)
			}

			return nil, deltaURL, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, item := range resp.GetValue() {
			if item.GetId() == nil {
				errs = multierror.Append(
					errs,
					errors.Errorf("contact with nil ID in folder %s", directoryID),
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

	return ids, deltaURL, errs.ErrorOrNil()
}

// FetchMessageIDsFromDirectory function that returns a list of  all the m365IDs of the exchange.Mail
// of the targeted directory
func FetchMessageIDsFromDirectory(
	ctx context.Context,
	gs graph.Servicer,
	user, directoryID, oldDelta string,
) ([]string, string, error) {
	var (
		errs     *multierror.Error
		ids      []string
		deltaURL string
	)

	options, err := optionsForFolderMessagesDelta([]string{"id"})
	if err != nil {
		return nil, deltaURL, errors.Wrap(err, "getting query options")
	}

	builder := gs.Client().
		UsersById(user).
		MailFoldersById(directoryID).
		Messages().
		Delta()

	// TODO(rkeepers): Awaiting full integration of incremental support, else this
	// will cause unexpected behavior/errors.
	// if len(oldDelta) > 0 {
	// 	builder = msuser.NewUsersItemMailFoldersItemMessagesDeltaRequestBuilder(oldDelta, gs.Adapter())
	// }

	for {
		resp, err := builder.Get(ctx, options)
		if err != nil {
			if hasErrorCode(err, errEmailFolderNotFound) {
				// The folder was deleted between the time we populated the container
				// cache and when we tried to fetch data for it. All we can do is
				// return no jobs because we've only pulled basic info about each
				// item.
				return nil, "", errors.WithStack(errContainerDeleted)
			}

			return nil, deltaURL, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
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

	return ids, deltaURL, errs.ErrorOrNil()
}
