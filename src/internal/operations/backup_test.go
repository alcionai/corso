package operations

import (
	"context"
	stdpath "path"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/mock"
	odConsts "github.com/alcionai/corso/src/internal/m365/onedrive/consts"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations/inject"
	ssmock "github.com/alcionai/corso/src/internal/streamstore/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
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

type restoreFunc func(
	id string,
	ps []path.RestorePaths,
) ([]data.RestoreCollection, error)

func (mr *mockRestoreProducer) buildRestoreFunc(
	t *testing.T,
	oid string,
	ops []path.Path,
) {
	mr.onRestore = func(
		id string,
		ps []path.RestorePaths,
	) ([]data.RestoreCollection, error) {
		gotPaths := make([]path.Path, 0, len(ps))

		for _, rp := range ps {
			gotPaths = append(gotPaths, rp.StoragePath)
		}

		assert.Equal(t, oid, id, "manifest id")
		checkPaths(t, ops, gotPaths)

		return mr.colls, mr.err
	}
}

func (mr *mockRestoreProducer) ProduceRestoreCollections(
	ctx context.Context,
	snapshotID string,
	paths []path.RestorePaths,
	bc kopia.ByteCounter,
	errs *fault.Bus,
) ([]data.RestoreCollection, error) {
	for _, ps := range paths {
		mr.gotPaths = append(mr.gotPaths, ps.StoragePath)
	}

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
		backupReasons []kopia.Reason,
		bases []kopia.IncrementalBase,
		cs []data.BackupCollection,
		tags map[string]string,
		buildTreeWithBase bool)
}

func (mbu mockBackupConsumer) ConsumeBackupCollections(
	ctx context.Context,
	backupReasons []kopia.Reason,
	bases []kopia.IncrementalBase,
	cs []data.BackupCollection,
	excluded prefixmatcher.StringSetReader,
	tags map[string]string,
	buildTreeWithBase bool,
	errs *fault.Bus,
) (*kopia.BackupStats, *details.Builder, kopia.DetailsMergeInfoer, error) {
	if mbu.checkFunc != nil {
		mbu.checkFunc(backupReasons, bases, cs, tags, buildTreeWithBase)
	}

	return &kopia.BackupStats{}, &details.Builder{}, nil, nil
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
	var (
		kw   = &kopia.Wrapper{}
		sw   = &store.Wrapper{}
		ctrl = &mock.Controller{}
		acct = account.Account{}
		now  = time.Now()
	)

	table := []struct {
		expectStatus OpStatus
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
				ctrl: &data.CollectionStats{Successes: 1},
			},
		},
		{
			expectStatus: Failed,
			expectErr:    assert.Error,
			fail:         assert.AnError,
			stats: backupStats{
				k:    &kopia.BackupStats{},
				ctrl: &data.CollectionStats{},
			},
		},
		{
			expectStatus: NoData,
			expectErr:    assert.NoError,
			stats: backupStats{
				k:    &kopia.BackupStats{},
				ctrl: &data.CollectionStats{},
			},
		},
	}
	for _, test := range table {
		suite.Run(test.expectStatus.String(), func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sel := selectors.Selector{}
			sel.DiscreteOwner = "bombadil"

			op, err := NewBackupOperation(
				ctx,
				control.Defaults(),
				kw,
				sw,
				ctrl,
				acct,
				sel,
				sel,
				evmock.NewBus())
			require.NoError(t, err, clues.ToCore(err))

			op.Errors.Fail(test.fail)

			test.expectErr(t, op.persistResults(now, &test.stats))

			assert.Equal(t, test.expectStatus.String(), op.Status.String(), "status")
			assert.Equal(t, test.stats.ctrl.Successes, op.Results.ItemsRead, "items read")
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
			path.EmailCategory.String())
		contactsBuilder = path.Builder{}.Append(
			tenant,
			path.ExchangeService.String(),
			resourceOwner,
			path.ContactsCategory.String())

		emailReason = kopia.NewReason(
			"",
			resourceOwner,
			path.ExchangeService,
			path.EmailCategory)
		contactsReason = kopia.NewReason(
			"",
			resourceOwner,
			path.ExchangeService,
			path.ContactsCategory)

		manifest1 = &snapshot.Manifest{
			ID: "id1",
		}
		manifest2 = &snapshot.Manifest{
			ID: "id2",
		}
	)

	table := []struct {
		name string
		// Backup model is untouched in this test so there's no need to populate it.
		input    kopia.BackupBases
		expected []kopia.IncrementalBase
	}{
		{
			name: "SingleManifestSingleReason",
			input: kopia.NewMockBackupBases().WithMergeBases(
				kopia.ManifestEntry{
					Manifest: manifest1,
					Reasons: []kopia.Reasoner{
						emailReason,
					},
				}).ClearMockAssistBases(),
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
			input: kopia.NewMockBackupBases().WithMergeBases(
				kopia.ManifestEntry{
					Manifest: manifest1,
					Reasons: []kopia.Reasoner{
						emailReason,
						contactsReason,
					},
				}).ClearMockAssistBases(),
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
			input: kopia.NewMockBackupBases().WithMergeBases(
				kopia.ManifestEntry{
					Manifest: manifest1,
					Reasons: []kopia.Reasoner{
						emailReason,
						contactsReason,
					},
				},
				kopia.ManifestEntry{
					Manifest: manifest2,
					Reasons: []kopia.Reasoner{
						emailReason,
						contactsReason,
					},
				}).ClearMockAssistBases(),
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
		{
			name: "Single Manifest Single Reason With Assist Base",
			input: kopia.NewMockBackupBases().WithMergeBases(
				kopia.ManifestEntry{
					Manifest: manifest1,
					Reasons: []kopia.Reasoner{
						emailReason,
					},
				}).WithAssistBases(
				kopia.ManifestEntry{
					Manifest: manifest2,
					Reasons: []kopia.Reasoner{
						contactsReason,
					},
				}),
			expected: []kopia.IncrementalBase{
				{
					Manifest: manifest1,
					SubtreePaths: []*path.Builder{
						emailBuilder,
					},
				},
				{
					Manifest: manifest2,
				},
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mbu := &mockBackupConsumer{
				checkFunc: func(
					backupReasons []kopia.Reason,
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
				test.input,
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
				odConsts.DrivesPathDir,
				"drive-id",
				odConsts.RootPathDir,
				"work",
				"item1",
			},
			true,
		)
		locationPath1 = path.Builder{}.Append(odConsts.RootPathDir, "work-display-name")
		itemPath2     = makePath(
			suite.T(),
			[]string{
				tenant,
				path.OneDriveService.String(),
				ro,
				path.FilesCategory.String(),
				odConsts.DrivesPathDir,
				"drive-id",
				odConsts.RootPathDir,
				"personal",
				"item2",
			},
			true,
		)
		locationPath2 = path.Builder{}.Append(odConsts.RootPathDir, "personal-display-name")
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

		pathReason1 = kopia.NewReason(
			"",
			itemPath1.ResourceOwner(),
			itemPath1.Service(),
			itemPath1.Category())
		pathReason3 = kopia.NewReason(
			"",
			itemPath3.ResourceOwner(),
			itemPath3.Service(),
			itemPath3.Category())
	)

	itemParents1, err := path.GetDriveFolderPath(itemPath1)
	require.NoError(suite.T(), err, clues.ToCore(err))

	itemParents1String := itemParents1.String()

	table := []struct {
		name             string
		populatedDetails map[string]*details.Details
		inputBackups     []kopia.BackupEntry
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
			name: "DetailsIDNotFound",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(itemPath1, itemPath1, locationPath1)

				return res
			}(),
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup.Backup{
						BaseModel: model.BaseModel{
							ID: backup1.ID,
						},
						DetailsID: "foo",
					},
					Reasons: []kopia.Reasoner{
						pathReason1,
					},
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
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup1,
					Reasons: []kopia.Reasoner{
						pathReason1,
					},
				},
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
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup1,
					Reasons: []kopia.Reasoner{
						pathReason1,
					},
				},
				{
					Backup: &backup1,
					Reasons: []kopia.Reasoner{
						pathReason1,
					},
				},
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
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup1,
					Reasons: []kopia.Reasoner{
						pathReason1,
					},
				},
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
										ParentPath: itemParents1String,
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
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup1,
					Reasons: []kopia.Reasoner{
						pathReason1,
					},
				},
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
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup1,
					Reasons: []kopia.Reasoner{
						pathReason1,
					},
				},
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
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup1,
					Reasons: []kopia.Reasoner{
						pathReason1,
					},
				},
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
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup1,
					Reasons: []kopia.Reasoner{
						pathReason1,
					},
				},
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
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup1,
					Reasons: []kopia.Reasoner{
						pathReason1,
					},
				},
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
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup1,
					Reasons: []kopia.Reasoner{
						pathReason1,
					},
				},
				{
					Backup: &backup2,
					Reasons: []kopia.Reasoner{
						pathReason3,
					},
				},
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
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mds := ssmock.Streamer{Deets: test.populatedDetails}
			deets := details.Builder{}
			writeStats := kopia.BackupStats{}

			err := mergeDetails(
				ctx,
				mds,
				test.inputBackups,
				test.mdm,
				&deets,
				&writeStats,
				path.OneDriveService,
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

		pathReason1 = kopia.NewReason(
			"",
			itemPath1.ResourceOwner(),
			itemPath1.Service(),
			itemPath1.Category())

		backup1 = kopia.BackupEntry{
			Backup: &backup.Backup{
				BaseModel: model.BaseModel{
					ID: "bid1",
				},
				DetailsID: "did1",
			},
			Reasons: []kopia.Reasoner{
				pathReason1,
			},
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

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		mds        = ssmock.Streamer{Deets: populatedDetails}
		deets      = details.Builder{}
		writeStats = kopia.BackupStats{}
	)

	err := mergeDetails(
		ctx,
		mds,
		[]kopia.BackupEntry{backup1},
		mdm,
		&deets,
		&writeStats,
		path.ExchangeService,
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

// ---------------------------------------------------------------------------
// integration tests
// ---------------------------------------------------------------------------

type BackupOpIntegrationSuite struct {
	tester.Suite
	user, site string
	ac         api.Client
}

func TestBackupOpIntegrationSuite(t *testing.T) {
	suite.Run(t, &BackupOpIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *BackupOpIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	graph.InitializeConcurrencyLimiter(ctx, true, 4)

	suite.user = tconfig.M365UserID(t)
	suite.site = tconfig.M365SiteID(t)

	a := tconfig.NewM365Account(t)

	creds, err := a.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	suite.ac, err = api.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *BackupOpIntegrationSuite) TestNewBackupOperation() {
	var (
		kw   = &kopia.Wrapper{}
		sw   = &store.Wrapper{}
		ctrl = &mock.Controller{}
		acct = tconfig.NewM365Account(suite.T())
		opts = control.Defaults()
	)

	table := []struct {
		name     string
		kw       *kopia.Wrapper
		sw       *store.Wrapper
		bp       inject.BackupProducer
		acct     account.Account
		targets  []string
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", kw, sw, ctrl, acct, nil, assert.NoError},
		{"missing kopia", nil, sw, ctrl, acct, nil, assert.Error},
		{"missing modelstore", kw, nil, ctrl, acct, nil, assert.Error},
		{"missing backup producer", kw, sw, nil, acct, nil, assert.Error},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			sel := selectors.Selector{DiscreteOwner: "test"}

			_, err := NewBackupOperation(
				ctx,
				opts,
				test.kw,
				test.sw,
				test.bp,
				test.acct,
				sel,
				sel,
				evmock.NewBus())
			test.errCheck(t, err, clues.ToCore(err))
		})
	}
}
