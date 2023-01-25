package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementAbstractComplexSettingInstance 
type DeviceManagementAbstractComplexSettingInstance struct {
    DeviceManagementSettingInstance
    // The definition ID for the chosen implementation of this complex setting
    implementationId *string
    // The values that make up the complex setting
    value []DeviceManagementSettingInstanceable
}
// NewDeviceManagementAbstractComplexSettingInstance instantiates a new DeviceManagementAbstractComplexSettingInstance and sets the default values.
func NewDeviceManagementAbstractComplexSettingInstance()(*DeviceManagementAbstractComplexSettingInstance) {
    m := &DeviceManagementAbstractComplexSettingInstance{
        DeviceManagementSettingInstance: *NewDeviceManagementSettingInstance(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementAbstractComplexSettingInstance";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementAbstractComplexSettingInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementAbstractComplexSettingInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementAbstractComplexSettingInstance(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementAbstractComplexSettingInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingInstance.GetFieldDeserializers()
    res["implementationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImplementationId(val)
        }
        return nil
    }
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
// GetImplementationId gets the implementationId property value. The definition ID for the chosen implementation of this complex setting
func (m *DeviceManagementAbstractComplexSettingInstance) GetImplementationId()(*string) {
    return m.implementationId
}
// GetValue gets the value property value. The values that make up the complex setting
func (m *DeviceManagementAbstractComplexSettingInstance) GetValue()([]DeviceManagementSettingInstanceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceManagementAbstractComplexSettingInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementSettingInstance.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("implementationId", m.GetImplementationId())
        if err != nil {
            return err
        }
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
// SetImplementationId sets the implementationId property value. The definition ID for the chosen implementation of this complex setting
func (m *DeviceManagementAbstractComplexSettingInstance) SetImplementationId(value *string)() {
    m.implementationId = value
}
// SetValue sets the value property value. The values that make up the complex setting
func (m *DeviceManagementAbstractComplexSettingInstance) SetValue(value []DeviceManagementSettingInstanceable)() {
    m.value = value
}
