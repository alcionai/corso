package mockconnector

import (
	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

// GetMockList returns a Listable object with generic
// information.
// Source: https://learn.microsoft.com/en-us/graph/api/list-create?view=graph-rest-1.0&tabs=go
func GetMockList(title string) models.Listable {
	requestBody := models.NewList()
	requestBody.SetDisplayName(&title)

	columnDef := models.NewColumnDefinition()
	name := "Author"
	text := models.NewTextColumn()

	columnDef.SetName(&name)
	columnDef.SetText(text)

	columnDef2 := models.NewColumnDefinition()
	name2 := "PageCount"
	number := models.NewNumberColumn()

	columnDef2.SetName(&name2)
	columnDef2.SetNumber(number)

	columns := []models.ColumnDefinitionable{
		columnDef,
		columnDef2,
	}
	requestBody.SetColumns(columns)

	aList := models.NewListInfo()
	template := "genericList"
	aList.SetTemplate(&template)
	requestBody.SetList(aList)

	return requestBody
}

// GetMockListBytes returns the byte representation of GetMockList
func GetMockListBytes(title string) ([]byte, error) {
	list := GetMockList(title)

	objectWriter := kw.NewJsonSerializationWriter()
	defer objectWriter.Close()

	err := objectWriter.WriteObjectValue("", list)
	if err != nil {
		return nil, err
	}

	return objectWriter.GetSerializedContent()
}
