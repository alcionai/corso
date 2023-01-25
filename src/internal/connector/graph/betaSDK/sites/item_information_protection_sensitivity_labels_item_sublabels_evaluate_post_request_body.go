package sites

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody provides operations to call the evaluate method.
type ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody instantiates a new ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody and sets the default values.
func NewItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody()(*ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody) {
    m := &ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetCurrentLabel gets the currentLabel property value. The currentLabel property
func (m *ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody) GetCurrentLabel()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CurrentLabelable) {
    return m.GetBackingStore().Get("currentLabel");
}
// GetDiscoveredSensitiveTypes gets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody) GetDiscoveredSensitiveTypes()([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DiscoveredSensitiveTypeable) {
    return m.GetBackingStore().Get("discoveredSensitiveTypes");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentLabel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateCurrentLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentLabel(val.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CurrentLabel))
        }
        return nil
    }
    res["discoveredSensitiveTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateDiscoveredSensitiveTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DiscoveredSensitiveType, len(val))
            for i, v := range val {
                res[i] = *(v.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DiscoveredSensitiveType))
            }
            m.SetDiscoveredSensitiveTypes(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("currentLabel", m.GetCurrentLabel())
        if err != nil {
            return err
        }
    }
    if m.GetDiscoveredSensitiveTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDiscoveredSensitiveTypes()))
        for i, v := range m.GetDiscoveredSensitiveTypes() {
            temp := v
            cast[i] = i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("discoveredSensitiveTypes", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetCurrentLabel sets the currentLabel property value. The currentLabel property
func (m *ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody) SetCurrentLabel(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CurrentLabelable)() {
    m.GetBackingStore().Set("currentLabel", value)
}
// SetDiscoveredSensitiveTypes sets the discoveredSensitiveTypes property value. The discoveredSensitiveTypes property
func (m *ItemInformationProtectionSensitivityLabelsItemSublabelsEvaluatePostRequestBody) SetDiscoveredSensitiveTypes(value []ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.DiscoveredSensitiveTypeable)() {
    m.GetBackingStore().Set("discoveredSensitiveTypes", value)
}
