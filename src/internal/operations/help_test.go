package operations

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"

	"github.com/alcionai/canario/src/internal/common/idname"
	"github.com/alcionai/canario/src/internal/m365"
	"github.com/alcionai/canario/src/pkg/account"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/selectors"
)

// A QoL builder for live instances that updates
// the selector's owner id and name in the process
// to help avoid gotchas.
func ControllerWithSelector(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	acct account.Account,
	sel selectors.Selector,
	ins idname.Cacher,
	onFail func(),
) (*m365.Controller, selectors.Selector) {
	ctrl, err := m365.NewController(
		ctx,
		acct,
		sel.PathService(),
		control.DefaultOptions(),
		count.New())
	if !assert.NoError(t, err, clues.ToCore(err)) {
		if onFail != nil {
			onFail()
		}

		t.FailNow()
	}

	resource, err := ctrl.PopulateProtectedResourceIDAndName(ctx, sel.DiscreteOwner, ins)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		if onFail != nil {
			onFail()
		}

		t.FailNow()
	}

	sel = sel.SetDiscreteOwnerIDName(resource.ID(), resource.Name())

	return ctrl, sel
}
