package exchange

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
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
	colls []data.Collection,
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
			items     = coll.Items()
			category  = coll.FullPath().Category()
		)

		for {
			select {
			case <-ctx.Done():
				return nil, errors.Wrap(ctx.Err(), "parsing collection metadata")

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
					return nil, errors.New("decoding metadata json")
				}

				switch item.UUID() {
				case graph.PreviousPathFileName:
					if _, ok := found[category]["path"]; ok {
						return nil, errors.Errorf("multiple versions of %s path metadata", category)
					}

					for k, p := range m {
						cdps.AddPath(k, p)
					}

					found[category]["path"] = struct{}{}

				case graph.DeltaURLsFileName:
					if _, ok := found[category]["delta"]; ok {
						return nil, errors.Errorf("multiple versions of %s delta metadata", category)
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
func DataCollection(
	ctx context.Context,
	selector selectors.Selector,
	metadata []data.Collection,
	userPNs []string,
	acct account.M365Config,
	su support.StatusUpdater,
	ctrlOpts control.Options,
) ([]data.Collection, error) {
	eb, err := selector.ToExchangeBackup()
	if err != nil {
		return nil, errors.Wrap(err, "exchangeDataCollection: parsing selector")
	}

	var (
		scopes      = eb.DiscreteScopes(userPNs)
		collections = []data.Collection{}
		errs        error
	)

	cdps, err := parseMetadataCollections(ctx, metadata)
	if err != nil {
		return nil, err
	}

	for _, scope := range scopes {
		dps := cdps[scope.Category().PathType()]

		dcs, err := createCollections(
			ctx,
			acct,
			scope,
			dps,
			control.Options{},
			su)
		if err != nil {
			user := scope.Get(selectors.ExchangeUser)
			return nil, support.WrapAndAppend(user[0], err, errs)
		}

		collections = append(collections, dcs...)
	}

	return collections, errs
}

// createCollections - utility function that retrieves M365
// IDs through Microsoft Graph API. The selectors.ExchangeScope
// determines the type of collections that are retrieved.
func createCollections(
	ctx context.Context,
	acct account.M365Config,
	scope selectors.ExchangeScope,
	dps deltaPaths,
	ctrlOpts control.Options,
	su support.StatusUpdater,
) ([]data.Collection, error) {
	var (
		errs           *multierror.Error
		users          = scope.Get(selectors.ExchangeUser)
		allCollections = make([]data.Collection, 0)
	)

	// Create collection of ExchangeDataCollection
	for _, user := range users {
		collections := make(map[string]data.Collection)

		qp := graph.QueryParams{
			Category:      scope.Category().PathType(),
			ResourceOwner: user,
			Credentials:   acct,
		}

		foldersComplete, closer := observe.MessageWithCompletion(fmt.Sprintf("∙ %s - %s:", qp.Category, user))
		defer closer()
		defer close(foldersComplete)

		resolver, err := populateExchangeContainerResolver(ctx, qp)
		if err != nil {
			return nil, errors.Wrap(err, "getting folder cache")
		}

		err = filterContainersAndFillCollections(
			ctx,
			qp,
			collections,
			su,
			resolver,
			scope,
			dps,
			ctrlOpts)

		if err != nil {
			return nil, errors.Wrap(err, "filling collections")
		}

		foldersComplete <- struct{}{}
		//TODO fillAllCollectiton
	}

	return allCollections, errs.ErrorOrNil()
}
