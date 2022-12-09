package model

import (
	"github.com/kopia/kopia/repo/manifest"
)

type (
	// StableID is used by BaseModel.ID to uniquely identify objects
	// stored in the modelStore.
	StableID string
	Schema   int
)

// Schema constants denote the type of model stored. The integer values of the
// constants can be changed without issue, but the string values should remain
// the same. If the string values are changed, additional code will be needed to
// transform from the old value to the new value.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=Schema
const (
	UnknownSchema = Schema(iota)
	BackupOpSchema
	RestoreOpSchema
	BackupSchema
	BackupDetailsSchema
)

// common tags for filtering
const (
	ServiceTag = "service"
)

// Valid returns true if the ModelType value fits within the iota range.
func (mt Schema) Valid() bool {
	return mt > 0 && mt < BackupDetailsSchema+1
}

type Model interface {
	// Returns a handle to the BaseModel for this model.
	Base() *BaseModel
}

// BaseModel defines required fields for models stored in ModelStore. Structs
// that wish to be stored should embed this struct. This struct also represents
// the common metadata the ModelStore will fill out/use.
type BaseModel struct {
	// ID is an identifier that other objects can use to refer to this
	// object in the ModelStore.
	// Once generated (during Put), it is guaranteed not to change. This field
	// should be treated as read-only by users.
	ID StableID `json:"ID,omitempty"`
	// ModelStoreID is an internal ID for the model in the store. If present it
	// can be used for efficient lookups, but should not be used by other models
	// to refer to this one. This field may change if the model is updated. This
	// field should be treated as read-only by users.
	ModelStoreID manifest.ID `json:"-"`
	// Version is a version number that can help track changes across models.
	// TODO(ashmrtn): Reference version control documentation.
	Version int `json:"-"`
	// Tags associated with this model in the store to facilitate lookup. Tags in
	// the struct are not serialized directly into the stored model, but are part
	// of the metadata for the model.
	Tags map[string]string `json:"-"`
}

func (bm *BaseModel) Base() *BaseModel {
	return bm
}

// GetID returns the baseModel.ID as a string rather than a model.StableID.
func (bm *BaseModel) GetID() string {
	return string(bm.ID)
}
