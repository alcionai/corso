package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidEnterpriseWiFiConfigurationable 
type AndroidEnterpriseWiFiConfigurationable interface {
    AndroidWiFiConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetEapType()(*AndroidEapType)
    GetIdentityCertificateForClientAuthentication()(AndroidCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType)
    GetInnerAuthenticationProtocolForPeap()(*NonEapAuthenticationMethodForPeap)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetPasswordFormatString()(*string)
    GetPreSharedKey()(*string)
    GetRootCertificateForServerValidation()(AndroidTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    GetUsernameFormatString()(*string)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetEapType(value *AndroidEapType)()
    SetIdentityCertificateForClientAuthentication(value AndroidCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetInnerAuthenticationProtocolForPeap(value *NonEapAuthenticationMethodForPeap)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetPasswordFormatString(value *string)()
    SetPreSharedKey(value *string)()
    SetRootCertificateForServerValidation(value AndroidTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
    SetUsernameFormatString(value *string)()
}
