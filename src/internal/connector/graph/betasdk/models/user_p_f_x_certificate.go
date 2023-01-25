package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserPFXCertificate entity that encapsulates all information required for a user's PFX certificates.
type UserPFXCertificate struct {
    Entity
    // Date/time when this PFX certificate was imported.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Encrypted PFX blob.
    encryptedPfxBlob []byte
    // Encrypted PFX password.
    encryptedPfxPassword *string
    // Certificate's validity expiration date/time.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Supported values for the intended purpose of a user PFX certificate.
    intendedPurpose *UserPfxIntendedPurpose
    // Name of the key (within the provider) used to encrypt the blob.
    keyName *string
    // Date/time when this PFX certificate was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Supported values for the padding scheme used by encryption provider.
    paddingScheme *UserPfxPaddingScheme
    // Crypto provider used to encrypt this blob.
    providerName *string
    // Certificate's validity start date/time.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // SHA-1 thumbprint of the PFX certificate.
    thumbprint *string
    // User Principal Name of the PFX certificate.
    userPrincipalName *string
}
// NewUserPFXCertificate instantiates a new userPFXCertificate and sets the default values.
func NewUserPFXCertificate()(*UserPFXCertificate) {
    m := &UserPFXCertificate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserPFXCertificateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserPFXCertificateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserPFXCertificate(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Date/time when this PFX certificate was imported.
func (m *UserPFXCertificate) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetEncryptedPfxBlob gets the encryptedPfxBlob property value. Encrypted PFX blob.
func (m *UserPFXCertificate) GetEncryptedPfxBlob()([]byte) {
    return m.encryptedPfxBlob
}
// GetEncryptedPfxPassword gets the encryptedPfxPassword property value. Encrypted PFX password.
func (m *UserPFXCertificate) GetEncryptedPfxPassword()(*string) {
    return m.encryptedPfxPassword
}
// GetExpirationDateTime gets the expirationDateTime property value. Certificate's validity expiration date/time.
func (m *UserPFXCertificate) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserPFXCertificate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["encryptedPfxBlob"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedPfxBlob(val)
        }
        return nil
    }
    res["encryptedPfxPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedPfxPassword(val)
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
    res["intendedPurpose"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserPfxIntendedPurpose)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntendedPurpose(val.(*UserPfxIntendedPurpose))
        }
        return nil
    }
    res["keyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyName(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["paddingScheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserPfxPaddingScheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPaddingScheme(val.(*UserPfxPaddingScheme))
        }
        return nil
    }
    res["providerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderName(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
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
// GetIntendedPurpose gets the intendedPurpose property value. Supported values for the intended purpose of a user PFX certificate.
func (m *UserPFXCertificate) GetIntendedPurpose()(*UserPfxIntendedPurpose) {
    return m.intendedPurpose
}
// GetKeyName gets the keyName property value. Name of the key (within the provider) used to encrypt the blob.
func (m *UserPFXCertificate) GetKeyName()(*string) {
    return m.keyName
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date/time when this PFX certificate was last modified.
func (m *UserPFXCertificate) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetPaddingScheme gets the paddingScheme property value. Supported values for the padding scheme used by encryption provider.
func (m *UserPFXCertificate) GetPaddingScheme()(*UserPfxPaddingScheme) {
    return m.paddingScheme
}
// GetProviderName gets the providerName property value. Crypto provider used to encrypt this blob.
func (m *UserPFXCertificate) GetProviderName()(*string) {
    return m.providerName
}
// GetStartDateTime gets the startDateTime property value. Certificate's validity start date/time.
func (m *UserPFXCertificate) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetThumbprint gets the thumbprint property value. SHA-1 thumbprint of the PFX certificate.
func (m *UserPFXCertificate) GetThumbprint()(*string) {
    return m.thumbprint
}
// GetUserPrincipalName gets the userPrincipalName property value. User Principal Name of the PFX certificate.
func (m *UserPFXCertificate) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *UserPFXCertificate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("encryptedPfxBlob", m.GetEncryptedPfxBlob())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("encryptedPfxPassword", m.GetEncryptedPfxPassword())
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
    if m.GetIntendedPurpose() != nil {
        cast := (*m.GetIntendedPurpose()).String()
        err = writer.WriteStringValue("intendedPurpose", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("keyName", m.GetKeyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetPaddingScheme() != nil {
        cast := (*m.GetPaddingScheme()).String()
        err = writer.WriteStringValue("paddingScheme", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("providerName", m.GetProviderName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("thumbprint", m.GetThumbprint())
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
// SetCreatedDateTime sets the createdDateTime property value. Date/time when this PFX certificate was imported.
func (m *UserPFXCertificate) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetEncryptedPfxBlob sets the encryptedPfxBlob property value. Encrypted PFX blob.
func (m *UserPFXCertificate) SetEncryptedPfxBlob(value []byte)() {
    m.encryptedPfxBlob = value
}
// SetEncryptedPfxPassword sets the encryptedPfxPassword property value. Encrypted PFX password.
func (m *UserPFXCertificate) SetEncryptedPfxPassword(value *string)() {
    m.encryptedPfxPassword = value
}
// SetExpirationDateTime sets the expirationDateTime property value. Certificate's validity expiration date/time.
func (m *UserPFXCertificate) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetIntendedPurpose sets the intendedPurpose property value. Supported values for the intended purpose of a user PFX certificate.
func (m *UserPFXCertificate) SetIntendedPurpose(value *UserPfxIntendedPurpose)() {
    m.intendedPurpose = value
}
// SetKeyName sets the keyName property value. Name of the key (within the provider) used to encrypt the blob.
func (m *UserPFXCertificate) SetKeyName(value *string)() {
    m.keyName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date/time when this PFX certificate was last modified.
func (m *UserPFXCertificate) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetPaddingScheme sets the paddingScheme property value. Supported values for the padding scheme used by encryption provider.
func (m *UserPFXCertificate) SetPaddingScheme(value *UserPfxPaddingScheme)() {
    m.paddingScheme = value
}
// SetProviderName sets the providerName property value. Crypto provider used to encrypt this blob.
func (m *UserPFXCertificate) SetProviderName(value *string)() {
    m.providerName = value
}
// SetStartDateTime sets the startDateTime property value. Certificate's validity start date/time.
func (m *UserPFXCertificate) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetThumbprint sets the thumbprint property value. SHA-1 thumbprint of the PFX certificate.
func (m *UserPFXCertificate) SetThumbprint(value *string)() {
    m.thumbprint = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User Principal Name of the PFX certificate.
func (m *UserPFXCertificate) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
