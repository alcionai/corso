package exchange

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/pkg/errs/core"
	"github.com/alcionai/canario/src/pkg/services/m365/api"
	"github.com/alcionai/canario/src/pkg/services/m365/api/graph"
	graphTD "github.com/alcionai/canario/src/pkg/services/m365/api/graph/testdata"
	"github.com/alcionai/canario/src/pkg/services/m365/api/mock"
)

type EnabledUnitSuite struct {
	tester.Suite
}

func TestEnabledUnitSuite(t *testing.T) {
	suite.Run(t, &EnabledUnitSuite{Suite: tester.NewUnitSuite(t)})
}

var _ getMailboxer = mockGMB{}

type mockGMB struct {
	mailbox         models.MailFolderable
	mailboxErr      error
	settings        models.Userable
	settingsErr     error
	inboxMessageErr error
}

func (m mockGMB) GetMailInbox(context.Context, string) (models.MailFolderable, error) {
	return m.mailbox, m.mailboxErr
}

func (m mockGMB) GetMailboxSettings(context.Context, string) (models.Userable, error) {
	return m.settings, m.settingsErr
}

func (m mockGMB) GetFirstInboxMessage(context.Context, string, string) error {
	return m.inboxMessageErr
}

func (suite *EnabledUnitSuite) TestIsServiceEnabled() {
	table := []struct {
		name      string
		mock      func(context.Context) getMailInboxer
		expect    assert.BoolAssertionFunc
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) getMailInboxer {
				return mockGMB{
					mailbox: models.NewMailFolder(),
				}
			},
			expect:    assert.True,
			expectErr: assert.NoError,
		},
		{
			name: "user has no mailbox",
			mock: func(ctx context.Context) getMailInboxer {
				odErr := graphTD.ODataErrWithMsg(string(graph.ResourceNotFound), "message")

				return mockGMB{
					mailboxErr: clues.Stack(odErr),
				}
			},
			expect:    assert.False,
			expectErr: assert.NoError,
		},
		{
			name: "user not found",
			mock: func(ctx context.Context) getMailInboxer {
				odErr := graphTD.ODataErrWithMsg(string(graph.RequestResourceNotFound), "message")

				return mockGMB{
					mailboxErr: clues.Stack(odErr),
				}
			},
			expect:    assert.False,
			expectErr: assert.Error,
		},
		{
			name: "overlapping resourcenotfound",
			mock: func(ctx context.Context) getMailInboxer {
				odErr := graphTD.ODataErrWithMsg(string(graph.ResourceNotFound), "User not found")
				err := clues.Stack(core.ErrNotFound, odErr)

				return mockGMB{
					mailboxErr: clues.StackWC(ctx, err),
				}
			},
			expect:    assert.False,
			expectErr: assert.NoError,
		},
		{
			name: "arbitrary error",
			mock: func(ctx context.Context) getMailInboxer {
				odErr := graphTD.ODataErrWithMsg("code", "message")

				return mockGMB{
					mailboxErr: clues.Stack(odErr),
				}
			},
			expect:    assert.False,
			expectErr: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			gmi := test.mock(ctx)

			ok, err := IsServiceEnabled(ctx, gmi, "resource_id")
			test.expect(t, ok, "has mailbox flag")
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}

func (suite *EnabledUnitSuite) TestGetMailboxInfo() {
	table := []struct {
		name      string
		mock      func(context.Context) getMailboxer
		expectErr func(*testing.T, error)
		expect    func(*testing.T) api.MailboxInfo
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) getMailboxer {
				return mockGMB{
					mailbox:  models.NewMailFolder(),
					settings: mock.UserSettings(),
				}
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expect: func(t *testing.T) api.MailboxInfo {
				return mock.UserMailboxInfo()
			},
		},
		{
			name: "user has no mailbox",
			mock: func(ctx context.Context) getMailboxer {
				err := graphTD.ODataErrWithMsg(string(graph.ResourceNotFound), "message")

				return mockGMB{
					mailboxErr: clues.StackWC(ctx, err),
				}
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expect: func(t *testing.T) api.MailboxInfo {
				mi := api.MailboxInfo{}
				mi.ErrGetMailBoxSetting = append(
					mi.ErrGetMailBoxSetting,
					api.ErrMailBoxNotFound)

				return mi
			},
		},
		{
			name: "settings access denied",
			mock: func(ctx context.Context) getMailboxer {
				err := graphTD.ODataErrWithMsg(string(graph.ErrorAccessDenied), "message")

				return mockGMB{
					mailbox:     models.NewMailFolder(),
					settingsErr: err,
				}
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expect: func(t *testing.T) api.MailboxInfo {
				mi := api.MailboxInfo{}
				mi.ErrGetMailBoxSetting = append(
					mi.ErrGetMailBoxSetting,
					api.ErrMailBoxSettingsAccessDenied)

				return mi
			},
		},
		{
			name: "error fetching settings",
			mock: func(ctx context.Context) getMailboxer {
				return mockGMB{
					mailbox:     models.NewMailFolder(),
					settingsErr: assert.AnError,
				}
			},
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
			expect: func(t *testing.T) api.MailboxInfo {
				return api.MailboxInfo{
					ErrGetMailBoxSetting: []error{},
				}
			},
		},
		{
			name: "mailbox quota exceeded",
			mock: func(ctx context.Context) getMailboxer {
				err := graphTD.ODataErrWithMsg(string(graph.QuotaExceeded), "message")
				return mockGMB{
					mailbox:         models.NewMailFolder(),
					settings:        mock.UserSettings(),
					inboxMessageErr: clues.StackWC(ctx, err),
				}
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expect: func(t *testing.T) api.MailboxInfo {
				mi := mock.UserMailboxInfo()
				mi.QuotaExceeded = true

				return mi
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			gmi := test.mock(ctx)

			mi, err := GetMailboxInfo(ctx, gmi, "resource_id")
			test.expectErr(t, err)
			assert.Equal(t, test.expect(t), mi)
		})
	}
}
