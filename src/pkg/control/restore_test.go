package control_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
)

type RestoreUnitSuite struct {
	tester.Suite
}

func TestRestoreUnitSuite(t *testing.T) {
	suite.Run(t, &RestoreUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *RestoreUnitSuite) TestEnsureRestoreConfigDefaults() {
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
		{
			name: "populated with slash prefix, then modified",
			input: control.RestoreConfig{
				OnCollision:       control.CollisionPolicy("batman"),
				ProtectedResource: "",
				Location:          "/smarfs",
				Drive:             "",
			},
			expect: control.RestoreConfig{
				OnCollision:       control.Skip,
				ProtectedResource: "",
				Location:          "smarfs",
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
