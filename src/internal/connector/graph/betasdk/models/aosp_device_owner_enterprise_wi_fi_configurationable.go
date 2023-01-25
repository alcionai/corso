package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AospDeviceOwnerEnterpriseWiFiConfigurationable 
type AospDeviceOwnerEnterpriseWiFiConfigurationable interface {
    AospDeviceOwnerWiFiConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetEapType()(*AndroidEapType)
    GetIdentityCertificateForClientAuthentication()(AospDeviceOwnerCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType)
    GetInnerAuthenticationProtocolForPeap()(*NonEapAuthenticationMethodForPeap)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetRootCertificateForServerValidation()(AospDeviceOwnerTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetEapType(value *AndroidEapType)()
    SetIdentityCertificateForClientAuthentication(value AospDeviceOwnerCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetInnerAuthenticationProtocolForPeap(value *NonEapAuthenticationMethodForPeap)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetRootCertificateForServerValidation(value AospDeviceOwnerTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
}
