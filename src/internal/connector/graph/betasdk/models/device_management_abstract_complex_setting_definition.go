package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementAbstractComplexSettingDefinition 
type DeviceManagementAbstractComplexSettingDefinition struct {
    DeviceManagementSettingDefinition
    // List of definition IDs for all possible implementations of this abstract complex setting
    implementations []string
}
// NewDeviceManagementAbstractComplexSettingDefinition instantiates a new DeviceManagementAbstractComplexSettingDefinition and sets the default values.
func NewDeviceManagementAbstractComplexSettingDefinition()(*DeviceManagementAbstractComplexSettingDefinition) {
    m := &DeviceManagementAbstractComplexSettingDefinition{
        DeviceManagementSettingDefinition: *NewDeviceManagementSettingDefinition(),
    }
    return m
}
// CreateDeviceManagementAbstractComplexSettingDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementAbstractComplexSettingDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementAbstractComplexSettingDefinition(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementAbstractComplexSettingDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementSettingDefinition.GetFieldDeserializers()
    res["implementations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetImplementations(res)
        }
        return nil
    }
    return res
}
// GetImplementations gets the implementations property value. List of definition IDs for all possible implementations of this abstract complex setting
func (m *DeviceManagementAbstractComplexSettingDefinition) GetImplementations()([]string) {
    return m.implementations
}
// Serialize serializes information the current object
func (m *DeviceManagementAbstractComplexSettingDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementSettingDefinition.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetImplementations() != nil {
        err = writer.WriteCollectionOfStringValues("implementations", m.GetImplementations())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetImplementations sets the implementations property value. List of definition IDs for all possible implementations of this abstract complex setting
func (m *DeviceManagementAbstractComplexSettingDefinition) SetImplementations(value []string)() {
    m.implementations = value
}
