package api

import (
	"context"

	"github.com/microsoft/kiota-abstractions-go/serialization"

	"github.com/alcionai/corso/src/pkg/fault"
)

type GetAndSerializeItemer[INFO any] interface {
	GetItemer[INFO]
	Serializer
}

type GetItemer[INFO any] interface {
	GetItem(
		ctx context.Context,
		protectedResource, itemID string,
		useImmutableIDs bool,
		errs *fault.Bus,
	) (serialization.Parsable, *INFO, error)
}

type Serializer interface {
	Serialize(
		ctx context.Context,
		item serialization.Parsable,
		protectedResource, itemID string,
	) ([]byte, error)
}
