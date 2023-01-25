package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CryptographySuite vPN Security Association Parameters
type CryptographySuite struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Authentication Transform Constants. Possible values are: md5_96, sha1_96, sha_256_128, aes128Gcm, aes192Gcm, aes256Gcm.
    authenticationTransformConstants *AuthenticationTransformConstant
    // Cipher Transform Constants. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
    cipherTransformConstants *VpnEncryptionAlgorithmType
    // Diffie Hellman Group. Possible values are: group1, group2, group14, ecp256, ecp384, group24.
    dhGroup *DiffieHellmanGroup
    // Encryption Method. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
    encryptionMethod *VpnEncryptionAlgorithmType
    // Integrity Check Method. Possible values are: sha2_256, sha1_96, sha1_160, sha2_384, sha2_512, md5.
    integrityCheckMethod *VpnIntegrityAlgorithmType
    // The OdataType property
    odataType *string
    // Perfect Forward Secrecy Group. Possible values are: pfs1, pfs2, pfs2048, ecp256, ecp384, pfsMM, pfs24.
    pfsGroup *PerfectForwardSecrecyGroup
}
// NewCryptographySuite instantiates a new cryptographySuite and sets the default values.
func NewCryptographySuite()(*CryptographySuite) {
    m := &CryptographySuite{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCryptographySuiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCryptographySuiteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCryptographySuite(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CryptographySuite) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAuthenticationTransformConstants gets the authenticationTransformConstants property value. Authentication Transform Constants. Possible values are: md5_96, sha1_96, sha_256_128, aes128Gcm, aes192Gcm, aes256Gcm.
func (m *CryptographySuite) GetAuthenticationTransformConstants()(*AuthenticationTransformConstant) {
    return m.authenticationTransformConstants
}
// GetCipherTransformConstants gets the cipherTransformConstants property value. Cipher Transform Constants. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
func (m *CryptographySuite) GetCipherTransformConstants()(*VpnEncryptionAlgorithmType) {
    return m.cipherTransformConstants
}
// GetDhGroup gets the dhGroup property value. Diffie Hellman Group. Possible values are: group1, group2, group14, ecp256, ecp384, group24.
func (m *CryptographySuite) GetDhGroup()(*DiffieHellmanGroup) {
    return m.dhGroup
}
// GetEncryptionMethod gets the encryptionMethod property value. Encryption Method. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
func (m *CryptographySuite) GetEncryptionMethod()(*VpnEncryptionAlgorithmType) {
    return m.encryptionMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CryptographySuite) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationTransformConstants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationTransformConstant)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationTransformConstants(val.(*AuthenticationTransformConstant))
        }
        return nil
    }
    res["cipherTransformConstants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnEncryptionAlgorithmType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCipherTransformConstants(val.(*VpnEncryptionAlgorithmType))
        }
        return nil
    }
    res["dhGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiffieHellmanGroup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDhGroup(val.(*DiffieHellmanGroup))
        }
        return nil
    }
    res["encryptionMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnEncryptionAlgorithmType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptionMethod(val.(*VpnEncryptionAlgorithmType))
        }
        return nil
    }
    res["integrityCheckMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVpnIntegrityAlgorithmType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegrityCheckMethod(val.(*VpnIntegrityAlgorithmType))
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
    res["pfsGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePerfectForwardSecrecyGroup)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPfsGroup(val.(*PerfectForwardSecrecyGroup))
        }
        return nil
    }
    return res
}
// GetIntegrityCheckMethod gets the integrityCheckMethod property value. Integrity Check Method. Possible values are: sha2_256, sha1_96, sha1_160, sha2_384, sha2_512, md5.
func (m *CryptographySuite) GetIntegrityCheckMethod()(*VpnIntegrityAlgorithmType) {
    return m.integrityCheckMethod
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CryptographySuite) GetOdataType()(*string) {
    return m.odataType
}
// GetPfsGroup gets the pfsGroup property value. Perfect Forward Secrecy Group. Possible values are: pfs1, pfs2, pfs2048, ecp256, ecp384, pfsMM, pfs24.
func (m *CryptographySuite) GetPfsGroup()(*PerfectForwardSecrecyGroup) {
    return m.pfsGroup
}
// Serialize serializes information the current object
func (m *CryptographySuite) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAuthenticationTransformConstants() != nil {
        cast := (*m.GetAuthenticationTransformConstants()).String()
        err := writer.WriteStringValue("authenticationTransformConstants", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCipherTransformConstants() != nil {
        cast := (*m.GetCipherTransformConstants()).String()
        err := writer.WriteStringValue("cipherTransformConstants", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDhGroup() != nil {
        cast := (*m.GetDhGroup()).String()
        err := writer.WriteStringValue("dhGroup", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEncryptionMethod() != nil {
        cast := (*m.GetEncryptionMethod()).String()
        err := writer.WriteStringValue("encryptionMethod", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIntegrityCheckMethod() != nil {
        cast := (*m.GetIntegrityCheckMethod()).String()
        err := writer.WriteStringValue("integrityCheckMethod", &cast)
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
    if m.GetPfsGroup() != nil {
        cast := (*m.GetPfsGroup()).String()
        err := writer.WriteStringValue("pfsGroup", &cast)
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
func (m *CryptographySuite) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAuthenticationTransformConstants sets the authenticationTransformConstants property value. Authentication Transform Constants. Possible values are: md5_96, sha1_96, sha_256_128, aes128Gcm, aes192Gcm, aes256Gcm.
func (m *CryptographySuite) SetAuthenticationTransformConstants(value *AuthenticationTransformConstant)() {
    m.authenticationTransformConstants = value
}
// SetCipherTransformConstants sets the cipherTransformConstants property value. Cipher Transform Constants. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
func (m *CryptographySuite) SetCipherTransformConstants(value *VpnEncryptionAlgorithmType)() {
    m.cipherTransformConstants = value
}
// SetDhGroup sets the dhGroup property value. Diffie Hellman Group. Possible values are: group1, group2, group14, ecp256, ecp384, group24.
func (m *CryptographySuite) SetDhGroup(value *DiffieHellmanGroup)() {
    m.dhGroup = value
}
// SetEncryptionMethod sets the encryptionMethod property value. Encryption Method. Possible values are: aes256, des, tripleDes, aes128, aes128Gcm, aes256Gcm, aes192, aes192Gcm, chaCha20Poly1305.
func (m *CryptographySuite) SetEncryptionMethod(value *VpnEncryptionAlgorithmType)() {
    m.encryptionMethod = value
}
// SetIntegrityCheckMethod sets the integrityCheckMethod property value. Integrity Check Method. Possible values are: sha2_256, sha1_96, sha1_160, sha2_384, sha2_512, md5.
func (m *CryptographySuite) SetIntegrityCheckMethod(value *VpnIntegrityAlgorithmType)() {
    m.integrityCheckMethod = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CryptographySuite) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPfsGroup sets the pfsGroup property value. Perfect Forward Secrecy Group. Possible values are: pfs1, pfs2, pfs2048, ecp256, ecp384, pfsMM, pfs24.
func (m *CryptographySuite) SetPfsGroup(value *PerfectForwardSecrecyGroup)() {
    m.pfsGroup = value
}
