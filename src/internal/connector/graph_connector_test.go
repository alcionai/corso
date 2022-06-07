package connector

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GraphConnectorTestSuite struct {
	suite.Suite
	connector *GraphConnector
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
		suite.connector, err = NewGraphConnector(tenant, client, secret)
		assert.Nil(suite.T(), err)
	}
}

func TestGraphConnectorSuite(t *testing.T) {
	suite.Run(t, new(GraphConnectorTestSuite))
}

func TestDisconnectedGraphSuite(t *testing.T) {
	suite.Run(t, new(DiconnectedGraphConnectorTestSuite))
}

func (suite *GraphConnectorTestSuite) TestGraphConnector() {
	if os.Getenv("INTEGRATION_TESTING") != "" {
		suite.T().Skip("Environmental Variables not set")
	}
	suite.True(suite.connector != nil)
}

// TestExchangeDataCollection is a call to the M365 backstore to very
func (suite *GraphConnectorTestSuite) TestExchangeDataCollection() {
	if os.Getenv("INTEGRATION_TESTING") != "" {
		suite.T().Skip("Environmental Variables not set")
	}
	exchangeData, err := suite.connector.ExchangeDataCollection("dustina@8qzvrj.onmicrosoft.com")
	assert.NotNil(suite.T(), exchangeData)
	assert.Error(suite.T(), err) // TODO Remove after https://github.com/alcionai/corso/issues/140
	suite.T().Logf("Missing Data: %s\n", err.Error())
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
		gc, err := NewGraphConnector(test.params[0], test.params[1], test.params[2])
		assert.Nil(suite.T(), gc, test.name+" failed")
		assert.NotNil(suite.T(), err, test.name+"failed")
	}
}
func Contains(elems []string, value string) bool {
	for _, s := range elems {
		if value == s {
			return true
		}
	}
	return false
}

func (suite *DiconnectedGraphConnectorTestSuite) TestBuild() {
	names := make(map[string]string)
	names["Al"] = "Bundy"
	names["Ellen"] = "Ripley"
	names["Axel"] = "Foley"
	first := buildFromMap(true, names)
	last := buildFromMap(false, names)
	suite.True(Contains(first, "Al"))
	suite.True(Contains(first, "Ellen"))
	suite.True(Contains(first, "Axel"))
	suite.True(Contains(last, "Bundy"))
	suite.True(Contains(last, "Ripley"))
	suite.True(Contains(last, "Foley"))

}

func (suite *DiconnectedGraphConnectorTestSuite) TestInterfaceAlignment() {
	var dc DataCollection
	concrete := NewExchangeDataCollection("Check", []string{"interface", "works"})
	dc = &concrete
	assert.NotNil(suite.T(), dc)
}
