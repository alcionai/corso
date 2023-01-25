package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAllDeviceCertificateState provides operations to manage the collection of site entities.
type ManagedAllDeviceCertificateState struct {
    Entity
    // Certificate expiry date
    certificateExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Enhanced Key Usage
    certificateExtendedKeyUsages *string
    // Issuance date
    certificateIssuanceDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Issuer
    certificateIssuerName *string
    // Key Usage
    certificateKeyUsages *int32
    // Certificate Revocation Status.
    certificateRevokeStatus *CertificateRevocationStatus
    // The time the revoke status was last changed
    certificateRevokeStatusLastChangeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Serial number
    certificateSerialNumber *string
    // Certificate subject name
    certificateSubjectName *string
    // Thumbprint
    certificateThumbprint *string
    // Device display name
    managedDeviceDisplayName *string
    // User principal name
    userPrincipalName *string
}
// NewManagedAllDeviceCertificateState instantiates a new managedAllDeviceCertificateState and sets the default values.
func NewManagedAllDeviceCertificateState()(*ManagedAllDeviceCertificateState) {
    m := &ManagedAllDeviceCertificateState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedAllDeviceCertificateStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAllDeviceCertificateStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedAllDeviceCertificateState(), nil
}
// GetCertificateExpirationDateTime gets the certificateExpirationDateTime property value. Certificate expiry date
func (m *ManagedAllDeviceCertificateState) GetCertificateExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.certificateExpirationDateTime
}
// GetCertificateExtendedKeyUsages gets the certificateExtendedKeyUsages property value. Enhanced Key Usage
func (m *ManagedAllDeviceCertificateState) GetCertificateExtendedKeyUsages()(*string) {
    return m.certificateExtendedKeyUsages
}
// GetCertificateIssuanceDateTime gets the certificateIssuanceDateTime property value. Issuance date
func (m *ManagedAllDeviceCertificateState) GetCertificateIssuanceDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.certificateIssuanceDateTime
}
// GetCertificateIssuerName gets the certificateIssuerName property value. Issuer
func (m *ManagedAllDeviceCertificateState) GetCertificateIssuerName()(*string) {
    return m.certificateIssuerName
}
// GetCertificateKeyUsages gets the certificateKeyUsages property value. Key Usage
func (m *ManagedAllDeviceCertificateState) GetCertificateKeyUsages()(*int32) {
    return m.certificateKeyUsages
}
// GetCertificateRevokeStatus gets the certificateRevokeStatus property value. Certificate Revocation Status.
func (m *ManagedAllDeviceCertificateState) GetCertificateRevokeStatus()(*CertificateRevocationStatus) {
    return m.certificateRevokeStatus
}
// GetCertificateRevokeStatusLastChangeDateTime gets the certificateRevokeStatusLastChangeDateTime property value. The time the revoke status was last changed
func (m *ManagedAllDeviceCertificateState) GetCertificateRevokeStatusLastChangeDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.certificateRevokeStatusLastChangeDateTime
}
// GetCertificateSerialNumber gets the certificateSerialNumber property value. Serial number
func (m *ManagedAllDeviceCertificateState) GetCertificateSerialNumber()(*string) {
    return m.certificateSerialNumber
}
// GetCertificateSubjectName gets the certificateSubjectName property value. Certificate subject name
func (m *ManagedAllDeviceCertificateState) GetCertificateSubjectName()(*string) {
    return m.certificateSubjectName
}
// GetCertificateThumbprint gets the certificateThumbprint property value. Thumbprint
func (m *ManagedAllDeviceCertificateState) GetCertificateThumbprint()(*string) {
    return m.certificateThumbprint
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAllDeviceCertificateState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["certificateExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateExpirationDateTime(val)
        }
        return nil
    }
    res["certificateExtendedKeyUsages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateExtendedKeyUsages(val)
        }
        return nil
    }
    res["certificateIssuanceDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateIssuanceDateTime(val)
        }
        return nil
    }
    res["certificateIssuerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateIssuerName(val)
        }
        return nil
    }
    res["certificateKeyUsages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateKeyUsages(val)
        }
        return nil
    }
    res["certificateRevokeStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCertificateRevocationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateRevokeStatus(val.(*CertificateRevocationStatus))
        }
        return nil
    }
    res["certificateRevokeStatusLastChangeDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateRevokeStatusLastChangeDateTime(val)
        }
        return nil
    }
    res["certificateSerialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSerialNumber(val)
        }
        return nil
    }
    res["certificateSubjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSubjectName(val)
        }
        return nil
    }
    res["certificateThumbprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateThumbprint(val)
        }
        return nil
    }
    res["managedDeviceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceDisplayName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetManagedDeviceDisplayName gets the managedDeviceDisplayName property value. Device display name
func (m *ManagedAllDeviceCertificateState) GetManagedDeviceDisplayName()(*string) {
    return m.managedDeviceDisplayName
}
// GetUserPrincipalName gets the userPrincipalName property value. User principal name
func (m *ManagedAllDeviceCertificateState) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *ManagedAllDeviceCertificateState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("certificateExpirationDateTime", m.GetCertificateExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateExtendedKeyUsages", m.GetCertificateExtendedKeyUsages())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("certificateIssuanceDateTime", m.GetCertificateIssuanceDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateIssuerName", m.GetCertificateIssuerName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("certificateKeyUsages", m.GetCertificateKeyUsages())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateRevokeStatus() != nil {
        cast := (*m.GetCertificateRevokeStatus()).String()
        err = writer.WriteStringValue("certificateRevokeStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("certificateRevokeStatusLastChangeDateTime", m.GetCertificateRevokeStatusLastChangeDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateSerialNumber", m.GetCertificateSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateSubjectName", m.GetCertificateSubjectName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("certificateThumbprint", m.GetCertificateThumbprint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceDisplayName", m.GetManagedDeviceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateExpirationDateTime sets the certificateExpirationDateTime property value. Certificate expiry date
func (m *ManagedAllDeviceCertificateState) SetCertificateExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateExpirationDateTime = value
}
// SetCertificateExtendedKeyUsages sets the certificateExtendedKeyUsages property value. Enhanced Key Usage
func (m *ManagedAllDeviceCertificateState) SetCertificateExtendedKeyUsages(value *string)() {
    m.certificateExtendedKeyUsages = value
}
// SetCertificateIssuanceDateTime sets the certificateIssuanceDateTime property value. Issuance date
func (m *ManagedAllDeviceCertificateState) SetCertificateIssuanceDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateIssuanceDateTime = value
}
// SetCertificateIssuerName sets the certificateIssuerName property value. Issuer
func (m *ManagedAllDeviceCertificateState) SetCertificateIssuerName(value *string)() {
    m.certificateIssuerName = value
}
// SetCertificateKeyUsages sets the certificateKeyUsages property value. Key Usage
func (m *ManagedAllDeviceCertificateState) SetCertificateKeyUsages(value *int32)() {
    m.certificateKeyUsages = value
}
// SetCertificateRevokeStatus sets the certificateRevokeStatus property value. Certificate Revocation Status.
func (m *ManagedAllDeviceCertificateState) SetCertificateRevokeStatus(value *CertificateRevocationStatus)() {
    m.certificateRevokeStatus = value
}
// SetCertificateRevokeStatusLastChangeDateTime sets the certificateRevokeStatusLastChangeDateTime property value. The time the revoke status was last changed
func (m *ManagedAllDeviceCertificateState) SetCertificateRevokeStatusLastChangeDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateRevokeStatusLastChangeDateTime = value
}
// SetCertificateSerialNumber sets the certificateSerialNumber property value. Serial number
func (m *ManagedAllDeviceCertificateState) SetCertificateSerialNumber(value *string)() {
    m.certificateSerialNumber = value
}
// SetCertificateSubjectName sets the certificateSubjectName property value. Certificate subject name
func (m *ManagedAllDeviceCertificateState) SetCertificateSubjectName(value *string)() {
    m.certificateSubjectName = value
}
// SetCertificateThumbprint sets the certificateThumbprint property value. Thumbprint
func (m *ManagedAllDeviceCertificateState) SetCertificateThumbprint(value *string)() {
    m.certificateThumbprint = value
}
// SetManagedDeviceDisplayName sets the managedDeviceDisplayName property value. Device display name
func (m *ManagedAllDeviceCertificateState) SetManagedDeviceDisplayName(value *string)() {
    m.managedDeviceDisplayName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name
func (m *ManagedAllDeviceCertificateState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
