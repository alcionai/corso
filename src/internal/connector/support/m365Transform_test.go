package support

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	ctesting "github.com/alcionai/corso/internal/testing"
)

type SupportTestSuite struct {
	suite.Suite
}

const (
	// File needs to be a single message .json
	// Use: https://developer.microsoft.com/en-us/graph/graph-explorer for details
	SUPPORT_FILE = "CORSO_TEST_SUPPORT_FILE"
)

func TestSupportTestSuite(t *testing.T) {
	evs, err := ctesting.GetRequiredEnvVars(SUPPORT_FILE)
	if err != nil {
		t.Skipf("Env not configured: %v\n", err)
	}
	_, err = os.Stat(evs[SUPPORT_FILE])
	if err != nil {
		t.Skip("Test object not available: Module Skipped")
	}
	suite.Run(t, new(SupportTestSuite))
}

func (suite *SupportTestSuite) TestToMessage() {
	bytes, err := ctesting.LoadAFile(os.Getenv(SUPPORT_FILE))
	if err != nil {
		suite.T().Errorf("Failed with %v\n", err)
	}
	require.NoError(suite.T(), err)
	message, err := CreateMessageFromBytes(bytes)
	require.NoError(suite.T(), err)
	clone := ToMessage(message)
	suite.Equal(message.GetBccRecipients(), clone.GetBccRecipients())
	suite.Equal(message.GetSubject(), clone.GetSubject())
	suite.Equal(message.GetSender(), clone.GetSender())
	suite.Equal(message.GetSentDateTime(), clone.GetSentDateTime())
	suite.NotEqual(message.GetId(), clone.GetId())

}
