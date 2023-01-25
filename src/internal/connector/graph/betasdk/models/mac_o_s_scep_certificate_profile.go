package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSScepCertificateProfile 
type MacOSScepCertificateProfile struct {
    MacOSCertificateProfileBase
    // AllowAllAppsAccess setting
    allowAllAppsAccess *bool
    // Target store certificate. Possible values are: user, machine.
    certificateStore *CertificateStore
    // Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
    customSubjectAlternativeNames []CustomSubjectAlternativeNameable
    // Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
    extendedKeyUsages []ExtendedKeyUsageable
    // SCEP Hash Algorithm. Possible values are: sha1, sha2.
    hashAlgorithm *HashAlgorithms
    // Key Size Options.
    keySize *KeySize
    // Key Usage Options.
    keyUsage *KeyUsages
    // Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
    managedDeviceCertificateStates []ManagedDeviceCertificateStateable
    // Trusted Root Certificate.
    rootCertificate MacOSTrustedRootCertificateable
    // SCEP Server Url(s).
    scepServerUrls []string
    // Custom String that defines the AAD Attribute.
    subjectAlternativeNameFormatString *string
    // Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
    subjectNameFormatString *string
}
// NewMacOSScepCertificateProfile instantiates a new MacOSScepCertificateProfile and sets the default values.
func NewMacOSScepCertificateProfile()(*MacOSScepCertificateProfile) {
    m := &MacOSScepCertificateProfile{
        MacOSCertificateProfileBase: *NewMacOSCertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.macOSScepCertificateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSScepCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSScepCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSScepCertificateProfile(), nil
}
// GetAllowAllAppsAccess gets the allowAllAppsAccess property value. AllowAllAppsAccess setting
func (m *MacOSScepCertificateProfile) GetAllowAllAppsAccess()(*bool) {
    return m.allowAllAppsAccess
}
// GetCertificateStore gets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *MacOSScepCertificateProfile) GetCertificateStore()(*CertificateStore) {
    return m.certificateStore
}
// GetCustomSubjectAlternativeNames gets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *MacOSScepCertificateProfile) GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable) {
    return m.customSubjectAlternativeNames
}
// GetExtendedKeyUsages gets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *MacOSScepCertificateProfile) GetExtendedKeyUsages()([]ExtendedKeyUsageable) {
    return m.extendedKeyUsages
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSScepCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MacOSCertificateProfileBase.GetFieldDeserializers()
    res["allowAllAppsAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAllAppsAccess(val)
        }
        return nil
    }
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
    res["customSubjectAlternativeNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCustomSubjectAlternativeNameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CustomSubjectAlternativeNameable, len(val))
            for i, v := range val {
                res[i] = v.(CustomSubjectAlternativeNameable)
            }
            m.SetCustomSubjectAlternativeNames(res)
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
        val, err := n.GetEnumValue(ParseHashAlgorithms)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHashAlgorithm(val.(*HashAlgorithms))
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
    res["managedDeviceCertificateStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceCertificateStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceCertificateStateable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceCertificateStateable)
            }
            m.SetManagedDeviceCertificateStates(res)
        }
        return nil
    }
    res["rootCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMacOSTrustedRootCertificateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootCertificate(val.(MacOSTrustedRootCertificateable))
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
    res["subjectAlternativeNameFormatString"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectAlternativeNameFormatString(val)
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
// GetHashAlgorithm gets the hashAlgorithm property value. SCEP Hash Algorithm. Possible values are: sha1, sha2.
func (m *MacOSScepCertificateProfile) GetHashAlgorithm()(*HashAlgorithms) {
    return m.hashAlgorithm
}
// GetKeySize gets the keySize property value. Key Size Options.
func (m *MacOSScepCertificateProfile) GetKeySize()(*KeySize) {
    return m.keySize
}
// GetKeyUsage gets the keyUsage property value. Key Usage Options.
func (m *MacOSScepCertificateProfile) GetKeyUsage()(*KeyUsages) {
    return m.keyUsage
}
// GetManagedDeviceCertificateStates gets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *MacOSScepCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    return m.managedDeviceCertificateStates
}
// GetRootCertificate gets the rootCertificate property value. Trusted Root Certificate.
func (m *MacOSScepCertificateProfile) GetRootCertificate()(MacOSTrustedRootCertificateable) {
    return m.rootCertificate
}
// GetScepServerUrls gets the scepServerUrls property value. SCEP Server Url(s).
func (m *MacOSScepCertificateProfile) GetScepServerUrls()([]string) {
    return m.scepServerUrls
}
// GetSubjectAlternativeNameFormatString gets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *MacOSScepCertificateProfile) GetSubjectAlternativeNameFormatString()(*string) {
    return m.subjectAlternativeNameFormatString
}
// GetSubjectNameFormatString gets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *MacOSScepCertificateProfile) GetSubjectNameFormatString()(*string) {
    return m.subjectNameFormatString
}
// Serialize serializes information the current object
func (m *MacOSScepCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MacOSCertificateProfileBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowAllAppsAccess", m.GetAllowAllAppsAccess())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateStore() != nil {
        cast := (*m.GetCertificateStore()).String()
        err = writer.WriteStringValue("certificateStore", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomSubjectAlternativeNames() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomSubjectAlternativeNames()))
        for i, v := range m.GetCustomSubjectAlternativeNames() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customSubjectAlternativeNames", cast)
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
        cast := (*m.GetHashAlgorithm()).String()
        err = writer.WriteStringValue("hashAlgorithm", &cast)
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
    if m.GetKeyUsage() != nil {
        cast := (*m.GetKeyUsage()).String()
        err = writer.WriteStringValue("keyUsage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedDeviceCertificateStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDeviceCertificateStates()))
        for i, v := range m.GetManagedDeviceCertificateStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDeviceCertificateStates", cast)
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
    if m.GetScepServerUrls() != nil {
        err = writer.WriteCollectionOfStringValues("scepServerUrls", m.GetScepServerUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectAlternativeNameFormatString", m.GetSubjectAlternativeNameFormatString())
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
// SetAllowAllAppsAccess sets the allowAllAppsAccess property value. AllowAllAppsAccess setting
func (m *MacOSScepCertificateProfile) SetAllowAllAppsAccess(value *bool)() {
    m.allowAllAppsAccess = value
}
// SetCertificateStore sets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *MacOSScepCertificateProfile) SetCertificateStore(value *CertificateStore)() {
    m.certificateStore = value
}
// SetCustomSubjectAlternativeNames sets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *MacOSScepCertificateProfile) SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)() {
    m.customSubjectAlternativeNames = value
}
// SetExtendedKeyUsages sets the extendedKeyUsages property value. Extended Key Usage (EKU) settings. This collection can contain a maximum of 500 elements.
func (m *MacOSScepCertificateProfile) SetExtendedKeyUsages(value []ExtendedKeyUsageable)() {
    m.extendedKeyUsages = value
}
// SetHashAlgorithm sets the hashAlgorithm property value. SCEP Hash Algorithm. Possible values are: sha1, sha2.
func (m *MacOSScepCertificateProfile) SetHashAlgorithm(value *HashAlgorithms)() {
    m.hashAlgorithm = value
}
// SetKeySize sets the keySize property value. Key Size Options.
func (m *MacOSScepCertificateProfile) SetKeySize(value *KeySize)() {
    m.keySize = value
}
// SetKeyUsage sets the keyUsage property value. Key Usage Options.
func (m *MacOSScepCertificateProfile) SetKeyUsage(value *KeyUsages)() {
    m.keyUsage = value
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *MacOSScepCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    m.managedDeviceCertificateStates = value
}
// SetRootCertificate sets the rootCertificate property value. Trusted Root Certificate.
func (m *MacOSScepCertificateProfile) SetRootCertificate(value MacOSTrustedRootCertificateable)() {
    m.rootCertificate = value
}
// SetScepServerUrls sets the scepServerUrls property value. SCEP Server Url(s).
func (m *MacOSScepCertificateProfile) SetScepServerUrls(value []string)() {
    m.scepServerUrls = value
}
// SetSubjectAlternativeNameFormatString sets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *MacOSScepCertificateProfile) SetSubjectAlternativeNameFormatString(value *string)() {
    m.subjectAlternativeNameFormatString = value
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *MacOSScepCertificateProfile) SetSubjectNameFormatString(value *string)() {
    m.subjectNameFormatString = value
}
