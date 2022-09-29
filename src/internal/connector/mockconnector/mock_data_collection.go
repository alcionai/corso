package mockconnector

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/path"
)

//nolint:lll
const (
	defaultMessageBody = "<span class=\\\"x_elementToProof ContentPasted0\\\" style=\\\"font-size:12pt;" +
		" margin:0px; background-color:rgb(255,255,255)\\\">Lidia,</span> <div class=\\\"x_elementToProof\\\" style=\\\"font-size:12pt; margin:0px; background-color:rgb(255,255,255)\\\"><br class=\\\"ContentPasted0\\\"></div><div class=\\\"x_elementToProof ContentPasted0\\\" style=\\\"font-size:12pt;" +
		" margin:0px; background-color:rgb(255,255,255)\\\">We have not received any reports on the development during Q2. It is in our best interest to have a new TPS Report by next Thursday prior to the retreat. If you have any questions, please let me know so I can address them.</div>" +
		"<div class=\\\"x_elementToProof\\\" style=\\\"font-size:12pt; margin:0px; background-color:rgb(255,255,255)\\\"><br class=\\\"ContentPasted0\\\"></div><div class=\\\"x_elementToProof ContentPasted0\\\" style=\\\"font-size:12pt; margin:0px; background-color:rgb(255,255,255)\\\">Thanking you in advance,</div>" +
		"<div class=\\\"x_elementToProof\\\" style=\\\"font-size:12pt; margin:0px; background-color:rgb(255,255,255)\\\"><br class=\\\"ContentPasted0\\\"></div><span class=\\\"x_elementToProof ContentPasted0\\\" style=\\\"font-size:12pt; margin:0px; background-color:rgb(255,255,255)\\\">Dustin</span><br>"
	defaultMessagePreview = "Lidia,\\n\\nWe have not received any reports on the development during Q2. It is in our best interest to have a new TPS Report by next Thursday prior to the retreat. If you have any questions, please let me know so I can address them.\\n" +
		"\\nThanking you in adv"

	// Order of fields to fill in:
	//   1. message body
	//   2. message preview
	//   3. sender user ID
	//   4. subject
	messageTmpl = "{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB3XwIkAAA=\",\"@odata.context\":\"https://graph.microsoft.com/v1.0/$metadata#users('a4a472f8-ccb0-43ec-bf52-3697a91b926c')/messages/$entity\"," +
		"\"@odata.etag\":\"W/\\\"CQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAAB2ZxqU\\\"\",\"categories\":[],\"changeKey\":\"CQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAAB2ZxqU\",\"createdDateTime\":\"2022-09-26T23:15:50Z\",\"lastModifiedDateTime\":\"2022-09-26T23:15:51Z\",\"bccRecipients\":[],\"body\":{\"content\":\"<html><head>" +
		"\\n<meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\"><style type=\\\"text/css\\\" style=\\\"display:none\\\">\\n<!--\\np\\n{margin-top:0;\\nmargin-bottom:0}\\n-->" +
		"\\n</style></head><body dir=\\\"ltr\\\"><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\">%s" +
		"</div></body></html>\",\"contentType\":\"html\"}," +
		"\"bodyPreview\":\"%s\"," +
		"\"ccRecipients\":[],\"conversationId\":\"AAQkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAQAK5nNWRdNWpGpLp7Xpb-m7A=\",\"conversationIndex\":\"AQHY0f3Ermc1ZF01akakuntelv+bsA==\",\"flag\":{\"flagStatus\":\"notFlagged\"}," +
		"\"from\":{\"emailAddress\":{\"address\":\"%s\",\"name\":\"A Stranger\"}},\"hasAttachments\":false,\"importance\":\"normal\",\"inferenceClassification\":\"focused\",\"internetMessageId\":\"<SJ0PR17MB562266A1E61A8EA12F5FB17BC3529@SJ0PR17MB5622.namprd17.prod.outlook.com>\"," +
		"\"isDeliveryReceiptRequested\":false,\"isDraft\":false,\"isRead\":false,\"isReadReceiptRequested\":false,\"parentFolderId\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAAA=\",\"receivedDateTime\":\"2022-09-26T23:15:50Z\"," +
		"\"replyTo\":[],\"sender\":{\"emailAddress\":{\"address\":\"foobar@8qzvrj.onmicrosoft.com\",\"name\":\"A Stranger\"}},\"sentDateTime\":\"2022-09-26T23:15:46Z\"," +
		"\"subject\":\"%s\",\"toRecipients\":[{\"emailAddress\":{\"address\":\"LidiaH@8qzvrj.onmicrosoft.com\",\"name\":\"Lidia Holloway\"}}]," +
		"\"webLink\":\"https://outlook.office365.com/owa/?ItemID=AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB3XwIkAAA%%3D&exvsurl=1&viewmodel=ReadMessageItem\"}"
)

// MockExchangeDataCollection represents a mock exchange mailbox
type MockExchangeDataCollection struct {
	fullPath     path.Path
	messageCount int
	Data         [][]byte
	Names        []string
}

var (
	_ data.Collection = &MockExchangeDataCollection{}
	_ data.Stream     = &MockExchangeData{}
	_ data.StreamInfo = &MockExchangeData{}
	_ data.StreamSize = &MockExchangeData{}
)

// NewMockExchangeDataCollection creates an data collection that will return the specified number of
// mock messages when iterated. Exchange type mail
func NewMockExchangeCollection(pathRepresentation path.Path, numMessagesToReturn int) *MockExchangeDataCollection {
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

// NewMockExchangeDataCollection creates an data collection that will return the specified number of
// mock messages when iterated. Exchange type mail
func NewMockContactCollection(pathRepresentation path.Path, numMessagesToReturn int) *MockExchangeDataCollection {
	c := &MockExchangeDataCollection{
		fullPath:     pathRepresentation,
		messageCount: numMessagesToReturn,
		Data:         [][]byte{},
		Names:        []string{},
	}

	rand.Seed(time.Now().UnixNano())

	middleNames := []string{
		"Argon",
		"Bernard",
		"Carleton",
		"Daphenius",
		"Ernesto",
		"Farraday",
		"Ghimley",
		"Irgot",
		"Jannes",
		"Knox",
		"Levi",
		"Milton",
	}

	for i := 0; i < c.messageCount; i++ {
		// We can plug in whatever data we want here (can be an io.Reader to a test data file if needed)
		c.Data = append(c.Data, GetMockContactBytes(middleNames[rand.Intn(len(middleNames))]))
		c.Names = append(c.Names, uuid.NewString())
	}

	return c
}

func (medc *MockExchangeDataCollection) FullPath() path.Path {
	return medc.fullPath
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
				size:   int64(len(medc.Data[i])),
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
	size    int64
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

func (med *MockExchangeData) Size() int64 {
	return med.size
}

// GetMockMessageBytes returns bytes for Messageable item.
// Contents verified as working with sample data from kiota-serialization-json-go v0.5.5
func GetMockMessageBytes(subject string) []byte {
	userID := "foobar@8qzvrj.onmicrosoft.com"
	timestamp := " " + common.FormatNow(common.SimpleDateTimeFormat)

	message := fmt.Sprintf(
		messageTmpl,
		defaultMessageBody,
		defaultMessagePreview,
		userID,
		"TPS Report "+subject+timestamp,
	)

	return []byte(message)
}

func GetMockMessageWithBodyBytes(subject, body string) []byte {
	userID := "foobar@8qzvrj.onmicrosoft.com"
	preview := body

	if len(preview) > 255 {
		preview = preview[:256]
	}

	message := fmt.Sprintf(
		messageTmpl,
		body,
		preview,
		userID,
		subject,
	)

	return []byte(message)
}

// GetMockContactBytes returns bytes for Contactable item.
// When hydrated: contact.GetGivenName() shows differences
func GetMockContactBytes(middleName string) []byte {
	phone := generatePhoneNumber()
	//nolint:lll
	contact := "{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEOAADSEBNbUIB9RL6ePDeF3FIYAABS7DZnAAA=\",\"@odata.context\":\"https://graph.microsoft.com/v1.0/$metadata#users('foobar%408qzvrj.onmicrosoft.com')/contacts/$entity\"," +
		"\"@odata.etag\":\"W/\\\"EQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAABSx4Tr\\\"\",\"categories\":[],\"changeKey\":\"EQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAABSx4Tr\",\"createdDateTime\":\"2019-08-04T06:55:33Z\",\"lastModifiedDateTime\":\"2019-08-04T06:55:33Z\",\"businessAddress\":{},\"businessPhones\":[],\"children\":[]," +
		"\"displayName\":\"Santiago Quail\",\"emailAddresses\":[],\"fileAs\":\"Quail, Santiago\",\"mobilePhone\": \"" + phone + "\"," +
		"\"givenName\":\"Santiago\",\"homeAddress\":{},\"homePhones\":[],\"imAddresses\":[],\"otherAddress\":{},\"parentFolderId\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9FIYAAAAAAEOAAA=\",\"personalNotes\":\"\",\"middleName\":\"" + middleName + "\",\"surname\":\"Quail\"}"

	return []byte(contact)
}

// generatePhoneNumber creates a random phone number
// @return string representation in format (xxx)xxx-xxxx
func generatePhoneNumber() string {
	numbers := make([]string, 0)

	for i := 0; i < 10; i++ {
		temp := rand.Intn(10)
		value := strconv.Itoa(temp)
		numbers = append(numbers, value)
	}

	area := strings.Join(numbers[:3], "")
	prefix := strings.Join(numbers[3:6], "")
	suffix := strings.Join(numbers[6:], "")
	phoneNo := "(" + area + ")" + prefix + "-" + suffix

	return phoneNo
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
		":[],\"body\":{\"content\":\"<html>\\r\\n<head>\\r\\n<meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\">\\r\\n</head>\\r\\n<body>\\r\\n" +
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

func GetMockEventWithAttendeesBytes(subject string) []byte {
	newTime := time.Now().AddDate(0, 0, 1)
	conversion := common.FormatTime(newTime)
	timeSlice := strings.Split(conversion, "T")

	//nolint:lll
	event := "{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAENAADSEBNbUIB9RL6ePDeF3FIYAABU_FdvAAA=\",\"@odata.etag\":\"W/\\\"0hATW1CAfUS+njw3hdxSGAAAVK7j9A==\\\"\"," +
		"\"calendar@odata.associationLink\":\"https://graph.microsoft.com/v1.0/users('a4a472f8-ccb0-43ec-bf52-3697a91b926c')/calendars('AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAENAAA=')/$ref\"," +
		"\"calendar@odata.navigationLink\":\"https://graph.microsoft.com/v1.0/users('a4a472f8-ccb0-43ec-bf52-3697a91b926c')/calendars('AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAENAAA=')\"," +
		"\"@odata.context\":\"https://graph.microsoft.com/v1.0/$metadata#users('a4a472f8-ccb0-43ec-bf52-3697a91b926c')/events/$entity\",\"categories\":[],\"changeKey\":\"0hATW1CAfUS+njw3hdxSGAAAVK7j9A==\",\"createdDateTime\":\"2022-08-06T12:47:56Z\",\"lastModifiedDateTime\":\"2022-08-06T12:49:59Z\",\"allowNewTimeProposals\":true," +
		"\"attendees\":[{\"emailAddress\":{\"address\":\"george.martinez@8qzvrj.onmicrosoft.com\",\"name\":\"George Martinez\"},\"type\":\"required\",\"status\":{\"response\":\"none\",\"time\":\"0001-01-01T00:00:00Z\"}},{\"emailAddress\":{\"address\":\"LeeG@8qzvrj.onmicrosoft.com\",\"name\":\"Lee Gu\"},\"type\":\"required\",\"status\":{\"response\":\"none\",\"time\":\"0001-01-01T00:00:00Z\"}}]," +
		"\"body\":{\"content\":\"<html>\\n<head>\\n<meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\">\\n</head>\\n<body>\\n<div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\\\">\\nDiscuss matters concerning stock options and early release of quarterly earnings.</div>\\n<br> " +
		"\\n<div style=\\\"width:100%; height:20px\\\"><span style=\\\"white-space:nowrap; color:#5F5F5F; opacity:.36\\\">________________________________________________________________________________</span>\\n</div>\\n<div class=\\\"me-email-text\\\" lang=\\\"en-GB\\\" style=\\\"color:#252424; font-family:'Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif\\\"> " +
		"\\n<div style=\\\"margin-top:24px; margin-bottom:20px\\\"><span style=\\\"font-size:24px; color:#252424\\\">Microsoft Teams meeting</span>\\n</div>\\n<div style=\\\"margin-bottom:20px\\\">\\n<div style=\\\"margin-top:0px; margin-bottom:0px; font-weight:bold\\\"><span style=\\\"font-size:14px; color:#252424\\\">Join on your computer or mobile app</span>" +
		"\\n</div>\\n<a href=\\\"https://teams.microsoft.com/l/meetup-join/19%3ameeting_YWNhMzAxZjItMzE2My00ZGQzLTkzMDUtNjQ3NTY0NjNjMTZi%40thread.v2/0?context=%7b%22Tid%22%3a%224d603060-18d6-4764-b9be-4cb794d32b69%22%2c%22Oid%22%3a%22a4a472f8-ccb0-43ec-bf52-3697a91b926c%22%7d\\\" class=\\\"me-email-headline\\\" style=\\\"font-size:14px; font-family:'Segoe UI Semibold','Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif; text-decoration:underline; color:#6264a7\\\">Click\\n here to join the meeting</a> </div>" +
		"\\n<div style=\\\"margin-bottom:20px; margin-top:20px\\\">\\n<div style=\\\"margin-bottom:4px\\\"><span style=\\\"font-size:14px; color:#252424\\\">Meeting ID:\\n<span style=\\\"font-size:16px; color:#252424\\\">292 784 521 247</span> </span><br>\\n<span style=\\\"font-size:14px; color:#252424\\\">Passcode: </span><span style=\\\"font-size:16px; color:#252424\\\">SzBkfK\\n</span>" +
		"\\n<div style=\\\"font-size:14px\\\"><a href=\\\"https://www.microsoft.com/en-us/microsoft-teams/download-app\\\" class=\\\"me-email-link\\\" style=\\\"font-size:14px; text-decoration:underline; color:#6264a7; font-family:'Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif\\\">Download\\n Teams</a> | <a href=\\\"https://www.microsoft.com/microsoft-teams/join-a-meeting\\\" class=\\\"me-email-link\\\" style=\\\"font-size:14px; text-decoration:underline; color:#6264a7; font-family:'Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif\\\">" +
		"\\nJoin on the web</a></div>\\n</div>\\n</div>\\n<div style=\\\"margin-bottom:24px; margin-top:20px\\\"><a href=\\\"https://aka.ms/JoinTeamsMeeting\\\" class=\\\"me-email-link\\\" style=\\\"font-size:14px; text-decoration:underline; color:#6264a7; font-family:'Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif\\\">Learn more</a>" +
		"\\n | <a href=\\\"https://teams.microsoft.com/meetingOptions/?organizerId=a4a472f8-ccb0-43ec-bf52-3697a91b926c&amp;tenantId=4d603060-18d6-4764-b9be-4cb794d32b69&amp;threadId=19_meeting_YWNhMzAxZjItMzE2My00ZGQzLTkzMDUtNjQ3NTY0NjNjMTZi@thread.v2&amp;messageId=0&amp;language=en-GB\\\" class=\\\"me-email-link\\\" style=\\\"font-size:14px; text-decoration:underline; color:#6264a7; font-family:'Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif\\\">" +
		"\\nMeeting options</a> </div>\\n</div>\\n<div style=\\\"font-size:14px; margin-bottom:4px; font-family:'Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif\\\">\\n</div>\\n<div style=\\\"font-size:12px\\\"></div>\\n<div></div>\\n<div style=\\\"width:100%; height:20px\\\"><span style=\\\"white-space:nowrap; color:#5F5F5F; opacity:.36\\\">________________________________________________________________________________</span>" +
		"\\n</div>\\n</body>\\n</html>\\n\",\"contentType\":\"html\"},\"bodyPreview\":\"Discuss matters concerning stock options and early release of quarterly earnings.\\n\\n\", " +
		"\"end\":{\"dateTime\":\"" + timeSlice[0] + "T16:00:00.0000000\",\"timeZone\":\"UTC\"},\"hasAttachments\":false,\"hideAttendees\":false,\"iCalUId\":\"040000008200E00074C5B7101A82E0080000000010A45EC092A9D801000000000000000010000000999C7C6281C2B24A91D5502392B8EF38\",\"importance\":\"normal\",\"isAllDay\":false,\"isCancelled\":false,\"isDraft\":false,\"isOnlineMeeting\":true,\"isOrganizer\":true,\"isReminderOn\":true," +
		"\"location\":{\"address\":{},\"coordinates\":{},\"displayName\":\"\",\"locationType\":\"default\",\"uniqueIdType\":\"unknown\"},\"locations\":[],\"onlineMeeting\":{\"joinUrl\":\"https://teams.microsoft.com/l/meetup-join/19%3ameeting_YWNhMzAxZjItMzE2My00ZGQzLTkzMDUtNjQ3NTY0NjNjMTZi%40thread.v2/0?context=%7b%22Tid%22%3a%224d603060-18d6-4764-b9be-4cb794d32b69%22%2c%22Oid%22%3a%22a4a472f8-ccb0-43ec-bf52-3697a91b926c%22%7d\"},\"onlineMeetingProvider\":\"teamsForBusiness\"," +
		"\"organizer\":{\"emailAddress\":{\"address\":\"LidiaH@8qzvrj.onmicrosoft.com\",\"name\":\"Lidia Holloway\"}},\"originalEndTimeZone\":\"Eastern Standard Time\",\"originalStartTimeZone\":\"Eastern Standard Time\",\"reminderMinutesBeforeStart\":15,\"responseRequested\":true,\"responseStatus\":{\"response\":\"organizer\",\"time\":\"0001-01-01T00:00:00Z\"},\"sensitivity\":\"normal\",\"showAs\":\"busy\"," +
		"\"start\":{\"dateTime\":\"" + timeSlice[0] + "T15:30:00.0000000\",\"timeZone\":\"UTC\"},\"subject\":\"Board " + subject + " Meeting\",\"transactionId\":\"28b36295-6cd3-952f-d8f5-deb313444a51\",\"type\":\"singleInstance\",\"webLink\":\"https://outlook.office365.com/owa/?itemid=AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAENAADSEBNbUIB9RL6ePDeF3FIYAABU%2BFdvAAA%3D&exvsurl=1&path=/calendar/item\"}"

	return []byte(event)
}

type errReader struct {
	readErr error
}

func (er errReader) Read([]byte) (int, error) {
	return 0, er.readErr
}
