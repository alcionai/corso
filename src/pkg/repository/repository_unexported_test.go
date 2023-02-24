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
	tester.Suite
}

func TestRepositoryModelSuite(t *testing.T) {
	suite.Run(t, &RepositoryModelSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs}),
	})
}

func (suite *RepositoryModelSuite) TestWriteGetModel() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t        = suite.T()
		s        = tester.NewPrefixedS3Storage(t)
		kopiaRef = kopia.NewConn(s)
	)

	require.NoError(t, kopiaRef.Initialize(ctx))
	require.NoError(t, kopiaRef.Connect(ctx))

	defer kopiaRef.Close(ctx)

	ms, err := kopia.NewModelStore(kopiaRef)
	require.NoError(t, err)

	defer ms.Close(ctx)

	require.NoError(t, newRepoModel(ctx, ms, "fnords"))

	got, err := getRepoModel(ctx, ms)
	require.NoError(t, err)
	assert.Equal(t, "fnords", string(got.ID))
}
