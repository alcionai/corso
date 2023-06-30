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
		return nil, nil, clues.New("no extensions supplied")
	}

	ctx = clues.Add(ctx, "num_extensions", len(factories))

	extInfo := &details.ExtensionInfo{
		Data: make(map[string]any),
	}

	logger.Ctx(ctx).Info("adding extensions")

	for _, factory := range factories {
		if factory == nil {
			return nil, nil, clues.New("nil extension factory")
		}

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
