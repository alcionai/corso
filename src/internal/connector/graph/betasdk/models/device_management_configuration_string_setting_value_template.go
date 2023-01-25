package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationStringSettingValueTemplate 
type DeviceManagementConfigurationStringSettingValueTemplate struct {
    DeviceManagementConfigurationSimpleSettingValueTemplate
    // String Setting Value Default Template.
    defaultValue DeviceManagementConfigurationStringSettingValueDefaultTemplateable
}
// NewDeviceManagementConfigurationStringSettingValueTemplate instantiates a new DeviceManagementConfigurationStringSettingValueTemplate and sets the default values.
func NewDeviceManagementConfigurationStringSettingValueTemplate()(*DeviceManagementConfigurationStringSettingValueTemplate) {
    m := &DeviceManagementConfigurationStringSettingValueTemplate{
        DeviceManagementConfigurationSimpleSettingValueTemplate: *NewDeviceManagementConfigurationSimpleSettingValueTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationStringSettingValueTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationStringSettingValueTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationStringSettingValueTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationStringSettingValueTemplate(), nil
}
// GetDefaultValue gets the defaultValue property value. String Setting Value Default Template.
func (m *DeviceManagementConfigurationStringSettingValueTemplate) GetDefaultValue()(DeviceManagementConfigurationStringSettingValueDefaultTemplateable) {
    return m.defaultValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationStringSettingValueTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSimpleSettingValueTemplate.GetFieldDeserializers()
    res["defaultValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationStringSettingValueDefaultTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultValue(val.(DeviceManagementConfigurationStringSettingValueDefaultTemplateable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationStringSettingValueTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSimpleSettingValueTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("defaultValue", m.GetDefaultValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultValue sets the defaultValue property value. String Setting Value Default Template.
func (m *DeviceManagementConfigurationStringSettingValueTemplate) SetDefaultValue(value DeviceManagementConfigurationStringSettingValueDefaultTemplateable)() {
    m.defaultValue = value
}
