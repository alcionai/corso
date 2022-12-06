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

type MetadataItem struct {
	name string
	data []byte
}

func NewMetadataItem(name string, itemData []byte) MetadataItem {
	return MetadataItem{
		name: name,
		data: itemData,
	}
}

func (mi MetadataItem) UUID() string {
	return mi.name
}

func (mi MetadataItem) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(mi.data))
}
