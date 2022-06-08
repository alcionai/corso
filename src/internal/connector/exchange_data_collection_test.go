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

func (suite *ExchangeDataCollectionSuite) TestExchangeDataReader() {
	m := []byte("test message")
	ed := &ExchangeData{message: m}

	// Read the message using the `ExchangeData` reader and validate it matches what we set
	buf := &bytes.Buffer{}
	buf.ReadFrom(ed.ToReader())
	assert.Equal(suite.T(), buf.Bytes(), m)
}

func (suite *ExchangeDataCollectionSuite) TestExchangeData_FullPath() {
	user := "a-user"
	fullPath := []string{"a-tenant", user, "emails"}
	edc := NewExchangeDataCollection(user, fullPath)

	assert.Equal(suite.T(), edc.FullPath(), fullPath)
}
