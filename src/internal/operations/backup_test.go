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

	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/events"
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
// unit
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
			op, err := NewBackupOperation(
				ctx,
				control.Options{},
				kw,
				sw,
				acct,
				selectors.Selector{},
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

func makeMetadataPath(
	t *testing.T,
	tenant string,
	service path.ServiceType,
	resourceOwner string,
	category path.CategoryType,
	fileName string,
) path.Path {
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

type mockBackuper struct {
	checkFunc func(
		bases []kopia.IncrementalBase,
		cs []data.Collection,
		service path.ServiceType,
		oc *kopia.OwnersCats,
		tags map[string]string,
	)
}

func (mbu mockBackuper) BackupCollections(
	ctx context.Context,
	bases []kopia.IncrementalBase,
	cs []data.Collection,
	service path.ServiceType,
	oc *kopia.OwnersCats,
	tags map[string]string,
) (*kopia.BackupStats, *details.Details, map[string]path.Path, error) {
	if mbu.checkFunc != nil {
		mbu.checkFunc(bases, cs, service, oc, tags)
	}

	return &kopia.BackupStats{}, &details.Details{}, nil, nil
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
			)
		})
	}
}

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

// Functions we need to implement but don't care about.
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

// TODO(ashmrtn): Really need to factor a function like this out into some
// common file that is only compiled for tests.
func makePath(
	t *testing.T,
	elements []string,
	isItem bool,
) path.Path {
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
) details.DetailsEntry {
	t.Helper()

	res := details.DetailsEntry{
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

func makeManifest(backupID model.StableID, incompleteReason string) *snapshot.Manifest {
	backupIDTagKey, _ := kopia.MakeTagKV(kopia.TagBackupID)

	return &snapshot.Manifest{
		Tags: map[string]string{
			backupIDTagKey: string(backupID),
		},
		IncompleteReason: incompleteReason,
	}
}

func (suite *BackupOpSuite) TestBackupOperation_MergeBackupDetails() {
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
		name             string
		populatedModels  map[model.StableID]backup.Backup
		populatedDetails map[string]*details.Details
		inputMans        []*kopia.ManifestEntry
		inputToMerge     map[string]path.Path

		errCheck        assert.ErrorAssertionFunc
		expectedEntries []details.DetailsEntry
	}{
		{
			name:     "NilToMerge",
			errCheck: assert.NoError,
			// Use empty slice so we don't error out on nil != empty.
			expectedEntries: []details.DetailsEntry{},
		},
		{
			name:         "EmptyToMerge",
			inputToMerge: map[string]path.Path{},
			errCheck:     assert.NoError,
			// Use empty slice so we don't error out on nil != empty.
			expectedEntries: []details.DetailsEntry{},
		},
		{
			name: "BackupIDNotFound",
			inputToMerge: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest("foo", ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "DetailsIDNotFound",
			inputToMerge: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(backup1.ID, ""),
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
			inputToMerge: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
				itemPath2.ShortRef(): itemPath2,
			},
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(backup1.ID, ""),
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
							makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "TooManyItems",
			inputToMerge: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
				{
					Manifest: makeManifest(backup1.ID, ""),
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
							makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "BadBaseRepoRef",
			inputToMerge: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(backup1.ID, ""),
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
			inputToMerge: map[string]path.Path{
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
					Manifest: makeManifest(backup1.ID, ""),
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
							makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.Error,
		},
		{
			name: "ItemMerged",
			inputToMerge: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(backup1.ID, ""),
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
							makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []details.DetailsEntry{
				makeDetailsEntry(suite.T(), itemPath1, 42, false),
			},
		},
		{
			name: "ItemMergedExtraItemsInBase",
			inputToMerge: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(backup1.ID, ""),
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
							makeDetailsEntry(suite.T(), itemPath1, 42, false),
							makeDetailsEntry(suite.T(), itemPath2, 84, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []details.DetailsEntry{
				makeDetailsEntry(suite.T(), itemPath1, 42, false),
			},
		},
		{
			name: "ItemMoved",
			inputToMerge: map[string]path.Path{
				itemPath1.ShortRef(): itemPath2,
			},
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(backup1.ID, ""),
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
							makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []details.DetailsEntry{
				makeDetailsEntry(suite.T(), itemPath2, 42, true),
			},
		},
		{
			name: "MultipleBases",
			inputToMerge: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
				itemPath3.ShortRef(): itemPath3,
			},
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
				{
					Manifest: makeManifest(backup2.ID, ""),
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
							makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
				backup2.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.DetailsEntry{
							// This entry should not be picked due to a mismatch on Reasons.
							makeDetailsEntry(suite.T(), itemPath1, 84, false),
							// This item should be picked.
							makeDetailsEntry(suite.T(), itemPath3, 37, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []details.DetailsEntry{
				makeDetailsEntry(suite.T(), itemPath1, 42, false),
				makeDetailsEntry(suite.T(), itemPath3, 37, false),
			},
		},
		{
			name: "SomeBasesIncomplete",
			inputToMerge: map[string]path.Path{
				itemPath1.ShortRef(): itemPath1,
			},
			inputMans: []*kopia.ManifestEntry{
				{
					Manifest: makeManifest(backup1.ID, ""),
					Reasons: []kopia.Reason{
						pathReason1,
					},
				},
				{
					Manifest: makeManifest(backup2.ID, "checkpoint"),
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
							makeDetailsEntry(suite.T(), itemPath1, 42, false),
						},
					},
				},
				backup2.DetailsID: {
					DetailsModel: details.DetailsModel{
						Entries: []details.DetailsEntry{
							// This entry should not be picked due to being incomplete.
							makeDetailsEntry(suite.T(), itemPath1, 84, false),
						},
					},
				},
			},
			errCheck: assert.NoError,
			expectedEntries: []details.DetailsEntry{
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

			deets := details.Details{}

			err := mergeDetails(
				ctx,
				w,
				mdr,
				test.inputMans,
				test.inputToMerge,
				&deets,
			)

			test.errCheck(t, err)
			if err != nil {
				return
			}

			assert.ElementsMatch(t, test.expectedEntries, deets.Entries)
		})
	}
}

// ---------------------------------------------------------------------------
// integration
// ---------------------------------------------------------------------------

//revive:disable:context-as-argument
func prepNewBackupOp(
	t *testing.T,
	ctx context.Context,
	bus events.Eventer,
	sel selectors.Selector,
) (BackupOperation, account.Account, *kopia.Wrapper, *kopia.ModelStore, func()) {
	//revive:enable:context-as-argument
	acct := tester.NewM365Account(t)

	// need to initialize the repository before we can test connecting to it.
	st := tester.NewPrefixedS3Storage(t)

	k := kopia.NewConn(st)
	require.NoError(t, k.Initialize(ctx))

	// kopiaRef comes with a count of 1 and Wrapper bumps it again so safe
	// to close here.
	closer := func() { k.Close(ctx) }

	kw, err := kopia.NewWrapper(k)
	if !assert.NoError(t, err) {
		closer()
		t.FailNow()
	}

	closer = func() {
		k.Close(ctx)
		kw.Close(ctx)
	}

	ms, err := kopia.NewModelStore(k)
	if !assert.NoError(t, err) {
		closer()
		t.FailNow()
	}

	closer = func() {
		k.Close(ctx)
		kw.Close(ctx)
		ms.Close(ctx)
	}

	sw := store.NewKopiaStore(ms)

	bo, err := NewBackupOperation(
		ctx,
		control.Options{},
		kw,
		sw,
		acct,
		sel,
		bus)
	if !assert.NoError(t, err) {
		closer()
		t.FailNow()
	}

	return bo, acct, kw, ms, closer
}

//revive:disable:context-as-argument
func checkMetadataFilesExist(
	t *testing.T,
	ctx context.Context,
	backupID model.StableID,
	kw *kopia.Wrapper,
	ms *kopia.ModelStore,
	tenant, user string,
	service path.ServiceType,
	category path.CategoryType,
	files []string,
) {
	//revive:enable:context-as-argument
	bup := &backup.Backup{}

	err := ms.Get(ctx, model.BackupSchema, backupID, bup)
	if !assert.NoError(t, err) {
		return
	}

	paths := []path.Path{}
	pathsByRef := map[string][]string{}

	for _, fName := range files {
		p, err := path.Builder{}.
			Append(fName).
			ToServiceCategoryMetadataPath(tenant, user, service, category, true)
		if !assert.NoError(t, err, "bad metadata path") {
			continue
		}

		dir, err := p.Dir()
		if !assert.NoError(t, err, "parent path") {
			continue
		}

		paths = append(paths, p)
		pathsByRef[dir.ShortRef()] = append(pathsByRef[dir.ShortRef()], fName)
	}

	cols, err := kw.RestoreMultipleItems(ctx, bup.SnapshotID, paths, nil)
	assert.NoError(t, err)

	for _, col := range cols {
		itemNames := []string{}

		for item := range col.Items() {
			assert.Implements(t, (*data.StreamSize)(nil), item)

			s := item.(data.StreamSize)
			assert.Greaterf(
				t,
				s.Size(),
				int64(0),
				"empty metadata file: %s/%s",
				col.FullPath(),
				item.UUID(),
			)

			itemNames = append(itemNames, item.UUID())
		}

		assert.ElementsMatchf(
			t,
			pathsByRef[col.FullPath().ShortRef()],
			itemNames,
			"collection %s missing expected files",
			col.FullPath(),
		)
	}
}

type BackupOpIntegrationSuite struct {
	suite.Suite
}

func TestBackupOpIntegrationSuite(t *testing.T) {
	if err := tester.RunOnAny(
		tester.CorsoCITests,
		tester.CorsoOperationTests,
		tester.CorsoOperationBackupTests,
	); err != nil {
		t.Skip(err)
	}

	suite.Run(t, new(BackupOpIntegrationSuite))
}

func (suite *BackupOpIntegrationSuite) SetupSuite() {
	_, err := tester.GetRequiredEnvSls(
		tester.AWSStorageCredEnvs,
		tester.M365AcctCredEnvs)
	require.NoError(suite.T(), err)
}

func (suite *BackupOpIntegrationSuite) TestNewBackupOperation() {
	kw := &kopia.Wrapper{}
	sw := &store.Wrapper{}
	acct := tester.NewM365Account(suite.T())

	table := []struct {
		name     string
		opts     control.Options
		kw       *kopia.Wrapper
		sw       *store.Wrapper
		acct     account.Account
		targets  []string
		errCheck assert.ErrorAssertionFunc
	}{
		{"good", control.Options{}, kw, sw, acct, nil, assert.NoError},
		{"missing kopia", control.Options{}, nil, sw, acct, nil, assert.Error},
		{"missing modelstore", control.Options{}, kw, nil, acct, nil, assert.Error},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			ctx, flush := tester.NewContext()
			defer flush()

			_, err := NewBackupOperation(
				ctx,
				test.opts,
				test.kw,
				test.sw,
				test.acct,
				selectors.Selector{},
				evmock.NewBus())
			test.errCheck(t, err)
		})
	}
}

// TestBackup_Run ensures that Integration Testing works
// for the following scopes: Contacts, Events, and Mail
func (suite *BackupOpIntegrationSuite) TestBackup_Run_exchange() {
	ctx, flush := tester.NewContext()
	defer flush()

	m365UserID := tester.M365UserID(suite.T())
	users := []string{m365UserID}

	tests := []struct {
		name          string
		selectFunc    func() *selectors.ExchangeBackup
		resourceOwner string
		category      path.CategoryType
		metadataFiles []string
	}{
		{
			name: "Mail",
			selectFunc: func() *selectors.ExchangeBackup {
				sel := selectors.NewExchangeBackup(users)
				sel.Include(sel.MailFolders(users, []string{exchange.DefaultMailFolder}, selectors.PrefixMatch()))
				return sel
			},
			resourceOwner: m365UserID,
			category:      path.EmailCategory,
			metadataFiles: exchange.MetadataFileNames(path.EmailCategory),
		},
		{
			name: "Contacts",
			selectFunc: func() *selectors.ExchangeBackup {
				sel := selectors.NewExchangeBackup(users)
				sel.Include(sel.ContactFolders(
					users,
					[]string{exchange.DefaultContactFolder},
					selectors.PrefixMatch()))

				return sel
			},
			resourceOwner: m365UserID,
			category:      path.ContactsCategory,
			metadataFiles: exchange.MetadataFileNames(path.ContactsCategory),
		},
		{
			name: "Calendar Events",
			selectFunc: func() *selectors.ExchangeBackup {
				sel := selectors.NewExchangeBackup(users)
				sel.Include(sel.EventCalendars(users, []string{exchange.DefaultCalendar}, selectors.PrefixMatch()))
				return sel
			},
			resourceOwner: m365UserID,
			category:      path.EventsCategory,
			metadataFiles: exchange.MetadataFileNames(path.EventsCategory),
		},
	}
	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			mb := evmock.NewBus()
			sel := test.selectFunc()
			bo, acct, kw, ms, closer := prepNewBackupOp(t, ctx, mb, sel.Selector)
			defer closer()

			failed := false

			require.NoError(t, bo.Run(ctx))
			require.NotEmpty(t, bo.Results)
			require.NotEmpty(t, bo.Results.BackupID)

			if !assert.Equalf(
				t,
				Completed,
				bo.Status,
				"backup status %s is not Completed",
				bo.Status,
			) {
				failed = true
			}

			if !assert.Less(t, 0, bo.Results.ItemsWritten) {
				failed = true
			}

			assert.Less(t, 0, bo.Results.ItemsRead)
			assert.Less(t, int64(0), bo.Results.BytesRead, "bytes read")
			assert.Less(t, int64(0), bo.Results.BytesUploaded, "bytes uploaded")
			assert.Equal(t, 1, bo.Results.ResourceOwners)
			assert.NoError(t, bo.Results.ReadErrors)
			assert.NoError(t, bo.Results.WriteErrors)
			assert.Equal(t, 1, mb.TimesCalled[events.BackupStart], "backup-start events")
			assert.Equal(t, 1, mb.TimesCalled[events.BackupEnd], "backup-end events")
			assert.Equal(t,
				mb.CalledWith[events.BackupStart][0][events.BackupID],
				bo.Results.BackupID, "backupID pre-declaration")

			// verify that we can find the new backup id in the manifests
			var (
				sck, scv = kopia.MakeServiceCat(sel.PathService(), test.category)
				oc       = &kopia.OwnersCats{
					ResourceOwners: map[string]struct{}{test.resourceOwner: {}},
					ServiceCats:    map[string]kopia.ServiceCat{sck: scv},
				}
				tags  = map[string]string{kopia.TagBackupCategory: ""}
				found bool
			)

			mans, err := kw.FetchPrevSnapshotManifests(ctx, oc, tags)
			assert.NoError(t, err)

			for _, man := range mans {
				tk, _ := kopia.MakeTagKV(kopia.TagBackupID)
				if man.Tags[tk] == string(bo.Results.BackupID) {
					found = true
					break
				}
			}

			assert.True(t, found, "backup retrieved by previous snapshot manifest")

			if failed {
				return
			}

			m365, err := acct.M365Config()
			require.NoError(t, err)

			checkMetadataFilesExist(
				t,
				ctx,
				bo.Results.BackupID,
				kw,
				ms,
				m365.AzureTenantID,
				m365UserID,
				path.ExchangeService,
				test.category,
				test.metadataFiles,
			)
		})
	}
}

func (suite *BackupOpIntegrationSuite) TestBackup_Run_oneDrive() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t          = suite.T()
		mb         = evmock.NewBus()
		m365UserID = tester.SecondaryM365UserID(t)
		sel        = selectors.NewOneDriveBackup([]string{m365UserID})
	)

	sel.Include(sel.Users([]string{m365UserID}))

	bo, _, _, _, closer := prepNewBackupOp(t, ctx, mb, sel.Selector)
	defer closer()

	require.NoError(t, bo.Run(ctx))
	require.NotEmpty(t, bo.Results)
	require.NotEmpty(t, bo.Results.BackupID)
	assert.Equalf(t, Completed, bo.Status, "backup status %s is not Completed", bo.Status)
	assert.Equal(t, bo.Results.ItemsRead, bo.Results.ItemsWritten)
	assert.Less(t, int64(0), bo.Results.BytesRead, "bytes read")
	assert.Less(t, int64(0), bo.Results.BytesUploaded, "bytes uploaded")
	assert.Equal(t, 1, bo.Results.ResourceOwners)
	assert.NoError(t, bo.Results.ReadErrors)
	assert.NoError(t, bo.Results.WriteErrors)
	assert.Equal(t, 1, mb.TimesCalled[events.BackupStart], "backup-start events")
	assert.Equal(t, 1, mb.TimesCalled[events.BackupEnd], "backup-end events")
	assert.Equal(t,
		mb.CalledWith[events.BackupStart][0][events.BackupID],
		bo.Results.BackupID, "backupID pre-declaration")
}

func (suite *BackupOpIntegrationSuite) TestBackup_Run_sharePoint() {
	ctx, flush := tester.NewContext()
	defer flush()

	var (
		t      = suite.T()
		mb     = evmock.NewBus()
		siteID = tester.M365SiteID(t)
		sel    = selectors.NewSharePointBackup([]string{siteID})
	)

	sel.Include(sel.Sites([]string{siteID}))

	bo, _, _, _, closer := prepNewBackupOp(t, ctx, mb, sel.Selector)
	defer closer()

	require.NoError(t, bo.Run(ctx))
	require.NotEmpty(t, bo.Results)
	require.NotEmpty(t, bo.Results.BackupID)
	assert.Equalf(t, Completed, bo.Status, "backup status %s is not Completed", bo.Status)
	assert.Equal(t, bo.Results.ItemsRead, bo.Results.ItemsWritten)
	assert.Less(t, int64(0), bo.Results.BytesRead, "bytes read")
	assert.Less(t, int64(0), bo.Results.BytesUploaded, "bytes uploaded")
	assert.Equal(t, 1, bo.Results.ResourceOwners)
	assert.NoError(t, bo.Results.ReadErrors)
	assert.NoError(t, bo.Results.WriteErrors)
	assert.Equal(t, 1, mb.TimesCalled[events.BackupStart], "backup-start events")
	assert.Equal(t, 1, mb.TimesCalled[events.BackupEnd], "backup-end events")
	assert.Equal(t,
		mb.CalledWith[events.BackupStart][0][events.BackupID],
		bo.Results.BackupID, "backupID pre-declaration")
}
