package m365

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/its"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
)

type M365IntgSuite struct {
	tester.Suite
	m365 its.M365IntgTestSetup
}

func TestM365IntgSuite(t *testing.T) {
	suite.Run(t, &M365IntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{}),
	})
}

func (suite *M365IntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *M365IntgSuite) TestNewM365Client() {
	table := []struct {
		name      string
		acct      func(t *testing.T) account.Account
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name: "Valid Credentials",
			acct: func(t *testing.T) account.Account {
				return suite.m365.Acct
			},
			expectErr: assert.NoError,
		},
		{
			name: "Invalid Credentials",
			acct: func(t *testing.T) account.Account {
				return tconfig.NewFakeM365Account(t)
			},
			expectErr: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			_, err := NewM365Client(ctx, test.acct(t))
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}
