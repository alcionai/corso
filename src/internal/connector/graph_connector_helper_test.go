package connector

import (
	"io"
	"reflect"
	"testing"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alcionai/corso/src/internal/connector/mockconnector"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/internal/data"
	"github.com/alcionai/corso/src/pkg/control"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/selectors"
)

func mustToDataLayerPath(
	t *testing.T,
	service path.ServiceType,
	tenant, user string,
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
		res, err = pb.ToDataLayerExchangePathForCategory(tenant, user, category, isItem)
	case path.OneDriveService:
		require.Equal(t, path.FilesCategory, category)

		res, err = pb.ToDataLayerOneDrivePath(tenant, user, isItem)

	default:
		err = errors.Errorf("bad service type %s", service.String())
	}

	require.NoError(t, err)

	return res
}

func emptyOrEqual[T any](t *testing.T, expected *T, got *T, msg string) {
	t.Helper()

	if expected == nil || got == nil {
		// Creates either the zero value or gets the value pointed to.
		assert.Equal(t, reflect.ValueOf(expected).Elem(), reflect.ValueOf(got).Elem())
		return
	}

	assert.Equal(t, *expected, *got, msg)
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
}

func checkMessage(
	t *testing.T,
	expected models.Messageable,
	got models.Messageable,
) {
	assert.Equal(t, expected.GetBccRecipients(), got.GetBccRecipients(), "BccRecipients")

	emptyOrEqual(t, expected.GetBody().GetContentType(), got.GetBody().GetContentType(), "Body.ContentType")

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

	assert.Equal(t, expected.GetFlag(), got.GetFlag(), "Flag")

	assert.Equal(t, expected.GetFrom(), got.GetFrom(), "From")

	emptyOrEqual(t, expected.GetHasAttachments(), got.GetHasAttachments(), "HasAttachments")

	// Skip Id as it's tied to this specific instance of the item.

	emptyOrEqual(t, expected.GetImportance(), got.GetImportance(), "Importance")

	emptyOrEqual(t, expected.GetInferenceClassification(), got.GetInferenceClassification(), "InferenceClassification")

	assert.Equal(t, expected.GetInternetMessageHeaders(), got.GetInternetMessageHeaders(), "InternetMessageHeaders")

	emptyOrEqual(t, expected.GetInternetMessageId(), got.GetInternetMessageId(), "InternetMessageId")

	emptyOrEqual(
		t,
		expected.GetIsDeliveryReceiptRequested(),
		got.GetIsDeliveryReceiptRequested(),
		"IsDeliverReceiptRequested",
	)

	emptyOrEqual(t, expected.GetIsDraft(), got.GetIsDraft(), "IsDraft")

	emptyOrEqual(t, expected.GetIsRead(), got.GetIsRead(), "IsRead")

	emptyOrEqual(t, expected.GetIsReadReceiptRequested(), got.GetIsReadReceiptRequested(), "IsReadReceiptRequested")

	// Skip LastModifiedDateTime as it's tied to this specific instance of the item.

	// Skip ParentFolderId as we restore to a different folder by default.

	emptyOrEqual(t, expected.GetReceivedDateTime(), got.GetReceivedDateTime(), "ReceivedDateTime")

	assert.Equal(t, expected.GetReplyTo(), got.GetReplyTo(), "ReplyTo")

	assert.Equal(t, expected.GetSender(), got.GetSender(), "Sender")

	emptyOrEqual(t, expected.GetSentDateTime(), got.GetSentDateTime(), "SentDateTime")

	emptyOrEqual(t, expected.GetSubject(), got.GetSubject(), "Subject")

	assert.Equal(t, expected.GetToRecipients(), got.GetToRecipients(), "ToRecipients")

	// Skip WebLink as it's tied to this specific instance of the item.

	assert.Equal(t, expected.GetUniqueBody(), got.GetUniqueBody(), "UniqueBody")
}

func checkContact(
	t *testing.T,
	expected models.Contactable,
	got models.Contactable,
) {
	emptyOrEqual(t, expected.GetAssistantName(), got.GetAssistantName(), "AssistantName")

	emptyOrEqual(t, expected.GetBirthday(), got.GetBirthday(), "Birthday")

	assert.Equal(t, expected.GetBusinessAddress(), got.GetBusinessAddress())

	emptyOrEqual(t, expected.GetBusinessHomePage(), got.GetBusinessHomePage(), "BusinessHomePage")

	assert.Equal(t, expected.GetBusinessPhones(), got.GetBusinessPhones())

	assert.Equal(t, expected.GetCategories(), got.GetCategories())

	// Skip ChangeKey as it's tied to this specific instance of the item.

	assert.Equal(t, expected.GetChildren(), got.GetChildren())

	emptyOrEqual(t, expected.GetCompanyName(), got.GetCompanyName(), "CompanyName")

	// Skip CreatedDateTime as it's tied to this specific instance of the item.

	emptyOrEqual(t, expected.GetDepartment(), got.GetDepartment(), "Department")

	emptyOrEqual(t, expected.GetDisplayName(), got.GetDisplayName(), "DisplayName")

	assert.Equal(t, expected.GetEmailAddresses(), got.GetEmailAddresses())

	emptyOrEqual(t, expected.GetFileAs(), got.GetFileAs(), "FileAs")

	emptyOrEqual(t, expected.GetGeneration(), got.GetGeneration(), "Generation")

	emptyOrEqual(t, expected.GetGivenName(), got.GetGivenName(), "GivenName")

	assert.Equal(t, expected.GetHomeAddress(), got.GetHomeAddress())

	assert.Equal(t, expected.GetHomePhones(), got.GetHomePhones())

	// Skip CreatedDateTime as it's tied to this specific instance of the item.

	assert.Equal(t, expected.GetImAddresses(), got.GetImAddresses())

	emptyOrEqual(t, expected.GetInitials(), got.GetInitials(), "Initials")

	emptyOrEqual(t, expected.GetJobTitle(), got.GetJobTitle(), "JobTitle")

	// Skip CreatedDateTime as it's tied to this specific instance of the item.

	emptyOrEqual(t, expected.GetManager(), got.GetManager(), "Manager")

	emptyOrEqual(t, expected.GetMiddleName(), got.GetMiddleName(), "MiddleName")

	emptyOrEqual(t, expected.GetMobilePhone(), got.GetMobilePhone(), "MobilePhone")

	emptyOrEqual(t, expected.GetNickName(), got.GetNickName(), "NickName")

	emptyOrEqual(t, expected.GetOfficeLocation(), got.GetOfficeLocation(), "OfficeLocation")

	assert.Equal(t, expected.GetOtherAddress(), got.GetOtherAddress())

	// Skip ParentFolderId as it's tied to this specific instance of the item.

	emptyOrEqual(t, expected.GetPersonalNotes(), got.GetPersonalNotes(), "PersonalNotes")

	assert.Equal(t, expected.GetPhoto(), got.GetPhoto())

	emptyOrEqual(t, expected.GetProfession(), got.GetProfession(), "Profession")

	emptyOrEqual(t, expected.GetSpouseName(), got.GetSpouseName(), "SpouseName")

	emptyOrEqual(t, expected.GetSurname(), got.GetSurname(), "Surname")

	emptyOrEqual(t, expected.GetTitle(), got.GetTitle(), "Title")

	emptyOrEqual(t, expected.GetYomiCompanyName(), got.GetYomiCompanyName(), "YomiCompanyName")

	emptyOrEqual(t, expected.GetYomiGivenName(), got.GetYomiGivenName(), "YomiGivenName")

	emptyOrEqual(t, expected.GetYomiSurname(), got.GetYomiSurname(), "YomiSurname")
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

func compareItem(
	t *testing.T,
	expected map[string][]byte,
	service path.ServiceType,
	category path.CategoryType,
	item data.Stream,
) {
	switch service {
	case path.ExchangeService:
		switch category {
		case path.EmailCategory:
			compareExchangeEmail(t, expected, item)
		case path.ContactsCategory:
			compareExchangeContact(t, expected, item)
		default:
			assert.FailNowf(t, "unexpected Exchange category: %s", category.String())
		}
	default:
		assert.FailNowf(t, "unexpected service: %s", service.String())
	}
}

func checkHasCollections(
	t *testing.T,
	expected map[string]map[string][]byte,
	got []data.Collection,
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

	assert.ElementsMatch(t, expectedNames, gotNames)
}

func checkCollections(
	t *testing.T,
	expectedItems int,
	expected map[string]map[string][]byte,
	got []data.Collection,
) {
	checkHasCollections(t, expected, got)

	gotItems := 0

	for _, returned := range got {
		service := returned.FullPath().Service()
		category := returned.FullPath().Category()
		expectedColData := expected[returned.FullPath().String()]

		// Need to iterate through all items even if we don't expect to find a match
		// because otherwise we'll deadlock waiting for GC status. Unexpected or
		// missing collection paths will be reported by checkHasCollections.
		for item := range returned.Items() {
			gotItems++

			if expectedColData == nil {
				continue
			}

			compareItem(t, expectedColData, service, category, item)
		}
	}

	assert.Equal(t, expectedItems, gotItems, "expected items")
}

func collectionsForInfo(
	t *testing.T,
	service path.ServiceType,
	tenant, user string,
	dest control.RestoreDestination,
	allInfo []colInfo,
) (int, []data.Collection, map[string]map[string][]byte) {
	collections := make([]data.Collection, 0, len(allInfo))
	expectedData := make(map[string]map[string][]byte, len(allInfo))
	totalItems := 0

	for _, info := range allInfo {
		pth := mustToDataLayerPath(
			t,
			service,
			tenant,
			user,
			info.category,
			info.pathElements,
			false,
		)
		c := mockconnector.NewMockExchangeCollection(pth, len(info.items))

		// TODO(ashmrtn): This will need expanded/broken up by service/category
		// depending on how restore for that service/category places data back in
		// M365.
		baseDestPath := mustToDataLayerPath(
			t,
			service,
			tenant,
			user,
			info.category,
			[]string{dest.ContainerName},
			false,
		)

		baseExpected := expectedData[baseDestPath.String()]
		if baseExpected == nil {
			expectedData[baseDestPath.String()] = make(map[string][]byte, len(info.items))
			baseExpected = expectedData[baseDestPath.String()]
		}

		for i := 0; i < len(info.items); i++ {
			c.Names[i] = info.items[i].name
			c.Data[i] = info.items[i].data

			baseExpected[info.items[i].lookupKey] = info.items[i].data
		}

		collections = append(collections, c)
		totalItems += len(info.items)
	}

	return totalItems, collections, expectedData
}

func getSelectorWith(service path.ServiceType) selectors.Selector {
	s := selectors.ServiceUnknown

	switch service {
	case path.ExchangeService:
		s = selectors.ServiceExchange
	case path.OneDriveService:
		s = selectors.ServiceOneDrive
	}

	return selectors.Selector{
		Service: s,
	}
}
