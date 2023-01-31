package graph

import (
	"net/http"
	"testing"
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

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
	adpt, err := CreateAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret)

	assert.NoError(t, err)
	assert.NotNil(t, adpt)
}

func (suite *GraphUnitSuite) TestHTTPClient() {
	table := []struct {
		name  string
		opts  []option
		check func(*testing.T, *http.Client)
	}{
		{
			name: "no options",
			opts: []option{},
			check: func(t *testing.T, c *http.Client) {
				assert.Equal(t, 90*time.Second, c.Timeout, "default timeout")
			},
		},
		{
			name: "no timeout",
			opts: []option{NoTimeout()},
			check: func(t *testing.T, c *http.Client) {
				assert.Equal(t, 0, int(c.Timeout), "unlimited timeout")
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			cli := HTTPClient(test.opts...)
			assert.NotNil(t, cli)
			test.check(t, cli)
		})
	}
}

func (suite *GraphUnitSuite) TestSerializationEndPoint() {
	t := suite.T()
	adpt, err := CreateAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret)
	require.NoError(t, err)

	serv := NewService(adpt)
	email := models.NewMessage()
	subject := "TestSerializationEndPoint"
	email.SetSubject(&subject)

	byteArray, err := serv.Serialize(email)
	assert.NoError(t, err)
	assert.NotNil(t, byteArray)
	t.Log(string(byteArray))
}
