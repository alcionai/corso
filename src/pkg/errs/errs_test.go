package errs

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/tester"
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
		{RepoAlreadyExists, clues.New("a repository was already initialized with that configuration")},
		{RepoAlreadyExists, errors.Wrapf(clues.New("a repository was already initialized with that configuration"), "error identifying resource owner")},
		{BackupNotFound, clues.New("no backup exists with that id")},
		{ServiceNotEnabled, clues.New("service is not enabled for that resource owner")},
		{ResourceOwnerNotFound, clues.New("resource owner not found in tenant")},
	}
	for _, test := range table {
		suite.Run(string(test.is), func() {
			assert.True(suite.T(), Is(test.input, test.is))
			assert.False(suite.T(), Is(assert.AnError, test.is))
		})
	}
}
