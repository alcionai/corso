package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSEnterpriseWiFiConfigurationable 
type MacOSEnterpriseWiFiConfigurationable interface {
    MacOSWiFiConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetEapFastConfiguration()(*EapFastConfiguration)
    GetEapType()(*EapType)
    GetIdentityCertificateForClientAuthentication()(MacOSCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetRootCertificateForServerValidation()(MacOSTrustedRootCertificateable)
    GetRootCertificatesForServerValidation()([]MacOSTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetEapFastConfiguration(value *EapFastConfiguration)()
    SetEapType(value *EapType)()
    SetIdentityCertificateForClientAuthentication(value MacOSCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetRootCertificateForServerValidation(value MacOSTrustedRootCertificateable)()
    SetRootCertificatesForServerValidation(value []MacOSTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
}
