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

func stripHiddenTags(tags map[string]string) {
	delete(tags, stableIDKey)
	delete(tags, manifest.TypeLabelKey)
}

func baseModelFromMetadata(m *manifest.EntryMetadata) (*model.BaseModel, error) {
	id, ok := m.Labels[stableIDKey]
	if !ok {
		return nil, errors.WithStack(errNoStableID)
	}

	res := &model.BaseModel{
		ModelStoreID: m.ID,
		StableID:     model.ID(id),
		Tags:         m.Labels,
	}

	stripHiddenTags(res.Tags)
	return res, nil
}

// GetIDsForType returns metadata for all models that match the given type and
// have the given tags. Returned IDs can be used in subsequent calls to Get,
// Update, or Delete.
func (ms *ModelStore) GetIDsForType(
	ctx context.Context,
	t modelType,
	tags map[string]string,
) ([]*model.BaseModel, error) {
	if _, ok := tags[stableIDKey]; ok {
		return nil, errors.WithStack(errBadTagKey)
	}

	tmpTags, err := tagsForModel(t, tags)
	if err != nil {
		return nil, errors.Wrap(err, "getting model metadata")
	}

	metadata, err := ms.wrapper.rep.FindManifests(ctx, tmpTags)
	if err != nil {
		return nil, errors.Wrap(err, "getting model metadata")
	}

	res := make([]*model.BaseModel, 0, len(metadata))
	for _, m := range metadata {
		bm, err := baseModelFromMetadata(m)
		if err != nil {
			return nil, errors.Wrap(err, "parsing model metadata")
		}

		res = append(res, bm)
	}

	return res, nil
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

// Get deserializes the model with the given StableID into data. Returns
// github.com/kopia/kopia/repo/manifest.ErrNotFound if no model was found.
func (ms *ModelStore) Get(
	ctx context.Context,
	id model.ID,
	data model.Model,
) error {
	modelID, err := ms.getModelStoreID(ctx, id)
	if err != nil {
		return err
	}

	return ms.GetWithModelStoreID(ctx, modelID, data)
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
	stripHiddenTags(base.Tags)
	base.ModelStoreID = id
	return nil
}

// checkPrevModelVersion compares the modelType and ModelStoreID in this model
// to model(s) previously stored in ModelStore that have the same StableID.
// Returns an error if no models or more than one model has the same StableID or
// the modelType or ModelStoreID differ between the stored model and the given
// model.
func (ms *ModelStore) checkPrevModelVersion(
	ctx context.Context,
	t modelType,
	b *model.BaseModel,
) error {
	id, err := ms.getModelStoreID(ctx, b.StableID)
	if err != nil {
		return err
	}

	// We actually got something back during our lookup.
	meta, err := ms.wrapper.rep.GetManifest(ctx, id, nil)
	if err != nil {
		return errors.Wrap(err, "getting previous model version")
	}

	if meta.ID != b.ModelStoreID {
		return errors.New("updated model has different ModelStoreID")
	}
	if meta.Labels[manifest.TypeLabelKey] != t.String() {
		return errors.New("updated model has different model type")
	}

	return nil
}

// Update adds the new version of the model with the given StableID to the model
// store and deletes the version of the model with old ModelStoreID if the old
// and new ModelStoreIDs do not match. Returns an error if another model has
// the same StableID but a different modelType or ModelStoreID or there is no
// previous version of the model. If an error occurs no visible changes will be
// made to the stored model.
func (ms *ModelStore) Update(
	ctx context.Context,
	t modelType,
	m model.Model,
) error {
	base := m.Base()
	if len(base.ModelStoreID) == 0 {
		return errors.WithStack(errNoModelStoreID)
	}

	// TODO(ashmrtnz): Can remove if bottleneck.
	if err := ms.checkPrevModelVersion(ctx, t, base); err != nil {
		return err
	}

	err := repo.WriteSession(
		ctx,
		ms.wrapper.rep,
		repo.WriteSessionOptions{Purpose: "ModelStoreUpdate"},
		func(innerCtx context.Context, w repo.RepositoryWriter) (innerErr error) {
			oldID := base.ModelStoreID

			defer func() {
				if innerErr != nil {
					// Restore the old ID if we failed.
					base.ModelStoreID = oldID
				}
			}()

			if innerErr = putInner(innerCtx, w, t, m, false); innerErr != nil {
				return innerErr
			}

			// If we fail at this point no changes will be made to the manifest store
			// in kopia, making it appear like nothing ever happened. At worst some
			// orphaned content blobs may be uploaded, but they should be garbage
			// collected the next time kopia maintenance is run.
			if oldID != base.ModelStoreID {
				innerErr = w.DeleteManifest(innerCtx, oldID)
			}

			return innerErr
		},
	)
	if err != nil {
		return errors.Wrap(err, "updating model")
	}

	return nil
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

// DeletWithModelStoreID deletes the model with the given ModelStoreID from the
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
