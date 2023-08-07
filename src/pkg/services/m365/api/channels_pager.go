package api

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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
// container pager
// ---------------------------------------------------------------------------

// EnumerateContainers iterates through all of the users teams
// channels, converting each to a graph.CacheFolder, and
// calling fn(cf) on each one.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Channels) EnumerateContainers(
	ctx context.Context,
	userID, baseContainerID string,
	fn func(graph.CachedContainer) error,
	errs *fault.Bus,
) error {
	var (
		el     = errs.Local()
		config = &teams.ItemChannelsRequestBuilderGetRequestConfiguration{
			QueryParameters: &teams.ItemChannelsRequestBuilderGetQueryParameters{
				Select: idAnd("name"),
			},
		}
		builder = c.Stable.
			Client().
			Teams().
			ByTeamId(userID).
			Channels()
	)

	for {
		if el.Failure() != nil {
			break
		}

		resp, err := builder.Get(ctx, config)
		if err != nil {
			return graph.Stack(ctx, err)
		}

		for _, channel := range resp.GetValue() {
			if el.Failure() != nil {
				break
			}

			cd := ChannelsDisplayable{Channelable: channel}
			if err := graph.CheckIDAndName(cd); err != nil {
				errs.AddRecoverable(ctx, graph.Stack(ctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}

			fctx := clues.Add(
				ctx,
				"container_id", ptr.Val(channel.GetId()),
				"container_name", ptr.Val(channel.GetDisplayName()))

			temp := graph.NewCacheFolder(
				// TODO: understand about it. Is path required and so on.
				cd,
				path.Builder{}.Append(ptr.Val(channel.GetId())),          // storage path
				path.Builder{}.Append(ptr.Val(channel.GetDisplayName()))) // display location
			if err := fn(&temp); err != nil {
				errs.AddRecoverable(ctx, graph.Stack(fctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = teams.NewItemChannelsRequestBuilder(link, c.Stable.Adapter())
	}

	return el.Failure()
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ itemPager[models.ChatMessageable] = &channelsPageCtrl{}

type channelsPageCtrl struct {
	gs      graph.Servicer
	builder *teams.ItemChannelsItemMessagesRequestBuilder
	options *teams.ItemChannelsItemMessagesRequestBuilderGetRequestConfiguration
}

func (c Channels) NewChannelsPager(
	teamID, containerID string,
	selectProps ...string,
) itemPager[models.ChatMessageable] {
	options := &teams.ItemChannelsItemMessagesRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
		QueryParameters: &teams.ItemChannelsItemMessagesRequestBuilderGetQueryParameters{
			// do NOT set Top.  It limits the total items received.
		},
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

	return &channelsPageCtrl{c.Stable, builder, options}
}

//lint:ignore U1000 False Positive
func (p *channelsPageCtrl) getPage(ctx context.Context) (PageLinkValuer[models.ChatMessageable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

//lint:ignore U1000 False Positive
func (p *channelsPageCtrl) setNext(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesRequestBuilder(nextLink, p.gs.Adapter())
}

// ---------------------------------------------------------------------------
// item ID pager
// ---------------------------------------------------------------------------

var _ itemIDPager = &channelIDPager{}

type channelIDPager struct {
	gs      graph.Servicer
	builder *teams.ItemChannelsItemMessagesRequestBuilder
	options *teams.ItemChannelsItemMessagesRequestBuilderGetRequestConfiguration
}

func (c Channels) NewChannelIDsPager(
	ctx context.Context,
	teamID, containerID string,
	immutableIDs bool,
) (itemIDPager, error) {
	options := &teams.ItemChannelsItemMessagesRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
		QueryParameters: &teams.ItemChannelsItemMessagesRequestBuilderGetQueryParameters{
			// do NOT set Top.  It limits the total items received.
		},
	}

	builder := c.Stable.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(containerID).
		Messages()

	return &channelIDPager{c.Stable, builder, options}, nil
}

func (p *channelIDPager) getPage(ctx context.Context) (DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return EmptyDeltaLinker[models.ChatMessageable]{PageLinkValuer: resp}, nil
}

func (p *channelIDPager) setNext(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesRequestBuilder(nextLink, p.gs.Adapter())
}

// non delta pagers don't need reset
func (p *channelIDPager) reset(context.Context) {}

func (p *channelIDPager) valuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.ChatMessageable](pl)
}

// ---------------------------------------------------------------------------
// delta item ID pager
// ---------------------------------------------------------------------------

var _ itemIDPager = &channelDeltaIDPager{}

type channelDeltaIDPager struct {
	gs          graph.Servicer
	userID      string
	containerID string
	builder     *teams.ItemChannelsItemMessagesDeltaRequestBuilder
	options     *teams.ItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration
}

func (c Channels) NewChannelsDeltaIDsPager(
	ctx context.Context,
	userID, containerID, oldDelta string,
	immutableIDs bool,
) (itemIDPager, error) {
	options := &teams.ItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(c.options.DeltaPageSize), preferImmutableIDs(immutableIDs)),
		QueryParameters: &teams.ItemChannelsItemMessagesDeltaRequestBuilderGetQueryParameters{
			// do NOT set Top.  It limits the total items received.
		},
	}

	var builder *teams.ItemChannelsItemMessagesDeltaRequestBuilder

	if oldDelta == "" {
		builder = getChannelsDeltaBuilder(ctx, c.Stable, userID, containerID, options)
	} else {
		builder = teams.NewItemChannelsItemMessagesDeltaRequestBuilder(oldDelta, c.Stable.Adapter())
	}

	return &channelDeltaIDPager{c.Stable, userID, containerID, builder, options}, nil
}

func getChannelsDeltaBuilder(
	ctx context.Context,
	gs graph.Servicer,
	teamID, containerID string,
	options *teams.ItemChannelsItemMessagesDeltaRequestBuilderGetRequestConfiguration,
) *teams.ItemChannelsItemMessagesDeltaRequestBuilder {

	builder := gs.
		Client().
		Teams().
		ByTeamId(teamID).
		Channels().
		ByChannelId(containerID).
		Messages().
		Delta()

	return builder
}

func (p *channelDeltaIDPager) getPage(ctx context.Context) (DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

func (p *channelDeltaIDPager) setNext(nextLink string) {
	p.builder = teams.NewItemChannelsItemMessagesDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *channelDeltaIDPager) reset(ctx context.Context) {
	p.builder = getChannelsDeltaBuilder(ctx, p.gs, p.userID, p.containerID, p.options)
}

func (p *channelDeltaIDPager) valuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Channelable](pl)
}

// ---------------------------------------------------------------------------
// helper funcs
// ---------------------------------------------------------------------------

// ChannelsDisplayable is a wrapper that complies with the
// models.Channelable interface with the graph.Container
// interfaces.
type ChannelsDisplayable struct {
	models.Channelable
}

// GetParentFolderId returns the default channe name address

//nolint:revive
func (c ChannelsDisplayable) GetParentFolderId() *string {
	return nil
}
