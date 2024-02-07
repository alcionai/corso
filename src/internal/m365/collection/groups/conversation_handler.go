package groups

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/internal/m365/collection/groups/metadata"
	"github.com/alcionai/canario/src/pkg/backup/details"
	deltaPath "github.com/alcionai/canario/src/pkg/backup/metadata"
	"github.com/alcionai/canario/src/pkg/path"
	"github.com/alcionai/canario/src/pkg/selectors"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
	"github.com/alcionai/canario/src/pkg/services/m365/api/pagers"
)

var _ backupHandler[models.Conversationable, models.Postable] = &conversationsBackupHandler{}

type conversationsBackupHandler struct {
	ac                api.Conversations
	protectedResource string
	// SMTP address for the group
	resourceEmail string
}

func NewConversationBackupHandler(
	protectedResource string,
	ac api.Conversations,
	resourceEmail string,
) conversationsBackupHandler {
	return conversationsBackupHandler{
		ac:                ac,
		protectedResource: protectedResource,
		resourceEmail:     resourceEmail,
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

//lint:ignore U1000 false linter issue due to generics
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
		postID)
}

//lint:ignore U1000 false linter issue due to generics
func (bh conversationsBackupHandler) getItemMetadata(
	ctx context.Context,
	c models.Conversationable,
) (io.ReadCloser, int, error) {
	meta := metadata.ConversationPostMetadata{
		Recipients: []string{bh.resourceEmail},
		Topic:      ptr.Val(c.GetTopic()),
	}

	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return nil, 0, clues.WrapWC(ctx, err, "serializing post metadata")
	}

	return io.NopCloser(bytes.NewReader(metaJSON)), len(metaJSON), nil
}

//lint:ignore U1000 false linter issue due to generics
func (bh conversationsBackupHandler) augmentItemInfo(
	dgi *details.GroupsInfo,
	c models.Conversationable,
) {
	// Posts are always sent to the group email address, along with additional
	// recipients if any. Currently we don't have a way to get the unique
	// recipient list for individual posts due to graph api limitations.
	//
	// Store the group mail address in details so that SDK users can use it.
	// This information will also be persisted in metadata files so that we
	// can use it during export & restore.
	dgi.Post.Recipients = []string{bh.resourceEmail}
	dgi.Post.Topic = ptr.Val(c.GetTopic())
}

//lint:ignore U1000 false linter issue due to generics
func (bh conversationsBackupHandler) supportsItemMetadata() bool {
	return true
}

func (bh conversationsBackupHandler) makeTombstones(
	dps deltaPath.DeltaPaths,
) (map[string]string, error) {
	r := make(map[string]string, len(dps))

	for id, v := range dps {
		// ID is of format conversationID/threadID. Tombstones are looked up
		// by conversationID only, so remove the threadID part. This is safe
		// because every conversation has only one thread.
		elems := path.Split(id)
		if len(elems) != 2 {
			return nil, clues.New("invalid prev path")
		}

		r[elems[0]] = v.Path
	}

	// We are assuming a 1:1 mapping between conversations and threads. While
	// this is true today, graph behavior may change in future. Throw an error
	// if the assumption is violated.
	//
	// We cannot catch this error with tests because creating conversations
	// requires delegated access.
	if len(dps) != len(r) {
		return nil, clues.New("multiple threads exist for a conversation")
	}

	return r, nil
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
