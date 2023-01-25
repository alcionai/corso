package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInformationProtectionSignDigestPostRequestBody provides operations to call the signDigest method.
type ItemInformationProtectionSignDigestPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemInformationProtectionSignDigestPostRequestBody instantiates a new ItemInformationProtectionSignDigestPostRequestBody and sets the default values.
func NewItemInformationProtectionSignDigestPostRequestBody()(*ItemInformationProtectionSignDigestPostRequestBody) {
    m := &ItemInformationProtectionSignDigestPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemInformationProtectionSignDigestPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInformationProtectionSignDigestPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInformationProtectionSignDigestPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionSignDigestPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionSignDigestPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetDigest gets the digest property value. The digest property
func (m *ItemInformationProtectionSignDigestPostRequestBody) GetDigest()([]byte) {
    return m.GetBackingStore().Get("digest");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInformationProtectionSignDigestPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["digest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDigest(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemInformationProtectionSignDigestPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("digest", m.GetDigest())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionSignDigestPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionSignDigestPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetDigest sets the digest property value. The digest property
func (m *ItemInformationProtectionSignDigestPostRequestBody) SetDigest(value []byte)() {
    m.GetBackingStore().Set("digest", value)
}
