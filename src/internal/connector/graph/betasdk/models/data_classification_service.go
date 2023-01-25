package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DataClassificationService 
type DataClassificationService struct {
    Entity
    // The classifyFileJobs property
    classifyFileJobs []JobResponseBaseable
    // The classifyTextJobs property
    classifyTextJobs []JobResponseBaseable
    // The evaluateDlpPoliciesJobs property
    evaluateDlpPoliciesJobs []JobResponseBaseable
    // The evaluateLabelJobs property
    evaluateLabelJobs []JobResponseBaseable
    // The exactMatchDataStores property
    exactMatchDataStores []ExactMatchDataStoreable
    // The exactMatchUploadAgents property
    exactMatchUploadAgents []ExactMatchUploadAgentable
    // The jobs property
    jobs []JobResponseBaseable
    // The sensitiveTypes property
    sensitiveTypes []SensitiveTypeable
    // The sensitivityLabels property
    sensitivityLabels []SensitivityLabelable
}
// NewDataClassificationService instantiates a new DataClassificationService and sets the default values.
func NewDataClassificationService()(*DataClassificationService) {
    m := &DataClassificationService{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDataClassificationServiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDataClassificationServiceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDataClassificationService(), nil
}
// GetClassifyFileJobs gets the classifyFileJobs property value. The classifyFileJobs property
func (m *DataClassificationService) GetClassifyFileJobs()([]JobResponseBaseable) {
    return m.classifyFileJobs
}
// GetClassifyTextJobs gets the classifyTextJobs property value. The classifyTextJobs property
func (m *DataClassificationService) GetClassifyTextJobs()([]JobResponseBaseable) {
    return m.classifyTextJobs
}
// GetEvaluateDlpPoliciesJobs gets the evaluateDlpPoliciesJobs property value. The evaluateDlpPoliciesJobs property
func (m *DataClassificationService) GetEvaluateDlpPoliciesJobs()([]JobResponseBaseable) {
    return m.evaluateDlpPoliciesJobs
}
// GetEvaluateLabelJobs gets the evaluateLabelJobs property value. The evaluateLabelJobs property
func (m *DataClassificationService) GetEvaluateLabelJobs()([]JobResponseBaseable) {
    return m.evaluateLabelJobs
}
// GetExactMatchDataStores gets the exactMatchDataStores property value. The exactMatchDataStores property
func (m *DataClassificationService) GetExactMatchDataStores()([]ExactMatchDataStoreable) {
    return m.exactMatchDataStores
}
// GetExactMatchUploadAgents gets the exactMatchUploadAgents property value. The exactMatchUploadAgents property
func (m *DataClassificationService) GetExactMatchUploadAgents()([]ExactMatchUploadAgentable) {
    return m.exactMatchUploadAgents
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DataClassificationService) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classifyFileJobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBaseable, len(val))
            for i, v := range val {
                res[i] = v.(JobResponseBaseable)
            }
            m.SetClassifyFileJobs(res)
        }
        return nil
    }
    res["classifyTextJobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBaseable, len(val))
            for i, v := range val {
                res[i] = v.(JobResponseBaseable)
            }
            m.SetClassifyTextJobs(res)
        }
        return nil
    }
    res["evaluateDlpPoliciesJobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBaseable, len(val))
            for i, v := range val {
                res[i] = v.(JobResponseBaseable)
            }
            m.SetEvaluateDlpPoliciesJobs(res)
        }
        return nil
    }
    res["evaluateLabelJobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBaseable, len(val))
            for i, v := range val {
                res[i] = v.(JobResponseBaseable)
            }
            m.SetEvaluateLabelJobs(res)
        }
        return nil
    }
    res["exactMatchDataStores"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExactMatchDataStoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactMatchDataStoreable, len(val))
            for i, v := range val {
                res[i] = v.(ExactMatchDataStoreable)
            }
            m.SetExactMatchDataStores(res)
        }
        return nil
    }
    res["exactMatchUploadAgents"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExactMatchUploadAgentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExactMatchUploadAgentable, len(val))
            for i, v := range val {
                res[i] = v.(ExactMatchUploadAgentable)
            }
            m.SetExactMatchUploadAgents(res)
        }
        return nil
    }
    res["jobs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateJobResponseBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]JobResponseBaseable, len(val))
            for i, v := range val {
                res[i] = v.(JobResponseBaseable)
            }
            m.SetJobs(res)
        }
        return nil
    }
    res["sensitiveTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitiveTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitiveTypeable, len(val))
            for i, v := range val {
                res[i] = v.(SensitiveTypeable)
            }
            m.SetSensitiveTypes(res)
        }
        return nil
    }
    res["sensitivityLabels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSensitivityLabelFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SensitivityLabelable, len(val))
            for i, v := range val {
                res[i] = v.(SensitivityLabelable)
            }
            m.SetSensitivityLabels(res)
        }
        return nil
    }
    return res
}
// GetJobs gets the jobs property value. The jobs property
func (m *DataClassificationService) GetJobs()([]JobResponseBaseable) {
    return m.jobs
}
// GetSensitiveTypes gets the sensitiveTypes property value. The sensitiveTypes property
func (m *DataClassificationService) GetSensitiveTypes()([]SensitiveTypeable) {
    return m.sensitiveTypes
}
// GetSensitivityLabels gets the sensitivityLabels property value. The sensitivityLabels property
func (m *DataClassificationService) GetSensitivityLabels()([]SensitivityLabelable) {
    return m.sensitivityLabels
}
// Serialize serializes information the current object
func (m *DataClassificationService) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassifyFileJobs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClassifyFileJobs()))
        for i, v := range m.GetClassifyFileJobs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("classifyFileJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetClassifyTextJobs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClassifyTextJobs()))
        for i, v := range m.GetClassifyTextJobs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("classifyTextJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvaluateDlpPoliciesJobs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvaluateDlpPoliciesJobs()))
        for i, v := range m.GetEvaluateDlpPoliciesJobs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("evaluateDlpPoliciesJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvaluateLabelJobs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvaluateLabelJobs()))
        for i, v := range m.GetEvaluateLabelJobs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("evaluateLabelJobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExactMatchDataStores() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExactMatchDataStores()))
        for i, v := range m.GetExactMatchDataStores() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("exactMatchDataStores", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExactMatchUploadAgents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExactMatchUploadAgents()))
        for i, v := range m.GetExactMatchUploadAgents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("exactMatchUploadAgents", cast)
        if err != nil {
            return err
        }
    }
    if m.GetJobs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetJobs()))
        for i, v := range m.GetJobs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("jobs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitiveTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSensitiveTypes()))
        for i, v := range m.GetSensitiveTypes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sensitiveTypes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSensitivityLabels() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSensitivityLabels()))
        for i, v := range m.GetSensitivityLabels() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sensitivityLabels", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassifyFileJobs sets the classifyFileJobs property value. The classifyFileJobs property
func (m *DataClassificationService) SetClassifyFileJobs(value []JobResponseBaseable)() {
    m.classifyFileJobs = value
}
// SetClassifyTextJobs sets the classifyTextJobs property value. The classifyTextJobs property
func (m *DataClassificationService) SetClassifyTextJobs(value []JobResponseBaseable)() {
    m.classifyTextJobs = value
}
// SetEvaluateDlpPoliciesJobs sets the evaluateDlpPoliciesJobs property value. The evaluateDlpPoliciesJobs property
func (m *DataClassificationService) SetEvaluateDlpPoliciesJobs(value []JobResponseBaseable)() {
    m.evaluateDlpPoliciesJobs = value
}
// SetEvaluateLabelJobs sets the evaluateLabelJobs property value. The evaluateLabelJobs property
func (m *DataClassificationService) SetEvaluateLabelJobs(value []JobResponseBaseable)() {
    m.evaluateLabelJobs = value
}
// SetExactMatchDataStores sets the exactMatchDataStores property value. The exactMatchDataStores property
func (m *DataClassificationService) SetExactMatchDataStores(value []ExactMatchDataStoreable)() {
    m.exactMatchDataStores = value
}
// SetExactMatchUploadAgents sets the exactMatchUploadAgents property value. The exactMatchUploadAgents property
func (m *DataClassificationService) SetExactMatchUploadAgents(value []ExactMatchUploadAgentable)() {
    m.exactMatchUploadAgents = value
}
// SetJobs sets the jobs property value. The jobs property
func (m *DataClassificationService) SetJobs(value []JobResponseBaseable)() {
    m.jobs = value
}
// SetSensitiveTypes sets the sensitiveTypes property value. The sensitiveTypes property
func (m *DataClassificationService) SetSensitiveTypes(value []SensitiveTypeable)() {
    m.sensitiveTypes = value
}
// SetSensitivityLabels sets the sensitivityLabels property value. The sensitivityLabels property
func (m *DataClassificationService) SetSensitivityLabels(value []SensitivityLabelable)() {
    m.sensitivityLabels = value
}
