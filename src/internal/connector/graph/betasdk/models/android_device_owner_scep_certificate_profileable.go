package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerScepCertificateProfileable 
type AndroidDeviceOwnerScepCertificateProfileable interface {
    AndroidDeviceOwnerCertificateProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateAccessType()(*AndroidDeviceOwnerCertificateAccessType)
    GetCertificateStore()(*CertificateStore)
    GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable)
    GetHashAlgorithm()(*HashAlgorithms)
    GetKeySize()(*KeySize)
    GetKeyUsage()(*KeyUsages)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetScepServerUrls()([]string)
    GetSilentCertificateAccessDetails()([]AndroidDeviceOwnerSilentCertificateAccessable)
    GetSubjectAlternativeNameFormatString()(*string)
    GetSubjectNameFormatString()(*string)
    SetCertificateAccessType(value *AndroidDeviceOwnerCertificateAccessType)()
    SetCertificateStore(value *CertificateStore)()
    SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)()
    SetHashAlgorithm(value *HashAlgorithms)()
    SetKeySize(value *KeySize)()
    SetKeyUsage(value *KeyUsages)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetScepServerUrls(value []string)()
    SetSilentCertificateAccessDetails(value []AndroidDeviceOwnerSilentCertificateAccessable)()
    SetSubjectAlternativeNameFormatString(value *string)()
    SetSubjectNameFormatString(value *string)()
}
