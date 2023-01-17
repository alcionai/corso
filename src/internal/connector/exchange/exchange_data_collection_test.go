package exchange

import (
	"bytes"
	"context"
	"testing"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

type mockItemer struct{}

func (mi mockItemer) GetItem(
	context.Context,
	string, string,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	return nil, nil, nil
}

func (mi mockItemer) Serialize(context.Context, serialization.Parsable, string, string) ([]byte, error) {
	return nil, nil
}

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

func (suite *ExchangeDataCollectionSuite) TestNewCollection_state() {
	fooP, err := path.Builder{}.
		Append("foo").
		ToDataLayerExchangePathForCategory("t", "u", path.EmailCategory, false)
	require.NoError(suite.T(), err)
	barP, err := path.Builder{}.
		Append("bar").
		ToDataLayerExchangePathForCategory("t", "u", path.EmailCategory, false)
	require.NoError(suite.T(), err)

	table := []struct {
		name   string
		prev   path.Path
		curr   path.Path
		expect data.CollectionState
	}{
		{
			name:   "new",
			curr:   fooP,
			expect: data.NewState,
		},
		{
			name:   "not moved",
			prev:   fooP,
			curr:   fooP,
			expect: data.NotMovedState,
		},
		{
			name:   "moved",
			prev:   fooP,
			curr:   barP,
			expect: data.MovedState,
		},
		{
			name:   "deleted",
			prev:   fooP,
			expect: data.DeletedState,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			c := NewCollection(
				"u",
				test.curr, test.prev,
				0,
				mockItemer{}, nil,
				control.Options{},
				false)
			assert.Equal(t, test.expect, c.State())
		})
	}
}
