package kopia

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/kopia/kopia/repo/manifest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/path"
)

type BackupCleanupUnitSuite struct {
	tester.Suite
}

func TestBackupCleanupUnitSuite(t *testing.T) {
	suite.Run(t, &BackupCleanupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

type mockManifestFinder struct {
	t         *testing.T
	manifests []*manifest.EntryMetadata
	err       error
}

func (mmf mockManifestFinder) FindManifests(
	ctx context.Context,
	tags map[string]string,
) ([]*manifest.EntryMetadata, error) {
	assert.Equal(
		mmf.t,
		map[string]string{"type": "snapshot"},
		tags,
		"snapshot search tags")

	return mmf.manifests, clues.Stack(mmf.err).OrNil()
}

type mockStorer struct {
	t *testing.T

	details    []*model.BaseModel
	detailsErr error

	backups       []backupRes
	backupListErr error

	expectDeleteIDs []manifest.ID
	deleteErr       error
}

func (ms mockStorer) Delete(context.Context, model.Schema, model.StableID) error {
	return clues.New("not implemented")
}

func (ms mockStorer) Get(context.Context, model.Schema, model.StableID, model.Model) error {
	return clues.New("not implemented")
}

func (ms mockStorer) Put(context.Context, model.Schema, model.Model) error {
	return clues.New("not implemented")
}

func (ms mockStorer) Update(context.Context, model.Schema, model.Model) error {
	return clues.New("not implemented")
}

func (ms mockStorer) GetIDsForType(
	_ context.Context,
	s model.Schema,
	tags map[string]string,
) ([]*model.BaseModel, error) {
	assert.Empty(ms.t, tags, "model search tags")

	switch s {
	case model.BackupDetailsSchema:
		return ms.details, clues.Stack(ms.detailsErr).OrNil()

	case model.BackupSchema:
		var bases []*model.BaseModel

		for _, b := range ms.backups {
			bases = append(bases, &b.bup.BaseModel)
		}

		return bases, clues.Stack(ms.backupListErr).OrNil()
	}

	return nil, clues.New(fmt.Sprintf("unknown type: %s", s.String()))
}

func (ms mockStorer) GetWithModelStoreID(
	_ context.Context,
	s model.Schema,
	id manifest.ID,
	m model.Model,
) error {
	assert.Equal(ms.t, model.BackupSchema, s, "model get schema")

	d := m.(*backup.Backup)

	for _, b := range ms.backups {
		if id == b.bup.ModelStoreID {
			*d = *b.bup
			return clues.Stack(b.err).OrNil()
		}
	}

	return clues.Stack(data.ErrNotFound)
}

func (ms mockStorer) DeleteWithModelStoreIDs(
	_ context.Context,
	ids ...manifest.ID,
) error {
	assert.ElementsMatch(ms.t, ms.expectDeleteIDs, ids, "model delete IDs")
	return clues.Stack(ms.deleteErr).OrNil()
}

// backupRes represents an individual return value for an item in GetIDsForType
// or the result of GetWithModelStoreID. err is used for GetWithModelStoreID
// only.
type backupRes struct {
	bup *backup.Backup
	err error
}

func (suite *BackupCleanupUnitSuite) TestCleanupOrphanedData() {
	backupTag, _ := makeTagKV(TagBackupCategory)

	// Current backup and snapshots.
	bupCurrent := func() *backup.Backup {
		return &backup.Backup{
			BaseModel: model.BaseModel{
				ID:           model.StableID("current-bup-id"),
				ModelStoreID: manifest.ID("current-bup-msid"),
			},
			SnapshotID:    "current-snap-msid",
			StreamStoreID: "current-deets-msid",
		}
	}

	snapCurrent := func() *manifest.EntryMetadata {
		return &manifest.EntryMetadata{
			ID: "current-snap-msid",
			Labels: map[string]string{
				backupTag: "0",
			},
		}
	}

	deetsCurrent := func() *manifest.EntryMetadata {
		return &manifest.EntryMetadata{
			ID: "current-deets-msid",
		}
	}

	bupCurrent2 := func() *backup.Backup {
		return &backup.Backup{
			BaseModel: model.BaseModel{
				ID:           model.StableID("current-bup-id-2"),
				ModelStoreID: manifest.ID("current-bup-msid-2"),
			},
			SnapshotID:    "current-snap-msid-2",
			StreamStoreID: "current-deets-msid-2",
		}
	}

	snapCurrent2 := func() *manifest.EntryMetadata {
		return &manifest.EntryMetadata{
			ID: "current-snap-msid-2",
			Labels: map[string]string{
				backupTag: "0",
			},
		}
	}

	deetsCurrent2 := func() *manifest.EntryMetadata {
		return &manifest.EntryMetadata{
			ID: "current-deets-msid-2",
		}
	}

	bupCurrent3 := func() *backup.Backup {
		return &backup.Backup{
			BaseModel: model.BaseModel{
				ID:           model.StableID("current-bup-id-3"),
				ModelStoreID: manifest.ID("current-bup-msid-3"),
			},
			SnapshotID:    "current-snap-msid-3",
			StreamStoreID: "current-deets-msid-3",
		}
	}

	snapCurrent3 := func() *manifest.EntryMetadata {
		return &manifest.EntryMetadata{
			ID: "current-snap-msid-3",
			Labels: map[string]string{
				backupTag: "0",
			},
		}
	}

	deetsCurrent3 := func() *manifest.EntryMetadata {
		return &manifest.EntryMetadata{
			ID: "current-deets-msid-3",
		}
	}

	// Legacy backup with details in separate model.
	bupLegacy := func() *backup.Backup {
		return &backup.Backup{
			BaseModel: model.BaseModel{
				ID:           model.StableID("legacy-bup-id"),
				ModelStoreID: manifest.ID("legacy-bup-msid"),
			},
			SnapshotID: "legacy-snap-msid",
			DetailsID:  "legacy-deets-msid",
		}
	}

	snapLegacy := func() *manifest.EntryMetadata {
		return &manifest.EntryMetadata{
			ID: "legacy-snap-msid",
			Labels: map[string]string{
				backupTag: "0",
			},
		}
	}

	deetsLegacy := func() *model.BaseModel {
		return &model.BaseModel{
			ID:           "legacy-deets-id",
			ModelStoreID: "legacy-deets-msid",
		}
	}

	// Incomplete backup missing data snapshot.
	bupNoSnapshot := func() *backup.Backup {
		return &backup.Backup{
			BaseModel: model.BaseModel{
				ID:           model.StableID("ns-bup-id"),
				ModelStoreID: manifest.ID("ns-bup-id-msid"),
			},
			StreamStoreID: "ns-deets-msid",
		}
	}

	deetsNoSnapshot := func() *manifest.EntryMetadata {
		return &manifest.EntryMetadata{
			ID: "ns-deets-msid",
		}
	}

	// Legacy incomplete backup missing data snapshot.
	bupLegacyNoSnapshot := func() *backup.Backup {
		return &backup.Backup{
			BaseModel: model.BaseModel{
				ID:           model.StableID("ns-legacy-bup-id"),
				ModelStoreID: manifest.ID("ns-legacy-bup-id-msid"),
			},
			DetailsID: "ns-legacy-deets-msid",
		}
	}

	deetsLegacyNoSnapshot := func() *model.BaseModel {
		return &model.BaseModel{
			ID:           "ns-legacy-deets-id",
			ModelStoreID: "ns-legacy-deets-msid",
		}
	}

	// Incomplete backup missing details.
	bupNoDetails := func() *backup.Backup {
		return &backup.Backup{
			BaseModel: model.BaseModel{
				ID:           model.StableID("nssid-bup-id"),
				ModelStoreID: manifest.ID("nssid-bup-msid"),
			},
			SnapshotID: "nssid-snap-msid",
		}
	}

	snapNoDetails := func() *manifest.EntryMetadata {
		return &manifest.EntryMetadata{
			ID: "nssid-snap-msid",
			Labels: map[string]string{
				backupTag: "0",
			},
		}
	}

	// Get some stable time so that we can do everything relative to this in the
	// tests. Mostly just makes reasoning/viewing times easier because the only
	// differences will be the changes we make.
	baseTime := time.Now()

	manifestWithTime := func(
		mt time.Time,
		m *manifest.EntryMetadata,
	) *manifest.EntryMetadata {
		res := *m
		res.ModTime = mt

		return &res
	}

	manifestWithReasons := func(
		m *manifest.EntryMetadata,
		tenantID string,
		reasons ...identity.Reasoner,
	) *manifest.EntryMetadata {
		res := *m

		if res.Labels == nil {
			res.Labels = map[string]string{}
		}

		res.Labels[kopiaPathLabel] = encodeAsPath(tenantID)

		// Add the given reasons.
		for _, r := range reasons {
			for _, k := range tagKeys(r) {
				key, _ := makeTagKV(k)
				res.Labels[key] = "0"
			}
		}

		// Also add other common reasons on item data snapshots.
		k, _ := makeTagKV(TagBackupCategory)
		res.Labels[k] = "0"

		return &res
	}

	backupWithTime := func(mt time.Time, b *backup.Backup) *backup.Backup {
		res := *b
		res.ModTime = mt
		res.CreationTime = mt

		return &res
	}

	backupWithResource := func(protectedResource string, isAssist bool, b *backup.Backup) *backup.Backup {
		res := *b
		res.ProtectedResourceID = protectedResource

		if isAssist {
			if res.Tags == nil {
				res.Tags = map[string]string{}
			}

			res.Tags[model.BackupTypeTag] = model.AssistBackup
		}

		return &res
	}

	table := []struct {
		name             string
		snapshots        []*manifest.EntryMetadata
		snapshotFetchErr error
		// only need BaseModel here since we never look inside the details items.
		detailsModels       []*model.BaseModel
		detailsModelListErr error
		backups             []backupRes
		backupListErr       error
		deleteErr           error
		time                time.Time
		buffer              time.Duration

		expectDeleteIDs []manifest.ID
		expectErr       assert.ErrorAssertionFunc
	}{
		{
			name:      "EmptyRepo",
			time:      baseTime,
			expectErr: assert.NoError,
		},
		{
			name: "OnlyCompleteBackups Noops",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent(),
				deetsCurrent(),
				snapLegacy(),
			},
			detailsModels: []*model.BaseModel{
				deetsLegacy(),
			},
			backups: []backupRes{
				{bup: bupCurrent()},
				{bup: bupLegacy()},
			},
			time:      baseTime,
			expectErr: assert.NoError,
		},
		{
			name: "MissingFieldsInBackup CausesCleanup",
			snapshots: []*manifest.EntryMetadata{
				snapNoDetails(),
				deetsNoSnapshot(),
			},
			detailsModels: []*model.BaseModel{
				deetsLegacyNoSnapshot(),
			},
			backups: []backupRes{
				{bup: bupNoSnapshot()},
				{bup: bupLegacyNoSnapshot()},
				{bup: bupNoDetails()},
			},
			expectDeleteIDs: []manifest.ID{
				manifest.ID(bupNoSnapshot().ModelStoreID),
				manifest.ID(bupLegacyNoSnapshot().ModelStoreID),
				manifest.ID(bupNoDetails().ModelStoreID),
				manifest.ID(deetsLegacyNoSnapshot().ModelStoreID),
				snapNoDetails().ID,
				deetsNoSnapshot().ID,
			},
			time:      baseTime,
			expectErr: assert.NoError,
		},
		{
			name: "MissingSnapshot CausesCleanup",
			snapshots: []*manifest.EntryMetadata{
				deetsCurrent(),
			},
			detailsModels: []*model.BaseModel{
				deetsLegacy(),
			},
			backups: []backupRes{
				{bup: bupCurrent()},
				{bup: bupLegacy()},
			},
			expectDeleteIDs: []manifest.ID{
				manifest.ID(bupCurrent().ModelStoreID),
				deetsCurrent().ID,
				manifest.ID(bupLegacy().ModelStoreID),
				manifest.ID(deetsLegacy().ModelStoreID),
			},
			time:      baseTime,
			expectErr: assert.NoError,
		},
		{
			name: "MissingDetails CausesCleanup",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent(),
				snapLegacy(),
			},
			backups: []backupRes{
				{bup: bupCurrent()},
				{bup: bupLegacy()},
			},
			expectDeleteIDs: []manifest.ID{
				manifest.ID(bupCurrent().ModelStoreID),
				manifest.ID(bupLegacy().ModelStoreID),
				snapCurrent().ID,
				snapLegacy().ID,
			},
			time:      baseTime,
			expectErr: assert.NoError,
		},
		// Tests with various errors from Storer.
		{
			name:             "SnapshotsListError Fails",
			snapshotFetchErr: assert.AnError,
			backups: []backupRes{
				{bup: bupCurrent()},
			},
			expectErr: assert.Error,
		},
		{
			name: "LegacyDetailsListError Fails",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent(),
			},
			detailsModelListErr: assert.AnError,
			backups: []backupRes{
				{bup: bupCurrent()},
			},
			time:      baseTime,
			expectErr: assert.Error,
		},
		{
			name: "BackupIDsListError Fails",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent(),
				deetsCurrent(),
			},
			backupListErr: assert.AnError,
			time:          baseTime,
			expectErr:     assert.Error,
		},
		{
			name: "BackupModelGetErrorNotFound CausesCleanup",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent(),
				deetsCurrent(),
				snapLegacy(),
				snapNoDetails(),
			},
			detailsModels: []*model.BaseModel{
				deetsLegacy(),
			},
			backups: []backupRes{
				{bup: bupCurrent()},
				{
					bup: bupLegacy(),
					err: data.ErrNotFound,
				},
				{
					bup: bupNoDetails(),
					err: data.ErrNotFound,
				},
			},
			// Backup IDs are still included in here because they're added to the
			// deletion set prior to attempting to fetch models. The model store
			// delete operation should ignore missing models though so there's no
			// issue.
			expectDeleteIDs: []manifest.ID{
				snapLegacy().ID,
				manifest.ID(deetsLegacy().ModelStoreID),
				manifest.ID(bupLegacy().ModelStoreID),
				snapNoDetails().ID,
				manifest.ID(bupNoDetails().ModelStoreID),
			},
			time:      baseTime,
			expectErr: assert.NoError,
		},
		{
			name: "BackupModelGetError Fails",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent(),
				deetsCurrent(),
				snapLegacy(),
				snapNoDetails(),
			},
			detailsModels: []*model.BaseModel{
				deetsLegacy(),
			},
			backups: []backupRes{
				{bup: bupCurrent()},
				{
					bup: bupLegacy(),
					err: assert.AnError,
				},
				{bup: bupNoDetails()},
			},
			time:      baseTime,
			expectErr: assert.Error,
		},
		{
			name: "DeleteError Fails",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent(),
				deetsCurrent(),
				snapLegacy(),
				snapNoDetails(),
			},
			detailsModels: []*model.BaseModel{
				deetsLegacy(),
			},
			backups: []backupRes{
				{bup: bupCurrent()},
				{bup: bupLegacy()},
				{bup: bupNoDetails()},
			},
			expectDeleteIDs: []manifest.ID{
				snapNoDetails().ID,
				manifest.ID(bupNoDetails().ModelStoreID),
			},
			deleteErr: assert.AnError,
			time:      baseTime,
			expectErr: assert.Error,
		},
		// Tests dealing with buffer times.
		{
			name: "MissingSnapshot BarelyTooYoungForCleanup Noops",
			snapshots: []*manifest.EntryMetadata{
				manifestWithTime(baseTime, deetsCurrent()),
			},
			backups: []backupRes{
				{bup: backupWithTime(baseTime, bupCurrent())},
			},
			time:      baseTime.Add(24 * time.Hour),
			buffer:    24 * time.Hour,
			expectErr: assert.NoError,
		},
		{
			name: "MissingSnapshot BarelyOldEnough CausesCleanup",
			snapshots: []*manifest.EntryMetadata{
				manifestWithTime(baseTime, deetsCurrent()),
			},
			backups: []backupRes{
				{bup: backupWithTime(baseTime, bupCurrent())},
			},
			expectDeleteIDs: []manifest.ID{
				deetsCurrent().ID,
				manifest.ID(bupCurrent().ModelStoreID),
			},
			time:      baseTime.Add((24 * time.Hour) + time.Second),
			buffer:    24 * time.Hour,
			expectErr: assert.NoError,
		},
		{
			name: "BackupGetErrorNotFound TooYoung Noops",
			snapshots: []*manifest.EntryMetadata{
				manifestWithTime(baseTime, snapCurrent()),
				manifestWithTime(baseTime, deetsCurrent()),
			},
			backups: []backupRes{
				{
					bup: backupWithTime(baseTime, bupCurrent()),
					err: data.ErrNotFound,
				},
			},
			time:      baseTime,
			buffer:    24 * time.Hour,
			expectErr: assert.NoError,
		},
		// Tests dealing with assist base cleanup.
		{
			// Test that even if we have multiple assist bases with the same
			// Reason(s), none of them are garbage collected if they are within the
			// buffer period used to exclude recently created backups from garbage
			// collection.
			name: "AssistBase NotYoungest InBufferTime Noops",
			snapshots: []*manifest.EntryMetadata{
				manifestWithReasons(
					manifestWithTime(baseTime, snapCurrent()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime, deetsCurrent()),

				manifestWithReasons(
					manifestWithTime(baseTime.Add(time.Second), snapCurrent2()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime.Add(time.Second), deetsCurrent2()),
			},
			backups: []backupRes{
				{bup: backupWithResource("ro", true, backupWithTime(baseTime, bupCurrent()))},
				{bup: backupWithResource("ro", true, backupWithTime(baseTime.Add(time.Second), bupCurrent2()))},
			},
			time:      baseTime,
			buffer:    24 * time.Hour,
			expectErr: assert.NoError,
		},
		{
			// Test that an assist base that has the same Reasons as a newer assist
			// base is garbage collected when it's outside the buffer period.
			name: "AssistBases NotYoungest CausesCleanup",
			snapshots: []*manifest.EntryMetadata{
				manifestWithReasons(
					manifestWithTime(baseTime, snapCurrent()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime, deetsCurrent()),

				manifestWithReasons(
					manifestWithTime(baseTime.Add(time.Second), snapCurrent2()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime.Add(time.Second), deetsCurrent2()),

				manifestWithReasons(
					manifestWithTime(baseTime.Add(time.Minute), snapCurrent3()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime.Add(time.Minute), deetsCurrent3()),
			},
			backups: []backupRes{
				{bup: backupWithResource("ro", true, backupWithTime(baseTime, bupCurrent()))},
				{bup: backupWithResource("ro", true, backupWithTime(baseTime.Add(time.Second), bupCurrent2()))},
				{bup: backupWithResource("ro", true, backupWithTime(baseTime.Add(time.Minute), bupCurrent3()))},
			},
			expectDeleteIDs: []manifest.ID{
				snapCurrent().ID,
				deetsCurrent().ID,
				manifest.ID(bupCurrent().ModelStoreID),
				snapCurrent2().ID,
				deetsCurrent2().ID,
				manifest.ID(bupCurrent2().ModelStoreID),
			},
			time:      baseTime.Add(48 * time.Hour),
			buffer:    24 * time.Hour,
			expectErr: assert.NoError,
		},
		{
			// Test that the most recent assist base is not garbage collected even if
			// there's a newer merge base that has the same Reasons as the assist
			// base. Also ensure assist bases with the same Reasons that are older
			// than the newest assist base are still garbage collected.
			name: "AssistBasesAndMergeBase NotYoungest CausesCleanupForAssistBase",
			snapshots: []*manifest.EntryMetadata{
				manifestWithReasons(
					manifestWithTime(baseTime, snapCurrent()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime, deetsCurrent()),

				manifestWithReasons(
					manifestWithTime(baseTime.Add(time.Second), snapCurrent2()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime.Add(time.Second), deetsCurrent2()),

				manifestWithReasons(
					manifestWithTime(baseTime.Add(time.Minute), snapCurrent3()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime.Add(time.Minute), deetsCurrent3()),
			},
			backups: []backupRes{
				{bup: backupAssist("ro", backupWithTime(baseTime, bupCurrent()))},
				{bup: backupAssist("ro", backupWithTime(baseTime.Add(time.Second), bupCurrent2()))},
				{bup: backupWithTime(baseTime.Add(time.Minute), bupCurrent3())},
			},
			expectDeleteIDs: []manifest.ID{
				snapCurrent().ID,
				deetsCurrent().ID,
				manifest.ID(bupCurrent().ModelStoreID),
			},
			time:      baseTime.Add(48 * time.Hour),
			buffer:    24 * time.Hour,
			expectErr: assert.NoError,
		},
		{
			// Test that an assist base that is not the most recent for Reason A but
			// is the most recent for Reason B is not garbage collected.
			name: "AssistBases YoungestInOneReason Noops",
			snapshots: []*manifest.EntryMetadata{
				manifestWithReasons(
					manifestWithTime(baseTime, snapCurrent()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory),
					NewReason("", "ro", path.ExchangeService, path.ContactsCategory)),
				manifestWithTime(baseTime, deetsCurrent()),

				manifestWithReasons(
					manifestWithTime(baseTime.Add(time.Second), snapCurrent2()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime.Add(time.Second), deetsCurrent2()),
			},
			backups: []backupRes{
				{bup: backupWithResource("ro", true, backupWithTime(baseTime, bupCurrent()))},
				{bup: backupWithResource("ro", true, backupWithTime(baseTime.Add(time.Second), bupCurrent2()))},
			},
			time:      baseTime.Add(48 * time.Hour),
			buffer:    24 * time.Hour,
			expectErr: assert.NoError,
		},
		{
			// Test that assist bases that have the same tenant, service, and category
			// but different protected resources are not garbage collected. This is
			// a test to ensure the Reason field is properly handled when finding the
			// most recent assist base.
			name: "AssistBases DifferentProtectedResources Noops",
			snapshots: []*manifest.EntryMetadata{
				manifestWithReasons(
					manifestWithTime(baseTime, snapCurrent()),
					"tenant1",
					NewReason("", "ro1", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime, deetsCurrent()),

				manifestWithReasons(
					manifestWithTime(baseTime.Add(time.Second), snapCurrent2()),
					"tenant1",
					NewReason("", "ro2", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime.Add(time.Second), deetsCurrent2()),
			},
			backups: []backupRes{
				{bup: backupWithResource("ro1", true, backupWithTime(baseTime, bupCurrent()))},
				{bup: backupWithResource("ro2", true, backupWithTime(baseTime.Add(time.Second), bupCurrent2()))},
			},
			time:      baseTime.Add(48 * time.Hour),
			buffer:    24 * time.Hour,
			expectErr: assert.NoError,
		},
		{
			// Test that assist bases that have the same protected resource, service,
			// and category but different tenants are not garbage collected. This is a
			// test to ensure the Reason field is properly handled when finding the
			// most recent assist base.
			name: "AssistBases DifferentTenants Noops",
			snapshots: []*manifest.EntryMetadata{
				manifestWithReasons(
					manifestWithTime(baseTime, snapCurrent()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime, deetsCurrent()),

				manifestWithReasons(
					manifestWithTime(baseTime.Add(time.Second), snapCurrent2()),
					"tenant2",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime.Add(time.Second), deetsCurrent2()),
			},
			backups: []backupRes{
				{bup: backupWithResource("ro", true, backupWithTime(baseTime, bupCurrent()))},
				{bup: backupWithResource("ro", true, backupWithTime(baseTime.Add(time.Second), bupCurrent2()))},
			},
			time:      baseTime.Add(48 * time.Hour),
			buffer:    24 * time.Hour,
			expectErr: assert.NoError,
		},
		{
			// Test that if the tenant is not available for a given assist base that
			// it's excluded from the garbage collection set. This behavior is
			// conservative because it's quite likely that we could garbage collect
			// the base without issue.
			name: "AssistBases NoTenant SkipsBackup",
			snapshots: []*manifest.EntryMetadata{
				manifestWithReasons(
					manifestWithTime(baseTime, snapCurrent()),
					"",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime, deetsCurrent()),

				manifestWithReasons(
					manifestWithTime(baseTime.Add(time.Second), snapCurrent2()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime.Add(time.Second), deetsCurrent2()),

				manifestWithReasons(
					manifestWithTime(baseTime.Add(time.Minute), snapCurrent3()),
					"tenant1",
					NewReason("", "ro", path.ExchangeService, path.EmailCategory)),
				manifestWithTime(baseTime.Add(time.Minute), deetsCurrent3()),
			},
			backups: []backupRes{
				{bup: backupWithResource("ro", true, backupWithTime(baseTime, bupCurrent()))},
				{bup: backupWithResource("ro", true, backupWithTime(baseTime.Add(time.Second), bupCurrent2()))},
				{bup: backupWithResource("ro", true, backupWithTime(baseTime.Add(time.Minute), bupCurrent3()))},
			},
			time:   baseTime.Add(48 * time.Hour),
			buffer: 24 * time.Hour,
			expectDeleteIDs: []manifest.ID{
				snapCurrent2().ID,
				deetsCurrent2().ID,
				manifest.ID(bupCurrent2().ModelStoreID),
			},
			expectErr: assert.NoError,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			mbs := mockStorer{
				t:               t,
				details:         test.detailsModels,
				detailsErr:      test.detailsModelListErr,
				backups:         test.backups,
				backupListErr:   test.backupListErr,
				expectDeleteIDs: test.expectDeleteIDs,
				deleteErr:       test.deleteErr,
			}

			mmf := mockManifestFinder{
				t:         t,
				manifests: test.snapshots,
				err:       test.snapshotFetchErr,
			}

			err := cleanupOrphanedData(
				ctx,
				mbs,
				mmf,
				test.buffer,
				func() time.Time { return test.time })
			test.expectErr(t, err, clues.ToCore(err))
		})
	}
}
