package control_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/pkg/control"
)

type OptionsSuite struct {
	suite.Suite
}

func TestOptionsSuite(t *testing.T) {
	suite.Run(t, new(OptionsSuite))
}

func (suite *OptionsSuite) TestNewOptions() {
	t := suite.T()

	o1 := control.NewOptions(true)
	assert.True(t, o1.FailFast, "failFast")

	o2 := control.NewOptions(false)
	assert.False(t, o2.FailFast, "failFast")
}
