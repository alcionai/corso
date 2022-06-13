package support

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	ctesting "github.com/alcionai/corso/internal/testing"
)

type SupportTestSuite struct {
	suite.Suite
}

const (
	SUPPORT_FILE = "TEST_SUPPORT_FILE" //message json with one message
)

func TestSupportTestSuite(t *testing.T) {
	file := "TEST_SUPPORT_FILE"
	evs, err := ctesting.GetRequiredEnvVars(file)
	if err != nil {
		t.Skipf("Env not configured: %v\n", err)
	}
	_, err = os.Stat(evs[file])
	if err != nil {
		t.Skip("Test object not available: Module Skipped")
	}
	suite.Run(t, new(SupportTestSuite))
}

// CreateMessageFromBytes
// Swap Message

func (suite *SupportTestSuite) TestSwapMessage() {
	bytes, err := LoadAFile(os.Getenv(SUPPORT_FILE))
	if err != nil {
		suite.T().Errorf("Failed with %v\n", err)
	}
	message, err := CreateMessageFromBytes(bytes)
	assert.NoError(suite.T(), err)
	clone := SwapMessage(message)
	suite.Equal(message.GetBccRecipients(), clone.GetBccRecipients())
	suite.Equal(message.GetSubject(), clone.GetSubject())
	suite.Equal(message.GetSender(), clone.GetSender())
	suite.Equal(message.GetSentDateTime(), clone.GetSentDateTime())
	suite.NotEqual(message.GetId(), clone.GetId())

}

func (suite *SupportTestSuite) TestCreateMessageFromBytes() {
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
