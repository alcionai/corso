package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerEnterpriseWiFiConfigurationable 
type AndroidDeviceOwnerEnterpriseWiFiConfigurationable interface {
    AndroidDeviceOwnerWiFiConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuthenticationMethod()(*WiFiAuthenticationMethod)
    GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable)
    GetEapType()(*AndroidEapType)
    GetIdentityCertificateForClientAuthentication()(AndroidDeviceOwnerCertificateProfileBaseable)
    GetInnerAuthenticationProtocolForEapTtls()(*NonEapAuthenticationMethodForEapTtlsType)
    GetInnerAuthenticationProtocolForPeap()(*NonEapAuthenticationMethodForPeap)
    GetOuterIdentityPrivacyTemporaryValue()(*string)
    GetRootCertificateForServerValidation()(AndroidDeviceOwnerTrustedRootCertificateable)
    GetTrustedServerCertificateNames()([]string)
    SetAuthenticationMethod(value *WiFiAuthenticationMethod)()
    SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)()
    SetEapType(value *AndroidEapType)()
    SetIdentityCertificateForClientAuthentication(value AndroidDeviceOwnerCertificateProfileBaseable)()
    SetInnerAuthenticationProtocolForEapTtls(value *NonEapAuthenticationMethodForEapTtlsType)()
    SetInnerAuthenticationProtocolForPeap(value *NonEapAuthenticationMethodForPeap)()
    SetOuterIdentityPrivacyTemporaryValue(value *string)()
    SetRootCertificateForServerValidation(value AndroidDeviceOwnerTrustedRootCertificateable)()
    SetTrustedServerCertificateNames(value []string)()
}
