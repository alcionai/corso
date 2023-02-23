package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/aw"
)

type RepositoryModelSuite struct {
	suite.Suite
}

func TestRepositoryModelSuite(t *testing.T) {
	tester.RunOnAny(
		t,
		tester.CorsoCITests,
		tester.CorsoRepositoryTests)

	suite.Run(t, new(RepositoryModelSuite))
}

// ensure all required env values are populated
func (suite *RepositoryModelSuite) SetupSuite() {
	tester.MustGetEnvSets(suite.T(), tester.AWSStorageCredEnvs, tester.M365AcctCredEnvs)
}

func (suite *RepositoryModelSuite) TestWriteGetModel() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t        = suite.T()
		s        = tester.NewPrefixedS3Storage(t)
		kopiaRef = kopia.NewConn(s)
	)

	aw.MustNoErr(t, kopiaRef.Initialize(ctx))
	aw.MustNoErr(t, kopiaRef.Connect(ctx))

	defer kopiaRef.Close(ctx)

	ms, err := kopia.NewModelStore(kopiaRef)
	aw.MustNoErr(t, err)

	defer ms.Close(ctx)

	aw.MustNoErr(t, newRepoModel(ctx, ms, "fnords"))

	got, err := getRepoModel(ctx, ms)
	aw.MustNoErr(t, err)
	assert.Equal(t, "fnords", string(got.ID))
}
