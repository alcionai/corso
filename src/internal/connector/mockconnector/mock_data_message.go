package mockconnector

import (
	"encoding/base64"
	"fmt"

	"github.com/alcionai/corso/src/internal/common"
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

	defaultMessageCreatedTime  = "2022-09-26T23:15:50Z"
	defaultMessageModifiedTime = "2022-09-26T23:15:51Z"
	defaultMessageReceivedTime = "2022-09-26T23:15:50Z"
	defaultMessageSentTime     = "2022-09-26T23:15:46Z"
	defaultMessageSender       = "foobar@8qzvrj.onmicrosoft.com"
	defaultMessageTo           = "LidiaH@8qzvrj.onmicrosoft.com"
	defaultMessageFrom         = "foobar@8qzvrj.onmicrosoft.com"

	// Order of fields to fill in:
	//   1. created datetime
	//   2. modified datetime
	//   3. message body
	//   4. message preview
	//   5. sender user ID
	//   6. received datetime
	//   7. sender email
	//   8. sent datetime
	//   9. subject
	//   10. recipient user addr
	messageTmpl = `{
		"id":"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB3XwIkAAA=",
		"@odata.context":"https://graph.microsoft.com/v1.0/$metadata#users('a4a472f8-ccb0-43ec-bf52-3697a91b926c')/messages/$entity",
		"@odata.etag":"W/\"CQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAAB2ZxqU\"",
		"categories":[],
		"changeKey":"CQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAAB2ZxqU",
		"createdDateTime":"%s",
		"lastModifiedDateTime":"%s",
		"bccRecipients":[],
		"body":{
			"content":"<html><head>\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><style type=\"text/css\" style=\"display:none\">\n<!--\np\n{margin-top:0;\nmargin-bottom:0}\n-->` +
		`\n</style></head><body dir=\"ltr\"><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\">` +
		`%s` +
		`</div></body></html>",
			"contentType":"html"
		},
		"bodyPreview":"%s",
		"ccRecipients":[],
		"conversationId":"AAQkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAQAK5nNWRdNWpGpLp7Xpb-m7A=",
		"conversationIndex":"AQHY0f3Ermc1ZF01akakuntelv+bsA==",
		"flag":{
			"flagStatus":"notFlagged"},
			"from":{
				"emailAddress":{
					"address":"%s",
					"name":"A Stranger"
				}
			},
		"hasAttachments":false,
		"importance":"normal",
		"inferenceClassification":"focused",
		"internetMessageId":"<SJ0PR17MB562266A1E61A8EA12F5FB17BC3529@SJ0PR17MB5622.namprd17.prod.outlook.com>",
		"isDeliveryReceiptRequested":false,
		"isDraft":false,
		"isRead":false,
		"isReadReceiptRequested":false,
		"parentFolderId":"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAAA=",
		"receivedDateTime":"%s",
		"replyTo":[],
		"sender":{
			"emailAddress":{
				"address":"%s",
				"name":"A Stranger"
			}
		},
		"sentDateTime":"%s",
		"subject":"%s",
		"toRecipients":[
			{
				"emailAddress":{
					"address":"%s",
					"name":"A Stranger"
				}
			}
		],
		"webLink":"https://outlook.office365.com/owa/?ItemID=AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB3XwIkAAA%%3D&exvsurl=1&viewmodel=ReadMessageItem"
	}`
)

// GetMockMessageBytes returns bytes for a Messageable item.
// Contents verified as working with sample data from kiota-serialization-json-go v0.5.5
func GetMockMessageBytes(subject string) []byte {
	return GetMockMessageWithBodyBytes(
		"TPS Report "+subject+" "+common.FormatNow(common.SimpleDateTime),
		defaultMessageBody, defaultMessagePreview)
}

// GetMockMessageBytes returns bytes for a Messageable item.
// Contents verified as working with sample data from kiota-serialization-json-go v0.5.5
// Body must contain a well-formatted string, consumable in a json payload.  IE: no unescaped newlines.
func GetMockMessageWithBodyBytes(subject, body, preview string) []byte {
	return GetMockMessageWith(
		defaultMessageTo,
		defaultMessageFrom,
		defaultMessageSender,
		subject,
		body,
		preview,
		defaultMessageCreatedTime,
		defaultMessageModifiedTime,
		defaultMessageSentTime,
		defaultMessageReceivedTime,
	)
}

// GetMockMessageWith returns bytes for a Messageable item.
// Contents verified as working with sample data from kiota-serialization-json-go v0.5.5
// created, modified, sent, and received should be in the format 2006-01-02T15:04:05Z
// Body must contain a well-formatted string, consumable in a json payload.  IE: no unescaped newlines.
func GetMockMessageWith(
	to, from, sender, // user PNs
	subject, body, preview, // arbitrary data
	created, modified, sent, received string, // legacy datetimes
) []byte {
	if len(preview) == 0 {
		preview = body
	}

	if len(preview) > 255 {
		preview = preview[:256]
	}

	message := fmt.Sprintf(
		messageTmpl,
		created,
		modified,
		body,
		preview,
		from,
		received,
		sender,
		sent,
		subject,
		to)

	return []byte(message)
}

// GetMockMessageWithDirectAttachment returns a message with inline attachment
// Serialized with: kiota-serialization-json-go v0.7.1
func GetMockMessageWithDirectAttachment(subject string) []byte {
	//nolint:lll
	message := "{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB4moqeAAA=\"," +
		"\"@odata.type\":\"#microsoft.graph.message\",\"@odata.etag\":\"W/\\\"CQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAAB3maFQ\\\"\",\"@odata.context\":\"https://graph.microsoft.com/v1.0/$metadata#users('a4a472f8-ccb0-43ec-bf52-3697a91b926c')/messages/$entity\",\"categories\":[]," +
		"\"changeKey\":\"CQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAAB3maFQ\",\"createdDateTime\":\"2022-09-29T17:39:06Z\",\"lastModifiedDateTime\":\"2022-09-29T17:39:08Z\"," +
		"\"attachments\":[{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB4moqeAAABEgAQANMmZLFhjWJJj4X9mj8piqg=\",\"@odata.type\":\"#microsoft.graph.fileAttachment\",\"@odata.mediaContentType\":\"application/octet-stream\"," +
		"\"contentType\":\"application/octet-stream\",\"isInline\":false,\"lastModifiedDateTime\":\"2022-09-29T17:39:06Z\",\"name\":\"database.db\",\"size\":11418," +
		"\"contentBytes\":\"U1FMaXRlIGZvcm1hdCAzAAQAAQEAQCAgAAAATQAAAAsAAAAEAAAACAAAAAsAAAAEAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABNAC3mBw0DZwACAg8AAxUCDwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACCAwMHFxUVAYNpdGFibGVkYXRhZGF0YQJDUkVBVEUgVEFCTEUgZGF0YSAoCiAgICAgICAgIGlkIGludGVnZXIgcHJpbWFyeSBrZXkgYXV0b2luY3JlbWVudCwKICAgICAgICAgbWVhbiB0ZXh0IG5vdCBudWxsLAogICAgICAgICBtYXggdGV4dCBub3QgbnVsbCwKICAgICAgICAgbWluIHRleHQgbm90IG51bGwsCiAgICAgIC" +
		"AgIGRhdGEgdGV4dCBub3QgbnVsbCwKICAgICAgICAgZ2VuZGVyIHRleHQgbm90IG51bGwsCgkgdmFsaWQgaW50ZWdlciBkZWZhdWx0IDEpUAIGFysrAVl0YWJsZXNxbGl0ZV9zZXF1ZW5jZXNxbGl0ZV9zZXF1ZW5jZQNDUkVBVEUgVEFCTEUgc3FsaXRlX3NlcXVlbmNlKG5hbWUsc2VxKQAAAJkAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA0AAAAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAANAAAAAAQAAAP2AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAcAAAAIAAAABgAAAAcAAAAFAAAACwAAAAoAAAAJAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
		"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=\"}]," +
		"\"bccRecipients\":[],\"body\":{\"content\":\"<html><head>\\r\\n<meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\"><style type=\\\"text/css\\\" style=\\\"display:none\\\">\\r\\n<!--\\r\\np\\r\\n\\t{margin-top:0;\\r\\n\\tmargin-bottom:0}\\r\\n-->\\r\\n</style></head><body dir=\\\"ltr\\\"><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\"><span class=\\\"x_elementToProof ContentPasted0\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\">Lidia,</span> <div class=\\\"x_elementToProof\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\"><br class=\\\"ContentPasted0\\\"></div><div class=\\\"x_elementToProof ContentPasted0\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\">I hope this message finds you well. I am researching a database construct for next quarter's review. SkyNet will<span class=\\\"ContentPasted0\\\">&nbsp;</span><span data-ogsb=\\\"rgb(255, 255, 0)\\\" class=\\\"ContentPasted0\\\" style=\\\"margin:0px; background-color:rgb(255,255,0)!important\\\">not</span><span class=\\\"ContentPasted0\\\">&nbsp;</span>be able to match our database process speeds if we utilize the formulae that are included.&nbsp;</div><div class=\\\"x_elementToProof\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\"><br class=\\\"ContentPasted0\\\"></div><div class=\\\"x_elementToProof ContentPasted0\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\">Please give me your thoughts on the implementation.</div><div class=\\\"x_elementToProof\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\"><br class=\\\"ContentPasted0\\\"></div><div class=\\\"x_elementToProof ContentPasted0\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\">Best,</div><div class=\\\"x_elementToProof\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\"><br class=\\\"ContentPasted0\\\"></div><span class=\\\"x_elementToProof ContentPasted0\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\">Dustin</span><br></div></body></html>\",\"contentType\":\"html\",\"@odata.type\":\"#microsoft.graph.itemBody\"}," +
		"\"bodyPreview\":\"Lidia,\\r\\n\\r\\nI hope this message finds you well. I am researching a database construct for next quarter's review. SkyNet will not be able to match our database process speeds if we utilize the formulae that are included.\\r\\n\\r\\nPlease give me your thoughts on th\",\"ccRecipients\":[]," +
		"\"conversationId\":\"AAQkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAQANPFOcy_BapBghezTzIIldI=\",\"conversationIndex\":\"AQHY1Cpb08U5zL4FqkGCF7NPMgiV0g==\",\"flag\":{\"flagStatus\":\"notFlagged\",\"@odata.type\":\"#microsoft.graph.followupFlag\"}," +
		"\"from\":{\"emailAddress\":{\"address\":\"dustina@8qzvrj.onmicrosoft.com\",\"name\":\"Dustin Abbot\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"hasAttachments\":true,\"importance\":\"normal\",\"inferenceClassification\":\"focused\"," +
		"\"internetMessageId\":\"<SJ0PR17MB56220C509D0006B8CC8FD952C3579@SJ0PR17MB5622.namprd17.prod.outlook.com>\",\"isDeliveryReceiptRequested\":false,\"isDraft\":false,\"isRead\":false,\"isReadReceiptRequested\":false," +
		"\"parentFolderId\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAAA=\",\"receivedDateTime\":\"2022-09-29T17:39:07Z\",\"replyTo\":[],\"sender\":{\"emailAddress\":{\"address\":\"dustina@8qzvrj.onmicrosoft.com\",\"name\":\"Dustin Abbot\"," +
		"\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"sentDateTime\":\"2022-09-29T17:39:02Z\"," +
		"\"subject\":\"" + subject + "\",\"toRecipients\":[{\"emailAddress\":{\"address\":\"LidiaH@8qzvrj.onmicrosoft.com\",\"name\":\"Lidia Holloway\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"}]," +
		"\"webLink\":\"https://outlook.office365.com/owa/?ItemID=AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB4moqeAAA%3D&exvsurl=1&viewmodel=ReadMessageItem\"}"

	return []byte(message)
}

// GetMockMessageWithDirectAttachment returns a message with a large attachment. This is derived from the message
// used in GetMockMessageWithDirectAttachment
// Serialized with: kiota-serialization-json-go v0.7.1
func GetMockMessageWithLargeAttachment(subject string) []byte {
	//nolint:lll
	messageFmt := "{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB4moqeAAA=\"," +
		"\"@odata.type\":\"#microsoft.graph.message\",\"@odata.etag\":\"W/\\\"CQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAAB3maFQ\\\"\",\"@odata.context\":\"https://graph.microsoft.com/v1.0/$metadata#users('a4a472f8-ccb0-43ec-bf52-3697a91b926c')/messages/$entity\",\"categories\":[]," +
		"\"changeKey\":\"CQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAAB3maFQ\",\"createdDateTime\":\"2022-09-29T17:39:06Z\",\"lastModifiedDateTime\":\"2022-09-29T17:39:08Z\"," +
		"\"attachments\":[{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB4moqeAAABEgAQANMmZLFhjWJJj4X9mj8piqg=\",\"@odata.type\":\"#microsoft.graph.fileAttachment\",\"@odata.mediaContentType\":\"application/octet-stream\"," +
		"\"contentType\":\"application/octet-stream\",\"isInline\":false,\"lastModifiedDateTime\":\"2022-09-29T17:39:06Z\",\"name\":\"database.db\",\"size\":%d," +
		"\"contentBytes\":\"%s\"}]," +
		"\"bccRecipients\":[],\"body\":{\"content\":\"<html><head>\\r\\n<meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\"><style type=\\\"text/css\\\" style=\\\"display:none\\\">\\r\\n<!--\\r\\np\\r\\n\\t{margin-top:0;\\r\\n\\tmargin-bottom:0}\\r\\n-->\\r\\n</style></head><body dir=\\\"ltr\\\"><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\"><span class=\\\"x_elementToProof ContentPasted0\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\">Lidia,</span> <div class=\\\"x_elementToProof\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\"><br class=\\\"ContentPasted0\\\"></div><div class=\\\"x_elementToProof ContentPasted0\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\">I hope this message finds you well. I am researching a database construct for next quarter's review. SkyNet will<span class=\\\"ContentPasted0\\\">&nbsp;</span><span data-ogsb=\\\"rgb(255, 255, 0)\\\" class=\\\"ContentPasted0\\\" style=\\\"margin:0px; background-color:rgb(255,255,0)!important\\\">not</span><span class=\\\"ContentPasted0\\\">&nbsp;</span>be able to match our database process speeds if we utilize the formulae that are included.&nbsp;</div><div class=\\\"x_elementToProof\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\"><br class=\\\"ContentPasted0\\\"></div><div class=\\\"x_elementToProof ContentPasted0\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\">Please give me your thoughts on the implementation.</div><div class=\\\"x_elementToProof\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\"><br class=\\\"ContentPasted0\\\"></div><div class=\\\"x_elementToProof ContentPasted0\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\">Best,</div><div class=\\\"x_elementToProof\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\"><br class=\\\"ContentPasted0\\\"></div><span class=\\\"x_elementToProof ContentPasted0\\\" data-ogsc=\\\"rgb(0, 0, 0)\\\" data-ogsb=\\\"rgb(255, 255, 255)\\\" style=\\\"font-size:12pt; margin:0px; color:rgb(0,0,0)!important; background-color:rgb(255,255,255)!important\\\">Dustin</span><br></div></body></html>\",\"contentType\":\"html\",\"@odata.type\":\"#microsoft.graph.itemBody\"}," +
		"\"bodyPreview\":\"Lidia,\\r\\n\\r\\nI hope this message finds you well. I am researching a database construct for next quarter's review. SkyNet will not be able to match our database process speeds if we utilize the formulae that are included.\\r\\n\\r\\nPlease give me your thoughts on th\",\"ccRecipients\":[]," +
		"\"conversationId\":\"AAQkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAQANPFOcy_BapBghezTzIIldI=\",\"conversationIndex\":\"AQHY1Cpb08U5zL4FqkGCF7NPMgiV0g==\",\"flag\":{\"flagStatus\":\"notFlagged\",\"@odata.type\":\"#microsoft.graph.followupFlag\"}," +
		"\"from\":{\"emailAddress\":{\"address\":\"dustina@8qzvrj.onmicrosoft.com\",\"name\":\"Dustin Abbot\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"hasAttachments\":true,\"importance\":\"normal\",\"inferenceClassification\":\"focused\"," +
		"\"internetMessageId\":\"<SJ0PR17MB56220C509D0006B8CC8FD952C3579@SJ0PR17MB5622.namprd17.prod.outlook.com>\",\"isDeliveryReceiptRequested\":false,\"isDraft\":false,\"isRead\":false,\"isReadReceiptRequested\":false," +
		"\"parentFolderId\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAAA=\",\"receivedDateTime\":\"2022-09-29T17:39:07Z\",\"replyTo\":[],\"sender\":{\"emailAddress\":{\"address\":\"dustina@8qzvrj.onmicrosoft.com\",\"name\":\"Dustin Abbot\"," +
		"\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"sentDateTime\":\"2022-09-29T17:39:02Z\"," +
		"\"subject\":\"" + subject + "\",\"toRecipients\":[{\"emailAddress\":{\"address\":\"LidiaH@8qzvrj.onmicrosoft.com\",\"name\":\"Lidia Holloway\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"}]," +
		"\"webLink\":\"https://outlook.office365.com/owa/?ItemID=AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB4moqeAAA%3D&exvsurl=1&viewmodel=ReadMessageItem\"}"

	attachmentSize := 3 * 1024 * 1024 // 3 MB
	attachmentBytes := make([]byte, attachmentSize)

	// Attachment content bytes are base64 encoded
	return []byte(fmt.Sprintf(messageFmt, attachmentSize, base64.StdEncoding.EncodeToString([]byte(attachmentBytes))))
}

// GetMessageWithOneDriveAttachment returns a message with an OneDrive attachment represented in bytes
// Serialized with: kiota-serialization-json-go v0.7.1
func GetMessageWithOneDriveAttachment(subject string) []byte {
	//nolint:lll
	message := "{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB4moqfAAA=\"," +
		"\"@odata.type\":\"#microsoft.graph.message\",\"@odata.etag\":\"W/\\\"CQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAAB3maFw\\\"\",\"@odata.context\":\"https://graph.microsoft.com/v1.0/$metadata#users('a4a472f8-ccb0-43ec-bf52-3697a91b926c')/messages/$entity\",\"categories\":[]," +
		"\"changeKey\":\"CQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAAB3maFw\",\"createdDateTime\":\"2022-09-29T17:43:25Z\",\"lastModifiedDateTime\":\"2022-09-29T17:43:42Z\"," +
		"\"attachments\":[{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB4moqfAAABEgAQAGYItMZBbGFNktPW57eZMCA=\",\"@odata.type\":\"#microsoft.graph.referenceAttachment\",\"contentType\":\"image/jpeg\",\"isInline\":true,\"lastModifiedDateTime\":\"2022-09-29T17:43:25Z\",\"name\":\"today.jpeg\",\"size\":1098}]," +
		"\"bccRecipients\":[],\"body\":{\"content\":\"<html><head>\\r\\n<meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\"><style type=\\\"text/css\\\" style=\\\"display:none\\\">\\r\\n<!--\\r\\np\\r\\n\\t{margin-top:0;\\r\\n\\tmargin-bottom:0}\\r\\n-->\\r\\n</style></head><body dir=\\\"ltr\\\"><div id=\\\"OwaReferenceAttachments\\\"><table style=\\\"padding-bottom:13px; border-width:0px; border-style:none\\\"><tbody><tr valign=\\\"top\\\"><td><table style=\\\"border-width:0px 0px 1px 0px; border-color:#C7C7C7; border-style:none none dotted none\\\"><tbody><tr valign=\\\"top\\\"><td style=\\\"padding-bottom:7px\\\"><table align=\\\"left\\\" style=\\\"padding-right:28px; border-width:0px; background-color:rgb(255,255,255); border-spacing:0px\\\"><tbody><tr valign=\\\"top\\\"><td style=\\\"padding:0px\\\"><div id=\\\"OwaReferenceAttachmentDescription\\\" style=\\\"padding-left:3px; font-size:14px; font-family:'Segoe UI','Segoe WP','Segoe UI WPC',Tahoma,Arial,sans-serif; color:rgb(102,102,102)\\\">Dustin Abbot has shared a OneDrive for Business file with you. To view it, click the link below. </div></td></tr></tbody></table></td></tr><tr valign=\\\"top\\\"><td><a href=\\\"https://8qzvrj-my.sharepoint.com/:i:/g/personal/dustina_8qzvrj_onmicrosoft_com/Ec0lO09-i3JNqugzLTbzvtkBNIc5aY_ltIC-Be9gIQOnZA\\\" target=\\\"_blank\\\"><table align=\\\"left\\\" style=\\\"padding-right:28px; padding-bottom:10px; border-width:0px; height:20px; background-color:rgb(255,255,255); border-spacing:0px\\\"><tbody><tr valign=\\\"top\\\"><td style=\\\"padding:0px\\\"><div style=\\\"background-color:rgb(255,255,255); height:20px; width:20px; max-height:20px\\\"><a href=\\\"https://8qzvrj-my.sharepoint.com/:i:/g/personal/dustina_8qzvrj_onmicrosoft_com/Ec0lO09-i3JNqugzLTbzvtkBNIc5aY_ltIC-Be9gIQOnZA\\\" target=\\\"_blank\\\"><img width=\\\"20\\\" src=\\\"https://r1.res.office365.com/owa/prem/images/dc-jpg_20.png\\\" style=\\\"border:0px\\\"></a></div></td><td><div id=\\\"OwaReferenceAttachmentFileName2\\\" style=\\\"padding:0px 0px 0px 5px; font-size:14px; font-family:'Segoe UI','Segoe WP','Segoe UI WPC',Tahoma,Arial,sans-serif; color:rgb(0,114,198)\\\"><a href=\\\"https://8qzvrj-my.sharepoint.com/:i:/g/personal/dustina_8qzvrj_onmicrosoft_com/Ec0lO09-i3JNqugzLTbzvtkBNIc5aY_ltIC-Be9gIQOnZA\\\" target=\\\"_blank\\\" style=\\\"text-decoration:none; margin:0px; font-size:14px; font-family:'Segoe UI','Segoe WP','Segoe UI WPC',Tahoma,Arial,sans-serif; color:rgb(0,114,198)\\\">today.jpeg</a></div></td><td width=\\\"0\\\" height=\\\"0\\\" style=\\\"display:none; visibility:hidden\\\"><img width=\\\"0\\\" height=\\\"0\\\" title=\\\"today.jpeg\\\" data-outlook-trace=\\\"F:1|T:1\\\" src=\\\"cid:6f03c4a1-cc7a-4f53-af15-bde934e529fd\\\" style=\\\"visibility:hidden; border:0px; display:none\\\"></td></tr></tbody></table></a></td></tr></tbody></table></td></tr></tbody></table></div><div id=\\\"OwaReferenceAttachmentsEnd\\\" style=\\\"display:none; visibility:hidden\\\"></div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\">Lidia,</div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\"><br></div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\">I have included a shared resource that keeps populating my OneDrive. I'm not sure if you are getting this image of this resort or not, but I cannot seem to get the rest of my team to it within OneDrive. Do you have any thoughts on how I should proceed?</div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\"><br></div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\">The object is only 320kb so I don't understand why it is presenting such a problem.</div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\"><br></div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\">Best,&nbsp;</div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\"><br></div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\">Dustin</div></body></html>\",\"contentType\":\"html\",\"@odata.type\":\"#microsoft.graph.itemBody\"}," +
		"\"bodyPreview\":\"Lidia,\\r\\n\\r\\nI have included a shared resource that keeps populating my OneDrive. I'm not sure if you are getting this image of this resort or not, but I cannot seem to get the rest of my team to it within OneDrive. Do you have any thoughts on how I should p\"," +
		"\"ccRecipients\":[],\"conversationId\":\"AAQkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAQANqW7u0YwUJGmoVx5-9afMM=\",\"conversationIndex\":\"AQHY1CqR2pbu7RjBQkaahXHn/1p8ww==\",\"flag\":{\"flagStatus\":\"notFlagged\",\"@odata.type\":\"#microsoft.graph.followupFlag\"}," +
		"\"from\":{\"emailAddress\":{\"address\":\"dustina@8qzvrj.onmicrosoft.com\",\"name\":\"Dustin Abbot\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"hasAttachments\":true,\"importance\":\"normal\",\"inferenceClassification\":\"focused\"," +
		"\"internetMessageId\":\"<SJ0PR17MB5622F8AE5C4518C23DCDCA25C3579@SJ0PR17MB5622.namprd17.prod.outlook.com>\",\"isDeliveryReceiptRequested\":false,\"isDraft\":false,\"isRead\":false,\"isReadReceiptRequested\":false,\"parentFolderId\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAAA=\"," +
		"\"receivedDateTime\":\"2022-09-29T17:43:25Z\",\"replyTo\":[],\"sender\":{\"emailAddress\":{\"address\":\"dustina@8qzvrj.onmicrosoft.com\",\"name\":\"Dustin Abbot\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"}," +
		"\"sentDateTime\":\"2022-09-29T17:43:20Z\",\"subject\":\"" + subject + "\",\"toRecipients\":[{\"emailAddress\":{\"address\":\"LidiaH@8qzvrj.onmicrosoft.com\",\"name\":\"Lidia Holloway\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"}]," +
		"\"webLink\":\"https://outlook.office365.com/owa/?ItemID=AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB4moqfAAA%3D&exvsurl=1&viewmodel=ReadMessageItem\"}"

	return []byte(message)
}

// GetMockMessageWithTwoAttachments returns byte representation of message with two attachments
// Serialized with: kiota-serialization-json-go v0.7.1
func GetMockMessageWithTwoAttachments(subject string) []byte {
	//nolint:lll
	message := "{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB6LpD0AAA=\",\"@odata.type\":\"#microsoft.graph.message\",\"@odata.etag\":\"W/\\\"CQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAAB5JBpO\\\"\",\"@odata.context\":\"https://graph.microsoft.com/v1.0/$metadata#users('a4a472f8-ccb0-43ec-bf52-3697a91b926c')/messages/$entity\",\"categories\":[],\"changeKey\":\"CQAAABYAAADSEBNbUIB9RL6ePDeF3FIYAAB5JBpO\"," +
		"\"createdDateTime\":\"2022-09-30T20:31:22Z\",\"lastModifiedDateTime\":\"2022-09-30T20:31:25Z\",\"attachments\":[{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB6LpD0AAABEgAQAMIBac0_D4pPgtgr9mhVWaM=\",\"@odata.type\":\"#microsoft.graph.fileAttachment\",\"@odata.mediaContentType\":\"text/plain\",\"contentType\":\"text/plain\",\"isInline\":false,\"lastModifiedDateTime\":\"2022-09-30T20:31:22Z\"," +
		"\"name\":\"sample.txt\",\"size\":198,\"contentBytes\":\"VFBTIFJlcG9ydHMgYXJlIGZvciB3aW5uZXJzCg==\"},{\"id\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB6LpD0AAABEgAQAHO2tnfyTF1HnQKqNSMCO7A=\",\"@odata.type\":\"#microsoft.graph.fileAttachment\",\"@odata.mediaContentType\":\"text/plain\",\"contentType\":\"text/plain\",\"isInline\":false,\"lastModifiedDateTime\":\"2022-09-30T20:31:22Z\"," +
		"\"name\":\"sample3.txt\",\"size\":234,\"contentBytes\":\"SWYgdGhlIGZvcmNlIGlzIHdpdGggeW91LCBpdCdzIHdpdGggeW91LiBOb3Qgb25seSBpbiBNYXkuCg==\"}],\"bccRecipients\":[],\"body\":{\"content\":\"<html><head>\\r\\n<meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\"><style type=\\\"text/css\\\" style=\\\"display:none\\\">\\r\\n<!--\\r\\np\\r\\n\\t{margin-top:0;\\r\\n\\tmargin-bottom:0}\\r\\n-->\\r\\n</style></head><body dir=\\\"ltr\\\"><div class=\\\"elementToProof\\\" " +
		"style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\">Lidia,</div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\"><br></div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\">We have to decide between two items for our speech writers to go over. Please let me know which is the best for the " +
		"upcoming retreat.&nbsp;</div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\"><br></div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\">Best,&nbsp;</div><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\"><br></div><div class=\\\"elementToProof\\\" " +
		"style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0)\\\">Dustin</div></body></html>\",\"contentType\":\"html\",\"@odata.type\":\"#microsoft.graph.itemBody\"},\"bodyPreview\":\"Lidia,\\r\\n\\r\\nWe have to decide between two items for our speech writers to go over. Please let me know which is the best for the upcoming retreat.\\r\\n\\r\\nBest,\\r\\n\\r\\nDustin\",\"ccRecipients\":[]," +
		"\"conversationId\":\"AAQkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAQANHUb9Zc-aBAvBW5io77k-g=\",\"conversationIndex\":\"AQHY1Qss0dRv1lz9oEC8FbmKjvuT+A==\",\"flag\":{\"flagStatus\":\"notFlagged\",\"@odata.type\":\"#microsoft.graph.followupFlag\"},\"from\":{\"emailAddress\":{\"address\":\"dustina@8qzvrj.onmicrosoft.com\",\"name\":\"Dustin Abbot\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"hasAttachments\":true," +
		"\"importance\":\"normal\",\"inferenceClassification\":\"focused\",\"internetMessageId\":\"<SJ0PR17MB5622DB7B5847BF4BE7965B32C3569@SJ0PR17MB5622.namprd17.prod.outlook.com>\",\"isDeliveryReceiptRequested\":false,\"isDraft\":false,\"isRead\":false,\"isReadReceiptRequested\":false,\"parentFolderId\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAAA=\",\"receivedDateTime\":\"2022-09-30T20:31:23Z\",\"replyTo\":[]," +
		"\"sender\":{\"emailAddress\":{\"address\":\"dustina@8qzvrj.onmicrosoft.com\",\"name\":\"Dustin Abbot\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"sentDateTime\":\"2022-09-30T20:31:19Z\"," +
		"\"subject\":\"" + subject + "\",\"toRecipients\":[{\"emailAddress\":{\"address\":\"LidiaH@8qzvrj.onmicrosoft.com\",\"name\":\"Lidia Holloway\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"}],\"webLink\":\"https://outlook.office365.com/owa/?ItemID=AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB6LpD0AAA%3D&exvsurl=1&viewmodel=ReadMessageItem\"}"

	return []byte(message)
}
