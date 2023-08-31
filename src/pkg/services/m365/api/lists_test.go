package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control/testdata"
)

type ListsAPIIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func (suite *ListsAPIIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func TestListsAPIIntgSuite(t *testing.T) {
	suite.Run(t, &ListsAPIIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *ListsAPIIntgSuite) TestLists_PostDrive() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acl       = suite.its.ac.Lists()
		driveName = testdata.DefaultRestoreConfig("list_api_post_drive").Location
		siteID    = suite.its.site.id
	)

	// first post, should have no errors
	list, err := acl.PostDrive(ctx, siteID, driveName)
	require.NoError(t, err, clues.ToCore(err))
	// the site name cannot be set when posting, only its DisplayName.
	// so we double check here that we're still getting the name we expect.
	assert.Equal(t, driveName, ptr.Val(list.GetName()))

	// second post, same name, should error on name conflict]
	_, err = acl.PostDrive(ctx, siteID, driveName)
	require.ErrorIs(t, err, graph.ErrItemAlreadyExistsConflict, clues.ToCore(err))
}
