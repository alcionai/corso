package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDefenderApplicationControlSupplementalPolicyAssignment a class containing the properties used for assignment of a WindowsDefenderApplicationControl supplemental policy to a group.
type WindowsDefenderApplicationControlSupplementalPolicyAssignment struct {
    Entity
    // The target group assignment defined by the admin.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewWindowsDefenderApplicationControlSupplementalPolicyAssignment instantiates a new windowsDefenderApplicationControlSupplementalPolicyAssignment and sets the default values.
func NewWindowsDefenderApplicationControlSupplementalPolicyAssignment()(*WindowsDefenderApplicationControlSupplementalPolicyAssignment) {
    m := &WindowsDefenderApplicationControlSupplementalPolicyAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsDefenderApplicationControlSupplementalPolicyAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDefenderApplicationControlSupplementalPolicyAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDefenderApplicationControlSupplementalPolicyAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetTarget gets the target property value. The target group assignment defined by the admin.
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetTarget sets the target property value. The target group assignment defined by the admin.
func (m *WindowsDefenderApplicationControlSupplementalPolicyAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
