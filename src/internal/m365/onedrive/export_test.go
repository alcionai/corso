package onedrive

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	odConsts "github.com/alcionai/corso/src/internal/m365/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
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
			assert.Equal(suite.T(), test.isMeta, isMetadataFile(test.id, test.backupVersion), "is metadata")
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
		tname         string
		id            string
		backupVersion int
		name          string
		fin           data.FetchItemByNamer
		errFunc       assert.ErrorAssertionFunc
	}{
		{
			tname:         "legacy",
			id:            "name",
			backupVersion: version.OneDrive1DataAndMetaFiles,
			name:          "name",
			errFunc:       assert.NoError,
		},
		{
			tname:         "name in filename",
			id:            "name.data",
			backupVersion: version.OneDrive4DirIncludesPermissions,
			name:          "name",
			errFunc:       assert.NoError,
		},
		{
			tname:         "name in metadata",
			id:            "name.data",
			backupVersion: version.Backup,
			name:          "name",
			fin:           finD{id: "name.data", name: "name"},
			errFunc:       assert.NoError,
		},
		{
			tname:         "name in metadata but error",
			id:            "name.data",
			backupVersion: version.Backup,
			name:          "",
			fin:           finD{err: assert.AnError},
			errFunc:       assert.Error,
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
				test.backupVersion,
				test.fin,
			)
			test.errFunc(t, err)

			assert.Equal(t, test.name, name, "name")
		})
	}
}

type mockRestoreCollection struct {
	path  path.Path
	items []data.Stream
}

func (rc mockRestoreCollection) Items(ctx context.Context, errs *fault.Bus) <-chan data.Stream {
	ch := make(chan data.Stream)

	go func() {
		defer close(ch)

		for _, item := range rc.items {
			ch <- item
		}
	}()

	return ch
}

func (rc mockRestoreCollection) FullPath() path.Path {
	return rc.path
}

type mockDataStream struct {
	id   string
	data string
}

func (ms mockDataStream) ToReader() io.ReadCloser {
	if ms.data != "" {
		return io.NopCloser(bytes.NewBufferString(ms.data))
	}

	return nil
}
func (ms mockDataStream) UUID() string  { return ms.id }
func (ms mockDataStream) Deleted() bool { return false }

func (suite *ExportUnitSuite) TestGetItems() {
	table := []struct {
		name               string
		version            int
		backingCollections []data.RestoreCollection
		expectedItems      []export.ExportItem
	}{
		{
			name:    "single item",
			version: 1,
			backingCollections: []data.RestoreCollection{
				data.NoFetchRestoreCollection{
					Collection: mockRestoreCollection{
						items: []data.Stream{
							mockDataStream{id: "name1", data: "body1"},
						},
					},
				},
			},
			expectedItems: []export.ExportItem{
				{
					ID: "name1",
					Data: export.ExportItemData{
						Name: "name1",
						Body: io.NopCloser((bytes.NewBufferString("body1"))),
					},
				},
			},
		},
		{
			name:    "multiple items",
			version: 1,
			backingCollections: []data.RestoreCollection{
				data.NoFetchRestoreCollection{
					Collection: mockRestoreCollection{
						items: []data.Stream{
							mockDataStream{id: "name1", data: "body1"},
							mockDataStream{id: "name2", data: "body2"},
						},
					},
				},
			},
			expectedItems: []export.ExportItem{
				{
					ID: "name1",
					Data: export.ExportItemData{
						Name: "name1",
						Body: io.NopCloser((bytes.NewBufferString("body1"))),
					},
				},
				{
					ID: "name2",
					Data: export.ExportItemData{
						Name: "name2",
						Body: io.NopCloser((bytes.NewBufferString("body2"))),
					},
				},
			},
		},
		{
			name:    "single item with data suffix",
			version: 2,
			backingCollections: []data.RestoreCollection{
				data.NoFetchRestoreCollection{
					Collection: mockRestoreCollection{
						items: []data.Stream{
							mockDataStream{id: "name1.data", data: "body1"},
						},
					},
				},
			},
			expectedItems: []export.ExportItem{
				{
					ID: "name1.data",
					Data: export.ExportItemData{
						Name: "name1",
						Body: io.NopCloser((bytes.NewBufferString("body1"))),
					},
				},
			},
		},
		{
			name:    "single item name from metadata",
			version: version.Backup,
			backingCollections: []data.RestoreCollection{
				data.FetchRestoreCollection{
					Collection: mockRestoreCollection{
						items: []data.Stream{
							mockDataStream{id: "id1.data", data: "body1"},
						},
					},
					FetchItemByNamer: finD{id: "id1.data", name: "name1"},
				},
			},
			expectedItems: []export.ExportItem{
				{
					ID: "id1.data",
					Data: export.ExportItemData{
						Name: "name1",
						Body: io.NopCloser((bytes.NewBufferString("body1"))),
					},
				},
			},
		},
		{
			name:    "single item name from metadata with error",
			version: version.Backup,
			backingCollections: []data.RestoreCollection{
				data.FetchRestoreCollection{
					Collection: mockRestoreCollection{
						items: []data.Stream{
							mockDataStream{id: "id1.data"},
						},
					},
					FetchItemByNamer: finD{err: assert.AnError},
				},
			},
			expectedItems: []export.ExportItem{
				{
					ID:    "id1.data",
					Error: assert.AnError,
				},
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ec := exportCollection{
				baseDir:            "",
				backingCollections: test.backingCollections,
				version:            test.version,
			}

			items := ec.GetItems(ctx)

			fitems := []export.ExportItem{}
			for item := range items {
				fitems = append(fitems, item)
			}

			assert.Len(t, fitems, len(test.expectedItems), "num of items")

			for i, item := range fitems {
				assert.Equal(t, test.expectedItems[i].ID, item.ID, "id")
				assert.Equal(t, test.expectedItems[i].Data.Name, item.Data.Name, "name")
				assert.Equal(t, test.expectedItems[i].Data.Body, item.Data.Body, "body")

				if test.expectedItems[i].Error != nil {
					assert.Contains(t, item.Error.Error(), test.expectedItems[i].Error.Error(), "error")
				}
			}
		})
	}
}

func (suite *ExportUnitSuite) TestExportRestoreCollections() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	dpb := odConsts.DriveFolderPrefixBuilder("driveID1")

	p, err := dpb.ToDataLayerOneDrivePath("t", "u", false)
	assert.NoError(t, err, "build path")

	dcs := []data.RestoreCollection{
		data.FetchRestoreCollection{
			Collection: mockRestoreCollection{
				path: p,
				items: []data.Stream{
					mockDataStream{id: "id1.data", data: "body1"},
				},
			},
			FetchItemByNamer: finD{id: "id1.data", name: "name1"},
		},
	}

	expectedItems := []export.ExportItem{
		{
			ID: "id1.data",
			Data: export.ExportItemData{
				Name: "name1",
				Body: io.NopCloser((bytes.NewBufferString("body1"))),
			},
		},
	}

	exportCfg := control.ExportConfig{}
	ecs, err := ExportRestoreCollections(ctx, int(version.Backup), exportCfg, control.Options{}, dcs, nil, fault.New(true))
	assert.NoError(t, err, "export collections error")

	assert.Len(t, ecs, 1, "num of collections")

	items := ecs[0].GetItems(ctx)

	fitems := []export.ExportItem{}
	for item := range items {
		fitems = append(fitems, item)
	}

	assert.Equal(t, expectedItems, fitems, "items")
}
