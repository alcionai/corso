package graph

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	_ data.BackupCollection = &MetadataCollection{}
	_ data.Item             = &metadataItem{}
)

// MetadataCollection in a simple collection that assumes all items to be
// returned are already resident in-memory and known when the collection is
// created. This collection has no logic for lazily fetching item data.
type MetadataCollection struct {
	fullPath      path.Path
	items         []metadataItem
	statusUpdater support.StatusUpdater
}

// MetadataCollectionEntry describes a file that should get added to a metadata
// collection.  The Data value will be encoded into json as part of a
// transformation into a MetadataItem.
type MetadataCollectionEntry struct {
	fileName string
	data     any
}

func NewMetadataEntry(fileName string, mData any) MetadataCollectionEntry {
	return MetadataCollectionEntry{fileName, mData}
}

func (mce MetadataCollectionEntry) toMetadataItem() (metadataItem, error) {
	if len(mce.fileName) == 0 {
		return metadataItem{}, clues.New("missing metadata filename")
	}

	if mce.data == nil {
		return metadataItem{}, clues.New("missing metadata")
	}

	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)

	if err := encoder.Encode(mce.data); err != nil {
		return metadataItem{}, clues.Wrap(err, "serializing metadata")
	}

	return metadataItem{
		Item: data.NewUnindexedPrefetchedItem(
			io.NopCloser(buf),
			mce.fileName,
			time.Now()),
		size: int64(buf.Len()),
	}, nil
}

// MakeMetadataCollection creates a metadata collection that has a file
// containing all the provided metadata as a single json object. Returns
// nil if the map does not have any entries.
func MakeMetadataCollection(
	pathPrefix path.Path,
	metadata []MetadataCollectionEntry,
	statusUpdater support.StatusUpdater,
) (data.BackupCollection, error) {
	if len(metadata) == 0 {
		return nil, nil
	}

	items := make([]metadataItem, 0, len(metadata))

	for _, md := range metadata {
		item, err := md.toMetadataItem()
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	coll := NewMetadataCollection(pathPrefix, items, statusUpdater)

	return coll, nil
}

func NewMetadataCollection(
	p path.Path,
	items []metadataItem,
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

// TODO(ashmrtn): Fill in with previous path once the Controller compares old
// and new folder hierarchies.
func (md MetadataCollection) PreviousPath() path.Path {
	return nil
}

// TODO(ashmrtn): Fill in once the Controller compares old and new folder
// hierarchies.
func (md MetadataCollection) State() data.CollectionState {
	return data.NewState
}

func (md MetadataCollection) DoNotMergeItems() bool {
	return false
}

func (md MetadataCollection) Items(
	ctx context.Context,
	_ *fault.Bus, // not used, just here for interface compliance
) <-chan data.Item {
	res := make(chan data.Item)

	go func() {
		totalBytes := int64(0)

		defer func() {
			// Need to report after the collection is created because otherwise
			// statusUpdater may not have accounted for the fact that this collection
			// will be running.
			status := support.CreateStatus(
				ctx,
				support.Backup,
				1,
				support.CollectionMetrics{
					Objects:   len(md.items),
					Successes: len(md.items),
					Bytes:     totalBytes,
				},
				md.fullPath.Folder(false))

			md.statusUpdater(status)
		}()
		defer close(res)

		for _, item := range md.items {
			totalBytes += item.size
			res <- item
		}
	}()

	return res
}

type metadataItem struct {
	data.Item
	size int64
}
