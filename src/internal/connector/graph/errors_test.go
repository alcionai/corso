package graph

import (
	"context"
	"net/http"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/tester"
)

type GraphErrorsUnitSuite struct {
	tester.Suite
}

func TestGraphErrorsUnitSuite(t *testing.T) {
	suite.Run(t, &GraphErrorsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func odErr(code string) *odataerrors.ODataError {
	odErr := &odataerrors.ODataError{}
	merr := odataerrors.MainError{}
	merr.SetCode(&code)
	odErr.SetError(&merr)

	return odErr
}

func (suite *GraphErrorsUnitSuite) TestIsErrDeletedInFlight() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "as",
			err:    ErrDeletedInFlight,
			expect: assert.True,
		},
		{
			name:   "non-matching oDataErr",
			err:    odErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "not-found oDataErr",
			err:    odErr(errCodeItemNotFound),
			expect: assert.True,
		},
		{
			name:   "sync-not-found oDataErr",
			err:    odErr(errCodeSyncFolderNotFound),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrDeletedInFlight(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrInvalidDelta() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "as",
			err:    ErrInvalidDelta,
			expect: assert.True,
		},
		{
			name:   "non-matching oDataErr",
			err:    odErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "resync-required oDataErr",
			err:    odErr(errCodeResyncRequired),
			expect: assert.True,
		},
		// next two tests are to make sure the checks are case insensitive
		{
			name:   "resync-required oDataErr camelcase",
			err:    odErr("resyncRequired"),
			expect: assert.True,
		},
		{
			name:   "resync-required oDataErr lowercase",
			err:    odErr("resyncrequired"),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrInvalidDelta(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrUserNotFound() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "non-matching oDataErr",
			err:    odErr("fnords"),
			expect: assert.False,
		},
		{
			name:   "request resource not found oDataErr",
			err:    odErr(errCodeRequestResourceNotFound),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrUserNotFound(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrTimeout() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name:   "as",
			err:    ErrTimeout,
			expect: assert.True,
		},
		{
			name:   "context deadline",
			err:    context.DeadlineExceeded,
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrTimeout(test.err))
		})
	}
}

func (suite *GraphErrorsUnitSuite) TestIsErrUnauthorized() {
	table := []struct {
		name   string
		err    error
		expect assert.BoolAssertionFunc
	}{
		{
			name:   "nil",
			err:    nil,
			expect: assert.False,
		},
		{
			name:   "non-matching",
			err:    assert.AnError,
			expect: assert.False,
		},
		{
			name: "as",
			err: clues.Stack(assert.AnError).
				Label(LabelStatus(http.StatusUnauthorized)),
			expect: assert.True,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), IsErrUnauthorized(test.err))
		})
	}
}
