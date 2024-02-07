package details

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/canario/src/internal/tester"
)

type ItemInfoUnitSuite struct {
	tester.Suite
}

func TestItemInfoUnitSuite(t *testing.T) {
	suite.Run(t, &ItemInfoUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *ItemInfoUnitSuite) TestItemInfo_IsDriveItem() {
	table := []struct {
		name   string
		ii     ItemInfo
		expect assert.BoolAssertionFunc
	}{
		{
			name: "onedrive item",
			ii: ItemInfo{
				OneDrive: &OneDriveInfo{
					ItemType: OneDriveItem,
				},
			},
			expect: assert.True,
		},
		{
			name: "sharepoint library",
			ii: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType: SharePointLibrary,
				},
			},
			expect: assert.True,
		},
		{
			name: "sharepoint page",
			ii: ItemInfo{
				SharePoint: &SharePointInfo{
					ItemType: SharePointPage,
				},
			},
			expect: assert.False,
		},
		{
			name: "groups library",
			ii: ItemInfo{
				Groups: &GroupsInfo{
					ItemType: SharePointLibrary,
				},
			},
			expect: assert.True,
		},
		{
			name: "groups channel message",
			ii: ItemInfo{
				Groups: &GroupsInfo{
					ItemType: GroupsChannelMessage,
				},
			},
			expect: assert.False,
		},
		{
			name: "exchange anything",
			ii: ItemInfo{
				Exchange: &ExchangeInfo{
					ItemType: ExchangeMail,
				},
			},
			expect: assert.False,
		},
		{
			name: "teams chat",
			ii: ItemInfo{
				TeamsChats: &TeamsChatsInfo{
					ItemType: TeamsChat,
				},
			},
			expect: assert.False,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			test.expect(suite.T(), test.ii.isDriveItem())
		})
	}
}
