package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BufferEncryptionResult 
type BufferEncryptionResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The encryptedBuffer property
    encryptedBuffer []byte
    // The OdataType property
    odataType *string
    // The publishingLicense property
    publishingLicense []byte
}
// NewBufferEncryptionResult instantiates a new bufferEncryptionResult and sets the default values.
func NewBufferEncryptionResult()(*BufferEncryptionResult) {
    m := &BufferEncryptionResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateBufferEncryptionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBufferEncryptionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBufferEncryptionResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BufferEncryptionResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetEncryptedBuffer gets the encryptedBuffer property value. The encryptedBuffer property
func (m *BufferEncryptionResult) GetEncryptedBuffer()([]byte) {
    return m.encryptedBuffer
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BufferEncryptionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BufferEncryptionResult) GetOdataType()(*string) {
    return m.odataType
}
// GetPublishingLicense gets the publishingLicense property value. The publishingLicense property
func (m *BufferEncryptionResult) GetPublishingLicense()([]byte) {
    return m.publishingLicense
}
// Serialize serializes information the current object
func (m *BufferEncryptionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("encryptedBuffer", m.GetEncryptedBuffer())
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
func (m *BufferEncryptionResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetEncryptedBuffer sets the encryptedBuffer property value. The encryptedBuffer property
func (m *BufferEncryptionResult) SetEncryptedBuffer(value []byte)() {
    m.encryptedBuffer = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BufferEncryptionResult) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPublishingLicense sets the publishingLicense property value. The publishingLicense property
func (m *BufferEncryptionResult) SetPublishingLicense(value []byte)() {
    m.publishingLicense = value
}
