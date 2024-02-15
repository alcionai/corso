package teamschats

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
	teamschatMock "github.com/alcionai/corso/src/internal/m365/service/teamschats/mock"
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

func (suite *ExportUnitSuite) TestExportRestoreCollections_chats() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		category      = path.ChatsCategory
		itemID        = "itemID"
		dii           = teamschatMock.ItemInfo()
		content       = `{"topic": "` + dii.TeamsChats.Chat.Topic + `"}`
		body          = io.NopCloser(bytes.NewBufferString(content))
		exportCfg     = control.ExportConfig{}
		expectedPath  = category.HumanString()
		expectedItems = []export.Item{
			{
				ID:   itemID,
				Name: itemID + ".json",
				// Body: body, not checked
			},
		}
	)

	p, err := path.BuildPrefix("t", "pr", path.TeamsChatsService, category)
	require.NoError(t, err, clues.ToCore(err))

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
			FetchItemByNamer: finD{
				id:   itemID,
				key:  "id",
				name: itemID,
			},
		},
	}

	stats := metrics.NewExportStats()

	ecs, err := NewTeamsChatsHandler(api.Client{}, nil).
		ProduceExportCollections(
			ctx,
			int(version.Backup),
			exportCfg,
			dcs,
			stats,
			fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.Len(t, ecs, 1, "num of collections")

	assert.Equal(t, expectedPath, ecs[0].BasePath(), "base dir")

	fitems := []export.Item{}

	size := 0

	for item := range ecs[0].Items(ctx) {
		b, err := io.ReadAll(item.Body)
		require.NoError(t, err, clues.ToCore(err))

		// count up size for tests
		size += len(b)

		// have to nil out body, otherwise assert fails due to
		// pointer memory location differences
		item.Body = nil
		fitems = append(fitems, item)
	}

	assert.Equal(t, expectedItems, fitems, "items")

	expectedStats := metrics.NewExportStats()
	expectedStats.UpdateBytes(category, int64(size))
	expectedStats.UpdateResourceCount(category)
	assert.Equal(t, expectedStats.GetStats(), stats.GetStats(), "stats")
}
