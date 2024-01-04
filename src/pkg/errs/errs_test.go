package errs

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/repository"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	graphTD "github.com/alcionai/corso/src/pkg/services/m365/api/graph/testdata"
)

type ErrUnitSuite struct {
	tester.Suite
}

func TestErrUnitSuite(t *testing.T) {
	suite.Run(t, &ErrUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ErrUnitSuite) TestInternal_errs() {
	table := []struct {
		get    *core.Err
		expect []error
	}{
		{
			get:    core.ErrRepoAlreadyExists,
			expect: []error{repository.ErrorRepoAlreadyExists},
		},
		{
			get:    core.ErrBackupNotFound,
			expect: []error{repository.ErrorBackupNotFound},
		},
		{
			get:    core.ErrResourceNotAccessible,
			expect: []error{graph.ErrResourceLocked},
		},
	}
	for _, test := range table {
		suite.Run(test.get.Error(), func() {
			// can't compare func signatures
			errs, _ := Internal(test.get)
			assert.ElementsMatch(suite.T(), test.expect, errs)
		})
	}
}

func (suite *ErrUnitSuite) TestInternal_checks() {
	table := []struct {
		get             *core.Err
		err             error
		expectHasChecks assert.ValueAssertionFunc
		expect          assert.BoolAssertionFunc
	}{
		{
			get:             core.ErrRepoAlreadyExists,
			err:             graphTD.ODataErr(string(graph.ApplicationThrottled)),
			expectHasChecks: assert.Empty,
			expect:          assert.False,
		},
		{
			get:             core.ErrBackupNotFound,
			err:             repository.ErrorBackupNotFound,
			expectHasChecks: assert.Empty,
			expect:          assert.False,
		},
		{
			get:             core.ErrResourceOwnerNotFound,
			err:             graphTD.ODataErr(string(graph.ItemNotFound)),
			expectHasChecks: assert.NotEmpty,
			expect:          assert.True,
		},
		{
			get:             core.ErrResourceOwnerNotFound,
			err:             graphTD.ODataErr(string(graph.ErrorItemNotFound)),
			expectHasChecks: assert.NotEmpty,
			expect:          assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.get.Error(), func() {
			t := suite.T()

			_, checks := Internal(test.get)

			test.expectHasChecks(t, checks)

			var result bool

			for _, check := range checks {
				if check(test.err) {
					result = true
					break
				}
			}

			test.expect(t, result)
		})
	}
}

func (suite *ErrUnitSuite) TestIs() {
	table := []struct {
		target *core.Err
		err    error
	}{
		{
			target: core.ErrRepoAlreadyExists,
			err:    repository.ErrorRepoAlreadyExists,
		},
		{
			target: core.ErrBackupNotFound,
			err:    repository.ErrorBackupNotFound,
		},
		{
			target: core.ErrResourceNotAccessible,
			err:    graph.ErrResourceLocked,
		},
	}
	for _, test := range table {
		suite.Run(test.target.Error(), func() {
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
