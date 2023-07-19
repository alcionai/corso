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
	table := []struct {
		name      string
		rco       *RestoreCfgOpts
		populated flags.PopulatedFlags
		expect    control.RestoreConfig
	}{
		{
			name: "not populated",
			rco: &RestoreCfgOpts{
				Collisions:  "collisions",
				Destination: "destination",
			},
			populated: flags.PopulatedFlags{},
			expect: control.RestoreConfig{
				OnCollision: control.Skip,
				Location:    "Corso_Restore_",
			},
		},
		{
			name: "collision populated",
			rco: &RestoreCfgOpts{
				Collisions:  "collisions",
				Destination: "destination",
			},
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
			rco: &RestoreCfgOpts{
				Collisions:  "collisions",
				Destination: "destination",
			},
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
			rco: &RestoreCfgOpts{
				Collisions:  "collisions",
				Destination: "destination",
			},
			populated: flags.PopulatedFlags{
				flags.CollisionsFN:  {},
				flags.DestinationFN: {},
			},
			expect: control.RestoreConfig{
				OnCollision: control.CollisionPolicy("collisions"),
				Location:    "destination",
			},
		},
		{
			name: "with restore permissions",
			rco: &RestoreCfgOpts{
				Collisions:         "collisions",
				Destination:        "destination",
				RestorePermissions: true,
			},
			populated: flags.PopulatedFlags{
				flags.CollisionsFN:  {},
				flags.DestinationFN: {},
			},
			expect: control.RestoreConfig{
				OnCollision:        control.CollisionPolicy("collisions"),
				Location:           "destination",
				IncludePermissions: true,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			opts := *test.rco
			opts.Populated = test.populated

			result := MakeRestoreConfig(ctx, opts)
			assert.Equal(t, test.expect.OnCollision, result.OnCollision)
			assert.Contains(t, result.Location, test.expect.Location)
			assert.Equal(t, test.expect.IncludePermissions, result.IncludePermissions)
		})
	}
}
