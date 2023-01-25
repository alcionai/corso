package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileEnterpriseWiFiConfigurationable 
type AndroidWorkProfileEnterpriseWiFiConfigurationable interface {
    AndroidWorkProfileWiFiConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetEapType()(*AndroidEapType)
    GetIdentityCertificateForClientAuthentication()(AndroidWorkProfileCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType)
    GetInnerAuthenticationProtocolForPeap()(*NonEapAuthenticationMethodForPeap)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetProxyAutomaticConfigurationUrl()(*string)
    GetProxySettings()(*WiFiProxySetting)
    GetRootCertificateForServerValidation()(AndroidWorkProfileTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetEapType(value *AndroidEapType)()
    SetIdentityCertificateForClientAuthentication(value AndroidWorkProfileCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetInnerAuthenticationProtocolForPeap(value *NonEapAuthenticationMethodForPeap)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetProxyAutomaticConfigurationUrl(value *string)()
    SetProxySettings(value *WiFiProxySetting)()
    SetRootCertificateForServerValidation(value AndroidWorkProfileTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
}
