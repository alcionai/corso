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

// ParseMetadataCollections produces two maps:
// 1- paths: folderID->filePath, used to look up previous folder pathing
// in case of a name change or relocation.
// 2- deltas: folderID->deltaToken, used to look up previous delta token
// retrievals.
func ParseMetadataCollections(
	ctx context.Context,
	colls []data.Collection,
) (map[string]string, map[string]string, error) {
	var (
		paths  = map[string]string{}
		deltas = map[string]string{}
	)

	for _, coll := range colls {
		items := coll.Items()

		for {
			var breakLoop bool

			select {
			case <-ctx.Done():
				return nil, nil, errors.Wrap(ctx.Err(), "parsing collection metadata")
			case item, ok := <-items:
				if !ok {
					breakLoop = true
					break
				}

				switch item.UUID() {
				// case graph.PreviousPathFileName:
				case graph.DeltaTokenFileName:
					err := json.NewDecoder(item.ToReader()).Decode(&deltas)
					if err != nil {
						return nil, nil, errors.New("parsing delta token map")
					}

					breakLoop = true
				}
			}

			if breakLoop {
				break
			}
		}
	}

	return paths, deltas, nil
}
