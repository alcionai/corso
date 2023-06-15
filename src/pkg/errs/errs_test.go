package errs

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/repository"
)

type ErrUnitSuite struct {
	tester.Suite
}

func TestErrUnitSuite(t *testing.T) {
	suite.Run(t, &ErrUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ErrUnitSuite) TestInternal() {
	table := []struct {
		get    errEnum
		expect []error
	}{
		{RepoAlreadyExists, []error{repository.ErrorRepoAlreadyExists}},
		{BackupNotFound, []error{repository.ErrorBackupNotFound}},
		{ServiceNotEnabled, []error{graph.ErrServiceNotEnabled}},
		{ResourceOwnerNotFound, []error{graph.ErrResourceOwnerNotFound}},
	}
	for _, test := range table {
		suite.Run(string(test.get), func() {
			assert.ElementsMatch(suite.T(), test.expect, Internal(test.get))
		})
	}
}

func (suite *ErrUnitSuite) TestIs() {
	table := []struct {
		target errEnum
		err    error
	}{
		{RepoAlreadyExists, repository.ErrorRepoAlreadyExists},
		{BackupNotFound, repository.ErrorBackupNotFound},
		{ServiceNotEnabled, graph.ErrServiceNotEnabled},
		{ResourceOwnerNotFound, graph.ErrResourceOwnerNotFound},
	}
	for _, test := range table {
		suite.Run(string(test.target), func() {
			var (
				w  = clues.Wrap(test.err, "wrap")
				s  = clues.Stack(test.err)
				es = clues.Stack(assert.AnError, test.err)
				se = clues.Stack(test.err, assert.AnError)
				sw = clues.Stack(assert.AnError, w)
				ws = clues.Stack(w, assert.AnError)
			)

			assert.True(suite.T(), Is(test.err, test.target))
			assert.True(suite.T(), Is(w, test.target))
			assert.True(suite.T(), Is(s, test.target))
			assert.True(suite.T(), Is(es, test.target))
			assert.True(suite.T(), Is(se, test.target))
			assert.True(suite.T(), Is(sw, test.target))
			assert.True(suite.T(), Is(ws, test.target))
			assert.False(suite.T(), Is(assert.AnError, test.target))
		})
	}
}
