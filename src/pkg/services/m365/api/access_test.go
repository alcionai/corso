package api_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type AccessAPIIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestAccessAPIIntgSuite(t *testing.T) {
	suite.Run(t, &AccessAPIIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *AccessAPIIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *AccessAPIIntgSuite) TestGetToken() {

	tests := []struct {
		name      string
		creds     func() account.M365Config
		expectErr require.ErrorAssertionFunc
	}{
		{
			name:      "good",
			creds:     func() account.M365Config { return suite.its.ac.Credentials },
			expectErr: require.NoError,
		},
		{
			name: "bad tenant ID",
			creds: func() account.M365Config {
				creds := suite.its.ac.Credentials
				creds.AzureTenantID = "ZIM"

				return creds
			},
			expectErr: require.Error,
		},
		{
			name: "missing tenant ID",
			creds: func() account.M365Config {
				creds := suite.its.ac.Credentials
				creds.AzureTenantID = ""

				return creds
			},
			expectErr: require.Error,
		},
		{
			name: "bad client ID",
			creds: func() account.M365Config {
				creds := suite.its.ac.Credentials
				creds.AzureClientID = "GIR"

				return creds
			},
			expectErr: require.Error,
		},
		{
			name: "missing client ID",
			creds: func() account.M365Config {
				creds := suite.its.ac.Credentials
				creds.AzureClientID = ""

				return creds
			},
			expectErr: require.Error,
		},
		{
			name: "bad client secret",
			creds: func() account.M365Config {
				creds := suite.its.ac.Credentials
				creds.AzureClientSecret = "MY TALLEST"

				return creds
			},
			expectErr: require.Error,
		},
		{
			name: "missing client secret",
			creds: func() account.M365Config {
				creds := suite.its.ac.Credentials
				creds.AzureClientSecret = ""

				return creds
			},
			expectErr: require.Error,
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ac, err := api.NewClient(suite.its.ac.Credentials, control.DefaultOptions())
			require.NoError(t, err, clues.ToCore(err))

			ac.Credentials = test.creds()

			err = ac.Access().GetToken(ctx)
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}
