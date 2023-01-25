package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSCertificateProfileBaseable 
type MacOSCertificateProfileBaseable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale)
    GetCertificateValidityPeriodValue()(*int32)
    GetRenewalThresholdPercentage()(*int32)
    GetSubjectAlternativeNameType()(*SubjectAlternativeNameType)
    GetSubjectNameFormat()(*AppleSubjectNameFormat)
    SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)()
    SetCertificateValidityPeriodValue(value *int32)()
    SetRenewalThresholdPercentage(value *int32)()
    SetSubjectAlternativeNameType(value *SubjectAlternativeNameType)()
    SetSubjectNameFormat(value *AppleSubjectNameFormat)()
}
