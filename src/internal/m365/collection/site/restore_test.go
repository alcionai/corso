package site

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/readers"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	siteMock "github.com/alcionai/corso/src/internal/m365/collection/site/mock"
	spMock "github.com/alcionai/corso/src/internal/m365/service/sharepoint/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/errs/core"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
)

type SharePointRestoreUnitSuite struct {
	tester.Suite
}

func TestSharePointRestoreUnitSuite(t *testing.T) {
	suite.Run(t, &SharePointRestoreUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SharePointCollectionUnitSuite) TestFormatListsRestoreDestination() {
	t := suite.T()

	dt := dttm.FormatNow(dttm.SafeForTesting)

	tests := []struct {
		name          string
		destName      string
		itemID        string
		getStoredList func() models.Listable
		expectedName  string
	}{
		{
			name:     "stored list has a display name",
			destName: "Corso_Restore_" + dt,
			itemID:   "someid",
			getStoredList: func() models.Listable {
				list := models.NewList()
				list.SetDisplayName(ptr.To("list1"))

				return list
			},
			expectedName: "Corso_Restore_" + dt + "_list1",
		},
		{
			name:     "stored list does not have a display name",
			destName: "Corso_Restore_" + dt,
			itemID:   "someid",
			getStoredList: func() models.Listable {
				return models.NewList()
			},
			expectedName: "Corso_Restore_" + dt + "_someid",
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			newName := formatListsRestoreDestination(test.destName, test.itemID, test.getStoredList())
			assert.Equal(t, test.expectedName, newName, "new name for list")
		})
	}
}

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

	var (
		listName     = "MockListing"
		listTemplate = "genericList"
		restoreCfg   = testdata.DefaultRestoreConfig("")
		destName     = restoreCfg.Location
		lrh          = NewListsRestoreHandler(suite.siteID, suite.ac.Lists())
		service      = createTestService(t, suite.creds)
		list         = createList(listTemplate, listName)
		mockData     = generateListData(t, service, list)
	)

	restoreCfg.OnCollision = control.Copy

	deets, err := restoreListItem(
		ctx,
		lrh,
		mockData,
		suite.siteID,
		restoreCfg,
		nil,
		count.New(),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.Equal(t, fmt.Sprintf("%s_%s", destName, listName), deets.SharePoint.List.Name)

	// Clean-Up
	deleteList(ctx, t, suite.siteID, lrh, deets)
}

func (suite *SharePointRestoreSuite) TestListCollection_Restore_invalidListTemplate() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		lrh        = NewListsRestoreHandler(suite.siteID, suite.ac.Lists())
		listName   = "MockListing"
		restoreCfg = testdata.DefaultRestoreConfig("")
		service    = createTestService(t, suite.creds)
	)

	restoreCfg.OnCollision = control.Copy

	tests := []struct {
		name   string
		list   models.Listable
		expect assert.ErrorAssertionFunc
	}{
		{
			name:   "list with template documentLibrary",
			list:   createList(api.DocumentLibraryListTemplate, listName),
			expect: assert.Error,
		},
		{
			name:   "list with template webTemplateExtensionsList",
			list:   createList(api.WebTemplateExtensionsListTemplate, listName),
			expect: assert.Error,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			listData := generateListData(t, service, test.list)

			_, err := restoreListItem(
				ctx,
				lrh,
				listData,
				suite.siteID,
				restoreCfg,
				nil,
				count.New(),
				fault.New(false))
			require.Error(t, err)
			assert.Contains(t, err.Error(), api.ErrSkippableListTemplate.Error())
		})
	}
}

func (suite *SharePointRestoreSuite) TestListCollection_RestoreInPlace_skip() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		listName     = "MockListing"
		listTemplate = "genericList"
		restoreCfg   = testdata.DefaultRestoreConfig("")
		lrh          = NewListsRestoreHandler(suite.siteID, suite.ac.Lists())
		service      = createTestService(t, suite.creds)
		list         = createList(listTemplate, listName)
		newList      = createList(listTemplate, listName)
		cl           = count.New()
	)

	mockData := generateListData(t, service, list)

	collisionKeyToItemID := map[string]string{
		api.ListCollisionKey(newList): "some-list-id",
	}

	deets, err := restoreListItem(
		ctx,
		lrh,
		mockData,
		suite.siteID,
		restoreCfg, // OnCollision is skip by default
		collisionKeyToItemID,
		cl,
		fault.New(true))
	require.Error(t, err, clues.ToCore(err))
	assert.True(t, errors.Is(core.ErrAlreadyExists, err))
	assert.Empty(t, deets)
	assert.Less(t, int64(0), cl.Get(count.CollisionSkip))
}

func (suite *SharePointRestoreSuite) TestListCollection_RestoreInPlace_copy() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		listName     = "MockListing"
		listTemplate = "genericList"
		listID       = "some-list-id"
		restoreCfg   = testdata.DefaultRestoreConfig("")
		lrh          = NewListsRestoreHandler(suite.siteID, suite.ac.Lists())
		service      = createTestService(t, suite.creds)
		cl           = count.New()
	)

	restoreCfg.OnCollision = control.Copy

	list := createList(listTemplate, listName)
	list.SetId(ptr.To(listID))

	newList := createList(listTemplate, listName)
	newList.SetId(ptr.To(listID))

	mockData := generateListData(t, service, list)

	collisionKeyToItemID := map[string]string{
		api.ListCollisionKey(newList): listID,
	}

	deets, err := restoreListItem(
		ctx,
		lrh,
		mockData,
		suite.siteID,
		restoreCfg,
		collisionKeyToItemID,
		cl,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.Zero(t, cl.Get(count.CollisionSkip))
	assert.Zero(t, cl.Get(count.CollisionReplace))

	// Clean-Up
	deleteList(ctx, t, suite.siteID, lrh, deets)
}

func (suite *SharePointRestoreSuite) TestListCollection_RestoreInPlace_replace() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		listName     = "MockListing"
		listTemplate = "genericList"
		restoreCfg   = testdata.DefaultRestoreConfig("")
		lrh          = NewListsRestoreHandler(suite.siteID, suite.ac.Lists())
		service      = createTestService(t, suite.creds)
		cl           = count.New()
		now          = time.Now()
		temMinAfter  = now.Add(10 * time.Minute)
	)

	restoreCfg.OnCollision = control.Replace

	list := createList(listTemplate, listName)
	list.SetLastModifiedDateTime(ptr.To(now))

	newList := createList(listTemplate, listName)
	newList.SetLastModifiedDateTime(ptr.To(temMinAfter))

	createdList, err := lrh.PostList(
		ctx,
		listName,
		newList,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	listID := ptr.Val(createdList.GetId())
	list.SetId(ptr.To(listID))

	mockData := generateListData(t, service, list)

	collisionKeyToItemID := map[string]string{
		api.ListCollisionKey(newList): listID,
	}

	deets, err := restoreListItem(
		ctx,
		lrh,
		mockData,
		suite.siteID,
		restoreCfg,
		collisionKeyToItemID,
		cl,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.Less(t, int64(0), cl.Get(count.CollisionReplace))

	// Clean-Up
	deleteList(ctx, t, suite.siteID, lrh, deets)
}

func (suite *SharePointRestoreSuite) TestListCollection_RestoreInPlace_replaceFails() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		listName     = "MockListing"
		listTemplate = "genericList"
		listID       = "some-list-id"
		restoreCfg   = testdata.DefaultRestoreConfig("")
		service      = createTestService(t, suite.creds)
		cl           = count.New()
	)

	restoreCfg.OnCollision = control.Replace

	list := createList(listTemplate, listName)
	list.SetId(ptr.To(listID))

	newList := createList(listTemplate, listName)
	newList.SetId(ptr.To(listID))

	collisionKeyToItemID := map[string]string{
		api.ListCollisionKey(newList): listID,
	}

	tests := []struct {
		name                 string
		lrh                  *siteMock.ListRestoreHandler
		expectErr            assert.ErrorAssertionFunc
		expectCollisionCount int64
	}{
		{
			name: "PostList fails for stored list",
			lrh: siteMock.NewListRestoreHandler(
				nil,
				errors.New("failed to create list"),
				nil),
			expectErr: assert.Error,
		},
		{
			name: "DeleteList fails",
			lrh: siteMock.NewListRestoreHandler(
				errors.New("failed to delete list"),
				nil,
				nil),
			expectErr: assert.Error,
		},
		{
			name: "PatchList fails",
			lrh: siteMock.NewListRestoreHandler(
				nil,
				nil,
				errors.New("failed to patch list")),
			expectErr: assert.Error,
		},
		{
			name: "PostList passes for stored list",
			lrh: siteMock.NewListRestoreHandler(
				nil,
				nil,
				nil),
			expectErr:            assert.NoError,
			expectCollisionCount: 1,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			mockData := generateListData(t, service, list)

			_, err := restoreListItem(
				ctx,
				test.lrh,
				mockData,
				suite.siteID,
				restoreCfg,
				collisionKeyToItemID,
				cl,
				fault.New(true))
			test.expectErr(t, err)
			assert.Equal(t, test.expectCollisionCount, cl.Get(count.CollisionReplace))
		})
	}
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
		if ptr.Val(l.GetDisplayName()) == deets.SharePoint.List.Name {
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

func generateListData(
	t *testing.T,
	service *graph.Service,
	list models.Listable,
) *dataMock.Item {
	listName := ptr.Val(list.GetDisplayName())

	byteArray, err := service.Serialize(list)
	require.NoError(t, err, clues.ToCore(err))

	listData, err := data.NewPrefetchedItemWithInfo(
		io.NopCloser(bytes.NewReader(byteArray)),
		listName,
		details.ItemInfo{SharePoint: api.ListToSPInfo(list)})
	require.NoError(t, err, clues.ToCore(err))

	r, err := readers.NewVersionedRestoreReader(listData.ToReader())
	require.NoError(t, err)

	mockData := &dataMock.Item{
		ItemID: listName,
		Reader: r,
	}

	return mockData
}

func createList(listTemplate, listDisplayName string) models.Listable {
	listInfo := models.NewListInfo()
	listInfo.SetTemplate(ptr.To(listTemplate))

	listing := spMock.ListDefault("Mock List")
	listing.SetDisplayName(ptr.To(listDisplayName))
	listing.SetList(listInfo)

	return listing
}
