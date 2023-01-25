package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmbeddedSIMActivationCodePoolAssignment the embedded SIM activation code pool assignment entity assigns a specific embeddedSIMActivationCodePool to an AAD device group.
type EmbeddedSIMActivationCodePoolAssignment struct {
    Entity
    // Base type for assignment targets.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewEmbeddedSIMActivationCodePoolAssignment instantiates a new embeddedSIMActivationCodePoolAssignment and sets the default values.
func NewEmbeddedSIMActivationCodePoolAssignment()(*EmbeddedSIMActivationCodePoolAssignment) {
    m := &EmbeddedSIMActivationCodePoolAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEmbeddedSIMActivationCodePoolAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmbeddedSIMActivationCodePoolAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmbeddedSIMActivationCodePoolAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmbeddedSIMActivationCodePoolAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetTarget gets the target property value. Base type for assignment targets.
func (m *EmbeddedSIMActivationCodePoolAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *EmbeddedSIMActivationCodePoolAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetTarget sets the target property value. Base type for assignment targets.
func (m *EmbeddedSIMActivationCodePoolAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
