package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/tester"
)

type CommonSlicesSuite struct {
	tester.Suite
}

func TestCommonSlicesSuite(t *testing.T) {
	s := &CommonSlicesSuite{Suite: tester.NewUnitSuite(t)}
	suite.Run(t, s)
}

func (suite *CommonSlicesSuite) TestContainsString() {
	t := suite.T()
	target := "fnords"
	good := []string{"fnords"}
	bad := []string{"foo", "bar"}

	assert.True(t, common.ContainsString(good, target))
	assert.False(t, common.ContainsString(bad, target))
}
