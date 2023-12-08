package groups

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
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

func (bh conversationsBackupHandler) canMakeDeltaQueries() bool {
	// not supported for conversations
	return false
}

//lint:ignore U1000 required for interface compliance
func (bh conversationsBackupHandler) getContainers(
	ctx context.Context,
	cc api.CallConfig,
) ([]container[models.Conversationable], error) {
	convs, err := bh.ac.GetConversations(ctx, bh.protectedResource, cc)
	if err != nil {
		return nil, clues.Wrap(err, "getting conversations")
	}

	results := []container[models.Conversationable]{}

	for _, conv := range convs {
		ictx := clues.Add(ctx, "conversation_id", ptr.Val(conv.GetId()))

		threads, err := bh.ac.GetConversationThreads(
			ictx,
			bh.protectedResource,
			ptr.Val(conv.GetId()),
			cc)
		if err != nil {
			return nil, clues.Wrap(err, "getting threads in conversation")
		}

		for _, thread := range threads {
			results = append(results, conversationThreadContainer(conv, thread))
		}
	}

	return results, nil
}

func (bh conversationsBackupHandler) getContainerItemIDs(
	ctx context.Context,
	containerPath path.Elements,
	_ string,
	cc api.CallConfig,
) (pagers.AddedAndRemoved, error) {
	return bh.ac.GetConversationThreadPostIDs(
		ctx,
		bh.protectedResource,
		containerPath[0],
		containerPath[1],
		cc)
}

//lint:ignore U1000 required for interface compliance
func (bh conversationsBackupHandler) includeContainer(
	conv models.Conversationable,
	scope selectors.GroupsScope,
) bool {
	return scope.Matches(selectors.GroupsConversation, ptr.Val(conv.GetTopic()))
}

func (bh conversationsBackupHandler) canonicalPath(
	storageDirFolders path.Elements,
	tenantID string,
) (path.Path, error) {
	return storageDirFolders.
		Builder().
		ToDataLayerPath(
			tenantID,
			bh.protectedResource,
			path.GroupsService,
			path.ConversationPostsCategory,
			false)
}

func (bh conversationsBackupHandler) PathPrefix(tenantID string) (path.Path, error) {
	return path.Build(
		tenantID,
		bh.protectedResource,
		path.GroupsService,
		path.ConversationPostsCategory,
		false)
}

//nolint:unused
func (bh conversationsBackupHandler) getItem(
	ctx context.Context,
	groupID string,
	containerIDs path.Elements, // expects: [conversationID, threadID]
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

//nolint:unused
func (bh conversationsBackupHandler) augmentItemInfo(
	dgi *details.GroupsInfo,
	c models.Conversationable,
) {
	dgi.Post.Topic = ptr.Val(c.GetTopic())
}

func conversationThreadContainer(
	c models.Conversationable,
	t models.ConversationThreadable,
) container[models.Conversationable] {
	return container[models.Conversationable]{
		storageDirFolders: path.Elements{ptr.Val(c.GetId()), ptr.Val(t.GetId())},
		// microsoft UX doesn't display any sort of container name that would make a reasonable
		// "location" for the posts in the conversation.  We may need to revisit this, perhaps
		// the subject (aka topic) is sufficiently acceptable.
		humanLocation:       path.Elements{ptr.Val(c.GetTopic())},
		canMakeDeltaQueries: false,
		container:           c,
	}
}
