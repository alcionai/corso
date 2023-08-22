package api

import (
	"context"
)

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

type MessageItemDeltaEnumerator interface {
	GetPage(context.Context) (DeltaPageLinker, error)
}

// TODO: implement
// var _ MessageItemDeltaEnumerator = &messagePageCtrl{}

// type messagePageCtrl struct {
// 	gs      graph.Servicer
// 	builder *teams.ItemChannelsItemMessagesRequestBuilder
// 	options *teams.ItemChannelsItemMessagesRequestBuilderGetRequestConfiguration
// }

// ---------------------------------------------------------------------------
// channel pager
// ---------------------------------------------------------------------------

type ChannelItemDeltaEnumerator interface {
	GetPage(context.Context) (DeltaPageLinker, error)
}

// TODO: implement
// var _ ChannelsItemDeltaEnumerator = &channelsPageCtrl{}

// type channelsPageCtrl struct {
// 	gs      graph.Servicer
// 	builder *teams.ItemChannelsChannelItemRequestBuilder
// 	options *teams.ItemChannelsChannelItemRequestBuilderGetRequestConfiguration
// }
