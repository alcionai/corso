package operations

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/common/ptr"
	strTD "github.com/alcionai/canario/src/internal/common/str/testdata"
	evmock "github.com/alcionai/canario/src/internal/events/mock"
	"github.com/alcionai/canario/src/internal/kopia"
	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/control/repository"
	storeTD "github.com/alcionai/canario/src/pkg/storage/testdata"
)

type PersistentConfigOpIntegrationSuite struct {
	tester.Suite
}

func TestPersistentConfigOpIntegrationSuite(t *testing.T) {
	suite.Run(t, &PersistentConfigOpIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *PersistentConfigOpIntegrationSuite) TestRepoPersistentConfig() {
	var (
		t = suite.T()
		// need to initialize the repository before we can test connecting to it.
		st           = storeTD.NewPrefixedS3Storage(t)
		k            = kopia.NewConn(st)
		repoNameHash = strTD.NewHashForRepoConfigName()
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	err := k.Initialize(ctx, repository.Options{}, repository.Retention{}, repoNameHash)
	require.NoError(t, err, clues.ToCore(err))

	kw, err := kopia.NewWrapper(k)
	// kopiaRef comes with a count of 1 and Wrapper bumps it again so safe
	// to close here.
	k.Close(ctx)

	require.NoError(t, err, clues.ToCore(err))

	defer kw.Close(ctx)

	// Only set extend locks parameter as other retention options require a bucket
	// with object locking enabled. There's more complete tests in the kopia
	// package.
	rco, err := NewPersistentConfigOperation(
		ctx,
		control.DefaultOptions(),
		kw,
		repository.PersistentConfig{
			MinEpochDuration: ptr.To(8 * time.Hour),
		},
		evmock.NewBus())
	require.NoError(t, err, clues.ToCore(err))

	err = rco.Run(ctx)
	assert.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, Completed, rco.Status)
	assert.NotZero(t, rco.Results.StartedAt)
	assert.NotZero(t, rco.Results.CompletedAt)
	assert.NotEqual(t, rco.Results.StartedAt, rco.Results.CompletedAt)
}
