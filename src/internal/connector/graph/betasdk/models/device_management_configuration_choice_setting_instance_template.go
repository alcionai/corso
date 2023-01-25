package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingInstanceTemplate 
type DeviceManagementConfigurationChoiceSettingInstanceTemplate struct {
    DeviceManagementConfigurationSettingInstanceTemplate
    // Choice Setting Value Template
    choiceSettingValueTemplate DeviceManagementConfigurationChoiceSettingValueTemplateable
}
// NewDeviceManagementConfigurationChoiceSettingInstanceTemplate instantiates a new DeviceManagementConfigurationChoiceSettingInstanceTemplate and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingInstanceTemplate()(*DeviceManagementConfigurationChoiceSettingInstanceTemplate) {
    m := &DeviceManagementConfigurationChoiceSettingInstanceTemplate{
        DeviceManagementConfigurationSettingInstanceTemplate: *NewDeviceManagementConfigurationSettingInstanceTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationChoiceSettingInstanceTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingInstanceTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingInstanceTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingInstanceTemplate(), nil
}
// GetChoiceSettingValueTemplate gets the choiceSettingValueTemplate property value. Choice Setting Value Template
func (m *DeviceManagementConfigurationChoiceSettingInstanceTemplate) GetChoiceSettingValueTemplate()(DeviceManagementConfigurationChoiceSettingValueTemplateable) {
    return m.choiceSettingValueTemplate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingInstanceTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstanceTemplate.GetFieldDeserializers()
    res["choiceSettingValueTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceManagementConfigurationChoiceSettingValueTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChoiceSettingValueTemplate(val.(DeviceManagementConfigurationChoiceSettingValueTemplateable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingInstanceTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstanceTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("choiceSettingValueTemplate", m.GetChoiceSettingValueTemplate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChoiceSettingValueTemplate sets the choiceSettingValueTemplate property value. Choice Setting Value Template
func (m *DeviceManagementConfigurationChoiceSettingInstanceTemplate) SetChoiceSettingValueTemplate(value DeviceManagementConfigurationChoiceSettingValueTemplateable)() {
    m.choiceSettingValueTemplate = value
}
