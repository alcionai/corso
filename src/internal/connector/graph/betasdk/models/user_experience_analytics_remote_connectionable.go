package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsRemoteConnectionable 
type UserExperienceAnalyticsRemoteConnectionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudPcFailurePercentage()(*float64)
    GetCloudPcRoundTripTime()(*float64)
    GetCloudPcSignInTime()(*float64)
    GetCoreBootTime()(*float64)
    GetCoreSignInTime()(*float64)
    GetDeviceCount()(*int32)
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetRemoteSignInTime()(*float64)
    GetUserPrincipalName()(*string)
    GetVirtualNetwork()(*string)
    SetCloudPcFailurePercentage(value *float64)()
    SetCloudPcRoundTripTime(value *float64)()
    SetCloudPcSignInTime(value *float64)()
    SetCoreBootTime(value *float64)()
    SetCoreSignInTime(value *float64)()
    SetDeviceCount(value *int32)()
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetRemoteSignInTime(value *float64)()
    SetUserPrincipalName(value *string)()
    SetVirtualNetwork(value *string)()
}
