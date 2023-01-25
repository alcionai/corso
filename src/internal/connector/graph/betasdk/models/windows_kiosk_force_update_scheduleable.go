package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsKioskForceUpdateScheduleable 
type WindowsKioskForceUpdateScheduleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDayofMonth()(*int32)
    GetDayofWeek()(*DayOfWeek)
    GetOdataType()(*string)
    GetRecurrence()(*Windows10AppsUpdateRecurrence)
    GetRunImmediatelyIfAfterStartDateTime()(*bool)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetDayofMonth(value *int32)()
    SetDayofWeek(value *DayOfWeek)()
    SetOdataType(value *string)()
    SetRecurrence(value *Windows10AppsUpdateRecurrence)()
    SetRunImmediatelyIfAfterStartDateTime(value *bool)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
