package api

import (
	"context"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/microsoftgraph/msgraph-sdk-go/teams"
)

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

type MessageItemDeltaEnumerator interface {
	GetPage(context.Context) (PageLinker, error)
	SetNext(nextLink string)
}

var _ MessageItemDeltaEnumerator = &messagePageCtrl{}

type messagePageCtrl struct {
	gs      graph.Servicer
	builder *teams.ItemChannelsItemMessagesDeltaRequestBuilder
	options *teams.ItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration
}

func (c Channels) NewMessagePager(
	teamID,
	channelID string,
	fields []string,
) *messagePageCtrl {
	requestConfig := &teams.ItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		QueryParameters: &teams.ItemChannelsItemMessagesDeltaRequestBuilderGetQueryParameters{
			Select: fields,
		},
	}

	res := &messagePageCtrl{
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

func (p *messagePageCtrl) SetNext(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *messagePageCtrl) GetPage(ctx context.Context) (PageLinker, error) {
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
