package api

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	graphTD "github.com/alcionai/corso/src/pkg/services/m365/api/graph/testdata"
)

type ListsPagerIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestListsPagerIntgSuite(t *testing.T) {
	suite.Run(t, &ListsPagerIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *ListsPagerIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *ListsPagerIntgSuite) TestEnumerateLists_withAssociatedRelationships() {
	var (
		t  = suite.T()
		ac = suite.its.gockAC.Lists()

		listID            = "fake-list-id"
		siteID            = suite.its.site.id
		textColumnDefID   = "fake-text-column-id"
		textColumnDefName = "itemName"
		numColumnDefID    = "fake-num-column-id"
		numColumnDefName  = "itemSize"
		colLinkID         = "fake-collink-id"
		cTypeID           = "fake-ctype-id"
		listItemID        = "fake-list-item-id"

		fieldsData = map[string]any{
			"itemName": "item1",
			"itemSize": 4,
		}
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	defer gock.Off()

	suite.setStubListAndItsRelationShip(listID,
		siteID,
		textColumnDefID,
		textColumnDefName,
		numColumnDefID,
		numColumnDefName,
		colLinkID,
		cTypeID,
		listItemID,
		fieldsData)

	lists, err := ac.GetLists(ctx, suite.its.site.id, CallConfig{})
	require.NoError(t, err)
	require.Equal(t, 1, len(lists))

	for _, list := range lists {
		suite.testEnumerateListItems(ctx, list, listItemID, fieldsData)
		suite.testEnumerateColumns(ctx, list, textColumnDefID)
		suite.testEnumerateContentTypes(ctx, list, cTypeID, colLinkID, numColumnDefID)
	}
}

func (suite *ListsPagerIntgSuite) testEnumerateListItems(
	ctx context.Context,
	list models.Listable,
	expectedListItemID string,
	setFieldsData map[string]any,
) []models.ListItemable {
	var listItems []models.ListItemable

	suite.Run("list item", func() {
		var (
			t   = suite.T()
			ac  = suite.its.gockAC.Lists()
			err error
		)

		listItems, err = ac.GetListItems(ctx, suite.its.site.id, *list.GetId(), CallConfig{})
		require.NoError(t, err, clues.ToCore(err))
		require.Equal(t, 1, len(listItems))

		for _, li := range listItems {
			assert.Equal(t, expectedListItemID, *li.GetId())

			fields := li.GetFields()
			require.NotEmpty(t, fields)

			fieldsData := fields.GetAdditionalData()
			expectedItemName := setFieldsData["itemName"].(string)
			actualItemName := ptr.Val(fieldsData["itemName"].(*string))
			expectedItemSize := setFieldsData["itemSize"].(int)
			actualItemSize := int(ptr.Val(fieldsData["itemSize"].(*float64)))

			assert.Equal(t, expectedItemName, actualItemName)
			assert.Equal(t, expectedItemSize, actualItemSize)
		}
	})

	return listItems
}

func (suite *ListsPagerIntgSuite) testEnumerateColumns(
	ctx context.Context,
	list models.Listable,
	expectedColumnID string,
) []models.ColumnDefinitionable {
	var columns []models.ColumnDefinitionable

	suite.Run("list columns", func() {
		var (
			t   = suite.T()
			ac  = suite.its.gockAC.Lists()
			err error
		)

		columns, err = ac.GetListColumns(ctx, suite.its.site.id, *list.GetId(), CallConfig{})
		require.NoError(t, err, clues.ToCore(err))
		require.Equal(t, 1, len(columns))

		for _, c := range columns {
			assert.Equal(suite.T(), expectedColumnID, *c.GetId())
		}
	})

	return columns
}

func (suite *ListsPagerIntgSuite) testEnumerateContentTypes(
	ctx context.Context,
	list models.Listable,
	expectedCTypeID,
	expectedColLinkID,
	expectedCTypeColID string,
) []models.ContentTypeable {
	var cTypes []models.ContentTypeable

	suite.Run("content type", func() {
		var (
			t   = suite.T()
			ac  = suite.its.gockAC.Lists()
			err error
		)

		cTypes, err = ac.GetContentTypes(ctx, suite.its.site.id, *list.GetId(), CallConfig{})
		require.NoError(t, err, clues.ToCore(err))
		require.Equal(t, 1, len(cTypes))

		for _, ct := range cTypes {
			assert.Equal(suite.T(), expectedCTypeID, *ct.GetId())

			suite.testEnumerateColumnLinks(ctx, list, ct, expectedColLinkID)
			suite.testEnumerateCTypeColumns(ctx, list, ct, expectedCTypeColID)
		}
	})

	return cTypes
}

func (suite *ListsPagerIntgSuite) testEnumerateColumnLinks(
	ctx context.Context,
	list models.Listable,
	cType models.ContentTypeable,
	expectedColLinkID string,
) []models.ColumnLinkable {
	var colLinks []models.ColumnLinkable

	suite.Run("column links", func() {
		var (
			t   = suite.T()
			ac  = suite.its.gockAC.Lists()
			err error
		)

		colLinks, err = ac.GetColumnLinks(ctx, suite.its.site.id, *list.GetId(), *cType.GetId(), CallConfig{})
		require.NoError(t, err, clues.ToCore(err))
		require.Equal(t, 1, len(colLinks))

		for _, cl := range colLinks {
			assert.Equal(suite.T(), expectedColLinkID, *cl.GetId())
		}
	})

	return colLinks
}

func (suite *ListsPagerIntgSuite) testEnumerateCTypeColumns(
	ctx context.Context,
	list models.Listable,
	cType models.ContentTypeable,
	expectedCTypeColID string,
) []models.ColumnDefinitionable {
	var cTypeCols []models.ColumnDefinitionable

	suite.Run("ctype columns", func() {
		var (
			t   = suite.T()
			ac  = suite.its.gockAC.Lists()
			err error
		)

		cTypeCols, err = ac.GetCTypesColumns(ctx, suite.its.site.id, *list.GetId(), *cType.GetId(), CallConfig{})
		require.NoError(t, err, clues.ToCore(err))
		require.Equal(t, 1, len(cTypeCols))

		for _, c := range cTypeCols {
			assert.Equal(suite.T(), expectedCTypeColID, *c.GetId())
		}
	})

	return cTypeCols
}

func (suite *ListsPagerIntgSuite) setStubListAndItsRelationShip(
	listID,
	siteID,
	textColumnDefID,
	textColumnDefName,
	numColumnDefID,
	numColumnDefName,
	colLinkID,
	cTypeID,
	listItemID string,
	fieldsData map[string]any,
) {
	list := models.NewList()
	list.SetId(&listID)

	listCol := models.NewListCollectionResponse()
	listCol.SetValue([]models.Listable{list})

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

	fields := models.NewFieldValueSet()
	fields.SetAdditionalData(fieldsData)

	listItem := models.NewListItem()
	listItem.SetId(&listItemID)
	listItem.SetFields(fields)

	listItemCol := models.NewListItemCollectionResponse()
	listItemCol.SetValue([]models.ListItemable{listItem})

	interceptV1Path(
		"sites",
		siteID,
		"lists").
		Reply(200).
		JSON(graphTD.ParseableToMap(suite.T(), listCol))

	interceptV1Path(
		"sites",
		siteID,
		"lists",
		listID,
		"columns").
		Reply(200).
		JSON(graphTD.ParseableToMap(suite.T(), columnDefCol))

	interceptV1Path(
		"sites",
		siteID,
		"lists",
		listID,
		"items").
		Reply(200).
		JSON(graphTD.ParseableToMap(suite.T(), listItemCol))

	interceptV1Path(
		"sites",
		siteID,
		"lists",
		listID,
		"contentTypes").
		Reply(200).
		JSON(graphTD.ParseableToMap(suite.T(), cTypesCol))

	interceptV1Path(
		"sites",
		siteID,
		"lists",
		listID,
		"contentTypes",
		cTypeID,
		"columns").
		Reply(200).
		JSON(graphTD.ParseableToMap(suite.T(), columnDefCol2))

	interceptV1Path(
		"sites",
		siteID,
		"lists",
		listID,
		"contentTypes",
		cTypeID,
		"columnLinks").
		Reply(200).
		JSON(graphTD.ParseableToMap(suite.T(), colLinkCol))
}
