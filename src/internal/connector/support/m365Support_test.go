package support

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	ctesting "github.com/alcionai/corso/internal/testing"
)

type DataSupportSuite struct {
	suite.Suite
}

const (
	// File needs to be a single message .json
	// Use: https://developer.microsoft.com/en-us/graph/graph-explorer for details
	support_file = "CORSO_TEST_SUPPORT_FILE"
)

func TestDataSupportSuite(t *testing.T) {
	err := ctesting.RunOnAny(support_file)
	if err != nil {
		t.Skipf("Skipping: %v\n", err)
	}
	suite.Run(t, new(DataSupportSuite))
}

func (suite *DataSupportSuite) TestCreateMessageFromBytes() {
	bytes, err := ctesting.LoadAFile(os.Getenv(SUPPORT_FILE))
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

func (suite *DataSupportSuite) TestDataSupport_TaskList() {
	tasks := NewTaskList()
	tasks.AddTask("person1", "Go to store")
	tasks.AddTask("person1", "drop off mail")
	values := tasks["person1"]
	suite.Equal(len(values), 2)
	nonValues := tasks["unknown"]
	suite.Zero(len(nonValues))
}
