package graph

import (
	"context"

	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// item response AdditionalData
// ---------------------------------------------------------------------------

const (
	// AddtlDataRemoved is the key value in the AdditionalData map
	// for when an item was deleted.
	//nolint:lll
	// https://learn.microsoft.com/en-us/graph/delta-query-overview?tabs=http#resource-representation-in-the-delta-query-response
	AddtlDataRemoved = "@removed"
)

// ---------------------------------------------------------------------------
// Metadata Files
// ---------------------------------------------------------------------------

const (
	// DeltaURLsFileName is the name of the file containing delta token(s) for a
	// given endpoint. The endpoint granularity varies by service.
	DeltaURLsFileName = "delta"

	// PreviousPathFileName is the name of the file containing previous path(s) for a
	// given endpoint.
	PreviousPathFileName = "previouspath"
)

// ---------------------------------------------------------------------------
// Runtime Configuration
// ---------------------------------------------------------------------------

type parallelism struct {
	collectionPool int
	itemFetch      int
	prefetchPool   int
}

func (p parallelism) CollectionPoolSize() int {
	return p.collectionPool
}

func (p parallelism) CollectionPoolOverride(ctx context.Context, override int) int {
	logger.Ctx(ctx).Infow(
		"collection pool parallelism",
		"default_parallelism", p.itemFetch,
		"requested_paralellism", override)

	if override < 1 || (p.collectionPool > 0 && override > p.collectionPool) {
		return p.collectionPool
	}

	return override
}

func (p parallelism) ItemFetchOverride(ctx context.Context, override int) int {
	logger.Ctx(ctx).Infow(
		"item fetch parallelism",
		"default_parallelism", p.itemFetch,
		"requested_paralellism", override)

	if override < 1 || (p.itemFetch > 0 && override > p.itemFetch) {
		return p.itemFetch
	}

	return override
}

func (p parallelism) ItemFetchSize() int {
	return p.itemFetch
}

func (p parallelism) PrefetchPoolOverride(ctx context.Context, override int) int {
	logger.Ctx(ctx).Infow(
		"item fetch parallelism",
		"default_parallelism", p.itemFetch,
		"requested_paralellism", override)

	if override < 1 || (p.prefetchPool > 0 && override > p.prefetchPool) {
		return p.prefetchPool
	}

	return override
}

func (p parallelism) PrefetchPoolSize() int {
	return p.prefetchPool
}

var sp = map[path.ServiceType]parallelism{
	path.ExchangeService: {
		collectionPool: 4,
		itemFetch:      4,
		prefetchPool:   4,
	},
	path.OneDriveService: {
		collectionPool: 5,
		itemFetch:      4,
		prefetchPool:   5,
	},
	path.SharePointService: {
		collectionPool: 5,
		itemFetch:      4,
		prefetchPool:   5,
	},
}

// Parallelism returns the Parallelism for the requested service.
func Parallelism(srv path.ServiceType) parallelism {
	return sp[srv]
}
