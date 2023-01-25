package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidForWorkEnterpriseWiFiConfigurationable 
type AndroidForWorkEnterpriseWiFiConfigurationable interface {
    AndroidForWorkWiFiConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetEapType()(*AndroidEapType)
    GetIdentityCertificateForClientAuthentication()(AndroidForWorkCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType)
    GetInnerAuthenticationProtocolForPeap()(*NonEapAuthenticationMethodForPeap)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetRootCertificateForServerValidation()(AndroidForWorkTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetEapType(value *AndroidEapType)()
    SetIdentityCertificateForClientAuthentication(value AndroidForWorkCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetInnerAuthenticationProtocolForPeap(value *NonEapAuthenticationMethodForPeap)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetRootCertificateForServerValidation(value AndroidForWorkTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
}
