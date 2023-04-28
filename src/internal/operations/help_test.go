package operations

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/connector"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// A QoL builder for live GC instances that updates
// the selector's owner id and name in the process
// to help avoid gotchas.
func GCWithSelector(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	acct account.Account,
	cr connector.Resource,
	sel selectors.Selector,
	ins idname.Cacher,
	onFail func(),
) (*connector.GraphConnector, selectors.Selector) {
	gc, err := connector.NewGraphConnector(ctx, acct, cr)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		if onFail != nil {
			onFail()
		}

		t.FailNow()
	}

	id, name, err := gc.PopulateOwnerIDAndNamesFrom(ctx, sel.DiscreteOwner, ins)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		if onFail != nil {
			onFail()
		}

		t.FailNow()
	}

	sel = sel.SetDiscreteOwnerIDName(id, name)

	return gc, sel
}
