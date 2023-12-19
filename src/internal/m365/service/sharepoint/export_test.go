package sharepoint

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
		driveID       = "driveID1"
		driveName     = "driveName1"
		itemName      = "name1"
		exportCfg     = control.ExportConfig{}
		dpb           = odConsts.DriveFolderPrefixBuilder(driveID)
		expectedPath  = path.LibrariesCategory.HumanString() + "/" + driveName
		expectedItems = []export.Item{
			{
				ID:   "id1.data",
				Name: itemName,
				Body: io.NopCloser((bytes.NewBufferString("body1"))),
			},
		}
	)

	p, err := dpb.ToDataLayerSharePointPath("t", "u", path.LibrariesCategory, false)
	assert.NoError(t, err, "build path")

	table := []struct {
		name     string
		itemInfo details.ItemInfo
	}{
		{
			name: "OneDriveLegacyItemInfo",
			itemInfo: details.ItemInfo{
				OneDrive: &details.OneDriveInfo{
					ItemType:  details.OneDriveItem,
					ItemName:  itemName,
					Size:      1,
					DriveName: driveName,
					DriveID:   driveID,
				},
			},
		},
		{
			name: "SharePointItemInfo",
			itemInfo: details.ItemInfo{
				SharePoint: &details.SharePointInfo{
					ItemType:  details.SharePointLibrary,
					ItemName:  itemName,
					Size:      1,
					DriveName: driveName,
					DriveID:   driveID,
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
						Path: p,
						ItemData: []data.Item{
							&dataMock.Item{
								ItemID: "id1.data",
								Reader: io.NopCloser(bytes.NewBufferString("body1")),
							},
						},
					},
					FetchItemByNamer: finD{id: "id1.meta", name: itemName},
				},
			}

			handler := NewSharePointHandler(control.DefaultOptions(), api.Client{}, nil)
			handler.CacheItemInfo(test.itemInfo)

			stats := metrics.ExportStats{}

			ecs, err := handler.ProduceExportCollections(
				ctx,
				int(version.Backup),
				exportCfg,
				dcs,
				&stats,
				fault.New(true))
			assert.NoError(t, err, "export collections error")
			assert.Len(t, ecs, 1, "num of collections")

			assert.Equal(t, expectedPath, ecs[0].BasePath(), "base dir")

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
		})
	}
}
