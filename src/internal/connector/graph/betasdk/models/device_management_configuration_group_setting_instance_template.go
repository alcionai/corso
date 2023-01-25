package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationGroupSettingInstanceTemplate 
type DeviceManagementConfigurationGroupSettingInstanceTemplate struct {
    DeviceManagementConfigurationSettingInstanceTemplate
    // Group Setting Value Template
    groupSettingValueTemplate DeviceManagementConfigurationGroupSettingValueTemplateable
}
// NewDeviceManagementConfigurationGroupSettingInstanceTemplate instantiates a new DeviceManagementConfigurationGroupSettingInstanceTemplate and sets the default values.
func NewDeviceManagementConfigurationGroupSettingInstanceTemplate()(*DeviceManagementConfigurationGroupSettingInstanceTemplate) {
    m := &DeviceManagementConfigurationGroupSettingInstanceTemplate{
        DeviceManagementConfigurationSettingInstanceTemplate: *NewDeviceManagementConfigurationSettingInstanceTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationGroupSettingInstanceTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationGroupSettingInstanceTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationGroupSettingInstanceTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationGroupSettingInstanceTemplate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationGroupSettingInstanceTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstanceTemplate.GetFieldDeserializers()
    res["groupSettingValueTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationGroupSettingValueTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupSettingValueTemplate(val.(DeviceManagementConfigurationGroupSettingValueTemplateable))
        }
        return nil
    }
    return res
}
// GetGroupSettingValueTemplate gets the groupSettingValueTemplate property value. Group Setting Value Template
func (m *DeviceManagementConfigurationGroupSettingInstanceTemplate) GetGroupSettingValueTemplate()(DeviceManagementConfigurationGroupSettingValueTemplateable) {
    return m.groupSettingValueTemplate
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationGroupSettingInstanceTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstanceTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("groupSettingValueTemplate", m.GetGroupSettingValueTemplate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroupSettingValueTemplate sets the groupSettingValueTemplate property value. Group Setting Value Template
func (m *DeviceManagementConfigurationGroupSettingInstanceTemplate) SetGroupSettingValueTemplate(value DeviceManagementConfigurationGroupSettingValueTemplateable)() {
    m.groupSettingValueTemplate = value
}
