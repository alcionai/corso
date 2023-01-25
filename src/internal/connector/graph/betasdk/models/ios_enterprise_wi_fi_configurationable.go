package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosEnterpriseWiFiConfigurationable 
type IosEnterpriseWiFiConfigurationable interface {
    IosWiFiConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable)
    GetEapFastConfiguration()(*EapFastConfiguration)
    GetEapType()(*EapType)
    GetIdentityCertificateForClientAuthentication()(IosCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetPasswordFormatString()(*string)
    GetRootCertificatesForServerValidation()([]IosTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    GetUsernameFormatString()(*string)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)()
    SetEapFastConfiguration(value *EapFastConfiguration)()
    SetEapType(value *EapType)()
    SetIdentityCertificateForClientAuthentication(value IosCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetPasswordFormatString(value *string)()
    SetRootCertificatesForServerValidation(value []IosTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
    SetUsernameFormatString(value *string)()
}
