package groups

import (
	"context"
	"io"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/groups/testdata"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

// ---------------------------------------------------------------------------
// mocks
// ---------------------------------------------------------------------------

var _ backupHandler[models.Channelable, models.ChatMessageable] = &mockChannelsBH{}

//lint:ignore U1000 false linter issue due to generics
type mockChannelsBH struct {
	channels      []models.Channelable
	conversations []models.Conversationable
	messageIDs    []string
	deletedMsgIDs []string
	messagesErr   error
	messages      map[string]models.ChatMessageable
	posts         map[string]models.Postable
	info          map[string]*details.GroupsInfo
	getMessageErr map[string]error
	doNotInclude  bool
}

//lint:ignore U1000 false linter issue due to generics
func (bh mockChannelsBH) augmentItemInfo(
	*details.GroupsInfo,
	models.Channelable,
) {
	// no-op
}

//lint:ignore U1000 false linter issue due to generics
func (bh mockChannelsBH) supportsItemMetadata() bool {
	return false
}

func (bh mockChannelsBH) canMakeDeltaQueries() bool {
	return true
}

func (bh mockChannelsBH) containers() []container[models.Channelable] {
	containers := make([]container[models.Channelable], 0, len(bh.channels))

	for _, ch := range bh.channels {
		containers = append(containers, channelContainer(ch))
	}

	return containers
}

//lint:ignore U1000 required for interface compliance
func (bh mockChannelsBH) getContainers(
	context.Context,
	api.CallConfig,
) ([]container[models.Channelable], error) {
	return bh.containers(), nil
}

func (bh mockChannelsBH) getContainerItemIDs(
	_ context.Context,
	_ path.Elements,
	_ string,
	_ api.CallConfig,
) (pagers.AddedAndRemoved, error) {
	idRes := make(map[string]time.Time, len(bh.messageIDs))

	for _, id := range bh.messageIDs {
		idRes[id] = time.Time{}
	}

	aar := pagers.AddedAndRemoved{
		Added:         idRes,
		Removed:       bh.deletedMsgIDs,
		ValidModTimes: true,
		DU:            pagers.DeltaUpdate{},
	}

	return aar, bh.messagesErr
}

//lint:ignore U1000 required for interface compliance
func (bh mockChannelsBH) includeContainer(
	models.Channelable,
	selectors.GroupsScope,
) bool {
	return !bh.doNotInclude
}

func (bh mockChannelsBH) canonicalPath(
	storageDirFolders path.Elements,
	tenantID string,
) (path.Path, error) {
	return storageDirFolders.
		Builder().
		ToDataLayerPath(
			tenantID,
			"protectedResource",
			path.GroupsService,
			path.ChannelMessagesCategory,
			false)
}

//lint:ignore U1000 false linter issue due to generics
func (bh mockChannelsBH) getItem(
	_ context.Context,
	_ string,
	_ path.Elements,
	itemID string,
) (models.ChatMessageable, *details.GroupsInfo, error) {
	return bh.messages[itemID], bh.info[itemID], bh.getMessageErr[itemID]
}

func (bh mockChannelsBH) getItemMetadata(
	_ context.Context,
	_ models.Channelable,
) (io.ReadCloser, int, error) {
	return nil, 0, errMetadataFilesNotSupported
}

//lint:ignore U1000 false linter issue due to generics
func (bh mockChannelsBH) makeTombstones(
	dps metadata.DeltaPaths,
) (map[string]string, error) {
	return makeTombstones(dps), nil
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

// ---------------------------------------------------------------------------
// Channels tests
// ---------------------------------------------------------------------------

func (suite *BackupUnitSuite) TestPopulateCollections() {
	var (
		qp = graph.QueryParams{
			Category:          path.ChannelMessagesCategory, // doesn't matter which one we use.
			ProtectedResource: inMock.NewProvider("group_id", "user_name"),
			TenantID:          suite.creds.AzureTenantID,
		}
		statusUpdater = func(*support.ControllerOperationStatus) {}
	)

	table := []struct {
		name                string
		mock                mockChannelsBH
		expectErr           require.ErrorAssertionFunc
		expectColls         int
		expectNewColls      int
		expectMetadataColls int
	}{
		{
			name: "happy path, one container",
			mock: mockChannelsBH{
				channels:   testdata.StubChannels("one"),
				messageIDs: []string{"msg-one"},
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      1,
			expectMetadataColls: 1,
		},
		{
			name: "happy path, one container, only deleted messages",
			mock: mockChannelsBH{
				channels:      testdata.StubChannels("one"),
				deletedMsgIDs: []string{"msg-one"},
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      1,
			expectMetadataColls: 1,
		},
		{
			name: "happy path, many containers",
			mock: mockChannelsBH{
				channels:   testdata.StubChannels("one", "two"),
				messageIDs: []string{"msg-one"},
			},
			expectErr:           require.NoError,
			expectColls:         3,
			expectNewColls:      2,
			expectMetadataColls: 1,
		},
		{
			name: "no containers pass scope",
			mock: mockChannelsBH{
				channels:     testdata.StubChannels("one"),
				doNotInclude: true,
			},
			expectErr:           require.NoError,
			expectColls:         1,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
		{
			name:                "no channels",
			mock:                mockChannelsBH{},
			expectErr:           require.NoError,
			expectColls:         1,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
		{
			name: "no channel messages",
			mock: mockChannelsBH{
				channels: testdata.StubChannels("one"),
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      1,
			expectMetadataColls: 1,
		},
		{
			name: "err: deleted in flight",
			mock: mockChannelsBH{
				channels:    testdata.StubChannels("one"),
				messagesErr: core.ErrNotFound,
			},
			expectErr:           require.Error,
			expectColls:         1,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
		{
			name: "err: other error",
			mock: mockChannelsBH{
				channels:    testdata.StubChannels("one"),
				messagesErr: assert.AnError,
			},
			expectErr:           require.Error,
			expectColls:         1,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrlOpts := control.Options{FailureHandling: control.FailFast}

			collections, err := populateCollections(
				ctx,
				qp,
				test.mock,
				statusUpdater,
				test.mock.containers(),
				selectors.NewGroupsBackup(nil).Channels(selectors.Any())[0],
				nil,
				false,
				ctrlOpts,
				count.New(),
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
		})
	}
}

func (suite *BackupUnitSuite) TestPopulateCollections_incremental() {
	var (
		qp = graph.QueryParams{
			Category:          path.ChannelMessagesCategory, // doesn't matter which one we use.
			ProtectedResource: inMock.NewProvider("group_id", "user_name"),
			TenantID:          suite.creds.AzureTenantID,
		}
		statusUpdater = func(*support.ControllerOperationStatus) {}
		allScope      = selectors.NewGroupsBackup(nil).Channels(selectors.Any())[0]
	)

	chanPath, err := path.Build("tid", "grp", path.GroupsService, path.ChannelMessagesCategory, false, "chan")
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		name                string
		mock                mockChannelsBH
		deltaPaths          metadata.DeltaPaths
		expectErr           require.ErrorAssertionFunc
		expectColls         int
		expectNewColls      int
		expectTombstoneCols int
		expectMetadataColls int
	}{
		{
			name: "non incremental",
			mock: mockChannelsBH{
				channels:   testdata.StubChannels("chan"),
				messageIDs: []string{"msg"},
			},
			deltaPaths:          metadata.DeltaPaths{},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      1,
			expectTombstoneCols: 0,
			expectMetadataColls: 1,
		},
		{
			name: "incremental",
			mock: mockChannelsBH{
				channels:      testdata.StubChannels("chan"),
				deletedMsgIDs: []string{"msg"},
			},
			deltaPaths: metadata.DeltaPaths{
				"chan": {
					Delta: "chan",
					Path:  chanPath.String(),
				},
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      0,
			expectTombstoneCols: 0,
			expectMetadataColls: 1,
		},
		{
			name: "incremental no new messages",
			mock: mockChannelsBH{
				channels: testdata.StubChannels("chan"),
			},
			deltaPaths: metadata.DeltaPaths{
				"chan": {
					Delta: "chan",
					Path:  chanPath.String(),
				},
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      0,
			expectTombstoneCols: 0,
			expectMetadataColls: 1,
		},
		{
			name: "incremental deleted channel",
			mock: mockChannelsBH{
				channels: testdata.StubChannels(),
			},
			deltaPaths: metadata.DeltaPaths{
				"chan": {
					Delta: "chan",
					Path:  chanPath.String(),
				},
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      0,
			expectTombstoneCols: 1,
			expectMetadataColls: 1,
		},
		{
			name: "incremental new and deleted channel",
			mock: mockChannelsBH{
				channels:   testdata.StubChannels("chan2"),
				messageIDs: []string{"msg"},
			},
			deltaPaths: metadata.DeltaPaths{
				"chan": {
					Delta: "chan",
					Path:  chanPath.String(),
				},
			},
			expectErr:           require.NoError,
			expectColls:         3,
			expectNewColls:      1,
			expectTombstoneCols: 1,
			expectMetadataColls: 1,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrlOpts := control.Options{FailureHandling: control.FailFast}

			collections, err := populateCollections(
				ctx,
				qp,
				test.mock,
				statusUpdater,
				test.mock.containers(),
				allScope,
				test.deltaPaths,
				false,
				ctrlOpts,
				count.New(),
				fault.New(true))
			test.expectErr(t, err, clues.ToCore(err))
			assert.Len(t, collections, test.expectColls, "number of collections")

			// collection assertions

			tombstones, news, metadatas, doNotMerges := 0, 0, 0, 0
			for _, c := range collections {
				if c.FullPath() != nil && c.FullPath().Service() == path.GroupsMetadataService {
					metadatas++
					continue
				}

				if c.State() == data.DeletedState {
					tombstones++
				}

				if c.State() == data.NewState {
					news++
				}

				if c.DoNotMergeItems() {
					doNotMerges++
				}
			}

			assert.Equal(t, test.expectNewColls, news, "new collections")
			assert.Equal(t, test.expectTombstoneCols, tombstones, "tombstone collections")
			assert.Equal(t, test.expectMetadataColls, metadatas, "metadata collections")
		})
	}
}

// ---------------------------------------------------------------------------
// Conversations tests
// ---------------------------------------------------------------------------

var _ backupHandler[models.Conversationable, models.Postable] = &mockConversationsBH{}

//lint:ignore U1000 false linter issue due to generics
type mockConversationsBH struct {
	conversations []models.Conversationable
	// Assume all conversations have the same thread object under them for simplicty.
	// It doesn't impact the tests.
	thread         models.ConversationThreadable
	postIDs        []string
	deletedPostIDs []string
	PostsErr       error
	Posts          map[string]models.Postable
	info           map[string]*details.GroupsInfo
	getPostErr     map[string]error
	doNotInclude   bool
}

//lint:ignore U1000 false linter issue due to generics
func (bh mockConversationsBH) augmentItemInfo(
	*details.GroupsInfo,
	models.Conversationable,
) {
	// no-op
}

func (bh mockConversationsBH) canMakeDeltaQueries() bool {
	return false
}

func (bh mockConversationsBH) containers() []container[models.Conversationable] {
	containers := make([]container[models.Conversationable], 0, len(bh.conversations))

	for _, ch := range bh.conversations {
		containers = append(containers, conversationThreadContainer(ch, bh.thread))
	}

	return containers
}

//lint:ignore U1000 required for interface compliance
func (bh mockConversationsBH) getContainers(
	context.Context,
	api.CallConfig,
) ([]container[models.Conversationable], error) {
	return bh.containers(), nil
}

func (bh mockConversationsBH) getContainerItemIDs(
	_ context.Context,
	_ path.Elements,
	_ string,
	_ api.CallConfig,
) (pagers.AddedAndRemoved, error) {
	idRes := make(map[string]time.Time, len(bh.postIDs))

	for _, id := range bh.postIDs {
		idRes[id] = time.Time{}
	}

	aar := pagers.AddedAndRemoved{
		Added:         idRes,
		Removed:       bh.deletedPostIDs,
		ValidModTimes: true,
		DU:            pagers.DeltaUpdate{},
	}

	return aar, bh.PostsErr
}

//lint:ignore U1000 required for interface compliance
func (bh mockConversationsBH) includeContainer(
	models.Conversationable,
	selectors.GroupsScope,
) bool {
	return !bh.doNotInclude
}

func (bh mockConversationsBH) canonicalPath(
	storageDirFolders path.Elements,
	tenantID string,
) (path.Path, error) {
	return storageDirFolders.
		Builder().
		ToDataLayerPath(
			tenantID,
			"protectedResource",
			path.GroupsService,
			path.ConversationPostsCategory,
			false)
}

//lint:ignore U1000 false linter issue due to generics
func (bh mockConversationsBH) getItem(
	_ context.Context,
	_ string,
	_ path.Elements,
	itemID string,
) (models.Postable, *details.GroupsInfo, error) {
	return bh.Posts[itemID], bh.info[itemID], bh.getPostErr[itemID]
}

//lint:ignore U1000 false linter issue due to generics
func (bh mockConversationsBH) supportsItemMetadata() bool {
	return true
}

func (bh mockConversationsBH) getItemMetadata(
	_ context.Context,
	_ models.Conversationable,
) (io.ReadCloser, int, error) {
	return nil, 0, nil
}

//lint:ignore U1000 false linter issue due to generics
func (bh mockConversationsBH) makeTombstones(
	dps metadata.DeltaPaths,
) (map[string]string, error) {
	r := make(map[string]string, len(dps))

	for id, v := range dps {
		elems := path.Split(id)
		if len(elems) != 2 {
			return nil, clues.New("invalid prev path")
		}

		r[elems[0]] = v.Path
	}

	return r, nil
}

func (suite *BackupUnitSuite) TestPopulateCollections_Conversations() {
	var (
		qp = graph.QueryParams{
			Category:          path.ConversationPostsCategory, // doesn't matter which one we use.
			ProtectedResource: inMock.NewProvider("group_id", "user_name"),
			TenantID:          suite.creds.AzureTenantID,
		}
		statusUpdater = func(*support.ControllerOperationStatus) {}
	)

	table := []struct {
		name                string
		mock                mockConversationsBH
		expectErr           require.ErrorAssertionFunc
		expectColls         int
		expectNewColls      int
		expectMetadataColls int
	}{
		{
			name: "happy path, one container",
			mock: mockConversationsBH{
				conversations: testdata.StubConversations("one"),
				thread:        testdata.StubConversationThreads("t-one")[0],
				postIDs:       []string{"msg-one"},
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      1,
			expectMetadataColls: 1,
		},
		{
			name: "happy path, one container, only deleted messages",
			mock: mockConversationsBH{
				conversations:  testdata.StubConversations("one"),
				thread:         testdata.StubConversationThreads("t-one")[0],
				deletedPostIDs: []string{"msg-one"},
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      1,
			expectMetadataColls: 1,
		},
		{
			name: "happy path, many containers",
			mock: mockConversationsBH{
				conversations: testdata.StubConversations("one", "two"),
				thread:        testdata.StubConversationThreads("t-one")[0],
				postIDs:       []string{"msg-one"},
			},
			expectErr:           require.NoError,
			expectColls:         3,
			expectNewColls:      2,
			expectMetadataColls: 1,
		},
		{
			name: "no containers pass scope",
			mock: mockConversationsBH{
				conversations: testdata.StubConversations("one"),
				thread:        testdata.StubConversationThreads("t-one")[0],
				doNotInclude:  true,
			},
			expectErr:           require.NoError,
			expectColls:         1,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
		{
			name:                "no conversations",
			mock:                mockConversationsBH{},
			expectErr:           require.NoError,
			expectColls:         1,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
		{
			name: "no conv posts",
			mock: mockConversationsBH{
				conversations: testdata.StubConversations("one"),
				thread:        testdata.StubConversationThreads("t-one")[0],
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      1,
			expectMetadataColls: 1,
		},
		{
			name: "err: deleted in flight",
			mock: mockConversationsBH{
				conversations: testdata.StubConversations("one"),
				thread:        testdata.StubConversationThreads("t-one")[0],
				PostsErr:      core.ErrNotFound,
			},
			expectErr:           require.Error,
			expectColls:         1,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
		{
			name: "err: other error",
			mock: mockConversationsBH{
				conversations: testdata.StubConversations("one"),
				thread:        testdata.StubConversationThreads("t-one")[0],
				PostsErr:      assert.AnError,
			},
			expectErr:           require.Error,
			expectColls:         1,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrlOpts := control.Options{FailureHandling: control.FailFast}

			collections, err := populateCollections(
				ctx,
				qp,
				test.mock,
				statusUpdater,
				test.mock.containers(),
				selectors.NewGroupsBackup(nil).Channels(selectors.Any())[0],
				nil,
				false,
				ctrlOpts,
				count.New(),
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
		})
	}
}

func (suite *BackupUnitSuite) TestPopulateCollections_ConversationsIncremental() {
	var (
		qp = graph.QueryParams{
			Category:          path.ConversationPostsCategory, // doesn't matter which one we use.
			ProtectedResource: inMock.NewProvider("group_id", "user_name"),
			TenantID:          suite.creds.AzureTenantID,
		}
		statusUpdater = func(*support.ControllerOperationStatus) {}
		allScope      = selectors.NewGroupsBackup(nil).Conversation(selectors.Any())[0]
	)

	convPath, err := path.Build("t", "g", path.GroupsService, path.ConversationPostsCategory, false, "conv0", "thread0")
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		name                string
		mock                mockConversationsBH
		deltaPaths          metadata.DeltaPaths
		expectErr           require.ErrorAssertionFunc
		expectColls         int
		expectNewColls      int
		expectTombstoneCols int
		expectMetadataColls int
	}{
		{
			name: "non incremental",
			mock: mockConversationsBH{
				conversations: testdata.StubConversations("conv0"),
				thread:        testdata.StubConversationThreads("t0")[0],
				postIDs:       []string{"msg"},
			},
			deltaPaths:          metadata.DeltaPaths{},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      1,
			expectTombstoneCols: 0,
			expectMetadataColls: 1,
		},
		{
			name: "incremental",
			mock: mockConversationsBH{
				conversations: testdata.StubConversations("conv0"),
				thread:        testdata.StubConversationThreads("t0")[0],
				postIDs:       []string{"msg"},
			},
			deltaPaths: metadata.DeltaPaths{
				"conv0/thread0": {
					Path: convPath.String(),
				},
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      1, // No delta support
			expectTombstoneCols: 0,
			expectMetadataColls: 1,
		},
		{
			name: "incremental no new posts",
			mock: mockConversationsBH{
				conversations: testdata.StubConversations("conv0"),
				thread:        testdata.StubConversationThreads("t0")[0],
			},
			deltaPaths: metadata.DeltaPaths{
				"conv0/thread0": {
					Path: convPath.String(),
				},
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      1, // No delta support
			expectTombstoneCols: 0,
			expectMetadataColls: 1,
		},
		{
			name: "incremental deleted conversation",
			mock: mockConversationsBH{
				conversations: testdata.StubConversations(),
			},
			deltaPaths: metadata.DeltaPaths{
				"conv0/thread0": {
					Path: convPath.String(),
				},
			},
			expectErr:           require.NoError,
			expectColls:         2,
			expectNewColls:      0,
			expectTombstoneCols: 1,
			expectMetadataColls: 1,
		},
		{
			name: "incremental new and deleted conversations",
			mock: mockConversationsBH{
				conversations: testdata.StubConversations("conv1"),
				thread:        testdata.StubConversationThreads("t1")[0],
				postIDs:       []string{"msg"},
			},
			deltaPaths: metadata.DeltaPaths{
				"conv0/thread0": {
					Path: convPath.String(),
				},
			},
			expectErr:           require.NoError,
			expectColls:         3,
			expectNewColls:      1,
			expectTombstoneCols: 1,
			expectMetadataColls: 1,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrlOpts := control.Options{FailureHandling: control.FailFast}

			collections, err := populateCollections(
				ctx,
				qp,
				test.mock,
				statusUpdater,
				test.mock.containers(),
				allScope,
				test.deltaPaths,
				false,
				ctrlOpts,
				count.New(),
				fault.New(true))
			test.expectErr(t, err, clues.ToCore(err))
			assert.Len(t, collections, test.expectColls, "number of collections")

			// collection assertions

			tombstones, news, metadatas, doNotMerges := 0, 0, 0, 0
			for _, c := range collections {
				if c.FullPath() != nil && c.FullPath().Service() == path.GroupsMetadataService {
					metadatas++
					continue
				}

				if c.State() == data.DeletedState {
					tombstones++
				}

				if c.State() == data.NewState {
					news++
				}

				if c.DoNotMergeItems() {
					doNotMerges++
				}
			}

			assert.Equal(t, test.expectNewColls, news, "new collections")
			assert.Equal(t, test.expectTombstoneCols, tombstones, "tombstone collections")
			assert.Equal(t, test.expectMetadataColls, metadatas, "metadata collections")
		})
	}
}

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

	suite.ac, err = api.NewClient(
		creds,
		control.DefaultOptions(),
		count.New())
	require.NoError(t, err, clues.ToCore(err))

	suite.tenantID = creds.AzureTenantID
}

func (suite *BackupIntgSuite) TestCreateCollections() {
	var (
		protectedResource = tconfig.M365TeamID(suite.T())
		resources         = []string{protectedResource}
		handler           = NewChannelBackupHandler(protectedResource, suite.ac.Channels())
	)

	tests := []struct {
		name         string
		scope        selectors.GroupsScope
		channelNames map[string]struct{}
	}{
		{
			name:  "channel messages",
			scope: selTD.GroupsBackupChannelScope(selectors.NewGroupsBackup(resources))[0],
			channelNames: map[string]struct{}{
				selTD.TestChannelName: {},
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrlOpts := control.DefaultOptions()

			sel := selectors.NewGroupsBackup([]string{protectedResource})
			sel.Include(selTD.GroupsBackupChannelScope(sel))

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: version.NoBackup,
				Options:           ctrlOpts,
				ProtectedResource: inMock.NewProvider(protectedResource, protectedResource),
				Selector:          sel.Selector,
			}

			collections, _, err := CreateCollections(
				ctx,
				bpc,
				handler,
				suite.tenantID,
				test.scope,
				func(status *support.ControllerOperationStatus) {},
				false,
				count.New(),
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			require.NotEmpty(t, collections, "must have at least one collection")

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
