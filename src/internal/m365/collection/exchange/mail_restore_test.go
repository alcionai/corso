package exchange

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/its"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

//nolint:lll
const TestDN = "/o=ExchangeLabs/ou=Exchange Administrative Group (FYDIBOHF23SPDLT)/cn=Recipients/cn=4eca0d46a2324036b0b326dc58cfc802-user"

type RestoreMailUnitSuite struct {
	tester.Suite
}

func TestRestoreMailUnitSuite(t *testing.T) {
	suite.Run(t, &RestoreMailUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RestoreMailUnitSuite) TestIsValidEmail() {
	table := []struct {
		name  string
		email string
		check assert.BoolAssertionFunc
	}{
		{
			name:  "valid email",
			email: "foo@bar.com",
			check: assert.True,
		},
		{
			name:  "invalid email, missing domain",
			email: "foo.com",
			check: assert.False,
		},
		{
			name:  "invalid email, random uuid",
			email: "12345678-abcd-90ef-88f8-2d95ef12fb66",
			check: assert.False,
		},
		{
			name:  "empty email",
			email: "",
			check: assert.False,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := isValidEmail(test.email)
			test.check(t, result)
		})
	}
}

func (suite *RestoreMailUnitSuite) TestIsValidDN() {
	table := []struct {
		name  string
		dn    string
		check assert.BoolAssertionFunc
	}{
		{
			name:  "valid DN",
			dn:    TestDN,
			check: assert.True,
		},
		{
			name:  "invalid DN",
			dn:    "random string",
			check: assert.False,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result := isValidDN(test.dn)
			test.check(t, result)
		})
	}
}

func (suite *RestoreMailUnitSuite) TestSetReplyTos() {
	t := suite.T()

	replyTos := make([]models.Recipientable, 0)

	emailAddresses := map[string]string{
		"foo.bar": "foo@bar.com",
		"foo.com": "foo.com",
		"empty":   "",
		"dn":      TestDN,
	}

	validEmailAddresses := map[string]string{
		"foo.bar": "foo@bar.com",
		"dn":      TestDN,
	}

	for k, v := range emailAddresses {
		emailAddress := models.NewEmailAddress()
		emailAddress.SetAddress(ptr.To(v))
		emailAddress.SetName(ptr.To(k))

		replyTo := models.NewRecipient()
		replyTo.SetEmailAddress(emailAddress)

		replyTos = append(replyTos, replyTo)
	}

	mailMessage := models.NewMessage()
	mailMessage.SetReplyTo(replyTos)

	setReplyTos(mailMessage)

	sanitizedReplyTos := mailMessage.GetReplyTo()
	require.Len(t, sanitizedReplyTos, len(validEmailAddresses))

	for _, sanitizedReplyTo := range sanitizedReplyTos {
		emailAddress := sanitizedReplyTo.GetEmailAddress()

		assert.Contains(t, validEmailAddresses, ptr.Val(emailAddress.GetName()))
		assert.Equal(t, validEmailAddresses[ptr.Val(emailAddress.GetName())], ptr.Val(emailAddress.GetAddress()))
	}
}

var _ mailRestorer = &mailRestoreMock{}

type mailRestoreMock struct {
	postItemErr       error
	calledPost        bool
	deleteItemErr     error
	calledDelete      bool
	postAttachmentErr error
}

func (m *mailRestoreMock) PostItem(
	_ context.Context,
	_, _ string,
	_ models.Messageable,
) (models.Messageable, error) {
	m.calledPost = true
	return models.NewMessage(), m.postItemErr
}

func (m *mailRestoreMock) DeleteItem(
	_ context.Context,
	_, _ string,
) error {
	m.calledDelete = true
	return m.deleteItemErr
}

func (m *mailRestoreMock) PostSmallAttachment(
	_ context.Context,
	_, _, _ string,
	_ models.Attachmentable,
) error {
	return m.postAttachmentErr
}

func (m *mailRestoreMock) PostLargeAttachment(
	_ context.Context,
	_, _, _, _ string,
	_ []byte,
) (string, error) {
	return uuid.NewString(), m.postAttachmentErr
}

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type MailRestoreIntgSuite struct {
	tester.Suite
	m365 its.M365IntgTestSetup
}

func TestMailRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &MailRestoreIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *MailRestoreIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *MailRestoreIntgSuite) TestCreateContainerDestination() {
	runCreateDestinationTest(
		suite.T(),
		newMailRestoreHandler(suite.m365.AC),
		path.EmailCategory,
		suite.m365.TenantID,
		suite.m365.User.ID,
		testdata.DefaultRestoreConfig("").Location,
		[]string{"Griffindor", "Croix"},
		[]string{"Griffindor", "Felicius"})
}

func (suite *MailRestoreIntgSuite) TestRestoreMail() {
	body := mock.MessageBytes("subject")

	stub, err := api.BytesToMessageable(body)
	require.NoError(suite.T(), err, clues.ToCore(err))

	collisionKey := api.MailCollisionKey(stub)

	type counts struct {
		skip    int64
		replace int64
		new     int64
	}

	table := []struct {
		name         string
		apiMock      *mailRestoreMock
		collisionMap map[string]string
		onCollision  control.CollisionPolicy
		expectErr    func(*testing.T, error)
		expectMock   func(*testing.T, *mailRestoreMock)
		expectCounts counts
	}{
		{
			name:         "no collision: skip",
			apiMock:      &mailRestoreMock{},
			collisionMap: map[string]string{},
			onCollision:  control.Copy,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.True(t, m.calledPost, "new item posted")
				assert.False(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name:         "no collision: copy",
			apiMock:      &mailRestoreMock{},
			collisionMap: map[string]string{},
			onCollision:  control.Skip,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.True(t, m.calledPost, "new item posted")
				assert.False(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name:         "no collision: replace",
			apiMock:      &mailRestoreMock{},
			collisionMap: map[string]string{},
			onCollision:  control.Replace,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.True(t, m.calledPost, "new item posted")
				assert.False(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name:         "collision: skip",
			apiMock:      &mailRestoreMock{},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Skip,
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, core.ErrAlreadyExists, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.False(t, m.calledPost, "new item posted")
				assert.False(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{1, 0, 0},
		},
		{
			name:         "collision: copy",
			apiMock:      &mailRestoreMock{},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Copy,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.True(t, m.calledPost, "new item posted")
				assert.False(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{0, 0, 1},
		},
		{
			name:         "collision: replace",
			apiMock:      &mailRestoreMock{},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Replace,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.True(t, m.calledPost, "new item posted")
				assert.True(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{0, 1, 0},
		},
		{
			name:         "collision: replace - err already deleted",
			apiMock:      &mailRestoreMock{deleteItemErr: core.ErrNotFound},
			collisionMap: map[string]string{collisionKey: "smarf"},
			onCollision:  control.Replace,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expectMock: func(t *testing.T, m *mailRestoreMock) {
				assert.True(t, m.calledPost, "new item posted")
				assert.True(t, m.calledDelete, "old item deleted")
			},
			expectCounts: counts{0, 1, 0},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctr := count.New()

			_, err := restoreMail(
				ctx,
				test.apiMock,
				body,
				suite.m365.User.ID,
				"destination",
				test.collisionMap,
				test.onCollision,
				fault.New(true),
				ctr)

			test.expectErr(t, err)
			test.expectMock(t, test.apiMock)
			assert.Equal(t, test.expectCounts.skip, ctr.Get(count.CollisionSkip), "skips")
			assert.Equal(t, test.expectCounts.replace, ctr.Get(count.CollisionReplace), "replaces")
			assert.Equal(t, test.expectCounts.new, ctr.Get(count.NewItemCreated), "new items")
		})
	}
}
