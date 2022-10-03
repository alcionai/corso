package common_test

import (
	"context"
	"testing"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CommonBucketsSuite struct {
	suite.Suite
	ctx context.Context
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
