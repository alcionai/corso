package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationPolicyAssignment the DeviceManagementConfigurationPolicyAssignment entity assigns a specific DeviceManagementConfigurationPolicy to an AAD group.
type DeviceManagementConfigurationPolicyAssignment struct {
    Entity
    // Represents source of assignment.
    source *DeviceAndAppManagementAssignmentSource
    // The identifier of the source of the assignment.
    sourceId *string
    // The assignment target for the DeviceManagementConfigurationPolicy.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewDeviceManagementConfigurationPolicyAssignment instantiates a new deviceManagementConfigurationPolicyAssignment and sets the default values.
func NewDeviceManagementConfigurationPolicyAssignment()(*DeviceManagementConfigurationPolicyAssignment) {
    m := &DeviceManagementConfigurationPolicyAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementConfigurationPolicyAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementConfigurationPolicyAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementConfigurationPolicyAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementConfigurationPolicyAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(*DeviceAndAppManagementAssignmentSource))
        }
        return nil
    }
    res["sourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceId(val)
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
// GetSource gets the source property value. Represents source of assignment.
func (m *DeviceManagementConfigurationPolicyAssignment) GetSource()(*DeviceAndAppManagementAssignmentSource) {
    return m.source
}
// GetSourceId gets the sourceId property value. The identifier of the source of the assignment.
func (m *DeviceManagementConfigurationPolicyAssignment) GetSourceId()(*string) {
    return m.sourceId
}
// GetTarget gets the target property value. The assignment target for the DeviceManagementConfigurationPolicy.
func (m *DeviceManagementConfigurationPolicyAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *DeviceManagementConfigurationPolicyAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sourceId", m.GetSourceId())
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
// SetSource sets the source property value. Represents source of assignment.
func (m *DeviceManagementConfigurationPolicyAssignment) SetSource(value *DeviceAndAppManagementAssignmentSource)() {
    m.source = value
}
// SetSourceId sets the sourceId property value. The identifier of the source of the assignment.
func (m *DeviceManagementConfigurationPolicyAssignment) SetSourceId(value *string)() {
    m.sourceId = value
}
// SetTarget sets the target property value. The assignment target for the DeviceManagementConfigurationPolicy.
func (m *DeviceManagementConfigurationPolicyAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
