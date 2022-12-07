package graph

import (
	"bytes"
	"context"
	"io"

	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	_ data.Collection = &MetadataCollection{}
	_ data.Stream     = &MetadataItem{}
)

// MetadataCollection in a simple collection that assumes all items to be
// returned are already resident in-memory and known when the collection is
// created. This collection has no logic for lazily fetching item data.
type MetadataCollection struct {
	fullPath      path.Path
	items         []MetadataItem
	statusUpdater support.StatusUpdater
}

func NewMetadataCollection(
	p path.Path,
	items []MetadataItem,
	statusUpdater support.StatusUpdater,
) *MetadataCollection {
	return &MetadataCollection{
		fullPath:      p,
		items:         items,
		statusUpdater: statusUpdater,
	}
}

func (md MetadataCollection) FullPath() path.Path {
	return md.fullPath
}

func (md MetadataCollection) Items() <-chan data.Stream {
	res := make(chan data.Stream)

	go func() {
		totalBytes := int64(0)

		defer func() {
			// Need to report after the collection is created because otherwise
			// statusUpdater may not have accounted for the fact that this collection
			// will be running.
			status := support.CreateStatus(
				context.TODO(),
				support.Backup,
				1,
				support.CollectionMetrics{
					Objects:    len(md.items),
					Successes:  len(md.items),
					TotalBytes: totalBytes,
				},
				nil,
				md.fullPath.Folder(),
			)

			md.statusUpdater(status)
		}()
		defer close(res)

		for _, item := range md.items {
			totalBytes += int64(len(item.data))
			res <- item
		}
	}()

	return res
}

// MetadataItem is an in-memory data.Stream implementation. MetadataItem does
// not implement additional interfaces like data.StreamInfo, so it should only
// be used for items with a small amount of content that don't need to be added
// to backup details.
//
// Currently the expected use-case for this struct are storing metadata for a
// backup like delta tokens or a mapping of container IDs to container paths.
type MetadataItem struct {
	// uuid is an ID that can be used to refer to the item.
	uuid string
	// data is a buffer of data that the item refers to.
	data []byte
}

func NewMetadataItem(uuid string, itemData []byte) MetadataItem {
	return MetadataItem{
		uuid: uuid,
		data: itemData,
	}
}

func (mi MetadataItem) UUID() string {
	return mi.uuid
}

func (mi MetadataItem) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(mi.data))
}
