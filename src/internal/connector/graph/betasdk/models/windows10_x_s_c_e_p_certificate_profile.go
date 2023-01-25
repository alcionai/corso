package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10XSCEPCertificateProfile 
type Windows10XSCEPCertificateProfile struct {
    Windows10XCertificateProfile
    // CertificateStore types
    certificateStore *CertificateStore
    // Certificate Validity Period Options.
    certificateValidityPeriodScale *CertificateValidityPeriodScale
    // Value for the Certificate Validity Period
    certificateValidityPeriodValue *int32
    // Extended Key Usage (EKU) settings.
    extendedKeyUsages []ExtendedKeyUsageable
    // SCEP Hash Algorithm.
    hashAlgorithm []HashAlgorithms
    // Key Size Options.
    keySize *KeySize
    // Key Storage Provider (KSP) Import Options.
    keyStorageProvider *KeyStorageProviderOption
    // Key Usage Options.
    keyUsage *KeyUsages
    // Certificate renewal threshold percentage
    renewalThresholdPercentage *int32
    // Trusted Root Certificate ID
    rootCertificateId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // SCEP Server Url(s).
    scepServerUrls []string
    // Custom AAD Attributes.
    subjectAlternativeNameFormats []Windows10XCustomSubjectAlternativeNameable
    // Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
    subjectNameFormatString *string
}
// NewWindows10XSCEPCertificateProfile instantiates a new Windows10XSCEPCertificateProfile and sets the default values.
func NewWindows10XSCEPCertificateProfile()(*Windows10XSCEPCertificateProfile) {
    m := &Windows10XSCEPCertificateProfile{
        Windows10XCertificateProfile: *NewWindows10XCertificateProfile(),
    }
    odataTypeValue := "#microsoft.graph.windows10XSCEPCertificateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows10XSCEPCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10XSCEPCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10XSCEPCertificateProfile(), nil
}
// GetCertificateStore gets the certificateStore property value. CertificateStore types
func (m *Windows10XSCEPCertificateProfile) GetCertificateStore()(*CertificateStore) {
    return m.certificateStore
}
// GetCertificateValidityPeriodScale gets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *Windows10XSCEPCertificateProfile) GetCertificateValidityPeriodScale()(*CertificateValidityPeriodScale) {
    return m.certificateValidityPeriodScale
}
// GetCertificateValidityPeriodValue gets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period
func (m *Windows10XSCEPCertificateProfile) GetCertificateValidityPeriodValue()(*int32) {
    return m.certificateValidityPeriodValue
}
// GetExtendedKeyUsages gets the extendedKeyUsages property value. Extended Key Usage (EKU) settings.
func (m *Windows10XSCEPCertificateProfile) GetExtendedKeyUsages()([]ExtendedKeyUsageable) {
    return m.extendedKeyUsages
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10XSCEPCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Windows10XCertificateProfile.GetFieldDeserializers()
    res["certificateStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateStore)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateStore(val.(*CertificateStore))
        }
        return nil
    }
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
    res["hashAlgorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseHashAlgorithms)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HashAlgorithms, len(val))
            for i, v := range val {
                res[i] = *(v.(*HashAlgorithms))
            }
            m.SetHashAlgorithm(res)
        }
        return nil
    }
    res["keySize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeySize)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeySize(val.(*KeySize))
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
    res["keyUsage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseKeyUsages)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyUsage(val.(*KeyUsages))
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
    res["rootCertificateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificateId(val)
        }
        return nil
    }
    res["scepServerUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetScepServerUrls(res)
        }
        return nil
    }
    res["subjectAlternativeNameFormats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindows10XCustomSubjectAlternativeNameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Windows10XCustomSubjectAlternativeNameable, len(val))
            for i, v := range val {
                res[i] = v.(Windows10XCustomSubjectAlternativeNameable)
            }
            m.SetSubjectAlternativeNameFormats(res)
        }
        return nil
    }
    res["subjectNameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectNameFormatString(val)
        }
        return nil
    }
    return res
}
// GetHashAlgorithm gets the hashAlgorithm property value. SCEP Hash Algorithm.
func (m *Windows10XSCEPCertificateProfile) GetHashAlgorithm()([]HashAlgorithms) {
    return m.hashAlgorithm
}
// GetKeySize gets the keySize property value. Key Size Options.
func (m *Windows10XSCEPCertificateProfile) GetKeySize()(*KeySize) {
    return m.keySize
}
// GetKeyStorageProvider gets the keyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *Windows10XSCEPCertificateProfile) GetKeyStorageProvider()(*KeyStorageProviderOption) {
    return m.keyStorageProvider
}
// GetKeyUsage gets the keyUsage property value. Key Usage Options.
func (m *Windows10XSCEPCertificateProfile) GetKeyUsage()(*KeyUsages) {
    return m.keyUsage
}
// GetRenewalThresholdPercentage gets the renewalThresholdPercentage property value. Certificate renewal threshold percentage
func (m *Windows10XSCEPCertificateProfile) GetRenewalThresholdPercentage()(*int32) {
    return m.renewalThresholdPercentage
}
// GetRootCertificateId gets the rootCertificateId property value. Trusted Root Certificate ID
func (m *Windows10XSCEPCertificateProfile) GetRootCertificateId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.rootCertificateId
}
// GetScepServerUrls gets the scepServerUrls property value. SCEP Server Url(s).
func (m *Windows10XSCEPCertificateProfile) GetScepServerUrls()([]string) {
    return m.scepServerUrls
}
// GetSubjectAlternativeNameFormats gets the subjectAlternativeNameFormats property value. Custom AAD Attributes.
func (m *Windows10XSCEPCertificateProfile) GetSubjectAlternativeNameFormats()([]Windows10XCustomSubjectAlternativeNameable) {
    return m.subjectAlternativeNameFormats
}
// GetSubjectNameFormatString gets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *Windows10XSCEPCertificateProfile) GetSubjectNameFormatString()(*string) {
    return m.subjectNameFormatString
}
// Serialize serializes information the current object
func (m *Windows10XSCEPCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Windows10XCertificateProfile.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCertificateStore() != nil {
        cast := (*m.GetCertificateStore()).String()
        err = writer.WriteStringValue("certificateStore", &cast)
        if err != nil {
            return err
        }
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
    if m.GetHashAlgorithm() != nil {
        err = writer.WriteCollectionOfStringValues("hashAlgorithm", SerializeHashAlgorithms(m.GetHashAlgorithm()))
        if err != nil {
            return err
        }
    }
    if m.GetKeySize() != nil {
        cast := (*m.GetKeySize()).String()
        err = writer.WriteStringValue("keySize", &cast)
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
    if m.GetKeyUsage() != nil {
        cast := (*m.GetKeyUsage()).String()
        err = writer.WriteStringValue("keyUsage", &cast)
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
        err = writer.WriteUUIDValue("rootCertificateId", m.GetRootCertificateId())
        if err != nil {
            return err
        }
    }
    if m.GetScepServerUrls() != nil {
        err = writer.WriteCollectionOfStringValues("scepServerUrls", m.GetScepServerUrls())
        if err != nil {
            return err
        }
    }
    if m.GetSubjectAlternativeNameFormats() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubjectAlternativeNameFormats()))
        for i, v := range m.GetSubjectAlternativeNameFormats() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("subjectAlternativeNameFormats", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectNameFormatString", m.GetSubjectNameFormatString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateStore sets the certificateStore property value. CertificateStore types
func (m *Windows10XSCEPCertificateProfile) SetCertificateStore(value *CertificateStore)() {
    m.certificateStore = value
}
// SetCertificateValidityPeriodScale sets the certificateValidityPeriodScale property value. Certificate Validity Period Options.
func (m *Windows10XSCEPCertificateProfile) SetCertificateValidityPeriodScale(value *CertificateValidityPeriodScale)() {
    m.certificateValidityPeriodScale = value
}
// SetCertificateValidityPeriodValue sets the certificateValidityPeriodValue property value. Value for the Certificate Validity Period
func (m *Windows10XSCEPCertificateProfile) SetCertificateValidityPeriodValue(value *int32)() {
    m.certificateValidityPeriodValue = value
}
// SetExtendedKeyUsages sets the extendedKeyUsages property value. Extended Key Usage (EKU) settings.
func (m *Windows10XSCEPCertificateProfile) SetExtendedKeyUsages(value []ExtendedKeyUsageable)() {
    m.extendedKeyUsages = value
}
// SetHashAlgorithm sets the hashAlgorithm property value. SCEP Hash Algorithm.
func (m *Windows10XSCEPCertificateProfile) SetHashAlgorithm(value []HashAlgorithms)() {
    m.hashAlgorithm = value
}
// SetKeySize sets the keySize property value. Key Size Options.
func (m *Windows10XSCEPCertificateProfile) SetKeySize(value *KeySize)() {
    m.keySize = value
}
// SetKeyStorageProvider sets the keyStorageProvider property value. Key Storage Provider (KSP) Import Options.
func (m *Windows10XSCEPCertificateProfile) SetKeyStorageProvider(value *KeyStorageProviderOption)() {
    m.keyStorageProvider = value
}
// SetKeyUsage sets the keyUsage property value. Key Usage Options.
func (m *Windows10XSCEPCertificateProfile) SetKeyUsage(value *KeyUsages)() {
    m.keyUsage = value
}
// SetRenewalThresholdPercentage sets the renewalThresholdPercentage property value. Certificate renewal threshold percentage
func (m *Windows10XSCEPCertificateProfile) SetRenewalThresholdPercentage(value *int32)() {
    m.renewalThresholdPercentage = value
}
// SetRootCertificateId sets the rootCertificateId property value. Trusted Root Certificate ID
func (m *Windows10XSCEPCertificateProfile) SetRootCertificateId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.rootCertificateId = value
}
// SetScepServerUrls sets the scepServerUrls property value. SCEP Server Url(s).
func (m *Windows10XSCEPCertificateProfile) SetScepServerUrls(value []string)() {
    m.scepServerUrls = value
}
// SetSubjectAlternativeNameFormats sets the subjectAlternativeNameFormats property value. Custom AAD Attributes.
func (m *Windows10XSCEPCertificateProfile) SetSubjectAlternativeNameFormats(value []Windows10XCustomSubjectAlternativeNameable)() {
    m.subjectAlternativeNameFormats = value
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *Windows10XSCEPCertificateProfile) SetSubjectNameFormatString(value *string)() {
    m.subjectNameFormatString = value
}
