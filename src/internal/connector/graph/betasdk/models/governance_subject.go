package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GovernanceSubject 
type GovernanceSubject struct {
    Entity
    // The display name of the subject.
    displayName *string
    // The email address of the user subject. If the subject is in other types, it is empty.
    email *string
    // The principal name of the user subject. If the subject is in other types, it is empty.
    principalName *string
    // The type of the subject. The value can be User, Group, and ServicePrincipal.
    type_escaped *string
}
// NewGovernanceSubject instantiates a new governanceSubject and sets the default values.
func NewGovernanceSubject()(*GovernanceSubject) {
    m := &GovernanceSubject{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGovernanceSubjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGovernanceSubjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGovernanceSubject(), nil
}
// GetDisplayName gets the displayName property value. The display name of the subject.
func (m *GovernanceSubject) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. The email address of the user subject. If the subject is in other types, it is empty.
func (m *GovernanceSubject) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GovernanceSubject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["principalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalName(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
// GetPrincipalName gets the principalName property value. The principal name of the user subject. If the subject is in other types, it is empty.
func (m *GovernanceSubject) GetPrincipalName()(*string) {
    return m.principalName
}
// GetType gets the type property value. The type of the subject. The value can be User, Group, and ServicePrincipal.
func (m *GovernanceSubject) GetType()(*string) {
    return m.type_escaped
}
// Serialize serializes information the current object
func (m *GovernanceSubject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalName", m.GetPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name of the subject.
func (m *GovernanceSubject) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. The email address of the user subject. If the subject is in other types, it is empty.
func (m *GovernanceSubject) SetEmail(value *string)() {
    m.email = value
}
// SetPrincipalName sets the principalName property value. The principal name of the user subject. If the subject is in other types, it is empty.
func (m *GovernanceSubject) SetPrincipalName(value *string)() {
    m.principalName = value
}
// SetType sets the type property value. The type of the subject. The value can be User, Group, and ServicePrincipal.
func (m *GovernanceSubject) SetType(value *string)() {
    m.type_escaped = value
}
