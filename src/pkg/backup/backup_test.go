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

type BackupSuite struct {
	tester.Suite
}

func TestBackupSuite(t *testing.T) {
	suite.Run(t, &BackupSuite{Suite: tester.NewUnitSuite(t)})
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
		Errors: fault.ErrorsData{
			Errs: []error{errors.New("read"), errors.New("write")},
		},
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
		"Resource Owner",
	}
	hs := b.Headers()
	assert.Equal(t, expectHs, hs)

	nowFmt := common.FormatTabularDisplayTime(now)
	expectVs := []string{
		nowFmt,
		"id",
		"status (2 errors)",
		"test",
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
	assert.Equal(t, b.BytesRead, result.BytesRead, "size")
	assert.Equal(t, b.BytesUploaded, result.BytesUploaded, "stored size")
}
