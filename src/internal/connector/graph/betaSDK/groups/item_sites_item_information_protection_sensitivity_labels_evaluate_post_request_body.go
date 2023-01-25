package groups

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody provides operations to call the evaluate method.
type ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The currentLabel property
    currentLabel ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CurrentLabelable
    // The discoveredSensitiveTypes property
    discoveredSensitiveTypes []ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DiscoveredSensitiveTypeable
}
// NewItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody instantiates a new ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody and sets the default values.
func NewItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody()(*ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) {
    m := &ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCurrentLabel gets the currentLabel property value. The currentLabel property
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) GetCurrentLabel()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CurrentLabelable) {
    return m.currentLabel
}
// GetDiscoveredSensitiveTypes gets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) GetDiscoveredSensitiveTypes()([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DiscoveredSensitiveTypeable) {
    return m.discoveredSensitiveTypes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateCurrentLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentLabel(val.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CurrentLabelable))
        }
        return nil
    }
    res["discoveredSensitiveTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateDiscoveredSensitiveTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DiscoveredSensitiveTypeable, len(val))
            for i, v := range val {
                res[i] = v.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DiscoveredSensitiveTypeable)
            }
            m.SetDiscoveredSensitiveTypes(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("currentLabel", m.GetCurrentLabel())
        if err != nil {
            return err
        }
    }
    if m.GetDiscoveredSensitiveTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDiscoveredSensitiveTypes()))
        for i, v := range m.GetDiscoveredSensitiveTypes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("discoveredSensitiveTypes", cast)
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
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCurrentLabel sets the currentLabel property value. The currentLabel property
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) SetCurrentLabel(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CurrentLabelable)() {
    m.currentLabel = value
}
// SetDiscoveredSensitiveTypes sets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) SetDiscoveredSensitiveTypes(value []ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DiscoveredSensitiveTypeable)() {
    m.discoveredSensitiveTypes = value
}
