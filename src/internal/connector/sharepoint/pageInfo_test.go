package sharepoint

import (
	"testing"

	bmodel "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
	"github.com/stretchr/testify/assert"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

func (suite *SharePointInfoSuite) TestSharePointInfo_Pages() {
	tests := []struct {
		name         string
		pageAndDeets func() (bmodel.SitePageable, *details.SharePointInfo)
	}{
		{
			name: "Empty Page",
			pageAndDeets: func() (bmodel.SitePageable, *details.SharePointInfo) {
				deets := &details.SharePointInfo{ItemType: details.SharePointItem}
				return bmodel.NewSitePage(), deets
			},
		},
		{
			name: "Only Name",
			pageAndDeets: func() (bmodel.SitePageable, *details.SharePointInfo) {
				title := "Blank Page"
				sPage := bmodel.NewSitePage()
				sPage.SetTitle(&title)
				deets := &details.SharePointInfo{
					ItemType: details.SharePointItem,
					ItemName: title,
				}

				return sPage, deets
			},
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			paged, expected := test.pageAndDeets()
			info := sharePointPageInfo(paged, 0)
			assert.Equal(t, expected.ItemType, info.ItemType)
			assert.Equal(t, expected.ItemName, info.ItemName)
			assert.Equal(t, expected.WebURL, info.WebURL)
		})
	}
}
