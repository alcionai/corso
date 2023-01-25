package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationSimpleSettingDefinitionable 
type DeviceManagementConfigurationSimpleSettingDefinitionable interface {
    DeviceManagementConfigurationSettingDefinitionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultValue()(DeviceManagementConfigurationSettingValueable)
    GetDependedOnBy()([]DeviceManagementConfigurationSettingDependedOnByable)
    GetDependentOn()([]DeviceManagementConfigurationDependentOnable)
    GetValueDefinition()(DeviceManagementConfigurationSettingValueDefinitionable)
    SetDefaultValue(value DeviceManagementConfigurationSettingValueable)()
    SetDependedOnBy(value []DeviceManagementConfigurationSettingDependedOnByable)()
    SetDependentOn(value []DeviceManagementConfigurationDependentOnable)()
    SetValueDefinition(value DeviceManagementConfigurationSettingValueDefinitionable)()
}
