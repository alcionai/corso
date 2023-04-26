package graph

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type CollectionsUnitSuite struct {
	tester.Suite
}

func TestCollectionsUnitSuite(t *testing.T) {
	suite.Run(t, &CollectionsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CollectionsUnitSuite) TestNewPrefixCollection() {
	t := suite.T()
	serv := path.OneDriveService
	cat := path.FilesCategory

	p1, err := path.ServicePrefix("t", "ro1", serv, cat)
	require.NoError(t, err, clues.ToCore(err))

	p2, err := path.ServicePrefix("t", "ro2", serv, cat)
	require.NoError(t, err, clues.ToCore(err))

	items, err := path.Build("t", "ro", serv, cat, true, "fld", "itm")
	require.NoError(t, err, clues.ToCore(err))

	folders, err := path.Build("t", "ro", serv, cat, false, "fld")
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name      string
		prev      path.Path
		full      path.Path
		expectErr require.ErrorAssertionFunc
	}{
		{
			name:      "not moved",
			prev:      p1,
			full:      p1,
			expectErr: require.NoError,
		},
		{
			name:      "moved",
			prev:      p1,
			full:      p2,
			expectErr: require.NoError,
		},
		{
			name:      "deleted",
			prev:      p1,
			full:      nil,
			expectErr: require.Error,
		},
		{
			name:      "new",
			prev:      nil,
			full:      p2,
			expectErr: require.Error,
		},
		{
			name:      "prev has items",
			prev:      items,
			full:      p1,
			expectErr: require.Error,
		},
		{
			name:      "prev has folders",
			prev:      folders,
			full:      p1,
			expectErr: require.Error,
		},
		{
			name:      "full has items",
			prev:      p1,
			full:      items,
			expectErr: require.Error,
		},
		{
			name:      "full has folders",
			prev:      p1,
			full:      folders,
			expectErr: require.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			_, err := NewPrefixCollection(test.prev, test.full, nil)
			test.expectErr(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *CollectionsUnitSuite) TestNewDeletedPrefixCollection() {
	t := suite.T()
	serv := path.OneDriveService
	cat := path.FilesCategory

	p1, err := path.ServicePrefix("t", "ro1", serv, cat)
	require.NoError(t, err, clues.ToCore(err))

	items, err := path.Build("t", "ro", serv, cat, true, "fld", "itm")
	require.NoError(t, err, clues.ToCore(err))

	folders, err := path.Build("t", "ro", serv, cat, false, "fld")
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name      string
		prev      path.Path
		expectErr require.ErrorAssertionFunc
	}{
		{
			name:      "nil",
			prev:      nil,
			expectErr: require.Error,
		},
		{
			name:      "deleted",
			prev:      p1,
			expectErr: require.NoError,
		},
		{
			name:      "prev has items",
			prev:      items,
			expectErr: require.Error,
		},
		{
			name:      "prev has folders",
			prev:      folders,
			expectErr: require.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			_, err := NewDeletedPrefixCollection(test.prev, nil)
			test.expectErr(suite.T(), err, clues.ToCore(err))
		})
	}
}
