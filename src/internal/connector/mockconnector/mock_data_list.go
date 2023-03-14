package mockconnector

import (
	"bytes"
	"context"
	"io"
	"testing"

	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	_ data.Stream           = &MockListData{}
	_ data.BackupCollection = &MockListCollection{}
)

type MockListCollection struct {
	fullPath path.Path
	Data     []*MockListData
	Names    []string
}

func (mlc *MockListCollection) SetPath(p path.Path) {
	mlc.fullPath = p
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

func (mlc *MockListCollection) Items(
	ctx context.Context,
	_ *fault.Bus, // unused
) <-chan data.Stream {
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

// GetMockList returns a Listable object with two columns.
// @param: Name of the displayable list
// @param: Column Name: Defines the 2nd Column Name of the created list the values from the map.
// The key values of the input map are used for the `Title` column.
// The values of the map are placed within the 2nd column.
// Source: https://learn.microsoft.com/en-us/graph/api/list-create?view=graph-rest-1.0&tabs=go
func GetMockList(title, columnName string, items map[string]string) models.Listable {
	requestBody := models.NewList()
	requestBody.SetDisplayName(&title)
	requestBody.SetName(&title)

	columnDef := models.NewColumnDefinition()
	name := columnName
	text := models.NewTextColumn()

	columnDef.SetName(&name)
	columnDef.SetText(text)

	columns := []models.ColumnDefinitionable{
		columnDef,
	}
	requestBody.SetColumns(columns)

	aList := models.NewListInfo()
	template := "genericList"
	aList.SetTemplate(&template)
	requestBody.SetList(aList)

	// item Creation
	itms := make([]models.ListItemable, 0)

	for k, v := range items {
		entry := map[string]interface{}{
			"Title":    k,
			columnName: v,
		}

		fields := models.NewFieldValueSet()
		fields.SetAdditionalData(entry)

		temp := models.NewListItem()
		temp.SetFields(fields)

		itms = append(itms, temp)
	}

	requestBody.SetItems(itms)

	return requestBody
}

// GetMockListDefault returns a two-list column list of
// Music lbums and the associated artist.
func GetMockListDefault(title string) models.Listable {
	return GetMockList(title, "Artist", getItems())
}

// GetMockListBytes returns the byte representation of GetMockList
func GetMockListBytes(title string) ([]byte, error) {
	list := GetMockListDefault(title)

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
func GetMockListStream(t *testing.T, title string, numOfItems int) *MockListData {
	byteArray, err := GetMockListBytes(title)
	require.NoError(t, err, clues.ToCore(err))

	listData := &MockListData{
		ID:     title,
		Reader: io.NopCloser(bytes.NewReader(byteArray)),
		size:   int64(len(byteArray)),
	}

	return listData
}

// getItems returns a map where key values are albums
// and values are the artist.
// Source: https://github.com/Currie32/500-Greatest-Albums/blob/master/albumlist.csv
func getItems() map[string]string {
	items := map[string]string{
		"London Calling":                  "The Clash",
		"Blonde on Blonde":                "Bob Dylan",
		"The Beatles '(The White Album)'": "The Beatles",
		"The Sun Sessions":                "Elvis Presley",
		"Kind of Blue":                    "Miles Davis",
		"The Velvet Underground & Nico":   "The Velvet Underground",
		"Abbey Road":                      "The Beatles",
		"Are You Experienced":             "The Jimi Hendrix Experience",
		"Blood on the Tracks":             "Bob Dylan",
		"Nevermind":                       "Nirvana",
		"Born to Run":                     "Bruce Springsteen",
		"Astral Weeks":                    "Van Morrison",
		"Thriller":                        "Michael Jackson",
		"The Great Twenty_Eight":          "Chuck Berry",
		"The Complete Recordings":         "Robert Johnson",
		"John Lennon/Plastic Ono Band":    "John Lennon / Plastic Ono Band",
		"Innervisions":                    "Stevie Wonder",
		"Live at the Apollo, 1962":        "James Brown",
		"Rumours":                         "Fleetwood Mac",
		"The Joshua Tree":                 "U2",
	}

	return items
}
