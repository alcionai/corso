package support

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/tester"
)

type SupportTestSuite struct {
	suite.Suite
}

func TestSupportTestSuite(t *testing.T) {
	evs, err := tester.GetRequiredEnvVars(tester.CorsoGraphConnectorTestSupportFile)
	if err != nil {
		t.Skipf("Env not configured: %v\n", err)
	}
	_, err = os.Stat(evs[tester.CorsoGraphConnectorTestSupportFile])
	if err != nil {
		t.Skip("Test object not available: Module Skipped")
	}
	suite.Run(t, new(SupportTestSuite))
}

func (suite *SupportTestSuite) TestToMessage() {
	bytes, err := tester.LoadAFile(os.Getenv(tester.CorsoGraphConnectorTestSupportFile))
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
