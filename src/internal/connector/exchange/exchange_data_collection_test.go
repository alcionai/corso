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
	edc := Collection{
		user:     user,
		fullPath: fullPath,
	}
	assert.Equal(suite.T(), edc.FullPath(), fullPath)
}

func (suite *ExchangeDataCollectionSuite) TestExchangeDataCollection_NewExchangeDataCollection() {
	name := "User"
	edc := Collection{
		user:     name,
		fullPath: []string{"Directory", "File", "task"},
	}
	suite.Equal(name, edc.user)
	suite.Contains(edc.FullPath(), "Directory")
	suite.Contains(edc.FullPath(), "File")
	suite.Contains(edc.FullPath(), "task")
}

func (suite *ExchangeDataCollectionSuite) TestExchangeCollection_AddJob() {
	eoc := Collection{
		user:     "Dexter",
		fullPath: []string{"Today", "is", "currently", "different"},
	}
	suite.Zero(len(eoc.jobs))
	shopping := []string{"tomotoes", "potatoes", "pasta", "ice tea"}
	for _, item := range shopping {
		eoc.AddJob(item)
	}
	suite.Equal(len(shopping), len(eoc.jobs))
}
