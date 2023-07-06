package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
)

type OptionsUnitSuite struct {
	tester.Suite
}

func TestOptionsUnitSuite(t *testing.T) {
	suite.Run(t, &OptionsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OptionsUnitSuite) TestEnsureRestoreConfigDefaults() {
	table := []struct {
		name   string
		input  control.RestoreConfig
		expect control.RestoreConfig
	}{
		{
			name: "populated",
			input: control.RestoreConfig{
				OnCollision:       control.Copy,
				ProtectedResource: "batman",
				Location:          "badman",
				Drive:             "hatman",
			},
			expect: control.RestoreConfig{
				OnCollision:       control.Copy,
				ProtectedResource: "batman",
				Location:          "badman",
				Drive:             "hatman",
			},
		},
		{
			name: "unpopulated",
			input: control.RestoreConfig{
				OnCollision:       control.Unknown,
				ProtectedResource: "",
				Location:          "",
				Drive:             "",
			},
			expect: control.RestoreConfig{
				OnCollision:       control.Skip,
				ProtectedResource: "",
				Location:          "",
				Drive:             "",
			},
		},
		{
			name: "populated, but modified",
			input: control.RestoreConfig{
				OnCollision:       control.CollisionPolicy("batman"),
				ProtectedResource: "",
				Location:          "/",
				Drive:             "",
			},
			expect: control.RestoreConfig{
				OnCollision:       control.Skip,
				ProtectedResource: "",
				Location:          "",
				Drive:             "",
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			result := control.EnsureRestoreConfigDefaults(ctx, test.input)
			assert.Equal(t, test.expect, result)
		})
	}
}
