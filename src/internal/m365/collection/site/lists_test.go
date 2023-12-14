package site

import (
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
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

	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, false, 4)
}

func TestListsUnitSuite(t *testing.T) {
	suite.Run(t, &ListsUnitSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
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
		},
		{
			name: "Only Name",
			listAndDeets: func() (models.Listable, *details.SharePointInfo) {
				aTitle := "Whole List"
				listTemplate := "genericList"
				listItemName := "listItem1"

				listInfo := models.NewListInfo()
				listInfo.SetTemplate(ptr.To(listTemplate))

				listing := models.NewList()
				listing.SetDisplayName(ptr.To(aTitle))
				listing.SetList(listInfo)

				li := models.NewListItem()
				li.SetId(ptr.To(listItemName))

				listing.SetItems([]models.ListItemable{li})
				i := &details.SharePointInfo{
					ItemType:     details.SharePointList,
					ItemName:     aTitle,
					ItemCount:    1,
					ItemTemplate: listTemplate,
				}

				return listing, i
			},
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			list, expected := test.listAndDeets()
			info := ListToSPInfo(list)
			assert.Equal(t, expected.ItemType, info.ItemType)
			assert.Equal(t, expected.ItemName, info.ItemName)
			assert.Equal(t, expected.ItemCount, info.ItemCount)
			assert.Equal(t, expected.WebURL, info.WebURL)
		})
	}
}
