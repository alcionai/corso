package kopia

import (
	"context"
	"strconv"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/kopia/kopia/repo"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
)

const (
	stableIDKey        = "stableID"
	modelVersionKey    = "storeVersion"
	globalModelVersion = 1
)

var (
	errNoModelStoreID     = errors.New("model has no ModelStoreID")
	errNoStableID         = errors.New("model has no StableID")
	errBadTagKey          = errors.New("tag key overlaps with required key")
	errModelTypeMismatch  = errors.New("model type doesn't match request")
	errUnrecognizedSchema = errors.New("unrecognized model schema")
)

func NewModelStore(c *conn) (*ModelStore, error) {
	if err := c.wrap(); err != nil {
		return nil, errors.Wrap(err, "creating ModelStore")
	}

	return &ModelStore{c: c, modelVersion: globalModelVersion}, nil
}

// ModelStore must not be accessed after the given KopiaWrapper is closed.
type ModelStore struct {
	c *conn
	// Stash a reference here so testing can easily change it.
	modelVersion int
}

func (ms *ModelStore) Close(ctx context.Context) error {
	if ms.c == nil {
		return nil
	}

	err := ms.c.Close(ctx)
	ms.c = nil

	return errors.Wrap(err, "closing ModelStore")
}

// tagsForModel creates a copy of tags and adds a tag for the model schema to it.
// Returns an error if another tag has the same key as the model schema or if a
// bad model type is given.
func tagsForModel(s model.Schema, tags map[string]string) (map[string]string, error) {
	if _, ok := tags[manifest.TypeLabelKey]; ok {
		return nil, clues.Stack(errBadTagKey)
	}

	res := make(map[string]string, len(tags)+1)
	res[manifest.TypeLabelKey] = s.String()

	maps.Copy(res, tags)

	return res, nil
}

// tagsForModelWithID creates a copy of tags and adds tags for the model type
// StableID to it. Returns an error if another tag has the same key as the model
// type or if a bad model type is given.
func tagsForModelWithID(
	s model.Schema,
	id model.StableID,
	version int,
	tags map[string]string,
) (map[string]string, error) {
	if !s.Valid() {
		return nil, clues.Stack(errUnrecognizedSchema)
	}

	if len(id) == 0 {
		return nil, clues.Stack(errNoStableID)
	}

	res, err := tagsForModel(s, tags)
	if err != nil {
		return nil, err
	}

	if _, ok := res[stableIDKey]; ok {
		return nil, clues.Stack(errBadTagKey)
	}

	res[stableIDKey] = string(id)

	if _, ok := res[modelVersionKey]; ok {
		return nil, clues.Stack(errBadTagKey)
	}

	res[modelVersionKey] = strconv.Itoa(version)

	return res, nil
}

// putInner contains logic for adding a model to the store. However, it does not
// issue a flush operation.
func putInner(
	ctx context.Context,
	w repo.RepositoryWriter,
	s model.Schema,
	m model.Model,
	create bool,
) error {
	if !s.Valid() {
		return clues.Stack(errUnrecognizedSchema).WithClues(ctx)
	}

	base := m.Base()
	if create && len(base.ID) == 0 {
		base.ID = model.StableID(uuid.NewString())
	}

	tmpTags, err := tagsForModelWithID(s, base.ID, base.Version, base.Tags)
	if err != nil {
		// Will be wrapped at a higher layer.
		return clues.Stack(err).WithClues(ctx)
	}

	id, err := w.PutManifest(ctx, tmpTags, m)
	if err != nil {
		// Will be wrapped at a higher layer.
		return clues.Stack(err).WithClues(ctx)
	}

	base.ModelStoreID = id

	return nil
}

// Put adds a model of the given type to the persistent model store. Any tags
// given to this function can later be used to help lookup the model.
func (ms *ModelStore) Put(
	ctx context.Context,
	s model.Schema,
	m model.Model,
) error {
	if !s.Valid() {
		return clues.Stack(errUnrecognizedSchema)
	}

	m.Base().Version = ms.modelVersion

	err := repo.WriteSession(
		ctx,
		ms.c,
		repo.WriteSessionOptions{Purpose: "ModelStorePut"},
		func(innerCtx context.Context, w repo.RepositoryWriter) error {
			err := putInner(innerCtx, w, s, m, true)
			if err != nil {
				return clues.Stack(err).WithClues(innerCtx)
			}

			return nil
		},
	)

	return clues.Wrap(err, "putting model").WithClues(ctx)
}

func stripHiddenTags(tags map[string]string) {
	delete(tags, stableIDKey)
	delete(tags, modelVersionKey)
	delete(tags, manifest.TypeLabelKey)
}

func (ms ModelStore) populateBaseModelFromMetadata(
	base *model.BaseModel,
	m *manifest.EntryMetadata,
) error {
	id, ok := m.Labels[stableIDKey]
	if !ok {
		return clues.Stack(errNoStableID)
	}

	v, err := strconv.Atoi(m.Labels[modelVersionKey])
	if err != nil {
		return errors.Wrap(err, "parsing model version")
	}

	if v != ms.modelVersion {
		return clues.Wrap(clues.New(m.Labels[modelVersionKey]), "bad model version")
	}

	base.ModelStoreID = m.ID
	base.ID = model.StableID(id)
	base.Version = v
	base.Tags = m.Labels

	stripHiddenTags(base.Tags)

	return nil
}

func (ms ModelStore) baseModelFromMetadata(
	m *manifest.EntryMetadata,
) (*model.BaseModel, error) {
	res := &model.BaseModel{}
	if err := ms.populateBaseModelFromMetadata(res, m); err != nil {
		return nil, clues.Stack(err).WithAll("metadata_id", m.ID, "metadata_modtime", m.ModTime)
	}

	return res, nil
}

// GetIDsForType returns metadata for all models that match the given type and
// have the given tags. Returned IDs can be used in subsequent calls to Get,
// Update, or Delete.
func (ms *ModelStore) GetIDsForType(
	ctx context.Context,
	s model.Schema,
	tags map[string]string,
) ([]*model.BaseModel, error) {
	if !s.Valid() {
		return nil, clues.Stack(errUnrecognizedSchema).WithClues(ctx)
	}

	if _, ok := tags[stableIDKey]; ok {
		return nil, clues.Stack(errBadTagKey).WithClues(ctx)
	}

	tmpTags, err := tagsForModel(s, tags)
	if err != nil {
		return nil, clues.Wrap(err, "getting model metadata").WithClues(ctx)
	}

	metadata, err := ms.c.FindManifests(ctx, tmpTags)
	if err != nil {
		return nil, clues.Wrap(err, "getting model metadata").WithClues(ctx)
	}

	res := make([]*model.BaseModel, 0, len(metadata))

	for _, m := range metadata {
		bm, err := ms.baseModelFromMetadata(m)
		if err != nil {
			return nil, clues.Wrap(err, "parsing model metadata").WithClues(ctx)
		}

		res = append(res, bm)
	}

	return res, nil
}

// getModelStoreID gets the ModelStoreID of the model with the given
// StableID. Returns an error if the given StableID is empty or more than
// one model has the same StableID.
func (ms *ModelStore) getModelStoreID(
	ctx context.Context,
	s model.Schema,
	id model.StableID,
) (manifest.ID, error) {
	if !s.Valid() {
		return "", clues.Stack(errUnrecognizedSchema).WithClues(ctx)
	}

	if len(id) == 0 {
		return "", clues.Stack(errNoStableID).WithClues(ctx)
	}

	tags := map[string]string{stableIDKey: string(id)}

	metadata, err := ms.c.FindManifests(ctx, tags)
	if err != nil {
		return "", clues.Wrap(err, "getting ModelStoreID").WithClues(ctx)
	}

	if len(metadata) == 0 {
		return "", clues.Wrap(data.ErrNotFound, "getting ModelStoreID").WithClues(ctx)
	}

	if len(metadata) != 1 {
		return "", clues.New("multiple models with same StableID").WithClues(ctx)
	}

	if metadata[0].Labels[manifest.TypeLabelKey] != s.String() {
		return "", clues.Stack(errModelTypeMismatch).WithClues(ctx)
	}

	return metadata[0].ID, nil
}

// Get deserializes the model with the given StableID into data.
// Returns an error if the persisted model has a different type
// than expected or if multiple models have the same StableID.
func (ms *ModelStore) Get(
	ctx context.Context,
	s model.Schema,
	id model.StableID,
	m model.Model,
) error {
	if !s.Valid() {
		return clues.Stack(errUnrecognizedSchema).WithClues(ctx)
	}

	modelID, err := ms.getModelStoreID(ctx, s, id)
	if err != nil {
		return err
	}

	return ms.GetWithModelStoreID(ctx, s, modelID, m)
}

// GetWithModelStoreID deserializes the model with the given ModelStoreID into
// data. Returns and error if the persisted model has a different type than
// expected.
func (ms *ModelStore) GetWithModelStoreID(
	ctx context.Context,
	s model.Schema,
	id manifest.ID,
	m model.Model,
) error {
	if !s.Valid() {
		return clues.Stack(errUnrecognizedSchema).WithClues(ctx)
	}

	if len(id) == 0 {
		return clues.Stack(errNoModelStoreID).WithClues(ctx)
	}

	metadata, err := ms.c.GetManifest(ctx, id, m)
	if err != nil {
		if errors.Is(err, manifest.ErrNotFound) {
			err = data.ErrNotFound
		}

		return clues.Wrap(err, "getting model data").WithClues(ctx)
	}

	mdlbl := metadata.Labels[manifest.TypeLabelKey]
	if mdlbl != s.String() {
		return clues.Stack(errModelTypeMismatch).
			WithClues(ctx).
			WithAll("expected_label", s, "got_label", mdlbl)
	}

	if err := ms.populateBaseModelFromMetadata(m.Base(), metadata); err != nil {
		return clues.Wrap(err, "getting model by ID").WithClues(ctx)
	}

	return nil
}

// checkPrevModelVersion compares the ModelType and ModelStoreID in this model
// to model(s) previously stored in ModelStore that have the same StableID.
// Returns an error if no models or more than one model has the same StableID or
// the ModelType or ModelStoreID differ between the stored model and the given
// model.
func (ms *ModelStore) checkPrevModelVersion(
	ctx context.Context,
	s model.Schema,
	b *model.BaseModel,
) error {
	if !s.Valid() {
		return clues.Stack(errUnrecognizedSchema).WithClues(ctx)
	}

	id, err := ms.getModelStoreID(ctx, s, b.ID)
	if err != nil {
		return clues.Stack(err).WithClues(ctx)
	}

	// We actually got something back during our lookup.
	meta, err := ms.c.GetManifest(ctx, id, nil)
	if err != nil {
		return clues.Wrap(err, "getting previous model version").WithClues(ctx)
	}

	if meta.ID != b.ModelStoreID {
		return clues.New("updated model has different ModelStoreID").
			WithClues(ctx).
			WithAll("expected_id", meta.ID, "model_store_id", b.ModelStoreID)
	}

	mdlbl := meta.Labels[manifest.TypeLabelKey]
	if mdlbl != s.String() {
		return clues.New("updated model has different model type").
			WithClues(ctx).
			WithAll("expected_label", s, "got_label", mdlbl)
	}

	return nil
}

// Update adds the new version of the model with the given StableID to the model
// store and deletes the version of the model with old ModelStoreID if the old
// and new ModelStoreIDs do not match. Returns an error if another model has
// the same StableID but a different ModelType or ModelStoreID or there is no
// previous version of the model. If an error occurs no visible changes will be
// made to the stored model.
func (ms *ModelStore) Update(
	ctx context.Context,
	s model.Schema,
	m model.Model,
) error {
	if !s.Valid() {
		return clues.Stack(errUnrecognizedSchema).WithClues(ctx)
	}

	base := m.Base()
	if len(base.ModelStoreID) == 0 {
		return clues.Stack(errNoModelStoreID).WithClues(ctx)
	}

	base.Version = ms.modelVersion

	// TODO(ashmrtnz): Can remove if bottleneck.
	if err := ms.checkPrevModelVersion(ctx, s, base); err != nil {
		return err
	}

	err := repo.WriteSession(
		ctx,
		ms.c,
		repo.WriteSessionOptions{Purpose: "ModelStoreUpdate"},
		func(innerCtx context.Context, w repo.RepositoryWriter) error {
			var (
				innerErr error
				oldID    = base.ModelStoreID
			)

			defer func() {
				if innerErr != nil {
					// Restore the old ID if we failed.
					base.ModelStoreID = oldID
				}
			}()

			if innerErr = putInner(innerCtx, w, s, m, false); innerErr != nil {
				return innerErr
			}

			// if equal, everything worked out fine.
			// if not, we handle the cleanup below.
			if oldID == base.ModelStoreID {
				return nil
			}

			// If we fail at this point no changes will be made to the manifest store
			// in kopia, making it appear like nothing ever happened. At worst some
			// orphaned content blobs may be uploaded, but they should be garbage
			// collected the next time kopia maintenance is run.
			innerErr = w.DeleteManifest(innerCtx, oldID)
			if innerErr != nil {
				return clues.Stack(innerErr).WithClues(ctx)
			}

			return nil
		},
	)
	if err != nil {
		return clues.Wrap(err, "updating model").WithClues(ctx)
	}

	return nil
}

// Delete deletes the model with the given StableID. Turns into a noop if id is
// not empty but the model does not exist. Returns an error if multiple models
// have the same StableID.
func (ms *ModelStore) Delete(ctx context.Context, s model.Schema, id model.StableID) error {
	if !s.Valid() {
		return clues.Stack(errUnrecognizedSchema).WithClues(ctx)
	}

	latest, err := ms.getModelStoreID(ctx, s, id)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
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
		return clues.Stack(errNoModelStoreID).WithClues(ctx)
	}

	opts := repo.WriteSessionOptions{Purpose: "ModelStoreDelete"}
	cb := func(innerCtx context.Context, w repo.RepositoryWriter) error {
		return w.DeleteManifest(innerCtx, id)
	}

	if err := repo.WriteSession(ctx, ms.c, opts, cb); err != nil {
		return clues.Wrap(err, "deleting model").WithClues(ctx)
	}

	return nil
}
