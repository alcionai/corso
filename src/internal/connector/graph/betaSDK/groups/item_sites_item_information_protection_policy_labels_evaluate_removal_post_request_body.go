package groups

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody provides operations to call the evaluateRemoval method.
type ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The contentInfo property
    contentInfo ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable
    // The downgradeJustification property
    downgradeJustification ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DowngradeJustificationable
}
// NewItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody instantiates a new ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody and sets the default values.
func NewItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody()(*ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) {
    m := &ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetContentInfo()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable) {
    return m.contentInfo
}
// GetDowngradeJustification gets the downgradeJustification property value. The downgradeJustification property
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetDowngradeJustification()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DowngradeJustificationable) {
    return m.downgradeJustification
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateContentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable))
        }
        return nil
    }
    res["downgradeJustification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateDowngradeJustificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDowngradeJustification(val.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DowngradeJustificationable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("downgradeJustification", m.GetDowngradeJustification())
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
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) SetContentInfo(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable)() {
    m.contentInfo = value
}
// SetDowngradeJustification sets the downgradeJustification property value. The downgradeJustification property
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) SetDowngradeJustification(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DowngradeJustificationable)() {
    m.downgradeJustification = value
}
