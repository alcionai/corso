package kopia

import (
	"fmt"
	"testing"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/path"
)

func makeManifest(id, incmpl, bID string, reasons ...identity.Reasoner) ManifestEntry {
	bIDKey, _ := makeTagKV(TagBackupID)

	return ManifestEntry{
		Manifest: &snapshot.Manifest{
			ID:               manifest.ID(id),
			IncompleteReason: incmpl,
			Tags:             map[string]string{bIDKey: bID},
		},
		Reasons: reasons,
	}
}

type BackupBasesUnitSuite struct {
	tester.Suite
}

func TestBackupBasesUnitSuite(t *testing.T) {
	suite.Run(t, &BackupBasesUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *BackupBasesUnitSuite) TestMinBackupVersion() {
	table := []struct {
		name            string
		bb              *backupBases
		expectedVersion int
	}{
		{
			name:            "Nil BackupBase",
			expectedVersion: version.NoBackup,
		},
		{
			name:            "No Backups",
			bb:              &backupBases{},
			expectedVersion: version.NoBackup,
		},
		{
			name: "Unsorted Backups",
			bb: &backupBases{
				backups: []BackupEntry{
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
			expectedVersion: 0,
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			assert.Equal(suite.T(), test.expectedVersion, test.bb.MinBackupVersion())
		})
	}
}

func (suite *BackupBasesUnitSuite) TestConvertToAssistBase() {
	backups := []BackupEntry{
		{Backup: &backup.Backup{SnapshotID: "1"}},
		{Backup: &backup.Backup{SnapshotID: "2"}},
		{Backup: &backup.Backup{SnapshotID: "3"}},
	}

	merges := []ManifestEntry{
		makeManifest("1", "", ""),
		makeManifest("2", "", ""),
		makeManifest("3", "", ""),
	}

	delID := manifest.ID("3")

	table := []struct {
		name string
		// Below indices specify which items to add from the defined sets above.
		backup       []int
		merge        []int
		assist       []int
		expectAssist []int
	}{
		{
			name:         "Not In Bases",
			backup:       []int{0, 1},
			merge:        []int{0, 1},
			assist:       []int{0, 1},
			expectAssist: []int{0, 1},
		},
		{
			name:         "Different Indexes",
			backup:       []int{2, 0, 1},
			merge:        []int{0, 2, 1},
			assist:       []int{0, 1},
			expectAssist: []int{0, 1, 2},
		},
		{
			name:         "First Item",
			backup:       []int{2, 0, 1},
			merge:        []int{2, 0, 1},
			assist:       []int{0, 1},
			expectAssist: []int{0, 1, 2},
		},
		{
			name:         "Middle Item",
			backup:       []int{0, 2, 1},
			merge:        []int{0, 2, 1},
			assist:       []int{0, 1},
			expectAssist: []int{0, 1, 2},
		},
		{
			name:         "Final Item",
			backup:       []int{0, 1, 2},
			merge:        []int{0, 1, 2},
			assist:       []int{0, 1},
			expectAssist: []int{0, 1, 2},
		},
		{
			name:         "Only In Backups",
			backup:       []int{0, 1, 2},
			merge:        []int{0, 1},
			assist:       []int{0, 1},
			expectAssist: []int{0, 1},
		},
		{
			name:         "Only In Merges",
			backup:       []int{0, 1},
			merge:        []int{0, 1, 2},
			assist:       []int{0, 1},
			expectAssist: []int{0, 1},
		},
		{
			name:         "Only In Assists Noops",
			backup:       []int{0, 1},
			merge:        []int{0, 1},
			assist:       []int{0, 1, 2},
			expectAssist: []int{0, 1, 2},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()
			bb := &backupBases{}

			for _, i := range test.backup {
				bb.backups = append(bb.backups, backups[i])
			}

			for _, i := range test.merge {
				bb.mergeBases = append(bb.mergeBases, merges[i])
			}

			for _, i := range test.assist {
				bb.assistBases = append(bb.assistBases, merges[i])
				bb.assistBackups = append(bb.assistBackups, backups[i])
			}

			expected := &backupBases{
				backups:    []BackupEntry{backups[0], backups[1]},
				mergeBases: []ManifestEntry{merges[0], merges[1]},
			}

			for _, i := range test.expectAssist {
				expected.assistBases = append(expected.assistBases, merges[i])
				expected.assistBackups = append(expected.assistBackups, backups[i])
			}

			bb.ConvertToAssistBase(delID)
			AssertBackupBasesEqual(t, expected, bb)
		})
	}
}

func (suite *BackupBasesUnitSuite) TestDisableMergeBases() {
	t := suite.T()

	merge := []BackupEntry{
		{Backup: &backup.Backup{BaseModel: model.BaseModel{ID: "m1"}}},
		{Backup: &backup.Backup{BaseModel: model.BaseModel{ID: "m2"}}},
	}
	assist := []BackupEntry{
		{Backup: &backup.Backup{BaseModel: model.BaseModel{ID: "a1"}}},
		{Backup: &backup.Backup{BaseModel: model.BaseModel{ID: "a2"}}},
	}

	bb := &backupBases{
		backups:       slices.Clone(merge),
		mergeBases:    make([]ManifestEntry, 2),
		assistBackups: slices.Clone(assist),
		assistBases:   make([]ManifestEntry, 2),
	}

	bb.DisableMergeBases()
	assert.Empty(t, bb.Backups())
	assert.Empty(t, bb.MergeBases())

	// Assist base set now has what used to be merge bases.
	assert.Len(t, bb.UniqueAssistBases(), 4)
	assert.Len(t, bb.SnapshotAssistBases(), 4)
	// Merge bases should appear in the set of backup bases used for details
	// merging.
	assert.ElementsMatch(
		t,
		append(slices.Clone(merge), assist...),
		bb.UniqueAssistBackups())
}

func (suite *BackupBasesUnitSuite) TestDisableAssistBases() {
	t := suite.T()
	bb := &backupBases{
		backups:       make([]BackupEntry, 2),
		mergeBases:    make([]ManifestEntry, 2),
		assistBases:   make([]ManifestEntry, 2),
		assistBackups: make([]BackupEntry, 2),
	}

	bb.DisableAssistBases()
	assert.Empty(t, bb.UniqueAssistBases())
	assert.Empty(t, bb.UniqueAssistBackups())
	assert.Empty(t, bb.SnapshotAssistBases())

	// Merge base should be unchanged.
	assert.Len(t, bb.Backups(), 2)
	assert.Len(t, bb.MergeBases(), 2)
}

func (suite *BackupBasesUnitSuite) TestMergeBackupBases() {
	ro := "resource_owner"

	type testInput struct {
		id  int
		cat []path.CategoryType
	}

	// Make a function so tests can modify things without messing with each other.
	makeBackupBases := func(mergeInputs []testInput, assistInputs []testInput) *backupBases {
		res := &backupBases{}

		for _, i := range mergeInputs {
			baseID := fmt.Sprintf("id%d", i.id)
			reasons := make([]identity.Reasoner, 0, len(i.cat))

			for _, c := range i.cat {
				reasons = append(reasons, NewReason("", ro, path.ExchangeService, c))
			}

			m := makeManifest(baseID, "", "b"+baseID, reasons...)

			b := BackupEntry{
				Backup: &backup.Backup{
					BaseModel:     model.BaseModel{ID: model.StableID("b" + baseID)},
					SnapshotID:    baseID,
					StreamStoreID: "ss" + baseID,
				},
				Reasons: reasons,
			}

			res.backups = append(res.backups, b)
			res.mergeBases = append(res.mergeBases, m)
		}

		for _, i := range assistInputs {
			baseID := fmt.Sprintf("id%d", i.id)

			reasons := make([]identity.Reasoner, 0, len(i.cat))

			for _, c := range i.cat {
				reasons = append(reasons, NewReason("", ro, path.ExchangeService, c))
			}

			m := makeManifest(baseID, "", "a"+baseID, reasons...)

			b := BackupEntry{
				Backup: &backup.Backup{
					BaseModel: model.BaseModel{
						ID:   model.StableID("a" + baseID),
						Tags: map[string]string{model.BackupTypeTag: model.AssistBackup},
					},
					SnapshotID:    baseID,
					StreamStoreID: "ss" + baseID,
				},
				Reasons: reasons,
			}

			res.assistBackups = append(res.assistBackups, b)
			res.assistBases = append(res.assistBases, m)
		}

		return res
	}

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

	makeMan := func(pct path.CategoryType, id, incmpl, bID string) ManifestEntry {
		r := NewReason("", ro, path.ExchangeService, pct)
		return makeManifest(id, incmpl, bID, r)
	}

	// Make a function so tests can modify things without messing with each other.
	validMail1 := func() *backupBases {
		return &backupBases{
			backups: []BackupEntry{
				{
					Backup: &backup.Backup{
						BaseModel: model.BaseModel{
							ID: "bid1",
						},
						SnapshotID:    "id1",
						StreamStoreID: "ssid1",
					},
				},
			},
			mergeBases: []ManifestEntry{
				makeMan(path.EmailCategory, "id1", "", "bid1"),
			},
			assistBackups: []BackupEntry{
				{
					Backup: &backup.Backup{
						BaseModel: model.BaseModel{
							ID:   "bid2",
							Tags: map[string]string{model.BackupTypeTag: model.AssistBackup},
						},
						SnapshotID:    "id2",
						StreamStoreID: "ssid2",
					},
				},
			},
			assistBases: []ManifestEntry{
				makeMan(path.EmailCategory, "id2", "", "bid2"),
			},
		}
	}

	table := []struct {
		name   string
		bb     *backupBases
		expect BackupBases
	}{
		{
			name: "empty BaseBackups",
			bb:   &backupBases{},
		},
		{
			name: "Merge Base Without Backup",
			bb: func() *backupBases {
				res := validMail1()
				res.backups = nil

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.mergeBases = nil
				res.backups = nil

				return res
			}(),
		},
		{
			name: "Merge Backup Missing Snapshot ID",
			bb: func() *backupBases {
				res := validMail1()
				res.backups[0].SnapshotID = ""

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.mergeBases = nil
				res.backups = nil

				return res
			}(),
		},
		{
			name: "Assist backup missing snapshot ID",
			bb: func() *backupBases {
				res := validMail1()
				res.assistBackups[0].SnapshotID = ""

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.assistBases = nil
				res.assistBackups = nil

				return res
			}(),
		},
		{
			name: "Merge backup missing deets ID",
			bb: func() *backupBases {
				res := validMail1()
				res.backups[0].StreamStoreID = ""

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.mergeBases = nil
				res.backups = nil

				return res
			}(),
		},
		{
			name: "Assist backup missing deets ID",
			bb: func() *backupBases {
				res := validMail1()
				res.assistBackups[0].StreamStoreID = ""

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.assistBases = nil
				res.assistBackups = nil

				return res
			}(),
		},
		{
			name: "Incomplete Snapshot",
			bb: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].IncompleteReason = "ir"
				res.assistBases[0].IncompleteReason = "ir"

				return res
			}(),
		},
		{
			name: "Duplicate Reason",
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
		},
		{
			name:   "Single Valid Entry",
			bb:     validMail1(),
			expect: validMail1(),
		},
		{
			name: "Single Valid Entry With Incomplete Assist With Same Reason",
			bb: func() *backupBases {
				res := validMail1()
				res.assistBases = append(
					res.assistBases,
					makeMan(path.EmailCategory, "id3", "checkpoint", "bid3"))

				return res
			}(),
			expect: validMail1(),
		},
		{
			name: "Single Valid Entry With Backup With Old Deets ID",
			bb: func() *backupBases {
				res := validMail1()
				res.backups[0].DetailsID = res.backups[0].StreamStoreID
				res.backups[0].StreamStoreID = ""

				res.assistBackups[0].DetailsID = res.assistBackups[0].StreamStoreID
				res.assistBackups[0].StreamStoreID = ""

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.backups[0].DetailsID = res.backups[0].StreamStoreID
				res.backups[0].StreamStoreID = ""

				res.assistBackups[0].DetailsID = res.assistBackups[0].StreamStoreID
				res.assistBackups[0].StreamStoreID = ""

				return res
			}(),
		},
		{
			name: "Single Valid Entry With Multiple Reasons",
			bb: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].Reasons = append(
					res.mergeBases[0].Reasons,
					NewReason("", ro, path.ExchangeService, path.ContactsCategory))

				res.assistBases[0].Reasons = append(
					res.assistBases[0].Reasons,
					NewReason("", ro, path.ExchangeService, path.ContactsCategory))

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.mergeBases[0].Reasons = append(
					res.mergeBases[0].Reasons,
					NewReason("", ro, path.ExchangeService, path.ContactsCategory))

				res.assistBases[0].Reasons = append(
					res.assistBases[0].Reasons,
					NewReason("", ro, path.ExchangeService, path.ContactsCategory))

				return res
			}(),
		},
		{
			name: "Two Entries Overlapping Reasons",
			bb: func() *backupBases {
				res := validMail1()
				res.mergeBases = append(
					res.mergeBases,
					makeMan(path.EmailCategory, "id3", "", "bid3"))

				res.assistBases = append(
					res.assistBases,
					makeMan(path.EmailCategory, "id4", "", "bid4"))

				return res
			}(),
		},
		{
			name: "Merge Backup, Three Entries One Invalid",
			bb: func() *backupBases {
				res := validMail1()
				res.backups = append(
					res.backups,
					BackupEntry{
						Backup: &backup.Backup{
							BaseModel: model.BaseModel{
								ID: "bid3",
							},
						},
					},
					BackupEntry{
						Backup: &backup.Backup{
							BaseModel: model.BaseModel{
								ID: "bid4",
							},
							SnapshotID:    "id4",
							StreamStoreID: "ssid4",
						},
					})
				res.mergeBases = append(
					res.mergeBases,
					makeMan(path.ContactsCategory, "id3", "checkpoint", "bid3"),
					makeMan(path.EventsCategory, "id4", "", "bid4"))

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.backups = append(
					res.backups,
					BackupEntry{
						Backup: &backup.Backup{
							BaseModel: model.BaseModel{
								ID: "bid4",
							},
							SnapshotID:    "id4",
							StreamStoreID: "ssid4",
						},
					})
				res.mergeBases = append(
					res.mergeBases,
					makeMan(path.EventsCategory, "id4", "", "bid4"))

				return res
			}(),
		},
		{
			name: "Assist Backup, Three Entries One Invalid",
			bb: func() *backupBases {
				res := validMail1()
				res.assistBackups = append(
					res.assistBackups,
					BackupEntry{
						Backup: &backup.Backup{
							BaseModel: model.BaseModel{
								ID:   "bid3",
								Tags: map[string]string{model.BackupTypeTag: model.AssistBackup},
							},
						},
					},
					BackupEntry{
						Backup: &backup.Backup{
							BaseModel: model.BaseModel{
								ID:   "bid4",
								Tags: map[string]string{model.BackupTypeTag: model.AssistBackup},
							},
							SnapshotID:    "id4",
							StreamStoreID: "ssid4",
						},
					})
				res.assistBases = append(
					res.assistBases,
					makeMan(path.ContactsCategory, "id3", "checkpoint", "bid3"),
					makeMan(path.EventsCategory, "id4", "", "bid4"))

				return res
			}(),
			expect: func() *backupBases {
				res := validMail1()
				res.assistBackups = append(
					res.assistBackups,
					BackupEntry{
						Backup: &backup.Backup{
							BaseModel: model.BaseModel{
								ID:   "bid4",
								Tags: map[string]string{model.BackupTypeTag: model.AssistBackup},
							},
							SnapshotID:    "id4",
							StreamStoreID: "ssid4",
						},
					})
				res.assistBases = append(
					res.assistBases,
					makeMan(path.EventsCategory, "id4", "", "bid4"))

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
