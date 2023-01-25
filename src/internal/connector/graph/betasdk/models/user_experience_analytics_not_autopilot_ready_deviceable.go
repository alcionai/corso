package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsNotAutopilotReadyDeviceable 
type UserExperienceAnalyticsNotAutopilotReadyDeviceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutoPilotProfileAssigned()(*bool)
    GetAutoPilotRegistered()(*bool)
    GetAzureAdJoinType()(*string)
    GetAzureAdRegistered()(*bool)
    GetDeviceName()(*string)
    GetManagedBy()(*string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetSerialNumber()(*string)
    SetAutoPilotProfileAssigned(value *bool)()
    SetAutoPilotRegistered(value *bool)()
    SetAzureAdJoinType(value *string)()
    SetAzureAdRegistered(value *bool)()
    SetDeviceName(value *string)()
    SetManagedBy(value *string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetSerialNumber(value *string)()
}
