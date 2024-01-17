package sharepoint

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
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

func (suite *ExportUnitSuite) TestExportRestoreCollections() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		driveID   = "driveID1"
		driveName = "driveName1"
		exportCfg = control.ExportConfig{}
		dpb       = odConsts.DriveFolderPrefixBuilder(driveID)
	)

	table := []struct {
		name          string
		itemName      string
		itemID        string
		itemInfo      details.ItemInfo
		getCollPath   func(t *testing.T) path.Path
		statsCat      path.CategoryType
		expectedItems []export.Item
		expectedPath  string
	}{
		{
			name:     "OneDriveLegacyItemInfo",
			itemName: "name1",
			itemID:   "id1.data",
			itemInfo: details.ItemInfo{
				OneDrive: &details.OneDriveInfo{
					ItemType:  details.OneDriveItem,
					ItemName:  "name1",
					Size:      1,
					DriveName: driveName,
					DriveID:   driveID,
				},
			},
			getCollPath: func(t *testing.T) path.Path {
				p, err := path.Build(
					"t",
					"u",
					path.SharePointService,
					path.LibrariesCategory,
					false,
					dpb.Elements()...)
				assert.NoError(t, err, "build path")

				return p
			},
			statsCat:     path.FilesCategory,
			expectedPath: path.LibrariesCategory.HumanString() + "/" + driveName,
			expectedItems: []export.Item{
				{
					ID:   "id1.data",
					Name: "name1",
					Body: io.NopCloser((bytes.NewBufferString("body1"))),
				},
			},
		},
		{
			name:     "SharePointItemInfo, Libraries Category",
			itemName: "name1",
			itemID:   "id1.data",
			itemInfo: details.ItemInfo{
				SharePoint: &details.SharePointInfo{
					ItemType:  details.SharePointLibrary,
					ItemName:  "name1",
					Size:      1,
					DriveName: driveName,
					DriveID:   driveID,
				},
			},
			getCollPath: func(t *testing.T) path.Path {
				p, err := path.Build(
					"t",
					"u",
					path.SharePointService,
					path.LibrariesCategory,
					false,
					dpb.Elements()...)
				assert.NoError(t, err, "build path")

				return p
			},
			statsCat:     path.FilesCategory,
			expectedPath: path.LibrariesCategory.HumanString() + "/" + driveName,
			expectedItems: []export.Item{
				{
					ID:   "id1.data",
					Name: "name1",
					Body: io.NopCloser((bytes.NewBufferString("body1"))),
				},
			},
		},
		{
			name:     "SharePointItemInfo, Lists Category",
			itemName: "list1",
			itemID:   "listid1",
			itemInfo: details.ItemInfo{
				SharePoint: &details.SharePointInfo{
					ItemType: details.SharePointList,
					List: &details.ListInfo{
						Name:      "list1",
						ItemCount: 10,
					},
				},
			},
			getCollPath: func(t *testing.T) path.Path {
				p, err := path.Build(
					"t",
					"u",
					path.SharePointService,
					path.ListsCategory,
					false,
					"listid1")
				assert.NoError(t, err, "build path")

				return p
			},
			statsCat:     path.ListsCategory,
			expectedPath: path.ListsCategory.HumanString() + "/listid1",
			expectedItems: []export.Item{
				{
					ID:   "listid1",
					Name: "listid1.json",
					Body: io.NopCloser((bytes.NewBufferString("body1"))),
				},
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			dcs := []data.RestoreCollection{
				data.FetchRestoreCollection{
					Collection: dataMock.Collection{
						Path: test.getCollPath(t),
						ItemData: []data.Item{
							&dataMock.Item{
								ItemID: test.itemID,
								Reader: io.NopCloser(bytes.NewBufferString("body1")),
							},
						},
					},
					FetchItemByNamer: finD{id: "id1.meta", name: test.itemName},
				},
			}

			handler := NewSharePointHandler(api.Client{}, nil)
			handler.CacheItemInfo(test.itemInfo)

			stats := metrics.NewExportStats()

			ecs, err := handler.ProduceExportCollections(
				ctx,
				int(version.Backup),
				exportCfg,
				dcs,
				stats,
				fault.New(true))
			require.NoError(t, err, "export collections error")
			assert.Len(t, ecs, 1, "num of collections")
			assert.Equal(t, test.expectedPath, ecs[0].BasePath(), "base dir")

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

			assert.Equal(t, test.expectedItems, fitems, "items")

			expectedStats := metrics.NewExportStats()
			expectedStats.UpdateBytes(test.statsCat, int64(size))
			expectedStats.UpdateResourceCount(test.statsCat)
			assert.Equal(t, expectedStats.GetStats(), stats.GetStats(), "stats")
		})
	}
}
