package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthRuntimeDetailsable 
type UserExperienceAnalyticsBatteryHealthRuntimeDetailsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveDevices()(*int32)
    GetBatteryRuntimeFair()(*int32)
    GetBatteryRuntimeGood()(*int32)
    GetBatteryRuntimePoor()(*int32)
    GetLastRefreshedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetActiveDevices(value *int32)()
    SetBatteryRuntimeFair(value *int32)()
    SetBatteryRuntimeGood(value *int32)()
    SetBatteryRuntimePoor(value *int32)()
    SetLastRefreshedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
