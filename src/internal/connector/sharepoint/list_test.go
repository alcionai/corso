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
	userID string
	creds  account.M365Config
}

func (suite *SharePointSuite) SetupSuite() {
	t := suite.T()
	suite.userID = tester.SecondaryM365UserID(suite.T())

	a := tester.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err)

	suite.creds = m365
}

func TestSharePointSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(SharePointSuite))
}

// Test LoadList --> Retrieves all data from backStore
// Functions tested:
// - fetchListItems()
// - fetchColumns()
// - fetchContentBaseTypes
// - fetchContentColumns()
// - fetchContentTypes()
// - fetchColumnLinks
// - fetchColumnPositions
func (suite *SharePointSuite) TestLoadList() {
	ctx, flush := tester.NewContext()
	defer flush()

	service, err := createTestService(suite.creds)
	require.NoError(suite.T(), err)

	lists, err := loadLists(ctx, service, "root")
	assert.Greater(suite.T(), len(lists), 0)
	assert.NoError(suite.T(), err)
}
