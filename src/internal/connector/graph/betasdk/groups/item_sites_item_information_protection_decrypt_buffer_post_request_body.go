package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemInformationProtectionDecryptBufferPostRequestBody provides operations to call the decryptBuffer method.
type ItemSitesItemInformationProtectionDecryptBufferPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The encryptedBuffer property
    encryptedBuffer []byte
    // The publishingLicense property
    publishingLicense []byte
}
// NewItemSitesItemInformationProtectionDecryptBufferPostRequestBody instantiates a new ItemSitesItemInformationProtectionDecryptBufferPostRequestBody and sets the default values.
func NewItemSitesItemInformationProtectionDecryptBufferPostRequestBody()(*ItemSitesItemInformationProtectionDecryptBufferPostRequestBody) {
    m := &ItemSitesItemInformationProtectionDecryptBufferPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemInformationProtectionDecryptBufferPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemInformationProtectionDecryptBufferPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemInformationProtectionDecryptBufferPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemInformationProtectionDecryptBufferPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEncryptedBuffer gets the encryptedBuffer property value. The encryptedBuffer property
func (m *ItemSitesItemInformationProtectionDecryptBufferPostRequestBody) GetEncryptedBuffer()([]byte) {
    return m.encryptedBuffer
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemInformationProtectionDecryptBufferPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["encryptedBuffer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptedBuffer(val)
        }
        return nil
    }
    res["publishingLicense"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublishingLicense(val)
        }
        return nil
    }
    return res
}
// GetPublishingLicense gets the publishingLicense property value. The publishingLicense property
func (m *ItemSitesItemInformationProtectionDecryptBufferPostRequestBody) GetPublishingLicense()([]byte) {
    return m.publishingLicense
}
// Serialize serializes information the current object
func (m *ItemSitesItemInformationProtectionDecryptBufferPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("encryptedBuffer", m.GetEncryptedBuffer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteByteArrayValue("publishingLicense", m.GetPublishingLicense())
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
func (m *ItemSitesItemInformationProtectionDecryptBufferPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEncryptedBuffer sets the encryptedBuffer property value. The encryptedBuffer property
func (m *ItemSitesItemInformationProtectionDecryptBufferPostRequestBody) SetEncryptedBuffer(value []byte)() {
    m.encryptedBuffer = value
}
// SetPublishingLicense sets the publishingLicense property value. The publishingLicense property
func (m *ItemSitesItemInformationProtectionDecryptBufferPostRequestBody) SetPublishingLicense(value []byte)() {
    m.publishingLicense = value
}
