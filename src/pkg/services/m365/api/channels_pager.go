package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/teams"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

// ---------------------------------------------------------------------------
// channel message pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.ChatMessageable] = &channelMessagePageCtrl{}

type channelMessagePageCtrl struct {
	resourceID, channelID string
	gs                    graph.Servicer
	builder               *teams.ItemChannelsItemMessagesRequestBuilder
	options               *teams.ItemChannelsItemMessagesRequestBuilderGetRequestConfiguration
}

func (p *channelMessagePageCtrl) SetNextLink(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *channelMessagePageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.ChatMessageable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *channelMessagePageCtrl) ValidModTimes() bool {
	return true
}

func (c Channels) NewChannelMessagePager(
	teamID, channelID string,
	cc CallConfig,
) *channelMessagePageCtrl {
	builder := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(channelID).
		Messages()

	options := &teams.ItemChannelsItemMessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: &teams.ItemChannelsItemMessagesRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
	}

	if len(cc.Select) > 0 {
		options.QueryParameters.Select = cc.Select
	}

	if len(cc.Expand) > 0 {
		options.QueryParameters.Expand = cc.Expand
	}

	return &channelMessagePageCtrl{
		resourceID: teamID,
		channelID:  channelID,
		builder:    builder,
		gs:         c.Stable,
		options:    options,
	}
}

// GetChannelMessages fetches a delta of all messages in the channel.
func (c Channels) GetChannelMessages(
	ctx context.Context,
	teamID, channelID string,
	cc CallConfig,
) ([]models.ChatMessageable, error) {
	ctx = clues.Add(ctx, "channel_id", channelID)
	pager := c.NewChannelMessagePager(teamID, channelID, cc)
	items, err := pagers.BatchEnumerateItems[models.ChatMessageable](ctx, pager)

	return items, graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// channel message delta pager
// ---------------------------------------------------------------------------

var _ pagers.DeltaHandler[models.ChatMessageable] = &channelMessageDeltaPageCtrl{}

type channelMessageDeltaPageCtrl struct {
	resourceID, channelID string
	gs                    graph.Servicer
	builder               *teams.ItemChannelsItemMessagesDeltaRequestBuilder
	options               *teams.ItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration
}

func (p *channelMessageDeltaPageCtrl) SetNextLink(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *channelMessageDeltaPageCtrl) GetPage(
	ctx context.Context,
) (pagers.DeltaLinkValuer[models.ChatMessageable], error) {
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

func (p *channelMessageDeltaPageCtrl) ValidModTimes() bool {
	return true
}

func (c Channels) NewChannelMessageDeltaPager(
	teamID, channelID, prevDelta string,
	selectProps ...string,
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
		QueryParameters: &teams.ItemChannelsItemMessagesDeltaRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(maxDeltaPageSize)),
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

// this is the message content for system chatMessage entities with type
// unknownFutureValue.
const channelMessageSystemMessageContent = "<systemEventMessage/>"

func filterOutSystemMessages(cm models.ChatMessageable) bool {
	if ptr.Val(cm.GetMessageType()) == models.SYSTEMEVENTMESSAGE_CHATMESSAGETYPE {
		return false
	}

	content := ""

	if cm.GetBody() != nil {
		content = ptr.Val(cm.GetBody().GetContent())
	}

	return !(ptr.Val(cm.GetMessageType()) == models.UNKNOWNFUTUREVALUE_CHATMESSAGETYPE &&
		content == channelMessageSystemMessageContent)
}

// GetChannelMessageIDs fetches a delta of all messages in the channel.
// returns two maps: addedItems, deletedItems
func (c Channels) GetChannelMessageIDs(
	ctx context.Context,
	teamID, channelID, prevDeltaLink string,
	cc CallConfig,
) (pagers.AddedAndRemoved, error) {
	aar, err := pagers.GetAddedAndRemovedItemIDs[models.ChatMessageable](
		ctx,
		c.NewChannelMessagePager(teamID, channelID, CallConfig{}),
		c.NewChannelMessageDeltaPager(teamID, channelID, prevDeltaLink),
		prevDeltaLink,
		cc.CanMakeDeltaQueries,
		pagers.AddedAndRemovedByDeletedDateTime[models.ChatMessageable],
		filterOutSystemMessages)

	return aar, clues.Stack(err).OrNil()
}

// ---------------------------------------------------------------------------
// channel message replies pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.ChatMessageable] = &channelMessageRepliesPageCtrl{}

type channelMessageRepliesPageCtrl struct {
	gs      graph.Servicer
	builder *teams.ItemChannelsItemMessagesItemRepliesRequestBuilder
	options *teams.ItemChannelsItemMessagesItemRepliesRequestBuilderGetRequestConfiguration
}

func (p *channelMessageRepliesPageCtrl) SetNextLink(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesItemRepliesRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *channelMessageRepliesPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.ChatMessageable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *channelMessageRepliesPageCtrl) GetOdataNextLink() *string {
	return ptr.To("")
}

func (p *channelMessageRepliesPageCtrl) ValidModTimes() bool {
	return true
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
			ByTeamId(teamID).
			Channels().
			ByChannelId(channelID).
			Messages().
			ByChatMessageId(messageID).
			Replies(),
	}

	return res
}

// GetChannels fetches the minimum valuable data from each reply in the message
func (c Channels) GetChannelMessageReplies(
	ctx context.Context,
	teamID, channelID, messageID string,
) ([]models.ChatMessageable, error) {
	return pagers.BatchEnumerateItems[models.ChatMessageable](
		ctx,
		c.NewChannelMessageRepliesPager(teamID, channelID, messageID))
}

// ---------------------------------------------------------------------------
// channel pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.Channelable] = &channelPageCtrl{}

type channelPageCtrl struct {
	gs      graph.Servicer
	builder *teams.ItemChannelsRequestBuilder
	options *teams.ItemChannelsRequestBuilderGetRequestConfiguration
}

func (p *channelPageCtrl) SetNextLink(nextLink string) {
	p.builder = teams.NewItemChannelsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *channelPageCtrl) GetPage(
	ctx context.Context,
) (pagers.NextLinkValuer[models.Channelable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *channelPageCtrl) ValidModTimes() bool {
	return false
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
	return pagers.BatchEnumerateItems[models.Channelable](ctx, c.NewChannelPager(teamID))
}
