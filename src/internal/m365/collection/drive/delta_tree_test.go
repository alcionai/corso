package drive

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type DeltaTreeUnitSuite struct {
	tester.Suite
}

func TestDeltaTreeUnitSuite(t *testing.T) {
	suite.Run(t, &DeltaTreeUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *DeltaTreeUnitSuite) TestNewFolderyMcFolderFace() {
	var (
		t      = suite.T()
		p, err = path.BuildPrefix("t", "r", path.OneDriveService, path.FilesCategory)
	)

	require.NoError(t, err, clues.ToCore(err))

	folderFace := newFolderyMcFolderFace(p)
	assert.Equal(t, p, folderFace.prefix)
	assert.Nil(t, folderFace.collections)
	assert.NotNil(t, folderFace.folderIDToNode)
	assert.NotNil(t, folderFace.tombstones)
	assert.NotNil(t, folderFace.excludeFileIDs)
}

func (suite *DeltaTreeUnitSuite) TestNewNodeyMcNodeFace() {
	var (
		t      = suite.T()
		parent = &nodeyMcNodeFace{}
		p, err = path.Build("t", "r", path.SharePointService, path.LibrariesCategory, false, "drive-id", "root:")
	)

	require.NoError(t, err, clues.ToCore(err))

	nodeFace := newNodeyMcNodeFace(parent, "id", "name", p, true)
	assert.Equal(t, parent, nodeFace.parent)
	assert.Equal(t, "id", nodeFace.id)
	assert.Equal(t, "name", nodeFace.name)
	assert.Equal(t, p, nodeFace.prev)
	assert.True(t, nodeFace.isPackage)
	assert.NotNil(t, nodeFace.childDirs)
	assert.NotNil(t, nodeFace.items)
}
