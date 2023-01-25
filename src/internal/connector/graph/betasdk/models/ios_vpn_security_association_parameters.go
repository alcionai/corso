package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVpnSecurityAssociationParameters vPN Security Association Parameters
type IosVpnSecurityAssociationParameters struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Lifetime (minutes)
    lifetimeInMinutes *int32
    // The OdataType property
    odataType *string
    // Diffie-Hellman Group
    securityDiffieHellmanGroup *int32
    // Encryption algorithm. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
    securityEncryptionAlgorithm *VpnEncryptionAlgorithmType
    // Integrity algorithm. Possible values are: sha2_256, sha1_96, sha1_160, sha2_384, sha2_512, md5.
    securityIntegrityAlgorithm *VpnIntegrityAlgorithmType
}
// NewIosVpnSecurityAssociationParameters instantiates a new iosVpnSecurityAssociationParameters and sets the default values.
func NewIosVpnSecurityAssociationParameters()(*IosVpnSecurityAssociationParameters) {
    m := &IosVpnSecurityAssociationParameters{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateIosVpnSecurityAssociationParametersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosVpnSecurityAssociationParametersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosVpnSecurityAssociationParameters(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IosVpnSecurityAssociationParameters) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosVpnSecurityAssociationParameters) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["lifetimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLifetimeInMinutes(val)
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
    res["securityDiffieHellmanGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityDiffieHellmanGroup(val)
        }
        return nil
    }
    res["securityEncryptionAlgorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnEncryptionAlgorithmType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityEncryptionAlgorithm(val.(*VpnEncryptionAlgorithmType))
        }
        return nil
    }
    res["securityIntegrityAlgorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnIntegrityAlgorithmType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityIntegrityAlgorithm(val.(*VpnIntegrityAlgorithmType))
        }
        return nil
    }
    return res
}
// GetLifetimeInMinutes gets the lifetimeInMinutes property value. Lifetime (minutes)
func (m *IosVpnSecurityAssociationParameters) GetLifetimeInMinutes()(*int32) {
    return m.lifetimeInMinutes
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IosVpnSecurityAssociationParameters) GetOdataType()(*string) {
    return m.odataType
}
// GetSecurityDiffieHellmanGroup gets the securityDiffieHellmanGroup property value. Diffie-Hellman Group
func (m *IosVpnSecurityAssociationParameters) GetSecurityDiffieHellmanGroup()(*int32) {
    return m.securityDiffieHellmanGroup
}
// GetSecurityEncryptionAlgorithm gets the securityEncryptionAlgorithm property value. Encryption algorithm. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
func (m *IosVpnSecurityAssociationParameters) GetSecurityEncryptionAlgorithm()(*VpnEncryptionAlgorithmType) {
    return m.securityEncryptionAlgorithm
}
// GetSecurityIntegrityAlgorithm gets the securityIntegrityAlgorithm property value. Integrity algorithm. Possible values are: sha2_256, sha1_96, sha1_160, sha2_384, sha2_512, md5.
func (m *IosVpnSecurityAssociationParameters) GetSecurityIntegrityAlgorithm()(*VpnIntegrityAlgorithmType) {
    return m.securityIntegrityAlgorithm
}
// Serialize serializes information the current object
func (m *IosVpnSecurityAssociationParameters) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("lifetimeInMinutes", m.GetLifetimeInMinutes())
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
        err := writer.WriteInt32Value("securityDiffieHellmanGroup", m.GetSecurityDiffieHellmanGroup())
        if err != nil {
            return err
        }
    }
    if m.GetSecurityEncryptionAlgorithm() != nil {
        cast := (*m.GetSecurityEncryptionAlgorithm()).String()
        err := writer.WriteStringValue("securityEncryptionAlgorithm", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityIntegrityAlgorithm() != nil {
        cast := (*m.GetSecurityIntegrityAlgorithm()).String()
        err := writer.WriteStringValue("securityIntegrityAlgorithm", &cast)
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
func (m *IosVpnSecurityAssociationParameters) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetLifetimeInMinutes sets the lifetimeInMinutes property value. Lifetime (minutes)
func (m *IosVpnSecurityAssociationParameters) SetLifetimeInMinutes(value *int32)() {
    m.lifetimeInMinutes = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IosVpnSecurityAssociationParameters) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSecurityDiffieHellmanGroup sets the securityDiffieHellmanGroup property value. Diffie-Hellman Group
func (m *IosVpnSecurityAssociationParameters) SetSecurityDiffieHellmanGroup(value *int32)() {
    m.securityDiffieHellmanGroup = value
}
// SetSecurityEncryptionAlgorithm sets the securityEncryptionAlgorithm property value. Encryption algorithm. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
func (m *IosVpnSecurityAssociationParameters) SetSecurityEncryptionAlgorithm(value *VpnEncryptionAlgorithmType)() {
    m.securityEncryptionAlgorithm = value
}
// SetSecurityIntegrityAlgorithm sets the securityIntegrityAlgorithm property value. Integrity algorithm. Possible values are: sha2_256, sha1_96, sha1_160, sha2_384, sha2_512, md5.
func (m *IosVpnSecurityAssociationParameters) SetSecurityIntegrityAlgorithm(value *VpnIntegrityAlgorithmType)() {
    m.securityIntegrityAlgorithm = value
}
