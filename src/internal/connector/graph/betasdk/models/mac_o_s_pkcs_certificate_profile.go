package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSPkcsCertificateProfile 
type MacOSPkcsCertificateProfile struct {
    MacOSCertificateProfileBase
    // AllowAllAppsAccess setting
    allowAllAppsAccess *bool
    // Target store certificate. Possible values are: user, machine.
    certificateStore *CertificateStore
    // PKCS certificate template name.
    certificateTemplateName *string
    // PKCS certification authority FQDN.
    certificationAuthority *string
    // PKCS certification authority Name.
    certificationAuthorityName *string
    // Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
    customSubjectAlternativeNames []CustomSubjectAlternativeNameable
    // Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
    managedDeviceCertificateStates []ManagedDeviceCertificateStateable
    // Format string that defines the subject alternative name.
    subjectAlternativeNameFormatString *string
    // Format string that defines the subject name. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
    subjectNameFormatString *string
}
// NewMacOSPkcsCertificateProfile instantiates a new MacOSPkcsCertificateProfile and sets the default values.
func NewMacOSPkcsCertificateProfile()(*MacOSPkcsCertificateProfile) {
    m := &MacOSPkcsCertificateProfile{
        MacOSCertificateProfileBase: *NewMacOSCertificateProfileBase(),
    }
    odataTypeValue := "#microsoft.graph.macOSPkcsCertificateProfile";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateMacOSPkcsCertificateProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSPkcsCertificateProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSPkcsCertificateProfile(), nil
}
// GetAllowAllAppsAccess gets the allowAllAppsAccess property value. AllowAllAppsAccess setting
func (m *MacOSPkcsCertificateProfile) GetAllowAllAppsAccess()(*bool) {
    return m.allowAllAppsAccess
}
// GetCertificateStore gets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *MacOSPkcsCertificateProfile) GetCertificateStore()(*CertificateStore) {
    return m.certificateStore
}
// GetCertificateTemplateName gets the certificateTemplateName property value. PKCS certificate template name.
func (m *MacOSPkcsCertificateProfile) GetCertificateTemplateName()(*string) {
    return m.certificateTemplateName
}
// GetCertificationAuthority gets the certificationAuthority property value. PKCS certification authority FQDN.
func (m *MacOSPkcsCertificateProfile) GetCertificationAuthority()(*string) {
    return m.certificationAuthority
}
// GetCertificationAuthorityName gets the certificationAuthorityName property value. PKCS certification authority Name.
func (m *MacOSPkcsCertificateProfile) GetCertificationAuthorityName()(*string) {
    return m.certificationAuthorityName
}
// GetCustomSubjectAlternativeNames gets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *MacOSPkcsCertificateProfile) GetCustomSubjectAlternativeNames()([]CustomSubjectAlternativeNameable) {
    return m.customSubjectAlternativeNames
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSPkcsCertificateProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *MacOSPkcsCertificateProfile) GetManagedDeviceCertificateStates()([]ManagedDeviceCertificateStateable) {
    return m.managedDeviceCertificateStates
}
// GetSubjectAlternativeNameFormatString gets the subjectAlternativeNameFormatString property value. Format string that defines the subject alternative name.
func (m *MacOSPkcsCertificateProfile) GetSubjectAlternativeNameFormatString()(*string) {
    return m.subjectAlternativeNameFormatString
}
// GetSubjectNameFormatString gets the subjectNameFormatString property value. Format string that defines the subject name. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *MacOSPkcsCertificateProfile) GetSubjectNameFormatString()(*string) {
    return m.subjectNameFormatString
}
// Serialize serializes information the current object
func (m *MacOSPkcsCertificateProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAllowAllAppsAccess sets the allowAllAppsAccess property value. AllowAllAppsAccess setting
func (m *MacOSPkcsCertificateProfile) SetAllowAllAppsAccess(value *bool)() {
    m.allowAllAppsAccess = value
}
// SetCertificateStore sets the certificateStore property value. Target store certificate. Possible values are: user, machine.
func (m *MacOSPkcsCertificateProfile) SetCertificateStore(value *CertificateStore)() {
    m.certificateStore = value
}
// SetCertificateTemplateName sets the certificateTemplateName property value. PKCS certificate template name.
func (m *MacOSPkcsCertificateProfile) SetCertificateTemplateName(value *string)() {
    m.certificateTemplateName = value
}
// SetCertificationAuthority sets the certificationAuthority property value. PKCS certification authority FQDN.
func (m *MacOSPkcsCertificateProfile) SetCertificationAuthority(value *string)() {
    m.certificationAuthority = value
}
// SetCertificationAuthorityName sets the certificationAuthorityName property value. PKCS certification authority Name.
func (m *MacOSPkcsCertificateProfile) SetCertificationAuthorityName(value *string)() {
    m.certificationAuthorityName = value
}
// SetCustomSubjectAlternativeNames sets the customSubjectAlternativeNames property value. Custom Subject Alternative Name Settings. This collection can contain a maximum of 500 elements.
func (m *MacOSPkcsCertificateProfile) SetCustomSubjectAlternativeNames(value []CustomSubjectAlternativeNameable)() {
    m.customSubjectAlternativeNames = value
}
// SetManagedDeviceCertificateStates sets the managedDeviceCertificateStates property value. Certificate state for devices. This collection can contain a maximum of 2147483647 elements.
func (m *MacOSPkcsCertificateProfile) SetManagedDeviceCertificateStates(value []ManagedDeviceCertificateStateable)() {
    m.managedDeviceCertificateStates = value
}
// SetSubjectAlternativeNameFormatString sets the subjectAlternativeNameFormatString property value. Format string that defines the subject alternative name.
func (m *MacOSPkcsCertificateProfile) SetSubjectAlternativeNameFormatString(value *string)() {
    m.subjectAlternativeNameFormatString = value
}
// SetSubjectNameFormatString sets the subjectNameFormatString property value. Format string that defines the subject name. Example: CN={{EmailAddress}},E={{EmailAddress}},OU=Enterprise Users,O=Contoso Corporation,L=Redmond,ST=WA,C=US
func (m *MacOSPkcsCertificateProfile) SetSubjectNameFormatString(value *string)() {
    m.subjectNameFormatString = value
}
