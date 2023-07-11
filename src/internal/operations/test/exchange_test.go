package test_test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/dttm"
	inMock "github.com/alcionai/corso/src/internal/common/idname/mock"
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/events"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/m365/exchange"
	exchMock "github.com/alcionai/corso/src/internal/m365/exchange/mock"
	exchTD "github.com/alcionai/corso/src/internal/m365/exchange/testdata"
	"github.com/alcionai/corso/src/internal/m365/resource"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	deeTD "github.com/alcionai/corso/src/pkg/backup/details/testdata"
	"github.com/alcionai/corso/src/pkg/control"
	ctrlTD "github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
	storeTD "github.com/alcionai/corso/src/pkg/storage/testdata"
)

type ExchangeBackupIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestExchangeBackupIntgSuite(t *testing.T) {
	suite.Run(t, &ExchangeBackupIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *ExchangeBackupIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

// TestBackup_Run ensures that Integration Testing works
// for the following scopes: Contacts, Events, and Mail
func (suite *ExchangeBackupIntgSuite) TestBackup_Run_exchange() {
	tests := []struct {
		name          string
		selector      func() *selectors.ExchangeBackup
		category      path.CategoryType
		metadataFiles []string
	}{
		{
			name: "Mail",
			selector: func() *selectors.ExchangeBackup {
				sel := selectors.NewExchangeBackup([]string{suite.its.userID})
				sel.Include(sel.MailFolders([]string{api.MailInbox}, selectors.PrefixMatch()))
				sel.DiscreteOwner = suite.its.userID

				return sel
			},
			category:      path.EmailCategory,
			metadataFiles: exchange.MetadataFileNames(path.EmailCategory),
		},
		{
			name: "Contacts",
			selector: func() *selectors.ExchangeBackup {
				sel := selectors.NewExchangeBackup([]string{suite.its.userID})
				sel.Include(sel.ContactFolders([]string{api.DefaultContacts}, selectors.PrefixMatch()))
				return sel
			},
			category:      path.ContactsCategory,
			metadataFiles: exchange.MetadataFileNames(path.ContactsCategory),
		},
		{
			name: "Calendar Events",
			selector: func() *selectors.ExchangeBackup {
				sel := selectors.NewExchangeBackup([]string{suite.its.userID})
				sel.Include(sel.EventCalendars([]string{api.DefaultCalendar}, selectors.PrefixMatch()))
				return sel
			},
			category:      path.EventsCategory,
			metadataFiles: exchange.MetadataFileNames(path.EventsCategory),
		},
	}
	for _, test := range tests {
		suite.Run(test.name, func() {
			t := suite.T()

			ctx, flush := tester.NewContext(t)
			defer flush()

			var (
				mb      = evmock.NewBus()
				sel     = test.selector().Selector
				opts    = control.Defaults()
				whatSet = deeTD.CategoryFromRepoRef
			)

			bo, bod := prepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup)
			defer bod.close(t, ctx)

			sel = bod.sel

			userID := sel.ID()

			m365, err := bod.acct.M365Config()
			require.NoError(t, err, clues.ToCore(err))

			// run the tests
			runAndCheckBackup(t, ctx, &bo, mb, false)
			checkBackupIsInManifests(
				t,
				ctx,
				bod.kw,
				bod.sw,
				&bo,
				sel,
				userID,
				test.category)
			checkMetadataFilesExist(
				t,
				ctx,
				bo.Results.BackupID,
				bod.kw,
				bod.kms,
				m365.AzureTenantID,
				userID,
				path.ExchangeService,
				map[path.CategoryType][]string{test.category: test.metadataFiles})

			_, expectDeets := deeTD.GetDeetsInBackup(
				t,
				ctx,
				bo.Results.BackupID,
				bod.acct.ID(),
				userID,
				path.ExchangeService,
				whatSet,
				bod.kms,
				bod.sss)
			deeTD.CheckBackupDetails(
				t,
				ctx,
				bo.Results.BackupID,
				whatSet,
				bod.kms,
				bod.sss,
				expectDeets,
				false)

			// Basic, happy path incremental test.  No changes are dictated or expected.
			// This only tests that an incremental backup is runnable at all, and that it
			// produces fewer results than the last backup.
			var (
				incMB = evmock.NewBus()
				incBO = newTestBackupOp(
					t,
					ctx,
					bod,
					incMB,
					opts)
			)

			runAndCheckBackup(t, ctx, &incBO, incMB, true)
			checkBackupIsInManifests(
				t,
				ctx,
				bod.kw,
				bod.sw,
				&incBO,
				sel,
				userID,
				test.category)
			checkMetadataFilesExist(
				t,
				ctx,
				incBO.Results.BackupID,
				bod.kw,
				bod.kms,
				m365.AzureTenantID,
				userID,
				path.ExchangeService,
				map[path.CategoryType][]string{test.category: test.metadataFiles})
			deeTD.CheckBackupDetails(
				t,
				ctx,
				incBO.Results.BackupID,
				whatSet,
				bod.kms,
				bod.sss,
				expectDeets,
				false)

			// do some additional checks to ensure the incremental dealt with fewer items.
			assert.Greater(t, bo.Results.ItemsWritten, incBO.Results.ItemsWritten, "incremental items written")
			assert.Greater(t, bo.Results.ItemsRead, incBO.Results.ItemsRead, "incremental items read")
			assert.Greater(t, bo.Results.BytesRead, incBO.Results.BytesRead, "incremental bytes read")
			assert.Greater(t, bo.Results.BytesUploaded, incBO.Results.BytesUploaded, "incremental bytes uploaded")
			assert.Equal(t, bo.Results.ResourceOwners, incBO.Results.ResourceOwners, "incremental backup resource owner")
			assert.NoError(t, incBO.Errors.Failure(), "incremental non-recoverable error", clues.ToCore(bo.Errors.Failure()))
			assert.Empty(t, incBO.Errors.Recovered(), "count incremental recoverable/iteration errors")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "incremental backup-start events")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
			assert.Equal(t,
				incMB.CalledWith[events.BackupStart][0][events.BackupID],
				incBO.Results.BackupID, "incremental backupID pre-declaration")
		})
	}
}

func (suite *ExchangeBackupIntgSuite) TestBackup_Run_incrementalExchange() {
	testExchangeContinuousBackups(suite, control.Toggles{})
}

func (suite *ExchangeBackupIntgSuite) TestBackup_Run_incrementalNonDeltaExchange() {
	testExchangeContinuousBackups(suite, control.Toggles{DisableDelta: true})
}

func testExchangeContinuousBackups(suite *ExchangeBackupIntgSuite, toggles control.Toggles) {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	tester.LogTimeOfTest(t)

	var (
		acct       = tconfig.NewM365Account(t)
		mb         = evmock.NewBus()
		now        = dttm.Now()
		service    = path.ExchangeService
		categories = map[path.CategoryType][]string{
			path.EmailCategory:    exchange.MetadataFileNames(path.EmailCategory),
			path.ContactsCategory: exchange.MetadataFileNames(path.ContactsCategory),
			// path.EventsCategory:   exchange.MetadataFileNames(path.EventsCategory),
		}
		container1      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 1, now)
		container2      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 2, now)
		container3      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 3, now)
		containerRename = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 4, now)

		// container3 and containerRename don't exist yet.  Those will get created
		// later on during the tests.  Putting their identifiers into the selector
		// at this point is harmless.
		containers = []string{container1, container2, container3, containerRename}
		sel        = selectors.NewExchangeBackup([]string{suite.its.userID})
		whatSet    = deeTD.CategoryFromRepoRef
		opts       = control.Defaults()
	)

	opts.ToggleFeatures = toggles
	ctrl, sels := ControllerWithSelector(t, ctx, acct, resource.Users, sel.Selector, nil, nil)
	sel.DiscreteOwner = sels.ID()
	sel.DiscreteOwnerName = sels.Name()

	uidn := inMock.NewProvider(sels.ID(), sels.Name())

	sel.Include(
		sel.ContactFolders(containers, selectors.PrefixMatch()),
		// sel.EventCalendars(containers, selectors.PrefixMatch()),
		sel.MailFolders(containers, selectors.PrefixMatch()))

	creds, err := acct.M365Config()
	require.NoError(t, err, clues.ToCore(err))

	ac, err := api.NewClient(creds)
	require.NoError(t, err, clues.ToCore(err))

	// generate 3 new folders with two items each.
	// Only the first two folders will be part of the initial backup and
	// incrementals.  The third folder will be introduced partway through
	// the changes.
	// This should be enough to cover most delta actions, since moving one
	// container into another generates a delta for both addition and deletion.
	type contDeets struct {
		containerID string
		locRef      string
		itemRefs    []string // cached for populating expected deets, otherwise not used
	}

	mailDBF := func(id, timeStamp, subject, body string) []byte {
		return exchMock.MessageWith(
			suite.its.userID, suite.its.userID, suite.its.userID,
			subject, body, body,
			now, now, now, now)
	}

	contactDBF := func(id, timeStamp, subject, body string) []byte {
		given, mid, sur := id[:8], id[9:13], id[len(id)-12:]

		return exchMock.ContactBytesWith(
			given+" "+sur,
			sur+", "+given,
			given, mid, sur,
			"123-456-7890")
	}

	eventDBF := func(id, timeStamp, subject, body string) []byte {
		return exchMock.EventWith(
			suite.its.userID, subject, body, body,
			exchMock.NoOriginalStartDate, now, now,
			exchMock.NoRecurrence, exchMock.NoAttendees,
			exchMock.NoAttachments, exchMock.NoCancelledOccurrences,
			exchMock.NoExceptionOccurrences)
	}

	// test data set
	dataset := map[path.CategoryType]struct {
		dbf   dataBuilderFunc
		dests map[string]contDeets
	}{
		path.EmailCategory: {
			dbf: mailDBF,
			dests: map[string]contDeets{
				container1: {},
				container2: {},
			},
		},
		path.ContactsCategory: {
			dbf: contactDBF,
			dests: map[string]contDeets{
				container1: {},
				container2: {},
			},
		},
		// path.EventsCategory: {
		// 	dbf: eventDBF,
		// 	dests: map[string]contDeets{
		// 		container1: {},
		// 		container2: {},
		// 	},
		// },
	}

	// populate initial test data
	for category, gen := range dataset {
		for destName := range gen.dests {
			deets := generateContainerOfItems(
				t,
				ctx,
				ctrl,
				service,
				category,
				selectors.NewExchangeRestore([]string{uidn.ID()}).Selector,
				creds.AzureTenantID,
				uidn.ID(),
				"",
				destName,
				2,
				version.Backup,
				gen.dbf)

			itemRefs := []string{}

			for _, ent := range deets.Entries {
				if ent.Exchange == nil || ent.Folder != nil {
					continue
				}

				if len(ent.ItemRef) > 0 {
					itemRefs = append(itemRefs, ent.ItemRef)
				}
			}

			// save the item ids for building expectedDeets later on
			cd := dataset[category].dests[destName]
			cd.itemRefs = itemRefs
			dataset[category].dests[destName] = cd
		}
	}

	bo, bod := prepNewTestBackupOp(t, ctx, mb, sel.Selector, opts, version.Backup)
	defer bod.close(t, ctx)

	// run the initial backup
	runAndCheckBackup(t, ctx, &bo, mb, false)

	rrPfx, err := path.ServicePrefix(acct.ID(), uidn.ID(), service, path.EmailCategory)
	require.NoError(t, err, clues.ToCore(err))

	// strip the category from the prefix; we primarily want the tenant and resource owner.
	expectDeets := deeTD.NewInDeets(rrPfx.ToBuilder().Dir().String())
	bupDeets, _ := deeTD.GetDeetsInBackup(
		t,
		ctx,
		bo.Results.BackupID,
		acct.ID(),
		uidn.ID(),
		service,
		whatSet,
		bod.kms,
		bod.sss)

	// update the datasets with their location refs
	for category, gen := range dataset {
		for destName, cd := range gen.dests {
			var longestLR string

			for _, ent := range bupDeets.Entries {
				// generated destinations should always contain items
				if ent.Folder != nil {
					continue
				}

				p, err := path.FromDataLayerPath(ent.RepoRef, false)
				require.NoError(t, err, clues.ToCore(err))

				// category must match, and the owning folder must be this destination
				if p.Category() != category || strings.HasSuffix(ent.LocationRef, destName) {
					continue
				}

				// emails, due to folder nesting and our design for populating data via restore,
				// will duplicate the dest folder as both the restore destination, and the "old parent
				// folder".  we'll get both a prefix/destName and a prefix/destName/destName folder.
				// since we want future comparison to only use the leaf dir, we select for the longest match.
				if len(ent.LocationRef) > len(longestLR) {
					longestLR = ent.LocationRef
				}
			}

			require.NotEmptyf(
				t,
				longestLR,
				"must find a details entry matching the generated %s container: %s",
				category,
				destName)

			cd.locRef = longestLR

			dataset[category].dests[destName] = cd
			expectDeets.AddLocation(category.String(), cd.locRef)

			for _, i := range dataset[category].dests[destName].itemRefs {
				expectDeets.AddItem(category.String(), cd.locRef, i)
			}
		}
	}

	// verify test data was populated, and track it for comparisons
	// TODO: this can be swapped out for InDeets checks if we add itemRefs to folder ents.
	for category, gen := range dataset {
		cr := exchTD.PopulateContainerCache(t, ctx, ac, category, uidn.ID(), fault.New(true))

		for destName, dest := range gen.dests {
			id, ok := cr.LocationInCache(dest.locRef)
			require.True(t, ok, "dir %s found in %s cache", dest.locRef, category)

			dest.containerID = id
			dataset[category].dests[destName] = dest
		}
	}

	// precheck to ensure the expectedDeets are correct.
	// if we fail here, the expectedDeets were populated incorrectly.
	deeTD.CheckBackupDetails(
		t,
		ctx,
		bo.Results.BackupID,
		whatSet,
		bod.kms,
		bod.sss,
		expectDeets,
		true)

	// Although established as a table, these tests are not isolated from each other.
	// Assume that every test's side effects cascade to all following test cases.
	// The changes are split across the table so that we can monitor the deltas
	// in isolation, rather than debugging one change from the rest of a series.
	table := []struct {
		name string
		// performs the incremental update required for the test.
		//revive:disable-next-line:context-as-argument
		updateUserData       func(t *testing.T, ctx context.Context)
		deltaItemsRead       int
		deltaItemsWritten    int
		nonDeltaItemsRead    int
		nonDeltaItemsWritten int
		nonMetaItemsWritten  int
	}{
		{
			name:                 "clean, no changes",
			updateUserData:       func(t *testing.T, ctx context.Context) {},
			deltaItemsRead:       0,
			deltaItemsWritten:    0,
			nonDeltaItemsRead:    8,
			nonDeltaItemsWritten: 0, // unchanged items are not counted towards write
			nonMetaItemsWritten:  4,
		},
		{
			name: "move an email folder to a subfolder",
			updateUserData: func(t *testing.T, ctx context.Context) {
				cat := path.EmailCategory

				// contacts and events cannot be sufoldered; this is an email-only change
				from := dataset[cat].dests[container2]
				to := dataset[cat].dests[container1]

				body := users.NewItemMailFoldersItemMovePostRequestBody()
				body.SetDestinationId(ptr.To(to.containerID))

				err := ac.Mail().MoveContainer(ctx, uidn.ID(), from.containerID, body)
				require.NoError(t, err, clues.ToCore(err))

				newLoc := expectDeets.MoveLocation(cat.String(), from.locRef, to.locRef)
				from.locRef = newLoc
			},
			deltaItemsRead:       0, // zero because we don't count container reads
			deltaItemsWritten:    2,
			nonDeltaItemsRead:    8,
			nonDeltaItemsWritten: 2,
			nonMetaItemsWritten:  6,
		},
		{
			name: "delete a folder",
			updateUserData: func(t *testing.T, ctx context.Context) {
				for category, d := range dataset {
					containerID := d.dests[container2].containerID

					switch category {
					case path.EmailCategory:
						err := ac.Mail().DeleteContainer(ctx, uidn.ID(), containerID)
						require.NoError(t, err, "deleting an email folder", clues.ToCore(err))
					case path.ContactsCategory:
						err := ac.Contacts().DeleteContainer(ctx, uidn.ID(), containerID)
						require.NoError(t, err, "deleting a contacts folder", clues.ToCore(err))
					case path.EventsCategory:
						err := ac.Events().DeleteContainer(ctx, uidn.ID(), containerID)
						require.NoError(t, err, "deleting a calendar", clues.ToCore(err))
					}

					expectDeets.RemoveLocation(category.String(), d.dests[container2].locRef)
				}
			},
			deltaItemsRead:       0,
			deltaItemsWritten:    0, // deletions are not counted as "writes"
			nonDeltaItemsRead:    4,
			nonDeltaItemsWritten: 0,
			nonMetaItemsWritten:  4,
		},
		{
			name: "add a new folder",
			updateUserData: func(t *testing.T, ctx context.Context) {
				for category, gen := range dataset {
					deets := generateContainerOfItems(
						t,
						ctx,
						ctrl,
						service,
						category,
						selectors.NewExchangeRestore([]string{uidn.ID()}).Selector,
						creds.AzureTenantID, suite.its.userID, "", container3,
						2,
						version.Backup,
						gen.dbf)

					expectedLocRef := container3
					if category == path.EmailCategory {
						expectedLocRef = path.Builder{}.Append(container3, container3).String()
					}

					cr := exchTD.PopulateContainerCache(t, ctx, ac, category, uidn.ID(), fault.New(true))

					id, ok := cr.LocationInCache(expectedLocRef)
					require.Truef(t, ok, "dir %s found in %s cache", expectedLocRef, category)

					dataset[category].dests[container3] = contDeets{
						containerID: id,
						locRef:      expectedLocRef,
						itemRefs:    nil, // not needed at this point
					}

					for _, ent := range deets.Entries {
						if ent.Folder == nil {
							expectDeets.AddItem(category.String(), expectedLocRef, ent.ItemRef)
						}
					}
				}
			},
			deltaItemsRead:       4,
			deltaItemsWritten:    4,
			nonDeltaItemsRead:    8,
			nonDeltaItemsWritten: 4,
			nonMetaItemsWritten:  8,
		},
		{
			name: "rename a folder",
			updateUserData: func(t *testing.T, ctx context.Context) {
				for category, d := range dataset {
					containerID := d.dests[container3].containerID
					newLoc := containerRename

					if category == path.EmailCategory {
						newLoc = path.Builder{}.Append(container3, containerRename).String()
					}

					d.dests[containerRename] = contDeets{
						containerID: containerID,
						locRef:      newLoc,
					}

					expectDeets.RenameLocation(
						category.String(),
						d.dests[container3].containerID,
						newLoc)

					switch category {
					case path.EmailCategory:
						body, err := ac.Mail().GetFolder(ctx, uidn.ID(), containerID)
						require.NoError(t, err, clues.ToCore(err))

						body.SetDisplayName(&containerRename)
						err = ac.Mail().PatchFolder(ctx, uidn.ID(), containerID, body)
						require.NoError(t, err, clues.ToCore(err))

					case path.ContactsCategory:
						body, err := ac.Contacts().GetFolder(ctx, uidn.ID(), containerID)
						require.NoError(t, err, clues.ToCore(err))

						body.SetDisplayName(&containerRename)
						err = ac.Contacts().PatchFolder(ctx, uidn.ID(), containerID, body)
						require.NoError(t, err, clues.ToCore(err))

					case path.EventsCategory:
						body, err := ac.Events().GetCalendar(ctx, uidn.ID(), containerID)
						require.NoError(t, err, clues.ToCore(err))

						body.SetName(&containerRename)
						err = ac.Events().PatchCalendar(ctx, uidn.ID(), containerID, body)
						require.NoError(t, err, clues.ToCore(err))
					}
				}
			},
			deltaItemsRead: 0, // containers are not counted as reads
			// Renaming a folder doesn't cause kopia changes as the folder ID doesn't
			// change.
			deltaItemsWritten:    0,
			nonDeltaItemsRead:    8,
			nonDeltaItemsWritten: 0,
			nonMetaItemsWritten:  4,
		},
		{
			name: "add a new item",
			updateUserData: func(t *testing.T, ctx context.Context) {
				for category, d := range dataset {
					containerID := d.dests[container1].containerID

					switch category {
					case path.EmailCategory:
						_, itemData := generateItemData(t, category, uidn.ID(), mailDBF)
						body, err := api.BytesToMessageable(itemData)
						require.NoErrorf(t, err, "transforming mail bytes to messageable: %+v", clues.ToCore(err))

						itm, err := ac.Mail().PostItem(ctx, uidn.ID(), containerID, body)
						require.NoError(t, err, clues.ToCore(err))

						expectDeets.AddItem(
							category.String(),
							d.dests[category.String()].locRef,
							ptr.Val(itm.GetId()))

					case path.ContactsCategory:
						_, itemData := generateItemData(t, category, uidn.ID(), contactDBF)
						body, err := api.BytesToContactable(itemData)
						require.NoErrorf(t, err, "transforming contact bytes to contactable: %+v", clues.ToCore(err))

						itm, err := ac.Contacts().PostItem(ctx, uidn.ID(), containerID, body)
						require.NoError(t, err, clues.ToCore(err))

						expectDeets.AddItem(
							category.String(),
							d.dests[category.String()].locRef,
							ptr.Val(itm.GetId()))

					case path.EventsCategory:
						_, itemData := generateItemData(t, category, uidn.ID(), eventDBF)
						body, err := api.BytesToEventable(itemData)
						require.NoErrorf(t, err, "transforming event bytes to eventable: %+v", clues.ToCore(err))

						itm, err := ac.Events().PostItem(ctx, uidn.ID(), containerID, body)
						require.NoError(t, err, clues.ToCore(err))

						expectDeets.AddItem(
							category.String(),
							d.dests[category.String()].locRef,
							ptr.Val(itm.GetId()))
					}
				}
			},
			deltaItemsRead:       2,
			deltaItemsWritten:    2,
			nonDeltaItemsRead:    10,
			nonDeltaItemsWritten: 2,
			nonMetaItemsWritten:  6,
		},
		{
			name: "delete an existing item",
			updateUserData: func(t *testing.T, ctx context.Context) {
				for category, d := range dataset {
					containerID := d.dests[container1].containerID

					switch category {
					case path.EmailCategory:
						ids, _, _, err := ac.Mail().GetAddedAndRemovedItemIDs(ctx, uidn.ID(), containerID, "", false, true)
						require.NoError(t, err, "getting message ids", clues.ToCore(err))
						require.NotEmpty(t, ids, "message ids in folder")

						err = ac.Mail().DeleteItem(ctx, uidn.ID(), ids[0])
						require.NoError(t, err, "deleting email item", clues.ToCore(err))

						expectDeets.RemoveItem(
							category.String(),
							d.dests[category.String()].locRef,
							ids[0])

					case path.ContactsCategory:
						ids, _, _, err := ac.Contacts().GetAddedAndRemovedItemIDs(ctx, uidn.ID(), containerID, "", false, true)
						require.NoError(t, err, "getting contact ids", clues.ToCore(err))
						require.NotEmpty(t, ids, "contact ids in folder")

						err = ac.Contacts().DeleteItem(ctx, uidn.ID(), ids[0])
						require.NoError(t, err, "deleting contact item", clues.ToCore(err))

						expectDeets.RemoveItem(
							category.String(),
							d.dests[category.String()].locRef,
							ids[0])

					case path.EventsCategory:
						ids, _, _, err := ac.Events().GetAddedAndRemovedItemIDs(ctx, uidn.ID(), containerID, "", false, true)
						require.NoError(t, err, "getting event ids", clues.ToCore(err))
						require.NotEmpty(t, ids, "event ids in folder")

						err = ac.Events().DeleteItem(ctx, uidn.ID(), ids[0])
						require.NoError(t, err, "deleting calendar", clues.ToCore(err))

						expectDeets.RemoveItem(
							category.String(),
							d.dests[category.String()].locRef,
							ids[0])
					}
				}
			},
			deltaItemsRead:       2,
			deltaItemsWritten:    0, // deletes are not counted as "writes"
			nonDeltaItemsRead:    8,
			nonDeltaItemsWritten: 0,
			nonMetaItemsWritten:  4,
		},
	}

	for _, test := range table {
		suite.Run(test.name, func() {
			var (
				t     = suite.T()
				incMB = evmock.NewBus()
				atid  = creds.AzureTenantID
			)

			ctx, flush := tester.WithContext(t, ctx)
			defer flush()

			incBO := newTestBackupOp(t, ctx, bod, incMB, opts)

			suite.Run("PreTestSetup", func() {
				t := suite.T()

				ctx, flush := tester.WithContext(t, ctx)
				defer flush()

				test.updateUserData(t, ctx)
			})

			err := incBO.Run(ctx)
			require.NoError(t, err, clues.ToCore(err))

			bupID := incBO.Results.BackupID

			checkBackupIsInManifests(
				t,
				ctx,
				bod.kw,
				bod.sw,
				&incBO,
				sels,
				uidn.ID(),
				maps.Keys(categories)...)
			checkMetadataFilesExist(
				t,
				ctx,
				bupID,
				bod.kw,
				bod.kms,
				atid,
				uidn.ID(),
				service,
				categories)
			deeTD.CheckBackupDetails(
				t,
				ctx,
				bupID,
				whatSet,
				bod.kms,
				bod.sss,
				expectDeets,
				true)

			// FIXME: commented tests are flaky due to delta calls retaining data that is
			// out of scope of the test data.
			// we need to find a better way to make isolated assertions here.
			// The addition of the deeTD package gives us enough coverage to comment
			// out the tests for now and look to their improvemeng later.

			// do some additional checks to ensure the incremental dealt with fewer items.
			// +4 on read/writes to account for metadata: 1 delta and 1 path for each type.
			// if !toggles.DisableDelta {
			// assert.Equal(t, test.deltaItemsRead+4, incBO.Results.ItemsRead, "incremental items read")
			// assert.Equal(t, test.deltaItemsWritten+4, incBO.Results.ItemsWritten, "incremental items written")
			// } else {
			// assert.Equal(t, test.nonDeltaItemsRead+4, incBO.Results.ItemsRead, "non delta items read")
			// assert.Equal(t, test.nonDeltaItemsWritten+4, incBO.Results.ItemsWritten, "non delta items written")
			// }
			// assert.Equal(t, test.nonMetaItemsWritten, incBO.Results.ItemsWritten, "non meta incremental items write")
			assert.NoError(t, incBO.Errors.Failure(), "incremental non-recoverable error", clues.ToCore(incBO.Errors.Failure()))
			assert.Empty(t, incBO.Errors.Recovered(), "incremental recoverable/iteration errors")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "incremental backup-start events")
			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
			assert.Equal(t,
				incMB.CalledWith[events.BackupStart][0][events.BackupID],
				bupID, "incremental backupID pre-declaration")
		})
	}
}

type ExchangeRestoreIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestExchangeRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &ExchangeRestoreIntgSuite{
		Suite: tester.NewIntegrationSuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *ExchangeRestoreIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

func (suite *ExchangeRestoreIntgSuite) TestRestore_Run_exchangeWithAdvancedOptions() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// a backup is required to run restores

	baseSel := selectors.NewExchangeBackup([]string{suite.its.userID})
	baseSel.Include(
		// events cannot be run, for the same reason as incremental backups: the user needs
		// to have their account recycled.
		// base_sel.EventCalendars([]string{api.DefaultCalendar}, selectors.PrefixMatch()),
		baseSel.ContactFolders([]string{api.DefaultContacts}, selectors.PrefixMatch()),
		baseSel.MailFolders([]string{api.MailInbox}, selectors.PrefixMatch()))

	baseSel.DiscreteOwner = suite.its.userID

	var (
		mb   = evmock.NewBus()
		opts = control.Defaults()
	)

	bo, bod := prepNewTestBackupOp(t, ctx, mb, baseSel.Selector, opts, version.Backup)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)

	rsel, err := baseSel.ToExchangeRestore()
	require.NoError(t, err, clues.ToCore(err))

	var (
		restoreCfg = ctrlTD.DefaultRestoreConfig("exchange_adv_restore")
		sel        = rsel.Selector
		userID     = sel.ID()
		cIDs       = map[path.CategoryType]string{
			path.ContactsCategory: "",
			path.EmailCategory:    "",
			path.EventsCategory:   "",
		}
		countItemsInRestore int
		collKeys            = map[path.CategoryType]map[string]string{}
		acCont              = suite.its.ac.Contacts()
		acMail              = suite.its.ac.Mail()
		// acEvts   = suite.its.ac.Events()
	)

	// initial restore

	suite.Run("baseline", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Copy

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		runAndCheckRestore(t, ctx, &ro, mb, -1, false)

		// get all files in folder, use these as the base
		// set of files to compare against.
		contGC, err := acCont.GetContainerByName(ctx, userID, "", restoreCfg.Location)
		require.NoError(t, err, clues.ToCore(err))

		cIDs[path.ContactsCategory] = ptr.Val(contGC.GetId())

		collKeys[path.ContactsCategory], err = acCont.GetItemsInContainerByCollisionKey(
			ctx,
			userID,
			cIDs[path.ContactsCategory])
		require.NoError(t, err, clues.ToCore(err))
		countItemsInRestore += len(collKeys[path.ContactsCategory])

		// gc, err = acEvts.GetContainerByName(ctx, userID, "", restoreCfg.Location)
		// require.NoError(t, err, clues.ToCore(err))

		// restoredContainerID[path.EventsCategory] = ptr.Val(gc.GetId())

		// collKeys[path.EventsCategory], err = acEvts.GetItemsInContainerByCollisionKey(
		// 	ctx,
		// 	userID,
		// 	cIDs[path.EventsCategory])
		// require.NoError(t, err, clues.ToCore(err))
		// countItemsInRestore += len(collKeys[path.EventsCategory])

		mailGC, err := acMail.GetContainerByName(ctx, userID, api.MsgFolderRoot, restoreCfg.Location)
		require.NoError(t, err, clues.ToCore(err))

		mailGC, err = acMail.GetContainerByName(ctx, userID, ptr.Val(mailGC.GetId()), api.MailInbox)
		require.NoError(t, err, clues.ToCore(err))

		cIDs[path.EmailCategory] = ptr.Val(mailGC.GetId())

		collKeys[path.EmailCategory], err = acMail.GetItemsInContainerByCollisionKey(
			ctx,
			userID,
			cIDs[path.EmailCategory])
		require.NoError(t, err, clues.ToCore(err))
		countItemsInRestore += len(collKeys[path.EmailCategory])

		checkRestoreCounts(t, ctr, 0, 0, countItemsInRestore)
	})

	// skip restore

	suite.Run("skip collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Skip

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		deets := runAndCheckRestore(t, ctx, &ro, mb, 0, false)

		assert.Zero(
			t,
			len(deets.Entries),
			"no items should have been restored")

		checkRestoreCounts(t, ctr, countItemsInRestore, 0, 0)

		// get all files in folder, use these as the base
		// set of files to compare against.
		result := filterCollisionKeyResults(
			t,
			ctx,
			userID,
			cIDs[path.ContactsCategory],
			GetItemsInContainerByCollisionKeyer[string](acCont),
			collKeys[path.ContactsCategory])

		// m = checkCollisionKeyResults(t, ctx, userID, cIDs[path.EventsCategory], acEvts, collKeys[path.EventsCategory])
		// maps.Copy(result, m)

		m := filterCollisionKeyResults(
			t,
			ctx,
			userID,
			cIDs[path.EmailCategory],
			GetItemsInContainerByCollisionKeyer[string](acMail),
			collKeys[path.EmailCategory])
		maps.Copy(result, m)

		assert.Len(t, result, 0, "no new items should get added")
	})

	// replace restore

	suite.Run("replace collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Replace

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		deets := runAndCheckRestore(t, ctx, &ro, mb, len(collKeys), false)
		filtEnts := []details.Entry{}

		for _, e := range deets.Entries {
			if e.Folder == nil {
				filtEnts = append(filtEnts, e)
			}
		}

		assert.Equal(
			t,
			len(filtEnts),
			len(collKeys),
			"every item should have been replaced: %+v", filtEnts)

		checkRestoreCounts(t, ctr, 0, countItemsInRestore, 0)

		result := filterCollisionKeyResults(
			t,
			ctx,
			userID,
			cIDs[path.ContactsCategory],
			GetItemsInContainerByCollisionKeyer[string](acCont),
			collKeys[path.ContactsCategory])

		// m = checkCollisionKeyResults(t, ctx, userID, cIDs[path.EventsCategory], acEvts, collKeys[path.EventsCategory])
		// maps.Copy(result, m)

		m := filterCollisionKeyResults(
			t,
			ctx,
			userID,
			cIDs[path.EmailCategory],
			GetItemsInContainerByCollisionKeyer[string](acMail),
			collKeys[path.EmailCategory])
		maps.Copy(result, m)

		assert.Len(t, result, 0, "all items should have been replaced")
	})

	// copy restore

	suite.Run("copy collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr := count.New()

		restoreCfg.OnCollision = control.Copy

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr,
			sel,
			opts,
			restoreCfg)

		deets := runAndCheckRestore(t, ctx, &ro, mb, len(collKeys), false)
		filtEnts := []details.Entry{}

		for _, e := range deets.Entries {
			if e.Folder == nil {
				filtEnts = append(filtEnts, e)
			}
		}

		assert.Equal(
			t,
			len(filtEnts),
			countItemsInRestore,
			"every item should have been copied: %+v", filtEnts)

		checkRestoreCounts(t, ctr, 0, 0, countItemsInRestore)

		result := filterCollisionKeyResults(
			t,
			ctx,
			userID,
			cIDs[path.ContactsCategory],
			GetItemsInContainerByCollisionKeyer[string](acCont),
			collKeys[path.ContactsCategory])

		// m = checkCollisionKeyResults(t, ctx, userID, cIDs[path.EventsCategory], acEvts, collKeys[path.EventsCategory])
		// maps.Copy(result, m)

		m := filterCollisionKeyResults(
			t,
			ctx,
			userID,
			cIDs[path.EmailCategory],
			GetItemsInContainerByCollisionKeyer[string](acMail),
			collKeys[path.EmailCategory])
		maps.Copy(result, m)

		// TODO: we have the option of modifying copy creations in exchange
		// so that the results don't collide.  But we haven't made that
		// decision yet.
		assert.Len(t, result, 0, "no items should have been added as copies")
	})
}
