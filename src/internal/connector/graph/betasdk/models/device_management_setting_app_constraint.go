package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingAppConstraint 
type DeviceManagementSettingAppConstraint struct {
    DeviceManagementConstraint
    // Acceptable app types to allow for this setting
    supportedTypes []string
}
// NewDeviceManagementSettingAppConstraint instantiates a new DeviceManagementSettingAppConstraint and sets the default values.
func NewDeviceManagementSettingAppConstraint()(*DeviceManagementSettingAppConstraint) {
    m := &DeviceManagementSettingAppConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementSettingAppConstraint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementSettingAppConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingAppConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingAppConstraint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingAppConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["supportedTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSupportedTypes(res)
        }
        return nil
    }
    return res
}
// GetSupportedTypes gets the supportedTypes property value. Acceptable app types to allow for this setting
func (m *DeviceManagementSettingAppConstraint) GetSupportedTypes()([]string) {
    return m.supportedTypes
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingAppConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSupportedTypes() != nil {
        err = writer.WriteCollectionOfStringValues("supportedTypes", m.GetSupportedTypes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetSupportedTypes sets the supportedTypes property value. Acceptable app types to allow for this setting
func (m *DeviceManagementSettingAppConstraint) SetSupportedTypes(value []string)() {
    m.supportedTypes = value
}
