package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosPkcsCertificateProfile 
type IosPkcsCertificateProfile struct {
    IosCertificateProfileBase
    // Target store certificate. Possible values are: user, machine.
    certificateStore *CertificateStore
    // PKCS Certificate Template Name.
    certificateTemplateName *string
    // PKCS Certification Authority.
    certificationAuthority *string
    // PKCS Certification Authority Name.
    certificationAuthorityName *string
    // Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
    customSubjectAlternativeNames []CustomSubjectAlternativeNameable
    // Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
    managedDeviceCertificateStates []ManagedDeviceCertificateStateable
    // Custom String that defines the AAD Attribute.
    subjectAlternativeNameFormatString *string
    // Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
    subjectNameFormatString *string
}
// NewIosPkcsCertificateProfile instantiates a new IosPkcsCertificateProfile and sets the default values.
func NewIosPkcsCertificateProfile()(*IosPkcsCertificateProfile) {
    m := &IosPkcsCertificateProfile{
        IosCertificateProfileBase: *NewIosCertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.iosPkcsCertificateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosPkcsCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosPkcsCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosPkcsCertificateProfile(), nil
}
// GetCertificateStore gets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *IosPkcsCertificateProfile) GetCertificateStore()(*CertificateStore) {
    return m.certificateStore
}
// GetCertificateTemplateName gets the certificateTemplateName property value. PKCS Certificate Template Name.
func (m *IosPkcsCertificateProfile) GetCertificateTemplateName()(*string) {
    return m.certificateTemplateName
}
// GetCertificationAuthority gets the certificationAuthority property value. PKCS Certification Authority.
func (m *IosPkcsCertificateProfile) GetCertificationAuthority()(*string) {
    return m.certificationAuthority
}
// GetCertificationAuthorityName gets the certificationAuthorityName property value. PKCS Certification Authority Name.
func (m *IosPkcsCertificateProfile) GetCertificationAuthorityName()(*string) {
    return m.certificationAuthorityName
}
// GetCustomSubjectAlternativeNames gets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *IosPkcsCertificateProfile) GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable) {
    return m.customSubjectAlternativeNames
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosPkcsCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosCertificateProfileBase.GetFieldDeserializers()
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
func (m *IosPkcsCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    return m.managedDeviceCertificateStates
}
// GetSubjectAlternativeNameFormatString gets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *IosPkcsCertificateProfile) GetSubjectAlternativeNameFormatString()(*string) {
    return m.subjectAlternativeNameFormatString
}
// GetSubjectNameFormatString gets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *IosPkcsCertificateProfile) GetSubjectNameFormatString()(*string) {
    return m.subjectNameFormatString
}
// Serialize serializes information the current object
func (m *IosPkcsCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosCertificateProfileBase.Serialize(writer)
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
// SetCertificateStore sets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *IosPkcsCertificateProfile) SetCertificateStore(value *CertificateStore)() {
    m.certificateStore = value
}
// SetCertificateTemplateName sets the certificateTemplateName property value. PKCS Certificate Template Name.
func (m *IosPkcsCertificateProfile) SetCertificateTemplateName(value *string)() {
    m.certificateTemplateName = value
}
// SetCertificationAuthority sets the certificationAuthority property value. PKCS Certification Authority.
func (m *IosPkcsCertificateProfile) SetCertificationAuthority(value *string)() {
    m.certificationAuthority = value
}
// SetCertificationAuthorityName sets the certificationAuthorityName property value. PKCS Certification Authority Name.
func (m *IosPkcsCertificateProfile) SetCertificationAuthorityName(value *string)() {
    m.certificationAuthorityName = value
}
// SetCustomSubjectAlternativeNames sets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *IosPkcsCertificateProfile) SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)() {
    m.customSubjectAlternativeNames = value
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *IosPkcsCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    m.managedDeviceCertificateStates = value
}
// SetSubjectAlternativeNameFormatString sets the subjectAlternativeNameFormatString property value. Custom String that defines the AAD Attribute.
func (m *IosPkcsCertificateProfile) SetSubjectAlternativeNameFormatString(value *string)() {
    m.subjectAlternativeNameFormatString = value
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Custom format to use with SubjectNameFormat = Custom. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *IosPkcsCertificateProfile) SetSubjectNameFormatString(value *string)() {
    m.subjectNameFormatString = value
}
