package api

import (
	"context"
	"fmt"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/pkg/path"
)

const eventBetaDeltaURLTemplate = "https://graph.microsoft.com/beta/users/%s/calendars/%s/events/delta"

// ---------------------------------------------------------------------------
// container pager
// ---------------------------------------------------------------------------

var _ Pager[models.Calendarable] = &eventsCalendarsPageCtrl{}

type eventsCalendarsPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemCalendarsRequestBuilder
	options *users.ItemCalendarsRequestBuilderGetRequestConfiguration
}

func (c Events) NewEventCalendarsPager(
	userID string,
	immutableIDs bool,
	selectProps ...string,
) Pager[models.Calendarable] {
	options := &users.ItemCalendarsRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
		QueryParameters: &users.ItemCalendarsRequestBuilderGetQueryParameters{},
		// do NOT set Top.  It limits the total items received.
	}

	if len(selectProps) > 0 {
		options.QueryParameters.Select = selectProps
	}

	builder := c.Stable.
		Client().
		Users().
		ByUserId(userID).
		Calendars()

	return &eventsCalendarsPageCtrl{c.Stable, builder, options}
}

func (p *eventsCalendarsPageCtrl) GetPage(
	ctx context.Context,
) (NextLinkValuer[models.Calendarable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *eventsCalendarsPageCtrl) SetNextLink(nextLink string) {
	p.builder = users.NewItemCalendarsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *eventsCalendarsPageCtrl) ValidModTimes() bool {
	return true
}

// EnumerateContainers retrieves all of the user's current mail folders.
func (c Events) EnumerateContainers(
	ctx context.Context,
	userID, _ string, // baseContainerID not needed here
	immutableIDs bool,
) ([]models.Calendarable, error) {
	containers, err := batchEnumerateItems(ctx, c.NewEventCalendarsPager(userID, immutableIDs))
	return containers, graph.Stack(ctx, err).OrNil()
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ Pager[models.Eventable] = &eventsPageCtrl{}

type eventsPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemCalendarsItemEventsRequestBuilder
	options *users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration
}

func (c Events) NewEventsPager(
	userID, containerID string,
	immutableIDs bool,
	selectProps ...string,
) Pager[models.Eventable] {
	options := &users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration{
		Headers:         newPreferHeaders(preferPageSize(maxNonDeltaPageSize), preferImmutableIDs(immutableIDs)),
		QueryParameters: &users.ItemCalendarsItemEventsRequestBuilderGetQueryParameters{},
		// do NOT set Top.  It limits the total items received.
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

func (p *eventsPageCtrl) GetPage(
	ctx context.Context,
) (NextLinkValuer[models.Eventable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *eventsPageCtrl) SetNextLink(nextLink string) {
	p.builder = users.NewItemCalendarsItemEventsRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *eventsPageCtrl) ValidModTimes() bool {
	return true
}

func (c Events) GetItemsInContainerByCollisionKey(
	ctx context.Context,
	userID, containerID string,
) (map[string]string, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewEventsPager(userID, containerID, false, eventCollisionKeyProps()...)

	items, err := batchEnumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating events")
	}

	m := map[string]string{}

	for _, item := range items {
		m[EventCollisionKey(item)] = ptr.Val(item.GetId())
	}

	return m, nil
}

func (c Events) GetItemIDsInContainer(
	ctx context.Context,
	userID, containerID string,
) (map[string]struct{}, error) {
	ctx = clues.Add(ctx, "container_id", containerID)
	pager := c.NewEventsPager(userID, containerID, false, idAnd()...)

	items, err := batchEnumerateItems(ctx, pager)
	if err != nil {
		return nil, graph.Wrap(ctx, err, "enumerating events")
	}

	m := map[string]struct{}{}

	for _, item := range items {
		m[ptr.Val(item.GetId())] = struct{}{}
	}

	return m, nil
}

// ---------------------------------------------------------------------------
// delta item ID pager
// ---------------------------------------------------------------------------

var _ DeltaPager[models.Eventable] = &eventDeltaPager{}

type eventDeltaPager struct {
	gs          graph.Servicer
	userID      string
	containerID string
	builder     *users.ItemCalendarsItemEventsDeltaRequestBuilder
	options     *users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration
}

func getEventDeltaBuilder(
	ctx context.Context,
	gs graph.Servicer,
	userID, containerID string,
) *users.ItemCalendarsItemEventsDeltaRequestBuilder {
	rawURL := fmt.Sprintf(eventBetaDeltaURLTemplate, userID, containerID)
	return users.NewItemCalendarsItemEventsDeltaRequestBuilder(rawURL, gs.Adapter())
}

func (c Events) NewEventsDeltaPager(
	ctx context.Context,
	userID, containerID, prevDeltaLink string,
	immutableIDs bool,
	selectProps ...string,
) DeltaPager[models.Eventable] {
	options := &users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration{
		// do NOT set Top.  It limits the total items received.
		QueryParameters: &users.ItemCalendarsItemEventsDeltaRequestBuilderGetQueryParameters{},
		Headers:         newPreferHeaders(preferPageSize(c.options.DeltaPageSize), preferImmutableIDs(immutableIDs)),
	}

	if len(selectProps) > 0 {
		options.QueryParameters.Select = selectProps
	}

	var builder *users.ItemCalendarsItemEventsDeltaRequestBuilder

	if len(prevDeltaLink) > 0 {
		builder = users.NewItemCalendarsItemEventsDeltaRequestBuilder(prevDeltaLink, c.Stable.Adapter())
	} else {
		builder = getEventDeltaBuilder(ctx, c.Stable, userID, containerID)
	}

	return &eventDeltaPager{c.Stable, userID, containerID, builder, options}
}

func (p *eventDeltaPager) GetPage(
	ctx context.Context,
) (DeltaLinkValuer[models.Eventable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, graph.Stack(ctx, err).OrNil()
}

func (p *eventDeltaPager) SetNextLink(nextLink string) {
	p.builder = users.NewItemCalendarsItemEventsDeltaRequestBuilder(nextLink, p.gs.Adapter())
}

func (p *eventDeltaPager) Reset(ctx context.Context) {
	p.builder = getEventDeltaBuilder(ctx, p.gs, p.userID, p.containerID)
}

func (p *eventDeltaPager) ValidModTimes() bool {
	return false
}

func (c Events) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	userID, containerID, prevDeltaLink string,
	immutableIDs bool,
	canMakeDeltaQueries bool,
) (map[string]time.Time, bool, []string, DeltaUpdate, error) {
	ctx = clues.Add(
		ctx,
		"data_category", path.EventsCategory,
		"container_id", containerID)

	deltaPager := c.NewEventsDeltaPager(
		ctx,
		userID,
		containerID,
		prevDeltaLink,
		immutableIDs,
		idAnd()...)
	pager := c.NewEventsPager(
		userID,
		containerID,
		immutableIDs,
		idAnd(lastModifiedDateTime)...)

	return getAddedAndRemovedItemIDs[models.Eventable](
		ctx,
		pager,
		deltaPager,
		prevDeltaLink,
		canMakeDeltaQueries,
		addedAndRemovedByAddtlData[models.Eventable])
}
