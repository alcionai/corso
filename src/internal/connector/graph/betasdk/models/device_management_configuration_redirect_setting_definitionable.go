package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementConfigurationRedirectSettingDefinitionable 
type DeviceManagementConfigurationRedirectSettingDefinitionable interface {
    DeviceManagementConfigurationSettingDefinitionable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeepLink()(*string)
    GetRedirectMessage()(*string)
    GetRedirectReason()(*string)
    SetDeepLink(value *string)()
    SetRedirectMessage(value *string)()
    SetRedirectReason(value *string)()
}
