package testdata

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/canario/src/internal/m365/collection/exchange"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph"
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
