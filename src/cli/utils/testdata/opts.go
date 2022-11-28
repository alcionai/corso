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
	"github.com/alcionai/corso/src/pkg/store"
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
			Name: "EmptyEmailReceiveAfter",
			Opts: utils.ExchangeOpts{
				EmailReceivedAfter: "",
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
			Name: "EmptyEmailReceiveBefore",
			Opts: utils.ExchangeOpts{
				EmailReceivedBefore: "",
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
			Name: "EmptyEventRecurs",
			Opts: utils.ExchangeOpts{
				EventRecurs: "",
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
			Name: "EmptyEventStartsAfter",
			Opts: utils.ExchangeOpts{
				EventStartsAfter: "",
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
		{
			Name: "EmptyEventStartsBefore",
			Opts: utils.ExchangeOpts{
				EventStartsBefore: "",
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
			Name:     "EmailsFolderPrefixMatchTrailingSlash",
			Expected: testdata.ExchangeEmailItems,
			Opts: utils.ExchangeOpts{
				EmailFolder: []string{testdata.ExchangeEmailInboxPath.Folder() + "/"},
			},
		},
		{
			Name: "EmailsFolderWithSlashPrefixMatch",
			Expected: []details.DetailsEntry{
				testdata.ExchangeEmailItems[1],
				testdata.ExchangeEmailItems[2],
			},
			Opts: utils.ExchangeOpts{
				EmailFolder: []string{testdata.ExchangeEmailBasePath2.Folder()},
			},
		},
		{
			Name: "EmailsFolderWithSlashPrefixMatchTrailingSlash",
			Expected: []details.DetailsEntry{
				testdata.ExchangeEmailItems[1],
				testdata.ExchangeEmailItems[2],
			},
			Opts: utils.ExchangeOpts{
				EmailFolder: []string{testdata.ExchangeEmailBasePath2.Folder() + "/"},
			},
		},
		{
			Name: "EmailsBySubject",
			Expected: []details.DetailsEntry{
				testdata.ExchangeEmailItems[0],
				testdata.ExchangeEmailItems[1],
			},
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
			Name: "MultipleMailShortRef",
			Expected: []details.DetailsEntry{
				testdata.ExchangeEmailItems[0],
				testdata.ExchangeEmailItems[1],
			},
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

type OneDriveOptionsTest struct {
	Name         string
	Opts         utils.OneDriveOpts
	BackupGetter *MockBackupGetter
	Expected     []details.DetailsEntry
}

var (
	// BadOneDriveOptionsFormats contains OneDriveOpts with flags that should
	// cause errors about the format of the input flag. Mocks are configured to
	// allow the system to run if it doesn't throw an error on formatting.
	BadOneDriveOptionsFormats = []OneDriveOptionsTest{
		{
			Name: "BadFileCreatedAfter",
			Opts: utils.OneDriveOpts{
				FileCreatedAfter: "foo",
				Populated: utils.PopulatedFlags{
					utils.FileCreatedAfterFN: struct{}{},
				},
			},
		},
		{
			Name: "EmptyFileCreatedAfter",
			Opts: utils.OneDriveOpts{
				FileCreatedAfter: "",
				Populated: utils.PopulatedFlags{
					utils.FileCreatedAfterFN: struct{}{},
				},
			},
		},
		{
			Name: "BadFileCreatedBefore",
			Opts: utils.OneDriveOpts{
				FileCreatedBefore: "foo",
				Populated: utils.PopulatedFlags{
					utils.FileCreatedBeforeFN: struct{}{},
				},
			},
		},
		{
			Name: "EmptyFileCreatedBefore",
			Opts: utils.OneDriveOpts{
				FileCreatedBefore: "",
				Populated: utils.PopulatedFlags{
					utils.FileCreatedBeforeFN: struct{}{},
				},
			},
		},
		{
			Name: "BadFileModifiedAfter",
			Opts: utils.OneDriveOpts{
				FileModifiedAfter: "foo",
				Populated: utils.PopulatedFlags{
					utils.FileModifiedAfterFN: struct{}{},
				},
			},
		},
		{
			Name: "EmptyFileModifiedAfter",
			Opts: utils.OneDriveOpts{
				FileModifiedAfter: "",
				Populated: utils.PopulatedFlags{
					utils.FileModifiedAfterFN: struct{}{},
				},
			},
		},
		{
			Name: "BadFileModifiedBefore",
			Opts: utils.OneDriveOpts{
				FileModifiedBefore: "foo",
				Populated: utils.PopulatedFlags{
					utils.FileModifiedBeforeFN: struct{}{},
				},
			},
		},
		{
			Name: "EmptyFileModifiedBefore",
			Opts: utils.OneDriveOpts{
				FileModifiedBefore: "",
				Populated: utils.PopulatedFlags{
					utils.FileModifiedBeforeFN: struct{}{},
				},
			},
		},
	}

	// OneDriveOptionDetailLookups contains flag inputs and expected results for
	// some choice input patterns. This set is not exhaustive. All inputs and
	// outputs are according to the data laid out in selectors/testdata. Mocks are
	// configured to return the full dataset listed in selectors/testdata.
	OneDriveOptionDetailLookups = []OneDriveOptionsTest{
		{
			Name:     "AllFiles",
			Expected: testdata.OneDriveItems,
			Opts: utils.OneDriveOpts{
				Paths: selectors.Any(),
			},
		},
		{
			Name:     "FolderPrefixMatch",
			Expected: testdata.OneDriveItems,
			Opts: utils.OneDriveOpts{
				Paths: []string{testdata.OneDriveFolderFolder},
			},
		},
		{
			Name:     "FolderPrefixMatchTrailingSlash",
			Expected: testdata.OneDriveItems,
			Opts: utils.OneDriveOpts{
				Paths: []string{testdata.OneDriveFolderFolder + "/"},
			},
		},
		{
			Name:     "FolderPrefixMatchTrailingSlash",
			Expected: testdata.OneDriveItems,
			Opts: utils.OneDriveOpts{
				Paths: []string{testdata.OneDriveFolderFolder + "/"},
			},
		},
		{
			Name: "ShortRef",
			Expected: []details.DetailsEntry{
				testdata.OneDriveItems[0],
				testdata.OneDriveItems[1],
			},
			Opts: utils.OneDriveOpts{
				Names: []string{
					testdata.OneDriveItems[0].ShortRef,
					testdata.OneDriveItems[1].ShortRef,
				},
			},
		},
		{
			Name:     "CreatedBefore",
			Expected: []details.DetailsEntry{testdata.OneDriveItems[1]},
			Opts: utils.OneDriveOpts{
				FileCreatedBefore: common.FormatTime(testdata.Time1.Add(time.Second)),
			},
		},
	}
)

type SharePointOptionsTest struct {
	Name         string
	Opts         utils.SharePointOpts
	BackupGetter *MockBackupGetter
	Expected     []details.DetailsEntry
}

var (
	// BadSharePointOptionsFormats contains SharePointOpts with flags that should
	// cause errors about the format of the input flag. Mocks are configured to
	// allow the system to run if it doesn't throw an error on formatting.
	BadSharePointOptionsFormats = []SharePointOptionsTest{
		// {
		// 	Name: "BadFileCreatedBefore",
		// 	Opts: utils.OneDriveOpts{
		// 		FileCreatedBefore: "foo",
		// 		Populated: utils.PopulatedFlags{
		// 			utils.FileCreatedBeforeFN: struct{}{},
		// 		},
		// 	},
		// },
		// {
		// 	Name: "EmptyFileCreatedBefore",
		// 	Opts: utils.OneDriveOpts{
		// 		FileCreatedBefore: "",
		// 		Populated: utils.PopulatedFlags{
		// 			utils.FileCreatedBeforeFN: struct{}{},
		// 		},
		// 	},
		// },
	}

	// SharePointOptionDetailLookups contains flag inputs and expected results for
	// some choice input patterns. This set is not exhaustive. All inputs and
	// outputs are according to the data laid out in selectors/testdata. Mocks are
	// configured to return the full dataset listed in selectors/testdata.
	SharePointOptionDetailLookups = []SharePointOptionsTest{
		{
			Name:     "AllLibraryItems",
			Expected: testdata.SharePointLibraryItems,
			Opts: utils.SharePointOpts{
				Libraries: selectors.Any(),
			},
		},
		{
			Name:     "FolderPrefixMatch",
			Expected: testdata.SharePointLibraryItems,
			Opts: utils.SharePointOpts{
				Libraries: []string{testdata.SharePointLibraryFolder},
			},
		},
		{
			Name:     "FolderPrefixMatchTrailingSlash",
			Expected: testdata.SharePointLibraryItems,
			Opts: utils.SharePointOpts{
				Libraries: []string{testdata.SharePointLibraryFolder + "/"},
			},
		},
		{
			Name:     "FolderPrefixMatchTrailingSlash",
			Expected: testdata.SharePointLibraryItems,
			Opts: utils.SharePointOpts{
				Libraries: []string{testdata.SharePointLibraryFolder + "/"},
			},
		},
		{
			Name: "ShortRef",
			Expected: []details.DetailsEntry{
				testdata.SharePointLibraryItems[0],
				testdata.SharePointLibraryItems[1],
			},
			Opts: utils.SharePointOpts{
				LibraryItems: []string{
					testdata.SharePointLibraryItems[0].ShortRef,
					testdata.SharePointLibraryItems[1].ShortRef,
				},
			},
		},
		// {
		// 	Name:     "CreatedBefore",
		// 	Expected: []details.DetailsEntry{testdata.SharePointLibraryItems[1]},
		// 	Opts: utils.SharePointOpts{
		// 		FileCreatedBefore: common.FormatTime(testdata.Time1.Add(time.Second)),
		// 	},
		// },
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

func (MockBackupGetter) Backups(context.Context, ...store.FilterOption) ([]*backup.Backup, error) {
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
