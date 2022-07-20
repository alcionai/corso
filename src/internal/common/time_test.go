package common_test

import (
	"testing"
	"time"

	"github.com/alcionai/corso/internal/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CommonTimeUnitSuite struct {
	suite.Suite
}

func TestCommonTimeUnitSuite(t *testing.T) {
	suite.Run(t, new(CommonTimeUnitSuite))
}

func (suite *CommonTimeUnitSuite) TestFormatTime() {
	t := suite.T()
	now := time.Now()
	result := common.FormatTime(now)
	assert.Equal(t, now.UTC().Format(time.RFC3339Nano), result)
}

func (suite *CommonTimeUnitSuite) TestParseTime() {
	t := suite.T()
	now := time.Now()
	nowStr := now.Format(time.RFC3339Nano)
	result := common.ParseTime(nowStr)
	assert.Equal(t, now.UTC(), result)
}
