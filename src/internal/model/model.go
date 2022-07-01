package model

import (
	"github.com/kopia/kopia/repo/manifest"
)

type ID string

type Model interface {
	// Returns a handle to the BaseModel for this model.
	Base() *BaseModel
}

// BaseModel defines required fields for models stored in ModelStore. Structs
// that wish to be stored should embed this struct. This struct also represents
// the common metadata the ModelStore will fill out/use.
type BaseModel struct {
	// StableID is an identifier that other objects can use to refer to this
	// object in the ModelStore.
	// Once generated (during Put), it is guaranteed not to change. This field
	// should be treated as read-only by users.
	StableID ID `json:"stableID,omitempty"`
	// ModelStoreID is an internal ID for the model in the store. If present it
	// can be used for efficient lookups, but should not be used by other models
	// to refer to this one. This field may change if the model is updated. This
	// field should be treated as read-only by users.
	ModelStoreID manifest.ID `json:"modelStoreID,omitempty"`
	// Tags associated with this model in the store to facilitate lookup. Tags in
	// the struct are not serialized directly into the stored model, but are part
	// of the metadata for the model.
	Tags map[string]string `json:"-"`
}

func (bm *BaseModel) Base() *BaseModel {
	return bm
}
