package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsModelScoresable 
type UserExperienceAnalyticsModelScoresable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppReliabilityScore()(*float64)
    GetBatteryHealthScore()(*float64)
    GetEndpointAnalyticsScore()(*float64)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetModelDeviceCount()(*int64)
    GetStartupPerformanceScore()(*float64)
    GetWorkFromAnywhereScore()(*float64)
    SetAppReliabilityScore(value *float64)()
    SetBatteryHealthScore(value *float64)()
    SetEndpointAnalyticsScore(value *float64)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetModelDeviceCount(value *int64)()
    SetStartupPerformanceScore(value *float64)()
    SetWorkFromAnywhereScore(value *float64)()
}
