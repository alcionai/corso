package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/jaytaylor/html2text"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
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

	preview, contentLen, err := getConversationPostContentPreview(post)
	if err != nil {
		preview = "malformed or unparseable html" + preview
	}

	// TODO(pandeyabs): Is this the best way to get embedded attachments?
	// Attachmentable has inline field also. what about that?
	if !ptr.Val(post.GetHasAttachments()) && !HasAttachments(post.GetBody()) {
		return post, conversationPostInfo(post, contentLen, preview), nil
	}

	attachments, totalSize, err := c.getAttachments(
		ctx,
		groupID,
		conversationID,
		threadID,
		postID)
	if err != nil {
		// A failure can be caused by having a lot of attachments.
		// If that happens, we can progres with a two-step approach of:
		// 1. getting all attachment IDs.
		// 2. fetching each attachment individually.
		logger.CtxErr(ctx, err).Info("falling back to fetching attachments by id")

		// attachments, totalSize, err = c.getAttachmentsIterated(ctx, userID, mailID, immutableIDs, errs)
		// if err != nil {
		// 	return nil, nil, clues.Stack(err)
		// }
	}

	contentLen += totalSize

	post.SetAttachments(attachments)

	return post, conversationPostInfo(post, contentLen, preview), graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func conversationPostInfo(
	post models.Postable,
	size int64,
	preview string,
) *details.GroupsInfo {
	if post == nil {
		return nil
	}

	var sender string
	if post.GetSender() != nil && post.GetSender().GetEmailAddress() != nil {
		sender = ptr.Val(post.GetSender().GetEmailAddress().GetAddress())
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

// getAttachments attempts to get all attachments, including their content, in a singe query.
// microsoft.graph.itemattachment/item" is returning 403.
// Compare with mail api. also try with attachments and Attachment() endpoint.
func (c Conversations) getAttachments(
	ctx context.Context,
	groupID, conversationID, threadID, postID string,
) ([]models.Attachmentable, int64, error) {
	var (
		result    = []models.Attachmentable{}
		totalSize int64
	)

	cfg := &groups.ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &groups.ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetQueryParameters{
			Expand: []string{"attachments"},
		},
	}

	post, err := c.LargeItem.
		Client().
		Groups().
		ByGroupId(groupID).
		Conversations().
		ByConversationId(conversationID).
		Threads().
		ByConversationThreadId(threadID).
		Posts().
		ByPostId(postID).
		Get(ctx, cfg)
	if err != nil {
		return nil, 0, graph.Stack(ctx, err)
	}

	attachments := post.GetAttachments()

	for _, a := range attachments {
		totalSize += int64(ptr.Val(a.GetSize()))
		result = append(result, a)
	}

	return result, totalSize, nil
}

// // getAttachments attempts to get all attachments, including their content, in a singe query.
// // microsoft.graph.itemattachment/item" is returning 403.
// func (c Conversations) getAttachments(
// 	ctx context.Context,
// 	groupID, conversationID, threadID, postID string,
// ) ([]models.Attachmentable, int64, error) {
// 	var (
// 		result    = []models.Attachmentable{}
// 		totalSize int64
// 	)

// 	cfg := &groups.ItemConversationsItemThreadsItemPostsItemAttachmentsRequestBuilderGetRequestConfiguration{
// 		QueryParameters: &groups.ItemConversationsItemThreadsItemPostsItemAttachmentsRequestBuilderGetQueryParameters{
// 			Expand: []string{"microsoft.graph.itemattachment/item"},
// 		},
// 	}

// 	attachments, err := c.LargeItem.
// 		Client().
// 		Groups().
// 		ByGroupId(groupID).
// 		Conversations().
// 		ByConversationId(conversationID).
// 		Threads().
// 		ByConversationThreadId(threadID).
// 		Posts().
// 		ByPostId(postID).
// 		Attachments().
// 		Get(ctx, cfg)
// 	if err != nil {
// 		return nil, 0, graph.Stack(ctx, err)
// 	}

// 	for _, a := range attachments.GetValue() {
// 		totalSize += int64(ptr.Val(a.GetSize()))
// 		result = append(result, a)
// 	}

// 	return result, totalSize, nil
// }
