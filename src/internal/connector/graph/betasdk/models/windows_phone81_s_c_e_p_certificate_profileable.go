package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81SCEPCertificateProfileable 
type WindowsPhone81SCEPCertificateProfileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    WindowsPhone81CertificateProfileBaseable
    GetHashAlgorithm()(*HashAlgorithms)
    GetKeySize()(*KeySize)
    GetKeyUsage()(*KeyUsages)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetRootCertificate()(WindowsPhone81TrustedRootCertificateable)
    GetScepServerUrls()([]string)
    GetSubjectAlternativeNameFormatString()(*string)
    GetSubjectNameFormatString()(*string)
    SetHashAlgorithm(value *HashAlgorithms)()
    SetKeySize(value *KeySize)()
    SetKeyUsage(value *KeyUsages)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetRootCertificate(value WindowsPhone81TrustedRootCertificateable)()
    SetScepServerUrls(value []string)()
    SetSubjectAlternativeNameFormatString(value *string)()
    SetSubjectNameFormatString(value *string)()
}
