package groups

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody provides operations to call the copyToDefaultContentLocation method.
type ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The destinationFileName property
    destinationFileName *string
    // The sourceFile property
    sourceFile ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemReferenceable
}
// NewItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody instantiates a new ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody and sets the default values.
func NewItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody()(*ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) {
    m := &ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDestinationFileName gets the destinationFileName property value. The destinationFileName property
func (m *ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetDestinationFileName()(*string) {
    return m.destinationFileName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
            m.SetSourceFile(val.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemReferenceable))
        }
        return nil
    }
    return res
}
// GetSourceFile gets the sourceFile property value. The sourceFile property
func (m *ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetSourceFile()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemReferenceable) {
    return m.sourceFile
}
// Serialize serializes information the current object
func (m *ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDestinationFileName sets the destinationFileName property value. The destinationFileName property
func (m *ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetDestinationFileName(value *string)() {
    m.destinationFileName = value
}
// SetSourceFile sets the sourceFile property value. The sourceFile property
func (m *ItemSitesItemListsItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetSourceFile(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemReferenceable)() {
    m.sourceFile = value
}
