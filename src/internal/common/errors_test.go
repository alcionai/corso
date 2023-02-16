package common_test

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/tester"
)

type testErr struct {
	common.Err
}

type testErr2 struct {
	common.Err
}

type ErrorsUnitSuite struct {
	tester.Suite
}

func TestErrorsUnitSuite(t *testing.T) {
	s := &ErrorsUnitSuite{Suite: tester.NewUnitSuite(t)}
	suite.Run(t, s)
}

func (suite *ErrorsUnitSuite) TestPropagatesCause() {
	err := assert.AnError
	te := testErr{*common.EncapsulateError(err)}
	te2 := testErr2{*common.EncapsulateError(te)}

	assert.Equal(suite.T(), assert.AnError, errors.Cause(te2))
}

func (suite *ErrorsUnitSuite) TestPropagatesIs() {
	err := assert.AnError
	te := testErr{*common.EncapsulateError(err)}
	te2 := testErr2{*common.EncapsulateError(te)}

	assert.True(suite.T(), errors.Is(te2, err))
}

func (suite *ErrorsUnitSuite) TestPropagatesAs() {
	err := assert.AnError
	te := testErr{*common.EncapsulateError(err)}
	te2 := testErr2{*common.EncapsulateError(te)}

	tmp := testErr{}
	assert.True(suite.T(), errors.As(te2, &tmp))
}

func (suite *ErrorsUnitSuite) TestAs() {
	err := assert.AnError
	te := testErr{*common.EncapsulateError(err)}
	te2 := testErr2{*common.EncapsulateError(te)}

	tmp := testErr2{}
	assert.True(suite.T(), errors.As(te2, &tmp))
}

func (suite *ErrorsUnitSuite) TestAsIsUnique() {
	err := assert.AnError
	te := testErr{*common.EncapsulateError(err)}

	tmp := testErr2{}
	assert.False(suite.T(), errors.As(te, &tmp))
}

func (suite *ErrorsUnitSuite) TestPrintsStack() {
	err := assert.AnError
	err = errors.Wrap(err, "wrapped error")
	te := testErr{*common.EncapsulateError(err)}
	te2 := testErr2{*common.EncapsulateError(te)}

	out := fmt.Sprintf("%+v", te2)

	// Stack trace should include a line noting that we're running testify.
	assert.Contains(suite.T(), out, "testify")
}
