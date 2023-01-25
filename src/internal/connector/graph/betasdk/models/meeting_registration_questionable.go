package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingRegistrationQuestionable 
type MeetingRegistrationQuestionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnswerInputType()(*AnswerInputType)
    GetAnswerOptions()([]string)
    GetDisplayName()(*string)
    GetIsRequired()(*bool)
    SetAnswerInputType(value *AnswerInputType)()
    SetAnswerOptions(value []string)()
    SetDisplayName(value *string)()
    SetIsRequired(value *bool)()
}
