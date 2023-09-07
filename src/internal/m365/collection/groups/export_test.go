package groups

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/export"
)

type ExportUnitSuite struct {
	tester.Suite
}

func TestExportUnitSuite(t *testing.T) {
	suite.Run(t, &ExportUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExportUnitSuite) TestStreamItems() {
	table := []struct {
		name        string
		backingColl dataMock.Collection
		expectName  string
		expectErr   assert.ErrorAssertionFunc
	}{
		{
			name: "no errors",
			backingColl: dataMock.Collection{
				ItemData: []*dataMock.Item{
					{ItemID: "zim"},
				},
			},
			expectName: "zim",
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
				ItemData: []*dataMock.Item{
					{ItemID: "gir"},
				},
				ItemsRecoverableErrs: []error{
					clues.New("I miss my cupcake."),
				},
			},
			expectName: "gir",
			expectErr:  assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ch := make(chan export.Item)

			streamItems(
				ctx,
				test.backingColl,
				version.NoBackup,
				ch)

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
