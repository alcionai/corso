package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleEnrollmentProfileAssignment an assignment of an Apple profile.
type AppleEnrollmentProfileAssignment struct {
    Entity
    // The assignment target for the Apple user initiated deployment profile.
    target DeviceAndAppManagementAssignmentTargetable
}
// NewAppleEnrollmentProfileAssignment instantiates a new appleEnrollmentProfileAssignment and sets the default values.
func NewAppleEnrollmentProfileAssignment()(*AppleEnrollmentProfileAssignment) {
    m := &AppleEnrollmentProfileAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAppleEnrollmentProfileAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleEnrollmentProfileAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppleEnrollmentProfileAssignment(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleEnrollmentProfileAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetTarget gets the target property value. The assignment target for the Apple user initiated deployment profile.
func (m *AppleEnrollmentProfileAssignment) GetTarget()(DeviceAndAppManagementAssignmentTargetable) {
    return m.target
}
// Serialize serializes information the current object
func (m *AppleEnrollmentProfileAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetTarget sets the target property value. The assignment target for the Apple user initiated deployment profile.
func (m *AppleEnrollmentProfileAssignment) SetTarget(value DeviceAndAppManagementAssignmentTargetable)() {
    m.target = value
}
