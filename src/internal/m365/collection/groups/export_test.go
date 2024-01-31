package groups

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/metrics"
	"github.com/alcionai/corso/src/pkg/path"
)

type ExportUnitSuite struct {
	tester.Suite
}

func TestExportUnitSuite(t *testing.T) {
	suite.Run(t, &ExportUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExportUnitSuite) TestStreamChannelMessages() {
	makeBody := func() io.ReadCloser {
		return io.NopCloser(bytes.NewReader([]byte("{}")))
	}

	table := []struct {
		name        string
		backingColl dataMock.Collection
		expectName  string
		expectErr   assert.ErrorAssertionFunc
	}{
		{
			name: "no errors",
			backingColl: dataMock.Collection{
				ItemData: []data.Item{
					&dataMock.Item{
						ItemID: "zim",
						Reader: makeBody(),
					},
				},
			},
			expectName: "zim.json",
			expectErr:  assert.NoError,
		},
		{
			name: "only recoverable errors",
			backingColl: dataMock.Collection{
				ItemsRecoverableErrs: []error{
					clues.New("The knowledge... it fills me! It is neat!"),
				},
			},
			expectErr: assert.Error,
		},
		{
			name: "items and recoverable errors",
			backingColl: dataMock.Collection{
				ItemData: []data.Item{
					&dataMock.Item{
						ItemID: "gir",
						Reader: makeBody(),
					},
				},
				ItemsRecoverableErrs: []error{
					clues.New("I miss my cupcake."),
				},
			},
			expectName: "gir.json",
			expectErr:  assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ch := make(chan export.Item)

			go streamChannelMessages(
				ctx,
				[]data.RestoreCollection{test.backingColl},
				version.NoBackup,
				control.DefaultExportConfig(),
				ch,
				&metrics.ExportStats{})

			var (
				itm export.Item
				err error
			)

			for i := range ch {
				if i.Error == nil {
					itm = i
				} else {
					err = i.Error
				}
			}

			test.expectErr(t, err, clues.ToCore(err))

			assert.Equal(t, test.expectName, itm.Name, "item name")
		})
	}
}

func (suite *ExportUnitSuite) TestStreamConversationPosts() {
	testPath, err := path.Build(
		"t",
		"g",
		path.GroupsService,
		path.ConversationPostsCategory,
		true,
		"convID",
		"threadID")
	require.NoError(suite.T(), err, clues.ToCore(err))

	makeBody := func() io.ReadCloser {
		rc := io.NopCloser(bytes.NewReader([]byte("{}")))

		return metrics.ReaderWithStats(
			rc,
			path.ConversationPostsCategory,
			&metrics.ExportStats{})
	}

	makeMeta := func() io.ReadCloser {
		return io.NopCloser(
			bytes.NewReader([]byte(`{"topic":"t", "recipients":["em@il"]}`)))
	}

	table := []struct {
		name        string
		backingColl dataMock.Collection
		expectItem  export.Item
		expectErr   assert.ErrorAssertionFunc
	}{
		{
			name: "no errors",
			backingColl: dataMock.Collection{
				ItemData: []data.Item{
					&dataMock.Item{
						ItemID: "zim.data",
						Reader: makeBody(),
					},
				},
				Path: testPath,
				AuxItems: map[string]data.Item{
					"zim.meta": &dataMock.Item{
						ItemID: "zim.meta",
						Reader: makeMeta(),
					},
				},
			},
			expectItem: export.Item{
				ID:   "zim.data",
				Name: "zim.eml",
				Body: makeBody(),
			},
			expectErr: assert.NoError,
		},
		{
			name: "only recoverable errors",
			backingColl: dataMock.Collection{
				ItemsRecoverableErrs: []error{
					clues.New("The knowledge... it fills me! It is neat!"),
				},
				Path: testPath,
			},
			expectErr: assert.Error,
		},
		{
			name: "items and recoverable errors",
			backingColl: dataMock.Collection{
				ItemData: []data.Item{
					&dataMock.Item{
						ItemID: "gir.data",
						Reader: makeBody(),
					},
				},
				ItemsRecoverableErrs: []error{
					clues.New("I miss my cupcake."),
				},
				Path: testPath,
				AuxItems: map[string]data.Item{
					"gir.meta": &dataMock.Item{
						ItemID: "gir.meta",
						Reader: makeMeta(),
					},
				},
			},
			expectItem: export.Item{
				ID:   "gir.data",
				Name: "gir.eml",
				Body: makeBody(),
			},
			expectErr: assert.Error,
		},
		{
			name: "missing metadata",
			backingColl: dataMock.Collection{
				ItemData: []data.Item{
					&dataMock.Item{
						ItemID: "mir.data",
						Reader: makeBody(),
					},
				},
				Path: testPath,
			},
			expectItem: export.Item{
				ID:    "mir.data",
				Error: assert.AnError,
			},
			expectErr: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ch := make(chan export.Item)

			go streamConversationPosts(
				ctx,
				[]data.RestoreCollection{test.backingColl},
				version.NoBackup,
				control.DefaultExportConfig(),
				ch,
				&metrics.ExportStats{})

			var (
				itm export.Item
				err error
			)

			for i := range ch {
				if i.Error == nil {
					itm = i
				} else {
					err = i.Error
				}
			}

			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.Equal(t, test.expectItem.ID, itm.ID, "item ID")
			assert.Equal(t, test.expectItem.Name, itm.Name, "item name")
			assert.NotNil(t, itm.Body, "body")

			_, err = io.ReadAll(itm.Body)
			require.NoError(t, err, clues.ToCore(err))

			itm.Body.Close()
		})
	}
}
