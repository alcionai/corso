package kopia

import (
	"context"
	"io"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/fs"

	"github.com/alcionai/canario/src/internal/common/readers"
	"github.com/alcionai/canario/src/internal/data"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/logger"
	"github.com/alcionai/canario/src/pkg/path"
)

var (
	_ data.RestoreCollection = &kopiaDataCollection{}
	_ data.Item              = &kopiaDataStream{}
	_ data.ItemSize          = &kopiaDataStream{}
)

type kopiaDataCollection struct {
	path            path.Path
	dir             fs.Directory
	items           []string
	counter         ByteCounter
	expectedVersion readers.SerializationVersion
}

func (kdc *kopiaDataCollection) Items(
	ctx context.Context,
	errs *fault.Bus,
) <-chan data.Item {
	var (
		res       = make(chan data.Item)
		el        = errs.Local()
		loadCount = 0
	)

	go func() {
		defer close(res)

		for _, item := range kdc.items {
			s, err := kdc.FetchItemByName(ctx, item)
			if err != nil {
				el.AddRecoverable(ctx, clues.WrapWC(ctx, err, "fetching item").
					Label(fault.LabelForceNoBackupCreation))

				continue
			}

			loadCount++
			if loadCount%1000 == 0 {
				logger.Ctx(ctx).Infow(
					"loading items from kopia",
					"loaded_items", loadCount)
			}

			res <- s
		}

		logger.Ctx(ctx).Infow(
			"done loading items from kopia",
			"loaded_items", loadCount)
	}()

	return res
}

func (kdc kopiaDataCollection) FullPath() path.Path {
	return kdc.path
}

// Fetch returns the file with the given name from the collection as a
// data.Item. Returns a data.ErrNotFound error if the file isn't in the
// collection.
func (kdc kopiaDataCollection) FetchItemByName(
	ctx context.Context,
	name string,
) (data.Item, error) {
	ctx = clues.Add(ctx, "item_name", clues.Hide(name))

	if kdc.dir == nil {
		return nil, clues.New("no snapshot directory")
	}

	if len(name) == 0 {
		return nil, clues.WrapWC(ctx, ErrNoRestorePath, "unknown item")
	}

	e, err := kdc.dir.Child(ctx, encodeAsPath(name))
	if err != nil {
		if isErrEntryNotFound(err) {
			err = clues.Stack(data.ErrNotFound, err)
		}

		return nil, clues.WrapWC(ctx, err, "getting item")
	}

	f, ok := e.(fs.File)
	if !ok {
		return nil, clues.NewWC(ctx, "object is not a file")
	}

	size := f.Size() - int64(readers.VersionFormatSize)
	if size < 0 {
		logger.Ctx(ctx).Infow("negative file size; resetting to 0", "file_size", size)

		size = 0
	}

	if kdc.counter != nil {
		kdc.counter.Count(size)
	}

	r, err := f.Open(ctx)
	if err != nil {
		return nil, clues.WrapWC(ctx, err, "opening file")
	}

	// TODO(ashmrtn): Remove this when individual services implement checks for
	// version and deleted items.
	rr, err := readers.NewVersionedRestoreReader(r)
	if err != nil {
		return nil, clues.StackWC(ctx, err)
	}

	if rr.Format().Version != kdc.expectedVersion {
		return nil, clues.NewWC(ctx, "unexpected data format").
			With(
				"read_version", rr.Format().Version,
				"expected_version", kdc.expectedVersion)
	}

	// This is a conservative check, but we shouldn't be seeing items that were
	// deleted in flight during restores because there's no way to select them.
	if rr.Format().DelInFlight {
		return nil, clues.NewWC(ctx, "selected item marked as deleted in flight")
	}

	return &kopiaDataStream{
		id:     name,
		reader: rr,
		size:   size,
	}, nil
}

type kopiaDataStream struct {
	reader io.ReadCloser
	id     string
	size   int64
}

func (kds kopiaDataStream) ToReader() io.ReadCloser {
	return kds.reader
}

func (kds kopiaDataStream) ID() string {
	return kds.id
}

func (kds kopiaDataStream) Deleted() bool {
	return false
}

func (kds kopiaDataStream) Size() int64 {
	return kds.size
}
