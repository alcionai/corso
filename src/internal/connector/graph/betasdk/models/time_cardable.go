package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeCardable 
type TimeCardable interface {
    ChangeTrackedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBreaks()([]TimeCardBreakable)
    GetClockInEvent()(TimeCardEventable)
    GetClockOutEvent()(TimeCardEventable)
    GetConfirmedBy()(*ConfirmedBy)
    GetNotes()(ItemBodyable)
    GetOriginalEntry()(TimeCardEntryable)
    GetState()(*TimeCardState)
    GetUserId()(*string)
    SetBreaks(value []TimeCardBreakable)()
    SetClockInEvent(value TimeCardEventable)()
    SetClockOutEvent(value TimeCardEventable)()
    SetConfirmedBy(value *ConfirmedBy)()
    SetNotes(value ItemBodyable)()
    SetOriginalEntry(value TimeCardEntryable)()
    SetState(value *TimeCardState)()
    SetUserId(value *string)()
}
