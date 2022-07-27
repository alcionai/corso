package exchange

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExchangeDataCollectionSuite struct {
	suite.Suite
}

func contains(elems []string, value string) bool {
	for _, s := range elems {
		if value == s {
			return true
		}
	}
	return false
}

func TestExchangeDataCollectionSuite(t *testing.T) {
	suite.Run(t, new(ExchangeDataCollectionSuite))
}

func (suite *ExchangeDataCollectionSuite) TestExchangeDataReader_Valid() {
	m := []byte("test message")
	description := "aFile"
	ed := &Stream{Id: description, Message: m}

	// Read the message using the `ExchangeData` reader and validate it matches what we set
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(ed.ToReader())
	assert.Nil(suite.T(), err, "received a buf.Read error")
	assert.Equal(suite.T(), buf.Bytes(), m)
	assert.Equal(suite.T(), description, ed.UUID())
}

func (suite *ExchangeDataCollectionSuite) TestExchangeDataReader_Empty() {
	var empty []byte
	expected := int64(0)
	ed := &Stream{Message: empty}
	buf := &bytes.Buffer{}
	received, err := buf.ReadFrom(ed.ToReader())
	suite.Equal(expected, received)
	assert.Nil(suite.T(), err, "received buf.Readfrom error ")
}
func (suite *ExchangeDataCollectionSuite) TestExchangeData_FullPath() {
	user := "a-user"
	fullPath := []string{"a-tenant", user, "emails"}
	edc := NewCollection(user, fullPath)
	assert.Equal(suite.T(), edc.FullPath(), fullPath)
}

func (suite *ExchangeDataCollectionSuite) TestExchangeDataCollection_NewExchangeDataCollection() {
	name := "User"
	edc := NewCollection(name, []string{"Directory", "File", "task"})
	suite.Equal(name, edc.User)
	suite.True(contains(edc.FullPath(), "Directory"))
	suite.True(contains(edc.FullPath(), "File"))
	suite.True(contains(edc.FullPath(), "task"))
}

func (suite *ExchangeDataCollectionSuite) TestExchangeDataCollection_PopulateCollection() {
	inputStrings := []string{"Jack", "and", "Jill", "went", "up", "the", "hill to",
		"fetch", "a", "pail", "of", "water"}
	expected := len(inputStrings) / 2 // We are using pairs
	edc := NewCollection("Fletcher", []string{"sugar", "horses", "painted red"})
	for i := 0; i < expected; i++ {
		edc.PopulateCollection(&Stream{Id: inputStrings[i*2], Message: []byte(inputStrings[i*2+1])})
	}
	suite.Equal(expected, len(edc.Data))
}

func (suite *ExchangeDataCollectionSuite) TestExchangeDataCollection_Items() {
	inputStrings := []string{"Jack", "and", "Jill", "went", "up", "the", "hill to",
		"fetch", "a", "pail", "of", "water"}
	expected := len(inputStrings) / 2 // We are using pairs
	edc := NewCollection("Fletcher", []string{"sugar", "horses", "painted red"})
	for i := 0; i < expected; i++ {
		edc.Data <- &Stream{Id: inputStrings[i*2], Message: []byte(inputStrings[i*2+1])}
	}
	close(edc.Data)
	suite.Equal(expected, len(edc.Data))
	streams := edc.Items()
	suite.Equal(expected, len(streams))
	count := 0
	for item := range streams {
		assert.NotNil(suite.T(), item)
		count++
	}
	suite.Equal(count, expected)
}
