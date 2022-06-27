package kopia

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// ---------------
// unit tests
// ---------------
type SingleItemCollectionUnitSuite struct {
	suite.Suite
}

func TestSingleItemCollectionUnitSuite(t *testing.T) {
	suite.Run(t, new(SingleItemCollectionUnitSuite))
}

func (suite *SingleItemCollectionUnitSuite) TestReturnsPath() {
	t := suite.T()

	path := []string{"some", "path", "for", "data"}

	c := singleItemCollection{
		stream: kopiaDataStream{},
		path:   path,
	}

	assert.Equal(t, c.FullPath(), path)
}

func (suite *SingleItemCollectionUnitSuite) TestReturnsOnlyOneItem() {
	t := suite.T()

	data := []byte("abcdefghijklmnopqrstuvwxyz")
	uuid := "a-file"
	stream := &kopiaDataStream{
		reader: io.NopCloser(bytes.NewReader(data)),
		uuid:   uuid,
	}

	c := singleItemCollection{
		stream: stream,
		path:   []string{},
	}

	count := 0
	for returnedStream := range c.Items() {
		assert.Equal(t, returnedStream.UUID(), uuid)

		buf, err := ioutil.ReadAll(returnedStream.ToReader())
		require.NoError(t, err)
		assert.Equal(t, buf, data)

		count++
	}

	assert.Equal(t, 1, count)
}
