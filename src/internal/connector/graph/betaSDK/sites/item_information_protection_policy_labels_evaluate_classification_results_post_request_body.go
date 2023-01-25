package sites

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody provides operations to call the evaluateClassificationResults method.
type ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody instantiates a new ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody and sets the default values.
func NewItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody()(*ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) {
    m := &ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetClassificationResults gets the classificationResults property value. The classificationResults property
func (m *ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetClassificationResults()([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ClassificationResultable) {
    return m.GetBackingStore().Get("classificationResults");
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetContentInfo()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable) {
    return m.GetBackingStore().Get("contentInfo");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["classificationResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateClassificationResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ClassificationResult, len(val))
            for i, v := range val {
                res[i] = *(v.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ClassificationResult))
            }
            m.SetClassificationResults(res)
        }
        return nil
    }
    res["contentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateContentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentInfo(val.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfo))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClassificationResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClassificationResults()))
        for i, v := range m.GetClassificationResults() {
            temp := v
            cast[i] = i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("classificationResults", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("contentInfo", m.GetContentInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetClassificationResults sets the classificationResults property value. The classificationResults property
func (m *ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) SetClassificationResults(value []ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ClassificationResultable)() {
    m.GetBackingStore().Set("classificationResults", value)
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *ItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) SetContentInfo(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable)() {
    m.GetBackingStore().Set("contentInfo", value)
}
