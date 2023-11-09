package eml

import (
	"fmt"
	"testing"
	"time"

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

func (suite *EMLUnitSuite) TestConvert_messageble_to_eml() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	body := []byte(testdata.EmailWithAttachments)

	out, err := FromJSON(ctx, body)
	assert.NoError(t, err, "converting to eml")

	msg, err := api.BytesToMessageable(body)
	require.NoError(t, err, "creating message")

	assert.Contains(t, out, fmt.Sprintf("Subject: %s", ptr.Val(msg.GetSubject())))
	assert.Contains(t, out, fmt.Sprintf("Date: %s", msg.GetSentDateTime().Format(time.RFC1123Z)))
	assert.Contains(
		t,
		out,
		fmt.Sprintf(
			`From: "%s" <%s>`,
			ptr.Val(msg.GetFrom().GetEmailAddress().GetName()),
			ptr.Val(msg.GetFrom().GetEmailAddress().GetAddress())))

	for _, addr := range msg.GetToRecipients() {
		assert.Contains(
			t,
			out,
			fmt.Sprintf(
				`To: "%s" <%s>`,
				ptr.Val(addr.GetEmailAddress().GetName()),
				ptr.Val(addr.GetEmailAddress().GetAddress())))
	}

	for _, addr := range msg.GetCcRecipients() {
		assert.Contains(
			t,
			out,
			fmt.Sprintf(
				`Cc: "%s" <%s>`,
				ptr.Val(addr.GetEmailAddress().GetName()),
				ptr.Val(addr.GetEmailAddress().GetAddress())))
	}

	for _, addr := range msg.GetBccRecipients() {
		assert.Contains(
			t,
			out,
			fmt.Sprintf(
				`Bcc: "%s" <%s>`,
				ptr.Val(addr.GetEmailAddress().GetName()),
				ptr.Val(addr.GetEmailAddress().GetAddress())))
	}

	// Only fist 30 chars as the .eml generator can introduce a
	// newline in between the text to limit the column width of the
	// output. It does not affect the data, but can break our tests and
	// so using 30 as a safe limit to test.
	assert.Contains(t, out, ptr.Val(msg.GetBody().GetContent())[:30], "body")
}
