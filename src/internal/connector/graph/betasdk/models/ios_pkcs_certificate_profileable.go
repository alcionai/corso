package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosPkcsCertificateProfileable 
type IosPkcsCertificateProfileable interface {
    IosCertificateProfileBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateStore()(*CertificateStore)
    GetCertificateTemplateName()(*string)
    GetCertificationAuthority()(*string)
    GetCertificationAuthorityName()(*string)
    GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable)
    GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable)
    GetSubjectAlternativeNameFormatString()(*string)
    GetSubjectNameFormatString()(*string)
    SetCertificateStore(value *CertificateStore)()
    SetCertificateTemplateName(value *string)()
    SetCertificationAuthority(value *string)()
    SetCertificationAuthorityName(value *string)()
    SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)()
    SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)()
    SetSubjectAlternativeNameFormatString(value *string)()
    SetSubjectNameFormatString(value *string)()
}
