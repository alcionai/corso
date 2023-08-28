package drive

import (
	"context"
	"strings"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
)

var _ export.Collection = &ExportCollection{}

// ExportCollection is the implementation of export.ExportCollection for OneDrive
type ExportCollection struct {
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

func NewExportCollection(
	baseDir string,
	backingCollection data.RestoreCollection,
	backupVersion int,
) ExportCollection {
	return ExportCollection{
		baseDir:           baseDir,
		backingCollection: backingCollection,
		backupVersion:     backupVersion,
	}
}

func (ec ExportCollection) BasePath() string {
	return ec.baseDir
}

func (ec ExportCollection) Items(ctx context.Context) <-chan export.Item {
	ch := make(chan export.Item)
	go items(ctx, ec, ch)

	return ch
}

// items converts items in backing collection to export items
func items(ctx context.Context, ec ExportCollection, ch chan<- export.Item) {
	defer close(ch)

	errs := fault.New(false)

	for item := range ec.backingCollection.Items(ctx, errs) {
		itemUUID := item.ID()
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

	eitems, erecovereable := errs.ItemsAndRecovered()

	// Return all the items that we failed to source from the persistence layer
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

		meta, err := FetchAndReadMetadata(ctx, fin, metaName)
		if err != nil {
			return "", clues.Wrap(err, "getting metadata").WithClues(ctx)
		}

		return meta.FileName, nil
	}

	return "", clues.New("invalid item id").WithClues(ctx)
}
