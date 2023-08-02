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
	expColl []Collection,
	errs *fault.Bus,
) error {
	el := errs.Local()

	for _, col := range expColl {
		if el.Failure() != nil {
			break
		}

		folder := filepath.Join(exportLocation, col.BasePath())
		ictx := clues.Add(ctx, "dir_name", folder)

		for item := range col.Items(ctx) {
			if item.Error != nil {
				el.AddRecoverable(ictx, clues.Wrap(item.Error, "getting item").WithClues(ctx))
			}

			if err := writeItem(ictx, item, folder); err != nil {
				el.AddRecoverable(
					ictx,
					clues.Wrap(err, "writing item").With("file_name", item.Data.Name).WithClues(ctx))
			}
		}
	}

	return el.Failure()
}

// writeItem writes an ExportItem to disk in the specified folder.
func writeItem(ctx context.Context, item Item, folder string) error {
	name := item.Data.Name
	fpath := filepath.Join(folder, name)

	progReader, pclose := observe.ItemSpinner(
		ctx,
		item.Data.Body,
		observe.ItemExportMsg,
		clues.Hide(name))

	defer item.Data.Body.Close()
	defer pclose()

	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return clues.Wrap(err, "creating directory")
	}

	// In case the user tries to restore to a non-clean
	// directory, we might run into collisions an fail.
	f, err := os.Create(fpath)
	if err != nil {
		return clues.Wrap(err, "creating file")
	}

	_, err = io.Copy(f, progReader)
	if err != nil {
		return clues.Wrap(err, "writing data")
	}

	return nil
}
