package streamstore

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/pkg/errors"
)

var _ Streamer = &streamDetails{}

type streamDetails struct {
	kw      *kopia.Wrapper
	tenant  string
	service path.ServiceType
}

// NewDetails creates a new storeStreamer for streaming
// details.Details structs.
func NewDetails(
	kw *kopia.Wrapper,
	tenant string,
	service path.ServiceType,
) *streamDetails {
	return &streamDetails{kw: kw, tenant: tenant, service: service}
}

const (
	// detailsItemName is the name of the stream used to store
	// backup details
	detailsItemName = "details"
	// collectionPurposeDetails is used to indicate
	// what the collection is being used for
	collectionPurposeDetails = "details"
)

// Write persists a `details.Details` object in the stream store
func (ss *streamDetails) Write(ctx context.Context, deets Marshaller, errs *fault.Bus) (string, error) {
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
	dbytes, err := deets.Marshal()
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

// Read reads a `details.Details` object from the kopia repository
func (ss *streamDetails) Read(
	ctx context.Context,
	detailsID string,
	umr Unmarshaller,
	errs *fault.Bus,
) error {
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
		return clues.Stack(err).WithClues(ctx)
	}

	dcs, err := ss.kw.RestoreMultipleItems(
		ctx,
		detailsID,
		[]path.Path{detailsPath},
		&stats.ByteCounter{},
		errs)
	if err != nil {
		return errors.Wrap(err, "retrieving backup details data")
	}

	// Expect only 1 data collection
	if len(dcs) != 1 {
		return clues.New("greater than 1 details collection found").
			WithClues(ctx).
			With("collection_count", len(dcs))
	}

	var (
		dc    = dcs[0]
		found = false
		items = dc.Items(ctx, errs)
	)

	for {
		select {
		case <-ctx.Done():
			return clues.New("context cancelled waiting for backup details data").WithClues(ctx)

		case itemData, ok := <-items:
			if !ok {
				if !found {
					return clues.New("no backup details found").WithClues(ctx)
				}

				return nil
			}

			if err := umr(itemData.ToReader()); err != nil {
				return clues.Wrap(err, "unmarshalling details data").WithClues(ctx)
			}

			found = true
		}
	}
}

// Delete deletes a `details.Details` object from the kopia repository
func (ss *streamDetails) Delete(ctx context.Context, detailsID string) error {
	err := ss.kw.DeleteSnapshot(ctx, detailsID)
	if err != nil {
		return errors.Wrap(err, "deleting backup details")
	}

	return nil
}
