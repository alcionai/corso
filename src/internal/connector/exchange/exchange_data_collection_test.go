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
	ed := &Stream{id: description, message: m}

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
	ed := &Stream{message: empty}
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
	suite.Equal(name, edc.user)
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
		edc.PopulateCollection(&Stream{id: inputStrings[i*2], message: []byte(inputStrings[i*2+1])})
	}
	suite.Equal(expected, len(edc.data))
}

func (suite *ExchangeDataCollectionSuite) TestExchangeDataCollection_Items() {
	inputStrings := []string{"Jack", "and", "Jill", "went", "up", "the", "hill to",
		"fetch", "a", "pail", "of", "water"}
	expected := len(inputStrings) / 2 // We are using pairs
	edc := NewCollection("Fletcher", []string{"sugar", "horses", "painted red"})
	for i := 0; i < expected; i++ {
		edc.data <- &Stream{id: inputStrings[i*2], message: []byte(inputStrings[i*2+1])}
	}
	close(edc.data)
	suite.Equal(expected, len(edc.data))
	streams := edc.Items()
	suite.Equal(expected, len(streams))
	count := 0
	for item := range streams {
		assert.NotNil(suite.T(), item)
		count++
	}
	suite.Equal(count, expected)
}

func (suite *ExchangeDataCollectionSuite) TestExchangeCollection_AddJob() {
	eoc := NewCollection("Dexter", []string{"Today", "is", "was", "different"})
	suite.Zero(len(eoc.jobs))
	shopping := []string{"tomotoes", "potatoes", "pasta", "ice tea"}
	for _, item := range shopping {
		eoc.AddJob(item)
	}
	suite.Equal(len(shopping), len(eoc.jobs))

}
