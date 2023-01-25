package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingParticipants 
type MeetingParticipants struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Information of the meeting attendees.
    attendees []MeetingParticipantInfoable
    // The contributors property
    contributors []MeetingParticipantInfoable
    // The OdataType property
    odataType *string
    // Information of the meeting organizer.
    organizer MeetingParticipantInfoable
    // The producers property
    producers []MeetingParticipantInfoable
}
// NewMeetingParticipants instantiates a new meetingParticipants and sets the default values.
func NewMeetingParticipants()(*MeetingParticipants) {
    m := &MeetingParticipants{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeetingParticipantsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingParticipantsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingParticipants(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingParticipants) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAttendees gets the attendees property value. Information of the meeting attendees.
func (m *MeetingParticipants) GetAttendees()([]MeetingParticipantInfoable) {
    return m.attendees
}
// GetContributors gets the contributors property value. The contributors property
func (m *MeetingParticipants) GetContributors()([]MeetingParticipantInfoable) {
    return m.contributors
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingParticipants) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attendees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingParticipantInfoable, len(val))
            for i, v := range val {
                res[i] = v.(MeetingParticipantInfoable)
            }
            m.SetAttendees(res)
        }
        return nil
    }
    res["contributors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingParticipantInfoable, len(val))
            for i, v := range val {
                res[i] = v.(MeetingParticipantInfoable)
            }
            m.SetContributors(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["organizer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMeetingParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizer(val.(MeetingParticipantInfoable))
        }
        return nil
    }
    res["producers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingParticipantInfoable, len(val))
            for i, v := range val {
                res[i] = v.(MeetingParticipantInfoable)
            }
            m.SetProducers(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MeetingParticipants) GetOdataType()(*string) {
    return m.odataType
}
// GetOrganizer gets the organizer property value. Information of the meeting organizer.
func (m *MeetingParticipants) GetOrganizer()(MeetingParticipantInfoable) {
    return m.organizer
}
// GetProducers gets the producers property value. The producers property
func (m *MeetingParticipants) GetProducers()([]MeetingParticipantInfoable) {
    return m.producers
}
// Serialize serializes information the current object
func (m *MeetingParticipants) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttendees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("attendees", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContributors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContributors()))
        for i, v := range m.GetContributors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("contributors", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("organizer", m.GetOrganizer())
        if err != nil {
            return err
        }
    }
    if m.GetProducers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProducers()))
        for i, v := range m.GetProducers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("producers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingParticipants) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAttendees sets the attendees property value. Information of the meeting attendees.
func (m *MeetingParticipants) SetAttendees(value []MeetingParticipantInfoable)() {
    m.attendees = value
}
// SetContributors sets the contributors property value. The contributors property
func (m *MeetingParticipants) SetContributors(value []MeetingParticipantInfoable)() {
    m.contributors = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MeetingParticipants) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrganizer sets the organizer property value. Information of the meeting organizer.
func (m *MeetingParticipants) SetOrganizer(value MeetingParticipantInfoable)() {
    m.organizer = value
}
// SetProducers sets the producers property value. The producers property
func (m *MeetingParticipants) SetProducers(value []MeetingParticipantInfoable)() {
    m.producers = value
}
