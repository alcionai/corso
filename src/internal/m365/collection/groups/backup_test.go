package groups

import (
	"context"
	"fmt"
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
	"github.com/alcionai/corso/src/pkg/backup/details"
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
	channels      []models.Channelable
	channelsErr   error
	messageIDs    map[string]struct{}
	deletedMsgIDs map[string]struct{}
	messagesErr   error
	messages      map[string]models.ChatMessageable
	info          map[string]*details.GroupsInfo
	getMessageErr map[string]error
	doNotInclude  bool
}

func (bh mockBackupHandler) getChannels(context.Context) ([]models.Channelable, error) {
	return bh.channels, bh.channelsErr
}

func (bh mockBackupHandler) getChannelMessageIDsDelta(
	_ context.Context,
	_, _ string,
) (map[string]struct{}, map[string]struct{}, api.DeltaUpdate, error) {
	return bh.messageIDs, bh.deletedMsgIDs, api.DeltaUpdate{}, bh.messagesErr
}

func (bh mockBackupHandler) includeContainer(
	context.Context,
	graph.QueryParams,
	models.Channelable,
	selectors.GroupsScope,
) bool {
	return !bh.doNotInclude
}

func (bh mockBackupHandler) canonicalPath(
	folders *path.Builder,
	tenantID string,
) (path.Path, error) {
	return folders.
		ToDataLayerPath(
			tenantID,
			"protectedResource",
			path.GroupsService,
			path.ChannelMessagesCategory,
			false)
}

func (bh mockBackupHandler) getChannelMessage(
	_ context.Context,
	_, _, itemID string,
) (models.ChatMessageable, *details.GroupsInfo, error) {
	return bh.messages[itemID], bh.info[itemID], bh.getMessageErr[itemID]
}

// ---------------------------------------------------------------------------
// Unit Suite
// ---------------------------------------------------------------------------

type BackupUnitSuite struct {
	tester.Suite
	creds account.M365Config
}

func TestServiceIteratorsUnitSuite(t *testing.T) {
	suite.Run(t, &BackupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *BackupUnitSuite) SetupSuite() {
	a := tconfig.NewFakeM365Account(suite.T())
	m365, err := a.M365Config()
	require.NoError(suite.T(), err, clues.ToCore(err))
	suite.creds = m365
}

func (suite *BackupUnitSuite) TestPopulateCollections() {
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
		expectErr             require.ErrorAssertionFunc
		expectColls           int
		expectNewColls        int
		expectMetadataColls   int
		expectDoNotMergeColls int
	}{
		{
			name: "happy path, one container",
			mock: mockBackupHandler{
				channels:   testdata.StubChannels("one"),
				messageIDs: map[string]struct{}{"msg-one": {}},
			},
			scope:                 allScope,
			expectErr:             require.NoError,
			expectColls:           1,
			expectNewColls:        1,
			expectMetadataColls:   0,
			expectDoNotMergeColls: 1,
		},
		{
			name: "happy path, one container, only deleted messages",
			mock: mockBackupHandler{
				channels:      testdata.StubChannels("one"),
				deletedMsgIDs: map[string]struct{}{"msg-one": {}},
			},
			scope:                 allScope,
			expectErr:             require.NoError,
			expectColls:           1,
			expectNewColls:        1,
			expectMetadataColls:   0,
			expectDoNotMergeColls: 1,
		},
		{
			name: "happy path, many containers",
			mock: mockBackupHandler{
				channels:   testdata.StubChannels("one", "two"),
				messageIDs: map[string]struct{}{"msg-one": {}},
			},
			scope:                 allScope,
			expectErr:             require.NoError,
			expectColls:           2,
			expectNewColls:        2,
			expectMetadataColls:   0,
			expectDoNotMergeColls: 2,
		},
		{
			name: "no containers pass scope",
			mock: mockBackupHandler{
				channels:     testdata.StubChannels("one"),
				doNotInclude: true,
			},
			scope:                 selectors.NewGroupsBackup(nil).Channels(selectors.None())[0],
			expectErr:             require.NoError,
			expectColls:           0,
			expectNewColls:        0,
			expectMetadataColls:   0,
			expectDoNotMergeColls: 0,
		},
		{
			name:                  "no channels",
			mock:                  mockBackupHandler{},
			scope:                 allScope,
			expectErr:             require.NoError,
			expectColls:           0,
			expectNewColls:        0,
			expectMetadataColls:   0,
			expectDoNotMergeColls: 0,
		},
		{
			name: "no channel messages",
			mock: mockBackupHandler{
				channels: testdata.StubChannels("one"),
			},
			scope:                 allScope,
			expectErr:             require.NoError,
			expectColls:           1,
			expectNewColls:        1,
			expectMetadataColls:   0,
			expectDoNotMergeColls: 1,
		},
		{
			name: "err: deleted in flight",
			mock: mockBackupHandler{
				channels:    testdata.StubChannels("one"),
				messagesErr: graph.ErrDeletedInFlight,
			},
			scope:                 allScope,
			expectErr:             require.Error,
			expectColls:           0,
			expectNewColls:        0,
			expectMetadataColls:   0,
			expectDoNotMergeColls: 0,
		},
		{
			name: "err: other error",
			mock: mockBackupHandler{
				channels:    testdata.StubChannels("one"),
				messagesErr: assert.AnError,
			},
			scope:                 allScope,
			expectErr:             require.Error,
			expectColls:           0,
			expectNewColls:        0,
			expectMetadataColls:   0,
			expectDoNotMergeColls: 0,
		},
	}
	for _, test := range table {
		// for _, canMakeDeltaQueries := range []bool{true, false} {
		name := test.name

		// 	if canMakeDeltaQueries {
		// 		name += "-delta"
		// 	} else {
		// 		name += "-non-delta"
		// 	}

		suite.Run(name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrlOpts := control.Options{FailureHandling: test.failFast}
			// ctrlOpts.ToggleFeatures.DisableDelta = !canMakeDeltaQueries

			collections, err := populateCollections(
				ctx,
				qp,
				test.mock,
				statusUpdater,
				test.mock.channels,
				test.scope,
				ctrlOpts,
				fault.New(true))
			test.expectErr(t, err, clues.ToCore(err))
			assert.Len(t, collections, test.expectColls, "number of collections")

			// collection assertions

			deleteds, news, metadatas, doNotMerges := 0, 0, 0, 0
			for _, c := range collections {
				if c.FullPath().Service() == path.GroupsMetadataService {
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

// }

// ---------------------------------------------------------------------------
// Integration tests
// ---------------------------------------------------------------------------

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

			sel := selectors.NewGroupsBackup([]string{protectedResource})
			sel.Include(selTD.GroupsBackupChannelScope(sel))

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: version.NoBackup,
				Options:           ctrlOpts,
				ProtectedResource: inMock.NewProvider(protectedResource, protectedResource),
				Selector:          sel.Selector,
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
			require.NotEmpty(t, collections, "must have at least one collection")

			for _, c := range collections {
				if c.FullPath().Service() == path.GroupsMetadataService {
					continue
				}

				require.NotEmpty(t, c.FullPath().Folder(false))

				fmt.Printf("\n-----\nfolder %+v\n-----\n", c.FullPath().Folder(false))

				// TODO(ashmrtn): Remove when LocationPath is made part of BackupCollection
				// interface.
				if !assert.Implements(t, (*data.LocationPather)(nil), c) {
					continue
				}

				loc := c.(data.LocationPather).LocationPath().String()

				fmt.Printf("\n-----\nloc %+v\n-----\n", c.(data.LocationPather).LocationPath().String())

				require.NotEmpty(t, loc)

				delete(test.channelNames, loc)
			}

			assert.Empty(t, test.channelNames)
		})
	}
}
