package onedrive

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/alcionai/corso/src/internal/data"
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

type metadataStream struct {
	id   string
	name string
}

func (ms metadataStream) ToReader() io.ReadCloser {
	return io.NopCloser(bytes.NewBufferString(`{"filename": "` + ms.name + `"}`))
}
func (ms metadataStream) UUID() string  { return ms.id }
func (ms metadataStream) Deleted() bool { return false }

type finD struct {
	id   string
	name string
	err  error
}

func (fd finD) FetchItemByName(ctx context.Context, name string) (data.Stream, error) {
	if fd.err != nil {
		return nil, fd.err
	}

	return metadataStream{id: fd.id, name: fd.name}, nil
}

func (suite *ExportUnitSuite) TestGetItemName() {
	table := []struct {
		tname   string
		id      string
		version int
		name    string
		fin     data.FetchItemByNamer
		errFunc assert.ErrorAssertionFunc
	}{
		{
			tname:   "legacy",
			id:      "name",
			version: 1,
			name:    "name",
			errFunc: assert.NoError,
		},
		{
			tname:   "name in filename",
			id:      "name.data",
			version: 4,
			name:    "name",
			errFunc: assert.NoError,
		},
		{
			tname:   "name in metadata",
			id:      "name.data",
			version: 5,
			name:    "name",
			fin:     finD{id: "name.data", name: "name"},
			errFunc: assert.NoError,
		},
		{
			tname:   "name in metadata but error",
			id:      "name.data",
			version: 5,
			name:    "",
			fin:     finD{err: assert.AnError},
			errFunc: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.tname, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			name, err := getItemName(
				ctx,
				test.id,
				test.version,
				test.fin,
			)
			test.errFunc(t, err)

			assert.Equal(t, test.name, name, "name")
		})
	}
}
