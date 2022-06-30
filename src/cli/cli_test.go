package cli_test

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/cli/backup"
	"github.com/alcionai/corso/cli/repo"
	"github.com/alcionai/corso/cli/restore"
)

type CliSuite struct {
	suite.Suite
}

type CLISuite struct {
	suite.Suite
}

func TestCLISuite(t *testing.T) {
	suite.Run(t, new(CLISuite))
}

func (suite *CLISuite) TestAddCommands_noPanics() {
	var test = &cobra.Command{
		Use:   "test",
		Short: "Protect your Microsoft 365 data.",
		Long:  `Reliable, secure, and efficient data protection for Microsoft 365.`,
		RunE:  func(c *cobra.Command, args []string) error { return nil },
	}

	repo.AddCommands(test)
	backup.AddCommands(test)
	restore.AddCommands(test)
}
