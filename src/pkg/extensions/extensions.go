package extensions

import (
	"bytes"
	"context"
	"io"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

// Extension client interface
type CorsoItemExtension interface {
	// TODO: iteminfo is duplicated across multiple calls. Add Initialize()?
	WrapItem(
		info details.ItemInfo,
		rc io.ReadCloser,
	) (io.ReadCloser, error)
	OutputData() map[string]any
}

type CorsoItemExtensionFactory func() CorsoItemExtension

var _ ExtensionHandler = &extensionHandler{}

type ExtensionHandler interface {
	io.ReadCloser
	GetExtensionDataer
}

type extensionHandler struct {
	info           details.ItemInfo
	innerRc        io.ReadCloser
	itemExtensions []CorsoItemExtension
	extensionStore map[string]any
}

type ExtensionHandlerFactory func(
	info details.ItemInfo,
	rc io.ReadCloser,
	factory []CorsoItemExtensionFactory,
) (ExtensionHandler, error)

func newExtensionHandler(
	info details.ItemInfo,
	rc io.ReadCloser,
	factory []CorsoItemExtensionFactory,
) (ExtensionHandler, error) {
	itemExtensions := make([]CorsoItemExtension, len(factory))

	for i, f := range factory {
		if f == nil {
			return nil, clues.New("nil extension factory")
		}

		itemExtensions[i] = f()
	}

	return &extensionHandler{
		info:           info,
		innerRc:        rc,
		itemExtensions: itemExtensions,
		extensionStore: map[string]any{},
	}, nil
}

// TODO: more robust error handling as per io.Reader guidelines
// and best practices followed by implementations of io.Reader
func (ef *extensionHandler) Read(p []byte) (int, error) {
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

func (ef *extensionHandler) Close() error {
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

func (ef *extensionHandler) GetExtensionData(
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
