package kopia

import (
	"context"

	"github.com/google/uuid"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/pkg/errors"
)

const (
	stableIDKey       = "stableID"
	errNoModelStoreID = "model has no ModelStoreID"
	errBadTagKey      = "tag key overlaps with required key"
)

//go:generate stringer -type=modelType
const (
	UnknownModel = modelType(iota)
	BackupOpModel
	RestoreOpModel
	RestorePointModel
)

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

// ID of the manifest in kopia. This is not guaranteed to be stable.
type modelType int

func NewModelStore(kw *KopiaWrapper) *ModelStore {
	return &ModelStore{wrapper: kw}
}

// ModelStore must not be accessed after the given KopiaWrapper is closed.
type ModelStore struct {
	wrapper *KopiaWrapper
}

// tagsForModel creates a copy of tags and adds a tag for the model type to it.
// Returns an error if another tag has the same key as the model type or if a
// bad model type is give.
func tagsForModel(t modelType, tags map[string]string) (map[string]string, error) {
	if t == UnknownModel {
		return nil, errors.New("bad model type")
	}

	if _, ok := tags[manifest.TypeLabelKey]; ok {
		return nil, errors.New(errBadTagKey)
	}

	res := make(map[string]string, len(tags)+1)
	res[manifest.TypeLabelKey] = t.String()
	for k, v := range tags {
		res[k] = v
	}

	return res, nil
}

// tagsForModelWithID creates a copy of tags and adds tags for the model type
// StableID to it. Returns an error if another tag has the same key as the model
// type or if a bad model type is give.
func tagsForModelWithID(
	t modelType,
	id ID,
	tags map[string]string,
) (map[string]string, error) {
	if len(id) == 0 {
		return nil, errors.New("missing ID for model")
	}

	res, err := tagsForModel(t, tags)
	if err != nil {
		return nil, err
	}

	if _, ok := res[stableIDKey]; ok {
		return nil, errors.New(errBadTagKey)
	}

	res[stableIDKey] = string(id)

	return res, nil
}

// putInner contains logic for adding a model to the store. However, it does not
// issue a flush operation.
func putInner(
	ctx context.Context,
	w repo.RepositoryWriter,
	t modelType,
	tags map[string]string,
	m Model,
	create bool,
) error {
	// ModelStoreID does not need to be persisted in the model itself.
	m.SetModelStoreID("")
	if create {
		m.SetStableID(ID(uuid.NewString()))
	}

	tmpTags, err := tagsForModelWithID(t, m.GetStableID(), tags)
	if err != nil {
		// Will be wrapped at a higher layer.
		return err
	}

	id, err := w.PutManifest(ctx, tmpTags, m)
	if err != nil {
		// Will be wrapped at a higher layer.
		return err
	}

	m.SetModelStoreID(id)
	return nil
}

// Put adds a model of the given type to the persistent model store. Any tags
// given to this function can later be used to help lookup the model.
func (ms *ModelStore) Put(
	ctx context.Context,
	t modelType,
	tags map[string]string,
	m Model,
) error {
	err := repo.WriteSession(
		ctx,
		ms.wrapper.rep,
		repo.WriteSessionOptions{Purpose: "ModelStorePut"},
		func(innerCtx context.Context, w repo.RepositoryWriter) error {
			err := putInner(innerCtx, w, t, tags, m, true)
			if err != nil {
				return err
			}

			return nil
		},
	)

	return errors.Wrap(err, "putting model")
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

// GetWithModelStoreID deserializes the model with the given ModelStoreID into
// data. Returns github.com/kopia/kopia/repo/manifest.ErrNotFound if no model
// was found.
func (ms *ModelStore) GetWithModelStoreID(ctx context.Context, id manifest.ID, data Model) error {
	if len(id) == 0 {
		return errors.New(errNoModelStoreID)
	}

	_, err := ms.wrapper.rep.GetManifest(ctx, id, data)
	// TODO(ashmrtnz): Should probably return some recognizable, non-kopia error
	// if not found. That way kopia doesn't need to be imported to higher layers.
	if err != nil {
		return errors.Wrap(err, "getting model data")
	}

	data.SetModelStoreID(id)
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
