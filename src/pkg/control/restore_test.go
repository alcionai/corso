package control_test

import (
	"fmt"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
)

type RestoreUnitSuite struct {
	tester.Suite
}

func TestRestoreUnitSuite(t *testing.T) {
	suite.Run(t, &RestoreUnitSuite{Suite: tester.NewUnitSuite(t)})
}

// set the clues hashing to mask for the span of this suite
func (suite *RestoreUnitSuite) SetupSuite() {
	clues.SetHasher(clues.HashCfg{HashAlg: clues.Flatmask})
}

// revert clues hashing to plaintext for all other tests
func (suite *RestoreUnitSuite) TeardownSuite() {
	clues.SetHasher(clues.NoHash())
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

func (suite *RestoreUnitSuite) TestRestoreConfig_piiHandling() {
	t := suite.T()

	p, err := path.Build("tid", "ro", path.ExchangeService, path.EmailCategory, true, "foo", "bar", "baz")
	require.NoError(t, err, clues.ToCore(err))

	cdrc := control.DefaultRestoreConfig(dttm.HumanReadable)

	table := []struct {
		name        string
		rc          control.RestoreConfig
		expectSafe  string
		expectPlain string
	}{
		{
			name:        "empty",
			expectSafe:  `{"onCollision":"","protectedResource":"","location":"","drive":"","includePermissions":false}`,
			expectPlain: `{"onCollision":"","protectedResource":"","location":"","drive":"","includePermissions":false}`,
		},
		{
			name:       "defaults",
			rc:         cdrc,
			expectSafe: `{"onCollision":"skip","protectedResource":"","location":"***","drive":"","includePermissions":false}`,
			expectPlain: `{"onCollision":"skip","protectedResource":"","location":"` +
				cdrc.Location + `","drive":"","includePermissions":false}`,
		},
		{
			name: "populated",
			rc: control.RestoreConfig{
				OnCollision:        control.Copy,
				ProtectedResource:  "snoob",
				Location:           p.String(),
				Drive:              "somedriveid",
				IncludePermissions: true,
			},
			expectSafe: `{"onCollision":"copy","protectedResource":"***","location":"***/exchange/***/email/***/***/***",` +
				`"drive":"***","includePermissions":true}`,
			expectPlain: `{"onCollision":"copy","protectedResource":"snoob","location":"tid/exchange/ro/email/foo/bar/baz",` +
				`"drive":"somedriveid","includePermissions":true}`,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			assert.Equal(t, test.expectSafe, test.rc.Conceal(), "conceal")
			assert.Equal(t, test.expectPlain, test.rc.String(), "string")
			assert.Equal(t, test.expectSafe, fmt.Sprintf("%s", test.rc), "fmt %%s")
			assert.Equal(t, test.expectSafe, fmt.Sprintf("%+v", test.rc), "fmt %%+v")
			assert.Equal(t, test.expectPlain, test.rc.PlainString(), "plain")
		})
	}
}
