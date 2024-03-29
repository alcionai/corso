package site

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

func parseListsMetadataCollections(
	ctx context.Context,
	cat path.CategoryType,
	colls []data.RestoreCollection,
) (metadata.DeltaPaths, bool, error) {
	cdp := metadata.CatDeltaPaths{
		cat: {},
	}

	found := map[path.CategoryType]map[string]struct{}{
		cat: {},
	}

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
				return nil, false, clues.WrapWC(ctx, ctx.Err(), "parsing collection metadata")

			case item, ok := <-items:
				if !ok || errs.Failure() != nil {
					breakLoop = true
					break
				}

				var (
					m                    = map[string]string{}
					cdps, wantedCategory = cdp[category]
				)

				if !wantedCategory {
					continue
				}

				err := json.NewDecoder(item.ToReader()).Decode(&m)
				if err != nil {
					return nil, false, clues.WrapWC(ctx, err, "decoding metadata json")
				}

				if item.ID() == metadata.PreviousPathFileName {
					if _, ok := found[category][metadata.PathKey]; ok {
						return nil, false, clues.WrapWC(ctx, err, "multiple versions of path metadata")
					}

					for k, p := range m {
						cdps.AddPath(k, p)
					}

					found[category][metadata.PathKey] = struct{}{}

					cdp[category] = cdps
				}
			}

			if breakLoop {
				break
			}
		}
	}

	if errs.Failure() != nil {
		logger.CtxErr(ctx, errs.Failure()).Info("reading metadata collection items")

		return metadata.DeltaPaths{}, false, nil
	}

	for _, dps := range cdp {
		for k, dp := range dps {
			if len(dp.Path) == 0 {
				delete(dps, k)
			}
		}
	}

	return cdp[cat], true, nil
}

func pathFromPrevString(ps string) (path.Path, error) {
	p, err := path.FromDataLayerPath(ps, false)
	if err != nil {
		return nil, clues.Wrap(err, "parsing previous path string")
	}

	return p, nil
}

func makeTombstones(dps metadata.DeltaPaths) map[string]string {
	r := make(map[string]string, len(dps))

	for id, v := range dps {
		r[id] = v.Path
	}

	return r
}
