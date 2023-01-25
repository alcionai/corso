package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingProfileConstraint 
type DeviceManagementSettingProfileConstraint struct {
    DeviceManagementConstraint
    // The source of the entity
    source *string
    // A collection of types this entity carries
    types []string
}
// NewDeviceManagementSettingProfileConstraint instantiates a new DeviceManagementSettingProfileConstraint and sets the default values.
func NewDeviceManagementSettingProfileConstraint()(*DeviceManagementSettingProfileConstraint) {
    m := &DeviceManagementSettingProfileConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementSettingProfileConstraint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementSettingProfileConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingProfileConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingProfileConstraint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingProfileConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
        }
        return nil
    }
    res["types"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTypes(res)
        }
        return nil
    }
    return res
}
// GetSource gets the source property value. The source of the entity
func (m *DeviceManagementSettingProfileConstraint) GetSource()(*string) {
    return m.source
}
// GetTypes gets the types property value. A collection of types this entity carries
func (m *DeviceManagementSettingProfileConstraint) GetTypes()([]string) {
    return m.types
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingProfileConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    if m.GetTypes() != nil {
        err = writer.WriteCollectionOfStringValues("types", m.GetTypes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSource sets the source property value. The source of the entity
func (m *DeviceManagementSettingProfileConstraint) SetSource(value *string)() {
    m.source = value
}
// SetTypes sets the types property value. A collection of types this entity carries
func (m *DeviceManagementSettingProfileConstraint) SetTypes(value []string)() {
    m.types = value
}
