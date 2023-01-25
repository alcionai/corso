package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileVpnConfigurationable 
type AndroidWorkProfileVpnConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlwaysOn()(*bool)
    GetAlwaysOnLockdown()(*bool)
    GetAuthenticationMethod()(*VpnAuthenticationMethod)
    GetConnectionName()(*string)
    GetConnectionType()(*AndroidWorkProfileVpnConnectionType)
    GetCustomData()([]KeyValueable)
    GetCustomKeyValueData()([]KeyValuePairable)
    GetFingerprint()(*string)
    GetIdentityCertificate()(AndroidWorkProfileCertificateProfileBaseable)
    GetMicrosoftTunnelSiteId()(*string)
    GetProxyServer()(VpnProxyServerable)
    GetRealm()(*string)
    GetRole()(*string)
    GetServers()([]VpnServerable)
    GetTargetedMobileApps()([]AppListItemable)
    GetTargetedPackageIds()([]string)
    SetAlwaysOn(value *bool)()
    SetAlwaysOnLockdown(value *bool)()
    SetAuthenticationMethod(value *VpnAuthenticationMethod)()
    SetConnectionName(value *string)()
    SetConnectionType(value *AndroidWorkProfileVpnConnectionType)()
    SetCustomData(value []KeyValueable)()
    SetCustomKeyValueData(value []KeyValuePairable)()
    SetFingerprint(value *string)()
    SetIdentityCertificate(value AndroidWorkProfileCertificateProfileBaseable)()
    SetMicrosoftTunnelSiteId(value *string)()
    SetProxyServer(value VpnProxyServerable)()
    SetRealm(value *string)()
    SetRole(value *string)()
    SetServers(value []VpnServerable)()
    SetTargetedMobileApps(value []AppListItemable)()
    SetTargetedPackageIds(value []string)()
}
