package connector

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	ctesting "github.com/alcionai/corso/internal/testing"
)

type GraphConnectorIntegrationSuite struct {
	suite.Suite
	connector *GraphConnector
}

func TestGraphConnectorSuite(t *testing.T) {
	if err := ctesting.RunOnAny(
		ctesting.CORSO_CI_TESTS,
		ctesting.CORSO_GRAPH_CONNECTOR_TESTS,
	); err != nil {
		t.Skip(err)
	}
	if err := ctesting.RunOnAny(
		"this-is-fake-it-forces-a-skip-until-we-fix-ci-details-here(rkeepers)",
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(GraphConnectorIntegrationSuite))
}

func (suite *GraphConnectorIntegrationSuite) SetupSuite() {
	evs, err := ctesting.GetRequiredEnvVars("TENANT_ID", "CLIENT_ID", "CLIENT_SECRET")
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.connector, err = NewGraphConnector(
		evs["TENANT_ID"],
		evs["CLIENT_ID"],
		evs["CLIENT_SECRET"])
	suite.NoError(err)
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector() {
	ctesting.LogTimeOfTest(suite.T())
	suite.NotNil(suite.connector)

}

// --------------------

type DiconnectedGraphConnectorSuite struct {
	suite.Suite
}

func TestDisconnectedGraphSuite(t *testing.T) {
	suite.Run(t, new(DiconnectedGraphConnectorSuite))
}

// TestExchangeDataCollection is a call to the M365 backstore to very
func (suite *GraphConnectorIntegrationSuite) TestExchangeDataCollection() {
	if os.Getenv("INTEGRATION_TESTING") != "" {
		suite.T().Skip("Environmental Variables not set")
	}
	exchangeData, err := suite.connector.ExchangeDataCollection("dustina@8qzvrj.onmicrosoft.com")
	assert.NotNil(suite.T(), exchangeData)
	assert.Error(suite.T(), err) // TODO Remove after https://github.com/alcionai/corso/issues/140
	suite.T().Logf("Missing Data: %s\n", err.Error())
}

func (suite *DiconnectedGraphConnectorSuite) TestBadConnection() {

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
		suite.T().Run(test.name, func(t *testing.T) {
			gc, err := NewGraphConnector(test.params[0], test.params[1], test.params[2])
			assert.Nil(t, gc, test.name+" failed")
			assert.NotNil(t, err, test.name+"failed")
		})
	}
}

// Contains is a helper method for verifying if element
// is contained within the slice
func Contains(elems []string, value string) bool {
	for _, s := range elems {
		if value == s {
			return true
		}
	}
	return false
}

func (suite *DiconnectedGraphConnectorSuite) TestBuild() {
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

func (suite *DiconnectedGraphConnectorSuite) TestInterfaceAlignment() {
	var dc DataCollection
	concrete := NewExchangeDataCollection("Check", []string{"interface", "works"})
	dc = &concrete
	assert.NotNil(suite.T(), dc)

}
