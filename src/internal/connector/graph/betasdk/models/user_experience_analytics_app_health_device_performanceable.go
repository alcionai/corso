package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthDevicePerformanceable 
type UserExperienceAnalyticsAppHealthDevicePerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppCrashCount()(*int32)
    GetAppHangCount()(*int32)
    GetCrashedAppCount()(*int32)
    GetDeviceAppHealthScore()(*float64)
    GetDeviceAppHealthStatus()(*string)
    GetDeviceDisplayName()(*string)
    GetDeviceId()(*string)
    GetDeviceManufacturer()(*string)
    GetDeviceModel()(*string)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetMeanTimeToFailureInMinutes()(*int32)
    GetProcessedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetAppCrashCount(value *int32)()
    SetAppHangCount(value *int32)()
    SetCrashedAppCount(value *int32)()
    SetDeviceAppHealthScore(value *float64)()
    SetDeviceAppHealthStatus(value *string)()
    SetDeviceDisplayName(value *string)()
    SetDeviceId(value *string)()
    SetDeviceManufacturer(value *string)()
    SetDeviceModel(value *string)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetMeanTimeToFailureInMinutes(value *int32)()
    SetProcessedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
