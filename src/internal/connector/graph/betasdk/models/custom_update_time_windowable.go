package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CustomUpdateTimeWindowable 
type CustomUpdateTimeWindowable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEndDay()(*DayOfWeek)
    GetEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetOdataType()(*string)
    GetStartDay()(*DayOfWeek)
    GetStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    SetEndDay(value *DayOfWeek)()
    SetEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetOdataType(value *string)()
    SetStartDay(value *DayOfWeek)()
    SetStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
}
