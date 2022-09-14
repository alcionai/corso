package exchange

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/path"
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
	var (
		empty    []byte
		expected int64
	)

	ed := &Stream{message: empty}
	buf := &bytes.Buffer{}
	received, err := buf.ReadFrom(ed.ToReader())

	suite.Equal(expected, received)
	assert.Nil(suite.T(), err, "received buf.Readfrom error ")
}

func (suite *ExchangeDataCollectionSuite) TestExchangeData_FullPath() {
	t := suite.T()
	tenant := "a-tenant"
	user := "a-user"
	folder := "a-folder"

	fullPath, err := path.Builder{}.Append(folder).ToDataLayerExchangePathForCategory(
		tenant,
		user,
		path.EmailCategory,
		false,
	)
	require.NoError(t, err)

	edc := Collection{
		user:     user,
		fullPath: fullPath,
	}

	assert.Equal(t, fullPath, edc.FullPath())
}

func (suite *ExchangeDataCollectionSuite) TestExchangeDataCollection_NewExchangeDataCollection() {
	tenant := "a-tenant"
	user := "a-user"
	folder := "a-folder"
	name := "User"

	fullPath, err := path.Builder{}.Append(folder).ToDataLayerExchangePathForCategory(
		tenant,
		user,
		path.EmailCategory,
		false,
	)
	require.NoError(suite.T(), err)

	edc := Collection{
		user:     name,
		fullPath: fullPath,
	}
	suite.Equal(name, edc.user)
	suite.Equal(fullPath, edc.FullPath())
}

func (suite *ExchangeDataCollectionSuite) TestExchangeCollection_AddJob() {
	eoc := Collection{
		user:     "Dexter",
		fullPath: nil,
	}
	suite.Zero(len(eoc.jobs))

	shopping := []string{"tomotoes", "potatoes", "pasta", "ice tea"}
	for _, item := range shopping {
		eoc.AddJob(item)
	}

	suite.Equal(len(shopping), len(eoc.jobs))
}
