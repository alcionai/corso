package sites

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody provides operations to call the evaluate method.
type ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody instantiates a new ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody and sets the default values.
func NewItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody()(*ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) {
    m := &ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetEvaluationInput gets the evaluationInput property value. The evaluationInput property
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetEvaluationInput()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpEvaluationInputable) {
    return m.GetBackingStore().Get("evaluationInput");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["evaluationInput"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateDlpEvaluationInputFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationInput(val.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpEvaluationInput))
        }
        return nil
    }
    res["notificationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateDlpNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationInfo(val.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpNotification))
        }
        return nil
    }
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
// GetNotificationInfo gets the notificationInfo property value. The notificationInfo property
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetNotificationInfo()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpNotificationable) {
    return m.GetBackingStore().Get("notificationInfo");
}
// GetTarget gets the target property value. The target property
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetTarget()(*string) {
    return m.GetBackingStore().Get("target");
}
// Serialize serializes information the current object
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("evaluationInput", m.GetEvaluationInput())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("notificationInfo", m.GetNotificationInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetEvaluationInput sets the evaluationInput property value. The evaluationInput property
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetEvaluationInput(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpEvaluationInputable)() {
    m.GetBackingStore().Set("evaluationInput", value)
}
// SetNotificationInfo sets the notificationInfo property value. The notificationInfo property
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetNotificationInfo(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpNotificationable)() {
    m.GetBackingStore().Set("notificationInfo", value)
}
// SetTarget sets the target property value. The target property
func (m *ItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetTarget(value *string)() {
    m.GetBackingStore().Set("target", value)
}
