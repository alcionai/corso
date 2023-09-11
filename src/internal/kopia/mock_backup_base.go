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

		return
	}

	if got == nil {
		if len(expect.Backups()) > 0 ||
			len(expect.MergeBases()) > 0 ||
			len(expect.UniqueAssistBackups()) > 0 ||
			len(expect.UniqueAssistBases()) > 0 {
			assert.Fail(t, "got was nil but expected non-nil result %v", expect)
		}

		return
	}

	assert.ElementsMatch(t, expect.Backups(), got.Backups(), "backups")
	assert.ElementsMatch(t, expect.MergeBases(), got.MergeBases(), "merge bases")
	assert.ElementsMatch(t, expect.UniqueAssistBackups(), got.UniqueAssistBackups(), "assist backups")
	assert.ElementsMatch(t, expect.UniqueAssistBases(), got.UniqueAssistBases(), "assist bases")
}

func NewMockBackupBases() *MockBackupBases {
	return &MockBackupBases{backupBases: &backupBases{}}
}

type MockBackupBases struct {
	*backupBases
}

func (bb *MockBackupBases) WithBackups(b ...BackupEntry) *MockBackupBases {
	bb.backupBases.backups = append(bb.Backups(), b...)
	return bb
}

func (bb *MockBackupBases) WithMergeBases(m ...ManifestEntry) *MockBackupBases {
	bb.backupBases.mergeBases = append(bb.MergeBases(), m...)
	bb.backupBases.assistBases = append(bb.UniqueAssistBases(), m...)

	return bb
}

func (bb *MockBackupBases) WithAssistBackups(b ...BackupEntry) *MockBackupBases {
	bb.backupBases.assistBackups = append(bb.UniqueAssistBackups(), b...)
	return bb
}

func (bb *MockBackupBases) WithAssistBases(m ...ManifestEntry) *MockBackupBases {
	bb.backupBases.assistBases = append(bb.UniqueAssistBases(), m...)
	return bb
}

func (bb *MockBackupBases) ClearMockAssistBases() *MockBackupBases {
	bb.backupBases.DisableAssistBases()
	return bb
}
