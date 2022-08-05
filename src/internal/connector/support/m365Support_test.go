package support

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

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

func (suite *DataSupportSuite) TestCreateMessageFromBytes() {
	bytes, err := tester.LoadAFile(os.Getenv(tester.CorsoGraphConnectorTestSupportFile))
	if err != nil {
		suite.T().Errorf("Failed with %v\n", err)
	}

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
			byteArray:   bytes,
			checkError:  assert.NoError,
			checkObject: assert.NotNil,
		},
	}
	for _, test := range table {
		result, err := CreateMessageFromBytes(test.byteArray)
		test.checkError(suite.T(), err)
		test.checkObject(suite.T(), result)
	}
}
