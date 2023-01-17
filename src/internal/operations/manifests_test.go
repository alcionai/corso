package operations

import (
	"context"
	"testing"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/path"
)

// ---------------------------------------------------------------------------
// interfaces
// ---------------------------------------------------------------------------

type mockManifestRestorer struct {
	mockRestorer
	mans  []*kopia.ManifestEntry
	mrErr error // err varname already claimed by mockRestorer
}

func (mmr mockManifestRestorer) FetchPrevSnapshotManifests(
	ctx context.Context,
	reasons []kopia.Reason,
	tags map[string]string,
) ([]*kopia.ManifestEntry, error) {
	return mmr.mans, mmr.mrErr
}

type mockGetDetailsIDer struct {
	detailsID string
	err       error
}

func (mg mockGetDetailsIDer) GetDetailsIDFromBackupID(
	ctx context.Context,
	backupID model.StableID,
) (string, *backup.Backup, error) {
	return mg.detailsID, nil, mg.err
}

// ---------------------------------------------------------------------------
// tests
// ---------------------------------------------------------------------------

type OperationsManifestsUnitSuite struct {
	suite.Suite
}

func TestOperationsManifestsUnitSuite(t *testing.T) {
	suite.Run(t, new(OperationsManifestsUnitSuite))
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
					p, err := emailPath.Append(f, true)
					assert.NoError(t, err)
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
					p, err := emailPath.Append(f, true)
					assert.NoError(t, err)
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
					p, err := emailPath.Append(f, true)
					assert.NoError(t, err)
					ps = append(ps, p)
					p, err = contactPath.Append(f, true)
					assert.NoError(t, err)
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
					p, err := emailPath.Append(f, true)
					assert.NoError(t, err)
					ps = append(ps, p)
					p, err = contactPath.Append(f, true)
					assert.NoError(t, err)
					ps = append(ps, p)
				}

				return ps
			},
			fileNames: []string{"a", "b"},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			paths := test.expectPaths(t, test.fileNames)

			mr := mockRestorer{err: test.expectErr}
			mr.buildRestoreFunc(t, test.manID, paths)

			man := &kopia.ManifestEntry{
				Manifest: &snapshot.Manifest{ID: manifest.ID(test.manID)},
				Reasons:  test.reasons,
			}

			_, err := collectMetadata(ctx, &mr, man, test.fileNames, tid)
			assert.ErrorIs(t, err, test.expectErr)
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
		suite.T().Run(test.name, func(t *testing.T) {
			err := verifyDistinctBases(test.mans)
			test.expect(t, err)
		})
	}
}

func (suite *OperationsManifestsUnitSuite) TestProduceManifestsAndMetadata() {
	const (
		ro  = "resourceowner"
		tid = "tenantid"
		did = "detailsid"
	)

	makeMan := func(pct path.CategoryType, ir, bid string) *kopia.ManifestEntry {
		return &kopia.ManifestEntry{
			Manifest: &snapshot.Manifest{
				IncompleteReason: ir,
				Tags:             map[string]string{"tag: " + kopia.TagBackupID: bid},
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
		gdi           mockGetDetailsIDer
		reasons       []kopia.Reason
		getMeta       bool
		assertErr     assert.ErrorAssertionFunc
		assertB       assert.BoolAssertionFunc
		expectDCS     []data.Collection
		expectNilMans bool
	}{
		{
			name: "don't get metadata, no mans",
			mr: mockManifestRestorer{
				mockRestorer: mockRestorer{},
				mans:         []*kopia.ManifestEntry{},
			},
			gdi:       mockGetDetailsIDer{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   false,
			assertErr: assert.NoError,
			assertB:   assert.False,
			expectDCS: nil,
		},
		{
			name: "don't get metadata",
			mr: mockManifestRestorer{
				mockRestorer: mockRestorer{},
				mans:         []*kopia.ManifestEntry{makeMan(path.EmailCategory, "", "")},
			},
			gdi:       mockGetDetailsIDer{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   false,
			assertErr: assert.NoError,
			assertB:   assert.False,
			expectDCS: nil,
		},
		{
			name: "don't get metadata, incomplete manifest",
			mr: mockManifestRestorer{
				mockRestorer: mockRestorer{},
				mans:         []*kopia.ManifestEntry{makeMan(path.EmailCategory, "ir", "")},
			},
			gdi:       mockGetDetailsIDer{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   false,
			assertErr: assert.NoError,
			assertB:   assert.False,
			expectDCS: nil,
		},
		{
			name: "fetch manifests errors",
			mr: mockManifestRestorer{
				mockRestorer: mockRestorer{},
				mrErr:        assert.AnError,
			},
			gdi:       mockGetDetailsIDer{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.Error,
			assertB:   assert.False,
			expectDCS: nil,
		},
		{
			name: "verify distinct bases fails",
			mr: mockManifestRestorer{
				mockRestorer: mockRestorer{},
				mans: []*kopia.ManifestEntry{
					makeMan(path.EmailCategory, "", ""),
					makeMan(path.EmailCategory, "", ""),
				},
			},
			gdi:       mockGetDetailsIDer{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.NoError, // No error, even though verify failed.
			assertB:   assert.False,
			expectDCS: nil,
		},
		{
			name: "no manifests",
			mr: mockManifestRestorer{
				mockRestorer: mockRestorer{},
				mans:         []*kopia.ManifestEntry{},
			},
			gdi:       mockGetDetailsIDer{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: nil,
		},
		{
			name: "only incomplete manifests",
			mr: mockManifestRestorer{
				mockRestorer: mockRestorer{},
				mans: []*kopia.ManifestEntry{
					makeMan(path.EmailCategory, "ir", ""),
					makeMan(path.ContactsCategory, "ir", ""),
				},
			},
			gdi:       mockGetDetailsIDer{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: nil,
		},
		{
			name: "man missing backup id",
			mr: mockManifestRestorer{
				mockRestorer: mockRestorer{colls: []data.Collection{}},
				mans:         []*kopia.ManifestEntry{makeMan(path.EmailCategory, "", "")},
			},
			gdi:           mockGetDetailsIDer{detailsID: did},
			reasons:       []kopia.Reason{},
			getMeta:       true,
			assertErr:     assert.Error,
			assertB:       assert.False,
			expectDCS:     []data.Collection{},
			expectNilMans: true,
		},
		{
			name: "backup missing details id",
			mr: mockManifestRestorer{
				mockRestorer: mockRestorer{colls: []data.Collection{}},
				mans:         []*kopia.ManifestEntry{makeMan(path.EmailCategory, "", "bid")},
			},
			gdi:           mockGetDetailsIDer{},
			reasons:       []kopia.Reason{},
			getMeta:       true,
			assertErr:     assert.Error,
			assertB:       assert.False,
			expectDCS:     []data.Collection{},
			expectNilMans: true,
		},
		{
			name: "one complete man, one incomplete",
			mr: mockManifestRestorer{
				mockRestorer: mockRestorer{colls: []data.Collection{}},
				mans: []*kopia.ManifestEntry{
					makeMan(path.EmailCategory, "", "bid"),
					makeMan(path.EmailCategory, "ir", ""),
				},
			},
			gdi:       mockGetDetailsIDer{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []data.Collection{},
		},
		{
			name: "happy path",
			mr: mockManifestRestorer{
				mockRestorer: mockRestorer{colls: []data.Collection{}},
				mans:         []*kopia.ManifestEntry{makeMan(path.EmailCategory, "", "bid")},
			},
			gdi:       mockGetDetailsIDer{detailsID: did},
			reasons:   []kopia.Reason{},
			getMeta:   true,
			assertErr: assert.NoError,
			assertB:   assert.True,
			expectDCS: []data.Collection{},
		},
		{
			name: "error collecting metadata",
			mr: mockManifestRestorer{
				mockRestorer: mockRestorer{err: assert.AnError},
				mans:         []*kopia.ManifestEntry{makeMan(path.EmailCategory, "", "bid")},
			},
			gdi:           mockGetDetailsIDer{detailsID: did},
			reasons:       []kopia.Reason{},
			getMeta:       true,
			assertErr:     assert.Error,
			assertB:       assert.False,
			expectDCS:     nil,
			expectNilMans: true,
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			mans, dcs, b, err := produceManifestsAndMetadata(
				ctx,
				&test.mr,
				&test.gdi,
				test.reasons,
				tid,
				test.getMeta,
			)
			test.assertErr(t, err)
			test.assertB(t, b)

			expectMans := test.mr.mans
			if test.expectNilMans {
				expectMans = nil
			}
			assert.Equal(t, expectMans, mans)

			assert.Len(t, dcs, len(test.expectDCS))
			for _, dc := range test.expectDCS {
				var (
					found bool
					s     = dc.FullPath().String()
				)

				for _, r := range dcs {
					if r.FullPath().String() == s {
						found = true
						break
					}
				}

				assert.True(t, found, "expected collection is present in results: "+s)
			}
		})
	}
}
