package extensions

import (
	"context"
	"hash/crc32"
	"io"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
)

// Temporary, testing purposes only
type MockExtension struct {
	// TODO: Add cumlulative crc32 checksum
	numBytes int
	crc32    uint32
	info     details.ItemInfo
	extInfo  *details.ExtensionInfo
	innerRc  io.ReadCloser
	ctx      context.Context
}

func (me *MockExtension) Read(p []byte) (int, error) {
	n, err := me.innerRc.Read(p)
	if err != nil && err != io.EOF {
		logger.CtxErr(me.ctx, err).Error("inner read error")
		return n, err
	}

	me.numBytes += n
	me.crc32 = crc32.Update(me.crc32, crc32.IEEETable, p[:n])

	if err == io.EOF {
		logger.Ctx(me.ctx).Info("mock extension reached EOF")
		me.extInfo.Data["numBytes"] = me.numBytes
		me.extInfo.Data["crc32"] = me.crc32
	}

	return n, err
}

func (me *MockExtension) Close() error {
	err := me.innerRc.Close()
	if err != nil {
		return err
	}

	me.extInfo.Data["numBytes"] = me.numBytes
	me.extInfo.Data["crc32"] = me.crc32
	logger.Ctx(me.ctx).Infow(
		"mock extension closed",
		"numBytes", me.numBytes, "crc32", me.crc32)

	return nil
}

func NewMockExtension(
	ctx context.Context,
	rc io.ReadCloser,
	info details.ItemInfo,
	extInfo *details.ExtensionInfo,
) (CorsoItemExtension, error) {
	return &MockExtension{
		ctx:     ctx,
		innerRc: rc,
		info:    info,
		extInfo: extInfo,
	}, nil
}

// Extension client interface
type CorsoItemExtension interface {
	io.ReadCloser
}

type CorsoItemExtensionFactory func(
	context.Context,
	io.ReadCloser,
	details.ItemInfo,
	*details.ExtensionInfo,
) (CorsoItemExtension, error)

// AddItemExtensions wraps provided readcloser with extensions
// supplied via factory
func AddItemExtensions(
	ctx context.Context,
	rc io.ReadCloser,
	info details.ItemInfo,
	factories []CorsoItemExtensionFactory,
) (CorsoItemExtension, *details.ExtensionInfo, error) {
	// TODO: move to validate
	if rc == nil {
		return nil, nil, clues.New("nil inner readcloser")
	}

	if len(factories) == 0 {
		return rc, nil, clues.New("no extensions supplied")
	}

	ctx = clues.Add(ctx, "num_extensions", len(factories))

	extInfo := &details.ExtensionInfo{
		Data: make(map[string]any),
	}

	logger.Ctx(ctx).Info("adding extensions")

	for _, factory := range factories {
		extRc, err := factory(ctx, rc, info, extInfo)
		if err != nil {
			return nil, nil, clues.Wrap(err, "creating extension")
		}

		rc = extRc
	}

	logger.Ctx(ctx).Info("added extensions")

	// TODO: Add an outermost extension for logging & metrics
	return rc, extInfo, nil
}

// var _ ExtensionHandler = &extensionHandler{}

// type ExtensionHandler interface {
// 	io.ReadCloser
// }

// type extensionHandler struct {
// 	info           details.ItemInfo
// 	extData        *details.ExtensionInfo
// 	innerRc        io.ReadCloser
// 	itemExtensions []CorsoItemExtension
// }

// type ExtensionHandlerFactory func(
// 	info details.ItemInfo,
// 	extData *details.ExtensionInfo,
// 	rc io.ReadCloser,
// 	factory []CorsoItemExtensionFactory,
// ) (ExtensionHandler, error)

// func NewExtensionHandler(
// 	info details.ItemInfo,
// 	extData *details.ExtensionInfo,
// 	rc io.ReadCloser,
// 	factory []CorsoItemExtensionFactory,
// ) (ExtensionHandler, error) {
// 	itemExtensions := make([]CorsoItemExtension, len(factory))

// 	for i, f := range factory {
// 		if f == nil {
// 			return nil, clues.New("nil extension factory")
// 		}

// 		itemExtensions[i] = f()
// 	}

// 	return &extensionHandler{
// 		info:           info,
// 		extData:        extData,
// 		innerRc:        rc,
// 		itemExtensions: itemExtensions,
// 	}, nil
// }

// // TODO: more robust error handling as per io.Reader guidelines
// // and best practices followed by implementations of io.Reader
// func (ef *extensionHandler) Read(p []byte) (int, error) {
// 	n, err := ef.innerRc.Read(p)
// 	if err != nil {
// 		return n, err
// 	}

// 	// TODO: Why not just send ioReader instead of ioReadCloser?
// 	rc := io.NopCloser(bytes.NewReader(p[:n]))

// 	// Call extensions iteratively
// 	for _, ext := range ef.itemExtensions {
// 		rc, err = ext.WrapItem(ef.info, rc)
// 		if err != nil {
// 			return 0, err
// 		}
// 	}

// 	return n, err
// }

// func (ef *extensionHandler) Close() error {
// 	err := ef.innerRc.Close()
// 	if err != nil {
// 		return err
// 	}

// 	// TODO: handle errors from extensions
// 	for _, ext := range ef.itemExtensions {
// 		for k, v := range ext.OutputData() {
// 			// Last writer wins on collisions
// 			ef.extData.Data[k] = v
// 		}
// 	}

// 	return nil
// }
