package crash_test

import (
	"testing"

	"github.com/alcionai/corso/src/internal/common/crash"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CrashTestDummySuite struct {
	tester.Suite
}

func TestCrashTestDummySuite(t *testing.T) {
	suite.Run(t, &CrashTestDummySuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CrashTestDummySuite) TestRecovery() {
	table := []struct {
		name   string
		v      any
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "no panic",
			v:      nil,
			expect: assert.NoError,
		},
		{
			name:   "error panic",
			v:      assert.AnError,
			expect: assert.Error,
		},
		{
			name:   "string panic",
			v:      "an error",
			expect: assert.Error,
		},
		{
			name:   "any panic",
			v:      map[string]string{"error": "yes"},
			expect: assert.Error,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			ctx, flush := tester.NewContext()

			defer func() {
				test.expect(t, crash.Recovery(ctx, recover()))
				flush()
			}()

			if test.v != nil {
				panic(test.v)
			}
		})
	}
}
