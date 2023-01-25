package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AospDeviceOwnerPkcsCertificateProfile 
type AospDeviceOwnerPkcsCertificateProfile struct {
    AospDeviceOwnerCertificateProfileBase
    // CertificateStore types
    certificateStore *CertificateStore
    // PKCS Certificate Template Name
    certificateTemplateName *string
    // PKCS Certification Authority
    certificationAuthority *string
    // PKCS Certification Authority Name
    certificationAuthorityName *string
    // Device Management Certification Authority Types.
    certificationAuthorityType *DeviceManagementCertificationAuthority
    // Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
    customSubjectAlternativeNames []CustomSubjectAlternativeNameable
    // Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
    managedDeviceCertificateStates []ManagedDeviceCertificateStateable
    // Custom String that defines the AAD Attribute.
    subjectAlternativeNameFormatString *string
    // Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
    subjectNameFormatString *string
}
// NewAospDeviceOwnerPkcsCertificateProfile instantiates a new AospDeviceOwnerPkcsCertificateProfile and sets the default values.
func NewAospDeviceOwnerPkcsCertificateProfile()(*AospDeviceOwnerPkcsCertificateProfile) {
    m := &AospDeviceOwnerPkcsCertificateProfile{
        AospDeviceOwnerCertificateProfileBase: *NewAospDeviceOwnerCertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.aospDeviceOwnerPkcsCertificateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAospDeviceOwnerPkcsCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAospDeviceOwnerPkcsCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAospDeviceOwnerPkcsCertificateProfile(), nil
}
// GetCertificateStore gets the certificateStore property value. CertificateStore types
func (m *AospDeviceOwnerPkcsCertificateProfile) GetCertificateStore()(*CertificateStore) {
    return m.certificateStore
}
// GetCertificateTemplateName gets the certificateTemplateName property value. PKCS Certificate Template Name
func (m *AospDeviceOwnerPkcsCertificateProfile) GetCertificateTemplateName()(*string) {
    return m.certificateTemplateName
}
// GetCertificationAuthority gets the certificationAuthority property value. PKCS Certification Authority
func (m *AospDeviceOwnerPkcsCertificateProfile) GetCertificationAuthority()(*string) {
    return m.certificationAuthority
}
// GetCertificationAuthorityName gets the certificationAuthorityName property value. PKCS Certification Authority Name
func (m *AospDeviceOwnerPkcsCertificateProfile) GetCertificationAuthorityName()(*string) {
    return m.certificationAuthorityName
}
// GetCertificationAuthorityType gets the certificationAuthorityType property value. Device Management Certification Authority Types.
func (m *AospDeviceOwnerPkcsCertificateProfile) GetCertificationAuthorityType()(*DeviceManagementCertificationAuthority) {
    return m.certificationAuthorityType
}
// GetCustomSubjectAlternativeNames gets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *AospDeviceOwnerPkcsCertificateProfile) GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable) {
    return m.customSubjectAlternativeNames
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AospDeviceOwnerPkcsCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AospDeviceOwnerCertificateProfileBase.GetFieldDeserializers()
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
    res["certificateTemplateName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateTemplateName(val)
        }
        return nil
    }
    res["certificationAuthority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthority(val)
        }
        return nil
    }
    res["certificationAuthorityName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthorityName(val)
        }
        return nil
    }
    res["certificationAuthorityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementCertificationAuthority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificationAuthorityType(val.(*DeviceManagementCertificationAuthority))
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
// GetManagedDeviceCertificateStates gets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *AospDeviceOwnerPkcsCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    return m.managedDeviceCertificateStates
}
// GetSubjectAlternativeNameFormatString gets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *AospDeviceOwnerPkcsCertificateProfile) GetSubjectAlternativeNameFormatString()(*string) {
    return m.subjectAlternativeNameFormatString
}
// GetSubjectNameFormatString gets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *AospDeviceOwnerPkcsCertificateProfile) GetSubjectNameFormatString()(*string) {
    return m.subjectNameFormatString
}
// Serialize serializes information the current object
func (m *AospDeviceOwnerPkcsCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AospDeviceOwnerCertificateProfileBase.Serialize(writer)
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
    {
        err = writer.WriteStringValue("certificateTemplateName", m.GetCertificateTemplateName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificationAuthority", m.GetCertificationAuthority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificationAuthorityName", m.GetCertificationAuthorityName())
        if err != nil {
            return err
        }
    }
    if m.GetCertificationAuthorityType() != nil {
        cast := (*m.GetCertificationAuthorityType()).String()
        err = writer.WriteStringValue("certificationAuthorityType", &cast)
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
// SetCertificateStore sets the certificateStore property value. CertificateStore types
func (m *AospDeviceOwnerPkcsCertificateProfile) SetCertificateStore(value *CertificateStore)() {
    m.certificateStore = value
}
// SetCertificateTemplateName sets the certificateTemplateName property value. PKCS Certificate Template Name
func (m *AospDeviceOwnerPkcsCertificateProfile) SetCertificateTemplateName(value *string)() {
    m.certificateTemplateName = value
}
// SetCertificationAuthority sets the certificationAuthority property value. PKCS Certification Authority
func (m *AospDeviceOwnerPkcsCertificateProfile) SetCertificationAuthority(value *string)() {
    m.certificationAuthority = value
}
// SetCertificationAuthorityName sets the certificationAuthorityName property value. PKCS Certification Authority Name
func (m *AospDeviceOwnerPkcsCertificateProfile) SetCertificationAuthorityName(value *string)() {
    m.certificationAuthorityName = value
}
// SetCertificationAuthorityType sets the certificationAuthorityType property value. Device Management Certification Authority Types.
func (m *AospDeviceOwnerPkcsCertificateProfile) SetCertificationAuthorityType(value *DeviceManagementCertificationAuthority)() {
    m.certificationAuthorityType = value
}
// SetCustomSubjectAlternativeNames sets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *AospDeviceOwnerPkcsCertificateProfile) SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)() {
    m.customSubjectAlternativeNames = value
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *AospDeviceOwnerPkcsCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    m.managedDeviceCertificateStates = value
}
// SetSubjectAlternativeNameFormatString sets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *AospDeviceOwnerPkcsCertificateProfile) SetSubjectAlternativeNameFormatString(value *string)() {
    m.subjectAlternativeNameFormatString = value
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *AospDeviceOwnerPkcsCertificateProfile) SetSubjectNameFormatString(value *string)() {
    m.subjectNameFormatString = value
}
