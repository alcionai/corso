package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingRegistrationable 
type MeetingRegistrationable interface {
    MeetingRegistrationBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomQuestions()([]MeetingRegistrationQuestionable)
    GetDescription()(*string)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRegistrationPageViewCount()(*int32)
    GetRegistrationPageWebUrl()(*string)
    GetSpeakers()([]MeetingSpeakerable)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubject()(*string)
    SetCustomQuestions(value []MeetingRegistrationQuestionable)()
    SetDescription(value *string)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRegistrationPageViewCount(value *int32)()
    SetRegistrationPageWebUrl(value *string)()
    SetSpeakers(value []MeetingSpeakerable)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubject(value *string)()
}
