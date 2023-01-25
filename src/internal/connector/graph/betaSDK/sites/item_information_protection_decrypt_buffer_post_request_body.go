package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInformationProtectionDecryptBufferPostRequestBody provides operations to call the decryptBuffer method.
type ItemInformationProtectionDecryptBufferPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemInformationProtectionDecryptBufferPostRequestBody instantiates a new ItemInformationProtectionDecryptBufferPostRequestBody and sets the default values.
func NewItemInformationProtectionDecryptBufferPostRequestBody()(*ItemInformationProtectionDecryptBufferPostRequestBody) {
    m := &ItemInformationProtectionDecryptBufferPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemInformationProtectionDecryptBufferPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInformationProtectionDecryptBufferPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInformationProtectionDecryptBufferPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionDecryptBufferPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionDecryptBufferPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetEncryptedBuffer gets the encryptedBuffer property value. The encryptedBuffer property
func (m *ItemInformationProtectionDecryptBufferPostRequestBody) GetEncryptedBuffer()([]byte) {
    return m.GetBackingStore().Get("encryptedBuffer");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInformationProtectionDecryptBufferPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemInformationProtectionDecryptBufferPostRequestBody) GetPublishingLicense()([]byte) {
    return m.GetBackingStore().Get("publishingLicense");
}
// Serialize serializes information the current object
func (m *ItemInformationProtectionDecryptBufferPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionDecryptBufferPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionDecryptBufferPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetEncryptedBuffer sets the encryptedBuffer property value. The encryptedBuffer property
func (m *ItemInformationProtectionDecryptBufferPostRequestBody) SetEncryptedBuffer(value []byte)() {
    m.GetBackingStore().Set("encryptedBuffer", value)
}
// SetPublishingLicense sets the publishingLicense property value. The publishingLicense property
func (m *ItemInformationProtectionDecryptBufferPostRequestBody) SetPublishingLicense(value []byte)() {
    m.GetBackingStore().Set("publishingLicense", value)
}
