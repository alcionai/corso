package exchange

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type intgTesterSetup struct {
	ac     api.Client
	creds  account.M365Config
	userID string
}

func newIntegrationTesterSetup(t *testing.T) intgTesterSetup {
	its := intgTesterSetup{}

	ctx, flush := tester.NewContext(t)
	defer flush()

	a := tconfig.NewM365Account(t)
	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	its.creds = creds

	its.ac, err = api.NewClient(creds, control.DefaultOptions())
	require.NoError(t, err, clues.ToCore(err))

	its.userID = tconfig.GetM365UserID(ctx)

	return its
}
