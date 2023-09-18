package drive

import (
	"context"
	"strings"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
)

func NewExportCollection(
	baseDir string,
	backingCollection []data.RestoreCollection,
	backupVersion int,
) export.Collectioner {
	return export.BaseCollection{
		BaseDir:           baseDir,
		BackingCollection: backingCollection,
		BackupVersion:     backupVersion,
		Stream:            streamItems,
	}
}

// streamItems streams the streamItems in the backingCollection into the export stream chan
func streamItems(
	ctx context.Context,
	drc []data.RestoreCollection,
	backupVersion int,
	cec control.ExportConfig,
	ch chan<- export.Item,
) {
	defer close(ch)

	errs := fault.New(false)

	for _, rc := range drc {
		for item := range rc.Items(ctx, errs) {
			itemUUID := item.ID()
			if isMetadataFile(itemUUID, backupVersion) {
				continue
			}

			name, err := getItemName(ctx, itemUUID, backupVersion, rc)

			ch <- export.Item{
				ID:    itemUUID,
				Name:  name,
				Body:  item.ToReader(),
				Error: err,
			}
		}

		items, recovered := errs.ItemsAndRecovered()

		// Return all the items that we failed to source from the persistence layer
		for _, err := range items {
			ch <- export.Item{
				ID:    err.ID,
				Error: &err,
			}
		}

		for _, err := range recovered {
			ch <- export.Item{
				Error: err,
			}
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
