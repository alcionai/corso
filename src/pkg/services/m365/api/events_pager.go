package api

import (
	"context"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	eventBetaDeltaURLTemplate = "https://graph.microsoft.com/beta/users/%s/calendars/%s/events/delta"
)

// ---------------------------------------------------------------------------
// container pager
// ---------------------------------------------------------------------------

// EnumerateContainers iterates through all of the users current
// calendars, converting each to a graph.CacheFolder, and
// calling fn(cf) on each one.
// Folder hierarchy is represented in its current state, and does
// not contain historical data.
func (c Events) EnumerateContainers(
	ctx context.Context,
	userID, baseContainerID string,
	fn func(graph.CachedContainer) error,
	errs *fault.Bus,
) error {
	var (
		el     = errs.Local()
		config = &users.ItemCalendarsRequestBuilderGetRequestConfiguration{
			QueryParameters: &users.ItemCalendarsRequestBuilderGetQueryParameters{
				Select: idAnd("name"),
			},
		}
		builder = c.Stable.
			Client().
			Users().
			ByUserId(userID).
			Calendars()
	)

	for {
		if el.Failure() != nil {
			break
		}

		resp, err := builder.Get(ctx, config)
		if err != nil {
			return graph.Stack(ctx, err)
		}

		for _, cal := range resp.GetValue() {
			if el.Failure() != nil {
				break
			}

			cd := CalendarDisplayable{Calendarable: cal}
			if err := graph.CheckIDAndName(cd); err != nil {
				errs.AddRecoverable(ctx, graph.Stack(ctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}

			fctx := clues.Add(
				ctx,
				"container_id", ptr.Val(cal.GetId()),
				"container_name", ptr.Val(cal.GetName()))

			temp := graph.NewCacheFolder(
				cd,
				path.Builder{}.Append(ptr.Val(cd.GetId())),          // storage path
				path.Builder{}.Append(ptr.Val(cd.GetDisplayName()))) // display location
			if err := fn(&temp); err != nil {
				errs.AddRecoverable(ctx, graph.Stack(fctx, err).Label(fault.LabelForceNoBackupCreation))
				continue
			}
		}

		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = users.NewItemCalendarsRequestBuilder(link, c.Stable.Adapter())
	}

	return el.Failure()
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ itemPager[models.Eventable] = &eventsPageCtrl{}

type eventsPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemCalendarsItemEventsRequestBuilder
	options *users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration
}

func (c Events) NewEventsPager(
	userID, containerID string,
	selectProps ...string,
) itemPager[models.Eventable] {
	options := &users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize)),
		QueryParameters: &users.ItemCalendarsItemEventsRequestBuilderGetQueryParameters{
			// do NOT set Top.  It limits the total items received.
		},
	}

	if len(selectProps) > 0 {
		options.QueryParameters.Select = selectProps
	}

	builder := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Calendars().
		ByCalendarId(containerID).
		Events()

	return &eventsPageCtrl{c.Stable, builder, options}
}

//lint:ignore U1000 False Positive
func (p *eventsPageCtrl) getPage(ctx context.Context) (PageLinkValuer[models.Eventable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

//lint:ignore U1000 False Positive
func (p *eventsPageCtrl) setNext(nextLink string) {
	p.builder = users.NewItemCalendarsItemEventsRequestBuilder(nextLink, p.gs.Adapter())
}

//lint:ignore U1000 False Positive
func (c Events) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	userID, containerID string,
) (map[string]string, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewEventsPager(userID, containerID, eventCollisionKeyProps()...)

	items, err := enumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating events")
	}

	m := map[string]string{}

	for _, item := range items {
		m[EventCollisionKey(item)] = ptr.Val(item.GetId())
	}

	return m, nil
}

// ---------------------------------------------------------------------------
// item ID pager
// ---------------------------------------------------------------------------

var _ DeltaPager[getIDAndAddtler] = &eventIDPager{}

type eventIDPager struct {
	gs      graph.Servicer
	builder *users.ItemCalendarsItemEventsRequestBuilder
	options *users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration
}

func (c Events) NewEventIDsPager(
	ctx context.Context,
	userID, containerID string,
	immutableIDs bool,
) (DeltaPager[getIDAndAddtler], error) {
	options := &users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
		QueryParameters: &users.ItemCalendarsItemEventsRequestBuilderGetQueryParameters{
			// do NOT set Top.  It limits the total items received.
		},
	}

	builder := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Calendars().
		ByCalendarId(containerID).
		Events()

	return &eventIDPager{c.Stable, builder, options}, nil
}

func (p *eventIDPager) GetPage(ctx context.Context) (DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return EmptyDeltaLinker[models.Eventable]{PageLinkValuer: resp}, nil
}

func (p *eventIDPager) SetNext(nextLink string) {
	p.builder = users.NewItemCalendarsItemEventsRequestBuilder(nextLink, p.gs.Adapter())
}

// non delta pagers don't need reset
func (p *eventIDPager) Reset(context.Context) {}

func (p *eventIDPager) ValuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Eventable](pl)
}

// ---------------------------------------------------------------------------
// delta item ID pager
// ---------------------------------------------------------------------------

var _ DeltaPager[getIDAndAddtler] = &eventDeltaIDPager{}

type eventDeltaIDPager struct {
	gs          graph.Servicer
	userID      string
	containerID string
	builder     *users.ItemCalendarsItemEventsDeltaRequestBuilder
	options     *users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration
}

func (c Events) NewEventDeltaIDsPager(
	ctx context.Context,
	userID, containerID, oldDelta string,
	immutableIDs bool,
) (DeltaPager[getIDAndAddtler], error) {
	options := &users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(c.options.DeltaPageSize), preferImmutableIDs(immutableIDs)),
		QueryParameters: &users.ItemCalendarsItemEventsDeltaRequestBuilderGetQueryParameters{
			// do NOT set Top.  It limits the total items received.
		},
	}

	var builder *users.ItemCalendarsItemEventsDeltaRequestBuilder

	if oldDelta == "" {
		builder = getEventDeltaBuilder(ctx, c.Stable, userID, containerID, options)
	} else {
		builder = users.NewItemCalendarsItemEventsDeltaRequestBuilder(oldDelta, c.Stable.Adapter())
	}

	return &eventDeltaIDPager{c.Stable, userID, containerID, builder, options}, nil
}

func getEventDeltaBuilder(
	ctx context.Context,
	gs graph.Servicer,
	userID, containerID string,
	options *users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration,
) *users.ItemCalendarsItemEventsDeltaRequestBuilder {
	// Graph SDK only supports delta queries against events on the beta version, so we're
	// manufacturing use of the beta version url to make the call instead.
	// See: https://learn.microsoft.com/ko-kr/graph/api/event-delta?view=graph-rest-beta&tabs=http
	// Note that the delta item body is skeletal compared to the actual event struct.  Lucky
	// for us, we only need the item ID.  As a result, even though we hacked the version, the
	// response body parses properly into the v1.0 structs and complies with our wanted interfaces.
	// Likewise, the NextLink and DeltaLink odata tags carry our hack forward, so the rest of the code
	// works as intended (until, at least, we want to _not_ call the beta anymore).
	rawURL := fmt.Sprintf(eventBetaDeltaURLTemplate, userID, containerID)
	builder := users.NewItemCalendarsItemEventsDeltaRequestBuilder(rawURL, gs.Adapter())

	return builder
}

func (p *eventDeltaIDPager) GetPage(ctx context.Context) (DeltaPageLinker, error) {
	resp, err := p.builder.Get(ctx, p.options)
	if err != nil {
		return nil, graph.Stack(ctx, err)
	}

	return resp, nil
}

func (p *eventDeltaIDPager) SetNext(nextLink string) {
	p.builder = users.NewItemCalendarsItemEventsDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *eventDeltaIDPager) Reset(ctx context.Context) {
	p.builder = getEventDeltaBuilder(ctx, p.gs, p.userID, p.containerID, p.options)
}

func (p *eventDeltaIDPager) ValuesIn(pl PageLinker) ([]getIDAndAddtler, error) {
	return toValues[models.Eventable](pl)
}

func (c Events) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	userID, containerID, oldDelta string,
	immutableIDs bool,
	canMakeDeltaQueries bool,
) ([]string, []string, DeltaUpdate, error) {
	ctx = clues.Add(ctx, "container_id", containerID)

	pager, err := c.NewEventIDsPager(ctx, userID, containerID, immutableIDs)
	if err != nil {
		return nil, nil, DeltaUpdate{}, graph.Wrap(ctx, err, "creating non-delta pager")
	}

	deltaPager, err := c.NewEventDeltaIDsPager(ctx, userID, containerID, oldDelta, immutableIDs)
	if err != nil {
		return nil, nil, DeltaUpdate{}, graph.Wrap(ctx, err, "creating delta pager")
	}

	return getAddedAndRemovedItemIDs(ctx, c.Stable, pager, deltaPager, oldDelta, canMakeDeltaQueries)
}
