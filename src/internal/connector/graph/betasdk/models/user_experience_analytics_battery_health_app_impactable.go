package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthAppImpactable 
type UserExperienceAnalyticsBatteryHealthAppImpactable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveDevices()(*int32)
    GetAppDisplayName()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetBatteryUsagePercentage()(*float64)
    GetIsForegroundApp()(*bool)
    SetActiveDevices(value *int32)()
    SetAppDisplayName(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetBatteryUsagePercentage(value *float64)()
    SetIsForegroundApp(value *bool)()
}
