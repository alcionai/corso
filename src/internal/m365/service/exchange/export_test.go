package exchange

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/converters/eml/testdata"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/collection/exchange"
	exchMock "github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
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

func (suite *ExportUnitSuite) TestGetItems() {
	emailBodyBytes := []byte(testdata.EmailWithAttachments)

	pb := path.Builder{}.Append("Inbox")
	p, err := pb.ToDataLayerPath("t", "r", path.ExchangeService, path.EmailCategory, false)
	assert.NoError(suite.T(), err, "build path")

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
					Path: p,
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
			name:    "single item with special characters",
			version: 1,
			backingCollection: data.NoFetchRestoreCollection{
				Collection: dataMock.Collection{
					Path: p,
					ItemData: []data.Item{
						&dataMock.Item{
							ItemID: "id1",
							Reader: io.NopCloser(bytes.NewReader(
								exchMock.MessageWithSpecialCharacters("special characters"))),
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
					Path: p,
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
					Path: p,
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

			stats := metrics.NewExportStats()
			ec := exchange.NewExportCollection(
				"",
				[]data.RestoreCollection{test.backingCollection},
				test.version,
				stats)

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

			var expectedStats metrics.ExportStats

			if size+count > 0 { // it is only initialized if we have something
				expectedStats = metrics.ExportStats{}
				expectedStats.UpdateBytes(path.EmailCategory, int64(size))

				for i := 0; i < count; i++ {
					expectedStats.UpdateResourceCount(path.EmailCategory)
				}
			}

			assert.Equal(t, expectedStats.GetStats(), stats.GetStats(), "stats")
		})
	}
}

func (suite *ExportUnitSuite) TestExportRestoreCollections() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	emailBodyBytes := []byte(testdata.EmailWithAttachments)

	pb := path.Builder{}.Append("Inbox")
	p, err := pb.ToDataLayerPath("t", "r", path.ExchangeService, path.EmailCategory, false)
	assert.NoError(t, err, "build path")

	p2, err := pb.ToDataLayerPath("t", "r", path.OneDriveService, path.FilesCategory, false)
	assert.NoError(t, err, "build path")

	tests := []struct {
		name          string
		dcs           []data.RestoreCollection
		expectedItems [][]export.Item
		hasErr        bool
	}{
		{
			name: "single item",
			dcs: []data.RestoreCollection{
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
			},
			expectedItems: [][]export.Item{
				{
					{
						ID:   "id1",
						Name: "id1.eml",
						Body: io.NopCloser(bytes.NewReader(emailBodyBytes)),
					},
				},
			},
		},
		{
			name: "multiple items",
			dcs: []data.RestoreCollection{
				data.FetchRestoreCollection{
					Collection: dataMock.Collection{
						Path: p,
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
			},
			expectedItems: [][]export.Item{
				{
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
		},
		{
			name: "items with success and fetch error",
			dcs: []data.RestoreCollection{
				data.FetchRestoreCollection{
					Collection: dataMock.Collection{
						Path: p,
						ItemData: []data.Item{
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
			},
			expectedItems: [][]export.Item{
				{
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
		},
		{
			name: "multiple collections",
			dcs: []data.RestoreCollection{
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
				data.FetchRestoreCollection{
					Collection: dataMock.Collection{
						Path: p,
						ItemData: []data.Item{
							&dataMock.Item{
								ItemID: "id2",
								Reader: io.NopCloser(bytes.NewReader(emailBodyBytes)),
							},
						},
					},
				},
			},
			expectedItems: [][]export.Item{
				{
					{
						ID:   "id1",
						Name: "id1.eml",
						Body: io.NopCloser(bytes.NewReader(emailBodyBytes)),
					},
				},
				{
					{
						ID:   "id2",
						Name: "id2.eml",
						Body: io.NopCloser(bytes.NewReader(emailBodyBytes)),
					},
				},
			},
		},
		{
			name: "collection without exchange category",
			dcs: []data.RestoreCollection{
				data.FetchRestoreCollection{
					Collection: dataMock.Collection{
						Path: p2,
						ItemData: []data.Item{
							&dataMock.Item{
								ItemID: "id1",
								Reader: io.NopCloser(bytes.NewReader(emailBodyBytes)),
							},
						},
					},
				},
			},
			expectedItems: [][]export.Item{},
			hasErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exportCfg := control.ExportConfig{}
			stats := metrics.NewExportStats()

			ecs, err := NewExchangeHandler(api.Client{}, nil).
				ProduceExportCollections(
					ctx,
					int(version.Backup),
					exportCfg,
					tt.dcs,
					stats,
					fault.New(true))

			if tt.hasErr {
				assert.Error(t, err, "export collections error")
				return
			}

			assert.NoError(t, err, "export collections error")
			assert.Len(t, ecs, len(tt.expectedItems), "num of collections")

			expectedStats := metrics.NewExportStats()

			// We are dependent on the order the collections are
			// returned in the test which is not necessary for the
			// correctness out the output.
			for c := range ecs {
				i := -1
				for item := range ecs[c].Items(ctx) {
					i++

					size := 0

					if item.Body == nil {
						assert.ErrorIs(t, item.Error, tt.expectedItems[c][i].Error)
						continue
					}

					// unwrap the body from stats reader
					b, err := io.ReadAll(item.Body)
					assert.NoError(t, err, clues.ToCore(err))

					size += len(b)

					expectedStats.UpdateBytes(path.EmailCategory, int64(size))
					expectedStats.UpdateResourceCount(path.EmailCategory)

					assert.Equal(t, tt.expectedItems[c][i].ID, item.ID, "id")
					assert.Equal(t, tt.expectedItems[c][i].Name, item.Name, "name")
					assert.NoError(t, item.Error, "error")

				}
			}

			assert.Equal(t, expectedStats.GetStats(), stats.GetStats(), "stats")
		})
	}
}
