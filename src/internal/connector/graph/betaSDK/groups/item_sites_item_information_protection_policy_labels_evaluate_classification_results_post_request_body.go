package groups

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody provides operations to call the evaluateClassificationResults method.
type ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The classificationResults property
    classificationResults []ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ClassificationResultable
    // The contentInfo property
    contentInfo ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable
}
// NewItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody instantiates a new ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody and sets the default values.
func NewItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody()(*ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) {
    m := &ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetClassificationResults gets the classificationResults property value. The classificationResults property
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetClassificationResults()([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ClassificationResultable) {
    return m.classificationResults
}
// GetContentInfo gets the contentInfo property value. The contentInfo property
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetContentInfo()(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable) {
    return m.contentInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["classificationResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateClassificationResultFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ClassificationResultable, len(val))
            for i, v := range val {
                res[i] = v.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ClassificationResultable)
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
            m.SetContentInfo(val.(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetClassificationResults() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClassificationResults()))
        for i, v := range m.GetClassificationResults() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetClassificationResults sets the classificationResults property value. The classificationResults property
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) SetClassificationResults(value []ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ClassificationResultable)() {
    m.classificationResults = value
}
// SetContentInfo sets the contentInfo property value. The contentInfo property
func (m *ItemSitesItemInformationProtectionPolicyLabelsEvaluateClassificationResultsPostRequestBody) SetContentInfo(value ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.ContentInfoable)() {
    m.contentInfo = value
}
