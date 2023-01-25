package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingCollectionDefinitionable 
type DeviceManagementConfigurationSimpleSettingCollectionDefinitionable interface {
    DeviceManagementConfigurationSimpleSettingDefinitionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMaximumCount()(*int32)
    GetMinimumCount()(*int32)
    SetMaximumCount(value *int32)()
    SetMinimumCount(value *int32)()
}
