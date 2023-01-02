package mockconnector

import (
	"bytes"
	"context"
	"io"
	"testing"

	kw "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/logger"
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
// information. The max amount of unique values is:
// If count exceeds max amount, the returned list will not be able to be loaded into
// the M365 back store.
// Source: https://learn.microsoft.com/en-us/graph/api/list-create?view=graph-rest-1.0&tabs=go
func GetMockList(title string, count int) models.Listable {
	requestBody := models.NewList()
	requestBody.SetDisplayName(&title)
	requestBody.SetName(&title)

	columnDef := models.NewColumnDefinition()
	name := "Artist"
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

	if count == 0 {
		return requestBody
	}

	// item Creation
	index := 0
	itms := make([]models.ListItemable, 0)

	itemMap := getItems()
	if count > len(itemMap) {
		logger.Ctx(context.TODO()).Info("returned mock list cannot be uploaded. Non-unique items included")
	}

	for k, v := range getItems() {
		temp := models.NewListItem()
		fields := models.NewFieldValueSet()
		entry := map[string]interface{}{
			"Title":  k,
			"Artist": v,
		}

		fields.SetAdditionalData(entry)

		itms = append(itms, temp)
		index++

		if index == count {
			break
		}
	}

	requestBody.SetItems(itms)

	return requestBody
}

// GetMockListBytes returns the byte representation of GetMockList
func GetMockListBytes(title string, numOfItems int) ([]byte, error) {
	list := GetMockList(title, numOfItems)

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
	byteArray, err := GetMockListBytes(title, numOfItems)
	require.NoError(t, err)

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
		"Who's Next":                      "The Who",
		"Led Zeppelin":                    "Led Zeppelin",
		"Blue":                            "Joni Mitchell",
		"Bringing It All Back Home":       "Bob Dylan",
		"Let It Bleed":                    "The Rolling Stones",
		"Ramones":                         "Ramones",
		"Music From Big Pink":             "The Band",
		"The Rise and Fall of Ziggy Stardust and the Spiders From Mars": "David Bowie",
		"Tapestry":         "Carole King",
		"Hotel California": "Eagles",
		"The Anthology":    "Muddy Waters",
		"Please Please Me": "The Beatles",
		"Forever Changes":  "Love",
		"Never Mind the Bollocks Here's the Sex Pistols": "Sex Pistols",
		"The Doors":                    "The Doors",
		"The Dark Side of the Moon":    "Pink Floyd",
		"Horses":                       "Patti Smith",
		"The Band `(The Brown Album)`": "The Band",
		"Legend: The Best of Bob Marley and The Wailers": "Bob Marley & The Wailers",
		"A Love Supreme": "John Coltrane",
		"It Takes a Nation of Millions to Hold Us Back": "Public Enemy",
		"At Fillmore East":                       "The Allman Brothers Band",
		"Here's Little Richard":                  "Little Richard",
		"Bridge Over Troubled Water":             "Simon & Garfunkel",
		"Greatest Hits":                          "Al Green",
		"Meet The Beatles!":                      "The Beatles",
		"The Birth of Soul":                      "Ray Charles",
		"Electric Ladyland":                      "The Jimi Hendrix Experience",
		"Elvis Presley":                          "Elvis Presley",
		"Songs in the Key of Life":               "Stevie Wonder",
		"Beggars Banquet":                        "The Rolling Stones",
		"Trout Mask Replica":                     "Captain Beefheart & His Magic Band",
		"Appetite for Destruction":               "Guns N' Roses",
		"Achtung Baby":                           "U2",
		"Sticky Fingers":                         "The Rolling Stones",
		"Back to Mono (1958-1969)":               "Phil Spector",
		"Moondance":                              "Van Morrison",
		"Kid A":                                  "Radiohead",
		"Off the Wall":                           "Michael Jackson",
		"[Led Zeppelin IV]":                      "Led Zeppelin",
		"The Stranger":                           "Billy Joel",
		"Graceland":                              "Paul Simon",
		"Superfly":                               "Curtis Mayfield",
		"Physical Graffiti":                      "Led Zeppelin",
		"After the Gold Rush":                    "Neil Young",
		"Star Time":                              "James Brown",
		"Purple Rain":                            "Prince and the Revolution",
		"Back in Black":                          "AC/DC",
		"Otis Blue: Otis Redding Sings Soul":     "Otis Redding",
		"Led Zeppelin II":                        "Led Zeppelin",
		"Imagine":                                "John Lennon",
		"The Clash":                              "The Clash",
		"Harvest":                                "Neil Young",
		"Axis: Bold as Love":                     "The Jimi Hendrix Experience",
		"I Never Loved a Man the Way I Love You": "Aretha Franklin",
		"Lady Soul":                              "Aretha Franklin",
		"Born in the U.S.A.":                     "Bruce Springsteen",
		"The Wall":                               "Pink Floyd",
		"At Folsom Prison":                       "Johnny Cash",
		"Dusty in Memphis":                       "Dusty Springfield",
		"Talking Book":                           "Stevie Wonder",
		"Goodbye Yellow Brick Road":              "Elton John",
		"20 Golden Greats":                       "Buddy Holly",
		"Sign 'Peace' the Times":                 "Prince",
		"40 Greatest Hits":                       "Hank Williams",
		"Tommy":                                  "The Who",
		"The Freewheelin' Bob Dylan":             "Bob Dylan",
		"This Year's Model":                      "Elvis Costello",
		"There's a Riot Goin' On":                "Sly & The Family Stone",
		"Odessey and Oracle":                     "The Zombies",
	}

	return items
}
