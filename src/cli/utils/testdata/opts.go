package testdata

import (
	"context"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	ftd "github.com/alcionai/corso/src/pkg/fault/testdata"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

type ExchangeOptionsTest struct {
	Name         string
	Opts         utils.ExchangeOpts
	BackupGetter *MockBackupGetter
	Expected     []details.Entry
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
				EmailFolder: []string{testdata.ExchangeEmailInboxPath.FolderLocation()},
			},
		},
		{
			Name:     "EmailsFolderPrefixMatchTrailingSlash",
			Expected: testdata.ExchangeEmailItems,
			Opts: utils.ExchangeOpts{
				EmailFolder: []string{testdata.ExchangeEmailInboxPath.FolderLocation() + "/"},
			},
		},
		{
			Name: "EmailsFolderWithSlashPrefixMatch",
			Expected: []details.Entry{
				testdata.ExchangeEmailItems[1],
				testdata.ExchangeEmailItems[2],
			},
			Opts: utils.ExchangeOpts{
				EmailFolder: []string{testdata.ExchangeEmailBasePath2.FolderLocation()},
			},
		},
		{
			Name: "EmailsFolderWithSlashPrefixMatchTrailingSlash",
			Expected: []details.Entry{
				testdata.ExchangeEmailItems[1],
				testdata.ExchangeEmailItems[2],
			},
			Opts: utils.ExchangeOpts{
				EmailFolder: []string{testdata.ExchangeEmailBasePath2.FolderLocation() + "/"},
			},
		},
		{
			Name: "EmailsBySubject",
			Expected: []details.Entry{
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
						[]details.Entry{},
						testdata.ExchangeEmailItems...,
					),
					testdata.ExchangeContactsItems...,
				),
				testdata.ExchangeEventsItems...,
			),
		},
		{
			Name:     "MailReceivedTime",
			Expected: []details.Entry{testdata.ExchangeEmailItems[0]},
			Opts: utils.ExchangeOpts{
				EmailReceivedBefore: dttm.Format(testdata.Time1.Add(time.Second)),
			},
		},
		{
			Name:     "MailShortRef",
			Expected: []details.Entry{testdata.ExchangeEmailItems[0]},
			Opts: utils.ExchangeOpts{
				Email: []string{testdata.ExchangeEmailItemPath1.RR.ShortRef()},
			},
		},
		{
			Name: "BadMailItemRef",
			// no matches are expected, since exchange ItemRefs
			// are not matched when using the CLI's selectors.
			Expected: []details.Entry{},
			Opts: utils.ExchangeOpts{
				Email: []string{testdata.ExchangeEmailItems[0].ItemRef},
			},
		},
		{
			Name: "MultipleMailShortRef",
			Expected: []details.Entry{
				testdata.ExchangeEmailItems[0],
				testdata.ExchangeEmailItems[1],
			},
			Opts: utils.ExchangeOpts{
				Email: []string{
					testdata.ExchangeEmailItemPath1.RR.ShortRef(),
					testdata.ExchangeEmailItemPath2.RR.ShortRef(),
				},
			},
		},
		{
			Name:     "AllEventsAndMailWithSubject",
			Expected: []details.Entry{testdata.ExchangeEmailItems[0]},
			Opts: utils.ExchangeOpts{
				EmailSubject: "foo",
				Event:        selectors.Any(),
			},
		},
		{
			Name:     "EventsAndMailWithSubject",
			Expected: []details.Entry{},
			Opts: utils.ExchangeOpts{
				EmailSubject: "foo",
				EventSubject: "foo",
			},
		},
		{
			Name: "EventsAndMailByShortRef",
			Expected: []details.Entry{
				testdata.ExchangeEmailItems[0],
				testdata.ExchangeEventsItems[0],
			},
			Opts: utils.ExchangeOpts{
				Email: []string{testdata.ExchangeEmailItemPath1.RR.ShortRef()},
				Event: []string{testdata.ExchangeEventsItemPath1.RR.ShortRef()},
			},
		},
	}
)

type OneDriveOptionsTest struct {
	Name         string
	Opts         utils.OneDriveOpts
	BackupGetter *MockBackupGetter
	Expected     []details.Entry
}

var (
	// BadOneDriveOptionsFormats contains OneDriveOpts with flags that should
	// cause errors about the format of the input flag. Mocks are configured to
	// allow the system to run if it doesn't throw an error on formatting.
	BadOneDriveOptionsFormats = []OneDriveOptionsTest{
		{
			Name: "BadFileCreatedAfter",
			Opts: utils.OneDriveOpts{
				Users:            selectors.Any(),
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
				FolderPath: selectors.Any(),
			},
		},
		{
			Name:     "FilesWithSingleSlash",
			Expected: testdata.OneDriveItems,
			Opts: utils.OneDriveOpts{
				FolderPath: []string{"/"},
			},
		},
		{
			Name:     "FolderPrefixMatch",
			Expected: testdata.OneDriveItems,
			Opts: utils.OneDriveOpts{
				FolderPath: []string{testdata.OneDriveFolderFolder},
			},
		},
		{
			Name:     "FolderPrefixMatchTrailingSlash",
			Expected: testdata.OneDriveItems,
			Opts: utils.OneDriveOpts{
				FolderPath: []string{testdata.OneDriveFolderFolder + "/"},
			},
		},
		{
			Name:     "FolderPrefixMatchTrailingSlash",
			Expected: testdata.OneDriveItems,
			Opts: utils.OneDriveOpts{
				FolderPath: []string{testdata.OneDriveFolderFolder + "/"},
			},
		},
		{
			Name:     "FolderRepoRefMatchesNothing",
			Expected: []details.Entry{},
			Opts: utils.OneDriveOpts{
				FolderPath: []string{testdata.OneDriveFolderPath.RR.Folder(true)},
			},
		},
		{
			Name: "ShortRef",
			Expected: []details.Entry{
				testdata.OneDriveItems[0],
				testdata.OneDriveItems[1],
			},
			Opts: utils.OneDriveOpts{
				FileName: []string{
					testdata.OneDriveItems[0].ShortRef,
					testdata.OneDriveItems[1].ShortRef,
				},
			},
		},
		{
			Name:     "SingleItem",
			Expected: []details.Entry{testdata.OneDriveItems[0]},
			Opts: utils.OneDriveOpts{
				FileName: []string{
					testdata.OneDriveItems[0].OneDrive.ItemName,
				},
			},
		},
		{
			Name: "MultipleItems",
			Expected: []details.Entry{
				testdata.OneDriveItems[0],
				testdata.OneDriveItems[1],
			},
			Opts: utils.OneDriveOpts{
				FileName: []string{
					testdata.OneDriveItems[0].OneDrive.ItemName,
					testdata.OneDriveItems[1].OneDrive.ItemName,
				},
			},
		},
		{
			Name:     "ItemRefMatchesNothing",
			Expected: []details.Entry{},
			Opts: utils.OneDriveOpts{
				FileName: []string{
					testdata.OneDriveItems[0].ItemRef,
				},
			},
		},
		{
			Name:     "CreatedBefore",
			Expected: []details.Entry{testdata.OneDriveItems[1]},
			Opts: utils.OneDriveOpts{
				FileCreatedBefore: dttm.Format(testdata.Time1.Add(time.Second)),
			},
		},
	}
)

type SharePointOptionsTest struct {
	Name         string
	Opts         utils.SharePointOpts
	BackupGetter *MockBackupGetter
	Expected     []details.Entry
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
				FolderPath: selectors.Any(),
			},
		},
		{
			Name:     "LibraryItemsWithSingleSlash",
			Expected: testdata.SharePointLibraryItems,
			Opts: utils.SharePointOpts{
				FolderPath: []string{"/"},
			},
		},
		{
			Name:     "FolderPrefixMatch",
			Expected: testdata.SharePointLibraryItems,
			Opts: utils.SharePointOpts{
				FolderPath: []string{testdata.SharePointLibraryFolder},
			},
		},
		{
			Name:     "FolderPrefixMatchTrailingSlash",
			Expected: testdata.SharePointLibraryItems,
			Opts: utils.SharePointOpts{
				FolderPath: []string{testdata.SharePointLibraryFolder + "/"},
			},
		},
		{
			Name:     "FolderPrefixMatchTrailingSlash",
			Expected: testdata.SharePointLibraryItems,
			Opts: utils.SharePointOpts{
				FolderPath: []string{testdata.SharePointLibraryFolder + "/"},
			},
		},
		{
			Name:     "FolderRepoRefMatchesNothing",
			Expected: []details.Entry{},
			Opts: utils.SharePointOpts{
				FolderPath: []string{testdata.SharePointLibraryPath.RR.Folder(true)},
			},
		},
		{
			Name: "ShortRef",
			Expected: []details.Entry{
				testdata.SharePointLibraryItems[0],
				testdata.SharePointLibraryItems[1],
			},
			Opts: utils.SharePointOpts{
				FileName: []string{
					testdata.SharePointLibraryItems[0].ShortRef,
					testdata.SharePointLibraryItems[1].ShortRef,
				},
			},
		},
		{
			Name:     "SingleItem",
			Expected: []details.Entry{testdata.SharePointLibraryItems[0]},
			Opts: utils.SharePointOpts{
				FileName: []string{
					testdata.SharePointLibraryItems[0].SharePoint.ItemName,
				},
			},
		},
		{
			Name: "MultipleItems",
			Expected: []details.Entry{
				testdata.SharePointLibraryItems[0],
				testdata.SharePointLibraryItems[1],
			},
			Opts: utils.SharePointOpts{
				FileName: []string{
					testdata.SharePointLibraryItems[0].SharePoint.ItemName,
					testdata.SharePointLibraryItems[1].SharePoint.ItemName,
				},
			},
		},
		{
			Name:     "ItemRefMatchesNothing",
			Expected: []details.Entry{},
			Opts: utils.SharePointOpts{
				FileName: []string{
					testdata.SharePointLibraryItems[0].ItemRef,
				},
			},
		},
		// {
		// 	Name:     "CreatedBefore",
		// 	Expected: []details.DetailsEntry{testdata.SharePointLibraryItems[1]},
		// 	Opts: utils.SharePointOpts{
		// 		FileCreatedBefore: dttm.Format(testdata.Time1.Add(time.Second)),
		// 	},
		// },
	}
)

// MockBackupGetter implements the repo.BackupGetter interface and returns
// (selectors/testdata.GetDetailsSet(), nil, nil) when BackupDetails is called
// on the nil instance. If an instance is given or Backups is called returns an
// error.
type MockBackupGetter struct {
	failure, recovered, skipped bool
}

func (MockBackupGetter) Backup(
	context.Context,
	string,
) (*backup.Backup, error) {
	return nil, clues.New("unexpected call to mock")
}

func (MockBackupGetter) Backups(
	context.Context,
	[]string,
) ([]*backup.Backup, *fault.Bus) {
	return nil, fault.New(false).Fail(clues.New("unexpected call to mock"))
}

func (MockBackupGetter) BackupsByTag(
	context.Context,
	...store.FilterOption,
) ([]*backup.Backup, error) {
	return nil, clues.New("unexpected call to mock")
}

func (bg *MockBackupGetter) GetBackupDetails(
	ctx context.Context,
	backupID string,
) (*details.Details, *backup.Backup, *fault.Bus) {
	if bg == nil {
		return testdata.GetDetailsSet(), nil, fault.New(true)
	}

	return nil, nil, fault.New(false).Fail(clues.New("unexpected call to mock"))
}

func (bg *MockBackupGetter) GetBackupErrors(
	ctx context.Context,
	backupID string,
) (*fault.Errors, *backup.Backup, *fault.Bus) {
	if bg == nil {
		fe := ftd.MakeErrors(bg.failure, bg.recovered, bg.skipped)
		return &fe, nil, fault.New(true)
	}

	return nil, nil, fault.New(false).Fail(clues.New("unexpected call to mock"))
}
