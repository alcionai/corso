package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationPolicyTemplateable 
type DeviceManagementConfigurationPolicyTemplateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowUnmanagedSettings()(*bool)
    GetBaseId()(*string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetDisplayVersion()(*string)
    GetLifecycleState()(*DeviceManagementTemplateLifecycleState)
    GetPlatforms()(*DeviceManagementConfigurationPlatforms)
    GetSettingTemplateCount()(*int32)
    GetSettingTemplates()([]DeviceManagementConfigurationSettingTemplateable)
    GetTechnologies()(*DeviceManagementConfigurationTechnologies)
    GetTemplateFamily()(*DeviceManagementConfigurationTemplateFamily)
    GetVersion()(*int32)
    SetAllowUnmanagedSettings(value *bool)()
    SetBaseId(value *string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetDisplayVersion(value *string)()
    SetLifecycleState(value *DeviceManagementTemplateLifecycleState)()
    SetPlatforms(value *DeviceManagementConfigurationPlatforms)()
    SetSettingTemplateCount(value *int32)()
    SetSettingTemplates(value []DeviceManagementConfigurationSettingTemplateable)()
    SetTechnologies(value *DeviceManagementConfigurationTechnologies)()
    SetTemplateFamily(value *DeviceManagementConfigurationTemplateFamily)()
    SetVersion(value *int32)()
}
