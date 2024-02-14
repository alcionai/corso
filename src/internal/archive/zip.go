package archive

import (
	"archive/zip"
	"context"
	"io"
	"path"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	// ZipCopyBufferSize is the size of the copy buffer for zip
	// write operations
	ZipCopyBufferSize = 5 * 1024 * 1024
)

type zipCollection struct {
	reader io.ReadCloser
}

func (z zipCollection) BasePath() string {
	return ""
}

func (z zipCollection) Items(ctx context.Context) <-chan export.Item {
	rc := make(chan export.Item, 1)
	defer close(rc)

	rc <- export.Item{
		Name: "Corso_Export_" + dttm.FormatNow(dttm.HumanReadable) + ".zip",
		Body: z.reader,
	}

	return rc
}

// ZipExportCollection takes a list of export collections and zips
// them into a single collection.
func ZipExportCollection(
	ctx context.Context,
	expCollections []export.Collectioner,
) (export.Collectioner, error) {
	if len(expCollections) == 0 {
		return nil, clues.New("no export collections provided")
	}

	reader, writer := io.Pipe()
	wr := zip.NewWriter(writer)

	go func() {
		defer writer.Close()
		defer wr.Close()

		buf := make([]byte, ZipCopyBufferSize)
		counted := 0
		log := logger.Ctx(ctx).
			With("collection_count", len(expCollections))

		for _, ec := range expCollections {
			folder := ec.BasePath()
			items := ec.Items(ctx)

			for item := range items {
				counted++

				// Log every 1000 items that are processed
				if counted%1000 == 0 {
					log.Infow("zipping export items", "count_items", counted)
				}

				err := item.Error
				if err != nil {
					writer.CloseWithError(clues.Wrap(err, "getting export item").With("id", item.ID))
					return
				}

				name := item.Name

				// We assume folder and name to not contain any path separators.
				// Also, this should always use `/` as this is
				// created within a zip file and not written to disk.
				// TODO(meain): Exchange paths might contain a path
				// separator and will have to have special handling.

				//nolint:forbidigo
				f, err := wr.Create(path.Join(folder, name))
				if err != nil {
					writer.CloseWithError(clues.Wrap(err, "creating zip entry").With("name", name).With("id", item.ID))
					return
				}

				_, err = io.CopyBuffer(f, item.Body, buf)
				if err != nil {
					writer.CloseWithError(clues.Wrap(err, "writing zip entry").With("name", name).With("id", item.ID))
					return
				}
			}
		}

		log.Infow("zipped export items", "count_items", counted)
	}()

	return zipCollection{reader}, nil
}
