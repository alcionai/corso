package eml

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/converters/eml/testdata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type EMLUnitSuite struct {
	tester.Suite
}

func TestEMLUnitSuite(t *testing.T) {
	suite.Run(t, &EMLUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *EMLUnitSuite) TestConvert_messageble_to_eml() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	msg, err := api.BytesToMessageable([]byte(testdata.EmailWithAttachments))
	require.NoError(t, err, "creating message")

	_, err = ToEml(ctx, msg)
	// TODO(meain): add more tests on the generated content
	// Cannot test output directly as it contains a random boundary
	assert.NoError(t, err, "converting to eml")
}
