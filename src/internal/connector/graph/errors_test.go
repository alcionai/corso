package graph

import (
	"context"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GraphErrorsUnitSuite struct {
	suite.Suite
}

func TestGraphErrorsUnitSuite(t *testing.T) {
	suite.Run(t, new(GraphErrorsUnitSuite))
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
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(t, IsErrDeletedInFlight(test.err))
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
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(t, IsErrInvalidDelta(test.err))
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
		suite.T().Run(test.name, func(t *testing.T) {
			test.expect(t, IsErrTimeout(test.err))
		})
	}
}
