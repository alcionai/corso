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
	defaultAlias               = "A Stranger"

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

// GetMockMessageWithDirectAttachment returns a message an attachment that contains n MB of data.
// Max limit on N is 35 (imposed by exchange) .
// Serialized with: kiota-serialization-json-go v0.7.1
func GetMockMessageWithSizedAttachment(subject string, n int) []byte {
	// I know we said 35, but after base64encoding, 24mb of base content
	// bloats up to 34mb (35 baloons to 49).  So we have to restrict n
	// appropriately.
	if n > 24 {
		n = 24
	}

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

	attachmentSize := n * 1024 * 1024 // n MB
	attachmentBytes := make([]byte, attachmentSize)

	// Attachment content bytes are base64 encoded
	return []byte(fmt.Sprintf(messageFmt, attachmentSize, base64.StdEncoding.EncodeToString([]byte(attachmentBytes))))
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
	return GetMockMessageWithSizedAttachment(subject, 3)
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
		"\"conversationId\":\"AAQkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAQANHUb9Zc-aBAvBW5io77k-g=\",\"conversationIndex\":\"AQHY1Qss0dRv1lz9oEC8FbmKjvuT+A==\",\"flag\":{\"flagStatus\":\"notFlagged\",\"@odata.type\":\"#microsoft.graph.followupFlag\"}," +
		"\"from\":{\"emailAddress\":{\"address\":\"" + defaultMessageFrom + "\",\"name\":\"" + defaultAlias + "\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"hasAttachments\":true," +
		"\"importance\":\"normal\",\"inferenceClassification\":\"focused\",\"internetMessageId\":\"<SJ0PR17MB5622DB7B5847BF4BE7965B32C3569@SJ0PR17MB5622.namprd17.prod.outlook.com>\",\"isDeliveryReceiptRequested\":false,\"isDraft\":false,\"isRead\":false,\"isReadReceiptRequested\":false,\"parentFolderId\":\"AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwAuAAAAAADCNgjhM9QmQYWNcI7hCpPrAQDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAAA=\",\"receivedDateTime\":\"2022-09-30T20:31:23Z\",\"replyTo\":[]," +
		"\"sender\":{\"emailAddress\":{\"address\":\"" + defaultMessageSender + "\",\"name\":\"" + defaultAlias + "\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"sentDateTime\":\"2022-09-30T20:31:19Z\"," +
		"\"subject\":\"" + subject + "\",\"toRecipients\":[{\"emailAddress\":{\"address\":\"" + defaultMessageTo + "\",\"name\":\"" + defaultAlias + "\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"}],\"webLink\":\"https://outlook.office365.com/owa/?ItemID=AAMkAGZmNjNlYjI3LWJlZWYtNGI4Mi04YjMyLTIxYThkNGQ4NmY1MwBGAAAAAADCNgjhM9QmQYWNcI7hCpPrBwDSEBNbUIB9RL6ePDeF3FIYAAAAAAEMAADSEBNbUIB9RL6ePDeF3FIYAAB6LpD0AAA%3D&exvsurl=1&viewmodel=ReadMessageItem\"}"

	return []byte(message)
}

// GetMockEventMessageResponse returns byte representation of EventMessageResponse
// Special Mock to ensure that EventMessageResponse emails are transformed properly
func GetMockEventMessageResponse(subject string) []byte {
	//nolint:lll
	message := "{\"id\":\"AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEMAAB8wYc0thTTTYl3RpEYIUq_AACL4y38AAA=\"," +
		"\"@odata.type\":\"#microsoft.graph.eventMessageResponse\",\"@odata.context\":\"https://graph.microsoft.com/v1.0/$metadata#users('dustina%408qzvrj.onmicrosoft.com')/messages/$entity\"," +
		"\"@odata.etag\":\"W/\\\"DAAAABYAAAB8wYc0thTTTYl3RpEYIUq+AACLqMzx\\\"\",\"categories\":[],\"changeKey\":\"DAAAABYAAAB8wYc0thTTTYl3RpEYIUq+AACLqMzx\",\"createdDateTime\":\"2022-11-04T19:52:34Z\"," +
		"\"lastModifiedDateTime\":\"2022-11-04T19:52:37Z\",\"bccRecipients\":[],\"body\":{\"content\":\"<html><head>\\r\\n<meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\"></head><body><div style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\\\"><br></div></body></html>\",\"contentType\":\"html\",\"@odata.type\":\"#microsoft.graph.itemBody\"}," +
		"\"bodyPreview\":\"\",\"ccRecipients\":[],\"conversationId\":\"AAQkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOAAQAMtB1_9d_x1MuCEXzNWAYMk=\",\"conversationIndex\":\"Adjwhp8cy0HX7137HUy4IRfM1YBgyQAAFyz/\",\"flag\":{\"flagStatus\":\"notFlagged\",\"@odata.type\":\"#microsoft.graph.followupFlag\"}," +
		"\"from\":{\"emailAddress\":{\"address\":\"" + defaultMessageFrom + "\",\"name\":\"" + defaultAlias + "\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"hasAttachments\":false,\"importance\":\"normal\",\"inferenceClassification\":\"focused\",\"internetMessageId\":\"<BLAPR17MB41793DBF298747F4A3042667B23B9@BLAPR17MB4179.namprd17.prod.outlook.com>\",\"isDraft\":false,\"isRead\":false,\"isReadReceiptRequested\":false,\"parentFolderId\":\"AQMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4ADVkZWQwNmNlMTgALgAAAw_9XBStqZdPuOVIalVTz7sBAHzBhzS2FNNNiXdGkRghSr4AAAIBDAAAAA==\",\"receivedDateTime\":\"2022-11-04T19:52:35Z\"," +
		"\"replyTo\":[],\"sender\":{\"emailAddress\":{\"address\":\"" + defaultMessageSender + "\",\"name\":\"" + defaultAlias + "\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"sentDateTime\":\"2022-11-04T19:52:29Z\"," +
		"\"subject\":\"" + subject + "\",\"toRecipients\":[{\"emailAddress\":{\"address\":\"" + defaultMessageTo + "\",\"name\":\"" + defaultAlias + "\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"}],\"webLink\":\"https://outlook.office365.com/owa/?ItemID=AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8%2B7BwB8wYc0thTTTYl3RpEYIUq%2BAAAAAAEMAAB8wYc0thTTTYl3RpEYIUq%2BAACL4y38AAA%3D&exvsurl=1&viewmodel=ReadMessageItem\",\"endDateTime\":{\"dateTime\":\"2022-11-26T16:30:00.0000000\",\"@odata.type\":\"#microsoft.graph.dateTimeTimeZone\",\"timeZone\":\"UTC\"},\"isAllDay\":false,\"isDelegated\":false,\"isOutOfDate\":false,\"meetingMessageType\":\"meetingTenativelyAccepted\",\"startDateTime\":{\"dateTime\":\"2022-11-26T16:00:00.0000000\",\"@odata.type\":\"#microsoft.graph.dateTimeTimeZone\",\"timeZone\":\"UTC\"},\"type\":\"singleInstance\",\"responseType\":\"tentativelyAccepted\"}"

	return []byte(message)
}

// GetMockEventMessageRequest returns byte representation of EventMessageRequest
// Special Mock to ensure that EventMessageRequests are transformed properly
func GetMockEventMessageRequest(subject string) []byte {
	//nolint:lll
	message := "{\"id\":\"AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEJAAB8wYc0thTTTYl3RpEYIUq_AACL5VwSAAA=\"," +
		"\"@odata.type\":\"#microsoft.graph.eventMessageRequest\",\"@odata.context\":\"https://graph.microsoft.com/v1.0/$metadata#users('dustina%408qzvrj.onmicrosoft.com')/messages/$entity\"," +
		"\"@odata.etag\":\"W/\\\"CwAAABYAAAB8wYc0thTTTYl3RpEYIUq+AACLqMzb\\\"\",\"categories\":[],\"changeKey\":\"CwAAABYAAAB8wYc0thTTTYl3RpEYIUq+AACLqMzb\",\"createdDateTime\":\"2022-11-04T19:49:59Z\"," +
		"\"lastModifiedDateTime\":\"2022-11-04T19:50:03Z\",\"bccRecipients\":[],\"body\":{\"content\":\"<html><head>\\r\\n<meta http-equiv=\\\"Content-Type\\\" content=\\\"text/html; charset=utf-8\\\"></head><body><div class=\\\"elementToProof\\\" style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\\\">Should get a sent mail of this&nbsp;</div><br><div style=\\\"width:100%; height:20px\\\"><span style=\\\"white-space:nowrap; color:#5F5F5F; opacity:.36\\\">________________________________________________________________________________</span> </div><div class=\\\"me-email-text\\\" lang=\\\"en-GB\\\" style=\\\"color:#252424; font-family:'Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif\\\"><div style=\\\"margin-top:24px; margin-bottom:20px\\\"><span style=\\\"font-size:24px; color:#252424\\\">Microsoft Teams meeting</span> </div><div style=\\\"margin-bottom:20px\\\"><div style=\\\"margin-top:0px; margin-bottom:0px; font-weight:bold\\\"><span style=\\\"font-size:14px; color:#252424\\\">Join on your computer or mobile app</span> </div><a class=\\\"me-email-headline\\\" href=\\\"https://teams.microsoft.com/l/meetup-join/19%3ameeting_NDVlMmMwMDEtMjdkOC00NGEyLWFkMDUtMDcxY2RmMzUzZWJm%40thread.v2/0?context=%7b%22Tid%22%3a%224d603060-18d6-4764-b9be-4cb794d32b69%22%2c%22Oid%22%3a%22f435c656-f8b2-4d71-93c3-6e092f52a167%22%7d\\\" target=\\\"_blank\\\" rel=\\\"noreferrer noopener\\\" style=\\\"font-size:14px; font-family:'Segoe UI Semibold','Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif; text-decoration:underline; color:#6264a7\\\">Click here to join the meeting</a> </div><div style=\\\"margin-bottom:20px; margin-top:20px\\\"><div style=\\\"margin-bottom:4px\\\"><span data-tid=\\\"meeting-code\\\" style=\\\"font-size:14px; color:#252424\\\">Meeting ID: <span style=\\\"font-size:16px; color:#252424\\\">220 529 763 834</span> </span><br><span style=\\\"font-size:14px; color:#252424\\\">Passcode: </span><span style=\\\"font-size:16px; color:#252424\\\">bayGtj </span><div style=\\\"font-size:14px\\\"><a class=\\\"me-email-link\\\" target=\\\"_blank\\\" href=\\\"https://www.microsoft.com/en-us/microsoft-teams/download-app\\\" rel=\\\"noreferrer noopener\\\" style=\\\"font-size:14px; text-decoration:underline; color:#6264a7; font-family:'Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif\\\">Download Teams</a> | <a class=\\\"me-email-link\\\" target=\\\"_blank\\\" href=\\\"https://www.microsoft.com/microsoft-teams/join-a-meeting\\\" rel=\\\"noreferrer noopener\\\" style=\\\"font-size:14px; text-decoration:underline; color:#6264a7; font-family:'Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif\\\">Join on the web</a></div></div></div><div style=\\\"margin-bottom:24px; margin-top:20px\\\"><a class=\\\"me-email-link\\\" target=\\\"_blank\\\" href=\\\"https://aka.ms/JoinTeamsMeeting\\\" rel=\\\"noreferrer noopener\\\" style=\\\"font-size:14px; text-decoration:underline; color:#6264a7; font-family:'Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif\\\">Learn more</a> | <a class=\\\"me-email-link\\\" target=\\\"_blank\\\" href=\\\"https://teams.microsoft.com/meetingOptions/?organizerId=f435c656-f8b2-4d71-93c3-6e092f52a167&amp;tenantId=4d603060-18d6-4764-b9be-4cb794d32b69&amp;threadId=19_meeting_NDVlMmMwMDEtMjdkOC00NGEyLWFkMDUtMDcxY2RmMzUzZWJm@thread.v2&amp;messageId=0&amp;language=en-GB\\\" rel=\\\"noreferrer noopener\\\" style=\\\"font-size:14px; text-decoration:underline; color:#6264a7; font-family:'Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif\\\">Meeting options</a> </div></div><div style=\\\"font-size:14px; margin-bottom:4px; font-family:'Segoe UI','Helvetica Neue',Helvetica,Arial,sans-serif\\\"></div><div style=\\\"font-size:12px\\\"></div><div></div><div style=\\\"width:100%; height:20px\\\"><span style=\\\"white-space:nowrap; color:#5F5F5F; opacity:.36\\\">________________________________________________________________________________</span> </div></body></html>\",\"contentType\":\"html\",\"@odata.type\":\"#microsoft.graph.itemBody\"},\"bodyPreview\":\"Should get a sent mail of this\\r\\n\\r\\n________________________________________________________________________________\\r\\nMicrosoft Teams meeting\\r\\nJoin on your computer or mobile app\\r\\nClick here to join the meeting\\r\\nMeeting ID: 220 529 763 834\\r\\nPasscode: bayGtj\",\"ccRecipients\":[],\"conversationId\":\"AAQkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOAAQAMtB1_9d_x1MuCEXzNWAYMk=\",\"conversationIndex\":\"Adjwhp8cy0HX7137HUy4IRfM1YBgyQ==\",\"flag\":{\"flagStatus\":\"notFlagged\",\"@odata.type\":\"#microsoft.graph.followupFlag\"}," +
		"\"from\":{\"emailAddress\":{\"address\":\"" + defaultMessageFrom + "\",\"name\":\"Dustin Abbot\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"hasAttachments\":false,\"importance\":\"normal\",\"inferenceClassification\":\"focused\",\"internetMessageId\":\"<SJ0PR17MB56221D8549729E3AFC63EFA1C33B9@SJ0PR17MB5622.namprd17.prod.outlook.com>\",\"isDraft\":false,\"isRead\":true,\"isReadReceiptRequested\":false,\"parentFolderId\":\"AQMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4ADVkZWQwNmNlMTgALgAAAw_9XBStqZdPuOVIalVTz7sBAHzBhzS2FNNNiXdGkRghSr4AAAIBCQAAAA==\",\"receivedDateTime\":\"2022-11-04T19:50:01Z\"," +
		"\"replyTo\":[],\"sender\":{\"emailAddress\":{\"address\":\"" + defaultMessageSender + "\",\"name\":\"" + defaultAlias + "\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"},\"sentDateTime\":\"2022-11-04T19:50:00Z\"," +
		"\"subject\":\"" + subject + "\",\"toRecipients\":[{\"emailAddress\":{\"address\":\"" + defaultMessageTo + "\",\"name\":\"" + defaultAlias + "\",\"@odata.type\":\"#microsoft.graph.emailAddress\"},\"@odata.type\":\"#microsoft.graph.recipient\"}],\"webLink\":\"https://outlook.office365.com/owa/?ItemID=AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8%2B7BwB8wYc0thTTTYl3RpEYIUq%2BAAAAAAEJAAB8wYc0thTTTYl3RpEYIUq%2BAACL5VwSAAA%3D&exvsurl=1&viewmodel=ReadMessageItem\",\"endDateTime\":{\"dateTime\":\"2022-11-26T16:30:00.0000000\",\"@odata.type\":\"#microsoft.graph.dateTimeTimeZone\",\"timeZone\":\"UTC\"},\"isAllDay\":false,\"isDelegated\":false,\"isOutOfDate\":false,\"meetingMessageType\":\"meetingRequest\",\"startDateTime\":{\"dateTime\":\"2022-11-26T16:00:00.0000000\",\"@odata.type\":\"#microsoft.graph.dateTimeTimeZone\",\"timeZone\":\"UTC\"},\"type\":\"singleInstance\",\"meetingRequestType\":\"newMeetingRequest\",\"responseRequested\":true}"

	return []byte(message)
}

func GetMockMessageWithItemAttachmentEvent(subject string) []byte {
	//nolint:lll
	message := "{\"id\":\"AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEMAAB8wYc0thTTTYl3RpEYIUq_AADFfThMAAA=\",\"@odata.type\":\"#microsoft.graph.message\"," +
		"\"@odata.etag\":\"W/\\\"CQAAABYAAAB8wYc0thTTTYl3RpEYIUq+AADFK3BH\\\"\",\"@odata.context\":\"https://graph.microsoft.com/v1.0/$metadata#users('dustina%408qzvrj.onmicrosoft.com')/messages/$entity\",\"categories\":[]," +
		"\"changeKey\":\"CQAAABYAAAB8wYc0thTTTYl3RpEYIUq+AADFK3BH\",\"createdDateTime\":\"2023-02-01T13:48:43Z\",\"lastModifiedDateTime\":\"2023-02-01T18:27:03Z\"," +
		"\"attachments\":[{\"id\":\"AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEMAAB8wYc0thTTTYl3RpEYIUq_AADFfThMAAABEgAQAKHxTL6mNCZPo71dbwrfKYM=\"," +
		"\"@odata.type\":\"#microsoft.graph.itemAttachment\",\"isInline\":false,\"lastModifiedDateTime\":\"2023-02-01T13:52:56Z\",\"name\":\"Holidayevent\",\"size\":2059,\"item\":{\"id\":\"\",\"@odata.type\":\"#microsoft.graph.event\"," +
		"\"createdDateTime\":\"2023-02-01T13:52:56Z\",\"lastModifiedDateTime\":\"2023-02-01T13:52:56Z\",\"body\":{\"content\":\"<html><head>\\r\\n<metahttp-equiv=\\\"Content-Type\\\"content=\\\"text/html;charset=utf-8\\\"></head><body>Let'slookforfunding!</body></html>\"," +
		"\"contentType\":\"html\"},\"end\":{\"dateTime\":\"2016-12-02T19:00:00.0000000Z\",\"timeZone\":\"UTC\"}," +
		"\"hasAttachments\":false,\"isAllDay\":false,\"isCancelled\":false,\"isDraft\":true,\"isOnlineMeeting\":false,\"isOrganizer\":true,\"isReminderOn\":false,\"organizer\":{\"emailAddress\":{\"address\":\"" + defaultMessageFrom + "\",\"name\":\"" + defaultAlias + "\"}}," +
		"\"originalEndTimeZone\":\"tzone://Microsoft/Utc\",\"originalStartTimeZone\":\"tzone://Microsoft/Utc\",\"reminderMinutesBeforeStart\":0,\"responseRequested\":true,\"start\":{\"dateTime\":\"2016-12-02T18:00:00.0000000Z\",\"timeZone\":\"UTC\"}," +
		"\"subject\":\"Discussgiftsforchildren\",\"type\":\"singleInstance\"}}],\"bccRecipients\":[],\"body\":{\"content\":\"<html><head>\\r\\n<metahttp-equiv=\\\"Content-Type\\\"content=\\\"text/html;charset=utf-8\\\"><styletype=\\\"text/css\\\"style=\\\"display:none\\\">\\r\\n<!--\\r\\np\\r\\n\\t{margin-top:0;\\r\\n\\tmargin-bottom:0}\\r\\n-->\\r\\n</style></head><bodydir=\\\"ltr\\\"><divclass=\\\"elementToProof\\\"style=\\\"font-family:Calibri,Arial,Helvetica,sans-serif;font-size:12pt;color:rgb(0,0,0);background-color:rgb(255,255,255)\\\">Lookingtodothis </div></body></html>\",\"contentType\":\"html\"}," +
		"\"bodyPreview\":\"Lookingtodothis\",\"ccRecipients\":[],\"conversationId\":\"AAQkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOAAQADGvj5ACBMdGpESX4xSOxCo=\",\"conversationIndex\":\"AQHZNkPmMa+PkAIEx0akRJfjFI7EKg==\",\"flag\":{\"flagStatus\":\"notFlagged\"}," +
		"\"from\":{\"emailAddress\":{\"address\":\"" + defaultMessageFrom + "\",\"name\":\"" + defaultAlias + "\"}},\"hasAttachments\":true,\"importance\":\"normal\",\"inferenceClassification\":\"focused\"," +
		"\"internetMessageId\":\"<SJ0PR17MB56220B4F6A443386A11D5154C3D19@SJ0PR17MB5622.namprd17.prod.outlook.com>\",\"isDeliveryReceiptRequested\":false,\"isDraft\":false,\"isRead\":true,\"isReadReceiptRequested\":false," +
		"\"parentFolderId\":\"AQMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4ADVkZWQwNmNlMTgALgAAAw_9XBStqZdPuOVIalVTz7sBAHzBhzS2FNNNiXdGkRghSr4AAAIBDAAAAA==\",\"receivedDateTime\":\"2023-02-01T13:48:47Z\",\"replyTo\":[]," +
		"\"sender\":{\"emailAddress\":{\"address\":\"" + defaultMessageSender + "\",\"name\":\"" + defaultAlias + "\"}},\"sentDateTime\":\"2023-02-01T13:48:46Z\"," +
		"\"subject\":\"" + subject + "\",\"toRecipients\":[{\"emailAddress\":{\"address\":\"" + defaultMessageTo + "\",\"name\":\"" + defaultAlias + "\"}}]," +
		"\"webLink\":\"https://outlook.office365.com/owa/?ItemID=AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8%2B7BwB8wYc0thTTTYl3RpEYIUq%2BAAAAAAEMAAB8wYc0thTTTYl3RpEYIUq%2BAADFfThMAAA%3D&exvsurl=1&viewmodel=ReadMessageItem\"}"

	return []byte(message)
}

func GetMockMessageWithItemAttachmentMail(subject string) []byte {
	//nolint:lll
	template := `{
		"@odata.context": "https://graph.microsoft.com/v1.0/$metadata#users('f435c656-f8b2-4d71-93c3-6e092f52a167')/messages(attachments())/$entity",
		"@odata.etag": "W/\"CQAAABYAAAB8wYc0thTTTYl3RpEYIUq+AADKTqGD\"",
		"id": "AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEMAAB8wYc0thTTTYl3RpEYIUq_AADKo35RAAA=",
		"createdDateTime": "2023-02-06T18:41:35Z",
		"lastModifiedDateTime": "2023-02-06T18:41:37Z",
		"changeKey": "CQAAABYAAAB8wYc0thTTTYl3RpEYIUq+AADKTqGD",
		"categories": [],
		"receivedDateTime": "2023-02-06T18:41:35Z",
		"sentDateTime": "2023-02-06T18:41:32Z",
		"hasAttachments": true,
		"internetMessageId": "<SJ0PR17MB56227C78330FF400E26FE856C3DA9@SJ0PR17MB5622.namprd17.prod.outlook.com>",
		"subject": "%[1]v",
		"bodyPreview": "One Item Attachment object type Email",
		"importance": "normal",
		"parentFolderId": "AQMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4ADVkZWQwNmNlMTgALgAAAw_9XBStqZdPuOVIalVTz7sBAHzBhzS2FNNNiXdGkRghSr4AAAIBDAAAAA==",
		"conversationId": "AAQkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOAAQAHfjAVIQ5hdDghDme0-iu3E=",
		"conversationIndex": "AQHZOlqMd+MBUhDmF0OCEOZ7T+K7cQ==",
		"isDeliveryReceiptRequested": false,
		"isReadReceiptRequested": false,
		"isRead": false,
		"isDraft": false,
		"webLink": "https://outlook.office365.com/owa/?ItemID=AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8%2B7BwB8wYc0thTTTYl3RpEYIUq%2BAAAAAAEMAAB8wYc0thTTTYl3RpEYIUq%2BAADKo35RAAA%3D&exvsurl=1&viewmodel=ReadMessageItem",
		"inferenceClassification": "focused",
		"body": {
			"contentType": "html",
			"content": "<html><head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><style type=\"text/css\" style=\"display:none\">\r\n<!--\r\np\r\n\t{margin-top:0;\r\n\tmargin-bottom:0}\r\n-->\r\n</style></head><body dir=\"ltr\"><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">One Item Attachment object type Email</div></body></html>"
		},
		"sender": {
			"emailAddress": {
				"name": "Dustin Abbot",
				"address": "dustina@8qzvrj.onmicrosoft.com"
			}
		},
		"from": {
			"emailAddress": {
				"name": "Dustin Abbot",
				"address": "dustina@8qzvrj.onmicrosoft.com"
			}
		},
		"toRecipients": [
			{
				"emailAddress": {
					"name": "Dustin Abbot",
					"address": "dustina@8qzvrj.onmicrosoft.com"
				}
			}
		],
		"ccRecipients": [],
		"bccRecipients": [],
		"replyTo": [],
		"flag": {
			"flagStatus": "notFlagged"
		},
		"attachments": [
				  {
				"@odata.context": "https://graph.microsoft.com/v1.0/$metadata#users('f435c656-f8b2-4d71-93c3-6e092f52a167')/messages('AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEMAAB8wYc0thTTTYl3RpEYIUq_AADKo35RAAA%3D')/attachments(microsoft.graph.itemAttachment/item())/$entity",
				"@odata.type": "#microsoft.graph.itemAttachment",
				"id": "AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEMAAB8wYc0thTTTYl3RpEYIUq_AADKo35RAAABEgAQADuOMl7I1J5ElEszCyRTu1g=",
				"lastModifiedDateTime": "2023-02-06T18:41:35Z",
				"name": "Diego Updated A Message",
				"contentType": null,
				"size": 28965,
				"isInline": false,
				"item@odata.associationLink": "https://graph.microsoft.com/v1.0/users('f435c656-f8b2-4d71-93c3-6e092f52a167')/messages('')/$ref",
				"item@odata.navigationLink": "https://graph.microsoft.com/v1.0/users('f435c656-f8b2-4d71-93c3-6e092f52a167')/messages('')",
				"item": {
					"@odata.type": "#microsoft.graph.message",
					"id": "",
					"createdDateTime": "2023-02-06T18:41:35Z",
					"lastModifiedDateTime": "2023-02-06T18:41:35Z",
					"receivedDateTime": "2022-07-22T19:36:53Z",
					"sentDateTime": "2022-07-22T19:36:53Z",
					"hasAttachments": false,
					"internetMessageId": "<99b8b235d160427998cd63bc9690d047-JFBVALKQOJXWILKCJQZFA7CPGM3DKTLFONZWCZ3FINSW45DFOJ6E2Q2ENFTWK43UL4YDIMJQGIZHYU3NORYA====@microsoft.com>",
					"subject": "Diego Updated A Message",
					"bodyPreview": "Message center announcements,\r\nApril 4-10, 2022\r\n8QZVRJ\r\nMajor updates\r\n(Updated) Teams Meeting Recordings Auto-Expiration in OneDrive and SharePoint\r\nAct by: March 25\r\nMC274188 | April 6 - Updated April 06, 2022: We have begun rolling this out and will b",
					"importance": "normal",
					"conversationId": "AAQkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOAAQADRbRXFBoltFumh-b67kQxM=",
					"conversationIndex": "AQHYngJlNFtFcUGiW0W6aH9vruRDEw==",
					"isDeliveryReceiptRequested": false,
					"isReadReceiptRequested": false,
					"isRead": true,
					"isDraft": false,
					"webLink": "https://outlook.office365.com/owa/?AttachmentItemID=AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8%2B7BwB8wYc0thTTTYl3RpEYIUq%2BAAAAAAEMAAB8wYc0thTTTYl3RpEYIUq%2BAADKo35RAAABEgAQADuOMl7I1J5ElEszCyRTu1g%3D&exvsurl=1&viewmodel=ItemAttachment",
					"body": {
						"contentType": "html",
						"content": "<html><head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><meta content=\"width=device-width, initial-scale=1\"><meta content=\"IE=edge\"><style><!--body, table, td{font-family:Segoe UI,Helvetica,Arial,sans-serif!important}a{color:#006CBE;text-decoration:none}--></style></head><body><div style=\"background:white; min-height:100vh; color:#323130; font-size:14px\"><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" height=\"100%\"><tbody><tr><td></td><td width=\"640\"><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" style=\"min-width:100%; background:white\"><tbody><tr><td style=\"padding:24px 24px 45px\"><img width=\"100\" height=\"21\" alt=\"Microsoft\" src=\"https://eus-contentstorage.osi.office.net/images/retailer.images/centralizeddeployment/logos/112fec798b78aa02.png\"> </td></tr><tr><td style=\"font-size:28px; padding:0 24px; font-weight:bold; color:#000000\">Message center announcements,<br>April 4-10, 2022</td></tr><tr><td style=\"color:#323130; padding:20px 24px 40px 24px\"><span style=\"font-weight:600\">8QZVRJ</span></td></tr><tr><td style=\"padding:0 24px 44px\"><div style=\"margin-bottom:20px\"><div style=\"font-weight:bold; font-size:16px; color:#323130\">Major updates</div><table><tbody><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/office_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">(Updated) Teams Meeting Recordings Auto-Expiration in OneDrive and SharePoint</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">Act by: March 25</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC274188 | April 6 - Updated April 06, 2022: We have begun rolling this out and will be monitoring customer reported issues to ensure a smooth deployment. To ensure the best experience we have postponed the start of final stage of this change until late March. The final stage is the part of the feature that actually stamps the expiration date on the file and physically deletes the file based on that stamped expiration date.For any tenant that does not have a custom policy in place already, we are updating the...</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC274188?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200821.001/assets/brand-icons/product-monoline/png/admin_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">(Updated) OAuth interface for Office 365 Reporting web service</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC318316 | April 5 - Updated April 05, 2022: We have updated the rollout timeline below and provided additional details for clarity. Thank you for your patience.Currently, users accessing Reporting Web service use “Basic Authentication” and must provide their credentials. With this feature update, Microsoft will improve the security of your tenant by replacing “Basic Authentication” access in favor of the recommended OAuth user interface which is where we will continue to invest our development resources.</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC318316?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr></tbody></table></div><div style=\"margin-bottom:20px\"><div style=\"font-weight:bold; font-size:16px; color:#323130\">Additional updates</div><table><tbody><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/office_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Microsoft Defender for Office 365: New default (URL click) alert policy</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC355215 | April 9 - Microsoft Defender for Office 365 offers many default alert policies. One of the current default alert policy, A potentially malicious URL click was detected, generates alerts when users click URLs (which are potentially malicious) in email messages. Currently, there are two scenarios that generate this alert.The URL verdict changes to bad after the user has already clicked it.A user clicks through (overrides the warning page) for a known, bad URL.As part of the work to expand the scope of the...</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC355215?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/sharepoint_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">SharePoint: Updates to channel site layouts and parent site settings</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC355214 | April 9 - Two updates are coming for the SharePoint sites that get created for every team and private channel in Microsoft Teams. The first update is to simplify the layout for all private channel sites to make them more useful and easier to navigate between the parent site and the associated team. The second update is to adjust the way site theme and navigation is inherited when the Teams-connected sites are added to a hub site.This message is associated with Microsoft 365 Roadmap ID 88963</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC355214?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200821.001/assets/brand-icons/product-monoline/png/admin_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Service health dashboard refresh and ability to show multiple services being impacted by a single issue</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC355213 | April 9 - When there is a broadly impacting issue that affects multiple services across the M365 services, Microsoft will create a single incident that will properly reflect the overall status of the different services if affects. In the message details of this incident, we will call out specific services impacted. With this change we will also be giving a refreshed user experience to the service health dashboard.This message is associated with Microsoft 365 Roadmap ID 88666.</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC355213?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200821.001/assets/brand-icons/product-monoline/png/admin_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Microsoft Information Protection: Updates to Sensitive Information Types (SITs) definitions</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC355212 | April 9 - In our continued efforts to improve the accuracy of out-of-the-box sensitive information types (SITs), we will be updating the definition of five existing SITs for better accuracy and to stay updated with the latest definition published by relevant authorities.</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC355212?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200821.001/assets/brand-icons/product-monoline/png/admin_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Insider Risk Management: New feature enhancements coming to public preview</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC354698 | April 7 - In April, we will be rolling out two new features in public preview to enhance your Insider risk management policies and workflows. Insider Risk Management in Microsoft 365 correlates various signals from the chip to the cloud to identify potential malicious or inadvertent insider risks, such as IP theft, security and policy violations, and more. Built with privacy by design, users are pseudonymized by default, and role-based access controls and audit logs are in place to help ensure user-level...</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC354698?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200821.001/assets/brand-icons/product-monoline/png/admin_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Microsoft 365 admin center: Prepare your users for Internet Explorer retirement with &quot;Reload in IE mode&quot; instructions</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">Act by: June 15</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC337246 | April 7 - Updated April 07, 2022: We have updated the rollout timeline below. Thank you for your patience.With the upcoming Internet Explorer 11 desktop application retirement (June 15th), for certain versions of Windows 10, we recommend that you move users to Microsoft Edge and explain how to access legacy websites by reloading those sites in Internet Explorer (IE) mode.This message is associated with Microsoft 365 Roadmap ID 88905.</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC337246?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/teams_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Microsoft Teams: Microsoft Teams Rooms on Windows store application 4.12 update</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC354488 | April 7 - The below message is for organizations using Microsoft Teams Rooms for Windows. If you are not using Teams Rooms for Windows, you can disregard this message.Teams Rooms on Windows is releasing Teams room Windows store application version 4.12 that includes improvements to existing functionality and additional controls for the device administrators to control application and device behavior: IT admins can enroll a Teams rooms device to receive public preview features through XML setting. Once...</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC354488?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/teams_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Teams admin center: Update to call management settings</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC354160 | April 6 - A new update to Teams admin center will extend the ability of an Admin to configure the call management capabilities for end-users. The existing call management capabilities of call delegation and group call pickup are being extended to also include the ability to manage an end-user's call forwarding and simultaneous ringing settings as well.</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC354160?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/sharepoint_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Microsoft Lists Calendar: Unscheduled Pane</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC354159 | April 6 - The new feature, Microsoft Lists Calendar: Unscheduled Pane, will allow users to review all the items not yet appearing in the Calendar view due to missing dates. These items will appear on the Unscheduled tab within the events pane to the right of the Calendar view.This message is associated with Microsoft 365 Roadmap ID 93223.</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC354159?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/office_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Microsoft Defender for Office 365: Email and collaboration reports : Mail latency report</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC353490 | April 5 - We're updating the chart view, list view and filter options to clarify email delivery latency and latency due to detonation of the attachments / URLs (for those items that are subject to Safe Attachments and Safe Links policies).This message is associated with Microsoft 365 Roadmap ID 93213</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC353490?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/office_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Microsoft Stream: View and edit video/audio file information</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC353485 | April 5 - We will be updating the web player for audio &amp; video files that are stored in OneDrive and SharePoint (including when those files are embedded in Teams, Yammer, and SharePoint web parts) with two new features called About and Custom Thumbnail.This message is associated with Microsoft 365 Roadmap ID 93224 and 93225</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC353485?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/office_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Announcing Microsoft Endpoint Manager remote help</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC353143 | April 5 - Remote help is now generally available, and you can get started from the Microsoft Endpoint Manager admin center. This secure, cloud-based remote assistance tool provides trusted helpdesk-to-user connection. The solution includes enterprise grade capabilities to enable IT administrators, helpdesk associates, and Windows users to resolve technical issues in real-time on enrolled and unenrolled devices. This message is associated with Microsoft 365 Roadmap ID 88504.</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC353143?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/office_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">OneDrive and SharePoint: Access your Teams standard and private channel files</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC336858 | April 5 - Updated April 05, 2022: We have updated the content below with additional details.With this new feature, users that navigate to a site in SharePoint or OneDrive will be able to access the files stored in the Teams standard and private channels associated with that site. Users will see an “In channels” section when you navigate to the default document library of a Teams-connected site. This functionality will be available in OneDrive &amp; SharePoint web, including experiences like the Move/Copy...</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC336858?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/office_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Microsoft and Office 365 users to edit assigned tasks in Project for the web project plans</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC352634 | April 4 - When Microsoft 365 and Office 365 E3 and E5 users are assigned tasks from a project plan built in Project for the web, they can update the progress and mark their tasks’ percent complete without a Project license. This editing capability will be rolling out in May 2022 to tenants with Project Plan 1, Project Plan 3, and Project Plan 5 subscriptions</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC352634?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/teams_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Soft focus and Adjust brightness in Teams video meetings</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC352623 | April 4 - Soft focus and Adjust brightness are video filters in Teams video meetings. Users will soon be able to access and apply both settings, before and during meetings. This message is associated with Microsoft 365 Roadmap ID 65944 Note: Soft focus is not available for EDU tenants.</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC352623?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr><tr><td style=\"padding-top:20px; padding-right:8px\"><img width=\"16\" height=\"16\" style=\"display:block\" src=\"https://static2.sharepointonline.com/files/fabric-cdn-prod_20200713.001/assets/brand-icons/product-monoline/png/teams_96x1.png\"></td><td style=\"font-weight:600; padding-top:20px; padding-bottom:3px\">Improved Meeting Support in Firefox Browser</td></tr><tr><td></td><td style=\"padding-top:3px; padding-bottom:3px\">MC352622 | April 4 - Microsoft Teams Meetings in the Firefox browser will now support full audio and screen sharing. This message is associated with Microsoft 365 Roadmap ID 83838</td></tr><tr><td></td><td><a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter/:/messages/MC352622?MCLinkSource=DigestMail\" target=\"_blank\" style=\"padding-top:3px\">View more</a></td></tr></tbody></table></div><div>To view all announcements, <a href=\"https://admin.microsoft.com/AdminPortal/home#/MessageCenter\" target=\"_blank\">sign in to Microsoft 365 admin center</a></div></td></tr><tr><td><table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" width=\"100%\" style=\"min-width:100%; background-color:#F3F2F1\"><tbody><tr><td style=\"padding:44px 24px 3px; font-size:10px; color:#484644\">You're subscribed to this email using danny.@8qzvrj.onmicrosoft.com. If you're an IT admin, you're subscribed by default, but you can <a href=\"https://admin.microsoft.com/adminportal/home#/MessageCenter/:/mcpreferences\" target=\"_blank\">unsubscribe at any time</a>. If you're not an IT admin, ask your admin to remove your email address from Microsoft 365 message center preferences. <br><br>This email might not include all Microsoft service updates from the past week. The content you see is based on the Microsoft services available to your organization, and the custom view and email options you (or your admin) select in <a href=\"https://admin.microsoft.com/adminportal/home#/MessageCenter/:/mcpreferences\" target=\"_blank\">Microsoft 365 message center preferences</a>. <br><br><a href=\"https://docs.microsoft.com/en-us/microsoft-365/admin/manage/language-translation-for-message-center-posts?view=o365-worldwide\" target=\"_blank\">How to view translated messages</a><br><a href=\"https://forms.office.com/Pages/ResponsePage.aspx?id=v4j5cvGGr0GRqy180BHbR-rMMzjkh-5MjuvzignA2eJUQURBR1k1MkRBUEY5MFRBSDNBWldWQlM4NC4u\" target=\"_blank\">Send us feedback about this email</a> <br><br></td></tr><tr><td style=\"padding:25px 24px 24px; font-size:12px\"><a href=\"https://go.microsoft.com/fwlink/?LinkId=521839\" target=\"_blank\" style=\"color:#696969; text-decoration:underline; text-decoration-color:#696969\">Privacy statement</a> <div style=\"color:#696969; margin-top:10px; margin-bottom:13px\">Microsoft Corporation, One Microsoft Way, Redmond WA 98052 USA</div><img width=\"94\" height=\"20\" alt=\"Microsoft\" src=\"https://eus-contentstorage.osi.office.net/images/retailer.images/centralizeddeployment/logos/112fec798b78aa02.png\"> </td></tr></tbody></table></td></tr></tbody></table></td><td></td></tr></tbody></table></div><img width=\"1\" height=\"1\" tabindex=\"-1\" aria-hidden=\"true\" alt=\"\" src=\"https://mucp.api.account.microsoft.com/m/v2/v?d=AIAADB2MK7POCPGPEDZSKDKBUO5KYKN6DTJ72CPAQBQ3KDU7UDXEXALE4M6TIXWJK7GC2RNNKAXRHWFXKYYSWJGRBINYATTS2JBPBGCB4IG7VRNOW2QWYN5I2ZIP36KKTI3VVFIS4ZIVNTUYN7DSQFBHVONJBDI&amp;i=AIAADPARFJJPFLMPTWNHE3UOC5LF5VEXU6GNERJPLOXJUNGY33QCNHPNPQ6GRONHCHKNO572L62HTG4Y5EWLIUNFNEX6XBS6UYHGEWBGZX2EKNMZSXNDX6APN2RY7DXUF6SKNI4N4EY24CFSV3HWUMC7XQOGKCHPSK4LNVKI2UJEPDYZAKF4FW47G3A6RYQUNP5PK2DE65ZYBOYC7EOGKR6FRZDVGKAJ4BXINGR72ZBCJZY54VSKRZ5L7IVGHGLKFUMA5LU43XKRQBFMDGKXFTYJTSSYQCQ\"> </body></html>"
					},
					"sender": {
						"emailAddress": {
							"name": "Microsoft 365 Message center",
							"address": "o365mc@microsoft.com"
						}
					},
					"from": {
						"emailAddress": {
							"name": "Microsoft 365 Message center",
							"address": "o365mc@microsoft.com"
						}
					},
					"toRecipients": [
						{
							"emailAddress": {
								"name": "Diego Siciliani",
								"address": "NotAMember@8qzvrj.onmicrosoft.com"
							}
						}
					],
					"flag": {
						"flagStatus": "notFlagged"
					}
				}
			}
		]
	}`

	message := fmt.Sprintf(
		template,
		subject,
	)

	return []byte(message)
}

func GetMockMessageWithNestedItemAttachmentEvent(subject string) []byte {
	//nolint:lll
	// Order of fields:
	// 1. subject
	// 2. alias
	// 3. sender address
	// 4. from address
	// 5. toRecipients email address
	template := `{
		"@odata.context": "https://graph.microsoft.com/v1.0/$metadata#users('f435c656-f8b2-4d71-93c3-6e092f52a167')/messages(attachments())/$entity",
		"@odata.etag": "W/\"CQAAABYAAAB8wYc0thTTTYl3RpEYIUq+AADFK782\"",
		"id": "AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEMAAB8wYc0thTTTYl3RpEYIUq_AADFfThSAAA=",
		"createdDateTime": "2023-02-02T21:38:27Z",
		"lastModifiedDateTime": "2023-02-02T22:42:49Z",
		"changeKey": "CQAAABYAAAB8wYc0thTTTYl3RpEYIUq+AADFK782",
		"categories": [],
		"receivedDateTime": "2023-02-02T21:38:27Z",
		"sentDateTime": "2023-02-02T21:38:24Z",
		"hasAttachments": true,
		"internetMessageId": "<SJ0PR17MB562287BE29A86751D6E77FE5C3D69@SJ0PR17MB5622.namprd17.prod.outlook.com>",
		"subject": "%[1]v",
		"bodyPreview": "Dustin,\r\n\r\nI'm here to see if we are still able to discover our object.",
		"importance": "normal",
		"parentFolderId": "AQMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4ADVkZWQwNmNlMTgALgAAAw_9XBStqZdPuOVIalVTz7sBAHzBhzS2FNNNiXdGkRghSr4AAAIBDAAAAA==",
		"conversationId": "AAQkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOAAQAB13OyMdkNJJqEaIrGi3Yjc=",
		"conversationIndex": "AQHZN06dHXc7Ix2Q0kmoRoisaLdiNw==",
		"isDeliveryReceiptRequested": false,
		"isReadReceiptRequested": false,
		"isRead": false,
		"isDraft": false,
		"webLink": "https://outlook.office365.com/owa/?ItemID=AAMkAGQ1NzTruncated",
		"inferenceClassification": "focused",
		"body": {
		  "contentType": "html",
		  "content": "<html><head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><style type=\"text/css\" style=\"display:none\">\r\n<!--\r\np\r\n\t{margin-top:0;\r\n\tmargin-bottom:0}\r\n-->\r\n</style></head><body dir=\"ltr\"><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">Dustin,</div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">I'm here to see if we are still able to discover our object.&nbsp;</div></body></html>"
		},
		"sender": {
		  "emailAddress": {
			"name": "%[2]s",
			"address": "%[3]s"
		  }
		},
		"from": {
		  "emailAddress": {
			"name": "%[2]s",
			"address": "%[4]s"
		  }
		},
		"toRecipients": [
		  {
			"emailAddress": {
			  "name": "%[2]s",
			  "address": "%[5]s"
			}
		  }
		],
		"ccRecipients": [],
		"bccRecipients": [],
		"replyTo": [],
		"flag": {
		  "flagStatus": "notFlagged"
		},
		"attachments": [
		  {
			"@odata.type": "#microsoft.graph.itemAttachment",
			"id": "AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEMAAB8wYc0thTTTYl3RpEYIUq_AADFfThSAAABEgAQAIyAgT1ZccRCjKKyF7VZ3dA=",
			"lastModifiedDateTime": "2023-02-02T21:38:27Z",
			"name": "Mail Item Attachment",
			"contentType": null,
			"size": 5362,
			"isInline": false,
			"item@odata.associationLink": "https://graph.microsoft.com/v1.0/users('f435c656-f8b2-4d71-93c3-6e092f52a167')/messages('')/$ref",
			"item@odata.navigationLink": "https://graph.microsoft.com/v1.0/users('f435c656-f8b2-4d71-93c3-6e092f52a167')/messages('')",
			"item": {
			  "@odata.type": "#microsoft.graph.message",
			  "id": "",
			  "createdDateTime": "2023-02-02T21:38:27Z",
			  "lastModifiedDateTime": "2023-02-02T21:38:27Z",
			  "receivedDateTime": "2023-02-01T13:48:47Z",
			  "sentDateTime": "2023-02-01T13:48:46Z",
			  "hasAttachments": true,
			  "internetMessageId": "<SJ0PR17MB56220B4F6A443386A11D5154C3D19@SJ0PR17MB5622.namprd17.prod.outlook.com>",
			  "subject": "Mail Item Attachment",
			  "bodyPreview": "Lookingtodothis",
			  "importance": "normal",
			  "conversationId": "AAQkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOAAQAMNK0NU7Kx5GhAaHdzhfSRU=",
			  "conversationIndex": "AQHZN02pw0rQ1TsrHkaEBod3OF9JFQ==",
			  "isDeliveryReceiptRequested": false,
			  "isReadReceiptRequested": false,
			  "isRead": true,
			  "isDraft": false,
			  "webLink": "https://outlook.office365.com/owa/?AttachmentItemID=AAMkAGQ1NzViZTdhLTEwMTM",
			  "body": {
				"contentType": "html",
				"content": "<html><head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><metahttp-equiv=\"Content-Type\"content=\"text html;charset=\"utf-8&quot;\"><styletype=\"text css?style=\"display:none\"><!--\r\np\r\n\t{margin-top:0;\r\n\tmargin-bottom:0}\r\n--><bodydir=\"ltr\"><divclass=\"elementToProof\"style=\"font-family:Calibri,Arial,Helvetica,sans-serif;font-size:12pt;color:rgb(0,0,0);background-color:rgb(255,255,255)\"></head><body>Lookingtodothis&nbsp; <div></div></body></html>"
			  },
			  "sender": {
				"emailAddress": {
				  "name": "A Stranger",
				  "address": "foobar@8qzvrj.onmicrosoft.com"
				}
			  },
			  "from": {
				"emailAddress": {
				  "name": "A Stranger",
				  "address": "foobar@8qzvrj.onmicrosoft.com"
				}
			  },
			  "toRecipients": [
				{
				  "emailAddress": {
					"name": "Direct Report",
					"address":  "notAvailable@8qzvrj.onmicrosoft.com"
				  }
				}
			  ],
			  "flag": {
				"flagStatus": "notFlagged"
			  },
			  "attachments": [
				{
				  "@odata.type": "#microsoft.graph.itemAttachment",
				  "id": "AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8_7BwB8wYc0thTTTYl3RpEYIUq_AAAAAAEMAAB8wYc0thTTTYl3RpEYIUq_AADFfThSAAACEgAQAIyAgT1ZccRCjKKyF7VZ3dASABAAuYCb3N2YZ02RpJrZPzCBFQ==",
				  "lastModifiedDateTime": "2023-02-02T21:38:27Z",
				  "name": "Holidayevent",
				  "contentType": null,
				  "size": 2331,
				  "isInline": false,
				  "item@odata.associationLink": "https://graph.microsoft.com/v1.0/users('f435c656-f8b2-4d71-93c3-6e092f52a167')/events('')/$ref",
				  "item@odata.navigationLink": "https://graph.microsoft.com/v1.0/users('f435c656-f8b2-4d71-93c3-6e092f52a167')/events('')",
				  "item": {
					"@odata.type": "#microsoft.graph.event",
					"id": "",
					"createdDateTime": "2023-02-02T21:38:27Z",
					"lastModifiedDateTime": "2023-02-02T21:38:27Z",
					"originalStartTimeZone": "tzone://Microsoft/Utc",
					"originalEndTimeZone": "tzone://Microsoft/Utc",
					"reminderMinutesBeforeStart": 0,
					"isReminderOn": false,
					"hasAttachments": false,
					"subject": "Discuss Gifts for Children",
					"isAllDay": false,
					"isCancelled": false,
					"isOrganizer": true,
					"responseRequested": true,
					"type": "singleInstance",
					"isOnlineMeeting": false,
					"isDraft": true,
					"body": {
					  "contentType": "html",
					  "content": "<html><head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><metahttp-equiv=\"Content-Type\"content=\"text html;charset=\"utf-8&quot;\"></head><body>Let'slookforfunding! </body></html>"
					},
					"start": {
					  "dateTime": "2016-12-02T18:00:00.0000000Z",
					  "timeZone": "UTC"
					},
					"end": {
					  "dateTime": "2016-12-02T19:00:00.0000000Z",
					  "timeZone": "UTC"
					},
					"organizer": {
					  "emailAddress": {
						"name": "Event Manager",
						"address": "philonis@8qzvrj.onmicrosoft.com"
					  }
					}
				  }
				}
			  ]
			}
		  }
		]
	  }`

	message := fmt.Sprintf(
		template,
		subject,
		defaultAlias,
		defaultMessageSender,
		defaultMessageFrom,
		defaultMessageTo,
	)

	return []byte(message)
}
