package sharepoint

import (
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

type SharePointInfoSuite struct {
	suite.Suite
}

func TestSharePointInfoSuite(t *testing.T) {
	suite.Run(t, new(SharePointInfoSuite))
}

func (suite *SharePointInfoSuite) TestSharePointInfo() {
	tests := []struct {
		name         string
		listAndDeets func() (models.Listable, *details.SharePointInfo)
	}{
		{
			name: "Empty List",
			listAndDeets: func() (models.Listable, *details.SharePointInfo) {
				i := &details.SharePointInfo{ItemType: details.SharePointItem}
				return models.NewList(), i
			},
		}, {
			name: "Only Name",
			listAndDeets: func() (models.Listable, *details.SharePointInfo) {
				aTitle := "Whole List"
				listing := models.NewList()
				listing.SetDisplayName(&aTitle)
				i := &details.SharePointInfo{
					ItemType: details.SharePointItem,
					ItemName: aTitle,
				}

				return listing, i
			},
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			list, expected := test.listAndDeets()
			info := sharePointListInfo(list, 10)
			assert.Equal(t, expected.ItemType, info.ItemType)
			assert.Equal(t, expected.ItemName, info.ItemName)
			assert.Equal(t, expected.WebURL, info.WebURL)
		})
	}
}
