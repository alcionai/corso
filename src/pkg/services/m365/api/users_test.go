package api

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
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

			err := validateUser(tt.args)
			tt.errCheck(t, err, clues.ToCore(err))
		})
	}
}
