package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSScepCertificateProfileable 
type MacOSScepCertificateProfileable interface {
    MacOSCertificateProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowAllAppsAccess()(*bool)
    GetCertificateStore()(*CertificateStore)
    GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable)
    GetExtendedKeyUsages()([]ExtendedKeyUsageable)
    GetHashAlgorithm()(*HashAlgorithms)
    GetKeySize()(*KeySize)
    GetKeyUsage()(*KeyUsages)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetRootCertificate()(MacOSTrustedRootCertificateable)
    GetScepServerUrls()([]string)
    GetSubjectAlternativeNameFormatString()(*string)
    GetSubjectNameFormatString()(*string)
    SetAllowAllAppsAccess(value *bool)()
    SetCertificateStore(value *CertificateStore)()
    SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)()
    SetExtendedKeyUsages(value []ExtendedKeyUsageable)()
    SetHashAlgorithm(value *HashAlgorithms)()
    SetKeySize(value *KeySize)()
    SetKeyUsage(value *KeyUsages)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetRootCertificate(value MacOSTrustedRootCertificateable)()
    SetScepServerUrls(value []string)()
    SetSubjectAlternativeNameFormatString(value *string)()
    SetSubjectNameFormatString(value *string)()
}
