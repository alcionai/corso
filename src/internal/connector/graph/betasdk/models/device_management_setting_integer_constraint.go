package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingIntegerConstraint 
type DeviceManagementSettingIntegerConstraint struct {
    DeviceManagementConstraint
    // The maximum permitted value
    maximumValue *int32
    // The minimum permitted value
    minimumValue *int32
}
// NewDeviceManagementSettingIntegerConstraint instantiates a new DeviceManagementSettingIntegerConstraint and sets the default values.
func NewDeviceManagementSettingIntegerConstraint()(*DeviceManagementSettingIntegerConstraint) {
    m := &DeviceManagementSettingIntegerConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementSettingIntegerConstraint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementSettingIntegerConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingIntegerConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingIntegerConstraint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingIntegerConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["maximumValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumValue(val)
        }
        return nil
    }
    res["minimumValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumValue(val)
        }
        return nil
    }
    return res
}
// GetMaximumValue gets the maximumValue property value. The maximum permitted value
func (m *DeviceManagementSettingIntegerConstraint) GetMaximumValue()(*int32) {
    return m.maximumValue
}
// GetMinimumValue gets the minimumValue property value. The minimum permitted value
func (m *DeviceManagementSettingIntegerConstraint) GetMinimumValue()(*int32) {
    return m.minimumValue
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingIntegerConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("maximumValue", m.GetMaximumValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumValue", m.GetMinimumValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMaximumValue sets the maximumValue property value. The maximum permitted value
func (m *DeviceManagementSettingIntegerConstraint) SetMaximumValue(value *int32)() {
    m.maximumValue = value
}
// SetMinimumValue sets the minimumValue property value. The minimum permitted value
func (m *DeviceManagementSettingIntegerConstraint) SetMinimumValue(value *int32)() {
    m.minimumValue = value
}
