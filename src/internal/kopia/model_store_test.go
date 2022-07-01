package kopia

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/model"
	ctesting "github.com/alcionai/corso/internal/testing"
)

type fooModel struct {
	model.BaseModel
	Bar string
}

func getModelStore(t *testing.T, ctx context.Context) *ModelStore {
	kw, err := openKopiaRepo(t, ctx)
	require.NoError(t, err)

	return NewModelStore(kw)
}

// ---------------
// integration tests that use kopia
// ---------------
type ModelStoreIntegrationSuite struct {
	suite.Suite
}

func TestModelStoreIntegrationSuite(t *testing.T) {
	if err := ctesting.RunOnAny(
		ctesting.CorsoCITests,
		ctesting.CorsoModelStoreTests,
	); err != nil {
		t.Skip()
	}

	suite.Run(t, new(ModelStoreIntegrationSuite))
}

func (suite *ModelStoreIntegrationSuite) SetupSuite() {
	_, err := ctesting.GetRequiredEnvVars(ctesting.AWSStorageCredEnvs...)
	require.NoError(suite.T(), err)
}

func (suite *ModelStoreIntegrationSuite) TestBadTagsErrors() {
	table := []struct {
		name string
		tags map[string]string
	}{
		{
			name: "StableIDTag",
			tags: map[string]string{
				stableIDKey: "foo",
			},
		},
		{
			name: "manifestTypeTag",
			tags: map[string]string{
				manifest.TypeLabelKey: "foo",
			},
		},
	}

	ctx := context.Background()
	t := suite.T()

	m := getModelStore(t, ctx)
	defer func() {
		assert.NoError(t, m.wrapper.Close(ctx))
	}()

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			foo := &fooModel{Bar: uuid.NewString()}
			foo.Tags = test.tags

			assert.Error(t, m.Put(ctx, BackupOpModel, foo))
		})
	}
}

func (suite *ModelStoreIntegrationSuite) TestNoIDsErrors() {
	ctx := context.Background()
	t := suite.T()

	m := getModelStore(t, ctx)
	defer func() {
		assert.NoError(t, m.wrapper.Close(ctx))
	}()

	noStableID := &fooModel{Bar: uuid.NewString()}
	noStableID.StableID = ""
	noStableID.ModelStoreID = manifest.ID(uuid.NewString())

	noModelStoreID := &fooModel{Bar: uuid.NewString()}
	noModelStoreID.StableID = model.ID(uuid.NewString())
	noModelStoreID.ModelStoreID = ""

	assert.Error(t, m.GetWithModelStoreID(ctx, "", nil))

	assert.Error(t, m.Delete(ctx, ""))
	assert.Error(t, m.DeleteWithModelStoreID(ctx, ""))
}

func (suite *ModelStoreIntegrationSuite) TestPutGet() {
	table := []struct {
		t      modelType
		check  require.ErrorAssertionFunc
		hasErr bool
	}{
		{
			t:      UnknownModel,
			check:  require.Error,
			hasErr: true,
		},
		{
			t:      BackupOpModel,
			check:  require.NoError,
			hasErr: false,
		},
		{
			t:      RestoreOpModel,
			check:  require.NoError,
			hasErr: false,
		},
		{
			t:      RestorePointModel,
			check:  require.NoError,
			hasErr: false,
		},
	}

	ctx := context.Background()
	t := suite.T()

	m := getModelStore(t, ctx)
	defer func() {
		assert.NoError(t, m.wrapper.Close(ctx))
	}()

	for _, test := range table {
		suite.T().Run(test.t.String(), func(t *testing.T) {
			foo := &fooModel{Bar: uuid.NewString()}
			// Avoid some silly test errors from comparing nil to empty map.
			foo.Tags = map[string]string{}

			err := m.Put(ctx, test.t, foo)
			test.check(t, err)

			if test.hasErr {
				return
			}

			require.NotEmpty(t, foo.ModelStoreID)
			require.NotEmpty(t, foo.StableID)

			returned := &fooModel{}
			err = m.GetWithModelStoreID(ctx, foo.ModelStoreID, returned)
			require.NoError(t, err)
			assert.Equal(t, foo, returned)
		})
	}
}

func (suite *ModelStoreIntegrationSuite) TestPutGet_WithTags() {
	ctx := context.Background()
	t := suite.T()

	m := getModelStore(t, ctx)
	defer func() {
		assert.NoError(t, m.wrapper.Close(ctx))
	}()

	foo := &fooModel{Bar: uuid.NewString()}
	foo.Tags = map[string]string{
		"bar": "baz",
	}

	require.NoError(t, m.Put(ctx, BackupOpModel, foo))

	require.NotEmpty(t, foo.ModelStoreID)
	require.NotEmpty(t, foo.StableID)

	returned := &fooModel{}
	err := m.GetWithModelStoreID(ctx, foo.ModelStoreID, returned)
	require.NoError(t, err)
	assert.Equal(t, foo, returned)
}

func (suite *ModelStoreIntegrationSuite) TestGet_NotFoundErrors() {
	ctx := context.Background()
	t := suite.T()

	m := getModelStore(t, ctx)
	defer func() {
		assert.NoError(t, m.wrapper.Close(ctx))
	}()

	assert.ErrorIs(t, m.GetWithModelStoreID(ctx, "baz", nil), manifest.ErrNotFound)
}

func (suite *ModelStoreIntegrationSuite) TestPutDelete() {
	ctx := context.Background()
	t := suite.T()

	m := getModelStore(t, ctx)
	defer func() {
		assert.NoError(t, m.wrapper.Close(ctx))
	}()

	foo := &fooModel{Bar: uuid.NewString()}

	require.NoError(t, m.Put(ctx, BackupOpModel, foo))

	require.NoError(t, m.Delete(ctx, foo.StableID))

	returned := &fooModel{}
	err := m.GetWithModelStoreID(ctx, foo.ModelStoreID, returned)
	assert.ErrorIs(t, err, manifest.ErrNotFound)
}

func (suite *ModelStoreIntegrationSuite) TestPutDelete_BadIDsNoop() {
	ctx := context.Background()
	t := suite.T()

	m := getModelStore(t, ctx)
	defer func() {
		assert.NoError(t, m.wrapper.Close(ctx))
	}()

	assert.NoError(t, m.Delete(ctx, "foo"))
	assert.NoError(t, m.DeleteWithModelStoreID(ctx, "foo"))
}
