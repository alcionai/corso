package mockconnector

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/alcionai/corso/src/internal/common"
)

// Order of fields to fill in:
// 1. body
// 2. body preview
// 3. endDate (rfc3339nano)
// 4. organizer email
// 5. start date (rfc3339nano)
// 6. subject
//nolint:lll
const (
	eventTmpl = `{
	"id":"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAENAADSEBNbUIB9RL6ePDeF3FIYAAAAAG76AAA=",
	"calendar@odata.navigationLink":"https://graph.microsoft.com/v1.0/users('foobar@8qzvrj.onmicrosoft.com')/calendars('AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAENAAA=')",
	"calendar@odata.associationLink":"https://graph.microsoft.com/v1.0/users('foobar@8qzvrj.onmicrosoft.com')/calendars('AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAENAAA=')/$ref",
	"@odata.etag":"W/\"0hATW1CAfUS+njw3hdxSGAAAJIxNug==\"",
	"@odata.context":"https://graph.microsoft.com/v1.0/$metadata#users('foobar%%408qzvrj.onmicrosoft.com')/events/$entity",
	"categories":[],
	"changeKey":"0hATW1CAfUS+njw3hdxSGAAAJIxNug==",
	"createdDateTime":"2022-03-28T03:42:03Z",
	"lastModifiedDateTime":"2022-05-26T19:25:58Z",
	"allowNewTimeProposals":true,
	"attendees":[],
	"body":{
		"content":"<html>\\r\\n<head>\\r\\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\">\\r\\n</head>\\r\\n<body>\\r\\n` +
		`<p>%s</p>\\r\\n</body>\\r\\n</html>\\r\\n",
		"contentType":"html"
	},
	"bodyPreview":"%s",
	"end":{
		"dateTime":"%s",
		"timeZone":"UTC"
	},
	"hasAttachments":false,
	"hideAttendees":false,
	"iCalUId":"040000008200E00074C5B7101A82E0080000000035723BC75542D801000000000000000010000000E1E7C8F785242E4894DA13AEFB947B85",
	"importance":"normal",
	"isAllDay":false,
	"isCancelled":false,
	"isDraft":false,
	"isOnlineMeeting":false,
	"isOrganizer":false,
	"isReminderOn":true,
	"location":{
		"displayName":"Umi Sake House (2230 1st Ave, Seattle, WA 98121 US)",
		"locationType":"default",
		"uniqueId":"Umi Sake House (2230 1st Ave, Seattle, WA 98121 US)",
		"uniqueIdType":"private"
	},
	"locations":[
		{
			"displayName":"Umi Sake House (2230 1st Ave, Seattle, WA 98121 US)",
			"locationType":"default",
			"uniqueId":"",
			"uniqueIdType":"unknown"
		}
	],
	"onlineMeetingProvider":"unknown",
	"organizer":{
		"emailAddress":{
			"address":"%s",
			"name":"Anu Pierson"
		}
	},
	"originalEndTimeZone":"UTC",
	"originalStartTimeZone":"UTC",
	"reminderMinutesBeforeStart":15,
	"responseRequested":true,
	"responseStatus":{
		"response":"notResponded",
		"time":"0001-01-01T00:00:00Z"
	},
	"sensitivity":"normal",
	"showAs":"tentative",
	"start":{
		"dateTime":"%s",
		"timeZone":"UTC"
	},
	"subject":"%s",
	"type":"singleInstance",
	"webLink":"https://outlook.office365.com/owa/?itemid=AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAENAADSEBNbUIB9RL6ePDeF3FIYAAAAAG76AAA%%3D&exvsurl=1&path=/calendar/item"
}`

	defaultEventBody        = "This meeting is to review the latest Tailspin Toys project proposal.<br>\\r\\nBut why not eat some sushi while we’re at it? :)"
	defaultEventBodyPreview = "This meeting is to review the latest Tailspin Toys project proposal.\\r\\nBut why not eat some sushi while we’re at it? :)"
	defaultEventOrganizer   = "foobar@8qzvrj.onmicrosoft.com"
)

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
func GetDefaultMockEventBytes(subject string) []byte {
	return GetMockEventWithSubjectBytes(" " + subject + " Review + Lunch")
}

func GetMockEventWithSubjectBytes(subject string) []byte {
	tomorrow := time.Now().UTC().AddDate(0, 0, 1)
	at := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), tomorrow.Hour(), 0, 0, 0, time.UTC)
	atTime := common.FormatTime(at)

	return GetMockEventWith(
		defaultEventOrganizer, subject,
		defaultEventBody, defaultEventBodyPreview,
		atTime, atTime,
	)
}

// GetMockEventWith returns bytes for an Eventable item.
// The event has no attendees.
// start and end times should be in the format 2006-01-02T15:04:05.0000000Z.
// The timezone (Z) will be automatically stripped.  A non-utc timezone may
// produce unexpected results.
// Body must contain a well-formatted string, consumable in a json payload.  IE: no unescaped newlines.
func GetMockEventWith(
	organizer, subject, body, bodyPreview,
	startDateTime, endDateTime string,
) []byte {
	startDateTime = strings.TrimSuffix(startDateTime, "Z")
	endDateTime = strings.TrimSuffix(endDateTime, "Z")

	if len(startDateTime) == len("2022-10-19T20:00:00") {
		startDateTime += ".0000000"
	}

	if len(endDateTime) == len("2022-10-19T20:00:00") {
		endDateTime += ".0000000"
	}

	return []byte(fmt.Sprintf(
		eventTmpl,
		body,
		bodyPreview,
		endDateTime,
		organizer,
		startDateTime,
		subject,
	))
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
