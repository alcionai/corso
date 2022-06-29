package kopia

import (
	"context"
)

// ID of the manifest in kopia. This is not guaranteed to be stable.
type ID string
type modelType int

//go:generate stringer -type=modelType
const (
	UnknownModel = modelType(iota)
	BackupOpModel
	RestoreOpModel
	RestorePointModel

	errStoreFlush = "flushing manifest store"
)

type ModelStore struct{}

// Put adds a model of the given type to the persistent model store. Any tags
// given to this function can later be used to help lookup the model. The
// returned ID can be used for subsequent Get, Update, or Delete calls.
func (ms *ModelStore) Put(
	ctx context.Context,
	t modelType,
	tags map[string]string,
	m any,
) (ID, error) {
	return "", nil
}

// GetIDsForType returns all IDs for models that match the given type and have
// the given tags. Returned IDs can be used in subsequent calls to Get, Update,
// or Delete.
func (ms *ModelStore) GetIDsForType(
	ctx context.Context,
	t modelType,
	tags map[string]string,
) ([]ID, error) {
	return nil, nil
}

// Get deserializes the model with the given ID into data.
func (ms *ModelStore) Get(ctx context.Context, id ID, data any) error {
	return nil
}

// Update adds the new version of the model to the model store and deletes the
// version of the model with oldID if the old and new IDs do not match. The new
// ID of the model is returned.
func (ms *ModelStore) Update(
	ctx context.Context,
	t modelType,
	oldID ID,
	tags map[string]string,
	m any,
) (ID, error) {
	return "", nil
}

// Delete deletes the model with the given ID from the model store.
func (ms *ModelStore) Delete(ctx context.Context, id ID) error {
	return nil
}
