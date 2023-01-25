package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VerifiedCustomDomainCertificatesMetadata 
type VerifiedCustomDomainCertificatesMetadata struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The expiry date of the custom domain certificate. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    expiryDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The issue date of the custom domain. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    issueDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The issuer name of the custom domain certificate.
    issuerName *string
    // The OdataType property
    odataType *string
    // The subject name of the custom domain certificate.
    subjectName *string
    // The thumbprint associated with the custom domain certificate.
    thumbprint *string
}
// NewVerifiedCustomDomainCertificatesMetadata instantiates a new verifiedCustomDomainCertificatesMetadata and sets the default values.
func NewVerifiedCustomDomainCertificatesMetadata()(*VerifiedCustomDomainCertificatesMetadata) {
    m := &VerifiedCustomDomainCertificatesMetadata{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVerifiedCustomDomainCertificatesMetadataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVerifiedCustomDomainCertificatesMetadataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVerifiedCustomDomainCertificatesMetadata(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifiedCustomDomainCertificatesMetadata) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetExpiryDate gets the expiryDate property value. The expiry date of the custom domain certificate. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *VerifiedCustomDomainCertificatesMetadata) GetExpiryDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expiryDate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VerifiedCustomDomainCertificatesMetadata) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["expiryDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiryDate(val)
        }
        return nil
    }
    res["issueDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssueDate(val)
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
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
    res["thumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbprint(val)
        }
        return nil
    }
    return res
}
// GetIssueDate gets the issueDate property value. The issue date of the custom domain. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *VerifiedCustomDomainCertificatesMetadata) GetIssueDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.issueDate
}
// GetIssuerName gets the issuerName property value. The issuer name of the custom domain certificate.
func (m *VerifiedCustomDomainCertificatesMetadata) GetIssuerName()(*string) {
    return m.issuerName
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VerifiedCustomDomainCertificatesMetadata) GetOdataType()(*string) {
    return m.odataType
}
// GetSubjectName gets the subjectName property value. The subject name of the custom domain certificate.
func (m *VerifiedCustomDomainCertificatesMetadata) GetSubjectName()(*string) {
    return m.subjectName
}
// GetThumbprint gets the thumbprint property value. The thumbprint associated with the custom domain certificate.
func (m *VerifiedCustomDomainCertificatesMetadata) GetThumbprint()(*string) {
    return m.thumbprint
}
// Serialize serializes information the current object
func (m *VerifiedCustomDomainCertificatesMetadata) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("expiryDate", m.GetExpiryDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("issueDate", m.GetIssueDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("issuerName", m.GetIssuerName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subjectName", m.GetSubjectName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("thumbprint", m.GetThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifiedCustomDomainCertificatesMetadata) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetExpiryDate sets the expiryDate property value. The expiry date of the custom domain certificate. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *VerifiedCustomDomainCertificatesMetadata) SetExpiryDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiryDate = value
}
// SetIssueDate sets the issueDate property value. The issue date of the custom domain. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *VerifiedCustomDomainCertificatesMetadata) SetIssueDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.issueDate = value
}
// SetIssuerName sets the issuerName property value. The issuer name of the custom domain certificate.
func (m *VerifiedCustomDomainCertificatesMetadata) SetIssuerName(value *string)() {
    m.issuerName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VerifiedCustomDomainCertificatesMetadata) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSubjectName sets the subjectName property value. The subject name of the custom domain certificate.
func (m *VerifiedCustomDomainCertificatesMetadata) SetSubjectName(value *string)() {
    m.subjectName = value
}
// SetThumbprint sets the thumbprint property value. The thumbprint associated with the custom domain certificate.
func (m *VerifiedCustomDomainCertificatesMetadata) SetThumbprint(value *string)() {
    m.thumbprint = value
}
