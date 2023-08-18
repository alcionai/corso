package site

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type SharePointPagesSuite struct {
	tester.Suite
}

func TestSharePointPagesSuite(t *testing.T) {
	suite.Run(t, &SharePointPagesSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *SharePointPagesSuite) SetupSuite() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, false, 4)
}

func (suite *SharePointPagesSuite) TestCollectPages() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		siteID = tconfig.M365SiteID(t)
		a      = tconfig.NewM365Account(t)
	)

	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	ac, err := api.NewClient(creds, control.DefaultOptions())
	require.NoError(t, err, clues.ToCore(err))

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: mock.NewProvider(siteID, siteID),
	}

	sel := selectors.NewSharePointBackup([]string{siteID})

	col, err := CollectPages(
		ctx,
		bpc,
		creds,
		ac,
		sel.Lists(selectors.Any())[0],
		(&MockGraphService{}).UpdateStatus,
		fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, col)
}
