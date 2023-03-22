package mock

import (
	"bytes"
	"context"
	"io"

	"github.com/pkg/errors"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/streamstore"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
)

var _ streamstore.Streamer = &Streamer{}

type Streamer struct {
	Deets  map[string]*details.Details
	Errors map[string]*fault.Errors
}

func (ms Streamer) Collect(context.Context, streamstore.Collectable) error {
	return clues.New("not implented")
}

func (ms Streamer) Read(
	ctx context.Context,
	snapshotID string,
	col streamstore.Collectable,
	errs *fault.Bus,
) error {
	var mr streamstore.Marshaller

	switch col.Type {
	case streamstore.DetailsType:
		mr = ms.Deets[snapshotID]
	case streamstore.FaultErrorsType:
		mr = ms.Errors[snapshotID]
	default:
		return clues.New("unknown type: " + col.Type)
	}

	if mr == nil {
		return errors.Errorf("no marshaller for %s ID %s", col.Type, snapshotID)
	}

	bs, err := mr.Marshal()
	if err != nil {
		return err
	}

	return col.Unmr(io.NopCloser(bytes.NewReader(bs)))
}

func (ms Streamer) Write(context.Context, *fault.Bus) (string, error) {
	return "", clues.New("not implented")
}

func (ms Streamer) Delete(context.Context, string) error {
	return clues.New("not implented")
}
