package onedrive

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
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/metrics"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type ExportUnitSuite struct {
	tester.Suite
}

func TestExportUnitSuite(t *testing.T) {
	suite.Run(t, &ExportUnitSuite{Suite: tester.NewUnitSuite(t)})
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

func (suite *ExportUnitSuite) TestGetItems() {
	table := []struct {
		name              string
		version           int
		backingCollection data.RestoreCollection
		expectedItems     []export.Item
	}{
		{
			name:    "single item",
			version: 1,
			backingCollection: data.NoFetchRestoreCollection{
				Collection: dataMock.Collection{
					ItemData: []data.Item{
						&dataMock.Item{
							ItemID: "name1",
							Reader: io.NopCloser(bytes.NewBufferString("body1")),
						},
					},
				},
			},
			expectedItems: []export.Item{
				{
					ID:   "name1",
					Name: "name1",
					Body: io.NopCloser((bytes.NewBufferString("body1"))),
				},
			},
		},
		{
			name:    "multiple items",
			version: 1,
			backingCollection: data.NoFetchRestoreCollection{
				Collection: dataMock.Collection{
					ItemData: []data.Item{
						&dataMock.Item{
							ItemID: "name1",
							Reader: io.NopCloser(bytes.NewBufferString("body1")),
						},
						&dataMock.Item{
							ItemID: "name2",
							Reader: io.NopCloser(bytes.NewBufferString("body2")),
						},
					},
				},
			},
			expectedItems: []export.Item{
				{
					ID:   "name1",
					Name: "name1",
					Body: io.NopCloser((bytes.NewBufferString("body1"))),
				},
				{
					ID:   "name2",
					Name: "name2",
					Body: io.NopCloser((bytes.NewBufferString("body2"))),
				},
			},
		},
		{
			name:    "single item with data suffix",
			version: 2,
			backingCollection: data.NoFetchRestoreCollection{
				Collection: dataMock.Collection{
					ItemData: []data.Item{
						&dataMock.Item{
							ItemID: "name1.data",
							Reader: io.NopCloser(bytes.NewBufferString("body1")),
						},
					},
				},
			},
			expectedItems: []export.Item{
				{
					ID:   "name1.data",
					Name: "name1",
					Body: io.NopCloser((bytes.NewBufferString("body1"))),
				},
			},
		},
		{
			name:    "single item name from metadata",
			version: version.Backup,
			backingCollection: data.FetchRestoreCollection{
				Collection: dataMock.Collection{
					ItemData: []data.Item{
						&dataMock.Item{
							ItemID: "id1.data",
							Reader: io.NopCloser(bytes.NewBufferString("body1")),
						},
					},
				},
				FetchItemByNamer: finD{id: "id1.meta", name: "name1"},
			},
			expectedItems: []export.Item{
				{
					ID:   "id1.data",
					Name: "name1",
					Body: io.NopCloser((bytes.NewBufferString("body1"))),
				},
			},
		},
		{
			name:    "single item name from metadata with error",
			version: version.Backup,
			backingCollection: data.FetchRestoreCollection{
				Collection: dataMock.Collection{
					ItemData: []data.Item{
						&dataMock.Item{ItemID: "id1.data"},
					},
				},
				FetchItemByNamer: finD{err: assert.AnError},
			},
			expectedItems: []export.Item{
				{
					ID:    "id1.data",
					Error: assert.AnError,
				},
			},
		},
		{
			name:    "items with success and metadata read error",
			version: version.Backup,
			backingCollection: data.FetchRestoreCollection{
				Collection: dataMock.Collection{
					ItemData: []data.Item{
						&dataMock.Item{
							ItemID: "missing.data",
						},
						&dataMock.Item{
							ItemID: "id1.data",
							Reader: io.NopCloser(bytes.NewBufferString("body1")),
						},
					},
				},
				FetchItemByNamer: finD{id: "id1.meta", name: "name1"},
			},
			expectedItems: []export.Item{
				{
					ID:    "missing.data",
					Error: assert.AnError,
				},
				{
					ID:   "id1.data",
					Name: "name1",
					Body: io.NopCloser(bytes.NewBufferString("body1")),
				},
			},
		},
		{
			name:    "items with success and fetch error",
			version: version.OneDrive1DataAndMetaFiles,
			backingCollection: data.FetchRestoreCollection{
				Collection: dataMock.Collection{
					ItemData: []data.Item{
						&dataMock.Item{
							ItemID: "name0",
							Reader: io.NopCloser(bytes.NewBufferString("body0")),
						},
						&dataMock.Item{
							ItemID:  "name1",
							ReadErr: assert.AnError,
						},
						&dataMock.Item{
							ItemID: "name2",
							Reader: io.NopCloser(bytes.NewBufferString("body2")),
						},
					},
				},
			},
			expectedItems: []export.Item{
				{
					ID:   "name0",
					Name: "name0",
					Body: io.NopCloser(bytes.NewBufferString("body0")),
				},
				{
					ID:   "name2",
					Name: "name2",
					Body: io.NopCloser(bytes.NewBufferString("body2")),
				},
				{
					ID:    "",
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

			stats := metrics.ExportStats{}
			ec := drive.NewExportCollection(
				"",
				[]data.RestoreCollection{test.backingCollection},
				test.version,
				&stats)

			items := ec.Items(ctx)

			count := 0
			size := 0
			fitems := []export.Item{}

			for item := range items {
				if item.Error == nil {
					count++
				}

				if item.Body != nil {
					b, err := io.ReadAll(item.Body)
					assert.NoError(t, err, clues.ToCore(err))

					size += len(b)
					item.Body = io.NopCloser(bytes.NewBuffer(b))
				}

				fitems = append(fitems, item)
			}

			assert.Len(t, fitems, len(test.expectedItems), "num of items")

			// We do not have any grantees about the ordering of the
			// items in the SDK, but leaving the test this way for now
			// to simplify testing.
			for i, item := range fitems {
				assert.Equal(t, test.expectedItems[i].ID, item.ID, "id")
				assert.Equal(t, test.expectedItems[i].Name, item.Name, "name")
				assert.Equal(t, test.expectedItems[i].Body, item.Body, "body")
				assert.ErrorIs(t, item.Error, test.expectedItems[i].Error)
			}

			var expectedStats metrics.ExportStats

			if size+count > 0 { // it is only initialized if we have something
				expectedStats = metrics.ExportStats{}
				expectedStats.UpdateBytes(path.FilesCategory, int64(size))

				for i := 0; i < count; i++ {
					expectedStats.UpdateResourceCount(path.FilesCategory)
				}
			}

			assert.Equal(t, expectedStats, stats, "stats")
		})
	}
}

func (suite *ExportUnitSuite) TestExportRestoreCollections() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		exportCfg     = control.ExportConfig{}
		dpb           = odConsts.DriveFolderPrefixBuilder("driveID1")
		expectedItems = []export.Item{
			{
				ID:   "id1.data",
				Name: "name1",
				Body: io.NopCloser((bytes.NewBufferString("body1"))),
			},
		}
	)

	p, err := dpb.ToDataLayerOneDrivePath("t", "u", false)
	assert.NoError(t, err, "build path")

	dcs := []data.RestoreCollection{
		data.FetchRestoreCollection{
			Collection: dataMock.Collection{
				Path: p,
				ItemData: []data.Item{
					&dataMock.Item{
						ItemID: "id1.data",
						Reader: io.NopCloser(bytes.NewBufferString("body1")),
					},
				},
			},
			FetchItemByNamer: finD{id: "id1.meta", name: "name1"},
		},
	}

	stats := metrics.ExportStats{}

	ecs, err := NewOneDriveHandler(control.DefaultOptions(), api.Client{}, nil).
		ProduceExportCollections(
			ctx,
			int(version.Backup),
			exportCfg,
			dcs,
			&stats,
			fault.New(true))
	assert.NoError(t, err, "export collections error")
	assert.Len(t, ecs, 1, "num of collections")

	fitems := []export.Item{}
	size := 0

	for item := range ecs[0].Items(ctx) {
		// unwrap the body from stats reader
		b, err := io.ReadAll(item.Body)
		assert.NoError(t, err, clues.ToCore(err))

		size += len(b)
		bitem := io.NopCloser(bytes.NewBuffer(b))
		item.Body = bitem

		fitems = append(fitems, item)
	}

	assert.Equal(t, expectedItems, fitems, "items")

	expectedStats := metrics.ExportStats{}
	expectedStats.UpdateBytes(path.FilesCategory, int64(size))
	expectedStats.UpdateResourceCount(path.FilesCategory)
	assert.Equal(t, expectedStats, stats, "stats")
}
