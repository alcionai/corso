package api

import (
	"context"
	"errors"
	"fmt"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/users"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

const eventBetaDeltaURLTemplate = "https://graph.microsoft.com/beta/users/%s/calendars/%s/events/delta"

// ---------------------------------------------------------------------------
// container pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.Calendarable] = &eventsCalendarsPageCtrl{}

type eventsCalendarsPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemCalendarsRequestBuilder
	options *users.ItemCalendarsRequestBuilderGetRequestConfiguration
}

func (c Events) NewEventCalendarsPager(
	userID string,
	selectProps ...string,
) pagers.NonDeltaHandler[models.Calendarable] {
	options := &users.ItemCalendarsRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(
			preferPageSize(maxNonDeltaPageSize),
			preferImmutableIDs(c.options.ToggleFeatures.ExchangeImmutableIDs)),
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
) (pagers.NextLinkValuer[models.Calendarable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, clues.Stack(err).OrNil()
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
) ([]models.Calendarable, error) {
	containers, err := pagers.BatchEnumerateItems(ctx, c.NewEventCalendarsPager(userID))
	return containers, clues.Stack(err).OrNil()
}

// ---------------------------------------------------------------------------
// item pager
// ---------------------------------------------------------------------------

var _ pagers.NonDeltaHandler[models.Eventable] = &eventsPageCtrl{}

type eventsPageCtrl struct {
	gs      graph.Servicer
	builder *users.ItemCalendarsItemEventsRequestBuilder
	options *users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration
}

func (c Events) NewEventsPager(
	userID, containerID string,
	selectProps ...string,
) pagers.NonDeltaHandler[models.Eventable] {
	options := &users.ItemCalendarsItemEventsRequestBuilderGetRequestConfiguration{
		Headers: newPreferHeaders(
			preferPageSize(maxNonDeltaPageSize),
			preferImmutableIDs(c.options.ToggleFeatures.ExchangeImmutableIDs)),
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
) (pagers.NextLinkValuer[models.Eventable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, clues.Stack(err).OrNil()
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
	pager := c.NewEventsPager(userID, containerID, eventCollisionKeyProps()...)

	items, err := pagers.BatchEnumerateItems(ctx, pager)
	if err != nil {
		return nil, clues.Wrap(err, "enumerating events")
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
	pager := c.NewEventsPager(userID, containerID, idAnd()...)

	items, err := pagers.BatchEnumerateItems(ctx, pager)
	if err != nil {
		return nil, clues.Wrap(err, "enumerating events")
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

var _ pagers.DeltaHandler[models.Eventable] = &eventDeltaPager{}

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
	selectProps ...string,
) pagers.DeltaHandler[models.Eventable] {
	return c.newEventsDeltaPagerWithPageSize(
		ctx,
		userID,
		containerID,
		prevDeltaLink,
		c.options.DeltaPageSize,
		selectProps...)
}

func (c Events) newEventsDeltaPagerWithPageSize(
	ctx context.Context,
	userID string,
	containerID string,
	prevDeltaLink string,
	pageSize int32,
	selectProps ...string,
) pagers.DeltaHandler[models.Eventable] {
	options := &users.ItemCalendarsItemEventsDeltaRequestBuilderGetRequestConfiguration{
		// do NOT set Top.  It limits the total items received.
		QueryParameters: &users.ItemCalendarsItemEventsDeltaRequestBuilderGetQueryParameters{},
		Headers: newPreferHeaders(
			preferPageSize(pageSize),
			preferImmutableIDs(c.options.ToggleFeatures.ExchangeImmutableIDs)),
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
) (pagers.DeltaLinkValuer[models.Eventable], error) {
	resp, err := p.builder.Get(ctx, p.options)
	return resp, clues.Stack(err).OrNil()
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
	config CallConfig,
) (pagers.AddedAndRemoved, error) {
	ctx = clues.Add(
		ctx,
		"data_category", path.EventsCategory,
		"container_id", containerID)

	deltaPager := c.NewEventsDeltaPager(
		ctx,
		userID,
		containerID,
		prevDeltaLink,
		idAnd()...)

	// Experiments showed that non-delta endpoint didn't have the same performance
	// degradation that delta endpoint did so a larger page size should be ok.
	pager := c.NewEventsPager(
		userID,
		containerID,
		idAnd(lastModifiedDateTime)...)

	// Try running the query with the given limits. If we fail to do a delta
	// enumeration with a 5xx error try rerunning the query with a smaller page
	// size. We've seen some resources where this consistently happens but a small
	// page size allows us to make progress.
	addedRemoved, err := pagers.GetAddedAndRemovedItemIDs[models.Eventable](
		ctx,
		pager,
		deltaPager,
		prevDeltaLink,
		config.CanMakeDeltaQueries,
		config.LimitResults,
		pagers.AddedAndRemovedByAddtlData[models.Eventable])
	if err == nil || !errors.Is(err, graph.ErrServiceUnavailableEmptyResp) {
		return addedRemoved, clues.Stack(err).OrNil()
	}

	effectivePageSize := minEventsDeltaPageSize

	logger.Ctx(ctx).Infow(
		"retrying event item query with reduced page size",
		"delta_pager_effective_page_size", effectivePageSize,
		"delta_pager_default_page_size", c.options.DeltaPageSize)

	deltaPager = c.newEventsDeltaPagerWithPageSize(
		ctx,
		userID,
		containerID,
		prevDeltaLink,
		effectivePageSize,
		idAnd()...)

	return pagers.GetAddedAndRemovedItemIDs[models.Eventable](
		ctx,
		pager,
		deltaPager,
		prevDeltaLink,
		config.CanMakeDeltaQueries,
		config.LimitResults,
		pagers.AddedAndRemovedByAddtlData[models.Eventable])
}
