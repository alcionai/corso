package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthDevicePerformanceable 
type UserExperienceAnalyticsBatteryHealthDevicePerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBatteryAgeInDays()(*int32)
    GetDeviceBatteryHealthScore()(*int32)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetEstimatedRuntimeInMinutes()(*int32)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetManufacturer()(*string)
    GetMaxCapacityPercentage()(*int32)
    GetModel()(*string)
    SetBatteryAgeInDays(value *int32)()
    SetDeviceBatteryHealthScore(value *int32)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetEstimatedRuntimeInMinutes(value *int32)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetManufacturer(value *string)()
    SetMaxCapacityPercentage(value *int32)()
    SetModel(value *string)()
}
