package model

import (
	"time"

	"github.com/kopia/kopia/repo/manifest"
)

type (
	// StableID is used by BaseModel.ID to uniquely identify objects
	// stored in the modelStore.
	StableID string
	Schema   int
)

func (id StableID) String() string {
	return string(id)
}

// Schema constants denote the type of model stored. The integer values of the
// constants can be changed without issue, but the string values should remain
// the same. If the string values are changed, additional code will be needed to
// transform from the old value to the new value.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=Schema
const (
	UnknownSchema       Schema = 0
	BackupOpSchema      Schema = 1
	RestoreOpSchema     Schema = 2
	BackupSchema        Schema = 3
	BackupDetailsSchema Schema = 4
	RepositorySchema    Schema = 5
)

// common tags for filtering
const (
	ServiceTag = "service"
	// BackupTypeTag is the key used to store the resulting type of backup from a
	// backup operation. The type of the backup is determined by a combination of
	// input options and if errors were encountered during the backup. When making
	// an incremental backup, previous backups' types are inspected to determine
	// if they can be used as a base.
	//
	// The backup type associated with this key should only be used for
	// determining if a backup is a valid base. Once the bases for a backup
	// operation have been found, structs like kopia.BackupBases should be used to
	// track the type of each base.
	BackupTypeTag = "backup-type"
	// AssistBackup denotes that this backup should only be used for kopia
	// assisted incrementals since it doesn't contain the complete set of data
	// being backed up.
	//
	// See comment on BackupTypeTag for more information.
	AssistBackup = "assist-backup"
	// MergeBackup denotes that this backup can be used as a merge base during an
	// incremental backup. It contains a complete snapshot of the data in the
	// external service. Merge bases can also be used as assist bases during an
	// incremental backup or demoted to being only an assist base.
	//
	// See comment on BackupTypeTag for more information.
	MergeBackup = "merge-backup"
	// PreviewBackup denotes that this backup contains a subset of information for
	// the protected resource. PreviewBackups are used to demonstrate value but
	// are not safe to use as merge bases for incremental backups. It's possible
	// they could be used as assist bases since the only difference from a regular
	// backup is the amount of data they contain.
	//
	// See comment on BackupTypeTag for more information.
	PreviewBackup = "preview-backup"
)

// Valid returns true if the ModelType value fits within the const range.
func (mt Schema) Valid() bool {
	return mt > 0 && mt < RepositorySchema+1
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
	// ModelVersion is a version number that can help track changes across models.
	// TODO(ashmrtn): Reference version control documentation.
	ModelVersion int `json:"-"`
	// Tags associated with this model in the store to facilitate lookup. Tags in
	// the struct are not serialized directly into the stored model, but are part
	// of the metadata for the model.
	Tags    map[string]string `json:"-"`
	ModTime time.Time         `json:"-"`
}

func (bm *BaseModel) Base() *BaseModel {
	return bm
}

// GetID returns the baseModel.ID as a string rather than a model.StableID.
func (bm *BaseModel) GetID() string {
	return string(bm.ID)
}
