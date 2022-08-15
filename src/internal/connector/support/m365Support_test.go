package support

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/mockconnector"
	"github.com/alcionai/corso/internal/tester"
)

type DataSupportSuite struct {
	suite.Suite
}

func TestDataSupportSuite(t *testing.T) {
	err := tester.RunOnAny(tester.CorsoGraphConnectorTestSupportFile)
	if err != nil {
		t.Skipf("Skipping: %v\n", err)
	}
	suite.Run(t, new(DataSupportSuite))
}

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
		suite.T().Run(test.name, func(t *testing.T) {
			result, err := CreateMessageFromBytes(test.byteArray)
			test.checkError(t, err)
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
			name:       "Empty Bytes",
			byteArray:  make([]byte, 0),
			checkError: assert.Error,
			isNil:      assert.Nil,
		},
		{
			name:       "Invalid Bytes",
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
		suite.T().Run(test.name, func(t *testing.T) {
			result, err := CreateContactFromBytes(test.byteArray)
			test.checkError(t, err)
			test.isNil(t, result)
		})
	}
}
