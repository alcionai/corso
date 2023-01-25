package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBaselineable 
type UserExperienceAnalyticsBaselineable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppHealthMetrics()(UserExperienceAnalyticsCategoryable)
    GetBatteryHealthMetrics()(UserExperienceAnalyticsCategoryable)
    GetBestPracticesMetrics()(UserExperienceAnalyticsCategoryable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceBootPerformanceMetrics()(UserExperienceAnalyticsCategoryable)
    GetDisplayName()(*string)
    GetIsBuiltIn()(*bool)
    GetRebootAnalyticsMetrics()(UserExperienceAnalyticsCategoryable)
    GetResourcePerformanceMetrics()(UserExperienceAnalyticsCategoryable)
    GetWorkFromAnywhereMetrics()(UserExperienceAnalyticsCategoryable)
    SetAppHealthMetrics(value UserExperienceAnalyticsCategoryable)()
    SetBatteryHealthMetrics(value UserExperienceAnalyticsCategoryable)()
    SetBestPracticesMetrics(value UserExperienceAnalyticsCategoryable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceBootPerformanceMetrics(value UserExperienceAnalyticsCategoryable)()
    SetDisplayName(value *string)()
    SetIsBuiltIn(value *bool)()
    SetRebootAnalyticsMetrics(value UserExperienceAnalyticsCategoryable)()
    SetResourcePerformanceMetrics(value UserExperienceAnalyticsCategoryable)()
    SetWorkFromAnywhereMetrics(value UserExperienceAnalyticsCategoryable)()
}
