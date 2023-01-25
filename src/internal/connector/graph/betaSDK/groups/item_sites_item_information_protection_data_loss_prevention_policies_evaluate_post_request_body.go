package groups

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody provides operations to call the evaluate method.
type ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The evaluationInput property
    evaluationInput ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpEvaluationInputable
    // The notificationInfo property
    notificationInfo ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpNotificationable
    // The target property
    target *string
}
// NewItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody instantiates a new ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody and sets the default values.
func NewItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody()(*ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) {
    m := &ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEvaluationInput gets the evaluationInput property value. The evaluationInput property
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetEvaluationInput()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpEvaluationInputable) {
    return m.evaluationInput
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["evaluationInput"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateDlpEvaluationInputFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationInput(val.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpEvaluationInputable))
        }
        return nil
    }
    res["notificationInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateDlpNotificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationInfo(val.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpNotificationable))
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
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetNotificationInfo()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpNotificationable) {
    return m.notificationInfo
}
// GetTarget gets the target property value. The target property
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) GetTarget()(*string) {
    return m.target
}
// Serialize serializes information the current object
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEvaluationInput sets the evaluationInput property value. The evaluationInput property
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetEvaluationInput(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpEvaluationInputable)() {
    m.evaluationInput = value
}
// SetNotificationInfo sets the notificationInfo property value. The notificationInfo property
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetNotificationInfo(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DlpNotificationable)() {
    m.notificationInfo = value
}
// SetTarget sets the target property value. The target property
func (m *ItemSitesItemInformationProtectionDataLossPreventionPoliciesEvaluatePostRequestBody) SetTarget(value *string)() {
    m.target = value
}
