package connector

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/account"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

func mustToDataLayerPath(
	t *testing.T,
	service path.ServiceType,
	tenant, resourceOwner string,
	category path.CategoryType,
	elements []string,
	isItem bool,
) path.Path {
	var (
		err error
		res path.Path
	)

	pb := path.Builder{}.Append(elements...)

	switch service {
	case path.ExchangeService:
		res, err = pb.ToDataLayerExchangePathForCategory(tenant, resourceOwner, category, isItem)
	case path.OneDriveService:
		require.Equal(t, path.FilesCategory, category)

		res, err = pb.ToDataLayerOneDrivePath(tenant, resourceOwner, isItem)
	case path.SharePointService:
		res, err = pb.ToDataLayerSharePointPath(tenant, resourceOwner, category, isItem)

	default:
		err = errors.Errorf("bad service type %s", service.String())
	}

	require.NoError(t, err)

	return res
}

func emptyOrEqual[T any](expected *T, got *T) bool {
	if expected == nil || got == nil {
		// Creates either the zero value or gets the value pointed to.
		return reflect.DeepEqual(reflect.ValueOf(expected).Elem(), reflect.ValueOf(got).Elem())
	}

	return reflect.DeepEqual(*expected, *got)
}

func testEmptyOrEqual[T any](t *testing.T, expected *T, got *T, msg string) {
	t.Helper()

	if expected == nil || got == nil {
		// Creates either the zero value or gets the value pointed to.
		assert.Equal(t, reflect.ValueOf(expected).Elem(), reflect.ValueOf(got).Elem())
		return
	}

	assert.Equal(t, *expected, *got, msg)
}

func testElementsMatch[T any](
	t *testing.T,
	expected []T,
	got []T,
	equalityCheck func(expectedItem, gotItem T) bool,
) {
	t.Helper()

	pending := make([]*T, len(expected))
	for i := 0; i < len(expected); i++ {
		pending[i] = &expected[i]
	}

	unexpected := []T{}

	for i := 0; i < len(got); i++ {
		found := false

		for j, maybe := range pending {
			if maybe == nil {
				// Already matched with something in got.
				continue
			}

			// Item matched, break out of inner loop and move to next item in got.
			if equalityCheck(*maybe, got[i]) {
				pending[j] = nil
				found = true

				break
			}
		}

		if !found {
			unexpected = append(unexpected, got[i])
		}
	}

	// Print differences.
	missing := []T{}

	for _, p := range pending {
		if p == nil {
			continue
		}

		missing = append(missing, *p)
	}

	if len(unexpected) == 0 && len(missing) == 0 {
		return
	}

	assert.Failf(
		t,
		"contain different elements",
		"missing items: (%T)%v\nunexpected items: (%T)%v\n",
		expected,
		missing,
		got,
		unexpected,
	)
}

type configInfo struct {
	acct           account.Account
	opts           control.Options
	resource       resource
	service        path.ServiceType
	tenant         string
	resourceOwners []string
	dest           control.RestoreDestination
}

type itemInfo struct {
	// lookupKey is a string that can be used to find this data from a set of
	// other data in the same collection. This key should be something that will
	// be the same before and after restoring the item in M365 and may not be
	// the M365 ID. When restoring items out of place, the item is assigned a
	// new ID making it unsuitable for a lookup key.
	lookupKey string
	name      string
	data      []byte
}

type colInfo struct {
	// Elements (in order) for the path representing this collection. Should
	// only contain elements after the prefix that corso uses for the path. For
	// example, a collection for the Inbox folder in exchange mail would just be
	// "Inbox".
	pathElements []string
	category     path.CategoryType
	items        []itemInfo
	// auxItems are items that can be retrieved with Fetch but won't be returned
	// by Items(). These files do not directly participate in comparisosn at the
	// end of a test.
	auxItems []itemInfo
}

type restoreBackupInfo struct {
	name        string
	service     path.ServiceType
	collections []colInfo
	resource    resource
}

type restoreBackupInfoMultiVersion struct {
	service             path.ServiceType
	collectionsLatest   []colInfo
	collectionsPrevious []colInfo
	resource            resource
	backupVersion       int
}

func attachmentEqual(
	expected models.Attachmentable,
	got models.Attachmentable,
) bool {
	// This is super hacky, but seems like it would be good to have a comparison
	// of the actual content. I think the only other way to really get it is to
	// serialize both structs to JSON and pull it from there or something though.
	expectedData := reflect.Indirect(reflect.ValueOf(expected)).FieldByName("contentBytes").Bytes()
	gotData := reflect.Indirect(reflect.ValueOf(got)).FieldByName("contentBytes").Bytes()

	if !reflect.DeepEqual(expectedData, gotData) {
		return false
	}

	if !emptyOrEqual(expected.GetContentType(), got.GetContentType()) {
		return false
	}

	// Skip Id as it's tied to this specific instance of the item.

	if !emptyOrEqual(expected.GetIsInline(), got.GetIsInline()) {
		return false
	}

	// Skip LastModifiedDateTime as it's tied to this specific instance of the item.

	if !emptyOrEqual(expected.GetName(), got.GetName()) {
		return false
	}

	// Skip Size as the server clobbers whatever value we give it. It's unknown
	// how they populate size though as it's not just the length of the byte
	// array backing the content.

	return true
}

func recipientEqual(
	expected models.Recipientable,
	got models.Recipientable,
) bool {
	// Don't compare names as M365 will override the name if the address is known.
	return emptyOrEqual(
		expected.GetEmailAddress().GetAddress(),
		got.GetEmailAddress().GetAddress(),
	)
}

func checkMessage(
	t *testing.T,
	expected models.Messageable,
	got models.Messageable,
) {
	testElementsMatch(t, expected.GetAttachments(), got.GetAttachments(), attachmentEqual)

	assert.Equal(t, expected.GetBccRecipients(), got.GetBccRecipients(), "BccRecipients")

	testEmptyOrEqual(t, expected.GetBody().GetContentType(), got.GetBody().GetContentType(), "Body.ContentType")

	// Skip Body.Content as there may be display formatting that changes.

	// Skip BodyPreview as it is auto-generated on the server side and isn't
	// always just the first 255 characters if the message is HTML and has
	// multiple paragraphs.

	assert.Equal(t, expected.GetCategories(), got.GetCategories(), "Categories")

	assert.Equal(t, expected.GetCcRecipients(), got.GetCcRecipients(), "CcRecipients")

	// Skip ChangeKey as it's tied to this specific instance of the item.

	// Skip ConversationId as it's tied to this specific instance of the item.

	// Skip ConversationIndex as it's tied to this specific instance of the item.

	// Skip CreatedDateTime as it's tied to this specific instance of the item.

	checkFlags(t, expected.GetFlag(), got.GetFlag())

	checkRecipentables(t, expected.GetFrom(), got.GetFrom())
	testEmptyOrEqual(t, expected.GetHasAttachments(), got.GetHasAttachments(), "HasAttachments")

	// Skip Id as it's tied to this specific instance of the item.

	testEmptyOrEqual(t, expected.GetImportance(), got.GetImportance(), "Importance")

	testEmptyOrEqual(t, expected.GetInferenceClassification(), got.GetInferenceClassification(), "InferenceClassification")

	assert.Equal(t, expected.GetInternetMessageHeaders(), got.GetInternetMessageHeaders(), "InternetMessageHeaders")

	testEmptyOrEqual(t, expected.GetInternetMessageId(), got.GetInternetMessageId(), "InternetMessageId")

	testEmptyOrEqual(
		t,
		expected.GetIsDeliveryReceiptRequested(),
		got.GetIsDeliveryReceiptRequested(),
		"IsDeliverReceiptRequested",
	)

	testEmptyOrEqual(t, expected.GetIsDraft(), got.GetIsDraft(), "IsDraft")

	testEmptyOrEqual(t, expected.GetIsRead(), got.GetIsRead(), "IsRead")

	testEmptyOrEqual(t, expected.GetIsReadReceiptRequested(), got.GetIsReadReceiptRequested(), "IsReadReceiptRequested")

	// Skip LastModifiedDateTime as it's tied to this specific instance of the item.

	// Skip ParentFolderId as we restore to a different folder by default.

	testEmptyOrEqual(t, expected.GetReceivedDateTime(), got.GetReceivedDateTime(), "ReceivedDateTime")

	assert.Equal(t, expected.GetReplyTo(), got.GetReplyTo(), "ReplyTo")

	checkRecipentables(t, expected.GetSender(), got.GetSender())

	testEmptyOrEqual(t, expected.GetSentDateTime(), got.GetSentDateTime(), "SentDateTime")

	testEmptyOrEqual(t, expected.GetSubject(), got.GetSubject(), "Subject")

	testElementsMatch(t, expected.GetToRecipients(), got.GetToRecipients(), recipientEqual)

	// Skip WebLink as it's tied to this specific instance of the item.

	assert.Equal(t, expected.GetUniqueBody(), got.GetUniqueBody(), "UniqueBody")
}

// checkFlags is a helper function to check equality of models.FollowupFlabables
// OdataTypes are omitted as these do change between msgraph-sdk-go versions
func checkFlags(
	t *testing.T,
	expected, got models.FollowupFlagable,
) {
	assert.Equal(t, expected.GetCompletedDateTime(), got.GetCompletedDateTime())
	assert.Equal(t, expected.GetDueDateTime(), got.GetDueDateTime())
	assert.Equal(t, expected.GetFlagStatus(), got.GetFlagStatus())
	assert.Equal(t, expected.GetStartDateTime(), got.GetStartDateTime())
	assert.Equal(t, expected.GetAdditionalData(), got.GetAdditionalData())
}

// checkRecipentables is a helper function to check equality between
// models.Recipientables. OdataTypes omitted.
func checkRecipentables(
	t *testing.T,
	expected, got models.Recipientable,
) {
	checkEmailAddressables(t, expected.GetEmailAddress(), got.GetEmailAddress())
	assert.Equal(t, expected.GetAdditionalData(), got.GetAdditionalData())
}

// checkEmailAddressables inspects EmailAddressables for equality
func checkEmailAddressables(
	t *testing.T,
	expected, got models.EmailAddressable,
) {
	assert.Equal(t, expected.GetAdditionalData(), got.GetAdditionalData())
	assert.Equal(t, *expected.GetAddress(), *got.GetAddress())
	assert.Equal(t, expected.GetName(), got.GetName())
}

func checkContact(
	t *testing.T,
	expected models.Contactable,
	got models.Contactable,
) {
	testEmptyOrEqual(t, expected.GetAssistantName(), got.GetAssistantName(), "AssistantName")

	testEmptyOrEqual(t, expected.GetBirthday(), got.GetBirthday(), "Birthday")
	// Not present in msgraph-beta-sdk/models
	// assert.Equal(t, expected.GetBusinessAddress(), got.GetBusinessAddress())
	// Not present in msgraph-beta-sdk/models
	// testEmptyOrEqual(t, expected.GetBusinessHomePage(), got.GetBusinessHomePage(), "BusinessHomePage")
	// Not present in msgraph-beta-sdk/models
	// assert.Equal(t, expected.GetBusinessPhones(), got.GetBusinessPhones())

	assert.Equal(t, expected.GetCategories(), got.GetCategories())

	// Skip ChangeKey as it's tied to this specific instance of the item.

	assert.Equal(t, expected.GetChildren(), got.GetChildren())

	testEmptyOrEqual(t, expected.GetCompanyName(), got.GetCompanyName(), "CompanyName")

	// Skip CreatedDateTime as it's tied to this specific instance of the item.

	testEmptyOrEqual(t, expected.GetDepartment(), got.GetDepartment(), "Department")

	testEmptyOrEqual(t, expected.GetDisplayName(), got.GetDisplayName(), "DisplayName")

	assert.Equal(t, expected.GetEmailAddresses(), got.GetEmailAddresses())

	testEmptyOrEqual(t, expected.GetFileAs(), got.GetFileAs(), "FileAs")

	testEmptyOrEqual(t, expected.GetGeneration(), got.GetGeneration(), "Generation")

	testEmptyOrEqual(t, expected.GetGivenName(), got.GetGivenName(), "GivenName")

	// Not present in msgraph-beta-sdk/models
	// assert.Equal(t, expected.GetHomeAddress(), got.GetHomeAddress())
	// Not present in msgraph-beta-sdk/models
	// assert.Equal(t, expected.GetHomePhones(), got.GetHomePhones())

	// Skip CreatedDateTime as it's tied to this specific instance of the item.

	assert.Equal(t, expected.GetImAddresses(), got.GetImAddresses())

	testEmptyOrEqual(t, expected.GetInitials(), got.GetInitials(), "Initials")

	testEmptyOrEqual(t, expected.GetJobTitle(), got.GetJobTitle(), "JobTitle")

	// Skip CreatedDateTime as it's tied to this specific instance of the item.

	testEmptyOrEqual(t, expected.GetManager(), got.GetManager(), "Manager")

	testEmptyOrEqual(t, expected.GetMiddleName(), got.GetMiddleName(), "MiddleName")

	// Not present in msgraph-beta-sdk/models
	// testEmptyOrEqual(t, expected.GetMobilePhone(), got.GetMobilePhone(), "MobilePhone")

	testEmptyOrEqual(t, expected.GetNickName(), got.GetNickName(), "NickName")

	testEmptyOrEqual(t, expected.GetOfficeLocation(), got.GetOfficeLocation(), "OfficeLocation")
	// Not present in msgraph-beta-sdk/models
	// assert.Equal(t, expected.GetOtherAddress(), got.GetOtherAddress())

	// Skip ParentFolderId as it's tied to this specific instance of the item.

	testEmptyOrEqual(t, expected.GetPersonalNotes(), got.GetPersonalNotes(), "PersonalNotes")

	assert.Equal(t, expected.GetPhoto(), got.GetPhoto())

	testEmptyOrEqual(t, expected.GetProfession(), got.GetProfession(), "Profession")

	testEmptyOrEqual(t, expected.GetSpouseName(), got.GetSpouseName(), "SpouseName")

	testEmptyOrEqual(t, expected.GetSurname(), got.GetSurname(), "Surname")

	testEmptyOrEqual(t, expected.GetTitle(), got.GetTitle(), "Title")

	testEmptyOrEqual(t, expected.GetYomiCompanyName(), got.GetYomiCompanyName(), "YomiCompanyName")

	testEmptyOrEqual(t, expected.GetYomiGivenName(), got.GetYomiGivenName(), "YomiGivenName")

	testEmptyOrEqual(t, expected.GetYomiSurname(), got.GetYomiSurname(), "YomiSurname")
}

func locationEqual(expected, got models.Locationable) bool {
	if !reflect.DeepEqual(expected.GetAddress(), got.GetAddress()) {
		return false
	}

	if !reflect.DeepEqual(expected.GetCoordinates(), got.GetCoordinates()) {
		return false
	}

	if !emptyOrEqual(expected.GetDisplayName(), got.GetDisplayName()) {
		return false
	}

	if !emptyOrEqual(expected.GetLocationEmailAddress(), got.GetLocationEmailAddress()) {
		return false
	}

	if !emptyOrEqual(expected.GetLocationType(), got.GetLocationType()) {
		return false
	}

	// Skip checking UniqueId as it's marked as for internal use only.

	// Skip checking UniqueIdType as it's marked as for internal use only.

	if !emptyOrEqual(expected.GetLocationUri(), got.GetLocationUri()) {
		return false
	}

	return true
}

func checkEvent(
	t *testing.T,
	expected models.Eventable,
	got models.Eventable,
) {
	testEmptyOrEqual(t, expected.GetAllowNewTimeProposals(), got.GetAllowNewTimeProposals(), "AllowNewTimeProposals")

	assert.Equal(t, expected.GetAttachments(), got.GetAttachments(), "Attachments")

	assert.Equal(t, expected.GetAttendees(), got.GetAttendees(), "Attendees")

	testEmptyOrEqual(t, expected.GetBody().GetContentType(), got.GetBody().GetContentType(), "Body.ContentType")

	// Skip checking Body.Content for now as M365 may have different formatting.

	// Skip checking BodyPreview for now as M365 may have different formatting.

	assert.Equal(t, expected.GetCalendar(), got.GetCalendar(), "Calendar")

	assert.Equal(t, expected.GetCategories(), got.GetCategories(), "Categories")

	// Skip ChangeKey as it's tied to this specific instance of the item.

	// Skip CreatedDateTime as it's tied to this specific instance of the item.

	assert.Equal(t, expected.GetEnd(), got.GetEnd(), "End")

	testEmptyOrEqual(t, expected.GetHasAttachments(), got.GetHasAttachments(), "HasAttachments")

	testEmptyOrEqual(t, expected.GetHideAttendees(), got.GetHideAttendees(), "HideAttendees")

	// TODO(ashmrtn): Uncomment when we figure out how to connect to the original
	// event.
	// testEmptyOrEqual(t, expected.GetICalUId(), got.GetICalUId(), "ICalUId")

	// Skip Id as it's tied to this specific instance of the item.

	testEmptyOrEqual(t, expected.GetImportance(), got.GetImportance(), "Importance")

	assert.Equal(t, expected.GetInstances(), got.GetInstances(), "Instances")

	testEmptyOrEqual(t, expected.GetIsAllDay(), got.GetIsAllDay(), "IsAllDay")

	testEmptyOrEqual(t, expected.GetIsCancelled(), got.GetIsCancelled(), "IsCancelled")

	testEmptyOrEqual(t, expected.GetIsDraft(), got.GetIsDraft(), "IsDraft")

	testEmptyOrEqual(t, expected.GetIsOnlineMeeting(), got.GetIsOnlineMeeting(), "IsOnlineMeeting")

	// TODO(ashmrtn): Uncomment when we figure out how to delegate event creation
	// to another user.
	// testEmptyOrEqual(t, expected.GetIsOrganizer(), got.GetIsOrganizer(), "IsOrganizer")

	testEmptyOrEqual(t, expected.GetIsReminderOn(), got.GetIsReminderOn(), "IsReminderOn")

	// Skip LastModifiedDateTime as it's tied to this specific instance of the item.

	// Cheating a little here in the name of code-reuse. model.Location needs
	// custom compare logic because it has fields marked as "internal use only"
	// that seem to change.
	testElementsMatch(
		t,
		[]models.Locationable{expected.GetLocation()},
		[]models.Locationable{got.GetLocation()},
		locationEqual,
	)

	testElementsMatch(t, expected.GetLocations(), got.GetLocations(), locationEqual)

	assert.Equal(t, expected.GetOnlineMeeting(), got.GetOnlineMeeting(), "OnlineMeeting")

	testEmptyOrEqual(t, expected.GetOnlineMeetingProvider(), got.GetOnlineMeetingProvider(), "OnlineMeetingProvider")

	testEmptyOrEqual(t, expected.GetOnlineMeetingUrl(), got.GetOnlineMeetingUrl(), "OnlineMeetingUrl")

	// TODO(ashmrtn): Uncomment when we figure out how to delegate event creation
	// to another user.
	// assert.Equal(t, expected.GetOrganizer(), got.GetOrganizer(), "Organizer")

	testEmptyOrEqual(t, expected.GetOriginalEndTimeZone(), got.GetOriginalEndTimeZone(), "OriginalEndTimeZone")

	testEmptyOrEqual(t, expected.GetOriginalStart(), got.GetOriginalStart(), "OriginalStart")

	testEmptyOrEqual(t, expected.GetOriginalStartTimeZone(), got.GetOriginalStartTimeZone(), "OriginalStartTimeZone")

	assert.Equal(t, expected.GetRecurrence(), got.GetRecurrence(), "Recurrence")

	testEmptyOrEqual(
		t,
		expected.GetReminderMinutesBeforeStart(),
		got.GetReminderMinutesBeforeStart(),
		"ReminderMinutesBeforeStart",
	)

	testEmptyOrEqual(t, expected.GetResponseRequested(), got.GetResponseRequested(), "ResponseRequested")

	// TODO(ashmrtn): Uncomment when we figure out how to connect to the original
	// event.
	// assert.Equal(t, expected.GetResponseStatus(), got.GetResponseStatus(), "ResponseStatus")

	testEmptyOrEqual(t, expected.GetSensitivity(), got.GetSensitivity(), "Sensitivity")

	testEmptyOrEqual(t, expected.GetSeriesMasterId(), got.GetSeriesMasterId(), "SeriesMasterId")

	testEmptyOrEqual(t, expected.GetShowAs(), got.GetShowAs(), "ShowAs")

	assert.Equal(t, expected.GetStart(), got.GetStart(), "Start")

	testEmptyOrEqual(t, expected.GetSubject(), got.GetSubject(), "Subject")

	testEmptyOrEqual(t, expected.GetTransactionId(), got.GetTransactionId(), "TransactionId")

	// Skip LastModifiedDateTime as it's tied to this specific instance of the item.

	testEmptyOrEqual(t, expected.GetType(), got.GetType(), "Type")
}

func compareExchangeEmail(
	t *testing.T,
	expected map[string][]byte,
	item data.Stream,
) {
	itemData, err := io.ReadAll(item.ToReader())
	if !assert.NoError(t, err, "reading collection item: %s", item.UUID()) {
		return
	}

	itemMessage, err := support.CreateMessageFromBytes(itemData)
	if !assert.NoError(t, err, "deserializing backed up message") {
		return
	}

	expectedBytes, ok := expected[*itemMessage.GetSubject()]
	if !assert.True(t, ok, "unexpected item with Subject %q", *itemMessage.GetSubject()) {
		return
	}

	expectedMessage, err := support.CreateMessageFromBytes(expectedBytes)
	assert.NoError(t, err, "deserializing source message")

	checkMessage(t, expectedMessage, itemMessage)
}

func compareExchangeContact(
	t *testing.T,
	expected map[string][]byte,
	item data.Stream,
) {
	itemData, err := io.ReadAll(item.ToReader())
	if !assert.NoError(t, err, "reading collection item: %s", item.UUID()) {
		return
	}

	itemContact, err := support.CreateContactFromBytes(itemData)
	if !assert.NoError(t, err, "deserializing backed up contact") {
		return
	}

	expectedBytes, ok := expected[*itemContact.GetMiddleName()]
	if !assert.True(t, ok, "unexpected item with middle name %q", *itemContact.GetMiddleName()) {
		return
	}

	expectedContact, err := support.CreateContactFromBytes(expectedBytes)
	assert.NoError(t, err, "deserializing source contact")

	checkContact(t, expectedContact, itemContact)
}

func compareExchangeEvent(
	t *testing.T,
	expected map[string][]byte,
	item data.Stream,
) {
	itemData, err := io.ReadAll(item.ToReader())
	if !assert.NoError(t, err, "reading collection item: %s", item.UUID()) {
		return
	}

	itemEvent, err := support.CreateEventFromBytes(itemData)
	if !assert.NoError(t, err, "deserializing backed up contact") {
		return
	}

	expectedBytes, ok := expected[*itemEvent.GetSubject()]
	if !assert.True(t, ok, "unexpected item with subject %q", *itemEvent.GetSubject()) {
		return
	}

	expectedEvent, err := support.CreateEventFromBytes(expectedBytes)
	assert.NoError(t, err, "deserializing source contact")

	checkEvent(t, expectedEvent, itemEvent)
}

func permissionEqual(expected onedrive.UserPermission, got onedrive.UserPermission) bool {
	if !strings.EqualFold(expected.Email, got.Email) {
		return false
	}

	if (expected.Expiration == nil && got.Expiration != nil) ||
		(expected.Expiration != nil && got.Expiration == nil) {
		return false
	}

	if expected.Expiration != nil &&
		got.Expiration != nil &&
		!expected.Expiration.Equal(*got.Expiration) {
		return false
	}

	if len(expected.Roles) != len(got.Roles) {
		return false
	}

	for _, r := range got.Roles {
		if !slices.Contains(expected.Roles, r) {
			return false
		}
	}

	return true
}

func compareOneDriveItem(
	t *testing.T,
	expected map[string][]byte,
	item data.Stream,
	dest control.RestoreDestination,
	restorePermissions bool,
) bool {
	// Skip OneDrive permissions in the folder that used to be the root. We don't
	// have a good way to materialize these in the test right now.
	if item.UUID() == dest.ContainerName+onedrive.DirMetaFileSuffix {
		return false
	}

	buf, err := io.ReadAll(item.ToReader())
	if !assert.NoError(t, err) {
		return true
	}

	name := item.UUID()

	if strings.HasSuffix(name, onedrive.MetaFileSuffix) ||
		strings.HasSuffix(name, onedrive.DirMetaFileSuffix) {
		var (
			itemMeta     onedrive.Metadata
			expectedMeta onedrive.Metadata
		)

		err = json.Unmarshal(buf, &itemMeta)
		if !assert.NoErrorf(t, err, "unmarshalling retrieved metadata for file %s", name) {
			return true
		}

		expectedData := expected[name]
		if !assert.NotNil(
			t,
			expectedData,
			"unexpected metadata file with name %s",
			name,
		) {
			return true
		}

		err = json.Unmarshal(expectedData, &expectedMeta)
		if !assert.NoError(t, err, "unmarshalling expected metadata") {
			return true
		}

		// Only compare file names if we're using a version that expects them to be
		// set.
		if len(expectedMeta.FileName) > 0 {
			assert.Equal(t, expectedMeta.FileName, itemMeta.FileName)
		}

		if !restorePermissions {
			assert.Equal(t, 0, len(itemMeta.Permissions))
			return true
		}

		testElementsMatch(
			t,
			expectedMeta.Permissions,
			itemMeta.Permissions,
			permissionEqual,
		)

		return true
	}

	var fileData testOneDriveData

	err = json.Unmarshal(buf, &fileData)
	if !assert.NoErrorf(t, err, "unmarshalling file data for file %s", name) {
		return true
	}

	expectedData := expected[fileData.FileName]
	if !assert.NotNil(t, expectedData, "unexpected file with name %s", name) {
		return true
	}

	// OneDrive data items are just byte buffers of the data. Nothing special to
	// interpret. May need to do chunked comparisons in the future if we test
	// large item equality.
	// Compare against the version with the file name embedded because that's what
	// the auto-generated expected data has.
	assert.Equal(t, expectedData, buf)

	return true
}

// compareItem compares the data returned by backup with the expected data.
// Returns true if a comparison was done else false. Bool return is mostly used
// to exclude OneDrive permissions for the root right now.
func compareItem(
	t *testing.T,
	expected map[string][]byte,
	service path.ServiceType,
	category path.CategoryType,
	item data.Stream,
	dest control.RestoreDestination,
	restorePermissions bool,
) bool {
	if mt, ok := item.(data.StreamModTime); ok {
		assert.NotZero(t, mt.ModTime())
	}

	switch service {
	case path.ExchangeService:
		switch category {
		case path.EmailCategory:
			compareExchangeEmail(t, expected, item)
		case path.ContactsCategory:
			compareExchangeContact(t, expected, item)
		case path.EventsCategory:
			compareExchangeEvent(t, expected, item)
		default:
			assert.FailNowf(t, "unexpected Exchange category: %s", category.String())
		}

	case path.OneDriveService:
		return compareOneDriveItem(t, expected, item, dest, restorePermissions)

	default:
		assert.FailNowf(t, "unexpected service: %s", service.String())
	}

	return true
}

func checkHasCollections(
	t *testing.T,
	expected map[string]map[string][]byte,
	got []data.BackupCollection,
) {
	t.Helper()

	expectedNames := make([]string, 0, len(expected))
	gotNames := make([]string, 0, len(got))

	for e := range expected {
		expectedNames = append(expectedNames, e)
	}

	for _, g := range got {
		gotNames = append(gotNames, g.FullPath().String())
	}

	assert.ElementsMatch(t, expectedNames, gotNames, "returned collections")
}

//revive:disable:context-as-argument
func checkCollections(
	t *testing.T,
	ctx context.Context,
	expectedItems int,
	expected map[string]map[string][]byte,
	got []data.BackupCollection,
	dest control.RestoreDestination,
	restorePermissions bool,
) int {
	//revive:enable:context-as-argument
	collectionsWithItems := []data.BackupCollection{}

	skipped := 0
	gotItems := 0

	for _, returned := range got {
		var (
			hasItems        bool
			service         = returned.FullPath().Service()
			category        = returned.FullPath().Category()
			expectedColData = expected[returned.FullPath().String()]
		)

		// Need to iterate through all items even if we don't expect to find a match
		// because otherwise we'll deadlock waiting for GC status. Unexpected or
		// missing collection paths will be reported by checkHasCollections.
		for item := range returned.Items(ctx, fault.New(true)) {
			// Skip metadata collections as they aren't directly related to items to
			// backup. Don't add them to the item count either since the item count
			// is for actual pull items.
			// TODO(ashmrtn): Should probably eventually check some data in metadata
			// collections.
			if service == path.ExchangeMetadataService ||
				service == path.OneDriveMetadataService ||
				service == path.SharePointMetadataService {
				skipped++
				continue
			}

			hasItems = true
			gotItems++

			if expectedColData == nil {
				continue
			}

			if !compareItem(t, expectedColData, service, category, item, dest, restorePermissions) {
				gotItems--
			}
		}

		if hasItems {
			collectionsWithItems = append(collectionsWithItems, returned)
		}
	}

	assert.Equal(t, expectedItems, gotItems, "expected items")
	checkHasCollections(t, expected, collectionsWithItems)

	// Return how many metadata files were skipped so we can account for it in the
	// check on GraphConnector status.
	return skipped
}

type destAndCats struct {
	resourceOwner string
	dest          string
	cats          map[path.CategoryType]struct{}
}

func makeExchangeBackupSel(
	t *testing.T,
	dests []destAndCats,
) selectors.Selector {
	toInclude := [][]selectors.ExchangeScope{}
	resourceOwners := map[string]struct{}{}

	for _, d := range dests {
		for c := range d.cats {
			resourceOwners[d.resourceOwner] = struct{}{}

			// nil owners here, but we'll need to stitch this together
			// below after the loops are complete.
			sel := selectors.NewExchangeBackup(nil)
			builder := sel.MailFolders

			switch c {
			case path.ContactsCategory:
				builder = sel.ContactFolders
			case path.EventsCategory:
				builder = sel.EventCalendars
			case path.EmailCategory: // already set
			}

			toInclude = append(toInclude, builder(
				[]string{d.dest},
				selectors.PrefixMatch(),
			))
		}
	}

	sel := selectors.NewExchangeBackup(maps.Keys(resourceOwners))
	sel.Include(toInclude...)

	return sel.Selector
}

func makeOneDriveBackupSel(
	t *testing.T,
	dests []destAndCats,
) selectors.Selector {
	toInclude := [][]selectors.OneDriveScope{}
	resourceOwners := map[string]struct{}{}

	for _, d := range dests {
		resourceOwners[d.resourceOwner] = struct{}{}

		// nil owners here, we'll need to stitch this together
		// below after the loops are complete.
		sel := selectors.NewOneDriveBackup(nil)

		toInclude = append(toInclude, sel.Folders(
			[]string{d.dest},
			selectors.PrefixMatch(),
		))
	}

	sel := selectors.NewOneDriveBackup(maps.Keys(resourceOwners))
	sel.Include(toInclude...)

	return sel.Selector
}

// backupSelectorForExpected creates a selector that can be used to backup the
// given items in expected based on the item paths. Fails the test if items from
// multiple services are in expected.
func backupSelectorForExpected(
	t *testing.T,
	service path.ServiceType,
	dests []destAndCats,
) selectors.Selector {
	require.NotEmpty(t, dests)

	switch service {
	case path.ExchangeService:
		return makeExchangeBackupSel(t, dests)

	case path.OneDriveService:
		return makeOneDriveBackupSel(t, dests)

	default:
		assert.FailNow(t, "unknown service type %s", service.String())
	}

	// Fix compile error about no return. Should not reach here.
	return selectors.Selector{}
}

// backupOutputPathFromRestore returns a path.Path denoting the location in
// kopia the data will be placed at. The location is a data-type specific
// combination of the location the data was recently restored to and where the
// data was originally in the hierarchy.
func backupOutputPathFromRestore(
	t *testing.T,
	restoreDest control.RestoreDestination,
	inputPath path.Path,
) path.Path {
	base := []string{restoreDest.ContainerName}

	// OneDrive has leading information like the drive ID.
	if inputPath.Service() == path.OneDriveService {
		folders := inputPath.Folders()
		base = append(append([]string{}, folders[:3]...), restoreDest.ContainerName)

		if len(folders) > 3 {
			base = append(base, folders[3:]...)
		}
	}

	if inputPath.Service() == path.ExchangeService && inputPath.Category() == path.EmailCategory {
		base = append(base, inputPath.Folders()...)
	}

	return mustToDataLayerPath(
		t,
		inputPath.Service(),
		inputPath.Tenant(),
		inputPath.ResourceOwner(),
		inputPath.Category(),
		base,
		false,
	)
}

// TODO(ashmrtn): Make this an actual mock class that can be used in other
// packages.
type mockRestoreCollection struct {
	data.Collection
	auxItems map[string]data.Stream
}

func (rc mockRestoreCollection) Fetch(
	ctx context.Context,
	name string,
) (data.Stream, error) {
	res := rc.auxItems[name]
	if res == nil {
		return nil, data.ErrNotFound
	}

	return res, nil
}

func collectionsForInfo(
	t *testing.T,
	service path.ServiceType,
	tenant, user string,
	dest control.RestoreDestination,
	allInfo []colInfo,
	backupVersion int,
) (int, int, []data.RestoreCollection, map[string]map[string][]byte) {
	var (
		collections  = make([]data.RestoreCollection, 0, len(allInfo))
		expectedData = make(map[string]map[string][]byte, len(allInfo))
		totalItems   = 0
		kopiaEntries = 0
	)

	for _, info := range allInfo {
		pth := mustToDataLayerPath(
			t,
			service,
			tenant,
			user,
			info.category,
			info.pathElements,
			false)

		mc := mockconnector.NewMockExchangeCollection(pth, pth, len(info.items))
		baseDestPath := backupOutputPathFromRestore(t, dest, pth)

		baseExpected := expectedData[baseDestPath.String()]
		if baseExpected == nil {
			expectedData[baseDestPath.String()] = make(map[string][]byte, len(info.items))
			baseExpected = expectedData[baseDestPath.String()]
		}

		for i := 0; i < len(info.items); i++ {
			mc.Names[i] = info.items[i].name
			mc.Data[i] = info.items[i].data

			baseExpected[info.items[i].lookupKey] = info.items[i].data

			// We do not count metadata files against item count
			if backupVersion == 0 || service != path.OneDriveService ||
				(service == path.OneDriveService &&
					strings.HasSuffix(info.items[i].name, onedrive.DataFileSuffix)) {
				totalItems++
			}
		}

		c := mockRestoreCollection{Collection: mc, auxItems: map[string]data.Stream{}}

		for _, aux := range info.auxItems {
			c.auxItems[aux.name] = &mockconnector.MockExchangeData{
				ID:     aux.name,
				Reader: io.NopCloser(bytes.NewReader(aux.data)),
			}
		}

		collections = append(collections, c)
		kopiaEntries += len(info.items)
	}

	return totalItems, kopiaEntries, collections, expectedData
}

func getSelectorWith(
	t *testing.T,
	service path.ServiceType,
	resourceOwners []string,
	forRestore bool,
) selectors.Selector {
	switch service {
	case path.ExchangeService:
		if forRestore {
			return selectors.NewExchangeRestore(resourceOwners).Selector
		}

		return selectors.NewExchangeBackup(resourceOwners).Selector

	case path.OneDriveService:
		if forRestore {
			return selectors.NewOneDriveRestore(resourceOwners).Selector
		}

		return selectors.NewOneDriveBackup(resourceOwners).Selector

	case path.SharePointService:
		if forRestore {
			return selectors.NewSharePointRestore(resourceOwners).Selector
		}

		return selectors.NewSharePointBackup(resourceOwners).Selector

	default:
		require.FailNow(t, "unknown path service")
		return selectors.Selector{}
	}
}

func loadConnector(ctx context.Context, t *testing.T, itemClient *http.Client, r resource) *GraphConnector {
	a := tester.NewM365Account(t)

	connector, err := NewGraphConnector(ctx, itemClient, a, r, fault.New(true))
	require.NoError(t, err)

	return connector
}
