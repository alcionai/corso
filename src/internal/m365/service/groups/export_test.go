package groups

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
	groupMock "github.com/alcionai/corso/src/internal/m365/service/groups/mock"
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
	key  string
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
			Reader: io.NopCloser(bytes.NewBufferString(`{"` + fd.key + `": "` + fd.name + `"}`)),
		}, nil
	}

	return nil, assert.AnError
}

func (suite *ExportUnitSuite) TestExportRestoreCollections_messages() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		itemID        = "itemID"
		containerName = "channelID"
		dii           = groupMock.ItemInfo()
		content       = `{"displayname": "` + dii.Groups.ItemName + `"}`
		body          = io.NopCloser(bytes.NewBufferString(content))
		exportCfg     = control.ExportConfig{}
		expectedPath  = path.ChannelMessagesCategory.HumanString() + "/" + containerName
		expectedItems = []export.Item{
			{
				ID:   itemID,
				Name: dii.Groups.ItemName + ".json",
				// Body: body, not checked
			},
		}
	)

	p, err := path.Build("t", "pr", path.GroupsService, path.ChannelMessagesCategory, false, containerName)
	assert.NoError(t, err, "build path")

	dcs := []data.RestoreCollection{
		data.FetchRestoreCollection{
			Collection: dataMock.Collection{
				Path: p,
				ItemData: []data.Item{
					&dataMock.Item{
						ItemID: itemID,
						Reader: body,
					},
				},
			},
			FetchItemByNamer: finD{id: itemID, key: "displayname", name: dii.Groups.ItemName},
		},
	}

	stats := metrics.NewExportStats()

	ecs, err := NewGroupsHandler(api.Client{}, nil).
		ProduceExportCollections(
			ctx,
			int(version.Backup),
			exportCfg,
			dcs,
			stats,
			fault.New(true))
	assert.NoError(t, err, "export collections error")
	assert.Len(t, ecs, 1, "num of collections")

	assert.Equal(t, expectedPath, ecs[0].BasePath(), "base dir")

	fitems := []export.Item{}

	size := 0

	for item := range ecs[0].Items(ctx) {
		b, err := io.ReadAll(item.Body)
		assert.NoError(t, err, clues.ToCore(err))

		// count up size for tests
		size += len(b)

		// have to nil out body, otherwise assert fails due to
		// pointer memory location differences
		item.Body = nil
		fitems = append(fitems, item)
	}

	assert.Equal(t, expectedItems, fitems, "items")

	expectedStats := metrics.NewExportStats()
	expectedStats.UpdateBytes(path.ChannelMessagesCategory, int64(size))
	expectedStats.UpdateResourceCount(path.ChannelMessagesCategory)
	assert.Equal(t, expectedStats.GetStats(), stats.GetStats(), "stats")
}

func (suite *ExportUnitSuite) TestExportRestoreCollections_libraries() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		siteID          = "siteID1"
		siteEscapedName = "siteName1"
		siteWebURL      = "https://site1.sharepoint.com/sites/" + siteEscapedName
		driveID         = "driveID1"
		driveName       = "driveName1"
		exportCfg       = control.ExportConfig{}
		dpb             = odConsts.DriveFolderPrefixBuilder(driveID)

		dii = details.ItemInfo{
			Groups: &details.GroupsInfo{
				ItemType:  details.SharePointLibrary,
				ItemName:  "name1",
				Size:      1,
				DriveName: driveName,
				DriveID:   driveID,
				SiteID:    siteID,
				WebURL:    siteWebURL,
			},
		}

		expectedPath  = "Libraries/" + siteEscapedName + "/" + driveName
		expectedItems = []export.Item{
			{
				ID:   "id1.data",
				Name: "name1",
				Body: io.NopCloser((bytes.NewBufferString("body1"))),
			},
		}
	)

	p, err := dpb.ToDataLayerPath(
		"t",
		"u",
		path.GroupsService,
		path.LibrariesCategory,
		false,
		odConsts.SitesPathDir,
		siteID)
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
			FetchItemByNamer: finD{id: "id1.meta", key: "filename", name: "name1"},
		},
	}

	handler := NewGroupsHandler(api.Client{}, nil)
	handler.CacheItemInfo(dii)

	stats := metrics.NewExportStats()

	ecs, err := handler.ProduceExportCollections(
		ctx,
		int(version.Backup),
		exportCfg,
		dcs,
		stats,
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

	expectedStats := metrics.NewExportStats()
	expectedStats.UpdateBytes(path.FilesCategory, int64(size))
	expectedStats.UpdateResourceCount(path.FilesCategory)
	assert.Equal(t, expectedStats.GetStats(), stats.GetStats(), "stats")
}
