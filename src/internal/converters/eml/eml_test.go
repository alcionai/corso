package eml

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/jhillyerd/enmime"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/converters/eml/testdata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type EMLUnitSuite struct {
	tester.Suite
}

func TestEMLUnitSuite(t *testing.T) {
	suite.Run(t, &EMLUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *EMLUnitSuite) TestFormatAddress() {
	t := suite.T()

	tests := []struct {
		tname    string
		name     string
		email    string
		expected string
	}{
		{
			tname:    "different name and email",
			name:     "John Doe",
			email:    "johndoe@provider.com",
			expected: `"John Doe" <johndoe@provider.com>`,
		},
		{
			tname:    "same name and email",
			name:     "johndoe@provider.com",
			email:    "johndoe@provider.com",
			expected: "johndoe@provider.com",
		},
		{
			tname:    "only email",
			name:     "",
			email:    "johndoe@provider.com",
			expected: "johndoe@provider.com",
		},
		{
			tname:    "only name",
			name:     "john doe",
			email:    "",
			expected: `"john doe"`,
		},
		{
			tname:    "neither mail or name",
			name:     "",
			email:    "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.tname, func(t *testing.T) {
			entity := models.NewEmailAddress()
			if len(tt.name) != 0 {
				entity.SetName(ptr.To(tt.name))
			}
			if len(tt.email) != 0 {
				entity.SetAddress(ptr.To(tt.email))
			}

			assert.Equal(t, tt.expected, formatAddress(entity))
		})
	}
}

func (suite *EMLUnitSuite) TestConvert_messageble_to_eml() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	body := []byte(testdata.EmailWithAttachments)

	out, err := FromJSON(ctx, body)
	assert.NoError(t, err, "converting to eml")

	msg, err := api.BytesToMessageable(body)
	require.NoError(t, err, "creating message")

	eml, err := enmime.ReadEnvelope(strings.NewReader(out))
	require.NoError(t, err, "reading created eml")

	assert.Equal(t, ptr.Val(msg.GetSubject()), eml.GetHeader("Subject"))
	assert.Equal(t, msg.GetSentDateTime().Format(time.RFC1123Z), eml.GetHeader("Date"))

	assert.Equal(t, formatAddress(msg.GetFrom().GetEmailAddress()), eml.GetHeader("From"))

	ccs := strings.Split(eml.GetHeader("Cc"), ", ")
	for _, cc := range msg.GetCcRecipients() {
		assert.Contains(t, ccs, formatAddress(cc.GetEmailAddress()))
	}

	bccs := strings.Split(eml.GetHeader("Bcc"), ", ")
	for _, bcc := range msg.GetBccRecipients() {
		assert.Contains(t, bccs, formatAddress(bcc.GetEmailAddress()))
	}

	tos := strings.Split(eml.GetHeader("To"), ", ")
	for _, to := range msg.GetToRecipients() {
		assert.Contains(t, tos, formatAddress(to.GetEmailAddress()))
	}

	source := strings.ReplaceAll(eml.HTML, "\n", "")
	target := strings.ReplaceAll(ptr.Val(msg.GetBody().GetContent()), "\n", "")

	// replace the cid with a constant value to make the comparison
	re := regexp.MustCompile(`src="cid:[^"]*"`)
	source = re.ReplaceAllString(source, `src="cid:replaced"`)
	target = re.ReplaceAllString(target, `src="cid:replaced"`)

	assert.Equal(t, source, target)
}

func (suite *EMLUnitSuite) TestConvert_edge_cases() {
	tests := []struct {
		name      string
		transform func(models.Messageable)
	}{
		{
			name: "incorrect address",
			transform: func(msg models.Messageable) {
				msg.GetFrom().GetEmailAddress().SetAddress(ptr.To("invalid"))
			},
		},
		{
			name: "empty attachment",
			transform: func(msg models.Messageable) {
				attachments := msg.GetAttachments()
				err := attachments[0].GetBackingStore().Set("contentBytes", []uint8{})
				require.NoError(suite.T(), err, "setting attachment content")
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			body := []byte(testdata.EmailWithAttachments)

			msg, err := api.BytesToMessageable(body)
			require.NoError(t, err, "creating message")

			test.transform(msg)

			writer := kjson.NewJsonSerializationWriter()

			defer writer.Close()

			err = writer.WriteObjectValue("", msg)
			require.NoError(t, err, "serializing message")

			nbody, err := writer.GetSerializedContent()
			require.NoError(t, err, "getting serialized content")

			_, err = FromJSON(ctx, nbody)
			assert.NoError(t, err, "converting to eml")
		})
	}
}
