package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationable 
type AndroidDeviceOwnerDerivedCredentialAuthenticationConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateAccessType()(*AndroidDeviceOwnerCertificateAccessType)
    GetDerivedCredentialSettings()(DeviceManagementDerivedCredentialSettingsable)
    GetSilentCertificateAccessDetails()([]AndroidDeviceOwnerSilentCertificateAccessable)
    SetCertificateAccessType(value *AndroidDeviceOwnerCertificateAccessType)()
    SetDerivedCredentialSettings(value DeviceManagementDerivedCredentialSettingsable)()
    SetSilentCertificateAccessDetails(value []AndroidDeviceOwnerSilentCertificateAccessable)()
}
