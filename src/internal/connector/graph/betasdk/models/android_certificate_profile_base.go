package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidCertificateProfileBase 
type AndroidCertificateProfileBase struct {
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
    rootCertificate AndroidTrustedRootCertificateable
    // Subject Alternative Name Options.
    subjectAlternativeNameType *SubjectAlternativeNameType
    // Subject Name Format Options.
    subjectNameFormat *SubjectNameFormat
}
// NewAndroidCertificateProfileBase instantiates a new AndroidCertificateProfileBase and sets the default values.
func NewAndroidCertificateProfileBase()(*AndroidCertificateProfileBase) {
    m := &AndroidCertificateProfileBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.androidCertificateProfileBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAndroidCertificateProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidCertificateProfileBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.androidForWorkImportedPFXCertificateProfile":
                        return NewAndroidForWorkImportedPFXCertificateProfile(), nil
                    case "#microsoft.graph.androidImportedPFXCertificateProfile":
                        return NewAndroidImportedPFXCertificateProfile(), nil
                    case "#microsoft.graph.androidPkcsCertificateProfile":
                        return NewAndroidPkcsCertificateProfile(), nil
                    case "#microsoft.graph.androidScepCertificateProfile":
                        return NewAndroidScepCertificateProfile(), nil
                }
            }
        }
    }
    return NewAndroidCertificateProfileBase(), nil
}
// GetCertificateValidityPeriodScale gets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *AndroidCertificateProfileBase) GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale) {
    return m.certificateValidityPeriodScale
}
// GetCertificateValidityPeriodValue gets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *AndroidCertificateProfileBase) GetCertificateValidityPeriodValue()(*int32) {
    return m.certificateValidityPeriodValue
}
// GetExtendedKeyUsages gets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *AndroidCertificateProfileBase) GetExtendedKeyUsages()([]ExtendedKeyUsageable) {
    return m.extendedKeyUsages
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidCertificateProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetObjectValue(CreateAndroidTrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificate(val.(AndroidTrustedRootCertificateable))
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
func (m *AndroidCertificateProfileBase) GetRenewalThresholdPercentage()(*int32) {
    return m.renewalThresholdPercentage
}
// GetRootCertificate gets the rootCertificate property value. Trusted Root Certificate.
func (m *AndroidCertificateProfileBase) GetRootCertificate()(AndroidTrustedRootCertificateable) {
    return m.rootCertificate
}
// GetSubjectAlternativeNameType gets the subjectAlternativeNameType property value. Subject Alternative Name Options.
func (m *AndroidCertificateProfileBase) GetSubjectAlternativeNameType()(*SubjectAlternativeNameType) {
    return m.subjectAlternativeNameType
}
// GetSubjectNameFormat gets the subjectNameFormat property value. Subject Name Format Options.
func (m *AndroidCertificateProfileBase) GetSubjectNameFormat()(*SubjectNameFormat) {
    return m.subjectNameFormat
}
// Serialize serializes information the current object
func (m *AndroidCertificateProfileBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AndroidCertificateProfileBase) SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)() {
    m.certificateValidityPeriodScale = value
}
// SetCertificateValidityPeriodValue sets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period.
func (m *AndroidCertificateProfileBase) SetCertificateValidityPeriodValue(value *int32)() {
    m.certificateValidityPeriodValue = value
}
// SetExtendedKeyUsages sets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *AndroidCertificateProfileBase) SetExtendedKeyUsages(value []ExtendedKeyUsageable)() {
    m.extendedKeyUsages = value
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *AndroidCertificateProfileBase) SetRenewalThresholdPercentage(value *int32)() {
    m.renewalThresholdPercentage = value
}
// SetRootCertificate sets the rootCertificate property value. Trusted Root Certificate.
func (m *AndroidCertificateProfileBase) SetRootCertificate(value AndroidTrustedRootCertificateable)() {
    m.rootCertificate = value
}
// SetSubjectAlternativeNameType sets the subjectAlternativeNameType property value. Subject Alternative Name Options.
func (m *AndroidCertificateProfileBase) SetSubjectAlternativeNameType(value *SubjectAlternativeNameType)() {
    m.subjectAlternativeNameType = value
}
// SetSubjectNameFormat sets the subjectNameFormat property value. Subject Name Format Options.
func (m *AndroidCertificateProfileBase) SetSubjectNameFormat(value *SubjectNameFormat)() {
    m.subjectNameFormat = value
}
