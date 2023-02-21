package graph

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	_ data.BackupCollection = &MetadataCollection{}
	_ data.Stream           = &MetadataItem{}
)

// MetadataCollection in a simple collection that assumes all items to be
// returned are already resident in-memory and known when the collection is
// created. This collection has no logic for lazily fetching item data.
type MetadataCollection struct {
	fullPath      path.Path
	items         []MetadataItem
	statusUpdater support.StatusUpdater
}

// MetadataCollecionEntry describes a file that should get added to a metadata
// collection.  The Data value will be encoded into json as part of a
// transformation into a MetadataItem.
type MetadataCollectionEntry struct {
	fileName string
	data     any
}

func NewMetadataEntry(fileName string, mData any) MetadataCollectionEntry {
	return MetadataCollectionEntry{fileName, mData}
}

func (mce MetadataCollectionEntry) toMetadataItem() (MetadataItem, error) {
	if len(mce.fileName) == 0 {
		return MetadataItem{}, errors.New("missing metadata filename")
	}

	if mce.data == nil {
		return MetadataItem{}, errors.New("missing metadata")
	}

	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)

	if err := encoder.Encode(mce.data); err != nil {
		return MetadataItem{}, errors.Wrap(err, "serializing metadata")
	}

	return NewMetadataItem(mce.fileName, buf.Bytes()), nil
}

// MakeMetadataCollection creates a metadata collection that has a file
// containing all the provided metadata as a single json object. Returns
// nil if the map does not have any entries.
func MakeMetadataCollection(
	tenant, resourceOwner string,
	service path.ServiceType,
	cat path.CategoryType,
	metadata []MetadataCollectionEntry,
	statusUpdater support.StatusUpdater,
) (data.BackupCollection, error) {
	if len(metadata) == 0 {
		return nil, nil
	}

	p, err := path.Builder{}.ToServiceCategoryMetadataPath(
		tenant,
		resourceOwner,
		service,
		cat,
		false,
	)
	if err != nil {
		return nil, errors.Wrap(err, "making metadata path")
	}

	items := make([]MetadataItem, 0, len(metadata))

	for _, md := range metadata {
		item, err := md.toMetadataItem()
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	coll := NewMetadataCollection(p, items, statusUpdater)

	return coll, nil
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

// TODO(ashmrtn): Fill in with previous path once GraphConnector compares old
// and new folder hierarchies.
func (md MetadataCollection) PreviousPath() path.Path {
	return nil
}

// TODO(ashmrtn): Fill in once GraphConnector compares old and new folder
// hierarchies.
func (md MetadataCollection) State() data.CollectionState {
	return data.NewState
}

func (md MetadataCollection) DoNotMergeItems() bool {
	return false
}

func (md MetadataCollection) Items(
	ctx context.Context,
	errs *fault.Bus,
) <-chan data.Stream {
	res := make(chan data.Stream)

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
				md.fullPath.Folder(false),
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

// TODO(ashmrtn): Fill in once we know how to handle this.
func (mi MetadataItem) Deleted() bool {
	return false
}

func (mi MetadataItem) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewReader(mi.data))
}
