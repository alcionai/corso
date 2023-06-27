package extensions

import (
	"bytes"
	"errors"
	"io"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

type BackupItemExtension interface {
	// TODO: Add error as a return value
	WrapItem(info details.ItemInfo, rc io.ReadCloser) io.ReadCloser
	OutputData() map[string]any
}

type BackupItemExtensionFactory func() BackupItemExtension

// TODO: Do we need thread safety here?
// Given that we don't process multiple chunks of the same item in
// parallel, I think we don't need it.
type extension struct {
	info       details.ItemInfo
	innerRc    io.ReadCloser
	extensions []BackupItemExtension
	extStore   map[string]any
}

func newExtension(
	info details.ItemInfo,
	rc io.ReadCloser,
	extFactories []BackupItemExtensionFactory,
) (*extension, error) {
	extensions := make([]BackupItemExtension, len(extFactories))

	for i, f := range extFactories {
		if f == nil {
			return nil, errors.New("nil extension factory")
		}

		extensions[i] = f()
	}

	return &extension{
			info:       info,
			innerRc:    rc,
			extensions: extensions,
			extStore:   map[string]any{},
		},
		nil
}

func (e *extension) Read(p []byte) (int, error) {
	n, err := e.innerRc.Read(p)
	// TODO: more robust error handling as per io.Reader guidelines
	if err != nil {
		return n, err
	}

	// TODO: handle EOF
	rc := io.NopCloser(bytes.NewReader(p[:n]))

	// Call extensions iteratively
	for _, ext := range e.extensions {
		rc = ext.WrapItem(e.info, rc)
	}

	return n, err
}

func (e *extension) Close() error {
	err := e.innerRc.Close()
	if err != nil {
		return err
	}

	// Call outputdata on extensions and store in map
	// TODO: handle collisions if we decide on flat hierarchy of ext kv store.
	// or we can do per extension kv store

	for _, ext := range e.extensions {
		for k, v := range ext.OutputData() {
			e.extStore[k] = v
		}
	}

	return nil
}

type GetExtensionDataer interface {
	GetExtensionData() (map[string]any, error)
}

func (e *extension) GetExtensionData() (map[string]any, error) {
	if e == nil {
		return nil, errors.New("nil extension")
	}

	if len(e.extStore) == 0 {
		return nil, errors.New("no extension data")
	}

	return e.extStore, nil
}
