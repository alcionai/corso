package export

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/observe"
	"github.com/alcionai/corso/src/pkg/fault"
)

func ConsumeExportCollections(
	ctx context.Context,
	exportLocation string,
	expColl []Collectioner,
	errs *fault.Bus,
) error {
	el := errs.Local()

	for _, col := range expColl {
		if el.Failure() != nil {
			break
		}

		folder := filepath.Join(exportLocation, col.BasePath())
		ictx := clues.Add(ctx, "dir_name", folder)

		for item := range col.Items(ictx) {
			if item.Error != nil {
				el.AddRecoverable(ictx, clues.Wrap(item.Error, "getting item"))
				continue
			}

			if err := writeItem(ictx, item, folder); err != nil {
				el.AddRecoverable(
					ictx,
					clues.Wrap(err, "writing item").With("file_name", item.Name))
			}
		}
	}

	return el.Failure()
}

// writeItem writes an ExportItem to disk in the specified folder.
func writeItem(ctx context.Context, item Item, folder string) error {
	name := item.Name
	fpath := filepath.Join(folder, name)

	progReader := observe.ItemSpinner(
		ctx,
		item.Body,
		observe.ItemExportMsg,
		clues.Hide(name))

	defer item.Body.Close()
	defer progReader.Close()

	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return clues.WrapWC(ctx, err, "creating directory")
	}

	// In case the user tries to restore to a non-clean
	// directory, we might run into collisions an fail.
	f, err := os.Create(fpath)
	if err != nil {
		return clues.WrapWC(ctx, err, "creating file")
	}

	_, err = io.Copy(f, progReader)
	if err != nil {
		return clues.WrapWC(ctx, err, "writing data")
	}

	return nil
}
