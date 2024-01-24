package backup_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/dttm"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type BackupUnitSuite struct {
	tester.Suite
}

func TestBackupUnitSuite(t *testing.T) {
	suite.Run(t, &BackupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func (suite *BackupUnitSuite) TestBackup_Type() {
	table := []struct {
		name string
		// versionStart is the first version to run this test on. Together with
		// versionEnd this allows more compact test definition when the behavior is
		// the same for multiple versions.
		versionStart int
		// versionEnd is the final version to run this test on. Together with
		// versionStart this allows more compact test definition when the behavior
		// is the same for multiple versions.
		versionEnd int
		// Take a map so that we can have check not having the tag at all vs. having
		// an empty value.
		inputTags map[string]string
		expect    string
	}{
		{
			name:         "NoVersion Returns Empty Type If Untagged",
			versionStart: version.NoBackup,
			versionEnd:   version.NoBackup,
		},
		{
			name:         "PreTag Versions Returns Merge Type If Untagged",
			versionStart: 0,
			versionEnd:   version.All8MigrateUserPNToID,
			expect:       model.MergeBackup,
		},
		{
			name:         "Tag Versions Returns Merge Type If Tagged",
			versionStart: version.All8MigrateUserPNToID,
			versionEnd:   version.Backup,
			inputTags: map[string]string{
				model.BackupTypeTag: model.MergeBackup,
			},
			expect: model.MergeBackup,
		},
		{
			name:         "Tag Versions Returns Assist Type If Tagged",
			versionStart: version.All8MigrateUserPNToID,
			versionEnd:   version.Backup,
			inputTags: map[string]string{
				model.BackupTypeTag: model.AssistBackup,
			},
			expect: model.AssistBackup,
		},
		{
			name:         "Tag Versions Returns Empty Type If Untagged",
			versionStart: version.Groups9Update,
			versionEnd:   version.Backup,
		},
		{
			name:         "All Versions Returns Empty Type If Empty Tag",
			versionStart: 0,
			versionEnd:   version.Backup,
			inputTags: map[string]string{
				model.BackupTypeTag: "",
			},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			for v := test.versionStart; v <= test.versionEnd; v++ {
				suite.Run(fmt.Sprintf("Version%d", v), func() {
					bup := backup.Backup{
						BaseModel: model.BaseModel{
							Tags: test.inputTags,
						},
						Version: v,
					}

					assert.Equal(suite.T(), test.expect, bup.Type())
				})
			}
		})
	}
}

func stubBackup(t time.Time, ownerID, ownerName string) backup.Backup {
	sel := selectors.NewExchangeBackup([]string{"test"})
	sel.Include(sel.AllData())

	return backup.Backup{
		BaseModel: model.BaseModel{
			ID: model.StableID("id"),
			Tags: map[string]string{
				model.ServiceTag: sel.PathService().String(),
			},
		},
		CreationTime:          t,
		SnapshotID:            "snapshot",
		DetailsID:             "details",
		ProtectedResourceID:   ownerID + "-pr",
		ProtectedResourceName: ownerName + "-pr",
		ResourceOwnerID:       ownerID + "-ro",
		ResourceOwnerName:     ownerName + "-ro",
		Status:                "status",
		Selector:              sel.Selector,
		ErrorCount:            2,
		Failure:               "read, write",
		ReadWrites: stats.ReadWrites{
			BytesRead:            301,
			BytesUploaded:        301,
			NonMetaBytesUploaded: 301,
			ItemsRead:            1,
			NonMetaItemsWritten:  1,
			ItemsWritten:         1,
		},
		StartAndEndTime: stats.StartAndEndTime{
			StartedAt:   t,
			CompletedAt: t.Add(1 * time.Minute),
		},
		SkippedCounts: stats.SkippedCounts{
			TotalSkippedItems: 1,
			SkippedMalware:    1,
		},
	}
}

func (suite *BackupUnitSuite) TestBackup_HeadersValues() {
	var (
		t        = suite.T()
		now      = time.Now()
		later    = now.Add(1 * time.Minute)
		b        = stubBackup(now, "id", "name")
		expectHs = []string{
			"ID",
			"Started at",
			"Duration",
			"Status",
			"Protected resource",
			"Data",
		}
		nowFmt   = dttm.FormatToTabularDisplay(now)
		expectVs = []string{
			"id",
			nowFmt,
			"1m0s",
			"status (2 errors, 1 skipped: 1 malware)",
			"name-pr",
			"Contacts,Emails,Events",
		}
	)

	b.StartAndEndTime.CompletedAt = later

	// single skipped malware
	hs := b.Headers(false)
	assert.Equal(t, expectHs, hs)

	vs := b.Values(false)
	assert.Equal(t, expectVs, vs)

	hs = b.Headers(true)
	assert.Equal(t, expectHs[1:], hs)

	vs = b.Values(true)
	assert.Equal(t, expectVs[1:], vs)
}

func (suite *BackupUnitSuite) TestBackup_HeadersValues_onlyResourceOwners() {
	var (
		t        = suite.T()
		now      = time.Now()
		later    = now.Add(1 * time.Minute)
		b        = stubBackup(now, "id", "name")
		expectHs = []string{
			"ID",
			"Started at",
			"Duration",
			"Status",
			"Protected resource",
			"Data",
		}
		nowFmt   = dttm.FormatToTabularDisplay(now)
		expectVs = []string{
			"id",
			nowFmt,
			"1m0s",
			"status (2 errors, 1 skipped: 1 malware)",
			"name-ro",
			"Contacts,Emails,Events",
		}
	)

	b.ProtectedResourceID = ""
	b.ProtectedResourceName = ""

	b.StartAndEndTime.CompletedAt = later

	// single skipped malware
	hs := b.Headers(false)
	assert.Equal(t, expectHs, hs)

	vs := b.Values(false)
	assert.Equal(t, expectVs, vs)

	hs = b.Headers(true)
	assert.Equal(t, expectHs[1:], hs)

	vs = b.Values(true)
	assert.Equal(t, expectVs[1:], vs)
}

func (suite *BackupUnitSuite) TestBackup_Values_statusVariations() {
	table := []struct {
		name   string
		bup    backup.Backup
		expect string
	}{
		{
			name:   "no extras",
			bup:    backup.Backup{Status: "test"},
			expect: "test",
		},
		{
			name: "errors",
			bup: backup.Backup{
				Status:     "test",
				ErrorCount: 42,
			},
			expect: "test (42 errors)",
		},
		{
			name: "malware",
			bup: backup.Backup{
				Status: "test",
				SkippedCounts: stats.SkippedCounts{
					TotalSkippedItems: 2,
					SkippedMalware:    1,
				},
			},
			expect: "test (2 skipped: 1 malware)",
		},
		{
			name: "errors and malware",
			bup: backup.Backup{
				Status:     "test",
				ErrorCount: 42,
				SkippedCounts: stats.SkippedCounts{
					TotalSkippedItems: 1,
					SkippedMalware:    1,
				},
			},
			expect: "test (42 errors, 1 skipped: 1 malware)",
		},
		{
			name: "errors and skipped",
			bup: backup.Backup{
				Status:     "test",
				ErrorCount: 42,
				SkippedCounts: stats.SkippedCounts{
					TotalSkippedItems: 1,
				},
			},
			expect: "test (42 errors, 1 skipped)",
		},
		{
			name: "errors and invalid OneNote",
			bup: backup.Backup{
				Status:     "test",
				ErrorCount: 42,
				SkippedCounts: stats.SkippedCounts{
					TotalSkippedItems:         1,
					SkippedInvalidOneNoteFile: 1,
				},
			},
			expect: "test (42 errors, 1 skipped: 1 invalid OneNote file)",
		},
		{
			name: "errors, malware, notFound, invalid OneNote",
			bup: backup.Backup{
				Status:     "test",
				ErrorCount: 42,
				SkippedCounts: stats.SkippedCounts{
					TotalSkippedItems:         1,
					SkippedMalware:            1,
					SkippedInvalidOneNoteFile: 1,
				},
			},
			expect: "test (42 errors, 1 skipped: 1 malware, 1 invalid OneNote file)",
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			result := test.bup.Values(false)
			assert.Equal(suite.T(), test.expect, result[3], "status value")

			result = test.bup.Values(true)
			assert.Equal(suite.T(), test.expect, result[2], "status value")
		})
	}
}

func (suite *BackupUnitSuite) TestBackup_MinimumPrintable() {
	t := suite.T()
	now := time.Now()
	b := stubBackup(now, "id", "name")

	resultIface := b.MinimumPrintable()
	result, ok := resultIface.(backup.Printable)
	require.True(t, ok)

	assert.Equal(t, b.ID, result.ID, "id")
	assert.Equal(t, 2, result.Stats.ErrorCount, "error count")
	assert.Equal(t, now, result.Stats.StartedAt, "started at")
	assert.Equal(t, b.Status, result.Status, "status")
	assert.Equal(t, b.BytesRead, result.Stats.BytesRead, "size")
	assert.Equal(t, b.NonMetaBytesUploaded, result.Stats.BytesUploaded, "stored size")
	assert.Equal(t, b.Selector.DiscreteOwner, result.Owner, "owner")
}

func (suite *BackupUnitSuite) TestStats() {
	var (
		t     = suite.T()
		start = time.Now()
		b     = stubBackup(start, "owner", "ownername")
		s     = b.ToPrintable().Stats
	)

	assert.Equal(t, b.BytesRead, s.BytesRead, "bytes read")
	assert.Equal(t, b.BytesUploaded, s.BytesUploaded, "bytes uploaded")
	assert.Equal(t, b.CompletedAt, s.EndedAt, "completion time")
	assert.Equal(t, b.ErrorCount, s.ErrorCount, "error count")
	assert.Equal(t, b.ItemsRead, s.ItemsRead, "items read")
	assert.Equal(t, b.TotalSkippedItems, s.ItemsSkipped, "items skipped")
	assert.Equal(t, b.ItemsWritten, s.ItemsWritten, "items written")
	assert.Equal(t, b.StartedAt, s.StartedAt, "started at")
}

func (suite *BackupUnitSuite) TestStats_headersValues() {
	var (
		t     = suite.T()
		start = time.Now()
		b     = stubBackup(start, "owner", "ownername")
		s     = b.ToPrintable().Stats
	)

	expectHeaders := []string{
		"ID",
		"Bytes uploaded",
		"Items uploaded",
		"Items skipped",
		"Errors",
	}

	assert.Equal(t, expectHeaders, s.Headers(false))
	assert.Equal(t, expectHeaders[1:], s.Headers(true))

	expectValues := []string{
		"id",
		humanize.Bytes(uint64(b.BytesUploaded)),
		strconv.Itoa(b.ItemsWritten),
		strconv.Itoa(b.TotalSkippedItems),
		strconv.Itoa(b.ErrorCount),
	}

	assert.Equal(t, expectValues, s.Values(false))
	assert.Equal(t, expectValues[1:], s.Values(true))
}
