package onedrive

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/collection/drive"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	odStub "github.com/alcionai/corso/src/internal/m365/service/onedrive/stub"
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

type mockRestoreCollection struct {
	path  path.Path
	items []*dataMock.Item
}

func (rc mockRestoreCollection) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	ch := make(chan data.Item)

	go func() {
		defer close(ch)

		el := errs.Local()

		for _, item := range rc.items {
			if item.ReadErr != nil {
				el.AddRecoverable(ctx, item.ReadErr)
				continue
			}

			ch <- item
		}
	}()

	return ch
}

func (rc mockRestoreCollection) FullPath() path.Path {
	return rc.path
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
				Collection: mockRestoreCollection{
					items: []*dataMock.Item{
						{
							ItemID: "name1",
							Reader: io.NopCloser(bytes.NewBufferString("body1")),
						},
					},
				},
			},
			expectedItems: []export.Item{
				{
					ID: "name1",
					Data: export.ItemData{
						Name: "name1",
						Body: io.NopCloser((bytes.NewBufferString("body1"))),
					},
				},
			},
		},
		{
			name:    "multiple items",
			version: 1,
			backingCollection: data.NoFetchRestoreCollection{
				Collection: mockRestoreCollection{
					items: []*dataMock.Item{
						{
							ItemID: "name1",
							Reader: io.NopCloser(bytes.NewBufferString("body1")),
						},
						{
							ItemID: "name2",
							Reader: io.NopCloser(bytes.NewBufferString("body2")),
						},
					},
				},
			},
			expectedItems: []export.Item{
				{
					ID: "name1",
					Data: export.ItemData{
						Name: "name1",
						Body: io.NopCloser((bytes.NewBufferString("body1"))),
					},
				},
				{
					ID: "name2",
					Data: export.ItemData{
						Name: "name2",
						Body: io.NopCloser((bytes.NewBufferString("body2"))),
					},
				},
			},
		},
		{
			name:    "single item with data suffix",
			version: 2,
			backingCollection: data.NoFetchRestoreCollection{
				Collection: mockRestoreCollection{
					items: []*dataMock.Item{
						{
							ItemID: "name1.data",
							Reader: io.NopCloser(bytes.NewBufferString("body1")),
						},
					},
				},
			},
			expectedItems: []export.Item{
				{
					ID: "name1.data",
					Data: export.ItemData{
						Name: "name1",
						Body: io.NopCloser((bytes.NewBufferString("body1"))),
					},
				},
			},
		},
		{
			name:    "single item name from metadata",
			version: version.Backup,
			backingCollection: data.FetchRestoreCollection{
				Collection: mockRestoreCollection{
					items: []*dataMock.Item{
						{
							ItemID: "id1.data",
							Reader: io.NopCloser(bytes.NewBufferString("body1")),
						},
					},
				},
				FetchItemByNamer: finD{id: "id1.meta", name: "name1"},
			},
			expectedItems: []export.Item{
				{
					ID: "id1.data",
					Data: export.ItemData{
						Name: "name1",
						Body: io.NopCloser((bytes.NewBufferString("body1"))),
					},
				},
			},
		},
		{
			name:    "single item name from metadata with error",
			version: version.Backup,
			backingCollection: data.FetchRestoreCollection{
				Collection: mockRestoreCollection{
					items: []*dataMock.Item{
						{ItemID: "id1.data"},
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
				Collection: mockRestoreCollection{
					items: []*dataMock.Item{
						{
							ItemID: "missing.data",
						},
						{
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
					ID: "id1.data",
					Data: export.ItemData{
						Name: "name1",
						Body: io.NopCloser(bytes.NewBufferString("body1")),
					},
				},
			},
		},
		{
			name:    "items with success and fetch error",
			version: version.OneDrive1DataAndMetaFiles,
			backingCollection: data.FetchRestoreCollection{
				Collection: mockRestoreCollection{
					items: []*dataMock.Item{
						{
							ItemID: "name0",
							Reader: io.NopCloser(bytes.NewBufferString("body0")),
						},
						{
							ItemID:  "name1",
							ReadErr: assert.AnError,
						},
						{
							ItemID: "name2",
							Reader: io.NopCloser(bytes.NewBufferString("body2")),
						},
					},
				},
			},
			expectedItems: []export.Item{
				{
					ID: "name0",
					Data: export.ItemData{
						Name: "name0",
						Body: io.NopCloser(bytes.NewBufferString("body0")),
					},
				},
				{
					ID: "name2",
					Data: export.ItemData{
						Name: "name2",
						Body: io.NopCloser(bytes.NewBufferString("body2")),
					},
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

			ec := drive.NewExportCollection("", test.backingCollection, test.version)

			items := ec.Items(ctx)

			fitems := []export.Item{}
			for item := range items {
				fitems = append(fitems, item)
			}

			assert.Len(t, fitems, len(test.expectedItems), "num of items")

			// We do not have any grantees about the ordering of the
			// items in the SDK, but leaving the test this way for now
			// to simplify testing.
			for i, item := range fitems {
				assert.Equal(t, test.expectedItems[i].ID, item.ID, "id")
				assert.Equal(t, test.expectedItems[i].Data.Name, item.Data.Name, "name")
				assert.Equal(t, test.expectedItems[i].Data.Body, item.Data.Body, "body")
				assert.ErrorIs(t, item.Error, test.expectedItems[i].Error)
			}
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
		dii           = odStub.DriveItemInfo()
		expectedItems = []export.Item{
			{
				ID: "id1.data",
				Data: export.ItemData{
					Name: "name1",
					Body: io.NopCloser((bytes.NewBufferString("body1"))),
				},
			},
		}
	)

	dii.OneDrive.ItemName = "name1"

	p, err := dpb.ToDataLayerOneDrivePath("t", "u", false)
	assert.NoError(t, err, "build path")

	dcs := []data.RestoreCollection{
		data.FetchRestoreCollection{
			Collection: mockRestoreCollection{
				path: p,
				items: []*dataMock.Item{
					{
						ItemID:   "id1.data",
						Reader:   io.NopCloser(bytes.NewBufferString("body1")),
						ItemInfo: dii,
					},
				},
			},
			FetchItemByNamer: finD{id: "id1.meta", name: "name1"},
		},
	}

	ecs, err := ProduceExportCollections(
		ctx,
		int(version.Backup),
		exportCfg,
		control.DefaultOptions(),
		dcs,
		nil,
		fault.New(true))
	assert.NoError(t, err, "export collections error")
	assert.Len(t, ecs, 1, "num of collections")

	fitems := []export.Item{}
	for item := range ecs[0].Items(ctx) {
		fitems = append(fitems, item)
	}

	assert.Equal(t, expectedItems, fitems, "items")
}
