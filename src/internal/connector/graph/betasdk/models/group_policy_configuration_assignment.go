package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupPolicyConfigurationAssignment the group policy configuration assignment entity assigns one or more AAD groups to a specific group policy configuration.
type GroupPolicyConfigurationAssignment struct {
    Entity
    // The date and time the entity was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The type of groups targeted the group policy configuration.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewGroupPolicyConfigurationAssignment instantiates a new groupPolicyConfigurationAssignment and sets the default values.
func NewGroupPolicyConfigurationAssignment()(*GroupPolicyConfigurationAssignment) {
    m := &GroupPolicyConfigurationAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupPolicyConfigurationAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupPolicyConfigurationAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupPolicyConfigurationAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupPolicyConfigurationAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
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
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyConfigurationAssignment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetTarget gets the target property value. The type of groups targeted the group policy configuration.
func (m *GroupPolicyConfigurationAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *GroupPolicyConfigurationAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time the entity was last modified.
func (m *GroupPolicyConfigurationAssignment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetTarget sets the target property value. The type of groups targeted the group policy configuration.
func (m *GroupPolicyConfigurationAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
