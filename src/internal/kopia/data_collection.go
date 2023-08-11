package kopia

import (
	"context"
	"io"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/fs"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	_ data.RestoreCollection = &kopiaDataCollection{}
	_ data.Item              = &kopiaDataStream{}
)

type kopiaDataCollection struct {
	path            path.Path
	dir             fs.Directory
	items           []string
	counter         ByteCounter
	expectedVersion uint32
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
				el.AddRecoverable(ctx, clues.Wrap(err, "fetching item").
					WithClues(ctx).
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
		return nil, clues.Wrap(ErrNoRestorePath, "unknown item").WithClues(ctx)
	}

	e, err := kdc.dir.Child(ctx, encodeAsPath(name))
	if err != nil {
		if isErrEntryNotFound(err) {
			err = clues.Stack(data.ErrNotFound, err)
		}

		return nil, clues.Wrap(err, "getting item").WithClues(ctx)
	}

	f, ok := e.(fs.File)
	if !ok {
		return nil, clues.New("object is not a file").WithClues(ctx)
	}

	size := f.Size() - int64(versionSize)
	if size < 0 {
		logger.Ctx(ctx).Infow("negative file size; resetting to 0", "file_size", size)

		size = 0
	}

	if kdc.counter != nil {
		kdc.counter.Count(size)
	}

	r, err := f.Open(ctx)
	if err != nil {
		return nil, clues.Wrap(err, "opening file").WithClues(ctx)
	}

	return &kopiaDataStream{
		id: name,
		reader: &restoreStreamReader{
			ReadCloser:      r,
			expectedVersion: kdc.expectedVersion,
		},
		size: size,
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
