package exchange

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/path"
)

// MetadataFileNames produces the category-specific set of filenames used to
// store graph metadata such as delta tokens and folderID->path references.
func MetadataFileNames(cat path.CategoryType) []string {
	switch cat {
	case path.EmailCategory, path.ContactsCategory:
		return []string{graph.DeltaTokenFileName, graph.PreviousPathFileName}
	default:
		return []string{graph.PreviousPathFileName}
	}
}

type CatDeltaPaths map[path.CategoryType]DeltaPaths

type DeltaPaths struct {
	deltas map[string]string
	paths  map[string]string
}

func makeDeltaPaths() DeltaPaths {
	return DeltaPaths{
		deltas: map[string]string{},
		paths:  map[string]string{},
	}
}

// ParseMetadataCollections produces a map of structs holding delta
// and path lookup maps.
func ParseMetadataCollections(
	ctx context.Context,
	colls []data.Collection,
) (CatDeltaPaths, error) {
	cdp := CatDeltaPaths{
		path.ContactsCategory: makeDeltaPaths(),
		path.EmailCategory:    makeDeltaPaths(),
		path.EventsCategory:   makeDeltaPaths(),
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

				m := map[string]string{}
				cdps := cdp[category]

				err := json.NewDecoder(item.ToReader()).Decode(&m)
				if err != nil {
					return nil, errors.New("decoding metadata json")
				}

				switch item.UUID() {
				case graph.PreviousPathFileName:
					if len(cdps.paths) > 0 {
						return nil, errors.Errorf("multiple versions of %s path metadata", category)
					}

					cdps.paths = m

				case graph.DeltaTokenFileName:
					if len(cdps.deltas) > 0 {
						return nil, errors.Errorf("multiple versions of %s delta metadata", category)
					}

					cdps.deltas = m
				}

				cdp[category] = cdps
			}

			if breakLoop {
				break
			}
		}
	}

	return cdp, nil
}
