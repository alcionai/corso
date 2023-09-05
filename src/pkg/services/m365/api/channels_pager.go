package api

import (
	"context"

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
		ByTeamIdString(p.resourceID).
		Channels().
		ByChannelIdString(p.channelID).
		Messages().
		Delta()
}

func (p *channelMessageDeltaPageCtrl) ValuesIn(l PageLinker) ([]models.ChatMessageable, error) {
	return getValues[models.ChatMessageable](l)
}

func (c Channels) NewChannelMessageDeltaPager(
	teamID, channelID, prevDelta string,
	selectProps ...string,
) *channelMessageDeltaPageCtrl {
	builder := c.Stable.
		Client().
		Teams().
		ByTeamIdString(teamID).
		Channels().
		ByChannelIdString(channelID).
		Messages().
		Delta()

	if len(prevDelta) > 0 {
		builder = teams.NewItemChannelsItemMessagesDeltaRequestBuilder(prevDelta, c.Stable.Adapter())
	}

	options := &teams.ItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: &teams.ItemChannelsItemMessagesDeltaRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(selectProps) > 0 {
		options.QueryParameters.Select = selectProps
	}

	return &channelMessageDeltaPageCtrl{
		resourceID: teamID,
		channelID:  channelID,
		builder:    builder,
		gs:         c.Stable,
		options:    options,
	}
}

// GetChannelMessageIDsDelta fetches a delta of all messages in the channel.
func (c Channels) GetChannelMessageIDsDelta(
	ctx context.Context,
	teamID, channelID, prevDelta string,
) (map[string]struct{}, DeltaUpdate, error) {
	var (
		vs = map[string]struct{}{}
		// select is not currently allowed on messages
		// this func will still isolate to the ID, however,
		// because we need the follow-up get request to gather
		// all replies to the message.
		// selectProps      = idAnd()
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
			vs = map[string]struct{}{}

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

		for _, v := range vals {
			vs[ptr.Val(v.GetId())] = struct{}{}
		}

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
// channel message replies pager
// ---------------------------------------------------------------------------

var _ Pager[models.ChatMessageable] = &channelMessageRepliesPageCtrl{}

type channelMessageRepliesPageCtrl struct {
	gs      graph.Servicer
	builder *teams.ItemChannelsItemMessagesItemRepliesRequestBuilder
	options *teams.ItemChannelsItemMessagesItemRepliesRequestBuilderGetRequestConfiguration
}

func (p *channelMessageRepliesPageCtrl) SetNext(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesItemRepliesRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *channelMessageRepliesPageCtrl) GetPage(
	ctx context.Context,
) (PageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *channelMessageRepliesPageCtrl) GetOdataNextLink() *string {
	return ptr.To("")
}

func (p *channelMessageRepliesPageCtrl) ValuesIn(l PageLinker) ([]models.ChatMessageable, error) {
	return getValues[models.ChatMessageable](l)
}

func (c Channels) NewChannelMessageRepliesPager(
	teamID, channelID, messageID string,
	selectProps ...string,
) *channelMessageRepliesPageCtrl {
	options := &teams.ItemChannelsItemMessagesItemRepliesRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(selectProps) > 0 {
		options.QueryParameters.Select = selectProps
	}

	res := &channelMessageRepliesPageCtrl{
		gs:      c.Stable,
		options: options,
		builder: c.Stable.
			Client().
			Teams().
			ByTeamIdString(teamID).
			Channels().
			ByChannelIdString(channelID).
			Messages().
			ByChatMessageIdString(messageID).
			Replies(),
	}

	return res
}

// GetChannels fetches the minimum valuable data from each reply in the message
func (c Channels) GetChannelMessageReplies(
	ctx context.Context,
	teamID, channelID, messageID string,
) ([]models.ChatMessageable, error) {
	var (
		vs = []models.ChatMessageable{}
		// select is not currently enabled for replies.
		// selectProps = idAnd(
		// 	"messageType",
		// 	"createdDateTime",
		// 	"from",
		// 	"body")
		pager = c.NewChannelMessageRepliesPager(teamID, channelID, messageID)
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

	logger.Ctx(ctx).Debugf("retrieved %d channel message replies", len(vs))

	return vs, nil
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
			ByTeamIdString(teamID).
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
