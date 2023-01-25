package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileScepCertificateProfileable 
type AndroidWorkProfileScepCertificateProfileable interface {
    AndroidWorkProfileCertificateProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateStore()(*CertificateStore)
    GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable)
    GetHashAlgorithm()(*HashAlgorithms)
    GetKeySize()(*KeySize)
    GetKeyUsage()(*KeyUsages)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetScepServerUrls()([]string)
    GetSubjectAlternativeNameFormatString()(*string)
    GetSubjectNameFormatString()(*string)
    SetCertificateStore(value *CertificateStore)()
    SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)()
    SetHashAlgorithm(value *HashAlgorithms)()
    SetKeySize(value *KeySize)()
    SetKeyUsage(value *KeyUsages)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetScepServerUrls(value []string)()
    SetSubjectAlternativeNameFormatString(value *string)()
    SetSubjectNameFormatString(value *string)()
}
