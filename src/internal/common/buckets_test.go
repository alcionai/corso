package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
)

type CommonBucketsSuite struct {
	suite.Suite
}

func TestCommonBucketsSuite(t *testing.T) {
	suite.Run(t, new(CommonBucketsSuite))
}

func (suite *CommonBucketsSuite) TestDoesThings() {
	t := suite.T()
	trimmablePrefixes := []string{"s3://"}

	for _, pfx := range trimmablePrefixes {
		assert.Equal(t, "fnords", common.NormalizeBucket(pfx+"fnords"))
		assert.Equal(t, "smarf", "smarf")
	}
}
