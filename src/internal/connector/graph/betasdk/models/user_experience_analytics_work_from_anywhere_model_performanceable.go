package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable 
type UserExperienceAnalyticsWorkFromAnywhereModelPerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudIdentityScore()(*float64)
    GetCloudManagementScore()(*float64)
    GetCloudProvisioningScore()(*float64)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetModelDeviceCount()(*int32)
    GetWindowsScore()(*float64)
    GetWorkFromAnywhereScore()(*float64)
    SetCloudIdentityScore(value *float64)()
    SetCloudManagementScore(value *float64)()
    SetCloudProvisioningScore(value *float64)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetModelDeviceCount(value *int32)()
    SetWindowsScore(value *float64)()
    SetWorkFromAnywhereScore(value *float64)()
}
