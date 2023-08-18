package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/teams"
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

// CreateContainer makes an channels with the name in the team
func (c Channels) CreateChannel(
	ctx context.Context,
	teamID, _, containerName string,
) (graph.Container, error) {
	body := models.NewChannel()
	body.SetDisplayName(&containerName)

	container, err := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating channel")
	}

	return ChannelsDisplayable{Channelable: container}, nil
}

// DeleteChannel removes a channel from user's M365 account
func (c Channels) DeleteChannel(
	ctx context.Context,
	teamID, containerID string,
) error {
	// deletes require unique http clients
	// https://github.com/alcionai/corso/issues/2707
	srv, err := NewService(c.Credentials)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	err = srv.Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(containerID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Stack(ctx, err)
	}

	return nil
}

// prefer GetChannelByID where possible.
// use this only in cases where the models.Channelable
// is required.
func (c Channels) GetChannel(
	ctx context.Context,
	teamID, containerID string,
) (models.Channelable, error) {
	config := &teams.ItemChannelsChannelItemRequestBuilderGetRequestConfiguration{
		QueryParameters: &teams.ItemChannelsChannelItemRequestBuilderGetQueryParameters{
			Select: idAnd("name", "owner"),
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

func (c Channels) GetChannelByID(
	ctx context.Context,
	teamID, containerID string,
) (graph.Container, error) {
	channel, err := c.GetChannel(ctx, teamID, containerID)
	if err != nil {
		return nil, err
	}

	return ChannelsDisplayable{Channelable: channel}, nil
}

// GetChannelByName fetches a calendar by name
func (c Channels) GetChannelByName(
	ctx context.Context,
	teamID, _, containerName string,
) (graph.Container, error) {

	ctx = clues.Add(ctx, "channel_name", containerName)

	resp, err := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		Get(ctx, nil)

	if err != nil {
		return nil, graph.Stack(ctx, err).WithClues(ctx)
	}

	gv := resp.GetValue()

	if len(gv) == 0 {
		return nil, clues.New("channel not found").WithClues(ctx)
	}

	// We only allow the api to match one calendar with the provided name.
	// If we match multiples, we'll eagerly return the first one.
	logger.Ctx(ctx).Debugw("calendars matched the name search", "calendar_count", len(gv))

	// Sanity check ID and name
	cal := gv[0]
	container := ChannelsDisplayable{Channelable: cal}

	if err := graph.CheckIDAndName(container); err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return container, nil
}

func (c Channels) PatchChannel(
	ctx context.Context,
	teamID, containerID string,
	body models.Channelable,
) error {
	_, err := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(containerID).
		Patch(ctx, body, nil)

	if err != nil {
		return graph.Wrap(ctx, err, "patching event calendar")
	}

	return nil
}

// ---------------------------------------------------------------------------
// message
// ---------------------------------------------------------------------------

// GetItem retrieves a Messageable item.
func (c Channels) GetMessage(
	ctx context.Context,
	teamID, channelID, itemID string,
	immutableIDs bool,
	errs *fault.Bus,
) (serialization.Parsable, *details.GroupsInfo, error) {
	var (
		size int64
	)

	// is preferImmutableIDs headers required here
	message, err := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(channelID).
		Messages().
		ByChatMessageId(itemID).
		Get(ctx, nil)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	return message, MessageInfo(message, size), nil
}

func (c Channels) PostMessage(
	ctx context.Context,
	teamID, containerID string,
	body models.ChatMessageable,
) (models.ChatMessageable, error) {
	itm, err := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(containerID).
		Messages().
		Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating mail message")
	}

	if itm == nil {
		return nil, clues.New("nil response mail message creation").WithClues(ctx)
	}

	return itm, nil
}

func (c Channels) DeleteMessage(
	ctx context.Context,
	teamID, itemID, containerID string,
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
		ByTeamId(teamID).
		Channels().
		ByChannelId(containerID).
		Messages().
		ByChatMessageId(itemID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting mail message")
	}

	return nil
}

// ---------------------------------------------------------------------------
// replies
// ---------------------------------------------------------------------------

// GetReplies retrieves a Messageable item.
func (c Channels) GetReplies(
	ctx context.Context,
	teamID, channelID, itemID string,
) (serialization.Parsable, error) {
	replies, err := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(channelID).
		Messages().
		ByChatMessageId(itemID).
		Replies().
		Get(ctx, nil)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return replies, nil
}

func (c Channels) PostReply(
	ctx context.Context,
	teamID, containerID, messageID string,
	body models.ChatMessageable,
) (models.ChatMessageable, error) {
	itm, err := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(containerID).
		Messages().
		ByChatMessageId(messageID).
		Replies().
		Post(ctx, body, nil)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "creating reply message")
	}

	if itm == nil {
		return nil, clues.New("nil response reply to message creation").WithClues(ctx)
	}

	return itm, nil
}

func (c Channels) DeleteReply(
	ctx context.Context,
	teamID, itemID, containerID, replyID string,
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
		ByTeamId(teamID).
		Channels().
		ByChannelId(containerID).
		Messages().
		ByChatMessageId(itemID).
		Replies().
		ByChatMessageId1(replyID).
		Delete(ctx, nil)
	if err != nil {
		return graph.Wrap(ctx, err, "deleting mail message")
	}

	return nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func MessageInfo(msg models.ChatMessageable, size int64) *details.GroupsInfo {
	var (
		created = ptr.Val(msg.GetCreatedDateTime())
	)

	return &details.GroupsInfo{
		ItemType: details.GroupChannel,
		Size:     size,
		Created:  created,
		Modified: ptr.OrNow(msg.GetLastModifiedDateTime()),
	}
}

// ---------------------------------------------------------------------------
// helper funcs
// ---------------------------------------------------------------------------

// ChannelsDisplayable is a wrapper that complies with the
// models.Channelable interface with the graph.Container
// interfaces.
type ChannelsDisplayable struct {
	models.Channelable
}

// GetParentFolderId returns the default channe name address

//nolint:revive
func (c ChannelsDisplayable) GetParentFolderId() *string {
	return nil
}
