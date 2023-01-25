package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingEnrollmentTypeConstraint 
type DeviceManagementSettingEnrollmentTypeConstraint struct {
    DeviceManagementConstraint
    // List of enrollment types
    enrollmentTypes []string
}
// NewDeviceManagementSettingEnrollmentTypeConstraint instantiates a new DeviceManagementSettingEnrollmentTypeConstraint and sets the default values.
func NewDeviceManagementSettingEnrollmentTypeConstraint()(*DeviceManagementSettingEnrollmentTypeConstraint) {
    m := &DeviceManagementSettingEnrollmentTypeConstraint{
        DeviceManagementConstraint: *NewDeviceManagementConstraint(),
    }
    odataTypeValue := "#microsoft.graph.deviceManagementSettingEnrollmentTypeConstraint";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateDeviceManagementSettingEnrollmentTypeConstraintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementSettingEnrollmentTypeConstraintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementSettingEnrollmentTypeConstraint(), nil
}
// GetEnrollmentTypes gets the enrollmentTypes property value. List of enrollment types
func (m *DeviceManagementSettingEnrollmentTypeConstraint) GetEnrollmentTypes()([]string) {
    return m.enrollmentTypes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementSettingEnrollmentTypeConstraint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementConstraint.GetFieldDeserializers()
    res["enrollmentTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEnrollmentTypes(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceManagementSettingEnrollmentTypeConstraint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementConstraint.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEnrollmentTypes() != nil {
        err = writer.WriteCollectionOfStringValues("enrollmentTypes", m.GetEnrollmentTypes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnrollmentTypes sets the enrollmentTypes property value. List of enrollment types
func (m *DeviceManagementSettingEnrollmentTypeConstraint) SetEnrollmentTypes(value []string)() {
    m.enrollmentTypes = value
}
