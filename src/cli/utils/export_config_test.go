package utils

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
)

type ExportCfgUnitSuite struct {
	tester.Suite
}

func TestExportCfgUnitSuite(t *testing.T) {
	suite.Run(t, &ExportCfgUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExportCfgUnitSuite) TestMakeExportConfig() {
	rco := &ExportCfgOpts{Archive: true}

	table := []struct {
		name      string
		populated flags.PopulatedFlags
		expect    control.ExportConfig
	}{
		{
			name: "archive populated",
			populated: flags.PopulatedFlags{
				flags.ArchiveFN: {},
			},
			expect: control.ExportConfig{
				Archive: true,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			opts := *rco
			opts.Populated = test.populated

			result := MakeExportConfig(ctx, opts)
			assert.Equal(t, test.expect.Archive, result.Archive)
		})
	}
}

func (suite *ExportCfgUnitSuite) TestValidateExportConfigFlags() {
	acceptedFormatTypes := []string{
		string(control.DefaultFormat),
		string(control.JSONFormat),
	}

	table := []struct {
		name         string
		input        ExportCfgOpts
		expectErr    assert.ErrorAssertionFunc
		expectFormat control.FormatType
	}{
		{
			name: "default",
			input: ExportCfgOpts{
				Format:    string(control.DefaultFormat),
				Populated: flags.PopulatedFlags{flags.FormatFN: struct{}{}},
			},
			expectErr:    assert.NoError,
			expectFormat: control.DefaultFormat,
		},
		{
			name: "json",
			input: ExportCfgOpts{
				Format:    string(control.JSONFormat),
				Populated: flags.PopulatedFlags{flags.FormatFN: struct{}{}},
			},
			expectErr:    assert.NoError,
			expectFormat: control.JSONFormat,
		},
		{
			name: "bad format",
			input: ExportCfgOpts{
				Format:    "smurfs",
				Populated: flags.PopulatedFlags{flags.FormatFN: struct{}{}},
			},
			expectErr:    assert.Error,
			expectFormat: control.DefaultFormat,
		},
		{
			name: "bad format unpopulated",
			input: ExportCfgOpts{
				Format: "smurfs",
			},
			expectErr:    assert.NoError,
			expectFormat: control.DefaultFormat,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			err := ValidateExportConfigFlags(&test.input, acceptedFormatTypes)

			test.expectErr(t, err, clues.ToCore(err))
			assert.Equal(t, test.expectFormat, control.FormatType(test.input.Format))
		})
	}
}
