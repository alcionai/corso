package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsQualityUpdateProfileAssignment this entity contains the properties used to assign a windows quality update profile to a group.
type WindowsQualityUpdateProfileAssignment struct {
    Entity
    // The assignment target that the quality update profile is assigned to.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewWindowsQualityUpdateProfileAssignment instantiates a new windowsQualityUpdateProfileAssignment and sets the default values.
func NewWindowsQualityUpdateProfileAssignment()(*WindowsQualityUpdateProfileAssignment) {
    m := &WindowsQualityUpdateProfileAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsQualityUpdateProfileAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsQualityUpdateProfileAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsQualityUpdateProfileAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsQualityUpdateProfileAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceAndAppManagementAssignmentTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(DeviceAndAppManagementAssignmentTargetable))
        }
        return nil
    }
    return res
}
// GetTarget gets the target property value. The assignment target that the quality update profile is assigned to.
func (m *WindowsQualityUpdateProfileAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *WindowsQualityUpdateProfileAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTarget sets the target property value. The assignment target that the quality update profile is assigned to.
func (m *WindowsQualityUpdateProfileAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
