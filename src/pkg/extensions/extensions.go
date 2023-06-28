package extensions

import (
	"bytes"
	"io"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

// Temporary, testing purposes only
type MockExtension struct {
	numBytes int
	Data     map[string]any
}

func (me *MockExtension) WrapItem(
	_ details.ItemInfo,
	rc io.ReadCloser,
) (io.ReadCloser, error) {
	p := make([]byte, 64*1024)

	for {
		n, err := rc.Read(p)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		me.numBytes += n
	}

	return rc, nil
}

func (me *MockExtension) OutputData() map[string]any {
	me.Data["numBytes"] = me.numBytes
	return me.Data
}

func NewMockExtension() CorsoItemExtension {
	return &MockExtension{
		Data: map[string]any{},
	}
}

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
}

type extensionHandler struct {
	info           details.ItemInfo
	extData        *details.ExtensionInfo
	innerRc        io.ReadCloser
	itemExtensions []CorsoItemExtension
}

type ExtensionHandlerFactory func(
	info details.ItemInfo,
	extData *details.ExtensionInfo,
	rc io.ReadCloser,
	factory []CorsoItemExtensionFactory,
) (ExtensionHandler, error)

func NewExtensionHandler(
	info details.ItemInfo,
	extData *details.ExtensionInfo,
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
		extData:        extData,
		innerRc:        rc,
		itemExtensions: itemExtensions,
	}, nil
}

// TODO: more robust error handling as per io.Reader guidelines
// and best practices followed by implementations of io.Reader
func (ef *extensionHandler) Read(p []byte) (int, error) {
	n, err := ef.innerRc.Read(p)
	if err != nil {
		return n, err
	}

	// TODO: Why not just send ioReader instead of ioReadCloser?
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

	// TODO: handle errors from extensions
	for _, ext := range ef.itemExtensions {
		for k, v := range ext.OutputData() {
			// Last writer wins on collisions
			ef.extData.Data[k] = v
		}
	}

	return nil
}
