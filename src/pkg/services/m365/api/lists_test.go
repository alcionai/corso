package api

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	graphTD "github.com/alcionai/corso/src/pkg/services/m365/api/graph/testdata"
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

func (suite *ListsAPIIntgSuite) TestLists_GetListByID() {
	var (
		listID            = "fake-list-id"
		siteID            = suite.its.site.id
		textColumnDefID   = "fake-text-column-id"
		textColumnDefName = "itemName"
		numColumnDefID    = "fake-num-column-id"
		numColumnDefName  = "itemSize"
		colLinkID         = "fake-collink-id"
		cTypeID           = "fake-ctype-id"
		listItemID        = "fake-list-item-id"
	)

	tests := []struct {
		name   string
		setupf func()
		expect assert.ErrorAssertionFunc
	}{
		{
			name: "",
			setupf: func() {
				list := models.NewList()
				list.SetId(&listID)

				txtColumnDef := models.NewColumnDefinition()
				txtColumnDef.SetId(&textColumnDefID)
				txtColumnDef.SetName(&textColumnDefName)
				textColumn := models.NewTextColumn()
				txtColumnDef.SetText(textColumn)
				columnDefCol := models.NewColumnDefinitionCollectionResponse()
				columnDefCol.SetValue([]models.ColumnDefinitionable{txtColumnDef})

				numColumnDef := models.NewColumnDefinition()
				numColumnDef.SetId(&numColumnDefID)
				numColumnDef.SetName(&numColumnDefName)
				numColumn := models.NewNumberColumn()
				numColumnDef.SetNumber(numColumn)
				columnDefCol2 := models.NewColumnDefinitionCollectionResponse()
				columnDefCol2.SetValue([]models.ColumnDefinitionable{numColumnDef})

				colLink := models.NewColumnLink()
				colLink.SetId(&colLinkID)
				colLinkCol := models.NewColumnLinkCollectionResponse()
				colLinkCol.SetValue([]models.ColumnLinkable{colLink})

				cTypes := models.NewContentType()
				cTypes.SetId(&cTypeID)
				cTypesCol := models.NewContentTypeCollectionResponse()
				cTypesCol.SetValue([]models.ContentTypeable{cTypes})

				listItem := models.NewListItem()
				listItem.SetId(&listItemID)
				listItemCol := models.NewListItemCollectionResponse()
				listItemCol.SetValue([]models.ListItemable{listItem})

				fields := models.NewFieldValueSet()
				fieldsData := map[string]any{
					"itemName": "item1",
					"itemSize": 4,
				}
				fields.SetAdditionalData(fieldsData)

				interceptV1Path(
					"sites", siteID,
					"lists", listID).
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), list))

				interceptV1Path(
					"sites", siteID,
					"lists", listID,
					"columns").
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), columnDefCol))

				interceptV1Path(
					"sites", siteID,
					"lists", listID,
					"items").
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), listItemCol))

				interceptV1Path(
					"sites", siteID,
					"lists", listID,
					"items", listItemID,
					"fields").
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), fields))

				interceptV1Path(
					"sites", siteID,
					"lists", listID,
					"contentTypes").
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), cTypesCol))

				interceptV1Path(
					"sites", siteID,
					"lists", listID,
					"contentTypes", cTypeID,
					"columns").
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), columnDefCol2))

				interceptV1Path(
					"sites", siteID,
					"lists", listID,
					"contentTypes", cTypeID,
					"columnLinks").
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), colLinkCol))
			},
			expect: assert.NoError,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			defer gock.Off()
			test.setupf()

			list, err := suite.its.gockAC.Lists().GetListByID(ctx, siteID, listID)
			test.expect(t, err)
			assert.Equal(t, listID, *list.GetId())

			items := list.GetItems()
			assert.Equal(t, 1, len(items))
			assert.Equal(t, listItemID, *items[0].GetId())

			expectedItemData := map[string]any{"itemName": ptr.To[string]("item1"), "itemSize": ptr.To[float64](float64(4))}
			itemFields := items[0].GetFields()
			itemData := itemFields.GetAdditionalData()
			assert.Equal(t, expectedItemData, itemData)

			columns := list.GetColumns()
			assert.Equal(t, 1, len(columns))
			assert.Equal(t, textColumnDefID, *columns[0].GetId())

			cTypes := list.GetContentTypes()
			assert.Equal(t, 1, len(cTypes))
			assert.Equal(t, cTypeID, *cTypes[0].GetId())

			colLinks := cTypes[0].GetColumnLinks()
			assert.Equal(t, 1, len(colLinks))
			assert.Equal(t, colLinkID, *colLinks[0].GetId())

			columns = cTypes[0].GetColumns()
			assert.Equal(t, 1, len(columns))
			assert.Equal(t, numColumnDefID, *columns[0].GetId())
		})
	}
}

func (suite *ListsAPIIntgSuite) TestLists_PostList() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acl      = suite.its.ac.Lists()
		siteID   = suite.its.site.id
		listName = testdata.DefaultRestoreConfig("list_api_post_list").Location
	)

	writer := kjson.NewJsonSerializationWriter()
	defer writer.Close()

	oldListID := "old-list"
	textColumnDefID := "list-col1"
	textColumnDefName := "itemName"
	template := "genericList"

	listInfo := models.NewListInfo()
	listInfo.SetTemplate(&template)

	textColumn := models.NewTextColumn()

	txtColumnDef := models.NewColumnDefinition()
	txtColumnDef.SetId(&textColumnDefID)
	txtColumnDef.SetName(&textColumnDefName)
	txtColumnDef.SetText(textColumn)

	list := models.NewList()
	list.SetId(&oldListID)
	list.SetColumns([]models.ColumnDefinitionable{txtColumnDef})
	list.SetList(listInfo)

	err := writer.WriteObjectValue("", list)
	require.NoError(t, err)

	oldListByteArray, err := writer.GetSerializedContent()
	require.NoError(t, err)

	newList, err := acl.PostList(ctx, siteID, listName, oldListByteArray)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, listName, ptr.Val(newList.GetDisplayName()))

	// clean up
	defer func(sID string, lst models.Listable) {
		err = acl.DeleteList(ctx, sID, ptr.Val(lst.GetId()))
		require.NoError(t, err)
	}(siteID, newList)

	_, err = acl.PostList(ctx, siteID, listName, oldListByteArray)
	require.Error(t, err)
}

func (suite *ListsAPIIntgSuite) TestLists_PostListItem() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acl      = suite.its.ac.Lists()
		siteID   = suite.its.site.id
		listName = testdata.DefaultRestoreConfig("list_api_post_list").Location
	)

	writer := kjson.NewJsonSerializationWriter()
	defer writer.Close()

	oldListID := "old-list"
	listItemID := "list-item1"
	textColumnDefID := "list-col1"
	textColumnDefName := "itemName"

	textColumn := models.NewTextColumn()

	txtColumnDef := models.NewColumnDefinition()
	txtColumnDef.SetId(&textColumnDefID)
	txtColumnDef.SetName(&textColumnDefName)
	txtColumnDef.SetText(textColumn)

	fields := models.NewFieldValueSet()
	fieldsData := map[string]any{
		textColumnDefName: "item1",
	}
	fields.SetAdditionalData(fieldsData)

	listItem := models.NewListItem()
	listItem.SetId(&listItemID)
	listItem.SetFields(fields)

	list := models.NewList()
	list.SetId(&oldListID)
	list.SetColumns([]models.ColumnDefinitionable{txtColumnDef})
	list.SetItems([]models.ListItemable{listItem})

	err := writer.WriteObjectValue("", list)
	require.NoError(t, err)

	oldListByteArray, err := writer.GetSerializedContent()
	require.NoError(t, err)

	newList, err := acl.PostList(ctx, siteID, listName, oldListByteArray)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, listName, ptr.Val(newList.GetDisplayName()))

	// clean up
	defer func(sID string, lst models.Listable) {
		err = acl.DeleteList(ctx, sID, ptr.Val(lst.GetId()))
		require.NoError(t, err)
	}(siteID, newList)

	newListItems, err := acl.PostListItem(ctx, siteID, ptr.Val(newList.GetId()), oldListByteArray)
	require.NoError(t, err, clues.ToCore(err))
	require.Less(t, 0, len(newListItems))

	newListItemFields := newListItems[0].GetFields()
	require.NotEmpty(t, newListItemFields)

	newListItemsData := newListItemFields.GetAdditionalData()
	require.NotEmpty(t, newListItemsData)

	for k, v := range newListItemsData {
		assert.Equal(t, fieldsData[k], ptr.Val(v.(*string)))
	}
}

func (suite *ListsAPIIntgSuite) TestLists_DeleteList() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acl      = suite.its.ac.Lists()
		siteID   = suite.its.site.id
		listName = testdata.DefaultRestoreConfig("list_api_post_list").Location
	)

	writer := kjson.NewJsonSerializationWriter()
	defer writer.Close()

	oldListID := "old-list"
	textColumnDefID := "list-col1"
	textColumnDefName := "itemName"
	template := "genericList"

	listInfo := models.NewListInfo()
	listInfo.SetTemplate(&template)

	textColumn := models.NewTextColumn()

	txtColumnDef := models.NewColumnDefinition()
	txtColumnDef.SetId(&textColumnDefID)
	txtColumnDef.SetName(&textColumnDefName)
	txtColumnDef.SetText(textColumn)

	list := models.NewList()
	list.SetId(&oldListID)
	list.SetColumns([]models.ColumnDefinitionable{txtColumnDef})
	list.SetList(listInfo)

	err := writer.WriteObjectValue("", list)
	require.NoError(t, err)

	oldListByteArray, err := writer.GetSerializedContent()
	require.NoError(t, err)

	newList, err := acl.PostList(ctx, siteID, listName, oldListByteArray)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, listName, ptr.Val(newList.GetDisplayName()))

	err = acl.DeleteList(ctx, siteID, ptr.Val(newList.GetId()))
	require.NoError(t, err)
}
