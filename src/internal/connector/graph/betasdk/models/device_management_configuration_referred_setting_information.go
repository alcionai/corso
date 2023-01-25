package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationReferredSettingInformation referred setting information about reusable setting
type DeviceManagementConfigurationReferredSettingInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataType property
    odataType *string
    // Setting definition id that is being referred to a setting. Applicable for reusable setting
    settingDefinitionId *string
}
// NewDeviceManagementConfigurationReferredSettingInformation instantiates a new deviceManagementConfigurationReferredSettingInformation and sets the default values.
func NewDeviceManagementConfigurationReferredSettingInformation()(*DeviceManagementConfigurationReferredSettingInformation) {
    m := &DeviceManagementConfigurationReferredSettingInformation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateDeviceManagementConfigurationReferredSettingInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationReferredSettingInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationReferredSettingInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceManagementConfigurationReferredSettingInformation) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationReferredSettingInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["settingDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingDefinitionId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationReferredSettingInformation) GetOdataType()(*string) {
    return m.odataType
}
// GetSettingDefinitionId gets the settingDefinitionId property value. Setting definition id that is being referred to a setting. Applicable for reusable setting
func (m *DeviceManagementConfigurationReferredSettingInformation) GetSettingDefinitionId()(*string) {
    return m.settingDefinitionId
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationReferredSettingInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("settingDefinitionId", m.GetSettingDefinitionId())
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
func (m *DeviceManagementConfigurationReferredSettingInformation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceManagementConfigurationReferredSettingInformation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSettingDefinitionId sets the settingDefinitionId property value. Setting definition id that is being referred to a setting. Applicable for reusable setting
func (m *DeviceManagementConfigurationReferredSettingInformation) SetSettingDefinitionId(value *string)() {
    m.settingDefinitionId = value
}
