package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/teams"
)

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

type ChannelMessageDeltaEnumerator interface {
	DeltaGetPager
	ValuesInPageLinker[models.ChatMessageable]
	SetNextLinker
}

var _ ChannelMessageDeltaEnumerator = &MessagePageCtrl{}

type MessagePageCtrl struct {
	gs      graph.Servicer
	builder *teams.ItemChannelsItemMessagesDeltaRequestBuilder
	options *teams.ItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration
}

func (c Channels) NewMessagePager(
	teamID,
	channelID string,
	fields []string,
) *MessagePageCtrl {
	requestConfig := &teams.ItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: &teams.ItemChannelsItemMessagesDeltaRequestBuilderGetQueryParameters{
			Select: fields,
		},
	}

	res := &MessagePageCtrl{
		gs:      c.Stable,
		options: requestConfig,
		builder: c.Stable.
			Client().
			Teams().
			ByTeamId(teamID).
			Channels().
			ByChannelId(channelID).
			Messages().
			Delta(),
	}

	return res
}

func (p *MessagePageCtrl) SetNext(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *MessagePageCtrl) GetPage(ctx context.Context) (DeltaPageLinker, error) {
	var (
		resp DeltaPageLinker
		err  error
	)

	resp, err = p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

func (p *MessagePageCtrl) ValuesIn(l PageLinker) ([]models.ChatMessageable, error) {
	return getValues[models.ChatMessageable](l)
}

type MessageItemIDType struct {
	ItemID string
}

type ChannelDeltaEnumerator interface {
	DeltaGetPager
	ValuesInPageLinker[models.Channelable]
	SetNextLinker
}

// TODO: implement
// var _ ChannelDeltaEnumerator = &channelsPageCtrl{}
type channelItemPageCtrl struct {
	gs      graph.Servicer
	builder *teams.ItemChannelsItemMessagesRequestBuilder
	options *teams.ItemChannelsItemMessagesRequestBuilderGetRequestConfiguration
}

func (c Channels) GetItemIDsInContainer(
	ctx context.Context,
	teamID, channelID string,
) (map[string]MessageItemIDType, error) {
	ctx = clues.Add(ctx, "channel_id", channelID)
	pager := c.NewChannelItemPager(teamID, channelID)

	items, err := enumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating contacts")
	}

	m := map[string]MessageItemIDType{}

	for _, item := range items {
		m[ptr.Val(item.GetId())] = MessageItemIDType{
			ItemID: ptr.Val(item.GetId()),
		}
	}

	return m, nil
}

func (c Channels) NewChannelItemPager(
	teamID, containerID string,
	selectProps ...string,
) itemPager[models.ChatMessageable] {
	options := &teams.ItemChannelsItemMessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &teams.ItemChannelsItemMessagesRequestBuilderGetQueryParameters{},
	}

	if len(selectProps) > 0 {
		options.QueryParameters.Select = selectProps
	}

	builder := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(containerID).
		Messages()

	return &channelItemPageCtrl{c.Stable, builder, options}
}

//lint:ignore U1000 False Positive
func (p *channelItemPageCtrl) getPage(ctx context.Context) (PageLinkValuer[models.ChatMessageable], error) {
	page, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return EmptyDeltaLinker[models.ChatMessageable]{PageLinkValuer: page}, nil
}

//lint:ignore U1000 False Positive
func (p *channelItemPageCtrl) setNext(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesRequestBuilder(nextLink, p.gs.Adapter())
}
