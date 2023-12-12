package site

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/sites"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/readers"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	spMock "github.com/alcionai/corso/src/internal/m365/service/sharepoint/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type SharePointRestoreSuite struct {
	tester.Suite
	siteID string
	creds  account.M365Config
	ac     api.Client
}

func (suite *SharePointRestoreSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()
	graph.InitializeConcurrencyLimiter(ctx, false, 4)

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

func TestSharePointRestoreSuite(t *testing.T) {
	suite.Run(t, &SharePointRestoreSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

// TestRestoreListCollection verifies Graph Restore API for the List Collection
func (suite *SharePointRestoreSuite) TestListCollection_Restore() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	service := createTestService(t, suite.creds)
	listing := spMock.ListDefault("Mock List")
	testName := "MockListing"
	listing.SetDisplayName(&testName)
	byteArray, err := service.Serialize(listing)
	require.NoError(t, err, clues.ToCore(err))

	destName := testdata.DefaultRestoreConfig("").Location

	listData, err := data.NewPrefetchedItemWithInfo(
		io.NopCloser(bytes.NewReader(byteArray)),
		testName,
		details.ItemInfo{SharePoint: ListToSPInfo(listing, int64(len(byteArray)))})
	require.NoError(t, err, clues.ToCore(err))

	r, err := readers.NewVersionedRestoreReader(listData.ToReader())
	require.NoError(t, err)

	mockData := &dataMock.Item{
		ItemID: testName,
		Reader: r,
	}

	lrh := NewListsRestoreHandler(suite.siteID, suite.ac.Lists())
	deets, err := restoreListItem(ctx, lrh, mockData, suite.siteID, destName)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, fmt.Sprintf("%s_%s", destName, testName), deets.SharePoint.ItemName)

	// Clean-Up
	var (
		builder  = service.Client().Sites().BySiteId(suite.siteID).Lists()
		isFound  bool
		deleteID string
	)

	for {
		resp, err := builder.Get(ctx, nil)
		assert.NoError(t, err, "getting site lists", clues.ToCore(err))

		for _, temp := range resp.GetValue() {
			if ptr.Val(temp.GetDisplayName()) == deets.SharePoint.ItemName {
				isFound = true
				deleteID = ptr.Val(temp.GetId())

				break
			}
		}
		// Get Next Link
		link, ok := ptr.ValOK(resp.GetOdataNextLink())
		if !ok {
			break
		}

		builder = sites.NewItemListsRequestBuilder(link, service.Adapter())
	}

	if isFound {
		err := DeleteList(ctx, service, suite.siteID, deleteID)
		assert.NoError(t, err, clues.ToCore(err))
	}
}
