package backup_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type BackupUnitSuite struct {
	tester.Suite
}

func TestBackupUnitSuite(t *testing.T) {
	suite.Run(t, &BackupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func stubBackup(t time.Time) backup.Backup {
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
		FailedItems: []fault.Item{*fault.FileErr(
			errors.New("read"),
			"id", "name",
			"containerID", "containerName",
		)},
		SkippedItems: []fault.Skipped{*fault.FileSkip(
			fault.SkipMalware,
			"id", "name",
			"containerID", "containerName",
		)},
		ReadWrites: stats.ReadWrites{
			BytesRead:     301,
			BytesUploaded: 301,
			ItemsRead:     1,
			ItemsWritten:  1,
		},
		StartAndEndTime: stats.StartAndEndTime{
			StartedAt:   t,
			CompletedAt: t,
		},
		// deprecated
		Errs: stats.Errs{
			ReadErrors:  errors.New("1"),
			WriteErrors: errors.New("1"),
		},
	}
}

func (suite *BackupUnitSuite) TestBackup_HeadersValues() {
	var (
		t        = suite.T()
		now      = time.Now()
		b        = stubBackup(now)
		expectHs = []string{
			"Started At",
			"ID",
			"Status",
			"Resource Owner",
		}
		nowFmt   = common.FormatTabularDisplayTime(now)
		expectVs = []string{
			nowFmt,
			"id",
			"status (2 errors, 1 item with malware detected and skipped)",
			"test",
		}
	)

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
				SkippedItems: []fault.Skipped{*fault.FileSkip(
					fault.SkipMalware,
					"id", "name",
					"containerID", "containerName",
				)},
			},
			expect: "test (1 not attempted: 1 malware)",
		},
		{
			name: "errors and malware",
			bup: backup.Backup{
				Status:     "test",
				ErrorCount: 42,
				SkippedItems: []fault.Skipped{*fault.FileSkip(
					fault.SkipMalware,
					"id", "name",
					"containerID", "containerName",
				)},
			},
			expect: "test (42 errors, 1 not attempted: 1 malware)",
		},
	}
	for _, test := range table {
		suite.Run(test.name, func() {
			result := test.bup.Values()
			assert.Equal(suite.T(), test.expect, result[2], "status value")
		})
	}
}

func (suite *BackupUnitSuite) TestBackup_MinimumPrintable() {
	t := suite.T()
	now := time.Now()
	b := stubBackup(now)

	resultIface := b.MinimumPrintable()
	result, ok := resultIface.(backup.Printable)
	require.True(t, ok)

	assert.Equal(t, b.ID, result.ID, "id")
	assert.Equal(t, 2, result.ErrorCount, "error count")
	assert.Equal(t, now, result.StartedAt, "started at")
	assert.Equal(t, b.Status, result.Status, "status")
	assert.Equal(t, b.BytesRead, result.BytesRead, "size")
	assert.Equal(t, b.BytesUploaded, result.BytesUploaded, "stored size")
	assert.Equal(t, b.Selector.DiscreteOwner, result.Owner, "owner")
}
