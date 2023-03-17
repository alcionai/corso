package graph

import (
	"net/http"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/clues"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

type GraphUnitSuite struct {
	tester.Suite
	credentials account.M365Config
}

func TestGraphUnitSuite(t *testing.T) {
	suite.Run(t, &GraphUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *GraphUnitSuite) SetupSuite() {
	t := suite.T()
	a := tester.NewMockM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.credentials = m365
}

func (suite *GraphUnitSuite) TestCreateAdapter() {
	t := suite.T()
	adpt, err := CreateAdapter(
		suite.credentials.AzureTenantID,
		suite.credentials.AzureClientID,
		suite.credentials.AzureClientSecret)

	assert.NoError(t, err, clues.ToCore(err))
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
				assert.Equal(t, defaultHTTPClientTimeout, c.Timeout, "default timeout")
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
		suite.Run(test.name, func() {
			t := suite.T()

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
	require.NoError(t, err, clues.ToCore(err))

	serv := NewService(adpt)
	email := models.NewMessage()
	subject := "TestSerializationEndPoint"
	email.SetSubject(&subject)

	byteArray, err := serv.Serialize(email)
	assert.NoError(t, err, clues.ToCore(err))
	assert.NotNil(t, byteArray)
	t.Log(string(byteArray))
}
