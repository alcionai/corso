package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsAppHealthDeviceModelPerformanceable 
type UserExperienceAnalyticsAppHealthDeviceModelPerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActiveDeviceCount()(*int32)
    GetDeviceManufacturer()(*string)
    GetDeviceModel()(*string)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetMeanTimeToFailureInMinutes()(*int32)
    GetModelAppHealthScore()(*float64)
    GetModelAppHealthStatus()(*string)
    SetActiveDeviceCount(value *int32)()
    SetDeviceManufacturer(value *string)()
    SetDeviceModel(value *string)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetMeanTimeToFailureInMinutes(value *int32)()
    SetModelAppHealthScore(value *float64)()
    SetModelAppHealthStatus(value *string)()
}
