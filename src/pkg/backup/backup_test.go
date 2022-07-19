package backup_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/zeebo/assert"

	"github.com/alcionai/corso/internal/model"
	"github.com/alcionai/corso/pkg/backup"
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
	}

	expectHs := []string{
		"Creation Time",
		"Stable ID",
		"Snapshot ID",
		"Details ID",
	}
	hs := b.Headers()
	assert.DeepEqual(t, expectHs, hs)

	expectVs := []string{
		now.Format(time.RFC3339Nano),
		"stable",
		"snapshot",
		"details",
	}
	vs := b.Values()
	assert.DeepEqual(t, expectVs, vs)
}

func (suite *BackupSuite) TestDetailsEntry_HeadersValues() {
	now := time.Now()
	nowStr := now.Format(time.RFC3339Nano)

	table := []struct {
		name     string
		entry    backup.DetailsEntry
		expectHs []string
		expectVs []string
	}{
		{
			name: "no info",
			entry: backup.DetailsEntry{
				RepoRef: "reporef",
			},
			expectHs: []string{"Repo Ref"},
			expectVs: []string{"reporef"},
		},
		{
			name: "exhange info",
			entry: backup.DetailsEntry{
				RepoRef: "reporef",
				ItemInfo: backup.ItemInfo{
					Exchange: &backup.ExchangeInfo{
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
			entry: backup.DetailsEntry{
				RepoRef: "reporef",
				ItemInfo: backup.ItemInfo{
					Sharepoint: &backup.SharepointInfo{},
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
		ents   []backup.DetailsEntry
		expect []string
	}{
		{
			name:   "nil entries",
			ents:   nil,
			expect: []string{},
		},
		{
			name: "single entry",
			ents: []backup.DetailsEntry{
				{RepoRef: "abcde"},
			},
			expect: []string{"abcde"},
		},
		{
			name: "multiple entries",
			ents: []backup.DetailsEntry{
				{RepoRef: "abcde"},
				{RepoRef: "12345"},
			},
			expect: []string{"abcde", "12345"},
		},
	}
	for _, test := range table {
		suite.T().Run(test.name, func(t *testing.T) {
			d := backup.Details{
				DetailsModel: backup.DetailsModel{
					Entries: test.ents,
				},
			}
			assert.DeepEqual(t, test.expect, d.Paths())
		})
	}
}
