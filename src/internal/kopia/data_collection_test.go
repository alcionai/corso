package kopia

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector"
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

	path := []string{"some", "path", "for", "data"}

	c := kopiaDataCollection{
		streams: []connector.DataStream{},
		path:    path,
	}

	assert.Equal(t, c.FullPath(), path)
}

func (suite *KopiaDataCollectionUnitSuite) TestReturnsStreams() {
	data := [][]byte{
		[]byte("abcdefghijklmnopqrstuvwxyz"),
		[]byte("zyxwvutsrqponmlkjihgfedcba"),
	}

	uuids := []string{
		"a-file",
		"another-file",
	}

	table := []struct {
		name    string
		streams []connector.DataStream
	}{
		{
			name: "SingleStream",
			streams: []connector.DataStream{
				&kopiaDataStream{
					reader: io.NopCloser(bytes.NewReader(data[0])),
					uuid:   uuids[0],
				},
			},
		},
		{
			name: "MultipleStreams",
			streams: []connector.DataStream{
				&kopiaDataStream{
					reader: io.NopCloser(bytes.NewReader(data[0])),
					uuid:   uuids[0],
				},
				&kopiaDataStream{
					reader: io.NopCloser(bytes.NewReader(data[1])),
					uuid:   uuids[1],
				},
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			c := kopiaDataCollection{
				streams: test.streams,
				path:    []string{},
			}

			count := 0
			for returnedStream := range c.Items() {
				require.Less(t, count, len(test.streams))

				assert.Equal(t, returnedStream.UUID(), uuids[count])

				buf, err := ioutil.ReadAll(returnedStream.ToReader())
				require.NoError(t, err)
				assert.Equal(t, buf, data[count])

				count++
			}

			assert.Equal(t, len(test.streams), count)
		})
	}
}
