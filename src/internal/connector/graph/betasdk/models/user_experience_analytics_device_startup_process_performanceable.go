package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDeviceStartupProcessPerformanceable 
type UserExperienceAnalyticsDeviceStartupProcessPerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceCount()(*int64)
    GetMedianImpactInMs()(*int32)
    GetMedianImpactInMs2()(*int64)
    GetProcessName()(*string)
    GetProductName()(*string)
    GetPublisher()(*string)
    GetTotalImpactInMs()(*int32)
    GetTotalImpactInMs2()(*int64)
    SetDeviceCount(value *int64)()
    SetMedianImpactInMs(value *int32)()
    SetMedianImpactInMs2(value *int64)()
    SetProcessName(value *string)()
    SetProductName(value *string)()
    SetPublisher(value *string)()
    SetTotalImpactInMs(value *int32)()
    SetTotalImpactInMs2(value *int64)()
}
