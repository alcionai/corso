package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate 
type DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate struct {
    DeviceManagementConfigurationStringSettingValueDefaultTemplate
    // Default Constant Value
    constantValue *string
}
// NewDeviceManagementConfigurationStringSettingValueConstantDefaultTemplate instantiates a new DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate and sets the default values.
func NewDeviceManagementConfigurationStringSettingValueConstantDefaultTemplate()(*DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate) {
    m := &DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate{
        DeviceManagementConfigurationStringSettingValueDefaultTemplate: *NewDeviceManagementConfigurationStringSettingValueDefaultTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationStringSettingValueConstantDefaultTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationStringSettingValueConstantDefaultTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationStringSettingValueConstantDefaultTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationStringSettingValueConstantDefaultTemplate(), nil
}
// GetConstantValue gets the constantValue property value. Default Constant Value
func (m *DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate) GetConstantValue()(*string) {
    return m.constantValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationStringSettingValueDefaultTemplate.GetFieldDeserializers()
    res["constantValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConstantValue(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationStringSettingValueDefaultTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("constantValue", m.GetConstantValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConstantValue sets the constantValue property value. Default Constant Value
func (m *DeviceManagementConfigurationStringSettingValueConstantDefaultTemplate) SetConstantValue(value *string)() {
    m.constantValue = value
}
