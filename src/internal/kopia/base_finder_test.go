package kopia

import (
	"context"
	"testing"
	"time"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	testCompleteMan   = false
	testIncompleteMan = !testCompleteMan
)

var (
	testT1 = time.Now()
	testT2 = testT1.Add(1 * time.Hour)

	testID1 = manifest.ID("snap1")
	testID2 = manifest.ID("snap2")

	testBackup1 = "backupID1"
	testBackup2 = "backupID2"

	testMail   = path.ExchangeService.String() + path.EmailCategory.String()
	testEvents = path.ExchangeService.String() + path.EventsCategory.String()

	testUser1 = "user1"
	testUser2 = "user2"
	testUser3 = "user3"

	testAllUsersAllCats = []identity.Reasoner{
		// User1 email and events.
		NewReason("", testUser1, path.ExchangeService, path.EmailCategory),
		NewReason("", testUser1, path.ExchangeService, path.EventsCategory),
		// User2 email and events.
		NewReason("", testUser2, path.ExchangeService, path.EmailCategory),
		NewReason("", testUser2, path.ExchangeService, path.EventsCategory),
		// User3 email and events.
		NewReason("", testUser3, path.ExchangeService, path.EmailCategory),
		NewReason("", testUser3, path.ExchangeService, path.EventsCategory),
	}
	testAllUsersMail = []identity.Reasoner{
		NewReason("", testUser1, path.ExchangeService, path.EmailCategory),
		NewReason("", testUser2, path.ExchangeService, path.EmailCategory),
		NewReason("", testUser3, path.ExchangeService, path.EmailCategory),
	}
	testUser1Mail = []identity.Reasoner{
		NewReason("", testUser1, path.ExchangeService, path.EmailCategory),
	}
)

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

func newManifestInfo(
	id manifest.ID,
	modTime time.Time,
	incomplete bool,
	backupID string,
	err error,
	tags ...string,
) manifestInfo {
	incompleteStr := ""
	if incomplete {
		incompleteStr = "checkpoint"
	}

	structTags := make(map[string]string, len(tags))

	for _, t := range tags {
		tk, _ := makeTagKV(t)
		structTags[tk] = ""
	}

	res := manifestInfo{
		tags: structTags,
		err:  err,
		metadata: &manifest.EntryMetadata{
			ID:      id,
			ModTime: modTime,
			Labels:  structTags,
		},
		man: &snapshot.Manifest{
			ID:               id,
			IncompleteReason: incompleteStr,
			Tags:             structTags,
		},
	}

	if len(backupID) > 0 {
		k, _ := makeTagKV(TagBackupID)
		res.metadata.Labels[k] = backupID
		res.man.Tags[k] = backupID
	}

	return res
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

func newBackupModel(
	id string,
	hasItemSnap bool,
	hasDetailsSnap bool,
	oldDetailsID bool,
	err error,
) backupInfo {
	res := backupInfo{
		b: backup.Backup{
			BaseModel: model.BaseModel{
				ID: model.StableID(id),
			},
			SnapshotID: "iid",
		},
		err: err,
	}

	if hasDetailsSnap {
		if !oldDetailsID {
			res.b.StreamStoreID = "ssid"
		} else {
			res.b.DetailsID = "ssid"
		}
	}

	return res
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
		NewReason("", "a-user", path.ExchangeService, path.EmailCategory),
	}

	bb := bf.FindBases(ctx, reasons, nil)
	assert.Empty(t, bb.MergeBases())
	assert.Empty(t, bb.AssistBases())
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
		NewReason("", "a-user", path.ExchangeService, path.EmailCategory),
	}

	bb := bf.FindBases(ctx, reasons, nil)
	assert.Empty(t, bb.MergeBases())
	assert.Empty(t, bb.AssistBases())
}

func (suite *BaseFinderUnitSuite) TestGetBases() {
	table := []struct {
		name         string
		input        []identity.Reasoner
		manifestData []manifestInfo
		// Use this to denote the Reasons a base backup or base manifest is
		// selected. The int maps to the index of the backup or manifest in data.
		expectedBaseReasons map[int][]identity.Reasoner
		// Use this to denote the Reasons a kopia assised incrementals manifest is
		// selected. The int maps to the index of the manifest in data.
		expectedAssistManifestReasons map[int][]identity.Reasoner
		backupData                    []backupInfo
	}{
		{
			name:  "Return Older Base If Fail To Get Manifest",
			input: testUser1Mail,
			manifestData: []manifestInfo{
				newManifestInfo(
					testID2,
					testT2,
					testCompleteMan,
					testBackup2,
					assert.AnError,
					testMail,
					testUser1,
				),
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testBackup1,
					nil,
					testMail,
					testUser1,
				),
			},
			expectedBaseReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
			expectedAssistManifestReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
			backupData: []backupInfo{
				newBackupModel(testBackup2, true, true, false, nil),
				newBackupModel(testBackup1, true, true, false, nil),
			},
		},
		{
			name:  "Return Older Base If Fail To Get Backup",
			input: testUser1Mail,
			manifestData: []manifestInfo{
				newManifestInfo(
					testID2,
					testT2,
					testCompleteMan,
					testBackup2,
					nil,
					testMail,
					testUser1,
				),
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testBackup1,
					nil,
					testMail,
					testUser1,
				),
			},
			expectedBaseReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
			expectedAssistManifestReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
				1: testUser1Mail,
			},
			backupData: []backupInfo{
				newBackupModel(testBackup2, false, false, false, assert.AnError),
				newBackupModel(testBackup1, true, true, false, nil),
			},
		},
		{
			name:  "Return Older Base If Missing Details",
			input: testUser1Mail,
			manifestData: []manifestInfo{
				newManifestInfo(
					testID2,
					testT2,
					testCompleteMan,
					testBackup2,
					nil,
					testMail,
					testUser1,
				),
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testBackup1,
					nil,
					testMail,
					testUser1,
				),
			},
			expectedBaseReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
			expectedAssistManifestReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
				1: testUser1Mail,
			},
			backupData: []backupInfo{
				newBackupModel(testBackup2, true, false, false, nil),
				newBackupModel(testBackup1, true, true, false, nil),
			},
		},
		{
			name:  "Old Backup Details Pointer",
			input: testUser1Mail,
			manifestData: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testBackup1,
					nil,
					testMail,
					testEvents,
					testUser1,
					testUser2,
					testUser3,
				),
			},
			expectedBaseReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
			expectedAssistManifestReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
			backupData: []backupInfo{
				newBackupModel(testBackup1, true, true, true, nil),
			},
		},
		{
			name:  "All One Snapshot",
			input: testAllUsersAllCats,
			manifestData: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testBackup1,
					nil,
					testMail,
					testEvents,
					testUser1,
					testUser2,
					testUser3,
				),
			},
			expectedBaseReasons: map[int][]identity.Reasoner{
				0: testAllUsersAllCats,
			},
			expectedAssistManifestReasons: map[int][]identity.Reasoner{
				0: testAllUsersAllCats,
			},
			backupData: []backupInfo{
				newBackupModel(testBackup1, true, true, false, nil),
			},
		},
		{
			name:  "Multiple Bases Some Overlapping Reasons",
			input: testAllUsersAllCats,
			manifestData: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testBackup1,
					nil,
					testMail,
					testEvents,
					testUser1,
					testUser2,
					testUser3,
				),
				newManifestInfo(
					testID2,
					testT2,
					testCompleteMan,
					testBackup2,
					nil,
					testEvents,
					testUser1,
					testUser2,
					testUser3,
				),
			},
			expectedBaseReasons: map[int][]identity.Reasoner{
				0: {
					NewReason("", testUser1, path.ExchangeService, path.EmailCategory),
					NewReason("", testUser2, path.ExchangeService, path.EmailCategory),
					NewReason("", testUser3, path.ExchangeService, path.EmailCategory),
				},
				1: {
					NewReason("", testUser1, path.ExchangeService, path.EventsCategory),
					NewReason("", testUser2, path.ExchangeService, path.EventsCategory),
					NewReason("", testUser3, path.ExchangeService, path.EventsCategory),
				},
			},
			expectedAssistManifestReasons: map[int][]identity.Reasoner{
				0: {
					NewReason("", testUser1, path.ExchangeService, path.EmailCategory),
					NewReason("", testUser2, path.ExchangeService, path.EmailCategory),
					NewReason("", testUser3, path.ExchangeService, path.EmailCategory),
				},
				1: {
					NewReason("", testUser1, path.ExchangeService, path.EventsCategory),
					NewReason("", testUser2, path.ExchangeService, path.EventsCategory),
					NewReason("", testUser3, path.ExchangeService, path.EventsCategory),
				},
			},
			backupData: []backupInfo{
				newBackupModel(testBackup1, true, true, false, nil),
				newBackupModel(testBackup2, true, true, false, nil),
			},
		},
		{
			name:  "Newer Incomplete Assist Snapshot",
			input: testUser1Mail,
			manifestData: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testBackup1,
					nil,
					testMail,
					testUser1,
				),
				newManifestInfo(
					testID2,
					testT2,
					testIncompleteMan,
					testBackup2,
					nil,
					testMail,
					testUser1,
				),
			},
			expectedBaseReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
			expectedAssistManifestReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
				1: testUser1Mail,
			},
			backupData: []backupInfo{
				newBackupModel(testBackup1, true, true, false, nil),
				// Shouldn't be returned but have here just so we can see.
				newBackupModel(testBackup2, true, true, false, nil),
			},
		},
		{
			name:  "Incomplete Older Than Complete",
			input: testUser1Mail,
			manifestData: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testIncompleteMan,
					testBackup1,
					nil,
					testMail,
					testUser1,
				),
				newManifestInfo(
					testID2,
					testT2,
					testCompleteMan,
					testBackup2,
					nil,
					testMail,
					testUser1,
				),
			},
			expectedBaseReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
			expectedAssistManifestReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
			backupData: []backupInfo{
				// Shouldn't be returned but have here just so we can see.
				newBackupModel(testBackup1, true, true, false, nil),
				newBackupModel(testBackup2, true, true, false, nil),
			},
		},
		{
			name:  "Newest Incomplete Only Incomplete",
			input: testUser1Mail,
			manifestData: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testIncompleteMan,
					testBackup1,
					nil,
					testMail,
					testUser1,
				),
				newManifestInfo(
					testID2,
					testT2,
					testIncompleteMan,
					testBackup2,
					nil,
					testMail,
					testUser1,
				),
			},
			expectedBaseReasons: map[int][]identity.Reasoner{},
			expectedAssistManifestReasons: map[int][]identity.Reasoner{
				1: testUser1Mail,
			},
			backupData: []backupInfo{
				// Shouldn't be returned but have here just so we can see.
				newBackupModel(testBackup1, true, true, false, nil),
				newBackupModel(testBackup2, true, true, false, nil),
			},
		},
		{
			name:  "Some Bases Not Found",
			input: testAllUsersMail,
			manifestData: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testBackup1,
					nil,
					testMail,
					testUser1,
				),
			},
			expectedBaseReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
			expectedAssistManifestReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
			backupData: []backupInfo{
				newBackupModel(testBackup1, true, true, false, nil),
			},
		},
		{
			name:  "Manifests Not Sorted",
			input: testAllUsersMail,
			// Manifests are currently returned in the order they're defined by the
			// mock.
			manifestData: []manifestInfo{
				newManifestInfo(
					testID2,
					testT2,
					testCompleteMan,
					testBackup2,
					nil,
					testMail,
					testUser1,
				),
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testBackup1,
					nil,
					testMail,
					testUser1,
				),
			},
			expectedBaseReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
			expectedAssistManifestReasons: map[int][]identity.Reasoner{
				0: testUser1Mail,
			},
			backupData: []backupInfo{
				newBackupModel(testBackup2, true, true, false, nil),
				// Shouldn't be returned but here just so we can check.
				newBackupModel(testBackup1, true, true, false, nil),
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			bf := baseFinder{
				sm: &mockSnapshotManager{data: test.manifestData},
				bg: &mockModelGetter{data: test.backupData},
			}

			bb := bf.FindBases(
				ctx,
				test.input,
				nil)

			checkBackupEntriesMatch(
				t,
				bb.Backups(),
				test.backupData,
				test.expectedBaseReasons)
			checkManifestEntriesMatch(
				t,
				bb.MergeBases(),
				test.manifestData,
				test.expectedBaseReasons)
			checkManifestEntriesMatch(
				t,
				bb.AssistBases(),
				test.manifestData,
				test.expectedAssistManifestReasons)
		})
	}
}

func (suite *BaseFinderUnitSuite) TestFindBases_CustomTags() {
	manifestData := []manifestInfo{
		newManifestInfo(
			testID1,
			testT1,
			testCompleteMan,
			testBackup1,
			nil,
			testMail,
			testUser1,
			"fnords",
			"smarf",
		),
	}
	backupData := []backupInfo{
		newBackupModel(testBackup1, true, true, false, nil),
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

			bf := baseFinder{
				sm: &mockSnapshotManager{data: manifestData},
				bg: &mockModelGetter{data: backupData},
			}

			bb := bf.FindBases(
				ctx,
				testAllUsersAllCats,
				test.tags)

			checkManifestEntriesMatch(
				t,
				bb.MergeBases(),
				manifestData,
				test.expectedIdxs)
		})
	}
}

func checkManifestEntriesMatch(
	t *testing.T,
	retSnaps []ManifestEntry,
	allExpected []manifestInfo,
	expectedIdxsAndReasons map[int][]identity.Reasoner,
) {
	// Check the proper snapshot manifests were returned.
	expected := make([]*snapshot.Manifest, 0, len(expectedIdxsAndReasons))
	for i := range expectedIdxsAndReasons {
		expected = append(expected, allExpected[i].man)
	}

	got := make([]*snapshot.Manifest, 0, len(retSnaps))
	for _, s := range retSnaps {
		got = append(got, s.Manifest)
	}

	assert.ElementsMatch(t, expected, got)

	// Check the reasons for selecting each manifest are correct.
	expectedReasons := make(map[manifest.ID][]identity.Reasoner, len(expectedIdxsAndReasons))
	for idx, reasons := range expectedIdxsAndReasons {
		expectedReasons[allExpected[idx].man.ID] = reasons
	}

	for _, found := range retSnaps {
		reasons, ok := expectedReasons[found.ID]
		if !ok {
			// Missing or extra snapshots will be reported by earlier checks.
			continue
		}

		assert.ElementsMatch(
			t,
			reasons,
			found.Reasons,
			"incorrect reasons for snapshot with ID %s",
			found.ID,
		)
	}
}

func checkBackupEntriesMatch(
	t *testing.T,
	retBups []BackupEntry,
	allExpected []backupInfo,
	expectedIdxsAndReasons map[int][]identity.Reasoner,
) {
	// Check the proper snapshot manifests were returned.
	expected := make([]*backup.Backup, 0, len(expectedIdxsAndReasons))
	for i := range expectedIdxsAndReasons {
		expected = append(expected, &allExpected[i].b)
	}

	got := make([]*backup.Backup, 0, len(retBups))
	for _, s := range retBups {
		got = append(got, s.Backup)
	}

	assert.ElementsMatch(t, expected, got)

	// Check the reasons for selecting each manifest are correct.
	expectedReasons := make(map[model.StableID][]identity.Reasoner, len(expectedIdxsAndReasons))
	for idx, reasons := range expectedIdxsAndReasons {
		expectedReasons[allExpected[idx].b.ID] = reasons
	}

	for _, found := range retBups {
		reasons, ok := expectedReasons[found.ID]
		if !ok {
			// Missing or extra snapshots will be reported by earlier checks.
			continue
		}

		assert.ElementsMatch(
			t,
			reasons,
			found.Reasons,
			"incorrect reasons for snapshot with ID %s",
			found.ID,
		)
	}
}
