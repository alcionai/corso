package connector_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	graph "github.com/alcionai/corso/internal/connector"
)

type GraphConnectorTestSuite struct {
	suite.Suite
	connector *graph.GraphConnector
}

type DiconnectedGraphConnectorTestSuite struct {
	suite.Suite
}

func (suite *GraphConnectorTestSuite) SetupSuite() {
	tenant := os.Getenv("TENANT_ID")
	client := os.Getenv("CLIENT_ID")
	secret := os.Getenv("CLIENT_SECRET")
	if os.Getenv("CI") == "" {
		var err error
		suite.connector, err = graph.NewGraphConnector(tenant, client, secret)
		assert.Nil(suite.T(), err)
	}
}

func TestGraphConnectorSuite(t *testing.T) {
	suite.Run(t, new(GraphConnectorTestSuite))
}

func TestDisconnectedGraphSuite(t *testing.T) {
	suite.Run(t, new(DiconnectedGraphConnectorTestSuite))
}

func (suite *DiconnectedGraphConnectorTestSuite) TestBadConnection() {
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

//
func (suite *GraphConnectorTestSuite) TestMailCount() {
	if os.Getenv("INTEGRATION_TESTING") != "" {
		suite.T().Skip("Environmental Variables not set")
	}
	exchangeData, err := suite.connector.ExchangeDataCollection("dustina@8qzvrj.onmicrosoft.com")
	assert.NotNil(suite.T(), exchangeData)
	suite.T().Logf("Missing Data: %s\n", err.Error())
}

func (suite *DiconnectedGraphConnectorTestSuite) TestInterfaceAlignment() {
	var dc graph.DataCollection
	concrete := graph.NewExchangeDataCollection("Check", 1, []string{"interface", "works"})
	dc = &concrete
	assert.NotNil(suite.T(), dc)
}
