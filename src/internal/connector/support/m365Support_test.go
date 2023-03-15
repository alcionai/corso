package support

import (
	"testing"

	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	bmodels "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/tester"
)

type DataSupportSuite struct {
	tester.Suite
}

func TestDataSupportSuite(t *testing.T) {
	suite.Run(t, &DataSupportSuite{Suite: tester.NewUnitSuite(t)})
}

var (
	empty   = "Empty Bytes"
	invalid = "Invalid Bytes"
)

// TestCreateMessageFromBytes verifies approved mockdata bytes can
// be successfully transformed into M365 Message data.
func (suite *DataSupportSuite) TestCreateMessageFromBytes() {
	table := []struct {
		name        string
		byteArray   []byte
		checkError  assert.ErrorAssertionFunc
		checkObject assert.ValueAssertionFunc
	}{
		{
			name:        "Empty Bytes",
			byteArray:   make([]byte, 0),
			checkError:  assert.Error,
			checkObject: assert.Nil,
		},
		{
			name:        "aMessage bytes",
			byteArray:   mockconnector.GetMockMessageBytes("m365 mail support test"),
			checkError:  assert.NoError,
			checkObject: assert.NotNil,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := CreateMessageFromBytes(test.byteArray)
			test.checkError(t, err, clues.ToCore(err))
			test.checkObject(t, result)
		})
	}
}

// TestCreateContactFromBytes verifies behavior of CreateContactFromBytes
// by ensuring correct error and object output.
func (suite *DataSupportSuite) TestCreateContactFromBytes() {
	table := []struct {
		name       string
		byteArray  []byte
		checkError assert.ErrorAssertionFunc
		isNil      assert.ValueAssertionFunc
	}{
		{
			name:       empty,
			byteArray:  make([]byte, 0),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       invalid,
			byteArray:  []byte("A random sentence doesn't make an object"),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Valid Contact",
			byteArray:  mockconnector.GetMockContactBytes("Support Test"),
			checkError: assert.NoError,
			isNil:      assert.NotNil,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := CreateContactFromBytes(test.byteArray)
			test.checkError(t, err, clues.ToCore(err))
			test.isNil(t, result)
		})
	}
}

func (suite *DataSupportSuite) TestCreateEventFromBytes() {
	tests := []struct {
		name       string
		byteArray  []byte
		checkError assert.ErrorAssertionFunc
		isNil      assert.ValueAssertionFunc
	}{
		{
			name:       empty,
			byteArray:  make([]byte, 0),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       invalid,
			byteArray:  []byte("Invalid byte stream \"subject:\" Not going to work"),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Valid Event",
			byteArray:  mockconnector.GetDefaultMockEventBytes("Event Test"),
			checkError: assert.NoError,
			isNil:      assert.NotNil,
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			result, err := CreateEventFromBytes(test.byteArray)
			test.checkError(t, err, clues.ToCore(err))
			test.isNil(t, result)
		})
	}
}

func (suite *DataSupportSuite) TestCreateListFromBytes() {
	listBytes, err := mockconnector.GetMockListBytes("DataSupportSuite")
	require.NoError(suite.T(), err)

	tests := []struct {
		name       string
		byteArray  []byte
		checkError assert.ErrorAssertionFunc
		isNil      assert.ValueAssertionFunc
	}{
		{
			name:       empty,
			byteArray:  make([]byte, 0),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       invalid,
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

func (suite *DataSupportSuite) TestCreatePageFromBytes() {
	tests := []struct {
		name       string
		checkError assert.ErrorAssertionFunc
		isNil      assert.ValueAssertionFunc
		getBytes   func(t *testing.T) []byte
	}{
		{
			empty,
			assert.Error,
			assert.Nil,
			func(t *testing.T) []byte {
				return make([]byte, 0)
			},
		},
		{
			invalid,
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
				err := pg.Serialize(writer)
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
		})
	}
}
