package exchange

import (
	"context"
	"fmt"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// FilterContainersAndFillCollections is a utility function
// that places the M365 object ids belonging to specific directories
// into a Collection. Messages outside of those directories are omitted.
// @param collection is filled with during this function.
// Supports all exchange applications: Contacts, Events, and Mail
func FilterContainersAndFillCollections(
	ctx context.Context,
	qp graph.QueryParams,
	collections map[string]*Collection,
	statusUpdater support.StatusUpdater,
	resolver graph.ContainerResolver,
) error {
	var (
		category       = graph.ScopeToPathCategory(qp.Scope)
		collectionType = categoryToOptionIdentifier(category)
		errs           error
	)

	logger.Ctx(ctx).Debugw("preparing collection", "user", qp.User)

	colProgress, closer := observe.CollectionProgress(
		qp.User,
		"Exchange",
		"Gathering Data")
	defer close(colProgress)

	go closer()

	for _, c := range resolver.Items() {
		dirPath, ok := pathAndMatch(qp, category, c)
		if ok {
			// Create only those that match
			service, err := createService(qp.Credentials, qp.FailFast)
			if err != nil {
				errs = support.WrapAndAppend(
					qp.User+" failed to create service during FilterContainerAndFillCollection",
					err,
					errs)

				if qp.FailFast {
					return errs
				}
			}

			edc := NewCollection(
				qp.User,
				dirPath,
				collectionType,
				service,
				statusUpdater,
			)
			collections[*c.GetId()] = &edc

			logger.Ctx(ctx).
				With("user", qp.User).
				With("dirPath", dirPath.Folder()).
				Debug("adding collection to backup")

			colProgress <- struct{}{}
		}
	}

	for directoryID, col := range collections {
		fetchFunc, err := getFetchIDFunc(category)
		if err != nil {
			errs = support.WrapAndAppend(
				qp.User,
				err,
				errs)

			if qp.FailFast {
				return errs
			}

			continue
		}

		jobs, err := fetchFunc(ctx, col.service, qp.User, directoryID)
		if err != nil {
			errs = support.WrapAndAppend(
				qp.User,
				err,
				errs,
			)
		}

		col.jobs = append(col.jobs, jobs...)
	}

	return errs
}

func IterativeCollectContactContainers(
	containers map[string]graph.Container,
	nameContains string,
	errUpdater func(string, error),
) func(any) bool {
	return func(entry any) bool {
		folder, ok := entry.(models.ContactFolderable)
		if !ok {
			errUpdater("", errors.New("casting item to models.ContactFolderable"))
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
			errUpdater("failure during IterativeCollectCalendarContainers",
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

// FetchIDFunc collection of helper functions which return a list of strings
// from a response.
type FetchIDFunc func(ctx context.Context, gs graph.Service, user, containerID string) ([]string, error)

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
	gs graph.Service,
	user, calendarID string,
) ([]string, error) {
	ids := []string{}

	logger.Ctx(ctx).
		With("user", user).
		With("containerID", calendarID).
		Debug("retrieving events")

	response, err := gs.Client().
		UsersById(user).
		CalendarsById(calendarID).
		Events().Get(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	pageIterator, err := msgraphgocore.NewPageIterator(
		response,
		gs.Adapter(),
		models.CreateEventCollectionResponseFromDiscriminatorValue,
	)
	if err != nil {
		return nil, errors.Wrap(err, "iterator creation failure during fetchEventIDs")
	}

	var errs *multierror.Error

	colProgress, closer := observe.CollectionProgress(
		user,
		"Events",
		calendarID)
	defer close(colProgress)

	go closer()

	err = pageIterator.Iterate(ctx, func(pageItem any) bool {
		entry, ok := pageItem.(graph.Idable)
		if !ok {
			errs = multierror.Append(errs, errors.New("item without GetId() call"))
			return true
		}

		if entry.GetId() == nil {
			errs = multierror.Append(errs, errors.New("item with nil ID"))
		}

		ids = append(ids, *entry.GetId())

		colProgress <- struct{}{}

		return true
	})

	if err != nil {
		return nil, errors.Wrap(
			err,
			support.ConnectorStackErrorTrace(err)+
				" :iterateFailure for fetching events from calendar "+calendarID,
		)
	}

	return ids, errs.ErrorOrNil()
}

// FetchContactIDsFromDirectory function that returns a list of  all the m365IDs of the contacts
// of the targeted directory
func FetchContactIDsFromDirectory(ctx context.Context, gs graph.Service, user, directoryID string) ([]string, error) {
	options, err := optionsForContactFoldersItem([]string{"parentFolderId"})
	if err != nil {
		return nil, err
	}

	ids := []string{}

	logger.Ctx(ctx).
		With("user", user).
		With("containerID", directoryID).
		Debug("retrieving contacts")

	response, err := gs.Client().
		UsersById(user).
		ContactFoldersById(directoryID).
		Contacts().
		Get(ctx, options)
	if err != nil {
		return nil, errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	pageIterator, err := msgraphgocore.NewPageIterator(
		response,
		gs.Adapter(),
		models.CreateContactCollectionResponseFromDiscriminatorValue,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failure to create iterator during FecthContactIDs")
	}

	var errs *multierror.Error

	colProgress, closer := observe.CollectionProgress(
		user,
		"Contacts",
		directoryID)
	defer close(colProgress)

	go closer()

	err = pageIterator.Iterate(ctx, func(pageItem any) bool {
		entry, ok := pageItem.(graph.Idable)
		if !ok {
			errs = multierror.Append(
				errs,
				errors.New("casting pageItem to models.Contactable"),
			)

			return true
		}

		ids = append(ids, *entry.GetId())

		colProgress <- struct{}{}

		return true
	})

	if err != nil {
		return nil,
			errors.Wrap(
				err,
				support.ConnectorStackErrorTrace(err)+
					" :iterate failure during fetching contactIDs from directory "+directoryID,
			)
	}

	return ids, errs.ErrorOrNil()
}

// FetchMessageIDsFromDirectory function that returns a list of  all the m365IDs of the exchange.Mail
// of the targeted directory
func FetchMessageIDsFromDirectory(
	ctx context.Context,
	gs graph.Service,
	user, directoryID string,
) ([]string, error) {
	ids := []string{}

	logger.Ctx(ctx).
		With("user", user).
		With("containerID", directoryID).
		Debug("retrieving emails")

	options, err := optionsForFolderMessages([]string{"id"})
	if err != nil {
		return nil, errors.Wrap(err, "getting query options")
	}

	response, err := gs.Client().
		UsersById(user).
		MailFoldersById(directoryID).
		Messages().
		Get(ctx, options)
	if err != nil {
		return nil, errors.Wrap(
			errors.Wrap(err, support.ConnectorStackErrorTrace(err)),
			"initial folder query",
		)
	}

	pageIter, err := msgraphgocore.NewPageIterator(
		response,
		gs.Adapter(),
		models.CreateMessageCollectionResponseFromDiscriminatorValue,
	)
	if err != nil {
		return nil, errors.Wrap(err, "creating graph iterator")
	}

	var errs *multierror.Error

	colProgress, closer := observe.CollectionProgress(
		user,
		"Emails",
		directoryID)
	defer close(colProgress)

	go closer()

	err = pageIter.Iterate(ctx, func(pageItem any) bool {
		item, ok := pageItem.(graph.Idable)
		if !ok {
			errs = multierror.Append(errs, errors.New("item without ID function"))
			return true
		}

		if item.GetId() == nil {
			errs = multierror.Append(errs, errors.New("item with nil ID"))
			return true
		}

		ids = append(ids, *item.GetId())

		colProgress <- struct{}{}

		return true
	})

	if err != nil {
		return nil, errors.Wrap(
			err,
			support.ConnectorStackErrorTrace(err)+
				" :iterateFailure for fetching messages from directory "+directoryID,
		)
	}

	return ids, errs.ErrorOrNil()
}
