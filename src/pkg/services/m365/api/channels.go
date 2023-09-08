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
		ByTeamIdString(teamID).
		Channels().
		ByChannelIdString(containerID).
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
		ByTeamIdString(teamID).
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
	var size int64

	message, err := c.Stable.
		Client().
		Teams().
		ByTeamIdString(teamID).
		Channels().
		ByChannelIdString(channelID).
		Messages().
		ByChatMessageIdString(messageID).
		Get(ctx, nil)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	info := ChannelMessageInfo(message, size)

	return message, info, nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func ChannelMessageInfo(
	msg models.ChatMessageable,
	size int64,
) *details.GroupsInfo {
	var (
		lastReply  time.Time
		modTime    = ptr.OrNow(msg.GetLastModifiedDateTime())
		msgCreator string
	)

	for _, r := range msg.GetReplies() {
		cdt := ptr.Val(r.GetCreatedDateTime())
		if cdt.After(lastReply) {
			lastReply = cdt
		}
	}

	// if the message hasn't been modified since before the most recent
	// reply, set the modified time to the most recent reply.  This ensures
	// we update the message contents to match changes in replies.
	if modTime.Before(lastReply) {
		modTime = lastReply
	}

	from := msg.GetFrom()

	switch true {
	case from == nil:
		// not all messages have a populated 'from'.  Namely, system messages do not.
	case from.GetApplication() != nil:
		msgCreator = ptr.Val(from.GetApplication().GetDisplayName())
	case from.GetDevice() != nil:
		msgCreator = ptr.Val(from.GetDevice().GetDisplayName())
	case from.GetUser() != nil:
		msgCreator = ptr.Val(from.GetUser().GetDisplayName())
	}

	return &details.GroupsInfo{
		ItemType:       details.GroupsChannelMessage,
		Created:        ptr.Val(msg.GetCreatedDateTime()),
		LastReplyAt:    lastReply,
		Modified:       modTime,
		MessageCreator: msgCreator,
		MessagePreview: str.Preview(ptr.Val(msg.GetBody().GetContent()), 16),
		ReplyCount:     len(msg.GetReplies()),
		Size:           size,
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
