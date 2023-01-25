package sites

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInformationProtectionEncryptBufferPostRequestBody provides operations to call the encryptBuffer method.
type ItemInformationProtectionEncryptBufferPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemInformationProtectionEncryptBufferPostRequestBody instantiates a new ItemInformationProtectionEncryptBufferPostRequestBody and sets the default values.
func NewItemInformationProtectionEncryptBufferPostRequestBody()(*ItemInformationProtectionEncryptBufferPostRequestBody) {
    m := &ItemInformationProtectionEncryptBufferPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemInformationProtectionEncryptBufferPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInformationProtectionEncryptBufferPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInformationProtectionEncryptBufferPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionEncryptBufferPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionEncryptBufferPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetBuffer gets the buffer property value. The buffer property
func (m *ItemInformationProtectionEncryptBufferPostRequestBody) GetBuffer()([]byte) {
    return m.GetBackingStore().Get("buffer");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInformationProtectionEncryptBufferPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["buffer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuffer(val)
        }
        return nil
    }
    res["labelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabelId(val)
        }
        return nil
    }
    return res
}
// GetLabelId gets the labelId property value. The labelId property
func (m *ItemInformationProtectionEncryptBufferPostRequestBody) GetLabelId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.GetBackingStore().Get("labelId");
}
// Serialize serializes information the current object
func (m *ItemInformationProtectionEncryptBufferPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("buffer", m.GetBuffer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("labelId", m.GetLabelId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionEncryptBufferPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionEncryptBufferPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetBuffer sets the buffer property value. The buffer property
func (m *ItemInformationProtectionEncryptBufferPostRequestBody) SetBuffer(value []byte)() {
    m.GetBackingStore().Set("buffer", value)
}
// SetLabelId sets the labelId property value. The labelId property
func (m *ItemInformationProtectionEncryptBufferPostRequestBody) SetLabelId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.GetBackingStore().Set("labelId", value)
}
