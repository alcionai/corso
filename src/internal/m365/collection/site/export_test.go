package site

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	kjson "github.com/microsoft/kiota-serialization-json-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/common/ptr"
	"github.com/alcionai/canario/src/internal/data"
	dataMock "github.com/alcionai/canario/src/internal/data/mock"
	"github.com/alcionai/canario/src/internal/tester"
	"github.com/alcionai/canario/src/internal/version"
	"github.com/alcionai/canario/src/pkg/control"
	"github.com/alcionai/canario/src/pkg/export"
	"github.com/alcionai/canario/src/pkg/metrics"
)

type ExportUnitSuite struct {
	tester.Suite
}

func TestExportUnitSuite(t *testing.T) {
	suite.Run(t, &ExportUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ExportUnitSuite) TestStreamItems() {
	t := suite.T()

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
						ItemID: "list1",
						Reader: makeListJSONReader(t, "list1"),
					},
				},
			},
			expectName: "list1.json",
			expectErr:  assert.NoError,
		},
		{
			name: "only recoverable errors",
			backingColl: dataMock.Collection{
				ItemsRecoverableErrs: []error{
					clues.New("some error"),
				},
			},
			expectErr: assert.Error,
		},
		{
			name: "items and recoverable errors",
			backingColl: dataMock.Collection{
				ItemData: []data.Item{
					&dataMock.Item{
						ItemID: "list2",
						Reader: makeListJSONReader(t, "list2"),
					},
				},
				ItemsRecoverableErrs: []error{
					clues.New("some error"),
				},
			},
			expectName: "list2.json",
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

func makeListJSONReader(t *testing.T, listName string) io.ReadCloser {
	listBytes := getListBytes(t, listName)
	return io.NopCloser(bytes.NewReader(listBytes))
}

func getListBytes(t *testing.T, listName string) []byte {
	writer := kjson.NewJsonSerializationWriter()
	defer writer.Close()

	list := models.NewList()
	list.SetId(ptr.To(listName))

	err := writer.WriteObjectValue("", list)
	require.NoError(t, err)

	storedListBytes, err := writer.GetSerializedContent()
	require.NoError(t, err)

	return storedListBytes
}
