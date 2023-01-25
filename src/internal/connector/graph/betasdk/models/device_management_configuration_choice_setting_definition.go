package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationChoiceSettingDefinition 
type DeviceManagementConfigurationChoiceSettingDefinition struct {
    DeviceManagementConfigurationSettingDefinition
    // Default option for choice setting
    defaultOptionId *string
    // Options for the setting that can be selected
    options []DeviceManagementConfigurationOptionDefinitionable
}
// NewDeviceManagementConfigurationChoiceSettingDefinition instantiates a new DeviceManagementConfigurationChoiceSettingDefinition and sets the default values.
func NewDeviceManagementConfigurationChoiceSettingDefinition()(*DeviceManagementConfigurationChoiceSettingDefinition) {
    m := &DeviceManagementConfigurationChoiceSettingDefinition{
        DeviceManagementConfigurationSettingDefinition: *NewDeviceManagementConfigurationSettingDefinition(),
    }
    return m
}
// CreateDeviceManagementConfigurationChoiceSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationChoiceSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.deviceManagementConfigurationChoiceSettingCollectionDefinition":
                        return NewDeviceManagementConfigurationChoiceSettingCollectionDefinition(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationChoiceSettingDefinition(), nil
}
// GetDefaultOptionId gets the defaultOptionId property value. Default option for choice setting
func (m *DeviceManagementConfigurationChoiceSettingDefinition) GetDefaultOptionId()(*string) {
    return m.defaultOptionId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationChoiceSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingDefinition.GetFieldDeserializers()
    res["defaultOptionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultOptionId(val)
        }
        return nil
    }
    res["options"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementConfigurationOptionDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementConfigurationOptionDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementConfigurationOptionDefinitionable)
            }
            m.SetOptions(res)
        }
        return nil
    }
    return res
}
// GetOptions gets the options property value. Options for the setting that can be selected
func (m *DeviceManagementConfigurationChoiceSettingDefinition) GetOptions()([]DeviceManagementConfigurationOptionDefinitionable) {
    return m.options
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationChoiceSettingDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("defaultOptionId", m.GetDefaultOptionId())
        if err != nil {
            return err
        }
    }
    if m.GetOptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOptions()))
        for i, v := range m.GetOptions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("options", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultOptionId sets the defaultOptionId property value. Default option for choice setting
func (m *DeviceManagementConfigurationChoiceSettingDefinition) SetDefaultOptionId(value *string)() {
    m.defaultOptionId = value
}
// SetOptions sets the options property value. Options for the setting that can be selected
func (m *DeviceManagementConfigurationChoiceSettingDefinition) SetOptions(value []DeviceManagementConfigurationOptionDefinitionable)() {
    m.options = value
}
