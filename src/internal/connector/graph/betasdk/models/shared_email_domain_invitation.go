package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedEmailDomainInvitation provides operations to manage the collection of site entities.
type SharedEmailDomainInvitation struct {
    Entity
    // The expiryTime property
    expiryTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The invitationDomain property
    invitationDomain *string
    // The invitationStatus property
    invitationStatus *string
}
// NewSharedEmailDomainInvitation instantiates a new sharedEmailDomainInvitation and sets the default values.
func NewSharedEmailDomainInvitation()(*SharedEmailDomainInvitation) {
    m := &SharedEmailDomainInvitation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSharedEmailDomainInvitationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedEmailDomainInvitationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharedEmailDomainInvitation(), nil
}
// GetExpiryTime gets the expiryTime property value. The expiryTime property
func (m *SharedEmailDomainInvitation) GetExpiryTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expiryTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedEmailDomainInvitation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expiryTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiryTime(val)
        }
        return nil
    }
    res["invitationDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitationDomain(val)
        }
        return nil
    }
    res["invitationStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitationStatus(val)
        }
        return nil
    }
    return res
}
// GetInvitationDomain gets the invitationDomain property value. The invitationDomain property
func (m *SharedEmailDomainInvitation) GetInvitationDomain()(*string) {
    return m.invitationDomain
}
// GetInvitationStatus gets the invitationStatus property value. The invitationStatus property
func (m *SharedEmailDomainInvitation) GetInvitationStatus()(*string) {
    return m.invitationStatus
}
// Serialize serializes information the current object
func (m *SharedEmailDomainInvitation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expiryTime", m.GetExpiryTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invitationDomain", m.GetInvitationDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("invitationStatus", m.GetInvitationStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpiryTime sets the expiryTime property value. The expiryTime property
func (m *SharedEmailDomainInvitation) SetExpiryTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiryTime = value
}
// SetInvitationDomain sets the invitationDomain property value. The invitationDomain property
func (m *SharedEmailDomainInvitation) SetInvitationDomain(value *string)() {
    m.invitationDomain = value
}
// SetInvitationStatus sets the invitationStatus property value. The invitationStatus property
func (m *SharedEmailDomainInvitation) SetInvitationStatus(value *string)() {
    m.invitationStatus = value
}
