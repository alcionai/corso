package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingRegistration 
type MeetingRegistration struct {
    MeetingRegistrationBase
    // Custom registration questions.
    customQuestions []MeetingRegistrationQuestionable
    // The description of the meeting.
    description *string
    // The meeting end time in UTC.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The number of times the registration page has been visited. Read-only.
    registrationPageViewCount *int32
    // The URL of the registration page. Read-only.
    registrationPageWebUrl *string
    // The meeting speaker's information.
    speakers []MeetingSpeakerable
    // The meeting start time in UTC.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject of the meeting.
    subject *string
}
// NewMeetingRegistration instantiates a new MeetingRegistration and sets the default values.
func NewMeetingRegistration()(*MeetingRegistration) {
    m := &MeetingRegistration{
        MeetingRegistrationBase: *NewMeetingRegistrationBase(),
    }
    odataTypeValue := "#microsoft.graph.meetingRegistration";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMeetingRegistrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingRegistrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingRegistration(), nil
}
// GetCustomQuestions gets the customQuestions property value. Custom registration questions.
func (m *MeetingRegistration) GetCustomQuestions()([]MeetingRegistrationQuestionable) {
    return m.customQuestions
}
// GetDescription gets the description property value. The description of the meeting.
func (m *MeetingRegistration) GetDescription()(*string) {
    return m.description
}
// GetEndDateTime gets the endDateTime property value. The meeting end time in UTC.
func (m *MeetingRegistration) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingRegistration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MeetingRegistrationBase.GetFieldDeserializers()
    res["customQuestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingRegistrationQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingRegistrationQuestionable, len(val))
            for i, v := range val {
                res[i] = v.(MeetingRegistrationQuestionable)
            }
            m.SetCustomQuestions(res)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["registrationPageViewCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationPageViewCount(val)
        }
        return nil
    }
    res["registrationPageWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegistrationPageWebUrl(val)
        }
        return nil
    }
    res["speakers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingSpeakerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingSpeakerable, len(val))
            for i, v := range val {
                res[i] = v.(MeetingSpeakerable)
            }
            m.SetSpeakers(res)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
// GetRegistrationPageViewCount gets the registrationPageViewCount property value. The number of times the registration page has been visited. Read-only.
func (m *MeetingRegistration) GetRegistrationPageViewCount()(*int32) {
    return m.registrationPageViewCount
}
// GetRegistrationPageWebUrl gets the registrationPageWebUrl property value. The URL of the registration page. Read-only.
func (m *MeetingRegistration) GetRegistrationPageWebUrl()(*string) {
    return m.registrationPageWebUrl
}
// GetSpeakers gets the speakers property value. The meeting speaker's information.
func (m *MeetingRegistration) GetSpeakers()([]MeetingSpeakerable) {
    return m.speakers
}
// GetStartDateTime gets the startDateTime property value. The meeting start time in UTC.
func (m *MeetingRegistration) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetSubject gets the subject property value. The subject of the meeting.
func (m *MeetingRegistration) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *MeetingRegistration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MeetingRegistrationBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCustomQuestions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomQuestions()))
        for i, v := range m.GetCustomQuestions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customQuestions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("registrationPageViewCount", m.GetRegistrationPageViewCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("registrationPageWebUrl", m.GetRegistrationPageWebUrl())
        if err != nil {
            return err
        }
    }
    if m.GetSpeakers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSpeakers()))
        for i, v := range m.GetSpeakers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("speakers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCustomQuestions sets the customQuestions property value. Custom registration questions.
func (m *MeetingRegistration) SetCustomQuestions(value []MeetingRegistrationQuestionable)() {
    m.customQuestions = value
}
// SetDescription sets the description property value. The description of the meeting.
func (m *MeetingRegistration) SetDescription(value *string)() {
    m.description = value
}
// SetEndDateTime sets the endDateTime property value. The meeting end time in UTC.
func (m *MeetingRegistration) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetRegistrationPageViewCount sets the registrationPageViewCount property value. The number of times the registration page has been visited. Read-only.
func (m *MeetingRegistration) SetRegistrationPageViewCount(value *int32)() {
    m.registrationPageViewCount = value
}
// SetRegistrationPageWebUrl sets the registrationPageWebUrl property value. The URL of the registration page. Read-only.
func (m *MeetingRegistration) SetRegistrationPageWebUrl(value *string)() {
    m.registrationPageWebUrl = value
}
// SetSpeakers sets the speakers property value. The meeting speaker's information.
func (m *MeetingRegistration) SetSpeakers(value []MeetingSpeakerable)() {
    m.speakers = value
}
// SetStartDateTime sets the startDateTime property value. The meeting start time in UTC.
func (m *MeetingRegistration) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetSubject sets the subject property value. The subject of the meeting.
func (m *MeetingRegistration) SetSubject(value *string)() {
    m.subject = value
}
