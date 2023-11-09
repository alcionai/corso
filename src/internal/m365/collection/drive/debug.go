package drive

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	bupMD "github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/store"
)

func DeserializeMetadataFiles(
	ctx context.Context,
	colls []data.RestoreCollection,
) ([]store.MetadataFile, error) {
	deltas, prevs, _, err := deserializeMetadata(ctx, colls)

	files := []store.MetadataFile{
		{
			Name: bupMD.PreviousPathFileName,
			Data: prevs,
		},
		{
			Name: bupMD.DeltaURLsFileName,
			Data: deltas,
		},
	}

	return files, clues.Stack(err).OrNil()
}
