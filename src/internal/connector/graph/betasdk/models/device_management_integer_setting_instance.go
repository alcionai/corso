package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementIntegerSettingInstance 
type DeviceManagementIntegerSettingInstance struct {
    DeviceManagementSettingInstance
    // The integer value
    value *int32
}
// NewDeviceManagementIntegerSettingInstance instantiates a new DeviceManagementIntegerSettingInstance and sets the default values.
func NewDeviceManagementIntegerSettingInstance()(*DeviceManagementIntegerSettingInstance) {
    m := &DeviceManagementIntegerSettingInstance{
        DeviceManagementSettingInstance: *NewDeviceManagementSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementIntegerSettingInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementIntegerSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementIntegerSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementIntegerSettingInstance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementIntegerSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingInstance.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
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
// GetValue gets the value property value. The integer value
func (m *DeviceManagementIntegerSettingInstance) GetValue()(*int32) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementIntegerSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The integer value
func (m *DeviceManagementIntegerSettingInstance) SetValue(value *int32)() {
    m.value = value
}
