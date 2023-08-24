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

// GetMessage retrieves a ChannelMessage item.
func (c Channels) GetMessage(
	ctx context.Context,
	teamID, channelID, itemID string,
	errs *fault.Bus,
) (serialization.Parsable, *details.GroupsInfo, error) {
	var (
		size int64
	)

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
		ItemType: details.TeamsChannelMessage,
		Size:     size,
		Created:  created,
		Modified: ptr.OrNow(msg.GetLastModifiedDateTime()),
	}
}

// ---------------------------------------------------------------------------
// helper funcs
// ---------------------------------------------------------------------------

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
