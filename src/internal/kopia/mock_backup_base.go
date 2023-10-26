package kopia

import (
	"fmt"
	"testing"

	"github.com/kopia/kopia/repo/manifest"
	"github.com/kopia/kopia/snapshot"
	"github.com/stretchr/testify/assert"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/path"
)

func basesMatch(t *testing.T, expect, got []BackupBase, dataType string) {
	expectBups := make([]*backup.Backup, 0, len(expect))
	expectMans := make([]*snapshot.Manifest, 0, len(expect))
	gotBups := make([]*backup.Backup, 0, len(got))
	gotMans := make([]*snapshot.Manifest, 0, len(got))
	gotBasesByID := map[model.StableID]BackupBase{}

	for _, e := range expect {
		expectBups = append(expectBups, e.Backup)
		expectMans = append(expectMans, e.ItemDataSnapshot)
	}

	for _, g := range got {
		gotBups = append(gotBups, g.Backup)
		gotMans = append(gotMans, g.ItemDataSnapshot)
		gotBasesByID[g.Backup.ID] = g
	}

	assert.ElementsMatch(t, expectBups, gotBups, dataType+" backup model")
	assert.ElementsMatch(t, expectMans, gotMans, dataType+" item data snapshot")

	// Need to compare Reasons separately since they're also a slice.
	for _, e := range expect {
		b, ok := gotBasesByID[e.Backup.ID]
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
		assert.Empty(t, got.MergeBases(), "merge bases")
		assert.Empty(t, got.UniqueAssistBases(), "assist bases")
		assert.Empty(t, got.SnapshotAssistBases(), "snapshot assist bases")

		return
	}

	if got == nil {
		if len(expect.MergeBases()) > 0 ||
			len(expect.UniqueAssistBases()) > 0 ||
			len(expect.SnapshotAssistBases()) > 0 {
			assert.Fail(t, "got was nil but expected non-nil result %v", expect)
		}

		return
	}

	basesMatch(t, expect.MergeBases(), got.MergeBases(), "merge bases")
	basesMatch(t, expect.UniqueAssistBases(), got.UniqueAssistBases(), "assist bases")
	basesMatch(t, expect.SnapshotAssistBases(), got.SnapshotAssistBases(), "snapshot assist bases")
}

func NewMockBackupBases() *MockBackupBases {
	return &MockBackupBases{backupBases: &backupBases{}}
}

type MockBackupBases struct {
	*backupBases
}

func (bb *MockBackupBases) WithMergeBases(b ...BackupBase) *MockBackupBases {
	bb.backupBases.mergeBases = append(bb.MergeBases(), b...)
	return bb
}

func (bb *MockBackupBases) WithAssistBases(b ...BackupBase) *MockBackupBases {
	bb.backupBases.assistBases = append(bb.UniqueAssistBases(), b...)
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
				identity.NewReason(
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

func (builder *BackupBaseBuilder) Build() BackupBase {
	return *builder.b
}

func (builder *BackupBaseBuilder) MarkAssistBase() *BackupBaseBuilder {
	if builder.b.Backup.Tags == nil {
		builder.b.Backup.Tags = map[string]string{}
	}

	builder.b.Backup.Tags[model.BackupTypeTag] = model.AssistBackup

	return builder
}

func (builder *BackupBaseBuilder) WithReasons(
	reasons ...identity.Reasoner,
) *BackupBaseBuilder {
	builder.b.Reasons = reasons
	return builder
}

func (builder *BackupBaseBuilder) AppendReasons(
	reasons ...identity.Reasoner,
) *BackupBaseBuilder {
	builder.b.Reasons = append(builder.b.Reasons, reasons...)
	return builder
}
