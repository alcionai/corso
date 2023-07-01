package m365

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/google/uuid"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/credentials"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type M365IntegrationSuite struct {
	tester.Suite
}

func TestM365IntegrationSuite(t *testing.T) {
	suite.Run(t, &M365IntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *M365IntegrationSuite) SetupSuite() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)
}

func (suite *M365IntegrationSuite) TestUsers() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	acct := tester.NewM365Account(suite.T())

	users, err := Users(ctx, acct, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, users)

	for _, u := range users {
		suite.Run("user_"+u.ID, func() {
			t := suite.T()

			assert.NotEmpty(t, u.ID)
			assert.NotEmpty(t, u.PrincipalName)
			assert.NotEmpty(t, u.Name)
			assert.NotEmpty(t, u.Info)
		})
	}
}

func (suite *M365IntegrationSuite) TestUsersCompat_HasNoInfo() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tester.NewM365Account(suite.T())

	users, err := UsersCompatNoInfo(ctx, acct)
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, users)

	for _, u := range users {
		suite.Run("user_"+u.ID, func() {
			t := suite.T()

			assert.NotEmpty(t, u.ID)
			assert.NotEmpty(t, u.PrincipalName)
			assert.NotEmpty(t, u.Name)
		})
	}
}

func (suite *M365IntegrationSuite) TestUserHasMailbox() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acct = tester.NewM365Account(t)
		uid  = tester.M365UserID(t)
	)

	enabled, err := UserHasMailbox(ctx, acct, uid)
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, enabled)
}

func (suite *M365IntegrationSuite) TestUserHasDrive() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acct = tester.NewM365Account(t)
		uid  = tester.M365UserID(t)
	)

	enabled, err := UserHasDrives(ctx, acct, uid)
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, enabled)
}

func (suite *M365IntegrationSuite) TestSites() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	acct := tester.NewM365Account(t)

	sites, err := Sites(ctx, acct, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, sites)

	for _, s := range sites {
		suite.Run("site_"+s.ID, func() {
			t := suite.T()
			assert.NotEmpty(t, s.WebURL)
			assert.NotEmpty(t, s.ID)
			assert.NotEmpty(t, s.DisplayName)
		})
	}
}

type m365UnitSuite struct {
	tester.Suite
}

func TestM365UnitSuite(t *testing.T) {
	suite.Run(t, &m365UnitSuite{Suite: tester.NewUnitSuite(t)})
}

type mockDGDD struct {
	response models.Driveable
	err      error
}

func (m mockDGDD) GetDefaultDrive(context.Context, string) (models.Driveable, error) {
	return m.response, m.err
}

func (suite *m365UnitSuite) TestCheckUserHasDrives() {
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
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To(string(graph.MysiteNotFound)))
				odErr.SetError(merr)

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
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To(string(graph.MysiteURLNotFound)))
				odErr.SetError(merr)

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
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To(string(graph.NoSPLicense)))
				odErr.SetError(merr)

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
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To(string(graph.RequestResourceNotFound)))
				merr.SetMessage(ptr.To("message"))
				odErr.SetError(merr)

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
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To("message"))
				odErr.SetError(merr)

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

			ok, err := checkUserHasDrives(ctx, dgdd, "foo")
			test.expect(t, ok, "has drives flag")
			test.expectErr(t, err)
		})
	}
}

type mockGAS struct {
	response []models.Siteable
	err      error
}

func (m mockGAS) GetAll(context.Context, *fault.Bus) ([]models.Siteable, error) {
	return m.response, m.err
}

func (suite *m365UnitSuite) TestGetAllSites() {
	table := []struct {
		name      string
		mock      func(context.Context) getAllSiteser
		expectErr func(*testing.T, error)
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) getAllSiteser {
				return mockGAS{[]models.Siteable{}, nil}
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "no sharepoint license",
			mock: func(ctx context.Context) getAllSiteser {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To(string(graph.NoSPLicense)))
				odErr.SetError(merr)

				return mockGAS{nil, graph.Stack(ctx, odErr)}
			},
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrServiceNotEnabled, clues.ToCore(err))
			},
		},
		{
			name: "arbitrary error",
			mock: func(ctx context.Context) getAllSiteser {
				odErr := odataerrors.NewODataError()
				merr := odataerrors.NewMainError()
				merr.SetCode(ptr.To("code"))
				merr.SetMessage(ptr.To("message"))
				odErr.SetError(merr)

				return mockGAS{nil, graph.Stack(ctx, odErr)}
			},
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

			gas := test.mock(ctx)

			_, err := getAllSites(ctx, gas)
			test.expectErr(t, err)
		})
	}
}

type DiscoveryIntgSuite struct {
	tester.Suite
	acct account.Account
}

func TestDiscoveryIntgSuite(t *testing.T) {
	suite.Run(t, &DiscoveryIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs}),
	})
}

func (suite *DiscoveryIntgSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	suite.acct = tester.NewM365Account(t)
}

func (suite *DiscoveryIntgSuite) TestUsers() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	errs := fault.New(true)

	users, err := Users(ctx, suite.acct, errs)
	assert.NoError(t, err, clues.ToCore(err))

	ferrs := errs.Errors()
	assert.Nil(t, ferrs.Failure)
	assert.Empty(t, ferrs.Recovered)
	assert.NotEmpty(t, users)
}

func (suite *DiscoveryIntgSuite) TestUsers_InvalidCredentials() {
	table := []struct {
		name string
		acct func(t *testing.T) account.Account
	}{
		{
			name: "Invalid Credentials",
			acct: func(t *testing.T) account.Account {
				a, err := account.NewAccount(
					account.ProviderM365,
					account.M365Config{
						M365: credentials.M365{
							AzureClientID:     "Test",
							AzureClientSecret: "without",
						},
						AzureTenantID: "data",
					},
				)
				require.NoError(t, err, clues.ToCore(err))

				return a
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			users, err := Users(ctx, test.acct(t), fault.New(true))
			assert.Empty(t, users, "returned some users")
			assert.NotNil(t, err)
		})
	}
}

func (suite *DiscoveryIntgSuite) TestSites_InvalidCredentials() {
	table := []struct {
		name string
		acct func(t *testing.T) account.Account
	}{
		{
			name: "Invalid Credentials",
			acct: func(t *testing.T) account.Account {
				a, err := account.NewAccount(
					account.ProviderM365,
					account.M365Config{
						M365: credentials.M365{
							AzureClientID:     "Test",
							AzureClientSecret: "without",
						},
						AzureTenantID: "data",
					},
				)
				require.NoError(t, err, clues.ToCore(err))

				return a
			},
		},
		{
			name: "Empty Credentials",
			acct: func(t *testing.T) account.Account {
				// intentionally swallowing the error here
				a, _ := account.NewAccount(account.ProviderM365)
				return a
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sites, err := Sites(ctx, test.acct(t), fault.New(true))
			assert.Empty(t, sites, "returned some sites")
			assert.NotNil(t, err)
		})
	}
}

func (suite *DiscoveryIntgSuite) TestGetUserInfo() {
	table := []struct {
		name      string
		user      string
		expect    *api.UserInfo
		expectErr require.ErrorAssertionFunc
	}{
		{
			name: "standard test user",
			user: tester.M365UserID(suite.T()),
			expect: &api.UserInfo{
				ServicesEnabled: map[path.ServiceType]struct{}{
					path.ExchangeService: {},
					path.OneDriveService: {},
				},
				Mailbox: api.MailboxInfo{
					Purpose:              "user",
					ErrGetMailBoxSetting: nil,
				},
			},
			expectErr: require.NoError,
		},
		{
			name: "user does not exist",
			user: uuid.NewString(),
			expect: &api.UserInfo{
				ServicesEnabled: map[path.ServiceType]struct{}{},
				Mailbox:         api.MailboxInfo{},
			},
			expectErr: require.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			result, err := GetUserInfo(ctx, suite.acct, test.user)
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, test.expect.ServicesEnabled, result.ServicesEnabled)
		})
	}
}

func (suite *DiscoveryIntgSuite) TestGetUserInfo_userWithoutDrive() {
	userID := tester.M365UserID(suite.T())

	table := []struct {
		name   string
		user   string
		expect *api.UserInfo
	}{
		{
			name: "user without drive and exchange",
			user: "a53c26f7-5100-4acb-a910-4d20960b2c19", // User: testevents@10rqc2.onmicrosoft.com
			expect: &api.UserInfo{
				ServicesEnabled: map[path.ServiceType]struct{}{},
				Mailbox: api.MailboxInfo{
					ErrGetMailBoxSetting: []error{api.ErrMailBoxSettingsNotFound},
				},
			},
		},
		{
			name: "user with drive and exchange",
			user: userID,
			expect: &api.UserInfo{
				ServicesEnabled: map[path.ServiceType]struct{}{
					path.ExchangeService: {},
					path.OneDriveService: {},
				},
				Mailbox: api.MailboxInfo{
					Purpose:              "user",
					ErrGetMailBoxSetting: []error{},
				},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			result, err := GetUserInfo(ctx, suite.acct, test.user)
			require.NoError(t, err, clues.ToCore(err))
			assert.Equal(t, test.expect.ServicesEnabled, result.ServicesEnabled)
			assert.Equal(t, test.expect.Mailbox.ErrGetMailBoxSetting, result.Mailbox.ErrGetMailBoxSetting)
			assert.Equal(t, test.expect.Mailbox.Purpose, result.Mailbox.Purpose)
		})
	}
}
