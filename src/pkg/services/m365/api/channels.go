package api

import (
	"context"
	"fmt"

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

// GetChannelByName fetches a channel by name
func (c Channels) GetChannelByName(
	ctx context.Context,
	teamID, containerName string,
) (graph.Container, error) {
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
	container := ChannelsDisplayable{Channelable: cal}

	if err := graph.CheckIDAndName(container); err != nil {
		return nil, clues.Stack(err).WithClues(ctx)
	}

	return container, nil
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

	return message, ChannelMessageInfo(message, size), nil
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

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func ChannelMessageInfo(msg models.ChatMessageable, size int64) *details.GroupsInfo {
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
