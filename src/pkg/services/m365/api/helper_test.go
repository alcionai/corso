package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type clientTesterSetup struct {
	ac     api.Client
	userID string
}

func newClientTesterSetup(t *testing.T) clientTesterSetup {
	cts := clientTesterSetup{}

	ctx, flush := tester.NewContext(t)
	defer flush()

	a := tester.NewM365Account(t)
	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	cts.ac, err = api.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	cts.userID = tester.GetM365UserID(ctx)

	return cts
}
