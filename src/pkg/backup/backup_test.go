package backup_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/zeebo/assert"

	"github.com/alcionai/corso/internal/common"
	"github.com/alcionai/corso/internal/model"
	"github.com/alcionai/corso/internal/stats"
	"github.com/alcionai/corso/pkg/backup"
	"github.com/alcionai/corso/pkg/backup/details"
	"github.com/alcionai/corso/pkg/selectors"
)

type BackupSuite struct {
	suite.Suite
}

func TestBackupSuite(t *testing.T) {
	suite.Run(t, new(BackupSuite))
}

func (suite *BackupSuite) TestBackup_HeadersValues() {
	t := suite.T()
	now := time.Now()

	b := backup.Backup{
		BaseModel: model.BaseModel{
			StableID: model.ID("stable"),
		},
		CreationTime: now,
		SnapshotID:   "snapshot",
		DetailsID:    "details",
		Status:       "status",
		Selectors:    selectors.Selector{},
		ReadWrites: stats.ReadWrites{
			ItemsRead:    1,
			ItemsWritten: 1,
			ReadErrors:   errors.New("1"),
			WriteErrors:  errors.New("1"),
		},
		StartAndEndTime: stats.StartAndEndTime{
			StartedAt:   now,
			CompletedAt: now,
		},
	}

	expectHs := []string{
		"Creation Time",
		"Stable ID",
		"Snapshot ID",
		"Details ID",
		"Status",
		"Selectors",
		"Items Read",
		"Items Written",
		"Read Errors",
		"Write Errors",
		"Started At",
		"Completed At",
	}
	hs := b.Headers()
	assert.DeepEqual(t, expectHs, hs)
	nowFmt := common.FormatTime(now)

	expectVs := []string{
		nowFmt,
		"stable",
		"snapshot",
		"details",
		"status",
		"{}",
		"1",
		"1",
		"1",
		"1",
		nowFmt,
		nowFmt,
	}
	vs := b.Values()
	assert.DeepEqual(t, expectVs, vs)
}

func (suite *BackupSuite) TestDetailsEntry_HeadersValues() {
	now := time.Now()
	nowStr := now.Format(time.RFC3339Nano)

	table := []struct {
		name     string
		entry    details.DetailsEntry
		expectHs []string
		expectVs []string
	}{
		{
			name: "no info",
			entry: details.DetailsEntry{
				RepoRef: "reporef",
			},
			expectHs: []string{"Repo Ref"},
			expectVs: []string{"reporef"},
		},
		{
			name: "exhange info",
			entry: details.DetailsEntry{
				RepoRef: "reporef",
				ItemInfo: details.ItemInfo{
					Exchange: &details.ExchangeInfo{
						Sender:   "sender",
						Subject:  "subject",
						Received: now,
					},
				},
			},
			expectHs: []string{"Repo Ref", "Sender", "Subject", "Received"},
			expectVs: []string{"reporef", "sender", "subject", nowStr},
		},
		{
			name: "sharepoint info",
			entry: details.DetailsEntry{
				RepoRef: "reporef",
				ItemInfo: details.ItemInfo{
					Sharepoint: &details.SharepointInfo{},
				},
			},
			expectHs: []string{"Repo Ref"},
			expectVs: []string{"reporef"},
		},
	}

	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			hs := test.entry.Headers()
			assert.DeepEqual(t, test.expectHs, hs)
			vs := test.entry.Values()
			assert.DeepEqual(t, test.expectVs, vs)
		})
	}
}

func (suite *BackupSuite) TestDetailsModel_Path() {
	table := []struct {
		name   string
		ents   []details.DetailsEntry
		expect []string
	}{
		{
			name:   "nil entries",
			ents:   nil,
			expect: []string{},
		},
		{
			name: "single entry",
			ents: []details.DetailsEntry{
				{RepoRef: "abcde"},
			},
			expect: []string{"abcde"},
		},
		{
			name: "multiple entries",
			ents: []details.DetailsEntry{
				{RepoRef: "abcde"},
				{RepoRef: "12345"},
			},
			expect: []string{"abcde", "12345"},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			d := details.Details{
				DetailsModel: details.DetailsModel{
					Entries: test.ents,
				},
			}
			assert.DeepEqual(t, test.expect, d.Paths())
		})
	}
}
