package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidVpnConfigurationable 
type AndroidVpnConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*VpnAuthenticationMethod)
    GetConnectionName()(*string)
    GetConnectionType()(*AndroidVpnConnectionType)
    GetCustomData()([]KeyValueable)
    GetCustomKeyValueData()([]KeyValuePairable)
    GetFingerprint()(*string)
    GetIdentityCertificate()(AndroidCertificateProfileBaseable)
    GetRealm()(*string)
    GetRole()(*string)
    GetServers()([]VpnServerable)
    SetAuthenticationMethod(value *VpnAuthenticationMethod)()
    SetConnectionName(value *string)()
    SetConnectionType(value *AndroidVpnConnectionType)()
    SetCustomData(value []KeyValueable)()
    SetCustomKeyValueData(value []KeyValuePairable)()
    SetFingerprint(value *string)()
    SetIdentityCertificate(value AndroidCertificateProfileBaseable)()
    SetRealm(value *string)()
    SetRole(value *string)()
    SetServers(value []VpnServerable)()
}
