package operations

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/config"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type MaintenanceOpIntegrationSuite struct {
	tester.Suite
}

func TestMaintenanceOpIntegrationSuite(t *testing.T) {
	suite.Run(t, &MaintenanceOpIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, config.M365AcctCredEnvs}),
	})
}

func (suite *MaintenanceOpIntegrationSuite) TestRepoMaintenance() {
	var (
		t = suite.T()
		// need to initialize the repository before we can test connecting to it.
		st = storeTD.NewPrefixedS3Storage(t)
		k  = kopia.NewConn(st)
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	err := k.Initialize(ctx, repository.Options{})
	require.NoError(t, err, clues.ToCore(err))

	kw, err := kopia.NewWrapper(k)
	// kopiaRef comes with a count of 1 and Wrapper bumps it again so safe
	// to close here.
	k.Close(ctx)

	require.NoError(t, err, clues.ToCore(err))

	defer kw.Close(ctx)

	mo, err := NewMaintenanceOperation(
		ctx,
		control.Defaults(),
		kw,
		repository.Maintenance{
			Type: repository.MetadataMaintenance,
		},
		evmock.NewBus())
	require.NoError(t, err, clues.ToCore(err))

	err = mo.Run(ctx)
	assert.NoError(t, err, clues.ToCore(err))
}
