package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeCardEntryable 
type TimeCardEntryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBreaks()([]TimeCardBreakable)
    GetClockInEvent()(TimeCardEventable)
    GetClockOutEvent()(TimeCardEventable)
    GetOdataType()(*string)
    SetBreaks(value []TimeCardBreakable)()
    SetClockInEvent(value TimeCardEventable)()
    SetClockOutEvent(value TimeCardEventable)()
    SetOdataType(value *string)()
}
