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
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type BackupSuite struct {
	suite.Suite
}

func TestBackupSuite(t *testing.T) {
	suite.Run(t, new(BackupSuite))
}

func stubBackup(t time.Time) backup.Backup {
	sel := selectors.NewExchangeBackup(selectors.Any())
	sel.Include(sel.Users(selectors.Any()))

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
		Selectors:    sel.Selector,
		Errs: stats.Errs{
			ReadErrors:  errors.New("1"),
			WriteErrors: errors.New("1"),
		},
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
	}
}

func (suite *BackupSuite) TestBackup_HeadersValues() {
	t := suite.T()
	now := time.Now()
	b := stubBackup(now)

	expectHs := []string{
		"Started At",
		"ID",
		"Status",
		"Selectors",
	}
	hs := b.Headers()
	assert.Equal(t, expectHs, hs)

	nowFmt := common.FormatTabularDisplayTime(now)
	expectVs := []string{
		nowFmt,
		"id",
		"status (2 errors)",
		selectors.All,
	}

	vs := b.Values()
	assert.Equal(t, expectVs, vs)
}

func (suite *BackupSuite) TestBackup_MinimumPrintable() {
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

	bselp := b.Selectors.ToPrintable()
	assert.Equal(t, bselp, result.Selectors, "selectors")
	assert.Equal(t, bselp.Resources(), result.Selectors.Resources(), "selector resources")

	assert.Equal(t, b.BytesRead, result.BytesRead, "size")
	assert.Equal(t, b.BytesUploaded, result.BytesUploaded, "stored size")
}
