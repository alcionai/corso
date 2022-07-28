// Package onedrive provides support for retrieving M365 OneDrive objects
package onedrive

import (
	"bytes"
	"io"

	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/pkg/backup/details"
)

var _ data.Collection = &Collection{}
var _ data.Stream = &Item{}
var _ data.StreamInfo = &Item{}

// Collection represents a set of OneDrive objects retreived from M365
type Collection struct {
	data       chan data.Stream
	folderPath string

	// file items in this collection
	driveItems []string
}

// NewCollection creates a Collection
func NewCollection(folderPath string) *Collection {
	return &Collection{
		folderPath: folderPath,
		driveItems: make([]string, 0),
	}
}

// Items() returns the channel containing M365 Exchange objects
func (oc *Collection) Items() <-chan data.Stream {
	return oc.data
}

func (oc *Collection) FullPath() []string {
	return []string{}
}

// Item represents a single item retrieved from OneDrive
type Item struct {
	id        string
	driveItem []byte                // TODO: Use DriveItem type?
	info      *details.OnedriveInfo //temporary change to bring populate function into directory
}

func (od *Item) UUID() string {
	return od.id
}

func (od *Item) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(od.driveItem))
}

func (od *Item) Info() details.ItemInfo {
	return details.ItemInfo{Onedrive: od.info}
}
