package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354 "github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
)

// ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody provides operations to call the evaluate method.
type ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The currentLabel property
    currentLabel ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CurrentLabelable
    // The discoveredSensitiveTypes property
    discoveredSensitiveTypes []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DiscoveredSensitiveTypeable
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
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) GetCurrentLabel()(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CurrentLabelable) {
    return m.currentLabel
}
// GetDiscoveredSensitiveTypes gets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) GetDiscoveredSensitiveTypes()([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DiscoveredSensitiveTypeable) {
    return m.discoveredSensitiveTypes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateCurrentLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentLabel(val.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CurrentLabelable))
        }
        return nil
    }
    res["discoveredSensitiveTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CreateDiscoveredSensitiveTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DiscoveredSensitiveTypeable, len(val))
            for i, v := range val {
                res[i] = v.(ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DiscoveredSensitiveTypeable)
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
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) SetCurrentLabel(value ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.CurrentLabelable)() {
    m.currentLabel = value
}
// SetDiscoveredSensitiveTypes sets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *ItemSitesItemInformationProtectionSensitivityLabelsEvaluatePostRequestBody) SetDiscoveredSensitiveTypes(value []ifda19816f54f079134d70c11e75d6b26799300cf72079e282f1d3bb9a6750354.DiscoveredSensitiveTypeable)() {
    m.discoveredSensitiveTypes = value
}
