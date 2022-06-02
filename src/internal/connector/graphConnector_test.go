package connector_test

import (
	"os"
	"testing"

	graph "github.com/alcionai/corso/internal/connector"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GraphConnectorTestSuite struct {
	suite.Suite
	connector *graph.GraphConnector
	err       error
}

func (suite *GraphConnectorTestSuite) SetupSuite() {
	tenant := os.Getenv("TENANT_ID")
	client := os.Getenv("CLIENT_ID")
	secret := os.Getenv("CLIENT_SECRET")
	if os.Getenv("CI") == "" {
		suite.connector, suite.err = graph.NewGraphConnector(tenant, client, secret)
	}
}

func TestGraphConnectorSuite(t *testing.T) {
	suite.Run(t, new(GraphConnectorTestSuite))
}

func (suite *GraphConnectorTestSuite) TestBadConnection() {
	gc, err := graph.NewGraphConnector("Test", "without", "data") //NOTE:[[o
	assert.NotNil(suite.T(), gc)
	assert.Nil(suite.T(), err)
	suite.Equal(len(gc.GetUsers()), 0)

}

func (suite *GraphConnectorTestSuite) TestGraphConnector() {
	if os.Getenv("INTEGRATION_TESTING") != "" {
		suite.T().Skip("Environmental Variables not set")
	}

	suite.False(suite.connector.HasConnectorErrors())
	suite.True(len(suite.connector.Users) > 0)
}
