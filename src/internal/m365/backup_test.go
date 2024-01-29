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
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/service/exchange"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/m365/service/sharepoint"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/its"
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
	m365 its.M365IntgTestSetup
}

func TestDataCollectionIntgSuite(t *testing.T) {
	suite.Run(t, &DataCollectionIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *DataCollectionIntgSuite) SetupSuite() {
	suite.m365 = its.GetM365(suite.T())
}

func (suite *DataCollectionIntgSuite) TestExchangeDataCollection() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	selUsers := []string{suite.m365.User.ID}

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
				sel.DiscreteOwner = suite.m365.User.ID
				return sel.Selector
			},
		},
		{
			name: "Contacts",
			getSelector: func(t *testing.T) selectors.Selector {
				sel := selectors.NewExchangeBackup(selUsers)
				sel.Include(sel.ContactFolders([]string{api.DefaultContacts}, selectors.PrefixMatch()))
				sel.DiscreteOwner = suite.m365.User.ID
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
					suite.m365.AC,
					suite.m365.Creds,
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

	selSites := []string{suite.m365.Site.ID}
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
				suite.m365.AC,
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
	m365      its.M365IntgTestSetup
}

func TestSPCollectionIntgSuite(t *testing.T) {
	suite.Run(t, &SPCollectionIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *SPCollectionIntgSuite) SetupSuite() {
	t := suite.T()
	suite.m365 = its.GetM365(t)

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.connector = newController(ctx, t, path.SharePointService)
}

func (suite *SPCollectionIntgSuite) TestCreateSharePointCollection_Libraries() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		ctrl    = newController(ctx, t, path.SharePointService)
		siteIDs = []string{suite.m365.Site.ID}
	)

	sel := selectors.NewSharePointBackup(siteIDs)
	sel.Include(sel.LibraryFolders([]string{"foo"}, selectors.PrefixMatch()))
	sel.SetDiscreteOwnerIDName(suite.m365.Site.ID, suite.m365.Site.WebURL)

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: suite.m365.Site.Provider,
		Selector:          sel.Selector,
	}

	cols, excludes, canUsePreviousBackup, err := ctrl.ProduceBackupCollections(
		ctx,
		bpc,
		count.New(),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, canUsePreviousBackup, "can use previous backup")
	require.Len(t, cols, 2) // 1 collection, 1 path prefix directory to ensure the root path exists.
	// No excludes yet as this isn't an incremental backup.
	assert.True(t, excludes.Empty())

	t.Logf("cols[0] Path: %s\n", cols[0].FullPath().String())
	assert.Equal(
		t,
		path.SharePointMetadataService.String(),
		cols[0].FullPath().Service().String())

	t.Logf("cols[1] Path: %s\n", cols[1].FullPath().String())
	assert.Equal(
		t,
		path.SharePointService.String(),
		cols[1].FullPath().Service().String())
}

func (suite *SPCollectionIntgSuite) TestCreateSharePointCollection_Lists() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		ctrl    = newController(ctx, t, path.SharePointService)
		siteIDs = []string{suite.m365.Site.ID}
	)

	sel := selectors.NewSharePointBackup(siteIDs)
	sel.Include(sel.Lists(selectors.Any()))
	sel.SetDiscreteOwnerIDName(suite.m365.Site.ID, suite.m365.Site.WebURL)

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: suite.m365.Site.Provider,
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
	m365      its.M365IntgTestSetup
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
	suite.m365 = its.GetM365(t)

	ctx, flush := tester.NewContext(t)
	defer flush()

	suite.connector = newController(ctx, t, path.GroupsService)
}

func (suite *GroupsCollectionIntgSuite) TestCreateGroupsCollection_SharePoint() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		ctrl     = newController(ctx, t, path.GroupsService)
		groupIDs = []string{suite.m365.Group.ID}
	)

	sel := selectors.NewGroupsBackup(groupIDs)
	sel.Include(sel.LibraryFolders([]string{"test"}, selectors.PrefixMatch()))
	sel.SetDiscreteOwnerIDName(suite.m365.Group.ID, suite.m365.Group.DisplayName)

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: suite.m365.Group.Provider,
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
		suite.m365.TenantID,
		suite.m365.Group.ID,
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
		ctrl     = newController(ctx, t, path.GroupsService)
		groupIDs = []string{suite.m365.Group.ID}
	)

	sel := selectors.NewGroupsBackup(groupIDs)
	sel.Include(sel.LibraryFolders([]string{"test"}, selectors.PrefixMatch()))
	sel.SetDiscreteOwnerIDName(suite.m365.Group.ID, suite.m365.Group.DisplayName)

	pth, err := path.Build(
		suite.m365.TenantID,
		suite.m365.Group.ID,
		path.GroupsService,
		path.LibrariesCategory,
		true,
		odConsts.SitesPathDir,
		suite.m365.Group.RootSite.ID)
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
		ProtectedResource:   suite.m365.Group.Provider,
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
		suite.m365.TenantID,
		suite.m365.Group.ID,
		path.GroupsService,
		path.LibrariesCategory,
		false)
	require.NoError(t, err, clues.ToCore(err))

	p, err = p.Append(false, odConsts.SitesPathDir)
	require.NoError(t, err, clues.ToCore(err))

	foundSitesMetadata := false
	foundRootTombstone := false

	sp, err := path.BuildPrefix(
		suite.m365.TenantID,
		suite.m365.Group.ID,
		path.GroupsService,
		path.LibrariesCategory)
	require.NoError(t, err, clues.ToCore(err))

	sp, err = sp.Append(false, odConsts.SitesPathDir, suite.m365.Site.ID)
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
