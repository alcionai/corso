package api

import (
	"testing"

	"github.com/alcionai/clues"
	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	bmodels "github.com/alcionai/corso/src/internal/m365/graph/betasdk/models"
	spMock "github.com/alcionai/corso/src/internal/m365/service/sharepoint/mock"
	"github.com/alcionai/corso/src/internal/tester"
)

type SerializationUnitSuite struct {
	tester.Suite
}

func TestDataSupportSuite(t *testing.T) {
	suite.Run(t, &SerializationUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SerializationUnitSuite) TestCreateListFromBytes() {
	listBytes, err := spMock.ListBytes("DataSupportSuite")
	require.NoError(suite.T(), err)

	tests := []struct {
		name       string
		byteArray  []byte
		checkError assert.ErrorAssertionFunc
		isNil      assert.ValueAssertionFunc
	}{
		{
			name:       "empty bytes",
			byteArray:  make([]byte, 0),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "invalid bytes",
			byteArray:  []byte("Invalid byte stream \"subject:\" Not going to work"),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Valid List",
			byteArray:  listBytes,
			checkError: assert.NoError,
			isNil:      assert.NotNil,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := CreateListFromBytes(test.byteArray)
			test.checkError(t, err, clues.ToCore(err))
			test.isNil(t, result)
		})
	}
}

func (suite *SerializationUnitSuite) TestCreatePageFromBytes() {
	tests := []struct {
		name       string
		checkError assert.ErrorAssertionFunc
		isNil      assert.ValueAssertionFunc
		getBytes   func(t *testing.T) []byte
	}{
		{
			"empty bytes",
			assert.Error,
			assert.Nil,
			func(t *testing.T) []byte {
				return make([]byte, 0)
			},
		},
		{
			"invalid bytes",
			assert.Error,
			assert.Nil,
			func(t *testing.T) []byte {
				return []byte("snarf")
			},
		},
		{
			"Valid Page",
			assert.NoError,
			assert.NotNil,
			func(t *testing.T) []byte {
				pg := bmodels.NewSitePage()
				title := "Tested"
				pg.SetTitle(&title)
				pg.SetName(&title)
				pg.SetWebUrl(&title)

				writer := kioser.NewJsonSerializationWriter()
				err := writer.WriteObjectValue("", pg)
				require.NoError(t, err, clues.ToCore(err))

				byteArray, err := writer.GetSerializedContent()
				require.NoError(t, err, clues.ToCore(err))

				return byteArray
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := CreatePageFromBytes(test.getBytes(t))
			test.checkError(t, err)
			test.isNil(t, result)
			if result != nil {
				assert.Equal(t, "Tested", *result.GetName(), "name")
				assert.Equal(t, "Tested", *result.GetTitle(), "title")
				assert.Equal(t, "Tested", *result.GetWebUrl(), "webURL")
			}
		})
	}
}
