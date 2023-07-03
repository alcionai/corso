package extensions

import (
	"context"
	"hash/crc32"
	"io"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
)

var _ io.ReadCloser = &MockExtension{}

type MockExtension struct {
	numBytes    int
	crc32       uint32
	info        details.ItemInfo
	extInfo     *details.ExtensionInfo
	innerRc     io.ReadCloser
	ctx         context.Context
	failOnRead  bool
	failOnClose bool
}

func (me *MockExtension) Read(p []byte) (int, error) {
	if me.failOnRead {
		return 0, clues.New("mock read error")
	}

	n, err := me.innerRc.Read(p)
	if err != nil && err != io.EOF {
		logger.CtxErr(me.ctx, err).Error("inner read error")
		return n, err
	}

	me.numBytes += n
	me.crc32 = crc32.Update(me.crc32, crc32.IEEETable, p[:n])

	if err == io.EOF {
		logger.Ctx(me.ctx).Debug("mock extension reached EOF")
		me.extInfo.Data["numBytes"] = me.numBytes
		me.extInfo.Data["crc32"] = me.crc32
	}

	return n, err
}

func (me *MockExtension) Close() error {
	if me.failOnClose {
		return clues.New("mock close error")
	}

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

type MockItemExtensionFactory struct {
	shouldReturnError bool
}

func (m *MockItemExtensionFactory) CreateItemExtension(
	ctx context.Context,
	rc io.ReadCloser,
	info details.ItemInfo,
	extInfo *details.ExtensionInfo,
) (io.ReadCloser, error) {
	if m.shouldReturnError {
		return nil, clues.New("factory error")
	}

	return &MockExtension{
		ctx:     ctx,
		innerRc: rc,
		info:    info,
		extInfo: extInfo,
	}, nil
}
