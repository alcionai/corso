package repository

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

type base struct {
	snapshotID manifest.ID
	reasons    []identity.Reasoner
}

func (b base) GetReasons() []identity.Reasoner {
	return b.reasons
}

func (b base) GetSnapshotID() manifest.ID {
	return b.snapshotID
}

type mdDeserialize func(
	ctx context.Context,
	metadataCollections []data.RestoreCollection,
) ([]store.MetadataFile, error)

// should probably turn into a NewDebug interface like we're
// doing with the other interfaces
type Debugger interface {
	GetBackupMetadata(
		ctx context.Context,
		sel selectors.Selector,
		backupID string,
		errs *fault.Bus,
	) ([]store.MetadataFile, error)
}

// Backups lists backups by ID. Returns as many backups as possible with
// errors for the backups it was unable to retrieve.
func (r repository) GetBackupMetadata(
	ctx context.Context,
	sel selectors.Selector,
	backupID string,
	errs *fault.Bus,
) ([]store.MetadataFile, error) {
	bup, err := r.Backup(ctx, backupID)
	if err != nil {
		return nil, clues.Wrap(err, "looking up backup")
	}

	sel = sel.SetDiscreteOwnerIDName(bup.ResourceOwnerID, bup.ResourceOwnerName)

	reasons, err := sel.Reasons(r.Account.ID(), false)
	if err != nil {
		return nil, clues.Wrap(err, "constructing lookup parameters")
	}

	var (
		rp = r.dataLayer
		dp = r.DataProvider()
	)

	paths, err := dp.GetMetadataPaths(
		ctx,
		rp,
		&base{manifest.ID(bup.SnapshotID), reasons},
		fault.New(true))
	if err != nil {
		return nil, clues.Wrap(err, "retrieving metadata files")
	}

	colls, err := rp.ProduceRestoreCollections(
		ctx,
		bup.SnapshotID,
		paths,
		nil,
		fault.New(true))
	if err != nil {
		return nil, clues.Wrap(err, "looking up metadata file content")
	}

	files, err := dp.DeserializeMetadataFiles(ctx, colls)

	return files, clues.Wrap(err, "deserializing metadata file content").OrNil()
}
