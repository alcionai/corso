package restore

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/cli/utils"
	ctesting "github.com/alcionai/corso/internal/testing"
)

type ExchangeSuite struct {
	suite.Suite
}

func TestExchangeSuite(t *testing.T) {
	suite.Run(t, new(ExchangeSuite))
}

func (suite *ExchangeSuite) TestValidateRestoreFlags() {
	table := []struct {
		name          string
		u, f, m, rpid string
		errCheck      assert.ErrorAssertionFunc
	}{
		{"all populated", "u", "f", "m", "rpid", assert.NoError},
		{"folder missing user", "", "f", "m", "rpid", assert.Error},
		{"folder with wildcard user", utils.Wildcard, "f", "m", "rpid", assert.Error},
		{"mail missing user", "", "", "m", "rpid", assert.Error},
		{"mail missing folder", "u", "", "m", "rpid", assert.Error},
		{"mail with wildcard folder", "u", utils.Wildcard, "m", "rpid", assert.Error},
		{"missing backup id", "u", "f", "m", "", assert.Error},
		{"all missing", "", "", "", "rpid", assert.NoError},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.errCheck(
				t,
				validateRestoreFlags(test.u, test.f, test.m, test.rpid),
			)
		})
	}
}

func (suite *ExchangeSuite) TestAddExchangeCommands() {
	expectUse := exchangeServiceCommand
	table := []struct {
		name        string
		use         string
		expectUse   string
		expectShort string
		expectRunE  func(*cobra.Command, []string) error
	}{
		{"restore exchange", restoreCommand, expectUse, exchangeRestoreCmd.Short, restoreExchangeCmd},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			cmd := &cobra.Command{Use: test.use}

			c := addExchangeCommands(cmd)
			require.NotNil(t, c)

			cmds := cmd.Commands()
			require.Len(t, cmds, 1)

			child := cmds[0]
			assert.Equal(t, test.expectUse, child.Use)
			assert.Equal(t, test.expectShort, child.Short)
			ctesting.AreSameFunc(t, test.expectRunE, child.RunE)
		})
	}
}
