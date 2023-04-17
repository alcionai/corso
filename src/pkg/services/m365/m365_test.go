package m365_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365"
)

type M365IntegrationSuite struct {
	tester.Suite
}

func TestM365IntegrationSuite(t *testing.T) {
	suite.Run(t, &M365IntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *M365IntegrationSuite) TestUsers() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t    = suite.T()
		acct = tester.NewM365Account(suite.T())
	)

	users, err := m365.Users(ctx, acct, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, users)

	for _, u := range users {
		suite.Run("user_"+u.ID, func() {
			t := suite.T()

			assert.NotEmpty(t, u.ID)
			assert.NotEmpty(t, u.PrincipalName)
			assert.NotEmpty(t, u.Name)
			assert.NotEmpty(t, u.ArchiveFolder)
			assert.NotEmpty(t, u.DateFormat)
			assert.NotEmpty(t, u.Timezone)
			assert.NotEmpty(t, u.DelegateMeetingMsgDeliveryOpt)
			assert.NotEmpty(t, u.TimeFormat)
			assert.NotEmpty(t, u.UserPurpose)
			assert.NotEmpty(t, u.HasMailBox)
			assert.NotEmpty(t, u.HasOnedrive)
		})
	}
}

func (suite *M365IntegrationSuite) TestGetUserInfo() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t    = suite.T()
		acct = tester.NewM365Account(t)
		uid  = tester.M365UserID(t)
	)

	info, err := m365.GetUserInfo(ctx, acct, uid)
	require.NoError(t, err, clues.ToCore(err))
	require.NotNil(t, info)
	require.NotEmpty(t, info)

	expect := &m365.UserInfo{
		ServicesEnabled: m365.ServiceAccess{
			Exchange: true,
		},
	}

	assert.Equal(t, expect, info)
}

func (suite *M365IntegrationSuite) TestSites() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t    = suite.T()
		acct = tester.NewM365Account(suite.T())
	)

	sites, err := m365.Sites(ctx, acct, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotEmpty(t, sites)

	for _, s := range sites {
		suite.Run("site", func() {
			t := suite.T()
			assert.NotEmpty(t, s.WebURL)
			assert.NotEmpty(t, s.ID)
			assert.NotEmpty(t, s.DisplayName)
		})
	}
}
