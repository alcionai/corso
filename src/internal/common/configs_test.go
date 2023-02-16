package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/tester"
)

type CommonConfigsSuite struct {
	tester.Suite
}

func TestCommonConfigsSuite(t *testing.T) {
	s := &CommonConfigsSuite{Suite: tester.NewUnitSuite(t)}
	suite.Run(t, s)
}

const (
	keyExpect  = "expect"
	keyExpect2 = "expect2"
)

type stringConfig struct {
	expectA string
	err     error
}

func (c stringConfig) StringConfig() (map[string]string, error) {
	return map[string]string{keyExpect: c.expectA}, c.err
}

type stringConfig2 struct {
	expectB string
	err     error
}

func (c stringConfig2) StringConfig() (map[string]string, error) {
	return map[string]string{keyExpect2: c.expectB}, c.err
}

func (suite *CommonConfigsSuite) TestUnionConfigs_string() {
	table := []struct {
		name     string
		ac       stringConfig
		bc       stringConfig2
		errCheck assert.ErrorAssertionFunc
	}{
		{"no error", stringConfig{keyExpect, nil}, stringConfig2{keyExpect2, nil}, assert.NoError},
		{"tc error", stringConfig{keyExpect, assert.AnError}, stringConfig2{keyExpect2, nil}, assert.Error},
		{"fc error", stringConfig{keyExpect, nil}, stringConfig2{keyExpect2, assert.AnError}, assert.Error},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			cs, err := common.UnionStringConfigs(test.ac, test.bc)
			test.errCheck(t, err)
			// remaining tests depend on error-free state
			if test.ac.err != nil || test.bc.err != nil {
				return
			}
			assert.Equalf(t,
				test.ac.expectA,
				cs[keyExpect],
				"expected unioned config to have value [%s] at key [%s], got [%s]", test.ac.expectA, keyExpect, cs[keyExpect])
			assert.Equalf(t,
				test.bc.expectB,
				cs[keyExpect2],
				"expected unioned config to have value [%s] at key [%s], got [%s]", test.bc.expectB, keyExpect2, cs[keyExpect2])
		})
	}
}
