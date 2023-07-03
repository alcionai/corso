package extensions

import (
	"context"
	"io"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
)

type CreateItemExtensioner interface {
	CreateItemExtension(
		context.Context,
		io.ReadCloser,
		details.ItemInfo,
		*details.ExtensionInfo,
	) (io.ReadCloser, error)
}

// AddItemExtensions wraps provided readcloser with extensions
// supplied via factory, with the first extension in slice being
// the innermost one.
func AddItemExtensions(
	ctx context.Context,
	rc io.ReadCloser,
	info details.ItemInfo,
	factories []CreateItemExtensioner,
) (io.ReadCloser, *details.ExtensionInfo, error) {
	if rc == nil {
		return nil, nil, clues.New("nil readcloser")
	}

	if len(factories) == 0 {
		return nil, nil, clues.New("no extensions supplied")
	}

	ctx = clues.Add(ctx, "num_extensions", len(factories))

	extInfo := &details.ExtensionInfo{
		Data: make(map[string]any),
	}

	for _, factory := range factories {
		if factory == nil {
			return nil, nil, clues.New("nil extension factory")
		}

		extRc, err := factory.CreateItemExtension(ctx, rc, info, extInfo)
		if err != nil {
			return nil, nil, clues.Wrap(err, "create item extension")
		}

		rc = extRc
	}

	logger.Ctx(ctx).Debug("added item extensions")

	return rc, extInfo, nil
}
