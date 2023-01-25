package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVpnConfigurationable 
type IosVpnConfigurationable interface {
    AppleVpnConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudName()(*string)
    GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable)
    GetExcludeList()([]string)
    GetIdentityCertificate()(IosCertificateProfileBaseable)
    GetMicrosoftTunnelSiteId()(*string)
    GetStrictEnforcement()(*bool)
    GetTargetedMobileApps()([]AppListItemable)
    GetUserDomain()(*string)
    SetCloudName(value *string)()
    SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)()
    SetExcludeList(value []string)()
    SetIdentityCertificate(value IosCertificateProfileBaseable)()
    SetMicrosoftTunnelSiteId(value *string)()
    SetStrictEnforcement(value *bool)()
    SetTargetedMobileApps(value []AppListItemable)()
    SetUserDomain(value *string)()
}
