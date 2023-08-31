package sharepoint

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/idname"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
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

func (suite *ExportUnitSuite) TestExportRestoreCollections() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		driveID   = "driveID1"
		driveName = "driveName1"
		exportCfg = control.ExportConfig{}
		dpb       = odConsts.DriveFolderPrefixBuilder(driveID)
		cache     = idname.NewCache(
			// Cache check with lowercased ids
			map[string]string{strings.ToLower(driveID): driveName},
		)
		dii           = odStub.DriveItemInfo()
		expectedPath  = "Libraries/" + driveName
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
		cache,
		nil,
		fault.New(true))
	assert.NoError(t, err, "export collections error")
	assert.Len(t, ecs, 1, "num of collections")

	assert.Equal(t, expectedPath, ecs[0].BasePath(), "base dir")

	fitems := []export.Item{}
	for item := range ecs[0].Items(ctx) {
		fitems = append(fitems, item)
	}

	assert.Equal(t, expectedItems, fitems, "items")
}
