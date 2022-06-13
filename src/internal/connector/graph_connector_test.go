package connector

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/internal/connector/support"
	ctesting "github.com/alcionai/corso/internal/testing"
	"github.com/alcionai/corso/pkg/credentials"
)

type GraphConnectorIntegrationSuite struct {
	suite.Suite
	connector *GraphConnector
}

func TestGraphConnectorSuite(t *testing.T) {
	if err := ctesting.RunOnAny(
		ctesting.CorsoCITests,
		ctesting.CorsoGraphConnectorTests,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(GraphConnectorIntegrationSuite))
}

func (suite *GraphConnectorIntegrationSuite) SetupSuite() {
	evs, err := ctesting.GetRequiredEnvVars(credentials.TenantID, credentials.ClientID, credentials.ClientSecret)
	require.NoError(suite.T(), err)

	suite.connector, err = NewGraphConnector(
		evs[credentials.TenantID],
		evs[credentials.ClientID],
		evs[credentials.ClientSecret])
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

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_setTenantUsers() {
	err := suite.connector.setTenantUsers()
	assert.NoError(suite.T(), err)
	suite.Greater(len(suite.connector.Users), 0)
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_ExchangeDataCollection() {
	if os.Getenv("INTEGRATION_TESTING") != "" {
		suite.T().Skip("Environmental Variables not set")
	}
	exchangeData, err := suite.connector.ExchangeDataCollection("lidiah@8qzvrj.onmicrosoft.com")
	assert.NotNil(suite.T(), exchangeData)
	assert.Error(suite.T(), err) // TODO Remove after https://github.com/alcionai/corso/issues/140
	if err != nil {
		suite.T().Logf("Missing Data: %s\n", err.Error())
	}
	suite.T().Logf("Full PathData: %s\n", exchangeData.FullPath())
}

func (suite *GraphConnectorIntegrationSuite) TestGraphConnector_restoreMessages() {
	user := "TEST_GRAPH_USER" // user.GetId()
	file := "TEST_GRAPH_FILE" // Test file should be sent or received by the user
	evs, err := ctesting.GetRequiredEnvVars(user, file)
	if err != nil {
		suite.T().Skipf("Environment not configured: %v\n", err)
	}
	bytes, err := support.LoadAFile(evs[file]) // TEST_GRAPH_FILE should have a single Message && not present in target inbox
	if err != nil {
		suite.T().Skipf("Support file not accessible: %v\n", err)
	}
	ds := ExchangeData{id: "test", message: bytes}
	edc := NewExchangeDataCollection("tenant", []string{evs[user], "dustina@@8qzvrj.onmicrosoft.com", "Inbox"})
	edc.PopulateCollection(ds)
	edc.FinishPopulation()
	err = suite.connector.restoreMessages(&edc)
	assert.NoError(suite.T(), err)
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
