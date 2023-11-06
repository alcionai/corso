package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/jaytaylor/html2text"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Conversations() Conversations {
	return Conversations{c}
}

// Conversations is an interface-compliant provider of the client.
type Conversations struct {
	Client
}

// ---------------------------------------------------------------------------
// Item (conversation thread post)
// ---------------------------------------------------------------------------

func (c Conversations) GetConversationPost(
	ctx context.Context,
	groupID, conversationID, threadID, postID string,
	cc CallConfig,
) (models.Postable, *details.GroupsInfo, error) {
	config := &groups.ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &groups.ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetQueryParameters{},
	}

	if len(cc.Select) > 0 {
		config.QueryParameters.Select = cc.Select
	}

	if len(cc.Expand) > 0 {
		config.QueryParameters.Expand = append(config.QueryParameters.Expand, cc.Expand...)
	}

	post, err := c.Stable.
		Client().
		Groups().
		ByGroupId(groupID).
		Conversations().
		ByConversationId(conversationID).
		Threads().
		ByConversationThreadId(threadID).
		Posts().
		ByPostId(postID).
		Get(ctx, config)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	return post, conversationPostInfo(post), graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func conversationPostInfo(
	post models.Postable,
) *details.GroupsInfo {
	if post == nil {
		return nil
	}

	preview, contentLen, err := getConversationPostContentPreview(post)
	if err != nil {
		preview = "malformed or unparseable html" + preview
	}

	var sender string
	if post.GetSender() != nil && post.GetSender().GetEmailAddress() != nil {
		sender = ptr.Val(post.GetSender().GetEmailAddress().GetAddress())
	}

	size := contentLen

	for _, a := range post.GetAttachments() {
		size += int64(ptr.Val(a.GetSize()))
	}

	cpi := details.ConversationPostInfo{
		CreatedAt: ptr.Val(post.GetCreatedDateTime()),
		Creator:   sender,
		Preview:   preview,
		Size:      size,
	}

	return &details.GroupsInfo{
		ItemType: details.GroupsConversationPost,
		Modified: ptr.Val(post.GetLastModifiedDateTime()),
		Post:     cpi,
	}
}

func getConversationPostContentPreview(post models.Postable) (string, int64, error) {
	content, origSize, err := stripConversationPostHTML(post)
	return str.Preview(content, 128), origSize, clues.Stack(err).OrNil()
}

func stripConversationPostHTML(post models.Postable) (string, int64, error) {
	var (
		content  string
		origSize int64
	)

	if post.GetBody() != nil {
		content = ptr.Val(post.GetBody().GetContent())
	}

	origSize = int64(len(content))

	content, err := html2text.FromString(content)

	return content, origSize, clues.Stack(err).OrNil()
}
