package groups

import (
	"bytes"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

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
	protectedResource := "a-protectedResource"
	folder := "a-folder"

	fullPath, err := path.Build(
		tenant,
		protectedResource,
		path.GroupsService,
		path.ChannelMessagesCategory,
		false,
		folder)
	require.NoError(t, err, clues.ToCore(err))

	edc := Collection{
		protectedResource: protectedResource,
		fullPath:          fullPath,
	}

	assert.Equal(t, fullPath, edc.FullPath())
}

func (suite *CollectionSuite) TestCollection_NewCollection() {
	t := suite.T()
	tenant := "a-tenant"
	protectedResource := "a-protectedResource"
	folder := "a-folder"
	name := "protectedResource"

	fullPath, err := path.Build(
		tenant,
		protectedResource,
		path.GroupsService,
		path.ChannelMessagesCategory,
		false,
		folder)
	require.NoError(t, err, clues.ToCore(err))

	edc := Collection{
		protectedResource: name,
		fullPath:          fullPath,
	}
	assert.Equal(t, name, edc.protectedResource)
	assert.Equal(t, fullPath, edc.FullPath())
}

func (suite *CollectionSuite) TestNewCollection_state() {
	fooP, err := path.Build("t", "u", path.GroupsService, path.ChannelMessagesCategory, false, "foo")
	require.NoError(suite.T(), err, clues.ToCore(err))
	barP, err := path.Build("t", "u", path.GroupsService, path.ChannelMessagesCategory, false, "bar")
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
				"g",
				test.curr, test.prev, test.loc,
				0,
				nil,
				control.DefaultOptions(),
				"chanID",
				"chanName")
			assert.Equal(t, test.expect, c.State(), "collection state")
			assert.Equal(t, test.curr, c.fullPath, "full path")
			assert.Equal(t, test.prev, c.prevPath, "prev path")
			assert.Equal(t, test.loc, c.locationPath, "location path")
		})
	}
}
