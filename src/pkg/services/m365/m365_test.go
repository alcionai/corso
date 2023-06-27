package m365

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/discovery"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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

func (suite *M365IntegrationSuite) TestUsers() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

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

func (suite *M365IntegrationSuite) TestGetUserInfo() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		acct = tester.NewM365Account(t)
		uid  = tester.M365UserID(t)
	)

	info, err := GetUserInfo(ctx, acct, uid)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, info)
	require.NotEmpty(t, info)

	expectEnabled := map[path.ServiceType]struct{}{
		path.ExchangeService: {},
		path.OneDriveService: {},
	}

	assert.NotEmpty(t, info.ServicesEnabled)
	assert.NotEmpty(t, info.Mailbox)
	assert.Equal(t, expectEnabled, info.ServicesEnabled)
	assert.Equal(t, "user", info.Mailbox.Purpose)
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
		mock      func(context.Context) discovery.GetDefaultDriver
		expect    assert.BoolAssertionFunc
		expectErr func(*testing.T, error)
	}{
		{
			name: "ok",
			mock: func(ctx context.Context) discovery.GetDefaultDriver {
				return mockDGDD{models.NewDrive(), nil}
			},
			expect: assert.True,
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "mysite not found",
			mock: func(ctx context.Context) discovery.GetDefaultDriver {
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
			mock: func(ctx context.Context) discovery.GetDefaultDriver {
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
			mock: func(ctx context.Context) discovery.GetDefaultDriver {
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
			mock: func(ctx context.Context) discovery.GetDefaultDriver {
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
			mock: func(ctx context.Context) discovery.GetDefaultDriver {
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
