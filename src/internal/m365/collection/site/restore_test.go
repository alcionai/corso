package site

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
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

	testName, lrh, destName, mockData := setupDependencies(
		suite,
		suite.ac,
		suite.siteID,
		suite.creds,
		"genericList")

	deets, err := restoreListItem(ctx, lrh, mockData, suite.siteID, destName)
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, fmt.Sprintf("%s_%s", destName, testName), deets.SharePoint.List.Name)

	// Clean-Up
	deleteList(ctx, t, suite.siteID, lrh, deets)
}

func (suite *SharePointRestoreSuite) TestListCollection_Restore_invalidListTemplate() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	_, lrh, destName, mockData := setupDependencies(
		suite,
		suite.ac,
		suite.siteID,
		suite.creds,
		api.WebTemplateExtensionsListTemplateName)

	_, err := restoreListItem(ctx, lrh, mockData, suite.siteID, destName)
	require.Error(t, err)
	assert.Contains(t, err.Error(), api.ErrCannotCreateWebTemplateExtension.Error())
}

func deleteList(
	ctx context.Context,
	t *testing.T,
	siteID string,
	lrh listsRestoreHandler,
	deets details.ItemInfo,
) {
	var (
		isFound  bool
		deleteID string
	)

	lists, err := lrh.ac.Client.
		Lists().
		GetLists(ctx, siteID, api.CallConfig{})
	assert.NoError(t, err, "getting site lists", clues.ToCore(err))

	for _, l := range lists {
		if ptr.Val(l.GetDisplayName()) == deets.SharePoint.ItemName {
			isFound = true
			deleteID = ptr.Val(l.GetId())

			break
		}
	}

	if isFound {
		err := lrh.DeleteList(ctx, deleteID)
		assert.NoError(t, err, clues.ToCore(err))
	}
}

func setupDependencies(
	suite tester.Suite,
	ac api.Client,
	siteID string,
	creds account.M365Config,
	listTemplate string) (
	string, listsRestoreHandler, string, *dataMock.Item,
) {
	t := suite.T()
	testName := "MockListing"

	lrh := NewListsRestoreHandler(siteID, ac.Lists())

	service := createTestService(t, creds)

	listInfo := models.NewListInfo()
	listInfo.SetTemplate(ptr.To(listTemplate))

	listing := spMock.ListDefault("Mock List")
	listing.SetDisplayName(&testName)
	listing.SetList(listInfo)

	byteArray, err := service.Serialize(listing)
	require.NoError(t, err, clues.ToCore(err))

	destName := testdata.DefaultRestoreConfig("").Location

	listData, err := data.NewPrefetchedItemWithInfo(
		io.NopCloser(bytes.NewReader(byteArray)),
		testName,
		details.ItemInfo{SharePoint: api.ListToSPInfo(listing)})
	require.NoError(t, err, clues.ToCore(err))

	r, err := readers.NewVersionedRestoreReader(listData.ToReader())
	require.NoError(t, err)

	mockData := &dataMock.Item{
		ItemID: testName,
		Reader: r,
	}

	return testName, lrh, destName, mockData
}
