package export

import (
	"context"
	"io"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/metrics"
)

// ---------------------------------------------------------------------------
// Collections
// ---------------------------------------------------------------------------

// Collectioner is the interface that is returned to the SDK consumer
type Collectioner interface {
	// BasePath gets the base path of the collection. This is derived
	// from FullPath, but trim out thing like drive id or any other part
	// that is not needed to show the path to the collection.
	BasePath() string

	// Items gets the items within the collection(folder)
	Items(context.Context) <-chan Item
}

type itemStreamer func(
	ctx context.Context,
	backingColls []data.RestoreCollection,
	backupVersion int,
	cfg control.ExportConfig,
	ch chan<- Item,
	stats *metrics.ExportStats)

// BaseCollection holds the foundational details of an export collection.
type BaseCollection struct {
	// BaseDir contains the destination path of the collection.
	BaseDir string

	// BackingCollection is the restore collection from which we will
	// create the export collection.
	BackingCollection []data.RestoreCollection

	// BackupVersion is the backupVersion of the data source.
	BackupVersion int

	Cfg control.ExportConfig

	Stream itemStreamer

	Stats *metrics.ExportStats
}

func (bc BaseCollection) BasePath() string {
	return bc.BaseDir
}

func (bc BaseCollection) Items(ctx context.Context) <-chan Item {
	ch := make(chan Item)
	go bc.Stream(ctx, bc.BackingCollection, bc.BackupVersion, bc.Cfg, ch, bc.Stats)

	return ch
}

// ---------------------------------------------------------------------------
// Items
// ---------------------------------------------------------------------------

// Item is the item that is returned to the SDK consumer
type Item struct {
	// ID will be a unique id for the item. This is same as the id
	// that is used to store the data. This is not the name and is
	// mostly used just for tracking.
	ID string

	// Name is the name of the item. This is the name that the item
	// would have had in the service.
	Name string

	// Body is the body of the item. This is an io.ReadCloser and the
	// SDK consumer is responsible for closing it.
	Body io.ReadCloser

	// Error will contain any error that happened while trying to get
	// the item/items like when trying to resolve the name of the item.
	// In case we have the error bound to a particular item, we will
	// also return the id of the item.
	Error error
}
