package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidDeviceOwnerCertificateProfileBase 
type AndroidDeviceOwnerCertificateProfileBase struct {
    DeviceConfiguration
    // Certificate Validity Period Options.
    certificateValidityPeriodScale *CertificateValidityPeriodScale
    // Value for the Certificate Validity Period.
    certificateValidityPeriodValue *int32
    // Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
    extendedKeyUsages []ExtendedKeyUsageable
    // Certificate renewal threshold percentage. Valid values 1 to 99
    renewalThresholdPercentage *int32
    // Trusted Root Certificate.
    rootCertificate AndroidDeviceOwnerTrustedRootCertificateable
    // Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
    subjectAlternativeNameType *SubjectAlternativeNameType
    // Certificate Subject Name Format. Possible values are: commonName, commonNameIncludingEmail, commonNameAsEmail, custom, commonNameAsIMEI, commonNameAsSerialNumber, commonNameAsAadDeviceId, commonNameAsIntuneDeviceId, commonNameAsDurableDeviceId.
    subjectNameFormat *SubjectNameFormat
}
// NewAndroidDeviceOwnerCertificateProfileBase instantiates a new AndroidDeviceOwnerCertificateProfileBase and sets the default values.
func NewAndroidDeviceOwnerCertificateProfileBase()(*AndroidDeviceOwnerCertificateProfileBase) {
    m := &AndroidDeviceOwnerCertificateProfileBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidDeviceOwnerCertificateProfileBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidDeviceOwnerCertificateProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidDeviceOwnerCertificateProfileBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.androidDeviceOwnerImportedPFXCertificateProfile":
                        return NewAndroidDeviceOwnerImportedPFXCertificateProfile(), nil
                    case "#microsoft.graph.androidDeviceOwnerPkcsCertificateProfile":
                        return NewAndroidDeviceOwnerPkcsCertificateProfile(), nil
                    case "#microsoft.graph.androidDeviceOwnerScepCertificateProfile":
                        return NewAndroidDeviceOwnerScepCertificateProfile(), nil
                }
            }
        }
    }
    return NewAndroidDeviceOwnerCertificateProfileBase(), nil
}
// GetCertificateValidityPeriodScale gets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *AndroidDeviceOwnerCertificateProfileBase) GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale) {
    return m.certificateValidityPeriodScale
}
// GetCertificateValidityPeriodValue gets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *AndroidDeviceOwnerCertificateProfileBase) GetCertificateValidityPeriodValue()(*int32) {
    return m.certificateValidityPeriodValue
}
// GetExtendedKeyUsages gets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerCertificateProfileBase) GetExtendedKeyUsages()([]ExtendedKeyUsageable) {
    return m.extendedKeyUsages
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidDeviceOwnerCertificateProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["certificateValidityPeriodScale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateValidityPeriodScale)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateValidityPeriodScale(val.(*CertificateValidityPeriodScale))
        }
        return nil
    }
    res["certificateValidityPeriodValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateValidityPeriodValue(val)
        }
        return nil
    }
    res["extendedKeyUsages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtendedKeyUsageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExtendedKeyUsageable, len(val))
            for i, v := range val {
                res[i] = v.(ExtendedKeyUsageable)
            }
            m.SetExtendedKeyUsages(res)
        }
        return nil
    }
    res["renewalThresholdPercentage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRenewalThresholdPercentage(val)
        }
        return nil
    }
    res["rootCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidDeviceOwnerTrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificate(val.(AndroidDeviceOwnerTrustedRootCertificateable))
        }
        return nil
    }
    res["subjectAlternativeNameType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectAlternativeNameType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectAlternativeNameType(val.(*SubjectAlternativeNameType))
        }
        return nil
    }
    res["subjectNameFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectNameFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectNameFormat(val.(*SubjectNameFormat))
        }
        return nil
    }
    return res
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *AndroidDeviceOwnerCertificateProfileBase) GetRenewalThresholdPercentage()(*int32) {
    return m.renewalThresholdPercentage
}
// GetRootCertificate gets the rootCertificate property value. Trusted Root Certificate.
func (m *AndroidDeviceOwnerCertificateProfileBase) GetRootCertificate()(AndroidDeviceOwnerTrustedRootCertificateable) {
    return m.rootCertificate
}
// GetSubjectAlternativeNameType gets the subjectAlternativeNameType property value. Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
func (m *AndroidDeviceOwnerCertificateProfileBase) GetSubjectAlternativeNameType()(*SubjectAlternativeNameType) {
    return m.subjectAlternativeNameType
}
// GetSubjectNameFormat gets the subjectNameFormat property value. Certificate Subject Name Format. Possible values are: commonName, commonNameIncludingEmail, commonNameAsEmail, custom, commonNameAsIMEI, commonNameAsSerialNumber, commonNameAsAadDeviceId, commonNameAsIntuneDeviceId, commonNameAsDurableDeviceId.
func (m *AndroidDeviceOwnerCertificateProfileBase) GetSubjectNameFormat()(*SubjectNameFormat) {
    return m.subjectNameFormat
}
// Serialize serializes information the current object
func (m *AndroidDeviceOwnerCertificateProfileBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateValidityPeriodScale() != nil {
        cast := (*m.GetCertificateValidityPeriodScale()).String()
        err = writer.WriteStringValue("certificateValidityPeriodScale", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("certificateValidityPeriodValue", m.GetCertificateValidityPeriodValue())
        if err != nil {
            return err
        }
    }
    if m.GetExtendedKeyUsages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtendedKeyUsages()))
        for i, v := range m.GetExtendedKeyUsages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extendedKeyUsages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("renewalThresholdPercentage", m.GetRenewalThresholdPercentage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rootCertificate", m.GetRootCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetSubjectAlternativeNameType() != nil {
        cast := (*m.GetSubjectAlternativeNameType()).String()
        err = writer.WriteStringValue("subjectAlternativeNameType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSubjectNameFormat() != nil {
        cast := (*m.GetSubjectNameFormat()).String()
        err = writer.WriteStringValue("subjectNameFormat", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateValidityPeriodScale sets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *AndroidDeviceOwnerCertificateProfileBase) SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)() {
    m.certificateValidityPeriodScale = value
}
// SetCertificateValidityPeriodValue sets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *AndroidDeviceOwnerCertificateProfileBase) SetCertificateValidityPeriodValue(value *int32)() {
    m.certificateValidityPeriodValue = value
}
// SetExtendedKeyUsages sets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *AndroidDeviceOwnerCertificateProfileBase) SetExtendedKeyUsages(value []ExtendedKeyUsageable)() {
    m.extendedKeyUsages = value
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *AndroidDeviceOwnerCertificateProfileBase) SetRenewalThresholdPercentage(value *int32)() {
    m.renewalThresholdPercentage = value
}
// SetRootCertificate sets the rootCertificate property value. Trusted Root Certificate.
func (m *AndroidDeviceOwnerCertificateProfileBase) SetRootCertificate(value AndroidDeviceOwnerTrustedRootCertificateable)() {
    m.rootCertificate = value
}
// SetSubjectAlternativeNameType sets the subjectAlternativeNameType property value. Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
func (m *AndroidDeviceOwnerCertificateProfileBase) SetSubjectAlternativeNameType(value *SubjectAlternativeNameType)() {
    m.subjectAlternativeNameType = value
}
// SetSubjectNameFormat sets the subjectNameFormat property value. Certificate Subject Name Format. Possible values are: commonName, commonNameIncludingEmail, commonNameAsEmail, custom, commonNameAsIMEI, commonNameAsSerialNumber, commonNameAsAadDeviceId, commonNameAsIntuneDeviceId, commonNameAsDurableDeviceId.
func (m *AndroidDeviceOwnerCertificateProfileBase) SetSubjectNameFormat(value *SubjectNameFormat)() {
    m.subjectNameFormat = value
}
