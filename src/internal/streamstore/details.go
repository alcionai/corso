package streamstore

import (
	"context"
	"encoding/json"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/pkg/errors"
)

type streamDetails struct {
	kw      *kopia.Wrapper
	tenant  string
	service path.ServiceType
}

func NewDetails(
	kw *kopia.Wrapper,
	tenant string,
	service path.ServiceType,
) *streamDetails {
	return &streamDetails{kw: kw, tenant: tenant, service: service}
}

// WriteBackupDetails persists a `details.Details`
// object in the stream store
func (ss *streamDetails) WriteBackupDetails(
	ctx context.Context,
	backupDetails *details.Details,
	errs *fault.Bus,
) (string, error) {
	// construct the path of the container for the `details` item
	p, err := path.Builder{}.
		ToStreamStorePath(
			ss.tenant,
			collectionPurposeDetails,
			ss.service,
			false)
	if err != nil {
		return "", clues.Stack(err).WithClues(ctx)
	}

	// TODO: We could use an io.Pipe here to avoid a double copy but that
	// makes error handling a bit complicated
	dbytes, err := json.Marshal(backupDetails)
	if err != nil {
		return "", clues.Wrap(err, "marshalling backup details").WithClues(ctx)
	}

	dc := &streamCollection{
		folderPath: p,
		item: &streamItem{
			name: detailsItemName,
			data: dbytes,
		},
	}

	backupStats, _, _, err := ss.kw.BackupCollections(
		ctx,
		nil,
		[]data.BackupCollection{dc},
		nil,
		nil,
		false,
		errs)
	if err != nil {
		return "", errors.Wrap(err, "storing details in repository")
	}

	return backupStats.SnapshotID, nil
}

// ReadBackupDetails reads the specified details object
// from the kopia repository
func (ss *streamDetails) ReadBackupDetails(
	ctx context.Context,
	detailsID string,
	errs *fault.Bus,
) (*details.Details, error) {
	// construct the path for the `details` item
	detailsPath, err := path.Builder{}.
		Append(detailsItemName).
		ToStreamStorePath(
			ss.tenant,
			collectionPurposeDetails,
			ss.service,
			true,
		)
	if err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	var bc stats.ByteCounter

	dcs, err := ss.kw.RestoreMultipleItems(ctx, detailsID, []path.Path{detailsPath}, &bc, errs)
	if err != nil {
		return nil, errors.Wrap(err, "retrieving backup details data")
	}

	// Expect only 1 data collection
	if len(dcs) != 1 {
		return nil, clues.New("greater than 1 details data collection found").
			WithClues(ctx).
			With("collection_count", len(dcs))
	}

	dc := dcs[0]

	var d details.Details

	found := false
	items := dc.Items(ctx, errs)

	for {
		select {
		case <-ctx.Done():
			return nil, clues.New("context cancelled waiting for backup details data").WithClues(ctx)

		case itemData, ok := <-items:
			if !ok {
				if !found {
					return nil, clues.New("no backup details found").WithClues(ctx)
				}

				return &d, nil
			}

			err := json.NewDecoder(itemData.ToReader()).Decode(&d)
			if err != nil {
				return nil, clues.Wrap(err, "decoding details data").WithClues(ctx)
			}

			found = true
		}
	}
}

// DeleteBackupDetails deletes the specified details object from the kopia repository
func (ss *streamDetails) DeleteBackupDetails(
	ctx context.Context,
	detailsID string,
) error {
	err := ss.kw.DeleteSnapshot(ctx, detailsID)
	if err != nil {
		return errors.Wrap(err, "deleting backup details")
	}

	return nil
}
