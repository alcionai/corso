package errs

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
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
		get    errEnum
		expect []error
	}{
		{
			get:    ApplicationThrottled,
			expect: []error{graph.ErrApplicationThrottled},
		},
		{
			get:    RepoAlreadyExists,
			expect: []error{repository.ErrorRepoAlreadyExists},
		},
		{
			get:    BackupNotFound,
			expect: []error{repository.ErrorBackupNotFound},
		},
		{
			get:    ServiceNotEnabled,
			expect: []error{graph.ErrServiceNotEnabled},
		},
		{
			get:    ResourceOwnerNotFound,
			expect: []error{graph.ErrResourceOwnerNotFound},
		},
		{
			get:    ResourceNotAccessible,
			expect: []error{graph.ErrResourceLocked},
		},
	}
	for _, test := range table {
		suite.Run(string(test.get), func() {
			// can't compare func signatures
			errs, _ := Internal(test.get)
			assert.ElementsMatch(suite.T(), test.expect, errs)
		})
	}
}

func (suite *ErrUnitSuite) TestInternal_checks() {
	table := []struct {
		get             errEnum
		err             error
		expectHasChecks assert.ValueAssertionFunc
		expect          assert.BoolAssertionFunc
	}{
		{
			get:             ApplicationThrottled,
			err:             graph.ErrApplicationThrottled,
			expectHasChecks: assert.NotEmpty,
			expect:          assert.True,
		},
		{
			get:             RepoAlreadyExists,
			err:             graph.ErrApplicationThrottled,
			expectHasChecks: assert.Empty,
			expect:          assert.False,
		},
		{
			get:             BackupNotFound,
			err:             repository.ErrorBackupNotFound,
			expectHasChecks: assert.Empty,
			expect:          assert.False,
		},
		{
			get:             ServiceNotEnabled,
			err:             graph.ErrServiceNotEnabled,
			expectHasChecks: assert.Empty,
			expect:          assert.False,
		},
		{
			get: ResourceOwnerNotFound,
			// won't match, checks itemNotFound, which isn't an error enum
			err:             graph.ErrResourceOwnerNotFound,
			expectHasChecks: assert.NotEmpty,
			expect:          assert.False,
		},
		{
			get:             ResourceNotAccessible,
			err:             graph.ErrResourceLocked,
			expectHasChecks: assert.NotEmpty,
			expect:          assert.True,
		},
		{
			get:             InsufficientAuthorization,
			err:             graphTD.ODataErr(string(graph.AuthorizationRequestDenied)),
			expectHasChecks: assert.NotEmpty,
			expect:          assert.True,
		},
	}
	for _, test := range table {
		suite.Run(string(test.get), func() {
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
		target errEnum
		err    error
	}{
		{
			target: ApplicationThrottled,
			err:    graph.ErrApplicationThrottled,
		},
		{
			target: RepoAlreadyExists,
			err:    repository.ErrorRepoAlreadyExists,
		},
		{
			target: BackupNotFound,
			err:    repository.ErrorBackupNotFound,
		},
		{
			target: ServiceNotEnabled,
			err:    graph.ErrServiceNotEnabled,
		},
		{
			target: ResourceOwnerNotFound,
			err:    graph.ErrResourceOwnerNotFound,
		},
		{
			target: ResourceNotAccessible,
			err:    graph.ErrResourceLocked,
		},
		{
			target: InsufficientAuthorization,
			err:    graphTD.ODataErr(string(graph.AuthorizationRequestDenied)),
		},
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
