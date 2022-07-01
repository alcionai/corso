package kopia

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	ctesting "github.com/alcionai/corso/internal/testing"
)

type fooModel struct {
	BaseModel
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

	foo := &fooModel{Bar: uuid.NewString()}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			assert.Error(t, m.Put(ctx, BackupOpModel, test.tags, foo))
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
	noStableID.SetStableID("")
	noStableID.SetModelStoreID(manifest.ID(uuid.NewString()))

	noModelStoreID := &fooModel{Bar: uuid.NewString()}
	noModelStoreID.SetStableID(ID(uuid.NewString()))
	noModelStoreID.SetModelStoreID("")

	assert.Error(t, m.GetWithModelStoreID(ctx, "", nil))
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

			err := m.Put(ctx, test.t, nil, foo)
			test.check(t, err)

			if test.hasErr {
				return
			}

			require.NotEmpty(t, foo.GetModelStoreID())
			require.NotEmpty(t, foo.GetStableID())

			returned := &fooModel{}
			err = m.GetWithModelStoreID(ctx, foo.GetModelStoreID(), returned)
			require.NoError(t, err)
			assert.Equal(t, foo, returned)
		})
	}
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
