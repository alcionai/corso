package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsImpactingProcessable 
type UserExperienceAnalyticsImpactingProcessable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategory()(*string)
    GetDescription()(*string)
    GetDeviceId()(*string)
    GetImpactValue()(*float64)
    GetProcessName()(*string)
    GetPublisher()(*string)
    SetCategory(value *string)()
    SetDescription(value *string)()
    SetDeviceId(value *string)()
    SetImpactValue(value *float64)()
    SetProcessName(value *string)()
    SetPublisher(value *string)()
}
