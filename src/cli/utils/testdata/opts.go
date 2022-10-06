package testdata

import (
	"context"
	"errors"
	"time"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common"
	"github.com/alcionai/corso/src/internal/model"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/selectors/testdata"
)

type ExchangeOptionsTest struct {
	Name         string
	Opts         utils.ExchangeOpts
	BackupGetter *MockBackupGetter
	Expected     []details.DetailsEntry
}

var (

	// BadExchangeOptionsFormats contains ExchangeOpts with flags that should
	// cause errors about the format of the input flag. Mocks are configured to
	// allow the system to run if it doesn't throw an error on formatting.
	BadExchangeOptionsFormats = []ExchangeOptionsTest{
		{
			Name: "BadEmailReceiveAfter",
			Opts: utils.ExchangeOpts{
				EmailReceivedAfter: "foo",
				Populated: utils.PopulatedFlags{
					utils.EmailReceivedAfterFN: struct{}{},
				},
			},
		},
		{
			Name: "BadEmailReceiveBefore",
			Opts: utils.ExchangeOpts{
				EmailReceivedBefore: "foo",
				Populated: utils.PopulatedFlags{
					utils.EmailReceivedBeforeFN: struct{}{},
				},
			},
		},
		{
			Name: "BadEventRecurs",
			Opts: utils.ExchangeOpts{
				EventRecurs: "foo",
				Populated: utils.PopulatedFlags{
					utils.EventRecursFN: struct{}{},
				},
			},
		},
		{
			Name: "BadEventStartsAfter",
			Opts: utils.ExchangeOpts{
				EventStartsAfter: "foo",
				Populated: utils.PopulatedFlags{
					utils.EventStartsAfterFN: struct{}{},
				},
			},
		},
		{
			Name: "BadEventStartsBefore",
			Opts: utils.ExchangeOpts{
				EventStartsBefore: "foo",
				Populated: utils.PopulatedFlags{
					utils.EventStartsBeforeFN: struct{}{},
				},
			},
		},
	}

	// ExchangeOptionDetailLookups contains flag inputs and expected results for
	// some choice input patterns. This set is not exhaustive. All inputs and
	// outputs are according to the data laid out in selectors/testdata. Mocks are
	// configured to return the full dataset listed in selectors/testdata.
	ExchangeOptionDetailLookups = []ExchangeOptionsTest{
		{
			Name:     "Emails",
			Expected: testdata.ExchangeEmailItems,
			Opts: utils.ExchangeOpts{
				Email: selectors.Any(),
			},
		},
		{
			Name:     "EmailsFolderPrefixMatch",
			Expected: testdata.ExchangeEmailItems,
			Opts: utils.ExchangeOpts{
				EmailFolder: []string{testdata.ExchangeEmailInboxPath.Folder()},
			},
		},
		{
			Name:     "EmailsBySubject",
			Expected: testdata.ExchangeEmailItems,
			Opts: utils.ExchangeOpts{
				EmailSender: "a-person",
			},
		},
		{
			Name: "AllExchange",
			Expected: append(
				append(
					append(
						[]details.DetailsEntry{},
						testdata.ExchangeEmailItems...,
					),
					testdata.ExchangeContactsItems...,
				),
				testdata.ExchangeEventsItems...,
			),
		},
		{
			Name:     "MailReceivedTime",
			Expected: []details.DetailsEntry{testdata.ExchangeEmailItems[0]},
			Opts: utils.ExchangeOpts{
				EmailReceivedBefore: common.FormatTime(testdata.Time1.Add(time.Second)),
			},
		},
		{
			Name:     "MailID",
			Expected: []details.DetailsEntry{testdata.ExchangeEmailItems[0]},
			Opts: utils.ExchangeOpts{
				Email: []string{testdata.ExchangeEmailItemPath1.Item()},
			},
		},
		{
			Name:     "MailShortRef",
			Expected: []details.DetailsEntry{testdata.ExchangeEmailItems[0]},
			Opts: utils.ExchangeOpts{
				Email: []string{testdata.ExchangeEmailItemPath1.ShortRef()},
			},
		},
		{
			Name:     "MultipleMailShortRef",
			Expected: testdata.ExchangeEmailItems,
			Opts: utils.ExchangeOpts{
				Email: []string{
					testdata.ExchangeEmailItemPath1.ShortRef(),
					testdata.ExchangeEmailItemPath2.ShortRef(),
				},
			},
		},
		{
			Name:     "AllEventsAndMailWithSubject",
			Expected: []details.DetailsEntry{testdata.ExchangeEmailItems[0]},
			Opts: utils.ExchangeOpts{
				EmailSubject: "foo",
				Event:        selectors.Any(),
			},
		},
		{
			Name:     "EventsAndMailWithSubject",
			Expected: []details.DetailsEntry{},
			Opts: utils.ExchangeOpts{
				EmailSubject: "foo",
				EventSubject: "foo",
			},
		},
		{
			Name: "EventsAndMailByShortRef",
			Expected: []details.DetailsEntry{
				testdata.ExchangeEmailItems[0],
				testdata.ExchangeEventsItems[0],
			},
			Opts: utils.ExchangeOpts{
				Email: []string{testdata.ExchangeEmailItemPath1.ShortRef()},
				Event: []string{testdata.ExchangeEventsItemPath1.ShortRef()},
			},
		},
	}
)

// MockBackupGetter implements the repo.BackupGetter interface and returns
// (selectors/testdata.GetDetailsSet(), nil, nil) when BackupDetails is called
// on the nil instance. If an instance is given or Backups is called returns an
// error.
type MockBackupGetter struct{}

func (MockBackupGetter) Backup(
	context.Context,
	model.StableID,
) (*backup.Backup, error) {
	return nil, errors.New("unexpected call to mock")
}

func (MockBackupGetter) Backups(context.Context) ([]backup.Backup, error) {
	return nil, errors.New("unexpected call to mock")
}

func (bg *MockBackupGetter) BackupDetails(
	ctx context.Context,
	backupID string,
) (*details.Details, *backup.Backup, error) {
	if bg == nil {
		return testdata.GetDetailsSet(), nil, nil
	}

	return nil, nil, errors.New("unexpected call to mock")
}
