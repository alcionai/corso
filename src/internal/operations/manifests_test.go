package operations

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup"
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

func (mc mockColl) Items(context.Context, *fault.Bus) <-chan data.Stream {
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
	reasons []kopia.Reason,
	_ map[string]string,
) kopia.BackupBases {
	if len(reasons) == 0 {
		return kopia.NewMockBackupBases()
	}

	if bf == nil {
		return kopia.NewMockBackupBases()
	}

	b := bf.data[reasons[0].ResourceOwner]
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

func (suite *OperationsManifestsUnitSuite) TestCollectMetadata() {
	const (
		ro  = "owner"
		tid = "tenantid"
	)

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
	)

	table := []struct {
		name        string
		manID       string
		reasons     []kopia.Reason
		fileNames   []string
		expectPaths func(*testing.T, []string) []path.Path
		expectErr   error
	}{
		{
			name:  "single reason, single file",
			manID: "single single",
			reasons: []kopia.Reason{
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.EmailCategory,
				},
			},
			expectPaths: func(t *testing.T, files []string) []path.Path {
				ps := make([]path.Path, 0, len(files))

				for _, f := range files {
					p, err := emailPath.AppendItem(f)
					assert.NoError(t, err, clues.ToCore(err))
					ps = append(ps, p)
				}

				return ps
			},
			fileNames: []string{"a"},
		},
		{
			name:  "single reason, multiple files",
			manID: "single multi",
			reasons: []kopia.Reason{
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.EmailCategory,
				},
			},
			expectPaths: func(t *testing.T, files []string) []path.Path {
				ps := make([]path.Path, 0, len(files))

				for _, f := range files {
					p, err := emailPath.AppendItem(f)
					assert.NoError(t, err, clues.ToCore(err))
					ps = append(ps, p)
				}

				return ps
			},
			fileNames: []string{"a", "b"},
		},
		{
			name:  "multiple reasons, single file",
			manID: "multi single",
			reasons: []kopia.Reason{
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.EmailCategory,
				},
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.ContactsCategory,
				},
			},
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
			fileNames: []string{"a"},
		},
		{
			name:  "multiple reasons, multiple file",
			manID: "multi multi",
			reasons: []kopia.Reason{
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.EmailCategory,
				},
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.ContactsCategory,
				},
			},
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
			fileNames: []string{"a", "b"},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			paths := test.expectPaths(t, test.fileNames)

			mr := mockRestoreProducer{err: test.expectErr}
			mr.buildRestoreFunc(t, test.manID, paths)

			man := kopia.ManifestEntry{
				Manifest: &snapshot.Manifest{ID: manifest.ID(test.manID)},
				Reasons:  test.reasons,
			}

			_, err := collectMetadata(ctx, &mr, man, test.fileNames, tid, fault.New(true))
			assert.ErrorIs(t, err, test.expectErr, clues.ToCore(err))
		})
	}
}

func buildReasons(
	ro string,
	service path.ServiceType,
	cats ...path.CategoryType,
) []kopia.Reason {
	var reasons []kopia.Reason

	for _, cat := range cats {
		reasons = append(
			reasons,
			kopia.Reason{
				ResourceOwner: ro,
				Service:       service,
				Category:      cat,
			})
	}

	return reasons
}

func (suite *OperationsManifestsUnitSuite) TestProduceManifestsAndMetadata() {
	const (
		ro  = "resourceowner"
		tid = "tenantid"
		did = "detailsid"
	)

	makeMan := func(id, incmpl string, cats ...path.CategoryType) kopia.ManifestEntry {
		return kopia.ManifestEntry{
			Manifest: &snapshot.Manifest{
				ID:               manifest.ID(id),
				IncompleteReason: incmpl,
			},
			Reasons: buildReasons(ro, path.ExchangeService, cats...),
		}
	}

	table := []struct {
		name        string
		bf          *mockBackupFinder
		rp          mockRestoreProducer
		reasons     []kopia.Reason
		getMeta     bool
		assertErr   assert.ErrorAssertionFunc
		assertB     assert.BoolAssertionFunc
		expectDCS   []mockColl
		expectPaths func(t *testing.T, gotPaths []path.Path)
		expectMans  kopia.BackupBases
	}{
		{
			name:       "don't get metadata, no mans",
			rp:         mockRestoreProducer{},
			reasons:    []kopia.Reason{},
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
					ro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan("id1", "", path.EmailCategory),
					),
				},
			},
			rp: mockRestoreProducer{},
			reasons: []kopia.Reason{
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.EmailCategory,
				},
			},
			getMeta:   false,
			assertErr: assert.NoError,
			assertB:   assert.False,
			expectDCS: nil,
			expectMans: kopia.NewMockBackupBases().WithAssistBases(
				makeMan("id1", "", path.EmailCategory),
			),
		},
		{
			name: "don't get metadata, incomplete manifest",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithAssistBases(
						makeMan("id1", "checkpoint", path.EmailCategory),
					),
				},
			},
			rp: mockRestoreProducer{},
			reasons: []kopia.Reason{
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.EmailCategory,
				},
			},
			getMeta:   true,
			assertErr: assert.NoError,
			// Doesn't matter if it's true or false as merge/assist bases are
			// distinct. A future PR can go and remove the requirement to pass the
			// flag to kopia and just pass it the bases instead.
			assertB:   assert.True,
			expectDCS: nil,
			expectMans: kopia.NewMockBackupBases().WithAssistBases(
				makeMan("id1", "checkpoint", path.EmailCategory),
			),
		},
		{
			name: "one valid man, multiple reasons",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan("id1", "", path.EmailCategory, path.ContactsCategory),
					),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					"id1": {data.NotFoundRestoreCollection{Collection: mockColl{id: "id1"}}},
				},
			},
			reasons: []kopia.Reason{
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.EmailCategory,
				},
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.ContactsCategory,
				},
			},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: "id1"}},
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
						"read data category doesn't match a given reason",
					)
				}
			},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				makeMan("id1", "", path.EmailCategory, path.ContactsCategory),
			),
		},
		{
			name: "one valid man, extra incomplete man",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan("id1", "", path.EmailCategory),
					).WithAssistBases(
						makeMan("id2", "checkpoint", path.EmailCategory),
					),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					"id1": {data.NotFoundRestoreCollection{Collection: mockColl{id: "id1"}}},
					"id2": {data.NotFoundRestoreCollection{Collection: mockColl{id: "id2"}}},
				},
			},
			reasons: []kopia.Reason{
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.EmailCategory,
				},
			},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: "id1"}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				makeMan("id1", "", path.EmailCategory),
			).WithAssistBases(
				makeMan("id2", "checkpoint", path.EmailCategory),
			),
		},
		{
			name: "multiple valid mans",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan("id1", "", path.EmailCategory),
						makeMan("id2", "", path.EmailCategory),
					),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					"id1": {data.NotFoundRestoreCollection{Collection: mockColl{id: "id1"}}},
					"id2": {data.NotFoundRestoreCollection{Collection: mockColl{id: "id2"}}},
				},
			},
			reasons: []kopia.Reason{
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.EmailCategory,
				},
			},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: "id1"}, {id: "id2"}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				makeMan("id1", "", path.EmailCategory),
				makeMan("id2", "", path.EmailCategory),
			),
		},
		{
			name: "error collecting metadata",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan("id1", "", path.EmailCategory),
					),
				},
			},
			rp: mockRestoreProducer{err: assert.AnError},
			reasons: []kopia.Reason{
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.EmailCategory,
				},
			},
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

			mans, dcs, b, err := produceManifestsAndMetadata(
				ctx,
				test.bf,
				&test.rp,
				test.reasons, nil,
				tid,
				test.getMeta)
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
					data.NotFoundRestoreCollection{},
					dc,
					"unexpected type returned [%T]",
					dc,
				) {
					continue
				}

				tmp := dc.(data.NotFoundRestoreCollection)

				if !assert.IsTypef(
					t,
					mockColl{},
					tmp.Collection,
					"unexpected type returned [%T]",
					tmp.Collection,
				) {
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
	const (
		ro   = "resourceowner"
		fbro = "fb_resourceowner"
		tid  = "tenantid"
		did  = "detailsid"
	)

	makeMan := func(ro, id, incmpl string, cats ...path.CategoryType) kopia.ManifestEntry {
		return kopia.ManifestEntry{
			Manifest: &snapshot.Manifest{
				ID:               manifest.ID(id),
				IncompleteReason: incmpl,
				Tags:             map[string]string{"tag:" + kopia.TagBackupID: id + "bup"},
			},
			Reasons: buildReasons(ro, path.ExchangeService, cats...),
		}
	}

	makeBackup := func(ro, snapID string, cats ...path.CategoryType) kopia.BackupEntry {
		return kopia.BackupEntry{
			Backup: &backup.Backup{
				BaseModel: model.BaseModel{
					ID: model.StableID(snapID + "bup"),
				},
				SnapshotID:    snapID,
				StreamStoreID: snapID + "store",
			},
			Reasons: buildReasons(ro, path.ExchangeService, cats...),
		}
	}

	emailReason := kopia.Reason{
		ResourceOwner: ro,
		Service:       path.ExchangeService,
		Category:      path.EmailCategory,
	}

	fbEmailReason := kopia.Reason{
		ResourceOwner: fbro,
		Service:       path.ExchangeService,
		Category:      path.EmailCategory,
	}

	table := []struct {
		name            string
		bf              *mockBackupFinder
		rp              mockRestoreProducer
		reasons         []kopia.Reason
		fallbackReasons []kopia.Reason
		getMeta         bool
		assertErr       assert.ErrorAssertionFunc
		assertB         assert.BoolAssertionFunc
		expectDCS       []mockColl
		expectMans      kopia.BackupBases
	}{
		{
			name: "don't get metadata, only fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					fbro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(fbro, "fb_id1", "", path.EmailCategory),
					).WithBackups(
						makeBackup(fbro, "fb_id1", path.EmailCategory),
					),
				},
			},
			rp:              mockRestoreProducer{},
			fallbackReasons: []kopia.Reason{fbEmailReason},
			getMeta:         false,
			assertErr:       assert.NoError,
			assertB:         assert.False,
			expectDCS:       nil,
			expectMans: kopia.NewMockBackupBases().WithAssistBases(
				makeMan(fbro, "fb_id1", "", path.EmailCategory),
			),
		},
		{
			name: "only fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					fbro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(fbro, "fb_id1", "", path.EmailCategory),
					).WithBackups(
						makeBackup(fbro, "fb_id1", path.EmailCategory),
					),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					"fb_id1": {data.NotFoundRestoreCollection{Collection: mockColl{id: "fb_id1"}}},
				},
			},
			fallbackReasons: []kopia.Reason{fbEmailReason},
			getMeta:         true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       []mockColl{{id: "fb_id1"}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				makeMan(fbro, "fb_id1", "", path.EmailCategory),
			).WithBackups(
				makeBackup(fbro, "fb_id1", path.EmailCategory),
			),
		},
		{
			name: "complete mans and fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(ro, "id1", "", path.EmailCategory),
					),
					fbro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(fbro, "fb_id1", "", path.EmailCategory),
					).WithBackups(
						makeBackup(fbro, "fb_id1", path.EmailCategory),
					),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					"id1":    {data.NotFoundRestoreCollection{Collection: mockColl{id: "id1"}}},
					"fb_id1": {data.NotFoundRestoreCollection{Collection: mockColl{id: "fb_id1"}}},
				},
			},
			reasons:         []kopia.Reason{emailReason},
			fallbackReasons: []kopia.Reason{fbEmailReason},
			getMeta:         true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       []mockColl{{id: "id1"}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				makeMan(ro, "id1", "", path.EmailCategory),
			),
		},
		{
			name: "incomplete mans and fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithAssistBases(
						makeMan(ro, "id2", "checkpoint", path.EmailCategory),
					),
					fbro: kopia.NewMockBackupBases().WithAssistBases(
						makeMan(fbro, "fb_id2", "checkpoint", path.EmailCategory),
					),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					"id2":    {data.NotFoundRestoreCollection{Collection: mockColl{id: "id2"}}},
					"fb_id2": {data.NotFoundRestoreCollection{Collection: mockColl{id: "fb_id2"}}},
				},
			},
			reasons:         []kopia.Reason{emailReason},
			fallbackReasons: []kopia.Reason{fbEmailReason},
			getMeta:         true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       nil,
			expectMans: kopia.NewMockBackupBases().WithAssistBases(
				makeMan(ro, "id2", "checkpoint", path.EmailCategory),
			),
		},
		{
			name: "complete and incomplete mans and fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(ro, "id1", "", path.EmailCategory),
					).WithAssistBases(
						makeMan(ro, "id2", "checkpoint", path.EmailCategory),
					),
					fbro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(fbro, "fb_id1", "", path.EmailCategory),
					).WithBackups(
						makeBackup(fbro, "fb_id1", path.EmailCategory),
					).WithAssistBases(
						makeMan(fbro, "fb_id2", "checkpoint", path.EmailCategory),
					),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					"id1":    {data.NotFoundRestoreCollection{Collection: mockColl{id: "id1"}}},
					"id2":    {data.NotFoundRestoreCollection{Collection: mockColl{id: "id2"}}},
					"fb_id1": {data.NotFoundRestoreCollection{Collection: mockColl{id: "fb_id1"}}},
					"fb_id2": {data.NotFoundRestoreCollection{Collection: mockColl{id: "fb_id2"}}},
				},
			},
			reasons:         []kopia.Reason{emailReason},
			fallbackReasons: []kopia.Reason{fbEmailReason},
			getMeta:         true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       []mockColl{{id: "id1"}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				makeMan(ro, "id1", "", path.EmailCategory),
			).WithAssistBases(
				makeMan(ro, "id2", "checkpoint", path.EmailCategory),
			),
		},
		{
			name: "incomplete mans and complete fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithAssistBases(
						makeMan(ro, "id2", "checkpoint", path.EmailCategory),
					),
					fbro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(fbro, "fb_id1", "", path.EmailCategory),
					).WithBackups(
						makeBackup(fbro, "fb_id1", path.EmailCategory),
					),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					"id2":    {data.NotFoundRestoreCollection{Collection: mockColl{id: "id2"}}},
					"fb_id1": {data.NotFoundRestoreCollection{Collection: mockColl{id: "fb_id1"}}},
				},
			},
			reasons:         []kopia.Reason{emailReason},
			fallbackReasons: []kopia.Reason{fbEmailReason},
			getMeta:         true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       []mockColl{{id: "fb_id1"}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				makeMan(fbro, "fb_id1", "", path.EmailCategory),
			).WithBackups(
				makeBackup(fbro, "fb_id1", path.EmailCategory),
			).WithAssistBases(
				makeMan(ro, "id2", "checkpoint", path.EmailCategory),
			),
		},
		{
			name: "complete mans and incomplete fallbacks",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(ro, "id1", "", path.EmailCategory),
					),
					fbro: kopia.NewMockBackupBases().WithAssistBases(
						makeMan(fbro, "fb_id2", "checkpoint", path.EmailCategory),
					),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					"id1":    {data.NotFoundRestoreCollection{Collection: mockColl{id: "id1"}}},
					"fb_id2": {data.NotFoundRestoreCollection{Collection: mockColl{id: "fb_id2"}}},
				},
			},
			reasons:         []kopia.Reason{emailReason},
			fallbackReasons: []kopia.Reason{fbEmailReason},
			getMeta:         true,
			assertErr:       assert.NoError,
			assertB:         assert.True,
			expectDCS:       []mockColl{{id: "id1"}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				makeMan(ro, "id1", "", path.EmailCategory),
			),
		},
		{
			name: "complete mans and complete fallbacks, multiple reasons",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(ro, "id1", "", path.EmailCategory, path.ContactsCategory),
					),
					fbro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(fbro, "fb_id1", "", path.EmailCategory, path.ContactsCategory),
					).WithBackups(
						makeBackup(fbro, "fb_id1", path.EmailCategory, path.ContactsCategory),
					),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					"id1":    {data.NotFoundRestoreCollection{Collection: mockColl{id: "id1"}}},
					"fb_id1": {data.NotFoundRestoreCollection{Collection: mockColl{id: "fb_id1"}}},
				},
			},
			reasons: []kopia.Reason{
				emailReason,
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.ContactsCategory,
				},
			},
			fallbackReasons: []kopia.Reason{
				fbEmailReason,
				{
					ResourceOwner: fbro,
					Service:       path.ExchangeService,
					Category:      path.ContactsCategory,
				},
			},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: "id1"}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				makeMan(ro, "id1", "", path.EmailCategory, path.ContactsCategory),
			),
		},
		{
			name: "complete mans and complete fallbacks, distinct reasons",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(ro, "id1", "", path.EmailCategory),
					),
					fbro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(fbro, "fb_id1", "", path.ContactsCategory),
					).WithBackups(
						makeBackup(fbro, "fb_id1", path.ContactsCategory),
					),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					"id1":    {data.NotFoundRestoreCollection{Collection: mockColl{id: "id1"}}},
					"fb_id1": {data.NotFoundRestoreCollection{Collection: mockColl{id: "fb_id1"}}},
				},
			},
			reasons: []kopia.Reason{emailReason},
			fallbackReasons: []kopia.Reason{
				{
					ResourceOwner: fbro,
					Service:       path.ExchangeService,
					Category:      path.ContactsCategory,
				},
			},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: "id1"}, {id: "fb_id1"}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				makeMan(ro, "id1", "", path.EmailCategory),
				makeMan(fbro, "fb_id1", "", path.ContactsCategory),
			).WithBackups(
				makeBackup(fbro, "fb_id1", path.ContactsCategory),
			),
		},
		{
			name: "complete mans and complete fallbacks, fallback has superset of reasons",
			bf: &mockBackupFinder{
				data: map[string]kopia.BackupBases{
					ro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(ro, "id1", "", path.EmailCategory),
					),
					fbro: kopia.NewMockBackupBases().WithMergeBases(
						makeMan(fbro, "fb_id1", "", path.EmailCategory, path.ContactsCategory),
					).WithBackups(
						makeBackup(fbro, "fb_id1", path.EmailCategory, path.ContactsCategory),
					),
				},
			},
			rp: mockRestoreProducer{
				collsByID: map[string][]data.RestoreCollection{
					"id1":    {data.NotFoundRestoreCollection{Collection: mockColl{id: "id1"}}},
					"fb_id1": {data.NotFoundRestoreCollection{Collection: mockColl{id: "fb_id1"}}},
				},
			},
			reasons: []kopia.Reason{
				emailReason,
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      path.ContactsCategory,
				},
			},
			fallbackReasons: []kopia.Reason{
				fbEmailReason,
				{
					ResourceOwner: fbro,
					Service:       path.ExchangeService,
					Category:      path.ContactsCategory,
				},
			},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: "id1"}, {id: "fb_id1"}},
			expectMans: kopia.NewMockBackupBases().WithMergeBases(
				makeMan(ro, "id1", "", path.EmailCategory),
				makeMan(fbro, "fb_id1", "", path.ContactsCategory),
			).WithBackups(
				makeBackup(fbro, "fb_id1", path.ContactsCategory),
			),
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			mans, dcs, b, err := produceManifestsAndMetadata(
				ctx,
				test.bf,
				&test.rp,
				test.reasons, test.fallbackReasons,
				tid,
				test.getMeta)
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
					dc,
				) {
					continue
				}

				tmp := dc.(data.NoFetchRestoreCollection)

				if !assert.IsTypef(
					t,
					mockColl{},
					tmp.Collection,
					"unexpected type returned [%T]",
					tmp.Collection,
				) {
					continue
				}

				mc := tmp.Collection.(mockColl)
				got = append(got, mc.id)
			}

			assert.ElementsMatch(t, expect, got, "expected collections are present")
		})
	}
}
