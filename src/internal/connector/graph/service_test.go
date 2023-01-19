package graph_test

import (
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

type GraphUnitSuite struct {
	suite.Suite
	credentials account.M365Config
}

func TestGraphUnitSuite(t *testing.T) {
	suite.Run(t, new(GraphUnitSuite))
}

func (suite *GraphUnitSuite) SetupSuite() {
	t := suite.T()
	a := tester.NewMockM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.credentials = m365
}

func (suite *GraphUnitSuite) TestCreateAdapter() {
	t := suite.T()
	adpt, err := graph.CreateAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret,
	)

	assert.NoError(t, err)
	assert.NotNil(t, adpt)
}

func (suite *GraphUnitSuite) TestBetaService() {
	t := suite.T()
	adpt, err := graph.CreateBetaAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret,
	)

	assert.NoError(t, err)
	require.NotNil(t, adpt)

	serv := graph.NewBetaService(adpt)
	assert.NotNil(t, serv)
}

func (suite *GraphUnitSuite) TestSerializationEndPoint() {
	t := suite.T()
	adpt, err := graph.CreateAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret,
	)
	require.NoError(t, err)

	serv := graph.NewService(adpt)
	email := models.NewMessage()
	subject := "TestSerializationEndPoint"
	email.SetSubject(&subject)

	byteArray, err := serv.Serialize(email)
	assert.NoError(t, err)
	assert.NotNil(t, byteArray)
	t.Log(string(byteArray))
}
