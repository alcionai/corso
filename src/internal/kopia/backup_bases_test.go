package kopia

import (
	"fmt"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/path"
)

func makeManifest(id, incmpl, bID string) *snapshot.Manifest {
	bIDKey, _ := makeTagKV(TagBackupID)

	return &snapshot.Manifest{
		ID:               manifest.ID(id),
		IncompleteReason: incmpl,
		Tags:             map[string]string{bIDKey: bID},
	}
}

type testInput struct {
	tenant            string
	protectedResource string
	id                int
	cat               []path.CategoryType
	isAssist          bool
	incompleteReason  string
}

func makeBase(ti testInput) BackupBase {
	baseID := fmt.Sprintf("id%d", ti.id)
	reasons := make([]identity.Reasoner, 0, len(ti.cat))

	for _, c := range ti.cat {
		reasons = append(
			reasons,
			identity.NewReason(
				ti.tenant,
				ti.protectedResource,
				path.ExchangeService,
				c))
	}

	return BackupBase{
		Backup: &backup.Backup{
			BaseModel:     model.BaseModel{ID: model.StableID("b" + baseID)},
			SnapshotID:    baseID,
			StreamStoreID: "ss" + baseID,
		},
		ItemDataSnapshot: makeManifest(baseID, ti.incompleteReason, "b"+baseID),
		Reasons:          reasons,
	}
}

// Make a function so tests can modify things without messing with each other.
func makeBackupBases(
	mergeInputs []testInput,
	assistInputs []testInput,
) *backupBases {
	res := &backupBases{}

	for _, i := range mergeInputs {
		res.mergeBases = append(res.mergeBases, makeBase(i))
	}

	for _, i := range assistInputs {
		i.isAssist = true
		res.assistBases = append(res.assistBases, makeBase(i))
	}

	return res
}

type BackupBasesUnitSuite struct {
	tester.Suite
}

func TestBackupBasesUnitSuite(t *testing.T) {
	suite.Run(t, &BackupBasesUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *BackupBasesUnitSuite) TestBackupBases_minVersions() {
	table := []struct {
		name                  string
		bb                    *backupBases
		expectedBackupVersion int
		expectedAssistVersion int
	}{
		{
			name:                  "Nil BackupBase",
			expectedBackupVersion: version.NoBackup,
			expectedAssistVersion: version.NoBackup,
		},
		{
			name:                  "No Backups",
			bb:                    &backupBases{},
			expectedBackupVersion: version.NoBackup,
			expectedAssistVersion: version.NoBackup,
		},
		{
			name: "Unsorted Backups",
			bb: &backupBases{
				mergeBases: []BackupBase{
					{
						Backup: &backup.Backup{
							Version: 4,
						},
					},
					{
						Backup: &backup.Backup{
							Version: 0,
						},
					},
					{
						Backup: &backup.Backup{
							Version: 2,
						},
					},
				},
			},
			expectedBackupVersion: 0,
			expectedAssistVersion: version.NoBackup,
		},
		{
			name: "Only Assist Bases",
			bb: &backupBases{
				assistBases: []BackupBase{
					{
						Backup: &backup.Backup{
							Version: 4,
						},
					},
					{
						Backup: &backup.Backup{
							Version: 0,
						},
					},
					{
						Backup: &backup.Backup{
							Version: 2,
						},
					},
				},
			},
			expectedBackupVersion: version.NoBackup,
			expectedAssistVersion: 0,
		},
		{
			name: "Assist and Merge Bases, min merge",
			bb: &backupBases{
				mergeBases: []BackupBase{
					{
						Backup: &backup.Backup{
							Version: 4,
						},
					},
					{
						Backup: &backup.Backup{
							Version: 2,
						},
					},
					{
						Backup: &backup.Backup{
							Version: 3,
						},
					},
				},
				assistBases: []BackupBase{
					{
						Backup: &backup.Backup{
							Version: 4,
						},
					},
					{
						Backup: &backup.Backup{
							Version: 1,
						},
					},
					{
						Backup: &backup.Backup{
							Version: 3,
						},
					},
				},
			},
			expectedBackupVersion: 2,
			expectedAssistVersion: 1,
		},
		{
			name: "Assist and Merge Bases, min assist",
			bb: &backupBases{
				mergeBases: []BackupBase{
					{
						Backup: &backup.Backup{
							Version: 4,
						},
					},
					{
						Backup: &backup.Backup{
							Version: 3,
						},
					},
					{
						Backup: &backup.Backup{
							Version: 2,
						},
					},
				},
				assistBases: []BackupBase{
					{
						Backup: &backup.Backup{
							Version: 4,
						},
					},
					{
						Backup: &backup.Backup{
							Version: 1,
						},
					},
					{
						Backup: &backup.Backup{
							Version: 3,
						},
					},
				},
			},
			expectedBackupVersion: 2,
			expectedAssistVersion: 1,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			assert.Equal(t, test.expectedBackupVersion, test.bb.MinBackupVersion(), "backup")
			assert.Equal(t, test.expectedAssistVersion, test.bb.MinAssistVersion(), "assist")
		})
	}
}

func (suite *BackupBasesUnitSuite) TestConvertToAssistBase() {
	bases := []BackupBase{
		{
			Backup: &backup.Backup{
				BaseModel: model.BaseModel{
					ID: "1",
				},
				SnapshotID: "its1",
			},
			ItemDataSnapshot: makeManifest("its1", "", ""),
		},
		{
			Backup: &backup.Backup{
				BaseModel: model.BaseModel{
					ID: "2",
				},
				SnapshotID: "its2",
			},
			ItemDataSnapshot: makeManifest("its2", "", ""),
		},
		{
			Backup: &backup.Backup{
				BaseModel: model.BaseModel{
					ID: "3",
				},
				SnapshotID: "its3",
			},
			ItemDataSnapshot: makeManifest("its3", "", ""),
		},
	}

	delID := model.StableID("3")

	table := []struct {
		name string
		// Below indices specify which items to add from the defined sets above.
		merge        []int
		assist       []int
		expectMerge  []int
		expectAssist []int
	}{
		{
			name:         "Not In Bases",
			merge:        []int{0, 1},
			assist:       []int{0, 1},
			expectMerge:  []int{0, 1},
			expectAssist: []int{0, 1},
		},
		{
			name:         "First Item",
			merge:        []int{2, 0, 1},
			assist:       []int{0, 1},
			expectMerge:  []int{0, 1},
			expectAssist: []int{0, 1, 2},
		},
		{
			name:         "Middle Item",
			merge:        []int{0, 2, 1},
			assist:       []int{0, 1},
			expectMerge:  []int{0, 1},
			expectAssist: []int{0, 1, 2},
		},
		{
			name:         "Final Item",
			merge:        []int{0, 1, 2},
			assist:       []int{0, 1},
			expectMerge:  []int{0, 1},
			expectAssist: []int{0, 1, 2},
		},
		{
			name:         "Only In Assists Noops",
			merge:        []int{0, 1},
			assist:       []int{0, 1, 2},
			expectMerge:  []int{0, 1},
			expectAssist: []int{0, 1, 2},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			bb := &backupBases{}

			for _, i := range test.merge {
				bb.mergeBases = append(bb.mergeBases, bases[i])
			}

			for _, i := range test.assist {
				bb.assistBases = append(bb.assistBases, bases[i])
			}

			expected := &backupBases{}

			for _, i := range test.expectMerge {
				expected.mergeBases = append(expected.mergeBases, bases[i])
			}

			for _, i := range test.expectAssist {
				expected.assistBases = append(expected.assistBases, bases[i])
			}

			bb.ConvertToAssistBase(delID)
			AssertBackupBasesEqual(t, expected, bb)
		})
	}
}

func (suite *BackupBasesUnitSuite) TestDisableMergeBases() {
	t := suite.T()

	merge := []BackupBase{
		{
			Backup:           &backup.Backup{BaseModel: model.BaseModel{ID: "m1"}},
			ItemDataSnapshot: &snapshot.Manifest{ID: "ms1"},
		},
		{
			Backup:           &backup.Backup{BaseModel: model.BaseModel{ID: "m2"}},
			ItemDataSnapshot: &snapshot.Manifest{ID: "ms2"},
		},
	}

	assist := []BackupBase{
		{
			Backup:           &backup.Backup{BaseModel: model.BaseModel{ID: "a1"}},
			ItemDataSnapshot: &snapshot.Manifest{ID: "as1"},
		},
		{
			Backup:           &backup.Backup{BaseModel: model.BaseModel{ID: "a2"}},
			ItemDataSnapshot: &snapshot.Manifest{ID: "as2"},
		},
	}

	bb := &backupBases{
		mergeBases:  merge,
		assistBases: assist,
	}

	bb.DisableMergeBases()
	assert.Empty(t, bb.MergeBases())

	// Merge bases should still appear in the assist base set passed in for kopia
	// snapshots and details merging.
	assert.ElementsMatch(
		t,
		append(slices.Clone(merge), assist...),
		bb.SnapshotAssistBases())

	assert.ElementsMatch(
		t,
		append(slices.Clone(merge), assist...),
		bb.UniqueAssistBases())
}

func (suite *BackupBasesUnitSuite) TestDisableAssistBases() {
	t := suite.T()
	bb := &backupBases{
		mergeBases:  make([]BackupBase, 2),
		assistBases: make([]BackupBase, 2),
	}

	bb.DisableAssistBases()
	assert.Empty(t, bb.UniqueAssistBases())
	assert.Empty(t, bb.SnapshotAssistBases())

	// Merge base should be unchanged.
	assert.Len(t, bb.MergeBases(), 2)
}

func (suite *BackupBasesUnitSuite) TestMergeBackupBases() {
	table := []struct {
		name        string
		merge       []testInput
		assist      []testInput
		otherMerge  []testInput
		otherAssist []testInput
		expect      func() *backupBases
	}{
		{
			name: "Other Empty",
			merge: []testInput{
				{cat: []path.CategoryType{path.EmailCategory}},
			},
			assist: []testInput{
				{cat: []path.CategoryType{path.EmailCategory}},
			},
			expect: func() *backupBases {
				bs := makeBackupBases([]testInput{
					{cat: []path.CategoryType{path.EmailCategory}},
				}, []testInput{
					{cat: []path.CategoryType{path.EmailCategory}},
				})

				return bs
			},
		},
		{
			name: "current Empty",
			otherMerge: []testInput{
				{cat: []path.CategoryType{path.EmailCategory}},
			},
			otherAssist: []testInput{
				{cat: []path.CategoryType{path.EmailCategory}},
			},
			expect: func() *backupBases {
				bs := makeBackupBases([]testInput{
					{cat: []path.CategoryType{path.EmailCategory}},
				}, []testInput{
					{cat: []path.CategoryType{path.EmailCategory}},
				})

				return bs
			},
		},
		{
			name: "Other overlaps merge and assist",
			merge: []testInput{
				{
					id:  1,
					cat: []path.CategoryType{path.EmailCategory},
				},
			},
			assist: []testInput{
				{
					id:  4,
					cat: []path.CategoryType{path.EmailCategory},
				},
			},
			otherMerge: []testInput{
				{
					id:  2,
					cat: []path.CategoryType{path.EmailCategory},
				},
				{
					id:  3,
					cat: []path.CategoryType{path.EmailCategory},
				},
			},
			otherAssist: []testInput{
				{
					id:  5,
					cat: []path.CategoryType{path.EmailCategory},
				},
			},
			expect: func() *backupBases {
				bs := makeBackupBases([]testInput{
					{
						id:  1,
						cat: []path.CategoryType{path.EmailCategory},
					},
				}, []testInput{
					{
						id:  4,
						cat: []path.CategoryType{path.EmailCategory},
					},
				})

				return bs
			},
		},
		{
			name: "Other overlaps merge",
			merge: []testInput{
				{
					id:  1,
					cat: []path.CategoryType{path.EmailCategory},
				},
			},
			otherMerge: []testInput{
				{
					id:  2,
					cat: []path.CategoryType{path.EmailCategory},
				},
			},
			expect: func() *backupBases {
				bs := makeBackupBases([]testInput{
					{
						id:  1,
						cat: []path.CategoryType{path.EmailCategory},
					},
				}, nil)

				return bs
			},
		},
		{
			name: "Current assist overlaps with Other merge",
			assist: []testInput{
				{
					id:  3,
					cat: []path.CategoryType{path.EmailCategory},
				},
			},
			otherMerge: []testInput{
				{
					id:  1,
					cat: []path.CategoryType{path.EmailCategory},
				},
			},
			otherAssist: []testInput{
				{
					id:  2,
					cat: []path.CategoryType{path.EmailCategory},
				},
			},

			expect: func() *backupBases {
				bs := makeBackupBases([]testInput{
					{
						id:  1,
						cat: []path.CategoryType{path.EmailCategory},
					},
				}, []testInput{
					{
						id:  3,
						cat: []path.CategoryType{path.EmailCategory},
					},
				})

				return bs
			},
		},
		{
			name: "Other Disjoint",
			merge: []testInput{
				{cat: []path.CategoryType{path.EmailCategory}},
				{
					id:  1,
					cat: []path.CategoryType{path.EmailCategory},
				},
			},
			otherMerge: []testInput{
				{
					id:  2,
					cat: []path.CategoryType{path.ContactsCategory},
				},
			},
			expect: func() *backupBases {
				bs := makeBackupBases([]testInput{
					{cat: []path.CategoryType{path.EmailCategory}},
					{
						id:  1,
						cat: []path.CategoryType{path.EmailCategory},
					},
					{
						id:  2,
						cat: []path.CategoryType{path.ContactsCategory},
					},
				}, nil)

				return bs
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			bb := makeBackupBases(test.merge, test.assist)
			other := makeBackupBases(test.otherMerge, test.otherAssist)
			expected := test.expect()

			ctx, flush := tester.NewContext(t)
			defer flush()

			got := bb.MergeBackupBases(
				ctx,
				other,
				func(r identity.Reasoner) string {
					return r.Service().String() + r.Category().String()
				})
			AssertBackupBasesEqual(t, expected, got)
		})
	}
}

func (suite *BackupBasesUnitSuite) TestFixupAndVerify() {
	ro := "resource_owner"

	// Make a function so tests can modify things without messing with each other.
	validMail1 := func() *backupBases {
		return makeBackupBases(
			[]testInput{
				{
					tenant:            "t",
					protectedResource: ro,
					id:                1,
					cat:               []path.CategoryType{path.EmailCategory},
				},
			},
			[]testInput{
				{
					tenant:            "t",
					protectedResource: ro,
					id:                2,
					cat:               []path.CategoryType{path.EmailCategory},
				},
			})
	}

	table := []struct {
		name   string
		bb     *backupBases
		expect BackupBases
	}{
		{
			name: "EmptyBaseBackups",
			bb:   &backupBases{},
		},
		{
			name: "MergeBase MissingBackup",
			bb: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].Backup = nil

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.mergeBases = nil

				return res
			}(),
		},
		{
			name: "MergeBase MissingSnapshot",
			bb: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].ItemDataSnapshot = nil

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.mergeBases = nil

				return res
			}(),
		},
		{
			name: "AssistBase MissingBackup",
			bb: func() *backupBases {
				res := validMail1()
				res.assistBases[0].Backup = nil

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.assistBases = nil

				return res
			}(),
		},
		{
			name: "AssistBase MissingSnapshot",
			bb: func() *backupBases {
				res := validMail1()
				res.assistBases[0].ItemDataSnapshot = nil

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.assistBases = nil

				return res
			}(),
		},
		{
			name: "MergeBase MissingDeetsID",
			bb: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].Backup.StreamStoreID = ""

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.mergeBases = nil

				return res
			}(),
		},
		{
			name: "AssistBase MissingDeetsID",
			bb: func() *backupBases {
				res := validMail1()
				res.assistBases[0].Backup.StreamStoreID = ""

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.assistBases = nil

				return res
			}(),
		},
		{
			name: "MergeAndAssistBase IncompleteSnapshot",
			bb: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].ItemDataSnapshot.IncompleteReason = "ir"
				res.assistBases[0].ItemDataSnapshot.IncompleteReason = "ir"

				return res
			}(),
		},
		{
			name: "MergeAndAssistBase DuplicateReasonInBase",
			bb: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].Reasons = append(
					res.mergeBases[0].Reasons,
					res.mergeBases[0].Reasons[0])

				res.assistBases[0].Reasons = append(
					res.assistBases[0].Reasons,
					res.assistBases[0].Reasons[0])

				return res
			}(),
			expect: validMail1(),
		},
		{
			name: "MergeAndAssistBase MissingReason",
			bb: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].Reasons = nil
				res.assistBases[0].Reasons = nil

				return res
			}(),
		},
		{
			name:   "SingleValidEntry",
			bb:     validMail1(),
			expect: validMail1(),
		},
		{
			name: "SingleValidEntry IncompleteAssistWithSameReason",
			bb: func() *backupBases {
				res := validMail1()
				res.assistBases = append(
					res.assistBases,
					makeBase(testInput{
						tenant:            "t",
						protectedResource: "ro",
						id:                3,
						cat:               []path.CategoryType{path.EmailCategory},
						isAssist:          true,
						incompleteReason:  "checkpoint",
					}))

				return res
			}(),
			expect: validMail1(),
		},
		{
			name: "SingleValidEntry BackupWithOldDeetsID",
			bb: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].Backup.DetailsID = res.mergeBases[0].Backup.StreamStoreID
				res.mergeBases[0].Backup.StreamStoreID = ""

				res.assistBases[0].Backup.DetailsID = res.assistBases[0].Backup.StreamStoreID
				res.assistBases[0].Backup.StreamStoreID = ""

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].Backup.DetailsID = res.mergeBases[0].Backup.StreamStoreID
				res.mergeBases[0].Backup.StreamStoreID = ""

				res.assistBases[0].Backup.DetailsID = res.assistBases[0].Backup.StreamStoreID
				res.assistBases[0].Backup.StreamStoreID = ""

				return res
			}(),
		},
		{
			name: "SingleValidEntry MultipleReasons",
			bb: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].Reasons = append(
					res.mergeBases[0].Reasons,
					identity.NewReason("t", ro, path.ExchangeService, path.ContactsCategory))

				res.assistBases[0].Reasons = append(
					res.assistBases[0].Reasons,
					identity.NewReason("t", ro, path.ExchangeService, path.ContactsCategory))

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].Reasons = append(
					res.mergeBases[0].Reasons,
					identity.NewReason("t", ro, path.ExchangeService, path.ContactsCategory))

				res.assistBases[0].Reasons = append(
					res.assistBases[0].Reasons,
					identity.NewReason("t", ro, path.ExchangeService, path.ContactsCategory))

				return res
			}(),
		},
		{
			name: "TwoEntries OverlappingReasons",
			bb: func() *backupBases {
				t1, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
				require.NoError(suite.T(), err, clues.ToCore(err))

				res := validMail1()
				res.mergeBases[0].Backup.CreationTime = t1
				res.assistBases[0].Backup.CreationTime = t1.Add(2 * time.Hour)

				other := makeBase(testInput{
					tenant:            "t",
					protectedResource: ro,
					id:                3,
					cat:               []path.CategoryType{path.EmailCategory},
				})
				other.Backup.CreationTime = t1.Add(-time.Minute)

				res.mergeBases = append(res.mergeBases, other)

				other = makeBase(testInput{
					tenant:            "t",
					protectedResource: ro,
					id:                4,
					cat:               []path.CategoryType{path.EmailCategory},
					isAssist:          true,
				})
				other.Backup.CreationTime = t1.Add(time.Hour)

				res.assistBases = append(res.assistBases, other)

				return res
			}(),
			expect: func() *backupBases {
				t1, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
				require.NoError(suite.T(), err, clues.ToCore(err))

				res := validMail1()
				res.mergeBases[0].Backup.CreationTime = t1
				res.assistBases[0].Backup.CreationTime = t1.Add(2 * time.Hour)

				return res
			}(),
		},
		{
			name: "TwoEntries OverlappingReasons OneInvalid",
			bb: func() *backupBases {
				t1, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
				require.NoError(suite.T(), err, clues.ToCore(err))

				res := validMail1()
				res.mergeBases[0].Backup.CreationTime = t1
				res.assistBases[0].Backup.CreationTime = t1.Add(2 * time.Hour)

				other := makeBase(testInput{
					tenant:            "t",
					protectedResource: ro,
					id:                3,
					cat:               []path.CategoryType{path.EmailCategory},
				})
				other.Backup.CreationTime = t1.Add(time.Minute)
				other.Backup.StreamStoreID = ""

				res.mergeBases = append(res.mergeBases, other)

				other = makeBase(testInput{
					tenant:            "t",
					protectedResource: ro,
					id:                4,
					cat:               []path.CategoryType{path.EmailCategory},
					isAssist:          true,
				})
				other.Backup.CreationTime = t1.Add(3 * time.Hour)
				other.Backup.StreamStoreID = ""

				res.assistBases = append(res.assistBases, other)

				return res
			}(),
			expect: func() *backupBases {
				t1, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
				require.NoError(suite.T(), err, clues.ToCore(err))

				res := validMail1()
				res.mergeBases[0].Backup.CreationTime = t1
				res.assistBases[0].Backup.CreationTime = t1.Add(2 * time.Hour)

				return res
			}(),
		},
		{
			name: "MergeBase ThreeEntriesOneInvalid",
			bb: func() *backupBases {
				invalid := makeBase(testInput{
					tenant:            "t",
					protectedResource: ro,
					id:                3,
					cat:               []path.CategoryType{path.ContactsCategory},
					incompleteReason:  "checkpoint",
				})
				invalid.Backup.StreamStoreID = ""
				invalid.Backup.SnapshotID = ""

				res := validMail1()
				res.mergeBases = append(
					res.mergeBases,
					invalid,
					makeBase(testInput{
						tenant:            "t",
						protectedResource: ro,
						id:                4,
						cat:               []path.CategoryType{path.EventsCategory},
					}))

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.mergeBases = append(
					res.mergeBases,
					makeBase(testInput{
						tenant:            "t",
						protectedResource: ro,
						id:                4,
						cat:               []path.CategoryType{path.EventsCategory},
					}))

				return res
			}(),
		},
		{
			name: "AssistBase ThreeEntriesOneInvalid",
			bb: func() *backupBases {
				invalid := makeBase(testInput{
					tenant:            "t",
					protectedResource: ro,
					id:                3,
					cat:               []path.CategoryType{path.ContactsCategory},
					isAssist:          true,
					incompleteReason:  "checkpoint",
				})
				invalid.Backup.StreamStoreID = ""
				invalid.Backup.SnapshotID = ""

				res := validMail1()
				res.assistBases = append(
					res.assistBases,
					invalid,
					makeBase(testInput{
						tenant:            "t",
						protectedResource: ro,
						id:                4,
						cat:               []path.CategoryType{path.EventsCategory},
						isAssist:          true,
					}))

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.assistBases = append(
					res.assistBases,
					makeBase(testInput{
						tenant:            "t",
						protectedResource: ro,
						id:                4,
						cat:               []path.CategoryType{path.EventsCategory},
						isAssist:          true,
					}))

				return res
			}(),
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			ctx, flush := tester.NewContext(suite.T())
			defer flush()

			test.bb.fixupAndVerify(ctx)
			AssertBackupBasesEqual(suite.T(), test.expect, test.bb)
		})
	}
}
