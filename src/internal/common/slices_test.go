package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
)

type CommonSlicesSuite struct {
	suite.Suite
}

func TestCommonSlicesSuite(t *testing.T) {
	suite.Run(t, new(CommonSlicesSuite))
}

func (suite *CommonSlicesSuite) TestContainsString() {
	t := suite.T()
	target := "fnords"
	good := []string{"fnords"}
	bad := []string{"foo", "bar"}

	assert.True(t, common.ContainsString(good, target))
	assert.False(t, common.ContainsString(bad, target))
}
