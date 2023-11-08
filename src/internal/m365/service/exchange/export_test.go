package exchange

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/collection/exchange"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

const testEmailPath = "../../../converters/eml/testdata/email-with-attachments.json"

type ExportUnitSuite struct {
	tester.Suite
}

func TestExportUnitSuite(t *testing.T) {
	suite.Run(t, &ExportUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExportUnitSuite) TestGetItems() {
	emailBodyBytes, err := os.ReadFile(testEmailPath)
	require.NoError(suite.T(), err, "read email file")

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
							ItemID: "id1",
							Reader: io.NopCloser(bytes.NewReader(emailBodyBytes)),
						},
					},
				},
			},
			expectedItems: []export.Item{
				{
					ID:   "id1",
					Name: "id1.eml",
					Body: io.NopCloser(bytes.NewReader(emailBodyBytes)),
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
							ItemID: "id1",
							Reader: io.NopCloser(bytes.NewReader(emailBodyBytes)),
						},
						&dataMock.Item{
							ItemID: "id2",
							Reader: io.NopCloser(bytes.NewReader(emailBodyBytes)),
						},
					},
				},
			},
			expectedItems: []export.Item{
				{
					ID:   "id1",
					Name: "id1.eml",
					Body: io.NopCloser(bytes.NewReader(emailBodyBytes)),
				},
				{
					ID:   "id2",
					Name: "id2.eml",
					Body: io.NopCloser(bytes.NewReader(emailBodyBytes)),
				},
			},
		},
		{
			name:    "items with success and fetch error",
			version: version.Groups9Update,
			backingCollection: data.FetchRestoreCollection{
				Collection: dataMock.Collection{
					ItemData: []data.Item{
						&dataMock.Item{
							ItemID: "id0",
							Reader: io.NopCloser(bytes.NewReader(emailBodyBytes)),
						},
						&dataMock.Item{
							ItemID:  "id1",
							ReadErr: assert.AnError,
						},
						&dataMock.Item{
							ItemID: "id2",
							Reader: io.NopCloser(bytes.NewReader(emailBodyBytes)),
						},
					},
				},
			},
			expectedItems: []export.Item{
				{
					ID:   "id0",
					Name: "id0.eml",
					Body: io.NopCloser(bytes.NewReader(emailBodyBytes)),
				},
				{
					ID:   "id2",
					Name: "id2.eml",
					Body: io.NopCloser(bytes.NewReader(emailBodyBytes)),
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

			stats := data.ExportStats{}
			ec := exchange.NewExportCollection(
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
				assert.ErrorIs(t, item.Error, test.expectedItems[i].Error)
			}

			var expectedStats data.ExportStats

			if size+count > 0 { // it is only initialized if we have something
				expectedStats = data.ExportStats{}
				expectedStats.UpdateBytes(path.EmailCategory, int64(size))

				for i := 0; i < count; i++ {
					expectedStats.UpdateResourceCount(path.EmailCategory)
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

	emailBodyBytes, err := os.ReadFile(testEmailPath)
	require.NoError(t, err, "read email file")

	var (
		exportCfg     = control.ExportConfig{}
		expectedItems = []export.Item{
			{
				ID:   "id1",
				Name: "id1.eml",
				Body: io.NopCloser(bytes.NewReader(emailBodyBytes)),
			},
		}
	)

	pb := path.Builder{}.Append("exchange")
	p, err := pb.ToDataLayerPath("t", "r", path.ExchangeService, path.EmailCategory, false)
	assert.NoError(t, err, "build path")

	dcs := []data.RestoreCollection{
		data.FetchRestoreCollection{
			Collection: dataMock.Collection{
				Path: p,
				ItemData: []data.Item{
					&dataMock.Item{
						ItemID: "id1",
						Reader: io.NopCloser(bytes.NewReader(emailBodyBytes)),
					},
				},
			},
		},
	}

	stats := data.ExportStats{}

	ecs, err := NewExchangeHandler(control.DefaultOptions()).
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

	for i, item := range expectedItems {
		assert.Equal(t, item.ID, fitems[i].ID, "id")
		assert.Equal(t, item.Name, fitems[i].Name, "name")
		assert.NoError(t, fitems[i].Error, "error")
	}

	expectedStats := data.ExportStats{}
	expectedStats.UpdateBytes(path.EmailCategory, int64(size))
	expectedStats.UpdateResourceCount(path.EmailCategory)
	assert.Equal(t, expectedStats, stats, "stats")
}
