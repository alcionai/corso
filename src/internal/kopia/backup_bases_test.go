package kopia

import (
	"testing"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

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

func (suite *BackupBasesUnitSuite) TestRemoveMergeBaseByManifestID() {
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

	expected := &backupBases{
		backups:     []BackupEntry{backups[0], backups[1]},
		mergeBases:  []ManifestEntry{merges[0], merges[1]},
		assistBases: []ManifestEntry{merges[0], merges[1]},
	}

	delID := manifest.ID("3")

	table := []struct {
		name string
		// Below indices specify which items to add from the defined sets above.
		backup []int
		merge  []int
		assist []int
	}{
		{
			name:   "Not In Bases",
			backup: []int{0, 1},
			merge:  []int{0, 1},
			assist: []int{0, 1},
		},
		{
			name:   "Different Indexes",
			backup: []int{2, 0, 1},
			merge:  []int{0, 2, 1},
			assist: []int{0, 1, 2},
		},
		{
			name:   "First Item",
			backup: []int{2, 0, 1},
			merge:  []int{2, 0, 1},
			assist: []int{2, 0, 1},
		},
		{
			name:   "Middle Item",
			backup: []int{0, 2, 1},
			merge:  []int{0, 2, 1},
			assist: []int{0, 2, 1},
		},
		{
			name:   "Final Item",
			backup: []int{0, 1, 2},
			merge:  []int{0, 1, 2},
			assist: []int{0, 1, 2},
		},
		{
			name:   "Only In Backups",
			backup: []int{0, 1, 2},
			merge:  []int{0, 1},
			assist: []int{0, 1},
		},
		{
			name:   "Only In Merges",
			backup: []int{0, 1},
			merge:  []int{0, 1, 2},
			assist: []int{0, 1},
		},
		{
			name:   "Only In Assists",
			backup: []int{0, 1},
			merge:  []int{0, 1},
			assist: []int{0, 1, 2},
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
			}

			bb.RemoveMergeBaseByManifestID(delID)
			AssertBackupBasesEqual(t, expected, bb)
		})
	}
}

func (suite *BackupBasesUnitSuite) TestClearMergeBases() {
	bb := &backupBases{
		backups:    make([]BackupEntry, 2),
		mergeBases: make([]ManifestEntry, 2),
	}

	bb.ClearMergeBases()
	assert.Empty(suite.T(), bb.Backups())
	assert.Empty(suite.T(), bb.MergeBases())
}

func (suite *BackupBasesUnitSuite) TestClearAssistBases() {
	bb := &backupBases{assistBases: make([]ManifestEntry, 2)}

	bb.ClearAssistBases()
	assert.Empty(suite.T(), bb.AssistBases())
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
				res.assistBases = res.mergeBases
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
				res.assistBases = res.mergeBases
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
			name: "Single Valid Entry",
			bb:   validMail1(),
			expect: func() *backupBases {
				res := validMail1()
				res.assistBases = append(res.mergeBases, res.assistBases...)

				return res
			}(),
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
			expect: func() *backupBases {
				res := validMail1()

				res.assistBases = append(res.mergeBases, res.assistBases...)
				return res
			}(),
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

				res.assistBases = append(res.mergeBases, res.assistBases...)

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

				res.assistBases = append(res.mergeBases, res.assistBases...)

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
				res.assistBases = append(res.mergeBases, res.assistBases...)

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

				res.assistBases = append(res.mergeBases, res.assistBases...)

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
