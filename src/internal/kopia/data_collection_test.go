package kopia

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/path"
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
	expected := []string{
		"a-tenant",
		path.ExchangeService.String(),
		"a-user",
		path.EmailCategory.String(),
		"some",
		"path",
		"for",
		"data",
	}

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

	// TODO(ashmrtn): Update when data.Collection.FullPath supports path.Path
	assert.Equal(t, expected, c.FullPath())
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
				},
			},
		},
		{
			name: "MultipleStreams",
			streams: []data.Stream{
				&kopiaDataStream{
					reader: io.NopCloser(bytes.NewReader(testData[0])),
					uuid:   uuids[0],
				},
				&kopiaDataStream{
					reader: io.NopCloser(bytes.NewReader(testData[1])),
					uuid:   uuids[1],
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

				buf, err := ioutil.ReadAll(returnedStream.ToReader())
				require.NoError(t, err)
				assert.Equal(t, buf, testData[count])

				count++
			}

			assert.Equal(t, len(test.streams), count)
		})
	}
}
