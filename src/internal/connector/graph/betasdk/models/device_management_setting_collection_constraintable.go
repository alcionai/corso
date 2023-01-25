package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementSettingCollectionConstraintable 
type DeviceManagementSettingCollectionConstraintable interface {
    DeviceManagementConstraintable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumLength()(*int32)
    GetMinimumLength()(*int32)
    SetMaximumLength(value *int32)()
    SetMinimumLength(value *int32)()
}
