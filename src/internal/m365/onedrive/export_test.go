package onedrive

import (
	"testing"

	"github.com/alcionai/corso/src/internal/m365/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExportUnitSuite struct {
	tester.Suite
}

func TestExportUnitSuite(t *testing.T) {
	suite.Run(t, &ExportUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExportUnitSuite) TestIsMetadataFile() {
	table := []struct {
		name    string
		id      string
		version int
		isMeta  bool
	}{
		{
			name:    "legacy",
			version: 1,
			isMeta:  false,
		},
		{
			name:    "metadata file",
			version: 2,
			id:      "name" + metadata.MetaFileSuffix,
			isMeta:  true,
		},
		{
			name:    "dir metadata file",
			version: 2,
			id:      "name" + metadata.DirMetaFileSuffix,
			isMeta:  true,
		},
		{
			name:    "non metadata file",
			version: 2,
			id:      "name" + metadata.DataFileSuffix,
			isMeta:  false,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			assert.Equal(suite.T(), test.isMeta, isMetadataFile(test.id, test.version), "is metadata")
		})
	}
}
