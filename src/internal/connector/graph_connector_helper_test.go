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

func notNilAndEq[T any](t *testing.T, expected *T, got *T, msg string) {
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

	notNilAndEq(t, expected.GetBody().GetContentType(), got.GetBody().GetContentType(), "Body.ContentType")

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

	notNilAndEq(t, expected.GetHasAttachments(), got.GetHasAttachments(), "HasAttachments")

	// Skip Id as it's tied to this specific instance of the item.

	notNilAndEq(t, expected.GetImportance(), got.GetImportance(), "Importance")

	notNilAndEq(t, expected.GetInferenceClassification(), got.GetInferenceClassification(), "InferenceClassification")

	assert.Equal(t, expected.GetInternetMessageHeaders(), got.GetInternetMessageHeaders(), "InternetMessageHeaders")

	notNilAndEq(t, expected.GetInternetMessageId(), got.GetInternetMessageId(), "InternetMessageId")

	notNilAndEq(
		t,
		expected.GetIsDeliveryReceiptRequested(),
		got.GetIsDeliveryReceiptRequested(),
		"IsDeliverReceiptRequested",
	)

	notNilAndEq(t, expected.GetIsDraft(), got.GetIsDraft(), "IsDraft")

	notNilAndEq(t, expected.GetIsRead(), got.GetIsRead(), "IsRead")

	notNilAndEq(t, expected.GetIsReadReceiptRequested(), got.GetIsReadReceiptRequested(), "IsReadReceiptRequested")

	// Skip LastModifiedDateTime as it's tied to this specific instance of the item.

	// Skip ParentFolderId as we restore to a different folder by default.

	notNilAndEq(t, expected.GetReceivedDateTime(), got.GetReceivedDateTime(), "ReceivedDateTime")

	assert.Equal(t, expected.GetReplyTo(), got.GetReplyTo(), "ReplyTo")

	assert.Equal(t, expected.GetSender(), got.GetSender(), "Sender")

	notNilAndEq(t, expected.GetSentDateTime(), got.GetSentDateTime(), "SentDateTime")

	notNilAndEq(t, expected.GetSubject(), got.GetSubject(), "Subject")

	assert.Equal(t, expected.GetToRecipients(), got.GetToRecipients(), "ToRecipients")

	// Skip WebLink as it's tied to this specific instance of the item.

	assert.Equal(t, expected.GetUniqueBody(), got.GetUniqueBody(), "UniqueBody")
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

		expectedData[baseDestPath.String()] = make(map[string][]byte, len(info.items))

		for i := 0; i < len(info.items); i++ {
			c.Names[i] = info.items[i].name
			c.Data[i] = info.items[i].data

			expectedData[baseDestPath.String()][info.items[i].lookupKey] = info.items[i].data
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
