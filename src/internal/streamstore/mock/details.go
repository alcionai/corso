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

var _ streamstore.Streamer = &DetailsStreamer{}

type DetailsStreamer struct {
	Entries map[string]*details.Details
}

func (ds DetailsStreamer) Read(
	ctx context.Context,
	detailsID string,
	umr streamstore.Unmarshaller,
	errs *fault.Bus,
) error {
	r := ds.Entries[detailsID]

	if r == nil {
		return errors.Errorf("no details for ID %s", detailsID)
	}

	bs, err := r.Marshal()
	if err != nil {
		return err
	}

	return umr(io.NopCloser(bytes.NewReader(bs)))
}

func (ds DetailsStreamer) Write(context.Context, streamstore.Marshaller, *fault.Bus) (string, error) {
	return "", clues.New("not implmented")
}

func (ds DetailsStreamer) Delete(context.Context, string) error {
	return clues.New("not implmented")
}
