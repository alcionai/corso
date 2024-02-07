package api

import (
	"context"
	"strings"

	"github.com/alcionai/clues"
	"github.com/jaytaylor/html2text"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/internal/common/sanitize"
	"github.com/alcionai/canario/src/internal/common/str"
	"github.com/alcionai/canario/src/pkg/backup/details"
	"github.com/alcionai/canario/src/pkg/logger"
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
) (models.Postable, *details.GroupsInfo, error) {
	config := &groups.ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &groups.ItemConversationsItemThreadsItemPostsPostItemRequestBuilderGetQueryParameters{},
	}

	// Expand inReplyTo property to additionally get the parent post contents.
	// This will be persisted as part of post data.
	//
	// This additional data will be useful for building a reply tree if we decide to
	// do in-order restore/export in future.
	config.QueryParameters.Expand = append(config.QueryParameters.Expand, "inReplyTo")

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
		return nil, nil, clues.Stack(err)
	}

	preview, contentLen, err := getConversationPostContentPreview(post)
	if err != nil {
		preview = "malformed or unparseable content body: " + preview
	}

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
		// Similar to exchange, a failure can happen if a post has a lot of attachments.
		// We don't have a fallback option here to fetch attachments one by one. See
		// issue #4991.
		//
		// Resort to failing the post backup for now since we don't know yet how this
		// error might manifest itself for posts.
		logger.CtxErr(ctx, err).Info("failed to get post attachments")

		return nil, nil, clues.Stack(err)
	}

	contentLen += totalSize

	post.SetAttachments(attachments)

	return post, conversationPostInfo(post, contentLen, preview), clues.Stack(err).OrNil()
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
		return nil, 0, clues.Stack(err)
	}

	attachments := post.GetAttachments()

	for _, a := range attachments {
		totalSize += int64(ptr.Val(a.GetSize()))
		result = append(result, a)
	}

	return result, totalSize, nil
}

func bytesToPostable(body []byte) (serialization.Parsable, error) {
	v, err := CreateFromBytes(body, models.CreatePostFromDiscriminatorValue)
	if err != nil {
		if !strings.Contains(err.Error(), invalidJSON) {
			return nil, clues.Wrap(err, "deserializing bytes to message")
		}

		// If the JSON was invalid try sanitizing and deserializing again.
		// Sanitizing should transform characters < 0x20 according to the spec where
		// possible. The resulting JSON may still be invalid though.
		body = sanitize.JSONBytes(body)
		v, err = CreateFromBytes(body, models.CreatePostFromDiscriminatorValue)
	}

	return v, clues.Stack(err).OrNil()
}

func BytesToPostable(body []byte) (models.Postable, error) {
	v, err := bytesToPostable(body)
	if err != nil {
		return nil, clues.Stack(err)
	}

	return v.(models.Postable), nil
}
