package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSettingGroupDefinitionable 
type DeviceManagementConfigurationSettingGroupDefinitionable interface {
    DeviceManagementConfigurationSettingDefinitionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildIds()([]string)
    GetDependedOnBy()([]DeviceManagementConfigurationSettingDependedOnByable)
    GetDependentOn()([]DeviceManagementConfigurationDependentOnable)
    SetChildIds(value []string)()
    SetDependedOnBy(value []DeviceManagementConfigurationSettingDependedOnByable)()
    SetDependentOn(value []DeviceManagementConfigurationDependentOnable)()
}
