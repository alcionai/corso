package utils

import (
	"testing"

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
