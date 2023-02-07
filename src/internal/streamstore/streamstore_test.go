package streamstore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

type StreamStoreIntegrationSuite struct {
	suite.Suite
}

func TestStreamStoreIntegrationSuite(t *testing.T) {
	tester.RunOnAny(t, tester.CorsoCITests)
	suite.Run(t, new(StreamStoreIntegrationSuite))
}

func (suite *StreamStoreIntegrationSuite) TestDetails() {
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

	deetsBuilder.Add("ref", "shortref", "parentref", "locationRef", true,
		details.ItemInfo{
			Exchange: &details.ExchangeInfo{
				Subject: "hello world",
			},
		})

	deets := deetsBuilder.Details()

	ss := New(kw, "tenant", path.ExchangeService)

	id, err := ss.WriteBackupDetails(ctx, deets)
	require.NoError(t, err)
	require.NotNil(t, id)

	readDeets, err := ss.ReadBackupDetails(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, readDeets)

	assert.Equal(t, len(deets.Entries), len(readDeets.Entries))
	assert.Equal(t, deets.Entries[0].ParentRef, readDeets.Entries[0].ParentRef)
	assert.Equal(t, deets.Entries[0].ShortRef, readDeets.Entries[0].ShortRef)
	assert.Equal(t, deets.Entries[0].RepoRef, readDeets.Entries[0].RepoRef)
	assert.Equal(t, deets.Entries[0].LocationRef, readDeets.Entries[0].LocationRef)
	assert.Equal(t, deets.Entries[0].Updated, readDeets.Entries[0].Updated)
	assert.NotNil(t, readDeets.Entries[0].Exchange)
	assert.Equal(t, *deets.Entries[0].Exchange, *readDeets.Entries[0].Exchange)
}
