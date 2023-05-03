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
	"github.com/alcionai/corso/src/internal/connector/exchange/api"
	"github.com/alcionai/corso/src/internal/connector/exchange/api/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
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
			assert.Equal(suite.T(), expected, api.MailInfo(msg))
		})
	}
}

type MailAPIIntgSuite struct {
	tester.Suite
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
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *MailAPIIntgSuite) SetupSuite() {
	t := suite.T()

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
				atts.SetValue([]models.Attachmentable{aitem})

				gock.New("https://graph.microsoft.com").
					Get("/v1.0/users/user/messages/" + mid + "/attachments").
					Reply(200).
					JSON(getJSONObject(suite.T(), atts))
			},
			attachmentCount: 1,
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
			expect:          assert.NoError,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			defer gock.Off()
			tt.setupf()

			item, _, err := suite.ac.Mail().GetItem(ctx, "user", mid, false, fault.New(true))
			tt.expect(suite.T(), err)

			it, ok := item.(models.Messageable)
			require.True(suite.T(), ok, "convert to messageable")

			assert.Equal(suite.T(), *it.GetId(), mid)
			assert.Equal(suite.T(), tt.attachmentCount, len(it.GetAttachments()), "attachment count")
			assert.True(suite.T(), gock.IsDone(), "made all requests")
		})
	}
}
