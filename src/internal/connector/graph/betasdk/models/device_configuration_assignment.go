package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationAssignment the device configuration assignment entity assigns an AAD group to a specific device configuration.
type DeviceConfigurationAssignment struct {
    Entity
    // The admin intent to apply or remove the profile. Possible values are: apply, remove.
    intent *DeviceConfigAssignmentIntent
    // Represents source of assignment.
    source *DeviceAndAppManagementAssignmentSource
    // The identifier of the source of the assignment. This property is read-only.
    sourceId *string
    // The assignment target for the device configuration.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewDeviceConfigurationAssignment instantiates a new deviceConfigurationAssignment and sets the default values.
func NewDeviceConfigurationAssignment()(*DeviceConfigurationAssignment) {
    m := &DeviceConfigurationAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["intent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceConfigAssignmentIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntent(val.(*DeviceConfigAssignmentIntent))
        }
        return nil
    }
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
// GetIntent gets the intent property value. The admin intent to apply or remove the profile. Possible values are: apply, remove.
func (m *DeviceConfigurationAssignment) GetIntent()(*DeviceConfigAssignmentIntent) {
    return m.intent
}
// GetSource gets the source property value. Represents source of assignment.
func (m *DeviceConfigurationAssignment) GetSource()(*DeviceAndAppManagementAssignmentSource) {
    return m.source
}
// GetSourceId gets the sourceId property value. The identifier of the source of the assignment. This property is read-only.
func (m *DeviceConfigurationAssignment) GetSourceId()(*string) {
    return m.sourceId
}
// GetTarget gets the target property value. The assignment target for the device configuration.
func (m *DeviceConfigurationAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *DeviceConfigurationAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIntent() != nil {
        cast := (*m.GetIntent()).String()
        err = writer.WriteStringValue("intent", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSource() != nil {
        cast := (*m.GetSource()).String()
        err = writer.WriteStringValue("source", &cast)
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
// SetIntent sets the intent property value. The admin intent to apply or remove the profile. Possible values are: apply, remove.
func (m *DeviceConfigurationAssignment) SetIntent(value *DeviceConfigAssignmentIntent)() {
    m.intent = value
}
// SetSource sets the source property value. Represents source of assignment.
func (m *DeviceConfigurationAssignment) SetSource(value *DeviceAndAppManagementAssignmentSource)() {
    m.source = value
}
// SetSourceId sets the sourceId property value. The identifier of the source of the assignment. This property is read-only.
func (m *DeviceConfigurationAssignment) SetSourceId(value *string)() {
    m.sourceId = value
}
// SetTarget sets the target property value. The assignment target for the device configuration.
func (m *DeviceConfigurationAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
