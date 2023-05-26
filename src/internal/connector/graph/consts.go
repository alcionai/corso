package graph

import (
	"context"

	"github.com/alcionai/corso/src/pkg/logger"
	"github.com/alcionai/corso/src/pkg/path"
)

const AttachmentChunkSize = 4 * 1024 * 1024

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
	// sets the collection buffer size before blocking.
	collectionBuffer int
	// sets the parallelism of item population within a collection.
	item int
}

func (p parallelism) CollectionBufferSize() int {
	if p.collectionBuffer == 0 {
		return 1
	}

	return p.collectionBuffer
}

func (p parallelism) CollectionBufferOverride(ctx context.Context, override int) int {
	logger.Ctx(ctx).Infow(
		"collection buffer parallelism",
		"default_parallelism", p.collectionBuffer,
		"requested_paralellism", override)

	if !isWithin(1, p.collectionBuffer, override) {
		return p.collectionBuffer
	}

	return override
}

func (p parallelism) ItemOverride(ctx context.Context, override int) int {
	logger.Ctx(ctx).Infow(
		"item-level parallelism",
		"default_parallelism", p.item,
		"requested_paralellism", override)

	if !isWithin(1, p.item, override) {
		return p.item
	}

	return override
}

func (p parallelism) Item() int {
	if p.item == 0 {
		return 1
	}

	return p.item
}

// returns low <= v <= high
// if high < low, returns low <= v
func isWithin(low, high, v int) bool {
	return v >= low && (high < low || v <= high)
}

var sp = map[path.ServiceType]parallelism{
	path.ExchangeService: {
		collectionBuffer: 4,
		item:             4,
	},
	path.OneDriveService: {
		collectionBuffer: 5,
		item:             4,
	},
	// sharepoint libraries are considered "onedrive" parallelism.
	// this only controls lists/pages.
	path.SharePointService: {
		collectionBuffer: 5,
		item:             4,
	},
}

// Parallelism returns the Parallelism for the requested service.
func Parallelism(srv path.ServiceType) parallelism {
	return sp[srv]
}
