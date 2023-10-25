package api

import (
	"context"
	"fmt"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/teams"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/str"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// controller
// ---------------------------------------------------------------------------

func (c Client) Channels() Channels {
	return Channels{c}
}

// Channels is an interface-compliant provider of the client.
type Channels struct {
	Client
}

// ---------------------------------------------------------------------------
// containers
// ---------------------------------------------------------------------------

func (c Channels) GetChannel(
	ctx context.Context,
	teamID, containerID string,
) (models.Channelable, error) {
	config := &teams.ItemChannelsChannelItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &teams.ItemChannelsChannelItemRequestBuilderGetQueryParameters{
			Select: idAnd("displayName"),
		},
	}

	resp, err := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(containerID).
		Get(ctx, config)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

// GetChannelByName fetches a channel by name
func (c Channels) GetChannelByName(
	ctx context.Context,
	teamID, containerName string,
) (models.Channelable, error) {
	ctx = clues.Add(ctx, "channel_name", containerName)

	filter := fmt.Sprintf("displayName eq '%s'", containerName)
	options := &teams.ItemChannelsRequestBuilderGetRequestConfiguration{
		QueryParameters: &teams.ItemChannelsRequestBuilderGetQueryParameters{
			Filter: &filter,
		},
	}

	resp, err := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		Get(ctx, options)
	if err != nil {
		return nil, graph.Stack(ctx, err).WithClues(ctx)
	}

	gv := resp.GetValue()
	if len(gv) == 0 {
		return nil, clues.New("channel not found").WithClues(ctx)
	}

	// We only allow the api to match one channel with the provided name.
	// If we match multiples, we'll eagerly return the first one.
	logger.Ctx(ctx).Debugw("channels matched the name search")

	// Sanity check ID and name
	cal := gv[0]

	if err := CheckIDAndName(cal); err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return cal, nil
}

// ---------------------------------------------------------------------------
// message
// ---------------------------------------------------------------------------

func (c Channels) GetChannelMessage(
	ctx context.Context,
	teamID, channelID, messageID string,
) (models.ChatMessageable, *details.GroupsInfo, error) {
	message, err := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(channelID).
		Messages().
		ByChatMessageId(messageID).
		Get(ctx, nil)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	replies, err := c.GetChannelMessageReplies(ctx, teamID, channelID, messageID)
	if err != nil {
		return nil, nil, graph.Wrap(ctx, err, "retrieving message replies")
	}

	message.SetReplies(replies)

	info := channelMessageInfo(message)

	return message, info, nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func channelMessageInfo(
	msg models.ChatMessageable,
) *details.GroupsInfo {
	var (
		lastReply   models.ChatMessageable
		lastReplyAt time.Time
		modTime     = ptr.OrNow(msg.GetLastModifiedDateTime())
	)

	replies := msg.GetReplies()

	for _, r := range replies {
		cdt := ptr.Val(r.GetCreatedDateTime())
		if cdt.After(lastReplyAt) {
			lastReply = r
			lastReplyAt = ptr.Val(r.GetCreatedDateTime())
		}
	}

	// if the message hasn't been modified since before the most recent
	// reply, set the modified time to the most recent reply.  This ensures
	// we update the message contents to match changes in replies.
	if modTime.Before(lastReplyAt) {
		modTime = lastReplyAt
	}

	preview, contentLen := GetChatMessageContentPreview(msg)

	message := details.ChannelMessageInfo{
		AttachmentNames: GetChatMessageAttachmentNames(msg),
		CreatedAt:       ptr.Val(msg.GetCreatedDateTime()),
		Creator:         GetChatMessageFrom(msg),
		Preview:         preview,
		ReplyCount:      len(replies),
		Size:            contentLen,
	}

	var lr details.ChannelMessageInfo

	if lastReply != nil {
		preview, contentLen = GetChatMessageContentPreview(lastReply)
		lr = details.ChannelMessageInfo{
			AttachmentNames: GetChatMessageAttachmentNames(lastReply),
			CreatedAt:       ptr.Val(lastReply.GetCreatedDateTime()),
			Creator:         GetChatMessageFrom(lastReply),
			Preview:         preview,
			Size:            contentLen,
		}
	}

	return &details.GroupsInfo{
		ItemType:  details.GroupsChannelMessage,
		Modified:  modTime,
		Message:   message,
		LastReply: lr,
	}
}

// CheckIDAndName is a validator that ensures the ID
// and name are populated and not zero valued.
func CheckIDAndName(c models.Channelable) error {
	if c == nil {
		return clues.New("nil container")
	}

	id := ptr.Val(c.GetId())
	if len(id) == 0 {
		return clues.New("container missing ID")
	}

	dn := ptr.Val(c.GetDisplayName())
	if len(dn) == 0 {
		return clues.New("container missing display name").With("container_id", id)
	}

	return nil
}

func GetChatMessageFrom(msg models.ChatMessageable) string {
	from := msg.GetFrom()

	switch true {
	case from == nil:
		return ""
	case from.GetApplication() != nil:
		return ptr.Val(from.GetApplication().GetDisplayName())
	case from.GetDevice() != nil:
		return ptr.Val(from.GetDevice().GetDisplayName())
	case from.GetUser() != nil:
		return ptr.Val(from.GetUser().GetDisplayName())
	}

	return ""
}

func GetChatMessageContentPreview(msg models.ChatMessageable) (string, int64) {
	var content string

	if msg.GetBody() != nil {
		content = ptr.Val(msg.GetBody().GetContent())
	}

	return str.Preview(content, 128), int64(len(content))
}

func GetChatMessageAttachmentNames(msg models.ChatMessageable) []string {
	names := make([]string, 0, len(msg.GetAttachments()))

	for _, a := range msg.GetAttachments() {
		if name := ptr.Val(a.GetName()); len(name) > 0 {
			names = append(names, name)
		}
	}

	return names
}
