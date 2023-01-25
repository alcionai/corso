package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSWiredNetworkConfigurationable 
type MacOSWiredNetworkConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetEapFastConfiguration()(*EapFastConfiguration)
    GetEapType()(*EapType)
    GetEnableOuterIdentityPrivacy()(*string)
    GetIdentityCertificateForClientAuthentication()(MacOSCertificateProfileBaseable)
    GetNetworkInterface()(*WiredNetworkInterface)
    GetNetworkName()(*string)
    GetNonEapAuthenticationMethodForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType)
    GetRootCertificateForServerValidation()(MacOSTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetEapFastConfiguration(value *EapFastConfiguration)()
    SetEapType(value *EapType)()
    SetEnableOuterIdentityPrivacy(value *string)()
    SetIdentityCertificateForClientAuthentication(value MacOSCertificateProfileBaseable)()
    SetNetworkInterface(value *WiredNetworkInterface)()
    SetNetworkName(value *string)()
    SetNonEapAuthenticationMethodForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetRootCertificateForServerValidation(value MacOSTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
}
