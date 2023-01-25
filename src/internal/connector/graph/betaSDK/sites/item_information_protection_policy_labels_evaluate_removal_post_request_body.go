package sites

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody provides operations to call the evaluateRemoval method.
type ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody instantiates a new ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody and sets the default values.
func NewItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody()(*ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) {
    m := &ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetContentInfo()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable) {
    return m.GetBackingStore().Get("contentInfo");
}
// GetDowngradeJustification gets the downgradeJustification property value. The downgradeJustification property
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetDowngradeJustification()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DowngradeJustificationable) {
    return m.GetBackingStore().Get("downgradeJustification");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateContentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfo))
        }
        return nil
    }
    res["downgradeJustification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateDowngradeJustificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDowngradeJustification(val.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DowngradeJustification))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) SetContentInfo(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable)() {
    m.GetBackingStore().Set("contentInfo", value)
}
// SetDowngradeJustification sets the downgradeJustification property value. The downgradeJustification property
func (m *ItemInformationProtectionPolicyLabelsEvaluateRemovalPostRequestBody) SetDowngradeJustification(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DowngradeJustificationable)() {
    m.GetBackingStore().Set("downgradeJustification", value)
}
