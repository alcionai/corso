package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementScriptGroupAssignment contains properties used to assign a device management script to a group.
type DeviceManagementScriptGroupAssignment struct {
    Entity
    // The Id of the Azure Active Directory group we are targeting the script to.
    targetGroupId *string
}
// NewDeviceManagementScriptGroupAssignment instantiates a new deviceManagementScriptGroupAssignment and sets the default values.
func NewDeviceManagementScriptGroupAssignment()(*DeviceManagementScriptGroupAssignment) {
    m := &DeviceManagementScriptGroupAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementScriptGroupAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementScriptGroupAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementScriptGroupAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementScriptGroupAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["targetGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetGroupId(val)
        }
        return nil
    }
    return res
}
// GetTargetGroupId gets the targetGroupId property value. The Id of the Azure Active Directory group we are targeting the script to.
func (m *DeviceManagementScriptGroupAssignment) GetTargetGroupId()(*string) {
    return m.targetGroupId
}
// Serialize serializes information the current object
func (m *DeviceManagementScriptGroupAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetGroupId", m.GetTargetGroupId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTargetGroupId sets the targetGroupId property value. The Id of the Azure Active Directory group we are targeting the script to.
func (m *DeviceManagementScriptGroupAssignment) SetTargetGroupId(value *string)() {
    m.targetGroupId = value
}
