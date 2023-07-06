package sharepoint

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
)

type ListsUnitSuite struct {
	tester.Suite
	creds account.M365Config
}

func (suite *ListsUnitSuite) SetupSuite() {
	t := suite.T()
	a := tconfig.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.creds = m365
}

func TestListsUnitSuite(t *testing.T) {
	suite.Run(t, &ListsUnitSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs},
		),
	})
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
func (suite *ListsUnitSuite) TestLoadList() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	service := createTestService(t, suite.creds)
	tuples, err := preFetchLists(ctx, service, "root")
	require.NoError(t, err, clues.ToCore(err))

	job := []string{tuples[0].id}
	lists, err := loadSiteLists(ctx, service, "root", job, fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	assert.Greater(t, len(lists), 0)
	t.Logf("Length: %d\n", len(lists))
}

func (suite *ListsUnitSuite) TestSharePointInfo() {
	tests := []struct {
		name         string
		listAndDeets func() (models.Listable, *details.SharePointInfo)
	}{
		{
			name: "Empty List",
			listAndDeets: func() (models.Listable, *details.SharePointInfo) {
				i := &details.SharePointInfo{ItemType: details.SharePointList}
				return models.NewList(), i
			},
		}, {
			name: "Only Name",
			listAndDeets: func() (models.Listable, *details.SharePointInfo) {
				aTitle := "Whole List"
				listing := models.NewList()
				listing.SetDisplayName(&aTitle)
				i := &details.SharePointInfo{
					ItemType: details.SharePointList,
					ItemName: aTitle,
				}

				return listing, i
			},
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			list, expected := test.listAndDeets()
			info := listToSPInfo(list, 10)
			assert.Equal(t, expected.ItemType, info.ItemType)
			assert.Equal(t, expected.ItemName, info.ItemName)
			assert.Equal(t, expected.WebURL, info.WebURL)
		})
	}
}
