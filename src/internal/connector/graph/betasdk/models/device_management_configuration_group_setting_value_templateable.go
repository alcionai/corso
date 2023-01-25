package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationGroupSettingValueTemplateable 
type DeviceManagementConfigurationGroupSettingValueTemplateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildren()([]DeviceManagementConfigurationSettingInstanceTemplateable)
    GetOdataType()(*string)
    GetSettingValueTemplateId()(*string)
    SetChildren(value []DeviceManagementConfigurationSettingInstanceTemplateable)()
    SetOdataType(value *string)()
    SetSettingValueTemplateId(value *string)()
}
