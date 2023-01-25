package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementScriptAssignment contains properties used to assign a device management script to a group.
type DeviceManagementScriptAssignment struct {
    Entity
    // The Id of the Azure Active Directory group we are targeting the script to.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewDeviceManagementScriptAssignment instantiates a new deviceManagementScriptAssignment and sets the default values.
func NewDeviceManagementScriptAssignment()(*DeviceManagementScriptAssignment) {
    m := &DeviceManagementScriptAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementScriptAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementScriptAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementScriptAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementScriptAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetTarget gets the target property value. The Id of the Azure Active Directory group we are targeting the script to.
func (m *DeviceManagementScriptAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetTarget sets the target property value. The Id of the Azure Active Directory group we are targeting the script to.
func (m *DeviceManagementScriptAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
