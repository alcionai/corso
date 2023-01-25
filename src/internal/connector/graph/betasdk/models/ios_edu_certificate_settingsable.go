package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosEduCertificateSettingsable 
type IosEduCertificateSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertFileName()(*string)
    GetCertificateTemplateName()(*string)
    GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale)
    GetCertificateValidityPeriodValue()(*int32)
    GetCertificationAuthority()(*string)
    GetCertificationAuthorityName()(*string)
    GetOdataType()(*string)
    GetRenewalThresholdPercentage()(*int32)
    GetTrustedRootCertificate()([]byte)
    SetCertFileName(value *string)()
    SetCertificateTemplateName(value *string)()
    SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)()
    SetCertificateValidityPeriodValue(value *int32)()
    SetCertificationAuthority(value *string)()
    SetCertificationAuthorityName(value *string)()
    SetOdataType(value *string)()
    SetRenewalThresholdPercentage(value *int32)()
    SetTrustedRootCertificate(value []byte)()
}
