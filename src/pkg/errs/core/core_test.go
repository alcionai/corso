package core_test

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/pkg/errs/core"
)

type ErrUnitSuite struct {
	tester.Suite
}

func TestErrUnitSuite(t *testing.T) {
	suite.Run(t, &ErrUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ErrUnitSuite) TestAs() {
	// shorthand reference for ease of reading
	cErr := core.ErrApplicationThrottled
	adHoc := &core.Err{}

	table := []struct {
		name      string
		err       error
		expectOK  assert.BoolAssertionFunc
		expectErr func(t *testing.T, ce *core.Err)
	}{
		{
			name:     "nil",
			err:      nil,
			expectOK: assert.False,
			expectErr: func(t *testing.T, ce *core.Err) {
				assert.Nil(t, ce)
			},
		},
		{
			name:     "non-matching",
			err:      assert.AnError,
			expectOK: assert.False,
			expectErr: func(t *testing.T, ce *core.Err) {
				assert.Nil(t, ce)
			},
		},
		{
			name:     "matching",
			err:      cErr,
			expectOK: assert.True,
			expectErr: func(t *testing.T, ce *core.Err) {
				assert.Equal(t, cErr, ce)
			},
		},
		{
			name:     "adHoc",
			err:      adHoc,
			expectOK: assert.True,
			expectErr: func(t *testing.T, ce *core.Err) {
				assert.Equal(t, adHoc, ce)
			},
		},
		{
			name:     "stacked",
			err:      clues.Stack(assert.AnError, cErr, assert.AnError),
			expectOK: assert.True,
			expectErr: func(t *testing.T, ce *core.Err) {
				assert.Equal(t, cErr, ce)
			},
		},
		{
			name:     "wrapped",
			err:      clues.Wrap(cErr, "wrapper"),
			expectOK: assert.True,
			expectErr: func(t *testing.T, ce *core.Err) {
				assert.Equal(t, cErr, ce)
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			err, ok := core.As(test.err)

			test.expectOK(t, ok)
			test.expectErr(t, err)
		})
	}
}
