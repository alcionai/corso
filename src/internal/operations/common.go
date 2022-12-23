package operations

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/store"
)

type detailsReader interface {
	ReadBackupDetails(ctx context.Context, detailsID string) (*details.Details, error)
}

func getBackupAndDetailsFromID(
	ctx context.Context,
	backupID model.StableID,
	ms *store.Wrapper,
	detailsStore detailsReader,
) (*backup.Backup, *details.Details, error) {
	dID, bup, err := ms.GetDetailsIDFromBackupID(ctx, backupID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting backup details ID")
	}

	deets, err := detailsStore.ReadBackupDetails(ctx, dID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting backup details data")
	}

	return bup, deets, nil
}
