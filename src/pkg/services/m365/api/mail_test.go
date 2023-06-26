package api_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	exchMock "github.com/alcionai/corso/src/internal/m365/exchange/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/mock"
)

type MailAPIUnitSuite struct {
	tester.Suite
}

func TestMailAPIUnitSuite(t *testing.T) {
	suite.Run(t, &MailAPIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *MailAPIUnitSuite) TestMailInfo() {
	initial := time.Now()

	tests := []struct {
		name     string
		msgAndRP func() (models.Messageable, *details.ExchangeInfo)
	}{
		{
			name: "Empty message",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				msg := models.NewMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)

				i := &details.ExchangeInfo{
					ItemType:  details.ExchangeMail,
					Recipient: []string{},
					Created:   initial,
					Modified:  initial,
				}
				return msg, i
			},
		},
		{
			name: "Just sender",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				sender := "foo@bar.com"
				sr := models.NewRecipient()
				sea := models.NewEmailAddress()
				msg := models.NewMessage()
				sea.SetAddress(&sender)
				sr.SetEmailAddress(sea)
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				msg.SetSender(sr)
				i := &details.ExchangeInfo{
					ItemType:  details.ExchangeMail,
					Recipient: []string{},
					Sender:    sender,
					Created:   initial,
					Modified:  initial,
				}
				return msg, i
			},
		},
		{
			name: "Just subject",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				subject := "Hello world"
				msg := models.NewMessage()
				msg.SetSubject(&subject)
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				i := &details.ExchangeInfo{
					ItemType:  details.ExchangeMail,
					Subject:   subject,
					Created:   initial,
					Recipient: []string{},
					Modified:  initial,
				}
				return msg, i
			},
		},
		{
			name: "Just receivedtime",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				now := time.Now()
				msg := models.NewMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				msg.SetReceivedDateTime(&now)
				i := &details.ExchangeInfo{
					ItemType:  details.ExchangeMail,
					Recipient: []string{},
					Received:  now,
					Created:   initial,
					Modified:  initial,
				}
				return msg, i
			},
		},
		{
			name: "All fields",
			msgAndRP: func() (models.Messageable, *details.ExchangeInfo) {
				sender := "foo@bar.com"
				receiver := "foofoo@bar.com"
				subject := "Hello world"
				now := time.Now()
				sr := models.NewRecipient()
				sea := models.NewEmailAddress()
				recv := models.NewRecipient()
				req := models.NewEmailAddress()
				recvs := make([]models.Recipientable, 0)
				msg := models.NewMessage()
				msg.SetCreatedDateTime(&initial)
				msg.SetLastModifiedDateTime(&initial)
				sea.SetAddress(&sender)
				sr.SetEmailAddress(sea)
				msg.SetSender(sr)
				req.SetAddress(&receiver)
				recv.SetEmailAddress(req)
				msg.SetSubject(&subject)
				msg.SetReceivedDateTime(&now)
				recvs = append(recvs, recv, sr)
				msg.SetToRecipients(recvs)
				i := &details.ExchangeInfo{
					ItemType:  details.ExchangeMail,
					Sender:    sender,
					Subject:   subject,
					Recipient: []string{receiver, sender},
					Received:  now,
					Created:   initial,
					Modified:  initial,
				}
				return msg, i
			},
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			msg, expected := tt.msgAndRP()
			assert.Equal(suite.T(), expected, api.MailInfo(msg, 0))
		})
	}
}

func (suite *MailAPIUnitSuite) TestBytesToMessagable() {
	table := []struct {
		name        string
		byteArray   []byte
		checkError  assert.ErrorAssertionFunc
		checkObject assert.ValueAssertionFunc
	}{
		{
			name:        "Empty Bytes",
			byteArray:   make([]byte, 0),
			checkError:  assert.Error,
			checkObject: assert.Nil,
		},
		{
			name:        "aMessage bytes",
			byteArray:   exchMock.MessageBytes("m365 mail support test"),
			checkError:  assert.NoError,
			checkObject: assert.NotNil,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := api.BytesToMessageable(test.byteArray)
			test.checkError(t, err, clues.ToCore(err))
			test.checkObject(t, result)
		})
	}
}

type MailAPIIntgSuite struct {
	tester.Suite
	cts         clientTesterSetup
	credentials account.M365Config
	ac          api.Client
	user        string
}

// We do end up mocking the actual request, but creating the rest
// similar to full integration tests.
func TestMailAPIIntgSuite(t *testing.T) {
	suite.Run(t, &MailAPIIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *MailAPIIntgSuite) SetupSuite() {
	t := suite.T()

	suite.cts = newClientTesterSetup(t)

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.credentials = m365
	suite.ac, err = mock.NewClient(m365)
	require.NoError(t, err, clues.ToCore(err))

	suite.user = tester.M365UserID(suite.T())
}

func getJSONObject(t *testing.T, thing serialization.Parsable) map[string]interface{} {
	sw := kjson.NewJsonSerializationWriter()

	err := sw.WriteObjectValue("", thing)
	require.NoError(t, err, "serialize")

	content, err := sw.GetSerializedContent()
	require.NoError(t, err, "serialize")

	var out map[string]interface{}
	err = json.Unmarshal([]byte(content), &out)
	require.NoError(t, err, "unmarshall")

	return out
}

func (suite *MailAPIIntgSuite) TestHugeAttachmentListDownload() {
	mid := "fake-message-id"
	aid := "fake-attachment-id"

	tests := []struct {
		name            string
		setupf          func()
		attachmentCount int
		size            int64
		expect          assert.ErrorAssertionFunc
	}{
		{
			name: "no attachments",
			setupf: func() {
				mitem := models.NewMessage()
				mitem.SetId(&mid)

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/user/messages/" + mid).
					Reply(200).
					JSON(getJSONObject(suite.T(), mitem))
			},
			expect: assert.NoError,
		},
		{
			name: "fetch with attachment",
			setupf: func() {
				mitem := models.NewMessage()
				mitem.SetId(&mid)
				mitem.SetHasAttachments(ptr.To(true))

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/user/messages/" + mid).
					Reply(200).
					JSON(getJSONObject(suite.T(), mitem))

				atts := models.NewAttachmentCollectionResponse()
				aitem := models.NewAttachment()

				asize := int32(50)
				aitem.SetSize(&asize)
				atts.SetValue([]models.Attachmentable{aitem})

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/user/messages/" + mid + "/attachments").
					Reply(200).
					JSON(getJSONObject(suite.T(), atts))
			},
			attachmentCount: 1,
			size:            50,
			expect:          assert.NoError,
		},
		{
			name: "fetch individual attachment",
			setupf: func() {
				truthy := true
				mitem := models.NewMessage()
				mitem.SetId(&mid)
				mitem.SetHasAttachments(&truthy)

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/user/messages/" + mid).
					Reply(200).
					JSON(getJSONObject(suite.T(), mitem))

				atts := models.NewAttachmentCollectionResponse()
				aitem := models.NewAttachment()
				aitem.SetId(&aid)

				asize := int32(200)
				aitem.SetSize(&asize)

				atts.SetValue([]models.Attachmentable{aitem})

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/user/messages/" + mid + "/attachments").
					Reply(503)

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/user/messages/" + mid + "/attachments").
					Reply(200).
					JSON(getJSONObject(suite.T(), atts))

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/user/messages/" + mid + "/attachments/" + aid).
					Reply(200).
					JSON(getJSONObject(suite.T(), aitem))
			},
			attachmentCount: 1,
			size:            200,
			expect:          assert.NoError,
		},
		{
			name: "fetch multiple individual attachments",
			setupf: func() {
				truthy := true
				mitem := models.NewMessage()
				mitem.SetId(&mid)
				mitem.SetHasAttachments(&truthy)

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/user/messages/" + mid).
					Reply(200).
					JSON(getJSONObject(suite.T(), mitem))

				atts := models.NewAttachmentCollectionResponse()
				aitem := models.NewAttachment()
				aitem.SetId(&aid)

				asize := int32(200)
				aitem.SetSize(&asize)

				atts.SetValue([]models.Attachmentable{aitem, aitem, aitem, aitem, aitem})

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/user/messages/" + mid + "/attachments").
					Reply(503)

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/user/messages/" + mid + "/attachments").
					Reply(200).
					JSON(getJSONObject(suite.T(), atts))

				for i := 0; i < 5; i++ {
					gock.New("https://graph.microsoft.com").
						Get("/v1.0/users/user/messages/" + mid + "/attachments/" + aid).
						Reply(200).
						JSON(getJSONObject(suite.T(), aitem))
				}
			},
			attachmentCount: 5,
			size:            200,
			expect:          assert.NoError,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			defer gock.Off()
			tt.setupf()

			item, _, err := suite.ac.Mail().GetItem(ctx, "user", mid, false, fault.New(true))
			tt.expect(t, err)

			it, ok := item.(models.Messageable)
			require.True(t, ok, "convert to messageable")

			var size int64
			mailBody := it.GetBody()
			if mailBody != nil {
				content := ptr.Val(mailBody.GetContent())
				if len(content) > 0 {
					size = int64(len(content))
				}
			}

			attachments := it.GetAttachments()
			for _, attachment := range attachments {
				size = +int64(*attachment.GetSize())
			}

			assert.Equal(t, *it.GetId(), mid)
			assert.Equal(t, tt.attachmentCount, len(attachments), "attachment count")
			assert.Equal(t, tt.size, size, "mail size")
			assert.True(t, gock.IsDone(), "made all requests")
		})
	}
}

func (suite *MailAPIIntgSuite) TestRestoreLargeAttachment() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	userID := tester.M365UserID(suite.T())

	folderName := testdata.DefaultRestoreConfig("maillargeattachmenttest").Location
	msgs := suite.ac.Mail()
	mailfolder, err := msgs.CreateMailFolder(ctx, userID, folderName)
	require.NoError(t, err, clues.ToCore(err))

	msg := models.NewMessage()
	msg.SetSubject(ptr.To("Mail with attachment"))

	item, err := msgs.PostItem(ctx, userID, ptr.Val(mailfolder.GetId()), msg)
	require.NoError(t, err, clues.ToCore(err))

	id, err := msgs.PostLargeAttachment(
		ctx,
		userID,
		ptr.Val(mailfolder.GetId()),
		ptr.Val(item.GetId()),
		"raboganm",
		[]byte("mangobar"),
	)
	require.NoError(t, err, clues.ToCore(err))
	require.NotEmpty(t, id, "empty id for large attachment")
}

func (suite *MailAPIIntgSuite) TestMail_malformedJSONResp() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	defer gock.Off()

	mid := "AAMkADMzZjZlNDg0LWIxNWUtNGUzOS1iNmViLWJhOGQ0NWFjOTJkNQBGAAAAAADb4v0HPlm8T5jPY51aTdZUBwCbSUNVDx_PTKATxbYuCKZlAAAAAAEMAACbSUNVDx_PTKATxbYuCKZlAABPQ75_AAA="

	gock.New("https://graph.microsoft.com").
		Get("/v1.0/users/rfinders@10rqc2.onmicrosoft.com/messages/"+mid).
		Reply(200).
		AddHeader("Content-Type", "application/json; odata.metadata=minimal; odata.streaming=true; IEEE754Compatible=false; charset=utf-8").
		AddHeader("Content-Encoding", "gzip").
		BodyString(string(m))

	_, _, err := suite.ac.Mail().GetItem(ctx, "rfinders@10rqc2.onmicrosoft.com", mid, false, fault.New(true))
	require.Error(t, err, clues.ToCore(err))
	require.Contains(t, err.Error(), "invalid json type")

	assert.True(t, gock.IsDone(), "made all requests")
}

const m = `{
    "@odata.context": "https://graph.microsoft.com/v1.0/$metadata#users('RFinders%4010rqc2.onmicrosoft.com')/messages/$entity",
    "@odata.etag": "W/\"CQAAABYAAACbSUNVDx+PTKATxbYuCKZlAABPIl1/\"",
    "id": "AAMkADMzZjZlNDg0LWIxNWUtNGUzOS1iNmViLWJhOGQ0NWFjOTJkNQBGAAAAAADb4v0HPlm8T5jPY51aTdZUBwCbSUNVDx_PTKATxbYuCKZlAAAAAAEMAACbSUNVDx_PTKATxbYuCKZlAABPQ75_AAA=",
    "createdDateTime": "2023-06-22T17:49:27Z",
    "lastModifiedDateTime": "2023-06-22T17:49:29Z",
    "changeKey": "CQAAABYAAACbSUNVDx+PTKATxbYuCKZlAABPIl1/",
    "categories": [],
    "receivedDateTime": "2023-06-22T17:49:28Z",
    "sentDateTime": "2023-06-22T17:49:21Z",
    "hasAttachments": false,
    "internetMessageId": "<PH8PR14MB7405D7605066409E4565A155D722A@PH8PR14MB7405.namprd14.prod.outlook.com>",
    "subject": "zzazil",
    "bodyPreview": "lizazz",
    "importance": "normal",
    "parentFolderId": "AQMkADMzAGY2ZTQ4NC1iMTVlLTRlMzktYjZlYi1iYThkNDVhYzkyZDUALgAAA9vi-Qc_WbxPmM9jnVpN1lQBAJtJQ1UPH49MoBPFti4IpmUAAAIBDAAAAA==",
    "conversationId": "AAQkADMzZjZlNDg0LWIxNWUtNGUzOS1iNmViLWJhOGQ0NWFjOTJkNQAQABMgDVOQI6NNs9ZEbpFe4rY=",
    "conversationIndex": "AQHZpTHUEyANU5Ajo02z1kRukV7itg==",
    "isDeliveryReceiptRequested": null,
    "isReadReceiptRequested": false,
    "isRead": false,
    "isDraft": false,
    "webLink": "https://outlook.office365.com/owa/?ItemID=AAMkADMzZjZlNDg0LWIxNWUtNGUzOS1iNmViLWJhOGQ0NWFjOTJkNQBGAAAAAADb4v0HPlm8T5jPY51aTdZUBwCbSUNVDx%2BPTKATxbYuCKZlAAAAAAEMAACbSUNVDx%2BPTKATxbYuCKZlAABPQ75%2BAAA%3D&exvsurl=1&viewmodel=ReadMessageItem",
    "inferenceClassification": "focused",
    "body": {
        "contentType": "html",
        "content": "<html><head>\r\n<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"><meta name=\"Generator\" content=\"Microsoft Word 15 (filtered medium)\"><style>\r\n<!--\r\n@font-face\r\n\t{font-family:\"Cambria Math\"}\r\n@font-face\r\n\t{font-family:Calibri}\r\np.MsoNormal, li.MsoNormal, div.MsoNormal\r\n\t{margin:0in;\r\n\tfont-size:11.0pt;\r\n\tfont-family:\"Calibri\",sans-serif}\r\nspan.EmailStyle17\r\n\t{font-family:\"Calibri\",sans-serif;\r\n\tcolor:windowtext}\r\n.MsoChpDefault\r\n\t{font-family:\"Calibri\",sans-serif}\r\n@page WordSection1\r\n\t{margin:1.0in 1.0in 1.0in 1.0in}\r\ndiv.WordSection1\r\n\t{}\r\n-->\r\n</style></head><body lang=\"EN-US\" link=\"#0563C1\" vlink=\"#954F72\" style=\"word-wrap:break-word\"><div class=\"WordSection1\"><p class=\"MsoNormal\">lizazz</p></div></body></html>"
    },
    "sender": {
        "emailAddress": {
            "name": "Ryan Keepers",
            "address": "rkeepers@alcion.ai"
        }
    },
    "from": {
        "emailAddress": {
            "name": "Ryan Keepers",
            "address": "rkeepers@alcion.ai"
        }
    },
    "toRecipients": [
        {
            "emailAddress": {
                "name": "Finders",
                "address": "RFinders@10rqc2.onmicrosoft.com"
            }
        }
    ],
    "ccRecipients": [],
    "bccRecipients": [],
    "replyTo": [],
    "flag": {
        "flagStatus": "notFlagged"
    }`

// should have an ending curly bracket.
// }`
