package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable 
type UserExperienceAnalyticsAppHealthAppPerformanceByOSVersionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveDeviceCount()(*int32)
    GetAppCrashCount()(*int32)
    GetAppDisplayName()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetAppUsageDuration()(*int32)
    GetMeanTimeToFailureInMinutes()(*int32)
    GetOsBuildNumber()(*string)
    GetOsVersion()(*string)
    SetActiveDeviceCount(value *int32)()
    SetAppCrashCount(value *int32)()
    SetAppDisplayName(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetAppUsageDuration(value *int32)()
    SetMeanTimeToFailureInMinutes(value *int32)()
    SetOsBuildNumber(value *string)()
    SetOsVersion(value *string)()
}
