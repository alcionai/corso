package operations

import (
	"context"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/store"
)

func getBackupFromID(
	ctx context.Context,
	backupID model.StableID,
	ms *store.Wrapper,
) (*backup.Backup, error) {
	bup, err := ms.GetBackup(ctx, backupID)
	if err != nil {
		return nil, clues.Wrap(err, "getting backup")
	}

	return bup, nil
}

func getBackupAndDetailsFromID(
	ctx context.Context,
	backupID model.StableID,
	ms *store.Wrapper,
	detailsStore streamstore.Reader,
	errs *fault.Bus,
) (*backup.Backup, *details.Details, error) {
	bup, err := ms.GetBackup(ctx, backupID)
	if err != nil {
		return nil, nil, clues.Wrap(err, "getting backup")
	}

	var (
		deets details.Details
		umt   = streamstore.DetailsReader(details.UnmarshalTo(&deets))
		ssid  = bup.StreamStoreID
	)

	if len(ssid) == 0 {
		ssid = bup.DetailsID
	}

	if len(ssid) == 0 {
		return bup, nil, clues.New("no details or errors in backup").WithClues(ctx)
	}

	if err := detailsStore.Read(ctx, ssid, umt, errs); err != nil {
		return nil, nil, clues.Wrap(err, "reading backup data from streamstore")
	}

	return bup, &deets, nil
}
