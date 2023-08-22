package test_test

import (
	"context"
	"testing"

	"github.com/alcionai/clues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/maps"

	"github.com/alcionai/corso/src/internal/common/ptr"
	evmock "github.com/alcionai/corso/src/internal/events/mock"
	"github.com/alcionai/corso/src/internal/m365/graph"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/internal/tester/tconfig"
	"github.com/alcionai/corso/src/internal/version"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/control"
	ctrlTD "github.com/alcionai/corso/src/pkg/control/testdata"
	"github.com/alcionai/corso/src/pkg/count"
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
// // for the following scopes: Contacts, Events, and Mail
// func (suite *ExchangeBackupIntgSuite) TestBackup_Run_exchange() {
// 	tests := []struct {
// 		name          string
// 		selector      func() *selectors.ExchangeBackup
// 		category      path.CategoryType
// 		metadataFiles []string
// 	}{
// 		{
// 			name: "Mail",
// 			selector: func() *selectors.ExchangeBackup {
// 				sel := selectors.NewExchangeBackup([]string{suite.its.user.ID})
// 				sel.Include(sel.MailFolders([]string{api.MailInbox}, selectors.PrefixMatch()))
// 				sel.DiscreteOwner = suite.its.user.ID

// 				return sel
// 			},
// 			category:      path.EmailCategory,
// 			metadataFiles: exchange.MetadataFileNames(path.EmailCategory),
// 		},
// 		{
// 			name: "Contacts",
// 			selector: func() *selectors.ExchangeBackup {
// 				sel := selectors.NewExchangeBackup([]string{suite.its.user.ID})
// 				sel.Include(sel.ContactFolders([]string{api.DefaultContacts}, selectors.PrefixMatch()))
// 				return sel
// 			},
// 			category:      path.ContactsCategory,
// 			metadataFiles: exchange.MetadataFileNames(path.ContactsCategory),
// 		},
// 		{
// 			name: "Calendar Events",
// 			selector: func() *selectors.ExchangeBackup {
// 				sel := selectors.NewExchangeBackup([]string{suite.its.user.ID})
// 				sel.Include(sel.EventCalendars([]string{api.DefaultCalendar}, selectors.PrefixMatch()))
// 				return sel
// 			},
// 			category:      path.EventsCategory,
// 			metadataFiles: exchange.MetadataFileNames(path.EventsCategory),
// 		},
// 	}
// 	for _, test := range tests {
// 		suite.Run(test.name, func() {
// 			t := suite.T()

// 			ctx, flush := tester.NewContext(t)
// 			defer flush()

// 			var (
// 				mb      = evmock.NewBus()
// 				sel     = test.selector().Selector
// 				opts    = control.DefaultOptions()
// 				whatSet = deeTD.CategoryFromRepoRef
// 			)

// 			bo, bod := prepNewTestBackupOp(t, ctx, mb, sel, opts, version.Backup)
// 			defer bod.close(t, ctx)

// 			sel = bod.sel

// 			userID := sel.ID()

// 			m365, err := bod.acct.M365Config()
// 			require.NoError(t, err, clues.ToCore(err))

// 			// run the tests
// 			runAndCheckBackup(t, ctx, &bo, mb, false)
// 			checkBackupIsInManifests(
// 				t,
// 				ctx,
// 				bod.kw,
// 				bod.sw,
// 				&bo,
// 				sel,
// 				userID,
// 				test.category)
// 			checkMetadataFilesExist(
// 				t,
// 				ctx,
// 				bo.Results.BackupID,
// 				bod.kw,
// 				bod.kms,
// 				m365.AzureTenantID,
// 				userID,
// 				path.ExchangeService,
// 				map[path.CategoryType][]string{test.category: test.metadataFiles})

// 			_, expectDeets := deeTD.GetDeetsInBackup(
// 				t,
// 				ctx,
// 				bo.Results.BackupID,
// 				bod.acct.ID(),
// 				userID,
// 				path.ExchangeService,
// 				whatSet,
// 				bod.kms,
// 				bod.sss)
// 			deeTD.CheckBackupDetails(
// 				t,
// 				ctx,
// 				bo.Results.BackupID,
// 				whatSet,
// 				bod.kms,
// 				bod.sss,
// 				expectDeets,
// 				false)

// 			// Basic, happy path incremental test.  No changes are dictated or expected.
// 			// This only tests that an incremental backup is runnable at all, and that it
// 			// produces fewer results than the last backup.
// 			var (
// 				incMB = evmock.NewBus()
// 				incBO = newTestBackupOp(
// 					t,
// 					ctx,
// 					bod,
// 					incMB,
// 					opts)
// 			)

// 			runAndCheckBackup(t, ctx, &incBO, incMB, true)
// 			checkBackupIsInManifests(
// 				t,
// 				ctx,
// 				bod.kw,
// 				bod.sw,
// 				&incBO,
// 				sel,
// 				userID,
// 				test.category)
// 			checkMetadataFilesExist(
// 				t,
// 				ctx,
// 				incBO.Results.BackupID,
// 				bod.kw,
// 				bod.kms,
// 				m365.AzureTenantID,
// 				userID,
// 				path.ExchangeService,
// 				map[path.CategoryType][]string{test.category: test.metadataFiles})
// 			deeTD.CheckBackupDetails(
// 				t,
// 				ctx,
// 				incBO.Results.BackupID,
// 				whatSet,
// 				bod.kms,
// 				bod.sss,
// 				expectDeets,
// 				false)

// 			// do some additional checks to ensure the incremental dealt with fewer items.
// 			assert.Greater(t, bo.Results.ItemsWritten, incBO.Results.ItemsWritten, "incremental items written")
// 			assert.Greater(t, bo.Results.ItemsRead, incBO.Results.ItemsRead, "incremental items read")
// 			assert.Greater(t, bo.Results.BytesRead, incBO.Results.BytesRead, "incremental bytes read")
// 			assert.Greater(t, bo.Results.BytesUploaded, incBO.Results.BytesUploaded, "incremental bytes uploaded")
// 			assert.Equal(t, bo.Results.ResourceOwners, incBO.Results.ResourceOwners, "incremental backup resource owner")
// 			assert.NoError(t, incBO.Errors.Failure(), "incremental non-recoverable error", clues.ToCore(bo.Errors.Failure()))
// 			assert.Empty(t, incBO.Errors.Recovered(), "count incremental recoverable/iteration errors")
// 			assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "incremental backup-start events")
// 			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
// 			assert.Equal(t,
// 				incMB.CalledWith[events.BackupStart][0][events.BackupID],
// 				incBO.Results.BackupID, "incremental backupID pre-declaration")
// 		})
// 	}
// }

// func (suite *ExchangeBackupIntgSuite) TestBackup_Run_incrementalExchange() {
// 	testExchangeContinuousBackups(suite, control.Toggles{})
// }

// func (suite *ExchangeBackupIntgSuite) TestBackup_Run_incrementalNonDeltaExchange() {
// 	testExchangeContinuousBackups(suite, control.Toggles{DisableDelta: true})
// }

// func testExchangeContinuousBackups(suite *ExchangeBackupIntgSuite, toggles control.Toggles) {
// 	t := suite.T()

// 	ctx, flush := tester.NewContext(t)
// 	defer flush()

// 	tester.LogTimeOfTest(t)

// 	var (
// 		acct       = tconfig.NewM365Account(t)
// 		mb         = evmock.NewBus()
// 		now        = dttm.Now()
// 		service    = path.ExchangeService
// 		categories = map[path.CategoryType][]string{
// 			path.EmailCategory:    exchange.MetadataFileNames(path.EmailCategory),
// 			path.ContactsCategory: exchange.MetadataFileNames(path.ContactsCategory),
// 			// path.EventsCategory:   exchange.MetadataFileNames(path.EventsCategory),
// 		}
// 		container1      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 1, now)
// 		container2      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 2, now)
// 		container3      = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 3, now)
// 		containerRename = fmt.Sprintf("%s%d_%s", incrementalsDestContainerPrefix, 4, now)

// 		// container3 and containerRename don't exist yet.  Those will get created
// 		// later on during the tests.  Putting their identifiers into the selector
// 		// at this point is harmless.
// 		containers = []string{container1, container2, container3, containerRename}
// 		sel        = selectors.NewExchangeBackup([]string{suite.its.user.ID})
// 		whatSet    = deeTD.CategoryFromRepoRef
// 		opts       = control.DefaultOptions()
// 	)

// 	opts.ToggleFeatures = toggles
// 	ctrl, sels := ControllerWithSelector(t, ctx, acct, resource.Users, sel.Selector, nil, nil)
// 	sel.DiscreteOwner = sels.ID()
// 	sel.DiscreteOwnerName = sels.Name()

// 	uidn := inMock.NewProvider(sels.ID(), sels.Name())

// 	sel.Include(
// 		sel.ContactFolders(containers, selectors.PrefixMatch()),
// 		// sel.EventCalendars(containers, selectors.PrefixMatch()),
// 		sel.MailFolders(containers, selectors.PrefixMatch()))

// 	creds, err := acct.M365Config()
// 	require.NoError(t, err, clues.ToCore(err))

// 	ac, err := api.NewClient(creds, control.DefaultOptions())
// 	require.NoError(t, err, clues.ToCore(err))

// 	// generate 3 new folders with two items each.
// 	// Only the first two folders will be part of the initial backup and
// 	// incrementals.  The third folder will be introduced partway through
// 	// the changes.
// 	// This should be enough to cover most delta actions, since moving one
// 	// container into another generates a delta for both addition and deletion.
// 	type contDeets struct {
// 		containerID string
// 		locRef      string
// 		itemRefs    []string // cached for populating expected deets, otherwise not used
// 	}

// 	mailDBF := func(id, timeStamp, subject, body string) []byte {
// 		return exchMock.MessageWith(
// 			suite.its.user.ID, suite.its.user.ID, suite.its.user.ID,
// 			subject, body, body,
// 			now, now, now, now)
// 	}

// 	contactDBF := func(id, timeStamp, subject, body string) []byte {
// 		given, mid, sur := id[:8], id[9:13], id[len(id)-12:]

// 		return exchMock.ContactBytesWith(
// 			given+" "+sur,
// 			sur+", "+given,
// 			given, mid, sur,
// 			"123-456-7890")
// 	}

// 	eventDBF := func(id, timeStamp, subject, body string) []byte {
// 		return exchMock.EventWith(
// 			suite.its.user.ID, subject, body, body,
// 			exchMock.NoOriginalStartDate, now, now,
// 			exchMock.NoRecurrence, exchMock.NoAttendees,
// 			exchMock.NoAttachments, exchMock.NoCancelledOccurrences,
// 			exchMock.NoExceptionOccurrences)
// 	}

// 	// test data set
// 	dataset := map[path.CategoryType]struct {
// 		dbf   dataBuilderFunc
// 		dests map[string]contDeets
// 	}{
// 		path.EmailCategory: {
// 			dbf: mailDBF,
// 			dests: map[string]contDeets{
// 				container1: {},
// 				container2: {},
// 			},
// 		},
// 		path.ContactsCategory: {
// 			dbf: contactDBF,
// 			dests: map[string]contDeets{
// 				container1: {},
// 				container2: {},
// 			},
// 		},
// 		// path.EventsCategory: {
// 		// 	dbf: eventDBF,
// 		// 	dests: map[string]contDeets{
// 		// 		container1: {},
// 		// 		container2: {},
// 		// 	},
// 		// },
// 	}

// 	rrPfx, err := path.BuildPrefix(acct.ID(), uidn.ID(), service, path.EmailCategory)
// 	require.NoError(t, err, clues.ToCore(err))

// 	// strip the category from the prefix; we primarily want the tenant and resource owner.
// 	expectDeets := deeTD.NewInDeets(rrPfx.ToBuilder().Dir().String())

// 	mustGetExpectedContainerItems := func(
// 		t *testing.T,
// 		category path.CategoryType,
// 		cr graph.ContainerResolver,
// 		destName string,
// 	) {
// 		locRef := path.Builder{}.Append(destName)

// 		if category == path.EmailCategory {
// 			locRef = locRef.Append(destName)
// 		}

// 		containerID, ok := cr.LocationInCache(locRef.String())
// 		require.True(t, ok, "dir %s found in %s cache", locRef.String(), category)

// 		var (
// 			err   error
// 			items []string
// 		)

// 		switch category {
// 		case path.EmailCategory:
// 			items, _, _, err = ac.Mail().GetAddedAndRemovedItemIDs(
// 				ctx,
// 				uidn.ID(),
// 				containerID,
// 				"",
// 				toggles.ExchangeImmutableIDs,
// 				true)

// 		case path.EventsCategory:
// 			items, _, _, err = ac.Events().GetAddedAndRemovedItemIDs(
// 				ctx,
// 				uidn.ID(),
// 				containerID,
// 				"",
// 				toggles.ExchangeImmutableIDs,
// 				true)

// 		case path.ContactsCategory:
// 			items, _, _, err = ac.Contacts().GetAddedAndRemovedItemIDs(
// 				ctx,
// 				uidn.ID(),
// 				containerID,
// 				"",
// 				toggles.ExchangeImmutableIDs,
// 				true)
// 		}

// 		require.NoError(
// 			t,
// 			err,
// 			"getting items for category %s, container %s",
// 			category,
// 			locRef.String())

// 		dest := dataset[category].dests[destName]
// 		dest.locRef = locRef.String()
// 		dest.containerID = containerID
// 		dest.itemRefs = items
// 		dataset[category].dests[destName] = dest

// 		// Add the directory and all its ancestors to the cache so we can compare
// 		// folders.
// 		for len(locRef.Elements()) > 0 {
// 			expectDeets.AddLocation(category.String(), locRef.String())
// 			locRef = locRef.Dir()
// 		}

// 		for _, i := range dataset[category].dests[destName].itemRefs {
// 			expectDeets.AddItem(category.String(), dest.locRef, i)
// 		}
// 	}

// 	// populate initial test data
// 	for category, gen := range dataset {
// 		for destName := range gen.dests {
// 			generateContainerOfItems(
// 				t,
// 				ctx,
// 				ctrl,
// 				service,
// 				category,
// 				selectors.NewExchangeRestore([]string{uidn.ID()}).Selector,
// 				creds.AzureTenantID,
// 				uidn.ID(),
// 				"",
// 				destName,
// 				2,
// 				version.Backup,
// 				gen.dbf)
// 		}

// 		cr := exchTD.PopulateContainerCache(
// 			t,
// 			ctx,
// 			ac,
// 			category,
// 			uidn.ID(),
// 			fault.New(true))

// 		for destName := range gen.dests {
// 			mustGetExpectedContainerItems(t, category, cr, destName)
// 		}
// 	}

// 	bo, bod := prepNewTestBackupOp(t, ctx, mb, sel.Selector, opts, version.Backup)
// 	defer bod.close(t, ctx)

// 	// run the initial backup
// 	runAndCheckBackup(t, ctx, &bo, mb, false)

// 	// precheck to ensure the expectedDeets are correct.
// 	// if we fail here, the expectedDeets were populated incorrectly.
// 	deeTD.CheckBackupDetails(
// 		t,
// 		ctx,
// 		bo.Results.BackupID,
// 		whatSet,
// 		bod.kms,
// 		bod.sss,
// 		expectDeets,
// 		true)

// 	// Although established as a table, these tests are not isolated from each other.
// 	// Assume that every test's side effects cascade to all following test cases.
// 	// The changes are split across the table so that we can monitor the deltas
// 	// in isolation, rather than debugging one change from the rest of a series.
// 	table := []struct {
// 		name string
// 		// performs the incremental update required for the test.
// 		//revive:disable-next-line:context-as-argument
// 		updateUserData       func(t *testing.T, ctx context.Context)
// 		deltaItemsRead       int
// 		deltaItemsWritten    int
// 		nonDeltaItemsRead    int
// 		nonDeltaItemsWritten int
// 		nonMetaItemsWritten  int
// 	}{
// 		{
// 			name:                 "clean, no changes",
// 			updateUserData:       func(t *testing.T, ctx context.Context) {},
// 			deltaItemsRead:       0,
// 			deltaItemsWritten:    0,
// 			nonDeltaItemsRead:    8,
// 			nonDeltaItemsWritten: 0, // unchanged items are not counted towards write
// 			nonMetaItemsWritten:  4,
// 		},
// 		{
// 			name: "move an email folder to a subfolder",
// 			updateUserData: func(t *testing.T, ctx context.Context) {
// 				cat := path.EmailCategory

// 				// contacts and events cannot be sufoldered; this is an email-only change
// 				from := dataset[cat].dests[container2]
// 				to := dataset[cat].dests[container1]

// 				body := users.NewItemMailFoldersItemMovePostRequestBody()
// 				body.SetDestinationId(ptr.To(to.containerID))

// 				err := ac.Mail().MoveContainer(ctx, uidn.ID(), from.containerID, body)
// 				require.NoError(t, err, clues.ToCore(err))

// 				newLoc := expectDeets.MoveLocation(cat.String(), from.locRef, to.locRef)

// 				// Remove ancestor folders of moved directory since they'll no longer
// 				// appear in details since we're not backing up items in them.
// 				pb, err := path.Builder{}.SplitUnescapeAppend(from.locRef)
// 				require.NoError(t, err, "getting Builder for location: %s", clues.ToCore(err))

// 				pb = pb.Dir()

// 				for len(pb.Elements()) > 0 {
// 					expectDeets.RemoveLocation(cat.String(), pb.String())
// 					pb = pb.Dir()
// 				}

// 				// Update cache with new location of container.
// 				from.locRef = newLoc
// 				dataset[cat].dests[container2] = from
// 			},
// 			deltaItemsRead:       0, // zero because we don't count container reads
// 			deltaItemsWritten:    2,
// 			nonDeltaItemsRead:    8,
// 			nonDeltaItemsWritten: 2,
// 			nonMetaItemsWritten:  6,
// 		},
// 		{
// 			name: "delete a folder",
// 			updateUserData: func(t *testing.T, ctx context.Context) {
// 				for category, d := range dataset {
// 					containerID := d.dests[container2].containerID

// 					switch category {
// 					case path.EmailCategory:
// 						err := ac.Mail().DeleteContainer(ctx, uidn.ID(), containerID)
// 						require.NoError(t, err, "deleting an email folder", clues.ToCore(err))
// 					case path.ContactsCategory:
// 						err := ac.Contacts().DeleteContainer(ctx, uidn.ID(), containerID)
// 						require.NoError(t, err, "deleting a contacts folder", clues.ToCore(err))
// 					case path.EventsCategory:
// 						err := ac.Events().DeleteContainer(ctx, uidn.ID(), containerID)
// 						require.NoError(t, err, "deleting a calendar", clues.ToCore(err))
// 					}

// 					expectDeets.RemoveLocation(category.String(), d.dests[container2].locRef)
// 				}
// 			},
// 			deltaItemsRead:       0,
// 			deltaItemsWritten:    0, // deletions are not counted as "writes"
// 			nonDeltaItemsRead:    4,
// 			nonDeltaItemsWritten: 0,
// 			nonMetaItemsWritten:  4,
// 		},
// 		{
// 			name: "add a new folder",
// 			updateUserData: func(t *testing.T, ctx context.Context) {
// 				for category, gen := range dataset {
// 					generateContainerOfItems(
// 						t,
// 						ctx,
// 						ctrl,
// 						service,
// 						category,
// 						selectors.NewExchangeRestore([]string{uidn.ID()}).Selector,
// 						creds.AzureTenantID, suite.its.user.ID, "", container3,
// 						2,
// 						version.Backup,
// 						gen.dbf)

// 					cr := exchTD.PopulateContainerCache(t, ctx, ac, category, uidn.ID(), fault.New(true))
// 					mustGetExpectedContainerItems(t, category, cr, container3)
// 				}
// 			},
// 			deltaItemsRead:       4,
// 			deltaItemsWritten:    4,
// 			nonDeltaItemsRead:    8,
// 			nonDeltaItemsWritten: 4,
// 			nonMetaItemsWritten:  8,
// 		},
// 		{
// 			name: "rename a folder",
// 			updateUserData: func(t *testing.T, ctx context.Context) {
// 				for category, d := range dataset {
// 					containerID := d.dests[container3].containerID
// 					newLoc := containerRename

// 					if category == path.EmailCategory {
// 						newLoc = path.Builder{}.Append(container3, containerRename).String()
// 					}

// 					d.dests[containerRename] = contDeets{
// 						containerID: containerID,
// 						locRef:      newLoc,
// 					}

// 					expectDeets.RenameLocation(
// 						category.String(),
// 						d.dests[container3].locRef,
// 						newLoc)

// 					switch category {
// 					case path.EmailCategory:
// 						body := models.NewMailFolder()
// 						body.SetDisplayName(ptr.To(containerRename))

// 						err := ac.Mail().PatchFolder(ctx, uidn.ID(), containerID, body)
// 						require.NoError(t, err, clues.ToCore(err))

// 					case path.ContactsCategory:
// 						body := models.NewContactFolder()
// 						body.SetDisplayName(ptr.To(containerRename))

// 						err = ac.Contacts().PatchFolder(ctx, uidn.ID(), containerID, body)
// 						require.NoError(t, err, clues.ToCore(err))

// 					case path.EventsCategory:
// 						body := models.NewCalendar()
// 						body.SetName(ptr.To(containerRename))

// 						err = ac.Events().PatchCalendar(ctx, uidn.ID(), containerID, body)
// 						require.NoError(t, err, clues.ToCore(err))
// 					}
// 				}
// 			},
// 			deltaItemsRead: 0, // containers are not counted as reads
// 			// Renaming a folder doesn't cause kopia changes as the folder ID doesn't
// 			// change.
// 			deltaItemsWritten:    0,
// 			nonDeltaItemsRead:    8,
// 			nonDeltaItemsWritten: 0,
// 			nonMetaItemsWritten:  4,
// 		},
// 		{
// 			name: "add a new item",
// 			updateUserData: func(t *testing.T, ctx context.Context) {
// 				for category, d := range dataset {
// 					containerID := d.dests[container1].containerID

// 					switch category {
// 					case path.EmailCategory:
// 						_, itemData := generateItemData(t, category, uidn.ID(), mailDBF)
// 						body, err := api.BytesToMessageable(itemData)
// 						require.NoErrorf(t, err, "transforming mail bytes to messageable: %+v", clues.ToCore(err))

// 						itm, err := ac.Mail().PostItem(ctx, uidn.ID(), containerID, body)
// 						require.NoError(t, err, clues.ToCore(err))

// 						expectDeets.AddItem(
// 							category.String(),
// 							d.dests[container1].locRef,
// 							ptr.Val(itm.GetId()))

// 					case path.ContactsCategory:
// 						_, itemData := generateItemData(t, category, uidn.ID(), contactDBF)
// 						body, err := api.BytesToContactable(itemData)
// 						require.NoErrorf(t, err, "transforming contact bytes to contactable: %+v", clues.ToCore(err))

// 						itm, err := ac.Contacts().PostItem(ctx, uidn.ID(), containerID, body)
// 						require.NoError(t, err, clues.ToCore(err))

// 						expectDeets.AddItem(
// 							category.String(),
// 							d.dests[container1].locRef,
// 							ptr.Val(itm.GetId()))

// 					case path.EventsCategory:
// 						_, itemData := generateItemData(t, category, uidn.ID(), eventDBF)
// 						body, err := api.BytesToEventable(itemData)
// 						require.NoErrorf(t, err, "transforming event bytes to eventable: %+v", clues.ToCore(err))

// 						itm, err := ac.Events().PostItem(ctx, uidn.ID(), containerID, body)
// 						require.NoError(t, err, clues.ToCore(err))

// 						expectDeets.AddItem(
// 							category.String(),
// 							d.dests[container1].locRef,
// 							ptr.Val(itm.GetId()))
// 					}
// 				}
// 			},
// 			deltaItemsRead:       2,
// 			deltaItemsWritten:    2,
// 			nonDeltaItemsRead:    10,
// 			nonDeltaItemsWritten: 2,
// 			nonMetaItemsWritten:  6,
// 		},
// 		{
// 			// Events and contacts have no Graph API call to move something between
// 			// containers. The calendars web UI does support moving events between
// 			// calendars though.
// 			name: "boomerang an email",
// 			updateUserData: func(t *testing.T, ctx context.Context) {
// 				containerInfo := dataset[path.EmailCategory].dests[container1]
// 				tempContainerID := dataset[path.EmailCategory].dests[container3].containerID

// 				ids := dataset[path.EmailCategory].dests[container1].itemRefs
// 				require.NotEmpty(t, ids, "message ids in folder")

// 				oldID := ids[0]

// 				newID, err := ac.Mail().MoveItem(
// 					ctx,
// 					uidn.ID(),
// 					containerInfo.containerID,
// 					tempContainerID,
// 					oldID)
// 				require.NoError(t, err, "moving to temp folder: %s", clues.ToCore(err))

// 				newID, err = ac.Mail().MoveItem(
// 					ctx,
// 					uidn.ID(),
// 					tempContainerID,
// 					containerInfo.containerID,
// 					newID)
// 				require.NoError(t, err, "moving back to original folder: %s", clues.ToCore(err))

// 				expectDeets.RemoveItem(
// 					path.EmailCategory.String(),
// 					containerInfo.locRef,
// 					oldID)
// 				expectDeets.AddItem(
// 					path.EmailCategory.String(),
// 					containerInfo.locRef,
// 					newID)

// 				// Will cause a different item to be deleted next.
// 				containerInfo.itemRefs = append(containerInfo.itemRefs[1:], newID)
// 				dataset[path.EmailCategory].dests[container1] = containerInfo
// 			},
// 			// TODO(ashmrtn): Below values need updated when we start checking them
// 			// again. Unclear what items would be considered the same as I'm not sure
// 			// about all the properties that change with a move.
// 			deltaItemsRead:       2,
// 			deltaItemsWritten:    2,
// 			nonDeltaItemsRead:    10,
// 			nonDeltaItemsWritten: 2,
// 			nonMetaItemsWritten:  6,
// 		},
// 		{
// 			name: "delete an existing item",
// 			updateUserData: func(t *testing.T, ctx context.Context) {
// 				for category, d := range dataset {
// 					containerInfo := d.dests[container1]
// 					require.NotEmpty(t, containerInfo.itemRefs)

// 					id := containerInfo.itemRefs[0]

// 					switch category {
// 					case path.EmailCategory:
// 						err = ac.Mail().DeleteItem(ctx, uidn.ID(), id)

// 					case path.ContactsCategory:
// 						err = ac.Contacts().DeleteItem(ctx, uidn.ID(), id)

// 					case path.EventsCategory:
// 						err = ac.Events().DeleteItem(ctx, uidn.ID(), id)
// 					}

// 					require.NoError(
// 						t,
// 						err,
// 						"deleting %s item: %s",
// 						category.String(),
// 						clues.ToCore(err))
// 					expectDeets.RemoveItem(category.String(), containerInfo.locRef, id)
// 				}
// 			},
// 			deltaItemsRead:       2,
// 			deltaItemsWritten:    0, // deletes are not counted as "writes"
// 			nonDeltaItemsRead:    8,
// 			nonDeltaItemsWritten: 0,
// 			nonMetaItemsWritten:  4,
// 		},
// 	}

// 	for _, test := range table {
// 		suite.Run(test.name, func() {
// 			var (
// 				t     = suite.T()
// 				incMB = evmock.NewBus()
// 				atid  = creds.AzureTenantID
// 			)

// 			ctx, flush := tester.WithContext(t, ctx)
// 			defer flush()

// 			incBO := newTestBackupOp(t, ctx, bod, incMB, opts)

// 			suite.Run("PreTestSetup", func() {
// 				t := suite.T()

// 				ctx, flush := tester.WithContext(t, ctx)
// 				defer flush()

// 				test.updateUserData(t, ctx)
// 			})

// 			err := incBO.Run(ctx)
// 			require.NoError(t, err, clues.ToCore(err))

// 			bupID := incBO.Results.BackupID

// 			checkBackupIsInManifests(
// 				t,
// 				ctx,
// 				bod.kw,
// 				bod.sw,
// 				&incBO,
// 				sels,
// 				uidn.ID(),
// 				maps.Keys(categories)...)
// 			checkMetadataFilesExist(
// 				t,
// 				ctx,
// 				bupID,
// 				bod.kw,
// 				bod.kms,
// 				atid,
// 				uidn.ID(),
// 				service,
// 				categories)
// 			deeTD.CheckBackupDetails(
// 				t,
// 				ctx,
// 				bupID,
// 				whatSet,
// 				bod.kms,
// 				bod.sss,
// 				expectDeets,
// 				true)

// 			// FIXME: commented tests are flaky due to delta calls retaining data that is
// 			// out of scope of the test data.
// 			// we need to find a better way to make isolated assertions here.
// 			// The addition of the deeTD package gives us enough coverage to comment
// 			// out the tests for now and look to their improvemeng later.

// 			// do some additional checks to ensure the incremental dealt with fewer items.
// 			// +4 on read/writes to account for metadata: 1 delta and 1 path for each type.
// 			// if !toggles.DisableDelta {
// 			// assert.Equal(t, test.deltaItemsRead+4, incBO.Results.ItemsRead, "incremental items read")
// 			// assert.Equal(t, test.deltaItemsWritten+4, incBO.Results.ItemsWritten, "incremental items written")
// 			// } else {
// 			// assert.Equal(t, test.nonDeltaItemsRead+4, incBO.Results.ItemsRead, "non delta items read")
// 			// assert.Equal(t, test.nonDeltaItemsWritten+4, incBO.Results.ItemsWritten, "non delta items written")
// 			// }
// 			// assert.Equal(t, test.nonMetaItemsWritten, incBO.Results.ItemsWritten, "non meta incremental items write")
// 			assert.NoError(t, incBO.Errors.Failure(), "incremental non-recoverable error", clues.ToCore(incBO.Errors.Failure()))
// 			assert.Empty(t, incBO.Errors.Recovered(), "incremental recoverable/iteration errors")
// 			assert.Equal(t, 1, incMB.TimesCalled[events.BackupStart], "incremental backup-start events")
// 			assert.Equal(t, 1, incMB.TimesCalled[events.BackupEnd], "incremental backup-end events")
// 			assert.Equal(t,
// 				incMB.CalledWith[events.BackupStart][0][events.BackupID],
// 				bupID, "incremental backupID pre-declaration")
// 		})
// 	}
// }

type ExchangeRestoreNightlyIntgSuite struct {
	tester.Suite
	its intgTesterSetup
}

func TestExchangeRestoreIntgSuite(t *testing.T) {
	suite.Run(t, &ExchangeRestoreNightlyIntgSuite{
		Suite: tester.NewNightlySuite(
			t,
			[][]string{tconfig.M365AcctCredEnvs, storeTD.AWSStorageCredEnvs}),
	})
}

func (suite *ExchangeRestoreNightlyIntgSuite) SetupSuite() {
	suite.its = newIntegrationTesterSetup(suite.T())
}

type clientItemPager interface {
	GetItemsInContainerByCollisionKeyer[string]
	GetItemIDsInContainer(
		ctx context.Context,
		userID, containerID string,
	) (map[string]struct{}, error)
	GetContainerByName(
		ctx context.Context,
		userID, parentContainerID, containerName string,
	) (graph.Container, error)
	GetItemsInContainerByCollisionKey(
		ctx context.Context,
		userID, containerID string,
	) (map[string]string, error)
	CreateContainer(
		ctx context.Context,
		userID, parentContainerID, containerName string,
	) (graph.Container, error)
}

func (suite *ExchangeRestoreNightlyIntgSuite) TestRestore_Run_exchangeWithAdvancedOptions() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// a backup is required to run restores

	baseSel := selectors.NewExchangeBackup([]string{suite.its.user.ID})
	baseSel.Include(
		// events cannot be run, for the same reason as incremental backups: the user needs
		// to have their account recycled.
		// base_sel.EventCalendars([]string{api.DefaultCalendar}, selectors.PrefixMatch()),
		//baseSel.ContactFolders([]string{api.DefaultContacts}, selectors.PrefixMatch()),
		baseSel.MailFolders([]string{api.MailInbox}, selectors.PrefixMatch()))

	baseSel.DiscreteOwner = suite.its.user.ID

	var (
		mb   = evmock.NewBus()
		opts = control.DefaultOptions()
	)

	bo, bod := prepNewTestBackupOp(t, ctx, mb, baseSel.Selector, opts, version.Backup)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)

	rsel, err := baseSel.ToExchangeRestore()
	require.NoError(t, err, clues.ToCore(err))

	var (
		restoreCfg          = ctrlTD.DefaultRestoreConfig("exchange_adv_restore")
		sel                 = rsel.Selector
		userID              = sel.ID()
		countItemsInRestore int

		itemIDs            = map[path.CategoryType]map[string]struct{}{}
		collisionKeys      = map[path.CategoryType]map[string]string{}
		containerIDs       = map[path.CategoryType]string{}
		parentContainerIDs = map[path.CategoryType]string{
			path.EmailCategory: api.MsgFolderRoot,
		}
		parentContainerNames = map[path.CategoryType][]string{
			path.EmailCategory:    {api.MailInbox},
			path.ContactsCategory: {},
			path.EventsCategory:   {},
		}

		testCategories = map[path.CategoryType]clientItemPager{
			//path.ContactsCategory: suite.its.ac.Contacts(),
			path.EmailCategory: suite.its.ac.Mail(),
			// path.EventsCategory: suite.its.ac.Events(),
		}
	)

	// initial restore

	suite.Run("baseline", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr1 := count.New()

		restoreCfg.OnCollision = control.Copy

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr1,
			sel,
			opts,
			restoreCfg)

		runAndCheckRestore(t, ctx, &ro, mb, false)

		// get all files in folder, use these as the base
		// set of files to compare against.

		for cat, ac := range testCategories {
			suite.Run(cat.String(), func() {
				t := suite.T()

				ctx, flush := tester.NewContext(t)
				defer flush()

				containers := append([]string{restoreCfg.Location}, parentContainerNames[cat]...)

				itemIDs[cat], collisionKeys[cat], containerIDs[cat] = getCollKeysAndItemIDs(
					t,
					ctx,
					ac,
					userID,
					parentContainerIDs[cat],
					containers...)

				countItemsInRestore += len(collisionKeys[cat])
			})
		}

		checkRestoreCounts(t, ctr1, 0, 0, countItemsInRestore)
	})

	// skip restore

	suite.Run("skip collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr2 := count.New()

		restoreCfg.OnCollision = control.Skip

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr2,
			sel,
			opts,
			restoreCfg)

		deets := runAndCheckRestore(t, ctx, &ro, mb, false)

		assert.Zero(
			t,
			len(deets.Entries),
			"no items should have been restored")

		checkRestoreCounts(t, ctr2, countItemsInRestore, 0, 0)

		result := map[string]string{}

		for cat, ac := range testCategories {
			suite.Run(cat.String(), func() {
				t := suite.T()

				ctx, flush := tester.NewContext(t)
				defer flush()

				m := filterCollisionKeyResults(
					t,
					ctx,
					userID,
					containerIDs[cat],
					GetItemsInContainerByCollisionKeyer[string](ac),
					collisionKeys[cat])
				maps.Copy(result, m)

				currentIDs, err := ac.GetItemIDsInContainer(ctx, userID, containerIDs[cat])
				require.NoError(t, err, clues.ToCore(err))

				assert.Equal(t, itemIDs[cat], currentIDs, "ids are equal")
			})
		}

		assert.Len(t, result, 0, "no new items should get added")
	})

	// replace restore

	suite.Run("replace collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr3 := count.New()

		restoreCfg.OnCollision = control.Replace

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr3,
			sel,
			opts,
			restoreCfg)

		deets := runAndCheckRestore(t, ctx, &ro, mb, false)
		filtEnts := []details.Entry{}

		for _, e := range deets.Entries {
			if e.Folder == nil {
				filtEnts = append(filtEnts, e)
			}
		}

		assert.Len(t, filtEnts, countItemsInRestore, "every item should have been replaced")

		checkRestoreCounts(t, ctr3, 0, countItemsInRestore, 0)

		result := map[string]string{}

		for cat, ac := range testCategories {
			suite.Run(cat.String(), func() {
				t := suite.T()

				ctx, flush := tester.NewContext(t)
				defer flush()

				m := filterCollisionKeyResults(
					t,
					ctx,
					userID,
					containerIDs[cat],
					GetItemsInContainerByCollisionKeyer[string](ac),
					collisionKeys[cat])
				maps.Copy(result, m)

				currentIDs, err := ac.GetItemIDsInContainer(ctx, userID, containerIDs[cat])
				require.NoError(t, err, clues.ToCore(err))

				assert.Equal(t, len(itemIDs[cat]), len(currentIDs), "count of ids are equal")
				for orig := range itemIDs[cat] {
					assert.NotContains(t, currentIDs, orig, "original item should not exist after replacement")
				}

				itemIDs[cat] = currentIDs
			})
		}

		assert.Len(t, result, 0, "all items should have been replaced")
	})

	// copy restore

	suite.Run("copy collisions", func() {
		t := suite.T()

		ctx, flush := tester.NewContext(t)
		defer flush()

		mb := evmock.NewBus()
		ctr4 := count.New()

		restoreCfg.OnCollision = control.Copy

		ro, _ := prepNewTestRestoreOp(
			t,
			ctx,
			bod.st,
			bo.Results.BackupID,
			mb,
			ctr4,
			sel,
			opts,
			restoreCfg)

		deets := runAndCheckRestore(t, ctx, &ro, mb, false)
		filtEnts := []details.Entry{}

		for _, e := range deets.Entries {
			if e.Folder == nil {
				filtEnts = append(filtEnts, e)
			}
		}

		assert.Len(t, filtEnts, countItemsInRestore, "every item should have been copied")

		checkRestoreCounts(t, ctr4, 0, 0, countItemsInRestore)

		result := map[string]string{}

		for cat, ac := range testCategories {
			suite.Run(cat.String(), func() {
				t := suite.T()

				ctx, flush := tester.NewContext(t)
				defer flush()

				m := filterCollisionKeyResults(
					t,
					ctx,
					userID,
					containerIDs[cat],
					GetItemsInContainerByCollisionKeyer[string](ac),
					collisionKeys[cat])
				maps.Copy(result, m)

				currentIDs, err := ac.GetItemIDsInContainer(ctx, userID, containerIDs[cat])
				require.NoError(t, err, clues.ToCore(err))

				assert.Equal(t, 2*len(itemIDs[cat]), len(currentIDs), "count of ids should be double from before")
				assert.Subset(t, maps.Keys(currentIDs), maps.Keys(itemIDs[cat]), "original item should exist after copy")
			})
		}

		// TODO: we have the option of modifying copy creations in exchange
		// so that the results don't collide.  But we haven't made that
		// decision yet.
		assert.Len(t, result, 0, "no items should have been added as copies")
	})
}

func (suite *ExchangeRestoreNightlyIntgSuite) TestRestore_Run_exchangeAlternateProtectedResource() {
	t := suite.T()

	ctx, flush := tester.NewContext(t)
	defer flush()

	// a backup is required to run restores

	baseSel := selectors.NewExchangeBackup([]string{suite.its.user.ID})
	baseSel.Include(
		// events cannot be run, for the same reason as incremental backups: the user needs
		// to have their account recycled.
		// base_sel.EventCalendars([]string{api.DefaultCalendar}, selectors.PrefixMatch()),
		baseSel.ContactFolders([]string{api.DefaultContacts}, selectors.PrefixMatch()),
		baseSel.MailFolders([]string{api.MailInbox}, selectors.PrefixMatch()))

	baseSel.DiscreteOwner = suite.its.user.ID

	var (
		mb   = evmock.NewBus()
		opts = control.DefaultOptions()
	)

	bo, bod := prepNewTestBackupOp(t, ctx, mb, baseSel.Selector, opts, version.Backup)
	defer bod.close(t, ctx)

	runAndCheckBackup(t, ctx, &bo, mb, false)

	rsel, err := baseSel.ToExchangeRestore()
	require.NoError(t, err, clues.ToCore(err))

	var (
		restoreCfg      = ctrlTD.DefaultRestoreConfig("exchange_restore_to_user")
		sel             = rsel.Selector
		userID          = suite.its.user.ID
		secondaryUserID = suite.its.secondaryUser.ID
		uid             = userID
		acCont          = suite.its.ac.Contacts()
		acMail          = suite.its.ac.Mail()
		// acEvts   = suite.its.ac.Events()
		firstCtr = count.New()
	)

	restoreCfg.OnCollision = control.Copy
	mb = evmock.NewBus()

	// first restore to the current user

	ro1, _ := prepNewTestRestoreOp(
		t,
		ctx,
		bod.st,
		bo.Results.BackupID,
		mb,
		firstCtr,
		sel,
		opts,
		restoreCfg)

	runAndCheckRestore(t, ctx, &ro1, mb, false)

	// get all files in folder, use these as the base
	// set of files to compare against.

	var (
		userItemIDs       = map[path.CategoryType]map[string]struct{}{}
		userCollisionKeys = map[path.CategoryType]map[string]string{}
	)

	// --- contacts
	cat := path.ContactsCategory
	userItemIDs[cat], userCollisionKeys[cat], _ = getCollKeysAndItemIDs(
		t,
		ctx,
		acCont,
		uid,
		"",
		restoreCfg.Location)

	// --- events
	// cat = path.EventsCategory
	// userItemIDs[cat], userCollisionKeys[cat], _ = getCollKeysAndItemIDs(
	// t,
	// ctx,
	// acEvts,
	// uid,
	// "",
	// restoreCfg.Location)

	// --- mail
	cat = path.EmailCategory
	userItemIDs[cat], userCollisionKeys[cat], _ = getCollKeysAndItemIDs(
		t,
		ctx,
		acMail,
		uid,
		"",
		restoreCfg.Location,
		api.MailInbox)

	// then restore to the secondary user

	uid = secondaryUserID
	mb = evmock.NewBus()
	secondCtr := count.New()
	restoreCfg.ProtectedResource = uid

	ro2, _ := prepNewTestRestoreOp(
		t,
		ctx,
		bod.st,
		bo.Results.BackupID,
		mb,
		secondCtr,
		sel,
		opts,
		restoreCfg)

	runAndCheckRestore(t, ctx, &ro2, mb, false)

	var (
		secondaryItemIDs       = map[path.CategoryType]map[string]struct{}{}
		secondaryCollisionKeys = map[path.CategoryType]map[string]string{}
	)

	// --- contacts
	cat = path.ContactsCategory
	secondaryItemIDs[cat], secondaryCollisionKeys[cat], _ = getCollKeysAndItemIDs(
		t,
		ctx,
		acCont,
		uid,
		"",
		restoreCfg.Location)

	// --- events
	// cat = path.EventsCategory
	// secondaryItemIDs[cat], secondaryCollisionKeys[cat], _ = getCollKeysAndItemIDs(
	// t,
	// ctx,
	// acEvts,
	// uid,
	// "",
	// restoreCfg.Location)

	// --- mail
	cat = path.EmailCategory
	secondaryItemIDs[cat], secondaryCollisionKeys[cat], _ = getCollKeysAndItemIDs(
		t,
		ctx,
		acMail,
		uid,
		"",
		restoreCfg.Location,
		api.MailInbox)

	// compare restore results
	for _, cat := range []path.CategoryType{path.ContactsCategory, path.EmailCategory, path.EventsCategory} {
		assert.Equal(t, len(userItemIDs[cat]), len(secondaryItemIDs[cat]))
		assert.ElementsMatch(t, maps.Keys(userCollisionKeys[cat]), maps.Keys(secondaryCollisionKeys[cat]))
	}
}

func getCollKeysAndItemIDs(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	cip clientItemPager,
	userID, parentContainerID string,
	containerNames ...string,
) (map[string]struct{}, map[string]string, string) {
	var (
		c   graph.Container
		err error
		cID = parentContainerID
	)

	for _, cn := range containerNames {
		c, err = cip.GetContainerByName(ctx, userID, cID, cn)
		require.NoError(t, err, clues.ToCore(err))

		cID = ptr.Val(c.GetId())
	}

	itemIDs, err := cip.GetItemIDsInContainer(ctx, userID, cID)
	require.NoError(t, err, clues.ToCore(err))

	collisionKeys, err := cip.GetItemsInContainerByCollisionKey(ctx, userID, cID)
	require.NoError(t, err, clues.ToCore(err))

	return itemIDs, collisionKeys, cID
}
