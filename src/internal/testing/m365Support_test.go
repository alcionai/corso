package testing

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DataSupportSuite struct {
	suite.Suite
}

const (
	// File needs to be a single message .json
	// Use: https://developer.microsoft.com/en-us/graph/graph-explorer for details
	SUPPORT_FILE = "CORSO_TEST_DATA_SUPPORT_FILE"
)

func TestDataSupportSuite(t *testing.T) {
	err := RunOnAny(SUPPORT_FILE)
	if err != nil {
		t.Skipf("Skipping: %v\n", err)
	}
	suite.Run(t, new(DataSupportSuite))
}

func (suite *DataSupportSuite) TestCreateMessageFromBytes() {
	bytes, err := LoadAFile(os.Getenv(SUPPORT_FILE))
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
