package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserAppInstallStatusable 
type UserAppInstallStatusable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApp()(MobileAppable)
    GetDeviceStatuses()([]MobileAppInstallStatusable)
    GetFailedDeviceCount()(*int32)
    GetInstalledDeviceCount()(*int32)
    GetNotInstalledDeviceCount()(*int32)
    GetUserName()(*string)
    GetUserPrincipalName()(*string)
    SetApp(value MobileAppable)()
    SetDeviceStatuses(value []MobileAppInstallStatusable)()
    SetFailedDeviceCount(value *int32)()
    SetInstalledDeviceCount(value *int32)()
    SetNotInstalledDeviceCount(value *int32)()
    SetUserName(value *string)()
    SetUserPrincipalName(value *string)()
}
