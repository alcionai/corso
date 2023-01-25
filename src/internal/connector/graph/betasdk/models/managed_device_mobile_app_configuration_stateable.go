package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceMobileAppConfigurationStateable 
type ManagedDeviceMobileAppConfigurationStateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetPlatformType()(*PolicyPlatformType)
    GetSettingCount()(*int32)
    GetSettingStates()([]ManagedDeviceMobileAppConfigurationSettingStateable)
    GetState()(*ComplianceStatus)
    GetUserId()(*string)
    GetUserPrincipalName()(*string)
    GetVersion()(*int32)
    SetDisplayName(value *string)()
    SetPlatformType(value *PolicyPlatformType)()
    SetSettingCount(value *int32)()
    SetSettingStates(value []ManagedDeviceMobileAppConfigurationSettingStateable)()
    SetState(value *ComplianceStatus)()
    SetUserId(value *string)()
    SetUserPrincipalName(value *string)()
    SetVersion(value *int32)()
}
