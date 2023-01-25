package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingIntegerConstraintable 
type DeviceManagementSettingIntegerConstraintable interface {
    DeviceManagementConstraintable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumValue()(*int32)
    GetMinimumValue()(*int32)
    SetMaximumValue(value *int32)()
    SetMinimumValue(value *int32)()
}
