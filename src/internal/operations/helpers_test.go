package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
)

type HelpersUnitSuite struct {
	tester.Suite
}

func TestHelpersUnitSuite(t *testing.T) {
	suite.Run(t, &HelpersUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *HelpersUnitSuite) TestFinalizeErrorHandling() {
	table := []struct {
		name      string
		errs      func() *fault.Bus
		opts      control.Options
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name: "no errors",
			errs: func() *fault.Bus {
				return fault.New(false)
			},
			opts: control.Options{
				FailureHandling: control.FailAfterRecovery,
			},
			expectErr: assert.NoError,
		},
		{
			name: "already failed",
			errs: func() *fault.Bus {
				fn := fault.New(false)
				fn.Fail(assert.AnError)
				return fn
			},
			opts: control.Options{
				FailureHandling: control.FailAfterRecovery,
			},
			expectErr: assert.Error,
		},
		{
			name: "best effort",
			errs: func() *fault.Bus {
				fn := fault.New(false)
				fn.AddRecoverable(assert.AnError)
				return fn
			},
			opts: control.Options{
				FailureHandling: control.BestEffort,
			},
			expectErr: assert.NoError,
		},
		{
			name: "recoverable errors produce hard fail",
			errs: func() *fault.Bus {
				fn := fault.New(false)
				fn.AddRecoverable(assert.AnError)
				return fn
			},
			opts: control.Options{
				FailureHandling: control.FailAfterRecovery,
			},
			expectErr: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext()
			defer flush()

			t := suite.T()
			errs := test.errs()

			finalizeErrorHandling(ctx, test.opts, errs, "test")
			test.expectErr(t, errs.Failure())
		})
	}
}
