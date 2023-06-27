package extensions

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

type CorsoItemExtension interface {
	// TODO: iteminfo is duplicated across multiple calls. Add Initialize()?
	WrapItem(
		info details.ItemInfo,
		rc io.ReadCloser,
	) (io.ReadCloser, error)
	OutputData() map[string]any
}

type CorsoItemExtensionFactory func() CorsoItemExtension

var _ io.ReadCloser = &extensionFramework{}

// TODO: need a better name
var _ GetExtensionDataer = &extensionFramework{}

// TODO: Do we need thread safety here?
// Given that we don't process multiple chunks of the same item in
// parallel, I think we don't need it.
type extensionFramework struct {
	info           details.ItemInfo
	innerRc        io.ReadCloser
	itemExtensions []CorsoItemExtension
	extensionStore map[string]any
}

func newExtension(
	info details.ItemInfo,
	rc io.ReadCloser,
	factory []CorsoItemExtensionFactory,
) (*extensionFramework, error) {
	itemExtensions := make([]CorsoItemExtension, len(factory))

	for i, f := range factory {
		if f == nil {
			return nil, errors.New("nil extension factory")
		}

		itemExtensions[i] = f()
	}

	return &extensionFramework{
		info:           info,
		innerRc:        rc,
		itemExtensions: itemExtensions,
		extensionStore: map[string]any{},
	}, nil
}

// TODO: more robust error handling as per io.Reader guidelines
func (ef *extensionFramework) Read(p []byte) (int, error) {
	n, err := ef.innerRc.Read(p)
	if err != nil {
		return n, err
	}

	// Why not just send ioReader instead of ioReadCloser?
	rc := io.NopCloser(bytes.NewReader(p[:n]))

	// Call extensions iteratively
	for _, ext := range ef.itemExtensions {
		rc, err = ext.WrapItem(ef.info, rc)
		if err != nil {
			return 0, err
		}
	}

	return n, err
}

func (ef *extensionFramework) Close() error {
	err := ef.innerRc.Close()
	if err != nil {
		return err
	}

	// Call outputdata on extensions and store in map
	// TODO: handle collisions if we decide on flat hierarchy of ext kv store.
	// or we can do per extension kv store

	for _, ext := range ef.itemExtensions {
		for k, v := range ext.OutputData() {
			ef.extensionStore[k] = v
		}
	}

	return nil
}

type GetExtensionDataer interface {
	GetExtensionData(
		ctx context.Context,
	) (map[string]any, error)
}

func (ef *extensionFramework) GetExtensionData(
	ctx context.Context,
) (map[string]any, error) {
	if ef == nil {
		return nil, clues.New("nil extension")
	}

	if len(ef.extensionStore) == 0 {
		return nil, clues.New("no extension data")
	}

	return ef.extensionStore, nil
}
