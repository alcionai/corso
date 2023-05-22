package testdata

import (
	"context"
	"testing"
	"time"

	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/cli/utils"
	"github.com/alcionai/corso/src/internal/common/dttm"
	"github.com/alcionai/corso/src/pkg/backup"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/fault"
	ftd "github.com/alcionai/corso/src/pkg/fault/testdata"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/store"
)

type ExchangeOptionsTest struct {
	Name         string
	Opts         func(t *testing.T, wantedVersion int) utils.ExchangeOpts
	BackupGetter *MockBackupGetter
	Expected     func(t *testing.T, wantedVersion int) []details.Entry
}

var (

	// BadExchangeOptionsFormats contains ExchangeOpts with flags that should
	// cause errors about the format of the input flag. Mocks are configured to
	// allow the system to run if it doesn't throw an error on formatting.
	BadExchangeOptionsFormats = []ExchangeOptionsTest{
		{
			Name: "BadEmailReceiveAfter",
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EmailReceivedAfter: "foo",
					Populated: utils.PopulatedFlags{
						utils.EmailReceivedAfterFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "EmptyEmailReceiveAfter",
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EmailReceivedAfter: "",
					Populated: utils.PopulatedFlags{
						utils.EmailReceivedAfterFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "BadEmailReceiveBefore",
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EmailReceivedBefore: "foo",
					Populated: utils.PopulatedFlags{
						utils.EmailReceivedBeforeFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "EmptyEmailReceiveBefore",
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EmailReceivedBefore: "",
					Populated: utils.PopulatedFlags{
						utils.EmailReceivedBeforeFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "BadEventRecurs",
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EventRecurs: "foo",
					Populated: utils.PopulatedFlags{
						utils.EventRecursFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "EmptyEventRecurs",
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EventRecurs: "",
					Populated: utils.PopulatedFlags{
						utils.EventRecursFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "BadEventStartsAfter",
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EventStartsAfter: "foo",
					Populated: utils.PopulatedFlags{
						utils.EventStartsAfterFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "EmptyEventStartsAfter",
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EventStartsAfter: "",
					Populated: utils.PopulatedFlags{
						utils.EventStartsAfterFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "BadEventStartsBefore",
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EventStartsBefore: "foo",
					Populated: utils.PopulatedFlags{
						utils.EventStartsBeforeFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "EmptyEventStartsBefore",
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EventStartsBefore: "",
					Populated: utils.PopulatedFlags{
						utils.EventStartsBeforeFN: struct{}{},
					},
				}
			},
		},
	}

	// ExchangeOptionDetailLookups contains flag inputs and expected results for
	// some choice input patterns. This set is not exhaustive. All inputs and
	// outputs are according to the data laid out in selectors/testdata. Mocks are
	// configured to return the full dataset listed in selectors/testdata.
	ExchangeOptionDetailLookups = []ExchangeOptionsTest{
		{
			Name: "Emails",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					Email: selectors.Any(),
				}
			},
		},
		{
			Name: "EmailsFolderPrefixMatch",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EmailFolder: []string{testdata.ExchangeEmailInboxPath.FolderLocation()},
				}
			},
		},
		{
			Name: "EmailsFolderPrefixMatchTrailingSlash",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EmailFolder: []string{testdata.ExchangeEmailInboxPath.FolderLocation() + "/"},
				}
			},
		},
		{
			Name: "EmailsFolderWithSlashPrefixMatch",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion,
					1, 2)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EmailFolder: []string{testdata.ExchangeEmailBasePath2.FolderLocation()},
				}
			},
		},
		{
			Name: "EmailsFolderWithSlashPrefixMatchTrailingSlash",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion,
					1, 2)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EmailFolder: []string{testdata.ExchangeEmailBasePath2.FolderLocation() + "/"},
				}
			},
		},
		{
			Name: "EmailsBySubject",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion,
					0, 1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EmailSender: "a-person",
				}
			},
		},
		{
			Name: "AllExchange",
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{}
			},
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return append(
					append(
						testdata.GetItemsForVersion(
							t,
							path.ExchangeService,
							path.EmailCategory,
							wantedVersion,
							-1),
						testdata.GetItemsForVersion(
							t,
							path.ExchangeService,
							path.EventsCategory,
							wantedVersion,
							-1)...),
					testdata.GetItemsForVersion(
						t,
						path.ExchangeService,
						path.ContactsCategory,
						wantedVersion,
						-1)...)
			},
		},
		{
			Name: "MailReceivedTime",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion,
					0)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EmailReceivedBefore: dttm.Format(testdata.Time1.Add(time.Second)),
				}
			},
		},
		{
			Name: "MailShortRef",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion,
					0)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				deets := testdata.GetDeetsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion)

				return utils.ExchangeOpts{
					Email: []string{deets[0].ShortRef},
				}
			},
		},
		{
			Name: "BadMailItemRef",
			// no matches are expected, since exchange ItemRefs
			// are not matched when using the CLI's selectors.
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return []details.Entry{}
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				deets := testdata.GetDeetsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion)

				return utils.ExchangeOpts{
					Email: []string{deets[0].ItemRef},
				}
			},
		},
		{
			Name: "MultipleMailShortRef",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion,
					0, 1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				deets := testdata.GetDeetsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion)

				return utils.ExchangeOpts{
					Email: []string{
						deets[0].ShortRef,
						deets[1].ShortRef,
					},
				}
			},
		},
		{
			Name: "AllEventsAndMailWithSubject",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion,
					0)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EmailSubject: "foo",
					Event:        selectors.Any(),
				}
			},
		},
		{
			Name: "EventsAndMailWithSubject",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return []details.Entry{}
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				return utils.ExchangeOpts{
					EmailSubject: "foo",
					EventSubject: "foo",
				}
			},
		},
		{
			Name: "EventsAndMailByShortRef",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return append(
					testdata.GetItemsForVersion(
						t,
						path.ExchangeService,
						path.EmailCategory,
						wantedVersion,
						0),
					testdata.GetItemsForVersion(
						t,
						path.ExchangeService,
						path.EventsCategory,
						wantedVersion,
						0)...)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.ExchangeOpts {
				emailDeets := testdata.GetDeetsForVersion(
					t,
					path.ExchangeService,
					path.EmailCategory,
					wantedVersion)

				eventDeets := testdata.GetDeetsForVersion(
					t,
					path.ExchangeService,
					path.EventsCategory,
					wantedVersion)

				return utils.ExchangeOpts{
					Email: []string{emailDeets[0].ShortRef},
					Event: []string{eventDeets[0].ShortRef},
				}
			},
		},
	}
)

type OneDriveOptionsTest struct {
	Name         string
	Opts         func(t *testing.T, wantedVersion int) utils.OneDriveOpts
	BackupGetter *MockBackupGetter
	Expected     func(t *testing.T, wantedVersion int) []details.Entry
}

var (
	// BadOneDriveOptionsFormats contains OneDriveOpts with flags that should
	// cause errors about the format of the input flag. Mocks are configured to
	// allow the system to run if it doesn't throw an error on formatting.
	BadOneDriveOptionsFormats = []OneDriveOptionsTest{
		{
			Name: "BadFileCreatedAfter",
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					Users:            selectors.Any(),
					FileCreatedAfter: "foo",
					Populated: utils.PopulatedFlags{
						utils.FileCreatedAfterFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "EmptyFileCreatedAfter",
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FileCreatedAfter: "",
					Populated: utils.PopulatedFlags{
						utils.FileCreatedAfterFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "BadFileCreatedBefore",
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FileCreatedBefore: "foo",
					Populated: utils.PopulatedFlags{
						utils.FileCreatedBeforeFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "EmptyFileCreatedBefore",
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FileCreatedBefore: "",
					Populated: utils.PopulatedFlags{
						utils.FileCreatedBeforeFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "BadFileModifiedAfter",
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FileModifiedAfter: "foo",
					Populated: utils.PopulatedFlags{
						utils.FileModifiedAfterFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "EmptyFileModifiedAfter",
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FileModifiedAfter: "",
					Populated: utils.PopulatedFlags{
						utils.FileModifiedAfterFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "BadFileModifiedBefore",
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FileModifiedBefore: "foo",
					Populated: utils.PopulatedFlags{
						utils.FileModifiedBeforeFN: struct{}{},
					},
				}
			},
		},
		{
			Name: "EmptyFileModifiedBefore",
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FileModifiedBefore: "",
					Populated: utils.PopulatedFlags{
						utils.FileModifiedBeforeFN: struct{}{},
					},
				}
			},
		},
	}

	// OneDriveOptionDetailLookups contains flag inputs and expected results for
	// some choice input patterns. This set is not exhaustive. All inputs and
	// outputs are according to the data laid out in selectors/testdata. Mocks are
	// configured to return the full dataset listed in selectors/testdata.
	OneDriveOptionDetailLookups = []OneDriveOptionsTest{
		{
			Name: "AllFiles",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FolderPath: selectors.Any(),
				}
			},
		},
		{
			Name: "FilesWithSingleSlash",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FolderPath: []string{"/"},
				}
			},
		},
		{
			Name: "FolderPrefixMatch",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FolderPath: []string{testdata.OneDriveFolderFolder},
				}
			},
		},
		{
			Name: "FolderPrefixMatchTrailingSlash",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FolderPath: []string{testdata.OneDriveFolderFolder + "/"},
				}
			},
		},
		{
			Name: "FolderPrefixMatchTrailingSlash",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FolderPath: []string{testdata.OneDriveFolderFolder + "/"},
				}
			},
		},
		{
			Name: "FolderRepoRefMatchesNothing",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return []details.Entry{}
			},
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FolderPath: []string{testdata.OneDriveFolderPath.RR.Folder(true)},
				}
			},
		},
		{
			Name: "ShortRef",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion,
					0, 1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				deets := testdata.GetDeetsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion)

				return utils.OneDriveOpts{
					FileName: []string{
						deets[0].ShortRef,
						deets[1].ShortRef,
					},
				}
			},
		},
		{
			Name: "SingleItem",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion,
					0)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				deets := testdata.GetDeetsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion)

				return utils.OneDriveOpts{
					FileName: []string{
						deets[0].OneDrive.ItemName,
					},
				}
			},
		},
		{
			Name: "MultipleItems",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion,
					0, 1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				deets := testdata.GetDeetsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion)

				return utils.OneDriveOpts{
					FileName: []string{
						deets[0].OneDrive.ItemName,
						deets[1].OneDrive.ItemName,
					},
				}
			},
		},
		{
			Name: "ItemRefMatchesNothing",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return []details.Entry{}
			},
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				deets := testdata.GetDeetsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion)

				return utils.OneDriveOpts{
					FileName: []string{
						deets[0].ItemRef,
					},
				}
			},
		},
		{
			Name: "CreatedBefore",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.OneDriveService,
					path.FilesCategory,
					wantedVersion,
					1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
				return utils.OneDriveOpts{
					FileCreatedBefore: dttm.Format(testdata.Time1.Add(time.Second)),
				}
			},
		},
	}
)

type SharePointOptionsTest struct {
	Name         string
	Opts         func(t *testing.T, wantedVersion int) utils.SharePointOpts
	BackupGetter *MockBackupGetter
	Expected     func(t *testing.T, wantedVersion int) []details.Entry
}

var (
	// BadSharePointOptionsFormats contains SharePointOpts with flags that should
	// cause errors about the format of the input flag. Mocks are configured to
	// allow the system to run if it doesn't throw an error on formatting.
	BadSharePointOptionsFormats = []SharePointOptionsTest{
		//{
		//	Name: "BadFileCreatedBefore",
		//	Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
		//		return utils.SharePointOpts{
		//			FileCreatedBefore: "foo",
		//			Populated: utils.PopulatedFlags{
		//				utils.FileCreatedBeforeFN: struct{}{},
		//			},
		//		}
		//	},
		//},
		//{
		//	Name: "EmptyFileCreatedBefore",
		//	Opts: func(t *testing.T, wantedVersion int) utils.OneDriveOpts {
		//		return utils.SharePointOpts{
		//			FileCreatedBefore: "",
		//			Populated: utils.PopulatedFlags{
		//				utils.FileCreatedBeforeFN: struct{}{},
		//			},
		//		}
		//	},
		//},
	}

	// SharePointOptionDetailLookups contains flag inputs and expected results for
	// some choice input patterns. This set is not exhaustive. All inputs and
	// outputs are according to the data laid out in selectors/testdata. Mocks are
	// configured to return the full dataset listed in selectors/testdata.
	SharePointOptionDetailLookups = []SharePointOptionsTest{
		{
			Name: "AllLibraryItems",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.SharePointService,
					path.LibrariesCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.SharePointOpts {
				return utils.SharePointOpts{
					FolderPath: selectors.Any(),
				}
			},
		},
		{
			Name: "LibraryItemsWithSingleSlash",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.SharePointService,
					path.LibrariesCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.SharePointOpts {
				return utils.SharePointOpts{
					FolderPath: []string{"/"},
				}
			},
		},
		{
			Name: "FolderPrefixMatch",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.SharePointService,
					path.LibrariesCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.SharePointOpts {
				return utils.SharePointOpts{
					FolderPath: []string{testdata.SharePointLibraryFolder},
				}
			},
		},
		{
			Name: "FolderPrefixMatchTrailingSlash",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.SharePointService,
					path.LibrariesCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.SharePointOpts {
				return utils.SharePointOpts{
					FolderPath: []string{testdata.SharePointLibraryFolder + "/"},
				}
			},
		},
		{
			Name: "FolderPrefixMatchTrailingSlash",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.SharePointService,
					path.LibrariesCategory,
					wantedVersion,
					-1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.SharePointOpts {
				return utils.SharePointOpts{
					FolderPath: []string{testdata.SharePointLibraryFolder + "/"},
				}
			},
		},
		{
			Name: "FolderRepoRefMatchesNothing",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return []details.Entry{}
			},
			Opts: func(t *testing.T, wantedVersion int) utils.SharePointOpts {
				return utils.SharePointOpts{
					FolderPath: []string{testdata.SharePointLibraryPath.RR.Folder(true)},
				}
			},
		},
		{
			Name: "ShortRef",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.SharePointService,
					path.LibrariesCategory,
					wantedVersion,
					0, 1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.SharePointOpts {
				deets := testdata.GetDeetsForVersion(
					t,
					path.SharePointService,
					path.LibrariesCategory,
					wantedVersion)

				return utils.SharePointOpts{
					FileName: []string{
						deets[0].ShortRef,
						deets[1].ShortRef,
					},
				}
			},
		},
		{
			Name: "SingleItem",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.SharePointService,
					path.LibrariesCategory,
					wantedVersion,
					0)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.SharePointOpts {
				deets := testdata.GetDeetsForVersion(
					t,
					path.SharePointService,
					path.LibrariesCategory,
					wantedVersion)

				return utils.SharePointOpts{
					FileName: []string{
						deets[0].SharePoint.ItemName,
					},
				}
			},
		},
		{
			Name: "MultipleItems",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return testdata.GetItemsForVersion(
					t,
					path.SharePointService,
					path.LibrariesCategory,
					wantedVersion,
					0, 1)
			},
			Opts: func(t *testing.T, wantedVersion int) utils.SharePointOpts {
				deets := testdata.GetDeetsForVersion(
					t,
					path.SharePointService,
					path.LibrariesCategory,
					wantedVersion)

				return utils.SharePointOpts{
					FileName: []string{
						deets[0].SharePoint.ItemName,
						deets[1].SharePoint.ItemName,
					},
				}
			},
		},
		{
			Name: "ItemRefMatchesNothing",
			Expected: func(t *testing.T, wantedVersion int) []details.Entry {
				return []details.Entry{}
			},
			Opts: func(t *testing.T, wantedVersion int) utils.SharePointOpts {
				deets := testdata.GetDeetsForVersion(
					t,
					path.SharePointService,
					path.LibrariesCategory,
					wantedVersion)

				return utils.SharePointOpts{
					FileName: []string{
						deets[0].ItemRef,
					},
				}
			},
		},
		//{
		//	Name: "CreatedBefore",
		//	Expected: func(t *testing.T, wantedVersion int) []details.DetailsEntry {
		//		return testdata.GetItemsForVersion(
		//			t,
		//			path.SharePointService,
		//			path.LibrariesCategory,
		//			wantedVersion,
		//			1)
		//	},
		//	Opts: func(t *testing.T, wantedVersion int) utils.SharePointOpts {
		//		return utils.SharePointOpts{
		//			FileCreatedBefore: dttm.Format(testdata.Time1.Add(time.Second)),
		//		}
		//	},
		//},
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

type VersionedBackupGetter struct {
	*MockBackupGetter
	Details *details.Details
}

func (bg VersionedBackupGetter) GetBackupDetails(
	ctx context.Context,
	backupID string,
) (*details.Details, *backup.Backup, *fault.Bus) {
	return bg.Details, nil, fault.New(true)
}
