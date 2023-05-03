package operations

import (
	"context"
	"fmt"
	stdpath "path"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/mock"
	"github.com/alcionai/corso/src/internal/data"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	ssmock "github.com/alcionai/corso/src/internal/streamstore/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

// ---------------------------------------------------------------------------
// mocks
// ---------------------------------------------------------------------------

// ----- restore producer

type mockRestoreProducer struct {
	gotPaths  []path.Path
	colls     []data.RestoreCollection
	collsByID map[string][]data.RestoreCollection // snapshotID: []RestoreCollection
	err       error
	onRestore restoreFunc
}

type restoreFunc func(id string, ps []path.Path) ([]data.RestoreCollection, error)

func (mr *mockRestoreProducer) buildRestoreFunc(
	t *testing.T,
	oid string,
	ops []path.Path,
) {
	mr.onRestore = func(id string, ps []path.Path) ([]data.RestoreCollection, error) {
		assert.Equal(t, oid, id, "manifest id")
		checkPaths(t, ops, ps)

		return mr.colls, mr.err
	}
}

func (mr *mockRestoreProducer) ProduceRestoreCollections(
	ctx context.Context,
	snapshotID string,
	paths []path.Path,
	bc kopia.ByteCounter,
	errs *fault.Bus,
) ([]data.RestoreCollection, error) {
	mr.gotPaths = append(mr.gotPaths, paths...)

	if mr.onRestore != nil {
		return mr.onRestore(snapshotID, paths)
	}

	if len(mr.collsByID) > 0 {
		return mr.collsByID[snapshotID], mr.err
	}

	return mr.colls, mr.err
}

func checkPaths(t *testing.T, expected, got []path.Path) {
	assert.ElementsMatch(t, expected, got)
}

// ----- backup consumer

type mockBackupConsumer struct {
	checkFunc func(
		bases []kopia.IncrementalBase,
		cs []data.BackupCollection,
		tags map[string]string,
		buildTreeWithBase bool)
}

func (mbu mockBackupConsumer) ConsumeBackupCollections(
	ctx context.Context,
	bases []kopia.IncrementalBase,
	cs []data.BackupCollection,
	excluded map[string]map[string]struct{},
	tags map[string]string,
	buildTreeWithBase bool,
	errs *fault.Bus,
) (*kopia.BackupStats, *details.Builder, kopia.DetailsMergeInfoer, error) {
	if mbu.checkFunc != nil {
		mbu.checkFunc(bases, cs, tags, buildTreeWithBase)
	}

	return &kopia.BackupStats{}, &details.Builder{}, nil, nil
}

// ----- model store for backups

type mockBackupStorer struct {
	// Only using this to store backup models right now.
	entries map[model.StableID]backup.Backup
}

func (mbs mockBackupStorer) Get(
	ctx context.Context,
	s model.Schema,
	id model.StableID,
	toPopulate model.Model,
) error {
	ctx = clues.Add(
		ctx,
		"model_schema", s,
		"model_id", id,
		"model_type", fmt.Sprintf("%T", toPopulate))

	if s != model.BackupSchema {
		return clues.New("unexpected schema").WithClues(ctx)
	}

	r, ok := mbs.entries[id]
	if !ok {
		return clues.New("model not found").WithClues(ctx)
	}

	bu, ok := toPopulate.(*backup.Backup)
	if !ok {
		return clues.New("bad population type").WithClues(ctx)
	}

	*bu = r

	return nil
}

func (mbs mockBackupStorer) Delete(context.Context, model.Schema, model.StableID) error {
	return clues.New("not implemented")
}

func (mbs mockBackupStorer) DeleteWithModelStoreID(context.Context, manifest.ID) error {
	return clues.New("not implemented")
}

func (mbs mockBackupStorer) GetIDsForType(
	context.Context,
	model.Schema,
	map[string]string,
) ([]*model.BaseModel, error) {
	return nil, clues.New("not implemented")
}

func (mbs mockBackupStorer) GetWithModelStoreID(
	context.Context,
	model.Schema,
	manifest.ID,
	model.Model,
) error {
	return clues.New("not implemented")
}

func (mbs mockBackupStorer) Put(context.Context, model.Schema, model.Model) error {
	return clues.New("not implemented")
}

func (mbs mockBackupStorer) Update(context.Context, model.Schema, model.Model) error {
	return clues.New("not implemented")
}

// ----- model store for backups

type mockDetailsMergeInfoer struct {
	repoRefs map[string]path.Path
	locs     map[string]*path.Builder
}

func (m *mockDetailsMergeInfoer) add(oldRef, newRef path.Path, newLoc *path.Builder) {
	oldPB := oldRef.ToBuilder()
	// Items are indexed individually.
	m.repoRefs[oldPB.ShortRef()] = newRef

	// Locations are indexed by directory.
	m.locs[oldPB.ShortRef()] = newLoc
}

func (m *mockDetailsMergeInfoer) GetNewPathRefs(
	oldRef *path.Builder,
	_ details.LocationIDer,
) (path.Path, *path.Builder, error) {
	return m.repoRefs[oldRef.ShortRef()], m.locs[oldRef.ShortRef()], nil
}

func (m *mockDetailsMergeInfoer) ItemsToMerge() int {
	if m == nil {
		return 0
	}

	return len(m.repoRefs)
}

func newMockDetailsMergeInfoer() *mockDetailsMergeInfoer {
	return &mockDetailsMergeInfoer{
		repoRefs: map[string]path.Path{},
		locs:     map[string]*path.Builder{},
	}
}

// ---------------------------------------------------------------------------
// helper funcs
// ---------------------------------------------------------------------------

// expects you to Append your own file
func makeMetadataBasePath(
	t *testing.T,
	tenant string,
	service path.ServiceType,
	resourceOwner string,
	category path.CategoryType,
) path.Path {
	t.Helper()

	p, err := path.Builder{}.ToServiceCategoryMetadataPath(
		tenant,
		resourceOwner,
		service,
		category,
		false)
	require.NoError(t, err, clues.ToCore(err))

	return p
}

func makeMetadataPath(
	t *testing.T,
	tenant string,
	service path.ServiceType,
	resourceOwner string,
	category path.CategoryType,
	fileName string,
) path.Path {
	t.Helper()

	p, err := path.Builder{}.Append(fileName).ToServiceCategoryMetadataPath(
		tenant,
		resourceOwner,
		service,
		category,
		true)
	require.NoError(t, err, clues.ToCore(err))

	return p
}

func makeFolderEntry(
	t *testing.T,
	pb, loc *path.Builder,
	size int64,
	modTime time.Time,
	dt details.ItemType,
) *details.Entry {
	t.Helper()

	return &details.Entry{
		RepoRef:     pb.String(),
		ShortRef:    pb.ShortRef(),
		ParentRef:   pb.Dir().ShortRef(),
		LocationRef: loc.Dir().String(),
		ItemInfo: details.ItemInfo{
			Folder: &details.FolderInfo{
				ItemType:    details.FolderItem,
				DisplayName: pb.LastElem(),
				Modified:    modTime,
				Size:        size,
				DataType:    dt,
			},
		},
	}
}

// TODO(ashmrtn): Really need to factor a function like this out into some
// common file that is only compiled for tests.
func makePath(t *testing.T, elements []string, isItem bool) path.Path {
	t.Helper()

	p, err := path.FromDataLayerPath(stdpath.Join(elements...), isItem)
	require.NoError(t, err, clues.ToCore(err))

	return p
}

func makeDetailsEntry(
	t *testing.T,
	p path.Path,
	l *path.Builder,
	size int,
	updated bool,
) *details.Entry {
	t.Helper()

	var lr string
	if l != nil {
		lr = l.String()
	}

	res := &details.Entry{
		RepoRef:     p.String(),
		ShortRef:    p.ShortRef(),
		ParentRef:   p.ToBuilder().Dir().ShortRef(),
		ItemRef:     p.Item(),
		LocationRef: lr,
		ItemInfo:    details.ItemInfo{},
		Updated:     updated,
	}

	switch p.Service() {
	case path.ExchangeService:
		if p.Category() != path.EmailCategory {
			assert.FailNowf(
				t,
				"category %s not supported in helper function",
				p.Category().String(),
			)
		}

		res.Exchange = &details.ExchangeInfo{
			ItemType:   details.ExchangeMail,
			Size:       int64(size),
			ParentPath: l.String(),
		}

	case path.OneDriveService:
		require.NotNil(t, l)

		res.OneDrive = &details.OneDriveInfo{
			ItemType:   details.OneDriveItem,
			ParentPath: l.PopFront().String(),
			Size:       int64(size),
			DriveID:    "drive-id",
			DriveName:  "drive-name",
		}

	default:
		assert.FailNowf(
			t,
			"service %s not supported in helper function",
			p.Service().String())
	}

	return res
}

// TODO(ashmrtn): This should belong to some code that lives in the kopia
// package that is only compiled when running tests.
func makeKopiaTagKey(k string) string {
	return "tag:" + k
}

func makeManifest(t *testing.T, backupID model.StableID, incompleteReason string) *snapshot.Manifest {
	t.Helper()

	tagKey := makeKopiaTagKey(kopia.TagBackupID)

	return &snapshot.Manifest{
		Tags: map[string]string{
			tagKey: string(backupID),
		},
		IncompleteReason: incompleteReason,
	}
}

// ---------------------------------------------------------------------------
// unit tests
// ---------------------------------------------------------------------------

type BackupOpUnitSuite struct {
	tester.Suite
}

func TestBackupOpUnitSuite(t *testing.T) {
	suite.Run(t, &BackupOpUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *BackupOpUnitSuite) TestBackupOperation_PersistResults() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		kw   = &kopia.Wrapper{}
		sw   = &store.Wrapper{}
		gc   = &mock.GraphConnector{}
		acct = account.Account{}
		now  = time.Now()
	)

	table := []struct {
		expectStatus opStatus
		expectErr    assert.ErrorAssertionFunc
		stats        backupStats
		fail         error
	}{
		{
			expectStatus: Completed,
			expectErr:    assert.NoError,
			stats: backupStats{
				resourceCount: 1,
				k: &kopia.BackupStats{
					TotalFileCount:     1,
					TotalHashedBytes:   1,
					TotalUploadedBytes: 1,
				},
				gc: &data.CollectionStats{Successes: 1},
			},
		},
		{
			expectStatus: Failed,
			expectErr:    assert.Error,
			fail:         assert.AnError,
			stats: backupStats{
				k:  &kopia.BackupStats{},
				gc: &data.CollectionStats{},
			},
		},
		{
			expectStatus: NoData,
			expectErr:    assert.NoError,
			stats: backupStats{
				k:  &kopia.BackupStats{},
				gc: &data.CollectionStats{},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.expectStatus.String(), func() {
			t := suite.T()
			sel := selectors.Selector{}
			sel.DiscreteOwner = "bombadil"

			op, err := NewBackupOperation(
				ctx,
				control.Defaults(),
				kw,
				sw,
				gc,
				acct,
				sel,
				sel,
				evmock.NewBus())
			require.NoError(t, err, clues.ToCore(err))

			op.Errors.Fail(test.fail)

			test.expectErr(t, op.persistResults(now, &test.stats))

			assert.Equal(t, test.expectStatus.String(), op.Status.String(), "status")
			assert.Equal(t, test.stats.gc.Successes, op.Results.ItemsRead, "items read")
			assert.Equal(t, test.stats.k.TotalFileCount, op.Results.ItemsWritten, "items written")
			assert.Equal(t, test.stats.k.TotalHashedBytes, op.Results.BytesRead, "bytes read")
			assert.Equal(t, test.stats.k.TotalUploadedBytes, op.Results.BytesUploaded, "bytes written")
			assert.Equal(t, test.stats.resourceCount, op.Results.ResourceOwners, "resource owners")
			assert.Equal(t, now, op.Results.StartedAt, "started at")
			assert.Less(t, now, op.Results.CompletedAt, "completed at")
		})
	}
}

func (suite *BackupOpUnitSuite) TestBackupOperation_ConsumeBackupDataCollections_Paths() {
	var (
		tenant        = "a-tenant"
		resourceOwner = "a-user"

		emailBuilder = path.Builder{}.Append(
			tenant,
			path.ExchangeService.String(),
			resourceOwner,
			path.EmailCategory.String(),
		)
		contactsBuilder = path.Builder{}.Append(
			tenant,
			path.ExchangeService.String(),
			resourceOwner,
			path.ContactsCategory.String(),
		)

		emailReason = kopia.Reason{
			ResourceOwner: resourceOwner,
			Service:       path.ExchangeService,
			Category:      path.EmailCategory,
		}
		contactsReason = kopia.Reason{
			ResourceOwner: resourceOwner,
			Service:       path.ExchangeService,
			Category:      path.ContactsCategory,
		}

		manifest1 = &snapshot.Manifest{
			ID: "id1",
		}
		manifest2 = &snapshot.Manifest{
			ID: "id2",
		}
	)

	table := []struct {
		name     string
		inputMan []*kopia.ManifestEntry
		expected []kopia.IncrementalBase
	}{
		{
			name: "SingleManifestSingleReason",
			inputMan: []*kopia.ManifestEntry{
				{
					Manifest: manifest1,
					Reasons: []kopia.Reason{
						emailReason,
					},
				},
			},
			expected: []kopia.IncrementalBase{
				{
					Manifest: manifest1,
					SubtreePaths: []*path.Builder{
						emailBuilder,
					},
				},
			},
		},
		{
			name: "SingleManifestMultipleReasons",
			inputMan: []*kopia.ManifestEntry{
				{
					Manifest: manifest1,
					Reasons: []kopia.Reason{
						emailReason,
						contactsReason,
					},
				},
			},
			expected: []kopia.IncrementalBase{
				{
					Manifest: manifest1,
					SubtreePaths: []*path.Builder{
						emailBuilder,
						contactsBuilder,
					},
				},
			},
		},
		{
			name: "MultipleManifestsMultipleReasons",
			inputMan: []*kopia.ManifestEntry{
				{
					Manifest: manifest1,
					Reasons: []kopia.Reason{
						emailReason,
						contactsReason,
					},
				},
				{
					Manifest: manifest2,
					Reasons: []kopia.Reason{
						emailReason,
						contactsReason,
					},
				},
			},
			expected: []kopia.IncrementalBase{
				{
					Manifest: manifest1,
					SubtreePaths: []*path.Builder{
						emailBuilder,
						contactsBuilder,
					},
				},
				{
					Manifest: manifest2,
					SubtreePaths: []*path.Builder{
						emailBuilder,
						contactsBuilder,
					},
				},
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			mbu := &mockBackupConsumer{
				checkFunc: func(
					bases []kopia.IncrementalBase,
					cs []data.BackupCollection,
					tags map[string]string,
					buildTreeWithBase bool,
				) {
					assert.ElementsMatch(t, test.expected, bases)
				},
			}

			//nolint:errcheck
			consumeBackupCollections(
				ctx,
				mbu,
				tenant,
				nil,
				test.inputMan,
				nil,
				nil,
				model.StableID(""),
				true,
				fault.New(true))
		})
	}
}

func (suite *BackupOpUnitSuite) TestBackupOperation_MergeBackupDetails_AddsItems() {
	var (
		tenant = "a-tenant"
		ro     = "a-user"

		itemPath1 = makePath(
			suite.T(),
			[]string{
				tenant,
				path.OneDriveService.String(),
				ro,
				path.FilesCategory.String(),
				"drives",
				"drive-id",
				"root:",
				"work",
				"item1",
			},
			true,
		)
		locationPath1 = path.Builder{}.Append("root:", "work-display-name")
		itemPath2     = makePath(
			suite.T(),
			[]string{
				tenant,
				path.OneDriveService.String(),
				ro,
				path.FilesCategory.String(),
				"drives",
				"drive-id",
				"root:",
				"personal",
				"item2",
			},
			true,
		)
		locationPath2 = path.Builder{}.Append("root:", "personal-display-name")
		itemPath3     = makePath(
			suite.T(),
			[]string{
				tenant,
				path.ExchangeService.String(),
				ro,
				path.EmailCategory.String(),
				"personal",
				"item3",
			},
			true,
		)
		locationPath3 = path.Builder{}.Append("personal-display-name")

		backup1 = backup.Backup{
			BaseModel: model.BaseModel{
				ID: "bid1",
			},
			DetailsID: "did1",
		}

		backup2 = backup.Backup{
			BaseModel: model.BaseModel{
				ID: "bid2",
			},
			DetailsID: "did2",
		}

		pathReason1 = kopia.Reason{
			ResourceOwner: itemPath1.ResourceOwner(),
			Service:       itemPath1.Service(),
			Category:      itemPath1.Category(),
		}
		pathReason3 = kopia.Reason{
			ResourceOwner: itemPath3.ResourceOwner(),
			Service:       itemPath3.Service(),
			Category:      itemPath3.Category(),
		}
	)

	itemParents1, err := path.GetDriveFolderPath(itemPath1)
	require.NoError(suite.T(), err, clues.ToCore(err))

	table := []struct {
		name             string
		populatedModels  map[model.StableID]backup.Backup
		populatedDetails map[string]*details.Details
		inputMans        []*kopia.ManifestEntry
		mdm              *mockDetailsMergeInfoer

		errCheck        assert.ErrorAssertionFunc
		expectedEntries []*details.Entry
	}{
		{
			name:     "NilShortRefsFromPrevBackup",
			errCheck: assert.NoError,
			// Use empty slice so we don't error out on nil != empty.
			expectedEntries: []*details.Entry{},
		},
		{
			name:     "EmptyShortRefsFromPrevBackup",
			mdm:      newMockDetailsMergeInfoer(),
			errCheck: assert.NoError,
			// Use empty slice so we don't error out on nil != empty.
			expectedEntries: []*details.Entry{},
		},
		{
			name: "BackupIDNotFound",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(itemPath1, itemPath1, locationPath1)

				return res
			}(),
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(suite.T(), "foo", ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "DetailsIDNotFound",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(itemPath1, itemPath1, locationPath1)

				return res
			}(),
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(suite.T(), backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
			},
			populatedModels: map[model.StableID]backup.Backup{
				backup1.ID: {
					BaseModel: model.BaseModel{
						ID: backup1.ID,
					},
					DetailsID: "foo",
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "BaseMissingItems",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(itemPath1, itemPath1, locationPath1)
				res.add(itemPath2, itemPath2, locationPath2)

				return res
			}(),
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(suite.T(), backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
			},
			populatedModels: map[model.StableID]backup.Backup{
				backup1.ID: backup1,
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "TooManyItems",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(itemPath1, itemPath1, locationPath1)

				return res
			}(),
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(suite.T(), backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
				{
					Manifest: makeManifest(suite.T(), backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
			},
			populatedModels: map[model.StableID]backup.Backup{
				backup1.ID: backup1,
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "BadBaseRepoRef",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(itemPath1, itemPath2, locationPath2)

				return res
			}(),
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(suite.T(), backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
			},
			populatedModels: map[model.StableID]backup.Backup{
				backup1.ID: backup1,
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							{
								RepoRef: stdpath.Join(
									append(
										[]string{
											itemPath1.Tenant(),
											itemPath1.Service().String(),
											itemPath1.ResourceOwner(),
											path.UnknownCategory.String(),
										},
										itemPath1.Folders()...,
									)...,
								),
								ItemInfo: details.ItemInfo{
									OneDrive: &details.OneDriveInfo{
										ItemType:   details.OneDriveItem,
										ParentPath: itemParents1,
										Size:       42,
									},
								},
							},
						},
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "BadOneDrivePath",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				p := makePath(
					suite.T(),
					[]string{
						itemPath1.Tenant(),
						path.OneDriveService.String(),
						itemPath1.ResourceOwner(),
						path.FilesCategory.String(),
						"personal",
						"item1",
					},
					true,
				)

				res.add(itemPath1, p, nil)

				return res
			}(),
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(suite.T(), backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
			},
			populatedModels: map[model.StableID]backup.Backup{
				backup1.ID: backup1,
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "ItemMerged",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(itemPath1, itemPath1, locationPath1)

				return res
			}(),
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(suite.T(), backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
			},
			populatedModels: map[model.StableID]backup.Backup{
				backup1.ID: backup1,
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.Entry{
				makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
			},
		},
		{
			name: "ItemMergedSameLocation",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(itemPath1, itemPath1, locationPath1)

				return res
			}(),
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(suite.T(), backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
			},
			populatedModels: map[model.StableID]backup.Backup{
				backup1.ID: backup1,
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.Entry{
				makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
			},
		},
		{
			name: "ItemMergedExtraItemsInBase",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(itemPath1, itemPath1, locationPath1)

				return res
			}(),
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(suite.T(), backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
			},
			populatedModels: map[model.StableID]backup.Backup{
				backup1.ID: backup1,
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
							*makeDetailsEntry(suite.T(), itemPath2, locationPath2, 84, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.Entry{
				makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
			},
		},
		{
			name: "ItemMoved",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(itemPath1, itemPath2, locationPath2)

				return res
			}(),
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(suite.T(), backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
			},
			populatedModels: map[model.StableID]backup.Backup{
				backup1.ID: backup1,
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.Entry{
				makeDetailsEntry(suite.T(), itemPath2, locationPath2, 42, true),
			},
		},
		{
			name: "MultipleBases",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(itemPath1, itemPath1, locationPath1)
				res.add(itemPath3, itemPath3, locationPath3)

				return res
			}(),
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(suite.T(), backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
				{
					Manifest: makeManifest(suite.T(), backup2.ID, ""),
					Reasons: []kopia.Reason{
						pathReason3,
					},
				},
			},
			populatedModels: map[model.StableID]backup.Backup{
				backup1.ID: backup1,
				backup2.ID: backup2,
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
						},
					},
				},
				backup2.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							// This entry should not be picked due to a mismatch on Reasons.
							*makeDetailsEntry(suite.T(), itemPath1, locationPath1, 84, false),
							// This item should be picked.
							*makeDetailsEntry(suite.T(), itemPath3, locationPath3, 37, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.Entry{
				makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
				makeDetailsEntry(suite.T(), itemPath3, locationPath3, 37, false),
			},
		},
		{
			name: "SomeBasesIncomplete",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(itemPath1, itemPath1, locationPath1)

				return res
			}(),
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(suite.T(), backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
				{
					Manifest: makeManifest(suite.T(), backup2.ID, "checkpoint"),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
			},
			populatedModels: map[model.StableID]backup.Backup{
				backup1.ID: backup1,
				backup2.ID: backup2,
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
						},
					},
				},
				backup2.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							// This entry should not be picked due to being incomplete.
							*makeDetailsEntry(suite.T(), itemPath1, locationPath1, 84, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.Entry{
				makeDetailsEntry(suite.T(), itemPath1, locationPath1, 42, false),
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext()
			defer flush()

			mds := ssmock.Streamer{Deets: test.populatedDetails}
			w := &store.Wrapper{Storer: mockBackupStorer{entries: test.populatedModels}}
			deets := details.Builder{}

			err := mergeDetails(
				ctx,
				w,
				mds,
				test.inputMans,
				test.mdm,
				&deets,
				fault.New(true))
			test.errCheck(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.ElementsMatch(t, test.expectedEntries, deets.Details().Items())
		})
	}
}

func (suite *BackupOpUnitSuite) TestBackupOperation_MergeBackupDetails_AddsFolders() {
	var (
		t = suite.T()

		tenant = "a-tenant"
		ro     = "a-user"

		pathElems = []string{
			tenant,
			path.ExchangeService.String(),
			ro,
			path.EmailCategory.String(),
			"work",
			"project8",
			"item1",
		}

		itemPath1 = makePath(
			t,
			pathElems,
			true)

		locPath1 = path.Builder{}.Append(itemPath1.Folders()...)

		backup1 = backup.Backup{
			BaseModel: model.BaseModel{
				ID: "bid1",
			},
			DetailsID: "did1",
		}

		pathReason1 = kopia.Reason{
			ResourceOwner: itemPath1.ResourceOwner(),
			Service:       itemPath1.Service(),
			Category:      itemPath1.Category(),
		}

		inputMans = []*kopia.ManifestEntry{
			{
				Manifest: makeManifest(t, backup1.ID, ""),
				Reasons: []kopia.Reason{
					pathReason1,
				},
			},
		}

		populatedModels = map[model.StableID]backup.Backup{
			backup1.ID: backup1,
		}

		itemSize = 42
		now      = time.Now()
		// later    = now.Add(42 * time.Minute)
	)

	mdm := newMockDetailsMergeInfoer()
	mdm.add(itemPath1, itemPath1, locPath1)

	itemDetails := makeDetailsEntry(t, itemPath1, locPath1, itemSize, false)
	// itemDetails.Exchange.Modified = now

	populatedDetails := map[string]*details.Details{
		backup1.DetailsID: {
			DetailsModel: details.DetailsModel{
				Entries: []details.Entry{*itemDetails},
			},
		},
	}

	expectedEntries := []details.Entry{*itemDetails}

	// update the details
	itemDetails.Exchange.Modified = now

	for i := 1; i <= len(locPath1.Elements()); i++ {
		expectedEntries = append(expectedEntries, *makeFolderEntry(
			t,
			// Include prefix elements in the RepoRef calculations.
			path.Builder{}.Append(pathElems[:4+i]...),
			path.Builder{}.Append(locPath1.Elements()[:i]...),
			int64(itemSize),
			itemDetails.Exchange.Modified,
			details.ExchangeMail))
	}

	ctx, flush := tester.NewContext()
	defer flush()

	var (
		mds   = ssmock.Streamer{Deets: populatedDetails}
		w     = &store.Wrapper{Storer: mockBackupStorer{entries: populatedModels}}
		deets = details.Builder{}
	)

	err := mergeDetails(
		ctx,
		w,
		mds,
		inputMans,
		mdm,
		&deets,
		fault.New(true))
	assert.NoError(t, err, clues.ToCore(err))
	compareDeetEntries(t, expectedEntries, deets.Details().Entries)
}

// compares two details slices.  Useful for tests where serializing the
// entries can produce minor variations in the time struct, causing
// assert.elementsMatch to fail.
func compareDeetEntries(
	t *testing.T,
	expect, result []details.Entry,
) {
	if !assert.Equal(t, len(expect), len(result), "entry slices should be equal len") {
		require.ElementsMatch(t, expect, result)
	}

	var (
		// repoRef -> modified time
		eMods = map[string]time.Time{}
		es    = make([]details.Entry, 0, len(expect))
		rs    = make([]details.Entry, 0, len(expect))
	)

	for _, e := range expect {
		eMods[e.RepoRef] = e.Modified()
		es = append(es, withoutModified(e))
	}

	for _, r := range result {
		// this comparison is an artifact of bad comparisons across time.Time
		// serialization using assert.ElementsMatch.  The time struct can produce
		// differences in its `ext` value across serialization while the actual time
		// reference remains the same.  assert handles this poorly, whereas the time
		// library provides successful comparison.
		assert.Truef(
			t,
			eMods[r.RepoRef].Equal(r.Modified()),
			"expected modified time %v, got %v", eMods[r.RepoRef], r.Modified())

		rs = append(rs, withoutModified(r))
	}

	assert.ElementsMatch(t, es, rs)
}

func withoutModified(de details.Entry) details.Entry {
	switch {
	case de.Exchange != nil:
		de.Exchange.Modified = time.Time{}

	case de.OneDrive != nil:
		de.OneDrive.Modified = time.Time{}

	case de.SharePoint != nil:
		de.SharePoint.Modified = time.Time{}

	case de.Folder != nil:
		de.Folder.Modified = time.Time{}
	}

	return de
}
