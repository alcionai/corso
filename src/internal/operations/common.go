package operations

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/store"
)

func getBackupAndDetailsFromID(
	ctx context.Context,
	tenant string,
	backupID model.StableID,
	service path.ServiceType,
	ms *store.Wrapper,
	kw *kopia.Wrapper,
) (*backup.Backup, *details.Details, error) {
	dID, bup, err := ms.GetDetailsIDFromBackupID(ctx, backupID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting backup details ID")
	}

	deets, err := streamstore.New(
		kw,
		tenant,
		service,
	).ReadBackupDetails(ctx, dID)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting backup details data")
	}

	return bup, deets, nil
}
