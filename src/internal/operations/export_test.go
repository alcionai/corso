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
	"github.com/alcionai/corso/src/internal/m365/mock"
	exchMock "github.com/alcionai/corso/src/internal/m365/service/exchange/mock"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

type ExportUnitSuite struct {
	tester.Suite
}

func TestExportUnitSuite(t *testing.T) {
	suite.Run(t, &ExportUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExportUnitSuite) TestExportOperation_Export() {
	var (
		kw        = &kopia.Wrapper{}
		sw        = store.NewWrapper(&kopia.ModelStore{})
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
			},
		},
		{
			expectStatus: Failed,
			expectErr:    assert.Error,
			fail:         assert.AnError,
			stats: exportStats{
				bytesRead: &stats.ByteCounter{},
			},
		},
		{
			expectStatus: NoData,
			expectErr:    assert.NoError,
			stats: exportStats{
				bytesRead: &stats.ByteCounter{},
				cs:        []data.RestoreCollection{},
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

func (suite *ExportUnitSuite) TestZipExports() {
	table := []struct {
		name          string
		inputColls    []export.Collectioner
		expectZipErr  assert.ErrorAssertionFunc
		expectReadErr assert.ErrorAssertionFunc
	}{
		{
			name:          "nothing",
			inputColls:    []export.Collectioner{},
			expectZipErr:  assert.Error,
			expectReadErr: assert.NoError,
		},
		{
			name: "empty",
			inputColls: []export.Collectioner{
				expCol{
					base:  "",
					items: []export.Item{},
				},
			},
			expectZipErr:  assert.NoError,
			expectReadErr: assert.NoError,
		},
		{
			name: "one item",
			inputColls: []export.Collectioner{
				expCol{
					base: "",
					items: []export.Item{
						{
							ID:   "id1",
							Name: "test",
							Body: NewReadSeekCloser([]byte("test")),
						},
					},
				},
			},
			expectZipErr:  assert.NoError,
			expectReadErr: assert.NoError,
		},
		{
			name: "multiple items",
			inputColls: []export.Collectioner{
				expCol{
					base: "",
					items: []export.Item{
						{
							ID:   "id1",
							Name: "test",
							Body: NewReadSeekCloser([]byte("test")),
						},
					},
				},
				expCol{
					base: "/fold",
					items: []export.Item{
						{
							ID:   "id2",
							Name: "test2",
							Body: NewReadSeekCloser([]byte("test2")),
						},
					},
				},
			},
			expectZipErr:  assert.NoError,
			expectReadErr: assert.NoError,
		},
		{
			name: "one item with err",
			inputColls: []export.Collectioner{
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
			expectZipErr:  assert.NoError,
			expectReadErr: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			zc, err := archive.ZipExportCollection(ctx, test.inputColls)
			test.expectZipErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Empty(t, zc.BasePath(), "base path")

			zippedItems := []export.Item{}

			count := 0
			for item := range zc.Items(ctx) {
				assert.True(t, strings.HasPrefix(item.Name, "Corso_Export_"), "name prefix")
				assert.True(t, strings.HasSuffix(item.Name, ".zip"), "name suffix")

				data, err := io.ReadAll(item.Body)
				test.expectReadErr(t, err, clues.ToCore(err))

				if err != nil {
					return
				}

				assert.NotEmpty(t, item.Name, "item name")

				item.Body.Close()

				reader, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
				require.NoError(t, err, clues.ToCore(err))

				for _, f := range reader.File {
					rc, err := f.Open()
					assert.NoError(t, err, clues.ToCore(err))

					data, err := io.ReadAll(rc)
					require.NoError(t, err, clues.ToCore(err))

					rc.Close()

					zippedItems = append(zippedItems, export.Item{
						Name: f.Name,
						Body: NewReadSeekCloser([]byte(data)),
					})
				}

				count++
			}

			assert.Equal(t, 1, count, "single item")

			expectedZippedItems := []export.Item{}

			for _, col := range test.inputColls {
				for item := range col.Items(ctx) {
					expected := export.Item{
						Name: item.Name,
						Body: item.Body,
					}

					if len(col.BasePath()) > 0 {
						expected.Name = strings.Join([]string{col.BasePath(), item.Name}, "/")
					}

					_, err := expected.Body.(io.ReadSeeker).Seek(0, io.SeekStart)
					require.NoError(t, err, clues.ToCore(err))

					expected.ID = ""

					expectedZippedItems = append(expectedZippedItems, expected)
				}
			}

			assert.Equal(t, expectedZippedItems, zippedItems, "items")
		})
	}
}
