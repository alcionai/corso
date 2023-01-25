package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemInformationProtectionSignDigestPostRequestBody provides operations to call the signDigest method.
type ItemSitesItemInformationProtectionSignDigestPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemSitesItemInformationProtectionSignDigestPostRequestBody instantiates a new ItemSitesItemInformationProtectionSignDigestPostRequestBody and sets the default values.
func NewItemSitesItemInformationProtectionSignDigestPostRequestBody()(*ItemSitesItemInformationProtectionSignDigestPostRequestBody) {
    m := &ItemSitesItemInformationProtectionSignDigestPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemInformationProtectionSignDigestPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemInformationProtectionSignDigestPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemInformationProtectionSignDigestPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemInformationProtectionSignDigestPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemSitesItemInformationProtectionSignDigestPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetDigest gets the digest property value. The digest property
func (m *ItemSitesItemInformationProtectionSignDigestPostRequestBody) GetDigest()([]byte) {
    return m.GetBackingStore().Get("digest");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemInformationProtectionSignDigestPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *ItemSitesItemInformationProtectionSignDigestPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("digest", m.GetDigest())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemInformationProtectionSignDigestPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemSitesItemInformationProtectionSignDigestPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetDigest sets the digest property value. The digest property
func (m *ItemSitesItemInformationProtectionSignDigestPostRequestBody) SetDigest(value []byte)() {
    m.GetBackingStore().Set("digest", value)
}
