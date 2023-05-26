package kopia

import (
	"context"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/fs"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/path"
)

const (
	testCompleteMan   = false
	testIncompleteMan = !testCompleteMan
)

var (
	testT1 = time.Now()
	testT2 = testT1.Add(1 * time.Hour)
	testT3 = testT2.Add(1 * time.Hour)

	testID1 = manifest.ID("snap1")
	testID2 = manifest.ID("snap2")
	testID3 = manifest.ID("snap3")

	testMail   = path.ExchangeService.String() + path.EmailCategory.String()
	testEvents = path.ExchangeService.String() + path.EventsCategory.String()

	testUser1 = "user1"
	testUser2 = "user2"
	testUser3 = "user3"

	testAllUsersAllCats = []Reason{
		{
			ResourceOwner: testUser1,
			Service:       path.ExchangeService,
			Category:      path.EmailCategory,
		},
		{
			ResourceOwner: testUser1,
			Service:       path.ExchangeService,
			Category:      path.EventsCategory,
		},
		{
			ResourceOwner: testUser2,
			Service:       path.ExchangeService,
			Category:      path.EmailCategory,
		},
		{
			ResourceOwner: testUser2,
			Service:       path.ExchangeService,
			Category:      path.EventsCategory,
		},
		{
			ResourceOwner: testUser3,
			Service:       path.ExchangeService,
			Category:      path.EmailCategory,
		},
		{
			ResourceOwner: testUser3,
			Service:       path.ExchangeService,
			Category:      path.EventsCategory,
		},
	}
	testAllUsersMail = []Reason{
		{
			ResourceOwner: testUser1,
			Service:       path.ExchangeService,
			Category:      path.EmailCategory,
		},
		{
			ResourceOwner: testUser2,
			Service:       path.ExchangeService,
			Category:      path.EmailCategory,
		},
		{
			ResourceOwner: testUser3,
			Service:       path.ExchangeService,
			Category:      path.EmailCategory,
		},
	}
)

type manifestInfo struct {
	// We don't currently use the values in the tags.
	tags     map[string]struct{}
	metadata *manifest.EntryMetadata
	man      *snapshot.Manifest
}

func newManifestInfo(
	id manifest.ID,
	modTime time.Time,
	incomplete bool,
	tags ...string,
) manifestInfo {
	incompleteStr := ""
	if incomplete {
		incompleteStr = "checkpoint"
	}

	structTags := make(map[string]struct{}, len(tags))

	for _, t := range tags {
		tk, _ := makeTagKV(t)
		structTags[tk] = struct{}{}
	}

	return manifestInfo{
		tags: structTags,
		metadata: &manifest.EntryMetadata{
			ID:      id,
			ModTime: modTime,
		},
		man: &snapshot.Manifest{
			ID:               id,
			StartTime:        fs.UTCTimestamp(modTime.UnixNano()),
			IncompleteReason: incompleteStr,
		},
	}
}

type mockSnapshotManager struct {
	data         []manifestInfo
	loadCallback func(ids []manifest.ID)
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

	res := []*manifest.EntryMetadata{}

	for _, mi := range msm.data {
		if matchesTags(mi, tags) {
			res = append(res, mi.metadata)
		}
	}

	return res, nil
}

func (msm *mockSnapshotManager) LoadSnapshots(
	ctx context.Context,
	ids []manifest.ID,
) ([]*snapshot.Manifest, error) {
	if msm == nil {
		return nil, assert.AnError
	}

	// Allow checking set of IDs passed in.
	if msm.loadCallback != nil {
		msm.loadCallback(ids)
	}

	res := []*snapshot.Manifest{}

	for _, id := range ids {
		for _, mi := range msm.data {
			if mi.man.ID == id {
				res = append(res, mi.man)
			}
		}
	}

	return res, nil
}

func (msm *mockSnapshotManager) LoadSnapshot(
	ctx context.Context,
	id manifest.ID,
) (*snapshot.Manifest, error) {
	return nil, clues.New("not implemented")
}

type SnapshotFetchUnitSuite struct {
	tester.Suite
}

func TestSnapshotFetchUnitSuite(t *testing.T) {
	suite.Run(t, &SnapshotFetchUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *SnapshotFetchUnitSuite) TestFetchPrevSnapshots() {
	table := []struct {
		name  string
		input []Reason
		data  []manifestInfo
		// Use this to denote which manifests in data should be expected. Allows
		// defining data in a table while not repeating things between data and
		// expected.
		expectedIdxs []int
		// Use this to denote the Reasons a manifest is selected. The int maps to
		// the index of the manifest in data.
		expectedReasons map[int][]Reason
		// Expected number of times a manifest should try to be loaded from kopia.
		// Used to check that caching is functioning properly.
		expectedLoadCounts map[manifest.ID]int
	}{
		{
			name:  "AllOneSnapshot",
			input: testAllUsersAllCats,
			data: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testMail,
					testEvents,
					testUser1,
					testUser2,
					testUser3,
				),
			},
			expectedIdxs: []int{0},
			expectedReasons: map[int][]Reason{
				0: {
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EventsCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EventsCategory,
					},
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EventsCategory,
					},
				},
			},
			expectedLoadCounts: map[manifest.ID]int{
				testID1: 1,
			},
		},
		{
			name:  "SplitByCategory",
			input: testAllUsersAllCats,
			data: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testMail,
					testUser1,
					testUser2,
					testUser3,
				),
				newManifestInfo(
					testID2,
					testT2,
					testCompleteMan,
					testEvents,
					testUser1,
					testUser2,
					testUser3,
				),
			},
			expectedIdxs: []int{0, 1},
			expectedReasons: map[int][]Reason{
				0: {
					{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
				1: {
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EventsCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EventsCategory,
					},
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EventsCategory,
					},
				},
			},
			expectedLoadCounts: map[manifest.ID]int{
				testID1: 1,
				testID2: 1,
			},
		},
		{
			name:  "IncompleteNewerThanComplete",
			input: testAllUsersMail,
			data: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testMail,
					testUser1,
					testUser2,
					testUser3,
				),
				newManifestInfo(
					testID2,
					testT2,
					testIncompleteMan,
					testMail,
					testUser1,
					testUser2,
					testUser3,
				),
			},
			expectedIdxs: []int{0, 1},
			expectedReasons: map[int][]Reason{
				0: {
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
				1: {
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
			},
			expectedLoadCounts: map[manifest.ID]int{
				testID1: 1,
				testID2: 3,
			},
		},
		{
			name:  "IncompleteOlderThanComplete",
			input: testAllUsersMail,
			data: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testIncompleteMan,
					testMail,
					testUser1,
					testUser2,
					testUser3,
				),
				newManifestInfo(
					testID2,
					testT2,
					testCompleteMan,
					testMail,
					testUser1,
					testUser2,
					testUser3,
				),
			},
			expectedIdxs: []int{1},
			expectedReasons: map[int][]Reason{
				1: {
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
			},
			expectedLoadCounts: map[manifest.ID]int{
				testID1: 1,
				testID2: 1,
			},
		},
		{
			name:  "OnlyIncomplete",
			input: testAllUsersMail,
			data: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testIncompleteMan,
					testMail,
					testUser1,
					testUser2,
					testUser3,
				),
			},
			expectedIdxs: []int{0},
			expectedReasons: map[int][]Reason{
				0: {
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
			},
			expectedLoadCounts: map[manifest.ID]int{
				testID1: 3,
			},
		},
		{
			name:  "NewestComplete",
			input: testAllUsersMail,
			data: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testMail,
					testUser1,
					testUser2,
					testUser3,
				),
				newManifestInfo(
					testID2,
					testT2,
					testCompleteMan,
					testMail,
					testUser1,
					testUser2,
					testUser3,
				),
			},
			expectedIdxs: []int{1},
			expectedReasons: map[int][]Reason{
				1: {
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
			},
			expectedLoadCounts: map[manifest.ID]int{
				testID1: 1,
				testID2: 1,
			},
		},
		{
			name:  "NewestIncomplete",
			input: testAllUsersMail,
			data: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testIncompleteMan,
					testMail,
					testUser1,
					testUser2,
					testUser3,
				),
				newManifestInfo(
					testID2,
					testT2,
					testIncompleteMan,
					testMail,
					testUser1,
					testUser2,
					testUser3,
				),
			},
			expectedIdxs: []int{1},
			expectedReasons: map[int][]Reason{
				1: {
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
			},
			expectedLoadCounts: map[manifest.ID]int{
				testID1: 3,
				testID2: 3,
			},
		},
		{
			name:  "SomeCachedSomeNewer",
			input: testAllUsersMail,
			data: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testMail,
					testUser1,
					testUser2,
					testUser3,
				),
				newManifestInfo(
					testID2,
					testT2,
					testCompleteMan,
					testMail,
					testUser3,
				),
			},
			expectedIdxs: []int{0, 1},
			expectedReasons: map[int][]Reason{
				0: {
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
				1: {
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
			},
			expectedLoadCounts: map[manifest.ID]int{
				testID1: 2,
				testID2: 1,
			},
		},
		{
			name:  "SomeCachedSomeNewerDifferentCategories",
			input: testAllUsersAllCats,
			data: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
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
					testMail,
					testUser3,
				),
			},
			expectedIdxs: []int{0, 1},
			expectedReasons: map[int][]Reason{
				0: {
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EventsCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EventsCategory,
					},
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EventsCategory,
					},
				},
				1: {
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
			},
			expectedLoadCounts: map[manifest.ID]int{
				testID1: 2,
				testID2: 1,
			},
		},
		{
			name:  "SomeCachedSomeNewerIncomplete",
			input: testAllUsersMail,
			data: []manifestInfo{
				newManifestInfo(
					testID1,
					testT1,
					testCompleteMan,
					testMail,
					testUser1,
					testUser2,
					testUser3,
				),
				newManifestInfo(
					testID2,
					testT2,
					testIncompleteMan,
					testMail,
					testUser3,
				),
			},
			expectedIdxs: []int{0, 1},
			expectedReasons: map[int][]Reason{
				0: {
					Reason{
						ResourceOwner: testUser1,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser2,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
				1: {
					Reason{
						ResourceOwner: testUser3,
						Service:       path.ExchangeService,
						Category:      path.EmailCategory,
					},
				},
			},
			expectedLoadCounts: map[manifest.ID]int{
				testID1: 1,
				testID2: 1,
			},
		},
		{
			name:         "NoMatches",
			input:        testAllUsersMail,
			data:         nil,
			expectedIdxs: nil,
			// Stop failure for nil-map comparison.
			expectedLoadCounts: map[manifest.ID]int{},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			msm := &mockSnapshotManager{
				data: test.data,
			}

			loadCounts := map[manifest.ID]int{}
			msm.loadCallback = func(ids []manifest.ID) {
				for _, id := range ids {
					loadCounts[id]++
				}
			}

			snaps := fetchPrevSnapshotManifests(ctx, msm, test.input, nil)

			// Check the proper snapshot manifests were returned.
			expected := make([]*snapshot.Manifest, 0, len(test.expectedIdxs))
			for _, i := range test.expectedIdxs {
				expected = append(expected, test.data[i].man)
			}

			got := make([]*snapshot.Manifest, 0, len(snaps))
			for _, s := range snaps {
				got = append(got, s.Manifest)
			}

			assert.ElementsMatch(t, expected, got)

			// Check the reasons for selecting each manifest are correct.
			expectedReasons := make(map[manifest.ID][]Reason, len(test.expectedReasons))
			for idx, reason := range test.expectedReasons {
				expectedReasons[test.data[idx].man.ID] = reason
			}

			for _, found := range snaps {
				reason, ok := expectedReasons[found.ID]
				if !ok {
					// Missing or extra snapshots will be reported by earlier checks.
					continue
				}

				assert.ElementsMatch(
					t,
					reason,
					found.Reasons,
					"incorrect reasons for snapshot with ID %s",
					found.ID,
				)
			}

			// Check number of loads to make sure caching is working properly.
			// Need to manually check because we don't know the order the
			// user/service/category labels will be iterated over. For some tests this
			// could cause more loads than the ideal case.
			assert.Len(t, loadCounts, len(test.expectedLoadCounts))
			for id, count := range loadCounts {
				assert.GreaterOrEqual(t, test.expectedLoadCounts[id], count)
			}
		})
	}
}

func (suite *SnapshotFetchUnitSuite) TestFetchPrevSnapshots_customTags() {
	data := []manifestInfo{
		newManifestInfo(
			testID1,
			testT1,
			false,
			testMail,
			testUser1,
			"fnords",
			"smarf",
		),
	}
	expectLoad1T1 := map[manifest.ID]int{
		testID1: 1,
	}

	table := []struct {
		name  string
		input []Reason
		tags  map[string]string
		// Use this to denote which manifests in data should be expected. Allows
		// defining data in a table while not repeating things between data and
		// expected.
		expectedIdxs []int
		// Expected number of times a manifest should try to be loaded from kopia.
		// Used to check that caching is functioning properly.
		expectedLoadCounts map[manifest.ID]int
	}{
		{
			name:               "no tags specified",
			tags:               nil,
			expectedIdxs:       []int{0},
			expectedLoadCounts: expectLoad1T1,
		},
		{
			name: "all custom tags",
			tags: map[string]string{
				"fnords": "",
				"smarf":  "",
			},
			expectedIdxs:       []int{0},
			expectedLoadCounts: expectLoad1T1,
		},
		{
			name:               "subset of custom tags",
			tags:               map[string]string{"fnords": ""},
			expectedIdxs:       []int{0},
			expectedLoadCounts: expectLoad1T1,
		},
		{
			name:               "custom tag mismatch",
			tags:               map[string]string{"bojangles": ""},
			expectedIdxs:       nil,
			expectedLoadCounts: nil,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			msm := &mockSnapshotManager{
				data: data,
			}

			loadCounts := map[manifest.ID]int{}
			msm.loadCallback = func(ids []manifest.ID) {
				for _, id := range ids {
					loadCounts[id]++
				}
			}

			snaps := fetchPrevSnapshotManifests(ctx, msm, testAllUsersAllCats, test.tags)

			expected := make([]*snapshot.Manifest, 0, len(test.expectedIdxs))
			for _, i := range test.expectedIdxs {
				expected = append(expected, data[i].man)
			}

			got := make([]*snapshot.Manifest, 0, len(snaps))
			for _, s := range snaps {
				got = append(got, s.Manifest)
			}

			assert.ElementsMatch(t, expected, got)

			// Need to manually check because we don't know the order the
			// user/service/category labels will be iterated over. For some tests this
			// could cause more loads than the ideal case.
			assert.Len(t, loadCounts, len(test.expectedLoadCounts))
			for id, count := range loadCounts {
				assert.GreaterOrEqual(t, test.expectedLoadCounts[id], count)
			}
		})
	}
}

// mockErrorSnapshotManager returns an error the first time LoadSnapshot and
// FindSnapshot are called. After that it passes the calls through to the
// contained snapshotManager.
type mockErrorSnapshotManager struct {
	retFindErr bool
	retLoadErr bool
	sm         snapshotManager
}

func (msm *mockErrorSnapshotManager) FindManifests(
	ctx context.Context,
	tags map[string]string,
) ([]*manifest.EntryMetadata, error) {
	if !msm.retFindErr {
		msm.retFindErr = true
		return nil, assert.AnError
	}

	return msm.sm.FindManifests(ctx, tags)
}

func (msm *mockErrorSnapshotManager) LoadSnapshots(
	ctx context.Context,
	ids []manifest.ID,
) ([]*snapshot.Manifest, error) {
	if !msm.retLoadErr {
		msm.retLoadErr = true
		return nil, assert.AnError
	}

	return msm.sm.LoadSnapshots(ctx, ids)
}

func (msm *mockErrorSnapshotManager) LoadSnapshot(
	ctx context.Context,
	id manifest.ID,
) (*snapshot.Manifest, error) {
	return nil, clues.New("not implemented")
}

func (suite *SnapshotFetchUnitSuite) TestFetchPrevSnapshots_withErrors() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	input := testAllUsersMail
	mockData := []manifestInfo{
		newManifestInfo(
			testID1,
			testT1,
			testCompleteMan,
			testMail,
			testUser1,
		),
		newManifestInfo(
			testID2,
			testT2,
			testCompleteMan,
			testMail,
			testUser2,
		),
		newManifestInfo(
			testID3,
			testT3,
			testCompleteMan,
			testMail,
			testUser3,
		),
	}

	msm := &mockErrorSnapshotManager{
		sm: &mockSnapshotManager{
			data: mockData,
		},
	}

	snaps := fetchPrevSnapshotManifests(ctx, msm, input, nil)

	// Only 1 snapshot should be chosen because the other two attempts fail.
	// However, which one is returned is non-deterministic because maps are used.
	assert.Len(t, snaps, 1)
}
