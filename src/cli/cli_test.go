package cli_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli"
)

type CLISuite struct {
	suite.Suite
}

func TestCLISuite(t *testing.T) {
	suite.Run(t, new(CLISuite))
}

func (suite *CLISuite) TestAddCommands_noPanics() {
	t := suite.T()

	test := &cobra.Command{
		Use:   "test",
		Short: "Free, Secure, Open-Source Backup for M365.",
		Long:  `Free, Secure, and Open-Source Backup for Microsoft 365.`,
		RunE:  func(c *cobra.Command, args []string) error { return nil },
	}

	assert.NotPanics(t, func() { cli.BuildCommandTree(test) })
}
