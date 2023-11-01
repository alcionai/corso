package kopia

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/path"
)

var (
	testT1 = time.Now()
	testT2 = testT1.Add(1 * time.Hour)
	testT3 = testT2.Add(1 * time.Hour)
	testT4 = testT3.Add(1 * time.Hour)

	testUser1 = "user1"
	testUser2 = "user2"
	testUser3 = "user3"

	testAllUsersAllCats = []identity.Reasoner{
		// User1 email and events.
		newTestReason(testUser1, path.EmailCategory),
		newTestReason(testUser1, path.EventsCategory),
		// User2 email and events.
		newTestReason(testUser2, path.EmailCategory),
		newTestReason(testUser2, path.EventsCategory),
		// User3 email and events.
		newTestReason(testUser3, path.EmailCategory),
		newTestReason(testUser3, path.EventsCategory),
	}
	testAllUsersMail = []identity.Reasoner{
		newTestReason(testUser1, path.EmailCategory),
		newTestReason(testUser2, path.EmailCategory),
		newTestReason(testUser3, path.EmailCategory),
	}
	testUser1Mail = []identity.Reasoner{
		newTestReason(testUser1, path.EmailCategory),
	}
)

// newTestReason is a helper function to make sure Reasons have a consistent
// tenant and service when created within a test.
func newTestReason(user string, category path.CategoryType) identity.Reasoner {
	return identity.NewReason("", user, path.ExchangeService, category)
}

// -----------------------------------------------------------------------------
// Empty mocks that return no data
// -----------------------------------------------------------------------------
type mockEmptySnapshotManager struct{}

func (sm mockEmptySnapshotManager) FindManifests(
	context.Context,
	map[string]string,
) ([]*manifest.EntryMetadata, error) {
	return nil, nil
}

func (sm mockEmptySnapshotManager) LoadSnapshot(
	context.Context,
	manifest.ID,
) (*snapshot.Manifest, error) {
	return nil, snapshot.ErrSnapshotNotFound
}

type mockEmptyModelGetter struct{}

func (mg mockEmptyModelGetter) GetBackup(
	context.Context,
	model.StableID,
) (*backup.Backup, error) {
	return nil, data.ErrNotFound
}

// -----------------------------------------------------------------------------
// Mocks that return data or errors
// -----------------------------------------------------------------------------
type manifestInfo struct {
	// We don't currently use the values in the tags.
	tags     map[string]string
	metadata *manifest.EntryMetadata
	man      *snapshot.Manifest
	err      error
}

type mockSnapshotManager struct {
	data    []manifestInfo
	findErr error
}

func matchesTags(mi manifestInfo, tags map[string]string) bool {
	for k := range tags {
		if _, ok := mi.tags[k]; !ok {
			return false
		}
	}

	return true
}

func (msm *mockSnapshotManager) FindManifests(
	ctx context.Context,
	tags map[string]string,
) ([]*manifest.EntryMetadata, error) {
	if msm == nil {
		return nil, assert.AnError
	}

	if msm.findErr != nil {
		return nil, msm.findErr
	}

	res := []*manifest.EntryMetadata{}

	for _, mi := range msm.data {
		if matchesTags(mi, tags) {
			res = append(res, mi.metadata)
		}
	}

	return res, nil
}

func (msm *mockSnapshotManager) LoadSnapshot(
	ctx context.Context,
	id manifest.ID,
) (*snapshot.Manifest, error) {
	if msm == nil {
		return nil, assert.AnError
	}

	for _, mi := range msm.data {
		if mi.man.ID == id {
			if mi.err != nil {
				return nil, mi.err
			}

			return mi.man, nil
		}
	}

	return nil, snapshot.ErrSnapshotNotFound
}

type backupInfo struct {
	b   backup.Backup
	err error
}

type mockModelGetter struct {
	data []backupInfo
}

func (mg mockModelGetter) GetBackup(
	_ context.Context,
	id model.StableID,
) (*backup.Backup, error) {
	// Use struct here so we return a copy of the struct just in case the caller
	// somehow ends up modifying it.
	var res backup.Backup

	for _, bi := range mg.data {
		if bi.b.ID != id {
			continue
		}

		if bi.err != nil {
			return nil, bi.err
		}

		res = bi.b

		return &res, nil
	}

	return nil, data.ErrNotFound
}

type baseInfo struct {
	manifest manifestInfo
	backup   backupInfo
}

type baseInfoBuilder struct {
	info baseInfo
}

// newBaseInfoBuilder returns a builder with the given ID and mod time that's
// valid. Use functions defined on the builder if an invalid or non-standard
// state is required.
func newBaseInfoBuilder(
	id int,
	modTime time.Time,
	reasons ...identity.Reasoner,
) *baseInfoBuilder {
	snapID := fmt.Sprintf("snap%d", id)
	bupID := fmt.Sprintf("backup%d", id)
	deetsID := fmt.Sprintf("details%d", id)

	manifestTags := map[string]string{}

	for _, r := range reasons {
		for _, k := range tagKeys(r) {
			mk, mv := makeTagKV(k)
			manifestTags[mk] = mv
		}
	}

	k, _ := makeTagKV(TagBackupID)
	manifestTags[k] = bupID

	return &baseInfoBuilder{
		info: baseInfo{
			manifest: manifestInfo{
				tags: manifestTags,
				metadata: &manifest.EntryMetadata{
					ID:      manifest.ID(snapID),
					ModTime: modTime,
					Labels:  manifestTags,
				},
				man: &snapshot.Manifest{
					ID:   manifest.ID(snapID),
					Tags: manifestTags,
				},
			},
			backup: backupInfo{
				b: backup.Backup{
					BaseModel: model.BaseModel{
						ID: model.StableID(bupID),
					},
					SnapshotID:    snapID,
					StreamStoreID: deetsID,
				},
			},
		},
	}
}

func (builder *baseInfoBuilder) build() baseInfo {
	return builder.info
}

func (builder *baseInfoBuilder) setBackupType(
	backupType string,
) *baseInfoBuilder {
	if builder.info.backup.b.Tags == nil {
		builder.info.backup.b.Tags = map[string]string{}
	}

	builder.info.backup.b.Tags[model.BackupTypeTag] = backupType

	return builder
}

func (builder *baseInfoBuilder) setSnapshotIncomplete(
	reason string,
) *baseInfoBuilder {
	builder.info.manifest.man.IncompleteReason = reason
	return builder
}

func (builder *baseInfoBuilder) legacyBackupDetails() *baseInfoBuilder {
	builder.info.backup.b.DetailsID = builder.info.backup.b.StreamStoreID
	builder.info.backup.b.StreamStoreID = ""

	return builder
}

func (builder *baseInfoBuilder) setBackupItemSnapshotID(
	id string,
) *baseInfoBuilder {
	builder.info.backup.b.SnapshotID = id
	return builder
}

func (builder *baseInfoBuilder) clearBackupDetails() *baseInfoBuilder {
	builder.info.backup.b.DetailsID = ""
	builder.info.backup.b.StreamStoreID = ""

	return builder
}

func (builder *baseInfoBuilder) appendSnapshotTagKeys(
	tags ...string,
) *baseInfoBuilder {
	kvs := make(map[string]string, len(tags))

	for _, t := range tags {
		tk, _ := makeTagKV(t)
		kvs[tk] = ""
	}

	if builder.info.manifest.metadata.Labels == nil {
		builder.info.manifest.metadata.Labels = map[string]string{}
	}

	if builder.info.manifest.man.Tags == nil {
		builder.info.manifest.man.Tags = map[string]string{}
	}

	maps.Copy(builder.info.manifest.metadata.Labels, kvs)
	maps.Copy(builder.info.manifest.man.Tags, kvs)

	return builder
}

func (builder *baseInfoBuilder) setBackupError(
	err error,
) *baseInfoBuilder {
	builder.info.backup.err = err
	return builder
}

func (builder *baseInfoBuilder) setSnapshotError(
	err error,
) *baseInfoBuilder {
	builder.info.manifest.err = err
	return builder
}

// -----------------------------------------------------------------------------
// Tests for getting bases
// -----------------------------------------------------------------------------
type BaseFinderUnitSuite struct {
	tester.Suite
}

func TestBaseFinderUnitSuite(t *testing.T) {
	suite.Run(t, &BaseFinderUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *BaseFinderUnitSuite) TestNoResult_NoBackupsOrSnapshots() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	bf := baseFinder{
		sm: mockEmptySnapshotManager{},
		bg: mockEmptyModelGetter{},
	}
	reasons := []identity.Reasoner{
		identity.NewReason("", "a-user", path.ExchangeService, path.EmailCategory),
	}

	bb := bf.FindBases(ctx, reasons, nil)
	assert.Empty(t, bb.MergeBases())
	assert.Empty(t, bb.UniqueAssistBases())
	assert.Empty(t, bb.SnapshotAssistBases())
}

func (suite *BaseFinderUnitSuite) TestNoResult_ErrorListingSnapshots() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	bf := baseFinder{
		sm: &mockSnapshotManager{findErr: assert.AnError},
		bg: mockEmptyModelGetter{},
	}
	reasons := []identity.Reasoner{
		identity.NewReason("", "a-user", path.ExchangeService, path.EmailCategory),
	}

	bb := bf.FindBases(ctx, reasons, nil)
	assert.Empty(t, bb.MergeBases())
	assert.Empty(t, bb.UniqueAssistBases())
	assert.Empty(t, bb.SnapshotAssistBases())
}

func (suite *BaseFinderUnitSuite) TestGetBases() {
	table := []struct {
		name  string
		input []identity.Reasoner
		data  []baseInfo
		// Use this to denote the Reasons a base backup or base manifest is
		// selected. The int maps to the index of the backup or manifest in data.
		expectedMergeReasons map[int][]identity.Reasoner
		// Use this to denote the Reasons a kopia assised incrementals manifest is
		// selected. The int maps to the index of the manifest in data.
		expectedAssistReasons map[int][]identity.Reasoner
	}{
		{
			name:  "Return Older Merge Base If Fail To Get Manifest",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					setSnapshotError(assert.AnError).
					build(),
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
		},
		{
			name:  "Return Older Assist Base If Fail To Get Manifest",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					setSnapshotError(assert.AnError).
					build(),
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					setBackupType(model.AssistBackup).
					build(),
			},
			expectedAssistReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
		},
		{
			name:  "Return Older Merge Base If Fail To Get Backup",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					setBackupError(assert.AnError).
					build(),
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
		},
		{
			name:  "Return Older Base If Missing Details",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					clearBackupDetails().
					build(),
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
		},
		{
			name:  "Return Older Base If Snapshot ID Mismatch",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					setBackupItemSnapshotID("foo").
					build(),
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
		},
		{
			name:  "Old Backup Details Pointer",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(1, testT1, testAllUsersAllCats...).
					legacyBackupDetails().
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
		},
		{
			name:  "All One Snapshot With Merge Base",
			input: testAllUsersAllCats,
			data: []baseInfo{
				newBaseInfoBuilder(1, testT1, testAllUsersAllCats...).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				0: testAllUsersAllCats,
			},
		},
		{
			name:  "All One Snapshot with Assist Base",
			input: testAllUsersAllCats,
			data: []baseInfo{
				newBaseInfoBuilder(1, testT1, testAllUsersAllCats...).
					setBackupType(model.AssistBackup).
					build(),
			},
			expectedAssistReasons: map[int][]identity.Reasoner{
				0: testAllUsersAllCats,
			},
		},
		{
			name:  "Multiple Bases Some Overlapping Reasons",
			input: testAllUsersAllCats,
			data: []baseInfo{
				newBaseInfoBuilder(1, testT1, testAllUsersAllCats...).
					build(),
				newBaseInfoBuilder(2, testT2, testAllUsersMail...).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				0: {
					newTestReason(testUser1, path.EventsCategory),
					newTestReason(testUser2, path.EventsCategory),
					newTestReason(testUser3, path.EventsCategory),
				},
				1: {
					newTestReason(testUser1, path.EmailCategory),
					newTestReason(testUser2, path.EmailCategory),
					newTestReason(testUser3, path.EmailCategory),
				},
			},
		},
		{
			name:  "Unique assist bases with common merge Base, overlapping reasons",
			input: testAllUsersAllCats,
			data: []baseInfo{
				newBaseInfoBuilder(
					3,
					testT3,
					newTestReason(testUser1, path.EventsCategory),
					newTestReason(testUser2, path.EventsCategory)).
					setBackupType(model.AssistBackup).
					build(),
				newBaseInfoBuilder(
					2,
					testT2,
					newTestReason(testUser1, path.EmailCategory),
					newTestReason(testUser2, path.EmailCategory)).
					setBackupType(model.AssistBackup).
					build(),
				newBaseInfoBuilder(
					1,
					testT1,
					newTestReason(testUser1, path.EventsCategory),
					newTestReason(testUser2, path.EventsCategory),
					newTestReason(testUser1, path.EmailCategory),
					newTestReason(testUser2, path.EmailCategory)).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				2: {
					newTestReason(testUser1, path.EmailCategory),
					newTestReason(testUser2, path.EmailCategory),
					newTestReason(testUser1, path.EventsCategory),
					newTestReason(testUser2, path.EventsCategory),
				},
			},
			expectedAssistReasons: map[int][]identity.Reasoner{
				0: {
					newTestReason(testUser1, path.EventsCategory),
					newTestReason(testUser2, path.EventsCategory),
				},
				1: {
					newTestReason(testUser1, path.EmailCategory),
					newTestReason(testUser2, path.EmailCategory),
				},
			},
		},
		{
			name:  "Newer Incomplete Assist Snapshot",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					build(),
				// This base shouldn't be returned.
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					setSnapshotIncomplete("checkpoint").
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
		},
		{
			name:  "Incomplete Older Than Complete",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					setSnapshotIncomplete("checkpoint").
					build(),
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
		},
		{
			name:  "Only Incomplete Returns Nothing",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					setSnapshotIncomplete("checkpoint").
					build(),
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					setSnapshotIncomplete("checkpoint").
					build(),
			},
		},
		{
			name:  "Some Bases Not Found",
			input: testAllUsersMail,
			data: []baseInfo{
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
		},
		{
			name:  "Manifests Not Sorted",
			input: testAllUsersMail,
			// Manifests are currently returned in the order they're defined by the
			// mock.
			data: []baseInfo{
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					build(),
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
		},
		{
			name:  "Return latest assist & merge base pair",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(4, testT4, testUser1Mail...).
					setBackupType(model.AssistBackup).
					build(),
				newBaseInfoBuilder(3, testT3, testUser1Mail...).
					setBackupType(model.AssistBackup).
					build(),
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					build(),
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				2: testUser1Mail,
			},
			expectedAssistReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
		},
		{
			name:  "Newer merge base than assist base",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					build(),
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					setBackupType(model.AssistBackup).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
		},
		{
			name:  "Only assist bases",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					setBackupType(model.AssistBackup).
					build(),
				newBaseInfoBuilder(1, testT1, testUser1Mail...).
					setBackupType(model.AssistBackup).
					build(),
			},
			expectedAssistReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
		},
		{
			name:  "Merge base with merge tag",
			input: testUser1Mail,
			data: []baseInfo{
				newBaseInfoBuilder(2, testT2, testUser1Mail...).
					setBackupType(model.MergeBackup).
					build(),
			},
			expectedMergeReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mans := make([]manifestInfo, 0, len(test.data))
			bups := make([]backupInfo, 0, len(test.data))

			for _, d := range test.data {
				mans = append(mans, d.manifest)
				bups = append(bups, d.backup)
			}

			bf := baseFinder{
				sm: &mockSnapshotManager{data: mans},
				bg: &mockModelGetter{data: bups},
			}

			bb := bf.FindBases(
				ctx,
				test.input,
				nil)

			checkBaseEntriesMatch(
				t,
				bb.MergeBases(),
				test.data,
				test.expectedMergeReasons)
			checkBaseEntriesMatch(
				t,
				bb.UniqueAssistBases(),
				test.data,
				test.expectedAssistReasons)
		})
	}
}

func (suite *BaseFinderUnitSuite) TestFindBases_CustomTags() {
	inputData := []baseInfo{
		newBaseInfoBuilder(1, testT1, testUser1Mail...).
			appendSnapshotTagKeys("fnords", "smarf").
			build(),
	}

	table := []struct {
		name  string
		input []identity.Reasoner
		tags  map[string]string
		// Use this to denote which manifests in data should be expected. Allows
		// defining data in a table while not repeating things between data and
		// expected.
		expectedIdxs map[int][]identity.Reasoner
	}{
		{
			name: "no tags specified",
			tags: nil,
			expectedIdxs: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
		},
		{
			name: "all custom tags",
			tags: map[string]string{
				"fnords": "",
				"smarf":  "",
			},
			expectedIdxs: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
		},
		{
			name: "subset of custom tags",
			tags: map[string]string{"fnords": ""},
			expectedIdxs: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
		},
		{
			name:         "custom tag mismatch",
			tags:         map[string]string{"bojangles": ""},
			expectedIdxs: nil,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mans := make([]manifestInfo, 0, len(inputData))
			bups := make([]backupInfo, 0, len(inputData))

			for _, d := range inputData {
				mans = append(mans, d.manifest)
				bups = append(bups, d.backup)
			}

			bf := baseFinder{
				sm: &mockSnapshotManager{data: mans},
				bg: &mockModelGetter{data: bups},
			}

			bb := bf.FindBases(
				ctx,
				testAllUsersAllCats,
				test.tags)

			checkBaseEntriesMatch(
				t,
				bb.MergeBases(),
				inputData,
				test.expectedIdxs)
		})
	}
}

func checkBaseEntriesMatch(
	t *testing.T,
	gotBases []BackupBase,
	allBases []baseInfo,
	expectedIdxsAndReasons map[int][]identity.Reasoner,
) {
	// Check the proper snapshot manifests were returned.
	expectedMans := make([]*snapshot.Manifest, 0, len(expectedIdxsAndReasons))
	expectedBups := make([]*backup.Backup, 0, len(expectedIdxsAndReasons))

	for i := range expectedIdxsAndReasons {
		expectedMans = append(expectedMans, allBases[i].manifest.man)
		expectedBups = append(expectedBups, &allBases[i].backup.b)
	}

	gotMans := make([]*snapshot.Manifest, 0, len(gotBases))
	gotBups := make([]*backup.Backup, 0, len(gotBases))

	for _, s := range gotBases {
		gotMans = append(gotMans, s.ItemDataSnapshot)
		gotBups = append(gotBups, s.Backup)
	}

	assert.ElementsMatch(t, expectedMans, gotMans, "item data manifests")
	assert.ElementsMatch(t, expectedBups, gotBups, "backup models")

	// Check the reasons for selecting each manifest are correct.
	expectedReasons := make(map[model.StableID][]identity.Reasoner, len(expectedIdxsAndReasons))

	for idx, reasons := range expectedIdxsAndReasons {
		expectedReasons[allBases[idx].backup.b.ID] = reasons
	}

	for _, found := range gotBases {
		reasons, ok := expectedReasons[found.Backup.ID]
		if !ok {
			// Missing or extra snapshots will be reported by earlier checks.
			continue
		}

		assert.ElementsMatch(
			t,
			reasons,
			found.Reasons,
			"incorrect reasons for backup with ID %s",
			found.Backup.ID)
	}
}
