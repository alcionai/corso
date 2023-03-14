package streamstore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type StreamFaultItemsIntegrationSuite struct {
	tester.Suite
}

func TestStreamFaultItemsIntegrationSuite(t *testing.T) {
	suite.Run(t, &StreamFaultItemsIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.AWSStorageCredEnvs}),
	})
}

func (suite *StreamFaultItemsIntegrationSuite) TestFaultItems() {
	t := suite.T()

	ctx, flush := tester.NewContext()
	defer flush()

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	k := kopia.NewConn(st)
	require.NoError(t, k.Initialize(ctx))

	defer k.Close(ctx)

	kw, err := kopia.NewWrapper(k)
	require.NoError(t, err)

	defer kw.Close(ctx)

	fe := fault.New(false)
	fe.AddRecoverable(fault.FileErr(assert.AnError, "id", "name", nil))
	fe.AddSkip(fault.FileSkip(fault.SkipMalware, "id2", "name2", nil))

	var (
		errs = fe.Errors()
		sd   = NewFaultErrors(kw, "tenant", path.ExchangeService)
	)

	id, err := sd.Write(ctx, errs, fault.New(true))
	require.NoError(t, err)
	require.NotNil(t, id)

	var result fault.Errors
	err = sd.Read(ctx, id, fault.UnmarshalErrorsTo(&result), fault.New(true))
	require.NoError(t, err)
	require.NotEmpty(t, result)

	assert.ElementsMatch(t, errs.Items, result.Items)
	assert.ElementsMatch(t, errs.Skipped, result.Skipped)
}
