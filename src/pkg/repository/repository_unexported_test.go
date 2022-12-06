package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/tester"
)

type RepositoryModelSuite struct {
	suite.Suite
}

func TestRepositoryModelSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoRepositoryTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(RepositoryModelSuite))
}

// ensure all required env values are populated
func (suite *RepositoryModelSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs)
	require.NoError(suite.T(), err)
}

func (suite *RepositoryModelSuite) TestNewRepoModel() {
	t := suite.T()

	rm := newRepoModel("smarf")
	assert.NotEmpty(t, rm)
	assert.Equal(t, repositoryID, rm.ID)
	assert.Equal(t, "smarf", rm.RepoID())
}

func (suite *RepositoryModelSuite) TestWriteGetModel() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t        = suite.T()
		s        = tester.NewPrefixedS3Storage(t)
		kopiaRef = kopia.NewConn(s)
		rm       = newRepoModel("fnords")
	)

	require.NoError(t, kopiaRef.Connect(ctx))
	defer kopiaRef.Close(ctx)

	ms, err := kopia.NewModelStore(kopiaRef)
	require.NoError(t, err)
	require.NoError(t, rm.write(ctx, ms))

	got, err := getRepoModel(ctx, ms)
	require.NoError(t, err)
	assert.Equal(t, "fnords", got.RepoID())
}
