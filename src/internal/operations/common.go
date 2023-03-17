package operations

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

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
		return nil, nil, errors.Wrap(err, "getting backup details ID")
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
		return nil, nil, errors.Wrap(err, "reading backup data from streamstore")
	}

	return bup, &deets, nil
}
