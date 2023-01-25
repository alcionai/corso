package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignmentFilterTypeAndEvaluationResult represents the filter type and evalaution result of the filter.
type AssignmentFilterTypeAndEvaluationResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Represents type of the assignment filter.
    assignmentFilterType *DeviceAndAppManagementAssignmentFilterType
    // Supported evaluation results for filter.
    evaluationResult *AssignmentFilterEvaluationResult
    // The OdataType property
    odataType *string
}
// NewAssignmentFilterTypeAndEvaluationResult instantiates a new assignmentFilterTypeAndEvaluationResult and sets the default values.
func NewAssignmentFilterTypeAndEvaluationResult()(*AssignmentFilterTypeAndEvaluationResult) {
    m := &AssignmentFilterTypeAndEvaluationResult{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssignmentFilterTypeAndEvaluationResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignmentFilterTypeAndEvaluationResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignmentFilterTypeAndEvaluationResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterTypeAndEvaluationResult) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetAssignmentFilterType gets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *AssignmentFilterTypeAndEvaluationResult) GetAssignmentFilterType()(*DeviceAndAppManagementAssignmentFilterType) {
    return m.assignmentFilterType
}
// GetEvaluationResult gets the evaluationResult property value. Supported evaluation results for filter.
func (m *AssignmentFilterTypeAndEvaluationResult) GetEvaluationResult()(*AssignmentFilterEvaluationResult) {
    return m.evaluationResult
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignmentFilterTypeAndEvaluationResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignmentFilterType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceAndAppManagementAssignmentFilterType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentFilterType(val.(*DeviceAndAppManagementAssignmentFilterType))
        }
        return nil
    }
    res["evaluationResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAssignmentFilterEvaluationResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEvaluationResult(val.(*AssignmentFilterEvaluationResult))
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AssignmentFilterTypeAndEvaluationResult) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AssignmentFilterTypeAndEvaluationResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignmentFilterType() != nil {
        cast := (*m.GetAssignmentFilterType()).String()
        err := writer.WriteStringValue("assignmentFilterType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvaluationResult() != nil {
        cast := (*m.GetEvaluationResult()).String()
        err := writer.WriteStringValue("evaluationResult", &cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignmentFilterTypeAndEvaluationResult) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetAssignmentFilterType sets the assignmentFilterType property value. Represents type of the assignment filter.
func (m *AssignmentFilterTypeAndEvaluationResult) SetAssignmentFilterType(value *DeviceAndAppManagementAssignmentFilterType)() {
    m.assignmentFilterType = value
}
// SetEvaluationResult sets the evaluationResult property value. Supported evaluation results for filter.
func (m *AssignmentFilterTypeAndEvaluationResult) SetEvaluationResult(value *AssignmentFilterEvaluationResult)() {
    m.evaluationResult = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignmentFilterTypeAndEvaluationResult) SetOdataType(value *string)() {
    m.odataType = value
}
