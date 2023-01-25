package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable 
type UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceId()(*string)
    GetEstimatedRuntimeInMinutes()(*int32)
    GetRuntimeDateTime()(*string)
    SetDeviceId(value *string)()
    SetEstimatedRuntimeInMinutes(value *int32)()
    SetRuntimeDateTime(value *string)()
}
