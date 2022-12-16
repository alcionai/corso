package selectors_test

import (
	"context"
	"fmt"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/selectors"
)

// ExampleNewSelector demonstrates creation and distribution of a Selector.
func Example_newSelector() {
	// Selectors should use application-specific constructors.
	// Generate a selector for backup operations.
	seb := selectors.NewExchangeBackup(nil)

	// Generate a selector for restore and 'backup details' operations.
	ser := selectors.NewExchangeRestore(nil)

	// Selectors specify the data that should be handled
	// in an operation by specifying the Scope of data.
	// Initially, the selector will ask for the resource
	// owners (users, in this example).  Only these users
	// will be involved in the backup.
	seb = selectors.NewExchangeBackup(
		[]string{"your-user-id", "foo-user-id", "bar-user-id"},
	)

	// The core selector can be passed around without slicing any
	// application-specific data.
	bSel := seb.Selector
	rSel := ser.Selector

	// And can be re-cast to the application instance again.
	seb, _ = bSel.ToExchangeBackup()
	ser, _ = rSel.ToExchangeRestore()

	// Casting the core selector to a different application will
	// result in an error.
	if _, err := bSel.ToOneDriveBackup(); err != nil {
		// this errors, because bSel is an Exchange selector.
		fmt.Println(err)
	}

	// You can inspect the selector.Service to know which application to use.
	switch bSel.Service {
	case selectors.ServiceExchange:
		//nolint
		bSel.ToExchangeBackup()
	case selectors.ServiceOneDrive:
		//nolint
		bSel.ToOneDriveBackup()
	}

	// Output: OneDrive service is not Exchange: wrong selector service type
}

// ExampleIncludeFoldersAndItems demonstrates how to select for granular data.
func Example_includeFoldersAndItems() {
	seb := selectors.NewExchangeBackup(
		[]string{"your-user-id", "foo-user-id", "bar-user-id"},
	)

	// Much of the data handled by Corso exists within an established hierarchy.
	// Resource Owner-level data (such as users) sits at the top, with Folder
	// structures and individual items below.  Higher level scopes will automatically
	// involve all descendant data in the hierarchy.

	// Users will select all Exchange data owned by the specified user.
	seb.Users([]string{"foo-user-id"})

	// Lower level Scopes are described on a per-data-type basis.  This scope will
	// select all email in the Inbox folder, for all users in the tenant.
	seb.MailFolders(selectors.Any(), []string{"Inbox"})

	// Folder-level scopes will, by default, include every folder whose name matches
	// the provided value, regardless of its position in the hierarchy.  If you want
	// to restrict the scope to a specific path, you can use the PrefixMatch option.
	// This scope selects all data in /foolder, but will skip /other/foolder.
	seb.MailFolders(
		selectors.Any(),
		[]string{"foolder"},
		selectors.PrefixMatch())

	// Individual items can be selected, too.  You don't have to use the Any()
	// selection for users and folders when specifying an item, but these ids are
	// usually unique, and have a low chance of collision.
	seb.Mails(
		selectors.Any(),
		selectors.Any(),
		[]string{"item-id-1", "item-id-2"},
	)
}

// ExampleFilters demonstrates selector filters.
func Example_filters() {
	ser := selectors.NewExchangeRestore(
		[]string{"your-user-id", "foo-user-id", "bar-user-id"},
	)

	// In addition to data ownership details (user, folder, itemID), certain operations
	// like `backup details` and restores allow items to be selected by filtering on
	// previously gathered metadata.

	// Unlike `Include()`, which will incorporate data so long as any Scope matches,
	// scopes in the `Filter()` category work as an intersection.  Data must pass
	// every filter to be selected.  The following selector will only include emails
	// received before the data, and with the given subject
	ser.Filter(
		// Note that the ReceivedBefore scope only accepts a single string instead of
		// a slice.  Since Filters act as intersections rather than unions, it wouldn't
		// make much sense to accept multiple values here.
		ser.MailReceivedBefore("2006-01-02"),
		// But you can still make a compound filter by adding each scope individually.
		ser.MailSubject("the answer to life, the universe, and everything"),
	)

	// Selectors can specify both Filter and Inclusion scopes.  Now, not only will the
	// data only include emails matching the filters above, it will only include emails
	// owned by this one user.
	ser.Include(ser.Users([]string{"foo-user-id"}))
}

var (
	//nolint
	ctxBG          = context.Background()
	exampleDetails = &details.Details{
		DetailsModel: details.DetailsModel{
			Entries: []details.DetailsEntry{
				{
					RepoRef:  "tID/exchange/your-user-id/email/example/itemID",
					ShortRef: "xyz",
					ItemInfo: details.ItemInfo{
						Exchange: &details.ExchangeInfo{
							ItemType: details.ExchangeMail,
							Subject:  "the answer to life, the universe, and everything",
						},
					},
				},
			},
		},
	}
)

// ExampleReduceDetails demonstrates how selectors are used to filter backup details.
func Example_reduceDetails() {
	ser := selectors.NewExchangeRestore(
		[]string{"your-user-id", "foo-user-id", "bar-user-id"},
	)

	// The Reduce() call is where our constructed selectors are applied to the data
	// from a previous backup record.
	filteredDetails := ser.Reduce(ctxBG, exampleDetails)

	// We haven't added any scopes to our selector yet, so none of the data is retained.
	fmt.Println("Before adding scopes:", len(filteredDetails.Entries))

	ser.Include(ser.Mails([]string{"your-user-id"}, []string{"example"}, []string{"xyz"}))
	ser.Filter(ser.MailSubject("the answer to life"))

	// Now that we've selected our data, we should find a result.
	filteredDetails = ser.Reduce(ctxBG, exampleDetails)
	fmt.Println("After adding scopes:", len(filteredDetails.Entries))

	// Output: Before adding scopes: 0
	// After adding scopes: 1
}

// ExampleScopeMatching demonstrates how to compare data against an individual scope.
func Example_scopeMatching() {
	// Just like sets of backup data can be filtered down using Reduce(), we can check
	// if an individual bit of data matches our scopes, too.
	scope := selectors.
		NewExchangeBackup(
			[]string{"your-user-id", "foo-user-id", "bar-user-id"},
		).
		Mails(
			[]string{"id-1"},
			[]string{"Inbox"},
			selectors.Any(),
		)[0]

	// To compare data against a scope, you need to specify the category of data,
	// and input the value to check.
	result := scope.Matches(selectors.ExchangeMailFolder, "inbox")
	fmt.Println("Matches the mail folder 'inbox':", result)

	// Non-matching values will return false.
	result = scope.Matches(selectors.ExchangeUser, "id-42")
	fmt.Println("Matches the user by id 'id-42':", result)

	// If you specify a category that doesn't belong to the expected
	// data type, the result is always false, even if the underlying
	// comparators match.
	result = scope.Matches(selectors.ExchangeContact, "id-1")
	fmt.Println("Matches the contact by id 'id-1':", result)

	// When in doubt, you can check the category of data in the scope
	// with the Category() method.
	cat := scope.Category()
	fmt.Println("Scope Category:", cat)

	// Output: Matches the mail folder 'inbox': true
	// Matches the user by id 'id-42': false
	// Matches the contact by id 'id-1': false
	// Scope Category: ExchangeMail
}
