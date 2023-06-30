package extensions

import (
	"context"
	"io"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
)

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

// Thin wrapper for runtime logging & metrics around extensions
type loggerExtension struct {
	info    details.ItemInfo
	innerRc io.ReadCloser
	ctx     context.Context
	extInfo *details.ExtensionInfo
}

func NewLoggerExtension(
	ctx context.Context,
	rc io.ReadCloser,
	info details.ItemInfo,
	extInfo *details.ExtensionInfo,
) (CorsoItemExtension, error) {
	return &loggerExtension{
		ctx:     ctx,
		innerRc: rc,
		info:    info,
		extInfo: extInfo,
	}, nil
}

func (l *loggerExtension) Read(p []byte) (int, error) {
	n, err := l.innerRc.Read(p)
	if err != nil && err != io.EOF {
		logger.CtxErr(l.ctx, err).Error("inner read")
		return n, err
	}

	if err == io.EOF {
		logger.Ctx(l.ctx).Debug("corso extensions: EOF")
	}

	return n, err
}

func (l *loggerExtension) Close() error {
	err := l.innerRc.Close()
	if err != nil {
		logger.CtxErr(l.ctx, err).Error("inner close")
		return err
	}

	logger.Ctx(l.ctx).Info("corso extensions: closed")

	return nil
}

type AddItemExtensioner interface {
	AddItemExtensions(
		context.Context,
		io.ReadCloser,
		details.ItemInfo,
		[]CorsoItemExtensionFactory,
	) (io.ReadCloser, *details.ExtensionInfo, error)
}

var _ AddItemExtensioner = &ItemExtensionHandler{}

type ItemExtensionHandler struct{}

// AddItemExtensions wraps provided readcloser with extensions
// supplied via factory, with the first extension in slice being
// the innermost one.
func (eh *ItemExtensionHandler) AddItemExtensions(
	ctx context.Context,
	rc io.ReadCloser,
	info details.ItemInfo,
	factories []CorsoItemExtensionFactory,
) (io.ReadCloser, *details.ExtensionInfo, error) {
	if rc == nil {
		return nil, nil, clues.New("nil inner readcloser")
	}

	if len(factories) == 0 {
		return nil, nil, clues.New("no extensions supplied")
	}

	factories = append(factories, NewLoggerExtension)
	ctx = clues.Add(ctx, "num_extensions", len(factories))

	extInfo := &details.ExtensionInfo{
		Data: make(map[string]any),
	}

	for _, factory := range factories {
		if factory == nil {
			return nil, nil, clues.New("nil extension factory")
		}

		extRc, err := factory(ctx, rc, info, extInfo)
		if err != nil {
			return nil, nil, clues.Wrap(err, "calling extension factory")
		}

		rc = extRc
	}

	logger.Ctx(ctx).Info("added extensions")

	return rc, extInfo, nil
}
