package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementScriptUserStateable 
type DeviceManagementScriptUserStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceRunStates()([]DeviceManagementScriptDeviceStateable)
    GetErrorDeviceCount()(*int32)
    GetSuccessDeviceCount()(*int32)
    GetUserPrincipalName()(*string)
    SetDeviceRunStates(value []DeviceManagementScriptDeviceStateable)()
    SetErrorDeviceCount(value *int32)()
    SetSuccessDeviceCount(value *int32)()
    SetUserPrincipalName(value *string)()
}
