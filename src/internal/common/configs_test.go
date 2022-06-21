package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/common"
)

type CommonConfigsSuite struct {
	suite.Suite
}

func TestCommonConfigsSuite(t *testing.T) {
	suite.Run(t, new(CommonConfigsSuite))
}

const (
	keyExpect  = "expect"
	keyExpect2 = "expect2"
)

type stringConfig struct {
	expectA string
	err     error
}

func (c stringConfig) Config() (common.Config[string], error) {
	return common.Config[string]{keyExpect: c.expectA}, c.err
}

type stringConfig2 struct {
	expectB string
	err     error
}

func (c stringConfig2) Config() (common.Config[string], error) {
	return common.Config[string]{keyExpect2: c.expectB}, c.err
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
		suite.T().Run(test.name, func(t *testing.T) {
			cs, err := common.UnionConfigs[string, common.Config[string]](test.ac, test.bc)
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

type anyConfig struct {
	expectA any
	err     error
}

func (c anyConfig) Config() (common.Config[any], error) {
	return common.Config[any]{keyExpect: c.expectA}, c.err
}

type anyConfig2 struct {
	expectB any
	err     error
}

func (c anyConfig2) Config() (common.Config[any], error) {
	return common.Config[any]{keyExpect2: c.expectB}, c.err
}

func (suite *CommonConfigsSuite) TestUnionConfigs_any() {
	table := []struct {
		name     string
		ac       anyConfig
		bc       anyConfig2
		errCheck assert.ErrorAssertionFunc
	}{
		{"no error", anyConfig{1, nil}, anyConfig2{2, nil}, assert.NoError},
		{"tc error", anyConfig{1, assert.AnError}, anyConfig2{2, nil}, assert.Error},
		{"fc error", anyConfig{1, nil}, anyConfig2{2, assert.AnError}, assert.Error},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			cs, err := common.UnionConfigs[any, common.Config[any]](test.ac, test.bc)
			test.errCheck(t, err)
			// remaining tests depend on error-free state
			if test.ac.err != nil || test.bc.err != nil {
				return
			}
			assert.Equalf(t,
				test.ac.expectA,
				cs[keyExpect],
				"expected unioned config to have value [%v] at key [%s], got [%v]", test.ac.expectA, keyExpect, cs[keyExpect])
			assert.Equalf(t,
				test.bc.expectB,
				cs[keyExpect2],
				"expected unioned config to have value [%v] at key [%s], got [%v]", test.bc.expectB, keyExpect2, cs[keyExpect2])
		})
	}
}
