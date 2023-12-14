package api

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	spMock "github.com/alcionai/corso/src/internal/m365/service/sharepoint/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	graphTD "github.com/alcionai/corso/src/pkg/services/m365/api/graph/testdata"
)

type ListsUnitSuite struct {
	tester.Suite
}

func TestListsUnitSuite(t *testing.T) {
	suite.Run(t, &ListsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ListsUnitSuite) TestBytesToListable() {
	listBytes, err := spMock.ListBytes("DataSupportSuite")
	require.NoError(suite.T(), err)

	tests := []struct {
		name       string
		byteArray  []byte
		checkError assert.ErrorAssertionFunc
		isNil      assert.ValueAssertionFunc
	}{
		{
			name:       "empty bytes",
			byteArray:  make([]byte, 0),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "invalid bytes",
			byteArray:  []byte("Invalid byte stream \"subject:\" Not going to work"),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Valid List",
			byteArray:  listBytes,
			checkError: assert.NoError,
			isNil:      assert.NotNil,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := BytesToListable(test.byteArray)
			test.checkError(t, err, clues.ToCore(err))
			test.isNil(t, result)
		})
	}
}

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
