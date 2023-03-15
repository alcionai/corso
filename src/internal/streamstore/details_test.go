package streamstore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type StreamDetailsIntegrationSuite struct {
	tester.Suite
}

func TestStreamDetailsIntegrationSuite(t *testing.T) {
	suite.Run(t, &StreamDetailsIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.AWSStorageCredEnvs}),
	})
}

func (suite *StreamDetailsIntegrationSuite) TestDetails() {
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

	deetsBuilder := &details.Builder{}

	require.NoError(
		t,
		deetsBuilder.Add("ref", "shortref", "parentref", "locationRef", true,
			details.ItemInfo{
				Exchange: &details.ExchangeInfo{
					Subject: "hello world",
				},
			}))

	var (
		deets = deetsBuilder.Details()
		sd    = NewDetails(kw, "tenant", path.ExchangeService)
	)

	id, err := sd.Write(ctx, deets, fault.New(true))
	require.NoError(t, err)
	require.NotNil(t, id)

	var readDeets details.Details
	err = sd.Read(ctx, id, details.UnmarshalTo(&readDeets), fault.New(true))
	require.NoError(t, err)
	require.NotEmpty(t, readDeets)

	assert.Equal(t, len(deets.Entries), len(readDeets.Entries))
	assert.Equal(t, deets.Entries[0].ParentRef, readDeets.Entries[0].ParentRef)
	assert.Equal(t, deets.Entries[0].ShortRef, readDeets.Entries[0].ShortRef)
	assert.Equal(t, deets.Entries[0].RepoRef, readDeets.Entries[0].RepoRef)
	assert.Equal(t, deets.Entries[0].LocationRef, readDeets.Entries[0].LocationRef)
	assert.Equal(t, deets.Entries[0].Updated, readDeets.Entries[0].Updated)
	assert.NotNil(t, readDeets.Entries[0].Exchange)
	assert.Equal(t, *deets.Entries[0].Exchange, *readDeets.Entries[0].Exchange)
}
