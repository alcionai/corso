package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerVpnConfigurationable 
type AndroidDeviceOwnerVpnConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    VpnConfigurationable
    GetAlwaysOn()(*bool)
    GetAlwaysOnLockdown()(*bool)
    GetConnectionType()(*AndroidVpnConnectionType)
    GetCustomData()([]KeyValueable)
    GetCustomKeyValueData()([]KeyValuePairable)
    GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable)
    GetIdentityCertificate()(AndroidDeviceOwnerCertificateProfileBaseable)
    GetMicrosoftTunnelSiteId()(*string)
    GetProxyServer()(VpnProxyServerable)
    GetTargetedMobileApps()([]AppListItemable)
    GetTargetedPackageIds()([]string)
    SetAlwaysOn(value *bool)()
    SetAlwaysOnLockdown(value *bool)()
    SetConnectionType(value *AndroidVpnConnectionType)()
    SetCustomData(value []KeyValueable)()
    SetCustomKeyValueData(value []KeyValuePairable)()
    SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)()
    SetIdentityCertificate(value AndroidDeviceOwnerCertificateProfileBaseable)()
    SetMicrosoftTunnelSiteId(value *string)()
    SetProxyServer(value VpnProxyServerable)()
    SetTargetedMobileApps(value []AppListItemable)()
    SetTargetedPackageIds(value []string)()
}
