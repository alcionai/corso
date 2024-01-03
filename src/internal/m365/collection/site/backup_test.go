package site

import (
	"errors"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/data"
	siteMock "github.com/alcionai/corso/src/internal/m365/collection/site/mock"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type SharePointBackupUnitSuite struct {
	tester.Suite
	creds account.M365Config
}

func TestSharePointBackupUnitSuite(t *testing.T) {
	suite.Run(t, &SharePointBackupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SharePointBackupUnitSuite) SetupSuite() {
	a := tconfig.NewFakeM365Account(suite.T())
	m365, err := a.M365Config()
	require.NoError(suite.T(), err, clues.ToCore(err))
	suite.creds = m365
}

func (suite *SharePointBackupUnitSuite) TestCollectLists() {
	t := suite.T()

	var (
		statusUpdater = func(*support.ControllerOperationStatus) {}
		siteID        = tconfig.M365SiteID(t)
		sel           = selectors.NewSharePointBackup([]string{siteID})
	)

	table := []struct {
		name                 string
		mock                 siteMock.ListHandler
		expectErr            require.ErrorAssertionFunc
		expectColls          int
		expectNewColls       int
		expectMetadataColls  int
		canUsePreviousBackup bool
	}{
		{
			name:                 "one list",
			mock:                 siteMock.NewListHandler(siteMock.StubLists("one"), siteID, nil),
			expectErr:            require.NoError,
			expectColls:          2,
			expectNewColls:       1,
			expectMetadataColls:  1,
			canUsePreviousBackup: true,
		},
		{
			name:                 "many lists",
			mock:                 siteMock.NewListHandler(siteMock.StubLists("one", "two"), siteID, nil),
			expectErr:            require.NoError,
			expectColls:          3,
			expectNewColls:       2,
			expectMetadataColls:  1,
			canUsePreviousBackup: true,
		},
		{
			name:                 "with error",
			mock:                 siteMock.NewListHandler(siteMock.StubLists("one"), siteID, errors.New("some error")),
			expectErr:            require.Error,
			expectColls:          0,
			expectNewColls:       0,
			expectMetadataColls:  0,
			canUsePreviousBackup: false,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext(t)
			defer flush()

			ac, err := api.NewClient(
				suite.creds,
				control.DefaultOptions(),
				count.New())
			require.NoError(t, err, clues.ToCore(err))

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: version.NoBackup,
				Options:           control.DefaultOptions(),
				ProtectedResource: mock.NewProvider(siteID, siteID),
			}

			cs, canUsePreviousBackup, err := CollectLists(
				ctx,
				test.mock,
				bpc,
				ac,
				suite.creds.AzureTenantID,
				sel.Lists(selectors.Any())[0],
				statusUpdater,
				count.New(),
				fault.New(false))

			test.expectErr(t, err, clues.ToCore(err))
			assert.Len(t, cs, test.expectColls, "number of collections")
			assert.Equal(t, test.canUsePreviousBackup, canUsePreviousBackup)

			newStates, metadatas := 0, 0
			for _, c := range cs {
				if c.FullPath() != nil && c.FullPath().Service() == path.SharePointMetadataService {
					metadatas++
					continue
				}

				if c.State() == data.NewState {
					newStates++
				}
			}

			assert.Equal(t, test.expectNewColls, newStates, "new collections")
			assert.Equal(t, test.expectMetadataColls, metadatas, "metadata collections")
		})
	}
}

func (suite *SharePointBackupUnitSuite) TestPopulateListsCollections_incremental() {
	t := suite.T()

	var (
		statusUpdater = func(*support.ControllerOperationStatus) {}
		siteID        = tconfig.M365SiteID(t)
		sel           = selectors.NewSharePointBackup([]string{siteID})
	)

	ac, err := api.NewClient(
		suite.creds,
		control.DefaultOptions(),
		count.New())
	require.NoError(t, err, clues.ToCore(err))

	listPathOne, err := path.Build(
		suite.creds.AzureTenantID,
		siteID,
		path.SharePointService,
		path.ListsCategory,
		false,
		"one")
	require.NoError(suite.T(), err, clues.ToCore(err))

	listPathTwo, err := path.Build(
		suite.creds.AzureTenantID,
		siteID,
		path.SharePointService,
		path.ListsCategory,
		false,
		"two")
	require.NoError(suite.T(), err, clues.ToCore(err))

	listPathThree, err := path.Build(
		suite.creds.AzureTenantID,
		siteID,
		path.SharePointService,
		path.ListsCategory,
		false,
		"three")
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		name                string
		lists               []models.Listable
		deltaPaths          metadata.DeltaPaths
		expectErr           require.ErrorAssertionFunc
		expectColls         int
		expectNewColls      int
		expectNotMovedColls int
		expectMetadataColls int
		expectTombstoneCols int
	}{
		{
			name:  "one list",
			lists: siteMock.StubLists("one"),
			deltaPaths: metadata.DeltaPaths{
				"one": {
					Path: listPathOne.String(),
				},
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNotMovedColls: 1,
			expectNewColls:      0,
			expectMetadataColls: 1,
			expectTombstoneCols: 0,
		},
		{
			name:  "one lists, one deleted",
			lists: siteMock.StubLists("two"),
			deltaPaths: metadata.DeltaPaths{
				"one": {
					Path: listPathOne.String(),
				},
			},
			expectErr:           require.NoError,
			expectColls:         3,
			expectNewColls:      1,
			expectMetadataColls: 1,
			expectTombstoneCols: 1,
		},
		{
			name:  "two lists, one deleted",
			lists: siteMock.StubLists("one", "two"),
			deltaPaths: metadata.DeltaPaths{
				"one": {
					Path: listPathOne.String(),
				},
				"three": {
					Path: listPathThree.String(),
				},
			},
			expectErr:           require.NoError,
			expectColls:         4,
			expectNotMovedColls: 1,
			expectNewColls:      1,
			expectMetadataColls: 1,
			expectTombstoneCols: 1,
		},
		{
			name:                "no previous paths",
			lists:               siteMock.StubLists("one", "two"),
			deltaPaths:          metadata.DeltaPaths{},
			expectErr:           require.NoError,
			expectColls:         3,
			expectNotMovedColls: 0,
			expectNewColls:      2,
			expectMetadataColls: 1,
			expectTombstoneCols: 0,
		},
		{
			name:  "two lists, unchanges",
			lists: siteMock.StubLists("one", "two"),
			deltaPaths: metadata.DeltaPaths{
				"one": {
					Path: listPathOne.String(),
				},
				"two": {
					Path: listPathTwo.String(),
				},
			},
			expectErr:           require.NoError,
			expectColls:         3,
			expectNotMovedColls: 2,
			expectNewColls:      0,
			expectMetadataColls: 1,
			expectTombstoneCols: 0,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext(t)
			defer flush()

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: version.NoBackup,
				Options:           control.DefaultOptions(),
				ProtectedResource: mock.NewProvider(siteID, siteID),
			}

			cs, err := populateListsCollections(
				ctx,
				siteMock.NewListHandler(test.lists, siteID, nil),
				bpc,
				ac,
				suite.creds.AzureTenantID,
				sel.Lists(selectors.Any())[0],
				statusUpdater,
				test.lists,
				test.deltaPaths,
				count.New(),
				fault.New(false))

			test.expectErr(t, err, clues.ToCore(err))
			assert.Len(t, cs, test.expectColls, "number of collections")

			newStates, notMovedStates, metadatas, tombstoned := 0, 0, 0, 0
			for _, c := range cs {
				if c.FullPath() != nil && c.FullPath().Service() == path.SharePointMetadataService {
					metadatas++
					continue
				}

				if c.State() == data.DeletedState {
					tombstoned++
				}

				if c.State() == data.NewState {
					newStates++
				}

				if c.State() == data.NotMovedState {
					notMovedStates++
				}
			}

			assert.Equal(t, test.expectNewColls, newStates, "new collections")
			assert.Equal(t, test.expectNotMovedColls, notMovedStates, "not moved collections")
			assert.Equal(t, test.expectMetadataColls, metadatas, "metadata collections")
			assert.Equal(t, test.expectTombstoneCols, tombstoned, "tombstone collections")
		})
	}
}

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

	col, _, err := CollectLists(
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
