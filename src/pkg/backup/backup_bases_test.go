package backup

import (
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

type BackupBasesUnitSuite struct {
	tester.Suite
}

func TestBackupBasesUnitSuite(t *testing.T) {
	suite.Run(t, &BackupBasesUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *BackupBasesUnitSuite) TestReasonSerialization() {
	table := []struct {
		name           string
		input          string
		expectErr      assert.ErrorAssertionFunc
		expectService  path.ServiceType
		expectCategory path.CategoryType
	}{
		{
			name:           "ProperFormat",
			input:          serviceCatString(path.ExchangeService, path.EmailCategory),
			expectErr:      assert.NoError,
			expectService:  path.ExchangeService,
			expectCategory: path.EmailCategory,
		},
		{
			name: "MissingPrefix",
			input: strings.TrimPrefix(
				serviceCatString(path.ExchangeService, path.EmailCategory),
				serviceCatPrefix),
			expectErr: assert.Error,
		},
		{
			name: "MissingSeparator",
			input: strings.ReplaceAll(
				serviceCatString(path.ExchangeService, path.EmailCategory),
				separator,
				""),
			expectErr: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			service, cat, err := serviceCatStringToTypes(test.input)
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, test.expectService, service)
			assert.Equal(t, test.expectCategory, cat)
		})
	}
}
