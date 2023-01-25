package groups

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody provides operations to call the copyToDefaultContentLocation method.
type ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The destinationFileName property
    destinationFileName *string
    // The sourceFile property
    sourceFile ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemReferenceable
}
// NewItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody instantiates a new ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody and sets the default values.
func NewItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody()(*ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) {
    m := &ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDestinationFileName gets the destinationFileName property value. The destinationFileName property
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetDestinationFileName()(*string) {
    return m.destinationFileName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetSourceFile()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemReferenceable) {
    return m.sourceFile
}
// Serialize serializes information the current object
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDestinationFileName sets the destinationFileName property value. The destinationFileName property
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetDestinationFileName(value *string)() {
    m.destinationFileName = value
}
// SetSourceFile sets the sourceFile property value. The sourceFile property
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetSourceFile(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ItemReferenceable)() {
    m.sourceFile = value
}
