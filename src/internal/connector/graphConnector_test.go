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
	table := []struct {
		name   string
		params []string
	}{
		{
			name:   "Invalid Credentials",
			params: []string{"Test", "without", "data"},
		},
		{
			name:   "Empty Credentials",
			params: []string{"", "", ""},
		},
	}
	for _, test := range table {
		gc, err := graph.NewGraphConnector(test.params[0], test.params[1], test.params[2])
		assert.Nil(suite.T(), gc, test.name+" failed")
		assert.NotNil(suite.T(), err, test.name+"failed")
	}
}

func (suite *GraphConnectorTestSuite) TestGraphConnector() {
	if os.Getenv("INTEGRATION_TESTING") != "" {
		suite.T().Skip("Environmental Variables not set")
	}
	suite.True(suite.connector != nil)
}
