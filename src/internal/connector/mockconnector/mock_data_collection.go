package mockconnector

import (
	"bytes"
	"io"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

// MockExchangeDataCollection represents a mock exchange mailbox
type MockExchangeDataCollection struct {
	fullPath     []string
	messageCount int
	Data         [][]byte
	Names        []string
}

var (
	_ data.Collection = &MockExchangeDataCollection{}
	_ data.Stream     = &MockExchangeData{}
	_ data.StreamInfo = &MockExchangeData{}
)

// NewMockExchangeDataCollection creates an data collection that will return the specified number of
// mock messages when iterated. Exchange type mail
func NewMockExchangeCollection(pathRepresentation []string, numMessagesToReturn int) *MockExchangeDataCollection {
	c := &MockExchangeDataCollection{
		fullPath:     pathRepresentation,
		messageCount: numMessagesToReturn,
		Data:         [][]byte{},
		Names:        []string{},
	}

	for i := 0; i < c.messageCount; i++ {
		// We can plug in whatever data we want here (can be an io.Reader to a test data file if needed)
		c.Data = append(c.Data, GetMockMessageBytes("From: NewMockExchangeCollection"))
		c.Names = append(c.Names, uuid.NewString())
	}

	return c
}

func (medc *MockExchangeDataCollection) FullPath() []string {
	return append([]string{}, medc.fullPath...)
}

// Items returns a channel that has the next items in the collection. The
// channel is closed when there are no more items available.
func (medc *MockExchangeDataCollection) Items() <-chan data.Stream {
	res := make(chan data.Stream)

	go func() {
		defer close(res)

		for i := 0; i < medc.messageCount; i++ {
			res <- &MockExchangeData{
				ID:     medc.Names[i],
				Reader: io.NopCloser(bytes.NewReader(medc.Data[i])),
			}
		}
	}()

	return res
}

// ExchangeData represents a single item retrieved from exchange
type MockExchangeData struct {
	ID      string
	Reader  io.ReadCloser
	ReadErr error
}

func (med *MockExchangeData) UUID() string {
	return med.ID
}

func (med *MockExchangeData) ToReader() io.ReadCloser {
	if med.ReadErr != nil {
		return io.NopCloser(errReader{med.ReadErr})
	}

	return med.Reader
}

func (med *MockExchangeData) Info() details.ItemInfo {
	return details.ItemInfo{
		Exchange: &details.ExchangeInfo{
			Sender:   "foo@bar.com",
			Subject:  "Hello world!",
			Received: time.Now(),
		},
	}
}

// GetMockMessageBytes returns bytes for Messageable item.
// Contents verified as working with sample data from kiota-serialization-json-go v0.5.5
func GetMockMessageBytes(subject string) []byte {
	userID := "foobar@8qzvrj.onmicrosoft.com"

	//nolint:lll
	message := "{\n    \"@odata.etag\": \"W/\\\"CQAAABYAAAB8wYc0thTTTYl3RpEYIUq+AAAZ0f0I\\\"\",\n    \"id\": \"AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEMAAB8wYc0thTTTYl3RpEYIUq_AAAZ3wG3AAA=\",\n    \"createdDateTime\": \"2022-04-08T18:08:02Z\",\n    \"lastModifiedDateTime\": \"2022-05-17T13:46:55Z\",\n    \"changeKey\": \"CQAAABYAAAB8wYc0thTTTYl3RpEYIUq+AAAZ0f0I\",\n    \"categories\": [],\n    \"receivedDateTime\": \"2022-04-08T18:08:02Z\",\n    \"sentDateTime\": \"2022-04-08T18:07:53Z\",\n    \"hasAttachments\": false,\n    \"internetMessageId\": \"<MWHPR1401MB1952C46D4A46B6398F562B0FA6E99@MWHPR1401MB1952.namprd14.prod.outlook.com>\",\n    \"subject\": \"" +
		//nolint:lll
		subject + " " + common.FormatNow(common.SimpleDateTimeFormat) + " Different\",\n    \"bodyPreview\": \"Who is coming to next week's party? I cannot imagine it is July soon\",\n    \"importance\": \"normal\",\n    \"parentFolderId\": \"AQMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4ADVkZWQwNmNlMTgALgAAAw_9XBStqZdPuOVIalVTz7sBAHzBhzS2FNNNiXdGkRghSr4AAAIBDAAAAA==\",\n    \"conversationId\": \"AAQkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOAAQAI7SSzmEPaRJsY-TWIALn1g=\",\n    \"conversationIndex\": \"AQHYS3N3jtJLOYQ9pEmxj9NYgAufWA==\",\n    \"isDeliveryReceiptRequested\": null,\n    \"isReadReceiptRequested\": false,\n    \"isRead\": true,\n    \"isDraft\": false,\n    \"webLink\": \"https://outlook.office365.com/owa/?ItemID=AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8%2B7BwB8wYc0thTTTYl3RpEYIUq%2BAAAAAAEMAAB8wYc0thTTTYl3RpEYIUq%2BAAAZ3wG3AAA%3D&exvsurl=1&viewmodel=ReadMessageItem\",\n    \"inferenceClassification\": \"focused\",\n    \"body\": {\n        \"contentType\": \"html\",\n        \"content\": \"<html><head><meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\"><meta name=\\\"Generator\\\" content=\\\"Microsoft Word 15 (filtered medium)\\\"><style><!--@font-face{font-family:\\\"Cambria Math\\\"}@font-face{font-family:Calibri}p.MsoNormal, li.MsoNormal, div.MsoNormal{margin:0in;font-size:11.0pt;font-family:\\\"Calibri\\\",sans-serif}span.EmailStyle17{font-family:\\\"Calibri\\\",sans-serif;color:windowtext}.MsoChpDefault{font-family:\\\"Calibri\\\",sans-serif}@page WordSection1{margin:1.0in 1.0in 1.0in 1.0in}div.WordSection1{}--></style></head><body lang=\\\"EN-US\\\" link=\\\"#0563C1\\\" vlink=\\\"#954F72\\\" style=\\\"word-wrap:break-word\\\"><div class=\\\"WordSection1\\\"><p class=\\\"MsoNormal\\\">I've been going through with the changing of messages. It shouldn't have the same calls, right? Call Me? </p><p class=\\\"MsoNormal\\\">&nbsp;</p><p class=\\\"MsoNormal\\\">We want to be able to send multiple messages and we want to be able to respond and do other things that make sense for our users. In this case. Let’s consider a Mailbox</p></div></body></html>\"\n    },\n    \"sender\": {\n        \"emailAddress\": {\n            \"name\": \"Lidia Holloway\",\n            \"address\": \"" + userID + "\"\n        }\n    },\n    \"from\": {\n        \"emailAddress\": {\n            \"name\": \"Lidia Holloway\",\n            \"address\": \"lidiah@8qzvrj.onmicrosoft.com\"\n        }\n    },\n    \"toRecipients\": [\n        {\n            \"emailAddress\": {\n                \"name\": \"Dustin Abbot\",\n                \"address\": \"dustina@8qzvrj.onmicrosoft.com\"\n            }\n        }\n    ],\n    \"ccRecipients\": [],\n    \"bccRecipients\": [],\n    \"replyTo\": [],\n    \"flag\": {\n        \"flagStatus\": \"notFlagged\"\n    }\n}\n"

	return []byte(message)
}

// GetMockContactBytes returns bytes for Contactable item.
func GetMockContactBytes(middleName string) []byte {
	//nolint:lll
	contact := "{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEOAADSEBNbUIB9RL6ePDeF3FIYAABS7DZnAAA=\",\"@odata.context\":\"https://graph.microsoft.com/v1.0/$metadata#users('foobar%408qzvrj.onmicrosoft.com')/contacts/$entity\",\"@odata.etag\":\"W/\\\"EQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAABSx4Tr\\\"\",\"categories\":[],\"changeKey\":\"EQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAABSx4Tr\",\"createdDateTime\":\"2019-08-04T06:55:33Z\",\"lastModifiedDateTime\":\"2019-08-04T06:55:33Z\",\"businessAddress\":{},\"businessPhones\":[],\"children\":[],\"displayName\":\"Santiago Quail\",\"emailAddresses\":[],\"fileAs\":\"Quail, Santiago\"," +
		//nolint:lll
		"\"givenName\":\"Santiago " + middleName + "\",\"homeAddress\":{},\"homePhones\":[],\"imAddresses\":[],\"otherAddress\":{},\"parentFolderId\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9FIYAAAAAAEOAAA=\",\"personalNotes\":\"\",\"surname\":\"Quail\"}"

	return []byte(contact)
}

// GetMockEventBytes returns test byte array representative of full Eventable item.
func GetMockEventBytes(subject string) []byte {
	newTime := time.Now().AddDate(0, 0, 1)
	conversion := common.FormatTime(newTime)
	timeSlice := strings.Split(conversion, "T")

	//nolint:lll
	event := "{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAENAADSEBNbUIB9RL6ePDeF3FIYAAAAAG76AAA=\",\"calendar@odata.navigationLink\":" +
		"\"https://graph.microsoft.com/v1.0/users('foobar@8qzvrj.onmicrosoft.com')/calendars('AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAENAAA=')\"," +
		"\"calendar@odata.associationLink\":\"https://graph.microsoft.com/v1.0/users('foobar@8qzvrj.onmicrosoft.com')/calendars('AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAENAAA=')/$ref\"," +
		"\"@odata.etag\":\"W/\\\"0hATW1CAfUS+njw3hdxSGAAAJIxNug==\\\"\",\"@odata.context\":\"https://graph.microsoft.com/v1.0/$metadata#users('foobar%408qzvrj.onmicrosoft.com')/events/$entity\",\"categories\":[],\"changeKey\":\"0hATW1CAfUS+njw3hdxSGAAAJIxNug==\"," +
		"\"createdDateTime\":\"2022-03-28T03:42:03Z\",\"lastModifiedDateTime\":\"2022-05-26T19:25:58Z\",\"allowNewTimeProposals\":true,\"attendees\"" +
		":[{\"emailAddress\":{\"address\":\"TestContact@8qzvrj.onmicrosoft.com\",\"name\":\"Fuu Gu\"},\"type\":\"required\",\"status\"" +
		":{\"response\":\"none\",\"time\":\"0001-01-01T00:00:00Z\"}},{\"emailAddress\":{\"address\":\"TestRoom@8qzvrj.onmicrosoft.com\",\"name\":\"Fuu Bar\"},\"type\":\"required\"" +
		",\"status\":{\"response\":\"none\",\"time\":\"0001-01-01T00:00:00Z\"}},{\"emailAddress\":{\"address\":\"testevents@8qzvrj.onmicrosoft.com\",\"name\":\"Ru Buu\"}," +
		"\"type\":\"required\",\"status\":{\"response\":\"none\",\"time\":\"0001-01-01T00:00:00Z\"}}],\"body\":{\"content\":\"<html>\\r\\n<head>\\r\\n<meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\">\\r\\n</head>\\r\\n<body>\\r\\n" +
		"<p>This meeting is to review the latest Tailspin Toys project proposal.<br>\\r\\nBut why not eat some sushi while we’re at it? :)</p>\\r\\n</body>\\r\\n</html>\\r\\n\"" +
		",\"contentType\":\"html\"},\"bodyPreview\":\"This meeting is to review the latest Tailspin Toys project proposal.\\r\\nBut why not eat some sushi while we’re at it? :)\"," +
		"\"end\":{\"dateTime\":\"" + timeSlice[0] + "T07:00:00.0000000\",\"timeZone\":\"UTC\"},\"hasAttachments\":false,\"hideAttendees\":false,\"iCalUId\":" +
		"\"040000008200E00074C5B7101A82E0080000000035723BC75542D801000000000000000010000000E1E7C8F785242E4894DA13AEFB947B85\",\"importance\":\"normal\",\"isAllDay\":false,\"isCancelled\":false," +
		"\"isDraft\":false,\"isOnlineMeeting\":false,\"isOrganizer\":false,\"isReminderOn\":true," +
		"\"location\":{\"displayName\":\"Umi Sake House (2230 1st Ave, Seattle, WA 98121 US)\",\"locationType\":\"default\",\"uniqueId\":\"Umi Sake House (2230 1st Ave, Seattle, WA 98121 US)\"," +
		"\"uniqueIdType\":\"private\"},\"locations\":[{\"displayName\":\"Umi Sake House (2230 1st Ave, Seattle, WA 98121 US)\",\"locationType\":\"default\",\"uniqueId\":\"\",\"uniqueIdType\":\"unknown\"}],\"onlineMeetingProvider\":\"unknown\",\"organizer\"" +
		":{\"emailAddress\":{\"address\":\"foobar3@8qzvrj.onmicrosoft.com\",\"name\":\"Anu Pierson\"}},\"originalEndTimeZone\":\"UTC\",\"originalStartTimeZone\":\"UTC\"," +
		"\"reminderMinutesBeforeStart\":15,\"responseRequested\":true,\"responseStatus\":{\"response\":\"notResponded\",\"time\":\"0001-01-01T00:00:00Z\"},\"sensitivity\":\"normal\",\"showAs\":\"tentative\"," +
		"\"start\":{\"dateTime\":\"" + timeSlice[0] + "T06:00:00.0000000\",\"timeZone\":\"UTC\"}," +
		"\"subject\":\" " + subject +
		" Review + Lunch\",\"type\":\"singleInstance\",\"webLink\":\"https://outlook.office365.com/owa/?itemid=AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAENAADSEBNbUIB9RL6ePDeF3FIYAAAAAG76AAA%3D&exvsurl=1&path=/calendar/item\"}"

	return []byte(event)
}

type errReader struct {
	readErr error
}

func (er errReader) Read([]byte) (int, error) {
	return 0, er.readErr
}
