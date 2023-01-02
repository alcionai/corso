package sharepoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
)

type SharePointSuite struct {
	suite.Suite
	creds account.M365Config
}

func (suite *SharePointSuite) SetupSuite() {
	t := suite.T()
	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.creds = m365
}

func TestSharePointSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoGraphConnectorSharePointTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(SharePointSuite))
}

// Test LoadList --> Retrieves all data from backStore
// Functions tested:
// - fetchListItems()
// - fetchColumns()
// - fetchContentColumns()
// - fetchContentTypes()
// - fetchColumnLinks
// TODO: upgrade passed github.com/microsoftgraph/msgraph-sdk-go v0.40.0
// to verify if these 2 calls are valid
// - fetchContentBaseTypes
// - fetchColumnPositions
func (suite *SharePointSuite) TestLoadList() {
	ctx, flush := tester.NewContext()
	defer flush()

	t := suite.T()
	service, err := createTestService(suite.creds)
	require.NoError(t, err)

	tuples, err := preFetchLists(ctx, service, "root")
	require.NoError(t, err)

	job := []string{tuples[0].id}
	lists, err := loadSiteLists(ctx, service, "root", job)
	assert.NoError(t, err)
	assert.Greater(t, len(lists), 0)
	t.Logf("Length: %d\n", len(lists))
}
