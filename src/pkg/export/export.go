package export

import (
	"context"
	"io"
)

// Collection is the interface that is returned to the SDK consumer
type Collection interface {
	// BasePath gets the base path of the collection. This is derived
	// from FullPath, but trim out thing like drive id or any other part
	// that is not needed to show the path to the collection.
	BasePath() string

	// Items gets the items within the collection(folder)
	Items(context.Context) <-chan Item
}

// ItemData is the data for an individual item.
type ItemData struct {
	// Name is the name of the item. This is the name that the item
	// would have had in the service.
	Name string

	// Body is the body of the item. This is an io.ReadCloser and the
	// SDK consumer is responsible for closing it.
	Body io.ReadCloser
}

// Item is the item that is returned to the SDK consumer
type Item struct {
	// ID will be a unique id for the item. This is same as the id
	// that is used to store the data. This is not the name and is
	// mostly used just for tracking.
	ID string

	// Data contains the actual data of the item. It will have both
	// the name of the item and an io.ReadCloser which contains the
	// body of the item.
	Data ItemData

	// Error will contain any error that happened while trying to get
	// the item/items like when trying to resolve the name of the item.
	// In case we have the error bound to a particular item, we will
	// also return the id of the item.
	Error error
}
