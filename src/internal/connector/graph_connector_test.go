package connector_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	graph "github.com/alcionai/corso/internal/connector"
	ctesting "github.com/alcionai/corso/internal/testing"
)

type GraphConnectorTestSuite struct {
	suite.Suite
	connector *graph.GraphConnector
}

func TestGraphConnectorSuite(t *testing.T) {
	if err := ctesting.RunOnAny(
		ctesting.CORSO_CI_TESTS,
		ctesting.CORSO_GRAPH_CONNECTOR_TESTS,
	); err != nil {
		t.Skip(err)
	}
	suite.Run(t, new(GraphConnectorTestSuite))
}

func (suite *GraphConnectorTestSuite) SetupSuite() {
	if os.Getenv("CI") == "" {
		evs, err := ctesting.RequireEnvVars("TENANT_ID", "CLIENT_ID", "CLIENT_SECRET")
		if err != nil {
			suite.T().Fatal(err)
		}
		suite.connector, err = graph.NewGraphConnector(
			evs["TENANT_ID"],
			evs["CLIENT_ID"],
			evs["CLIENT_SECRET"])
		assert.Nil(suite.T(), err)
	}
}

func (suite *GraphConnectorTestSuite) TestGraphConnector() {
	suite.True(suite.connector != nil)
}

// --------------------

type DiconnectedGraphConnectorTestSuite struct {
	suite.Suite
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
