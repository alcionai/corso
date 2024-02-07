package teamschats

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/canario/src/internal/data"
	"github.com/alcionai/canario/src/pkg/store"
)

func DeserializeMetadataFiles(
	ctx context.Context,
	colls []data.RestoreCollection,
) ([]store.MetadataFile, error) {
	return nil, clues.New("no metadata stored for this service/category")
}
