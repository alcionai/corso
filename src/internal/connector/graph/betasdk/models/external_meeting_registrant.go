package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalMeetingRegistrant 
type ExternalMeetingRegistrant struct {
    MeetingRegistrantBase
    // The tenant ID of this registrant if in Azure Active Directory.
    tenantId *string
    // The user ID of this registrant if in Azure Active Directory.
    userId *string
}
// NewExternalMeetingRegistrant instantiates a new ExternalMeetingRegistrant and sets the default values.
func NewExternalMeetingRegistrant()(*ExternalMeetingRegistrant) {
    m := &ExternalMeetingRegistrant{
        MeetingRegistrantBase: *NewMeetingRegistrantBase(),
    }
    odataTypeValue := "#microsoft.graph.externalMeetingRegistrant";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExternalMeetingRegistrantFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalMeetingRegistrantFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalMeetingRegistrant(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalMeetingRegistrant) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MeetingRegistrantBase.GetFieldDeserializers()
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetTenantId gets the tenantId property value. The tenant ID of this registrant if in Azure Active Directory.
func (m *ExternalMeetingRegistrant) GetTenantId()(*string) {
    return m.tenantId
}
// GetUserId gets the userId property value. The user ID of this registrant if in Azure Active Directory.
func (m *ExternalMeetingRegistrant) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *ExternalMeetingRegistrant) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MeetingRegistrantBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTenantId sets the tenantId property value. The tenant ID of this registrant if in Azure Active Directory.
func (m *ExternalMeetingRegistrant) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetUserId sets the userId property value. The user ID of this registrant if in Azure Active Directory.
func (m *ExternalMeetingRegistrant) SetUserId(value *string)() {
    m.userId = value
}
