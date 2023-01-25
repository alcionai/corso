package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementScriptRunSummaryable 
type DeviceManagementScriptRunSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetErrorDeviceCount()(*int32)
    GetErrorUserCount()(*int32)
    GetSuccessDeviceCount()(*int32)
    GetSuccessUserCount()(*int32)
    SetErrorDeviceCount(value *int32)()
    SetErrorUserCount(value *int32)()
    SetSuccessDeviceCount(value *int32)()
    SetSuccessUserCount(value *int32)()
}
