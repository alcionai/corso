package archive

import (
	"archive/zip"
	"context"
	"io"
	"path"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/export"
)

const (
	// ZipCopyBufferSize is the size of the copy buffer for zip
	// write operations
	// TODO(meain): tweak this value
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
		Data: export.ItemData{
			Name: "Corso_Export_" + dttm.FormatNow(dttm.HumanReadable) + ".zip",
			Body: z.reader,
		},
	}

	return rc
}

// ZipExportCollection takes a list of export collections and zips
// them into a single collection.
func ZipExportCollection(
	ctx context.Context,
	expCollections []export.Collection,
) (export.Collection, error) {
	if len(expCollections) == 0 {
		return nil, clues.New("no export collections provided")
	}

	reader, writer := io.Pipe()
	wr := zip.NewWriter(writer)

	go func() {
		defer writer.Close()
		defer wr.Close()

		buf := make([]byte, ZipCopyBufferSize)

		for _, ec := range expCollections {
			folder := ec.BasePath()
			items := ec.Items(ctx)

			for item := range items {
				err := item.Error
				if err != nil {
					writer.CloseWithError(clues.Wrap(err, "getting export item").With("id", item.ID))
					return
				}

				name := item.Data.Name

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

				_, err = io.CopyBuffer(f, item.Data.Body, buf)
				if err != nil {
					writer.CloseWithError(clues.Wrap(err, "writing zip entry").With("name", name).With("id", item.ID))
					return
				}
			}
		}
	}()

	return zipCollection{reader}, nil
}
