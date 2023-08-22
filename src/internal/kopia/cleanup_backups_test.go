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
	bupCurrent := &backup.Backup{
		BaseModel: model.BaseModel{
			ID:           model.StableID("current-bup-id"),
			ModelStoreID: manifest.ID("current-bup-msid"),
		},
		SnapshotID:    "current-snap-msid",
		StreamStoreID: "current-deets-msid",
	}

	snapCurrent := &manifest.EntryMetadata{
		ID: "current-snap-msid",
		Labels: map[string]string{
			backupTag: "0",
		},
	}

	deetsCurrent := &manifest.EntryMetadata{
		ID: "current-deets-msid",
	}

	// Legacy backup with details in separate model.
	bupLegacy := &backup.Backup{
		BaseModel: model.BaseModel{
			ID:           model.StableID("legacy-bup-id"),
			ModelStoreID: manifest.ID("legacy-bup-msid"),
		},
		SnapshotID: "legacy-snap-msid",
		DetailsID:  "legacy-deets-msid",
	}

	snapLegacy := &manifest.EntryMetadata{
		ID: "legacy-snap-msid",
		Labels: map[string]string{
			backupTag: "0",
		},
	}

	deetsLegacy := &model.BaseModel{
		ID:           "legacy-deets-id",
		ModelStoreID: "legacy-deets-msid",
	}

	// Incomplete backup missing data snapshot.
	bupNoSnapshot := &backup.Backup{
		BaseModel: model.BaseModel{
			ID:           model.StableID("ns-bup-id"),
			ModelStoreID: manifest.ID("ns-bup-id-msid"),
		},
		StreamStoreID: "ns-deets-msid",
	}

	deetsNoSnapshot := &manifest.EntryMetadata{
		ID: "ns-deets-msid",
	}

	// Legacy incomplete backup missing data snapshot.
	bupLegacyNoSnapshot := &backup.Backup{
		BaseModel: model.BaseModel{
			ID:           model.StableID("ns-legacy-bup-id"),
			ModelStoreID: manifest.ID("ns-legacy-bup-id-msid"),
		},
		DetailsID: "ns-legacy-deets-msid",
	}

	deetsLegacyNoSnapshot := &model.BaseModel{
		ID:           "ns-legacy-deets-id",
		ModelStoreID: "ns-legacy-deets-msid",
	}

	// Incomplete backup missing details.
	bupNoDetails := &backup.Backup{
		BaseModel: model.BaseModel{
			ID:           model.StableID("nssid-bup-id"),
			ModelStoreID: manifest.ID("nssid-bup-msid"),
		},
		SnapshotID: "nssid-snap-msid",
	}

	snapNoDetails := &manifest.EntryMetadata{
		ID: "nssid-snap-msid",
		Labels: map[string]string{
			backupTag: "0",
		},
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

	backupWithTime := func(mt time.Time, b *backup.Backup) *backup.Backup {
		res := *b
		res.ModTime = mt

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
				snapCurrent,
				deetsCurrent,
				snapLegacy,
			},
			detailsModels: []*model.BaseModel{
				deetsLegacy,
			},
			backups: []backupRes{
				{bup: bupCurrent},
				{bup: bupLegacy},
			},
			time:      baseTime,
			expectErr: assert.NoError,
		},
		{
			name: "MissingFieldsInBackup CausesCleanup",
			snapshots: []*manifest.EntryMetadata{
				snapNoDetails,
				deetsNoSnapshot,
			},
			detailsModels: []*model.BaseModel{
				deetsLegacyNoSnapshot,
			},
			backups: []backupRes{
				{bup: bupNoSnapshot},
				{bup: bupLegacyNoSnapshot},
				{bup: bupNoDetails},
			},
			expectDeleteIDs: []manifest.ID{
				manifest.ID(bupNoSnapshot.ModelStoreID),
				manifest.ID(bupLegacyNoSnapshot.ModelStoreID),
				manifest.ID(bupNoDetails.ModelStoreID),
				manifest.ID(deetsLegacyNoSnapshot.ModelStoreID),
				snapNoDetails.ID,
				deetsNoSnapshot.ID,
			},
			time:      baseTime,
			expectErr: assert.NoError,
		},
		{
			name: "MissingSnapshot CausesCleanup",
			snapshots: []*manifest.EntryMetadata{
				deetsCurrent,
			},
			detailsModels: []*model.BaseModel{
				deetsLegacy,
			},
			backups: []backupRes{
				{bup: bupCurrent},
				{bup: bupLegacy},
			},
			expectDeleteIDs: []manifest.ID{
				manifest.ID(bupCurrent.ModelStoreID),
				deetsCurrent.ID,
				manifest.ID(bupLegacy.ModelStoreID),
				manifest.ID(deetsLegacy.ModelStoreID),
			},
			time:      baseTime,
			expectErr: assert.NoError,
		},
		{
			name: "MissingDetails CausesCleanup",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent,
				snapLegacy,
			},
			backups: []backupRes{
				{bup: bupCurrent},
				{bup: bupLegacy},
			},
			expectDeleteIDs: []manifest.ID{
				manifest.ID(bupCurrent.ModelStoreID),
				manifest.ID(bupLegacy.ModelStoreID),
				snapCurrent.ID,
				snapLegacy.ID,
			},
			time:      baseTime,
			expectErr: assert.NoError,
		},
		{
			name:             "SnapshotsListError Fails",
			snapshotFetchErr: assert.AnError,
			backups: []backupRes{
				{bup: bupCurrent},
			},
			expectErr: assert.Error,
		},
		{
			name: "LegacyDetailsListError Fails",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent,
			},
			detailsModelListErr: assert.AnError,
			backups: []backupRes{
				{bup: bupCurrent},
			},
			time:      baseTime,
			expectErr: assert.Error,
		},
		{
			name: "BackupIDsListError Fails",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent,
				deetsCurrent,
			},
			backupListErr: assert.AnError,
			time:          baseTime,
			expectErr:     assert.Error,
		},
		{
			name: "BackupModelGetErrorNotFound CausesCleanup",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent,
				deetsCurrent,
				snapLegacy,
				snapNoDetails,
			},
			detailsModels: []*model.BaseModel{
				deetsLegacy,
			},
			backups: []backupRes{
				{bup: bupCurrent},
				{
					bup: bupLegacy,
					err: data.ErrNotFound,
				},
				{
					bup: bupNoDetails,
					err: data.ErrNotFound,
				},
			},
			// Backup IDs are still included in here because they're added to the
			// deletion set prior to attempting to fetch models. The model store
			// delete operation should ignore missing models though so there's no
			// issue.
			expectDeleteIDs: []manifest.ID{
				snapLegacy.ID,
				manifest.ID(deetsLegacy.ModelStoreID),
				manifest.ID(bupLegacy.ModelStoreID),
				snapNoDetails.ID,
				manifest.ID(bupNoDetails.ModelStoreID),
			},
			time:      baseTime,
			expectErr: assert.NoError,
		},
		{
			name: "BackupModelGetError Fails",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent,
				deetsCurrent,
				snapLegacy,
				snapNoDetails,
			},
			detailsModels: []*model.BaseModel{
				deetsLegacy,
			},
			backups: []backupRes{
				{bup: bupCurrent},
				{
					bup: bupLegacy,
					err: assert.AnError,
				},
				{bup: bupNoDetails},
			},
			time:      baseTime,
			expectErr: assert.Error,
		},
		{
			name: "DeleteError Fails",
			snapshots: []*manifest.EntryMetadata{
				snapCurrent,
				deetsCurrent,
				snapLegacy,
				snapNoDetails,
			},
			detailsModels: []*model.BaseModel{
				deetsLegacy,
			},
			backups: []backupRes{
				{bup: bupCurrent},
				{bup: bupLegacy},
				{bup: bupNoDetails},
			},
			expectDeleteIDs: []manifest.ID{
				snapNoDetails.ID,
				manifest.ID(bupNoDetails.ModelStoreID),
			},
			deleteErr: assert.AnError,
			time:      baseTime,
			expectErr: assert.Error,
		},
		{
			name: "MissingSnapshot BarelyTooYoungForCleanup Noops",
			snapshots: []*manifest.EntryMetadata{
				manifestWithTime(baseTime, deetsCurrent),
			},
			backups: []backupRes{
				{bup: backupWithTime(baseTime, bupCurrent)},
			},
			time:      baseTime.Add(24 * time.Hour),
			buffer:    24 * time.Hour,
			expectErr: assert.NoError,
		},
		{
			name: "MissingSnapshot BarelyOldEnough CausesCleanup",
			snapshots: []*manifest.EntryMetadata{
				manifestWithTime(baseTime, deetsCurrent),
			},
			backups: []backupRes{
				{bup: backupWithTime(baseTime, bupCurrent)},
			},
			expectDeleteIDs: []manifest.ID{
				deetsCurrent.ID,
				manifest.ID(bupCurrent.ModelStoreID),
			},
			time:      baseTime.Add((24 * time.Hour) + time.Second),
			buffer:    24 * time.Hour,
			expectErr: assert.NoError,
		},
		{
			name: "BackupGetErrorNotFound TooYoung Noops",
			snapshots: []*manifest.EntryMetadata{
				manifestWithTime(baseTime, snapCurrent),
				manifestWithTime(baseTime, deetsCurrent),
			},
			backups: []backupRes{
				{
					bup: backupWithTime(baseTime, bupCurrent),
					err: data.ErrNotFound,
				},
			},
			time:      baseTime,
			buffer:    24 * time.Hour,
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
