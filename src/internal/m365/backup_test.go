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
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/internal/m365/service/exchange"
	"github.com/alcionai/corso/src/internal/m365/service/sharepoint"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/control"
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

	suite.ac, err = api.NewClient(creds, control.DefaultOptions())
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *DataCollectionIntgSuite) TestExchangeDataCollection() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	selUsers := []string{suite.user}

	ctrl := newController(ctx, suite.T(), resource.Users, path.ExchangeService)
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

				collections, excludes, canUsePreviousBackup, err := exchange.ProduceBackupCollections(
					ctx,
					bpc,
					suite.ac,
					suite.tenantID,
					ctrl.UpdateStatus,
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
	ctrl := newController(ctx, suite.T(), resource.Users, path.ExchangeService)
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
	ctrl := newController(ctx, suite.T(), resource.Sites, path.SharePointService)
	tests := []struct {
		name        string
		expected    int
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
			name:     "Lists",
			expected: 0,
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

			collections, excludes, canUsePreviousBackup, err := sharepoint.ProduceBackupCollections(
				ctx,
				bpc,
				suite.ac,
				ctrl.credentials,
				ctrl.UpdateStatus,
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			assert.True(t, canUsePreviousBackup, "can use previous backup")
			// Not expecting excludes as this isn't an incremental backup.
			assert.True(t, excludes.Empty())

			for range collections {
				ctrl.incrementAwaitingMessages()
			}

			// we don't know an exact count of drives this will produce,
			// but it should be more than one.
			assert.Less(t, test.expected, len(collections))

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
			[][]string{tconfig.M365AcctCredEnvs},
		),
	})
}

func (suite *SPCollectionIntgSuite) SetupSuite() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	suite.connector = newController(ctx, suite.T(), resource.Sites, path.SharePointService)
	suite.user = tconfig.M365UserID(suite.T())

	tester.LogTimeOfTest(suite.T())
}

func (suite *SPCollectionIntgSuite) TestCreateSharePointCollection_Libraries() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		siteID  = tconfig.M365SiteID(t)
		ctrl    = newController(ctx, t, resource.Sites, path.SharePointService)
		siteIDs = []string{siteID}
	)

	id, name, err := ctrl.PopulateProtectedResourceIDAndName(ctx, siteID, nil)
	require.NoError(t, err, clues.ToCore(err))

	sel := selectors.NewSharePointBackup(siteIDs)
	sel.Include(sel.LibraryFolders([]string{"foo"}, selectors.PrefixMatch()))

	sel.SetDiscreteOwnerIDName(id, name)

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: inMock.NewProvider(id, name),
		Selector:          sel.Selector,
	}

	cols, excludes, canUsePreviousBackup, err := ctrl.ProduceBackupCollections(
		ctx,
		bpc,
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
		siteID  = tconfig.M365SiteID(t)
		ctrl    = newController(ctx, t, resource.Sites, path.SharePointService)
		siteIDs = []string{siteID}
	)

	id, name, err := ctrl.PopulateProtectedResourceIDAndName(ctx, siteID, nil)
	require.NoError(t, err, clues.ToCore(err))

	sel := selectors.NewSharePointBackup(siteIDs)
	sel.Include(sel.Lists(selectors.Any()))

	sel.SetDiscreteOwnerIDName(id, name)

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: inMock.NewProvider(id, name),
		Selector:          sel.Selector,
	}

	cols, excludes, canUsePreviousBackup, err := ctrl.ProduceBackupCollections(
		ctx,
		bpc,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, canUsePreviousBackup, "can use previous backup")
	assert.Less(t, 0, len(cols))
	// No excludes yet as this isn't an incremental backup.
	assert.True(t, excludes.Empty())

	for _, collection := range cols {
		t.Logf("Path: %s\n", collection.FullPath().String())

		for item := range collection.Items(ctx, fault.New(true)) {
			t.Log("File: " + item.ID())

			bs, err := io.ReadAll(item.ToReader())
			require.NoError(t, err, clues.ToCore(err))
			t.Log(string(bs))
		}
	}
}

// ---------------------------------------------------------------------------
// CreateGroupsCollection tests
// ---------------------------------------------------------------------------

type GroupsCollectionIntgSuite struct {
	tester.Suite
	connector *Controller
	user      string
}

func TestGroupsCollectionIntgSuite(t *testing.T) {
	suite.Run(t, &GroupsCollectionIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs},
		),
	})
}

func (suite *GroupsCollectionIntgSuite) SetupSuite() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	suite.connector = newController(ctx, suite.T(), resource.Sites, path.GroupsService)
	suite.user = tconfig.M365UserID(suite.T())

	tester.LogTimeOfTest(suite.T())
}

func (suite *GroupsCollectionIntgSuite) TestCreateGroupsCollection_SharePoint() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		groupID  = tconfig.M365GroupID(t)
		ctrl     = newController(ctx, t, resource.Groups, path.GroupsService)
		groupIDs = []string{groupID}
	)

	id, name, err := ctrl.PopulateProtectedResourceIDAndName(ctx, groupID, nil)
	require.NoError(t, err, clues.ToCore(err))

	sel := selectors.NewGroupsBackup(groupIDs)
	// TODO(meain): make use of selectors
	sel.Include(sel.LibraryFolders([]string{"test"}, selectors.PrefixMatch()))

	sel.SetDiscreteOwnerIDName(id, name)

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: inMock.NewProvider(id, name),
		Selector:          sel.Selector,
	}

	collections, excludes, canUsePreviousBackup, err := ctrl.ProduceBackupCollections(
		ctx,
		bpc,
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))
	assert.True(t, canUsePreviousBackup, "can use previous backup")
	// No excludes yet as this isn't an incremental backup.
	assert.True(t, excludes.Empty())

	// we don't know an exact count of drives this will produce,
	// but it should be more than one.
	assert.Greater(t, len(collections), 1)

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
}
