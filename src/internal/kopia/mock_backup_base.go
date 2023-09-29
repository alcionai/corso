package kopia

import (
	"testing"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
)

// TODO(ashmrtn): Temp function until all PRs in the series merge.
func backupsMatch(t *testing.T, expect, got []BackupEntry, dataType string) {
	expectBups := make([]*backup.Backup, 0, len(expect))
	gotBups := make([]*backup.Backup, 0, len(got))
	gotBasesByID := map[model.StableID]BackupEntry{}

	for _, e := range expect {
		if e.Backup != nil {
			expectBups = append(expectBups, e.Backup)
		}
	}

	for _, g := range got {
		if g.Backup != nil {
			gotBups = append(gotBups, g.Backup)
			gotBasesByID[g.Backup.ID] = g
		}
	}

	assert.ElementsMatch(t, expectBups, gotBups, dataType+" backup model")

	// Need to compare Reasons separately since they're also a slice.
	for _, e := range expect {
		if e.Backup == nil {
			continue
		}

		b, ok := gotBasesByID[e.Backup.ID]
		if !ok {
			// Missing bases will be reported above.
			continue
		}

		assert.ElementsMatch(t, e.Reasons, b.Reasons)
	}
}

// TODO(ashmrtn): Temp function until all PRs in the series merge.
func manifestsMatch(t *testing.T, expect, got []ManifestEntry, dataType string) {
	expectMans := make([]*snapshot.Manifest, 0, len(expect))
	gotMans := make([]*snapshot.Manifest, 0, len(got))
	gotBasesByID := map[manifest.ID]ManifestEntry{}

	for _, e := range expect {
		if e.Manifest != nil {
			expectMans = append(expectMans, e.Manifest)
		}
	}

	for _, g := range got {
		if g.Manifest != nil {
			gotMans = append(gotMans, g.Manifest)
			gotBasesByID[g.Manifest.ID] = g
		}
	}

	assert.ElementsMatch(t, expectMans, gotMans, dataType+" item data snapshot")

	// Need to compare Reasons separately since they're also a slice.
	for _, e := range expect {
		if e.Manifest == nil {
			continue
		}

		b, ok := gotBasesByID[e.Manifest.ID]
		if !ok {
			// Missing bases will be reported above.
			continue
		}

		assert.ElementsMatch(t, e.Reasons, b.Reasons)
	}
}

func AssertBackupBasesEqual(t *testing.T, expect, got BackupBases) {
	if expect == nil && got == nil {
		return
	}

	if expect == nil {
		assert.Empty(t, got.Backups(), "backups")
		assert.Empty(t, got.MergeBases(), "merge bases")
		assert.Empty(t, got.UniqueAssistBackups(), "assist backups")
		assert.Empty(t, got.UniqueAssistBases(), "assist bases")
		assert.Empty(t, got.SnapshotAssistBases(), "snapshot assist bases")

		return
	}

	if got == nil {
		if len(expect.Backups()) > 0 ||
			len(expect.MergeBases()) > 0 ||
			len(expect.UniqueAssistBackups()) > 0 ||
			len(expect.UniqueAssistBases()) > 0 ||
			len(expect.SnapshotAssistBases()) > 0 {
			assert.Fail(t, "got was nil but expected non-nil result %v", expect)
		}

		return
	}

	backupsMatch(t, expect.Backups(), got.Backups(), "merge backups")
	manifestsMatch(t, expect.MergeBases(), got.MergeBases(), "merge manifests")
	backupsMatch(t, expect.UniqueAssistBackups(), got.UniqueAssistBackups(), "assist backups")
	manifestsMatch(t, expect.UniqueAssistBases(), got.UniqueAssistBases(), "assist manifests")
	manifestsMatch(t, expect.SnapshotAssistBases(), got.SnapshotAssistBases(), "snapshot assist bases")
}

func NewMockBackupBases() *MockBackupBases {
	return &MockBackupBases{backupBases: &backupBases{}}
}

type MockBackupBases struct {
	*backupBases
}

func (bb *MockBackupBases) WithBackups(b ...BackupEntry) *MockBackupBases {
	bases := make([]BackupBase, 0, len(b))
	for _, base := range b {
		bases = append(bases, BackupBase{
			Backup:  base.Backup,
			Reasons: base.Reasons,
		})
	}

	bb.backupBases.mergeBases = append(bb.NewMergeBases(), bases...)

	return bb
}

func (bb *MockBackupBases) WithMergeBases(m ...ManifestEntry) *MockBackupBases {
	bases := make([]BackupBase, 0, len(m))
	for _, base := range m {
		bases = append(bases, BackupBase{
			ItemDataSnapshot: base.Manifest,
			Reasons:          base.Reasons,
		})
	}

	bb.backupBases.mergeBases = append(bb.NewMergeBases(), bases...)

	return bb
}

func (bb *MockBackupBases) WithAssistBackups(b ...BackupEntry) *MockBackupBases {
	bases := make([]BackupBase, 0, len(b))
	for _, base := range b {
		bases = append(bases, BackupBase{
			Backup:  base.Backup,
			Reasons: base.Reasons,
		})
	}

	bb.backupBases.assistBases = append(bb.NewUniqueAssistBases(), bases...)

	return bb
}

func (bb *MockBackupBases) WithAssistBases(m ...ManifestEntry) *MockBackupBases {
	bases := make([]BackupBase, 0, len(m))
	for _, base := range m {
		bases = append(bases, BackupBase{
			ItemDataSnapshot: base.Manifest,
			Reasons:          base.Reasons,
		})
	}

	bb.backupBases.assistBases = append(bb.NewUniqueAssistBases(), bases...)

	return bb
}

func (bb *MockBackupBases) NewWithMergeBases(b ...BackupBase) *MockBackupBases {
	bb.backupBases.mergeBases = append(bb.NewMergeBases(), b...)
	return bb
}

func (bb *MockBackupBases) MockDisableAssistBases() *MockBackupBases {
	bb.DisableAssistBases()
	return bb
}

func (bb *MockBackupBases) MockDisableMergeBases() *MockBackupBases {
	bb.DisableMergeBases()
	return bb
}

// -----------------------------------------------------------------------------
// Functions for BackupBase creation
// -----------------------------------------------------------------------------

func NewBackupBaseBuilder(idPrefix string, id int) *BackupBaseBuilder {
	bIDKey, _ := makeTagKV(TagBackupID)
	baseID := fmt.Sprintf("%sID%d", idPrefix, id)

	return &BackupBaseBuilder{
		b: &BackupBase{
			Backup: &backup.Backup{
				BaseModel: model.BaseModel{
					ID: model.StableID(baseID + "-backup"),
				},
				SnapshotID:    baseID + "-item-data",
				StreamStoreID: baseID + "-stream-store",
			},
			ItemDataSnapshot: &snapshot.Manifest{
				ID:   manifest.ID(baseID + "-item-data"),
				Tags: map[string]string{bIDKey: baseID + "-backup"},
			},
			Reasons: []identity.Reasoner{
				NewReason(
					"tenant",
					"protected_resource",
					path.ExchangeService,
					path.EmailCategory),
			},
		},
	}
}

type BackupBaseBuilder struct {
	b *BackupBase
}

func (bbb *BackupBaseBuilder) Build() BackupBase {
	return *bbb.b
}

func (bbb *BackupBaseBuilder) MarkAssistBase() *BackupBaseBuilder {
	if bbb.b.Backup.Tags == nil {
		bbb.b.Backup.Tags = map[string]string{}
	}

	bbb.b.Backup.Tags[model.BackupTypeTag] = model.AssistBackup

	return bbb
}

func (bbb *BackupBaseBuilder) WithReasons(reasons ...identity.Reasoner) *BackupBaseBuilder {
	bbb.b.Reasons = reasons
	return bbb
}

func (bbb *BackupBaseBuilder) AppendReasons(reasons ...identity.Reasoner) *BackupBaseBuilder {
	bbb.b.Reasons = append(bbb.b.Reasons, reasons...)
	return bbb
}
