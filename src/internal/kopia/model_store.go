package kopia

import (
	"context"

	"github.com/google/uuid"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/internal/model"
)

const stableIDKey = "stableID"

var (
	errNoModelStoreID = errors.New("model has no ModelStoreID")
	errNoStableID     = errors.New("model has no StableID")
	errBadTagKey      = errors.New("tag key overlaps with required key")
)

type modelType int

//go:generate stringer -type=modelType
const (
	UnknownModel = modelType(iota)
	BackupOpModel
	RestoreOpModel
	RestorePointModel
)

func NewModelStore(kw *KopiaWrapper) *ModelStore {
	return &ModelStore{wrapper: kw}
}

// ModelStore must not be accessed after the given KopiaWrapper is closed.
type ModelStore struct {
	wrapper *KopiaWrapper
}

// tagsForModel creates a copy of tags and adds a tag for the model type to it.
// Returns an error if another tag has the same key as the model type or if a
// bad model type is given.
func tagsForModel(t modelType, tags map[string]string) (map[string]string, error) {
	if t == UnknownModel {
		return nil, errors.New("bad model type")
	}

	if _, ok := tags[manifest.TypeLabelKey]; ok {
		return nil, errors.WithStack(errBadTagKey)
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
// type or if a bad model type is given.
func tagsForModelWithID(
	t modelType,
	id model.ID,
	tags map[string]string,
) (map[string]string, error) {
	if len(id) == 0 {
		return nil, errors.WithStack(errNoStableID)
	}

	res, err := tagsForModel(t, tags)
	if err != nil {
		return nil, err
	}

	if _, ok := res[stableIDKey]; ok {
		return nil, errors.WithStack(errBadTagKey)
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
	m model.Model,
	create bool,
) error {
	base := m.Base()
	if create {
		base.StableID = model.ID(uuid.NewString())
	}

	tmpTags, err := tagsForModelWithID(t, base.StableID, base.Tags)
	if err != nil {
		// Will be wrapped at a higher layer.
		return err
	}

	id, err := w.PutManifest(ctx, tmpTags, m)
	if err != nil {
		// Will be wrapped at a higher layer.
		return err
	}

	base.ModelStoreID = id
	return nil
}

// Put adds a model of the given type to the persistent model store. Any tags
// given to this function can later be used to help lookup the model.
func (ms *ModelStore) Put(
	ctx context.Context,
	t modelType,
	m model.Model,
) error {
	err := repo.WriteSession(
		ctx,
		ms.wrapper.rep,
		repo.WriteSessionOptions{Purpose: "ModelStorePut"},
		func(innerCtx context.Context, w repo.RepositoryWriter) error {
			err := putInner(innerCtx, w, t, m, true)
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
) ([]model.ID, error) {
	return nil, nil
}

// getModelStoreID gets the ModelStoreID of the model with the given
// StableID. Returns github.com/kopia/kopia/repo/manifest.ErrNotFound if no
// model was found. Returns an error if the given StableID is empty or more than
// one model has the same StableID.
func (ms *ModelStore) getModelStoreID(ctx context.Context, id model.ID) (manifest.ID, error) {
	if len(id) == 0 {
		return "", errors.WithStack(errNoStableID)
	}

	tags := map[string]string{stableIDKey: string(id)}
	metadata, err := ms.wrapper.rep.FindManifests(ctx, tags)
	if err != nil {
		return "", errors.Wrap(err, "getting ModelStoreID")
	}

	if len(metadata) == 0 {
		return "", errors.Wrap(manifest.ErrNotFound, "getting ModelStoreID")
	}
	if len(metadata) != 1 {
		return "", errors.New("multiple models with same StableID")
	}

	return metadata[0].ID, nil
}

// Get deserializes the model with the given ID into data.
func (ms *ModelStore) Get(ctx context.Context, id model.ID, data any) error {
	return nil
}

// GetWithModelStoreID deserializes the model with the given ModelStoreID into
// data. Returns github.com/kopia/kopia/repo/manifest.ErrNotFound if no model
// was found.
func (ms *ModelStore) GetWithModelStoreID(
	ctx context.Context,
	id manifest.ID,
	data model.Model,
) error {
	if len(id) == 0 {
		return errors.WithStack(errNoModelStoreID)
	}

	metadata, err := ms.wrapper.rep.GetManifest(ctx, id, data)
	// TODO(ashmrtnz): Should probably return some recognizable, non-kopia error
	// if not found. That way kopia doesn't need to be imported to higher layers.
	if err != nil {
		return errors.Wrap(err, "getting model data")
	}

	base := data.Base()
	base.Tags = metadata.Labels
	// Hide the fact that StableID and modelType are just a tag from the user.
	delete(base.Tags, stableIDKey)
	delete(base.Tags, manifest.TypeLabelKey)
	base.ModelStoreID = id
	return nil
}

// Update adds the new version of the model to the model store and deletes the
// version of the model with oldID if the old and new IDs do not match. The new
// ID of the model is returned.
func (ms *ModelStore) Update(
	ctx context.Context,
	t modelType,
	oldID model.ID,
	tags map[string]string,
	m any,
) (model.ID, error) {
	return "", nil
}

// Delete deletes the model with the given StableID. Turns into a noop if id is
// not empty but the model does not exist.
func (ms *ModelStore) Delete(ctx context.Context, id model.ID) error {
	latest, err := ms.getModelStoreID(ctx, id)
	if err != nil {
		if errors.Is(err, manifest.ErrNotFound) {
			return nil
		}

		return err
	}

	return ms.DeleteWithModelStoreID(ctx, latest)
}

// DeletWithModelStoreIDe deletes the model with the given ModelStoreID from the
// model store. Turns into a noop if id is not empty but the model does not
// exist.
func (ms *ModelStore) DeleteWithModelStoreID(ctx context.Context, id manifest.ID) error {
	if len(id) == 0 {
		return errors.WithStack(errNoModelStoreID)
	}

	err := repo.WriteSession(
		ctx,
		ms.wrapper.rep,
		repo.WriteSessionOptions{Purpose: "ModelStoreDelete"},
		func(innerCtx context.Context, w repo.RepositoryWriter) error {
			return w.DeleteManifest(innerCtx, id)
		},
	)

	return errors.Wrap(err, "deleting model")
}
