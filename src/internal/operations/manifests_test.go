package operations

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

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

type mockManifestRestorer struct {
	mockRestoreProducer
	mans  []*kopia.ManifestEntry
	mrErr error // err varname already claimed by mockRestoreProducer
}

func (mmr mockManifestRestorer) FetchPrevSnapshotManifests(
	ctx context.Context,
	reasons []kopia.Reason,
	tags map[string]string,
) ([]*kopia.ManifestEntry, error) {
	mans := map[string]*kopia.ManifestEntry{}

	for _, r := range reasons {
		for _, m := range mmr.mans {
			for _, mr := range m.Reasons {
				if mr.ResourceOwner == r.ResourceOwner {
					mans[string(m.ID)] = m
					break
				}
			}
		}
	}

	if len(mans) == 0 && len(reasons) == 0 {
		return mmr.mans, mmr.mrErr
	}

	return maps.Values(mans), mmr.mrErr
}

type mockGetBackuper struct {
	detailsID     string
	streamstoreID string
	err           error
}

func (mg mockGetBackuper) GetBackup(
	ctx context.Context,
	backupID model.StableID,
) (*backup.Backup, error) {
	return &backup.Backup{
		DetailsID:     mg.detailsID,
		StreamStoreID: mg.streamstoreID,
	}, mg.err
}

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

			man := &kopia.ManifestEntry{
				Manifest: &snapshot.Manifest{ID: manifest.ID(test.manID)},
				Reasons:  test.reasons,
			}

			_, err := collectMetadata(ctx, &mr, man, test.fileNames, tid, fault.New(true))
			assert.ErrorIs(t, err, test.expectErr, clues.ToCore(err))
		})
	}
}

func (suite *OperationsManifestsUnitSuite) TestVerifyDistinctBases() {
	ro := "resource_owner"

	table := []struct {
		name   string
		mans   []*kopia.ManifestEntry
		expect assert.ErrorAssertionFunc
	}{
		{
			name: "one manifest, one reason",
			mans: []*kopia.ManifestEntry{
				{
					Manifest: &snapshot.Manifest{},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: ro,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
					},
				},
			},
			expect: assert.NoError,
		},
		{
			name: "one incomplete manifest",
			mans: []*kopia.ManifestEntry{
				{
					Manifest: &snapshot.Manifest{IncompleteReason: "ir"},
				},
			},
			expect: assert.NoError,
		},
		{
			name: "one manifest, multiple reasons",
			mans: []*kopia.ManifestEntry{
				{
					Manifest: &snapshot.Manifest{},
					Reasons: []kopia.Reason{
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
				},
			},
			expect: assert.NoError,
		},
		{
			name: "one manifest, duplicate reasons",
			mans: []*kopia.ManifestEntry{
				{
					Manifest: &snapshot.Manifest{},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: ro,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
						{
							ResourceOwner: ro,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
					},
				},
			},
			expect: assert.Error,
		},
		{
			name: "two manifests, non-overlapping reasons",
			mans: []*kopia.ManifestEntry{
				{
					Manifest: &snapshot.Manifest{},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: ro,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
					},
				},
				{
					Manifest: &snapshot.Manifest{},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: ro,
							Service:       path.ExchangeService,
							Category:      path.ContactsCategory,
						},
					},
				},
			},
			expect: assert.NoError,
		},
		{
			name: "two manifests, overlapping reasons",
			mans: []*kopia.ManifestEntry{
				{
					Manifest: &snapshot.Manifest{},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: ro,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
					},
				},
				{
					Manifest: &snapshot.Manifest{},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: ro,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
					},
				},
			},
			expect: assert.Error,
		},
		{
			name: "two manifests, overlapping reasons, one snapshot incomplete",
			mans: []*kopia.ManifestEntry{
				{
					Manifest: &snapshot.Manifest{},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: ro,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
					},
				},
				{
					Manifest: &snapshot.Manifest{IncompleteReason: "ir"},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: ro,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
					},
				},
			},
			expect: assert.NoError,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext(suite.T())
			defer flush()

			err := verifyDistinctBases(ctx, test.mans)
			test.expect(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *OperationsManifestsUnitSuite) TestProduceManifestsAndMetadata() {
	const (
		ro  = "resourceowner"
		tid = "tenantid"
		did = "detailsid"
	)

	makeMan := func(pct path.CategoryType, id, incmpl, bid string) *kopia.ManifestEntry {
		tags := map[string]string{}
		if len(bid) > 0 {
			tags = map[string]string{"tag:" + kopia.TagBackupID: bid}
		}

		return &kopia.ManifestEntry{
			Manifest: &snapshot.Manifest{
				ID:               manifest.ID(id),
				IncompleteReason: incmpl,
				Tags:             tags,
			},
			Reasons: []kopia.Reason{
				{
					ResourceOwner: ro,
					Service:       path.ExchangeService,
					Category:      pct,
				},
			},
		}
	}

	table := []struct {
		name          string
		mr            mockManifestRestorer
		gb            mockGetBackuper
		reasons       []kopia.Reason
		getMeta       bool
		assertErr     assert.ErrorAssertionFunc
		assertB       assert.BoolAssertionFunc
		expectDCS     []mockColl
		expectNilMans bool
	}{
		{
			name: "don't get metadata, no mans",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{},
				mans:                []*kopia.ManifestEntry{},
			},
			gb:        mockGetBackuper{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   false,
			assertErr: assert.NoError,
			assertB:   assert.False,
			expectDCS: nil,
		},
		{
			name: "don't get metadata",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{},
				mans:                []*kopia.ManifestEntry{makeMan(path.EmailCategory, "id1", "", "")},
			},
			gb:        mockGetBackuper{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   false,
			assertErr: assert.NoError,
			assertB:   assert.False,
			expectDCS: nil,
		},
		{
			name: "don't get metadata, incomplete manifest",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{},
				mans:                []*kopia.ManifestEntry{makeMan(path.EmailCategory, "id1", "ir", "")},
			},
			gb:        mockGetBackuper{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   false,
			assertErr: assert.NoError,
			assertB:   assert.False,
			expectDCS: nil,
		},
		{
			name: "fetch manifests errors",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{},
				mrErr:               assert.AnError,
			},
			gb:        mockGetBackuper{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.Error,
			assertB:   assert.False,
			expectDCS: nil,
		},
		{
			name: "verify distinct bases fails",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{},
				mans: []*kopia.ManifestEntry{
					makeMan(path.EmailCategory, "id1", "", ""),
					makeMan(path.EmailCategory, "id2", "", ""),
				},
			},
			gb:        mockGetBackuper{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.NoError, // No error, even though verify failed.
			assertB:   assert.False,
			expectDCS: nil,
		},
		{
			name: "no manifests",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{},
				mans:                []*kopia.ManifestEntry{},
			},
			gb:        mockGetBackuper{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: nil,
		},
		{
			name: "only incomplete manifests",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{},
				mans: []*kopia.ManifestEntry{
					makeMan(path.EmailCategory, "id1", "ir", ""),
					makeMan(path.ContactsCategory, "id2", "ir", ""),
				},
			},
			gb:        mockGetBackuper{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: nil,
		},
		{
			name: "man missing backup id",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{
					collsByID: map[string][]data.RestoreCollection{
						"id": {data.NotFoundRestoreCollection{Collection: mockColl{id: "id_coll"}}},
					},
				},
				mans: []*kopia.ManifestEntry{makeMan(path.EmailCategory, "id", "", "")},
			},
			gb:            mockGetBackuper{detailsID: did},
			reasons:       []kopia.Reason{},
			getMeta:       true,
			assertErr:     assert.Error,
			assertB:       assert.False,
			expectNilMans: true,
		},
		{
			name: "backup missing details id",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{},
				mans:                []*kopia.ManifestEntry{makeMan(path.EmailCategory, "id1", "", "bid")},
			},
			gb:        mockGetBackuper{},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.False,
		},
		{
			name: "one complete, one incomplete",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{
					collsByID: map[string][]data.RestoreCollection{
						"id":        {data.NotFoundRestoreCollection{Collection: mockColl{id: "id_coll"}}},
						"incmpl_id": {data.NotFoundRestoreCollection{Collection: mockColl{id: "incmpl_id_coll"}}},
					},
				},
				mans: []*kopia.ManifestEntry{
					makeMan(path.EmailCategory, "id", "", "bid"),
					makeMan(path.EmailCategory, "incmpl_id", "ir", ""),
				},
			},
			gb:        mockGetBackuper{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: "id_coll"}},
		},
		{
			name: "single valid man",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{
					collsByID: map[string][]data.RestoreCollection{
						"id": {data.NotFoundRestoreCollection{Collection: mockColl{id: "id_coll"}}},
					},
				},
				mans: []*kopia.ManifestEntry{makeMan(path.EmailCategory, "id", "", "bid")},
			},
			gb:        mockGetBackuper{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{{id: "id_coll"}},
		},
		{
			name: "multiple valid mans",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{
					collsByID: map[string][]data.RestoreCollection{
						"mail":    {data.NotFoundRestoreCollection{Collection: mockColl{id: "mail_coll"}}},
						"contact": {data.NotFoundRestoreCollection{Collection: mockColl{id: "contact_coll"}}},
					},
				},
				mans: []*kopia.ManifestEntry{
					makeMan(path.EmailCategory, "mail", "", "bid"),
					makeMan(path.ContactsCategory, "contact", "", "bid"),
				},
			},
			gb:        mockGetBackuper{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []mockColl{
				{id: "mail_coll"},
				{id: "contact_coll"},
			},
		},
		{
			name: "error collecting metadata",
			mr: mockManifestRestorer{
				mockRestoreProducer: mockRestoreProducer{err: assert.AnError},
				mans:                []*kopia.ManifestEntry{makeMan(path.EmailCategory, "id1", "", "bid")},
			},
			gb:            mockGetBackuper{detailsID: did},
			reasons:       []kopia.Reason{},
			getMeta:       true,
			assertErr:     assert.Error,
			assertB:       assert.False,
			expectDCS:     nil,
			expectNilMans: true,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mans, dcs, b, err := produceManifestsAndMetadata(
				ctx,
				&test.mr,
				&test.gb,
				test.reasons, nil,
				tid,
				test.getMeta)
			test.assertErr(t, err, clues.ToCore(err))
			test.assertB(t, b)

			expectMans := test.mr.mans
			if test.expectNilMans {
				expectMans = nil
			}

			assert.ElementsMatch(t, expectMans, mans)

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
		})
	}
}

func (suite *OperationsManifestsUnitSuite) TestProduceManifestsAndMetadata_fallbackReasons() {
	const (
		ro            = "resourceowner"
		manComplete   = "complete"
		manIncomplete = "incmpl"

		fbro         = "fb_resourceowner"
		fbComplete   = "fb_complete"
		fbIncomplete = "fb_incmpl"
	)

	makeMan := func(id, incmpl string, reasons []kopia.Reason) *kopia.ManifestEntry {
		return &kopia.ManifestEntry{
			Manifest: &snapshot.Manifest{
				ID:               manifest.ID(id),
				IncompleteReason: incmpl,
				Tags:             map[string]string{},
			},
			Reasons: reasons,
		}
	}

	type testInput struct {
		id         string
		incomplete bool
	}

	table := []struct {
		name            string
		man             []testInput
		fallback        []testInput
		reasons         []kopia.Reason
		fallbackReasons []kopia.Reason
		manCategories   []path.CategoryType
		fbCategories    []path.CategoryType
		assertErr       assert.ErrorAssertionFunc
		expectManIDs    []string
		expectNilMans   bool
		expectReasons   map[string][]path.CategoryType
	}{
		{
			name: "only mans, no fallbacks",
			man: []testInput{
				{
					id: manComplete,
				},
				{
					id:         manIncomplete,
					incomplete: true,
				},
			},
			manCategories: []path.CategoryType{path.EmailCategory},
			fbCategories:  []path.CategoryType{path.EmailCategory},
			expectManIDs:  []string{manComplete, manIncomplete},
			expectReasons: map[string][]path.CategoryType{
				manComplete:   {path.EmailCategory},
				manIncomplete: {path.EmailCategory},
			},
		},
		{
			name: "no mans, only fallbacks",
			fallback: []testInput{
				{
					id: fbComplete,
				},
				{
					id:         fbIncomplete,
					incomplete: true,
				},
			},
			manCategories: []path.CategoryType{path.EmailCategory},
			fbCategories:  []path.CategoryType{path.EmailCategory},
			expectManIDs:  []string{fbComplete, fbIncomplete},
			expectReasons: map[string][]path.CategoryType{
				fbComplete:   {path.EmailCategory},
				fbIncomplete: {path.EmailCategory},
			},
		},
		{
			name: "complete mans and fallbacks",
			man: []testInput{
				{
					id: manComplete,
				},
			},
			fallback: []testInput{
				{
					id: fbComplete,
				},
			},
			manCategories: []path.CategoryType{path.EmailCategory},
			fbCategories:  []path.CategoryType{path.EmailCategory},
			expectManIDs:  []string{manComplete},
			expectReasons: map[string][]path.CategoryType{
				manComplete: {path.EmailCategory},
			},
		},
		{
			name: "incomplete mans and fallbacks",
			man: []testInput{
				{
					id:         manIncomplete,
					incomplete: true,
				},
			},
			fallback: []testInput{
				{
					id:         fbIncomplete,
					incomplete: true,
				},
			},
			manCategories: []path.CategoryType{path.EmailCategory},
			fbCategories:  []path.CategoryType{path.EmailCategory},
			expectManIDs:  []string{manIncomplete},
			expectReasons: map[string][]path.CategoryType{
				manIncomplete: {path.EmailCategory},
			},
		},
		{
			name: "complete and incomplete mans and fallbacks",
			man: []testInput{
				{
					id: manComplete,
				},
				{
					id:         manIncomplete,
					incomplete: true,
				},
			},
			fallback: []testInput{
				{
					id: fbComplete,
				},
				{
					id:         fbIncomplete,
					incomplete: true,
				},
			},
			manCategories: []path.CategoryType{path.EmailCategory},
			fbCategories:  []path.CategoryType{path.EmailCategory},
			expectManIDs:  []string{manComplete, manIncomplete},
			expectReasons: map[string][]path.CategoryType{
				manComplete:   {path.EmailCategory},
				manIncomplete: {path.EmailCategory},
			},
		},
		{
			name: "incomplete mans, complete fallbacks",
			man: []testInput{
				{
					id:         manIncomplete,
					incomplete: true,
				},
			},
			fallback: []testInput{
				{
					id: fbComplete,
				},
			},
			manCategories: []path.CategoryType{path.EmailCategory},
			fbCategories:  []path.CategoryType{path.EmailCategory},
			expectManIDs:  []string{fbComplete, manIncomplete},
			expectReasons: map[string][]path.CategoryType{
				fbComplete:    {path.EmailCategory},
				manIncomplete: {path.EmailCategory},
			},
		},
		{
			name: "complete mans, incomplete fallbacks",
			man: []testInput{
				{
					id: manComplete,
				},
			},
			fallback: []testInput{
				{
					id:         fbIncomplete,
					incomplete: true,
				},
			},
			manCategories: []path.CategoryType{path.EmailCategory},
			fbCategories:  []path.CategoryType{path.EmailCategory},
			expectManIDs:  []string{manComplete},
			expectReasons: map[string][]path.CategoryType{
				manComplete: {path.EmailCategory},
			},
		},
		{
			name: "complete mans, complete fallbacks, multiple reasons",
			man: []testInput{
				{
					id: manComplete,
				},
			},
			fallback: []testInput{
				{
					id: fbComplete,
				},
			},
			manCategories: []path.CategoryType{path.EmailCategory, path.ContactsCategory},
			fbCategories:  []path.CategoryType{path.EmailCategory, path.ContactsCategory},
			expectManIDs:  []string{manComplete},
			expectReasons: map[string][]path.CategoryType{
				manComplete: {path.EmailCategory, path.ContactsCategory},
			},
		},
		{
			name: "complete mans, complete fallbacks, distinct reasons",
			man: []testInput{
				{
					id: manComplete,
				},
			},
			fallback: []testInput{
				{
					id: fbComplete,
				},
			},
			manCategories: []path.CategoryType{path.ContactsCategory},
			fbCategories:  []path.CategoryType{path.EmailCategory},
			expectManIDs:  []string{manComplete, fbComplete},
			expectReasons: map[string][]path.CategoryType{
				manComplete: {path.ContactsCategory},
				fbComplete:  {path.EmailCategory},
			},
		},
		{
			name: "fb has superset of mans reasons",
			man: []testInput{
				{
					id: manComplete,
				},
			},
			fallback: []testInput{
				{
					id: fbComplete,
				},
			},
			manCategories: []path.CategoryType{path.ContactsCategory},
			fbCategories:  []path.CategoryType{path.EmailCategory, path.ContactsCategory, path.EventsCategory},
			expectManIDs:  []string{manComplete, fbComplete},
			expectReasons: map[string][]path.CategoryType{
				manComplete: {path.ContactsCategory},
				fbComplete:  {path.EmailCategory, path.EventsCategory},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mainReasons := []kopia.Reason{}
			fbReasons := []kopia.Reason{}

			for _, cat := range test.manCategories {
				mainReasons = append(
					mainReasons,
					kopia.Reason{
						ResourceOwner: ro,
						Service:       path.ExchangeService,
						Category:      cat,
					})
			}

			for _, cat := range test.fbCategories {
				fbReasons = append(
					fbReasons,
					kopia.Reason{
						ResourceOwner: fbro,
						Service:       path.ExchangeService,
						Category:      cat,
					})
			}

			mans := []*kopia.ManifestEntry{}

			for _, m := range test.man {
				incomplete := ""
				if m.incomplete {
					incomplete = "ir"
				}

				mans = append(mans, makeMan(m.id, incomplete, mainReasons))
			}

			for _, m := range test.fallback {
				incomplete := ""
				if m.incomplete {
					incomplete = "ir"
				}

				mans = append(mans, makeMan(m.id, incomplete, fbReasons))
			}

			mr := mockManifestRestorer{mans: mans}

			mans, _, b, err := produceManifestsAndMetadata(
				ctx,
				&mr,
				nil,
				mainReasons,
				fbReasons,
				"tid",
				false)
			require.NoError(t, err, clues.ToCore(err))
			assert.False(t, b, "no-metadata is forced for this test")

			manIDs := []string{}

			for _, m := range mans {
				manIDs = append(manIDs, string(m.ID))

				reasons := test.expectReasons[string(m.ID)]

				mrs := []path.CategoryType{}
				for _, r := range m.Reasons {
					mrs = append(mrs, r.Category)
				}

				assert.ElementsMatch(t, reasons, mrs)
			}

			assert.ElementsMatch(t, test.expectManIDs, manIDs)
		})
	}
}

// ---------------------------------------------------------------------------
// older tests
// ---------------------------------------------------------------------------

type BackupManifestUnitSuite struct {
	tester.Suite
}

func TestBackupManifestUnitSuite(t *testing.T) {
	suite.Run(t, &BackupManifestUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *BackupManifestUnitSuite) TestBackupOperation_VerifyDistinctBases() {
	const user = "a-user"

	table := []struct {
		name     string
		input    []*kopia.ManifestEntry
		errCheck assert.ErrorAssertionFunc
	}{
		{
			name: "SingleManifestMultipleReasons",
			input: []*kopia.ManifestEntry{
				{
					Manifest: &snapshot.Manifest{
						ID: "id1",
					},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: user,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
						{
							ResourceOwner: user,
							Service:       path.ExchangeService,
							Category:      path.EventsCategory,
						},
					},
				},
			},
			errCheck: assert.NoError,
		},
		{
			name: "MultipleManifestsDistinctReason",
			input: []*kopia.ManifestEntry{
				{
					Manifest: &snapshot.Manifest{
						ID: "id1",
					},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: user,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
					},
				},
				{
					Manifest: &snapshot.Manifest{
						ID: "id2",
					},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: user,
							Service:       path.ExchangeService,
							Category:      path.EventsCategory,
						},
					},
				},
			},
			errCheck: assert.NoError,
		},
		{
			name: "MultipleManifestsSameReason",
			input: []*kopia.ManifestEntry{
				{
					Manifest: &snapshot.Manifest{
						ID: "id1",
					},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: user,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
					},
				},
				{
					Manifest: &snapshot.Manifest{
						ID: "id2",
					},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: user,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "MultipleManifestsSameReasonOneIncomplete",
			input: []*kopia.ManifestEntry{
				{
					Manifest: &snapshot.Manifest{
						ID: "id1",
					},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: user,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
					},
				},
				{
					Manifest: &snapshot.Manifest{
						ID:               "id2",
						IncompleteReason: "checkpoint",
					},
					Reasons: []kopia.Reason{
						{
							ResourceOwner: user,
							Service:       path.ExchangeService,
							Category:      path.EmailCategory,
						},
					},
				},
			},
			errCheck: assert.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext(suite.T())
			defer flush()

			err := verifyDistinctBases(ctx, test.input)
			test.errCheck(suite.T(), err, clues.ToCore(err))
		})
	}
}

func (suite *BackupManifestUnitSuite) TestBackupOperation_CollectMetadata() {
	var (
		tenant        = "a-tenant"
		resourceOwner = "a-user"
		fileNames     = []string{
			"delta",
			"paths",
		}

		emailDeltaPath = makeMetadataPath(
			suite.T(),
			tenant,
			path.ExchangeService,
			resourceOwner,
			path.EmailCategory,
			fileNames[0],
		)
		emailPathsPath = makeMetadataPath(
			suite.T(),
			tenant,
			path.ExchangeService,
			resourceOwner,
			path.EmailCategory,
			fileNames[1],
		)
		contactsDeltaPath = makeMetadataPath(
			suite.T(),
			tenant,
			path.ExchangeService,
			resourceOwner,
			path.ContactsCategory,
			fileNames[0],
		)
		contactsPathsPath = makeMetadataPath(
			suite.T(),
			tenant,
			path.ExchangeService,
			resourceOwner,
			path.ContactsCategory,
			fileNames[1],
		)
	)

	table := []struct {
		name       string
		inputMan   *kopia.ManifestEntry
		inputFiles []string
		expected   []path.Path
	}{
		{
			name: "SingleReasonSingleFile",
			inputMan: &kopia.ManifestEntry{
				Manifest: &snapshot.Manifest{},
				Reasons: []kopia.Reason{
					{
						ResourceOwner: resourceOwner,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
			},
			inputFiles: []string{fileNames[0]},
			expected:   []path.Path{emailDeltaPath},
		},
		{
			name: "SingleReasonMultipleFiles",
			inputMan: &kopia.ManifestEntry{
				Manifest: &snapshot.Manifest{},
				Reasons: []kopia.Reason{
					{
						ResourceOwner: resourceOwner,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
			},
			inputFiles: fileNames,
			expected:   []path.Path{emailDeltaPath, emailPathsPath},
		},
		{
			name: "MultipleReasonsMultipleFiles",
			inputMan: &kopia.ManifestEntry{
				Manifest: &snapshot.Manifest{},
				Reasons: []kopia.Reason{
					{
						ResourceOwner: resourceOwner,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					{
						ResourceOwner: resourceOwner,
						Service:       path.ExchangeService,
						Category:      path.ContactsCategory,
					},
				},
			},
			inputFiles: fileNames,
			expected: []path.Path{
				emailDeltaPath,
				emailPathsPath,
				contactsDeltaPath,
				contactsPathsPath,
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mr := &mockRestoreProducer{}

			_, err := collectMetadata(ctx, mr, test.inputMan, test.inputFiles, tenant, fault.New(true))
			assert.NoError(t, err, clues.ToCore(err))

			checkPaths(t, test.expected, mr.gotPaths)
		})
	}
}
