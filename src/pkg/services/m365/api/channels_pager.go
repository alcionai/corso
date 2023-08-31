package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/teams"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
)

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

type ChannelMessageDeltaEnumerator interface {
	DeltaGetPager
	ValuesInPageLinker[models.ChatMessageable]
	SetNextLinker
}

var _ ChannelMessageDeltaEnumerator = &ChannelMessageDeltaPageCtrl{}

type ChannelMessageDeltaPageCtrl struct {
	gs      graph.Servicer
	builder *teams.ItemChannelsItemMessagesDeltaRequestBuilder
	options *teams.ItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration
}

func (c Channels) NewMessagePager(
	teamID, channelID string,
	fields []string,
) *ChannelMessageDeltaPageCtrl {
	res := &ChannelMessageDeltaPageCtrl{
		gs:      c.Stable,
		options: nil,
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

func (p *ChannelMessageDeltaPageCtrl) SetNext(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *ChannelMessageDeltaPageCtrl) GetPage(ctx context.Context) (DeltaPageLinker, error) {
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

func (p *ChannelMessageDeltaPageCtrl) ValuesIn(l PageLinker) ([]models.ChatMessageable, error) {
	return getValues[models.ChatMessageable](l)
}

// ---------------------------------------------------------------------------
// non delta channel message pager
// ---------------------------------------------------------------------------

type MessageItemIDType struct {
	ItemID string
}

type channelMessagePageCtrl struct {
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

	return &channelMessagePageCtrl{c.Stable, builder, options}
}

func (p *channelMessagePageCtrl) getPage(ctx context.Context) (PageLinkValuer[models.ChatMessageable], error) {
	page, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return EmptyDeltaLinker[models.ChatMessageable]{PageLinkValuer: page}, nil
}

//lint:ignore U1000 False Positive
func (p *channelMessagePageCtrl) setNext(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesRequestBuilder(nextLink, p.gs.Adapter())
}

// ---------------------------------------------------------------------------
// channel pager
// ---------------------------------------------------------------------------

type ChannelEnumerator interface {
	PageLinker
	ValuesInPageLinker[models.Channelable]
	SetNextLinker
}

var _ ChannelEnumerator = &channelPageCtrl{}

type channelPageCtrl struct {
	gs      graph.Servicer
	builder *teams.ItemChannelsRequestBuilder
	options *teams.ItemChannelsRequestBuilderGetRequestConfiguration
}

func (c Channels) NewChannelPager(
	teamID,
	channelID string,
	fields []string,
) *channelPageCtrl {
	res := &channelPageCtrl{
		gs:      c.Stable,
		options: nil,
		builder: c.Stable.
			Client().
			Teams().
			ByTeamId(teamID).
			Channels(),
	}

	return res
}

func (p *channelPageCtrl) SetNext(nextLink string) {
	p.builder = teams.NewItemChannelsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *channelPageCtrl) GetPage(ctx context.Context) (PageLinker, error) {
	var (
		resp PageLinker
		err  error
	)

	resp, err = p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

func (p *channelPageCtrl) ValuesIn(l PageLinker) ([]models.Channelable, error) {
	return getValues[models.Channelable](l)
}

func (p *channelPageCtrl) GetOdataNextLink() *string {
	// No next link preent in the API result
	emptyString := ""
	return &emptyString
}
