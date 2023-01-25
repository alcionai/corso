package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingInstanceTemplate 
type DeviceManagementConfigurationSimpleSettingInstanceTemplate struct {
    DeviceManagementConfigurationSettingInstanceTemplate
    // Simple Setting Value Template
    simpleSettingValueTemplate DeviceManagementConfigurationSimpleSettingValueTemplateable
}
// NewDeviceManagementConfigurationSimpleSettingInstanceTemplate instantiates a new DeviceManagementConfigurationSimpleSettingInstanceTemplate and sets the default values.
func NewDeviceManagementConfigurationSimpleSettingInstanceTemplate()(*DeviceManagementConfigurationSimpleSettingInstanceTemplate) {
    m := &DeviceManagementConfigurationSimpleSettingInstanceTemplate{
        DeviceManagementConfigurationSettingInstanceTemplate: *NewDeviceManagementConfigurationSettingInstanceTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSimpleSettingInstanceTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationSimpleSettingInstanceTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSimpleSettingInstanceTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSimpleSettingInstanceTemplate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSimpleSettingInstanceTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstanceTemplate.GetFieldDeserializers()
    res["simpleSettingValueTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationSimpleSettingValueTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimpleSettingValueTemplate(val.(DeviceManagementConfigurationSimpleSettingValueTemplateable))
        }
        return nil
    }
    return res
}
// GetSimpleSettingValueTemplate gets the simpleSettingValueTemplate property value. Simple Setting Value Template
func (m *DeviceManagementConfigurationSimpleSettingInstanceTemplate) GetSimpleSettingValueTemplate()(DeviceManagementConfigurationSimpleSettingValueTemplateable) {
    return m.simpleSettingValueTemplate
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSimpleSettingInstanceTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstanceTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("simpleSettingValueTemplate", m.GetSimpleSettingValueTemplate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSimpleSettingValueTemplate sets the simpleSettingValueTemplate property value. Simple Setting Value Template
func (m *DeviceManagementConfigurationSimpleSettingInstanceTemplate) SetSimpleSettingValueTemplate(value DeviceManagementConfigurationSimpleSettingValueTemplateable)() {
    m.simpleSettingValueTemplate = value
}
