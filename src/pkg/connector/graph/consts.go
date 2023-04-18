package graph

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
