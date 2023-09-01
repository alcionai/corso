package operations

import (
	"context"
	"encoding/json"
	stdpath "path"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/cli/config"
	"github.com/alcionai/corso/src/internal/common/prefixmatcher"
	"github.com/alcionai/corso/src/internal/data"
	dataMock "github.com/alcionai/corso/src/internal/data/mock"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/m365/mock"
	odConsts "github.com/alcionai/corso/src/internal/m365/service/onedrive/consts"
	odMock "github.com/alcionai/corso/src/internal/m365/service/onedrive/mock"
	odStub "github.com/alcionai/corso/src/internal/m365/service/onedrive/stub"
	"github.com/alcionai/corso/src/internal/m365/support"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/operations/inject"
	"github.com/alcionai/corso/src/internal/streamstore"
	ssmock "github.com/alcionai/corso/src/internal/streamstore/mock"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	deeTD "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/control/repository"
	"github.com/alcionai/corso/src/pkg/extensions"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	selTD "github.com/alcionai/corso/src/pkg/selectors/testdata"
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
		backupReasons []identity.Reasoner,
		bases kopia.BackupBases,
		cs []data.BackupCollection,
		tags map[string]string,
		buildTreeWithBase bool)
}

func (mbu mockBackupConsumer) ConsumeBackupCollections(
	ctx context.Context,
	backupReasons []identity.Reasoner,
	bases kopia.BackupBases,
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
	modTimes map[string]time.Time
}

func (m *mockDetailsMergeInfoer) add(oldRef, newRef path.Path, newLoc *path.Builder) {
	oldPB := oldRef.ToBuilder()
	// Items are indexed individually.
	m.repoRefs[oldPB.ShortRef()] = newRef

	// Locations are indexed by directory.
	m.locs[oldPB.ShortRef()] = newLoc
}

func (m *mockDetailsMergeInfoer) addWithModTime(
	oldRef path.Path,
	modTime time.Time,
	newRef path.Path,
	newLoc *path.Builder,
) {
	oldPB := oldRef.ToBuilder()
	// Items are indexed individually.
	m.repoRefs[oldPB.ShortRef()] = newRef
	m.modTimes[oldPB.ShortRef()] = modTime

	// Locations are indexed by directory.
	m.locs[oldPB.ShortRef()] = newLoc
}

func (m *mockDetailsMergeInfoer) GetNewPathRefs(
	oldRef *path.Builder,
	modTime time.Time,
	_ details.LocationIDer,
) (path.Path, *path.Builder, error) {
	// Return no match if the modTime was set and it wasn't what was passed in.
	if mt, ok := m.modTimes[oldRef.ShortRef()]; ok && !mt.Equal(modTime) {
		return nil, nil, nil
	}

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
		modTimes: map[string]time.Time{},
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

func makeDetailsEntryWithModTime(
	t *testing.T,
	p path.Path,
	l *path.Builder,
	size int,
	updated bool,
	modTime time.Time,
) *details.Entry {
	t.Helper()

	res := makeDetailsEntry(t, p, l, size, updated)

	switch {
	case res.Exchange != nil:
		res.Exchange.Modified = modTime
	case res.OneDrive != nil:
		res.OneDrive.Modified = modTime
	case res.SharePoint != nil:
		res.SharePoint.Modified = modTime
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
		sw   = store.NewWrapper(&kopia.ModelStore{})
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
				control.DefaultOptions(),
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
		t = suite.T()

		tenant        = "a-tenant"
		resourceOwner = "a-user"

		emailReason = kopia.NewReason(
			tenant,
			resourceOwner,
			path.ExchangeService,
			path.EmailCategory)
		contactsReason = kopia.NewReason(
			tenant,
			resourceOwner,
			path.ExchangeService,
			path.ContactsCategory)

		reasons = []identity.Reasoner{
			emailReason,
			contactsReason,
		}

		manifest1 = &snapshot.Manifest{
			ID: "id1",
		}
		manifest2 = &snapshot.Manifest{
			ID: "id2",
		}

		bases = kopia.NewMockBackupBases().WithMergeBases(
			kopia.ManifestEntry{
				Manifest: manifest1,
				Reasons: []identity.Reasoner{
					emailReason,
				},
			}).WithAssistBases(
			kopia.ManifestEntry{
				Manifest: manifest2,
				Reasons: []identity.Reasoner{
					contactsReason,
				},
			})

		backupID     = model.StableID("foo")
		expectedTags = map[string]string{
			kopia.TagBackupID:       string(backupID),
			kopia.TagBackupCategory: "",
		}
	)

	mbu := &mockBackupConsumer{
		checkFunc: func(
			backupReasons []identity.Reasoner,
			gotBases kopia.BackupBases,
			cs []data.BackupCollection,
			gotTags map[string]string,
			buildTreeWithBase bool,
		) {
			kopia.AssertBackupBasesEqual(t, bases, gotBases)
			assert.Equal(t, expectedTags, gotTags)
			assert.ElementsMatch(t, reasons, backupReasons)
		},
	}

	ctx, flush := tester.NewContext(t)
	defer flush()

	//nolint:errcheck
	consumeBackupCollections(
		ctx,
		mbu,
		tenant,
		reasons,
		bases,
		nil,
		nil,
		backupID,
		true,
		fault.New(true))
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

		time1 = time.Now()
		time2 = time1.Add(time.Hour)

		exchangeItemPath1 = makePath(
			suite.T(),
			[]string{
				tenant,
				path.ExchangeService.String(),
				ro,
				path.EmailCategory.String(),
				"work",
				"item1",
			},
			true)
		exchangeLocationPath1 = path.Builder{}.Append("work-display-name")
		exchangePathReason1   = kopia.NewReason(
			"",
			exchangeItemPath1.ResourceOwner(),
			exchangeItemPath1.Service(),
			exchangeItemPath1.Category())
	)

	itemParents1, err := path.GetDriveFolderPath(itemPath1)
	require.NoError(suite.T(), err, clues.ToCore(err))

	itemParents1String := itemParents1.String()

	table := []struct {
		name               string
		populatedDetails   map[string]*details.Details
		inputBackups       []kopia.BackupEntry
		inputAssistBackups []kopia.BackupEntry
		mdm                *mockDetailsMergeInfoer

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
					Reasons: []identity.Reasoner{
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
					Reasons: []identity.Reasoner{
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
					Reasons: []identity.Reasoner{
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
					Reasons: []identity.Reasoner{
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
					Reasons: []identity.Reasoner{
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
			name: "ExchangeItemMerged",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.add(exchangeItemPath1, exchangeItemPath1, exchangeLocationPath1)

				return res
			}(),
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup1,
					Reasons: []identity.Reasoner{
						exchangePathReason1,
					},
				},
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntry(suite.T(), exchangeItemPath1, exchangeLocationPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.Entry{
				makeDetailsEntry(suite.T(), exchangeItemPath1, exchangeLocationPath1, 42, false),
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
					Reasons: []identity.Reasoner{
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
					Reasons: []identity.Reasoner{
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
					Reasons: []identity.Reasoner{
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
					Reasons: []identity.Reasoner{
						pathReason1,
					},
				},
				{
					Backup: &backup2,
					Reasons: []identity.Reasoner{
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
		{
			name: "MergeAndAssistBases SameItems",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.addWithModTime(itemPath1, time1, itemPath1, locationPath1)
				res.addWithModTime(itemPath3, time2, itemPath3, locationPath3)

				return res
			}(),
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup1,
					Reasons: []identity.Reasoner{
						pathReason1,
						pathReason3,
					},
				},
			},
			inputAssistBackups: []kopia.BackupEntry{
				{Backup: &backup2},
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 42, false, time1),
							*makeDetailsEntryWithModTime(suite.T(), itemPath3, locationPath3, 37, false, time2),
						},
					},
				},
				backup2.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 42, false, time1),
							*makeDetailsEntryWithModTime(suite.T(), itemPath3, locationPath3, 37, false, time2),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.Entry{
				makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 42, false, time1),
				makeDetailsEntryWithModTime(suite.T(), itemPath3, locationPath3, 37, false, time2),
			},
		},
		{
			name: "MergeAndAssistBases AssistBaseHasNewerItems",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.addWithModTime(itemPath1, time2, itemPath1, locationPath1)

				return res
			}(),
			inputBackups: []kopia.BackupEntry{
				{
					Backup: &backup1,
					Reasons: []identity.Reasoner{
						pathReason1,
					},
				},
			},
			inputAssistBackups: []kopia.BackupEntry{
				{Backup: &backup2},
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 42, false, time1),
						},
					},
				},
				backup2.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 84, false, time2),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.Entry{
				makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 84, false, time2),
			},
		},
		{
			name: "AssistBases ConcurrentAssistBasesPicksMatchingVersion1",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.addWithModTime(itemPath1, time2, itemPath1, locationPath1)

				return res
			}(),
			inputAssistBackups: []kopia.BackupEntry{
				{Backup: &backup1},
				{Backup: &backup2},
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 42, false, time1),
						},
					},
				},
				backup2.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 84, false, time2),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.Entry{
				makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 84, false, time2),
			},
		},
		{
			name: "AssistBases ConcurrentAssistBasesPicksMatchingVersion2",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.addWithModTime(itemPath1, time1, itemPath1, locationPath1)

				return res
			}(),
			inputAssistBackups: []kopia.BackupEntry{
				{Backup: &backup1},
				{Backup: &backup2},
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 42, false, time1),
						},
					},
				},
				backup2.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 84, false, time2),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.Entry{
				makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 42, false, time1),
			},
		},
		{
			name: "AssistBases SameItemVersion",
			mdm: func() *mockDetailsMergeInfoer {
				res := newMockDetailsMergeInfoer()
				res.addWithModTime(itemPath1, time1, itemPath1, locationPath1)

				return res
			}(),
			inputAssistBackups: []kopia.BackupEntry{
				{Backup: &backup1},
				{Backup: &backup2},
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 42, false, time1),
						},
					},
				},
				backup2.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 42, false, time1),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.Entry{
				makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 42, false, time1),
			},
		},
		{
			name: "AssistBase ItemDeleted",
			mdm: func() *mockDetailsMergeInfoer {
				return newMockDetailsMergeInfoer()
			}(),
			inputAssistBackups: []kopia.BackupEntry{
				{Backup: &backup1},
			},
			populatedDetails: map[string]*details.Details{
				backup1.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.Entry{
							*makeDetailsEntryWithModTime(suite.T(), itemPath1, locationPath1, 42, false, time1),
						},
					},
				},
			},
			errCheck:        assert.NoError,
			expectedEntries: []*details.Entry{},
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

			bb := kopia.NewMockBackupBases().
				WithBackups(test.inputBackups...).
				WithAssistBackups(test.inputAssistBackups...)

			err := mergeDetails(
				ctx,
				mds,
				bb,
				test.mdm,
				&deets,
				&writeStats,
				path.OneDriveService,
				fault.New(true))
			test.errCheck(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			// Check the JSON output format of things because for some reason it's not
			// using the proper comparison for time.Time and failing due to that.
			checkJSONOutputs(t, test.expectedEntries, deets.Details().Items())
		})
	}
}

func checkJSONOutputs(
	t *testing.T,
	expected []*details.Entry,
	got []*details.Entry,
) {
	t.Helper()

	expectedJSON, err := json.Marshal(expected)
	require.NoError(t, err, "marshalling expected data")

	gotJSON, err := json.Marshal(got)
	require.NoError(t, err, "marshalling got data")

	assert.JSONEq(t, string(expectedJSON), string(gotJSON))
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
			Reasons: []identity.Reasoner{
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
		kopia.NewMockBackupBases().WithBackups(backup1),
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

	suite.ac, err = api.NewClient(creds, control.DefaultOptions())
	require.NoError(t, err, clues.ToCore(err))
}

func (suite *BackupOpIntegrationSuite) TestNewBackupOperation() {
	var (
		kw   = &kopia.Wrapper{}
		sw   = store.NewWrapper(&kopia.ModelStore{})
		ctrl = &mock.Controller{}
		acct = tconfig.NewM365Account(suite.T())
		opts = control.DefaultOptions()
	)

	table := []struct {
		name     string
		kw       *kopia.Wrapper
		sw       store.BackupStorer
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

type AssistBackupIntegrationSuite struct {
	tester.Suite
	kopiaCloser func(ctx context.Context)
	acct        account.Account
	kw          *kopia.Wrapper
	sw          store.BackupStorer
	ms          *kopia.ModelStore
}

func TestAssistBackupIntegrationSuite(t *testing.T) {
	suite.Run(t, &AssistBackupIntegrationSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{storeTD.AWSStorageCredEnvs, tconfig.M365AcctCredEnvs}),
	})
}

func (suite *AssistBackupIntegrationSuite) SetupSuite() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	var (
		st = storeTD.NewPrefixedS3Storage(t)
		k  = kopia.NewConn(st)
	)

	suite.acct = tconfig.NewM365Account(t)

	err := k.Initialize(ctx, repository.Options{}, repository.Retention{})
	require.NoError(t, err, clues.ToCore(err))

	suite.kopiaCloser = func(ctx context.Context) {
		k.Close(ctx)
	}

	kw, err := kopia.NewWrapper(k)
	require.NoError(t, err, clues.ToCore(err))

	suite.kw = kw

	ms, err := kopia.NewModelStore(k)
	require.NoError(t, err, clues.ToCore(err))

	suite.ms = ms

	sw := store.NewWrapper(ms)
	suite.sw = sw
}

func (suite *AssistBackupIntegrationSuite) TearDownSuite() {
	ctx, flush := tester.NewContext(suite.T())
	defer flush()

	if suite.ms != nil {
		suite.ms.Close(ctx)
	}

	if suite.kw != nil {
		suite.kw.Close(ctx)
	}

	if suite.kopiaCloser != nil {
		suite.kopiaCloser(ctx)
	}
}

var _ inject.BackupProducer = &mockBackupProducer{}

type mockBackupProducer struct {
	colls                   []data.BackupCollection
	dcs                     data.CollectionStats
	injectNonRecoverableErr bool
}

func (mbp *mockBackupProducer) ProduceBackupCollections(
	context.Context,
	inject.BackupProducerConfig,
	*fault.Bus,
) ([]data.BackupCollection, prefixmatcher.StringSetReader, bool, error) {
	if mbp.injectNonRecoverableErr {
		return nil, nil, false, clues.New("non-recoverable error")
	}

	return mbp.colls, nil, true, nil
}

func (mbp *mockBackupProducer) IsBackupRunnable(
	context.Context,
	path.ServiceType,
	string,
) (bool, error) {
	return true, nil
}

func (mbp *mockBackupProducer) Wait() *data.CollectionStats {
	return &mbp.dcs
}

func makeBackupCollection(
	p path.Path,
	locPath *path.Builder,
	items []dataMock.Item,
) data.BackupCollection {
	streams := make([]data.Item, len(items))

	for i := range items {
		streams[i] = &items[i]
	}

	return &mock.BackupCollection{
		Path:    p,
		Loc:     locPath,
		Streams: streams,
	}
}

func makeMetadataCollectionEntries(
	deltaURL, driveID, folderID string,
	p path.Path,
) []graph.MetadataCollectionEntry {
	return []graph.MetadataCollectionEntry{
		graph.NewMetadataEntry(
			graph.DeltaURLsFileName,
			map[string]string{driveID: deltaURL},
		),
		graph.NewMetadataEntry(
			graph.PreviousPathFileName,
			map[string]map[string]string{
				driveID: {
					folderID: p.PlainString(),
				},
			},
		),
	}
}

const (
	userID    = "user-id"
	driveID   = "drive-id"
	driveName = "drive-name"
	folderID  = "folder-id"
)

func makeMockItem(
	fileID string,
	extData *details.ExtensionData,
	modTime time.Time,
	del bool,
	readErr error,
) dataMock.Item {
	rc := odMock.FileRespReadCloser(odMock.DriveFilePayloadData)
	if extData != nil {
		rc = odMock.FileRespWithExtensions(odMock.DriveFilePayloadData, extData)
	}

	dmi := dataMock.Item{
		DeletedFlag:  del,
		ItemID:       fileID,
		ItemInfo:     odStub.DriveItemInfo(),
		ItemSize:     100,
		ModifiedTime: modTime,
		Reader:       rc,
		ReadErr:      readErr,
	}

	dmi.ItemInfo.OneDrive.DriveID = driveID
	dmi.ItemInfo.OneDrive.DriveName = driveName
	dmi.ItemInfo.OneDrive.Modified = modTime
	dmi.ItemInfo.Extension = extData

	return dmi
}

// Check what kind of backup is produced for a given failurePolicy/observed fault
// bus combination.
//
// It's currently using errors generated during mockBackupProducer phase.
// Ideally we would test with errors generated in various phases of backup, but
// that needs putting produceManifestsAndMetadata and mergeDetails behind mockable
// interfaces.
//
// Note: Tests are incremental since we are reusing kopia repo between tests,
// but this is irrelevant here.

func (suite *AssistBackupIntegrationSuite) TestBackupTypesForFailureModes() {
	var (
		acct     = tconfig.NewM365Account(suite.T())
		tenantID = acct.Config[config.AzureTenantIDKey]
		opts     = control.DefaultOptions()
		osel     = selectors.NewOneDriveBackup([]string{userID})
	)

	osel.Include(selTD.OneDriveBackupFolderScope(osel))

	pathElements := []string{odConsts.DrivesPathDir, "drive-id", odConsts.RootPathDir, folderID}

	tmp, err := path.Build(tenantID, userID, path.OneDriveService, path.FilesCategory, false, pathElements...)
	require.NoError(suite.T(), err, clues.ToCore(err))

	locPath := path.Builder{}.Append(tmp.Folders()...)

	table := []struct {
		name                    string
		collFunc                func() []data.BackupCollection
		injectNonRecoverableErr bool
		failurePolicy           control.FailurePolicy
		expectRunErr            assert.ErrorAssertionFunc
		expectBackupTag         string
		expectFaults            func(t *testing.T, errs *fault.Bus)
	}{
		{
			name: "fail fast, no errors",
			collFunc: func() []data.BackupCollection {
				bc := []data.BackupCollection{
					makeBackupCollection(
						tmp,
						locPath,
						[]dataMock.Item{
							makeMockItem("file1", nil, time.Now(), false, nil),
						}),
				}

				return bc
			},
			failurePolicy:   control.FailFast,
			expectRunErr:    assert.NoError,
			expectBackupTag: model.MergeBackup,
			expectFaults: func(t *testing.T, errs *fault.Bus) {
				assert.NoError(t, errs.Failure(), clues.ToCore(errs.Failure()))
				assert.Empty(t, errs.Recovered(), "recovered errors")
			},
		},
		{
			name: "fail fast, any errors",
			collFunc: func() []data.BackupCollection {
				bc := []data.BackupCollection{
					makeBackupCollection(
						tmp,
						locPath,
						[]dataMock.Item{
							makeMockItem("file1", nil, time.Now(), false, assert.AnError),
						}),
				}
				return bc
			},
			failurePolicy:   control.FailFast,
			expectRunErr:    assert.Error,
			expectBackupTag: "",
			expectFaults: func(t *testing.T, errs *fault.Bus) {
				assert.Error(t, errs.Failure(), clues.ToCore(errs.Failure()))
			},
		},
		{
			name: "best effort, no errors",
			collFunc: func() []data.BackupCollection {
				bc := []data.BackupCollection{
					makeBackupCollection(
						tmp,
						locPath,
						[]dataMock.Item{
							makeMockItem("file1", nil, time.Now(), false, nil),
						}),
				}

				return bc
			},
			failurePolicy:   control.BestEffort,
			expectRunErr:    assert.NoError,
			expectBackupTag: model.MergeBackup,
			expectFaults: func(t *testing.T, errs *fault.Bus) {
				assert.NoError(t, errs.Failure(), clues.ToCore(errs.Failure()))
				assert.Empty(t, errs.Recovered(), "recovered errors")
			},
		},
		{
			name: "best effort, non-recoverable errors",
			collFunc: func() []data.BackupCollection {
				return nil
			},
			injectNonRecoverableErr: true,
			failurePolicy:           control.BestEffort,
			expectRunErr:            assert.Error,
			expectBackupTag:         "",
			expectFaults: func(t *testing.T, errs *fault.Bus) {
				assert.Error(t, errs.Failure(), clues.ToCore(errs.Failure()))
			},
		},
		{
			name: "best effort, recoverable errors",
			collFunc: func() []data.BackupCollection {
				bc := []data.BackupCollection{
					makeBackupCollection(
						tmp,
						locPath,
						[]dataMock.Item{
							makeMockItem("file1", nil, time.Now(), false, assert.AnError),
						}),
				}

				return bc
			},
			failurePolicy:   control.BestEffort,
			expectRunErr:    assert.NoError,
			expectBackupTag: model.MergeBackup,
			expectFaults: func(t *testing.T, errs *fault.Bus) {
				assert.NoError(t, errs.Failure(), clues.ToCore(errs.Failure()))
				assert.Greater(t, len(errs.Recovered()), 0, "recovered errors")
			},
		},
		{
			name: "fail after recovery, no errors",
			collFunc: func() []data.BackupCollection {
				bc := []data.BackupCollection{
					makeBackupCollection(
						tmp,
						locPath,
						[]dataMock.Item{
							makeMockItem("file1", nil, time.Now(), false, nil),
							makeMockItem("file2", nil, time.Now(), false, nil),
						}),
				}

				return bc
			},
			failurePolicy:   control.FailAfterRecovery,
			expectRunErr:    assert.NoError,
			expectBackupTag: model.MergeBackup,
			expectFaults: func(t *testing.T, errs *fault.Bus) {
				assert.NoError(t, errs.Failure(), clues.ToCore(errs.Failure()))
				assert.Empty(t, errs.Recovered(), "recovered errors")
			},
		},
		{
			name: "fail after recovery, non-recoverable errors",
			collFunc: func() []data.BackupCollection {
				return nil
			},
			injectNonRecoverableErr: true,
			failurePolicy:           control.FailAfterRecovery,
			expectRunErr:            assert.Error,
			expectBackupTag:         "",
			expectFaults: func(t *testing.T, errs *fault.Bus) {
				assert.Error(t, errs.Failure(), clues.ToCore(errs.Failure()))
			},
		},
		{
			name: "fail after recovery, recoverable errors",
			collFunc: func() []data.BackupCollection {
				bc := []data.BackupCollection{
					makeBackupCollection(
						tmp,
						locPath,
						[]dataMock.Item{
							makeMockItem("file1", nil, time.Now(), false, nil),
							makeMockItem("file2", nil, time.Now(), false, assert.AnError),
						}),
				}

				return bc
			},
			failurePolicy:   control.FailAfterRecovery,
			expectRunErr:    assert.Error,
			expectBackupTag: model.AssistBackup,
			expectFaults: func(t *testing.T, errs *fault.Bus) {
				assert.Error(t, errs.Failure(), clues.ToCore(errs.Failure()))
				assert.Greater(t, len(errs.Recovered()), 0, "recovered errors")
			},
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			cs := test.collFunc()

			pathPrefix, err := path.Builder{}.ToServiceCategoryMetadataPath(
				tenantID,
				userID,
				path.OneDriveService,
				path.FilesCategory,
				false)
			require.NoError(t, err, clues.ToCore(err))

			mc, err := graph.MakeMetadataCollection(
				pathPrefix,
				makeMetadataCollectionEntries("url/1", driveID, folderID, tmp),
				func(*support.ControllerOperationStatus) {})
			require.NoError(t, err, clues.ToCore(err))

			cs = append(cs, mc)
			bp := &mockBackupProducer{
				colls:                   cs,
				injectNonRecoverableErr: test.injectNonRecoverableErr,
			}

			opts.FailureHandling = test.failurePolicy

			bo, err := NewBackupOperation(
				ctx,
				opts,
				suite.kw,
				suite.sw,
				bp,
				acct,
				osel.Selector,
				selectors.Selector{DiscreteOwner: userID},
				evmock.NewBus())
			require.NoError(t, err, clues.ToCore(err))

			err = bo.Run(ctx)
			test.expectRunErr(t, err, clues.ToCore(err))
			test.expectFaults(t, bo.Errors)

			if len(test.expectBackupTag) == 0 {
				return
			}

			bID := bo.Results.BackupID
			require.NotEmpty(t, bID)

			bup := backup.Backup{}

			err = suite.ms.Get(ctx, model.BackupSchema, bID, &bup)
			require.NoError(t, err, clues.ToCore(err))

			require.Equal(t, test.expectBackupTag, bup.Tags[model.BackupTypeTag])
		})
	}
}

func selectFilesFromDeets(d details.Details) map[string]details.Entry {
	files := make(map[string]details.Entry)

	for _, ent := range d.Entries {
		if ent.Folder != nil {
			continue
		}

		files[ent.ItemRef] = ent
	}

	return files
}

// TestExtensionsIncrementals tests presence of corso extension data in details
// Note that since we are mocking out backup producer here, corso extensions can't be
// attached as they would in prod. However, this is fine here, since we are more interested
// in testing whether deets get carried over correctly for various scenarios.
func (suite *AssistBackupIntegrationSuite) TestExtensionsIncrementals() {
	var (
		acct     = tconfig.NewM365Account(suite.T())
		tenantID = acct.Config[config.AzureTenantIDKey]
		opts     = control.DefaultOptions()
		osel     = selectors.NewOneDriveBackup([]string{userID})
		// Default policy used by SDK clients
		failurePolicy = control.FailAfterRecovery
		T1            = time.Now().Truncate(0)
		T2            = T1.Add(time.Hour).Truncate(0)
		T3            = T2.Add(time.Hour).Truncate(0)
		extData       = make(map[int]*details.ExtensionData)
	)

	for i := 0; i < 3; i++ {
		d := make(map[string]any)
		extData[i] = &details.ExtensionData{
			Data: d,
		}
	}

	osel.Include(selTD.OneDriveBackupFolderScope(osel))

	sss := streamstore.NewStreamer(
		suite.kw,
		suite.acct.ID(),
		osel.PathService())

	pathElements := []string{odConsts.DrivesPathDir, "drive-id", odConsts.RootPathDir, folderID}

	tmp, err := path.Build(tenantID, userID, path.OneDriveService, path.FilesCategory, false, pathElements...)
	require.NoError(suite.T(), err, clues.ToCore(err))

	locPath := path.Builder{}.Append(tmp.Folders()...)

	table := []struct {
		name          string
		collFunc      func() []data.BackupCollection
		expectRunErr  assert.ErrorAssertionFunc
		validateDeets func(t *testing.T, gotDeets details.Details)
	}{
		{
			name: "Assist backup, 1 new deets",
			collFunc: func() []data.BackupCollection {
				bc := []data.BackupCollection{
					makeBackupCollection(
						tmp,
						locPath,
						[]dataMock.Item{
							makeMockItem("file1", extData[0], T1, false, nil),
							makeMockItem("file2", extData[1], T1, false, assert.AnError),
						}),
				}

				return bc
			},
			expectRunErr: assert.Error,
			validateDeets: func(t *testing.T, d details.Details) {
				files := selectFilesFromDeets(d)
				require.Len(t, files, 1)

				f := files["file1"]
				require.NotNil(t, f)

				require.True(t, T1.Equal(f.Modified()))
				require.NotZero(t, f.Extension.Data[extensions.KNumBytes])
			},
		},
		{
			name: "Assist backup after assist backup, 1 existing, 1 new deets",
			collFunc: func() []data.BackupCollection {
				bc := []data.BackupCollection{
					makeBackupCollection(
						tmp,
						locPath,
						[]dataMock.Item{
							makeMockItem("file1", extData[0], T1, false, nil),
							makeMockItem("file2", extData[1], T2, false, nil),
							makeMockItem("file3", extData[2], T2, false, assert.AnError),
						}),
				}

				return bc
			},
			expectRunErr: assert.Error,
			validateDeets: func(t *testing.T, d details.Details) {
				files := selectFilesFromDeets(d)
				require.Len(t, files, 2)

				for _, f := range files {
					switch f.ItemRef {
					case "file1":
						require.True(t, T1.Equal(f.Modified()))
						require.NotZero(t, f.Extension.Data[extensions.KNumBytes])
					case "file2":
						require.True(t, T2.Equal(f.Modified()))
						require.NotZero(t, f.Extension.Data[extensions.KNumBytes])
					default:
						require.Fail(t, "unexpected file", f.ItemRef)
					}
				}
			},
		},
		{
			name: "Merge backup, 2 existing deets, 1 new deet",
			collFunc: func() []data.BackupCollection {
				bc := []data.BackupCollection{
					makeBackupCollection(
						tmp,
						locPath,
						[]dataMock.Item{
							makeMockItem("file1", extData[0], T1, false, nil),
							makeMockItem("file2", extData[1], T2, false, nil),
							makeMockItem("file3", extData[2], T3, false, nil),
						}),
				}

				return bc
			},
			expectRunErr: assert.NoError,
			validateDeets: func(t *testing.T, d details.Details) {
				files := selectFilesFromDeets(d)
				require.Len(t, files, 3)

				for _, f := range files {
					switch f.ItemRef {
					case "file1":
						require.True(t, T1.Equal(f.Modified()))
						require.NotZero(t, f.Extension.Data[extensions.KNumBytes])
					case "file2":
						require.True(t, T2.Equal(f.Modified()))
						require.NotZero(t, f.Extension.Data[extensions.KNumBytes])
					case "file3":
						require.True(t, T3.Equal(f.Modified()))
						require.NotZero(t, f.Extension.Data[extensions.KNumBytes])
					default:
						require.Fail(t, "unexpected file", f.ItemRef)
					}
				}
			},
		},
		{
			// Reset state so we can reuse the same test data
			name: "All files deleted",
			collFunc: func() []data.BackupCollection {
				bc := []data.BackupCollection{
					makeBackupCollection(
						tmp,
						locPath,
						[]dataMock.Item{
							makeMockItem("file1", extData[0], T1, true, nil),
							makeMockItem("file2", extData[1], T2, true, nil),
							makeMockItem("file3", extData[2], T3, true, nil),
						}),
				}

				return bc
			},
			expectRunErr: assert.NoError,
			validateDeets: func(t *testing.T, d details.Details) {
				files := selectFilesFromDeets(d)
				require.Len(t, files, 0)
			},
		},
		{
			name: "Merge backup, 1 new deets",
			collFunc: func() []data.BackupCollection {
				bc := []data.BackupCollection{
					makeBackupCollection(
						tmp,
						locPath,
						[]dataMock.Item{
							makeMockItem("file1", extData[0], T1, false, nil),
						}),
				}

				return bc
			},
			expectRunErr: assert.NoError,
			validateDeets: func(t *testing.T, d details.Details) {
				files := selectFilesFromDeets(d)
				require.Len(t, files, 1)

				for _, f := range files {
					switch f.ItemRef {
					case "file1":
						require.True(t, T1.Equal(f.Modified()))
						require.NotZero(t, f.Extension.Data[extensions.KNumBytes])
					default:
						require.Fail(t, "unexpected file", f.ItemRef)
					}
				}
			},
		},
		// This test fails currently, need to rerun with Ashlie's PR.
		{
			name: "Assist backup after merge backup, 1 new deets, 1 existing deet",
			collFunc: func() []data.BackupCollection {
				bc := []data.BackupCollection{
					makeBackupCollection(
						tmp,
						locPath,
						[]dataMock.Item{
							makeMockItem("file1", extData[0], T1, false, nil),
							makeMockItem("file2", extData[1], T2, false, nil),
							makeMockItem("file3", extData[2], T3, false, assert.AnError),
						}),
				}

				return bc
			},
			expectRunErr: assert.Error,
			validateDeets: func(t *testing.T, d details.Details) {
				files := selectFilesFromDeets(d)
				require.Len(t, files, 2)

				for _, f := range files {
					switch f.ItemRef {
					case "file1":
						require.True(t, T1.Equal(f.Modified()))
						require.NotZero(t, f.Extension.Data[extensions.KNumBytes])

					case "file2":
						require.True(t, T2.Equal(f.Modified()))
						require.NotZero(t, f.Extension.Data[extensions.KNumBytes])
					default:
						require.Fail(t, "unexpected file", f.ItemRef)
					}
				}
			},
		},

		// TODO(pandeyabs): Remaining tests.
		// 1. Deets updated in assist backup. Following backup should have updated deets.
		// 2. Concurrent overlapping reasons.
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			cs := test.collFunc()

			pathPrefix, err := path.Builder{}.ToServiceCategoryMetadataPath(
				tenantID,
				userID,
				path.OneDriveService,
				path.FilesCategory,
				false)
			require.NoError(t, err, clues.ToCore(err))

			mc, err := graph.MakeMetadataCollection(
				pathPrefix,
				makeMetadataCollectionEntries("url/1", driveID, folderID, tmp),
				func(*support.ControllerOperationStatus) {})
			require.NoError(t, err, clues.ToCore(err))

			cs = append(cs, mc)
			bp := &mockBackupProducer{
				colls: cs,
			}

			opts.FailureHandling = failurePolicy

			bo, err := NewBackupOperation(
				ctx,
				opts,
				suite.kw,
				suite.sw,
				bp,
				acct,
				osel.Selector,
				selectors.Selector{DiscreteOwner: userID},
				evmock.NewBus())
			require.NoError(t, err, clues.ToCore(err))

			err = bo.Run(ctx)
			test.expectRunErr(t, err, clues.ToCore(err))

			assert.NotEmpty(t, bo.Results.BackupID)

			deets, _ := deeTD.GetDeetsInBackup(
				t,
				ctx,
				bo.Results.BackupID,
				tenantID,
				userID,
				path.OneDriveService,
				deeTD.DriveIDFromRepoRef,
				suite.ms,
				sss)
			assert.NotNil(t, deets)

			test.validateDeets(t, deets)

			// Clear extension data between test runs
			for i := 0; i < 3; i++ {
				d := make(map[string]any)
				extData[i] = &details.ExtensionData{
					Data: d,
				}
			}
		})
	}
}
