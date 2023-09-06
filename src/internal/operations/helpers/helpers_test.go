package helpers

import (
	"context"
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
		errs      func(context.Context) *fault.Bus
		opts      control.Options
		expectErr assert.ErrorAssertionFunc
	}{
		{
			name: "no errors",
			errs: func(ctx context.Context) *fault.Bus {
				return fault.New(false)
			},
			opts: control.Options{
				FailureHandling: control.FailAfterRecovery,
			},
			expectErr: assert.NoError,
		},
		{
			name: "already failed",
			errs: func(ctx context.Context) *fault.Bus {
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
			errs: func(ctx context.Context) *fault.Bus {
				fn := fault.New(false)
				fn.AddRecoverable(ctx, assert.AnError)
				return fn
			},
			opts: control.Options{
				FailureHandling: control.BestEffort,
			},
			expectErr: assert.NoError,
		},
		{
			name: "recoverable errors produce hard fail",
			errs: func(ctx context.Context) *fault.Bus {
				fn := fault.New(false)
				fn.AddRecoverable(ctx, assert.AnError)
				return fn
			},
			opts: control.Options{
				FailureHandling: control.FailAfterRecovery,
			},
			expectErr: assert.Error,
		},
		{
			name: "multiple recoverable errors produce hard fail",
			errs: func(ctx context.Context) *fault.Bus {
				fn := fault.New(false)
				fn.AddRecoverable(ctx, assert.AnError)
				fn.AddRecoverable(ctx, assert.AnError)
				fn.AddRecoverable(ctx, assert.AnError)
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
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			errs := test.errs(ctx)

			FinalizeErrorHandling(ctx, test.opts, errs, "test")
			test.expectErr(t, errs.Failure())
		})
	}
}
