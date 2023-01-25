package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthDeviceAppImpactable 
type UserExperienceAnalyticsBatteryHealthDeviceAppImpactable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDisplayName()(*string)
    GetAppName()(*string)
    GetAppPublisher()(*string)
    GetBatteryUsagePercentage()(*float64)
    GetDeviceId()(*string)
    GetIsForegroundApp()(*bool)
    SetAppDisplayName(value *string)()
    SetAppName(value *string)()
    SetAppPublisher(value *string)()
    SetBatteryUsagePercentage(value *float64)()
    SetDeviceId(value *string)()
    SetIsForegroundApp(value *bool)()
}
