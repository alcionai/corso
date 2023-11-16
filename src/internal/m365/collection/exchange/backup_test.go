package exchange

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/microsoft/kiota-abstractions-go/serialization"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/common/readers"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	"github.com/alcionai/corso/src/pkg/services/m365/api/graph"
	"github.com/alcionai/corso/src/pkg/services/m365/api/pagers"
)

// ---------------------------------------------------------------------------
// mocks
// ---------------------------------------------------------------------------

var (
	_ backupHandler        = &mockBackupHandler{}
	_ itemGetterSerializer = mockItemGetter{}
)

// mockItemGetter implmenets the basics required to allow calls to
// Collection.Items(). However, it returns static data.
type mockItemGetter struct{}

func (ig mockItemGetter) GetItem(
	context.Context,
	string,
	string,
	bool,
	*fault.Bus,
) (serialization.Parsable, *details.ExchangeInfo, error) {
	return models.NewMessage(), &details.ExchangeInfo{}, nil
}

func (ig mockItemGetter) Serialize(
	context.Context,
	serialization.Parsable,
	string,
	string,
) ([]byte, error) {
	return []byte("foo"), nil
}

type mockBackupHandler struct {
	mg              mockGetter
	fg              containerGetter
	category        path.CategoryType
	ac              api.Client
	userID          string
	previewIncludes []string
	previewExcludes []string
}

func (bh mockBackupHandler) itemEnumerator() addedAndRemovedItemGetter { return bh.mg }
func (bh mockBackupHandler) itemHandler() itemGetterSerializer         { return mockItemGetter{} }
func (bh mockBackupHandler) folderGetter() containerGetter             { return bh.fg }
func (bh mockBackupHandler) previewIncludeContainers() []string        { return bh.previewIncludes }
func (bh mockBackupHandler) previewExcludeContainers() []string        { return bh.previewExcludes }

func (bh mockBackupHandler) NewContainerCache(
	userID string,
) (string, graph.ContainerResolver) {
	return BackupHandlers(bh.ac)[bh.category].NewContainerCache(bh.userID)
}

var _ addedAndRemovedItemGetter = &mockGetter{}

type (
	mockGetter struct {
		noReturnDelta bool
		results       map[string]mockGetterResults
	}
	mockGetterResults struct {
		added    []string
		removed  []string
		newDelta pagers.DeltaUpdate
		err      error
	}
)

func (mg mockGetter) GetAddedAndRemovedItemIDs(
	ctx context.Context,
	userID, cID, prevDelta string,
	config api.CallConfig,
) (pagers.AddedAndRemoved, error) {
	results, ok := mg.results[cID]
	if !ok {
		return pagers.AddedAndRemoved{}, clues.New("mock not found for " + cID)
	}

	delta := results.newDelta
	if mg.noReturnDelta {
		delta.URL = ""
	}

	toAdd := config.LimitResults
	if toAdd == 0 || toAdd > len(results.added) {
		toAdd = len(results.added)
	}

	resAdded := make(map[string]time.Time, toAdd)
	for _, add := range results.added[:toAdd] {
		resAdded[add] = time.Time{}
	}

	aar := pagers.AddedAndRemoved{
		Added:         resAdded,
		Removed:       results.removed,
		ValidModTimes: false,
		DU:            delta,
	}

	return aar, results.err
}

var (
	_ graph.ContainerResolver = &mockResolver{}
	_ containerGetter         = &mockResolver{}
)

type mockResolver struct {
	items []graph.CachedContainer
	added map[string]string
}

func newMockResolver(items ...mockContainer) mockResolver {
	is := make([]graph.CachedContainer, 0, len(items))

	for _, i := range items {
		is = append(is, i)
	}

	return mockResolver{items: is}
}

func (m mockResolver) ItemByID(id string) graph.CachedContainer {
	for _, c := range m.items {
		if ptr.Val(c.GetId()) == id {
			return c
		}
	}

	return nil
}

// GetContainerByID returns the given container if it exists in the resolver.
// This is kind of merging functionality that we generally assume is separate,
// but it does allow for easier test setup.
func (m mockResolver) GetContainerByID(
	ctx context.Context,
	userID, dirID string,
) (graph.Container, error) {
	c := m.ItemByID(dirID)
	if c == nil {
		return nil, data.ErrNotFound
	}

	return c, nil
}

func (m mockResolver) Items() []graph.CachedContainer {
	return m.items
}

func (m mockResolver) AddToCache(ctx context.Context, ctrl graph.Container) error {
	if len(m.added) == 0 {
		m.added = map[string]string{}
	}

	m.added[ptr.Val(ctrl.GetDisplayName())] = ptr.Val(ctrl.GetId())

	return nil
}
func (m mockResolver) DestinationNameToID(dest string) string { return m.added[dest] }
func (m mockResolver) IDToPath(context.Context, string) (*path.Builder, *path.Builder, error) {
	return nil, nil, nil
}
func (m mockResolver) PathInCache(string) (string, bool)                             { return "", false }
func (m mockResolver) LocationInCache(string) (string, bool)                         { return "", false }
func (m mockResolver) Populate(context.Context, *fault.Bus, string, ...string) error { return nil }

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
		name                 string
		data                 []fileValues
		expect               map[string]metadata.DeltaPath
		canUsePreviousBackup bool
		expectError          assert.ErrorAssertionFunc
	}{
		{
			name: "delta urls only",
			data: []fileValues{
				{metadata.DeltaURLsFileName, "delta-link"},
			},
			expect:               map[string]metadata.DeltaPath{},
			canUsePreviousBackup: true,
			expectError:          assert.NoError,
		},
		{
			name: "multiple delta urls",
			data: []fileValues{
				{metadata.DeltaURLsFileName, "delta-link"},
				{metadata.DeltaURLsFileName, "delta-link-2"},
			},
			canUsePreviousBackup: false,
			expectError:          assert.Error,
		},
		{
			name: "previous path only",
			data: []fileValues{
				{metadata.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]metadata.DeltaPath{
				"key": {
					Delta: "delta-link",
					Path:  "prev-path",
				},
			},
			canUsePreviousBackup: true,
			expectError:          assert.NoError,
		},
		{
			name: "multiple previous paths",
			data: []fileValues{
				{metadata.PreviousPathFileName, "prev-path"},
				{metadata.PreviousPathFileName, "prev-path-2"},
			},
			canUsePreviousBackup: false,
			expectError:          assert.Error,
		},
		{
			name: "delta urls and previous paths",
			data: []fileValues{
				{metadata.DeltaURLsFileName, "delta-link"},
				{metadata.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]metadata.DeltaPath{
				"key": {
					Delta: "delta-link",
					Path:  "prev-path",
				},
			},
			canUsePreviousBackup: true,
			expectError:          assert.NoError,
		},
		{
			name: "delta urls and empty previous paths",
			data: []fileValues{
				{metadata.DeltaURLsFileName, "delta-link"},
				{metadata.PreviousPathFileName, ""},
			},
			expect:               map[string]metadata.DeltaPath{},
			canUsePreviousBackup: true,
			expectError:          assert.NoError,
		},
		{
			name: "empty delta urls and previous paths",
			data: []fileValues{
				{metadata.DeltaURLsFileName, ""},
				{metadata.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]metadata.DeltaPath{
				"key": {
					Delta: "delta-link",
					Path:  "prev-path",
				},
			},
			canUsePreviousBackup: true,
			expectError:          assert.NoError,
		},
		{
			name: "delta urls with special chars",
			data: []fileValues{
				{metadata.DeltaURLsFileName, "`!@#$%^&*()_[]{}/\"\\"},
				{metadata.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]metadata.DeltaPath{
				"key": {
					Delta: "`!@#$%^&*()_[]{}/\"\\",
					Path:  "prev-path",
				},
			},
			canUsePreviousBackup: true,
			expectError:          assert.NoError,
		},
		{
			name: "delta urls with escaped chars",
			data: []fileValues{
				{metadata.DeltaURLsFileName, `\n\r\t\b\f\v\0\\`},
				{metadata.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]metadata.DeltaPath{
				"key": {
					Delta: "\\n\\r\\t\\b\\f\\v\\0\\\\",
					Path:  "prev-path",
				},
			},
			canUsePreviousBackup: true,
			expectError:          assert.NoError,
		},
		{
			name: "delta urls with newline char runes",
			data: []fileValues{
				// rune(92) = \, rune(110) = n.  Ensuring it's not possible to
				// error in serializing/deserializing and produce a single newline
				// character from those two runes.
				{metadata.DeltaURLsFileName, string([]rune{rune(92), rune(110)})},
				{metadata.PreviousPathFileName, "prev-path"},
			},
			expect: map[string]metadata.DeltaPath{
				"key": {
					Delta: "\\n",
					Path:  "prev-path",
				},
			},
			canUsePreviousBackup: true,
			expectError:          assert.NoError,
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

			pathPrefix, err := path.BuildMetadata(
				"t", "u",
				path.ExchangeService,
				path.EmailCategory,
				false)
			require.NoError(t, err, "path prefix")

			coll, err := graph.MakeMetadataCollection(
				pathPrefix,
				entries,
				func(cos *support.ControllerOperationStatus) {})
			require.NoError(t, err, clues.ToCore(err))

			cdps, canUsePreviousBackup, err := ParseMetadataCollections(ctx, []data.RestoreCollection{
				dataMock.NewUnversionedRestoreCollection(t, data.NoFetchRestoreCollection{Collection: coll}),
			})
			test.expectError(t, err, clues.ToCore(err))

			assert.Equal(t, test.canUsePreviousBackup, canUsePreviousBackup, "can use previous backup")

			emails := cdps[path.EmailCategory]

			assert.Len(t, emails, len(test.expect))

			for k, v := range emails {
				assert.Equal(t, v.Delta, emails[k].Delta, "delta")
				assert.Equal(t, v.Path, emails[k].Path, "path")
			}
		})
	}
}

type failingColl struct {
	t *testing.T
}

func (f failingColl) Items(ctx context.Context, errs *fault.Bus) <-chan data.Item {
	ic := make(chan data.Item)
	defer close(ic)

	errs.AddRecoverable(ctx, assert.AnError)

	return ic
}

func (f failingColl) FullPath() path.Path {
	tmp, err := path.Build(
		"tenant",
		"user",
		path.ExchangeService,
		path.EmailCategory,
		false,
		"inbox")
	require.NoError(f.t, err, clues.ToCore(err))

	return tmp
}

func (f failingColl) FetchItemByName(context.Context, string) (data.Item, error) {
	// no fetch calls will be made
	return nil, nil
}

// This check is to ensure that we don't error out, but still return
// canUsePreviousBackup as false on read errors
func (suite *DataCollectionsUnitSuite) TestParseMetadataCollections_ReadFailure() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	fc := failingColl{t}

	_, canUsePreviousBackup, err := ParseMetadataCollections(ctx, []data.RestoreCollection{fc})
	require.NoError(t, err)
	require.False(t, canUsePreviousBackup)
}

// ---------------------------------------------------------------------------
// Integration tests
// ---------------------------------------------------------------------------

func newStatusUpdater(t *testing.T, wg *sync.WaitGroup) func(status *support.ControllerOperationStatus) {
	updater := func(status *support.ControllerOperationStatus) {
		defer wg.Done()
	}

	return updater
}

type BackupIntgSuite struct {
	tester.Suite
	user     string
	site     string
	tenantID string
	ac       api.Client
}

func TestBackupIntgSuite(t *testing.T) {
	suite.Run(t, &BackupIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs}),
	})
}

func (suite *BackupIntgSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	suite.user = tconfig.M365UserID(t)
	suite.site = tconfig.M365SiteID(t)

	acct := tconfig.NewM365Account(t)
	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.ac, err = api.NewClient(
		creds,
		control.DefaultOptions(),
		count.New())
	require.NoError(t, err, clues.ToCore(err))

	suite.tenantID = creds.AzureTenantID

	tester.LogTimeOfTest(t)
}

func (suite *BackupIntgSuite) TestMailFetch() {
	var (
		userID   = tconfig.M365UserID(suite.T())
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
				[]string{api.MailInbox},
				selectors.PrefixMatch())[0],
			folderNames: map[string]struct{}{
				api.MailInbox: {},
			},
			canMakeDeltaQueries: true,
		},
		{
			name: "Folder Iterative Check Mail Non-Delta",
			scope: selectors.NewExchangeBackup(users).MailFolders(
				[]string{api.MailInbox},
				selectors.PrefixMatch())[0],
			folderNames: map[string]struct{}{
				api.MailInbox: {},
			},
			canMakeDeltaQueries: false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			ctrlOpts := control.DefaultOptions()
			ctrlOpts.ToggleFeatures.DisableDelta = !test.canMakeDeltaQueries

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: version.NoBackup,
				Options:           ctrlOpts,
				ProtectedResource: inMock.NewProvider(userID, userID),
			}

			collections, err := CreateCollections(
				ctx,
				bpc,
				handlers,
				suite.tenantID,
				test.scope,
				metadata.DeltaPaths{},
				func(status *support.ControllerOperationStatus) {},
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

func (suite *BackupIntgSuite) TestDelta() {
	var (
		userID   = tconfig.M365UserID(suite.T())
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
				[]string{api.MailInbox},
				selectors.PrefixMatch())[0],
		},
		{
			name: "Contacts",
			scope: selectors.NewExchangeBackup(users).ContactFolders(
				[]string{api.DefaultContacts},
				selectors.PrefixMatch())[0],
		},
		{
			name: "Events",
			scope: selectors.NewExchangeBackup(users).EventCalendars(
				[]string{api.DefaultCalendar},
				selectors.PrefixMatch())[0],
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
				ProtectedResource: inMock.NewProvider(userID, userID),
			}

			// get collections without providing any delta history (ie: full backup)
			collections, err := CreateCollections(
				ctx,
				bpc,
				handlers,
				suite.tenantID,
				test.scope,
				metadata.DeltaPaths{},
				func(status *support.ControllerOperationStatus) {},
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

			cdps, canUsePreviousBackup, err := ParseMetadataCollections(ctx, []data.RestoreCollection{
				dataMock.NewUnversionedRestoreCollection(t, data.NoFetchRestoreCollection{Collection: metadata}),
			})
			require.NoError(t, err, clues.ToCore(err))
			assert.True(t, canUsePreviousBackup, "can use previous backup")

			dps := cdps[test.scope.Category().PathType()]

			// now do another backup with the previous delta tokens,
			// which should only contain the difference.
			_, err = CreateCollections(
				ctx,
				bpc,
				handlers,
				suite.tenantID,
				test.scope,
				dps,
				func(status *support.ControllerOperationStatus) {},
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
		})
	}
}

// TestMailSerializationRegression verifies that all mail data stored in the
// test account can be successfully downloaded into bytes and restored into
// M365 mail objects
func (suite *BackupIntgSuite) TestMailSerializationRegression() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		wg       sync.WaitGroup
		users    = []string{suite.user}
		handlers = BackupHandlers(suite.ac)
	)

	sel := selectors.NewExchangeBackup(users)
	sel.Include(sel.MailFolders([]string{api.MailInbox}, selectors.PrefixMatch()))

	bpc := inject.BackupProducerConfig{
		LastBackupVersion: version.NoBackup,
		Options:           control.DefaultOptions(),
		ProtectedResource: inMock.NewProvider(suite.user, suite.user),
		Selector:          sel.Selector,
	}

	collections, err := CreateCollections(
		ctx,
		bpc,
		handlers,
		suite.tenantID,
		sel.Scopes()[0],
		metadata.DeltaPaths{},
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

				rr, err := readers.NewVersionedRestoreReader(stream.ToReader())
				require.NoError(t, err, clues.ToCore(err))

				assert.Equal(t, readers.DefaultSerializationVersion, rr.Format().Version)

				read, err := buf.ReadFrom(rr)
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
// and to store contact within prefetchCollection. Downloaded contacts are run through
// a regression test to ensure that downloaded items can be uploaded.
func (suite *BackupIntgSuite) TestContactSerializationRegression() {
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
			scope: selectors.
				NewExchangeBackup(users).
				ContactFolders([]string{api.DefaultContacts}, selectors.PrefixMatch())[0],
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var wg sync.WaitGroup

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: version.NoBackup,
				Options:           control.DefaultOptions(),
				ProtectedResource: inMock.NewProvider(suite.user, suite.user),
			}

			edcs, err := CreateCollections(
				ctx,
				bpc,
				handlers,
				suite.tenantID,
				test.scope,
				metadata.DeltaPaths{},
				newStatusUpdater(t, &wg),
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			wg.Add(len(edcs))

			require.GreaterOrEqual(t, len(edcs), 1, "expected 1 <= num collections <= 2")
			require.GreaterOrEqual(t, 2, len(edcs), "expected 1 <= num collections <= 2")

			for _, edc := range edcs {
				var (
					isMetadata = edc.FullPath().Service() == path.ExchangeMetadataService
					count      = 0
				)

				for stream := range edc.Items(ctx, fault.New(true)) {
					buf := &bytes.Buffer{}

					rr, err := readers.NewVersionedRestoreReader(stream.ToReader())
					require.NoError(t, err, clues.ToCore(err))

					assert.Equal(t, readers.DefaultSerializationVersion, rr.Format().Version)

					read, err := buf.ReadFrom(rr)
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
					api.DefaultContacts)
				assert.NotZero(t, count)
			}

			wg.Wait()
		})
	}
}

// TestEventsSerializationRegression ensures functionality of createCollections
// to be able to successfully query, download and restore event objects
func (suite *BackupIntgSuite) TestEventsSerializationRegression() {
	var (
		users    = []string{suite.user}
		handlers = BackupHandlers(suite.ac)
	)

	tests := []struct {
		name, expectedContainerName string
		scope                       selectors.ExchangeScope
	}{
		{
			name:                  "Default Event Calendar",
			expectedContainerName: api.DefaultCalendar,
			scope: selectors.NewExchangeBackup(users).EventCalendars(
				[]string{api.DefaultCalendar},
				selectors.PrefixMatch())[0],
		},
		{
			name:                  "Birthday Calendar",
			expectedContainerName: "Birthdays",
			scope: selectors.NewExchangeBackup(users).EventCalendars(
				[]string{"Birthdays"},
				selectors.PrefixMatch())[0],
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var wg sync.WaitGroup

			bpc := inject.BackupProducerConfig{
				LastBackupVersion: version.NoBackup,
				Options:           control.DefaultOptions(),
				ProtectedResource: inMock.NewProvider(suite.user, suite.user),
			}

			collections, err := CreateCollections(
				ctx,
				bpc,
				handlers,
				suite.tenantID,
				test.scope,
				metadata.DeltaPaths{},
				newStatusUpdater(t, &wg),
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))
			require.Len(t, collections, 2)

			wg.Add(len(collections))

			for _, edc := range collections {
				dlp, isDLP := edc.(data.LocationPather)

				var isMetadata bool

				if edc.FullPath().Service() == path.ExchangeService {
					require.True(t, isDLP, "must be a location pather")
					assert.Contains(t, dlp.LocationPath().Elements(), test.expectedContainerName)
				} else {
					isMetadata = true
					assert.Empty(t, edc.FullPath().Folder(false))
				}

				for item := range edc.Items(ctx, fault.New(true)) {
					buf := &bytes.Buffer{}

					rr, err := readers.NewVersionedRestoreReader(item.ToReader())
					require.NoError(t, err, clues.ToCore(err))

					assert.Equal(t, readers.DefaultSerializationVersion, rr.Format().Version)

					read, err := buf.ReadFrom(rr)
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

type CollectionPopulationSuite struct {
	tester.Suite
	creds account.M365Config
}

func TestServiceIteratorsUnitSuite(t *testing.T) {
	suite.Run(t, &CollectionPopulationSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *CollectionPopulationSuite) SetupSuite() {
	a := tconfig.NewFakeM365Account(suite.T())
	m365, err := a.M365Config()
	require.NoError(suite.T(), err, clues.ToCore(err))
	suite.creds = m365
}

func (suite *CollectionPopulationSuite) TestPopulateCollections() {
	var (
		qp = graph.QueryParams{
			Category:          path.EmailCategory, // doesn't matter which one we use.
			ProtectedResource: inMock.NewProvider("user_id", "user_name"),
			TenantID:          suite.creds.AzureTenantID,
		}
		statusUpdater = func(*support.ControllerOperationStatus) {}
		allScope      = selectors.NewExchangeBackup(nil).MailFolders(selectors.Any())[0]
		dps           = metadata.DeltaPaths{} // incrementals are tested separately
		commonResult  = mockGetterResults{
			added:    []string{"a1", "a2", "a3"},
			removed:  []string{"r1", "r2", "r3"},
			newDelta: pagers.DeltaUpdate{URL: "delta_url"},
		}
		errorResult = mockGetterResults{
			added:    []string{"a1", "a2", "a3"},
			removed:  []string{"r1", "r2", "r3"},
			newDelta: pagers.DeltaUpdate{URL: "delta_url"},
			err:      assert.AnError,
		}
		deletedInFlightResult = mockGetterResults{
			added:    []string{"a1", "a2", "a3"},
			removed:  []string{"r1", "r2", "r3"},
			newDelta: pagers.DeltaUpdate{URL: "delta_url"},
			err:      graph.ErrDeletedInFlight,
		}
		container1 = mockContainer{
			id:          strPtr("1"),
			displayName: strPtr("display_name_1"),
			p:           path.Builder{}.Append("1"),
			l:           path.Builder{}.Append("display_name_1"),
		}
		container2 = mockContainer{
			id:          strPtr("2"),
			displayName: strPtr("display_name_2"),
			p:           path.Builder{}.Append("2"),
			l:           path.Builder{}.Append("display_name_2"),
		}
	)

	table := []struct {
		name                  string
		getter                mockGetter
		resolver              graph.ContainerResolver
		scope                 selectors.ExchangeScope
		failFast              control.FailurePolicy
		expectErr             assert.ErrorAssertionFunc
		expectNewColls        int
		expectMetadataColls   int
		expectDoNotMergeColls int
	}{
		{
			name: "happy path, one container",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResult,
				},
			},
			resolver:            newMockResolver(container1),
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      1,
			expectMetadataColls: 1,
		},
		{
			name: "happy path, many containers",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResult,
					"2": commonResult,
				},
			},
			resolver:            newMockResolver(container1, container2),
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      2,
			expectMetadataColls: 1,
		},
		{
			name: "no containers pass scope",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResult,
					"2": commonResult,
				},
			},
			resolver:            newMockResolver(container1, container2),
			scope:               selectors.NewExchangeBackup(nil).MailFolders(selectors.None())[0],
			expectErr:           assert.NoError,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
		{
			name: "err: deleted in flight",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": deletedInFlightResult,
				},
			},
			resolver:              newMockResolver(container1),
			scope:                 allScope,
			expectErr:             assert.NoError,
			expectNewColls:        1,
			expectMetadataColls:   1,
			expectDoNotMergeColls: 1,
		},
		{
			name: "err: other error",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": errorResult,
				},
			},
			resolver:            newMockResolver(container1),
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      0,
			expectMetadataColls: 1,
		},
		{
			name: "half collections error: deleted in flight",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": deletedInFlightResult,
					"2": commonResult,
				},
			},
			resolver:              newMockResolver(container1, container2),
			scope:                 allScope,
			expectErr:             assert.NoError,
			expectNewColls:        2,
			expectMetadataColls:   1,
			expectDoNotMergeColls: 1,
		},
		{
			name: "half collections error: other error",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": errorResult,
					"2": commonResult,
				},
			},
			resolver:            newMockResolver(container1, container2),
			scope:               allScope,
			expectErr:           assert.NoError,
			expectNewColls:      1,
			expectMetadataColls: 1,
		},
		{
			name: "half collections error: deleted in flight, fail fast",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": deletedInFlightResult,
					"2": commonResult,
				},
			},
			resolver:              newMockResolver(container1, container2),
			scope:                 allScope,
			failFast:              control.FailFast,
			expectErr:             assert.NoError,
			expectNewColls:        2,
			expectMetadataColls:   1,
			expectDoNotMergeColls: 1,
		},
		{
			name: "half collections error: other error, fail fast",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": errorResult,
					"2": commonResult,
				},
			},
			resolver:            newMockResolver(container1, container2),
			scope:               allScope,
			failFast:            control.FailFast,
			expectErr:           assert.Error,
			expectNewColls:      0,
			expectMetadataColls: 0,
		},
	}
	for _, test := range table {
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

				ctrlOpts := control.Options{FailureHandling: test.failFast}
				ctrlOpts.ToggleFeatures.DisableDelta = !canMakeDeltaQueries

				mbh := mockBackupHandler{
					mg:       test.getter,
					category: qp.Category,
				}

				collections, err := populateCollections(
					ctx,
					qp,
					mbh,
					statusUpdater,
					test.resolver,
					test.scope,
					dps,
					ctrlOpts,
					fault.New(test.failFast == control.FailFast))
				test.expectErr(t, err, clues.ToCore(err))

				// collection assertions

				deleteds, news, metadatas, doNotMerges := 0, 0, 0, 0
				for _, c := range collections {
					if c.FullPath().Service() == path.ExchangeMetadataService {
						metadatas++
						continue
					}

					if c.State() == data.DeletedState {
						deleteds++
					}

					if c.State() == data.NewState {
						news++
					}

					if c.DoNotMergeItems() {
						doNotMerges++
					}
				}

				assert.Zero(t, deleteds, "deleted collections")
				assert.Equal(t, test.expectNewColls, news, "new collections")
				assert.Equal(t, test.expectMetadataColls, metadatas, "metadata collections")
				assert.Equal(t, test.expectDoNotMergeColls, doNotMerges, "doNotMerge collections")

				// items in collections assertions
				for k, expect := range test.getter.results {
					coll := collections[k]

					if coll == nil {
						continue
					}

					exColl, ok := coll.(*prefetchCollection)
					require.True(t, ok, "collection is an *exchange.prefetchCollection")

					ids := [][]string{
						make([]string, 0, len(exColl.added)),
						make([]string, 0, len(exColl.removed)),
					}

					for id := range exColl.added {
						ids[0] = append(ids[0], id)
					}

					for id := range exColl.removed {
						ids[1] = append(ids[1], id)
					}

					assert.ElementsMatch(t, expect.added, ids[0], "added items")
					assert.ElementsMatch(t, expect.removed, ids[1], "removed items")
				}
			})
		}
	}
}

func checkMetadata(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	cat path.CategoryType,
	expect metadata.DeltaPaths,
	c data.BackupCollection,
) {
	catPaths, _, err := ParseMetadataCollections(
		ctx,
		[]data.RestoreCollection{
			dataMock.NewUnversionedRestoreCollection(t, data.NoFetchRestoreCollection{Collection: c}),
		})
	if !assert.NoError(t, err, "getting metadata", clues.ToCore(err)) {
		return
	}

	assert.Equal(t, expect, catPaths[cat])
}

func (suite *CollectionPopulationSuite) TestFilterContainersAndFillCollections_DuplicateFolders() {
	type scopeCat struct {
		scope selectors.ExchangeScope
		cat   path.CategoryType
	}

	var (
		qp = graph.QueryParams{
			ProtectedResource: inMock.NewProvider("user_id", "user_name"),
			TenantID:          suite.creds.AzureTenantID,
		}

		statusUpdater = func(*support.ControllerOperationStatus) {}

		dataTypes = []scopeCat{
			{
				scope: selectors.NewExchangeBackup(nil).MailFolders(selectors.Any())[0],
				cat:   path.EmailCategory,
			},
			{
				scope: selectors.NewExchangeBackup(nil).ContactFolders(selectors.Any())[0],
				cat:   path.ContactsCategory,
			},
			{
				scope: selectors.NewExchangeBackup(nil).EventCalendars(selectors.Any())[0],
				cat:   path.EventsCategory,
			},
		}

		location = path.Builder{}.Append("foo", "bar")

		result1 = mockGetterResults{
			added:    []string{"a1", "a2", "a3"},
			removed:  []string{"r1", "r2", "r3"},
			newDelta: pagers.DeltaUpdate{URL: "delta_url"},
		}
		result2 = mockGetterResults{
			added:    []string{"a4", "a5", "a6"},
			removed:  []string{"r4", "r5", "r6"},
			newDelta: pagers.DeltaUpdate{URL: "delta_url2"},
		}

		container1 = mockContainer{
			id:          strPtr("1"),
			displayName: strPtr("bar"),
			p:           path.Builder{}.Append("1"),
			l:           location,
		}
		container2 = mockContainer{
			id:          strPtr("2"),
			displayName: strPtr("bar"),
			p:           path.Builder{}.Append("2"),
			l:           location,
		}
	)

	oldPath1 := func(t *testing.T, cat path.CategoryType) path.Path {
		res, err := location.Append("1").ToDataLayerPath(
			suite.creds.AzureTenantID,
			qp.ProtectedResource.ID(),
			path.ExchangeService,
			cat,
			false)
		require.NoError(t, err, clues.ToCore(err))

		return res
	}

	oldPath2 := func(t *testing.T, cat path.CategoryType) path.Path {
		res, err := location.Append("2").ToDataLayerPath(
			suite.creds.AzureTenantID,
			qp.ProtectedResource.ID(),
			path.ExchangeService,
			cat,
			false)
		require.NoError(t, err, clues.ToCore(err))

		return res
	}

	idPath1 := func(t *testing.T, cat path.CategoryType) path.Path {
		res, err := path.Builder{}.Append("1").ToDataLayerPath(
			suite.creds.AzureTenantID,
			qp.ProtectedResource.ID(),
			path.ExchangeService,
			cat,
			false)
		require.NoError(t, err, clues.ToCore(err))

		return res
	}

	idPath2 := func(t *testing.T, cat path.CategoryType) path.Path {
		res, err := path.Builder{}.Append("2").ToDataLayerPath(
			suite.creds.AzureTenantID,
			qp.ProtectedResource.ID(),
			path.ExchangeService,
			cat,
			false)
		require.NoError(t, err, clues.ToCore(err))

		return res
	}

	table := []struct {
		name           string
		getter         mockGetter
		resolver       graph.ContainerResolver
		inputMetadata  func(t *testing.T, cat path.CategoryType) metadata.DeltaPaths
		expectNewColls int
		expectDeleted  int
		expectMetadata func(t *testing.T, cat path.CategoryType) metadata.DeltaPaths
	}{
		{
			name: "1 moved to duplicate",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": result1,
					"2": result2,
				},
			},
			resolver: newMockResolver(container1, container2),
			inputMetadata: func(t *testing.T, cat path.CategoryType) metadata.DeltaPaths {
				return metadata.DeltaPaths{
					"1": metadata.DeltaPath{
						Delta: "old_delta",
						Path:  oldPath1(t, cat).String(),
					},
					"2": metadata.DeltaPath{
						Delta: "old_delta",
						Path:  idPath2(t, cat).String(),
					},
				}
			},
			expectMetadata: func(t *testing.T, cat path.CategoryType) metadata.DeltaPaths {
				return metadata.DeltaPaths{
					"1": metadata.DeltaPath{
						Delta: "delta_url",
						Path:  idPath1(t, cat).String(),
					},
					"2": metadata.DeltaPath{
						Delta: "delta_url2",
						Path:  idPath2(t, cat).String(),
					},
				}
			},
		},
		{
			name: "both move to duplicate",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": result1,
					"2": result2,
				},
			},
			resolver: newMockResolver(container1, container2),
			inputMetadata: func(t *testing.T, cat path.CategoryType) metadata.DeltaPaths {
				return metadata.DeltaPaths{
					"1": metadata.DeltaPath{
						Delta: "old_delta",
						Path:  oldPath1(t, cat).String(),
					},
					"2": metadata.DeltaPath{
						Delta: "old_delta",
						Path:  oldPath2(t, cat).String(),
					},
				}
			},
			expectMetadata: func(t *testing.T, cat path.CategoryType) metadata.DeltaPaths {
				return metadata.DeltaPaths{
					"1": metadata.DeltaPath{
						Delta: "delta_url",
						Path:  idPath1(t, cat).String(),
					},
					"2": metadata.DeltaPath{
						Delta: "delta_url2",
						Path:  idPath2(t, cat).String(),
					},
				}
			},
		},
		{
			name: "both new",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": result1,
					"2": result2,
				},
			},
			resolver: newMockResolver(container1, container2),
			inputMetadata: func(t *testing.T, cat path.CategoryType) metadata.DeltaPaths {
				return metadata.DeltaPaths{}
			},
			expectNewColls: 2,
			expectMetadata: func(t *testing.T, cat path.CategoryType) metadata.DeltaPaths {
				return metadata.DeltaPaths{
					"1": metadata.DeltaPath{
						Delta: "delta_url",
						Path:  idPath1(t, cat).String(),
					},
					"2": metadata.DeltaPath{
						Delta: "delta_url2",
						Path:  idPath2(t, cat).String(),
					},
				}
			},
		},
		{
			name: "add 1 remove 2",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": result1,
				},
			},
			resolver: newMockResolver(container1),
			inputMetadata: func(t *testing.T, cat path.CategoryType) metadata.DeltaPaths {
				return metadata.DeltaPaths{
					"2": metadata.DeltaPath{
						Delta: "old_delta",
						Path:  idPath2(t, cat).String(),
					},
				}
			},
			expectNewColls: 1,
			expectDeleted:  1,
			expectMetadata: func(t *testing.T, cat path.CategoryType) metadata.DeltaPaths {
				return metadata.DeltaPaths{
					"1": metadata.DeltaPath{
						Delta: "delta_url",
						Path:  idPath1(t, cat).String(),
					},
				}
			},
		},
	}

	for _, sc := range dataTypes {
		suite.Run(sc.cat.String(), func() {
			qp.Category = sc.cat

			for _, test := range table {
				suite.Run(test.name, func() {
					t := suite.T()

					ctx, flush := tester.NewContext(t)
					defer flush()

					mbh := mockBackupHandler{
						mg:       test.getter,
						category: qp.Category,
					}

					collections, err := populateCollections(
						ctx,
						qp,
						mbh,
						statusUpdater,
						test.resolver,
						sc.scope,
						test.inputMetadata(t, qp.Category),
						control.Options{FailureHandling: control.FailFast},
						fault.New(true))
					require.NoError(t, err, "getting collections", clues.ToCore(err))

					// collection assertions

					deleteds, news, metadatas := 0, 0, 0
					for _, c := range collections {
						if c.State() == data.DeletedState {
							deleteds++
							continue
						}

						if c.FullPath().Service() == path.ExchangeMetadataService {
							metadatas++
							checkMetadata(t, ctx, qp.Category, test.expectMetadata(t, qp.Category), c)
							continue
						}

						if c.State() == data.NewState {
							news++
						}
					}

					assert.Equal(t, test.expectDeleted, deleteds, "deleted collections")
					assert.Equal(t, test.expectNewColls, news, "new collections")
					assert.Equal(t, 1, metadatas, "metadata collections")

					// items in collections assertions
					for k, expect := range test.getter.results {
						coll := collections[k]

						if coll == nil {
							continue
						}

						exColl, ok := coll.(*prefetchCollection)
						require.True(t, ok, "collection is an *exchange.prefetchCollection")

						ids := [][]string{
							make([]string, 0, len(exColl.added)),
							make([]string, 0, len(exColl.removed)),
						}

						for id := range exColl.added {
							ids[0] = append(ids[0], id)
						}

						for id := range exColl.removed {
							ids[1] = append(ids[1], id)
						}

						assert.ElementsMatch(t, expect.added, ids[0], "added items")
						assert.ElementsMatch(t, expect.removed, ids[1], "removed items")
					}
				})
			}
		})
	}
}

func (suite *CollectionPopulationSuite) TestFilterContainersAndFillCollections_repeatedItems() {
	newDelta := pagers.DeltaUpdate{URL: "delta_url"}

	table := []struct {
		name          string
		getter        mockGetter
		expectAdded   map[string]struct{}
		expectRemoved map[string]struct{}
	}{
		{
			name: "repeated adds",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": {
						added:    []string{"a1", "a2", "a3", "a1"},
						newDelta: newDelta,
					},
				},
			},
			expectAdded: map[string]struct{}{
				"a1": {},
				"a2": {},
				"a3": {},
			},
			expectRemoved: map[string]struct{}{},
		},
		{
			name: "repeated removes",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": {
						removed:  []string{"r1", "r2", "r3", "r1"},
						newDelta: newDelta,
					},
				},
			},
			expectAdded: map[string]struct{}{},
			expectRemoved: map[string]struct{}{
				"r1": {},
				"r2": {},
				"r3": {},
			},
		},
		{
			name: "remove for same item wins",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": {
						added:    []string{"i1", "a2", "a3"},
						removed:  []string{"i1", "r2", "r3"},
						newDelta: newDelta,
					},
				},
			},
			expectAdded: map[string]struct{}{
				"a2": {},
				"a3": {},
			},
			expectRemoved: map[string]struct{}{
				"i1": {},
				"r2": {},
				"r3": {},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				qp = graph.QueryParams{
					Category:          path.EmailCategory, // doesn't matter which one we use.
					ProtectedResource: inMock.NewProvider("user_id", "user_name"),
					TenantID:          suite.creds.AzureTenantID,
				}
				statusUpdater = func(*support.ControllerOperationStatus) {}
				allScope      = selectors.NewExchangeBackup(nil).MailFolders(selectors.Any())[0]
				dps           = metadata.DeltaPaths{} // incrementals are tested separately
				container1    = mockContainer{
					id:          strPtr("1"),
					displayName: strPtr("display_name_1"),
					p:           path.Builder{}.Append("1"),
					l:           path.Builder{}.Append("display_name_1"),
				}
				resolver = newMockResolver(container1)
				mbh      = mockBackupHandler{
					mg:       test.getter,
					category: qp.Category,
				}
			)

			require.Equal(t, "user_id", qp.ProtectedResource.ID(), qp.ProtectedResource)
			require.Equal(t, "user_name", qp.ProtectedResource.Name(), qp.ProtectedResource)

			collections, err := populateCollections(
				ctx,
				qp,
				mbh,
				statusUpdater,
				resolver,
				allScope,
				dps,
				control.Options{FailureHandling: control.FailFast},
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			// collection assertions

			deleteds, news, metadatas, doNotMerges := 0, 0, 0, 0
			for _, c := range collections {
				if c.FullPath().Service() == path.ExchangeMetadataService {
					metadatas++
					continue
				}

				if c.State() == data.DeletedState {
					deleteds++
				}

				if c.State() == data.NewState {
					news++
				}

				if c.DoNotMergeItems() {
					doNotMerges++
				}
			}

			assert.Zero(t, deleteds, "deleted collections")
			assert.Equal(t, 1, news, "new collections")
			assert.Equal(t, 1, metadatas, "metadata collections")
			assert.Zero(t, doNotMerges, "doNotMerge collections")

			// items in collections assertions
			for k := range test.getter.results {
				coll := collections[k]
				if !assert.NotNilf(t, coll, "missing collection for path %s", k) {
					continue
				}

				exColl, ok := coll.(*prefetchCollection)
				require.True(t, ok, "collection is an *exchange.prefetchCollection")

				assert.ElementsMatch(
					t,
					maps.Keys(test.expectAdded),
					maps.Keys(exColl.added),
					"added items")
				assert.Equal(t, test.expectRemoved, exColl.removed, "removed items")
			}
		})
	}
}

func (suite *CollectionPopulationSuite) TestFilterContainersAndFillCollections_PreviewBackup() {
	type itemContainer struct {
		container mockContainer
		added     []string
		removed   []string
	}

	type expected struct {
		mustHave  []itemContainer
		maybeHave []itemContainer
		// numItems is the total number of added items to expect. Needed because
		// some tests can return one of a set of items depending on the order
		// containers are processed in.
		numItems int
	}

	var (
		containers []mockContainer
		newDelta   = pagers.DeltaUpdate{URL: "delta_url"}
	)

	for i := 0; i < 10; i++ {
		id := fmt.Sprintf("%d", i)
		name := fmt.Sprintf("display_name_%d", i)
		containers = append(containers, mockContainer{
			id:          strPtr(id),
			displayName: strPtr(name),
			p:           path.Builder{}.Append(id),
			l:           path.Builder{}.Append(name),
		})
	}

	table := []struct {
		name     string
		limits   control.PreviewItemLimits
		data     []itemContainer
		includes []string
		excludes []string
		expect   expected
	}{
		{
			name: "IncludeContainer NoItemLimit ContainerLimit",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        1,
			},
			data: []itemContainer{
				{
					container: containers[0],
					added:     []string{"a1", "a2", "a3", "a4", "a5"},
				},
				{
					container: containers[1],
					added:     []string{"a6", "a7", "a8", "a9", "a10"},
				},
				{
					container: containers[2],
					added:     []string{"a11", "a12", "a13", "a14", "a15"},
				},
			},
			includes: []string{ptr.Val(containers[1].GetId())},
			expect: expected{
				mustHave: []itemContainer{
					{
						container: containers[1],
						added:     []string{"a6", "a7", "a8", "a9", "a10"},
					},
				},
				numItems: 5,
			},
		},
		{
			name: "IncludeContainer ItemLimit ContainerLimit",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             3,
				MaxItemsPerContainer: 999,
				MaxContainers:        1,
			},
			data: []itemContainer{
				{
					container: containers[0],
					added:     []string{"a1", "a2", "a3", "a4", "a5"},
				},
				{
					container: containers[1],
					added:     []string{"a6", "a7", "a8", "a9", "a10"},
				},
				{
					container: containers[2],
					added:     []string{"a11", "a12", "a13", "a14", "a15"},
				},
			},
			includes: []string{ptr.Val(containers[1].GetId())},
			expect: expected{
				maybeHave: []itemContainer{
					{
						container: containers[1],
						added:     []string{"a6", "a7", "a8", "a9", "a10"},
					},
				},
				numItems: 3,
			},
		},
		{
			name: "IncludeContainer ItemLimit NoContainerLimit",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             8,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
			},
			data: []itemContainer{
				{
					container: containers[0],
					added:     []string{"a1", "a2", "a3", "a4", "a5"},
				},
				{
					container: containers[1],
					added:     []string{"a6", "a7", "a8", "a9", "a10"},
				},
				{
					container: containers[2],
					added:     []string{"a11", "a12", "a13", "a14", "a15"},
				},
			},
			includes: []string{ptr.Val(containers[1].GetId())},
			expect: expected{
				mustHave: []itemContainer{
					{
						container: containers[1],
						added:     []string{"a6", "a7", "a8", "a9", "a10"},
					},
				},
				maybeHave: []itemContainer{
					{
						container: containers[0],
						added:     []string{"a1", "a2", "a3", "a4", "a5"},
					},
					{
						container: containers[2],
						added:     []string{"a11", "a12", "a13", "a14", "a15"},
					},
				},
				numItems: 8,
			},
		},
		{
			name: "PerContainerItemLimit NoContainerLimit",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 3,
				MaxContainers:        999,
			},
			data: []itemContainer{
				{
					container: containers[0],
					added:     []string{"a1", "a2", "a3", "a4", "a5"},
				},
				{
					container: containers[1],
					added:     []string{"a6", "a7", "a8", "a9", "a10"},
				},
				{
					container: containers[2],
					added:     []string{"a11", "a12", "a13", "a14", "a15"},
				},
			},
			expect: expected{
				// The test isn't setup to handle partial containers so the best we can
				// do is check that all items are expected and the item limit is hit.
				maybeHave: []itemContainer{
					{
						container: containers[1],
						added:     []string{"a6", "a7", "a8", "a9", "a10"},
					},
					{
						container: containers[0],
						added:     []string{"a1", "a2", "a3", "a4", "a5"},
					},
					{
						container: containers[2],
						added:     []string{"a11", "a12", "a13", "a14", "a15"},
					},
				},
				numItems: 9,
			},
		},
		{
			name: "ExcludeContainer NoLimits",
			limits: control.PreviewItemLimits{
				Enabled:              true,
				MaxItems:             999,
				MaxItemsPerContainer: 999,
				MaxContainers:        999,
			},
			excludes: []string{ptr.Val(containers[1].GetId())},
			data: []itemContainer{
				{
					container: containers[0],
					added:     []string{"a1", "a2", "a3", "a4", "a5"},
				},
				{
					container: containers[1],
					added:     []string{"a6", "a7", "a8", "a9", "a10"},
				},
				{
					container: containers[2],
					added:     []string{"a11", "a12", "a13", "a14", "a15"},
				},
			},
			expect: expected{
				// The test isn't setup to handle partial containers so the best we can
				// do is check that all items are expected and the item limit is hit.
				maybeHave: []itemContainer{
					{
						container: containers[0],
						added:     []string{"a1", "a2", "a3", "a4", "a5"},
					},
					{
						container: containers[2],
						added:     []string{"a11", "a12", "a13", "a14", "a15"},
					},
				},
				numItems: 10,
			},
		},
		{
			name: "NotPreview IgnoresLimits",
			limits: control.PreviewItemLimits{
				MaxItems:             1,
				MaxItemsPerContainer: 1,
				MaxContainers:        1,
			},
			excludes: []string{ptr.Val(containers[1].GetId())},
			data: []itemContainer{
				{
					container: containers[0],
					added:     []string{"a1", "a2", "a3", "a4", "a5"},
				},
				{
					container: containers[1],
					added:     []string{"a6", "a7", "a8", "a9", "a10"},
				},
				{
					container: containers[2],
					added:     []string{"a11", "a12", "a13", "a14", "a15"},
				},
			},
			expect: expected{
				mustHave: []itemContainer{
					{
						container: containers[0],
						added:     []string{"a1", "a2", "a3", "a4", "a5"},
					},
					{
						container: containers[1],
						added:     []string{"a6", "a7", "a8", "a9", "a10"},
					},
					{
						container: containers[2],
						added:     []string{"a11", "a12", "a13", "a14", "a15"},
					},
				},
				numItems: 15,
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				qp = graph.QueryParams{
					Category:          path.EmailCategory, // doesn't matter which one we use.
					ProtectedResource: inMock.NewProvider("user_id", "user_name"),
					TenantID:          suite.creds.AzureTenantID,
				}
				statusUpdater = func(*support.ControllerOperationStatus) {}
				allScope      = selectors.NewExchangeBackup(nil).MailFolders(selectors.Any())[0]
				dps           = metadata.DeltaPaths{} // incrementals are tested separately
			)

			inputContainers := make([]mockContainer, 0, len(test.data))
			inputItems := map[string]mockGetterResults{}

			for _, item := range test.data {
				inputContainers = append(inputContainers, item.container)
				inputItems[ptr.Val(item.container.GetId())] = mockGetterResults{
					added:    item.added,
					removed:  item.removed,
					newDelta: newDelta,
				}
			}

			// Make sure concurrency limit is initialized to a non-zero value or we'll
			// deadlock.
			opts := control.DefaultOptions()
			opts.FailureHandling = control.FailFast
			opts.PreviewLimits = test.limits

			resolver := newMockResolver(inputContainers...)
			getter := mockGetter{results: inputItems}
			mbh := mockBackupHandler{
				mg:              getter,
				fg:              resolver,
				category:        qp.Category,
				previewIncludes: test.includes,
				previewExcludes: test.excludes,
			}

			require.Equal(t, "user_id", qp.ProtectedResource.ID(), qp.ProtectedResource)
			require.Equal(t, "user_name", qp.ProtectedResource.Name(), qp.ProtectedResource)

			collections, err := populateCollections(
				ctx,
				qp,
				mbh,
				statusUpdater,
				resolver,
				allScope,
				dps,
				opts,
				fault.New(true))
			require.NoError(t, err, clues.ToCore(err))

			var totalItems int

			// collection assertions
			for _, c := range collections {
				if c.FullPath().Service() == path.ExchangeMetadataService {
					continue
				}

				// We don't expect any deleted containers in this test.
				if !assert.NotEqual(
					t,
					data.DeletedState,
					c.State(),
					"container marked deleted") {
					continue
				}

				// TODO(ashmrtn): Remove when we make LocationPath part of the
				// Collection interface.
				lp := c.(data.LocationPather)
				mustHave := map[string]struct{}{}
				maybeHave := map[string]struct{}{}

				containerKey := lp.LocationPath().String()

				for _, item := range test.expect.mustHave {
					// Get the right container of items.
					if containerKey != item.container.l.String() {
						continue
					}

					for _, id := range item.added {
						mustHave[id] = struct{}{}
					}
				}

				for _, item := range test.expect.maybeHave {
					// Get the right container of items.
					if containerKey != item.container.l.String() {
						continue
					}

					for _, id := range item.added {
						maybeHave[id] = struct{}{}
					}
				}

				errs := fault.New(true)

				for item := range c.Items(ctx, errs) {
					// We don't expect deleted items in the test or in practice because we
					// never reuse delta tokens for preview backups.
					if item.Deleted() {
						continue
					}

					totalItems++

					var found bool

					if _, found = mustHave[item.ID()]; found {
						delete(mustHave, item.ID())
						continue
					}

					if _, found = maybeHave[item.ID()]; found {
						delete(maybeHave, item.ID())
						continue
					}

					assert.True(t, found, "unexpected item %v", item.ID())
				}
				require.NoError(t, errs.Failure())

				assert.Empty(
					t,
					mustHave,
					"container %v missing required items",
					lp.LocationPath().String())
			}

			assert.Equal(
				t,
				test.expect.numItems,
				totalItems,
				"total items seen across collections")
		})
	}
}

func (suite *CollectionPopulationSuite) TestFilterContainersAndFillCollections_incrementals_nondelta() {
	var (
		userID   = "user_id"
		tenantID = suite.creds.AzureTenantID
		cat      = path.EmailCategory // doesn't matter which one we use,
		qp       = graph.QueryParams{
			Category:          cat,
			ProtectedResource: inMock.NewProvider("user_id", "user_name"),
			TenantID:          suite.creds.AzureTenantID,
		}
		statusUpdater = func(*support.ControllerOperationStatus) {}
		allScope      = selectors.NewExchangeBackup(nil).MailFolders(selectors.Any())[0]
		commonResults = mockGetterResults{
			added:    []string{"added"},
			newDelta: pagers.DeltaUpdate{URL: "new_delta_url"},
		}
		expiredResults = mockGetterResults{
			added: []string{"added"},
			newDelta: pagers.DeltaUpdate{
				URL:   "new_delta_url",
				Reset: true,
			},
		}
	)

	prevPath := func(t *testing.T, at ...string) path.Path {
		p, err := path.Build(tenantID, userID, path.ExchangeService, cat, false, at...)
		require.NoError(t, err, clues.ToCore(err))

		return p
	}

	type endState struct {
		state      data.CollectionState
		doNotMerge bool
	}

	table := []struct {
		name                  string
		getter                mockGetter
		resolver              graph.ContainerResolver
		dps                   metadata.DeltaPaths
		expect                map[string]endState
		skipWhenForcedNoDelta bool
	}{
		{
			name: "new container",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("new"),
				p:           path.Builder{}.Append("1", "new"),
				l:           path.Builder{}.Append("1", "new"),
			}),
			dps: metadata.DeltaPaths{},
			expect: map[string]endState{
				"1": {data.NewState, false},
			},
		},
		{
			name: "not moved container",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("not_moved"),
				p:           path.Builder{}.Append("1", "not_moved"),
				l:           path.Builder{}.Append("1", "not_moved"),
			}),
			dps: metadata.DeltaPaths{
				"1": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "not_moved").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.NotMovedState, false},
			},
		},
		{
			name: "moved container",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("moved"),
				p:           path.Builder{}.Append("1", "moved"),
				l:           path.Builder{}.Append("1", "moved"),
			}),
			dps: metadata.DeltaPaths{
				"1": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "prev").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.MovedState, false},
			},
		},
		{
			name: "deleted container",
			getter: mockGetter{
				results: map[string]mockGetterResults{},
			},
			resolver: newMockResolver(),
			dps: metadata.DeltaPaths{
				"1": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "deleted").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.DeletedState, false},
			},
		},
		{
			name: "one deleted, one new",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"2": commonResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("2"),
				displayName: strPtr("new"),
				p:           path.Builder{}.Append("2", "new"),
				l:           path.Builder{}.Append("2", "new"),
			}),
			dps: metadata.DeltaPaths{
				"1": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "deleted").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.DeletedState, false},
				"2": {data.NewState, false},
			},
		},
		{
			name: "one deleted, one new, same path",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"2": commonResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("2"),
				displayName: strPtr("same"),
				p:           path.Builder{}.Append("2", "same"),
				l:           path.Builder{}.Append("2", "same"),
			}),
			dps: metadata.DeltaPaths{
				"1": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "same").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.DeletedState, false},
				"2": {data.NewState, false},
			},
		},
		{
			name: "one moved, one new, same path",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResults,
					"2": commonResults,
				},
			},
			resolver: newMockResolver(
				mockContainer{
					id:          strPtr("1"),
					displayName: strPtr("moved"),
					p:           path.Builder{}.Append("1", "moved"),
					l:           path.Builder{}.Append("1", "moved"),
				},
				mockContainer{
					id:          strPtr("2"),
					displayName: strPtr("prev"),
					p:           path.Builder{}.Append("2", "prev"),
					l:           path.Builder{}.Append("2", "prev"),
				}),
			dps: metadata.DeltaPaths{
				"1": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "prev").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.MovedState, false},
				"2": {data.NewState, false},
			},
		},
		{
			name: "bad previous path strings",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("not_moved"),
				p:           path.Builder{}.Append("1", "not_moved"),
				l:           path.Builder{}.Append("1", "not_moved"),
			}),
			dps: metadata.DeltaPaths{
				"1": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  "1/fnords/mc/smarfs",
				},
				"2": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  "2/fnords/mc/smarfs",
				},
			},
			expect: map[string]endState{
				"1": {data.NewState, false},
			},
		},
		{
			name: "delta expiration",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": expiredResults,
				},
			},
			resolver: newMockResolver(mockContainer{
				id:          strPtr("1"),
				displayName: strPtr("same"),
				p:           path.Builder{}.Append("1", "same"),
				l:           path.Builder{}.Append("1", "same"),
			}),
			dps: metadata.DeltaPaths{
				"1": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "1", "same").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.NotMovedState, true},
			},
			skipWhenForcedNoDelta: true, // this is not a valid test for non-delta
		},
		{
			name: "a little bit of everything",
			getter: mockGetter{
				results: map[string]mockGetterResults{
					"1": commonResults,  // new
					"2": commonResults,  // notMoved
					"3": commonResults,  // moved
					"4": expiredResults, // moved
					// "5" gets deleted
				},
			},
			resolver: newMockResolver(
				mockContainer{
					id:          strPtr("1"),
					displayName: strPtr("new"),
					p:           path.Builder{}.Append("1", "new"),
					l:           path.Builder{}.Append("1", "new"),
				},
				mockContainer{
					id:          strPtr("2"),
					displayName: strPtr("not_moved"),
					p:           path.Builder{}.Append("2", "not_moved"),
					l:           path.Builder{}.Append("2", "not_moved"),
				},
				mockContainer{
					id:          strPtr("3"),
					displayName: strPtr("moved"),
					p:           path.Builder{}.Append("3", "moved"),
					l:           path.Builder{}.Append("3", "moved"),
				},
				mockContainer{
					id:          strPtr("4"),
					displayName: strPtr("moved"),
					p:           path.Builder{}.Append("4", "moved"),
					l:           path.Builder{}.Append("4", "moved"),
				}),
			dps: metadata.DeltaPaths{
				"2": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "2", "not_moved").String(),
				},
				"3": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "3", "prev").String(),
				},
				"4": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "4", "prev").String(),
				},
				"5": metadata.DeltaPath{
					Delta: "old_delta_url",
					Path:  prevPath(suite.T(), "5", "deleted").String(),
				},
			},
			expect: map[string]endState{
				"1": {data.NewState, false},
				"2": {data.NotMovedState, false},
				"3": {data.MovedState, false},
				"4": {data.MovedState, true},
				"5": {data.DeletedState, false},
			},
			skipWhenForcedNoDelta: true,
		},
	}
	for _, test := range table {
		for _, deltaBefore := range []bool{true, false} {
			for _, deltaAfter := range []bool{true, false} {
				name := test.name

				if deltaAfter {
					name += "-delta"
				} else {
					if test.skipWhenForcedNoDelta {
						suite.T().Skip("intentionally skipped non-delta case")
					}
					name += "-non-delta"
				}

				suite.Run(name, func() {
					t := suite.T()

					ctx, flush := tester.NewContext(t)
					defer flush()

					ctrlOpts := control.DefaultOptions()
					ctrlOpts.ToggleFeatures.DisableDelta = !deltaAfter

					getter := test.getter
					if !deltaAfter {
						getter.noReturnDelta = false
					}

					mbh := mockBackupHandler{
						mg:       test.getter,
						category: qp.Category,
					}

					dps := test.dps
					if !deltaBefore {
						for k, dp := range dps {
							dp.Delta = ""
							dps[k] = dp
						}
					}

					collections, err := populateCollections(
						ctx,
						qp,
						mbh,
						statusUpdater,
						test.resolver,
						allScope,
						test.dps,
						ctrlOpts,
						fault.New(true))
					assert.NoError(t, err, clues.ToCore(err))

					metadatas := 0
					for _, c := range collections {
						p := c.FullPath()
						if p == nil {
							p = c.PreviousPath()
						}

						require.NotNil(t, p)

						if p.Service() == path.ExchangeMetadataService {
							metadatas++
							continue
						}

						p0 := p.Folders()[0]

						expect, ok := test.expect[p0]
						assert.True(t, ok, "collection is expected in result")

						assert.Equalf(t, expect.state, c.State(), "collection %s state", p0)
						assert.Equalf(t, expect.doNotMerge, c.DoNotMergeItems(), "collection %s DoNotMergeItems", p0)
					}

					assert.Equal(t, 1, metadatas, "metadata collections")
				})
			}
		}
	}
}
