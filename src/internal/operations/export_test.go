package operations

import (
	"archive/zip"
	"bytes"
	"context"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/archive"
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
				control.DefaultOptions(),
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

			err = op.finalizeMetrics(ctx, now, &test.stats)
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

// ReadSeekCloser implements io.ReadSeekCloser.
type ReadSeekCloser struct {
	*bytes.Reader
}

// NewReadSeekCloser creates a new ReadSeekCloser from a byte slice.
func NewReadSeekCloser(byts []byte) *ReadSeekCloser {
	return &ReadSeekCloser{
		Reader: bytes.NewReader(byts),
	}
}

// Close implements the io.Closer interface.
func (r *ReadSeekCloser) Close() error {
	// Nothing to close for a byte slice.
	return nil
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
								Body: NewReadSeekCloser([]byte("test")),
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
								Body: NewReadSeekCloser([]byte("test")),
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
								Body: NewReadSeekCloser([]byte("test2")),
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

			zc, err := archive.ZipExportCollection(ctx, test.collection)

			if test.shouldErr {
				assert.Error(t, err, "error")
				return
			}

			require.NoError(t, err, "error")
			assert.Empty(t, zc.BasePath(), "base path")

			zippedItems := []export.ItemData{}

			count := 0
			for item := range zc.Items(ctx) {
				assert.True(t, strings.HasPrefix(item.Data.Name, "Corso_Export_"), "name prefix")
				assert.True(t, strings.HasSuffix(item.Data.Name, ".zip"), "name suffix")

				data, err := io.ReadAll(item.Data.Body)
				if test.readErr {
					assert.Error(t, err, "read error")
					return
				}

				size := int64(len(data))

				item.Data.Body.Close()

				reader, err := zip.NewReader(bytes.NewReader(data), size)
				require.NoError(t, err, "zip reader")

				for _, f := range reader.File {
					rc, err := f.Open()
					assert.NoError(t, err, "open file in zip")

					data, err := io.ReadAll(rc)
					require.NoError(t, err, "read zip file content")

					rc.Close()

					zippedItems = append(zippedItems, export.ItemData{
						Name: f.Name,
						Body: NewReadSeekCloser([]byte(data)),
					})
				}

				count++
			}

			assert.Equal(t, 1, count, "single item")

			expectedZippedItems := []export.ItemData{}
			for _, col := range test.collection {
				for item := range col.Items(ctx) {
					if col.BasePath() != "" {
						item.Data.Name = strings.Join([]string{col.BasePath(), item.Data.Name}, "/")
					}
					_, err := item.Data.Body.(io.ReadSeeker).Seek(0, io.SeekStart)
					require.NoError(t, err, "seek")
					expectedZippedItems = append(expectedZippedItems, item.Data)
				}
			}
			assert.Equal(t, expectedZippedItems, zippedItems, "items")
		})
	}
}
