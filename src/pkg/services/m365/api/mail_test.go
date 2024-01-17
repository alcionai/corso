package api

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/sanitize"
	exchMock "github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	graphTD "github.com/alcionai/corso/src/pkg/services/m365/api/graph/testdata"
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
			assert.Equal(suite.T(), expected, MailInfo(msg, 0))
		})
	}
}

// TestBytesToMessagable_InvalidError tests that the error message kiota returns
// for invalid JSON matches what we check for. This helps keep things in sync
// when kiota is updated.
func (suite *MailAPIUnitSuite) TestBytesToMessagable_InvalidError() {
	t := suite.T()
	input := exchMock.MessageWithSpecialCharacters("m365 mail support test")

	_, err := CreateFromBytes(input, models.CreateMessageFromDiscriminatorValue)
	require.Error(t, err, clues.ToCore(err))

	assert.Contains(t, err.Error(), invalidJSON)
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
		{
			name:        "malformed JSON bytes passes sanitization",
			byteArray:   exchMock.MessageWithSpecialCharacters("m365 mail support test"),
			checkError:  assert.NoError,
			checkObject: assert.NotNil,
		},
		{
			name: "invalid JSON bytes",
			byteArray: append(
				exchMock.MessageWithSpecialCharacters("m365 mail support test"),
				[]byte("}")...),
			checkError:  assert.Error,
			checkObject: assert.Nil,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := BytesToMessageable(test.byteArray)
			test.checkError(t, err, clues.ToCore(err))
			test.checkObject(t, result)
		})
	}
}

type MailAPIIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

// We do end up mocking the actual request, but creating the rest
// similar to full integration tests.
func TestMailAPIIntgSuite(t *testing.T) {
	suite.Run(t, &MailAPIIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *MailAPIIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *MailAPIIntgSuite) TestMail_attachmentListDownload() {
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

				interceptV1Path("users", "user", "messages", mid).
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), mitem))
			},
			expect: assert.NoError,
		},
		{
			name: "fetch with attachment",
			setupf: func() {
				email := models.NewMessage()
				email.SetId(&mid)
				email.SetHasAttachments(ptr.To(true))

				interceptV1Path("users", "user", "messages", mid).
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), email))

				atts := models.NewAttachmentCollectionResponse()
				attch := models.NewAttachment()

				size := int32(50)
				attch.SetSize(&size)
				atts.SetValue([]models.Attachmentable{attch})

				interceptV1Path("users", "user", "messages", mid, "attachments").
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), atts))
			},
			attachmentCount: 1,
			size:            50,
			expect:          assert.NoError,
		},
		{
			name: "fetch individual attachment",
			setupf: func() {
				email := models.NewMessage()
				email.SetId(&mid)
				email.SetHasAttachments(ptr.To(true))

				interceptV1Path("users", "user", "messages", mid).
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), email))

				atts := models.NewAttachmentCollectionResponse()
				attch := models.NewAttachment()
				attch.SetId(&aid)

				size := int32(200)
				attch.SetSize(&size)

				atts.SetValue([]models.Attachmentable{attch})

				interceptV1Path("users", "user", "messages", mid, "attachments").
					Reply(503)

				interceptV1Path("users", "user", "messages", mid, "attachments").
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), atts))

				interceptV1Path("users", "user", "messages", mid, "attachments", aid).
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), attch))
			},
			attachmentCount: 1,
			size:            200,
			expect:          assert.NoError,
		},
		{
			name: "fetch multiple individual attachments",
			setupf: func() {
				truthy := true
				email := models.NewMessage()
				email.SetId(&mid)
				email.SetHasAttachments(&truthy)

				interceptV1Path("users", "user", "messages", mid).
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), email))

				atts := models.NewAttachmentCollectionResponse()
				attch := models.NewAttachment()
				attch.SetId(&aid)

				asize := int32(200)
				attch.SetSize(&asize)

				atts.SetValue([]models.Attachmentable{attch, attch, attch, attch, attch})

				interceptV1Path("users", "user", "messages", mid, "attachments").
					Reply(503)

				interceptV1Path("users", "user", "messages", mid, "attachments").
					Reply(200).
					JSON(graphTD.ParseableToMap(suite.T(), atts))

				for i := 0; i < 5; i++ {
					interceptV1Path("users", "user", "messages", mid, "attachments", aid).
						Reply(200).
						JSON(graphTD.ParseableToMap(suite.T(), attch))
				}
			},
			attachmentCount: 5,
			size:            1000,
			expect:          assert.NoError,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			defer gock.Off()
			test.setupf()

			item, _, err := suite.its.gockAC.Mail().GetItem(
				ctx,
				"user",
				mid,
				fault.New(true))
			test.expect(t, err)

			it, ok := item.(models.Messageable)
			require.True(t, ok, "convert to messageable")

			var size int64

			if it.GetBody() != nil {
				content := ptr.Val(it.GetBody().GetContent())
				size = int64(len(content))
			}

			attachments := it.GetAttachments()
			for _, attachment := range attachments {
				size += int64(*attachment.GetSize())
			}

			assert.Equal(t, *it.GetId(), mid)
			assert.Equal(t, test.attachmentCount, len(attachments), "attachment count")
			assert.Equal(t, test.size, size, "mail size")
			assert.True(t, gock.IsDone(), "made all requests")
		})
	}
}

type attachment struct {
	name    string
	data    string
	isLarge bool
}

func createMailWithAttachment(
	ctx context.Context,
	t *testing.T,
	ac Client,
	userID string,
	mailFolder graph.Container,
	attachments ...attachment,
) models.Messageable {
	msg := models.NewMessage()

	msg.SetSubject(ptr.To("attachment test"))

	item, err := ac.Mail().PostItem(
		ctx,
		userID,
		ptr.Val(mailFolder.GetId()),
		msg)
	require.NoError(t, err, clues.ToCore(err))

	for _, attach := range attachments {
		if attach.isLarge {
			id, err := ac.Mail().PostLargeAttachment(
				ctx,
				userID,
				ptr.Val(mailFolder.GetId()),
				ptr.Val(item.GetId()),
				attach.name,
				[]byte(attach.data))
			require.NoError(t, err, clues.ToCore(err))
			require.NotEmpty(t, id, "empty id for large attachment")
		} else {
			att := models.NewFileAttachment()

			att.SetName(ptr.To(attach.name))
			att.SetContentBytes([]byte(attach.data))

			err = ac.Mail().PostSmallAttachment(
				ctx,
				userID,
				ptr.Val(mailFolder.GetId()),
				ptr.Val(item.GetId()),
				att)
			require.NoError(t, err, clues.ToCore(err))
		}
	}

	return item
}

func (suite *MailAPIIntgSuite) TestMail_PostAndGetAttachments() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	userID := tconfig.M365UserID(t)
	folderName := testdata.DefaultRestoreConfig("getattachmentstest").Location
	mailFolder, err := suite.its.ac.Mail().CreateContainer(
		ctx,
		userID,
		MsgFolderRoot,
		folderName)
	require.NoError(t, err, clues.ToCore(err))

	tests := []struct {
		name          string
		createMessage func(
			ctx context.Context,
			t *testing.T,
			userID string) models.Messageable
		verify func(t *testing.T, item models.Messageable)
	}{
		{
			name: "Single large attachment",
			createMessage: func(
				ctx context.Context,
				t *testing.T,
				userID string,
			) models.Messageable {
				return createMailWithAttachment(
					ctx,
					t,
					suite.its.ac,
					userID,
					mailFolder,
					attachment{
						name:    "abcd",
						data:    "1234567",
						isLarge: true,
					})
			},
			verify: func(t *testing.T, item models.Messageable) {
				assert.Equal(t, 1, len(item.GetAttachments()))
				assert.Equal(t, "abcd", ptr.Val(item.GetAttachments()[0].GetName()))

				// GetSize doesn't return the size of attachment content. Skip checking it.
				contentBytes, err := item.GetAttachments()[0].GetBackingStore().Get("contentBytes")
				require.NoError(t, err, clues.ToCore(err))
				assert.Equal(t, "1234567", string(contentBytes.([]byte)))
			},
		},
		{
			name: "Two attachments, one large, one small",
			createMessage: func(
				ctx context.Context,
				t *testing.T,
				userID string,
			) models.Messageable {
				return createMailWithAttachment(
					ctx,
					t,
					suite.its.ac,
					userID,
					mailFolder,
					attachment{
						name:    "abcd",
						data:    "1234567",
						isLarge: true,
					},
					attachment{
						name: "efgh",
						data: "7654321",
					})
			},
			verify: func(t *testing.T, item models.Messageable) {
				assert.Equal(t, 2, len(item.GetAttachments()))
				assert.Equal(t, "abcd", ptr.Val(item.GetAttachments()[0].GetName()))
				assert.Equal(t, "efgh", ptr.Val(item.GetAttachments()[1].GetName()))

				contentBytes, err := item.GetAttachments()[0].GetBackingStore().Get("contentBytes")
				require.NoError(t, err, clues.ToCore(err))
				assert.Equal(t, "1234567", string(contentBytes.([]byte)))

				contentBytes, err = item.GetAttachments()[1].GetBackingStore().Get("contentBytes")
				require.NoError(t, err, clues.ToCore(err))
				assert.Equal(t, "7654321", string(contentBytes.([]byte)))
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			m := test.createMessage(ctx, t, userID)

			item, _, err := suite.its.ac.Mail().GetItem(
				ctx,
				userID,
				ptr.Val(m.GetId()),
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			msg, ok := item.(models.Messageable)
			require.True(t, ok, "convert to messageable")

			test.verify(t, msg)
		})
	}
}

func (suite *MailAPIIntgSuite) TestMail_GetContainerByName() {
	var (
		t   = suite.T()
		acm = suite.its.ac.Mail()
		rc  = testdata.DefaultRestoreConfig("mail_get_container_by_name")
	)

	ctx, flush := tester.NewContext(t)
	defer flush()

	parent, err := acm.CreateContainer(ctx, suite.its.user.id, "msgfolderroot", rc.Location)
	require.NoError(t, err, clues.ToCore(err))

	table := []struct {
		name              string
		parentContainerID string
		expectErr         assert.ErrorAssertionFunc
	}{
		{
			name:      MailInbox,
			expectErr: assert.NoError,
		},
		{
			name:      "smarfs",
			expectErr: assert.Error,
		},
		{
			name:              rc.Location,
			parentContainerID: ptr.Val(parent.GetId()),
			expectErr:         assert.Error,
		},
		{
			name:              "Inbox",
			parentContainerID: ptr.Val(parent.GetId()),
			expectErr:         assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			_, err := acm.GetContainerByName(ctx, suite.its.user.id, test.parentContainerID, test.name)
			test.expectErr(t, err, clues.ToCore(err))
		})
	}

	suite.Run("child folder with same name", func() {
		pid := ptr.Val(parent.GetId())
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		child, err := acm.CreateContainer(ctx, suite.its.user.id, pid, rc.Location)
		require.NoError(t, err, clues.ToCore(err))

		result, err := acm.GetContainerByName(ctx, suite.its.user.id, pid, rc.Location)
		assert.NoError(t, err, clues.ToCore(err))
		assert.Equal(t, ptr.Val(child.GetId()), ptr.Val(result.GetId()))
	})
}

func (suite *MailAPIIntgSuite) TestMail_GetContainerByName_mocked() {
	mf := models.NewMailFolder()
	mf.SetId(ptr.To("id"))
	mf.SetDisplayName(ptr.To("display name"))

	table := []struct {
		name      string
		results   func(*testing.T) map[string]any
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name: "zero",
			results: func(t *testing.T) map[string]any {
				return graphTD.ParseableToMap(t, models.NewMailFolderCollectionResponse())
			},
			expectErr: assert.Error,
		},
		{
			name: "one",
			results: func(t *testing.T) map[string]any {
				mfcr := models.NewMailFolderCollectionResponse()
				mfcr.SetValue([]models.MailFolderable{mf})

				return graphTD.ParseableToMap(t, mfcr)
			},
			expectErr: assert.NoError,
		},
		{
			name: "two",
			results: func(t *testing.T) map[string]any {
				mfcr := models.NewMailFolderCollectionResponse()
				mfcr.SetValue([]models.MailFolderable{mf, mf})

				return graphTD.ParseableToMap(t, mfcr)
			},
			expectErr: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			ctx, flush := tester.NewContext(t)

			defer flush()
			defer gock.Off()

			interceptV1Path("users", "u", "mailFolders").
				Reply(200).
				JSON(test.results(t))

			_, err := suite.its.gockAC.
				Mail().
				GetContainerByName(ctx, "u", "", test.name)
			test.expectErr(t, err, clues.ToCore(err))
			assert.True(t, gock.IsDone())
		})
	}
}

func sendItemWithBodyAndGetSerialized(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	msgs Mail,
	userID string,
	mailFolderID string,
	subject string,
	bodyContent string,
	contentType models.BodyType,
) []byte {
	msg := models.NewMessage()
	msg.SetSubject(ptr.To(subject))

	body := models.NewItemBody()
	body.SetContent(ptr.To(bodyContent))
	body.SetContentType(ptr.To(contentType))

	msg.SetBody(body)

	item, err := msgs.PostItem(ctx, userID, mailFolderID, msg)
	require.NoError(t, err, clues.ToCore(err))

	fetched, _, err := msgs.GetItem(
		ctx,
		userID,
		ptr.Val(item.GetId()),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	serialized, err := msgs.Serialize(
		ctx,
		fetched,
		userID,
		ptr.Val(item.GetId()))
	require.NoError(t, err, clues.ToCore(err))

	return serialized
}

func sendSerializedItemAndGetSerialized(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	msgs Mail,
	userID string,
	mailFolderID string,
	serializedInput []byte,
) []byte {
	msg, err := BytesToMessageable(serializedInput)
	require.NoError(t, err, clues.ToCore(err))

	item, err := msgs.PostItem(ctx, userID, mailFolderID, msg)
	require.NoError(t, err, clues.ToCore(err))

	fetched, _, err := msgs.GetItem(
		ctx,
		userID,
		ptr.Val(item.GetId()),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	serialized, err := msgs.Serialize(
		ctx,
		fetched,
		userID,
		ptr.Val(item.GetId()))
	require.NoError(t, err, clues.ToCore(err))

	return serialized
}

func (suite *MailAPIIntgSuite) TestMail_WithSpecialCharacters() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	contentRegex := regexp.MustCompile(`"content": ?"(.*?)"(?:, ?"contentType"|}}|},)`)

	htmlStub := "<html><head>\r\n<meta http-equiv=\"Content-Type\" " +
		"content=\"text/html; charset=utf-8\"><style type=\"text/css\" " +
		"style=\"display:none\">\r\n<!--\r\np\r\n\t{margin-top:0;\r\n\t" +
		"margin-bottom:0}\r\n-->\r\n</style></head>" +
		"<body dir=\"ltr\"><div " +
		"class=\"elementToProof\" style=\"font-family:Aptos,Aptos_EmbeddedFont," +
		"Aptos_MSFontService,Calibri,Helvetica,sans-serif; font-size:12pt; " +
		"color:rgb(0,0,0)\">%s</div></body></html>"

	userID := tconfig.M365UserID(suite.T())

	folderName := testdata.DefaultRestoreConfig("EscapeCharacters").Location
	msgs := suite.its.ac.Mail()
	mailfolder, err := msgs.CreateContainer(ctx, userID, MsgFolderRoot, folderName)
	require.NoError(t, err, clues.ToCore(err))

	escapeCharRanges := [][]int{
		{0x0, 0x20},
		{0x22, 0x23},
		{0x5c, 0x5d},
	}
	testSequences := []string{
		// Character code for backspace
		`\u0008`,
		`\\u0008`,
		"u0008",
		// Character code for \
		`\u005c`,
		`\\u005c`,
		"u005c",
		// Character code for "
		`\u0022`,
		`\\u0022`,
		"u0022",
		// Character code for B
		`\u0042`,
		`\\u0042`,
		"u0042",
		// G clef character (U+1D11E) (from JSON RFC).
		// TODO(ashmrtn): Uncomment this once the golang graph SDK is fixed. Right
		// now it can't properly send these code points.
		//`\uD834\uDD1E`,
		"\\n",
		"\\\n",
		"abcdef\b\b",
		"n" + string(rune(0)),
		"n" + string(rune(0)) + "n",
	}

	table := []struct {
		name        string
		contentTmpl string
		contentType models.BodyType
	}{
		{
			name:        "PlainText",
			contentTmpl: "%s",
			contentType: models.TEXT_BODYTYPE,
		},
		{
			name:        "HTML",
			contentTmpl: htmlStub,
			contentType: models.HTML_BODYTYPE,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			for _, charRange := range escapeCharRanges {
				for i := charRange[0]; i < charRange[1]; i++ {
					subject := fmt.Sprintf("character %x", i)

					bodyContent := fmt.Sprintf(test.contentTmpl, string(rune(i)))

					serialized := sendItemWithBodyAndGetSerialized(
						t,
						ctx,
						msgs,
						userID,
						ptr.Val(mailfolder.GetId()),
						subject,
						bodyContent,
						test.contentType)

					matches := contentRegex.FindAllSubmatch(serialized, -1)

					switch {
					case len(matches) == 0:
						t.Logf("character 0x%x wasn't found", i)

					case len(matches[0]) < 2:
						t.Logf("character 0x%x was removed", i)

					case bodyContent != string(matches[0][1]):
						t.Logf("character 0x%x has been transformed to %s", i, matches[0][1])
					}

					sanitized := sanitize.JSONBytes(serialized)
					newSerialized := sendSerializedItemAndGetSerialized(
						t,
						ctx,
						msgs,
						userID,
						ptr.Val(mailfolder.GetId()),
						sanitized)

					newMatches := contentRegex.FindAllSubmatch(newSerialized, -1)

					switch {
					case len(newMatches) == 0:
						t.Logf("sanitized character 0x%x wasn't found", i)

					case len(newMatches[0]) < 2:
						t.Logf("sanitized character 0x%x was removed", i)

					case bodyContent != string(newMatches[0][1]):
						t.Logf(
							"sanitized character 0x%x has been transformed to %s",
							i,
							newMatches[0][1])
					}

					assert.Equal(t, matches[0][1], newMatches[0][1])
				}
			}

			for i, sequence := range testSequences {
				subject := fmt.Sprintf("sequence %d", i)

				bodyContent := fmt.Sprintf(test.contentTmpl, sequence)

				serialized := sendItemWithBodyAndGetSerialized(
					t,
					ctx,
					msgs,
					userID,
					ptr.Val(mailfolder.GetId()),
					subject,
					bodyContent,
					test.contentType)

				matches := contentRegex.FindAllSubmatch(serialized, -1)

				switch {
				case len(matches) == 0:
					t.Logf("sequence %d wasn't found", i)

				case len(matches[0]) < 2:
					t.Logf("sequence %d was removed", i)

				case sequence != string(matches[0][1]):
					t.Logf("sequence %d has been transformed to %s", i, matches[0][1])
				}

				sanitized := sanitize.JSONBytes(serialized)
				newSerialized := sendSerializedItemAndGetSerialized(
					t,
					ctx,
					msgs,
					userID,
					ptr.Val(mailfolder.GetId()),
					sanitized)

				newMatches := contentRegex.FindAllSubmatch(newSerialized, -1)

				switch {
				case len(newMatches) == 0:
					t.Logf("sanitized sequence %d wasn't found", i)

				case len(newMatches[0]) < 2:
					t.Logf("sanitized sequence %d was removed", i)

				case sequence != string(newMatches[0][1]):
					t.Logf(
						"sanitized sequence %d has been transformed to %s",
						i,
						newMatches[0][1])
				}

				assert.Equal(t, matches[0][1], newMatches[0][1])
			}
		})
	}
}
