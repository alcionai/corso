package groups

import (
	"context"
	"sync"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/groups/testdata"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// mocks
// ---------------------------------------------------------------------------

var _ backupHandler = &mockBackupHandler{}

type mockBackupHandler struct {
	protectedResource string
	channels          []models.Channelable
	channelsErr       error
	messages          []models.ChatMessageable
	messagesErr       error
	include           bool
}

func (bh mockBackupHandler) getChannels(context.Context) ([]models.Channelable, error) {
	return bh.channels, bh.channelsErr
}
func (bh mockBackupHandler) getChannelMessagesDelta(
	_ context.Context,
	_, _ string,
) ([]models.ChatMessageable, api.DeltaUpdate, error) {
	return bh.messages, api.DeltaUpdate{}, bh.messagesErr
}

func (bh mockBackupHandler) includeContainer(
	context.Context,
	graph.QueryParams,
	models.Channelable,
	selectors.GroupsScope,
) bool {
	return bh.include
}

func (bh mockBackupHandler) canonicalPath(
	folders *path.Builder,
	tenantID string,
) (path.Path, error) {
	return folders.
		ToDataLayerPath(
			tenantID,
			bh.protectedResource,
			path.GroupsService,
			path.ChannelMessagesCategory,
			false)
}

// ---------------------------------------------------------------------------
// Unit Suite
// ---------------------------------------------------------------------------

type UnitSuite struct {
	tester.Suite
	creds account.M365Config
}

func TestServiceIteratorsUnitSuite(t *testing.T) {
	suite.Run(t, &UnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *UnitSuite) SetupSuite() {
	a := tconfig.NewFakeM365Account(suite.T())
	m365, err := a.M365Config()
	require.NoError(suite.T(), err, clues.ToCore(err))
	suite.creds = m365
}

func (suite *UnitSuite) TestPopulateCollections() {
	var (
		qp = graph.QueryParams{
			Category:          path.ChannelMessagesCategory, // doesn't matter which one we use.
			ProtectedResource: inMock.NewProvider("group_id", "user_name"),
			TenantID:          suite.creds.AzureTenantID,
		}
		statusUpdater = func(*support.ControllerOperationStatus) {}
		allScope      = selectors.NewGroupsBackup(nil).Channels(selectors.Any())[0]
	)

	table := []struct {
		name                  string
		mock                  mockBackupHandler
		scope                 selectors.GroupsScope
		failFast              control.FailurePolicy
		expectErr             assert.ErrorAssertionFunc
		expectNewColls        int
		expectMetadataColls   int
		expectDoNotMergeColls int
	}{
		{
			name: "happy path, one container",
			mock: mockBackupHandler{
				channels: testdata.StubChannels("one"),
			},
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      1,
			expectMetadataColls: 0,
		},
		{
			name: "happy path, many containers",
			mock: mockBackupHandler{
				channels: testdata.StubChannels("one", "two"),
			},
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      2,
			expectMetadataColls: 0,
		},
		{
			name: "no containers pass scope",
			mock: mockBackupHandler{
				channels: testdata.StubChannels("one"),
			},
			scope:               selectors.NewGroupsBackup(nil).Channels(selectors.None())[0],
			expectErr:           assert.NoError,
			expectNewColls:      0,
			expectMetadataColls: 0,
		},
		{
			name: "err: deleted in flight",
			mock: mockBackupHandler{
				channelsErr: graph.ErrDeletedInFlight,
			},
			scope:                 allScope,
			expectErr:             assert.NoError,
			expectNewColls:        1,
			expectMetadataColls:   0,
			expectDoNotMergeColls: 1,
		},
		{
			name: "err: other error",
			mock: mockBackupHandler{
				channelsErr: assert.AnError,
			},
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      0,
			expectMetadataColls: 0,
		},
	}
	for _, test := range table {
		for _, canMakeDeltaQueries := range []bool{true, false} {
			name := test.name

			if canMakeDeltaQueries {
				name += "-delta"
			} else {
				name += "-non-delta"
			}

			suite.Run(name, func() {
				t := suite.T()

				ctx, flush := tester.NewContext(t)
				defer flush()

				ctrlOpts := control.Options{FailureHandling: test.failFast}
				ctrlOpts.ToggleFeatures.DisableDelta = !canMakeDeltaQueries

				collections, err := populateCollections(
					ctx,
					qp,
					test.mock,
					statusUpdater,
					nil,
					test.scope,
					ctrlOpts,
					fault.New(test.failFast == control.FailFast))
				test.expectErr(t, err, clues.ToCore(err))

				// collection assertions

				deleteds, news, metadatas, doNotMerges := 0, 0, 0, 0
				for _, c := range collections {
					if c.FullPath().Service() == path.ExchangeMetadataService {
						metadatas++
						continue
					}

					if c.State() == data.DeletedState {
						deleteds++
					}

					if c.State() == data.NewState {
						news++
					}

					if c.DoNotMergeItems() {
						doNotMerges++
					}
				}

				assert.Zero(t, deleteds, "deleted collections")
				assert.Equal(t, test.expectNewColls, news, "new collections")
				assert.Equal(t, test.expectMetadataColls, metadatas, "metadata collections")
				assert.Equal(t, test.expectDoNotMergeColls, doNotMerges, "doNotMerge collections")
			})
		}
	}
}

// ---------------------------------------------------------------------------
// Integration tests
// ---------------------------------------------------------------------------

func newStatusUpdater(t *testing.T, wg *sync.WaitGroup) func(status *support.ControllerOperationStatus) {
	updater := func(status *support.ControllerOperationStatus) {
		defer wg.Done()
	}

	return updater
}

type BackupIntgSuite struct {
	tester.Suite
	resource string
	tenantID string
	ac       api.Client
}

func TestBackupIntgSuite(t *testing.T) {
	suite.Run(t, &BackupIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *BackupIntgSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	suite.resource = tconfig.M365TeamID(t)

	acct := tconfig.NewM365Account(t)
	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.ac, err = api.NewClient(creds, control.DefaultOptions())
	require.NoError(t, err, clues.ToCore(err))

	suite.tenantID = creds.AzureTenantID
}

func (suite *BackupIntgSuite) TestCreateCollections() {
	var (
		protectedResource = tconfig.M365GroupID(suite.T())
		resources         = []string{protectedResource}
		handler           = NewChannelBackupHandler(protectedResource, suite.ac.Channels())
	)

	tests := []struct {
		name                string
		scope               selectors.GroupsScope
		channelNames        map[string]struct{}
		canMakeDeltaQueries bool
	}{
		{
			name:  "channel messages non-delta",
			scope: selTD.GroupsBackupChannelScope(selectors.NewGroupsBackup(resources))[0],
			channelNames: map[string]struct{}{
				selTD.TestChannelName: {},
			},
			canMakeDeltaQueries: false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrlOpts := control.DefaultOptions()
			ctrlOpts.ToggleFeatures.DisableDelta = !test.canMakeDeltaQueries

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: version.NoBackup,
				Options:           ctrlOpts,
				ProtectedResource: inMock.NewProvider(protectedResource, protectedResource),
			}

			collections, err := CreateCollections(
				ctx,
				bpc,
				handler,
				suite.tenantID,
				test.scope,
				func(status *support.ControllerOperationStatus) {},
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			for _, c := range collections {
				if c.FullPath().Service() == path.GroupsMetadataService {
					continue
				}

				require.NotEmpty(t, c.FullPath().Folder(false))

				// TODO(ashmrtn): Remove when LocationPath is made part of BackupCollection
				// interface.
				if !assert.Implements(t, (*data.LocationPather)(nil), c) {
					continue
				}

				loc := c.(data.LocationPather).LocationPath().String()

				require.NotEmpty(t, loc)

				delete(test.channelNames, loc)
			}

			assert.Empty(t, test.channelNames)
		})
	}
}
