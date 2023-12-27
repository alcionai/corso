package site

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	kioser "github.com/microsoft/kiota-serialization-json-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/m365/collection/site/mock"
	betaAPI "github.com/alcionai/corso/src/internal/m365/service/sharepoint/api"
	spMock "github.com/alcionai/corso/src/internal/m365/service/sharepoint/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type SharePointCollectionSuite struct {
	tester.Suite
	siteID string
	creds  account.M365Config
	ac     api.Client
}

func (suite *SharePointCollectionSuite) SetupSuite() {
	t := suite.T()

	suite.siteID = tconfig.M365SiteID(t)
	a := tconfig.NewM365Account(t)
	m365, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.creds = m365

	ac, err := api.NewClient(
		m365,
		control.DefaultOptions(),
		count.New())
	require.NoError(t, err, clues.ToCore(err))

	suite.ac = ac
}

func TestSharePointCollectionSuite(t *testing.T) {
	suite.Run(t, &SharePointCollectionSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

// TestListCollection tests basic functionality to create
// SharePoint collection and to use the data stream channel.
func (suite *SharePointCollectionSuite) TestCollection_Items() {
	var (
		tenant  = "some"
		user    = "user"
		dirRoot = "directory"
	)

	sel := selectors.NewSharePointBackup([]string{"site"})

	tables := []struct {
		name, itemName string
		scope          selectors.SharePointScope
		getter         getItemByIDer
		getDir         func(t *testing.T) path.Path
		getItem        func(t *testing.T, itemName string) data.Item
	}{
		{
			name:     "List",
			itemName: "MockListing",
			scope:    sel.Lists(selectors.Any())[0],
			getter:   &mock.ListHandler{},
			getDir: func(t *testing.T) path.Path {
				dir, err := path.Build(
					tenant,
					user,
					path.SharePointService,
					path.ListsCategory,
					false,
					dirRoot)
				require.NoError(t, err, clues.ToCore(err))

				return dir
			},
			getItem: func(t *testing.T, name string) data.Item {
				ow := kioser.NewJsonSerializationWriter()
				listing := spMock.ListDefault(name)
				listing.SetDisplayName(&name)

				err := ow.WriteObjectValue("", listing)
				require.NoError(t, err, clues.ToCore(err))

				byteArray, err := ow.GetSerializedContent()
				require.NoError(t, err, clues.ToCore(err))

				info := &details.SharePointInfo{
					ItemName: name,
				}

				data, err := data.NewPrefetchedItemWithInfo(
					io.NopCloser(bytes.NewReader(byteArray)),
					name,
					details.ItemInfo{SharePoint: info})
				require.NoError(t, err, clues.ToCore(err))

				return data
			},
		},
		{
			name:     "Pages",
			itemName: "MockPages",
			scope:    sel.Pages(selectors.Any())[0],
			getter:   nil,
			getDir: func(t *testing.T) path.Path {
				dir, err := path.Build(
					tenant,
					user,
					path.SharePointService,
					path.PagesCategory,
					false,
					dirRoot)
				require.NoError(t, err, clues.ToCore(err))

				return dir
			},
			getItem: func(t *testing.T, itemName string) data.Item {
				byteArray := spMock.Page(itemName)
				page, err := betaAPI.BytesToSitePageable(byteArray)
				require.NoError(t, err, clues.ToCore(err))

				data, err := data.NewPrefetchedItemWithInfo(
					io.NopCloser(bytes.NewReader(byteArray)),
					itemName,
					details.ItemInfo{SharePoint: betaAPI.PageInfo(page, int64(len(byteArray)))})
				require.NoError(t, err, clues.ToCore(err))

				return data
			},
		},
	}

	for _, test := range tables {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			col := NewCollection(
				test.getter,
				test.getDir(t),
				suite.ac,
				test.scope,
				nil,
				control.DefaultOptions())
			col.stream <- test.getItem(t, test.itemName)

			readItems := []data.Item{}

			for item := range col.Items(ctx, fault.New(true)) {
				readItems = append(readItems, item)
			}

			require.Equal(t, len(readItems), 1)
			item := readItems[0]
			shareInfo, ok := item.(data.ItemInfo)
			require.True(t, ok)

			info, err := shareInfo.Info()
			require.NoError(t, err, clues.ToCore(err))

			assert.NotNil(t, info)
			assert.NotNil(t, info.SharePoint)
			assert.Equal(t, test.itemName, info.SharePoint.ItemName)
		})
	}
}
