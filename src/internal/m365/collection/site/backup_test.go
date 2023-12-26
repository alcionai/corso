package site

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type SharePointSuite struct {
	tester.Suite
}

func TestSharePointSuite(t *testing.T) {
	suite.Run(t, &SharePointSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *SharePointSuite) SetupSuite() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, false, 4)
}

func (suite *SharePointSuite) TestCollectPages() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		siteID  = tconfig.M365SiteID(t)
		a       = tconfig.NewM365Account(t)
		counter = count.New()
	)

	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	ac, err := api.NewClient(
		creds,
		control.DefaultOptions(),
		counter)
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
		counter,
		fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, col)
}

func (suite *SharePointSuite) TestCollectLists() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		siteID  = tconfig.M365SiteID(t)
		a       = tconfig.NewM365Account(t)
		counter = count.New()
	)

	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	ac, err := api.NewClient(
		creds,
		control.DefaultOptions(),
		counter)
	require.NoError(t, err, clues.ToCore(err))

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: mock.NewProvider(siteID, siteID),
	}

	sel := selectors.NewSharePointBackup([]string{siteID})

	bh := NewListsBackupHandler(bpc.ProtectedResource.ID(), ac.Lists())

	col, err := CollectLists(
		ctx,
		bh,
		bpc,
		ac,
		creds.AzureTenantID,
		sel.Lists(selectors.Any())[0],
		(&MockGraphService{}).UpdateStatus,
		count.New(),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	metadataFound := false

	for _, c := range col {
		if c.FullPath().Service() == path.SharePointMetadataService {
			metadataFound = true
			break
		}
	}

	assert.Less(t, 0, len(col))
	assert.True(t, metadataFound)
}
