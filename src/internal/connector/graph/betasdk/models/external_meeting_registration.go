package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExternalMeetingRegistration 
type ExternalMeetingRegistration struct {
    MeetingRegistrationBase
}
// NewExternalMeetingRegistration instantiates a new ExternalMeetingRegistration and sets the default values.
func NewExternalMeetingRegistration()(*ExternalMeetingRegistration) {
    m := &ExternalMeetingRegistration{
        MeetingRegistrationBase: *NewMeetingRegistrationBase(),
    }
    odataTypeValue := "#microsoft.graph.externalMeetingRegistration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateExternalMeetingRegistrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalMeetingRegistrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalMeetingRegistration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalMeetingRegistration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MeetingRegistrationBase.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ExternalMeetingRegistration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MeetingRegistrationBase.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
