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

	// backingCollections will, in most cases contain just one
	// collection. However, in cases where we have to combine multiple
	// collections, like when generating pst files for outlook, this
	// will contain multiple collections.
	backingCollections []data.RestoreCollection

	// version is the version of the backup this collection was part
	// of. This is required to figure out how to get the name of the
	// item.
	version int
}

func (ec exportCollection) GetBasePath() string {
	return ec.baseDir
}

func (ec exportCollection) GetItems(ctx context.Context) <-chan export.Item {
	ch := make(chan export.Item)

	go func() {
		defer close(ch)

		errs := fault.New(false)

		// There will only be a single item in the backingCollections
		// for OneDrive
		for _, c := range ec.backingCollections {
			for item := range c.Items(ctx, errs) {
				itemUUID := item.UUID()
				if isMetadataFile(itemUUID, ec.version) {
					continue
				}

				name, err := getItemName(ctx, itemUUID, ec.version, c)

				ch <- export.Item{
					ID: itemUUID,
					Data: export.ItemData{
						Name: name,
						Body: item.ToReader(),
					},
					Error: err,
				}
			}
		}

		// Return all the items that we failed to get from kopia at the end
		for _, err := range errs.Errors().Items {
			ch <- export.Item{
				ID:    err.ID,
				Error: &err,
			}
		}
	}()

	return ch
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

// getItemName is used to get the name of the item as how we get the
// name depends on the version of the backup.
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
	el := errs.Local()

	ec := make([]export.Collection, 0, len(dcs))

	for _, dc := range dcs {
		drivePath, err := path.ToDrivePath(dc.FullPath())
		if err != nil {
			return nil, clues.Wrap(err, "creating drive path").WithClues(ctx)
		}

		baseDir := &path.Builder{}
		baseDir = baseDir.Append(drivePath.Folders...)

		ec = append(ec, exportCollection{
			baseDir:            baseDir.String(),
			backingCollections: []data.RestoreCollection{dc},
			version:            backupVersion,
		})
	}

	return ec, el.Failure()
}
