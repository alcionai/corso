package operations

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/m365"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	"github.com/alcionai/corso/src/internal/operations/inject/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/backup/metadata"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

type mockColl struct {
	id string // for comparisons
	p  path.Path
}

func (mc mockColl) Items(context.Context, *fault.Bus) <-chan data.Item {
	return nil
}

func (mc mockColl) FullPath() path.Path {
	return mc.p
}

type mockBackupFinder struct {
	// ResourceOwner -> returned set of data for call to FindBases. We can just
	// switch on the ResourceOwner as the passed in Reasons should be the same
	// beyond that and results are returned for the union of the reasons anyway.
	// This does assume that the return data is properly constructed to return a
	// union of the reasons etc.
	data map[string]kopia.BackupBases
}

func (bf *mockBackupFinder) FindBases(
	_ context.Context,
	reasons []identity.Reasoner,
	_ map[string]string,
) kopia.BackupBases {
	if len(reasons) == 0 {
		return kopia.NewMockBackupBases()
	}

	if bf == nil {
		return kopia.NewMockBackupBases()
	}

	b := bf.data[reasons[0].ProtectedResource()]
	if b == nil {
		return kopia.NewMockBackupBases()
	}

	return b
}

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type OperationsManifestsUnitSuite struct {
	tester.Suite
}

func TestOperationsManifestsUnitSuite(t *testing.T) {
	suite.Run(t, &OperationsManifestsUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *OperationsManifestsUnitSuite) TestGetMetadataPaths() {
	const (
		ro  = "owner"
		tid = "tenantid"
	)

	t := suite.T()

	var (
		emailPath = makeMetadataBasePath(
			suite.T(),
			tid,
			path.ExchangeService,
			ro,
			path.EmailCategory)
		contactPath = makeMetadataBasePath(
			suite.T(),
			tid,
			path.ExchangeService,
			ro,
			path.ContactsCategory)
		spLibsPath = makeMetadataBasePath(
			suite.T(),
			tid,
			path.SharePointService,
			ro,
			path.LibrariesCategory)
		messagesPath = makeMetadataBasePath(
			suite.T(),
			tid,
			path.GroupsService,
			ro,
			path.ChannelMessagesCategory)
		groupLibsPath = makeMetadataBasePath(
			suite.T(),
			tid,
			path.GroupsService,
			ro,
			path.LibrariesCategory)
	)

	groupLibsSitesPath, err := groupLibsPath.Append(false, odConsts.SitesPathDir)
	assert.NoError(t, err, clues.ToCore(err))

	groupLibsSite1Path, err := groupLibsSitesPath.Append(false, "site1")
	assert.NoError(t, err, clues.ToCore(err))

	groupLibsSite2Path, err := groupLibsSitesPath.Append(false, "site2")
	assert.NoError(t, err, clues.ToCore(err))

	getRestorePaths := func(t *testing.T, base path.Path, paths []string) []path.RestorePaths {
		ps := []path.RestorePaths{}

		for _, f := range paths {
			p, err := base.AppendItem(f)
			assert.NoError(t, err, clues.ToCore(err))

			ps = append(ps, path.RestorePaths{StoragePath: p, RestorePath: base})
		}

		return ps
	}

	table := []struct {
		name               string
		manID              string
		reasons            []identity.Reasoner
		preFetchPaths      []string
		preFetchCollection []data.RestoreCollection
		expectPaths        func(*testing.T, []string) []path.Path
		restorePaths       []path.RestorePaths
		expectErr          error
	}{
		{
			name:  "single reason",
			manID: "single",
			reasons: []identity.Reasoner{
				identity.NewReason(tid, ro, path.ExchangeService, path.EmailCategory),
			},
			preFetchPaths: []string{},
			expectPaths: func(t *testing.T, files []string) []path.Path {
				ps := make([]path.Path, 0, len(files))

				for _, f := range files {
					p, err := emailPath.AppendItem(f)
					assert.NoError(t, err, clues.ToCore(err))
					ps = append(ps, p)
				}

				return ps
			},
			restorePaths: getRestorePaths(t, emailPath, metadata.AllMetadataFileNames()),
		},
		{
			name:  "multiple reasons",
			manID: "multi",
			reasons: []identity.Reasoner{
				identity.NewReason(tid, ro, path.ExchangeService, path.EmailCategory),
				identity.NewReason(tid, ro, path.ExchangeService, path.ContactsCategory),
			},
			preFetchPaths: []string{},
			expectPaths: func(t *testing.T, files []string) []path.Path {
				ps := make([]path.Path, 0, len(files))

				for _, f := range files {
					p, err := emailPath.AppendItem(f)
					assert.NoError(t, err, clues.ToCore(err))
					ps = append(ps, p)
					p, err = contactPath.AppendItem(f)
					assert.NoError(t, err, clues.ToCore(err))
					ps = append(ps, p)
				}

				return ps
			},
			restorePaths: append(
				getRestorePaths(t, emailPath, metadata.AllMetadataFileNames()),
				getRestorePaths(t, contactPath, metadata.AllMetadataFileNames())...),
		},
		{
			name:  "single reason sp libraries",
			manID: "single-sp-libraries",
			reasons: []identity.Reasoner{
				identity.NewReason(tid, ro, path.SharePointService, path.LibrariesCategory),
			},
			preFetchPaths: []string{},
			expectPaths: func(t *testing.T, files []string) []path.Path {
				ps := make([]path.Path, 0, len(files))

				for _, f := range files {
					p, err := spLibsPath.AppendItem(f)
					assert.NoError(t, err, clues.ToCore(err))
					ps = append(ps, p)
				}

				return ps
			},
			restorePaths: getRestorePaths(t, spLibsPath, metadata.AllMetadataFileNames()),
		},
		{
			name:  "single reason groups messages",
			manID: "single-groups-messages",
			reasons: []identity.Reasoner{
				identity.NewReason(tid, ro, path.GroupsService, path.ChannelMessagesCategory),
			},
			preFetchPaths: []string{},
			expectPaths: func(t *testing.T, files []string) []path.Path {
				ps := make([]path.Path, 0, len(files))

				for _, f := range files {
					p, err := messagesPath.AppendItem(f)
					assert.NoError(t, err, clues.ToCore(err))
					ps = append(ps, p)
				}

				return ps
			},
			restorePaths: getRestorePaths(t, messagesPath, metadata.AllMetadataFileNames()),
		},
		{
			name:  "single reason groups libraries",
			manID: "single-groups-libraries",
			reasons: []identity.Reasoner{
				identity.NewReason(tid, ro, path.GroupsService, path.LibrariesCategory),
			},
			preFetchPaths: []string{"previouspath"},
			expectPaths: func(t *testing.T, files []string) []path.Path {
				ps := make([]path.Path, 0, len(files))

				assert.NoError(t, err, clues.ToCore(err))
				for _, f := range files {
					p, err := groupLibsSitesPath.AppendItem(f)
					assert.NoError(t, err, clues.ToCore(err))
					ps = append(ps, p)
				}

				return ps
			},
			restorePaths: append(
				getRestorePaths(t, groupLibsSite1Path, metadata.AllMetadataFileNames()),
				getRestorePaths(t, groupLibsSite2Path, metadata.AllMetadataFileNames())...),
			preFetchCollection: []data.RestoreCollection{dataMock.Collection{
				ItemData: []data.Item{
					&dataMock.Item{
						ItemID: "previouspath",
						Reader: io.NopCloser(bytes.NewReader(
							[]byte(`{"site1": "/path/does/not/matter", "site2": "/path/does/not/matter"}`))),
					},
				},
			}},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			paths := test.expectPaths(t, test.preFetchPaths)

			mr := mockRestoreProducer{err: test.expectErr, colls: test.preFetchCollection}
			mr.buildRestoreFunc(t, test.manID, paths)

			base := kopia.BackupBase{
				ItemDataSnapshot: &snapshot.Manifest{ID: manifest.ID(test.manID)},
				Reasons:          test.reasons,
			}

			controller := m365.Controller{}
			pths, err := controller.GetMetadataPaths(ctx, &mr, base, fault.New(true))
			assert.ErrorIs(t, err, test.expectErr, clues.ToCore(err))
			assert.ElementsMatch(t, test.restorePaths, pths, "restore paths")
		})
	}
}

func (suite *OperationsManifestsUnitSuite) TestProduceManifestsAndMetadata() {
	var (
		ro          = "resourceowner"
		tid         = "tenantid"
		emailReason = kopia.NewReason(tid, ro, path.ExchangeService, path.EmailCategory)

		baseBuilder = func(id int) *kopia.BackupBaseBuilder {
			return kopia.NewBackupBaseBuilder("", id).
				WithReasons(emailReason)
		}
		colID = func(id int) string {
			return fmt.Sprintf("ID%d-item-data", id)
		}
	)

	table := []struct {
		name        string
		bf          *mockBackupFinder
		rp          mockRestoreProducer
		reasons     []identity.Reasoner
		getMeta     bool
		dropAssist  bool
		assertErr   assert.ErrorAssertionFunc
		assertB     assert.BoolAssertionFunc
		expectDCS   []mockColl
		expectPaths func(t *testing.T, gotPaths []path.Path)
		expectMans  kopia.BackupBases
	}{
		{
			name:       "don't get metadata, no mans",
			rp:         mockRestoreProducer{},
			reasons:    []identity.Reasoner{},
			getMeta:    false,
			assertErr:  assert.NoError,
			assertB:    assert.False,
			expectDCS:  nil,
			expectMans: kopia.NewMockBackupBases(),
		},
		{
			name: "don't get metadata",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithMergeBases(baseBuilder(1).Build()),
				},
			},
			rp:        mockRestoreProducer{},
			reasons:   []identity.Reasoner{emailReason},
			getMeta:   false,
			assertErr: assert.NoError,
			assertB:   assert.False,
			expectDCS: nil,
			expectMans: kopia.NewMockBackupBases().
				WithMergeBases(baseBuilder(1).Build()).MockDisableMergeBases(),
		},
		{
			name: "don't get metadata, assist base",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithAssistBases(baseBuilder(1).MarkAssistBase().Build()),
				},
			},
			rp:        mockRestoreProducer{},
			reasons:   []identity.Reasoner{emailReason},
			getMeta:   true,
			assertErr: assert.NoError,
			// Doesn't matter if it's true or false as merge/assist bases are
			// distinct. A future PR can go and remove the requirement to pass the
			// flag to kopia and just pass it the bases instead.
			assertB:   assert.True,
			expectDCS: nil,
			expectMans: kopia.NewMockBackupBases().
				WithAssistBases(baseBuilder(1).MarkAssistBase().Build()),
		},
		{
			name: "one valid man, multiple reasons",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithMergeBases(baseBuilder(1).AppendReasons(
							kopia.NewReason(tid, ro, path.ExchangeService, path.EmailCategory)).
							Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID(1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID(1)}}},
				},
			},
			reasons: []identity.Reasoner{
				emailReason,
				identity.NewReason(tid, ro, path.ExchangeService, path.ContactsCategory),
			},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: colID(1)}},
			expectPaths: func(t *testing.T, gotPaths []path.Path) {
				for _, p := range gotPaths {
					assert.Equal(
						t,
						path.ExchangeMetadataService,
						p.Service(),
						"read data service")

					assert.Contains(
						t,
						[]path.CategoryType{
							path.EmailCategory,
							path.ContactsCategory,
						},
						p.Category(),
						"read data category doesn't match a given reason")
				}
			},
			expectMans: kopia.NewMockBackupBases().
				WithMergeBases(baseBuilder(1).AppendReasons(
					kopia.NewReason(tid, ro, path.ExchangeService, path.EmailCategory)).
					Build()),
		},
		{
			name: "one valid man, extra incomplete man",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithMergeBases(baseBuilder(1).Build()).
						WithAssistBases(baseBuilder(2).MarkAssistBase().Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID(1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID(1)}}},
					colID(2): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID(2)}}},
				},
			},
			reasons:   []identity.Reasoner{emailReason},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: colID(1)}},
			expectMans: kopia.NewMockBackupBases().
				WithMergeBases(baseBuilder(1).Build()).
				WithAssistBases(baseBuilder(2).MarkAssistBase().Build()),
		},
		{
			name: "one valid man, extra incomplete man, drop assist bases",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithMergeBases(baseBuilder(1).Build()).
						WithAssistBases(baseBuilder(2).MarkAssistBase().Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID(1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID(1)}}},
					colID(2): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID(2)}}},
				},
			},
			reasons:    []identity.Reasoner{emailReason},
			getMeta:    true,
			dropAssist: true,
			assertErr:  assert.NoError,
			assertB:    assert.True,
			expectDCS:  []mockColl{{id: colID(1)}},
			expectMans: kopia.NewMockBackupBases().
				WithMergeBases(baseBuilder(1).Build()).
				MockDisableAssistBases(),
		},
		{
			name: "multiple valid mans",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithMergeBases(
						baseBuilder(1).Build(),
						baseBuilder(2).
							WithReasons(
								kopia.NewReason(tid, ro, path.ExchangeService, path.EventsCategory)).
							Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID(1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID(1)}}},
					colID(2): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID(2)}}},
				},
			},
			reasons: []identity.Reasoner{
				emailReason,
				identity.NewReason(tid, ro, path.ExchangeService, path.EventsCategory),
			},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: colID(1)}, {id: colID(2)}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				baseBuilder(1).Build(),
				baseBuilder(2).
					WithReasons(
						identity.NewReason(tid, ro, path.ExchangeService, path.EventsCategory)).
					Build()),
		},
		{
			name: "error collecting metadata",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithMergeBases(baseBuilder(1).Build()),
				},
			},
			rp:         mockRestoreProducer{err: assert.AnError},
			reasons:    []identity.Reasoner{emailReason},
			getMeta:    true,
			assertErr:  assert.Error,
			assertB:    assert.False,
			expectDCS:  nil,
			expectMans: nil,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			emptyMockBackpuProducer := mock.NewMockBackupProducer(nil, data.CollectionStats{}, false)
			mans, dcs, b, err := produceManifestsAndMetadata(
				ctx,
				test.bf,
				&emptyMockBackpuProducer,
				&test.rp,
				test.reasons, nil,
				tid,
				test.getMeta,
				test.dropAssist)
			test.assertErr(t, err, clues.ToCore(err))
			test.assertB(t, b)

			kopia.AssertBackupBasesEqual(t, test.expectMans, mans)

			expect, got := []string{}, []string{}

			for _, dc := range test.expectDCS {
				expect = append(expect, dc.id)
			}

			for _, dc := range dcs {
				if !assert.IsTypef(
					t,
					data.NoFetchRestoreCollection{},
					dc,
					"unexpected type returned [%T]",
					dc) {
					continue
				}

				tmp := dc.(data.NoFetchRestoreCollection)

				if !assert.IsTypef(
					t,
					mockColl{},
					tmp.Collection,
					"unexpected type returned [%T]",
					tmp.Collection) {
					continue
				}

				mc := tmp.Collection.(mockColl)
				got = append(got, mc.id)
			}

			assert.ElementsMatch(t, expect, got, "expected collections are present")

			if test.expectPaths != nil {
				test.expectPaths(t, test.rp.gotPaths)
			}
		})
	}
}

func (suite *OperationsManifestsUnitSuite) TestProduceManifestsAndMetadata_FallbackReasons() {
	var (
		ro   = "resourceowner"
		fbro = "fb_resourceowner"
		tid  = "tenantid"

		emailReason = kopia.NewReason(
			tid,
			ro,
			path.ExchangeService,
			path.EmailCategory)

		fbEmailReason = kopia.NewReason(
			tid,
			fbro,
			path.ExchangeService,
			path.EmailCategory)

		baseBuilder = func(id int) *kopia.BackupBaseBuilder {
			return kopia.NewBackupBaseBuilder("", id).
				WithReasons(emailReason)
		}
		fbBaseBuilder = func(id int) *kopia.BackupBaseBuilder {
			return kopia.NewBackupBaseBuilder("fb", id).
				WithReasons(fbEmailReason)
		}
		colID = func(prefix string, id int) string {
			return fmt.Sprintf("%sID%d-item-data", prefix, id)
		}
	)

	table := []struct {
		name            string
		bf              *mockBackupFinder
		rp              mockRestoreProducer
		reasons         []identity.Reasoner
		fallbackReasons []identity.Reasoner
		getMeta         bool
		dropAssist      bool
		assertErr       assert.ErrorAssertionFunc
		assertB         assert.BoolAssertionFunc
		expectDCS       []mockColl
		expectMans      kopia.BackupBases
	}{
		{
			name: "don't get metadata, only fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					fbro: kopia.NewMockBackupBases().
						WithMergeBases(fbBaseBuilder(1).Build()),
				},
			},
			rp:              mockRestoreProducer{},
			fallbackReasons: []identity.Reasoner{fbEmailReason},
			getMeta:         false,
			assertErr:       assert.NoError,
			assertB:         assert.False,
			expectDCS:       nil,
			expectMans: kopia.NewMockBackupBases().
				WithMergeBases(fbBaseBuilder(1).Build()).
				MockDisableMergeBases(),
		},
		{
			name: "only fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					fbro: kopia.NewMockBackupBases().
						WithMergeBases(fbBaseBuilder(1).Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID("fb", 1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("fb", 1)}}},
				},
			},
			fallbackReasons: []identity.Reasoner{fbEmailReason},
			getMeta:         true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       []mockColl{{id: colID("fb", 1)}},
			expectMans: kopia.NewMockBackupBases().
				WithMergeBases(fbBaseBuilder(1).Build()),
		},
		{
			name: "only fallbacks, drop assist",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					fbro: kopia.NewMockBackupBases().
						WithMergeBases(fbBaseBuilder(1).Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID("fb", 1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("fb", 1)}}},
				},
			},
			fallbackReasons: []identity.Reasoner{fbEmailReason},
			getMeta:         true,
			dropAssist:      true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       []mockColl{{id: colID("fb", 1)}},
			expectMans: kopia.NewMockBackupBases().
				WithMergeBases(fbBaseBuilder(1).Build()).
				MockDisableAssistBases(),
		},
		{
			name: "complete mans and fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithMergeBases(baseBuilder(1).Build()),
					fbro: kopia.NewMockBackupBases().
						WithMergeBases(fbBaseBuilder(1).Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID("", 1):   {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("", 1)}}},
					colID("fb", 1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("fb", 1)}}},
				},
			},
			reasons:         []identity.Reasoner{emailReason},
			fallbackReasons: []identity.Reasoner{fbEmailReason},
			getMeta:         true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       []mockColl{{id: colID("", 1)}},
			expectMans: kopia.NewMockBackupBases().
				WithMergeBases(baseBuilder(1).Build()),
		},
		{
			name: "incomplete mans and fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithAssistBases(
							baseBuilder(2).MarkAssistBase().Build()),
					fbro: kopia.NewMockBackupBases().
						WithAssistBases(fbBaseBuilder(2).MarkAssistBase().Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID("", 2):   {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("", 2)}}},
					colID("fb", 2): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("fb", 2)}}},
				},
			},
			reasons:         []identity.Reasoner{emailReason},
			fallbackReasons: []identity.Reasoner{fbEmailReason},
			getMeta:         true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       nil,
			expectMans: kopia.NewMockBackupBases().
				WithAssistBases(baseBuilder(2).MarkAssistBase().Build()),
		},
		{
			name: "complete and incomplete mans and fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithMergeBases(baseBuilder(1).Build()).
						WithAssistBases(baseBuilder(2).MarkAssistBase().Build()),
					fbro: kopia.NewMockBackupBases().
						WithMergeBases(fbBaseBuilder(1).Build()).
						WithAssistBases(fbBaseBuilder(2).MarkAssistBase().Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID("", 1):   {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("", 1)}}},
					colID("", 2):   {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("", 2)}}},
					colID("fb", 1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("fb", 1)}}},
					colID("fb", 2): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("fb", 2)}}},
				},
			},
			reasons:         []identity.Reasoner{emailReason},
			fallbackReasons: []identity.Reasoner{fbEmailReason},
			getMeta:         true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       []mockColl{{id: colID("", 1)}},
			expectMans: kopia.NewMockBackupBases().
				WithMergeBases(baseBuilder(1).Build()).
				WithAssistBases(baseBuilder(2).MarkAssistBase().Build()),
		},
		{
			name: "incomplete mans and complete fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithAssistBases(
							baseBuilder(2).MarkAssistBase().Build()),
					fbro: kopia.NewMockBackupBases().
						WithMergeBases(fbBaseBuilder(1).Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID("", 2):   {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("", 2)}}},
					colID("fb", 1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("fb", 1)}}},
				},
			},
			reasons:         []identity.Reasoner{emailReason},
			fallbackReasons: []identity.Reasoner{fbEmailReason},
			getMeta:         true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       []mockColl{{id: colID("fb", 1)}},
			expectMans: kopia.NewMockBackupBases().
				WithMergeBases(fbBaseBuilder(1).Build()).
				WithAssistBases(baseBuilder(2).MarkAssistBase().Build()),
		},
		{
			name: "incomplete mans and complete fallbacks, no assist bases",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithAssistBases(baseBuilder(2).MarkAssistBase().Build()),
					fbro: kopia.NewMockBackupBases().
						WithMergeBases(fbBaseBuilder(1).Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID("", 2):   {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("", 2)}}},
					colID("fb", 1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("fb", 1)}}},
				},
			},
			reasons:         []identity.Reasoner{emailReason},
			fallbackReasons: []identity.Reasoner{fbEmailReason},
			getMeta:         true,
			dropAssist:      true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       []mockColl{{id: colID("fb", 1)}},
			expectMans: kopia.NewMockBackupBases().
				WithMergeBases(fbBaseBuilder(1).Build()).
				MockDisableAssistBases(),
		},
		{
			name: "complete mans and incomplete fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithMergeBases(baseBuilder(1).Build()),
					fbro: kopia.NewMockBackupBases().
						WithAssistBases(fbBaseBuilder(2).MarkAssistBase().Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID("", 1):   {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("", 1)}}},
					colID("fb", 2): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("fb", 2)}}},
				},
			},
			reasons:         []identity.Reasoner{emailReason},
			fallbackReasons: []identity.Reasoner{fbEmailReason},
			getMeta:         true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       []mockColl{{id: colID("", 1)}},
			expectMans: kopia.NewMockBackupBases().
				WithMergeBases(baseBuilder(1).Build()),
		},
		{
			name: "complete mans and complete fallbacks, multiple reasons",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithMergeBases(
						baseBuilder(1).
							AppendReasons(kopia.NewReason(tid, ro, path.ExchangeService, path.ContactsCategory)).
							Build()),
					fbro: kopia.NewMockBackupBases().WithMergeBases(
						fbBaseBuilder(1).
							AppendReasons(kopia.NewReason(tid, fbro, path.ExchangeService, path.ContactsCategory)).
							Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID("", 1):   {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("", 1)}}},
					colID("fb", 1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("fb", 1)}}},
				},
			},
			reasons: []identity.Reasoner{
				emailReason,
				identity.NewReason(tid, ro, path.ExchangeService, path.ContactsCategory),
			},
			fallbackReasons: []identity.Reasoner{
				fbEmailReason,
				identity.NewReason(tid, fbro, path.ExchangeService, path.ContactsCategory),
			},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: colID("", 1)}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				baseBuilder(1).
					AppendReasons(kopia.NewReason(tid, ro, path.ExchangeService, path.ContactsCategory)).
					Build()),
		},
		{
			name: "complete mans and complete fallbacks, distinct reasons",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithMergeBases(baseBuilder(1).Build()),
					fbro: kopia.NewMockBackupBases().WithMergeBases(
						fbBaseBuilder(1).
							WithReasons(kopia.NewReason(tid, fbro, path.ExchangeService, path.ContactsCategory)).
							Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID("", 1):   {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("", 1)}}},
					colID("fb", 1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("fb", 1)}}},
				},
			},
			reasons: []identity.Reasoner{emailReason},
			fallbackReasons: []identity.Reasoner{
				identity.NewReason(tid, fbro, path.ExchangeService, path.ContactsCategory),
			},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: colID("", 1)}, {id: colID("fb", 1)}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				baseBuilder(1).Build(),
				fbBaseBuilder(1).
					WithReasons(kopia.NewReason(tid, fbro, path.ExchangeService, path.ContactsCategory)).
					Build()),
		},
		{
			name: "complete mans and complete fallbacks, fallback has superset of reasons",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().
						WithMergeBases(baseBuilder(1).Build()),
					fbro: kopia.NewMockBackupBases().WithMergeBases(
						fbBaseBuilder(1).
							AppendReasons(kopia.NewReason(tid, fbro, path.ExchangeService, path.ContactsCategory)).
							Build()),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					colID("", 1):   {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("", 1)}}},
					colID("fb", 1): {data.NoFetchRestoreCollection{Collection: mockColl{id: colID("fb", 1)}}},
				},
			},
			reasons: []identity.Reasoner{
				emailReason,
				identity.NewReason(tid, ro, path.ExchangeService, path.ContactsCategory),
			},
			fallbackReasons: []identity.Reasoner{
				fbEmailReason,
				identity.NewReason(tid, fbro, path.ExchangeService, path.ContactsCategory),
			},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: colID("", 1)}, {id: colID("fb", 1)}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				baseBuilder(1).Build(),
				fbBaseBuilder(1).
					WithReasons(kopia.NewReason(tid, fbro, path.ExchangeService, path.ContactsCategory)).
					Build()),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mbp := mock.NewMockBackupProducer(nil, data.CollectionStats{}, false)
			mans, dcs, b, err := produceManifestsAndMetadata(
				ctx,
				test.bf,
				&mbp,
				&test.rp,
				test.reasons, test.fallbackReasons,
				tid,
				test.getMeta,
				test.dropAssist)
			test.assertErr(t, err, clues.ToCore(err))
			test.assertB(t, b)

			kopia.AssertBackupBasesEqual(t, test.expectMans, mans)

			expect, got := []string{}, []string{}

			for _, dc := range test.expectDCS {
				expect = append(expect, dc.id)
			}

			for _, dc := range dcs {
				if !assert.IsTypef(
					t,
					data.NoFetchRestoreCollection{},
					dc,
					"unexpected type returned [%T]",
					dc) {
					continue
				}

				tmp := dc.(data.NoFetchRestoreCollection)

				if !assert.IsTypef(
					t,
					mockColl{},
					tmp.Collection,
					"unexpected type returned [%T]",
					tmp.Collection) {
					continue
				}

				mc := tmp.Collection.(mockColl)
				got = append(got, mc.id)
			}

			assert.ElementsMatch(t, expect, got, "expected collections are present")
		})
	}
}
