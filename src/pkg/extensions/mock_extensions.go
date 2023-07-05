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
	NumBytes    int64
	Crc32       uint32
	Info        details.ItemInfo
	ExtInfo     *details.ExtensionInfo
	InnerRc     io.ReadCloser
	Ctx         context.Context
	FailOnRead  bool
	FailOnClose bool
}

func (me *MockExtension) Read(p []byte) (int, error) {
	if me.FailOnRead {
		return 0, clues.New("mock read error")
	}

	n, err := me.InnerRc.Read(p)
	if err != nil && err != io.EOF {
		logger.CtxErr(me.Ctx, err).Error("inner read error")
		return n, err
	}

	me.NumBytes += int64(n)
	me.Crc32 = crc32.Update(me.Crc32, crc32.IEEETable, p[:n])

	if err == io.EOF {
		me.ExtInfo.Data["NumBytes"] = me.NumBytes
		me.ExtInfo.Data["Crc32"] = me.Crc32
	}

	return n, err
}

func (me *MockExtension) Close() error {
	if me.FailOnClose {
		return clues.New("mock close error")
	}

	err := me.InnerRc.Close()
	if err != nil {
		return err
	}

	me.ExtInfo.Data["NumBytes"] = me.NumBytes
	me.ExtInfo.Data["Crc32"] = me.Crc32
	logger.Ctx(me.Ctx).Infow(
		"mock extension closed",
		"NumBytes", me.NumBytes, "Crc32", me.Crc32)

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
		Ctx:     ctx,
		InnerRc: rc,
		Info:    info,
		ExtInfo: extInfo,
	}, nil
}
