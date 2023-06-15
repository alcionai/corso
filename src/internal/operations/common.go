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

	deets, err := getDetailsFromBackup(ctx, bup, detailsStore, errs)
	if err != nil {
		return nil, nil, clues.Stack(err)
	}

	return bup, deets, nil
}

func getDetailsFromBackup(
	ctx context.Context,
	bup *backup.Backup,
	detailsStore streamstore.Reader,
	errs *fault.Bus,
) (*details.Details, error) {
	var (
		deets details.Details
		umt   = streamstore.DetailsReader(details.UnmarshalTo(&deets))
		ssid  = bup.StreamStoreID
	)

	if len(ssid) == 0 {
		ssid = bup.DetailsID
	}

	if len(ssid) == 0 {
		return nil, clues.New("no details or errors in backup").WithClues(ctx)
	}

	if err := detailsStore.Read(ctx, ssid, umt, errs); err != nil {
		return nil, clues.Wrap(err, "reading backup data from streamstore")
	}

	return &deets, nil
}
