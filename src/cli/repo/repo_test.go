package repo

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type RepoUnitSuite struct {
	tester.Suite
}

func TestRepoUnitSuite(t *testing.T) {
	suite.Run(t, &RepoUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RepoUnitSuite) TestAddRepoCommands() {
	t := suite.T()
	cmd := &cobra.Command{}

	AddCommands(cmd)

	var found bool

	// This is the repo command.
	repoCmds := cmd.Commands()
	require.Len(t, repoCmds, 1)

	for _, c := range repoCmds[0].Commands() {
		if c.Use == maintenanceCommand {
			found = true
		}
	}

	assert.True(t, found, "looking for maintenance command")
}
