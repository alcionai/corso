package exchange

import (
	"bytes"
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/graph"
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
	bool,
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

type CollectionSuite struct {
	tester.Suite
}

func TestCollectionSuite(t *testing.T) {
	suite.Run(t, &CollectionSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CollectionSuite) TestReader_Valid() {
	m := []byte("test message")
	description := "aFile"
	ed := &Item{id: description, message: m}

	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(ed.ToReader())
	assert.NoError(suite.T(), err, clues.ToCore(err))
	assert.Equal(suite.T(), buf.Bytes(), m)
	assert.Equal(suite.T(), description, ed.ID())
}

func (suite *CollectionSuite) TestReader_Empty() {
	var (
		empty    []byte
		expected int64
		t        = suite.T()
	)

	ed := &Item{message: empty}
	buf := &bytes.Buffer{}
	received, err := buf.ReadFrom(ed.ToReader())

	assert.Equal(t, expected, received)
	assert.NoError(t, err, clues.ToCore(err))
}

func (suite *CollectionSuite) TestColleciton_FullPath() {
	t := suite.T()
	tenant := "a-tenant"
	user := "a-user"
	folder := "a-folder"

	fullPath, err := path.Build(
		tenant,
		[]path.ServiceResource{{
			Service:           path.ExchangeService,
			ProtectedResource: user,
		}},
		path.EmailCategory,
		false,
		folder)
	require.NoError(t, err, clues.ToCore(err))

	edc := Collection{
		user:     user,
		fullPath: fullPath,
	}

	assert.Equal(t, fullPath, edc.FullPath())
}

func (suite *CollectionSuite) TestCollection_NewCollection() {
	t := suite.T()
	tenant := "a-tenant"
	user := "a-user"
	folder := "a-folder"
	name := "User"

	fullPath, err := path.Build(
		tenant,
		[]path.ServiceResource{{
			Service:           path.ExchangeService,
			ProtectedResource: user,
		}},
		path.EmailCategory,
		false,
		folder)
	require.NoError(t, err, clues.ToCore(err))

	edc := Collection{
		user:     name,
		fullPath: fullPath,
	}
	assert.Equal(t, name, edc.user)
	assert.Equal(t, fullPath, edc.FullPath())
}

func (suite *CollectionSuite) TestNewCollection_state() {
	fooP, err := path.Build(
		"t",
		[]path.ServiceResource{{
			Service:           path.ExchangeService,
			ProtectedResource: "u",
		}},
		path.EmailCategory,
		false,
		"foo")
	require.NoError(suite.T(), err, clues.ToCore(err))
	barP, err := path.Build(
		"t",
		[]path.ServiceResource{{
			Service:           path.ExchangeService,
			ProtectedResource: "u",
		}},
		path.EmailCategory,
		false,
		"bar")
	require.NoError(suite.T(), err, clues.ToCore(err))

	locPB := path.Builder{}.Append("human-readable")

	table := []struct {
		name   string
		prev   path.Path
		curr   path.Path
		loc    *path.Builder
		expect data.CollectionState
	}{
		{
			name:   "new",
			curr:   fooP,
			loc:    locPB,
			expect: data.NewState,
		},
		{
			name:   "not moved",
			prev:   fooP,
			curr:   fooP,
			loc:    locPB,
			expect: data.NotMovedState,
		},
		{
			name:   "moved",
			prev:   fooP,
			curr:   barP,
			loc:    locPB,
			expect: data.MovedState,
		},
		{
			name:   "deleted",
			prev:   fooP,
			expect: data.DeletedState,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			c := NewCollection(
				"u",
				test.curr, test.prev, test.loc,
				0,
				&mockItemer{}, nil,
				control.DefaultOptions(),
				false)
			assert.Equal(t, test.expect, c.State(), "collection state")
			assert.Equal(t, test.curr, c.fullPath, "full path")
			assert.Equal(t, test.prev, c.prevPath, "prev path")
			assert.Equal(t, test.loc, c.locationPath, "location path")
		})
	}
}

func (suite *CollectionSuite) TestGetItemWithRetries() {
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
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectGetCalls: 1,
		},
		{
			name:  "an error",
			items: &mockItemer{getErr: assert.AnError},
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
			expectGetCalls: 3,
		},
		{
			name: "deleted in flight",
			items: &mockItemer{
				getErr: graph.ErrDeletedInFlight,
			},
			expectErr: func(t *testing.T, err error) {
				assert.True(t, graph.IsErrDeletedInFlight(err), "is ErrDeletedInFlight")
			},
			expectGetCalls: 1,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			// itemer is mocked, so only the errors are configured atm.
			_, _, err := test.items.GetItem(ctx, "userID", "itemID", false, fault.New(true))
			test.expectErr(t, err)
		})
	}
}
