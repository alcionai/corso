package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationOptionDefinitionable 
type DeviceManagementConfigurationOptionDefinitionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDependedOnBy()([]DeviceManagementConfigurationSettingDependedOnByable)
    GetDependentOn()([]DeviceManagementConfigurationDependentOnable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetHelpText()(*string)
    GetItemId()(*string)
    GetName()(*string)
    GetOdataType()(*string)
    GetOptionValue()(DeviceManagementConfigurationSettingValueable)
    SetDependedOnBy(value []DeviceManagementConfigurationSettingDependedOnByable)()
    SetDependentOn(value []DeviceManagementConfigurationDependentOnable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetHelpText(value *string)()
    SetItemId(value *string)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetOptionValue(value DeviceManagementConfigurationSettingValueable)()
}
