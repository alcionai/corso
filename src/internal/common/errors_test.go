package common_test

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/common"
)

type testErr struct {
	common.BaseError
}

type testErr2 struct {
	common.BaseError
}

type ErrorsUnitSuite struct {
	suite.Suite
}

func TestErrorsUnitSuite(t *testing.T) {
	suite.Run(t, new(ErrorsUnitSuite))
}

func (suite *ErrorsUnitSuite) TestPropagatesCause() {
	err := assert.AnError
	te := testErr{common.BaseError{err}}
	te2 := testErr2{common.BaseError{te}}

	assert.Equal(suite.T(), assert.AnError, errors.Cause(te2))
}

func (suite *ErrorsUnitSuite) TestPropagatesIs() {
	err := assert.AnError
	te := testErr{common.BaseError{err}}
	te2 := testErr2{common.BaseError{te}}

	assert.True(suite.T(), errors.Is(te2, err))
}

func (suite *ErrorsUnitSuite) TestPropagatesAs() {
	err := assert.AnError
	te := testErr{common.BaseError{err}}
	te2 := testErr2{common.BaseError{te}}

	var tmp testErr
	assert.True(suite.T(), errors.As(te2, &tmp))
}

func (suite *ErrorsUnitSuite) TestAs() {
	err := assert.AnError
	te := testErr{common.BaseError{err}}
	te2 := testErr2{common.BaseError{te}}

	var tmp testErr2
	assert.True(suite.T(), errors.As(te2, &tmp))
}

func (suite *ErrorsUnitSuite) TestAsIsUnique() {
	err := assert.AnError
	te := testErr{common.BaseError{err}}

	var tmp testErr2
	assert.False(suite.T(), errors.As(te, &tmp))
}

func (suite *ErrorsUnitSuite) TestPrintsStack() {
	err := assert.AnError
	err = errors.Wrap(err, "wrapped error")
	te := testErr{common.BaseError{err}}
	te2 := testErr2{common.BaseError{te}}

	out := fmt.Sprintf("%+v", te2)

	// Stack trace should include a line noting that we're running testify.
	assert.Contains(suite.T(), out, "testify")
}
