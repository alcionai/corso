package vcf

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/converters/vcf/testdata"
	"github.com/alcionai/corso/src/internal/tester"
)

type VCFUnitSuite struct {
	tester.Suite
}

func TestVCFUnitSuite(t *testing.T) {
	suite.Run(t, &VCFUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *VCFUnitSuite) TestConvert_messageble_to_eml() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	body := []byte(testdata.ContactsInput)

	bytes, err := FromJSON(ctx, body)
	require.NoError(t, err, "convert")

	out := strings.ReplaceAll(string(bytes), "\r", "") // output contains \r
	assert.Equal(t, strings.TrimSpace(testdata.ContactsOutput), strings.TrimSpace(string(out)))
}
