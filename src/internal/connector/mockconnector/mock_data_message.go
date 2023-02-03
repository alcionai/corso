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

func GetMockMessageWithNestedItemAttachmentEvent(subject string) []byte {
	//nolint:lll
	// Order of fields:
	// 1. subject
	// 2. alias
	// 3. sender address
	// 4. from alias
	// 5. from address
	// 6. toRecipients alias
	// 7. toRecipients email address
	// 8. organizer alias
	// 9. organizer email address
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
		"subject": "%s",
		"bodyPreview": "Dustin,\r\n\r\nI'm here to see if we are still able to discover our object.",
		"importance": "normal",
		"parentFolderId": "AQMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4ADVkZWQwNmNlMTgALgAAAw_9XBStqZdPuOVIalVTz7sBAHzBhzS2FNNNiXdGkRghSr4AAAIBDAAAAA==",
		"conversationId": "AAQkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOAAQAB13OyMdkNJJqEaIrGi3Yjc=",
		"conversationIndex": "AQHZN06dHXc7Ix2Q0kmoRoisaLdiNw==",
		"isDeliveryReceiptRequested": false,
		"isReadReceiptRequested": false,
		"isRead": false,
		"isDraft": false,
		"webLink": "https://outlook.office365.com/owa/?ItemID=AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8%2B7BwB8wYc0thTTTYl3RpEYIUq%2BAAAAAAEMAAB8wYc0thTTTYl3RpEYIUq%2BAADFfThSAAA%3D&exvsurl=1&viewmodel=ReadMessageItem",
		"inferenceClassification": "focused",
		"body": {
		  "contentType": "html",
		  "content": "<html><head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><style type=\"text/css\" style=\"display:none\">\r\n<!--\r\np\r\n\t{margin-top:0;\r\n\tmargin-bottom:0}\r\n-->\r\n</style></head><body dir=\"ltr\"><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">Dustin,</div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\"><br></div><div class=\"elementToProof\" style=\"font-family:Calibri,Arial,Helvetica,sans-serif; font-size:12pt; color:rgb(0,0,0); background-color:rgb(255,255,255)\">I'm here to see if we are still able to discover our object.&nbsp;</div></body></html>"
		},
		"sender": {
		  "emailAddress": {
			"name": "%s",
			"address": "%s"
		  }
		},
		"from": {
		  "emailAddress": {
			"name": "%s",
			"address": "%s"
		  }
		},
		"toRecipients": [
		  {
			"emailAddress": {
			  "name": "%s",
			  "address": "%s"
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
			  "webLink": "https://outlook.office365.com/owa/?AttachmentItemID=AAMkAGQ1NzViZTdhLTEwMTMtNGJjNi05YWI2LTg4NWRlZDA2Y2UxOABGAAAAAAAPvVwUramXT7jlSGpVU8%2B7BwB8wYc0thTTTYl3RpEYIUq%2BAAAAAAEMAAB8wYc0thTTTYl3RpEYIUq%2BAADFfThSAAABEgAQAIyAgT1ZccRCjKKyF7VZ3dA%3D&exvsurl=1&viewmodel=ItemAttachment",
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
						"name": "%s",
						"address": "%s"
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
		defaultAlias,
		defaultMessageFrom,
		defaultAlias,
		defaultMessageTo,
		defaultAlias,
		defaultEventOrganizer,
	)

	return []byte(message)
}
