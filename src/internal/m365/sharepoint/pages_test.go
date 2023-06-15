package sharepoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/m365/graph/betasdk/models"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

type PagesUnitSuite struct {
	tester.Suite
}

func TestPagesUnitSuite(t *testing.T) {
	suite.Run(t, &PagesUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *PagesUnitSuite) TestSharePointInfo_Pages() {
	tests := []struct {
		name         string
		pageAndDeets func() (models.SitePageable, *details.SharePointInfo)
	}{
		{
			name: "Empty Page",
			pageAndDeets: func() (models.SitePageable, *details.SharePointInfo) {
				deets := &details.SharePointInfo{ItemType: details.SharePointPage}
				return models.NewSitePage(), deets
			},
		},
		{
			name: "Only Name",
			pageAndDeets: func() (models.SitePageable, *details.SharePointInfo) {
				title := "Blank Page"
				sPage := models.NewSitePage()
				sPage.SetTitle(&title)
				deets := &details.SharePointInfo{
					ItemType: details.SharePointPage,
					ItemName: title,
				}

				return sPage, deets
			},
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			paged, expected := test.pageAndDeets()
			info := pageToSPInfo(paged, "", 0)
			assert.Equal(t, expected.ItemType, info.ItemType)
			assert.Equal(t, expected.ItemName, info.ItemName)
			assert.Equal(t, expected.WebURL, info.WebURL)
		})
	}
}
