package model

import (
	"github.com/kopia/kopia/repo/manifest"
)

type ID string

type Model interface {
	GetStableID() ID
	SetStableID(id ID)
	GetModelStoreID() manifest.ID
	SetModelStoreID(id manifest.ID)
}

// BaseModel defines required fields for models stored in ModelStore. Structs
// that wish to be stored should embed this struct.
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
}

func (bm *BaseModel) GetStableID() ID {
	return bm.StableID
}

func (bm *BaseModel) SetStableID(id ID) {
	bm.StableID = id
}

func (bm *BaseModel) GetModelStoreID() manifest.ID {
	return bm.ModelStoreID
}

func (bm *BaseModel) SetModelStoreID(id manifest.ID) {
	bm.ModelStoreID = id
}
