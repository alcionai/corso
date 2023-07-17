package operations

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	exchMock "github.com/alcionai/corso/src/internal/m365/exchange/mock"
	"github.com/alcionai/corso/src/internal/m365/mock"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

type ExportOpSuite struct {
	tester.Suite
}

func TestExportOpSuite(t *testing.T) {
	suite.Run(t, &ExportOpSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExportOpSuite) TestExportOperation_PersistResults() {
	var (
		kw        = &kopia.Wrapper{}
		sw        = &store.Wrapper{}
		ctrl      = &mock.Controller{}
		now       = time.Now()
		exportCfg = control.DefaultExportConfig()
	)

	table := []struct {
		expectStatus OpStatus
		expectErr    assert.ErrorAssertionFunc
		stats        exportStats
		fail         error
	}{
		{
			expectStatus: Completed,
			expectErr:    assert.NoError,
			stats: exportStats{
				resourceCount: 1,
				bytesRead: &stats.ByteCounter{
					NumBytes: 42,
				},
				cs: []data.RestoreCollection{
					data.NoFetchRestoreCollection{
						Collection: &exchMock.DataCollection{},
					},
				},
				ctrl: &data.CollectionStats{
					Objects:   1,
					Successes: 1,
				},
			},
		},
		{
			expectStatus: Failed,
			expectErr:    assert.Error,
			fail:         assert.AnError,
			stats: exportStats{
				bytesRead: &stats.ByteCounter{},
				ctrl:      &data.CollectionStats{},
			},
		},
		{
			expectStatus: NoData,
			expectErr:    assert.NoError,
			stats: exportStats{
				bytesRead: &stats.ByteCounter{},
				cs:        []data.RestoreCollection{},
				ctrl:      &data.CollectionStats{},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.expectStatus.String(), func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			op, err := NewExportOperation(
				ctx,
				control.Defaults(),
				kw,
				sw,
				ctrl,
				account.Account{},
				"foo",
				selectors.Selector{DiscreteOwner: "test"},
				exportCfg,
				evmock.NewBus())
			require.NoError(t, err, clues.ToCore(err))

			op.Errors.Fail(test.fail)

			err = op.persistResults(ctx, now, &test.stats)
			test.expectErr(t, err, clues.ToCore(err))

			assert.Equal(t, test.expectStatus.String(), op.Status.String(), "status")
			assert.Equal(t, len(test.stats.cs), op.Results.ItemsRead, "items read")
			assert.Equal(t, test.stats.bytesRead.NumBytes, op.Results.BytesRead, "resource owners")
			assert.Equal(t, test.stats.resourceCount, op.Results.ResourceOwners, "resource owners")
			assert.Equal(t, now, op.Results.StartedAt, "started at")
			assert.Less(t, now, op.Results.CompletedAt, "completed at")
		})
	}
}

type expCol struct {
	base  string
	items []export.Item
}

func (ec expCol) BasePath() string { return ec.base }
func (ec expCol) Items(ctx context.Context) <-chan export.Item {
	ch := make(chan export.Item)

	go func() {
		defer close(ch)

		for _, item := range ec.items {
			ch <- item
		}
	}()

	return ch
}

func (suite *ExportOpSuite) TestZipExports() {
	table := []struct {
		name       string
		collection []export.Collection
		shouldErr  bool
		readErr    bool
	}{
		{
			name:       "nothing",
			collection: []export.Collection{},
			shouldErr:  true,
		},
		{
			name: "empty",
			collection: []export.Collection{
				expCol{
					base:  "",
					items: []export.Item{},
				},
			},
		},
		{
			name: "one item",
			collection: []export.Collection{
				expCol{
					base: "",
					items: []export.Item{
						{
							ID: "id1",
							Data: export.ItemData{
								Name: "test",
								Body: io.NopCloser(bytes.NewBufferString("test")),
							},
						},
					},
				},
			},
		},
		{
			name: "multiple items",
			collection: []export.Collection{
				expCol{
					base: "",
					items: []export.Item{
						{
							ID: "id1",
							Data: export.ItemData{
								Name: "test",
								Body: io.NopCloser(bytes.NewBufferString("test")),
							},
						},
					},
				},
				expCol{
					base: "/fold",
					items: []export.Item{
						{
							ID: "id2",
							Data: export.ItemData{
								Name: "test2",
								Body: io.NopCloser(bytes.NewBufferString("test2")),
							},
						},
					},
				},
			},
		},
		{
			name: "one item with err",
			collection: []export.Collection{
				expCol{
					base: "",
					items: []export.Item{
						{
							ID:    "id3",
							Error: assert.AnError,
						},
					},
				},
			},
			readErr: true,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			zc, err := zipExportCollection(ctx, test.collection)

			if test.shouldErr {
				assert.Error(t, err, "error")
				return
			}

			require.NoError(t, err, "error")

			assert.Empty(t, zc.BasePath(), "base path")

			count := 0
			for item := range zc.Items(ctx) {
				assert.Equal(t, "export.zip", item.Data.Name, "name")

				_, err := io.Copy(io.Discard, item.Data.Body)
				if test.readErr {
					assert.Error(t, err, "read error")
					return
				}

				require.NoError(t, err, "read item")

				item.Data.Body.Close()

				count++
			}

			assert.Equal(t, 1, count, "single item")
		})
	}
}
