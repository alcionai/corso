package m365

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/mock"
)

type CommonM365UnitSuite struct {
	tester.Suite
}

func TestM365UnitSuite(t *testing.T) {
	suite.Run(t, &CommonM365UnitSuite{Suite: tester.NewUnitSuite(t)})
}

var _ getDefaultDriver = mockDGDD{}

type mockDGDD struct {
	response models.Driveable
	err      error
}

func (m mockDGDD) GetDefaultDrive(context.Context, string) (models.Driveable, error) {
	return m.response, m.err
}

// Copied from src/internal/m365/graph/errors_test.go
func odErrMsg(code, message string) *odataerrors.ODataError {
	odErr := odataerrors.NewODataError()
	merr := odataerrors.NewMainError()
	merr.SetCode(&code)
	merr.SetMessage(&message)
	odErr.SetErrorEscaped(merr)

	return odErr
}

func (suite *CommonM365UnitSuite) TestIsOneDriveServiceEnabled() {
	table := []struct {
		name      string
		mock      func(context.Context) getDefaultDriver
		expect    assert.BoolAssertionFunc
		expectErr func(*testing.T, error)
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) getDefaultDriver {
				return mockDGDD{models.NewDrive(), nil}
			},
			expect: assert.True,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "mysite not found",
			mock: func(ctx context.Context) getDefaultDriver {
				odErr := odErrMsg("code", string(graph.MysiteNotFound))
				return mockDGDD{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "mysite URL not found",
			mock: func(ctx context.Context) getDefaultDriver {
				odErr := odErrMsg("code", string(graph.MysiteURLNotFound))
				return mockDGDD{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "no sharepoint license",
			mock: func(ctx context.Context) getDefaultDriver {
				odErr := odErrMsg("code", string(graph.NoSPLicense))
				return mockDGDD{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "user not found",
			mock: func(ctx context.Context) getDefaultDriver {
				odErr := odErrMsg(string(graph.RequestResourceNotFound), "message")
				return mockDGDD{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
		{
			name: "arbitrary error",
			mock: func(ctx context.Context) getDefaultDriver {
				odErr := odErrMsg("code", "message")
				return mockDGDD{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			dgdd := test.mock(ctx)

			ok, err := IsOneDriveServiceEnabled(ctx, dgdd, "resource_id")
			test.expect(t, ok, "has drives flag")
			test.expectErr(t, err)
		})
	}
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

func (suite *CommonM365UnitSuite) TestIsExchangeServiceEnabled() {
	table := []struct {
		name      string
		mock      func(context.Context) getMailInboxer
		expect    assert.BoolAssertionFunc
		expectErr func(*testing.T, error)
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) getMailInboxer {
				return mockGMB{
					mailbox: models.NewMailFolder(),
				}
			},
			expect: assert.True,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "user has no mailbox",
			mock: func(ctx context.Context) getMailInboxer {
				odErr := odErrMsg(string(graph.ResourceNotFound), "message")

				return mockGMB{
					mailboxErr: graph.Stack(ctx, odErr),
				}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "user not found",
			mock: func(ctx context.Context) getMailInboxer {
				odErr := odErrMsg(string(graph.RequestResourceNotFound), "message")

				return mockGMB{
					mailboxErr: graph.Stack(ctx, odErr),
				}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
		{
			name: "overlapping resourcenotfound",
			mock: func(ctx context.Context) getMailInboxer {
				odErr := odErrMsg(string(graph.ResourceNotFound), "User not found")

				return mockGMB{
					mailboxErr: graph.Stack(ctx, odErr),
				}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
		{
			name: "arbitrary error",
			mock: func(ctx context.Context) getMailInboxer {
				odErr := odErrMsg("code", "message")

				return mockGMB{
					mailboxErr: graph.Stack(ctx, odErr),
				}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			gmi := test.mock(ctx)

			ok, err := IsExchangeServiceEnabled(ctx, gmi, "resource_id")
			test.expect(t, ok, "has mailbox flag")
			test.expectErr(t, err)
		})
	}
}

var _ getSiteRooter = mockGSR{}

type mockGSR struct {
	response models.Siteable
	err      error
}

func (m mockGSR) GetRoot(context.Context) (models.Siteable, error) {
	return m.response, m.err
}

func (suite *CommonM365UnitSuite) TestIsSharePointServiceEnabled() {
	table := []struct {
		name      string
		mock      func(context.Context) getSiteRooter
		expect    assert.BoolAssertionFunc
		expectErr func(*testing.T, error)
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) getSiteRooter {
				return mockGSR{models.NewSite(), nil}
			},
			expect: assert.True,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "no sharepoint license",
			mock: func(ctx context.Context) getSiteRooter {
				odErr := odErrMsg("code", string(graph.NoSPLicense))

				return mockGSR{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "arbitrary error",
			mock: func(ctx context.Context) getSiteRooter {
				odErr := odErrMsg("code", "message")

				return mockGSR{nil, graph.Stack(ctx, odErr)}
			},
			expect: assert.False,
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			gsr := test.mock(ctx)

			ok, err := IsSharePointServiceEnabled(ctx, gsr, "resource_id")
			test.expect(t, ok, "has sites flag")
			test.expectErr(t, err)
		})
	}
}

func (suite *CommonM365UnitSuite) TestGetMailboxInfo() {
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
				err := odErrMsg(string(graph.ResourceNotFound), "message")

				return mockGMB{
					mailboxErr: graph.Stack(ctx, err),
				}
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
			expect: func(t *testing.T) api.MailboxInfo {
				mi := api.MailboxInfo{}
				mi.ErrGetMailBoxSetting = append(
					mi.ErrGetMailBoxSetting,
					api.ErrMailBoxSettingsNotFound)

				return mi
			},
		},
		{
			name: "settings access denied",
			mock: func(ctx context.Context) getMailboxer {
				err := odErrMsg(string(graph.ErrorAccessDenied), "message")

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
				err := odErrMsg(string(graph.QuotaExceeded), "message")
				return mockGMB{
					mailbox:         models.NewMailFolder(),
					settings:        mock.UserSettings(),
					inboxMessageErr: graph.Stack(ctx, err),
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
