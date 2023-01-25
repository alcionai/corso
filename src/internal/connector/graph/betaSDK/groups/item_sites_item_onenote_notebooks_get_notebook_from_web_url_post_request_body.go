package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody provides operations to call the getNotebookFromWebUrl method.
type ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody instantiates a new ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody and sets the default values.
func NewItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody()(*ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) {
    m := &ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetWebUrl gets the webUrl property value. The webUrl property
func (m *ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) GetWebUrl()(*string) {
    return m.GetBackingStore().Get("webUrl");
}
// Serialize serializes information the current object
func (m *ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetWebUrl sets the webUrl property value. The webUrl property
func (m *ItemSitesItemOnenoteNotebooksGetNotebookFromWebUrlPostRequestBody) SetWebUrl(value *string)() {
    m.GetBackingStore().Set("webUrl", value)
}
