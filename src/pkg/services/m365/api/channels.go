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

//  ---------------------------------------------------------------------------
// Currently implemented-
// - Channels CRUD
// - Items i.e. messages CRUD
// Pending
// - teams CRUD

// ---------------------------------------------------------------------------

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
func (c Channels) CreateContainer(
	ctx context.Context,
	// parentContainerID needed for iface, doesn't apply to events
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

// DeleteContainer removes a channel from user's M365 account
func (c Channels) DeleteContainer(
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

// prefer GetContainerByID where possible.
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

// interface-compliant wrapper of GetCalendar
func (c Channels) GetContainerByID(
	ctx context.Context,
	teamID, containerID string,
) (graph.Container, error) {
	channel, err := c.GetChannel(ctx, teamID, containerID)
	if err != nil {
		return nil, err
	}

	return ChannelsDisplayable{Channelable: channel}, nil
}

// GetContainerByName fetches a calendar by name
func (c Channels) GetContainerByName(
	ctx context.Context,
	// parentContainerID needed for iface, doesn't apply to events
	teamID, _, containerName string,
) (graph.Container, error) {

	// TODO: check container filter
	// filter := fmt.Sprintf("name eq '%s'", containerName)
	// options := &teams.ItemChannelsChannelItemRequestBuilderGetRequestConfiguration{
	// 	QueryParameters: &teams.ItemChannelsChannelItemRequestBuilderGetQueryParameters{
	// 		Filter: &filter,
	// 	},
	// }

	ctx = clues.Add(ctx, "container_name", containerName)

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
		return nil, clues.New("container not found").WithClues(ctx)
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
// items
// ---------------------------------------------------------------------------

// GetItem retrieves a Messageable item.  If the item contains an attachment, that
// attachment is also downloaded.
func (c Channels) GetItem(
	ctx context.Context,
	teamID, channelID, itemID string,
	immutableIDs bool,
	errs *fault.Bus,
) (serialization.Parsable, *details.ChannelsInfo, error) {
	var (
		size int64
	)

	// is preferImmutableIDs headers required here
	message, err := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(itemID).
		Messages().
		ByChatMessageId("").
		Get(ctx, nil)
	if err != nil {
		return nil, nil, graph.Stack(ctx, err)
	}

	return message, ChannelsInfo(message, size), nil
}

func (c Channels) PostItem(
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

func (c Channels) DeleteItem(
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
// Helpers
// ---------------------------------------------------------------------------

func ChannelsInfo(msg models.ChatMessageable, size int64) *details.ChannelsInfo {
	var (
		sender  = ptr.Val(msg.GetFrom().GetUser().GetDisplayName())
		created = ptr.Val(msg.GetCreatedDateTime())
	)

	return &details.ChannelsInfo{
		ItemType: details.ExchangeMail,
		Sender:   sender,
		Size:     size,
		Created:  created,
		Modified: ptr.OrNow(msg.GetLastModifiedDateTime()),
	}
}
