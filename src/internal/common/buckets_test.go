package common_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/tester"
)

type CommonBucketsSuite struct {
	tester.Suite
}

func TestCommonBucketsSuite(t *testing.T) {
	s := &CommonBucketsSuite{Suite: tester.NewUnitSuite(t)}
	suite.Run(t, s)
}

func (suite *CommonBucketsSuite) TestBucketPrefix() {
	t := suite.T()
	trimmablePrefixes := []string{"s3://"}

	for _, pfx := range trimmablePrefixes {
		assert.Equal(t, "fnords", common.NormalizeBucket(pfx+"fnords"))
		assert.Equal(t, "smarf", "smarf")
	}
}

func (suite *CommonBucketsSuite) TestPrefixSuffix() {
	t := suite.T()

	prefixBase := "repo-prefix"
	properPrefix := prefixBase + "/"

	assert.Equal(t, properPrefix, common.NormalizePrefix(prefixBase), "Trailing '/' should be added")
	assert.Equal(t, properPrefix, common.NormalizePrefix(properPrefix), "Properly formatted prefix should not change")
	assert.Equal(t, properPrefix, common.NormalizePrefix(prefixBase+"///"), "Only one trailing / should exist")
	assert.Equal(t, properPrefix+"/sub/", common.NormalizePrefix(properPrefix+"/sub"), "Only affect trailing /")
	assert.Equal(t, "", common.NormalizePrefix(""), "Only normalize actual prefix.")
	assert.Equal(t, "", common.NormalizePrefix("//"), "Only normalize actual prefix.")
}
