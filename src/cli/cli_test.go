package cli_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli"
	"github.com/alcionai/corso/src/internal/tester"
)

type CLIUnitSuite struct {
	tester.Suite
}

func TestCLIUnitSuite(t *testing.T) {
	suite.Run(t, &CLIUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CLIUnitSuite) TestAddCommands_noPanics() {
	t := suite.T()

	test := &cobra.Command{
		Use:   "test",
		Short: "Free, Secure, Open-Source Backup for M365.",
		Long:  `Free, Secure, and Open-Source Backup for Microsoft 365.`,
		RunE:  func(c *cobra.Command, args []string) error { return nil },
	}

	assert.NotPanics(t, func() { cli.BuildCommandTree(test) })
}
