package sharepoint

import (
	"io"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/path"
)

type Target int

//go:generate stringer -type=Target
const (
	Unknown Target = iota
	List
	Drive
)

var (
	_ data.Collection = &Collection{}
	_ data.Stream     = &Item{}
)

type Collection struct {
	data chan data.Stream
	// folderPath indicates the hierarchy within the collection
	folderPath path.Path
	// M365 IDs of the items of this collection
	service       graph.Service
	statusUpdater support.StatusUpdater
}

func (sc *Collection) FullPath() path.Path {
	return sc.FullPath()
}

func (sc *Collection) Items() <-chan data.Stream {
	return sc.data
}

type Item struct {
	id   string
	data io.ReadCloser
}

func (sd *Item) UUID() string {
	return sd.id
}

func (sd *Item) ToReader() io.ReadCloser {
	return sd.data
}
