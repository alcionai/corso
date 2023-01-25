package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrivilegedSignupStatus 
type PrivilegedSignupStatus struct {
    Entity
    // The isRegistered property
    isRegistered *bool
    // The status property
    status *SetupStatus
}
// NewPrivilegedSignupStatus instantiates a new PrivilegedSignupStatus and sets the default values.
func NewPrivilegedSignupStatus()(*PrivilegedSignupStatus) {
    m := &PrivilegedSignupStatus{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrivilegedSignupStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrivilegedSignupStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrivilegedSignupStatus(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrivilegedSignupStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isRegistered"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRegistered(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSetupStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SetupStatus))
        }
        return nil
    }
    return res
}
// GetIsRegistered gets the isRegistered property value. The isRegistered property
func (m *PrivilegedSignupStatus) GetIsRegistered()(*bool) {
    return m.isRegistered
}
// GetStatus gets the status property value. The status property
func (m *PrivilegedSignupStatus) GetStatus()(*SetupStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *PrivilegedSignupStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isRegistered", m.GetIsRegistered())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsRegistered sets the isRegistered property value. The isRegistered property
func (m *PrivilegedSignupStatus) SetIsRegistered(value *bool)() {
    m.isRegistered = value
}
// SetStatus sets the status property value. The status property
func (m *PrivilegedSignupStatus) SetStatus(value *SetupStatus)() {
    m.status = value
}
