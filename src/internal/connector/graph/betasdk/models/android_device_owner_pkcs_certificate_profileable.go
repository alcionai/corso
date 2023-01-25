package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerPkcsCertificateProfileable 
type AndroidDeviceOwnerPkcsCertificateProfileable interface {
    AndroidDeviceOwnerCertificateProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateAccessType()(*AndroidDeviceOwnerCertificateAccessType)
    GetCertificateStore()(*CertificateStore)
    GetCertificateTemplateName()(*string)
    GetCertificationAuthority()(*string)
    GetCertificationAuthorityName()(*string)
    GetCertificationAuthorityType()(*DeviceManagementCertificationAuthority)
    GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetSilentCertificateAccessDetails()([]AndroidDeviceOwnerSilentCertificateAccessable)
    GetSubjectAlternativeNameFormatString()(*string)
    GetSubjectNameFormatString()(*string)
    SetCertificateAccessType(value *AndroidDeviceOwnerCertificateAccessType)()
    SetCertificateStore(value *CertificateStore)()
    SetCertificateTemplateName(value *string)()
    SetCertificationAuthority(value *string)()
    SetCertificationAuthorityName(value *string)()
    SetCertificationAuthorityType(value *DeviceManagementCertificationAuthority)()
    SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetSilentCertificateAccessDetails(value []AndroidDeviceOwnerSilentCertificateAccessable)()
    SetSubjectAlternativeNameFormatString(value *string)()
    SetSubjectNameFormatString(value *string)()
}
