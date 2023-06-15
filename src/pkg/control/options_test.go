package control

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
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
		input  RestoreConfig
		expect RestoreConfig
	}{
		{
			name: "populated",
			input: RestoreConfig{
				OnCollision:       Copy,
				ProtectedResource: "batman",
				Location:          "badman",
				Drive:             "hatman",
			},
			expect: RestoreConfig{
				OnCollision:       Copy,
				ProtectedResource: "batman",
				Location:          "badman",
				Drive:             "hatman",
			},
		},
		{
			name: "unpopulated",
			input: RestoreConfig{
				OnCollision:       Unknown,
				ProtectedResource: "",
				Location:          "",
				Drive:             "",
			},
			expect: RestoreConfig{
				OnCollision:       Skip,
				ProtectedResource: "",
				Location:          "",
				Drive:             "",
			},
		},
		{
			name: "populated, but modified",
			input: RestoreConfig{
				OnCollision:       CollisionPolicy("batman"),
				ProtectedResource: "",
				Location:          "/",
				Drive:             "",
			},
			expect: RestoreConfig{
				OnCollision:       Skip,
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

			result := EnsureRestoreConfigDefaults(ctx, test.input)
			assert.Equal(t, test.expect, result)
		})
	}
}
