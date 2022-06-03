package connector_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	graph "github.com/alcionai/corso/internal/connector"
	ctesting "github.com/alcionai/corso/internal/testing"
)

type GraphConnectorIntegrationSuite struct {
	suite.Suite
	connector *graph.GraphConnector
}

func TestGraphConnectorSuite(t *testing.T) {
	if err := ctesting.RunOnAny(
		ctesting.CORSO_CI_TESTS,
		ctesting.CORSO_GRAPH_CONNECTOR_TESTS,
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
	suite.connector, err = graph.NewGraphConnector(
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
			gc, err := graph.NewGraphConnector(test.params[0], test.params[1], test.params[2])
			assert.Nil(t, gc, test.name+" failed")
			assert.NotNil(t, err, test.name+"failed")
		})
	}
}
