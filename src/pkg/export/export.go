package export

import (
	"context"
	"io"
)

// ExportCollection is the interface that is returned to the SDK consumer
type ExportCollection interface {
	// GetBasePath gets the base path of the collection
	GetBasePath() string

	// GetItems gets the items within the collection(folder)
	GetItems(context.Context) <-chan ExportItem
}

// ExportItemData is the data for an individual item.
type ExportItemData struct {
	// Name is the name of the item. This is the name that the item
	// would have had in the service.
	Name string

	// Body is the body of the item. This is an io.ReadCloser and the
	// SDK consumer is responsible for closing it.
	Body io.ReadCloser
}

// ExportItem is the item that is returned to the SDK consumer
type ExportItem struct {
	// ID will be a unique id for the item. This is same as the id
	// that is used to store the data. This is not the name and is
	// mostly used just for tracking.
	ID string

	// Data contains the actual data of the item. It will have both
	// the name of the item and an io.ReadCloser which contains the
	// body of the item.
	Data ExportItemData

	// Error will contain any error that happened while trying to get
	// the item like when trying to resolve the name of the item.
	Error error
}
