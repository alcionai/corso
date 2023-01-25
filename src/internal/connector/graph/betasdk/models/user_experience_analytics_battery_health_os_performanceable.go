package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthOsPerformanceable 
type UserExperienceAnalyticsBatteryHealthOsPerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveDevices()(*int32)
    GetAverageBatteryAgeInDays()(*int32)
    GetAverageEstimatedRuntimeInMinutes()(*int32)
    GetAverageMaxCapacityPercentage()(*int32)
    GetOsBuildNumber()(*string)
    GetOsVersion()(*string)
    SetActiveDevices(value *int32)()
    SetAverageBatteryAgeInDays(value *int32)()
    SetAverageEstimatedRuntimeInMinutes(value *int32)()
    SetAverageMaxCapacityPercentage(value *int32)()
    SetOsBuildNumber(value *string)()
    SetOsVersion(value *string)()
}
