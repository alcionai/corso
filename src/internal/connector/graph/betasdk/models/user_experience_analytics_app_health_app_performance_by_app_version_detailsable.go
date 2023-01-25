package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable 
type UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetailsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppCrashCount()(*int32)
    GetAppDisplayName()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetAppVersion()(*string)
    GetDeviceCountWithCrashes()(*int32)
    GetIsLatestUsedVersion()(*bool)
    GetIsMostUsedVersion()(*bool)
    SetAppCrashCount(value *int32)()
    SetAppDisplayName(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetAppVersion(value *string)()
    SetDeviceCountWithCrashes(value *int32)()
    SetIsLatestUsedVersion(value *bool)()
    SetIsMostUsedVersion(value *bool)()
}
