package groups

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

// ParseMetadataCollections produces a map of structs holding delta
// and path lookup maps.
func parseMetadataCollections(
	ctx context.Context,
	colls []data.RestoreCollection,
) (metadata.CatDeltaPaths, bool, error) {
	// cdp stores metadata
	cdp := metadata.CatDeltaPaths{
		path.ChannelMessagesCategory: {},
	}

	// found tracks the metadata we've loaded, to make sure we don't
	// fetch overlapping copies.
	found := map[path.CategoryType]map[string]struct{}{
		path.ChannelMessagesCategory: {},
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
					m                    = map[string]string{}
					cdps, wantedCategory = cdp[category]
				)

				// avoid sharepoint site deltapaths
				if !wantedCategory {
					continue
				}

				err := json.NewDecoder(item.ToReader()).Decode(&m)
				if err != nil {
					return nil, false, clues.New("decoding metadata json").WithClues(ctx)
				}

				switch item.ID() {
				case metadata.PreviousPathFileName:
					// no-op at this time, previous paths not needed

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
			path.ChannelMessagesCategory: {},
		}, false, nil
	}

	// Do not remove entries that contain only a path or a delta, but not both.
	// This condition is expected.  Channels only record their path.  Messages
	// only record their deltas.

	return cdp, true, nil
}

// produces a set of id:path pairs from the deltapaths map.
// Each entry in the set will, if not removed, produce a collection
// that will delete the tombstone by path.
func makeTombstones(dps metadata.DeltaPaths) map[string]string {
	r := make(map[string]string, len(dps))

	for id, v := range dps {
		r[id] = v.Path
	}

	return r
}

func pathFromPrevString(ps string) (path.Path, error) {
	p, err := path.FromDataLayerPath(ps, false)
	if err != nil {
		return nil, clues.Wrap(err, "parsing previous path string")
	}

	return p, nil
}
