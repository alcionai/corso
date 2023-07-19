package onedrive

import (
	"context"
	"strings"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ export.Collection = &exportCollection{}

// exportCollection is the implementation of export.ExportCollection for OneDrive
type exportCollection struct {
	// baseDir contains the path of the collection
	baseDir string

	// backingCollection is the restore collection from which we will
	// create the export collection.
	backingCollection data.RestoreCollection

	// backupVersion is the backupVersion of the backup this collection was part
	// of. This is required to figure out how to get the name of the
	// item.
	backupVersion int
}

func (ec exportCollection) BasePath() string {
	return ec.baseDir
}

func (ec exportCollection) Items(ctx context.Context) <-chan export.Item {
	ch := make(chan export.Item)
	go items(ctx, ec, ch)

	return ch
}

// items converts items in backing collection to export items
func items(ctx context.Context, ec exportCollection, ch chan<- export.Item) {
	defer close(ch)

	errs := fault.New(false)

	// There will only be a single item in the backingCollections
	// for OneDrive
	for item := range ec.backingCollection.Items(ctx, errs) {
		itemUUID := item.UUID()
		if isMetadataFile(itemUUID, ec.backupVersion) {
			continue
		}

		name, err := getItemName(ctx, itemUUID, ec.backupVersion, ec.backingCollection)

		ch <- export.Item{
			ID: itemUUID,
			Data: export.ItemData{
				Name: name,
				Body: item.ToReader(),
			},
			Error: err,
		}
	}

	eitems, erecovereable := errs.ErrorEntries()

	// Return all the items that we failed to get from kopia at the end
	for _, err := range eitems {
		ch <- export.Item{
			ID:    err.ID,
			Error: &err,
		}
	}

	for _, ec := range erecovereable {
		ch <- export.Item{
			Error: ec,
		}
	}
}

// isMetadataFile is used to determine if a path corresponds to a
// metadata file.  This is OneDrive specific logic and depends on the
// version of the backup unlike metadata.IsMetadataFile which only has
// to be concerned about the current version.
func isMetadataFile(id string, backupVersion int) bool {
	if backupVersion < version.OneDrive1DataAndMetaFiles {
		return false
	}

	return strings.HasSuffix(id, metadata.MetaFileSuffix) ||
		strings.HasSuffix(id, metadata.DirMetaFileSuffix)
}

// getItemName is used to get the name of the item.
// How we get the name depends on the version of the backup.
func getItemName(
	ctx context.Context,
	id string,
	backupVersion int,
	fin data.FetchItemByNamer,
) (string, error) {
	if backupVersion < version.OneDrive1DataAndMetaFiles {
		return id, nil
	}

	if backupVersion < version.OneDrive5DirMetaNoName {
		return strings.TrimSuffix(id, metadata.DataFileSuffix), nil
	}

	if strings.HasSuffix(id, metadata.DataFileSuffix) {
		trimmedName := strings.TrimSuffix(id, metadata.DataFileSuffix)
		metaName := trimmedName + metadata.MetaFileSuffix

		meta, err := fetchAndReadMetadata(ctx, fin, metaName)
		if err != nil {
			return "", clues.Wrap(err, "getting metadata").WithClues(ctx)
		}

		return meta.FileName, nil
	}

	return "", clues.New("invalid item id").WithClues(ctx)
}

// ExportRestoreCollections will create the export collections for the
// given restore collections.
func ExportRestoreCollections(
	ctx context.Context,
	backupVersion int,
	exportCfg control.ExportConfig,
	opts control.Options,
	dcs []data.RestoreCollection,
	deets *details.Builder,
	errs *fault.Bus,
) ([]export.Collection, error) {
	var (
		el = errs.Local()
		ec = make([]export.Collection, 0, len(dcs))
	)

	for _, dc := range dcs {
		drivePath, err := path.ToDrivePath(dc.FullPath())
		if err != nil {
			return nil, clues.Wrap(err, "transforming path to drive path").WithClues(ctx)
		}

		baseDir := path.Builder{}.Append(drivePath.Folders...)

		ec = append(ec, exportCollection{
			baseDir:           baseDir.String(),
			backingCollection: dc,
			backupVersion:     backupVersion,
		})
	}

	return ec, el.Failure()
}
