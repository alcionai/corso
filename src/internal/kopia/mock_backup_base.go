package kopia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

	assert.ElementsMatch(t, expect.Backups(), got.Backups(), "backups")
	assert.ElementsMatch(t, expect.MergeBases(), got.MergeBases(), "merge bases")
	assert.ElementsMatch(t, expect.UniqueAssistBackups(), got.UniqueAssistBackups(), "assist backups")
	assert.ElementsMatch(t, expect.UniqueAssistBases(), got.UniqueAssistBases(), "assist bases")
	assert.ElementsMatch(t, expect.SnapshotAssistBases(), got.SnapshotAssistBases(), "snapshot assist bases")
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
