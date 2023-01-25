package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeOffable 
type TimeOffable interface {
    ChangeTrackedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDraftTimeOff()(TimeOffItemable)
    GetIsStagedForDeletion()(*bool)
    GetSharedTimeOff()(TimeOffItemable)
    GetUserId()(*string)
    SetDraftTimeOff(value TimeOffItemable)()
    SetIsStagedForDeletion(value *bool)()
    SetSharedTimeOff(value TimeOffItemable)()
    SetUserId(value *string)()
}
