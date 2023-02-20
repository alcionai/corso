package exchange

import (
	"bytes"
	"context"
	"testing"

	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

type mockItemer struct {
	getCount       int
	serializeCount int
	getErr         error
	serializeErr   error
}

func (mi *mockItemer) GetItem(
	context.Context,
	string, string,
	*fault.Bus,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	mi.getCount++
	return nil, nil, mi.getErr
}

func (mi *mockItemer) Serialize(
	context.Context,
	serialization.Parsable,
	string, string,
) ([]byte, error) {
	mi.serializeCount++
	return nil, mi.serializeErr
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
	locP, err := path.Builder{}.
		Append("human-readable").
		ToDataLayerExchangePathForCategory("t", "u", path.EmailCategory, false)
	require.NoError(suite.T(), err)

	table := []struct {
		name   string
		prev   path.Path
		curr   path.Path
		loc    path.Path
		expect data.CollectionState
	}{
		{
			name:   "new",
			curr:   fooP,
			loc:    locP,
			expect: data.NewState,
		},
		{
			name:   "not moved",
			prev:   fooP,
			curr:   fooP,
			loc:    locP,
			expect: data.NotMovedState,
		},
		{
			name:   "moved",
			prev:   fooP,
			curr:   barP,
			loc:    locP,
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
				test.curr, test.prev, test.loc,
				0,
				&mockItemer{}, nil,
				control.Options{},
				false)
			assert.Equal(t, test.expect, c.State(), "collection state")
			assert.Equal(t, test.curr, c.fullPath, "full path")
			assert.Equal(t, test.prev, c.prevPath, "prev path")
			assert.Equal(t, test.loc, c.locationPath, "location path")
		})
	}
}

func (suite *ExchangeDataCollectionSuite) TestGetItemWithRetries() {
	table := []struct {
		name           string
		items          *mockItemer
		expectErr      func(*testing.T, error)
		expectGetCalls int
	}{
		{
			name:  "happy",
			items: &mockItemer{},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
			expectGetCalls: 1,
		},
		{
			name:  "an error",
			items: &mockItemer{getErr: assert.AnError},
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
			expectGetCalls: 3,
		},
		{
			name: "deleted in flight",
			items: &mockItemer{
				getErr: graph.ErrDeletedInFlight{
					Err: *common.EncapsulateError(assert.AnError),
				},
			},
			expectErr: func(t *testing.T, err error) {
				assert.True(t, graph.IsErrDeletedInFlight(err), "is ErrDeletedInFlight")
			},
			expectGetCalls: 1,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			// itemer is mocked, so only the errors are configured atm.
			_, _, err := getItemWithRetries(ctx, "userID", "itemID", test.items, fault.New(true))
			test.expectErr(t, err)
		})
	}
}
