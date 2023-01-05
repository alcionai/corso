package operations

import (
	"context"
	stdpath "path"
	"testing"
	"time"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/kopia"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

// ---------------------------------------------------------------------------
// mocks
// ---------------------------------------------------------------------------

// ----- restore producer

type mockRestorer struct {
	gotPaths []path.Path
}

func (mr *mockRestorer) RestoreMultipleItems(
	ctx context.Context,
	snapshotID string,
	paths []path.Path,
	bc kopia.ByteCounter,
) ([]data.Collection, error) {
	mr.gotPaths = append(mr.gotPaths, paths...)

	return nil, nil
}

func (mr mockRestorer) checkPaths(t *testing.T, expected []path.Path) {
	t.Helper()

	assert.ElementsMatch(t, expected, mr.gotPaths)
}

// ----- backup producer

type mockBackuper struct {
	checkFunc func(
		bases []kopia.IncrementalBase,
		cs []data.Collection,
		service path.ServiceType,
		oc *kopia.OwnersCats,
		tags map[string]string,
		buildTreeWithBase bool,
	)
}

func (mbu mockBackuper) BackupCollections(
	ctx context.Context,
	bases []kopia.IncrementalBase,
	cs []data.Collection,
	service path.ServiceType,
	oc *kopia.OwnersCats,
	tags map[string]string,
	buildTreeWithBase bool,
) (*kopia.BackupStats, *details.Builder, map[string]path.Path, error) {
	if mbu.checkFunc != nil {
		mbu.checkFunc(bases, cs, service, oc, tags, buildTreeWithBase)
	}

	return &kopia.BackupStats{}, &details.Builder{}, nil, nil
}

// ----- details

type mockDetailsReader struct {
	entries map[string]*details.Details
}

func (mdr mockDetailsReader) ReadBackupDetails(
	ctx context.Context,
	detailsID string,
) (*details.Details, error) {
	r := mdr.entries[detailsID]

	if r == nil {
		return nil, errors.Errorf("no details for ID %s", detailsID)
	}

	return r, nil
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
	if s != model.BackupSchema {
		return errors.Errorf("unexpected schema %s", s)
	}

	r, ok := mbs.entries[id]
	if !ok {
		return errors.Errorf("model with id %s not found", id)
	}

	bu, ok := toPopulate.(*backup.Backup)
	if !ok {
		return errors.Errorf("bad input type %T", toPopulate)
	}

	*bu = r

	return nil
}

func (mbs mockBackupStorer) Delete(context.Context, model.Schema, model.StableID) error {
	return errors.New("not implemented")
}

func (mbs mockBackupStorer) DeleteWithModelStoreID(context.Context, manifest.ID) error {
	return errors.New("not implemented")
}

func (mbs mockBackupStorer) GetIDsForType(
	context.Context,
	model.Schema,
	map[string]string,
) ([]*model.BaseModel, error) {
	return nil, errors.New("not implemented")
}

func (mbs mockBackupStorer) GetWithModelStoreID(
	context.Context,
	model.Schema,
	manifest.ID,
	model.Model,
) error {
	return errors.New("not implemented")
}

func (mbs mockBackupStorer) Put(context.Context, model.Schema, model.Model) error {
	return errors.New("not implemented")
}

func (mbs mockBackupStorer) Update(context.Context, model.Schema, model.Model) error {
	return errors.New("not implemented")
}

// ---------------------------------------------------------------------------
// helper funcs
// ---------------------------------------------------------------------------

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
		true,
	)
	require.NoError(t, err)

	return p
}

func makeFolderEntry(
	t *testing.T,
	pb *path.Builder,
	size int64,
	modTime time.Time,
) *details.DetailsEntry {
	t.Helper()

	return &details.DetailsEntry{
		RepoRef:   pb.String(),
		ShortRef:  pb.ShortRef(),
		ParentRef: pb.Dir().ShortRef(),
		ItemInfo: details.ItemInfo{
			Folder: &details.FolderInfo{
				ItemType:    details.FolderItem,
				DisplayName: pb.Elements()[len(pb.Elements())-1],
				Modified:    modTime,
				Size:        size,
			},
		},
	}
}

// TODO(ashmrtn): Really need to factor a function like this out into some
// common file that is only compiled for tests.
func makePath(t *testing.T, elements []string, isItem bool) path.Path {
	t.Helper()

	p, err := path.FromDataLayerPath(stdpath.Join(elements...), isItem)
	require.NoError(t, err)

	return p
}

func makeDetailsEntry(
	t *testing.T,
	p path.Path,
	size int,
	updated bool,
) *details.DetailsEntry {
	t.Helper()

	res := &details.DetailsEntry{
		RepoRef:   p.String(),
		ShortRef:  p.ShortRef(),
		ParentRef: p.ToBuilder().Dir().ShortRef(),
		ItemInfo:  details.ItemInfo{},
		Updated:   updated,
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
			ItemType: details.ExchangeMail,
			Size:     int64(size),
		}

	case path.OneDriveService:
		parent, err := path.GetDriveFolderPath(p)
		require.NoError(t, err)

		res.OneDrive = &details.OneDriveInfo{
			ItemType:   details.OneDriveItem,
			ParentPath: parent,
			Size:       int64(size),
		}

	default:
		assert.FailNowf(
			t,
			"service %s not supported in helper function",
			p.Service().String(),
		)
	}

	return res
}

func makeManifest(t *testing.T, backupID model.StableID, incompleteReason string) *snapshot.Manifest {
	t.Helper()

	backupIDTagKey, _ := kopia.MakeTagKV(kopia.TagBackupID)

	return &snapshot.Manifest{
		Tags: map[string]string{
			backupIDTagKey: string(backupID),
		},
		IncompleteReason: incompleteReason,
	}
}

// ---------------------------------------------------------------------------
// unit tests
// ---------------------------------------------------------------------------

type BackupOpSuite struct {
	suite.Suite
}

func TestBackupOpSuite(t *testing.T) {
	suite.Run(t, new(BackupOpSuite))
}

func (suite *BackupOpSuite) TestBackupOperation_PersistResults() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		kw   = &kopia.Wrapper{}
		sw   = &store.Wrapper{}
		acct = account.Account{}
		now  = time.Now()
	)

	table := []struct {
		expectStatus opStatus
		expectErr    assert.ErrorAssertionFunc
		stats        backupStats
	}{
		{
			expectStatus: Completed,
			expectErr:    assert.NoError,
			stats: backupStats{
				started:       true,
				resourceCount: 1,
				k: &kopia.BackupStats{
					TotalFileCount:     1,
					TotalHashedBytes:   1,
					TotalUploadedBytes: 1,
				},
				gc: &support.ConnectorOperationStatus{
					Successful: 1,
				},
			},
		},
		{
			expectStatus: Failed,
			expectErr:    assert.Error,
			stats: backupStats{
				started: false,
				k:       &kopia.BackupStats{},
				gc:      &support.ConnectorOperationStatus{},
			},
		},
		{
			expectStatus: NoData,
			expectErr:    assert.NoError,
			stats: backupStats{
				started: true,
				k:       &kopia.BackupStats{},
				gc:      &support.ConnectorOperationStatus{},
			},
		},
	}
	for _, test := range table {
		suite.T().Run(test.expectStatus.String(), func(t *testing.T) {
			sel := selectors.Selector{}
			sel.DiscreteOwner = "bombadil"

			op, err := NewBackupOperation(
				ctx,
				control.Options{},
				kw,
				sw,
				acct,
				sel,
				evmock.NewBus())
			require.NoError(t, err)
			test.expectErr(t, op.persistResults(now, &test.stats))

			assert.Equal(t, test.expectStatus.String(), op.Status.String(), "status")
			assert.Equal(t, test.stats.gc.Successful, op.Results.ItemsRead, "items read")
			assert.Equal(t, test.stats.readErr, op.Results.ReadErrors, "read errors")
			assert.Equal(t, test.stats.k.TotalFileCount, op.Results.ItemsWritten, "items written")
			assert.Equal(t, test.stats.k.TotalHashedBytes, op.Results.BytesRead, "bytes read")
			assert.Equal(t, test.stats.k.TotalUploadedBytes, op.Results.BytesUploaded, "bytes written")
			assert.Equal(t, test.stats.resourceCount, op.Results.ResourceOwners, "resource owners")
			assert.Equal(t, test.stats.writeErr, op.Results.WriteErrors, "write errors")
			assert.Equal(t, now, op.Results.StartedAt, "started at")
			assert.Less(t, now, op.Results.CompletedAt, "completed at")
		})
	}
}

func (suite *BackupOpSuite) TestBackupOperation_CollectMetadata() {
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
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			mr := &mockRestorer{}

			_, err := collectMetadata(ctx, mr, test.inputMan, test.inputFiles, tenant)
			assert.NoError(t, err)

			mr.checkPaths(t, test.expected)
		})
	}
}

func (suite *BackupOpSuite) TestBackupOperation_ConsumeBackupDataCollections_Paths() {
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

		sel = selectors.NewExchangeBackup([]string{resourceOwner}).Selector
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
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			mbu := &mockBackuper{
				checkFunc: func(
					bases []kopia.IncrementalBase,
					cs []data.Collection,
					service path.ServiceType,
					oc *kopia.OwnersCats,
					tags map[string]string,
					buildTreeWithBase bool,
				) {
					assert.ElementsMatch(t, test.expected, bases)
				},
			}

			//nolint:errcheck
			consumeBackupDataCollections(
				ctx,
				mbu,
				tenant,
				sel,
				nil,
				test.inputMan,
				nil,
				model.StableID(""),
				true,
			)
		})
	}
}

func (suite *BackupOpSuite) TestBackupOperation_MergeBackupDetails_AddsItems() {
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
		itemPath2 = makePath(
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
		itemPath3 = makePath(
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
	require.NoError(suite.T(), err)

	table := []struct {
		name                         string
		populatedModels              map[model.StableID]backup.Backup
		populatedDetails             map[string]*details.Details
		inputMans                    []*kopia.ManifestEntry
		inputShortRefsFromPrevBackup map[string]path.Path

		errCheck        assert.ErrorAssertionFunc
		expectedEntries []*details.DetailsEntry
	}{
		{
			name:     "NilShortRefsFromPrevBackup",
			errCheck: assert.NoError,
			// Use empty slice so we don't error out on nil != empty.
			expectedEntries: []*details.DetailsEntry{},
		},
		{
			name:                         "EmptyShortRefsFromPrevBackup",
			inputShortRefsFromPrevBackup: map[string]path.Path{},
			errCheck:                     assert.NoError,
			// Use empty slice so we don't error out on nil != empty.
			expectedEntries: []*details.DetailsEntry{},
		},
		{
			name: "BackupIDNotFound",
			inputShortRefsFromPrevBackup: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
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
			inputShortRefsFromPrevBackup: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
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
			inputShortRefsFromPrevBackup: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
				itemPath2.ShortRef(): itemPath2,
			},
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
						Entries: []details.DetailsEntry{
							*makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "TooManyItems",
			inputShortRefsFromPrevBackup: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
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
						Entries: []details.DetailsEntry{
							*makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "BadBaseRepoRef",
			inputShortRefsFromPrevBackup: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
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
						Entries: []details.DetailsEntry{
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
			inputShortRefsFromPrevBackup: map[string]path.Path{
				itemPath1.ShortRef(): makePath(
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
				),
			},
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
						Entries: []details.DetailsEntry{
							*makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "ItemMerged",
			inputShortRefsFromPrevBackup: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
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
						Entries: []details.DetailsEntry{
							*makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.DetailsEntry{
				makeDetailsEntry(suite.T(), itemPath1, 42, false),
			},
		},
		{
			name: "ItemMergedExtraItemsInBase",
			inputShortRefsFromPrevBackup: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
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
						Entries: []details.DetailsEntry{
							*makeDetailsEntry(suite.T(), itemPath1, 42, false),
							*makeDetailsEntry(suite.T(), itemPath2, 84, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.DetailsEntry{
				makeDetailsEntry(suite.T(), itemPath1, 42, false),
			},
		},
		{
			name: "ItemMoved",
			inputShortRefsFromPrevBackup: map[string]path.Path{
				itemPath1.ShortRef(): itemPath2,
			},
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
						Entries: []details.DetailsEntry{
							*makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.DetailsEntry{
				makeDetailsEntry(suite.T(), itemPath2, 42, true),
			},
		},
		{
			name: "MultipleBases",
			inputShortRefsFromPrevBackup: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
				itemPath3.ShortRef(): itemPath3,
			},
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
						Entries: []details.DetailsEntry{
							*makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
				backup2.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.DetailsEntry{
							// This entry should not be picked due to a mismatch on Reasons.
							*makeDetailsEntry(suite.T(), itemPath1, 84, false),
							// This item should be picked.
							*makeDetailsEntry(suite.T(), itemPath3, 37, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.DetailsEntry{
				makeDetailsEntry(suite.T(), itemPath1, 42, false),
				makeDetailsEntry(suite.T(), itemPath3, 37, false),
			},
		},
		{
			name: "SomeBasesIncomplete",
			inputShortRefsFromPrevBackup: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
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
						Entries: []details.DetailsEntry{
							*makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
				backup2.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.DetailsEntry{
							// This entry should not be picked due to being incomplete.
							*makeDetailsEntry(suite.T(), itemPath1, 84, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []*details.DetailsEntry{
				makeDetailsEntry(suite.T(), itemPath1, 42, false),
			},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			mdr := mockDetailsReader{entries: test.populatedDetails}
			w := &store.Wrapper{Storer: mockBackupStorer{entries: test.populatedModels}}

			deets := details.Builder{}

			err := mergeDetails(
				ctx,
				w,
				mdr,
				test.inputMans,
				test.inputShortRefsFromPrevBackup,
				&deets,
			)

			test.errCheck(t, err)
			if err != nil {
				return
			}

			assert.ElementsMatch(t, test.expectedEntries, deets.Details().Items())
		})
	}
}

func (suite *BackupOpSuite) TestBackupOperation_MergeBackupDetails_AddsFolders() {
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
			"item1",
		}

		itemPath1 = makePath(
			t,
			pathElems,
			true,
		)

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

		inputToMerge = map[string]path.Path{
			itemPath1.ShortRef(): itemPath1,
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

		itemSize    = 42
		itemDetails = makeDetailsEntry(t, itemPath1, itemSize, false)

		populatedDetails = map[string]*details.Details{
			backup1.DetailsID: {
				DetailsModel: details.DetailsModel{
					Entries: []details.DetailsEntry{
						*itemDetails,
					},
				},
			},
		}

		expectedEntries = []details.DetailsEntry{
			*itemDetails,
		}
	)

	itemDetails.Exchange.Modified = time.Now()

	for i := 1; i < len(pathElems); i++ {
		expectedEntries = append(expectedEntries, *makeFolderEntry(
			t,
			path.Builder{}.Append(pathElems[:i]...),
			int64(itemSize),
			itemDetails.Exchange.Modified,
		))
	}

	ctx, flush := tester.NewContext()
	defer flush()

	mdr := mockDetailsReader{entries: populatedDetails}
	w := &store.Wrapper{Storer: mockBackupStorer{entries: populatedModels}}

	deets := details.Builder{}

	err := mergeDetails(
		ctx,
		w,
		mdr,
		inputMans,
		inputToMerge,
		&deets,
	)

	assert.NoError(t, err)
	assert.ElementsMatch(t, expectedEntries, deets.Details().Entries)
}
