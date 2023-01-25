package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSecretSettingValue 
type DeviceManagementConfigurationSecretSettingValue struct {
    DeviceManagementConfigurationSimpleSettingValue
    // Value of the secret setting.
    value *string
    // type tracking the encryption state of a secret setting value
    valueState *DeviceManagementConfigurationSecretSettingValueState
}
// NewDeviceManagementConfigurationSecretSettingValue instantiates a new DeviceManagementConfigurationSecretSettingValue and sets the default values.
func NewDeviceManagementConfigurationSecretSettingValue()(*DeviceManagementConfigurationSecretSettingValue) {
    m := &DeviceManagementConfigurationSecretSettingValue{
        DeviceManagementConfigurationSimpleSettingValue: *NewDeviceManagementConfigurationSimpleSettingValue(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationSecretSettingValue";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationSecretSettingValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationSecretSettingValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationSecretSettingValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationSecretSettingValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConfigurationSimpleSettingValue.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    res["valueState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementConfigurationSecretSettingValueState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueState(val.(*DeviceManagementConfigurationSecretSettingValueState))
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. Value of the secret setting.
func (m *DeviceManagementConfigurationSecretSettingValue) GetValue()(*string) {
    return m.value
}
// GetValueState gets the valueState property value. type tracking the encryption state of a secret setting value
func (m *DeviceManagementConfigurationSecretSettingValue) GetValueState()(*DeviceManagementConfigurationSecretSettingValueState) {
    return m.valueState
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationSecretSettingValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConfigurationSimpleSettingValue.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    if m.GetValueState() != nil {
        cast := (*m.GetValueState()).String()
        err = writer.WriteStringValue("valueState", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. Value of the secret setting.
func (m *DeviceManagementConfigurationSecretSettingValue) SetValue(value *string)() {
    m.value = value
}
// SetValueState sets the valueState property value. type tracking the encryption state of a secret setting value
func (m *DeviceManagementConfigurationSecretSettingValue) SetValueState(value *DeviceManagementConfigurationSecretSettingValueState)() {
    m.valueState = value
}
