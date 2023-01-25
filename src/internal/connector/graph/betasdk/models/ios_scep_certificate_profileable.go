package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosScepCertificateProfileable 
type IosScepCertificateProfileable interface {
    IosCertificateProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateStore()(*CertificateStore)
    GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable)
    GetExtendedKeyUsages()([]ExtendedKeyUsageable)
    GetKeySize()(*KeySize)
    GetKeyUsage()(*KeyUsages)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetRootCertificate()(IosTrustedRootCertificateable)
    GetScepServerUrls()([]string)
    GetSubjectAlternativeNameFormatString()(*string)
    GetSubjectNameFormatString()(*string)
    SetCertificateStore(value *CertificateStore)()
    SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)()
    SetExtendedKeyUsages(value []ExtendedKeyUsageable)()
    SetKeySize(value *KeySize)()
    SetKeyUsage(value *KeyUsages)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetRootCertificate(value IosTrustedRootCertificateable)()
    SetScepServerUrls(value []string)()
    SetSubjectAlternativeNameFormatString(value *string)()
    SetSubjectNameFormatString(value *string)()
}
