package groups

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/export"
	"github.com/alcionai/corso/src/pkg/metrics"
)

type ExportUnitSuite struct {
	tester.Suite
}

func TestExportUnitSuite(t *testing.T) {
	suite.Run(t, &ExportUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExportUnitSuite) TestStreamItems() {
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

			go streamItems(
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
