package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationStringSettingValue 
type DeviceManagementConfigurationStringSettingValue struct {
    DeviceManagementConfigurationSimpleSettingValue
    // Value of the string setting.
    value *string
}
// NewDeviceManagementConfigurationStringSettingValue instantiates a new DeviceManagementConfigurationStringSettingValue and sets the default values.
func NewDeviceManagementConfigurationStringSettingValue()(*DeviceManagementConfigurationStringSettingValue) {
    m := &DeviceManagementConfigurationStringSettingValue{
        DeviceManagementConfigurationSimpleSettingValue: *NewDeviceManagementConfigurationSimpleSettingValue(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementConfigurationStringSettingValue";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementConfigurationStringSettingValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationStringSettingValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.deviceManagementConfigurationReferenceSettingValue":
                        return NewDeviceManagementConfigurationReferenceSettingValue(), nil
                }
            }
        }
    }
    return NewDeviceManagementConfigurationStringSettingValue(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationStringSettingValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetValue gets the value property value. Value of the string setting.
func (m *DeviceManagementConfigurationStringSettingValue) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationStringSettingValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetValue sets the value property value. Value of the string setting.
func (m *DeviceManagementConfigurationStringSettingValue) SetValue(value *string)() {
    m.value = value
}
