package exchange

import (
	"context"
	"encoding/json"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// MetadataFileNames produces the category-specific set of filenames used to
// store graph metadata such as delta tokens and folderID->path references.
func MetadataFileNames(cat path.CategoryType) [][]string {
	switch cat {
	// TODO: should this include events?
	case path.EmailCategory, path.ContactsCategory:
		return [][]string{{metadata.DeltaURLsFileName}, {metadata.PreviousPathFileName}}
	default:
		return [][]string{{metadata.PreviousPathFileName}}
	}
}

// ParseMetadataCollections produces a map of structs holding delta
// and path lookup maps.
func ParseMetadataCollections(
	ctx context.Context,
	colls []data.RestoreCollection,
) (metadata.CatDeltaPaths, bool, error) {
	// cdp stores metadata
	cdp := metadata.CatDeltaPaths{
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

	// errors from metadata items should not stop the backup,
	// but it should prevent us from using previous backups
	errs := fault.New(true)

	for _, coll := range colls {
		var (
			breakLoop bool
			items     = coll.Items(ctx, errs)
			category  = coll.FullPath().Category()
		)

		for {
			select {
			case <-ctx.Done():
				return nil, false, clues.Wrap(ctx.Err(), "parsing collection metadata").WithClues(ctx)

			case item, ok := <-items:
				if !ok || errs.Failure() != nil {
					breakLoop = true
					break
				}

				var (
					m    = map[string]string{}
					cdps = cdp[category]
				)

				err := json.NewDecoder(item.ToReader()).Decode(&m)
				if err != nil {
					return nil, false, clues.New("decoding metadata json").WithClues(ctx)
				}

				switch item.ID() {
				case metadata.PreviousPathFileName:
					if _, ok := found[category][metadata.PathKey]; ok {
						return nil, false, clues.Wrap(clues.New(category.String()), "multiple versions of path metadata").WithClues(ctx)
					}

					for k, p := range m {
						cdps.AddPath(k, p)
					}

					found[category][metadata.PathKey] = struct{}{}

				case metadata.DeltaURLsFileName:
					if _, ok := found[category][metadata.DeltaKey]; ok {
						return nil, false, clues.Wrap(clues.New(category.String()), "multiple versions of delta metadata").WithClues(ctx)
					}

					for k, d := range m {
						cdps.AddDelta(k, d)
					}

					found[category][metadata.DeltaKey] = struct{}{}
				}

				cdp[category] = cdps
			}

			if breakLoop {
				break
			}
		}
	}

	if errs.Failure() != nil {
		logger.CtxErr(ctx, errs.Failure()).Info("reading metadata collection items")

		return metadata.CatDeltaPaths{
			path.ContactsCategory: {},
			path.EmailCategory:    {},
			path.EventsCategory:   {},
		}, false, nil
	}

	// Remove any entries that contain a path or a delta, but not both.
	// That metadata is considered incomplete, and needs to incur a
	// complete backup on the next run.
	for _, dps := range cdp {
		for k, dp := range dps {
			if len(dp.Path) == 0 {
				delete(dps, k)
			}
		}
	}

	return cdp, true, nil
}
