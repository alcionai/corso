package streamstore

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/pkg/errors"
)

var _ Streamer = &streamFaultErrors{}

type streamFaultErrors struct {
	kw      *kopia.Wrapper
	tenant  string
	service path.ServiceType
}

// NewFaultErrors creates a new storeStreamer for streaming
// fault.Error structs.
func NewFaultErrors(
	kw *kopia.Wrapper,
	tenant string,
	service path.ServiceType,
) *streamFaultErrors {
	return &streamFaultErrors{kw: kw, tenant: tenant, service: service}
}

const (
	// faultErrorsName is the name of the stream used to store fault errors
	faultErrorsName = "fault_error"
	// collectionPurposeFaultErrors is used to indicate
	// what the collection is being used for
	collectionPurposeFaultErrors = "fault_error"
)

// Write persists a slice of `fault.Error` objects in the stream store
func (ss *streamFaultErrors) Write(ctx context.Context, mr Marshaller, errs *fault.Bus) (string, error) {
	id, err := write(ctx, ss.tenant, ss.service, collectionPurposeFaultErrors, faultErrorsName, ss.kw, mr, errs)
	if err != nil {
		return "", clues.Wrap(err, "fault errors")
	}

	return id, nil
}

// Read reads a slice of `fault.Error` objects from the kopia repository
func (ss *streamFaultErrors) Read(ctx context.Context, id string, umr Unmarshaller, errs *fault.Bus) error {
	err := read(ctx, ss.tenant, ss.service, collectionPurposeFaultErrors, faultErrorsName, id, ss.kw, umr, errs)
	if err != nil {
		return clues.Wrap(err, "fault errors")
	}

	return nil
}

// Delete deletes a slice of `fault.Error` objects from the kopia repository
func (ss *streamFaultErrors) Delete(ctx context.Context, detailsID string) error {
	err := ss.kw.DeleteSnapshot(ctx, detailsID)
	if err != nil {
		return errors.Wrap(err, "deleting backup fault errors")
	}

	return nil
}
