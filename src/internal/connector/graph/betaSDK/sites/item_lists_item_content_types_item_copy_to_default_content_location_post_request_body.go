package sites

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody provides operations to call the copyToDefaultContentLocation method.
type ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody instantiates a new ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody and sets the default values.
func NewItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody()(*ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) {
    m := &ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetDestinationFileName gets the destinationFileName property value. The destinationFileName property
func (m *ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetDestinationFileName()(*string) {
    return m.GetBackingStore().Get("destinationFileName");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destinationFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationFileName(val)
        }
        return nil
    }
    res["sourceFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateItemReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceFile(val.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemReference))
        }
        return nil
    }
    return res
}
// GetSourceFile gets the sourceFile property value. The sourceFile property
func (m *ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetSourceFile()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemReferenceable) {
    return m.GetBackingStore().Get("sourceFile");
}
// Serialize serializes information the current object
func (m *ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationFileName", m.GetDestinationFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sourceFile", m.GetSourceFile())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetDestinationFileName sets the destinationFileName property value. The destinationFileName property
func (m *ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetDestinationFileName(value *string)() {
    m.GetBackingStore().Set("destinationFileName", value)
}
// SetSourceFile sets the sourceFile property value. The sourceFile property
func (m *ItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetSourceFile(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemReferenceable)() {
    m.GetBackingStore().Set("sourceFile", value)
}
