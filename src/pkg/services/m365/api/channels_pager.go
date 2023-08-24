package api

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

type ChannelMessageDeltaEnumerator interface {
	DeltaGetPager
	ValuesInPageLinker[models.ChatMessageable]
	SetNextLinker
}

// TODO: implement
// var _ ChannelMessageDeltaEnumerator = &messagePageCtrl{}

// type messagePageCtrl struct {
// 	gs      graph.Servicer
// 	builder *teams.ItemChannelsItemMessagesRequestBuilder
// 	options *teams.ItemChannelsItemMessagesRequestBuilderGetRequestConfiguration
// }

// ---------------------------------------------------------------------------
// channel pager
// ---------------------------------------------------------------------------

type ChannelDeltaEnumerator interface {
	DeltaGetPager
	ValuesInPageLinker[models.Channelable]
	SetNextLinker
}

// TODO: implement
// var _ ChannelDeltaEnumerator = &channelsPageCtrl{}

// type channelsPageCtrl struct {
// 	gs      graph.Servicer
// 	builder *teams.ItemChannelsChannelItemRequestBuilder
// 	options *teams.ItemChannelsChannelItemRequestBuilderGetRequestConfiguration
// }
