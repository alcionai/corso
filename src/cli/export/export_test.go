package export

import (
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/export"
)

type ExportE2ESuite struct {
	tester.Suite
	called bool
}

func TestExportE2ESuite(t *testing.T) {
	suite.Run(t, &ExportE2ESuite{Suite: tester.NewE2ESuite(t, nil)})
}

func (suite *ExportE2ESuite) SetupSuite() {
	suite.called = true
}

type mockExportCollection struct {
	path  string
	items []export.Item
}

func (mec mockExportCollection) BasePath() string { return mec.path }
func (mec mockExportCollection) Items(context.Context) <-chan export.Item {
	ch := make(chan export.Item)

	go func() {
		defer close(ch)

		for _, item := range mec.items {
			ch <- item
		}
	}()

	return ch
}

func (suite *ExportE2ESuite) TestWriteExportCollection() {
	type ei struct {
		name string
		body string
	}

	type i struct {
		path  string
		items []ei
	}

	table := []struct {
		name string
		cols []i
	}{
		{
			name: "single root collection single item",
			cols: []i{
				{
					path: "",
					items: []ei{
						{
							name: "name1",
							body: "body1",
						},
					},
				},
			},
		},
		{
			name: "single root collection multiple items",
			cols: []i{
				{
					path: "",
					items: []ei{
						{
							name: "name1",
							body: "body1",
						},
						{
							name: "name2",
							body: "body2",
						},
					},
				},
			},
		},
		{
			name: "multiple collections multiple items",
			cols: []i{
				{
					path: "",
					items: []ei{
						{
							name: "name1",
							body: "body1",
						},
						{
							name: "name2",
							body: "body2",
						},
					},
				},
				{
					path: "folder",
					items: []ei{
						{
							name: "name3",
							body: "body3",
						},
					},
				},
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ecs := []export.Collection{}
			for _, col := range test.cols {
				items := []export.Item{}
				for _, item := range col.items {
					items = append(items, export.Item{
						Data: export.ItemData{
							Name: item.name,
							Body: io.NopCloser((bytes.NewBufferString(item.body))),
						},
					})
				}

				ecs = append(ecs, mockExportCollection{
					path:  col.path,
					items: items,
				})
			}

			dir, err := os.MkdirTemp("", "export-test")
			require.NoError(t, err)
			defer os.RemoveAll(dir)

			err = writeExportCollections(ctx, dir, ecs)
			require.NoError(t, err, "writing data")

			for _, col := range test.cols {
				for _, item := range col.items {
					f, err := os.Open(filepath.Join(dir, col.path, item.name))
					require.NoError(t, err, "opening file")

					buf := new(bytes.Buffer)

					_, err = buf.ReadFrom(f)
					require.NoError(t, err, "reading file")

					assert.Equal(t, item.body, buf.String(), "file contents")
				}
			}
		})
	}
}
