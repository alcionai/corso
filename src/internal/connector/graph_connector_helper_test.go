package connector

import (
	"context"
	"encoding/json"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/alcionai/clues"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/exchange"
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/internal/tester"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

func testElementsMatch[T any](
	t *testing.T,
	expected []T,
	got []T,
	subset bool,
	equalityCheck func(expectedItem, gotItem T) bool,
) {
	t.Helper()

	pending := make([]*T, len(expected))

	for i := range expected {
		ei := expected[i]
		pending[i] = &ei
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

	if subset && len(missing) == 0 && len(unexpected) > 0 {
		return
	}

	assert.Failf(
		t,
		"elements differ",
		"expected: (%T)%+v\ngot: (%T)%+v\nmissing: %+v\nextra: %+v\n",
		expected,
		expected,
		got,
		got,
		missing,
		unexpected)
}

type restoreBackupInfo struct {
	name        string
	service     path.ServiceType
	collections []ColInfo
	resource    Resource
}

type restoreBackupInfoMultiVersion struct {
	service             path.ServiceType
	collectionsLatest   []ColInfo
	collectionsPrevious []ColInfo
	resource            Resource
	backupVersion       int
}

func attachmentEqual(
	expected models.Attachmentable,
	got models.Attachmentable,
) bool {
	expectedData, err := exchange.GetAttachmentBytes(expected)
	if err != nil {
		return false
	}

	gotData, err := exchange.GetAttachmentBytes(got)
	if err != nil {
		return false
	}

	if !reflect.DeepEqual(expectedData, gotData) {
		return false
	}

	if !reflect.DeepEqual(ptr.Val(expected.GetContentType()), ptr.Val(got.GetContentType())) {
		return false
	}

	// Skip Id as it's tied to this specific instance of the item.

	if !reflect.DeepEqual(ptr.Val(expected.GetIsInline()), ptr.Val(got.GetIsInline())) {
		return false
	}

	// Skip LastModifiedDateTime as it's tied to this specific instance of the item.

	if !reflect.DeepEqual(ptr.Val(expected.GetName()), ptr.Val(got.GetName())) {
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
	return reflect.DeepEqual(
		ptr.Val(expected.GetEmailAddress().GetAddress()),
		ptr.Val(got.GetEmailAddress().GetAddress()),
	)
}

func checkMessage(
	t *testing.T,
	expected models.Messageable,
	got models.Messageable,
) {
	testElementsMatch(t, expected.GetAttachments(), got.GetAttachments(), false, attachmentEqual)

	assert.Equal(t, expected.GetBccRecipients(), got.GetBccRecipients(), "BccRecipients")

	assert.Equal(
		t,
		ptr.Val(expected.GetBody().GetContentType()),
		ptr.Val(got.GetBody().GetContentType()),
		"Body.ContentType")

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
	assert.Equal(t, ptr.Val(expected.GetHasAttachments()), ptr.Val(got.GetHasAttachments()), "HasAttachments")

	// Skip Id as it's tied to this specific instance of the item.

	assert.Equal(t, ptr.Val(expected.GetImportance()), ptr.Val(got.GetImportance()), "Importance")

	assert.Equal(
		t,
		ptr.Val(expected.GetInferenceClassification()),
		ptr.Val(got.GetInferenceClassification()),
		"InferenceClassification")

	assert.Equal(t, expected.GetInternetMessageHeaders(), got.GetInternetMessageHeaders(), "InternetMessageHeaders")

	assert.Equal(t, ptr.Val(expected.GetInternetMessageId()), ptr.Val(got.GetInternetMessageId()), "InternetMessageId")

	assert.Equal(
		t,
		ptr.Val(expected.GetIsDeliveryReceiptRequested()),
		ptr.Val(got.GetIsDeliveryReceiptRequested()),
		"IsDeliverReceiptRequested",
	)

	assert.Equal(t, ptr.Val(expected.GetIsDraft()), ptr.Val(got.GetIsDraft()), "IsDraft")

	assert.Equal(t, ptr.Val(expected.GetIsRead()), ptr.Val(got.GetIsRead()), "IsRead")

	assert.Equal(
		t,
		ptr.Val(expected.GetIsReadReceiptRequested()),
		ptr.Val(got.GetIsReadReceiptRequested()),
		"IsReadReceiptRequested")

	// Skip LastModifiedDateTime as it's tied to this specific instance of the item.

	// Skip ParentFolderId as we restore to a different folder by default.

	assert.Equal(t, ptr.Val(expected.GetReceivedDateTime()), ptr.Val(got.GetReceivedDateTime()), "ReceivedDateTime")

	assert.Equal(t, expected.GetReplyTo(), got.GetReplyTo(), "ReplyTo")

	checkRecipentables(t, expected.GetSender(), got.GetSender())

	assert.Equal(t, ptr.Val(expected.GetSentDateTime()), ptr.Val(got.GetSentDateTime()), "SentDateTime")

	assert.Equal(t, ptr.Val(expected.GetSubject()), ptr.Val(got.GetSubject()), "Subject")

	testElementsMatch(t, expected.GetToRecipients(), got.GetToRecipients(), false, recipientEqual)

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
	colPath path.Path,
	expected models.Contactable,
	got models.Contactable,
) {
	assert.Equal(t, ptr.Val(expected.GetAssistantName()), ptr.Val(got.GetAssistantName()), "AssistantName")

	assert.Equal(t, ptr.Val(expected.GetBirthday()), ptr.Val(got.GetBirthday()), "Birthday")
	// Not present in msgraph-beta-sdk/models
	// assert.Equal(t, expected.GetBusinessAddress(), got.GetBusinessAddress())
	// Not present in msgraph-beta-sdk/models
	// assert.Equal(t, ptr.Val(expected.GetBusinessHomePage()), ptr.Val(got.GetBusinessHomePage()), "BusinessHomePage")
	// Not present in msgraph-beta-sdk/models
	// assert.Equal(t, expected.GetBusinessPhones(), got.GetBusinessPhones())

	// TODO(ashmrtn): Remove this when we properly set and handle categories in
	// addition to folders for contacts.
	folders := colPath.Folder(false)
	gotCategories := []string{}

	for _, cat := range got.GetCategories() {
		// Don't add a category for the current folder since we didn't create the
		// item with it and it throws off our comparisons.
		if cat == folders {
			continue
		}

		gotCategories = append(gotCategories, cat)
	}

	assert.ElementsMatch(t, expected.GetCategories(), gotCategories, "Categories")

	// Skip ChangeKey as it's tied to this specific instance of the item.

	assert.Equal(t, expected.GetChildren(), got.GetChildren())

	assert.Equal(t, ptr.Val(expected.GetCompanyName()), ptr.Val(got.GetCompanyName()), "CompanyName")

	// Skip CreatedDateTime as it's tied to this specific instance of the item.

	assert.Equal(t, ptr.Val(expected.GetDepartment()), ptr.Val(got.GetDepartment()), "Department")

	assert.Equal(t, ptr.Val(expected.GetDisplayName()), ptr.Val(got.GetDisplayName()), "DisplayName")

	assert.Equal(t, expected.GetEmailAddresses(), got.GetEmailAddresses())

	assert.Equal(t, ptr.Val(expected.GetFileAs()), ptr.Val(got.GetFileAs()), "FileAs")

	assert.Equal(t, ptr.Val(expected.GetGeneration()), ptr.Val(got.GetGeneration()), "Generation")

	assert.Equal(t, ptr.Val(expected.GetGivenName()), ptr.Val(got.GetGivenName()), "GivenName")

	// Not present in msgraph-beta-sdk/models
	// assert.Equal(t, expected.GetHomeAddress(), got.GetHomeAddress())
	// Not present in msgraph-beta-sdk/models
	// assert.Equal(t, expected.GetHomePhones(), got.GetHomePhones())

	// Skip CreatedDateTime as it's tied to this specific instance of the item.

	assert.Equal(t, expected.GetImAddresses(), got.GetImAddresses())

	assert.Equal(t, ptr.Val(expected.GetInitials()), ptr.Val(got.GetInitials()), "Initials")

	assert.Equal(t, ptr.Val(expected.GetJobTitle()), ptr.Val(got.GetJobTitle()), "JobTitle")

	// Skip CreatedDateTime as it's tied to this specific instance of the item.

	assert.Equal(t, ptr.Val(expected.GetManager()), ptr.Val(got.GetManager()), "Manager")

	assert.Equal(t, ptr.Val(expected.GetMiddleName()), ptr.Val(got.GetMiddleName()), "MiddleName")

	// Not present in msgraph-beta-sdk/models
	// assert.Equal(t, ptr.Val(expected.GetMobilePhone()), ptr.Val(got.GetMobilePhone()), "MobilePhone")

	assert.Equal(t, ptr.Val(expected.GetNickName()), ptr.Val(got.GetNickName()), "NickName")

	assert.Equal(t, ptr.Val(expected.GetOfficeLocation()), ptr.Val(got.GetOfficeLocation()), "OfficeLocation")
	// Not present in msgraph-beta-sdk/models
	// assert.Equal(t, expected.GetOtherAddress(), got.GetOtherAddress())

	// Skip ParentFolderId as it's tied to this specific instance of the item.

	assert.Equal(t, ptr.Val(expected.GetPersonalNotes()), ptr.Val(got.GetPersonalNotes()), "PersonalNotes")

	assert.Equal(t, expected.GetPhoto(), got.GetPhoto())

	assert.Equal(t, ptr.Val(expected.GetProfession()), ptr.Val(got.GetProfession()), "Profession")

	assert.Equal(t, ptr.Val(expected.GetSpouseName()), ptr.Val(got.GetSpouseName()), "SpouseName")

	assert.Equal(t, ptr.Val(expected.GetSurname()), ptr.Val(got.GetSurname()), "Surname")

	assert.Equal(t, ptr.Val(expected.GetTitle()), ptr.Val(got.GetTitle()), "Title")

	assert.Equal(t, ptr.Val(expected.GetYomiCompanyName()), ptr.Val(got.GetYomiCompanyName()), "YomiCompanyName")

	assert.Equal(t, ptr.Val(expected.GetYomiGivenName()), ptr.Val(got.GetYomiGivenName()), "YomiGivenName")

	assert.Equal(t, ptr.Val(expected.GetYomiSurname()), ptr.Val(got.GetYomiSurname()), "YomiSurname")
}

func locationEqual(expected, got models.Locationable) bool {
	if !reflect.DeepEqual(expected.GetAddress(), got.GetAddress()) {
		return false
	}

	if !reflect.DeepEqual(expected.GetCoordinates(), got.GetCoordinates()) {
		return false
	}

	if !reflect.DeepEqual(ptr.Val(expected.GetDisplayName()), ptr.Val(got.GetDisplayName())) {
		return false
	}

	if !reflect.DeepEqual(ptr.Val(expected.GetLocationEmailAddress()), ptr.Val(got.GetLocationEmailAddress())) {
		return false
	}

	if !reflect.DeepEqual(ptr.Val(expected.GetLocationType()), ptr.Val(got.GetLocationType())) {
		return false
	}

	// Skip checking UniqueId as it's marked as for internal use only.

	// Skip checking UniqueIdType as it's marked as for internal use only.

	if !reflect.DeepEqual(ptr.Val(expected.GetLocationUri()), ptr.Val(got.GetLocationUri())) {
		return false
	}

	return true
}

func checkEvent(
	t *testing.T,
	expected models.Eventable,
	got models.Eventable,
) {
	assert.Equal(
		t,
		ptr.Val(expected.GetAllowNewTimeProposals()),
		ptr.Val(got.GetAllowNewTimeProposals()),
		"AllowNewTimeProposals")

	assert.Equal(t, expected.GetAttachments(), got.GetAttachments(), "Attachments")

	assert.Equal(t, expected.GetAttendees(), got.GetAttendees(), "Attendees")

	assert.Equal(
		t,
		ptr.Val(expected.GetBody().GetContentType()),
		ptr.Val(got.GetBody().GetContentType()),
		"Body.ContentType")

	// Skip checking Body.Content for now as M365 may have different formatting.

	// Skip checking BodyPreview for now as M365 may have different formatting.

	assert.Equal(t, expected.GetCalendar(), got.GetCalendar(), "Calendar")

	assert.Equal(t, expected.GetCategories(), got.GetCategories(), "Categories")

	// Skip ChangeKey as it's tied to this specific instance of the item.

	// Skip CreatedDateTime as it's tied to this specific instance of the item.

	assert.Equal(t, expected.GetEnd(), got.GetEnd(), "End")

	assert.Equal(t, ptr.Val(expected.GetHasAttachments()), ptr.Val(got.GetHasAttachments()), "HasAttachments")

	assert.Equal(t, ptr.Val(expected.GetHideAttendees()), ptr.Val(got.GetHideAttendees()), "HideAttendees")

	// TODO(ashmrtn): Uncomment when we figure out how to connect to the original
	// event.
	// assert.Equal(t, ptr.Val(expected.GetICalUId()), ptr.Val(got.GetICalUId()), "ICalUId")

	// Skip Id as it's tied to this specific instance of the item.

	assert.Equal(t, ptr.Val(expected.GetImportance()), ptr.Val(got.GetImportance()), "Importance")

	assert.Equal(t, expected.GetInstances(), got.GetInstances(), "Instances")

	assert.Equal(t, ptr.Val(expected.GetIsAllDay()), ptr.Val(got.GetIsAllDay()), "IsAllDay")

	assert.Equal(t, ptr.Val(expected.GetIsCancelled()), ptr.Val(got.GetIsCancelled()), "IsCancelled")

	assert.Equal(t, ptr.Val(expected.GetIsDraft()), ptr.Val(got.GetIsDraft()), "IsDraft")

	assert.Equal(t, ptr.Val(expected.GetIsOnlineMeeting()), ptr.Val(got.GetIsOnlineMeeting()), "IsOnlineMeeting")

	// TODO(ashmrtn): Uncomment when we figure out how to delegate event creation
	// to another user.
	// assert.Equal(t, ptr.Val(expected.GetIsOrganizer()), ptr.Val(got.GetIsOrganizer()), "IsOrganizer")

	assert.Equal(t, ptr.Val(expected.GetIsReminderOn()), ptr.Val(got.GetIsReminderOn()), "IsReminderOn")

	// Skip LastModifiedDateTime as it's tied to this specific instance of the item.

	// Cheating a little here in the name of code-reuse. model.Location needs
	// custom compare logic because it has fields marked as "internal use only"
	// that seem to change.
	testElementsMatch(
		t,
		[]models.Locationable{expected.GetLocation()},
		[]models.Locationable{got.GetLocation()},
		false,
		locationEqual)

	testElementsMatch(t, expected.GetLocations(), got.GetLocations(), false, locationEqual)

	assert.Equal(t, expected.GetOnlineMeeting(), got.GetOnlineMeeting(), "OnlineMeeting")

	assert.Equal(
		t,
		ptr.Val(expected.GetOnlineMeetingProvider()),
		ptr.Val(got.GetOnlineMeetingProvider()),
		"OnlineMeetingProvider")

	assert.Equal(
		t,
		ptr.Val(expected.GetOnlineMeetingUrl()),
		ptr.Val(got.GetOnlineMeetingUrl()),
		"OnlineMeetingUrl")

	// TODO(ashmrtn): Uncomment when we figure out how to delegate event creation
	// to another user.
	// assert.Equal(t, expected.GetOrganizer(), got.GetOrganizer(), "Organizer")

	assert.Equal(
		t,
		ptr.Val(expected.GetOriginalEndTimeZone()),
		ptr.Val(got.GetOriginalEndTimeZone()),
		"OriginalEndTimeZone")

	assert.Equal(
		t,
		ptr.Val(expected.GetOriginalStart()),
		ptr.Val(got.GetOriginalStart()),
		"OriginalStart")

	assert.Equal(
		t,
		ptr.Val(expected.GetOriginalStartTimeZone()),
		ptr.Val(got.GetOriginalStartTimeZone()),
		"OriginalStartTimeZone")

	assert.Equal(t, expected.GetRecurrence(), got.GetRecurrence(), "Recurrence")

	assert.Equal(
		t,
		ptr.Val(expected.GetReminderMinutesBeforeStart()),
		ptr.Val(got.GetReminderMinutesBeforeStart()),
		"ReminderMinutesBeforeStart",
	)

	assert.Equal(
		t,
		ptr.Val(expected.GetResponseRequested()),
		ptr.Val(got.GetResponseRequested()),
		"ResponseRequested")

	// TODO(ashmrtn): Uncomment when we figure out how to connect to the original
	// event.
	// assert.Equal(t, expected.GetResponseStatus(), got.GetResponseStatus(), "ResponseStatus")

	assert.Equal(t, ptr.Val(expected.GetSensitivity()), ptr.Val(got.GetSensitivity()), "Sensitivity")

	assert.Equal(t, ptr.Val(expected.GetSeriesMasterId()), ptr.Val(got.GetSeriesMasterId()), "SeriesMasterId")

	assert.Equal(t, ptr.Val(expected.GetShowAs()), ptr.Val(got.GetShowAs()), "ShowAs")

	assert.Equal(t, expected.GetStart(), got.GetStart(), "Start")

	assert.Equal(t, ptr.Val(expected.GetSubject()), ptr.Val(got.GetSubject()), "Subject")

	assert.Equal(t, ptr.Val(expected.GetTransactionId()), ptr.Val(got.GetTransactionId()), "TransactionId")

	// Skip LastModifiedDateTime as it's tied to this specific instance of the item.

	assert.Equal(t, ptr.Val(expected.GetType()), ptr.Val(got.GetType()), "Type")
}

func compareExchangeEmail(
	t *testing.T,
	expected map[string][]byte,
	item data.Stream,
) {
	itemData, err := io.ReadAll(item.ToReader())
	if !assert.NoError(t, err, "reading collection item", item.UUID(), clues.ToCore(err)) {
		return
	}

	itemMessage, err := support.CreateMessageFromBytes(itemData)
	if !assert.NoError(t, err, "deserializing backed up message", clues.ToCore(err)) {
		return
	}

	expectedBytes, ok := expected[ptr.Val(itemMessage.GetSubject())]
	if !assert.True(t, ok, "unexpected item with Subject %q", ptr.Val(itemMessage.GetSubject())) {
		return
	}

	expectedMessage, err := support.CreateMessageFromBytes(expectedBytes)
	assert.NoError(t, err, "deserializing source message", clues.ToCore(err))

	checkMessage(t, expectedMessage, itemMessage)
}

func compareExchangeContact(
	t *testing.T,
	colPath path.Path,
	expected map[string][]byte,
	item data.Stream,
) {
	itemData, err := io.ReadAll(item.ToReader())
	if !assert.NoError(t, err, "reading collection item", item.UUID(), clues.ToCore(err)) {
		return
	}

	itemContact, err := support.CreateContactFromBytes(itemData)
	if !assert.NoError(t, err, "deserializing backed up contact", clues.ToCore(err)) {
		return
	}

	expectedBytes, ok := expected[ptr.Val(itemContact.GetMiddleName())]
	if !assert.True(t, ok, "unexpected item with middle name %q", ptr.Val(itemContact.GetMiddleName())) {
		return
	}

	expectedContact, err := support.CreateContactFromBytes(expectedBytes)
	if !assert.NoError(t, err, "deserializing source contact") {
		return
	}

	checkContact(t, colPath, expectedContact, itemContact)
}

func compareExchangeEvent(
	t *testing.T,
	expected map[string][]byte,
	item data.Stream,
) {
	itemData, err := io.ReadAll(item.ToReader())
	if !assert.NoError(t, err, "reading collection item", item.UUID(), clues.ToCore(err)) {
		return
	}

	itemEvent, err := support.CreateEventFromBytes(itemData)
	if !assert.NoError(t, err, "deserializing backed up contact", clues.ToCore(err)) {
		return
	}

	expectedBytes, ok := expected[ptr.Val(itemEvent.GetSubject())]
	if !assert.True(t, ok, "unexpected item with subject %q", ptr.Val(itemEvent.GetSubject())) {
		return
	}

	expectedEvent, err := support.CreateEventFromBytes(expectedBytes)
	assert.NoError(t, err, "deserializing source contact", clues.ToCore(err))

	checkEvent(t, expectedEvent, itemEvent)
}

func permissionEqual(expected metadata.Permission, got metadata.Permission) bool {
	if !strings.EqualFold(expected.Email, got.Email) {
		return false
	}

	if (expected.Expiration == nil && got.Expiration != nil) ||
		(expected.Expiration != nil && got.Expiration == nil) {
		return false
	}

	if expected.Expiration != nil &&
		got.Expiration != nil &&
		!expected.Expiration.Equal(ptr.Val(got.Expiration)) {
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

func compareDriveItem(
	t *testing.T,
	expected map[string][]byte,
	item data.Stream,
	config ConfigInfo,
	rootDir bool,
) bool {
	// Skip Drive permissions in the folder that used to be the root. We don't
	// have a good way to materialize these in the test right now.
	if rootDir && item.UUID() == metadata.DirMetaFileSuffix {
		return false
	}

	buf, err := io.ReadAll(item.ToReader())
	if !assert.NoError(t, err, clues.ToCore(err)) {
		return true
	}

	var (
		displayName string
		name        = item.UUID()
		isMeta      = metadata.HasMetaSuffix(name)
	)

	if !isMeta {
		oitem := item.(*onedrive.Item)
		info := oitem.Info()

		if info.OneDrive != nil {
			displayName = oitem.Info().OneDrive.ItemName

			// Don't need to check SharePoint because it was added after we stopped
			// adding meta files to backup details.
			assert.False(t, oitem.Info().OneDrive.IsMeta, "meta marker for non meta item %s", name)
		} else if info.SharePoint != nil {
			displayName = oitem.Info().SharePoint.ItemName
		} else {
			assert.Fail(t, "ItemInfo is not SharePoint or OneDrive")
		}
	}

	if isMeta {
		var itemType *metadata.Item

		assert.IsType(t, itemType, item)

		var (
			itemMeta     metadata.Metadata
			expectedMeta metadata.Metadata
		)

		err = json.Unmarshal(buf, &itemMeta)
		if !assert.NoError(t, err, "unmarshalling retrieved metadata for file", name, clues.ToCore(err)) {
			return true
		}

		key := name

		if strings.HasSuffix(name, metadata.MetaFileSuffix) {
			key = itemMeta.FileName
		}

		expectedData := expected[key]

		if !assert.NotNil(
			t,
			expectedData,
			"unexpected metadata file with name %s",
			name,
		) {
			return true
		}

		err = json.Unmarshal(expectedData, &expectedMeta)
		if !assert.NoError(t, err, "unmarshalling expected metadata", clues.ToCore(err)) {
			return true
		}

		// Only compare file names if we're using a version that expects them to be
		// set.
		if len(expectedMeta.FileName) > 0 {
			assert.Equal(t, expectedMeta.FileName, itemMeta.FileName)
		}

		if !config.Opts.RestorePermissions {
			assert.Equal(t, 0, len(itemMeta.Permissions))
			return true
		}

		// We cannot restore owner permissions, so skip checking them
		itemPerms := []metadata.Permission{}

		for _, p := range itemMeta.Permissions {
			if p.Roles[0] != "owner" {
				itemPerms = append(itemPerms, p)
			}
		}

		testElementsMatch(
			t,
			expectedMeta.Permissions,
			itemPerms,
			// sharepoint retrieves a superset of permissions
			// (all site admins, site groups, built in by default)
			// relative to the permissions changed by the test.
			config.Service == path.SharePointService,
			permissionEqual)

		return true
	}

	var fileData testOneDriveData

	err = json.Unmarshal(buf, &fileData)
	if !assert.NoError(t, err, "unmarshalling file data for file", name, clues.ToCore(err)) {
		return true
	}

	expectedData := expected[fileData.FileName]
	if !assert.NotNil(t, expectedData, "unexpected file with name", name) {
		return true
	}

	// OneDrive data items are just byte buffers of the data. Nothing special to
	// interpret. May need to do chunked comparisons in the future if we test
	// large item equality.
	// Compare against the version with the file name embedded because that's what
	// the auto-generated expected data has.
	assert.Equal(t, expectedData, buf)
	// Display name in ItemInfo should match the name the file was given in the
	// test. Name used for the lookup key has a `.data` suffix to make it unique
	// from the metadata files' lookup keys.
	assert.Equal(t, fileData.FileName, displayName+metadata.DataFileSuffix)

	return true
}

// compareItem compares the data returned by backup with the expected data.
// Returns true if a comparison was done else false. Bool return is mostly used
// to exclude OneDrive permissions for the root right now.
func compareItem(
	t *testing.T,
	colPath path.Path,
	expected map[string][]byte,
	service path.ServiceType,
	category path.CategoryType,
	item data.Stream,
	config ConfigInfo,
	rootDir bool,
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
			compareExchangeContact(t, colPath, expected, item)
		case path.EventsCategory:
			compareExchangeEvent(t, expected, item)
		default:
			assert.FailNowf(t, "unexpected Exchange category: %s", category.String())
		}

	case path.OneDriveService:
		return compareDriveItem(t, expected, item, config, rootDir)

	case path.SharePointService:
		if category != path.LibrariesCategory {
			assert.FailNowf(t, "unsupported SharePoint category: %s", category.String())
		}

		// SharePoint libraries reuses OneDrive code.
		return compareDriveItem(t, expected, item, config, rootDir)

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
		// TODO(ashmrtn): Remove when LocationPath is made part of BackupCollection
		// interface.
		if !assert.Implements(t, (*data.LocationPather)(nil), g) {
			continue
		}

		fp := g.FullPath()
		loc := g.(data.LocationPather).LocationPath()

		if fp.Service() == path.OneDriveService ||
			(fp.Service() == path.SharePointService && fp.Category() == path.LibrariesCategory) {
			dp, err := path.ToDrivePath(fp)
			if !assert.NoError(t, err, clues.ToCore(err)) {
				continue
			}

			loc = path.BuildDriveLocation(dp.DriveID, loc.Elements()...)
		}

		p, err := loc.ToDataLayerPath(
			fp.Tenant(),
			fp.ResourceOwner(),
			fp.Service(),
			fp.Category(),
			false)
		if !assert.NoError(t, err, clues.ToCore(err)) {
			continue
		}

		gotNames = append(gotNames, p.String())
	}

	assert.ElementsMatch(t, expectedNames, gotNames, "returned collections")
}

func checkCollections(
	t *testing.T,
	ctx context.Context, //revive:disable-line:context-as-argument
	expectedItems int,
	expected map[string]map[string][]byte,
	got []data.BackupCollection,
	config ConfigInfo,
) int {
	collectionsWithItems := []data.BackupCollection{}

	skipped := 0
	gotItems := 0

	for _, returned := range got {
		var (
			hasItems        bool
			service         = returned.FullPath().Service()
			category        = returned.FullPath().Category()
			expectedColData = expected[returned.FullPath().String()]
			folders         = returned.FullPath().Elements()
			rootDir         = folders[len(folders)-1] == config.Dest.ContainerName
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

			if !compareItem(
				t,
				returned.FullPath(),
				expectedColData,
				service,
				category,
				item,
				config,
				rootDir) {
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

func makeSharePointBackupSel(
	t *testing.T,
	dests []destAndCats,
) selectors.Selector {
	toInclude := [][]selectors.SharePointScope{}
	resourceOwners := map[string]struct{}{}

	for _, d := range dests {
		for c := range d.cats {
			if c != path.LibrariesCategory {
				assert.FailNowf(t, "unsupported category type %s", c.String())
			}

			resourceOwners[d.resourceOwner] = struct{}{}

			// nil owners here, we'll need to stitch this together
			// below after the loops are complete.
			sel := selectors.NewSharePointBackup(nil)

			toInclude = append(toInclude, sel.LibraryFolders(
				[]string{d.dest},
				selectors.PrefixMatch(),
			))
		}
	}

	sel := selectors.NewSharePointBackup(maps.Keys(resourceOwners))
	sel.Include(toInclude...)

	return sel.Selector
}

// backupSelectorForExpected creates a selector that can be used to backup the
// given dests based on the item paths. Fails the test if items from
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

	case path.SharePointService:
		return makeSharePointBackupSel(t, dests)

	default:
		assert.FailNow(t, "unknown service type %s", service.String())
	}

	// Fix compile error about no return. Should not reach here.
	return selectors.Selector{}
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

func loadConnector(ctx context.Context, t *testing.T, r Resource) *GraphConnector {
	a := tester.NewM365Account(t)

	connector, err := NewGraphConnector(ctx, a, r)
	require.NoError(t, err, clues.ToCore(err))

	return connector
}
