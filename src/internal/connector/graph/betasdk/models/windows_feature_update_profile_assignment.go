package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsFeatureUpdateProfileAssignment this entity contains the properties used to assign a windows feature update profile to a group.
type WindowsFeatureUpdateProfileAssignment struct {
    Entity
    // The assignment target that the feature update profile is assigned to.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewWindowsFeatureUpdateProfileAssignment instantiates a new windowsFeatureUpdateProfileAssignment and sets the default values.
func NewWindowsFeatureUpdateProfileAssignment()(*WindowsFeatureUpdateProfileAssignment) {
    m := &WindowsFeatureUpdateProfileAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsFeatureUpdateProfileAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsFeatureUpdateProfileAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsFeatureUpdateProfileAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsFeatureUpdateProfileAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetTarget gets the target property value. The assignment target that the feature update profile is assigned to.
func (m *WindowsFeatureUpdateProfileAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *WindowsFeatureUpdateProfileAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetTarget sets the target property value. The assignment target that the feature update profile is assigned to.
func (m *WindowsFeatureUpdateProfileAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
