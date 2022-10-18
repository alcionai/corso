package exchange

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/tester"
)

type ContactFolderCacheIntegrationSuite struct {
	suite.Suite
	gs graph.Service
}

func (suite *ContactFolderCacheIntegrationSuite) SetupSuite() {
	t := suite.T()

	_, err := tester.GetRequiredEnvVars(tester.M365AcctCredEnvs...)
	require.NoError(t, err)

	a := tester.NewM365Account(t)
	require.NoError(t, err)

	m365, err := a.M365Config()
	require.NoError(t, err)

	service, err := createService(m365, false)
	require.NoError(t, err)

	suite.gs = service
}

func TestContactFolderCacheIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorTests,
		tester.CorsoGraphConnectorExchangeTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(ContactFolderCacheIntegrationSuite))
}

func (suite *ContactFolderCacheIntegrationSuite) TestPopulate() {
	ctx, flush := tester.NewContext()
	defer flush()

	cfc := contactFolderCache{
		userID: tester.M365UserID(suite.T()),
		gs:     suite.gs,
	}

	tests := []struct {
		name       string
		folderName string
		basePath   string
		canFind    assert.BoolAssertionFunc
	}{
		{
			name:       "Default Contact Cache",
			folderName: DefaultContactFolder,
			basePath:   DefaultContactFolder,
			canFind:    assert.True,
		},
		{
			name:       "Default Contact Hidden",
			folderName: DefaultContactFolder,
			canFind:    assert.False,
		},
		{
			name:       "Name Not in Cache",
			folderName: "testFooBarWhoBar",
			canFind:    assert.False,
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			require.NoError(t, cfc.Populate(ctx, DefaultContactFolder, test.basePath))
			_, isFound := cfc.PathInCache(test.folderName)
			test.canFind(t, isFound)
		})
	}
}
