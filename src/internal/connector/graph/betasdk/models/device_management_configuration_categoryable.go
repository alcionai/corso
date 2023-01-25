package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationCategoryable 
type DeviceManagementConfigurationCategoryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategoryDescription()(*string)
    GetChildCategoryIds()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetHelpText()(*string)
    GetName()(*string)
    GetParentCategoryId()(*string)
    GetPlatforms()(*DeviceManagementConfigurationPlatforms)
    GetRootCategoryId()(*string)
    GetSettingUsage()(*DeviceManagementConfigurationSettingUsage)
    GetTechnologies()(*DeviceManagementConfigurationTechnologies)
    SetCategoryDescription(value *string)()
    SetChildCategoryIds(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetHelpText(value *string)()
    SetName(value *string)()
    SetParentCategoryId(value *string)()
    SetPlatforms(value *DeviceManagementConfigurationPlatforms)()
    SetRootCategoryId(value *string)()
    SetSettingUsage(value *DeviceManagementConfigurationSettingUsage)()
    SetTechnologies(value *DeviceManagementConfigurationTechnologies)()
}
