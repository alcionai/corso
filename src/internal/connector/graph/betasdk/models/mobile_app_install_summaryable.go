package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppInstallSummaryable 
type MobileAppInstallSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFailedDeviceCount()(*int32)
    GetFailedUserCount()(*int32)
    GetInstalledDeviceCount()(*int32)
    GetInstalledUserCount()(*int32)
    GetNotApplicableDeviceCount()(*int32)
    GetNotApplicableUserCount()(*int32)
    GetNotInstalledDeviceCount()(*int32)
    GetNotInstalledUserCount()(*int32)
    GetPendingInstallDeviceCount()(*int32)
    GetPendingInstallUserCount()(*int32)
    SetFailedDeviceCount(value *int32)()
    SetFailedUserCount(value *int32)()
    SetInstalledDeviceCount(value *int32)()
    SetInstalledUserCount(value *int32)()
    SetNotApplicableDeviceCount(value *int32)()
    SetNotApplicableUserCount(value *int32)()
    SetNotInstalledDeviceCount(value *int32)()
    SetNotInstalledUserCount(value *int32)()
    SetPendingInstallDeviceCount(value *int32)()
    SetPendingInstallUserCount(value *int32)()
}
