package operations

import (
	"context"

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
	dID, bup, err := ms.GetDetailsIDFromBackupID(ctx, backupID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting backup details ID")
	}

	var (
		deets details.Details
		umt   = details.UnmarshalTo(&deets)
	)

	if err := detailsStore.Read(ctx, dID, umt, errs); err != nil {
		return nil, nil, errors.Wrap(err, "getting backup details data")
	}

	return bup, &deets, nil
}
