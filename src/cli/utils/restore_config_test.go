package utils

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/flags"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
)

type RestoreCfgUnitSuite struct {
	tester.Suite
}

func TestRestoreCfgUnitSuite(t *testing.T) {
	suite.Run(t, &RestoreCfgUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RestoreCfgUnitSuite) TestValidateRestoreConfigFlags() {
	table := []struct {
		name   string
		fv     string
		opts   RestoreCfgOpts
		expect assert.ErrorAssertionFunc
	}{
		{
			name: "no error",
			fv:   string(control.Skip),
			opts: RestoreCfgOpts{
				Collisions: string(control.Skip),
				Populated: flags.PopulatedFlags{
					flags.CollisionsFN: {},
				},
			},
			expect: assert.NoError,
		},
		{
			name: "bad but not populated",
			fv:   "foo",
			opts: RestoreCfgOpts{
				Collisions: "foo",
				Populated:  flags.PopulatedFlags{},
			},
			expect: assert.NoError,
		},
		{
			name: "error",
			fv:   "foo",
			opts: RestoreCfgOpts{
				Collisions: "foo",
				Populated: flags.PopulatedFlags{
					flags.CollisionsFN: {},
				},
			},
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			err := validateRestoreConfigFlags(test.fv, test.opts)
			test.expect(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *RestoreCfgUnitSuite) TestMakeRestoreConfig() {
	rco := &RestoreCfgOpts{
		Collisions:  "collisions",
		Destination: "destination",
	}

	table := []struct {
		name      string
		populated flags.PopulatedFlags
		expect    control.RestoreConfig
	}{
		{
			name:      "not populated",
			populated: flags.PopulatedFlags{},
			expect: control.RestoreConfig{
				OnCollision: control.Skip,
				Location:    "Corso_Restore_",
			},
		},
		{
			name: "collision populated",
			populated: flags.PopulatedFlags{
				flags.CollisionsFN: {},
			},
			expect: control.RestoreConfig{
				OnCollision: control.CollisionPolicy("collisions"),
				Location:    "Corso_Restore_",
			},
		},
		{
			name: "destination populated",
			populated: flags.PopulatedFlags{
				flags.DestinationFN: {},
			},
			expect: control.RestoreConfig{
				OnCollision: control.Skip,
				Location:    "destination",
			},
		},
		{
			name: "both populated",
			populated: flags.PopulatedFlags{
				flags.CollisionsFN:  {},
				flags.DestinationFN: {},
			},
			expect: control.RestoreConfig{
				OnCollision: control.CollisionPolicy("collisions"),
				Location:    "destination",
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

			result := MakeRestoreConfig(ctx, opts, dttm.HumanReadable)
			assert.Equal(t, test.expect.OnCollision, result.OnCollision)
			assert.Contains(t, result.Location, test.expect.Location)
		})
	}
}
