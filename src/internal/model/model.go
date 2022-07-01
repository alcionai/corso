package model

import "github.com/kopia/kopia/repo/manifest"

type ID string

type Model interface {
	GetStableID() ID
	// StableID is an identifier that can be used to link objects in ModelStore.
	// Once generated (during Put), it is guaranteed not to change.
	SetStableID(id ID)
	GetModelStoreID() manifest.ID
	// ModelStoreID is and opaque field in models. This field may change if the
	// model is updated. The field housing this should always have `omitempty`.
	SetModelStoreID(id manifest.ID)
}

// BaseModel defines required fields for models stored in ModelStore. Structs
// that wish to be stored should embed this struct.
type BaseModel struct {
	StableID ID `json:"stableID,omitempty"`
	// ModelStoreID is an opaque field.
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
