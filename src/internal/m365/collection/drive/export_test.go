package drive

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
)

type ExportUnitSuite struct {
	tester.Suite
}

func TestExportUnitSuite(t *testing.T) {
	suite.Run(t, &ExportUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExportUnitSuite) TestIsMetadataFile() {
	table := []struct {
		name          string
		id            string
		backupVersion int
		isMeta        bool
	}{
		{
			name:          "legacy",
			backupVersion: version.OneDrive1DataAndMetaFiles,
			isMeta:        false,
		},
		{
			name:          "metadata file",
			backupVersion: version.OneDrive3IsMetaMarker,
			id:            "name" + metadata.MetaFileSuffix,
			isMeta:        true,
		},
		{
			name:          "dir metadata file",
			backupVersion: version.OneDrive3IsMetaMarker,
			id:            "name" + metadata.DirMetaFileSuffix,
			isMeta:        true,
		},
		{
			name:          "non metadata file",
			backupVersion: version.OneDrive3IsMetaMarker,
			id:            "name" + metadata.DataFileSuffix,
			isMeta:        false,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			assert.Equal(
				suite.T(),
				test.isMeta,
				isMetadataFile(test.id, test.backupVersion),
				"is metadata")
		})
	}
}

type finD struct {
	id   string
	name string
	err  error
}

func (fd finD) FetchItemByName(ctx context.Context, name string) (data.Item, error) {
	if fd.err != nil {
		return nil, fd.err
	}

	if name == fd.id {
		return &dataMock.Item{
			ItemID: fd.id,
			Reader: io.NopCloser(bytes.NewBufferString(`{"filename": "` + fd.name + `"}`)),
		}, nil
	}

	return nil, assert.AnError
}

func (suite *ExportUnitSuite) TestGetItemName() {
	table := []struct {
		name          string
		id            string
		backupVersion int
		expectName    string
		fin           data.FetchItemByNamer
		expectErr     assert.ErrorAssertionFunc
	}{
		{
			name:          "legacy",
			id:            "name",
			backupVersion: version.OneDrive1DataAndMetaFiles,
			expectName:    "name",
			expectErr:     assert.NoError,
		},
		{
			name:          "name in filename",
			id:            "name.data",
			backupVersion: version.OneDrive4DirIncludesPermissions,
			expectName:    "name",
			expectErr:     assert.NoError,
		},
		{
			name:          "name in metadata",
			id:            "id.data",
			backupVersion: version.Backup,
			expectName:    "name",
			fin:           finD{id: "id.meta", name: "name"},
			expectErr:     assert.NoError,
		},
		{
			name:          "name in metadata but error",
			id:            "id.data",
			backupVersion: version.Backup,
			expectName:    "",
			fin:           finD{err: assert.AnError},
			expectErr:     assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			name, err := getItemName(
				ctx,
				test.id,
				test.backupVersion,
				test.fin)
			test.expectErr(t, err, clues.ToCore(err))

			assert.Equal(t, test.expectName, name, "name")
		})
	}
}
