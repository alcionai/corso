package api_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/h2non/gock"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type UsersUnitSuite struct {
	tester.Suite
}

func TestUsersUnitSuite(t *testing.T) {
	suite.Run(t, &UsersUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *UsersUnitSuite) TestValidateUser() {
	name := "testuser"
	email := "testuser@foo.com"
	id := "testID"
	user := models.NewUser()
	user.SetUserPrincipalName(&email)
	user.SetDisplayName(&name)
	user.SetId(&id)

	tests := []struct {
		name     string
		args     models.Userable
		errCheck assert.ErrorAssertionFunc
	}{
		{
			name:     "No ID",
			args:     models.NewUser(),
			errCheck: assert.Error,
		},
		{
			name: "No user principal name",
			args: func() *models.User {
				u := models.NewUser()
				u.SetId(&id)
				return u
			}(),
			errCheck: assert.Error,
		},
		{
			name:     "Valid User",
			args:     user,
			errCheck: assert.NoError,
		},
	}
	for _, tt := range tests {
		suite.Run(tt.name, func() {
			t := suite.T()

			err := api.ValidateUser(tt.args)
			tt.errCheck(t, err, clues.ToCore(err))
		})
	}
}

func (suite *UsersUnitSuite) TestEvaluateMailboxError() {
	table := []struct {
		name   string
		err    error
		expect func(t *testing.T, err error)
	}{
		{
			name: "nil",
			err:  nil,
			expect: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "mail inbox err - user not found",
			err:  odErr(string(graph.RequestResourceNotFound)),
			expect: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrResourceOwnerNotFound, clues.ToCore(err))
			},
		},
		{
			name: "mail inbox err - user not found",
			err:  odErr(string(graph.MailboxNotEnabledForRESTAPI)),
			expect: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "mail inbox err - authenticationError",
			err:  odErr(string(graph.AuthenticationError)),
			expect: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "mail inbox err - other error",
			err:  odErrMsg("somecode", "somemessage"),
			expect: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), api.EvaluateMailboxError(test.err))
		})
	}
}

type UsersIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestUsersIntgSuite(t *testing.T) {
	suite.Run(t, &UsersIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *UsersIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *UsersIntgSuite) TestUsers_GetInfo_errors() {
	table := []struct {
		name      string
		setGocks  func(t *testing.T)
		expectErr func(t *testing.T, err error)
	}{
		{
			name: "default drive err - mysite not found",
			setGocks: func(t *testing.T) {
				interceptV1Path("users", "user", "drive").
					Reply(400).
					JSON(parseableToMap(t, odErrMsg("anycode", string(graph.MysiteNotFound))))
				interceptV1Path("users", "user", "mailFolders", api.MailInbox).
					Reply(400).
					JSON(parseableToMap(t, odErr(string(graph.ResourceNotFound))))
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "default drive err - no sharepoint license",
			setGocks: func(t *testing.T) {
				interceptV1Path("users", "user", "drive").
					Reply(400).
					JSON(parseableToMap(t, odErrMsg("anycode", string(graph.NoSPLicense))))
				interceptV1Path("users", "user", "mailFolders", api.MailInbox).
					Reply(400).
					JSON(parseableToMap(t, odErr(string(graph.ResourceNotFound))))
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "default drive err - other error",
			setGocks: func(t *testing.T) {
				interceptV1Path("users", "user", "drive").
					Reply(400).
					JSON(parseableToMap(t, odErrMsg("somecode", "somemessage")))
			},
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err, clues.ToCore(err))
			},
		},
		{
			name: "mail inbox err - user not found",
			setGocks: func(t *testing.T) {
				interceptV1Path("users", "user", "drive").
					Reply(200).
					JSON(parseableToMap(t, models.NewDrive()))
				interceptV1Path("users", "user", "mailFolders", api.MailInbox).
					Reply(400).
					JSON(parseableToMap(t, odErr(string(graph.RequestResourceNotFound))))
			},
			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, graph.ErrResourceOwnerNotFound, clues.ToCore(err))
			},
		},
		{
			name: "mail inbox err - user not found",
			setGocks: func(t *testing.T) {
				interceptV1Path("users", "user", "drive").
					Reply(200).
					JSON(parseableToMap(t, models.NewDrive()))
				interceptV1Path("users", "user", "mailFolders", api.MailInbox).
					Reply(400).
					JSON(parseableToMap(t, odErr(string(graph.MailboxNotEnabledForRESTAPI))))
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "mail inbox err - authenticationError",
			setGocks: func(t *testing.T) {
				interceptV1Path("users", "user", "drive").
					Reply(200).
					JSON(parseableToMap(t, models.NewDrive()))
				interceptV1Path("users", "user", "mailFolders", api.MailInbox).
					Reply(400).
					JSON(parseableToMap(t, odErr(string(graph.AuthenticationError))))
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err, clues.ToCore(err))
			},
		},
		{
			name: "mail inbox err - other error",
			setGocks: func(t *testing.T) {
				interceptV1Path("users", "user", "drive").
					Reply(200).
					JSON(parseableToMap(t, models.NewDrive()))
				interceptV1Path("users", "user", "mailFolders", api.MailInbox).
					Reply(400).
					JSON(parseableToMap(t, odErrMsg("somecode", "somemessage")))
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
			defer gock.Off()

			test.setGocks(t)

			_, err := suite.its.gockAC.Users().GetInfo(ctx, "user")
			test.expectErr(t, err)
		})
	}
}

func (suite *UsersIntgSuite) TestUsers_GetInfo_QuotaExceeded() {
	t := suite.T()
	ctx, flush := tester.NewContext(t)

	defer flush()
	defer gock.Off()

	gock.EnableNetworking()
	gock.New(graphAPIHostURL).
		// Wildcard match on the inbox folder ID.
		Get(v1APIURLPath("users", suite.its.userID, "mailFolders", "(.*)", "messages", "delta")).
		Reply(403).
		SetHeaders(
			map[string]string{
				"Content-Type": "application/json; odata.metadata=minimal; " +
					"odata.streaming=true; IEEE754Compatible=false; charset=utf-8",
			},
		).
		BodyString(`{"error":{"code":"ErrorQuotaExceeded","message":"The process failed to get the correct properties."}}`)

	output, err := suite.its.gockAC.Users().GetInfo(ctx, suite.its.userID)
	require.NoError(t, err, clues.ToCore(err))

	assert.True(t, output.Mailbox.QuotaExceeded)
}
