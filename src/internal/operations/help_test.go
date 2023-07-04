package operations

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/m365"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// A QoL builder for live instances that updates
// the selector's owner id and name in the process
// to help avoid gotchas.
func ControllerWithSelector(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	acct account.Account,
	cr resource.Category,
	sel selectors.Selector,
	ins idname.Cacher,
	onFail func(),
) (*m365.Controller, selectors.Selector) {
	ctrl, err := m365.NewController(ctx, acct, cr, sel.PathService(), control.Options{})
	if !assert.NoError(t, err, clues.ToCore(err)) {
		if onFail != nil {
			onFail()
		}

		t.FailNow()
	}

	id, name, err := ctrl.PopulateOwnerIDAndNamesFrom(ctx, sel.DiscreteOwner, ins)
	if !assert.NoError(t, err, clues.ToCore(err)) {
		if onFail != nil {
			onFail()
		}

		t.FailNow()
	}

	sel = sel.SetDiscreteOwnerIDName(id, name)

	return ctrl, sel
}
