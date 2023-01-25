package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody provides operations to call the addCopyFromContentTypeHub method.
type ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody instantiates a new ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody and sets the default values.
func NewItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody()(*ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) {
    m := &ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetContentTypeId gets the contentTypeId property value. The contentTypeId property
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) GetContentTypeId()(*string) {
    return m.GetBackingStore().Get("contentTypeId");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentTypeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentTypeId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentTypeId", m.GetContentTypeId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetContentTypeId sets the contentTypeId property value. The contentTypeId property
func (m *ItemListsItemContentTypesAddCopyFromContentTypeHubPostRequestBody) SetContentTypeId(value *string)() {
    m.GetBackingStore().Set("contentTypeId", value)
}
