package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationStringSettingValueDefinition 
type DeviceManagementConfigurationStringSettingValueDefinition struct {
    DeviceManagementConfigurationSettingValueDefinition
    // Supported file types for this setting.
    fileTypes []string
    // The format property
    format *DeviceManagementConfigurationStringFormat
    // Regular expression or any xml or json schema that the input string should match
    inputValidationSchema *string
    // Specifies whether the setting needs to be treated as a secret. Settings marked as yes will be encrypted in transit and at rest and will be displayed as asterisks when represented in the UX.
    isSecret *bool
    // Maximum length of string
    maximumLength *int64
    // Minimum length of string
    minimumLength *int64
}
// NewDeviceManagementConfigurationStringSettingValueDefinition instantiates a new DeviceManagementConfigurationStringSettingValueDefinition and sets the default values.
func NewDeviceManagementConfigurationStringSettingValueDefinition()(*DeviceManagementConfigurationStringSettingValueDefinition) {
    m := &DeviceManagementConfigurationStringSettingValueDefinition{
        DeviceManagementConfigurationSettingValueDefinition: *NewDeviceManagementConfigurationSettingValueDefinition(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationStringSettingValueDefinition";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationStringSettingValueDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationStringSettingValueDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationStringSettingValueDefinition(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSettingValueDefinition.GetFieldDeserializers()
    res["fileTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFileTypes(res)
        }
        return nil
    }
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationStringFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val.(*DeviceManagementConfigurationStringFormat))
        }
        return nil
    }
    res["inputValidationSchema"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInputValidationSchema(val)
        }
        return nil
    }
    res["isSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSecret(val)
        }
        return nil
    }
    res["maximumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumLength(val)
        }
        return nil
    }
    res["minimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumLength(val)
        }
        return nil
    }
    return res
}
// GetFileTypes gets the fileTypes property value. Supported file types for this setting.
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetFileTypes()([]string) {
    return m.fileTypes
}
// GetFormat gets the format property value. The format property
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetFormat()(*DeviceManagementConfigurationStringFormat) {
    return m.format
}
// GetInputValidationSchema gets the inputValidationSchema property value. Regular expression or any xml or json schema that the input string should match
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetInputValidationSchema()(*string) {
    return m.inputValidationSchema
}
// GetIsSecret gets the isSecret property value. Specifies whether the setting needs to be treated as a secret. Settings marked as yes will be encrypted in transit and at rest and will be displayed as asterisks when represented in the UX.
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetIsSecret()(*bool) {
    return m.isSecret
}
// GetMaximumLength gets the maximumLength property value. Maximum length of string
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetMaximumLength()(*int64) {
    return m.maximumLength
}
// GetMinimumLength gets the minimumLength property value. Minimum length of string
func (m *DeviceManagementConfigurationStringSettingValueDefinition) GetMinimumLength()(*int64) {
    return m.minimumLength
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationStringSettingValueDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSettingValueDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetFileTypes() != nil {
        err = writer.WriteCollectionOfStringValues("fileTypes", m.GetFileTypes())
        if err != nil {
            return err
        }
    }
    if m.GetFormat() != nil {
        cast := (*m.GetFormat()).String()
        err = writer.WriteStringValue("format", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("inputValidationSchema", m.GetInputValidationSchema())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSecret", m.GetIsSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("maximumLength", m.GetMaximumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("minimumLength", m.GetMinimumLength())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFileTypes sets the fileTypes property value. Supported file types for this setting.
func (m *DeviceManagementConfigurationStringSettingValueDefinition) SetFileTypes(value []string)() {
    m.fileTypes = value
}
// SetFormat sets the format property value. The format property
func (m *DeviceManagementConfigurationStringSettingValueDefinition) SetFormat(value *DeviceManagementConfigurationStringFormat)() {
    m.format = value
}
// SetInputValidationSchema sets the inputValidationSchema property value. Regular expression or any xml or json schema that the input string should match
func (m *DeviceManagementConfigurationStringSettingValueDefinition) SetInputValidationSchema(value *string)() {
    m.inputValidationSchema = value
}
// SetIsSecret sets the isSecret property value. Specifies whether the setting needs to be treated as a secret. Settings marked as yes will be encrypted in transit and at rest and will be displayed as asterisks when represented in the UX.
func (m *DeviceManagementConfigurationStringSettingValueDefinition) SetIsSecret(value *bool)() {
    m.isSecret = value
}
// SetMaximumLength sets the maximumLength property value. Maximum length of string
func (m *DeviceManagementConfigurationStringSettingValueDefinition) SetMaximumLength(value *int64)() {
    m.maximumLength = value
}
// SetMinimumLength sets the minimumLength property value. Minimum length of string
func (m *DeviceManagementConfigurationStringSettingValueDefinition) SetMinimumLength(value *int64)() {
    m.minimumLength = value
}
