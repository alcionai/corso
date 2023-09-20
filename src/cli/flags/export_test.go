package flags

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
)

type ExportUnitSuite struct {
	tester.Suite
}

func TestExportUnitSuite(t *testing.T) {
	suite.Run(t, &ExportUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExportUnitSuite) TestValidateExportConfigFlags() {
	t := suite.T()

	FormatFV = ""

	err := ValidateExportConfigFlags()
	assert.NoError(t, err, clues.ToCore(err))

	FormatFV = "json"

	err = ValidateExportConfigFlags()
	assert.NoError(t, err, clues.ToCore(err))

	FormatFV = "JsoN"

	err = ValidateExportConfigFlags()
	assert.NoError(t, err, clues.ToCore(err))

	FormatFV = "fnerds"

	err = ValidateExportConfigFlags()
	assert.Error(t, err, clues.ToCore(err))
}
