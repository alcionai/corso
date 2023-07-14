package extensions

import (
	"context"
	"errors"
	"hash/crc32"
	"io"
	"sync/atomic"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
)

const (
	KNumBytes = "NumBytes"
	KCrc32    = "Crc32"
)

var _ io.ReadCloser = &MockExtension{}

type MockExtension struct {
	NumBytes    int64
	Crc32       uint32
	Info        details.ItemInfo
	ExtData     *details.ExtensionData
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
	if err != nil && !errors.Is(err, io.EOF) {
		logger.CtxErr(me.Ctx, err).Error("inner read error")
		return n, clues.Stack(err)
	}

	atomic.AddInt64(&me.NumBytes, int64(n))

	me.Crc32 = crc32.Update(me.Crc32, crc32.IEEETable, p[:n])

	if errors.Is(err, io.EOF) {
		me.ExtData.Data[KNumBytes] = me.NumBytes
		me.ExtData.Data[KCrc32] = me.Crc32
	}

	return n, err
}

func (me *MockExtension) Close() error {
	if me.FailOnClose {
		return clues.New("mock close error")
	}

	err := me.InnerRc.Close()
	if err != nil {
		return clues.Stack(err)
	}

	me.ExtData.Data[KNumBytes] = me.NumBytes
	me.ExtData.Data[KCrc32] = me.Crc32
	logger.Ctx(me.Ctx).Infow(
		"mock extension closed",
		KNumBytes, me.NumBytes, KCrc32, me.Crc32)

	return nil
}

type MockItemExtensionFactory struct {
	FailOnFactoryCreation bool
	FailOnRead            bool
	FailOnClose           bool
}

func (m *MockItemExtensionFactory) CreateItemExtension(
	ctx context.Context,
	rc io.ReadCloser,
	info details.ItemInfo,
	extData *details.ExtensionData,
) (io.ReadCloser, error) {
	if m.FailOnFactoryCreation {
		return nil, clues.New("factory error")
	}

	return &MockExtension{
		Ctx:         ctx,
		InnerRc:     rc,
		Info:        info,
		ExtData:     extData,
		FailOnRead:  m.FailOnRead,
		FailOnClose: m.FailOnClose,
	}, nil
}
