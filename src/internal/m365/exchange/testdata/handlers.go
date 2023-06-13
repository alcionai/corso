package testdata

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/m365/exchange"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func PopulateContainerCache(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	ac api.Client,
	category path.CategoryType,
	resourceOwnerID string,
	errs *fault.Bus,
) graph.ContainerResolver {
	handler, ok := exchange.BackupHandlers(ac)[category]
	require.Truef(t, ok, "container resolver registered for category %s", category)

	root, cc := handler.NewContainerCache(resourceOwnerID)

	err := cc.Populate(ctx, errs, root)
	require.NoError(t, err, clues.ToCore(err))

	return cc
}
