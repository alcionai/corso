package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementStringSettingInstance 
type DeviceManagementStringSettingInstance struct {
    DeviceManagementSettingInstance
    // The string value
    value *string
}
// NewDeviceManagementStringSettingInstance instantiates a new DeviceManagementStringSettingInstance and sets the default values.
func NewDeviceManagementStringSettingInstance()(*DeviceManagementStringSettingInstance) {
    m := &DeviceManagementStringSettingInstance{
        DeviceManagementSettingInstance: *NewDeviceManagementSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementStringSettingInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementStringSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementStringSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementStringSettingInstance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementStringSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingInstance.GetFieldDeserializers()
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
// GetValue gets the value property value. The string value
func (m *DeviceManagementStringSettingInstance) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementStringSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementSettingInstance.Serialize(writer)
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
// SetValue sets the value property value. The string value
func (m *DeviceManagementStringSettingInstance) SetValue(value *string)() {
    m.value = value
}
