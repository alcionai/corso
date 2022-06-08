package connector

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExchangeDataCollectionSuite struct {
	suite.Suite
}

func TestExchangeDataCollectionSuite(t *testing.T) {
	suite.Run(t, new(ExchangeDataCollectionSuite))
}

func (suite *ExchangeDataCollectionSuite) TestExchangeDataReader_Valid() {
	m := []byte("test message")
	description := "aFile"
	ed := &ExchangeData{id: description, message: m}

	// Read the message using the `ExchangeData` reader and validate it matches what we set
	buf := &bytes.Buffer{}
	buf.ReadFrom(ed.ToReader())
	assert.Equal(suite.T(), buf.Bytes(), m)
	assert.Equal(suite.T(), description, ed.UUID())
}

func (suite *ExchangeDataCollectionSuite) TestExchangeDataReader_Invalid() {
	var empty []byte
	expected := int64(0)
	ed := &ExchangeData{message: empty}
	buf := &bytes.Buffer{}
	received, err := buf.ReadFrom(ed.ToReader())
	suite.Equal(expected, received)
	assert.Nil(suite.T(), err, "received buf.Readfrom error ")
	received, err = buf.ReadFrom(ed.ToReader())
	suite.T().Logf("Received2: %v err: %v", received, err)
}

func (suite *ExchangeDataCollectionSuite) TestExchangeDataCollection_NewExchangeDataCollection() {
	name := "User"
	edc := NewExchangeDataCollection(name, []string{"Directory", "File", "task"})
	suite.Equal(name, edc.user)
	suite.True(Contains(edc.FullPath, "Directory"))
	suite.True(Contains(edc.FullPath, "File"))
	suite.True(Contains(edc.FullPath, "task"))
	suite.Zero(edc.Length())
}
