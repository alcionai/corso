package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementCollectionSettingInstance 
type DeviceManagementCollectionSettingInstance struct {
    DeviceManagementSettingInstance
    // The collection of values
    value []DeviceManagementSettingInstanceable
}
// NewDeviceManagementCollectionSettingInstance instantiates a new DeviceManagementCollectionSettingInstance and sets the default values.
func NewDeviceManagementCollectionSettingInstance()(*DeviceManagementCollectionSettingInstance) {
    m := &DeviceManagementCollectionSettingInstance{
        DeviceManagementSettingInstance: *NewDeviceManagementSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementCollectionSettingInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementCollectionSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementCollectionSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementCollectionSettingInstance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementCollectionSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetValue gets the value property value. The collection of values
func (m *DeviceManagementCollectionSettingInstance) GetValue()([]DeviceManagementSettingInstanceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementCollectionSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetValue sets the value property value. The collection of values
func (m *DeviceManagementCollectionSettingInstance) SetValue(value []DeviceManagementSettingInstanceable)() {
    m.value = value
}
