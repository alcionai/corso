package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingXmlConstraint 
type DeviceManagementSettingXmlConstraint struct {
    DeviceManagementConstraint
}
// NewDeviceManagementSettingXmlConstraint instantiates a new DeviceManagementSettingXmlConstraint and sets the default values.
func NewDeviceManagementSettingXmlConstraint()(*DeviceManagementSettingXmlConstraint) {
    m := &DeviceManagementSettingXmlConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementSettingXmlConstraint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementSettingXmlConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingXmlConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingXmlConstraint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingXmlConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingXmlConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
