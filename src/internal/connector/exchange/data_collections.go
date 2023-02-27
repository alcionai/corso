package exchange

import (
	"context"
	"encoding/json"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// MetadataFileNames produces the category-specific set of filenames used to
// store graph metadata such as delta tokens and folderID->path references.
func MetadataFileNames(cat path.CategoryType) []string {
	switch cat {
	case path.EmailCategory, path.ContactsCategory:
		return []string{graph.DeltaURLsFileName, graph.PreviousPathFileName}
	default:
		return []string{graph.PreviousPathFileName}
	}
}

type CatDeltaPaths map[path.CategoryType]DeltaPaths

type DeltaPaths map[string]DeltaPath

func (dps DeltaPaths) AddDelta(k, d string) {
	dp, ok := dps[k]
	if !ok {
		dp = DeltaPath{}
	}

	dp.delta = d
	dps[k] = dp
}

func (dps DeltaPaths) AddPath(k, p string) {
	dp, ok := dps[k]
	if !ok {
		dp = DeltaPath{}
	}

	dp.path = p
	dps[k] = dp
}

type DeltaPath struct {
	delta string
	path  string
}

// ParseMetadataCollections produces a map of structs holding delta
// and path lookup maps.
func parseMetadataCollections(
	ctx context.Context,
	colls []data.RestoreCollection,
	errs *fault.Bus,
) (CatDeltaPaths, error) {
	// cdp stores metadata
	cdp := CatDeltaPaths{
		path.ContactsCategory: {},
		path.EmailCategory:    {},
		path.EventsCategory:   {},
	}

	// found tracks the metadata we've loaded, to make sure we don't
	// fetch overlapping copies.
	found := map[path.CategoryType]map[string]struct{}{
		path.ContactsCategory: {},
		path.EmailCategory:    {},
		path.EventsCategory:   {},
	}

	for _, coll := range colls {
		var (
			breakLoop bool
			items     = coll.Items(ctx, errs)
			category  = coll.FullPath().Category()
		)

		for {
			select {
			case <-ctx.Done():
				return nil, clues.Wrap(ctx.Err(), "parsing collection metadata").WithClues(ctx)

			case item, ok := <-items:
				if !ok {
					breakLoop = true
					break
				}

				var (
					m    = map[string]string{}
					cdps = cdp[category]
				)

				err := json.NewDecoder(item.ToReader()).Decode(&m)
				if err != nil {
					return nil, clues.New("decoding metadata json").WithClues(ctx)
				}

				switch item.UUID() {
				case graph.PreviousPathFileName:
					if _, ok := found[category]["path"]; ok {
						return nil, clues.Wrap(clues.New(category.String()), "multiple versions of path metadata").WithClues(ctx)
					}

					for k, p := range m {
						cdps.AddPath(k, p)
					}

					found[category]["path"] = struct{}{}

				case graph.DeltaURLsFileName:
					if _, ok := found[category]["delta"]; ok {
						return nil, clues.Wrap(clues.New(category.String()), "multiple versions of delta metadata").WithClues(ctx)
					}

					for k, d := range m {
						cdps.AddDelta(k, d)
					}

					found[category]["delta"] = struct{}{}
				}

				cdp[category] = cdps
			}

			if breakLoop {
				break
			}
		}
	}

	// Remove any entries that contain a path or a delta, but not both.
	// That metadata is considered incomplete, and needs to incur a
	// complete backup on the next run.
	for _, dps := range cdp {
		for k, dp := range dps {
			if len(dp.delta) == 0 || len(dp.path) == 0 {
				delete(dps, k)
			}
		}
	}

	return cdp, nil
}

// DataCollections returns a DataCollection which the caller can
// use to read mailbox data out for the specified user
// Assumption: User exists
//
//	Add iota to this call -> mail, contacts, calendar,  etc.
func DataCollections(
	ctx context.Context,
	selector selectors.Selector,
	metadata []data.RestoreCollection,
	acct account.M365Config,
	su support.StatusUpdater,
	ctrlOpts control.Options,
	errs *fault.Bus,
) ([]data.BackupCollection, map[string]struct{}, error) {
	eb, err := selector.ToExchangeBackup()
	if err != nil {
		return nil, nil, clues.Wrap(err, "exchange dataCollection selector").WithClues(ctx)
	}

	var (
		user        = selector.DiscreteOwner
		collections = []data.BackupCollection{}
		el          = errs.Local()
	)

	cdps, err := parseMetadataCollections(ctx, metadata, errs)
	if err != nil {
		return nil, nil, err
	}

	for _, scope := range eb.Scopes() {
		if el.Failure() != nil {
			break
		}

		dcs, err := createCollections(
			ctx,
			acct,
			user,
			scope,
			cdps[scope.Category().PathType()],
			ctrlOpts,
			su,
			errs)
		if err != nil {
			el.AddRecoverable(err)
			continue
		}

		collections = append(collections, dcs...)
	}

	return collections, nil, el.Failure()
}

func getterByType(ac api.Client, category path.CategoryType) (addedAndRemovedItemIDsGetter, error) {
	switch category {
	case path.EmailCategory:
		return ac.Mail(), nil
	case path.EventsCategory:
		return ac.Events(), nil
	case path.ContactsCategory:
		return ac.Contacts(), nil
	default:
		return nil, clues.New("no api client registered for category")
	}
}

// createCollections - utility function that retrieves M365
// IDs through Microsoft Graph API. The selectors.ExchangeScope
// determines the type of collections that are retrieved.
func createCollections(
	ctx context.Context,
	creds account.M365Config,
	user string,
	scope selectors.ExchangeScope,
	dps DeltaPaths,
	ctrlOpts control.Options,
	su support.StatusUpdater,
	errs *fault.Bus,
) ([]data.BackupCollection, error) {
	var (
		allCollections = make([]data.BackupCollection, 0)
		ac             = api.Client{Credentials: creds}
		category       = scope.Category().PathType()
	)

	ctx = clues.Add(ctx, "category", category)

	getter, err := getterByType(ac, category)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	// Create collection of ExchangeDataCollection
	collections := make(map[string]data.BackupCollection)

	qp := graph.QueryParams{
		Category:      category,
		ResourceOwner: user,
		Credentials:   creds,
	}

	foldersComplete, closer := observe.MessageWithCompletion(ctx, observe.Bulletf(
		"%s - %s",
		observe.Safe(qp.Category.String()),
		observe.PII(user)))
	defer closer()
	defer close(foldersComplete)

	resolver, err := PopulateExchangeContainerResolver(ctx, qp, errs)
	if err != nil {
		return nil, errors.Wrap(err, "populating container cache")
	}

	err = filterContainersAndFillCollections(
		ctx,
		qp,
		getter,
		collections,
		su,
		resolver,
		scope,
		dps,
		ctrlOpts,
		errs)
	if err != nil {
		return nil, errors.Wrap(err, "filling collections")
	}

	foldersComplete <- struct{}{}

	for _, coll := range collections {
		allCollections = append(allCollections, coll)
	}

	return allCollections, nil
}
