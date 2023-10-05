package backup

import (
	"strconv"
	"testing"
	"time"

	"github.com/alcionai/clues"
	"github.com/dustin/go-humanize"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/internal/common/errs"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/internal/stats"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/backup/identity"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

type BackupUnitSuite struct {
	tester.Suite
}

func TestBackupUnitSuite(t *testing.T) {
	suite.Run(t, &BackupUnitSuite{Suite: tester.NewUnitSuite(t)})
}

func stubBackup(t time.Time, ownerID, ownerName string) Backup {
	sel := selectors.NewExchangeBackup([]string{"test"})
	sel.Include(sel.AllData())

	return Backup{
		BaseModel: model.BaseModel{
			ID: model.StableID("id"),
			Tags: map[string]string{
				model.ServiceTag: sel.PathService().String(),
			},
		},
		CreationTime:          t,
		SnapshotID:            "snapshot",
		DetailsID:             "details",
		ProtectedResourceID:   ownerID + "-ro",
		ProtectedResourceName: ownerName + "-ro",
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

func (suite *BackupUnitSuite) TestBackup_Bases() {
	const (
		mergeID  model.StableID = "merge-backup-id"
		assistID model.StableID = "assist-backup-id"
		userID                  = "user-id"
	)

	stub := stubBackup(time.Now(), userID, "user-name")

	defaultEmailReason := identity.NewReason(
		"",
		stub.ProtectedResourceID,
		path.ExchangeService,
		path.EmailCategory)
	defaultContactsReason := identity.NewReason(
		"",
		stub.ProtectedResourceID,
		path.ExchangeService,
		path.ContactsCategory)

	table := []struct {
		name         string
		getBackup    func() *Backup
		expectErr    assert.ErrorAssertionFunc
		expectMerge  map[model.StableID][]identity.Reasoner
		expectAssist map[model.StableID][]identity.Reasoner
	}{
		{
			name: "MergeAndAssist SameReasonEach",
			getBackup: func() *Backup {
				res := stub
				res.MergeBases = map[model.StableID][]string{}
				res.AssistBases = map[model.StableID][]string{}

				res.MergeBases[mergeID] = []string{
					serviceCatString(
						defaultEmailReason.Service(),
						defaultEmailReason.Category()),
				}

				res.AssistBases[assistID] = []string{
					serviceCatString(
						defaultEmailReason.Service(),
						defaultEmailReason.Category()),
				}

				return &res
			},
			expectErr: assert.NoError,
			expectMerge: map[model.StableID][]identity.Reasoner{
				mergeID: {defaultEmailReason},
			},
			expectAssist: map[model.StableID][]identity.Reasoner{
				assistID: {defaultEmailReason},
			},
		},
		{
			name: "MergeAndAssist DifferentReasonEach",
			getBackup: func() *Backup {
				res := stub
				res.MergeBases = map[model.StableID][]string{}
				res.AssistBases = map[model.StableID][]string{}

				res.MergeBases[mergeID] = []string{
					serviceCatString(
						defaultEmailReason.Service(),
						defaultEmailReason.Category()),
				}

				res.AssistBases[assistID] = []string{
					serviceCatString(
						defaultContactsReason.Service(),
						defaultContactsReason.Category()),
				}

				return &res
			},
			expectErr: assert.NoError,
			expectMerge: map[model.StableID][]identity.Reasoner{
				mergeID: {defaultEmailReason},
			},
			expectAssist: map[model.StableID][]identity.Reasoner{
				assistID: {defaultContactsReason},
			},
		},
		{
			name: "MergeAndAssist MultipleReasonsEach",
			getBackup: func() *Backup {
				res := stub
				res.MergeBases = map[model.StableID][]string{}
				res.AssistBases = map[model.StableID][]string{}

				res.MergeBases[mergeID] = []string{
					serviceCatString(
						defaultEmailReason.Service(),
						defaultEmailReason.Category()),
					serviceCatString(
						defaultContactsReason.Service(),
						defaultContactsReason.Category()),
				}

				res.AssistBases[assistID] = []string{
					serviceCatString(
						defaultEmailReason.Service(),
						defaultEmailReason.Category()),
					serviceCatString(
						defaultContactsReason.Service(),
						defaultContactsReason.Category()),
				}

				return &res
			},
			expectErr: assert.NoError,
			expectMerge: map[model.StableID][]identity.Reasoner{
				mergeID: {
					defaultEmailReason,
					defaultContactsReason,
				},
			},
			expectAssist: map[model.StableID][]identity.Reasoner{
				assistID: {
					defaultEmailReason,
					defaultContactsReason,
				},
			},
		},
		{
			name: "OnlyMerge SingleReason",
			getBackup: func() *Backup {
				res := stub
				res.MergeBases = map[model.StableID][]string{}

				res.MergeBases[mergeID] = []string{
					serviceCatString(
						defaultEmailReason.Service(),
						defaultEmailReason.Category()),
				}

				return &res
			},
			expectErr: assert.NoError,
			expectMerge: map[model.StableID][]identity.Reasoner{
				mergeID: {defaultEmailReason},
			},
		},
		{
			name: "OnlyAssist SingleReason",
			getBackup: func() *Backup {
				res := stub
				res.AssistBases = map[model.StableID][]string{}

				res.AssistBases[mergeID] = []string{
					serviceCatString(
						defaultEmailReason.Service(),
						defaultEmailReason.Category()),
				}

				return &res
			},
			expectErr: assert.NoError,
			expectAssist: map[model.StableID][]identity.Reasoner{
				mergeID: {defaultEmailReason},
			},
		},
		{
			name: "BadReasonFormat",
			getBackup: func() *Backup {
				res := stub
				res.AssistBases = map[model.StableID][]string{}

				res.AssistBases[mergeID] = []string{"foo"}

				return &res
			},
			expectErr: assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			bup := test.getBackup()

			got, err := bup.Bases()
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			// Since the result contains a slice of Reasons directly calling Equals
			// will fail because we want ElementsMatch on the internal slices.
			assert.ElementsMatch(
				t,
				maps.Keys(test.expectMerge),
				maps.Keys(got.Merge),
				"merge base keys")
			assert.ElementsMatch(
				t,
				maps.Keys(test.expectAssist),
				maps.Keys(got.Assist),
				"assist base keys")

			for id, e := range test.expectMerge {
				assert.ElementsMatch(t, e, got.Merge[id], "merge bases")
			}

			for id, e := range test.expectAssist {
				assert.ElementsMatch(t, e, got.Assist[id], "assist bases")
			}
		})
	}
}

func (suite *BackupUnitSuite) TestBackup_Tenant() {
	const tenant = "tenant-id"

	stub := stubBackup(time.Now(), "user-id", "user-name")

	table := []struct {
		name       string
		inputKey   string
		inputValue string
		expectErr  assert.ErrorAssertionFunc
		expect     string
	}{
		{
			name:       "ProperlyFormatted",
			inputKey:   tenantIDKey,
			inputValue: tenant,
			expectErr:  assert.NoError,
			expect:     tenant,
		},
		{
			name:       "WrongKey",
			inputKey:   "foo",
			inputValue: tenant,
			expectErr:  assert.Error,
		},
		{
			name:       "EmptyValue",
			inputKey:   tenantIDKey,
			inputValue: "",
			expectErr:  assert.Error,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			b := stub
			b.Tags = map[string]string{test.inputKey: test.inputValue}

			gotTenant, err := b.Tenant()
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				assert.ErrorIs(t, err, errs.NotFound)
				return
			}

			assert.Equal(t, test.expect, gotTenant)
		})
	}
}

func (suite *BackupUnitSuite) TestBackup_Reasons() {
	const (
		tenantID = "tenant-id"
		userID   = "user-id"
	)

	stub := stubBackup(time.Now(), userID, "user-name")

	defaultEmailReason := identity.NewReason(
		tenantID,
		stub.ProtectedResourceID,
		path.ExchangeService,
		path.EmailCategory)
	defaultContactsReason := identity.NewReason(
		tenantID,
		stub.ProtectedResourceID,
		path.ExchangeService,
		path.ContactsCategory)

	table := []struct {
		name      string
		getBackup func() *Backup
		expectErr assert.ErrorAssertionFunc
		expect    []identity.Reasoner
	}{
		{
			name: "SingleReason",
			getBackup: func() *Backup {
				res := stub
				res.Tags = map[string]string{}

				for k, v := range reasonTags(defaultEmailReason) {
					res.Tags[k] = v
				}

				return &res
			},
			expectErr: assert.NoError,
			expect:    []identity.Reasoner{defaultEmailReason},
		},
		{
			name: "MultipleReasons",
			getBackup: func() *Backup {
				res := stub
				res.Tags = map[string]string{}

				for _, reason := range []identity.Reasoner{defaultEmailReason, defaultContactsReason} {
					for k, v := range reasonTags(reason) {
						res.Tags[k] = v
					}
				}

				return &res
			},
			expectErr: assert.NoError,
			expect: []identity.Reasoner{
				defaultEmailReason,
				defaultContactsReason,
			},
		},
		{
			name: "SingleReason OtherTags",
			getBackup: func() *Backup {
				res := stub
				res.Tags = map[string]string{}

				for k, v := range reasonTags(defaultEmailReason) {
					res.Tags[k] = v
				}

				res.Tags["foo"] = "bar"

				return &res
			},
			expectErr: assert.NoError,
			expect:    []identity.Reasoner{defaultEmailReason},
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			t := suite.T()

			bup := test.getBackup()

			got, err := bup.Reasons()
			test.expectErr(t, err, clues.ToCore(err))

			if err != nil {
				return
			}

			assert.ElementsMatch(t, test.expect, got)
		})
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
			"name-ro",
		}
	)

	b.StartAndEndTime.CompletedAt = later

	// single skipped malware
	hs := b.Headers()
	assert.Equal(t, expectHs, hs)

	vs := b.Values()
	assert.Equal(t, expectVs, vs)
}

func (suite *BackupUnitSuite) TestBackup_HeadersValues_onlyResourceOwners() {
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
			"name-ro",
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
		bup    Backup
		expect string
	}{
		{
			name:   "no extras",
			bup:    Backup{Status: "test"},
			expect: "test",
		},
		{
			name: "errors",
			bup: Backup{
				Status:     "test",
				ErrorCount: 42,
			},
			expect: "test (42 errors)",
		},
		{
			name: "malware",
			bup: Backup{
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
			bup: Backup{
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
			bup: Backup{
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
			bup: Backup{
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
			bup: Backup{
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
	result, ok := resultIface.(Printable)
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
