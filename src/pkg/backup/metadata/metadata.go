package metadata

import "github.com/alcionai/corso/src/pkg/path"

const (
	// DeltaURLsFileName is the name of the file containing delta token(s) for a
	// given endpoint. The endpoint granularity varies by service.
	DeltaURLsFileName = "delta"

	// PreviousPathFileName is the name of the file containing previous path(s) for a
	// given endpoint.
	PreviousPathFileName = "previouspath"

	PathKey  = "path"
	DeltaKey = "delta"
)

type (
	CatDeltaPaths map[path.CategoryType]DeltaPaths
	DeltaPaths    map[string]DeltaPath
	DeltaPath     struct {
		Delta string
		Path  string
	}
)

func (dps DeltaPaths) AddDelta(k, d string) {
	dp, ok := dps[k]
	if !ok {
		dp = DeltaPath{}
	}

	dp.Delta = d
	dps[k] = dp
}

func (dps DeltaPaths) AddPath(k, p string) {
	dp, ok := dps[k]
	if !ok {
		dp = DeltaPath{}
	}

	dp.Path = p
	dps[k] = dp
}

// AllMetadataFileNames produces the standard set of filenames used to store graph
// metadata such as delta tokens and folderID->path references.
func AllMetadataFileNames() []string {
	return []string{DeltaURLsFileName, PreviousPathFileName}
}
