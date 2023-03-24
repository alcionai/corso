package exchange

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

var ErrFolderNotFound = errors.New("folder not found")

func createService(credentials account.M365Config) (*graph.Service, error) {
	adapter, err := graph.CreateAdapter(
		credentials.AzureTenantID,
		credentials.AzureClientID,
		credentials.AzureClientSecret)
	if err != nil {
		return nil, errors.Wrap(err, "creating microsoft graph service for exchange")
	}

	return graph.NewService(adapter), nil
}

// populateExchangeContainerResolver gets a folder resolver if one is available for
// this category of data. If one is not available, returns nil so that other
// logic in the caller can complete as long as they check if the resolver is not
// nil. If an error occurs populating the resolver, returns an error.
func PopulateExchangeContainerResolver(
	ctx context.Context,
	qp graph.QueryParams,
	errs *fault.Bus,
) (graph.ContainerResolver, error) {
	var (
		res       graph.ContainerResolver
		cacheRoot string
	)

	ac, err := api.NewClient(qp.Credentials)
	if err != nil {
		return nil, err
	}

	switch qp.Category {
	case path.EmailCategory:
		acm := ac.Mail()
		res = &mailFolderCache{
			userID: qp.ResourceOwner,
			getter: acm,
			enumer: acm,
		}
		cacheRoot = rootFolderAlias

	case path.ContactsCategory:
		acc := ac.Contacts()
		res = &contactFolderCache{
			userID: qp.ResourceOwner,
			getter: acc,
			enumer: acc,
		}
		cacheRoot = DefaultContactFolder

	case path.EventsCategory:
		ecc := ac.Events()
		res = &eventCalendarCache{
			userID: qp.ResourceOwner,
			getter: ecc,
			enumer: ecc,
		}
		cacheRoot = DefaultCalendar

	default:
		return nil, clues.New("no container resolver registered for category").WithClues(ctx)
	}

	if err := res.Populate(ctx, errs, cacheRoot); err != nil {
		return nil, clues.Wrap(err, "populating directory resolver").WithClues(ctx)
	}

	return res, nil
}

// Returns true if the container passes the scope comparison and should be included.
// Returns:
// - the path representing the directory as it should be stored in the repository.
// - the human-readable path using display names.
// - true if the path passes the scope comparison.
func includeContainer(
	qp graph.QueryParams,
	c graph.CachedContainer,
	scope selectors.ExchangeScope,
) (path.Path, path.Path, bool) {
	var (
		directory string
		locPath   path.Path
		category  = scope.Category().PathType()
		pb        = c.Path()
		loc       = c.Location()
	)

	// Clause ensures that DefaultContactFolder is inspected properly
	if category == path.ContactsCategory && ptr.Val(c.GetDisplayName()) == DefaultContactFolder {
		pb = pb.Append(DefaultContactFolder)

		if loc != nil {
			loc = loc.Append(DefaultContactFolder)
		}
	}

	dirPath, err := pb.ToDataLayerExchangePathForCategory(
		qp.Credentials.AzureTenantID,
		qp.ResourceOwner,
		category,
		false)
	// Containers without a path (e.g. Root mail folder) always err here.
	if err != nil {
		return nil, nil, false
	}

	directory = dirPath.Folder(false)

	if loc != nil {
		locPath, err = loc.ToDataLayerExchangePathForCategory(
			qp.Credentials.AzureTenantID,
			qp.ResourceOwner,
			category,
			false)
		// Containers without a path (e.g. Root mail folder) always err here.
		if err != nil {
			return nil, nil, false
		}

		directory = locPath.Folder(false)
	}

	var ok bool

	switch category {
	case path.EmailCategory:
		ok = scope.Matches(selectors.ExchangeMailFolder, directory)
	case path.ContactsCategory:
		ok = scope.Matches(selectors.ExchangeContactFolder, directory)
	case path.EventsCategory:
		ok = scope.Matches(selectors.ExchangeEventCalendar, directory)
	default:
		return nil, nil, false
	}

	return dirPath, locPath, ok
}
