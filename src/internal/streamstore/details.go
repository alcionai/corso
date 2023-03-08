package streamstore

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/kopia"
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
func (ss *streamDetails) Write(ctx context.Context, mr Marshaller, errs *fault.Bus) (string, error) {
	id, err := write(ctx, ss.tenant, ss.service, collectionPurposeDetails, detailsItemName, ss.kw, mr, errs)
	if err != nil {
		return "", clues.Wrap(err, "backup details")
	}

	return id, nil
}

// Read reads a `details.Details` object from the kopia repository
func (ss *streamDetails) Read(ctx context.Context, id string, umr Unmarshaller, errs *fault.Bus) error {
	err := read(ctx, ss.tenant, ss.service, collectionPurposeDetails, detailsItemName, id, ss.kw, umr, errs)
	if err != nil {
		return clues.Wrap(err, "backup details")
	}

	return nil
}

// Delete deletes a `details.Details` object from the kopia repository
func (ss *streamDetails) Delete(ctx context.Context, detailsID string) error {
	err := ss.kw.DeleteSnapshot(ctx, detailsID)
	if err != nil {
		return errors.Wrap(err, "deleting backup details")
	}

	return nil
}
