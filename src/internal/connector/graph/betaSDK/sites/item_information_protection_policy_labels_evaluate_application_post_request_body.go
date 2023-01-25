package sites

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody provides operations to call the evaluateApplication method.
type ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody instantiates a new ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody and sets the default values.
func NewItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody()(*ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) {
    m := &ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) GetContentInfo()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable) {
    return m.GetBackingStore().Get("contentInfo");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["labelingOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateLabelingOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelingOptions(val.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.LabelingOptions))
        }
        return nil
    }
    return res
}
// GetLabelingOptions gets the labelingOptions property value. The labelingOptions property
func (m *ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) GetLabelingOptions()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.LabelingOptionsable) {
    return m.GetBackingStore().Get("labelingOptions");
}
// Serialize serializes information the current object
func (m *ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("labelingOptions", m.GetLabelingOptions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) SetContentInfo(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable)() {
    m.GetBackingStore().Set("contentInfo", value)
}
// SetLabelingOptions sets the labelingOptions property value. The labelingOptions property
func (m *ItemInformationProtectionPolicyLabelsEvaluateApplicationPostRequestBody) SetLabelingOptions(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.LabelingOptionsable)() {
    m.GetBackingStore().Set("labelingOptions", value)
}
