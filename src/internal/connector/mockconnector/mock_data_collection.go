package mockconnector

import (
	"bytes"
	"io"
	"time"

	"github.com/google/uuid"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/internal/data"
	"github.com/alcionai/corso/internal/tester"
	"github.com/alcionai/corso/pkg/backup/details"
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
				medc.Names[i],
				io.NopCloser(bytes.NewReader(medc.Data[i])),
			}
		}
	}()

	return res
}

// ExchangeData represents a single item retrieved from exchange
type MockExchangeData struct {
	ID     string
	Reader io.ReadCloser
}

func (med *MockExchangeData) UUID() string {
	return med.ID
}

func (med *MockExchangeData) ToReader() io.ReadCloser {
	return med.Reader
}

func (med *MockExchangeData) Info() details.ItemInfo {
	return details.ItemInfo{Exchange: &details.ExchangeInfo{Sender: "foo@bar.com", Subject: "Hello world!", Received: time.Now()}}
}

// GetMockMessageBytes returns bytes for Messageable item.
// Contents verified as working with sample data from kiota-serialization-json-go v0.5.5
func GetMockMessageBytes(subject string) []byte {

	userID, err := tester.M365UserID()
	if err != nil {
		userID = "lidiah@8qzvrj.onmicrosoft.com"
	}

	message := "{\n    \"@odata.etag\": \"W/\\\"CQAAABYAAAB8wYc0thTTTYl3RpEYIUq+AAAZ0f0I\\\"\",\n    \"id\": \"AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEMAAB8wYc0thTTTYl3RpEYIUq_AAAZ3wG3AAA=\",\n    \"createdDateTime\": \"2022-04-08T18:08:02Z\",\n    \"lastModifiedDateTime\": \"2022-05-17T13:46:55Z\",\n    \"changeKey\": \"CQAAABYAAAB8wYc0thTTTYl3RpEYIUq+AAAZ0f0I\",\n    \"categories\": [],\n    \"receivedDateTime\": \"2022-04-08T18:08:02Z\",\n    \"sentDateTime\": \"2022-04-08T18:07:53Z\",\n    \"hasAttachments\": false,\n    \"internetMessageId\": \"<MWHPR1401MB1952C46D4A46B6398F562B0FA6E99@MWHPR1401MB1952.namprd14.prod.outlook.com>\",\n    \"subject\": \"" +
		subject + " " + common.FormatNow(common.SimpleDateTimeFormat) + " Different\",\n    \"bodyPreview\": \"Who is coming to next week's party? I cannot imagine it is July soon\",\n    \"importance\": \"normal\",\n    \"parentFolderId\": \"AQMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4ADVkZWQwNmNlMTgALgAAAw_9XBStqZdPuOVIalVTz7sBAHzBhzS2FNNNiXdGkRghSr4AAAIBDAAAAA==\",\n    \"conversationId\": \"AAQkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOAAQAI7SSzmEPaRJsY-TWIALn1g=\",\n    \"conversationIndex\": \"AQHYS3N3jtJLOYQ9pEmxj9NYgAufWA==\",\n    \"isDeliveryReceiptRequested\": null,\n    \"isReadReceiptRequested\": false,\n    \"isRead\": true,\n    \"isDraft\": false,\n    \"webLink\": \"https://outlook.office365.com/owa/?ItemID=AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8%2B7BwB8wYc0thTTTYl3RpEYIUq%2BAAAAAAEMAAB8wYc0thTTTYl3RpEYIUq%2BAAAZ3wG3AAA%3D&exvsurl=1&viewmodel=ReadMessageItem\",\n    \"inferenceClassification\": \"focused\",\n    \"body\": {\n        \"contentType\": \"html\",\n        \"content\": \"<html><head><meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\"><meta name=\\\"Generator\\\" content=\\\"Microsoft Word 15 (filtered medium)\\\"><style><!--@font-face{font-family:\\\"Cambria Math\\\"}@font-face{font-family:Calibri}p.MsoNormal, li.MsoNormal, div.MsoNormal{margin:0in;font-size:11.0pt;font-family:\\\"Calibri\\\",sans-serif}span.EmailStyle17{font-family:\\\"Calibri\\\",sans-serif;color:windowtext}.MsoChpDefault{font-family:\\\"Calibri\\\",sans-serif}@page WordSection1{margin:1.0in 1.0in 1.0in 1.0in}div.WordSection1{}--></style></head><body lang=\\\"EN-US\\\" link=\\\"#0563C1\\\" vlink=\\\"#954F72\\\" style=\\\"word-wrap:break-word\\\"><div class=\\\"WordSection1\\\"><p class=\\\"MsoNormal\\\">I've been going through with the changing of messages. It shouldn't have the same calls, right? Call Me? </p><p class=\\\"MsoNormal\\\">&nbsp;</p><p class=\\\"MsoNormal\\\">We want to be able to send multiple messages and we want to be able to respond and do other things that make sense for our users. In this case. Let’s consider a Mailbox</p></div></body></html>\"\n    },\n    \"sender\": {\n        \"emailAddress\": {\n            \"name\": \"Lidia Holloway\",\n            \"address\": \"" + userID + "\"\n        }\n    },\n    \"from\": {\n        \"emailAddress\": {\n            \"name\": \"Lidia Holloway\",\n            \"address\": \"lidiah@8qzvrj.onmicrosoft.com\"\n        }\n    },\n    \"toRecipients\": [\n        {\n            \"emailAddress\": {\n                \"name\": \"Dustin Abbot\",\n                \"address\": \"dustina@8qzvrj.onmicrosoft.com\"\n            }\n        }\n    ],\n    \"ccRecipients\": [],\n    \"bccRecipients\": [],\n    \"replyTo\": [],\n    \"flag\": {\n        \"flagStatus\": \"notFlagged\"\n    }\n}\n"

	return []byte(message)
}
