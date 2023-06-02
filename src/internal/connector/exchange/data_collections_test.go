package exchange

import (
	"bytes"
	"sync"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// Unit tests
// ---------------------------------------------------------------------------

type DataCollectionsUnitSuite struct {
	tester.Suite
}

func TestDataCollectionsUnitSuite(t *testing.T) {
	suite.Run(t, &DataCollectionsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *DataCollectionsUnitSuite) TestParseMetadataCollections() {
	type fileValues struct {
		fileName string
		value    string
	}

	table := []struct {
		name        string
		data        []fileValues
		expect      map[string]DeltaPath
		expectError assert.ErrorAssertionFunc
	}{
		{
			name: "delta urls only",
			data: []fileValues{
				{graph.DeltaURLsFileName, "delta-link"},
			},
			expect:      map[string]DeltaPath{},
			expectError: assert.NoError,
		},
		{
			name: "multiple delta urls",
			data: []fileValues{
				{graph.DeltaURLsFileName, "delta-link"},
				{graph.DeltaURLsFileName, "delta-link-2"},
			},
			expectError: assert.Error,
		},
		{
			name: "previous path only",
			data: []fileValues{
				{graph.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]DeltaPath{
				"key": {
					Delta: "delta-link",
					Path:  "prev-path",
				},
			},
			expectError: assert.NoError,
		},
		{
			name: "multiple previous paths",
			data: []fileValues{
				{graph.PreviousPathFileName, "prev-path"},
				{graph.PreviousPathFileName, "prev-path-2"},
			},
			expectError: assert.Error,
		},
		{
			name: "delta urls and previous paths",
			data: []fileValues{
				{graph.DeltaURLsFileName, "delta-link"},
				{graph.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]DeltaPath{
				"key": {
					Delta: "delta-link",
					Path:  "prev-path",
				},
			},
			expectError: assert.NoError,
		},
		{
			name: "delta urls and empty previous paths",
			data: []fileValues{
				{graph.DeltaURLsFileName, "delta-link"},
				{graph.PreviousPathFileName, ""},
			},
			expect:      map[string]DeltaPath{},
			expectError: assert.NoError,
		},
		{
			name: "empty delta urls and previous paths",
			data: []fileValues{
				{graph.DeltaURLsFileName, ""},
				{graph.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]DeltaPath{
				"key": {
					Delta: "delta-link",
					Path:  "prev-path",
				},
			},
			expectError: assert.NoError,
		},
		{
			name: "delta urls with special chars",
			data: []fileValues{
				{graph.DeltaURLsFileName, "`!@#$%^&*()_[]{}/\"\\"},
				{graph.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]DeltaPath{
				"key": {
					Delta: "`!@#$%^&*()_[]{}/\"\\",
					Path:  "prev-path",
				},
			},
			expectError: assert.NoError,
		},
		{
			name: "delta urls with escaped chars",
			data: []fileValues{
				{graph.DeltaURLsFileName, `\n\r\t\b\f\v\0\\`},
				{graph.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]DeltaPath{
				"key": {
					Delta: "\\n\\r\\t\\b\\f\\v\\0\\\\",
					Path:  "prev-path",
				},
			},
			expectError: assert.NoError,
		},
		{
			name: "delta urls with newline char runes",
			data: []fileValues{
				// rune(92) = \, rune(110) = n.  Ensuring it's not possible to
				// error in serializing/deserializing and produce a single newline
				// character from those two runes.
				{graph.DeltaURLsFileName, string([]rune{rune(92), rune(110)})},
				{graph.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]DeltaPath{
				"key": {
					Delta: "\\n",
					Path:  "prev-path",
				},
			},
			expectError: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			entries := []graph.MetadataCollectionEntry{}

			for _, d := range test.data {
				entries = append(
					entries,
					graph.NewMetadataEntry(d.fileName, map[string]string{"key": d.value}))
			}

			coll, err := graph.MakeMetadataCollection(
				"t", "u",
				path.ExchangeService,
				path.EmailCategory,
				entries,
				func(cos *support.ConnectorOperationStatus) {},
			)
			require.NoError(t, err, clues.ToCore(err))

			cdps, err := parseMetadataCollections(ctx, []data.RestoreCollection{
				data.NoFetchRestoreCollection{Collection: coll},
			}, fault.New(true))
			test.expectError(t, err, clues.ToCore(err))

			emails := cdps[path.EmailCategory]

			assert.Len(t, emails, len(test.expect))

			for k, v := range emails {
				assert.Equal(t, v.Delta, emails[k].Delta, "delta")
				assert.Equal(t, v.Path, emails[k].Path, "path")
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Integration tests
// ---------------------------------------------------------------------------

func newStatusUpdater(t *testing.T, wg *sync.WaitGroup) func(status *support.ConnectorOperationStatus) {
	updater := func(status *support.ConnectorOperationStatus) {
		defer wg.Done()
	}

	return updater
}

type DataCollectionsIntegrationSuite struct {
	tester.Suite
	user     string
	site     string
	tenantID string
	ac       api.Client
}

func TestDataCollectionsIntegrationSuite(t *testing.T) {
	suite.Run(t, &DataCollectionsIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tester.M365AcctCredEnvs},
		),
	})
}

func (suite *DataCollectionsIntegrationSuite) SetupSuite() {
	suite.user = tester.M365UserID(suite.T())
	suite.site = tester.M365SiteID(suite.T())

	acct := tester.NewM365Account(suite.T())
	creds, err := acct.M365Config()
	require.NoError(suite.T(), err, clues.ToCore(err))

	suite.ac, err = api.NewClient(creds)
	require.NoError(suite.T(), err, clues.ToCore(err))

	suite.tenantID = creds.AzureTenantID

	tester.LogTimeOfTest(suite.T())
}

func (suite *DataCollectionsIntegrationSuite) TestMailFetch() {
	var (
		userID   = tester.M365UserID(suite.T())
		users    = []string{userID}
		handlers = BackupHandlers(suite.ac)
	)

	tests := []struct {
		name                string
		scope               selectors.ExchangeScope
		folderNames         map[string]struct{}
		canMakeDeltaQueries bool
	}{
		{
			name: "Folder Iterative Check Mail",
			scope: selectors.NewExchangeBackup(users).MailFolders(
				[]string{DefaultMailFolder},
				selectors.PrefixMatch(),
			)[0],
			folderNames: map[string]struct{}{
				DefaultMailFolder: {},
			},
			canMakeDeltaQueries: true,
		},
		{
			name: "Folder Iterative Check Mail Non-Delta",
			scope: selectors.NewExchangeBackup(users).MailFolders(
				[]string{DefaultMailFolder},
				selectors.PrefixMatch(),
			)[0],
			folderNames: map[string]struct{}{
				DefaultMailFolder: {},
			},
			canMakeDeltaQueries: false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrlOpts := control.Defaults()
			ctrlOpts.ToggleFeatures.DisableDelta = !test.canMakeDeltaQueries

			collections, err := createCollections(
				ctx,
				handlers,
				suite.tenantID,
				inMock.NewProvider(userID, userID),
				test.scope,
				DeltaPaths{},
				ctrlOpts,
				func(status *support.ConnectorOperationStatus) {},
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			for _, c := range collections {
				if c.FullPath().Service() == path.ExchangeMetadataService {
					continue
				}

				require.NotEmpty(t, c.FullPath().Folder(false))

				// TODO(ashmrtn): Remove when LocationPath is made part of BackupCollection
				// interface.
				if !assert.Implements(t, (*data.LocationPather)(nil), c) {
					continue
				}

				loc := c.(data.LocationPather).LocationPath().String()

				require.NotEmpty(t, loc)

				delete(test.folderNames, loc)
			}

			assert.Empty(t, test.folderNames)
		})
	}
}

func (suite *DataCollectionsIntegrationSuite) TestDelta() {
	var (
		userID   = tester.M365UserID(suite.T())
		users    = []string{userID}
		handlers = BackupHandlers(suite.ac)
	)

	tests := []struct {
		name  string
		scope selectors.ExchangeScope
	}{
		{
			name: "Mail",
			scope: selectors.NewExchangeBackup(users).MailFolders(
				[]string{DefaultMailFolder},
				selectors.PrefixMatch(),
			)[0],
		},
		{
			name: "Contacts",
			scope: selectors.NewExchangeBackup(users).ContactFolders(
				[]string{DefaultContactFolder},
				selectors.PrefixMatch(),
			)[0],
		},
		{
			name: "Events",
			scope: selectors.NewExchangeBackup(users).EventCalendars(
				[]string{DefaultCalendar},
				selectors.PrefixMatch(),
			)[0],
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			// get collections without providing any delta history (ie: full backup)
			collections, err := createCollections(
				ctx,
				handlers,
				suite.tenantID,
				inMock.NewProvider(userID, userID),
				test.scope,
				DeltaPaths{},
				control.Defaults(),
				func(status *support.ConnectorOperationStatus) {},
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			assert.Less(t, 1, len(collections), "retrieved metadata and data collections")

			var metadata data.BackupCollection

			for _, coll := range collections {
				if coll.FullPath().Service() == path.ExchangeMetadataService {
					metadata = coll
				}
			}

			require.NotNil(t, metadata, "collections contains a metadata collection")

			cdps, err := parseMetadataCollections(ctx, []data.RestoreCollection{
				data.NoFetchRestoreCollection{Collection: metadata},
			}, fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			dps := cdps[test.scope.Category().PathType()]

			// now do another backup with the previous delta tokens,
			// which should only contain the difference.
			collections, err = createCollections(
				ctx,
				handlers,
				suite.tenantID,
				inMock.NewProvider(userID, userID),
				test.scope,
				dps,
				control.Defaults(),
				func(status *support.ConnectorOperationStatus) {},
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			// TODO(keepers): this isn't a very useful test at the moment.  It needs to
			// investigate the items in the original and delta collections to at least
			// assert some minimum assumptions, such as "deltas should retrieve fewer items".
			// Delta usage is commented out at the moment, anyway.  So this is currently
			// a sanity check that the minimum behavior won't break.
			for _, coll := range collections {
				if coll.FullPath().Service() != path.ExchangeMetadataService {
					ec, ok := coll.(*Collection)
					require.True(t, ok, "collection is *Collection")
					assert.NotNil(t, ec)
				}
			}
		})
	}
}

// TestMailSerializationRegression verifies that all mail data stored in the
// test account can be successfully downloaded into bytes and restored into
// M365 mail objects
func (suite *DataCollectionsIntegrationSuite) TestMailSerializationRegression() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		wg       sync.WaitGroup
		users    = []string{suite.user}
		handlers = BackupHandlers(suite.ac)
	)

	sel := selectors.NewExchangeBackup(users)
	sel.Include(sel.MailFolders([]string{DefaultMailFolder}, selectors.PrefixMatch()))

	collections, err := createCollections(
		ctx,
		handlers,
		suite.tenantID,
		inMock.NewProvider(suite.user, suite.user),
		sel.Scopes()[0],
		DeltaPaths{},
		control.Defaults(),
		newStatusUpdater(t, &wg),
		fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	wg.Add(len(collections))

	for _, edc := range collections {
		suite.Run(edc.FullPath().String(), func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			isMetadata := edc.FullPath().Service() == path.ExchangeMetadataService
			streamChannel := edc.Items(ctx, fault.New(true))

			// Verify that each message can be restored
			for stream := range streamChannel {
				buf := &bytes.Buffer{}

				read, err := buf.ReadFrom(stream.ToReader())
				assert.NoError(t, err, clues.ToCore(err))
				assert.NotZero(t, read)

				if isMetadata {
					continue
				}

				message, err := api.BytesToMessageable(buf.Bytes())
				assert.NotNil(t, message)
				assert.NoError(t, err, clues.ToCore(err))
			}
		})
	}

	wg.Wait()
}

// TestContactSerializationRegression verifies ability to query contact items
// and to store contact within Collection. Downloaded contacts are run through
// a regression test to ensure that downloaded items can be uploaded.
func (suite *DataCollectionsIntegrationSuite) TestContactSerializationRegression() {
	var (
		users    = []string{suite.user}
		handlers = BackupHandlers(suite.ac)
	)

	tests := []struct {
		name  string
		scope selectors.ExchangeScope
	}{
		{
			name: "Default Contact Folder",
			scope: selectors.NewExchangeBackup(users).ContactFolders(
				[]string{DefaultContactFolder},
				selectors.PrefixMatch())[0],
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var wg sync.WaitGroup

			edcs, err := createCollections(
				ctx,
				handlers,
				suite.tenantID,
				inMock.NewProvider(suite.user, suite.user),
				test.scope,
				DeltaPaths{},
				control.Defaults(),
				newStatusUpdater(t, &wg),
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			wg.Add(len(edcs))

			require.GreaterOrEqual(t, len(edcs), 1, "expected 1 <= num collections <= 2")
			require.GreaterOrEqual(t, 2, len(edcs), "expected 1 <= num collections <= 2")

			for _, edc := range edcs {
				isMetadata := edc.FullPath().Service() == path.ExchangeMetadataService
				count := 0

				for stream := range edc.Items(ctx, fault.New(true)) {
					buf := &bytes.Buffer{}
					read, err := buf.ReadFrom(stream.ToReader())
					assert.NoError(t, err, clues.ToCore(err))
					assert.NotZero(t, read)

					if isMetadata {
						continue
					}

					contact, err := api.BytesToContactable(buf.Bytes())
					assert.NotNil(t, contact)
					assert.NoError(t, err, "converting contact bytes: "+buf.String(), clues.ToCore(err))
					count++
				}

				if isMetadata {
					continue
				}

				// TODO(ashmrtn): Remove when LocationPath is made part of BackupCollection
				// interface.
				if !assert.Implements(t, (*data.LocationPather)(nil), edc) {
					continue
				}

				assert.Equal(
					t,
					edc.(data.LocationPather).LocationPath().String(),
					DefaultContactFolder)
				assert.NotZero(t, count)
			}

			wg.Wait()
		})
	}
}

// TestEventsSerializationRegression ensures functionality of createCollections
// to be able to successfully query, download and restore event objects
func (suite *DataCollectionsIntegrationSuite) TestEventsSerializationRegression() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		users    = []string{suite.user}
		handlers = BackupHandlers(suite.ac)
		calID    string
		bdayID   string
	)

	fn := func(gcf graph.CachedContainer) error {
		if ptr.Val(gcf.GetDisplayName()) == DefaultCalendar {
			calID = ptr.Val(gcf.GetId())
		}

		if ptr.Val(gcf.GetDisplayName()) == "Birthdays" {
			bdayID = ptr.Val(gcf.GetId())
		}

		return nil
	}

	err := suite.ac.Events().EnumerateContainers(ctx, suite.user, DefaultCalendar, fn, fault.New(true))
	require.NoError(t, err, clues.ToCore(err))

	tests := []struct {
		name, expected string
		scope          selectors.ExchangeScope
	}{
		{
			name:     "Default Event Calendar",
			expected: calID,
			scope: selectors.NewExchangeBackup(users).EventCalendars(
				[]string{DefaultCalendar},
				selectors.PrefixMatch(),
			)[0],
		},
		{
			name:     "Birthday Calendar",
			expected: bdayID,
			scope: selectors.NewExchangeBackup(users).EventCalendars(
				[]string{"Birthdays"},
				selectors.PrefixMatch(),
			)[0],
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var wg sync.WaitGroup

			collections, err := createCollections(
				ctx,
				handlers,
				suite.tenantID,
				inMock.NewProvider(suite.user, suite.user),
				test.scope,
				DeltaPaths{},
				control.Defaults(),
				newStatusUpdater(t, &wg),
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			require.Len(t, collections, 2)

			wg.Add(len(collections))

			for _, edc := range collections {
				var isMetadata bool

				if edc.FullPath().Service() != path.ExchangeMetadataService {
					isMetadata = true
					assert.Equal(t, test.expected, edc.FullPath().Folder(false))
				} else {
					assert.Equal(t, "", edc.FullPath().Folder(false))
				}

				for item := range edc.Items(ctx, fault.New(true)) {
					buf := &bytes.Buffer{}

					read, err := buf.ReadFrom(item.ToReader())
					assert.NoError(t, err, clues.ToCore(err))
					assert.NotZero(t, read)

					if isMetadata {
						continue
					}

					event, err := api.BytesToEventable(buf.Bytes())
					assert.NotNil(t, event)
					assert.NoError(t, err, "creating event from bytes: "+buf.String(), clues.ToCore(err))
				}
			}

			wg.Wait()
		})
	}
}
