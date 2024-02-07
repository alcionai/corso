package teamschats

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	inMock "github.com/alcionai/canario/src/internal/common/idname/mock"
	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/internal/data"
	"github.com/alcionai/canario/src/internal/m365/collection/teamschats/testdata"
	"github.com/alcionai/canario/src/internal/m365/support"
	"github.com/alcionai/canario/src/internal/operations/inject"
	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/internal/tester/its"
	"github.com/alcionai/canario/src/internal/tester/tconfig"
	"github.com/alcionai/canario/src/internal/version"
	"github.com/alcionai/canario/src/pkg/account"
	"github.com/alcionai/canario/src/pkg/backup/details"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/count"
	"github.com/alcionai/canario/src/pkg/errs/core"
	"github.com/alcionai/canario/src/pkg/fault"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/selectors"
	selTD "github.com/alcionai/canario/src/pkg/selectors/testdata"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph"
)

// ---------------------------------------------------------------------------
// mocks
// ---------------------------------------------------------------------------

var _ backupHandler[models.Chatable] = &mockBackupHandler{}

//lint:ignore U1000 false linter issue due to generics
type mockBackupHandler struct {
	chatsErr        error
	chats           []models.Chatable
	chatMessagesErr error
	chatMessages    map[string][]models.ChatMessageable
	info            map[string]*details.TeamsChatsInfo
	getMessageErr   map[string]error
	doNotInclude    bool
}

//lint:ignore U1000 false linter issue due to generics
func (bh mockBackupHandler) augmentItemInfo(
	*details.TeamsChatsInfo,
	models.Chatable,
) {
	// no-op
}

func (bh mockBackupHandler) container() container[models.Chatable] {
	return chatContainer()
}

//lint:ignore U1000 required for interface compliance
func (bh mockBackupHandler) getContainer(
	context.Context,
	api.CallConfig,
) (container[models.Chatable], error) {
	return chatContainer(), nil
}

func (bh mockBackupHandler) getItemIDs(
	_ context.Context,
	_ api.CallConfig,
) ([]models.Chatable, error) {
	return bh.chats, bh.chatsErr
}

//lint:ignore U1000 required for interface compliance
func (bh mockBackupHandler) includeItem(
	models.Chatable,
	selectors.TeamsChatsScope,
) bool {
	return !bh.doNotInclude
}

func (bh mockBackupHandler) CanonicalPath() (path.Path, error) {
	return path.BuildPrefix(
		"tenant",
		"protectedResource",
		path.TeamsChatsService,
		path.ChatsCategory)
}

//lint:ignore U1000 false linter issue due to generics
func (bh mockBackupHandler) getItem(
	_ context.Context,
	_ string,
	itemID string,
) (models.Chatable, *details.TeamsChatsInfo, error) {
	chat := models.NewChat()

	chat.SetId(ptr.To(itemID))
	chat.SetTopic(ptr.To(itemID))
	chat.SetMessages(bh.chatMessages[itemID])

	return chat, bh.info[itemID], bh.getMessageErr[itemID]
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
			Category:          path.ChatsCategory, // doesn't matter which one we use.
			ProtectedResource: inMock.NewProvider("user_id", "user_name"),
			TenantID:          suite.creds.AzureTenantID,
		}
		statusUpdater = func(*support.ControllerOperationStatus) {}
	)

	table := []struct {
		name       string
		mock       mockBackupHandler
		expectErr  require.ErrorAssertionFunc
		expectColl require.ValueAssertionFunc
	}{
		{
			name: "happy path, one chat",
			mock: mockBackupHandler{
				chats: testdata.StubChats("one"),
				chatMessages: map[string][]models.ChatMessageable{
					"one": testdata.StubChatMessages("msg-one"),
				},
			},
			expectErr:  require.NoError,
			expectColl: require.NotNil,
		},
		{
			name: "happy path, many chats",
			mock: mockBackupHandler{
				chats: testdata.StubChats("one", "two"),
				chatMessages: map[string][]models.ChatMessageable{
					"one": testdata.StubChatMessages("msg-one"),
					"two": testdata.StubChatMessages("msg-two"),
				},
			},
			expectErr:  require.NoError,
			expectColl: require.NotNil,
		},
		{
			name: "no chats pass scope",
			mock: mockBackupHandler{
				chats:        testdata.StubChats("one"),
				doNotInclude: true,
			},
			expectErr:  require.NoError,
			expectColl: require.NotNil,
		},
		{
			name:       "no chats",
			mock:       mockBackupHandler{},
			expectErr:  require.NoError,
			expectColl: require.NotNil,
		},
		{
			name: "no chat messages",
			mock: mockBackupHandler{
				chats: testdata.StubChats("one"),
			},
			expectErr:  require.NoError,
			expectColl: require.NotNil,
		},
		{
			name: "err: deleted in flight",
			mock: mockBackupHandler{
				chats:    testdata.StubChats("one"),
				chatsErr: core.ErrNotFound,
			},
			expectErr:  require.Error,
			expectColl: require.Nil,
		},
		{
			name: "err enumerating chats",
			mock: mockBackupHandler{
				chats:    testdata.StubChats("one"),
				chatsErr: assert.AnError,
			},
			expectErr:  require.Error,
			expectColl: require.Nil,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrlOpts := control.Options{FailureHandling: control.FailFast}

			result, err := populateCollection(
				ctx,
				qp,
				test.mock,
				statusUpdater,
				test.mock.container(),
				selectors.NewTeamsChatsBackup(nil).Chats(selectors.Any())[0],
				false,
				ctrlOpts,
				count.New(),
				fault.New(true))
			test.expectErr(t, err, clues.ToCore(err))
			test.expectColl(t, result)

			if err != nil || result == nil {
				return
			}

			// collection assertions

			assert.NotEqual(
				t,
				result.FullPath().Service(),
				path.TeamsChatsMetadataService,
				"should not contain metadata collections")
			assert.NotEqual(t, result.State(), data.DeletedState, "no tombstones should be produced")
			assert.Equal(t, result.State(), data.NotMovedState)
			assert.False(t, result.DoNotMergeItems(), "doNotMergeItems should always be false")
		})
	}
}

// ---------------------------------------------------------------------------
// Integration tests
// ---------------------------------------------------------------------------

type BackupIntgSuite struct {
	tester.Suite
	m365 its.M365IntgTestSetup
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
	suite.m365 = its.GetM365(t)

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)
}

func (suite *BackupIntgSuite) TestCreateCollections() {
	var (
		tenant            = suite.m365.TenantID
		protectedResource = suite.m365.Group.ID
		resources         = []string{protectedResource}
		handler           = NewUsersChatsBackupHandler(tenant, protectedResource, suite.m365.AC.Chats())
	)

	tests := []struct {
		name      string
		scope     selectors.TeamsChatsScope
		chatNames map[string]struct{}
	}{
		{
			name:  "chat messages",
			scope: selTD.TeamsChatsBackupChatScope(selectors.NewTeamsChatsBackup(resources))[0],
			chatNames: map[string]struct{}{
				selTD.TestChatTopic: {},
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrlOpts := control.DefaultOptions()

			sel := selectors.NewTeamsChatsBackup([]string{suite.m365.Group.ID})
			sel.Include(selTD.TeamsChatsBackupChatScope(sel))

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: version.NoBackup,
				Options:           ctrlOpts,
				ProtectedResource: suite.m365.Group.Provider,
				Selector:          sel.Selector,
			}

			collections, _, err := CreateCollections(
				ctx,
				bpc,
				handler,
				suite.m365.TenantID,
				test.scope,
				func(status *support.ControllerOperationStatus) {},
				false,
				count.New(),
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			require.NotEmpty(t, collections, "must have at least one collection")

			for _, c := range collections {
				if c.FullPath().Service() == path.TeamsChatsMetadataService {
					continue
				}

				require.Empty(t, c.FullPath().Folder(false), "all items should be stored at the root")

				locp, ok := c.(data.LocationPather)

				if ok {
					loc := locp.LocationPath().String()
					require.Empty(t, loc, "no items should have locations")
				}
			}

			assert.Len(t, collections, 2, "should have the root folder collection and metadata collection")
		})
	}
}
