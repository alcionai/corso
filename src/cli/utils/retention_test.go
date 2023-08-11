package utils_test

import (
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control/repository"
)

type RetentionCfgUnitSuite struct {
	tester.Suite
}

func TestRetentionCfgUnitSuite(t *testing.T) {
	suite.Run(t, &RetentionCfgUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RetentionCfgUnitSuite) TestMakeRetentionOpts() {
	table := []struct {
		name      string
		flags     map[string]string
		expectErr assert.ErrorAssertionFunc
		expect    repository.Retention
	}{
		{
			name:      "Nothing Set",
			expectErr: assert.NoError,
		},
		{
			name: "Invalid Mode",
			flags: map[string]string{
				flags.RetentionModeFN: "foo",
			},
			expectErr: assert.Error,
		},
		{
			name: "Negative Duration",
			flags: map[string]string{
				flags.RetentionDurationFN: "-5h",
			},
			expectErr: assert.Error,
		},
		{
			name: "Only Governance Mode",
			flags: map[string]string{
				flags.RetentionModeFN: "governance",
			},
			expectErr: assert.NoError,
			expect: repository.Retention{
				Mode: ptr.To(repository.GovernanceRetention),
			},
		},
		{
			name: "Only Compliance Mode",
			flags: map[string]string{
				flags.RetentionModeFN: "compliance",
			},
			expectErr: assert.NoError,
			expect: repository.Retention{
				Mode: ptr.To(repository.ComplianceRetention),
			},
		},
		{
			name: "Only No Retention Mode",
			flags: map[string]string{
				flags.RetentionModeFN: "none",
			},
			expectErr: assert.NoError,
			expect: repository.Retention{
				Mode: ptr.To(repository.NoRetention),
			},
		},
		{
			name: "Mode And Duration",
			flags: map[string]string{
				flags.RetentionModeFN:     "governance",
				flags.RetentionDurationFN: "48h",
			},
			expectErr: assert.NoError,
			expect: repository.Retention{
				Mode:     ptr.To(repository.GovernanceRetention),
				Duration: ptr.To(time.Hour * 48),
			},
		},
		{
			name: "Mode And Extend",
			flags: map[string]string{
				flags.RetentionModeFN:   "governance",
				flags.ExtendRetentionFN: "false",
			},
			expectErr: assert.NoError,
			expect: repository.Retention{
				Mode:   ptr.To(repository.GovernanceRetention),
				Extend: ptr.To(false),
			},
		},
		{
			name: "Duration And Extend",
			flags: map[string]string{
				flags.RetentionDurationFN: "48h",
				flags.ExtendRetentionFN:   "true",
			},
			expectErr: assert.NoError,
			expect: repository.Retention{
				Duration: ptr.To(time.Hour * 48),
				Extend:   ptr.To(true),
			},
		},
		{
			name: "All Set",
			flags: map[string]string{
				flags.RetentionModeFN:     "governance",
				flags.RetentionDurationFN: "48h",
				flags.ExtendRetentionFN:   "true",
			},
			expectErr: assert.NoError,
			expect: repository.Retention{
				Mode:     ptr.To(repository.GovernanceRetention),
				Duration: ptr.To(time.Hour * 48),
				Extend:   ptr.To(true),
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cmd := &cobra.Command{}
			flags.AddRetentionConfigFlags(cmd)
			fs := cmd.Flags()

			for fn, fv := range test.flags {
				require.NoError(t, fs.Set(fn, fv), "setting flag values")
			}

			result, err := utils.MakeRetentionOpts(cmd)
			test.expectErr(t, err, "parsing flags into struct: %v", clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, test.expect, result)
		})
	}
}
