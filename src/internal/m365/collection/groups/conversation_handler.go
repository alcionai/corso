package groups

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

var _ backupHandler[models.Conversationable, models.Postable] = &conversationsBackupHandler{}

type conversationsBackupHandler struct {
	ac                api.Conversations
	protectedResource string
}

func NewConversationBackupHandler(
	protectedResource string,
	ac api.Conversations,
) conversationsBackupHandler {
	return conversationsBackupHandler{
		ac:                ac,
		protectedResource: protectedResource,
	}
}

func (bh conversationsBackupHandler) canMakeDeltaQueries(models.Conversationable) bool {
	return false
}

func (bh conversationsBackupHandler) getContainers(
	ctx context.Context,
) ([]models.Conversationable, error) {
	return bh.ac.GetConversations(ctx, bh.protectedResource, api.CallConfig{})
}

func (bh conversationsBackupHandler) getContainerItemIDs(
	ctx context.Context,
	conversationID, prevDelta string,
	cc api.CallConfig,
) (pagers.AddedAndRemoved, error) {
	return bh.ac.GetConversationThreadPostIDs(
		ctx,
		bh.protectedResource,
		conversationID,
		prevDelta,
		cc)
}

func (bh conversationsBackupHandler) includeContainer(
	ctx context.Context,
	qp graph.QueryParams,
	conv models.Conversationable,
	scope selectors.GroupsScope,
) bool {
	return scope.Matches(selectors.GroupsConversation, ptr.Val(conv.GetTopic()))
}

func (bh conversationsBackupHandler) canonicalPath(
	folders *path.Builder,
	tenantID string,
) (path.Path, error) {
	return folders.
		ToDataLayerPath(
			tenantID,
			bh.protectedResource,
			path.GroupsService,
			path.ConversationPostsCategory,
			false)
}

func (bh conversationsBackupHandler) locationPath(c models.Conversationable) *path.Builder {
	return &path.Builder{}
}

func (bh conversationsBackupHandler) PathPrefix(tenantID string) (path.Path, error) {
	return path.Build(
		tenantID,
		bh.protectedResource,
		path.GroupsService,
		path.ConversationPostsCategory,
		false)
}

func (bh conversationsBackupHandler) GetItem(
	ctx context.Context,
	groupID string,
	containerIDs path.Elements,
	postID string,
) (models.Postable, *details.GroupsInfo, error) {
	return bh.ac.GetConversationPost(
		ctx,
		groupID,
		containerIDs[0],
		containerIDs[1],
		postID,
		api.CallConfig{})
}
