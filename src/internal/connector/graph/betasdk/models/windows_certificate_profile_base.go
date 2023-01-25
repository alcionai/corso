package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsCertificateProfileBase 
type WindowsCertificateProfileBase struct {
    DeviceConfiguration
    // Certificate Validity Period Options.
    certificateValidityPeriodScale *CertificateValidityPeriodScale
    // Value for the Certificate Validity Period
    certificateValidityPeriodValue *int32
    // Key Storage Provider (KSP) Import Options.
    keyStorageProvider *KeyStorageProviderOption
    // Certificate renewal threshold percentage. Valid values 1 to 99
    renewalThresholdPercentage *int32
    // Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
    subjectAlternativeNameType *SubjectAlternativeNameType
    // Subject Name Format Options.
    subjectNameFormat *SubjectNameFormat
}
// NewWindowsCertificateProfileBase instantiates a new WindowsCertificateProfileBase and sets the default values.
func NewWindowsCertificateProfileBase()(*WindowsCertificateProfileBase) {
    m := &WindowsCertificateProfileBase{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsCertificateProfileBase";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindowsCertificateProfileBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsCertificateProfileBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.windows10CertificateProfileBase":
                        return NewWindows10CertificateProfileBase(), nil
                    case "#microsoft.graph.windows10ImportedPFXCertificateProfile":
                        return NewWindows10ImportedPFXCertificateProfile(), nil
                    case "#microsoft.graph.windows10PkcsCertificateProfile":
                        return NewWindows10PkcsCertificateProfile(), nil
                    case "#microsoft.graph.windows81CertificateProfileBase":
                        return NewWindows81CertificateProfileBase(), nil
                    case "#microsoft.graph.windows81SCEPCertificateProfile":
                        return NewWindows81SCEPCertificateProfile(), nil
                    case "#microsoft.graph.windowsPhone81ImportedPFXCertificateProfile":
                        return NewWindowsPhone81ImportedPFXCertificateProfile(), nil
                }
            }
        }
    }
    return NewWindowsCertificateProfileBase(), nil
}
// GetCertificateValidityPeriodScale gets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *WindowsCertificateProfileBase) GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale) {
    return m.certificateValidityPeriodScale
}
// GetCertificateValidityPeriodValue gets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period
func (m *WindowsCertificateProfileBase) GetCertificateValidityPeriodValue()(*int32) {
    return m.certificateValidityPeriodValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsCertificateProfileBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["keyStorageProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeyStorageProviderOption)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyStorageProvider(val.(*KeyStorageProviderOption))
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
// GetKeyStorageProvider gets the keyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *WindowsCertificateProfileBase) GetKeyStorageProvider()(*KeyStorageProviderOption) {
    return m.keyStorageProvider
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *WindowsCertificateProfileBase) GetRenewalThresholdPercentage()(*int32) {
    return m.renewalThresholdPercentage
}
// GetSubjectAlternativeNameType gets the subjectAlternativeNameType property value. Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
func (m *WindowsCertificateProfileBase) GetSubjectAlternativeNameType()(*SubjectAlternativeNameType) {
    return m.subjectAlternativeNameType
}
// GetSubjectNameFormat gets the subjectNameFormat property value. Subject Name Format Options.
func (m *WindowsCertificateProfileBase) GetSubjectNameFormat()(*SubjectNameFormat) {
    return m.subjectNameFormat
}
// Serialize serializes information the current object
func (m *WindowsCertificateProfileBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetKeyStorageProvider() != nil {
        cast := (*m.GetKeyStorageProvider()).String()
        err = writer.WriteStringValue("keyStorageProvider", &cast)
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
func (m *WindowsCertificateProfileBase) SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)() {
    m.certificateValidityPeriodScale = value
}
// SetCertificateValidityPeriodValue sets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period
func (m *WindowsCertificateProfileBase) SetCertificateValidityPeriodValue(value *int32)() {
    m.certificateValidityPeriodValue = value
}
// SetKeyStorageProvider sets the keyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *WindowsCertificateProfileBase) SetKeyStorageProvider(value *KeyStorageProviderOption)() {
    m.keyStorageProvider = value
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. Certificate renewal threshold percentage. Valid values 1 to 99
func (m *WindowsCertificateProfileBase) SetRenewalThresholdPercentage(value *int32)() {
    m.renewalThresholdPercentage = value
}
// SetSubjectAlternativeNameType sets the subjectAlternativeNameType property value. Certificate Subject Alternative Name Type. Possible values are: none, emailAddress, userPrincipalName, customAzureADAttribute, domainNameService, universalResourceIdentifier.
func (m *WindowsCertificateProfileBase) SetSubjectAlternativeNameType(value *SubjectAlternativeNameType)() {
    m.subjectAlternativeNameType = value
}
// SetSubjectNameFormat sets the subjectNameFormat property value. Subject Name Format Options.
func (m *WindowsCertificateProfileBase) SetSubjectNameFormat(value *SubjectNameFormat)() {
    m.subjectNameFormat = value
}
