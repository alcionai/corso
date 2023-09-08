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

// CreateChannel makes an channels with the name in the team
func (c Channels) CreateChannel(
	ctx context.Context,
	teamID, channelName string,
) (models.Channelable, error) {
	body := models.NewChannel()
	body.SetDisplayName(&channelName)

	channel, err := c.Stable.
		Client().
		Teams().
		ByTeamIdString(teamID).
		Channels().
		Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating channel")
	}

	return channel, nil
}

// DeleteChannel removes a channel from team with given teamID
func (c Channels) DeleteChannel(
	ctx context.Context,
	teamID, channelID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := NewService(c.Credentials)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.Client().
		Teams().
		ByTeamIdString(teamID).
		Channels().
		ByChannelIdString(channelID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	return nil
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

func (c Channels) PostChannelMessage(
	ctx context.Context,
	teamID, channelID string,
	body *models.ItemBody,
) (models.ChatMessageable, error) {
	requestBody := models.NewChatMessage()
	requestBody.SetBody(body)

	itm, err := c.Stable.
		Client().
		Teams().
		ByTeamIdString(teamID).
		Channels().
		ByChannelIdString(channelID).
		Messages().
		Post(ctx, requestBody, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating chamailnnel message")
	}

	if itm == nil {
		return nil, clues.New("nil response channel message creation").WithClues(ctx)
	}

	return itm, nil
}

func (c Channels) DeleteChannelMessage(
	ctx context.Context,
	teamID, messageID, channelID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := NewService(c.Credentials)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.
		Client().
		Teams().
		ByTeamIdString(teamID).
		Channels().
		ByChannelIdString(channelID).
		Messages().
		ByChatMessageIdString(messageID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting channel message")
	}

	return nil
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
