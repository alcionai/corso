package api

import (
	"reflect"
	"testing"

	"github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/stretchr/testify/suite"
)

type UsersUnitSuite struct {
	suite.Suite
}

func TestUsersUnitSuite(t *testing.T) {
	suite.Run(t, new(UsersUnitSuite))
}

func (suite *UsersUnitSuite) TestValidateUser() {
	t := suite.T()

	name := "testuser"
	email := "testuser@foo.com"
	id := "testID"
	user := models.NewUser()
	user.SetUserPrincipalName(&email)
	user.SetDisplayName(&name)
	user.SetId(&id)

	tests := []struct {
		name    string
		args    interface{}
		want    models.Userable
		wantErr bool
	}{
		{
			name:    "Invalid type",
			args:    string("invalid type"),
			wantErr: true,
		},
		{
			name:    "No ID",
			args:    models.NewUser(),
			wantErr: true,
		},
		{
			name: "No user principal name",
			args: func() *models.User {
				u := models.NewUser()
				u.SetId(&id)
				return u
			}(),
			wantErr: true,
		},
		{
			name: "Valid User",
			args: user,
			want: user,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := validateUser(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
