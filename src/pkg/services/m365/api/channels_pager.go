package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/teams"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/logger"
)

// ---------------------------------------------------------------------------
// channel message pager
// ---------------------------------------------------------------------------

var _ DeltaPager[models.ChatMessageable] = &channelMessageDeltaPageCtrl{}

type channelMessageDeltaPageCtrl struct {
	resourceID, channelID string
	gs                    graph.Servicer
	builder               *teams.ItemChannelsItemMessagesDeltaRequestBuilder
	options               *teams.ItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration
}

func (p *channelMessageDeltaPageCtrl) SetNext(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *channelMessageDeltaPageCtrl) GetPage(
	ctx context.Context,
) (DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *channelMessageDeltaPageCtrl) Reset(context.Context) {
	p.builder = p.gs.
		Client().
		Teams().
		ByTeamId(p.resourceID).
		Channels().
		ByChannelId(p.channelID).
		Messages().
		Delta()
}

func (p *channelMessageDeltaPageCtrl) ValuesIn(l PageLinker) ([]models.ChatMessageable, error) {
	return getValues[models.ChatMessageable](l)
}

type ChannelDeltaEnumerator interface {
	PageLinker
	ValuesInPageLinker[models.Channelable]
	SetNextLinker
	GetPage(ctx context.Context) (PageLinker, error)
}

func (c Channels) NewChannelMessageDeltaPager(
	teamID, channelID, prevDelta string,
) *channelMessageDeltaPageCtrl {
	builder := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(channelID).
		Messages().
		Delta()

	if len(prevDelta) > 0 {
		builder = teams.NewItemChannelsItemMessagesDeltaRequestBuilder(prevDelta, c.Stable.Adapter())
	}

	options := &teams.ItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	return &channelMessageDeltaPageCtrl{
		resourceID: teamID,
		channelID:  channelID,
		builder:    builder,
		gs:         c.Stable,
		options:    options,
	}
}

// var _ ChannelMessageDeltaEnumerator = &ChannelMessageDeltaPageCtrl{}
var _ Pager[models.Channelable] = &channelPageCtrl{}

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
) (map[string]struct{}, error) {
	ctx = clues.Add(ctx, "channel_id", channelID)
	pager := c.NewChannelItemPager(teamID, channelID)

	items, err := enumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating contacts")
	}

	m := map[string]struct{}{}

	for _, item := range items {
		m[ptr.Val(item.GetId())] = struct{}{}
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

//lint:ignore U1000 False Positive
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

// GetChannelMessagesDelta fetches a delta of all messages in the channel.
func (c Channels) GetChannelMessagesDelta(
	ctx context.Context,
	teamID, channelID, prevDelta string,
) ([]models.ChatMessageable, DeltaUpdate, error) {
	var (
		vs               = []models.ChatMessageable{}
		pager            = c.NewChannelMessageDeltaPager(teamID, channelID, prevDelta)
		invalidPrevDelta = len(prevDelta) == 0
		newDeltaLink     string
	)

	// Loop through all pages returned by Graph API.
	for {
		page, err := pager.GetPage(graph.ConsumeNTokens(ctx, graph.SingleGetOrDeltaLC))
		if graph.IsErrInvalidDelta(err) {
			logger.Ctx(ctx).Infow("Invalid previous delta", "delta_link", prevDelta)

			invalidPrevDelta = true
			vs = []models.ChatMessageable{}

			pager.Reset(ctx)

			continue
		}

		if err != nil {
			return nil, DeltaUpdate{}, graph.Wrap(ctx, err, "retrieving page of channel messages")
		}

		vals, err := pager.ValuesIn(page)
		if err != nil {
			return nil, DeltaUpdate{}, graph.Wrap(ctx, err, "extracting channel messages from response")
		}

		vs = append(vs, vals...)

		nextLink, deltaLink := NextAndDeltaLink(page)

		if len(deltaLink) > 0 {
			newDeltaLink = deltaLink
		}

		if len(nextLink) == 0 {
			break
		}

		pager.SetNext(nextLink)
	}

	logger.Ctx(ctx).Debugf("retrieved %d channel messages", len(vs))

	du := DeltaUpdate{
		URL:   newDeltaLink,
		Reset: invalidPrevDelta,
	}

	return vs, du, nil
}

// ---------------------------------------------------------------------------
// channel pager
// ---------------------------------------------------------------------------

var _ Pager[models.Channelable] = &channelPageCtrl{}

type channelPageCtrl struct {
	gs      graph.Servicer
	builder *teams.ItemChannelsRequestBuilder
	options *teams.ItemChannelsRequestBuilderGetRequestConfiguration
}

func (p *channelPageCtrl) SetNext(nextLink string) {
	p.builder = teams.NewItemChannelsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *channelPageCtrl) GetPage(
	ctx context.Context,
) (PageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *channelPageCtrl) ValuesIn(l PageLinker) ([]models.Channelable, error) {
	return getValues[models.Channelable](l)
}

func (c Channels) NewChannelPager(
	teamID string,
) *channelPageCtrl {
	requestConfig := &teams.ItemChannelsRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	res := &channelPageCtrl{
		gs:      c.Stable,
		options: requestConfig,
		builder: c.Stable.
			Client().
			Teams().
			ByTeamId(teamID).
			Channels(),
	}

	return res
}

// GetChannels fetches all channels in the team.
func (c Channels) GetChannels(
	ctx context.Context,
	teamID string,
) ([]models.Channelable, error) {
	var (
		vs    = []models.Channelable{}
		pager = c.NewChannelPager(teamID)
	)

	// Loop through all pages returned by Graph API.
	for {
		page, err := pager.GetPage(ctx)
		if err != nil {
			return nil, graph.Wrap(ctx, err, "retrieving page of channels")
		}

		vals, err := pager.ValuesIn(page)
		if err != nil {
			return nil, graph.Wrap(ctx, err, "extracting channels from response")
		}

		vs = append(vs, vals...)

		nextLink := ptr.Val(page.GetOdataNextLink())
		if len(nextLink) == 0 {
			break
		}

		pager.SetNext(nextLink)
	}

	logger.Ctx(ctx).Debugf("retrieved %d channels", len(vs))

	return vs, nil
}
