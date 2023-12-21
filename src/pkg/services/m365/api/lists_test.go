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

func (suite *ListsUnitSuite) TestColumnDefinitionable_GetValidation() {
	tests := []struct {
		name    string
		getOrig func() models.ColumnDefinitionable
		expect  assert.ValueAssertionFunc
	}{
		{
			name: "column validation not set",
			getOrig: func() models.ColumnDefinitionable {
				textColumn := models.NewTextColumn()

				cd := models.NewColumnDefinition()
				cd.SetText(textColumn)

				return cd
			},
			expect: assert.Nil,
		},
		{
			name: "column validation set",
			getOrig: func() models.ColumnDefinitionable {
				textColumn := models.NewTextColumn()

				colValidation := models.NewColumnValidation()

				cd := models.NewColumnDefinition()
				cd.SetText(textColumn)
				cd.SetValidation(colValidation)

				return cd
			},
			expect: assert.NotNil,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			orig := test.getOrig()
			newCd := cloneColumnDefinitionable(orig)

			require.NotEmpty(t, newCd)

			test.expect(t, newCd.GetValidation())
		})
	}
}

func (suite *ListsUnitSuite) TestColumnDefinitionable_GetDefaultValue() {
	tests := []struct {
		name    string
		getOrig func() models.ColumnDefinitionable
		expect  func(t *testing.T, cd models.ColumnDefinitionable)
	}{
		{
			name: "column default value not set",
			getOrig: func() models.ColumnDefinitionable {
				textColumn := models.NewTextColumn()

				cd := models.NewColumnDefinition()
				cd.SetText(textColumn)

				return cd
			},
			expect: func(t *testing.T, cd models.ColumnDefinitionable) {
				assert.Nil(t, cd.GetDefaultValue())
			},
		},
		{
			name: "column default value set",
			getOrig: func() models.ColumnDefinitionable {
				defaultVal := "some-val"

				textColumn := models.NewTextColumn()

				colDefaultVal := models.NewDefaultColumnValue()
				colDefaultVal.SetValue(ptr.To(defaultVal))

				cd := models.NewColumnDefinition()
				cd.SetText(textColumn)
				cd.SetDefaultValue(colDefaultVal)

				return cd
			},
			expect: func(t *testing.T, cd models.ColumnDefinitionable) {
				assert.NotNil(t, cd.GetDefaultValue())
				assert.Equal(t, "some-val", ptr.Val(cd.GetDefaultValue().GetValue()))
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			orig := test.getOrig()
			newCd := cloneColumnDefinitionable(orig)

			require.NotEmpty(t, newCd)
			test.expect(t, newCd)
		})
	}
}

func (suite *ListsUnitSuite) TestColumnDefinitionable_ColumnType() {
	tests := []struct {
		name    string
		getOrig func() models.ColumnDefinitionable
		checkFn func(models.ColumnDefinitionable) bool
	}{
		{
			name: "column type should be number",
			getOrig: func() models.ColumnDefinitionable {
				numColumn := models.NewNumberColumn()

				cd := models.NewColumnDefinition()
				cd.SetNumber(numColumn)

				return cd
			},
			checkFn: func(cd models.ColumnDefinitionable) bool {
				return cd.GetNumber() != nil
			},
		},
		{
			name: "column type should be person or group",
			getOrig: func() models.ColumnDefinitionable {
				pgColumn := models.NewPersonOrGroupColumn()

				cd := models.NewColumnDefinition()
				cd.SetPersonOrGroup(pgColumn)

				return cd
			},
			checkFn: func(cd models.ColumnDefinitionable) bool {
				return cd.GetPersonOrGroup() != nil
			},
		},
		{
			name: "column type should default to text",
			getOrig: func() models.ColumnDefinitionable {
				return models.NewColumnDefinition()
			},
			checkFn: func(cd models.ColumnDefinitionable) bool {
				return cd.GetText() != nil
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			orig := test.getOrig()
			newCd := cloneColumnDefinitionable(orig)

			require.NotEmpty(t, newCd)
			assert.True(t, test.checkFn(newCd))
		})
	}
}

func (suite *ListsUnitSuite) TestColumnDefinitionable_LegacyColumns() {
	listName := "test-list"
	textColumnName := "ItemName"
	textColumnDisplayName := "Item Name"
	titleColumnName := "Title"
	titleColumnDisplayName := "Title"
	readOnlyColumnName := "TestColumn"
	readOnlyColumnDisplayName := "Test Column"

	contentTypeCd := models.NewColumnDefinition()
	contentTypeCd.SetName(ptr.To(ContentTypeColumnName))
	contentTypeCd.SetDisplayName(ptr.To(ContentTypeColumnDisplayName))

	attachmentCd := models.NewColumnDefinition()
	attachmentCd.SetName(ptr.To(AttachmentsColumnName))
	attachmentCd.SetDisplayName(ptr.To(AttachmentsColumnName))

	editCd := models.NewColumnDefinition()
	editCd.SetName(ptr.To(EditColumnName))
	editCd.SetDisplayName(ptr.To(EditColumnName))

	textCol := models.NewTextColumn()
	titleCol := models.NewTextColumn()
	roCol := models.NewTextColumn()

	textCd := models.NewColumnDefinition()
	textCd.SetName(ptr.To(textColumnName))
	textCd.SetDisplayName(ptr.To(textColumnDisplayName))
	textCd.SetText(textCol)

	titleCd := models.NewColumnDefinition()
	titleCd.SetName(ptr.To(titleColumnName))
	titleCd.SetDisplayName(ptr.To(titleColumnDisplayName))
	titleCd.SetText(titleCol)

	roCd := models.NewColumnDefinition()
	roCd.SetName(ptr.To(readOnlyColumnName))
	roCd.SetDisplayName(ptr.To(readOnlyColumnDisplayName))
	roCd.SetText(roCol)
	roCd.SetReadOnly(ptr.To(true))

	tests := []struct {
		name    string
		getList func() *models.List
		length  int
	}{
		{
			name: "all legacy columns",
			getList: func() *models.List {
				lst := models.NewList()
				lst.SetColumns([]models.ColumnDefinitionable{
					contentTypeCd,
					attachmentCd,
					editCd,
				})
				return lst
			},
			length: 0,
		},
		{
			name: "title and legacy columns",
			getList: func() *models.List {
				lst := models.NewList()
				lst.SetColumns([]models.ColumnDefinitionable{
					contentTypeCd,
					attachmentCd,
					editCd,
					titleCd,
				})
				return lst
			},
			length: 0,
		},
		{
			name: "readonly and legacy columns",
			getList: func() *models.List {
				lst := models.NewList()
				lst.SetColumns([]models.ColumnDefinitionable{
					contentTypeCd,
					attachmentCd,
					editCd,
					roCd,
				})
				return lst
			},
			length: 0,
		},
		{
			name: "legacy and a text column",
			getList: func() *models.List {
				lst := models.NewList()
				lst.SetColumns([]models.ColumnDefinitionable{
					contentTypeCd,
					attachmentCd,
					editCd,
					textCd,
				})
				return lst
			},
			length: 1,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			clonedList := ToListable(test.getList(), listName)
			require.NotEmpty(t, clonedList)

			cols := clonedList.GetColumns()
			assert.Len(t, cols, test.length)
		})
	}
}

func (suite *ListsUnitSuite) TestFieldValueSetable() {
	t := suite.T()

	additionalData := map[string]any{
		DescoratorFieldNamePrefix + "odata.etag":            "14fe12b2-e180-49f7-8fc3-5936f3dcf5d2,1",
		ReadOnlyOrHiddenFieldNamePrefix + "UIVersionString": "1.0",
		AuthorLookupIDColumnName:                            "6",
		EditorLookupIDColumnName:                            "6",
		"Item" + ChildCountFieldNamePart:                    "0",
		"Folder" + ChildCountFieldNamePart:                  "0",
		ModifiedColumnName:                                  "2023-12-13T15:47:51Z",
		CreatedColumnName:                                   "2023-12-13T15:47:51Z",
		EditColumnName:                                      "",
		LinkTitleFieldNamePart + "NoMenu":                   "Person1",
	}

	origFs := models.NewFieldValueSet()
	origFs.SetAdditionalData(additionalData)

	fs := retrieveFieldData(origFs)
	fsAdditionalData := fs.GetAdditionalData()
	assert.Empty(t, fsAdditionalData)

	additionalData["itemName"] = "item-1"
	origFs = models.NewFieldValueSet()
	origFs.SetAdditionalData(additionalData)

	fs = retrieveFieldData(origFs)
	fsAdditionalData = fs.GetAdditionalData()
	assert.NotEmpty(t, fsAdditionalData)

	val, ok := fsAdditionalData["itemName"]
	assert.True(t, ok)
	assert.Equal(t, "item-1", val)
}

func (suite *ListsUnitSuite) TestFieldValueSetable_Location() {
	t := suite.T()

	additionalData := map[string]any{
		"MyAddress": map[string]any{
			AddressFieldName: map[string]any{
				"city":            "Tagaytay",
				"countryOrRegion": "Philippines",
				"postalCode":      "4120",
				"state":           "Calabarzon",
				"street":          "Prime Residences CityLand 1852",
			},
			CoordinatesFieldName: map[string]any{
				"latitude":  "14.1153",
				"longitude": "120.962",
			},
			DisplayNameFieldName: "B123 Unit 1852 Prime Residences Tagaytay",
			LocationURIFieldName: "https://www.bingapis.com/api/v6/localbusinesses/YN8144x496766267081923032",
			UniqueIDFieldName:    "https://www.bingapis.com/api/v6/localbusinesses/YN8144x496766267081923032",
		},
		CountryOrRegionFieldName: "Philippines",
		StateFieldName:           "Calabarzon",
		CityFieldName:            "Tagaytay",
		PostalCodeFieldName:      "4120",
		StreetFieldName:          "Prime Residences CityLand 1852",
		GeoLocFieldName: map[string]any{
			"latitude":  14.1153,
			"longitude": 120.962,
		},
		DispNameFieldName: "B123 Unit 1852 Prime Residences Tagaytay",
	}

	expectedData := map[string]any{
		"MyAddress": map[string]any{
			AddressFieldName: map[string]any{
				"city":            "Tagaytay",
				"countryOrRegion": "Philippines",
				"postalCode":      "4120",
				"state":           "Calabarzon",
				"street":          "Prime Residences CityLand 1852",
			},
			CoordinatesFieldName: map[string]any{
				"latitude":  "14.1153",
				"longitude": "120.962",
			},
			DisplayNameFieldName: "B123 Unit 1852 Prime Residences Tagaytay",
			LocationURIFieldName: "https://www.bingapis.com/api/v6/localbusinesses/YN8144x496766267081923032",
			UniqueIDFieldName:    "https://www.bingapis.com/api/v6/localbusinesses/YN8144x496766267081923032",
		},
	}

	origFs := models.NewFieldValueSet()
	origFs.SetAdditionalData(additionalData)

	fs := retrieveFieldData(origFs)
	fsAdditionalData := fs.GetAdditionalData()
	assert.Equal(t, expectedData, fsAdditionalData)
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

	fieldsData, list := getFieldsDataAndList()

	err := writer.WriteObjectValue("", list)
	require.NoError(t, err)

	oldListByteArray, err := writer.GetSerializedContent()
	require.NoError(t, err)

	newList, err := acl.PostList(ctx, siteID, listName, oldListByteArray)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, listName, ptr.Val(newList.GetDisplayName()))

	_, err = acl.PostList(ctx, siteID, listName, oldListByteArray)
	require.Error(t, err)

	newListItems := newList.GetItems()
	require.Less(t, 0, len(newListItems))

	newListItemFields := newListItems[0].GetFields()
	require.NotEmpty(t, newListItemFields)

	newListItemsData := newListItemFields.GetAdditionalData()
	require.NotEmpty(t, newListItemsData)

	for k, v := range newListItemsData {
		assert.Equal(t, fieldsData[k], ptr.Val(v.(*string)))
	}

	err = acl.DeleteList(ctx, siteID, ptr.Val(newList.GetId()))
	require.NoError(t, err)
}

func (suite *ListsAPIIntgSuite) TestLists_PostList_invalidTemplate() {
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

	overrideListInfo := models.NewListInfo()
	overrideListInfo.SetTemplate(ptr.To(WebTemplateExtensionsListTemplateName))

	_, list := getFieldsDataAndList()
	list.SetList(overrideListInfo)

	err := writer.WriteObjectValue("", list)
	require.NoError(t, err)

	oldListByteArray, err := writer.GetSerializedContent()
	require.NoError(t, err)

	_, err = acl.PostList(ctx, siteID, listName, oldListByteArray)
	require.Error(t, err)
	assert.Equal(t, ErrCannotCreateWebTemplateExtension.Error(), err.Error())
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

	_, list := getFieldsDataAndList()

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

func getFieldsDataAndList() (map[string]any, *models.List) {
	oldListID := "old-list"
	listItemID := "list-item1"
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
	list.SetList(listInfo)
	list.SetColumns([]models.ColumnDefinitionable{txtColumnDef})
	list.SetItems([]models.ListItemable{listItem})

	return fieldsData, list
}
