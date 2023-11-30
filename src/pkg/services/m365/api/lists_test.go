package api

import (
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
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

func (suite *ListsAPIIntgSuite) TestLists_GetListById() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acl            = suite.its.ac.Lists()
		siteID         = suite.its.site.id
		listID         = "integration-test-list"
		textColumnName = "ItemName"
		itemName       = "new item"
	)

	// <------ Setup ---->
	// create list
	list := models.NewList()
	list.SetDisplayName(&listID)

	_, err := acl.Stable.Client().Sites().BySiteId(siteID).Lists().Post(ctx, list, nil)
	if err != nil {
		require.Contains(t, strings.ToLower(err.Error()), "name already exists", clues.ToCore(err))
	}

	// create column(s)
	textColumn := models.NewColumnDefinition()
	textColumn.SetName(&textColumnName)

	text := models.NewTextColumn()
	textColumn.SetText(text)

	_, err = acl.Stable.Client().Sites().BySiteId(siteID).Lists().ByListId(listID).Columns().Post(ctx, textColumn, nil)
	if err != nil {
		require.Contains(t, strings.ToLower(err.Error()), "name already exists", clues.ToCore(err))
	}

	// create list item(s)
	fields := models.NewFieldValueSet()
	additionalData := map[string]any{
		textColumnName: itemName,
	}
	fields.SetAdditionalData(additionalData)

	listItem := models.NewListItem()
	listItemName := "FirstEntry"
	listItem.SetName(&listItemName)
	listItem.SetFields(fields)

	_, err = acl.Stable.Client().Sites().BySiteId(siteID).Lists().ByListId(listID).Items().Post(ctx, listItem, nil)
	require.NoError(t, err, clues.ToCore(err))

	// <------ Test ---->

	fetchedList, err := acl.GetListByID(ctx, siteID, listID)
	require.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, fetchedList)

	cols, _, items, err := acl.getListContents(ctx, siteID, listID)
	require.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, cols)
	assert.NotEmpty(t, items)
}
