package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemListsItemContentTypesAddCopyPostRequestBody provides operations to call the addCopy method.
type ItemSitesItemListsItemContentTypesAddCopyPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemSitesItemListsItemContentTypesAddCopyPostRequestBody instantiates a new ItemSitesItemListsItemContentTypesAddCopyPostRequestBody and sets the default values.
func NewItemSitesItemListsItemContentTypesAddCopyPostRequestBody()(*ItemSitesItemListsItemContentTypesAddCopyPostRequestBody) {
    m := &ItemSitesItemListsItemContentTypesAddCopyPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemListsItemContentTypesAddCopyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemListsItemContentTypesAddCopyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemListsItemContentTypesAddCopyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemListsItemContentTypesAddCopyPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemSitesItemListsItemContentTypesAddCopyPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetContentType gets the contentType property value. The contentType property
func (m *ItemSitesItemListsItemContentTypesAddCopyPostRequestBody) GetContentType()(*string) {
    return m.GetBackingStore().Get("contentType");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemListsItemContentTypesAddCopyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemSitesItemListsItemContentTypesAddCopyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemListsItemContentTypesAddCopyPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemSitesItemListsItemContentTypesAddCopyPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetContentType sets the contentType property value. The contentType property
func (m *ItemSitesItemListsItemContentTypesAddCopyPostRequestBody) SetContentType(value *string)() {
    m.GetBackingStore().Set("contentType", value)
}
