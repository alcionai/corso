package kopia

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------
// unit tests
// ---------------
type KopiaDataCollectionUnitSuite struct {
	suite.Suite
}

func TestKopiaDataCollectionUnitSuite(t *testing.T) {
	suite.Run(t, new(KopiaDataCollectionUnitSuite))
}

func (suite *KopiaDataCollectionUnitSuite) TestReturnsPath() {
	t := suite.T()

	b := path.Builder{}.Append("some", "path", "for", "data")
	pth, err := b.ToDataLayerExchangePathForCategory(
		"a-tenant",
		"a-user",
		path.EmailCategory,
		false,
	)
	require.NoError(t, err)

	c := kopiaDataCollection{
		streams: []data.Stream{},
		path:    pth,
	}

	assert.Equal(t, pth, c.FullPath())
}

func (suite *KopiaDataCollectionUnitSuite) TestReturnsStreams() {
	testData := [][]byte{
		[]byte("abcdefghijklmnopqrstuvwxyz"),
		[]byte("zyxwvutsrqponmlkjihgfedcba"),
	}

	uuids := []string{
		"a-file",
		"another-file",
	}

	table := []struct {
		name    string
		streams []data.Stream
	}{
		{
			name: "SingleStream",
			streams: []data.Stream{
				&kopiaDataStream{
					reader: io.NopCloser(bytes.NewReader(testData[0])),
					uuid:   uuids[0],
					size:   int64(len(testData[0])),
				},
			},
		},
		{
			name: "MultipleStreams",
			streams: []data.Stream{
				&kopiaDataStream{
					reader: io.NopCloser(bytes.NewReader(testData[0])),
					uuid:   uuids[0],
					size:   int64(len(testData[0])),
				},
				&kopiaDataStream{
					reader: io.NopCloser(bytes.NewReader(testData[1])),
					uuid:   uuids[1],
					size:   int64(len(testData[1])),
				},
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			c := kopiaDataCollection{
				streams: test.streams,
				path:    nil,
			}

			count := 0
			for returnedStream := range c.Items() {
				require.Less(t, count, len(test.streams))

				assert.Equal(t, returnedStream.UUID(), uuids[count])

				buf, err := io.ReadAll(returnedStream.ToReader())
				require.NoError(t, err)
				assert.Equal(t, buf, testData[count])
				require.Implements(t, (*data.StreamSize)(nil), returnedStream)
				ss := returnedStream.(data.StreamSize)
				assert.Equal(t, len(buf), int(ss.Size()))

				count++
			}

			assert.Equal(t, len(test.streams), count)
		})
	}
}
