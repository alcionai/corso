package mockconnector

import (
	"bytes"
	"io"
	"testing"

	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	_ data.Stream     = &MockListData{}
	_ data.Collection = &MockListCollection{}
)

type MockListCollection struct {
	fullPath path.Path
	Data     []*MockListData
	Names    []string
}

func (mlc *MockListCollection) State() data.CollectionState {
	return data.NewState
}

func (mlc *MockListCollection) FullPath() path.Path {
	return mlc.fullPath
}

func (mlc *MockListCollection) DoNotMergeItems() bool {
	return false
}

func (mlc *MockListCollection) PreviousPath() path.Path {
	return nil
}

func (mlc *MockListCollection) Items() <-chan data.Stream {
	res := make(chan data.Stream)

	go func() {
		defer close(res)

		for _, stream := range mlc.Data {
			res <- stream
		}
	}()

	return res
}

type MockListData struct {
	ID      string
	Reader  io.ReadCloser
	ReadErr error
	size    int64
	deleted bool
}

func (mld *MockListData) UUID() string {
	return mld.ID
}

func (mld MockListData) Deleted() bool {
	return mld.deleted
}

func (mld *MockListData) ToReader() io.ReadCloser {
	return mld.Reader
}

// GetMockList returns a Listable object with generic
// information.
// Source: https://learn.microsoft.com/en-us/graph/api/list-create?view=graph-rest-1.0&tabs=go
func GetMockList(title string) models.Listable {
	requestBody := models.NewList()
	requestBody.SetDisplayName(&title)
	requestBody.SetName(&title)

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

// GetMockListStream returns the data.Stream representation
// of the Mocked SharePoint List
func GetMockListStream(t *testing.T, title string) *MockListData {
	byteArray, err := GetMockListBytes(title)
	require.NoError(t, err)

	listData := &MockListData{
		ID:     title,
		Reader: io.NopCloser(bytes.NewReader(byteArray)),
		size:   int64(len(byteArray)),
	}

	return listData
}
