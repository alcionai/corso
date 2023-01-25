package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate 
type DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate struct {
    DeviceManagementConfigurationSettingInstanceTemplate
    // Linked policy may append values which are not present in the template.
    allowUnmanagedValues *bool
    // Group Setting Collection Value Template
    groupSettingCollectionValueTemplate []DeviceManagementConfigurationGroupSettingValueTemplateable
}
// NewDeviceManagementConfigurationGroupSettingCollectionInstanceTemplate instantiates a new DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate and sets the default values.
func NewDeviceManagementConfigurationGroupSettingCollectionInstanceTemplate()(*DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate) {
    m := &DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate{
        DeviceManagementConfigurationSettingInstanceTemplate: *NewDeviceManagementConfigurationSettingInstanceTemplate(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationGroupSettingCollectionInstanceTemplate";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationGroupSettingCollectionInstanceTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationGroupSettingCollectionInstanceTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationGroupSettingCollectionInstanceTemplate(), nil
}
// GetAllowUnmanagedValues gets the allowUnmanagedValues property value. Linked policy may append values which are not present in the template.
func (m *DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate) GetAllowUnmanagedValues()(*bool) {
    return m.allowUnmanagedValues
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["groupSettingCollectionValueTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationGroupSettingValueTemplateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationGroupSettingValueTemplateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationGroupSettingValueTemplateable)
            }
            m.SetGroupSettingCollectionValueTemplate(res)
        }
        return nil
    }
    return res
}
// GetGroupSettingCollectionValueTemplate gets the groupSettingCollectionValueTemplate property value. Group Setting Collection Value Template
func (m *DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate) GetGroupSettingCollectionValueTemplate()([]DeviceManagementConfigurationGroupSettingValueTemplateable) {
    return m.groupSettingCollectionValueTemplate
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetGroupSettingCollectionValueTemplate() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupSettingCollectionValueTemplate()))
        for i, v := range m.GetGroupSettingCollectionValueTemplate() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupSettingCollectionValueTemplate", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowUnmanagedValues sets the allowUnmanagedValues property value. Linked policy may append values which are not present in the template.
func (m *DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate) SetAllowUnmanagedValues(value *bool)() {
    m.allowUnmanagedValues = value
}
// SetGroupSettingCollectionValueTemplate sets the groupSettingCollectionValueTemplate property value. Group Setting Collection Value Template
func (m *DeviceManagementConfigurationGroupSettingCollectionInstanceTemplate) SetGroupSettingCollectionValueTemplate(value []DeviceManagementConfigurationGroupSettingValueTemplateable)() {
    m.groupSettingCollectionValueTemplate = value
}
