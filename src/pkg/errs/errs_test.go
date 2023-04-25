package errs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/repository"
)

type ErrUnitSuite struct {
	tester.Suite
}

func TestErrUnitSuite(t *testing.T) {
	suite.Run(t, &ErrUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ErrUnitSuite) TestIs() {
	table := []struct {
		is    errEnum
		input error
	}{
		{RepoAlreadyExists, repository.ErrorRepoAlreadyExists},
		{BackupNotFound, repository.ErrorBackupNotFound},
		{ServiceNotEnabled, graph.ErrServiceNotEnabled},
		{ResourceOwnerNotFound, graph.ErrResourceOwnerNotFound},
	}
	for _, test := range table {
		suite.Run(string(test.is), func() {
			assert.True(suite.T(), Is(test.input, test.is))
			assert.False(suite.T(), Is(assert.AnError, test.is))
		})
	}
}
