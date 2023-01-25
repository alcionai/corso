package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileCertificateProfileBaseable 
type AndroidWorkProfileCertificateProfileBaseable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale)
    GetCertificateValidityPeriodValue()(*int32)
    GetExtendedKeyUsages()([]ExtendedKeyUsageable)
    GetRenewalThresholdPercentage()(*int32)
    GetRootCertificate()(AndroidWorkProfileTrustedRootCertificateable)
    GetSubjectAlternativeNameType()(*SubjectAlternativeNameType)
    GetSubjectNameFormat()(*SubjectNameFormat)
    SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)()
    SetCertificateValidityPeriodValue(value *int32)()
    SetExtendedKeyUsages(value []ExtendedKeyUsageable)()
    SetRenewalThresholdPercentage(value *int32)()
    SetRootCertificate(value AndroidWorkProfileTrustedRootCertificateable)()
    SetSubjectAlternativeNameType(value *SubjectAlternativeNameType)()
    SetSubjectNameFormat(value *SubjectNameFormat)()
}
