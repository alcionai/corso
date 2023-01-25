package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RoleScopeTagAutoAssignment contains the properties for auto-assigning a Role Scope Tag to a group to be applied to Devices.
type RoleScopeTagAutoAssignment struct {
    Entity
    // The auto-assignment target for the specific Role Scope Tag.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewRoleScopeTagAutoAssignment instantiates a new roleScopeTagAutoAssignment and sets the default values.
func NewRoleScopeTagAutoAssignment()(*RoleScopeTagAutoAssignment) {
    m := &RoleScopeTagAutoAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRoleScopeTagAutoAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRoleScopeTagAutoAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRoleScopeTagAutoAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RoleScopeTagAutoAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetTarget gets the target property value. The auto-assignment target for the specific Role Scope Tag.
func (m *RoleScopeTagAutoAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *RoleScopeTagAutoAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetTarget sets the target property value. The auto-assignment target for the specific Role Scope Tag.
func (m *RoleScopeTagAutoAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
