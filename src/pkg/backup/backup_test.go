package backup_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type BackupUnitSuite struct {
	tester.Suite
}

func TestBackupUnitSuite(t *testing.T) {
	suite.Run(t, &BackupUnitSuite{Suite: tester.NewUnitSuite(t)})
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
		CreationTime: t,
		SnapshotID:   "snapshot",
		DetailsID:    "details",
		Status:       "status",
		Selector:     sel.Selector,
		ErrorCount:   2,
		Failure:      "read, write",
		ReadWrites: stats.ReadWrites{
			BytesRead:     301,
			BytesUploaded: 301,
			ItemsRead:     1,
			ItemsWritten:  1,
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
			"Started At",
			"Duration",
			"Status",
			"Resource Owner",
		}
		nowFmt   = dttm.FormatToTabularDisplay(now)
		expectVs = []string{
			"id",
			nowFmt,
			"1m0s",
			"status (2 errors, 1 skipped: 1 malware)",
			"test",
		}
	)

	b.StartAndEndTime.CompletedAt = later

	// single skipped malware
	hs := b.Headers()
	assert.Equal(t, expectHs, hs)

	vs := b.Values()
	assert.Equal(t, expectVs, vs)
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
			name: "not found",
			bup: backup.Backup{
				Status: "test",
				SkippedCounts: stats.SkippedCounts{
					TotalSkippedItems: 2,
					SkippedNotFound:   1,
				},
			},
			expect: "test (2 skipped: 1 not found)",
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
			name: "errors and not found",
			bup: backup.Backup{
				Status:     "test",
				ErrorCount: 42,
				SkippedCounts: stats.SkippedCounts{
					TotalSkippedItems: 1,
					SkippedNotFound:   1,
				},
			},
			expect: "test (42 errors, 1 skipped: 1 not found)",
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
					SkippedNotFound:           1,
					SkippedInvalidOneNoteFile: 1,
				},
			},
			expect: "test (42 errors, 1 skipped: 1 malware, 1 not found, 1 invalid OneNote file)",
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			result := test.bup.Values()
			assert.Equal(suite.T(), test.expect, result[3], "status value")
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
	assert.Equal(t, b.BytesUploaded, result.Stats.BytesUploaded, "stored size")
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
		"Bytes Uploaded",
		"Items Uploaded",
		"Items Skipped",
		"Errors",
	}

	assert.Equal(t, expectHeaders, s.Headers())

	expectValues := []string{
		"id",
		humanize.Bytes(uint64(b.BytesUploaded)),
		strconv.Itoa(b.ItemsWritten),
		strconv.Itoa(b.TotalSkippedItems),
		strconv.Itoa(b.ErrorCount),
	}

	assert.Equal(t, expectValues, s.Values())
}
