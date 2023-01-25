package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate 
type DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate struct {
    DeviceManagementConfigurationSettingInstanceTemplate
    // Linked policy may append values which are not present in the template.
    allowUnmanagedValues *bool
    // Choice Setting Collection Value Template
    choiceSettingCollectionValueTemplate []DeviceManagementConfigurationChoiceSettingValueTemplateable
}
// NewDeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate instantiates a new DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate()(*DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) {
    m := &DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate{
        DeviceManagementConfigurationSettingInstanceTemplate: *NewDeviceManagementConfigurationSettingInstanceTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationChoiceSettingCollectionInstanceTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingCollectionInstanceTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingCollectionInstanceTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate(), nil
}
// GetAllowUnmanagedValues gets the allowUnmanagedValues property value. Linked policy may append values which are not present in the template.
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) GetAllowUnmanagedValues()(*bool) {
    return m.allowUnmanagedValues
}
// GetChoiceSettingCollectionValueTemplate gets the choiceSettingCollectionValueTemplate property value. Choice Setting Collection Value Template
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) GetChoiceSettingCollectionValueTemplate()([]DeviceManagementConfigurationChoiceSettingValueTemplateable) {
    return m.choiceSettingCollectionValueTemplate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingInstanceTemplate.GetFieldDeserializers()
    res["allowUnmanagedValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUnmanagedValues(val)
        }
        return nil
    }
    res["choiceSettingCollectionValueTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationChoiceSettingValueTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationChoiceSettingValueTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationChoiceSettingValueTemplateable)
            }
            m.SetChoiceSettingCollectionValueTemplate(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingInstanceTemplate.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowUnmanagedValues", m.GetAllowUnmanagedValues())
        if err != nil {
            return err
        }
    }
    if m.GetChoiceSettingCollectionValueTemplate() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChoiceSettingCollectionValueTemplate()))
        for i, v := range m.GetChoiceSettingCollectionValueTemplate() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("choiceSettingCollectionValueTemplate", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowUnmanagedValues sets the allowUnmanagedValues property value. Linked policy may append values which are not present in the template.
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) SetAllowUnmanagedValues(value *bool)() {
    m.allowUnmanagedValues = value
}
// SetChoiceSettingCollectionValueTemplate sets the choiceSettingCollectionValueTemplate property value. Choice Setting Collection Value Template
func (m *DeviceManagementConfigurationChoiceSettingCollectionInstanceTemplate) SetChoiceSettingCollectionValueTemplate(value []DeviceManagementConfigurationChoiceSettingValueTemplateable)() {
    m.choiceSettingCollectionValueTemplate = value
}
