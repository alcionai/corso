package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementComplexSettingInstance 
type DeviceManagementComplexSettingInstance struct {
    DeviceManagementSettingInstance
    // The values that make up the complex setting
    value []DeviceManagementSettingInstanceable
}
// NewDeviceManagementComplexSettingInstance instantiates a new DeviceManagementComplexSettingInstance and sets the default values.
func NewDeviceManagementComplexSettingInstance()(*DeviceManagementComplexSettingInstance) {
    m := &DeviceManagementComplexSettingInstance{
        DeviceManagementSettingInstance: *NewDeviceManagementSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementComplexSettingInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementComplexSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementComplexSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementComplexSettingInstance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementComplexSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingInstance.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementSettingInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementSettingInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementSettingInstanceable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The values that make up the complex setting
func (m *DeviceManagementComplexSettingInstance) GetValue()([]DeviceManagementSettingInstanceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementComplexSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The values that make up the complex setting
func (m *DeviceManagementComplexSettingInstance) SetValue(value []DeviceManagementSettingInstanceable)() {
    m.value = value
}
