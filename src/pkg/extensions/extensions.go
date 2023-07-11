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
		*details.ExtensionData,
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
) (io.ReadCloser, *details.ExtensionData, error) {
	if rc == nil {
		return nil, nil, clues.New("nil readcloser")
	}

	// If no extensions were supplied, return the original readcloser
	if len(factories) == 0 {
		return rc, &details.ExtensionData{}, nil
	}

	ctx = clues.Add(ctx, "num_extensions", len(factories))

	extData := &details.ExtensionData{
		Data: make(map[string]any),
	}

	for _, factory := range factories {
		if factory == nil {
			return nil, nil, clues.New("nil extension factory")
		}

		extRc, err := factory.CreateItemExtension(ctx, rc, info, extData)
		if err != nil {
			return nil, nil, clues.Wrap(err, "create item extension")
		}

		rc = extRc
	}

	logger.Ctx(ctx).Debug("added item extensions")

	return rc, extData, nil
}
