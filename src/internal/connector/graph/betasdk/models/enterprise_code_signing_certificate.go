package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnterpriseCodeSigningCertificate provides operations to manage the collection of site entities.
type EnterpriseCodeSigningCertificate struct {
    Entity
    // The Windows Enterprise Code-Signing Certificate in the raw data format.
    content []byte
    // The Cert Expiration Date.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Issuer value for the cert.
    issuer *string
    // The Issuer Name for the cert.
    issuerName *string
    // The status property
    status *CertificateStatus
    // The Subject Value for the cert.
    subject *string
    // The Subject Name for the cert.
    subjectName *string
    // The date time of CodeSigning Cert when it is uploaded.
    uploadDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewEnterpriseCodeSigningCertificate instantiates a new enterpriseCodeSigningCertificate and sets the default values.
func NewEnterpriseCodeSigningCertificate()(*EnterpriseCodeSigningCertificate) {
    m := &EnterpriseCodeSigningCertificate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEnterpriseCodeSigningCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnterpriseCodeSigningCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnterpriseCodeSigningCertificate(), nil
}
// GetContent gets the content property value. The Windows Enterprise Code-Signing Certificate in the raw data format.
func (m *EnterpriseCodeSigningCertificate) GetContent()([]byte) {
    return m.content
}
// GetExpirationDateTime gets the expirationDateTime property value. The Cert Expiration Date.
func (m *EnterpriseCodeSigningCertificate) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnterpriseCodeSigningCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["issuerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerName(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*CertificateStatus))
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["subjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectName(val)
        }
        return nil
    }
    res["uploadDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUploadDateTime(val)
        }
        return nil
    }
    return res
}
// GetIssuer gets the issuer property value. The Issuer value for the cert.
func (m *EnterpriseCodeSigningCertificate) GetIssuer()(*string) {
    return m.issuer
}
// GetIssuerName gets the issuerName property value. The Issuer Name for the cert.
func (m *EnterpriseCodeSigningCertificate) GetIssuerName()(*string) {
    return m.issuerName
}
// GetStatus gets the status property value. The status property
func (m *EnterpriseCodeSigningCertificate) GetStatus()(*CertificateStatus) {
    return m.status
}
// GetSubject gets the subject property value. The Subject Value for the cert.
func (m *EnterpriseCodeSigningCertificate) GetSubject()(*string) {
    return m.subject
}
// GetSubjectName gets the subjectName property value. The Subject Name for the cert.
func (m *EnterpriseCodeSigningCertificate) GetSubjectName()(*string) {
    return m.subjectName
}
// GetUploadDateTime gets the uploadDateTime property value. The date time of CodeSigning Cert when it is uploaded.
func (m *EnterpriseCodeSigningCertificate) GetUploadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.uploadDateTime
}
// Serialize serializes information the current object
func (m *EnterpriseCodeSigningCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuerName", m.GetIssuerName())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subjectName", m.GetSubjectName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("uploadDateTime", m.GetUploadDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The Windows Enterprise Code-Signing Certificate in the raw data format.
func (m *EnterpriseCodeSigningCertificate) SetContent(value []byte)() {
    m.content = value
}
// SetExpirationDateTime sets the expirationDateTime property value. The Cert Expiration Date.
func (m *EnterpriseCodeSigningCertificate) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetIssuer sets the issuer property value. The Issuer value for the cert.
func (m *EnterpriseCodeSigningCertificate) SetIssuer(value *string)() {
    m.issuer = value
}
// SetIssuerName sets the issuerName property value. The Issuer Name for the cert.
func (m *EnterpriseCodeSigningCertificate) SetIssuerName(value *string)() {
    m.issuerName = value
}
// SetStatus sets the status property value. The status property
func (m *EnterpriseCodeSigningCertificate) SetStatus(value *CertificateStatus)() {
    m.status = value
}
// SetSubject sets the subject property value. The Subject Value for the cert.
func (m *EnterpriseCodeSigningCertificate) SetSubject(value *string)() {
    m.subject = value
}
// SetSubjectName sets the subjectName property value. The Subject Name for the cert.
func (m *EnterpriseCodeSigningCertificate) SetSubjectName(value *string)() {
    m.subjectName = value
}
// SetUploadDateTime sets the uploadDateTime property value. The date time of CodeSigning Cert when it is uploaded.
func (m *EnterpriseCodeSigningCertificate) SetUploadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.uploadDateTime = value
}
