package m365

import (
	"bytes"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/service/exchange"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/service/sharepoint"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// DataCollection tests
// ---------------------------------------------------------------------------

type DataCollectionIntgSuite struct {
	tester.Suite
	user     string
	site     string
	tenantID string
	ac       api.Client
}

func TestDataCollectionIntgSuite(t *testing.T) {
	suite.Run(t, &DataCollectionIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *DataCollectionIntgSuite) SetupSuite() {
	t := suite.T()

	suite.user = tconfig.M365UserID(t)
	suite.site = tconfig.M365SiteID(t)

	acct := tconfig.NewM365Account(t)
	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.tenantID = creds.AzureTenantID

	suite.ac, err = api.NewClient(
		creds,
		control.DefaultOptions(),
		count.New())
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *DataCollectionIntgSuite) TestExchangeDataCollection() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	selUsers := []string{suite.user}

	ctrl := newController(ctx, suite.T(), path.ExchangeService)
	tests := []struct {
		name        string
		getSelector func(t *testing.T) selectors.Selector
	}{
		{
			name: "Email",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup(selUsers)
				sel.Include(sel.MailFolders([]string{api.MailInbox}, selectors.PrefixMatch()))
				sel.DiscreteOwner = suite.user
				return sel.Selector
			},
		},
		{
			name: "Contacts",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup(selUsers)
				sel.Include(sel.ContactFolders([]string{api.DefaultContacts}, selectors.PrefixMatch()))
				sel.DiscreteOwner = suite.user
				return sel.Selector
			},
		},
		// {
		// 	name: "Events",
		// 	getSelector: func(t *testing.T) selectors.Selector {
		// 		sel := selectors.NewExchangeBackup(selUsers)
		// 		sel.Include(sel.EventCalendars([]string{exchange.DefaultCalendar}, selectors.PrefixMatch()))
		// 		sel.DiscreteOwner = suite.user
		// 		return sel.Selector
		// 	},
		// },
	}

	for _, test := range tests {
		for _, canMakeDeltaQueries := range []bool{true, false} {
			name := test.name

			if canMakeDeltaQueries {
				name += "-delta"
			} else {
				name += "-non-delta"
			}

			suite.Run(name, func() {
				t := suite.T()

				ctx, flush := tester.NewContext(t)
				defer flush()

				sel := test.getSelector(t)
				uidn := inMock.NewProvider(sel.ID(), sel.Name())

				ctrlOpts := control.DefaultOptions()
				ctrlOpts.ToggleFeatures.DisableDelta = !canMakeDeltaQueries

				bpc := inject.BackupProducerConfig{
					// exchange doesn't have any changes based on backup version yet.
					LastBackupVersion: version.NoBackup,
					Options:           ctrlOpts,
					ProtectedResource: uidn,
					Selector:          sel,
				}

				collections, excludes, canUsePreviousBackup, err := exchange.NewBackup().ProduceBackupCollections(
					ctx,
					bpc,
					suite.ac,
					suite.ac.Credentials,
					ctrl.UpdateStatus,
					count.New(),
					fault.New(true))
				require.NoError(t, err, clues.ToCore(err))
				assert.True(t, canUsePreviousBackup, "can use previous backup")
				assert.True(t, excludes.Empty())

				for range collections {
					ctrl.incrementAwaitingMessages()
				}

				// Categories with delta endpoints will produce a collection for metadata
				// as well as the actual data pulled, and the "temp" root collection.
				assert.LessOrEqual(t, 1, len(collections), "expected 1 <= num collections <= 3")
				assert.GreaterOrEqual(t, 3, len(collections), "expected 1 <= num collections <= 3")

				for _, col := range collections {
					for object := range col.Items(ctx, fault.New(true)) {
						buf := &bytes.Buffer{}
						_, err := buf.ReadFrom(object.ToReader())
						assert.NoError(t, err, "received a buf.Read error", clues.ToCore(err))
					}
				}

				status := ctrl.Wait()
				assert.NotZero(t, status.Successes)
				t.Log(status.String())
			})
		}
	}
}

// TestInvalidUserForDataCollections ensures verification process for users
func (suite *DataCollectionIntgSuite) TestDataCollections_invalidResourceOwner() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	owners := []string{"snuffleupagus"}
	ctrl := newController(ctx, suite.T(), path.ExchangeService)
	tests := []struct {
		name        string
		getSelector func(t *testing.T) selectors.Selector
	}{
		{
			name: "invalid exchange backup user",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup(owners)
				sel.Include(sel.MailFolders(selectors.Any()))
				return sel.Selector
			},
		},
		{
			name: "Invalid onedrive backup user",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup(owners)
				sel.Include(selTD.OneDriveBackupFolderScope(sel))
				return sel.Selector
			},
		},
		{
			name: "Invalid sharepoint backup site",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup(owners)
				sel.Include(selTD.SharePointBackupFolderScope(sel))
				return sel.Selector
			},
		},
		{
			name: "missing exchange backup user",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup(owners)
				sel.Include(sel.MailFolders(selectors.Any()))
				sel.DiscreteOwner = ""
				return sel.Selector
			},
		},
		{
			name: "missing onedrive backup user",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewOneDriveBackup(owners)
				sel.Include(selTD.OneDriveBackupFolderScope(sel))
				sel.DiscreteOwner = ""
				return sel.Selector
			},
		},
		{
			name: "missing sharepoint backup site",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewSharePointBackup(owners)
				sel.Include(selTD.SharePointBackupFolderScope(sel))
				sel.DiscreteOwner = ""
				return sel.Selector
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: version.NoBackup,
				Options:           control.DefaultOptions(),
				ProtectedResource: test.getSelector(t),
			}

			collections, excludes, canUsePreviousBackup, err := ctrl.ProduceBackupCollections(
				ctx,
				bpc,
				count.New(),
				fault.New(true))
			assert.Error(t, err, clues.ToCore(err))
			assert.False(t, canUsePreviousBackup, "can use previous backup")
			assert.Empty(t, collections)
			assert.Nil(t, excludes)
		})
	}
}

func (suite *DataCollectionIntgSuite) TestSharePointDataCollection() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	selSites := []string{suite.site}
	ctrl := newController(ctx, suite.T(), path.SharePointService)
	tests := []struct {
		name        string
		getSelector func() selectors.Selector
	}{
		{
			name: "Libraries",
			getSelector: func() selectors.Selector {
				sel := selectors.NewSharePointBackup(selSites)
				sel.Include(selTD.SharePointBackupFolderScope(sel))
				return sel.Selector
			},
		},
		{
			name: "Lists",
			getSelector: func() selectors.Selector {
				sel := selectors.NewSharePointBackup(selSites)
				sel.Include(sel.Lists(selectors.Any()))
				return sel.Selector
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sel := test.getSelector()

			bpc := inject.BackupProducerConfig{
				Options:           control.DefaultOptions(),
				ProtectedResource: sel,
				Selector:          sel,
			}

			collections, excludes, canUsePreviousBackup, err := sharepoint.NewBackup().ProduceBackupCollections(
				ctx,
				bpc,
				suite.ac,
				ctrl.credentials,
				ctrl.UpdateStatus,
				count.New(),
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			assert.True(t, canUsePreviousBackup, "can use previous backup")
			// Not expecting excludes as this isn't an incremental backup.
			assert.True(t, excludes.Empty())

			for range collections {
				ctrl.incrementAwaitingMessages()
			}

			// we don't know an exact count of drives this will produce,
			// but it should be more than zero.
			assert.NotEmpty(t, collections)

			for _, coll := range collections {
				for object := range coll.Items(ctx, fault.New(true)) {
					buf := &bytes.Buffer{}
					_, err := buf.ReadFrom(object.ToReader())
					assert.NoError(t, err, "reading item", clues.ToCore(err))
				}
			}

			status := ctrl.Wait()
			assert.NotZero(t, status.Successes)
			t.Log(status.String())
		})
	}
}

// ---------------------------------------------------------------------------
// CreateSharePointCollection tests
// ---------------------------------------------------------------------------

type SPCollectionIntgSuite struct {
	tester.Suite
	connector *Controller
	user      string
}

func TestSPCollectionIntgSuite(t *testing.T) {
	suite.Run(t, &SPCollectionIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *SPCollectionIntgSuite) SetupSuite() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	suite.connector = newController(ctx, suite.T(), path.SharePointService)
	suite.user = tconfig.M365UserID(suite.T())

	tester.LogTimeOfTest(suite.T())
}

func (suite *SPCollectionIntgSuite) TestCreateSharePointCollection_Libraries() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		siteID  = tconfig.M365SiteID(t)
		ctrl    = newController(ctx, t, path.SharePointService)
		siteIDs = []string{siteID}
	)

	site, err := ctrl.PopulateProtectedResourceIDAndName(ctx, siteID, nil)
	require.NoError(t, err, clues.ToCore(err))

	sel := selectors.NewSharePointBackup(siteIDs)
	sel.Include(sel.LibraryFolders([]string{"foo"}, selectors.PrefixMatch()))
	sel.Include(sel.Library("Documents"))

	sel.SetDiscreteOwnerIDName(site.ID(), site.Name())

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: site,
		Selector:          sel.Selector,
	}

	cols, excludes, canUsePreviousBackup, err := ctrl.ProduceBackupCollections(
		ctx,
		bpc,
		count.New(),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, canUsePreviousBackup, "can use previous backup")

	var (
		hasDocumentsColl bool
		hasMetadataColl  bool
	)

	documentsColl, err := path.BuildPrefix(
		suite.connector.tenant,
		siteID,
		path.SharePointService,
		path.LibrariesCategory)
	require.NoError(t, err, clues.ToCore(err))

	metadataColl, err := path.BuildMetadata(
		suite.connector.tenant,
		siteID,
		path.SharePointService,
		path.LibrariesCategory,
		false)
	require.NoError(t, err, clues.ToCore(err))

	for i, col := range cols {
		fp := col.FullPath()
		t.Logf("Collection %d: %s", i, fp)

		hasDocumentsColl = hasDocumentsColl || fp.Equal(documentsColl)
		hasMetadataColl = hasMetadataColl || fp.Equal(metadataColl)
	}

	require.Truef(t, hasDocumentsColl, "found documents collection %s", documentsColl)
	require.Truef(t, hasMetadataColl, "found metadata collection %s", metadataColl)

	// No excludes yet as this isn't an incremental backup.
	assert.True(t, excludes.Empty())
}

func (suite *SPCollectionIntgSuite) TestCreateSharePointCollection_Lists() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		siteID  = tconfig.M365SiteID(t)
		ctrl    = newController(ctx, t, path.SharePointService)
		siteIDs = []string{siteID}
	)

	site, err := ctrl.PopulateProtectedResourceIDAndName(ctx, siteID, nil)
	require.NoError(t, err, clues.ToCore(err))

	sel := selectors.NewSharePointBackup(siteIDs)
	sel.Include(sel.Lists(selectors.Any()))

	sel.SetDiscreteOwnerIDName(site.ID(), site.Name())

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: site,
		Selector:          sel.Selector,
	}

	cols, excludes, canUsePreviousBackup, err := ctrl.ProduceBackupCollections(
		ctx,
		bpc,
		count.New(),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, canUsePreviousBackup, "can use previous backup")
	assert.Less(t, 0, len(cols))
	// No excludes yet as this isn't an incremental backup.
	assert.True(t, excludes.Empty())

	for _, collection := range cols {
		assert.True(t, path.SharePointService == collection.FullPath().Service() ||
			path.SharePointMetadataService == collection.FullPath().Service())
		assert.Equal(t, path.ListsCategory, collection.FullPath().Category())

		for item := range collection.Items(ctx, fault.New(true)) {
			t.Log("File: " + item.ID())

			_, err := io.ReadAll(item.ToReader())
			require.NoError(t, err, clues.ToCore(err))
		}
	}
}

// ---------------------------------------------------------------------------
// CreateGroupsCollection tests
// ---------------------------------------------------------------------------

type GroupsCollectionIntgSuite struct {
	tester.Suite
	connector *Controller
	tenantID  string
	user      string
}

func TestGroupsCollectionIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsCollectionIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *GroupsCollectionIntgSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.connector = newController(ctx, t, path.GroupsService)
	suite.user = tconfig.M365UserID(t)

	acct := tconfig.NewM365Account(t)
	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.tenantID = creds.AzureTenantID

	tester.LogTimeOfTest(t)
}

func (suite *GroupsCollectionIntgSuite) TestCreateGroupsCollection_SharePoint() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		groupID  = tconfig.M365TeamID(t)
		ctrl     = newController(ctx, t, path.GroupsService)
		groupIDs = []string{groupID}
	)

	group, err := ctrl.PopulateProtectedResourceIDAndName(ctx, groupID, nil)
	require.NoError(t, err, clues.ToCore(err))

	sel := selectors.NewGroupsBackup(groupIDs)
	sel.Include(sel.LibraryFolders([]string{"test"}, selectors.PrefixMatch()))

	sel.SetDiscreteOwnerIDName(group.ID(), group.Name())

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: group,
		Selector:          sel.Selector,
	}

	collections, excludes, canUsePreviousBackup, err := ctrl.ProduceBackupCollections(
		ctx,
		bpc,
		count.New(),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, canUsePreviousBackup, "can use previous backup")
	// No excludes yet as this isn't an incremental backup.
	assert.True(t, excludes.Empty())

	// we don't know an exact count of drives this will produce,
	// but it should be more than one.
	assert.Greater(t, len(collections), 1)

	p, err := path.BuildMetadata(
		suite.tenantID,
		groupID,
		path.GroupsService,
		path.LibrariesCategory,
		false)
	require.NoError(t, err, clues.ToCore(err))

	p, err = p.Append(false, odConsts.SitesPathDir)
	require.NoError(t, err, clues.ToCore(err))

	foundSitesMetadata := false

	for _, coll := range collections {
		sitesMetadataCollection := coll.FullPath().String() == p.String()

		for object := range coll.Items(ctx, fault.New(true)) {
			if object.ID() == "previouspath" && sitesMetadataCollection {
				foundSitesMetadata = true
			}

			buf := &bytes.Buffer{}
			_, err := buf.ReadFrom(object.ToReader())
			assert.NoError(t, err, "reading item", clues.ToCore(err))
		}
	}

	assert.True(t, foundSitesMetadata, "missing sites metadata")

	status := ctrl.Wait()
	assert.NotZero(t, status.Successes)
	t.Log(status.String())
}

func (suite *GroupsCollectionIntgSuite) TestCreateGroupsCollection_SharePoint_InvalidMetadata() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		groupID  = tconfig.M365TeamID(t)
		ctrl     = newController(ctx, t, path.GroupsService)
		groupIDs = []string{groupID}
	)

	group, err := ctrl.PopulateProtectedResourceIDAndName(ctx, groupID, nil)
	require.NoError(t, err, clues.ToCore(err))

	sel := selectors.NewGroupsBackup(groupIDs)
	sel.Include(sel.LibraryFolders([]string{"test"}, selectors.PrefixMatch()))

	sel.SetDiscreteOwnerIDName(group.ID(), group.Name())

	site, err := suite.connector.AC.Groups().GetRootSite(ctx, groupID)
	require.NoError(t, err, clues.ToCore(err))

	pth, err := path.Build(
		suite.tenantID,
		groupID,
		path.GroupsService,
		path.LibrariesCategory,
		true,
		odConsts.SitesPathDir,
		ptr.Val(site.GetId()))
	require.NoError(t, err, clues.ToCore(err))

	mmc := []data.RestoreCollection{
		mock.Collection{
			Path: pth,
			ItemData: []data.Item{
				&mock.Item{
					ItemID: "previouspath",
					Reader: io.NopCloser(bytes.NewReader([]byte("invalid"))),
				},
			},
		},
	}

	bpc := inject.BackupProducerConfig{
		LastBackupVersion:   version.NoBackup,
		Options:             control.DefaultOptions(),
		ProtectedResource:   group,
		Selector:            sel.Selector,
		MetadataCollections: mmc,
	}

	collections, excludes, canUsePreviousBackup, err := ctrl.ProduceBackupCollections(
		ctx,
		bpc,
		count.New(),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, canUsePreviousBackup, "can use previous backup")
	// No excludes yet as this isn't an incremental backup.
	assert.True(t, excludes.Empty())

	// we don't know an exact count of drives this will produce,
	// but it should be more than one.
	assert.Greater(t, len(collections), 1)

	p, err := path.BuildMetadata(
		suite.tenantID,
		groupID,
		path.GroupsService,
		path.LibrariesCategory,
		false)
	require.NoError(t, err, clues.ToCore(err))

	p, err = p.Append(false, odConsts.SitesPathDir)
	require.NoError(t, err, clues.ToCore(err))

	foundSitesMetadata := false
	foundRootTombstone := false

	sp, err := path.BuildPrefix(
		suite.tenantID,
		groupID,
		path.GroupsService,
		path.LibrariesCategory)
	require.NoError(t, err, clues.ToCore(err))

	sp, err = sp.Append(false, odConsts.SitesPathDir, ptr.Val(site.GetId()))
	require.NoError(t, err, clues.ToCore(err))

	for _, coll := range collections {
		if coll.State() == data.DeletedState {
			if coll.PreviousPath() != nil && coll.PreviousPath().String() == sp.String() {
				foundRootTombstone = true
			}

			continue
		}

		sitesMetadataCollection := coll.FullPath().String() == p.String()

		for object := range coll.Items(ctx, fault.New(true)) {
			if object.ID() == "previouspath" && sitesMetadataCollection {
				foundSitesMetadata = true
			}

			buf := &bytes.Buffer{}
			_, err := buf.ReadFrom(object.ToReader())
			assert.NoError(t, err, "reading item", clues.ToCore(err))
		}
	}

	assert.True(t, foundSitesMetadata, "missing sites metadata")
	assert.True(t, foundRootTombstone, "missing root tombstone")

	status := ctrl.Wait()
	assert.NotZero(t, status.Successes)
	t.Log(status.String())
}
